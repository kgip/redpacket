package po

var GrabRedPacketRecordModel = &GrabRedPacketRecord{}

// GrabRedPacketRecord 用户抢红包记录
type GrabRedPacketRecord struct {
	Base
	UserId      uint    `gorm:"type:int(10);not null;index:user_redpacket_idx"`
	RedPacketId uint    `gorm:"type:int(10);not null;index:user_redpacket_idx"`
	Amount      float64 `gorm:"not null;default 0"`
}
