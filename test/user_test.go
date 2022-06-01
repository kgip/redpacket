package test

import (
	"github.com/gin-gonic/gin"
	gin_mock "github.com/kgip/mock-utils/core/gin-mock"
	gorm_mock "github.com/kgip/mock-utils/core/gorm-mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"redpacket/global"
	"redpacket/initialize"
	"redpacket/model/vo"
	"testing"
)

var (
	router *gin.Engine
)

func init() {
	////1.初始化配置文件ls
	//initialize.Config(fmt.Sprintf("../%s", global.ConfigPath))
	////2.初始化zap日志
	global.LOG = initialize.Zap()
	initialize.Service()
	router = initialize.Router()
}

func TestAddUser(t *testing.T) {
	global.DB = gorm_mock.NewDBMockCreator().Insert(1, 1).Create()

	r := gin_mock.DoRequest(router, http.MethodPost, "/redpacket/v1/user", &vo.UserAddVo{Username: "aaa", Balance: 111})
	assert.Equal(t, http.StatusOK, r.Code)
}
