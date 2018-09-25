package ethereum

import (
	"context"
	"math/big"
	"time"

	"github.com/andyvauliln/prediction-markets-api/app"
	"github.com/andyvauliln/prediction-markets-api/interfaces"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthereumProvider struct {
	Client interfaces.EthereumClient
	Config interfaces.EthereumConfig
}

func NewEthereumProvider(c interfaces.EthereumClient) *EthereumProvider {
	url := app.Config.Ethereum["http_url"]
	config := NewEthereumConfig(url)

	return &EthereumProvider{
		Client: c,
		Config: config,
	}
}

func NewDefaultEthereumProvider() *EthereumProvider {
	url := app.Config.Ethereum["http_url"]

	conn, err := rpc.DialHTTP(app.Config.Ethereum["http_url"])
	if err != nil {
		panic(err)
	}

	client := ethclient.NewClient(conn)
	config := NewEthereumConfig(url)

	return &EthereumProvider{
		Client: client,
		Config: config,
	}
}

func NewWebsocketProvider() *EthereumProvider {
	url := app.Config.Ethereum["ws_url"]

	conn, err := rpc.DialWebsocket(context.Background(), url, "")
	if err != nil {
		panic(err)
	}

	client := ethclient.NewClient(conn)
	config := NewEthereumConfig(url)

	return &EthereumProvider{
		Client: client,
		Config: config,
	}
}

func NewSimulatedEthereumProvider(accs []common.Address) *EthereumProvider {
	url := app.Config.Ethereum["http_url"]

	config := NewEthereumConfig(url)
	client := NewSimulatedClient(accs)

	return &EthereumProvider{
		Client: client,
		Config: config,
	}
}

func (e *EthereumProvider) WaitMined(hash common.Hash) (*eth.Receipt, error) {
	ctx := context.Background()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		receipt, _ := e.Client.TransactionReceipt(ctx, hash)
		if receipt != nil {
			return receipt, nil
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
		}
	}
}

func (e *EthereumProvider) GetBalanceAt(a common.Address) (*big.Int, error) {
	ctx := context.Background()
	nonce, err := e.Client.BalanceAt(ctx, a, nil)
	if err != nil {
		logger.Error(err)
		return big.NewInt(0), err
	}

	return nonce, nil
}

func (e *EthereumProvider) GetPendingNonceAt(a common.Address) (uint64, error) {
	ctx := context.Background()
	nonce, err := e.Client.PendingNonceAt(ctx, a)
	if err != nil {
		logger.Error(err)
		return 0, err
	}

	return nonce, nil
}

