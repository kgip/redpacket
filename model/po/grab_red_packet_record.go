package po

var GrabRedPacketRecordModel = &GrabRedPacketRecord{}

// GrabRedPacketRecord 用户抢红包记录
type GrabRedPacketRecord struct {
	Base
	RedPacketId uint    `gorm:"type:int(10);not null;index:redpacket_user_idx"`
	UserId      uint    `gorm:"type:int(10);not null;index:redpacket_user_idx"`
	Amount      float64 `gorm:"not null;default 0"`
}
