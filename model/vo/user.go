package vo

import (
	"redpacket/model/common"
)

type UserVo struct {
	ID        uint            `json:"id"`
	Username  string          `json:"username"`
	Balance   float64         `json:"balance"`
	CreatedAt common.JSONTime `json:"createdAt"`
}

type UserAddVo struct {
	Username string  `json:"username" binding:"required"`
	Balance  float64 `json:"balance" binding:"required,gte=0"`
}
