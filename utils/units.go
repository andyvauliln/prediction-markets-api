package utils

import (
	"math/big"

	"github.com/andyvauliln/prediction-markets-api/utils/math"
)

func Ethers(value int64) *big.Int {
	return math.Mul(big.NewInt(1e18), big.NewInt(value))
}
