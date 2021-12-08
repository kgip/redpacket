package service

import (
	"redpacket/service/red_packet"
	"redpacket/service/user"
)

type ServiceGroup struct {
	UserServiceGroup      user.ServiceGroup
	RedPacketServiceGroup red_packet.ServiceGroup
}

var ServiceGroups = &ServiceGroup{}
