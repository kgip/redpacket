package impl

import (
	"redpacket/exception"
	"redpacket/global"
	"redpacket/model/common"
	"redpacket/model/po"
	"redpacket/model/vo"
	"redpacket/utils"
)

type UserService struct {
}

func (*UserService) GetUserList(page *common.Page) *common.Page {
	list := make([]*po.User, page.Limit)
	var total int64
	if page.Page == 0 {
		exception.TryThrow(global.DB.Find(&list).Error)
		total = int64(len(list))
	} else {
		exception.TryThrow(global.DB.Offset((page.Page - 1) * page.Limit).Limit(page.Limit).Find(&list).Error)
		exception.TryThrow(global.DB.Count(&total).Error)
	}
	userVoList := make([]*vo.UserVo, len(list))
	for i, user := range list {
		userVo := &vo.UserVo{}
		utils.BeanCopy(user, userVo)
		userVo.CreatedAt = common.JSONTime(user.CreatedAt)
		userVoList[i] = userVo
	}
	page.Total = int(total)
	page.Data = userVoList
	return page
}

func (*UserService) AddUsers(users []*po.User) bool {
	if len(users) >= 0 {
		exception.TryThrow(global.DB.CreateInBatches(users, len(users)).Error)
	}
	return true
}

func (*UserService) DeleteUserById(id int) bool {
	exception.TryThrow(global.DB.Delete(&po.User{}, id).Error)
	return true
}

func (*UserService) UserTest() bool {
	return true
}
