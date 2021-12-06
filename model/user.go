package model

type User struct {
	Base
	Username string  `gorm:"type:VARCHAR(50);NOT NULL"`
	Balance  float64 `gorm:"type:INT(10);NOT NULL;DEFAULT 0;COMMENT '账号余额'"`
}
