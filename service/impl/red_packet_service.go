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
			r.addRedPacketToRedisHSet(redpacket.ID, redpacket.Count, redpacket.Balance)
		} else {
			return ex.InsufficientBalanceException
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
	global.Redis.Expire(context.Background(), key, constant.RedPacketExpireDuration)
}

// GrabPacket todo 方案性能比较：1.直接加锁操作 2.使用双重检查锁操作
func (r *RedPacketService) GrabPacket(id, userId uint) {
	//1.检查包是否为空
	r.isRedPacketEmpty(id)
	lockKey := fmt.Sprintf("%s%d", constant.RedPacketLockKeyPrefix, id)
	global.LockOperator.Lock(lockKey, lock.Context())
	defer global.LockOperator.Unlock(lockKey)
	//2.再次检查包是否为空
	userCount, count := r.isRedPacketEmpty(id)
	var rpRecordCount int64
	//防止redis记录删除失败，检查数据库中的红包记录看红包是否过期
	if global.DB.Model(po.RedPacketModel).Where("id = ? and is_expire = ?", id, 0).Count(&rpRecordCount); rpRecordCount <= 0 {
		panic(ex.PacketIsExpireException)
	}
	//红包不为空
	//3.获取红包余额
	strBalance, err := global.Redis.HGet(context.Background(), fmt.Sprintf("%s%d", constant.RedPacketKeyPrefix, id), constant.RedPacketHSetBalanceField).Result()
	ex.TryThrow(err)
	balance, err := strconv.ParseFloat(strBalance, 64)
	ex.TryThrow(err)
	var grabBalance = balance //用户抢到的红包金额
	ex.TryThrow(err)
	//如果不是最后一个红包，需要计算红包金额
	if count > userCount+1 {
		grabBalance = utils.CalculateRedPacketBalance(count-userCount, balance)
	}

	//保存红包记录，并转账给用户
	ex.TryThrow(global.DB.Transaction(func(tx *gorm.DB) error {
		//插入抢红包记录
		ex.TryThrow(ex.HandleDbError(tx.Create(&po.GrabRedPacketRecord{RedPacketId: id, UserId: userId, Amount: grabBalance})))
		//给用户转账并生成转账记录
		ex.TryThrow(ex.HandleDbError(tx.Model(po.UserModel).Where("id = ?", userId).Update("balance", gorm.Expr("balance + ?", grabBalance))))
		ex.TryThrow(ex.HandleDbError(tx.Create(&po.TransferRecord{SenderId: 00000, ReceiverId: userId, Amount: grabBalance})))
		//在redis中保存抢红包相关信息
		ex.TryThrow(global.Redis.ZAdd(context.Background(), fmt.Sprintf("%s%d", constant.GrabRedPacketUserSetKeyPrefix, id), &redis.Z{Member: userId, Score: grabBalance}).Err())
		ex.TryThrow(global.Redis.HSet(context.Background(), fmt.Sprintf("%s%d", constant.RedPacketKeyPrefix, id), map[string]interface{}{
			constant.RedPacketHSetBalanceField: balance - grabBalance,
		}).Err())
		return nil
	}))
}

//判断红包是否为空，如果不为空，返回已经抢红包的人数和红包总数
func (*RedPacketService) isRedPacketEmpty(redPacketId uint) (userCount, redPacketCount int64) {
	if r := global.Redis.HGet(context.Background(), fmt.Sprintf("%s%d", constant.RedPacketKeyPrefix, redPacketId), constant.RedPacketHSetCountField); r.Err() != nil {
		panic(r.Err())
	} else {
		strCount, _ := r.Result()
		count, err := strconv.ParseInt(strCount, 10, 64)
		ex.TryThrow(err)
		//2.检查红包是否已经被抢完
		if userCount, err = global.Redis.ZCard(context.Background(), fmt.Sprintf("%s%d", constant.GrabRedPacketUserSetKeyPrefix, redPacketId)).Result(); err != nil {
			panic(err)
		} else if userCount >= count {
			panic(ex.PacketIsEmptyException)
		}
		return userCount, count
	}
}
