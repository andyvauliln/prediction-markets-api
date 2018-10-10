package contracts_test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"

	contracts "github.com/andyvauliln/prediction-markets-api/contracts/augurcontracts"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestMarketes(t *testing.T) {

	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0x990b2d2af7e87cd015a607c3a95d7622c9bbede1")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(3000419),
		ToBlock:   big.NewInt(3065419),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	// Create a new instance of the Inbox contract bound to a specific deployed contract
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.AugurABI)))
	if err != nil {
		log.Fatal(err)
	}

	event := &contracts.AugurMarketCreated{}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		err := contractAbi.Unpack(&event, "MarketCreated", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(event.MinPrice) // foo
		fmt.Println(event.MaxPrice) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("AugurMarketCreated(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex())

}
