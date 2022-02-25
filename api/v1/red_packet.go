package v1

import (
	"github.com/gin-gonic/gin"
	"redpacket/ex"
	"redpacket/model/common"
	"redpacket/model/vo"
	"redpacket/service"
)

type RedPacketApi struct {
	RedPacketService service.RedPacketService
}

func (r *RedPacketApi) SendPacket(c *gin.Context) {
	packetVo := &vo.SendPacketVo{}
	ex.TryThrow(c.ShouldBind(packetVo), ex.RequestParamsException)
	r.RedPacketService.SendPacket(packetVo)
	common.Ok(c)
}

func (r *RedPacketApi) GrabPacket(c *gin.Context) {

}
