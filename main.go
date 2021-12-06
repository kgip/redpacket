package main

import (
	"fmt"
	"os"
	"os/signal"
	"redpacket/global"
	"redpacket/initialize"
)

//服务被杀死后的回调
func stopCallback() {
	global.LOG.Error("服务被杀死")
}

func main() {
	//1.初始化配置文件ls
	initialize.Config(global.CONFIG_PATH)
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
		signal.Notify(signalChan, os.Kill, os.Interrupt)
		<-signalChan
		stopCallback()
	}()
	server := initialize.Server()
	server.Run(fmt.Sprintf("%s:%d", global.Config.Server.Host, global.Config.Server.Port))
}
