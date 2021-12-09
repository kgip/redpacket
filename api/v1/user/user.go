package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redpacket/exception"
	"redpacket/model/common"
	"redpacket/service"
	"strconv"
)

type UserApi struct {
}

var userService = service.ServiceGroups.UserServiceGroup.UserService

func (*UserApi) Test(c *gin.Context) {
	i := c.Param("i")
	intI, err := strconv.ParseInt(i, 10, 64)
	exception.TryThrow(err)
	fmt.Println(intI + 1)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}

func (*UserApi) GetUserList(c *gin.Context) {
	page := &common.Page{}
	c.BindJSON(page)
	page = userService.GetUserList(page)
	common.OkWithData(page, c)
}
