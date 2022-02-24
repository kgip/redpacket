package initialize

import (
	"github.com/gin-gonic/gin"
	"redpacket/middleware"
	v1 "redpacket/router/v1"
)

func Router() *gin.Engine {
	router := gin.Default()
	group := router.Group("redpacket")

	v1Group := group.Group("v1")
	v1Group.Use(middleware.ErrorHandler)
	{
		v1.RouterGroups.UserRouterGroup.InitRouter(v1Group)
		v1.RouterGroups.RedPacketRouterGroup.InitRedPacketRouter(v1Group)
	}
	return router
}
