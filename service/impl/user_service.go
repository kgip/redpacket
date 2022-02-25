package impl

import (
	"redpacket/ex"
	"redpacket/global"
	"redpacket/model/common"
	"redpacket/model/po"
	"redpacket/model/vo"
	"redpacket/utils"
)

type UserService struct{}

func (*UserService) GetUserList(page *common.Page) *common.Page {
	list := make([]*po.User, page.Limit)
	var total int64
	if page.Page == 0 {
		ex.TryThrow(global.DB.Model(po.UserModel).Find(&list).Error)
		total = int64(len(list))
	} else {
		ex.TryThrow(global.DB.Model(po.UserModel).Offset((page.Page - 1) * page.Limit).Limit(page.Limit).Find(&list).Error)
		ex.TryThrow(global.DB.Model(po.UserModel).Count(&total).Error)
	}
	userVoList := make([]*vo.UserVo, len(list))
	for i, user := range list {
		userVo := &vo.UserVo{CreatedAt: common.JSONTime(user.CreatedAt)}
		utils.BeanCopy(user, userVo, "CreatedAt")
		userVo.CreatedAt = common.JSONTime(user.CreatedAt)
		userVoList[i] = userVo
	}
	page.Total = int(total)
	page.Data = userVoList
	return page
}

func (*UserService) AddUser(userVo *vo.UserAddVo) bool {
	if userVo != nil {
		ex.TryThrow(global.DB.Create(&po.User{Username: userVo.Username, Balance: userVo.Balance}).Error)
	}
	return true
}
