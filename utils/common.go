package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// MockServices is a that tolds different mock services to be passed
// around easily for testing setup

// UintToPaddedString converts an int to string of length 19 by padding with 0
func UintToPaddedString(num int64) string {
	return fmt.Sprintf("%019d", num)
}

func GetMarketChannelID(bt, qt common.Address) string {
	return strings.ToLower(fmt.Sprintf("%s::%s", bt.Hex(), qt.Hex()))
}

func PrintJSON(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Print(string(b), "\n")
}

func PrintError(msg string, err error) {
	log.Printf("\n%v: %v\n", msg, err)
}

// Util function to handle unused variables while testing
func Use(...interface{}) {

}
