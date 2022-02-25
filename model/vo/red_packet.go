package vo

type SendPacketVo struct {
	Count   int     `json:"count" binding:"required,gt=0"`
	Balance float64 `json:"balance" binding:"required,gt=1"`
}
