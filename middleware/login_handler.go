package middleware

import (
	"github.com/gin-gonic/gin"
	"redpacket/ex"
)

func LoginHandler(c *gin.Context) {
	id := c.GetHeader("userId")
	if id != "" {
		c.Set("userId", id)
	} else {
		panic(ex.LoginException)
	}
	c.Next()
}
