package user

import (
	"github.com/gin-gonic/gin"
	v1 "redpacket/api/v1"
)

type UserRouter struct {
}

func (*UserRouter) InitRouter(group *gin.RouterGroup) {
	userRouterGroup := group.Group("user")
	var userApi = v1.ApiGroups.UserApiGroup.UserApi
	{
		userRouterGroup.GET("list", userApi.GetUserList)
	}
}
