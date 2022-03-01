package service

import (
	"redpacket/model/vo"
)

type RedPacketService interface {
	SendPacket(vo *vo.SendPacketVo, id uint)
	GrabPacket(id, userId uint)
}
