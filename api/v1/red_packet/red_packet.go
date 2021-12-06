package red_packet

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RedPacketApi struct {
}

func (*RedPacketApi) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}
