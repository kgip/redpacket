package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"redpacket/global"
	"redpacket/initialize"
	"syscall"
)

//服务被杀死后的回调
func stopCallback() {
	zap.S().Error("服务被杀死")
}

func main() {
	//1.初始化配置文件ls
	initialize.Config(global.ConfigPath)
	//2.初始化zap日志
	global.LOG = initialize.Zap()
	//3.初始化redis
	global.Redis = initialize.Redis()
	//4.初始化gorm
	global.DB = initialize.Gorm()
	//5.初始化server
	go func() {
		//接收服务异常退出信息
		signalChan := make(chan os.Signal)
		signal.Notify(signalChan, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-signalChan
		stopCallback()
	}()
	server := initialize.Router()
	server.Run(fmt.Sprintf("%s:%d", global.Config.Server.Host, global.Config.Server.Port))
}
