package user

import (
	"redpacket/exception"
	"redpacket/global"
	"redpacket/model"
	"redpacket/model/common"
	"redpacket/model/response"
	"redpacket/utils"
)

type UserService struct {
}

func (*UserService) GetUserList(page *common.Page) *common.Page {
	list := make([]*model.User, page.Limit)
	var total int64
	if page.Page == 0 {
		exception.TryThrow(global.DB.Find(&list).Error)
		total = int64(len(list))
	} else {
		exception.TryThrow(global.DB.Offset((page.Page - 1) * page.Limit).Limit(page.Limit).Find(&list).Error)
		exception.TryThrow(global.DB.Count(&total).Error)
	}
	userVoList := make([]*response.UserVo, len(list))
	for i, user := range list {
		userVo := &response.UserVo{}
		utils.BeanCopy(user, userVo)
		userVo.CreatedAt = common.JSONTime(user.CreatedAt)
		userVoList[i] = userVo
	}
	page.Total = int(total)
	page.Data = userVoList
	return page
}
