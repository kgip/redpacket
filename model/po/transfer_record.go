package po

// TransferRecord 用户转账记录
type TransferRecord struct {
	Base
	SenderId   uint    `gorm:"type:INT(10);NOT NULL"`
	ReceiverId uint    `gorm:"type:INT(10);NOT NULL"`
	Amount     float64 `gorm:"type:INT(10);NOT NULL;DEFAULT 0;COMMENT '转账金额'"`
}
