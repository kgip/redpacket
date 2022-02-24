package service

import (
	"redpacket/model/common"
	"redpacket/model/po"
)

type UserService interface {
	GetUserList(page *common.Page) *common.Page
	AddUsers(users []*po.User) bool
	DeleteUserById(id int) bool
	UserTest() bool
}
