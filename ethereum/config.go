package ethereum

import (
	"github.com/andyvauliln/prediction-markets-api/utils"
	"github.com/ethereum/go-ethereum/common"
)

var logger = utils.Logger

type EthereumConfig struct {
	url             string
	exchangeAddress common.Address
	wethAddress     common.Address
}

func NewEthereumConfig(url string) *EthereumConfig {
	return &EthereumConfig{
		url: url,
	}
}

func (c *EthereumConfig) GetURL() string {
	return c.url
}

func (c *EthereumConfig) ExchangeAddress() common.Address {
	return c.exchangeAddress
}

func (c *EthereumConfig) WethAddress() common.Address {
	return c.wethAddress
}
