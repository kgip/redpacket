package vo

type SendPacketVo struct {
	Count   int     `json:"count" binding:"required,gt=0,le=100"`
	Balance float64 `json:"balance" binding:"required,gt=1,le=200"`
}

type RedPacketIdVo struct {
	Id uint `json:"id" binding:"required,gt=0"`
}
