package impl

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"redpacket/ex"
	"redpacket/global"
	"redpacket/model/po"
	"redpacket/model/vo"
	"redpacket/utils/constant"
	"strconv"
	"time"
)

type RedPacketService struct{}

func (r *RedPacketService) SendPacket(vo *vo.SendPacketVo, id interface{}) {
	intId, err := strconv.ParseInt(id.(string), 10, 64)
	ex.TryThrow(err)
	ex.TryThrow(global.DB.Transaction(func(tx *gorm.DB) (err error) {
		//1.查询是否金额足够
		var balance float64
		if tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(po.UserModel).Select("balance").Where("id = ?", intId).First(&balance); balance >= vo.Balance {
			//2.扣减用户金额
			ex.TryThrow(ex.HandleDbError(tx.Where("id = ? and balance >= ?", id, vo.Balance).Update("balance", gorm.Expr("balance - ?", vo.Balance))))
			//3.创建红包记录
			redpacket := &po.RedPacket{
				Amount:      vo.Balance,
				Count:       uint(vo.Count),
				Balance:     vo.Balance,
				RemainCount: uint(vo.Count),
				UserId:      uint(intId),
				ExpireAt:    time.Now().AddDate(0, 0, 7), //红包过期时间为7天
			}
			ex.HandleDbError(tx.Create(redpacket))
			//4.在redis中保存红包记录
			r.addRedPacketToRedisHSet(redpacket.ID, redpacket.Count, redpacket.Balance)
			//5.发送定时消息到本地消息队列，红包过期后返还剩余金额
			global.MQ.SendMessage(constant.ReturnRedPacketBalanceTopic, redpacket.ID, time.Hour*24*7)
		} else {
			return ex.InsufficientBalance
		}
		return
	}))
}

func (*RedPacketService) addRedPacketToRedisHSet(id, count uint, balance float64) {
	key := fmt.Sprintf("%s%d", constant.RedPacketKeyPrefix, id)
	ex.TryThrow(global.Redis.HMSet(context.Background(), key, map[string]interface{}{
		constant.RedPacketHSetCountField:   count,
		constant.RedPacketHSetBalanceField: balance,
	}).Err())
	global.Redis.Expire(context.Background(), key, time.Hour*24*7)
}
