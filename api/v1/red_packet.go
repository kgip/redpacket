package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redpacket/service"
)

type RedPacketApi struct {
	RedPacketService service.RedPacketService
}

func (*RedPacketApi) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}
