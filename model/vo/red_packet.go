package vo

type SendPacketVo struct {
	Count   int     `json:"count" binding:"required,gt=0,lte=100"`
	Balance float64 `json:"balance" binding:"required,gte=10,lte=200"`
}

type RedPacketIdVo struct {
	Id uint `json:"id" binding:"required,gt=0"`
}
