package middleware

import (
	"github.com/gin-gonic/gin"
	"redpacket/exception"
	"redpacket/model/common"
)

func ErrorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ex, ok := err.(exception.Exception); ok {
				common.FailWithMessageCode(ex.Error(), ex.GetCode(), c)
			} else if e, ok := err.(error); ok {
				common.FailWithMessage(e.Error(), c)
			} else if msg, ok := err.(string); ok {
				common.FailWithMessage(msg, c)
			} else {
				common.FailWithMessage("Unknown Internal Error!", c)
			}
			c.Abort()
		}
	}()
	c.Next()
}
