package red_packet

import (
	"github.com/gin-gonic/gin"
	v1 "redpacket/api/v1"
)

type RedPacketRouter struct {
}

func (*RedPacketRouter) InitRedPacketRouter(group *gin.RouterGroup) {
	redPacketRouterGroup := group.Group("redpacket")
	redPacketApi := v1.ApiGroups.RedPacketApiGroup.RedPacketApi
	{
		redPacketRouterGroup.GET("test", redPacketApi.Test)
	}
}
