package testutils

import (
	"encoding/json"
	"testing"

	"github.com/Proofsuite/amp-matching-engine/types"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert" 
	"gopkg.in/mgo.v2/dbtest"
)

func CompareOrder(t *testing.T, a, b *types.Market) {
	assert.Equal(t, a.ID, b.ID)
	assert.Equal(t, a.MarketID, b.MarketID)
	assert.Equal(t, a.Universe, b.Universe)
	assert.Equal(t, a.MarketType, b.MarketType)
	assert.Equal(t, a.NumOutcomes, b.NumOutcomes)
	assert.Equal(t, a.MinPrice, b.MinPrice)
	assert.Equal(t, a.MaxPrice, b.MaxPrice)
	assert.Equal(t, a.CumulativeScale, b.CumulativeScale)
	assert.Equal(t, a.Author, b.Author)
	assert.Equal(t, a.CreationTime, b.CreationTime)
	assert.Equal(t, a.CreationBlock, b.CreationBlock)
	assert.Equal(t, a.CreationFee, b.CreationFee)
	assert.Equal(t, a.SettlementFee, b.SettlementFee)
	assert.Equal(t, a.ReportingFeeRate, b.ReportingFeeRate)
	assert.Equal(t, a.MarketCreatorFeeRate, b.MarketCreatorFeeRate)
	assert.Equal(t, a.MarketCreatorFeesBalance, b.MarketCreatorFeesBalance)
	assert.Equal(t, a.MarketCreatorMailbox, b.MarketCreatorMailbox)
	assert.Equal(t, a.MarketCreatorMailboxOwner, b.MarketCreatorMailboxOwner)
	assert.Equal(t, a.InitialReportSize, b.InitialReportSize)
	assert.Equal(t, a.Category, b.Category)
	assert.Equal(t, a.Volume, b.Volume)
	assert.Equal(t, a.Tags, b.Tags)
	assert.Equal(t, a.OpenInterest, b.OpenInterest)
	assert.Equal(t, a.OutstandingShares, b.OutstandingShares)
	assert.Equal(t, a.ReportingState, b.ReportingState)
	assert.Equal(t, a.Forking, b.Forking)
	assert.Equal(t, a.NeedsMigration, b.NeedsMigration)
	assert.Equal(t, a.FeeWindow, b.FeeWindow)
	assert.Equal(t, a.EndTime, b.EndTime)
	assert.Equal(t, a.FinalizationBlockNumber, b.FinalizationBlockNumber)
	assert.Equal(t, a.FinalizationTime, b.FinalizationTime)
	assert.Equal(t, a.LastTradeTime, b.LastTradeTime)
	assert.Equal(t, a.LastTradeBlockNumber, b.LastTradeBlockNumber)
	assert.Equal(t, a.Description, b.Description)
	assert.Equal(t, a.Details, b.Details)
	assert.Equal(t, a.ScalarDenomination, b.ScalarDenomination)
	assert.Equal(t, a.DesignatedReporter, b.DesignatedReporter)
	assert.Equal(t, a.DesignatedReportStake, b.DesignatedReportStake)
	assert.Equal(t, a.ResolutionSource, b.ResolutionSource)
	assert.Equal(t, a.NumTicks, b.NumTicks)
	assert.Equal(t, a.TickSize, b.TickSize)
	assert.Equal(t, a.DisputeRounds, b.DisputeRounds)
	assert.Equal(t, a.Consensus, b.Consensus)
	assert.Equal(t, a.Outcomes, b.Outcomes)

}

func Compare(t *testing.T, expected interface{}, value interface{}) {
	expectedBytes, _ := json.Marshal(expected)
	bytes, _ := json.Marshal(value)

	assert.JSONEqf(t, string(expectedBytes), string(bytes), "")
}

func CompareStructs(t *testing.T, expected interface{}, order interface{}) {
	diff := deep.Equal(expected, order)
	if diff != nil {
		t.Errorf("\n%+v\nGot: \n%+v\n\n", expected, order)
	}
}

func NewDBTestServer() *dbtest.DBServer {
	return &dbtest.DBServer{}
}
