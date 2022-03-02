package po

import "time"

var RedPacketModel = &RedPacket{}

// RedPacket 用户发的红包记录
type RedPacket struct {
	Base
	Amount      float64   `gorm:"type:decimal(6,2);NOT NULL;DEFAULT:0.00;COMMENT '总金额'"`
	Count       uint      `gorm:"type:INT(5);NOT NULL;COMMENT '数量'"`
	Balance     float64   `gorm:"type:decimal(6,2);NOT NULL;DEFAULT:0.00;COMMENT '剩余金额'"`
	RemainCount uint      `gorm:"type:INT(5);NOT NULL;COMMENT '剩余数量'"`
	UserId      uint      `gorm:"type:INT(10);NOT NULL;index:user"`
	ExpireAt    time.Time `gorm:"NOT NULL;COMMENT '过期时间'"`
	IsExpire    byte      `gorm:"type:tinyint(1);NOT NULL;default 0;COMMENT '红包是否过期'"`
}
