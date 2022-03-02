package utils

import (
	"github.com/shopspring/decimal"
	"math/rand"
	"redpacket/ex"
	"redpacket/utils/constant"
	"time"
)

// GenericRedPackets 二倍均值法
func GenericRedPackets(count uint, balance float64) []interface{} {
	var isRetry bool
	minBalance := decimal.NewFromFloat(constant.MinSingleRedPacketBalance)
	for retry := 0; retry < 3; retry++ {
		totalAmount := decimal.NewFromFloat(balance)
		var redPackets = make([]interface{}, count)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < int(count-1); i++ {
			limit := totalAmount.Div(decimal.NewFromInt(int64(count)).Sub(decimal.NewFromInt(int64(i)))).Mul(decimal.NewFromInt(2))
			grabBalance := decimal.NewFromFloat(rand.Float64()).Mul(limit)
			if grabBalance.Cmp(minBalance) < 0 {
				grabBalance = minBalance
			}
			redPackets[i], _ = grabBalance.RoundDown(2).Float64()
			totalAmount = totalAmount.Sub(decimal.NewFromFloat(redPackets[i].(float64)))
			if totalAmount.Cmp(minBalance) < 0 {
				isRetry = true
				break
			}
		}
		if isRetry {
			isRetry = false
			continue
		}
		redPackets[count-1], _ = totalAmount.Float64()
		return redPackets
	}
	panic(ex.GenericRedPacketException)
}
