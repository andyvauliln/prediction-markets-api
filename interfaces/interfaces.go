package interfaces

import (
	"context"
	"math/big"

	"github.com/andyvauliln/prediction-markets-api/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/mgo.v2/bson"
)

type MarketService interface {
	Create(market *types.Market) error
	GetByID(id bson.ObjectId) (*types.Market, error)
	GetAll() ([]types.Market, error)
	GetByMarketAddress(marketId common.Address) (*types.Market, error)
}

type MarketDao interface {
	Create(o *types.Market) error
	GetAll() ([]types.Market, error)
	GetByID(id bson.ObjectId) (*types.Market, error)
	GetByMarketAddress(marketID common.Address) (*types.Market, error)
}

type EthereumService interface {
	WaitMined(hash common.Hash) (*eth.Receipt, error)
	GetBalanceAt(a common.Address) (*big.Int, error)
	GetPendingNonceAt(a common.Address) (uint64, error)
}

type EthereumConfig interface {
	GetURL() string
	ExchangeAddress() common.Address
	WethAddress() common.Address
}

type EthereumClient interface {
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*eth.Receipt, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
	SendTransaction(ctx context.Context, tx *eth.Transaction) error
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	BalanceAt(ctx context.Context, contract common.Address, blockNumber *big.Int) (*big.Int, error)
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]eth.Log, error)
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- eth.Log) (ethereum.Subscription, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
}

type EthereumProvider interface {
	WaitMined(hash common.Hash) (*eth.Receipt, error)
	GetBalanceAt(a common.Address) (*big.Int, error)
	GetPendingNonceAt(a common.Address) (uint64, error)
	BalanceOf(owner common.Address, token common.Address) (*big.Int, error)
}
