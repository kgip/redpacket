package response

import (
	"redpacket/model/common"
)

type UserVo struct {
	ID        uint            `json:"id"`
	Username  string          `json:"username"`
	Balance   float64         `json:"balance"`
	CreatedAt common.JSONTime `json:"createdAt"`
}
