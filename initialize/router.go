package initialize

import (
	"github.com/gin-gonic/gin"
	v1 "redpacket/router/v1"
)

func Router(group *gin.RouterGroup) {
	v1Group := group.Group("v1")
	v1.RouterGroups.UserRouterGroup.InitRouter(v1Group)
	v1.RouterGroups.RedPacketRouterGroup.InitRedPacketRouter(v1Group)
}
