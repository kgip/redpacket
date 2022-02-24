package service

import "redpacket/model/po"

type RedPacketService interface {
	List() []*po.RedPacket
}
