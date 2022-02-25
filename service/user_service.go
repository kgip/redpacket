package service

import (
	"redpacket/model/common"
	"redpacket/model/vo"
)

type UserService interface {
	GetUserList(page *common.Page) *common.Page
	AddUser(user *vo.UserAddVo) bool
}
