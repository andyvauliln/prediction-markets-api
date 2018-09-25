package testutils

import (
	"github.com/andyvauliln/prediction-markets-api/types"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
)

func GetMockMarket() *types.Market {

	return &types.Market{
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
		InitialReportSize:         0,
		Category:                  types.Category{Name: "Money", Popularity: 2},
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
		Consensus: types.NormalizedPayout{
			IsInvalid: false,
			Payout:    []int{1, 23, 14},
		},
		Outcomes: []types.OutcomeInfo{
			{
				ID:          4,
				Volume:      0.123,
				Price:       3.123,
				Description: "test description",
			},
			{
				ID:          3,
				Volume:      0.113,
				Price:       3.1213,
				Description: "test description",
			},
		},
	}
}
