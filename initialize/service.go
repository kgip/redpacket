package initialize

import (
	"redpacket/global/service"
	"redpacket/service/impl"
)

func Service() {
	service.UserService = &impl.UserService{}
	service.RedPacketService = &impl.RedPacketSerivce{}
}
