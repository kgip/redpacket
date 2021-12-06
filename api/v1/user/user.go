package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserApi struct {
}

func (*UserApi) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}
