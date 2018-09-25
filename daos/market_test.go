package daos

import (
	"io/ioutil"
	"testing"

	"github.com/andyvauliln/prediction-markets-api/types"
	"github.com/andyvauliln/prediction-markets-api/utils/testutils"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	temp, _ := ioutil.TempDir("", "test")
	server.SetPath(temp)

	session := server.Session()
	db = &Database{session}
}

func TestMarketDao(t *testing.T) {
	dao := NewMarketDao()

	market := &types.Market{
		ID:                        bson.NewObjectId(),
		MarketID:                  common.HexToAddress("0x9ad256e6681a7b3c7df29dac0c27f5b4c1d84ae7"),
		Universe:                  common.HexToAddress("0x02149d40d255fceac54a3ee3899807b0539bad60"),
		MarketType:                "yesNo",
		NumOutcomes:               2,
		MinPrice:                  0,
		MaxPrice:                  1,
		CumulativeScale:           1,
		Author:                    common.HexToAddress("0x113b462d14c542d208f5262d82e2eafd7cffd88a"),
		CreationTime:              2687212,
		CreationBlock:             2687212,
		CreationFee:               0.01,
		SettlementFee:             0.0001,
		ReportingFeeRate:          0.0001,
		MarketCreatorFeeRate:      0.01,
		MarketCreatorFeesBalance:  0,
		MarketCreatorMailbox:      common.HexToAddress("0x5acf8f0dcf807dcff537d7c034e020eafea533b7"),
		MarketCreatorMailboxOwner: common.HexToAddress("0x5acf8f0dcf807dcff537d7c034e020eafea533b7"),
		InitialReportSize:         nil,
		Category:                  "climate",
		Volume:                    0.015,
		Tags:                      []string{"Antarctica", "Warming"},
		OpenInterest:              0.012,
		OutstandingShares:         0.012,
		ReportingState:            "PRE_REPORTING",
		Forking:                   false,
		NeedsMigration:            false,
		FeeWindow:                 common.HexToAddress("0x0000000000000000000000000000000000000000"),
		EndTime:                   1575187200,
		FinalizationBlockNumber:   1575183200,
		FinalizationTime:          1575185200,
		LastTradeTime:             1575186200,
		LastTradeBlockNumber:      2825832,
		Description:               "Will the Larsen B ice shelf collapse by the end of November 2019?",
		Details:                   "Long description",
		ScalarDenomination:        "degrees Fahrenheit",
		DesignatedReporter:        common.HexToAddress("0x113b462d14c542d208f5262d82e2eafd7cffd88a"),
		DesignatedReportStake:     0.349680582682291667,
		ResolutionSource:          "tut.by",
		NumTicks:                  10000,
		TickSize:                  0.0001,
		DisputeRounds:             0,
		Consensus: NormalizedPayout{
			IsInvalid: 1,
			Payout:    []int{1, 23, 14},
		},
		Outcomes: []OutcomeInfo{{
			ID:          4,
			Volume:      0.123,
			Price:       3.123,
			Description: "test description",
		}, {
			ID:          3,
			Volume:      0.113,
			Price:       3.1213,
			Description: "test description",
		}},
	}

	err := dao.Create(market)
	if err != nil {
		t.Errorf("Could not create market object: %+v", err)
	}

	all, err := dao.GetAll()
	if err != nil {
		t.Errorf("Could not get markets: %+v", err)
	}

	testutils.CompareMarket(t, market, &all[0])

	byID, err := dao.GetByID(market.ID)
	if err != nil {
		t.Errorf("Could not get market by ID: %v", err)
	}

	testutils.CompareMarket(t, market, byID)

}
