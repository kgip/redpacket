package v1

import (
	"redpacket/router/v1/red_packet"
	"redpacket/router/v1/user"
)

type RouterGroup struct {
	RedPacketRouterGroup red_packet.RouterGroup
	UserRouterGroup      user.UserRouter
}

var RouterGroups = &RouterGroup{}
