package v1

import (
	"redpacket/api/v1/red_packet"
	"redpacket/api/v1/user"
)

type ApiGroup struct {
	RedPacketApiGroup red_packet.ApiGroup
	UserApiGroup      user.ApiGroup
}

var ApiGroups = &ApiGroup{}
