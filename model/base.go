package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint           `gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
