package testutils

import (
	"time"

	"github.com/andyvauliln/prediction-markets-api/ethereum"
	"github.com/andyvauliln/prediction-markets-api/utils/testutils/mocks"
)

func Mine(client *ethereum.SimulatedClient) {
	nextTime := time.Now()
	nextTime = nextTime.Add(500 * time.Millisecond)
	time.Sleep(time.Until(nextTime))

	client.Commit()
	go Mine(client)
}

type MockServices struct {
	MarketService   *mocks.MarketService
	EthereumService *mocks.EthereumService
}

type MockDaos struct {
	MarketDao *mocks.MarketDao
}

func NewMockServices() *MockServices {
	return &MockServices{
		MarketService:   new(mocks.MarketService),
		EthereumService: new(mocks.EthereumService),
	}
}

func NewMockDaos() *MockDaos {
	return &MockDaos{
		MarketDao: new(mocks.MarketDao),
	}
}
