package po

var UserModel = &User{}

type User struct {
	Base
	Username string  `gorm:"type:VARCHAR(50);NOT NULL;index:unique"`
	Balance  float64 `gorm:"type:decimal(20,2);NOT NULL;DEFAULT 0.00;COMMENT '账号余额'"`
}
