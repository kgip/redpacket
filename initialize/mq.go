package initialize

import (
	"gorm.io/gorm"
	"redpacket/global"
	"redpacket/model/po"
	"redpacket/utils/constant"
	"redpacket/utils/mq"
)

var returnExpireRedPacketBalance = func(msg interface{}) {
	redPacketId := msg.(int64)
	global.DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(po.RedPacketModel).Where("id = ?", redPacketId)
		return nil
	})
}

func MQ() mq.MqOperator {
	operator := &mq.LocalMQ{}
	operator.AddQueue(constant.ReturnRedPacketBalanceTopic, 1000)

	operator.RegistryMessageHandler([]string{constant.ReturnRedPacketBalanceTopic}, returnExpireRedPacketBalance)
	return operator
}
