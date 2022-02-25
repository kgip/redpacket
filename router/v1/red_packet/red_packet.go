package red_packet

import (
	"github.com/gin-gonic/gin"
	v1 "redpacket/api/v1"
	"redpacket/global/service"
)

type RedPacketRouter struct{}

func (*RedPacketRouter) InitRedPacketRouter(group *gin.RouterGroup) {
	redPacketRouterGroup := group.Group("redpacket")
	redPacketApi := &v1.RedPacketApi{RedPacketService: service.RedPacketService}
	{
		redPacketRouterGroup.POST("send", redPacketApi.SendPacket)
		redPacketRouterGroup.POST("grab", redPacketApi.GrabPacket)
	}
}
