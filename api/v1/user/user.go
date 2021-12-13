package user

import (
	"github.com/gin-gonic/gin"
	"redpacket/exception"
	"redpacket/model/common"
	"redpacket/service"
)

type UserApi struct {
}

var userService = service.ServiceGroups.UserServiceGroup.UserService

func (*UserApi) GetUserList(c *gin.Context) {
	page := &common.Page{}
	exception.TryThrow(c.ShouldBind(page))
	page = userService.GetUserList(page)
	common.OkWithData(page, c)
}
