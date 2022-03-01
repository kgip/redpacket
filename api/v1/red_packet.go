package v1

import (
	"github.com/gin-gonic/gin"
	"redpacket/ex"
	"redpacket/model/common"
	"redpacket/model/vo"
	"redpacket/service"
	"strconv"
)

type RedPacketApi struct {
	RedPacketService service.RedPacketService
}

func (r *RedPacketApi) SendPacket(c *gin.Context) {
	packetVo := &vo.SendPacketVo{}
	ex.TryThrow(c.ShouldBind(packetVo), ex.RequestParamsException)
	id, _ := c.Get("userId")
	intId, err := strconv.ParseUint(id.(string), 10, 64)
	ex.TryThrow(err, ex.RequestParamsException)
	r.RedPacketService.SendPacket(packetVo, uint(intId))
	common.Ok(c)
}

func (r *RedPacketApi) GrabPacket(c *gin.Context) {
	idVo := &vo.RedPacketIdVo{}
	ex.TryThrow(c.ShouldBind(idVo), ex.RequestParamsException)
	userId, _ := c.Get("userId")
	intUserId, err := strconv.ParseUint(userId.(string), 10, 64)
	ex.TryThrow(err, ex.RequestParamsException)
	r.RedPacketService.GrabPacket(idVo.Id, uint(intUserId))
	common.Ok(c)
}
