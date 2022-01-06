package middlewares

import (
	"math"
	"math/big"
)

var ether = new(big.Float)

func WeiToEther(wei *big.Int) *big.Float {

	ether.SetString(wei.String())
	return new(big.Float).Quo(ether, big.NewFloat(math.Pow10(18)))
}
