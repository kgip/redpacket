package po

var TransferRecordModel = &TransferRecord{}

// TransferRecord 用户转账记录
type TransferRecord struct {
	Base
	SenderId   uint    `gorm:"type:INT(10);NOT NULL;index:sender"`
	ReceiverId uint    `gorm:"type:INT(10);NOT NULL;index:receiver"`
	Amount     float64 `gorm:"type:INT(10);NOT NULL;DEFAULT 0;COMMENT '转账金额'"`
}
