package po

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Base struct {
	ID uint
	//ID string
	CreatedAt time.Time             `gorm:"not null"`
	Deleted   soft_delete.DeletedAt `gorm:"type:smallint(1);softDelete:flag"`
}
