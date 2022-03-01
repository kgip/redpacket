package constant

import "time"

const (
	ReturnRedPacketBalanceTopic = "redpacket"
	RedPacketExpireDuration     = time.Hour * 24 * 7 // 7天
)
