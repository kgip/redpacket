package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redpacket/global"
	"redpacket/model/common"
	"redpacket/service"
)

type UserApi struct {
	UserService service.UserService
}

func (u *UserApi) GetUserList(c *gin.Context) {
	page := &common.Page{}
	c.ShouldBind(page)
	page2 := &common.Page{}
	if err := c.ShouldBind(page2); err != nil {
		fmt.Println(err)
	}
	page = u.UserService.GetUserList(page)
	common.OkWithData(page, c)
}

func (u *UserApi) UserTest(c *gin.Context) {
	if u.UserService.UserTest() {
		global.LOG.Info("gin单元测试")
	}
	common.Ok(c)
}
