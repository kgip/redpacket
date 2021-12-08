package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redpacket/service"
)

type UserApi struct {
}

var userService = service.ServiceGroups.UserServiceGroup.UserService

func (*UserApi) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}

func (*UserApi) GetUserList(c *gin.Context) {
	//c.BindJSON()
	//userService.GetUserList()
}
