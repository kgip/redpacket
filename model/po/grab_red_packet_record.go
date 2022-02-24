package po

// GrabRedPacketRecord 用户抢红包记录
type GrabRedPacketRecord struct {
	Base
	RedPacketId uint    `gorm:"type:int(10);not null"`
	UserId      uint    `gorm:"type:int(10);not null"`
	Amount      float64 `gorm:"not null;default 0"`
}
