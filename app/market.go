package app

import (
	. "github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ReportingState int

const (
	PRE_REPORTING                ReportingState = iota // Market’s end time has not yet come to pass.
	DESIGNATED_REPORTING                               // Market’s end time has occurred, and it is pending a Designated Report.
	OPEN_REPORTING                                     // The Designated Reporter failed to submit a Designated Report within the allotted time, causing the Market to enter the Open Reporting Phase.
	CROWDSOURCING_DISPUTE                              //  An Initial Report for the Market has been submitted, and the Market’s Tentative Outcome is open to being Disputed.
	AWAITING_NEXT_WINDOW                               //  Either the Market had an Initial Report submitted in the current Fee Window, or one of the Market’s Dispute Crowdsourcers received enough REP to Challenge the Market’s Tentative Outcome. In either case, the Market is awaiting the next Fee Window in order to enter another Dispute Round.
	AWAITING_FINALIZATION                              // The Market has been Finalized, but the Post-Finalization Waiting Period has not elapsed.
	FINALIZED                                          // An Outcome for the Market has been determined.
	FORKING                                            // The Dispute Crowdsourcer for one of the Market’s Outcomes received enough REP to reach the Fork Threshold, causing a fork. Users can migrate their REP to the Universe of their choice.
	AWAITING_NO_REPORT_MIGRATION                       // Either the Designated Report was Disputed, or the Designated Reporter failed to submit a Report, and the Market is waiting for the next reporting phase to begin.
	AWAITING_FORK_MIGRATION                            //  Market is waiting for another Market’s Fork to be resolved. This means its Tentative Outcome has been reset to the Outcome submitted in the Initial Report, and all Stake in the Market’s Dispute Crowdsourcers has been refunded to the users who Staked on them.
)

type MarketType int

const (
	YES_NO MarketType = iota
	CATEGORICAL
	SCALAR
)

type NormalizedPayout struct {
	Id     string `json:"id,omitempty"`     //  Market Outcome ID
	Volume string `json:"volume,omitempty"` // Trading volume for this Outcome.
	Price  string `json:"price,omitempty"`
	// Last price at which the outcome was traded. If no trades have taken place in
	// the Market, this value is set to the Market midpoint. If there is no volume on this Outcome,
	// but there is volume on another Outcome in the Market, price is set to 0 for
	// Yes/No Markets and Categorical Markets.
	Description string `json:"Description,omitempty"` // Description for the Outcome.
}

type OutcomeInfo struct {
	IsInvalid bool  `json:"isInvalid,omitempty"` // Whether the Outcome is Invalid.
	Payout    []int `json:"payout,omitempty"`    // Payout Set for the Dispute Crowdsourcer.
}

type Market struct {
	Id                        string           `json:"id,omitempty"`                        // Address of a Market, as a hexadecimal string.
	Universe                  string           `json:"universe,omitempty"`                  // Address of a Universe, as a hexadecimal string
	MarketType                MarketType       `json:"marketType,omitempty"`                // Type of Market (“yesNo”, “categorical”, or “scalar”).
	NumOutcomes               int16            `json:"numOutcomes,omitempty"`               // Total possible Outcomes for the Market.
	MinPrice                  int              `json:"minPrice,omitempty"`                  // Minimum price allowed for a share on a Market, in ETH. For Yes/No & Categorical Markets, this is 0 ETH. For Scalar Markets, this is the bottom end of the range set by the Market creator.
	MaxPrice                  int              `json:"maxPrice,omitempty"`                  // Maximum price allowed for a share on a Market, in ETH. For Yes/No & Categorical Markets, this is 1 ETH. For Scalar Markets, this is the top end of the range set by the Market creator.
	CumulativeScale           string           `json:"cumulativeScale,omitempty"`           // Difference between maxPrice and minPrice.
	Author                    Address          `json:"author,omitempty"`                    // Ethereum address of the creator of the Market, as a hexadecimal string.
	CreationTime              *big.Int         `json:"creationTime,omitempty"`              // Timestamp when the Ethereum block containing the Market creation was created, in seconds.
	CreationBlock             *big.Int         `json:"creationBlock,omitempty"`             // Number of the Ethereum block containing the Market creation.
	CreationFee               string           `json:"creationFee,omitempty"`               // Fee paid by the Market Creator to create the Market, in ETH.
	SettlementFee             string           `json:"settlementFee,omitempty"`             // Fee extracted when a Complete Set is Settled. It is the combination of the Creator Fee and the Reporting Fee.
	ReportingFeeRate          string           `json:"reportingFeeRate,omitempty"`          // Percentage rate of ETH sent to the Fee Window containing the Market whenever shares are settled. Reporting Fees are later used to pay REP holders for Reporting on the Outcome of Markets.
	MarketCreatorFeeRate      string           `json:"marketCreatorFeeRate,omitempty"`      // Percentage rate of ETH paid to the Market creator whenever shares are settled.
	MarketCreatorFeesBalance  string           `json:"marketCreatorFeesBalance,omitempty"`  // Amount of claimable fees the Market creator has not collected from the Market, in ETH.
	MarketCreatorMailbox      string           `json:"marketCreatorMailbox,omitempty"`      // Ethereum address of the Market Creator, as a hexadecimal string.
	MarketCreatorMailboxOwner string           `json:"marketCreatorMailboxOwner,omitempty"` // Ethereum address of the Market Creator Mailbox, as a hexadecimal string.
	InitialReportSize         string           `json:"initialReportSize,omitempty"`         // Size of the No-Show Bond (if the Initial Report was submitted by a First Public Reporter instead of the Designated Reporter).
	Category                  string           `json:"category,omitempty"`                  // Name of the category the Market is in.
	Volume                    string           `json:"volume,omitempty"`                    // Trading volume for this Outcome.
	Tags                      []string         `json:"tags,omitempty"`                      // Names with which the Market has been tagged.
	OpenInterest              string           `json:"openInterest,omitempty"`              // Open interest for the Market.
	OutstandingShares         string           `json:"outstandingShares,omitempty"`         // Number of Complete Sets in the Market.
	ReportingState            ReportingState   `json:"reportingState,omitempty"`            // Reporting state name.
	Forking                   bool             `json:"forking,omitempty"`                   // Whether the Market has Forked.
	NeedsMigration            bool             `json:"needsMigration,omitempty"`            // Whether the Market needs to be migrated to its
	FeeWindow                 string           `json:"feeWindow,omitempty"`                 // Contract address of the Fee Window the Market is in, as a hexadecimal string.
	EndTime                   *big.Int         `json:"endTime,omitempty"`                   // Timestamp when the Market event ends, in seconds.
	FinalizationBlockNumber   *big.Int         `json:"finalizationBlockNumber,omitempty"`   // Ethereum block number in which the Market was Finalized.
	FinalizationTime          *big.Int         `json:"finalizationTime,omitempty"`          // Timestamp when the Market was finalized (if any), in seconds.
	LastTradeTime             *big.Int         `json:"lastTradeTime,omitempty"`             // Unix timestamp when the last trade occurred in this Market.
	LastTradeBlockNumber      *big.Int         `json:"lastTradeBlockNumber,omitempty"`      // Ethereum block number in which the last trade occurred for this Market.
	Description               string           `json:"description,omitempty"`               // Description of the Market.
	Details                   string           `json:"details,omitempty"`                   // Stringified JSON object containing resolutionSource, tags,
	ScalarDenomination        string           `json:"scalarDenomination,omitempty"`        // Denomination used for the numerical range of a Scalar Market (e.g., dollars, degrees Fahrenheit, parts-per-billion).
	DesignatedReporter        string           `json:"designatedReporter,omitempty"`        // Ethereum address of the Market’s designated report, as a hexadecimal string.
	DesignatedReportStake     string           `json:"designatedReportStake,omitempty"`     // Size of the Designated Reporter Stake, in attoETH, that the Designated Reporter must pay to submit the Designated Report for this Market.
	ResolutionSource          string           `json:"resolutionSource,omitempty"`          // Reference source used to determine the Outcome of the Market event.
	NumTicks                  string           `json:"numTicks,omitempty"`                  // Number of possible prices, or ticks, between a Market’s minimum price and maximum price.
	TickSize                  string           `json:"tickSize,omitempty"`                  // Smallest recognized amount by which a price of a security or future may fluctuate in the Market.
	Consensus                 NormalizedPayout `json:"consensus,omitempty"`                 // Consensus Outcome for the Market.
	Outcomes                  []OutcomeInfo    `json:"outcomes,omitempty"`                  // Array of OutcomeInfo objects.
	events                    chan *Event
}
