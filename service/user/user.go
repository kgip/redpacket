package user

import (
	"redpacket/exception"
	"redpacket/global"
	"redpacket/model"
	"redpacket/model/common"
)

type UserService struct {
}

func (*UserService) GetUserList(page *common.Page) *common.Page {
	list := make([]*model.User, page.Limit)
	exception.TryThrow(global.DB.Offset(page.Page).Limit(page.Limit).Find(list).Error)
	page.Data = list
	return page
}
