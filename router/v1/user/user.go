package user

import (
	"github.com/gin-gonic/gin"
	v1 "redpacket/api/v1"
	"redpacket/global/service"
)

type UserRouter struct{}

func (*UserRouter) InitRouter(group *gin.RouterGroup) {
	userRouterGroup := group.Group("user")
	var userApi = &v1.UserApi{UserService: service.UserService}
	{
		userRouterGroup.GET("list", userApi.GetUserList)
		userRouterGroup.POST("", userApi.AddUser)
	}
}
