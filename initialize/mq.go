package initialize

import (
	"context"
	"fmt"
	"github.com/kgip/redis-lock/lock"
	"gorm.io/gorm"
	"redpacket/ex"
	"redpacket/global"
	"redpacket/model/po"
	"redpacket/utils/constant"
	"redpacket/utils/mq"
)

//红包过期后返还给用户红包剩余金额
var returnExpireRedPacketBalance = func(msg interface{}) {
	redPacketId := msg.(int64)
	//1.删除redis中的红包记录和用户记录
	if err := global.Redis.Del(context.Background(), fmt.Sprintf("%s%d", constant.RedPacketKeyPrefix, redPacketId)).Err(); err != nil {
		global.LOG.Error(err.Error())
	}
	if err := global.Redis.Del(context.Background(), fmt.Sprintf("%s%d", constant.GrabRedPacketUserSetKeyPrefix, redPacketId)).Err(); err != nil {
		global.LOG.Error(err.Error())
	}

	lockKey := fmt.Sprintf("%s%d", constant.RedPacketLockKeyPrefix, redPacketId)
	global.LockOperator.Lock(lockKey, lock.Context())
	defer global.LockOperator.Unlock(lockKey)
	//2.获取红包领取的总金额
	var grabTotalAmount float64
	ex.TryThrow(global.DB.Model(po.GrabRedPacketRecordModel).Select("sum(amount)").Where("red_packet_id = ?", redPacketId).Find(&grabTotalAmount).Error)
	//3.获取红包总金额
	rp := &po.RedPacket{}
	ex.TryThrow(global.DB.Model(po.RedPacketModel).Select("amount", "user_id").Where("id = ?", redPacketId).Find(rp).Error)
	//未查询到记录，不进行后续处理
	if rp.UserId <= 0 {
		return
	}
	if rp.Amount > grabTotalAmount {
		ex.TryThrow(global.DB.Transaction(func(tx *gorm.DB) error {
			//4.更新红包记录状态为过期
			ex.TryThrow(tx.Model(po.RedPacketModel).Where("id = ?", redPacketId).Update("is_expire", 1).Error)
			//5.将剩余的金额返还用户账号
			ex.TryThrow(tx.Model(po.UserModel).Where("id = ?", rp.UserId).Update("balance = ?", gorm.Expr("balance + ?", rp.Amount-grabTotalAmount)).Error)
			return nil
		}))
	} else {
		//4.更新红包记录状态为过期
		ex.TryThrow(global.DB.Model(po.RedPacketModel).Where("id = ?", redPacketId).Update("is_expire", 1).Error)
	}
}

func MQ() mq.MqOperator {
	operator := &mq.LocalMQ{}
	operator.AddQueue(constant.ReturnRedPacketBalanceTopic, 1000)

	operator.RegistryMessageHandler([]string{constant.ReturnRedPacketBalanceTopic}, returnExpireRedPacketBalance)
	return operator
}
