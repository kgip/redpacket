package model

// RedPacket 用户发的红包记录
type RedPacket struct {
	Base
	Amount      float64 `gorm:"NOT NULL;DEFAULT:0;COMMENT '总金额'"`
	Count       uint    `gorm:"type:INT(5);NOT NULL;COMMENT '数量'"`
	Balance     float64 `gorm:"NOT NULL;DEFAULT:0;COMMENT '剩余金额'"`
	RemainCount uint    `gorm:"type:INT(5);NOT NULL;COMMENT '剩余数量'"`
	UserId      uint    `gorm:"type:INT(10);NOT NULL"`
}
