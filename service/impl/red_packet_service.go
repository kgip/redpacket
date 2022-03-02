package impl

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kgip/redis-lock/lock"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"redpacket/ex"
	"redpacket/global"
	"redpacket/model/po"
	"redpacket/model/vo"
	"redpacket/utils"
	"redpacket/utils/constant"
	"strconv"
	"time"
)

type RedPacketService struct{}

func (r *RedPacketService) SendPacket(vo *vo.SendPacketVo, id uint) {
	ex.TryThrow(global.DB.Transaction(func(tx *gorm.DB) (err error) {
		//1.查询是否金额足够
		var balance float64
		if tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(po.UserModel).Select("balance").Where("id = ?", id).First(&balance); balance >= vo.Balance {
			//2.扣减用户金额
			ex.TryThrow(ex.HandleDbError(tx.Where("id = ? and balance >= ?", id, vo.Balance).Update("balance", gorm.Expr("balance - ?", vo.Balance))))
			//3.创建红包记录
			redpacket := &po.RedPacket{
				Amount:      vo.Balance,
				Count:       uint(vo.Count),
				Balance:     vo.Balance,
				RemainCount: uint(vo.Count),
				UserId:      id,
				ExpireAt:    time.Now().Add(constant.RedPacketExpireDuration), //红包过期时间为7天
			}
			ex.HandleDbError(tx.Create(redpacket))
			//4.发送定时消息到本地消息队列，红包过期后返还剩余金额
			ex.TryThrow(global.MQ.SendMessage(constant.ReturnRedPacketBalanceTopic, redpacket.ID, constant.RedPacketExpireDuration))
			//5.在redis中保存红包记录
			r.addRedPacketToRedis(redpacket.ID, redpacket.Count, redpacket.Balance)
		} else {
			return ex.InsufficientBalanceException
		}
		return
	}))
}

func (*RedPacketService) addRedPacketToRedis(id, count uint, balance float64) {
	key := fmt.Sprintf("%s%d", constant.RedPacketKeyPrefix, id)
	redPackets := utils.GenericRedPackets(count, balance)
	ex.TryThrow(global.Redis.RPush(context.Background(), key, redPackets...).Err())
}

func (r *RedPacketService) GrabPacket(id, userId uint) {
	key := fmt.Sprintf("%s%d", constant.GrabRedPacketUserSetKeyPrefix, id)
	//加分布式锁
	lockKey := fmt.Sprintf("%s%d", constant.RedPacketLockKeyPrefix, id)
	global.LockOperator.Lock(lockKey, lock.Context())
	defer global.LockOperator.Unlock(lockKey)

	//todo lua脚本
	//检查用户是否已经抢过红包
	if _, err := global.Redis.ZRank(context.Background(), key, fmt.Sprintf("%d", userId)).Result(); err == nil {
		panic(ex.RepeatGrabRedPacketException)
	}
	//redis List中获取红包
	strAmount, err := global.Redis.LPop(context.Background(), fmt.Sprintf("%s%d", constant.RedPacketKeyPrefix, id)).Result()

	ex.TryThrow(err, ex.PacketIsEmptyException)
	grabBalance, err := strconv.ParseFloat(strAmount, 64)
	ex.TryThrow(err)
	var rpRecordCount int64
	//防止redis记录删除失败，检查数据库中的红包记录看红包是否过期
	if global.DB.Model(po.RedPacketModel).Where("id = ? and is_expire = ?", id, 0).Count(&rpRecordCount); rpRecordCount <= 0 {
		panic(ex.PacketIsExpireException)
	}
	//保存红包记录，并转账给用户
	ex.TryThrow(global.DB.Transaction(func(tx *gorm.DB) error {
		//插入抢红包记录
		ex.TryThrow(ex.HandleDbError(tx.Create(&po.GrabRedPacketRecord{RedPacketId: id, UserId: userId, Amount: grabBalance})))
		//给用户转账并生成转账记录
		ex.TryThrow(ex.HandleDbError(tx.Model(po.UserModel).Where("id = ?", userId).Update("balance", gorm.Expr("balance + ?", grabBalance))))
		ex.TryThrow(ex.HandleDbError(tx.Create(&po.TransferRecord{SenderId: 00000, ReceiverId: userId, Amount: grabBalance})))
		//在redis中保存抢红包相关信息
		ex.TryThrow(global.Redis.ZAdd(context.Background(), key, &redis.Z{Member: userId, Score: grabBalance}).Err())
		return nil
	}))
}
