package v1

import (
	"github.com/gin-gonic/gin"
	"redpacket/ex"
	"redpacket/model/common"
	"redpacket/model/vo"
	"redpacket/service"
)

type UserApi struct {
	UserService service.UserService
}

func (u *UserApi) GetUserList(c *gin.Context) {
	page := &common.Page{}
	ex.TryThrow(c.ShouldBind(page), ex.RequestParamsException)
	page = u.UserService.GetUserList(page)
	common.OkWithData(page, c)
}

func (u *UserApi) AddUser(c *gin.Context) {
	userVo := &vo.UserAddVo{}
	ex.TryThrow(c.ShouldBind(userVo), ex.RequestParamsException)
	u.UserService.AddUser(userVo)
	common.Ok(c)
}
