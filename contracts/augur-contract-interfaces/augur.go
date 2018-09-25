package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// IAugurABI is the input ABI used to generate the binding from.
const IAugurABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"name\":\"_repReceived\",\"type\":\"uint256\"},{\"name\":\"_reportingFeesReceived\",\"type\":\"uint256\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"logDisputeCrowdsourcerRedeemed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_extraInfo\",\"type\":\"string\"},{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_marketCreator\",\"type\":\"address\"},{\"name\":\"_outcomes\",\"type\":\"bytes32[]\"},{\"name\":\"_minPrice\",\"type\":\"int256\"},{\"name\":\"_maxPrice\",\"type\":\"int256\"},{\"name\":\"_marketType\",\"type\":\"uint8\"}],\"name\":\"logMarketCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"uint8\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_creator\",\"type\":\"address\"},{\"name\":\"_moneyEscrowed\",\"type\":\"uint256\"},{\"name\":\"_sharesEscrowed\",\"type\":\"uint256\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_shareToken\",\"type\":\"address\"}],\"name\":\"logOrderCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerTokensBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"logReportingParticipantDisavowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_originalUniverse\",\"type\":\"address\"}],\"name\":\"logMarketMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"logMarketTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logFeeTokenMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"logMarketFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"logMarketMailboxTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logFeeTokenTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"name\":\"_reportingFeesReceived\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowRedeemed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logReputationTokenBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"logUniverseForked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_disputeCrowdsourcer\",\"type\":\"address\"}],\"name\":\"logDisputeCrowdsourcerCompleted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerTokensTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_disputeCrowdsourcer\",\"type\":\"address\"},{\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\"}],\"name\":\"logDisputeCrowdsourcerContribution\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_shareToken\",\"type\":\"address\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_numShares\",\"type\":\"uint256\"},{\"name\":\"_numPayoutTokens\",\"type\":\"uint256\"},{\"name\":\"_finalTokenBalance\",\"type\":\"uint256\"}],\"name\":\"logTradingProceedsClaimed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_shareToken\",\"type\":\"address\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"name\":\"_orderType\",\"type\":\"uint8\"},{\"name\":\"_tokenRefund\",\"type\":\"uint256\"},{\"name\":\"_sharesRefund\",\"type\":\"uint256\"}],\"name\":\"logOrderCanceled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logReputationTokenMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_disputeCrowdsourcer\",\"type\":\"address\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_size\",\"type\":\"uint256\"},{\"name\":\"_invalid\",\"type\":\"bool\"}],\"name\":\"disputeCrowdsourcerCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_shareToken\",\"type\":\"address\"},{\"name\":\"_filler\",\"type\":\"address\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"name\":\"_numCreatorShares\",\"type\":\"uint256\"},{\"name\":\"_numCreatorTokens\",\"type\":\"uint256\"},{\"name\":\"_numFillerShares\",\"type\":\"uint256\"},{\"name\":\"_numFillerTokens\",\"type\":\"uint256\"},{\"name\":\"_marketCreatorFees\",\"type\":\"uint256\"},{\"name\":\"_reporterFees\",\"type\":\"uint256\"},{\"name\":\"_amountFilled\",\"type\":\"uint256\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"logOrderFilled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logShareTokensTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"},{\"name\":\"_parentPayoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_parentInvalid\",\"type\":\"bool\"}],\"name\":\"createChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"isKnownUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_feeWindow\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logFeeTokenBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"name\":\"_isDesignatedReporter\",\"type\":\"bool\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_invalid\",\"type\":\"bool\"},{\"name\":\"description\",\"type\":\"string\"}],\"name\":\"logInitialReportSubmitted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logShareTokenBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logShareTokenMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerTokensMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_extraInfo\",\"type\":\"string\"},{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_marketCreator\",\"type\":\"address\"},{\"name\":\"_minPrice\",\"type\":\"int256\"},{\"name\":\"_maxPrice\",\"type\":\"int256\"},{\"name\":\"_marketType\",\"type\":\"uint8\"}],\"name\":\"logMarketCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_numCompleteSets\",\"type\":\"uint256\"}],\"name\":\"logCompleteSetsPurchased\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"logMarketParticipantsDisavowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newTimestamp\",\"type\":\"uint256\"}],\"name\":\"logTimestampSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"logInitialReporterTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"name\":\"_repReceived\",\"type\":\"uint256\"},{\"name\":\"_reportingFeesReceived\",\"type\":\"uint256\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"logInitialReporterRedeemed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"trustedTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logReputationTokensTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_numCompleteSets\",\"type\":\"uint256\"}],\"name\":\"logCompleteSetsSold\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAugurBin is the compiled bytecode used for deploying new contracts.
const IAugurBin = `0x`

// DeployIAugur deploys a new Ethereum contract, binding an instance of IAugur to it.
func DeployIAugur(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IAugur, error) {
	parsed, err := abi.JSON(strings.NewReader(IAugurABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IAugurBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IAugur{IAugurCaller: IAugurCaller{contract: contract}, IAugurTransactor: IAugurTransactor{contract: contract}, IAugurFilterer: IAugurFilterer{contract: contract}}, nil
}

// IAugur is an auto generated Go binding around an Ethereum contract.
type IAugur struct {
	IAugurCaller     // Read-only binding to the contract
	IAugurTransactor // Write-only binding to the contract
	IAugurFilterer   // Log filterer for contract events
}

// IAugurCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAugurCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAugurTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAugurTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAugurFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAugurFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAugurSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAugurSession struct {
	Contract     *IAugur           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAugurCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAugurCallerSession struct {
	Contract *IAugurCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAugurTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAugurTransactorSession struct {
	Contract     *IAugurTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAugurRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAugurRaw struct {
	Contract *IAugur // Generic contract binding to access the raw methods on
}

// IAugurCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAugurCallerRaw struct {
	Contract *IAugurCaller // Generic read-only contract binding to access the raw methods on
}

// IAugurTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAugurTransactorRaw struct {
	Contract *IAugurTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAugur creates a new instance of IAugur, bound to a specific deployed contract.
func NewIAugur(address common.Address, backend bind.ContractBackend) (*IAugur, error) {
	contract, err := bindIAugur(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAugur{IAugurCaller: IAugurCaller{contract: contract}, IAugurTransactor: IAugurTransactor{contract: contract}, IAugurFilterer: IAugurFilterer{contract: contract}}, nil
}

// NewIAugurCaller creates a new read-only instance of IAugur, bound to a specific deployed contract.
func NewIAugurCaller(address common.Address, caller bind.ContractCaller) (*IAugurCaller, error) {
	contract, err := bindIAugur(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAugurCaller{contract: contract}, nil
}

// NewIAugurTransactor creates a new write-only instance of IAugur, bound to a specific deployed contract.
func NewIAugurTransactor(address common.Address, transactor bind.ContractTransactor) (*IAugurTransactor, error) {
	contract, err := bindIAugur(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAugurTransactor{contract: contract}, nil
}

// NewIAugurFilterer creates a new log filterer instance of IAugur, bound to a specific deployed contract.
func NewIAugurFilterer(address common.Address, filterer bind.ContractFilterer) (*IAugurFilterer, error) {
	contract, err := bindIAugur(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAugurFilterer{contract: contract}, nil
}

// bindIAugur binds a generic wrapper to an already deployed contract.
func bindIAugur(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAugurABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAugur *IAugurRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IAugur.Contract.IAugurCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAugur *IAugurRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAugur.Contract.IAugurTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAugur *IAugurRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAugur.Contract.IAugurTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAugur *IAugurCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IAugur.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAugur *IAugurTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAugur.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAugur *IAugurTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAugur.Contract.contract.Transact(opts, method, params...)
}

// IsKnownUniverse is a free data retrieval call binding the contract method 0x8cfb8f21.
//
// Solidity: function isKnownUniverse(_universe address) constant returns(bool)
func (_IAugur *IAugurCaller) IsKnownUniverse(opts *bind.CallOpts, _universe common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IAugur.contract.Call(opts, out, "isKnownUniverse", _universe)
	return *ret0, err
}

// IsKnownUniverse is a free data retrieval call binding the contract method 0x8cfb8f21.
//
// Solidity: function isKnownUniverse(_universe address) constant returns(bool)
func (_IAugur *IAugurSession) IsKnownUniverse(_universe common.Address) (bool, error) {
	return _IAugur.Contract.IsKnownUniverse(&_IAugur.CallOpts, _universe)
}

// IsKnownUniverse is a free data retrieval call binding the contract method 0x8cfb8f21.
//
// Solidity: function isKnownUniverse(_universe address) constant returns(bool)
func (_IAugur *IAugurCallerSession) IsKnownUniverse(_universe common.Address) (bool, error) {
	return _IAugur.Contract.IsKnownUniverse(&_IAugur.CallOpts, _universe)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x8892bb73.
//
// Solidity: function createChildUniverse(_parentPayoutDistributionHash bytes32, _parentPayoutNumerators uint256[], _parentInvalid bool) returns(address)
func (_IAugur *IAugurTransactor) CreateChildUniverse(opts *bind.TransactOpts, _parentPayoutDistributionHash [32]byte, _parentPayoutNumerators []*big.Int, _parentInvalid bool) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "createChildUniverse", _parentPayoutDistributionHash, _parentPayoutNumerators, _parentInvalid)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x8892bb73.
//
// Solidity: function createChildUniverse(_parentPayoutDistributionHash bytes32, _parentPayoutNumerators uint256[], _parentInvalid bool) returns(address)
func (_IAugur *IAugurSession) CreateChildUniverse(_parentPayoutDistributionHash [32]byte, _parentPayoutNumerators []*big.Int, _parentInvalid bool) (*types.Transaction, error) {
	return _IAugur.Contract.CreateChildUniverse(&_IAugur.TransactOpts, _parentPayoutDistributionHash, _parentPayoutNumerators, _parentInvalid)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x8892bb73.
//
// Solidity: function createChildUniverse(_parentPayoutDistributionHash bytes32, _parentPayoutNumerators uint256[], _parentInvalid bool) returns(address)
func (_IAugur *IAugurTransactorSession) CreateChildUniverse(_parentPayoutDistributionHash [32]byte, _parentPayoutNumerators []*big.Int, _parentInvalid bool) (*types.Transaction, error) {
	return _IAugur.Contract.CreateChildUniverse(&_IAugur.TransactOpts, _parentPayoutDistributionHash, _parentPayoutNumerators, _parentInvalid)
}

// DisputeCrowdsourcerCreated is a paid mutator transaction binding the contract method 0x7d4c7806.
//
// Solidity: function disputeCrowdsourcerCreated(_universe address, _market address, _disputeCrowdsourcer address, _payoutNumerators uint256[], _size uint256, _invalid bool) returns(bool)
func (_IAugur *IAugurTransactor) DisputeCrowdsourcerCreated(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _size *big.Int, _invalid bool) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "disputeCrowdsourcerCreated", _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _size, _invalid)
}

// DisputeCrowdsourcerCreated is a paid mutator transaction binding the contract method 0x7d4c7806.
//
// Solidity: function disputeCrowdsourcerCreated(_universe address, _market address, _disputeCrowdsourcer address, _payoutNumerators uint256[], _size uint256, _invalid bool) returns(bool)
func (_IAugur *IAugurSession) DisputeCrowdsourcerCreated(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _size *big.Int, _invalid bool) (*types.Transaction, error) {
	return _IAugur.Contract.DisputeCrowdsourcerCreated(&_IAugur.TransactOpts, _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _size, _invalid)
}

// DisputeCrowdsourcerCreated is a paid mutator transaction binding the contract method 0x7d4c7806.
//
// Solidity: function disputeCrowdsourcerCreated(_universe address, _market address, _disputeCrowdsourcer address, _payoutNumerators uint256[], _size uint256, _invalid bool) returns(bool)
func (_IAugur *IAugurTransactorSession) DisputeCrowdsourcerCreated(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _size *big.Int, _invalid bool) (*types.Transaction, error) {
	return _IAugur.Contract.DisputeCrowdsourcerCreated(&_IAugur.TransactOpts, _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _size, _invalid)
}

// LogCompleteSetsPurchased is a paid mutator transaction binding the contract method 0xc509d0b2.
//
// Solidity: function logCompleteSetsPurchased(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogCompleteSetsPurchased(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logCompleteSetsPurchased", _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsPurchased is a paid mutator transaction binding the contract method 0xc509d0b2.
//
// Solidity: function logCompleteSetsPurchased(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_IAugur *IAugurSession) LogCompleteSetsPurchased(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogCompleteSetsPurchased(&_IAugur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsPurchased is a paid mutator transaction binding the contract method 0xc509d0b2.
//
// Solidity: function logCompleteSetsPurchased(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogCompleteSetsPurchased(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogCompleteSetsPurchased(&_IAugur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsSold is a paid mutator transaction binding the contract method 0xed654fd7.
//
// Solidity: function logCompleteSetsSold(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogCompleteSetsSold(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logCompleteSetsSold", _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsSold is a paid mutator transaction binding the contract method 0xed654fd7.
//
// Solidity: function logCompleteSetsSold(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_IAugur *IAugurSession) LogCompleteSetsSold(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogCompleteSetsSold(&_IAugur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsSold is a paid mutator transaction binding the contract method 0xed654fd7.
//
// Solidity: function logCompleteSetsSold(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogCompleteSetsSold(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogCompleteSetsSold(&_IAugur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogDisputeCrowdsourcerCompleted is a paid mutator transaction binding the contract method 0x4c0019c3.
//
// Solidity: function logDisputeCrowdsourcerCompleted(_universe address, _market address, _disputeCrowdsourcer address) returns(bool)
func (_IAugur *IAugurTransactor) LogDisputeCrowdsourcerCompleted(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logDisputeCrowdsourcerCompleted", _universe, _market, _disputeCrowdsourcer)
}

// LogDisputeCrowdsourcerCompleted is a paid mutator transaction binding the contract method 0x4c0019c3.
//
// Solidity: function logDisputeCrowdsourcerCompleted(_universe address, _market address, _disputeCrowdsourcer address) returns(bool)
func (_IAugur *IAugurSession) LogDisputeCrowdsourcerCompleted(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerCompleted(&_IAugur.TransactOpts, _universe, _market, _disputeCrowdsourcer)
}

// LogDisputeCrowdsourcerCompleted is a paid mutator transaction binding the contract method 0x4c0019c3.
//
// Solidity: function logDisputeCrowdsourcerCompleted(_universe address, _market address, _disputeCrowdsourcer address) returns(bool)
func (_IAugur *IAugurTransactorSession) LogDisputeCrowdsourcerCompleted(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerCompleted(&_IAugur.TransactOpts, _universe, _market, _disputeCrowdsourcer)
}

// LogDisputeCrowdsourcerContribution is a paid mutator transaction binding the contract method 0x5f872147.
//
// Solidity: function logDisputeCrowdsourcerContribution(_universe address, _reporter address, _market address, _disputeCrowdsourcer address, _amountStaked uint256, description string) returns(bool)
func (_IAugur *IAugurTransactor) LogDisputeCrowdsourcerContribution(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _amountStaked *big.Int, description string) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logDisputeCrowdsourcerContribution", _universe, _reporter, _market, _disputeCrowdsourcer, _amountStaked, description)
}

// LogDisputeCrowdsourcerContribution is a paid mutator transaction binding the contract method 0x5f872147.
//
// Solidity: function logDisputeCrowdsourcerContribution(_universe address, _reporter address, _market address, _disputeCrowdsourcer address, _amountStaked uint256, description string) returns(bool)
func (_IAugur *IAugurSession) LogDisputeCrowdsourcerContribution(_universe common.Address, _reporter common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _amountStaked *big.Int, description string) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerContribution(&_IAugur.TransactOpts, _universe, _reporter, _market, _disputeCrowdsourcer, _amountStaked, description)
}

// LogDisputeCrowdsourcerContribution is a paid mutator transaction binding the contract method 0x5f872147.
//
// Solidity: function logDisputeCrowdsourcerContribution(_universe address, _reporter address, _market address, _disputeCrowdsourcer address, _amountStaked uint256, description string) returns(bool)
func (_IAugur *IAugurTransactorSession) LogDisputeCrowdsourcerContribution(_universe common.Address, _reporter common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _amountStaked *big.Int, description string) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerContribution(&_IAugur.TransactOpts, _universe, _reporter, _market, _disputeCrowdsourcer, _amountStaked, description)
}

// LogDisputeCrowdsourcerRedeemed is a paid mutator transaction binding the contract method 0x06ba8e42.
//
// Solidity: function logDisputeCrowdsourcerRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_IAugur *IAugurTransactor) LogDisputeCrowdsourcerRedeemed(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logDisputeCrowdsourcerRedeemed", _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogDisputeCrowdsourcerRedeemed is a paid mutator transaction binding the contract method 0x06ba8e42.
//
// Solidity: function logDisputeCrowdsourcerRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_IAugur *IAugurSession) LogDisputeCrowdsourcerRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerRedeemed(&_IAugur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogDisputeCrowdsourcerRedeemed is a paid mutator transaction binding the contract method 0x06ba8e42.
//
// Solidity: function logDisputeCrowdsourcerRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_IAugur *IAugurTransactorSession) LogDisputeCrowdsourcerRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerRedeemed(&_IAugur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogDisputeCrowdsourcerTokensBurned is a paid mutator transaction binding the contract method 0x10561361.
//
// Solidity: function logDisputeCrowdsourcerTokensBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogDisputeCrowdsourcerTokensBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logDisputeCrowdsourcerTokensBurned", _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensBurned is a paid mutator transaction binding the contract method 0x10561361.
//
// Solidity: function logDisputeCrowdsourcerTokensBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogDisputeCrowdsourcerTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerTokensBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensBurned is a paid mutator transaction binding the contract method 0x10561361.
//
// Solidity: function logDisputeCrowdsourcerTokensBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogDisputeCrowdsourcerTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerTokensBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensMinted is a paid mutator transaction binding the contract method 0xb14823e1.
//
// Solidity: function logDisputeCrowdsourcerTokensMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogDisputeCrowdsourcerTokensMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logDisputeCrowdsourcerTokensMinted", _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensMinted is a paid mutator transaction binding the contract method 0xb14823e1.
//
// Solidity: function logDisputeCrowdsourcerTokensMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogDisputeCrowdsourcerTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerTokensMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensMinted is a paid mutator transaction binding the contract method 0xb14823e1.
//
// Solidity: function logDisputeCrowdsourcerTokensMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogDisputeCrowdsourcerTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerTokensMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensTransferred is a paid mutator transaction binding the contract method 0x50bd5cb9.
//
// Solidity: function logDisputeCrowdsourcerTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogDisputeCrowdsourcerTokensTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logDisputeCrowdsourcerTokensTransferred", _universe, _from, _to, _value)
}

// LogDisputeCrowdsourcerTokensTransferred is a paid mutator transaction binding the contract method 0x50bd5cb9.
//
// Solidity: function logDisputeCrowdsourcerTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurSession) LogDisputeCrowdsourcerTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerTokensTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogDisputeCrowdsourcerTokensTransferred is a paid mutator transaction binding the contract method 0x50bd5cb9.
//
// Solidity: function logDisputeCrowdsourcerTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogDisputeCrowdsourcerTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogDisputeCrowdsourcerTokensTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogFeeTokenBurned is a paid mutator transaction binding the contract method 0x979141ea.
//
// Solidity: function logFeeTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogFeeTokenBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logFeeTokenBurned", _universe, _target, _amount)
}

// LogFeeTokenBurned is a paid mutator transaction binding the contract method 0x979141ea.
//
// Solidity: function logFeeTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogFeeTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeTokenBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogFeeTokenBurned is a paid mutator transaction binding the contract method 0x979141ea.
//
// Solidity: function logFeeTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogFeeTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeTokenBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogFeeTokenMinted is a paid mutator transaction binding the contract method 0x2698eec9.
//
// Solidity: function logFeeTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogFeeTokenMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logFeeTokenMinted", _universe, _target, _amount)
}

// LogFeeTokenMinted is a paid mutator transaction binding the contract method 0x2698eec9.
//
// Solidity: function logFeeTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogFeeTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeTokenMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogFeeTokenMinted is a paid mutator transaction binding the contract method 0x2698eec9.
//
// Solidity: function logFeeTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogFeeTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeTokenMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogFeeTokenTransferred is a paid mutator transaction binding the contract method 0x37227c07.
//
// Solidity: function logFeeTokenTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogFeeTokenTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logFeeTokenTransferred", _universe, _from, _to, _value)
}

// LogFeeTokenTransferred is a paid mutator transaction binding the contract method 0x37227c07.
//
// Solidity: function logFeeTokenTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurSession) LogFeeTokenTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeTokenTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogFeeTokenTransferred is a paid mutator transaction binding the contract method 0x37227c07.
//
// Solidity: function logFeeTokenTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogFeeTokenTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeTokenTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogFeeWindowBurned is a paid mutator transaction binding the contract method 0x542e9b18.
//
// Solidity: function logFeeWindowBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogFeeWindowBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logFeeWindowBurned", _universe, _target, _amount)
}

// LogFeeWindowBurned is a paid mutator transaction binding the contract method 0x542e9b18.
//
// Solidity: function logFeeWindowBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogFeeWindowBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogFeeWindowBurned is a paid mutator transaction binding the contract method 0x542e9b18.
//
// Solidity: function logFeeWindowBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogFeeWindowBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogFeeWindowCreated is a paid mutator transaction binding the contract method 0x8d1b2afd.
//
// Solidity: function logFeeWindowCreated(_feeWindow address, _id uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogFeeWindowCreated(opts *bind.TransactOpts, _feeWindow common.Address, _id *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logFeeWindowCreated", _feeWindow, _id)
}

// LogFeeWindowCreated is a paid mutator transaction binding the contract method 0x8d1b2afd.
//
// Solidity: function logFeeWindowCreated(_feeWindow address, _id uint256) returns(bool)
func (_IAugur *IAugurSession) LogFeeWindowCreated(_feeWindow common.Address, _id *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowCreated(&_IAugur.TransactOpts, _feeWindow, _id)
}

// LogFeeWindowCreated is a paid mutator transaction binding the contract method 0x8d1b2afd.
//
// Solidity: function logFeeWindowCreated(_feeWindow address, _id uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogFeeWindowCreated(_feeWindow common.Address, _id *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowCreated(&_IAugur.TransactOpts, _feeWindow, _id)
}

// LogFeeWindowMinted is a paid mutator transaction binding the contract method 0x60fe103e.
//
// Solidity: function logFeeWindowMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogFeeWindowMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logFeeWindowMinted", _universe, _target, _amount)
}

// LogFeeWindowMinted is a paid mutator transaction binding the contract method 0x60fe103e.
//
// Solidity: function logFeeWindowMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogFeeWindowMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogFeeWindowMinted is a paid mutator transaction binding the contract method 0x60fe103e.
//
// Solidity: function logFeeWindowMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogFeeWindowMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogFeeWindowRedeemed is a paid mutator transaction binding the contract method 0x3b186bfc.
//
// Solidity: function logFeeWindowRedeemed(_universe address, _reporter address, _amountRedeemed uint256, _reportingFeesReceived uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogFeeWindowRedeemed(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _amountRedeemed *big.Int, _reportingFeesReceived *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logFeeWindowRedeemed", _universe, _reporter, _amountRedeemed, _reportingFeesReceived)
}

// LogFeeWindowRedeemed is a paid mutator transaction binding the contract method 0x3b186bfc.
//
// Solidity: function logFeeWindowRedeemed(_universe address, _reporter address, _amountRedeemed uint256, _reportingFeesReceived uint256) returns(bool)
func (_IAugur *IAugurSession) LogFeeWindowRedeemed(_universe common.Address, _reporter common.Address, _amountRedeemed *big.Int, _reportingFeesReceived *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowRedeemed(&_IAugur.TransactOpts, _universe, _reporter, _amountRedeemed, _reportingFeesReceived)
}

// LogFeeWindowRedeemed is a paid mutator transaction binding the contract method 0x3b186bfc.
//
// Solidity: function logFeeWindowRedeemed(_universe address, _reporter address, _amountRedeemed uint256, _reportingFeesReceived uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogFeeWindowRedeemed(_universe common.Address, _reporter common.Address, _amountRedeemed *big.Int, _reportingFeesReceived *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowRedeemed(&_IAugur.TransactOpts, _universe, _reporter, _amountRedeemed, _reportingFeesReceived)
}

// LogFeeWindowTransferred is a paid mutator transaction binding the contract method 0x788873ea.
//
// Solidity: function logFeeWindowTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogFeeWindowTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logFeeWindowTransferred", _universe, _from, _to, _value)
}

// LogFeeWindowTransferred is a paid mutator transaction binding the contract method 0x788873ea.
//
// Solidity: function logFeeWindowTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurSession) LogFeeWindowTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogFeeWindowTransferred is a paid mutator transaction binding the contract method 0x788873ea.
//
// Solidity: function logFeeWindowTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogFeeWindowTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogFeeWindowTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogInitialReportSubmitted is a paid mutator transaction binding the contract method 0x9a0d59aa.
//
// Solidity: function logInitialReportSubmitted(_universe address, _reporter address, _market address, _amountStaked uint256, _isDesignatedReporter bool, _payoutNumerators uint256[], _invalid bool, description string) returns(bool)
func (_IAugur *IAugurTransactor) LogInitialReportSubmitted(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _amountStaked *big.Int, _isDesignatedReporter bool, _payoutNumerators []*big.Int, _invalid bool, description string) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logInitialReportSubmitted", _universe, _reporter, _market, _amountStaked, _isDesignatedReporter, _payoutNumerators, _invalid, description)
}

// LogInitialReportSubmitted is a paid mutator transaction binding the contract method 0x9a0d59aa.
//
// Solidity: function logInitialReportSubmitted(_universe address, _reporter address, _market address, _amountStaked uint256, _isDesignatedReporter bool, _payoutNumerators uint256[], _invalid bool, description string) returns(bool)
func (_IAugur *IAugurSession) LogInitialReportSubmitted(_universe common.Address, _reporter common.Address, _market common.Address, _amountStaked *big.Int, _isDesignatedReporter bool, _payoutNumerators []*big.Int, _invalid bool, description string) (*types.Transaction, error) {
	return _IAugur.Contract.LogInitialReportSubmitted(&_IAugur.TransactOpts, _universe, _reporter, _market, _amountStaked, _isDesignatedReporter, _payoutNumerators, _invalid, description)
}

// LogInitialReportSubmitted is a paid mutator transaction binding the contract method 0x9a0d59aa.
//
// Solidity: function logInitialReportSubmitted(_universe address, _reporter address, _market address, _amountStaked uint256, _isDesignatedReporter bool, _payoutNumerators uint256[], _invalid bool, description string) returns(bool)
func (_IAugur *IAugurTransactorSession) LogInitialReportSubmitted(_universe common.Address, _reporter common.Address, _market common.Address, _amountStaked *big.Int, _isDesignatedReporter bool, _payoutNumerators []*big.Int, _invalid bool, description string) (*types.Transaction, error) {
	return _IAugur.Contract.LogInitialReportSubmitted(&_IAugur.TransactOpts, _universe, _reporter, _market, _amountStaked, _isDesignatedReporter, _payoutNumerators, _invalid, description)
}

// LogInitialReporterRedeemed is a paid mutator transaction binding the contract method 0xec18e2f1.
//
// Solidity: function logInitialReporterRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_IAugur *IAugurTransactor) LogInitialReporterRedeemed(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logInitialReporterRedeemed", _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogInitialReporterRedeemed is a paid mutator transaction binding the contract method 0xec18e2f1.
//
// Solidity: function logInitialReporterRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_IAugur *IAugurSession) LogInitialReporterRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogInitialReporterRedeemed(&_IAugur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogInitialReporterRedeemed is a paid mutator transaction binding the contract method 0xec18e2f1.
//
// Solidity: function logInitialReporterRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_IAugur *IAugurTransactorSession) LogInitialReporterRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogInitialReporterRedeemed(&_IAugur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogInitialReporterTransferred is a paid mutator transaction binding the contract method 0xe3142e90.
//
// Solidity: function logInitialReporterTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_IAugur *IAugurTransactor) LogInitialReporterTransferred(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logInitialReporterTransferred", _universe, _market, _from, _to)
}

// LogInitialReporterTransferred is a paid mutator transaction binding the contract method 0xe3142e90.
//
// Solidity: function logInitialReporterTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_IAugur *IAugurSession) LogInitialReporterTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogInitialReporterTransferred(&_IAugur.TransactOpts, _universe, _market, _from, _to)
}

// LogInitialReporterTransferred is a paid mutator transaction binding the contract method 0xe3142e90.
//
// Solidity: function logInitialReporterTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_IAugur *IAugurTransactorSession) LogInitialReporterTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogInitialReporterTransferred(&_IAugur.TransactOpts, _universe, _market, _from, _to)
}

// LogMarketCreated is a paid mutator transaction binding the contract method 0xbc339f41.
//
// Solidity: function logMarketCreated(_topic bytes32, _description string, _extraInfo string, _universe address, _market address, _marketCreator address, _minPrice int256, _maxPrice int256, _marketType uint8) returns(bool)
func (_IAugur *IAugurTransactor) LogMarketCreated(opts *bind.TransactOpts, _topic [32]byte, _description string, _extraInfo string, _universe common.Address, _market common.Address, _marketCreator common.Address, _minPrice *big.Int, _maxPrice *big.Int, _marketType uint8) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logMarketCreated", _topic, _description, _extraInfo, _universe, _market, _marketCreator, _minPrice, _maxPrice, _marketType)
}

// LogMarketCreated is a paid mutator transaction binding the contract method 0xbc339f41.
//
// Solidity: function logMarketCreated(_topic bytes32, _description string, _extraInfo string, _universe address, _market address, _marketCreator address, _minPrice int256, _maxPrice int256, _marketType uint8) returns(bool)
func (_IAugur *IAugurSession) LogMarketCreated(_topic [32]byte, _description string, _extraInfo string, _universe common.Address, _market common.Address, _marketCreator common.Address, _minPrice *big.Int, _maxPrice *big.Int, _marketType uint8) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketCreated(&_IAugur.TransactOpts, _topic, _description, _extraInfo, _universe, _market, _marketCreator, _minPrice, _maxPrice, _marketType)
}

// LogMarketCreated is a paid mutator transaction binding the contract method 0xbc339f41.
//
// Solidity: function logMarketCreated(_topic bytes32, _description string, _extraInfo string, _universe address, _market address, _marketCreator address, _minPrice int256, _maxPrice int256, _marketType uint8) returns(bool)
func (_IAugur *IAugurTransactorSession) LogMarketCreated(_topic [32]byte, _description string, _extraInfo string, _universe common.Address, _market common.Address, _marketCreator common.Address, _minPrice *big.Int, _maxPrice *big.Int, _marketType uint8) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketCreated(&_IAugur.TransactOpts, _topic, _description, _extraInfo, _universe, _market, _marketCreator, _minPrice, _maxPrice, _marketType)
}

// LogMarketFinalized is a paid mutator transaction binding the contract method 0x27816ffc.
//
// Solidity: function logMarketFinalized(_universe address) returns(bool)
func (_IAugur *IAugurTransactor) LogMarketFinalized(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logMarketFinalized", _universe)
}

// LogMarketFinalized is a paid mutator transaction binding the contract method 0x27816ffc.
//
// Solidity: function logMarketFinalized(_universe address) returns(bool)
func (_IAugur *IAugurSession) LogMarketFinalized(_universe common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketFinalized(&_IAugur.TransactOpts, _universe)
}

// LogMarketFinalized is a paid mutator transaction binding the contract method 0x27816ffc.
//
// Solidity: function logMarketFinalized(_universe address) returns(bool)
func (_IAugur *IAugurTransactorSession) LogMarketFinalized(_universe common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketFinalized(&_IAugur.TransactOpts, _universe)
}

// LogMarketMailboxTransferred is a paid mutator transaction binding the contract method 0x339594f9.
//
// Solidity: function logMarketMailboxTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_IAugur *IAugurTransactor) LogMarketMailboxTransferred(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logMarketMailboxTransferred", _universe, _market, _from, _to)
}

// LogMarketMailboxTransferred is a paid mutator transaction binding the contract method 0x339594f9.
//
// Solidity: function logMarketMailboxTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_IAugur *IAugurSession) LogMarketMailboxTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketMailboxTransferred(&_IAugur.TransactOpts, _universe, _market, _from, _to)
}

// LogMarketMailboxTransferred is a paid mutator transaction binding the contract method 0x339594f9.
//
// Solidity: function logMarketMailboxTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_IAugur *IAugurTransactorSession) LogMarketMailboxTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketMailboxTransferred(&_IAugur.TransactOpts, _universe, _market, _from, _to)
}

// LogMarketMigrated is a paid mutator transaction binding the contract method 0x17674e4d.
//
// Solidity: function logMarketMigrated(_market address, _originalUniverse address) returns(bool)
func (_IAugur *IAugurTransactor) LogMarketMigrated(opts *bind.TransactOpts, _market common.Address, _originalUniverse common.Address) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logMarketMigrated", _market, _originalUniverse)
}

// LogMarketMigrated is a paid mutator transaction binding the contract method 0x17674e4d.
//
// Solidity: function logMarketMigrated(_market address, _originalUniverse address) returns(bool)
func (_IAugur *IAugurSession) LogMarketMigrated(_market common.Address, _originalUniverse common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketMigrated(&_IAugur.TransactOpts, _market, _originalUniverse)
}

// LogMarketMigrated is a paid mutator transaction binding the contract method 0x17674e4d.
//
// Solidity: function logMarketMigrated(_market address, _originalUniverse address) returns(bool)
func (_IAugur *IAugurTransactorSession) LogMarketMigrated(_market common.Address, _originalUniverse common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketMigrated(&_IAugur.TransactOpts, _market, _originalUniverse)
}

// LogMarketParticipantsDisavowed is a paid mutator transaction binding the contract method 0xc67af5cc.
//
// Solidity: function logMarketParticipantsDisavowed(_universe address) returns(bool)
func (_IAugur *IAugurTransactor) LogMarketParticipantsDisavowed(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logMarketParticipantsDisavowed", _universe)
}

// LogMarketParticipantsDisavowed is a paid mutator transaction binding the contract method 0xc67af5cc.
//
// Solidity: function logMarketParticipantsDisavowed(_universe address) returns(bool)
func (_IAugur *IAugurSession) LogMarketParticipantsDisavowed(_universe common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketParticipantsDisavowed(&_IAugur.TransactOpts, _universe)
}

// LogMarketParticipantsDisavowed is a paid mutator transaction binding the contract method 0xc67af5cc.
//
// Solidity: function logMarketParticipantsDisavowed(_universe address) returns(bool)
func (_IAugur *IAugurTransactorSession) LogMarketParticipantsDisavowed(_universe common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketParticipantsDisavowed(&_IAugur.TransactOpts, _universe)
}

// LogMarketTransferred is a paid mutator transaction binding the contract method 0x23290737.
//
// Solidity: function logMarketTransferred(_universe address, _from address, _to address) returns(bool)
func (_IAugur *IAugurTransactor) LogMarketTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logMarketTransferred", _universe, _from, _to)
}

// LogMarketTransferred is a paid mutator transaction binding the contract method 0x23290737.
//
// Solidity: function logMarketTransferred(_universe address, _from address, _to address) returns(bool)
func (_IAugur *IAugurSession) LogMarketTransferred(_universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketTransferred(&_IAugur.TransactOpts, _universe, _from, _to)
}

// LogMarketTransferred is a paid mutator transaction binding the contract method 0x23290737.
//
// Solidity: function logMarketTransferred(_universe address, _from address, _to address) returns(bool)
func (_IAugur *IAugurTransactorSession) LogMarketTransferred(_universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogMarketTransferred(&_IAugur.TransactOpts, _universe, _from, _to)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0x6e1636bb.
//
// Solidity: function logOrderCanceled(_universe address, _shareToken address, _sender address, _orderId bytes32, _orderType uint8, _tokenRefund uint256, _sharesRefund uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogOrderCanceled(opts *bind.TransactOpts, _universe common.Address, _shareToken common.Address, _sender common.Address, _orderId [32]byte, _orderType uint8, _tokenRefund *big.Int, _sharesRefund *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logOrderCanceled", _universe, _shareToken, _sender, _orderId, _orderType, _tokenRefund, _sharesRefund)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0x6e1636bb.
//
// Solidity: function logOrderCanceled(_universe address, _shareToken address, _sender address, _orderId bytes32, _orderType uint8, _tokenRefund uint256, _sharesRefund uint256) returns(bool)
func (_IAugur *IAugurSession) LogOrderCanceled(_universe common.Address, _shareToken common.Address, _sender common.Address, _orderId [32]byte, _orderType uint8, _tokenRefund *big.Int, _sharesRefund *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogOrderCanceled(&_IAugur.TransactOpts, _universe, _shareToken, _sender, _orderId, _orderType, _tokenRefund, _sharesRefund)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0x6e1636bb.
//
// Solidity: function logOrderCanceled(_universe address, _shareToken address, _sender address, _orderId bytes32, _orderType uint8, _tokenRefund uint256, _sharesRefund uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogOrderCanceled(_universe common.Address, _shareToken common.Address, _sender common.Address, _orderId [32]byte, _orderType uint8, _tokenRefund *big.Int, _sharesRefund *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogOrderCanceled(&_IAugur.TransactOpts, _universe, _shareToken, _sender, _orderId, _orderType, _tokenRefund, _sharesRefund)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x0ae41574.
//
// Solidity: function logOrderCreated(_orderType uint8, _amount uint256, _price uint256, _creator address, _moneyEscrowed uint256, _sharesEscrowed uint256, _tradeGroupId bytes32, _orderId bytes32, _universe address, _shareToken address) returns(bool)
func (_IAugur *IAugurTransactor) LogOrderCreated(opts *bind.TransactOpts, _orderType uint8, _amount *big.Int, _price *big.Int, _creator common.Address, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _tradeGroupId [32]byte, _orderId [32]byte, _universe common.Address, _shareToken common.Address) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logOrderCreated", _orderType, _amount, _price, _creator, _moneyEscrowed, _sharesEscrowed, _tradeGroupId, _orderId, _universe, _shareToken)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x0ae41574.
//
// Solidity: function logOrderCreated(_orderType uint8, _amount uint256, _price uint256, _creator address, _moneyEscrowed uint256, _sharesEscrowed uint256, _tradeGroupId bytes32, _orderId bytes32, _universe address, _shareToken address) returns(bool)
func (_IAugur *IAugurSession) LogOrderCreated(_orderType uint8, _amount *big.Int, _price *big.Int, _creator common.Address, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _tradeGroupId [32]byte, _orderId [32]byte, _universe common.Address, _shareToken common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogOrderCreated(&_IAugur.TransactOpts, _orderType, _amount, _price, _creator, _moneyEscrowed, _sharesEscrowed, _tradeGroupId, _orderId, _universe, _shareToken)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x0ae41574.
//
// Solidity: function logOrderCreated(_orderType uint8, _amount uint256, _price uint256, _creator address, _moneyEscrowed uint256, _sharesEscrowed uint256, _tradeGroupId bytes32, _orderId bytes32, _universe address, _shareToken address) returns(bool)
func (_IAugur *IAugurTransactorSession) LogOrderCreated(_orderType uint8, _amount *big.Int, _price *big.Int, _creator common.Address, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _tradeGroupId [32]byte, _orderId [32]byte, _universe common.Address, _shareToken common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogOrderCreated(&_IAugur.TransactOpts, _orderType, _amount, _price, _creator, _moneyEscrowed, _sharesEscrowed, _tradeGroupId, _orderId, _universe, _shareToken)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x80d5398e.
//
// Solidity: function logOrderFilled(_universe address, _shareToken address, _filler address, _orderId bytes32, _numCreatorShares uint256, _numCreatorTokens uint256, _numFillerShares uint256, _numFillerTokens uint256, _marketCreatorFees uint256, _reporterFees uint256, _amountFilled uint256, _tradeGroupId bytes32) returns(bool)
func (_IAugur *IAugurTransactor) LogOrderFilled(opts *bind.TransactOpts, _universe common.Address, _shareToken common.Address, _filler common.Address, _orderId [32]byte, _numCreatorShares *big.Int, _numCreatorTokens *big.Int, _numFillerShares *big.Int, _numFillerTokens *big.Int, _marketCreatorFees *big.Int, _reporterFees *big.Int, _amountFilled *big.Int, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logOrderFilled", _universe, _shareToken, _filler, _orderId, _numCreatorShares, _numCreatorTokens, _numFillerShares, _numFillerTokens, _marketCreatorFees, _reporterFees, _amountFilled, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x80d5398e.
//
// Solidity: function logOrderFilled(_universe address, _shareToken address, _filler address, _orderId bytes32, _numCreatorShares uint256, _numCreatorTokens uint256, _numFillerShares uint256, _numFillerTokens uint256, _marketCreatorFees uint256, _reporterFees uint256, _amountFilled uint256, _tradeGroupId bytes32) returns(bool)
func (_IAugur *IAugurSession) LogOrderFilled(_universe common.Address, _shareToken common.Address, _filler common.Address, _orderId [32]byte, _numCreatorShares *big.Int, _numCreatorTokens *big.Int, _numFillerShares *big.Int, _numFillerTokens *big.Int, _marketCreatorFees *big.Int, _reporterFees *big.Int, _amountFilled *big.Int, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _IAugur.Contract.LogOrderFilled(&_IAugur.TransactOpts, _universe, _shareToken, _filler, _orderId, _numCreatorShares, _numCreatorTokens, _numFillerShares, _numFillerTokens, _marketCreatorFees, _reporterFees, _amountFilled, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x80d5398e.
//
// Solidity: function logOrderFilled(_universe address, _shareToken address, _filler address, _orderId bytes32, _numCreatorShares uint256, _numCreatorTokens uint256, _numFillerShares uint256, _numFillerTokens uint256, _marketCreatorFees uint256, _reporterFees uint256, _amountFilled uint256, _tradeGroupId bytes32) returns(bool)
func (_IAugur *IAugurTransactorSession) LogOrderFilled(_universe common.Address, _shareToken common.Address, _filler common.Address, _orderId [32]byte, _numCreatorShares *big.Int, _numCreatorTokens *big.Int, _numFillerShares *big.Int, _numFillerTokens *big.Int, _marketCreatorFees *big.Int, _reporterFees *big.Int, _amountFilled *big.Int, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _IAugur.Contract.LogOrderFilled(&_IAugur.TransactOpts, _universe, _shareToken, _filler, _orderId, _numCreatorShares, _numCreatorTokens, _numFillerShares, _numFillerTokens, _marketCreatorFees, _reporterFees, _amountFilled, _tradeGroupId)
}

// LogReportingParticipantDisavowed is a paid mutator transaction binding the contract method 0x17570e80.
//
// Solidity: function logReportingParticipantDisavowed(_universe address, _market address) returns(bool)
func (_IAugur *IAugurTransactor) LogReportingParticipantDisavowed(opts *bind.TransactOpts, _universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logReportingParticipantDisavowed", _universe, _market)
}

// LogReportingParticipantDisavowed is a paid mutator transaction binding the contract method 0x17570e80.
//
// Solidity: function logReportingParticipantDisavowed(_universe address, _market address) returns(bool)
func (_IAugur *IAugurSession) LogReportingParticipantDisavowed(_universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogReportingParticipantDisavowed(&_IAugur.TransactOpts, _universe, _market)
}

// LogReportingParticipantDisavowed is a paid mutator transaction binding the contract method 0x17570e80.
//
// Solidity: function logReportingParticipantDisavowed(_universe address, _market address) returns(bool)
func (_IAugur *IAugurTransactorSession) LogReportingParticipantDisavowed(_universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _IAugur.Contract.LogReportingParticipantDisavowed(&_IAugur.TransactOpts, _universe, _market)
}

// LogReputationTokenBurned is a paid mutator transaction binding the contract method 0x4405a339.
//
// Solidity: function logReputationTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogReputationTokenBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logReputationTokenBurned", _universe, _target, _amount)
}

// LogReputationTokenBurned is a paid mutator transaction binding the contract method 0x4405a339.
//
// Solidity: function logReputationTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogReputationTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogReputationTokenBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogReputationTokenBurned is a paid mutator transaction binding the contract method 0x4405a339.
//
// Solidity: function logReputationTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogReputationTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogReputationTokenBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogReputationTokenMinted is a paid mutator transaction binding the contract method 0x79fff7a9.
//
// Solidity: function logReputationTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogReputationTokenMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logReputationTokenMinted", _universe, _target, _amount)
}

// LogReputationTokenMinted is a paid mutator transaction binding the contract method 0x79fff7a9.
//
// Solidity: function logReputationTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogReputationTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogReputationTokenMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogReputationTokenMinted is a paid mutator transaction binding the contract method 0x79fff7a9.
//
// Solidity: function logReputationTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogReputationTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogReputationTokenMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogReputationTokensTransferred is a paid mutator transaction binding the contract method 0xec37a6e4.
//
// Solidity: function logReputationTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogReputationTokensTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logReputationTokensTransferred", _universe, _from, _to, _value)
}

// LogReputationTokensTransferred is a paid mutator transaction binding the contract method 0xec37a6e4.
//
// Solidity: function logReputationTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurSession) LogReputationTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogReputationTokensTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogReputationTokensTransferred is a paid mutator transaction binding the contract method 0xec37a6e4.
//
// Solidity: function logReputationTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogReputationTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogReputationTokensTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogShareTokenBurned is a paid mutator transaction binding the contract method 0xa1b7887f.
//
// Solidity: function logShareTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogShareTokenBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logShareTokenBurned", _universe, _target, _amount)
}

// LogShareTokenBurned is a paid mutator transaction binding the contract method 0xa1b7887f.
//
// Solidity: function logShareTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogShareTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogShareTokenBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogShareTokenBurned is a paid mutator transaction binding the contract method 0xa1b7887f.
//
// Solidity: function logShareTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogShareTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogShareTokenBurned(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogShareTokenMinted is a paid mutator transaction binding the contract method 0xa1dfe545.
//
// Solidity: function logShareTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogShareTokenMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logShareTokenMinted", _universe, _target, _amount)
}

// LogShareTokenMinted is a paid mutator transaction binding the contract method 0xa1dfe545.
//
// Solidity: function logShareTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) LogShareTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogShareTokenMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogShareTokenMinted is a paid mutator transaction binding the contract method 0xa1dfe545.
//
// Solidity: function logShareTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogShareTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogShareTokenMinted(&_IAugur.TransactOpts, _universe, _target, _amount)
}

// LogShareTokensTransferred is a paid mutator transaction binding the contract method 0x86b9a1f4.
//
// Solidity: function logShareTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogShareTokensTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logShareTokensTransferred", _universe, _from, _to, _value)
}

// LogShareTokensTransferred is a paid mutator transaction binding the contract method 0x86b9a1f4.
//
// Solidity: function logShareTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurSession) LogShareTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogShareTokensTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogShareTokensTransferred is a paid mutator transaction binding the contract method 0x86b9a1f4.
//
// Solidity: function logShareTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogShareTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogShareTokensTransferred(&_IAugur.TransactOpts, _universe, _from, _to, _value)
}

// LogTimestampSet is a paid mutator transaction binding the contract method 0xc8e6b2a8.
//
// Solidity: function logTimestampSet(_newTimestamp uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogTimestampSet(opts *bind.TransactOpts, _newTimestamp *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logTimestampSet", _newTimestamp)
}

// LogTimestampSet is a paid mutator transaction binding the contract method 0xc8e6b2a8.
//
// Solidity: function logTimestampSet(_newTimestamp uint256) returns(bool)
func (_IAugur *IAugurSession) LogTimestampSet(_newTimestamp *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogTimestampSet(&_IAugur.TransactOpts, _newTimestamp)
}

// LogTimestampSet is a paid mutator transaction binding the contract method 0xc8e6b2a8.
//
// Solidity: function logTimestampSet(_newTimestamp uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogTimestampSet(_newTimestamp *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogTimestampSet(&_IAugur.TransactOpts, _newTimestamp)
}

// LogTradingProceedsClaimed is a paid mutator transaction binding the contract method 0x6051fa2c.
//
// Solidity: function logTradingProceedsClaimed(_universe address, _shareToken address, _sender address, _market address, _numShares uint256, _numPayoutTokens uint256, _finalTokenBalance uint256) returns(bool)
func (_IAugur *IAugurTransactor) LogTradingProceedsClaimed(opts *bind.TransactOpts, _universe common.Address, _shareToken common.Address, _sender common.Address, _market common.Address, _numShares *big.Int, _numPayoutTokens *big.Int, _finalTokenBalance *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logTradingProceedsClaimed", _universe, _shareToken, _sender, _market, _numShares, _numPayoutTokens, _finalTokenBalance)
}

// LogTradingProceedsClaimed is a paid mutator transaction binding the contract method 0x6051fa2c.
//
// Solidity: function logTradingProceedsClaimed(_universe address, _shareToken address, _sender address, _market address, _numShares uint256, _numPayoutTokens uint256, _finalTokenBalance uint256) returns(bool)
func (_IAugur *IAugurSession) LogTradingProceedsClaimed(_universe common.Address, _shareToken common.Address, _sender common.Address, _market common.Address, _numShares *big.Int, _numPayoutTokens *big.Int, _finalTokenBalance *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogTradingProceedsClaimed(&_IAugur.TransactOpts, _universe, _shareToken, _sender, _market, _numShares, _numPayoutTokens, _finalTokenBalance)
}

// LogTradingProceedsClaimed is a paid mutator transaction binding the contract method 0x6051fa2c.
//
// Solidity: function logTradingProceedsClaimed(_universe address, _shareToken address, _sender address, _market address, _numShares uint256, _numPayoutTokens uint256, _finalTokenBalance uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) LogTradingProceedsClaimed(_universe common.Address, _shareToken common.Address, _sender common.Address, _market common.Address, _numShares *big.Int, _numPayoutTokens *big.Int, _finalTokenBalance *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.LogTradingProceedsClaimed(&_IAugur.TransactOpts, _universe, _shareToken, _sender, _market, _numShares, _numPayoutTokens, _finalTokenBalance)
}

// LogUniverseForked is a paid mutator transaction binding the contract method 0x4a812023.
//
// Solidity: function logUniverseForked() returns(bool)
func (_IAugur *IAugurTransactor) LogUniverseForked(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "logUniverseForked")
}

// LogUniverseForked is a paid mutator transaction binding the contract method 0x4a812023.
//
// Solidity: function logUniverseForked() returns(bool)
func (_IAugur *IAugurSession) LogUniverseForked() (*types.Transaction, error) {
	return _IAugur.Contract.LogUniverseForked(&_IAugur.TransactOpts)
}

// LogUniverseForked is a paid mutator transaction binding the contract method 0x4a812023.
//
// Solidity: function logUniverseForked() returns(bool)
func (_IAugur *IAugurTransactorSession) LogUniverseForked() (*types.Transaction, error) {
	return _IAugur.Contract.LogUniverseForked(&_IAugur.TransactOpts)
}

// TrustedTransfer is a paid mutator transaction binding the contract method 0xec238994.
//
// Solidity: function trustedTransfer(_token address, _from address, _to address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactor) TrustedTransfer(opts *bind.TransactOpts, _token common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.contract.Transact(opts, "trustedTransfer", _token, _from, _to, _amount)
}

// TrustedTransfer is a paid mutator transaction binding the contract method 0xec238994.
//
// Solidity: function trustedTransfer(_token address, _from address, _to address, _amount uint256) returns(bool)
func (_IAugur *IAugurSession) TrustedTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.TrustedTransfer(&_IAugur.TransactOpts, _token, _from, _to, _amount)
}

// TrustedTransfer is a paid mutator transaction binding the contract method 0xec238994.
//
// Solidity: function trustedTransfer(_token address, _from address, _to address, _amount uint256) returns(bool)
func (_IAugur *IAugurTransactorSession) TrustedTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IAugur.Contract.TrustedTransfer(&_IAugur.TransactOpts, _token, _from, _to, _amount)
}

// AugurABI is the input ABI used to generate the binding from.
const AugurABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"name\":\"_repReceived\",\"type\":\"uint256\"},{\"name\":\"_reportingFeesReceived\",\"type\":\"uint256\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"logDisputeCrowdsourcerRedeemed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_extraInfo\",\"type\":\"string\"},{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_marketCreator\",\"type\":\"address\"},{\"name\":\"_outcomes\",\"type\":\"bytes32[]\"},{\"name\":\"_minPrice\",\"type\":\"int256\"},{\"name\":\"_maxPrice\",\"type\":\"int256\"},{\"name\":\"_marketType\",\"type\":\"uint8\"}],\"name\":\"logMarketCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"uint8\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_creator\",\"type\":\"address\"},{\"name\":\"_moneyEscrowed\",\"type\":\"uint256\"},{\"name\":\"_sharesEscrowed\",\"type\":\"uint256\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_shareToken\",\"type\":\"address\"}],\"name\":\"logOrderCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerTokensBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"logReportingParticipantDisavowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_originalUniverse\",\"type\":\"address\"}],\"name\":\"logMarketMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"logMarketTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logFeeTokenMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"logMarketFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"logMarketMailboxTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logFeeTokenTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"name\":\"_reportingFeesReceived\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowRedeemed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logReputationTokenBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"logUniverseForked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_disputeCrowdsourcer\",\"type\":\"address\"}],\"name\":\"logDisputeCrowdsourcerCompleted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerTokensTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_disputeCrowdsourcer\",\"type\":\"address\"},{\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"logDisputeCrowdsourcerContribution\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_shareToken\",\"type\":\"address\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_numShares\",\"type\":\"uint256\"},{\"name\":\"_numPayoutTokens\",\"type\":\"uint256\"},{\"name\":\"_finalTokenBalance\",\"type\":\"uint256\"}],\"name\":\"logTradingProceedsClaimed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_shareToken\",\"type\":\"address\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"name\":\"_orderType\",\"type\":\"uint8\"},{\"name\":\"_tokenRefund\",\"type\":\"uint256\"},{\"name\":\"_sharesRefund\",\"type\":\"uint256\"}],\"name\":\"logOrderCanceled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logReputationTokenMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_disputeCrowdsourcer\",\"type\":\"address\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_size\",\"type\":\"uint256\"},{\"name\":\"_invalid\",\"type\":\"bool\"}],\"name\":\"disputeCrowdsourcerCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_shareToken\",\"type\":\"address\"},{\"name\":\"_filler\",\"type\":\"address\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"name\":\"_numCreatorShares\",\"type\":\"uint256\"},{\"name\":\"_numCreatorTokens\",\"type\":\"uint256\"},{\"name\":\"_numFillerShares\",\"type\":\"uint256\"},{\"name\":\"_numFillerTokens\",\"type\":\"uint256\"},{\"name\":\"_marketCreatorFees\",\"type\":\"uint256\"},{\"name\":\"_reporterFees\",\"type\":\"uint256\"},{\"name\":\"_amountFilled\",\"type\":\"uint256\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"logOrderFilled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isKnownShareToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logShareTokensTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"},{\"name\":\"_parentPayoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_parentInvalid\",\"type\":\"bool\"}],\"name\":\"createChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"isKnownUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_feeWindow\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"logFeeWindowCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGenesisUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logFeeTokenBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"name\":\"_isDesignatedReporter\",\"type\":\"bool\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_invalid\",\"type\":\"bool\"},{\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"logInitialReportSubmitted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logShareTokenBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logShareTokenMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerTokensMinted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_crowdsourcer\",\"type\":\"address\"}],\"name\":\"isKnownCrowdsourcer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_extraInfo\",\"type\":\"string\"},{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_marketCreator\",\"type\":\"address\"},{\"name\":\"_minPrice\",\"type\":\"int256\"},{\"name\":\"_maxPrice\",\"type\":\"int256\"},{\"name\":\"_marketType\",\"type\":\"uint8\"}],\"name\":\"logMarketCreated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_numCompleteSets\",\"type\":\"uint256\"}],\"name\":\"logCompleteSetsPurchased\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"logMarketParticipantsDisavowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newTimestamp\",\"type\":\"uint256\"}],\"name\":\"logTimestampSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"logInitialReporterTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"name\":\"_repReceived\",\"type\":\"uint256\"},{\"name\":\"_reportingFeesReceived\",\"type\":\"uint256\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"logInitialReporterRedeemed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"trustedTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"logReputationTokensTransferred\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_numCompleteSets\",\"type\":\"uint256\"}],\"name\":\"logCompleteSetsSold\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"extraInfo\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"marketCreator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"outcomes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"marketCreationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minPrice\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"maxPrice\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"marketType\",\"type\":\"uint8\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isDesignatedReporter\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"invalid\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"description\",\"type\":\"string\"}],\"name\":\"InitialReportSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"disputeCrowdsourcer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"invalid\",\"type\":\"bool\"}],\"name\":\"DisputeCrowdsourcerCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"disputeCrowdsourcer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"description\",\"type\":\"string\"}],\"name\":\"DisputeCrowdsourcerContribution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"disputeCrowdsourcer\",\"type\":\"address\"}],\"name\":\"DisputeCrowdsourcerCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"repReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reportingFeesReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"InitialReporterRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"disputeCrowdsourcer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"repReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reportingFeesReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"DisputeCrowdsourcerRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reportingParticipant\",\"type\":\"address\"}],\"name\":\"ReportingParticipantDisavowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketParticipantsDisavowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"feeWindow\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reportingFeesReceived\",\"type\":\"uint256\"}],\"name\":\"FeeWindowRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"originalUniverse\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newUniverse\",\"type\":\"address\"}],\"name\":\"MarketMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"}],\"name\":\"UniverseForked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"parentUniverse\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"childUniverse\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"invalid\",\"type\":\"bool\"}],\"name\":\"UniverseCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"shareToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"tokenRefund\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesRefund\",\"type\":\"uint256\"}],\"name\":\"OrderCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"moneyEscrowed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesEscrowed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tradeGroupId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"shareToken\",\"type\":\"address\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"shareToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"filler\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"numCreatorShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numCreatorTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numFillerShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numFillerTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"marketCreatorFees\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reporterFees\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"numCompleteSets\",\"type\":\"uint256\"}],\"name\":\"CompleteSetsPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"numCompleteSets\",\"type\":\"uint256\"}],\"name\":\"CompleteSetsSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"shareToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"numShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numPayoutTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"finalTokenBalance\",\"type\":\"uint256\"}],\"name\":\"TradingProceedsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"market\",\"type\":\"address\"}],\"name\":\"TokensTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"market\",\"type\":\"address\"}],\"name\":\"TokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"market\",\"type\":\"address\"}],\"name\":\"TokensBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"feeWindow\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"FeeWindowCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"InitialReporterTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"MarketTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"mailbox\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"MarketMailboxTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"isOn\",\"type\":\"bool\"}],\"name\":\"EscapeHatchChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"name\":\"TimestampSet\",\"type\":\"event\"}]"

// AugurBin is the compiled bytecode used for deploying new contracts.
const AugurBin = `0x608060405260008054600160a060020a03191633179055613f8e806100256000396000f30060806040526004361061022c5763ffffffff60e060020a60003504166306ba8e42811461023157806307c1880a146102c35780630ae41574146103c9578063105613611461041457806317570e801461043e57806317674e4d14610465578063232907371461048c5780632698eec9146104b957806327816ffc146104e35780633018205f14610504578063339594f91461053557806337227c07146105685780633b186bfc146105985780634405a339146105c55780634a812023146105ef5780634c0019c31461060457806350bd5cb914610631578063542e9b18146106615780635f8721471461068b5780636051fa2c1461070a57806360fe103e146107465780636e1636bb14610770578063788873ea146107ac57806379fff7a9146107dc5780637d4c78061461080657806380d5398e1461088257806382e79da8146108ce57806386b9a1f4146108ef5780638892bb731461091f5780638cfb8f211461097d5780638d1b2afd1461099e57806392eefe9b146109c25780639684da1a146109e3578063979141ea146109f85780639a0d59aa14610a22578063a1b7887f14610ae1578063a1dfe54514610b0b578063b14823e114610b35578063b70da7dc14610b5f578063bc339f4114610b80578063c509d0b214610c4c578063c67af5cc14610c7c578063c8e6b2a814610c9d578063e3142e9014610cb5578063ec18e2f114610ce8578063ec23899414610d66578063ec37a6e414610d96578063ed654fd714610dc6575b600080fd5b34801561023d57600080fd5b506040805160c435600481810135602081810285810182019096528185526102af95600160a060020a0384358116966024803583169760443590931696606435966084359660a435963696939560e495909401929182919085019084908082843750949750610df69650505050505050565b604080519115158252519081900360200190f35b3480156102cf57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526102af95833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437505060408051602060608901358a01803582810280850184018652818552999c8b35600160a060020a039081169d858e013582169d978801359091169b919a50985060809095019650929450810192829185019084908082843750949750508435955050506020830135926040013560ff169150610ee19050565b3480156103d557600080fd5b506102af60ff60043516602435604435600160a060020a036064358116906084359060a4359060c4359060e43590610104358116906101243516611133565b34801561042057600080fd5b506102af600160a060020a0360043581169060243516604435611261565b34801561044a57600080fd5b506102af600160a060020a0360043581169060243516611357565b34801561047157600080fd5b506102af600160a060020a0360043581169060243516611451565b34801561049857600080fd5b506102af600160a060020a03600435811690602435811690604435166114bc565b3480156104c557600080fd5b506102af600160a060020a03600435811690602435166044356115b3565b3480156104ef57600080fd5b506102af600160a060020a03600435166116d7565b34801561051057600080fd5b506105196117c0565b60408051600160a060020a039092168252519081900360200190f35b34801561054157600080fd5b506102af600160a060020a03600435811690602435811690604435811690606435166117cf565b34801561057457600080fd5b506102af600160a060020a0360043581169060243581169060443516606435611950565b3480156105a457600080fd5b506102af600160a060020a0360043581169060243516604435606435611a8f565b3480156105d157600080fd5b506102af600160a060020a0360043581169060243516604435611b93565b3480156105fb57600080fd5b506102af611c77565b34801561061057600080fd5b506102af600160a060020a0360043581169060243581169060443516611cc6565b34801561063d57600080fd5b506102af600160a060020a0360043581169060243581169060443516606435611db1565b34801561066d57600080fd5b506102af600160a060020a0360043581169060243516604435611eba565b34801561069757600080fd5b50604080516020600460a43581810135601f81018490048402850184019095528484526102af948235600160a060020a039081169560248035831696604435841696606435909416956084359536959460c49490939201918190840183828082843750949750611f9e9650505050505050565b34801561071657600080fd5b506102af600160a060020a036004358116906024358116906044358116906064351660843560a43560c435612120565b34801561075257600080fd5b506102af600160a060020a0360043581169060243516604435612216565b34801561077c57600080fd5b506102af600160a060020a036004358116906024358116906044351660643560ff6084351660a43560c4356122fa565b3480156107b857600080fd5b506102af600160a060020a0360043581169060243581169060443516606435612416565b3480156107e857600080fd5b506102af600160a060020a0360043581169060243516604435612513565b34801561081257600080fd5b5060408051606435600481810135602081810285810182019096528185526102af95600160a060020a0384358116966024803583169760443590931696369690956084959394909201929091829190850190849080828437509497505084359550505050506020013515156125f7565b34801561088e57600080fd5b506102af600160a060020a036004358116906024358116906044351660643560843560a43560c43560e43561010435610124356101443561016435612762565b3480156108da57600080fd5b506102af600160a060020a03600435166128b5565b3480156108fb57600080fd5b506102af600160a060020a03600435811690602435811690604435166064356128d3565b34801561092b57600080fd5b5060408051602060046024803582810135848102808701860190975280865261051996843596369660449591949091019291829185019084908082843750949750505050913515159250612959915050565b34801561098957600080fd5b506102af600160a060020a0360043516612985565b3480156109aa57600080fd5b506102af600160a060020a03600435166024356129a3565b3480156109ce57600080fd5b506102af600160a060020a0360043516612b15565b3480156109ef57600080fd5b50610519612b5f565b348015610a0457600080fd5b506102af600160a060020a0360043581169060243516604435612b83565b348015610a2e57600080fd5b506040805160a435600481810135602081810285810182019096528185526102af95600160a060020a0384358116966024803583169760443590931696606435966084351515963696929560c49594019282918501908490808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b8a3515159b909a909994019750919550918201935091508190840183828082843750949750612c799650505050505050565b348015610aed57600080fd5b506102af600160a060020a0360043581169060243516604435612e41565b348015610b1757600080fd5b506102af600160a060020a0360043581169060243516604435612ec6565b348015610b4157600080fd5b506102af600160a060020a0360043581169060243516604435612f4b565b348015610b6b57600080fd5b506102af600160a060020a0360043516612fd0565b348015610b8c57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526102af95833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975050508335600160a060020a039081169550602085013581169460408101359091169350606081013592506080810135915060a0013560ff16612fee565b348015610c5857600080fd5b506102af600160a060020a036004358116906024358116906044351660643561326a565b348015610c8857600080fd5b506102af600160a060020a0360043516613355565b348015610ca957600080fd5b506102af60043561343e565b348015610cc157600080fd5b506102af600160a060020a036004358116906024358116906044358116906064351661353f565b348015610cf457600080fd5b506040805160c435600481810135602081810285810182019096528185526102af95600160a060020a0384358116966024803583169760443590931696606435966084359660a435963696939560e4959094019291829190850190849080828437509497506136bc9650505050505050565b348015610d7257600080fd5b506102af600160a060020a0360043581169060243581169060443516606435613823565b348015610da257600080fd5b506102af600160a060020a0360043581169060243581169060443516606435613971565b348015610dd257600080fd5b506102af600160a060020a0360043581169060243581169060443516606435613a6e565b600033610e0281612fd0565b1515610e0d57600080fd5b86600160a060020a031688600160a060020a03168a600160a060020a03167f450bd662d3b1e236c8f344457690d257aeae5dca1add336752839ac206613cc0848a8a8a8a6040518086600160a060020a0316600160a060020a0316815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015610ebb578181015183820152602001610ea3565b50505050905001965050505050505060405180910390a450600198975050505050505050565b6000610eec88612985565b1515610ef757600080fd5b600160a060020a0388163314610f0c57600080fd5b610f1587613b59565b5085600160a060020a031688600160a060020a03168c600019167fb2e65de73007eef46316e4f18ab1f301b4d0e31aa56733387b469612f90894df8d8d8c8b8f600160a060020a031663ec86fdbd6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610f9257600080fd5b505af1158015610fa6573d6000803e3d6000fd5b505050506040513d6020811015610fbc57600080fd5b505160408051600160a060020a038516918101919091526080810182905260a081018e905260c081018d90528d908d908d9080602081016060820160e0830185600281111561100757fe5b60ff16815260200184810384528c818151815260200191508051906020019080838360005b8381101561104457818101518382015260200161102c565b50505050905090810190601f1680156110715780820380516001836020036101000a031916815260200191505b5084810383528b5181528b516020918201918d019080838360005b838110156110a457818101518382015260200161108c565b50505050905090810190601f1680156110d15780820380516001836020036101000a031916815260200191505b508481038252895181528951602091820191808c01910280838360005b838110156111065781810151838201526020016110ee565b505050509050019b50505050505050505050505060405180910390a45060019a9950505050505050505050565b6000805460408051600080516020613f4383398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b15801561118757600080fd5b505af115801561119b573d6000803e3d6000fd5b505050506040513d60208110156111b157600080fd5b505115156111be57600080fd5b81600160a060020a031683600160a060020a031689600160a060020a03167f32d554e498d0c7f2a5c7fd8b6b234bfc4e1dfb5290466d998af09a813db32f318e8e8e8d8d8d8d6040518088600181111561121457fe5b60ff1681526020810197909752506040808701959095526060860193909352608085019190915260a084015260c0830152519081900360e0019150a45060019a9950505050505050505050565b60003361126d81612fd0565b151561127857600080fd5b83600160a060020a031633600160a060020a031686600160a060020a0316600080516020613ee383398151915286600286600160a060020a031663f1be16796040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156112e657600080fd5b505af11580156112fa573d6000803e3d6000fd5b505050506040513d602081101561131057600080fd5b50516040518381526020810183600481111561132857fe5b60ff168152600160a060020a0390921660208301525060408051918290030192509050a4506001949350505050565b600061136283612985565b151561136d57600080fd5b604080517ff76514c70000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0385169163f76514c79160248083019260209291908290030181600087803b1580156113ce57600080fd5b505af11580156113e2573d6000803e3d6000fd5b505050506040513d60208110156113f857600080fd5b5051151561140557600080fd5b604080513381529051600160a060020a0380851692908616917fb20adf682c8f82b94a135452f54ac4483c9ee8c9b2324e946120696ab1d034b49181900360200190a350600192915050565b60003361145d81612985565b151561146857600080fd5b80600160a060020a031683600160a060020a031685600160a060020a03167fc3cf07f8fa0fafc25a9dd0bad2cd6b961c55dad41b42c8ef8f931bc40e41e08c60405160405180910390a45060019392505050565b6000806114c885612985565b15156114d357600080fd5b5060408051600080516020613ec3833981519152815233600482018190529151600160a060020a03871691639f7e1bf69160248083019260209291908290030181600087803b15801561152557600080fd5b505af1158015611539573d6000803e3d6000fd5b505050506040513d602081101561154f57600080fd5b5051151561155c57600080fd5b60408051600160a060020a0386811682528581166020830152825181851693918916927f55f2a7bfa32e835c3f3c3cff653a3d11c077ce1b00c5a41c6aaf09eedc1ac3b292908290030190a3506001949350505050565b60006115be84612985565b15156115c957600080fd5b604080517f26d16bc90000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a038616916326d16bc99160248083019260209291908290030181600087803b15801561162a57600080fd5b505af115801561163e573d6000803e3d6000fd5b505050506040513d602081101561165457600080fd5b5051151561166157600080fd5b82600160a060020a031633600160a060020a031685600160a060020a0316600080516020613f238339815191528560046000604051808481526020018360048111156116a957fe5b60ff16815260200182600160a060020a03168152602001935050505060405180910390a45060019392505050565b6000806116e383612985565b15156116ee57600080fd5b5060408051600080516020613ec3833981519152815233600482018190529151600160a060020a03851691639f7e1bf69160248083019260209291908290030181600087803b15801561174057600080fd5b505af1158015611754573d6000803e3d6000fd5b505050506040513d602081101561176a57600080fd5b5051151561177757600080fd5b80600160a060020a031683600160a060020a03167f014ce4e12965529d7d31e11411d7a23b1778d448ab763ffc4d55830cbb4919d760405160405180910390a350600192915050565b600054600160a060020a031690565b60006117da85612985565b15156117e557600080fd5b84600160a060020a0316639f7e1bf6856040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561184057600080fd5b505af1158015611854573d6000803e3d6000fd5b505050506040513d602081101561186a57600080fd5b5051151561187757600080fd5b83600160a060020a031663ed23378b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156118b557600080fd5b505af11580156118c9573d6000803e3d6000fd5b505050506040513d60208110156118df57600080fd5b5051600160a060020a031633146118f557600080fd5b60408051600160a060020a03858116825284811660208301528251339388831693928a16927f8a34ec183bf620d74d0b52e71165bb4255b0591d1c8e9d07c707a7f1d763158d929081900390910190a4506001949350505050565b600061195b85612985565b151561196657600080fd5b604080517f26d16bc90000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a038716916326d16bc99160248083019260209291908290030181600087803b1580156119c757600080fd5b505af11580156119db573d6000803e3d6000fd5b505050506040513d60208110156119f157600080fd5b505115156119fe57600080fd5b83600160a060020a031633600160a060020a031686600160a060020a0316600080516020613ea38339815191528686600460006040518085600160a060020a0316600160a060020a03168152602001848152602001836004811115611a5f57fe5b60ff16815260200182600160a060020a0316815260200194505050505060405180910390a4506001949350505050565b6000611a9a85612985565b1515611aa557600080fd5b60408051600080516020613f0383398151915281523360048201529051600160a060020a0387169163c7c88d709160248083019260209291908290030181600087803b158015611af457600080fd5b505af1158015611b08573d6000803e3d6000fd5b505050506040513d6020811015611b1e57600080fd5b50511515611b2b57600080fd5b33600160a060020a031684600160a060020a031686600160a060020a03167fc62cff53848fe243adb6130140cfe557ce16e8006861abd50adfe425150ba6c58686604051808381526020018281526020019250505060405180910390a4506001949350505050565b6000611b9e84612985565b1515611ba957600080fd5b33600160a060020a031684600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611bf157600080fd5b505af1158015611c05573d6000803e3d6000fd5b505050506040513d6020811015611c1b57600080fd5b5051600160a060020a031614611c3057600080fd5b82600160a060020a031633600160a060020a031685600160a060020a0316600080516020613ee383398151915285600080604051808481526020018360048111156116a957fe5b3360009081526001602052604081205460ff161515611c9557600080fd5b60405133907fd4d990bbdf9b9a4383a394341465060ccb75513432ceee3d5fcd8788ab1a507f90600090a250600190565b6000611cd184612985565b1515611cdc57600080fd5b60408051600080516020613ec383398151915281523360048201529051600160a060020a03861691639f7e1bf69160248083019260209291908290030181600087803b158015611d2b57600080fd5b505af1158015611d3f573d6000803e3d6000fd5b505050506040513d6020811015611d5557600080fd5b50511515611d6257600080fd5b60408051600160a060020a0384811682529151828616928716917fec05f094139821aeb3220a0837f5d14eb02aa619179aadf3b316ed95b3648abb919081900360200190a35060019392505050565b600033611dbd81612fd0565b1515611dc857600080fd5b84600160a060020a031633600160a060020a031687600160a060020a0316600080516020613ea38339815191528787600287600160a060020a031663f1be16796040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611e3757600080fd5b505af1158015611e4b573d6000803e3d6000fd5b505050506040513d6020811015611e6157600080fd5b505160408051600160a060020a038616815260208101859052908101836004811115611e8957fe5b60ff168152600160a060020a039092166020830152506040805191829003019350915050a450600195945050505050565b6000611ec584612985565b1515611ed057600080fd5b60408051600080516020613f0383398151915281523360048201529051600160a060020a0386169163c7c88d709160248083019260209291908290030181600087803b158015611f1f57600080fd5b505af1158015611f33573d6000803e3d6000fd5b505050506040513d6020811015611f4957600080fd5b50511515611f5657600080fd5b82600160a060020a031633600160a060020a031685600160a060020a0316600080516020613ee38339815191528560036000604051808481526020018360048111156116a957fe5b6000611fa987612985565b1515611fb457600080fd5b60408051600080516020613ec383398151915281523360048201529051600160a060020a03891691639f7e1bf69160248083019260209291908290030181600087803b15801561200357600080fd5b505af1158015612017573d6000803e3d6000fd5b505050506040513d602081101561202d57600080fd5b5051151561203a57600080fd5b84600160a060020a031686600160a060020a031688600160a060020a03167f1ba97c2894f2b4eb21d849bdb2c4b2007b3562407a13d5581e8cc603ccfc70aa8787876040518084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156120d75781810151838201526020016120bf565b50505050905090810190601f1680156121045780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a45060019695505050505050565b6000805460408051600080516020613f4383398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b15801561217457600080fd5b505af1158015612188573d6000803e3d6000fd5b505050506040513d602081101561219e57600080fd5b505115156121ab57600080fd5b60408051600160a060020a03878116825260208201879052818301869052606082018590529151828916928a811692908c16917fa7e9373569caad2b7871ecb4d498619fc1c42840a6c0dbeb8dff20b131721e509181900360800190a4506001979650505050505050565b600061222184612985565b151561222c57600080fd5b60408051600080516020613f0383398151915281523360048201529051600160a060020a0386169163c7c88d709160248083019260209291908290030181600087803b15801561227b57600080fd5b505af115801561228f573d6000803e3d6000fd5b505050506040513d60208110156122a557600080fd5b505115156122b257600080fd5b82600160a060020a031633600160a060020a031685600160a060020a0316600080516020613f238339815191528560036000604051808481526020018360048111156116a957fe5b6000805460408051600080516020613f4383398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b15801561234e57600080fd5b505af1158015612362573d6000803e3d6000fd5b505050506040513d602081101561237857600080fd5b5051151561238557600080fd5b85600160a060020a031687600160a060020a031689600160a060020a03167f513d029ff62330c16d8d4b36b28fab53f09d10bb51b56fe121ab710ca2d1af80888888886040518085600019166000191681526020018460018111156123e657fe5b60ff16815260200183815260200182815260200194505050505060405180910390a4506001979650505050505050565b600061242185612985565b151561242c57600080fd5b60408051600080516020613f0383398151915281523360048201529051600160a060020a0387169163c7c88d709160248083019260209291908290030181600087803b15801561247b57600080fd5b505af115801561248f573d6000803e3d6000fd5b505050506040513d60208110156124a557600080fd5b505115156124b257600080fd5b83600160a060020a031633600160a060020a031686600160a060020a0316600080516020613ea38339815191528686600360006040518085600160a060020a0316600160a060020a03168152602001848152602001836004811115611a5f57fe5b600061251e84612985565b151561252957600080fd5b33600160a060020a031684600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561257157600080fd5b505af1158015612585573d6000803e3d6000fd5b505050506040513d602081101561259b57600080fd5b5051600160a060020a0316146125b057600080fd5b82600160a060020a031633600160a060020a031685600160a060020a0316600080516020613f2383398151915285600080604051808481526020018360048111156116a957fe5b600061260287612985565b151561260d57600080fd5b60408051600080516020613ec383398151915281523360048201529051600160a060020a03891691639f7e1bf69160248083019260209291908290030181600087803b15801561265c57600080fd5b505af1158015612670573d6000803e3d6000fd5b505050506040513d602081101561268657600080fd5b5051151561269357600080fd5b600160a060020a038086166000818152600260209081526040808320805460ff19166001179055805193845283018790528515156060840152608083820181815289519185019190915288518b8616958d16947fccc07058358a9411a6acb3cd58bf6d0b398c3ff1f0b2c8e97a6dbdbbe74eae41948c948c948c948c9493919260a085019288820192909102908190849084905b8381101561273f578181015183820152602001612727565b505050509050019550505050505060405180910390a35060019695505050505050565b6000805460408051600080516020613f4383398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b1580156127b657600080fd5b505af11580156127ca573d6000803e3d6000fd5b505050506040513d60208110156127e057600080fd5b505115156127ed57600080fd5b8b600160a060020a03168d600160a060020a03167fabb970462c1f0de9e237d127ad47c01c4e69caa179fd850d076ae9bfc529176e8d8d8d8d8d8d8d8d8d8d604051808b600160a060020a0316600160a060020a031681526020018a6000191660001916815260200189815260200188815260200187815260200186815260200185815260200184815260200183815260200182600019166000191681526020019a505050505050505050505060405180910390a35060019c9b505050505050505050505050565b600160a060020a031660009081526003602052604090205460ff1690565b6000336128df816128b5565b15156128ea57600080fd5b84600160a060020a031633600160a060020a031687600160a060020a0316600080516020613ea38339815191528787600187600160a060020a031663f1be16796040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611e3757600080fd5b60003361296581612985565b151561297057600080fd5b61297c81868686613c89565b95945050505050565b600160a060020a031660009081526001602052604090205460ff1690565b3360009081526001602052604081205460ff1615156129c157600080fd5b33600160a060020a03167fbaba17e31bb9fbfbc0b794111d2b1236ed4e36067a5e0d7c3c3433ad66c99f9d8485600160a060020a031663c828371e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612a2b57600080fd5b505af1158015612a3f573d6000803e3d6000fd5b505050506040513d6020811015612a5557600080fd5b5051604080517f439f5ac20000000000000000000000000000000000000000000000000000000081529051600160a060020a0389169163439f5ac29160048083019260209291908290030181600087803b158015612ab257600080fd5b505af1158015612ac6573d6000803e3d6000fd5b505050506040513d6020811015612adc57600080fd5b505160408051600160a060020a03909416845260208401929092528282015260608201869052519081900360800190a250600192915050565b60008054600160a060020a03163314612b2d57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60408051600080825260208201909252612b7e90829081906000613c89565b905090565b6000612b8e84612985565b1515612b9957600080fd5b604080517f26d16bc90000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a038616916326d16bc99160248083019260209291908290030181600087803b158015612bfa57600080fd5b505af1158015612c0e573d6000803e3d6000fd5b505050506040513d6020811015612c2457600080fd5b50511515612c3157600080fd5b82600160a060020a031633600160a060020a031685600160a060020a0316600080516020613ee38339815191528560046000604051808481526020018360048111156116a957fe5b6000612c8489612985565b1515612c8f57600080fd5b60408051600080516020613ec383398151915281523360048201529051600160a060020a038b1691639f7e1bf69160248083019260209291908290030181600087803b158015612cde57600080fd5b505af1158015612cf2573d6000803e3d6000fd5b505050506040513d6020811015612d0857600080fd5b50511515612d1557600080fd5b86600160a060020a031688600160a060020a03168a600160a060020a03167f5f28d90f610b97bb029fe74dfab936367e76c0fe576309e0b65cf72a975beac489898989896040518086815260200185151515158152602001806020018415151515815260200180602001838103835286818151815260200191508051906020019060200280838360005b83811015612db7578181015183820152602001612d9f565b50505050905001838103825284818151815260200191508051906020019080838360005b83811015612df3578181015183820152602001612ddb565b50505050905090810190601f168015612e205780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a450600198975050505050505050565b600033612e4d816128b5565b1515612e5857600080fd5b83600160a060020a031633600160a060020a031686600160a060020a0316600080516020613ee383398151915286600186600160a060020a031663f1be16796040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156112e657600080fd5b600033612ed2816128b5565b1515612edd57600080fd5b83600160a060020a031633600160a060020a031686600160a060020a0316600080516020613f2383398151915286600186600160a060020a031663f1be16796040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156112e657600080fd5b600033612f5781612fd0565b1515612f6257600080fd5b83600160a060020a031633600160a060020a031686600160a060020a0316600080516020613f2383398151915286600286600160a060020a031663f1be16796040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156112e657600080fd5b600160a060020a031660009081526002602052604090205460ff1690565b6000612ff987612985565b151561300457600080fd5b600160a060020a038716331461301957600080fd5b61302286613b59565b5084600160a060020a031687600160a060020a03168b600019167fb2e65de73007eef46316e4f18ab1f301b4d0e31aa56733387b469612f90894df8c8c8b600060405190808252806020026020018201604052801561308b578160200160208202803883390190505b508e600160a060020a031663ec86fdbd6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156130ca57600080fd5b505af11580156130de573d6000803e3d6000fd5b505050506040513d60208110156130f457600080fd5b505160408051600160a060020a038516918101919091526080810182905260a081018e905260c081018d90528d908d908d9080602081016060820160e0830185600281111561313f57fe5b60ff16815260200184810384528c818151815260200191508051906020019080838360005b8381101561317c578181015183820152602001613164565b50505050905090810190601f1680156131a95780820380516001836020036101000a031916815260200191505b5084810383528b5181528b516020918201918d019080838360005b838110156131dc5781810151838201526020016131c4565b50505050905090810190601f1680156132095780820380516001836020036101000a031916815260200191505b508481038252895181528951602091820191808c01910280838360005b8381101561323e578181015183820152602001613226565b505050509050019b50505050505050505050505060405180910390a45060019998505050505050505050565b6000805460408051600080516020613f4383398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b1580156132be57600080fd5b505af11580156132d2573d6000803e3d6000fd5b505050506040513d60208110156132e857600080fd5b505115156132f557600080fd5b82600160a060020a031684600160a060020a031686600160a060020a03167f349ab20f76ba930a00da1936627d07400af6bb7cd2e2b4c68bcab93ca8aff418856040518082815260200191505060405180910390a4506001949350505050565b60008061336183612985565b151561336c57600080fd5b5060408051600080516020613ec3833981519152815233600482018190529151600160a060020a03851691639f7e1bf69160248083019260209291908290030181600087803b1580156133be57600080fd5b505af11580156133d2573d6000803e3d6000fd5b505050506040513d60208110156133e857600080fd5b505115156133f557600080fd5b80600160a060020a031683600160a060020a03167f3b4f3db017516414df2695e5b0052661779d7163a6cd4368fd74313be73fa0b860405160405180910390a350600192915050565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f54696d650000000000000000000000000000000000000000000000000000000060048201529051600160a060020a039092169163f39ec1f79160248082019260209290919082900301818787803b1580156134c457600080fd5b505af11580156134d8573d6000803e3d6000fd5b505050506040513d60208110156134ee57600080fd5b5051600160a060020a0316331461350457600080fd5b6040805183815290517f11dda748f0bd3af85a073da0088a0acb827d9584a4fdb825c81f1232a53095389181900360200190a1506001919050565b600061354a85612985565b151561355557600080fd5b84600160a060020a0316639f7e1bf6856040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156135b057600080fd5b505af11580156135c4573d6000803e3d6000fd5b505050506040513d60208110156135da57600080fd5b505115156135e757600080fd5b83600160a060020a031663c7600cde6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561362557600080fd5b505af1158015613639573d6000803e3d6000fd5b505050506040513d602081101561364f57600080fd5b5051600160a060020a0316331461366557600080fd5b60408051600160a060020a0385811682528481166020830152825181881693918916927fee62c58e2603b92f96a002e012f4f3bd5748102cfa3b711f6d778c6237fcaa9692908290030190a3506001949350505050565b60006136c788612985565b15156136d257600080fd5b604080517ff76514c70000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a038a169163f76514c79160248083019260209291908290030181600087803b15801561373357600080fd5b505af1158015613747573d6000803e3d6000fd5b505050506040513d602081101561375d57600080fd5b5051151561376a57600080fd5b85600160a060020a031687600160a060020a031689600160a060020a03167fdd0dca2d338dc86ba5431017bdf6f3ad45247d608b0a38d866e3131a876be2cc888888886040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156137ff5781810151838201526020016137e7565b505050509050019550505050505060405180910390a4506001979650505050505050565b6000805460408051600080516020613f4383398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b15801561387757600080fd5b505af115801561388b573d6000803e3d6000fd5b505050506040513d60208110156138a157600080fd5b505115156138ae57600080fd5b600082116138bb57600080fd5b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301528581166024830152604482018590529151918716916323b872dd916064808201926020929091908290030181600087803b15801561392f57600080fd5b505af1158015613943573d6000803e3d6000fd5b505050506040513d602081101561395957600080fd5b5051151561396657600080fd5b506001949350505050565b600061397c85612985565b151561398757600080fd5b33600160a060020a031685600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156139cf57600080fd5b505af11580156139e3573d6000803e3d6000fd5b505050506040513d60208110156139f957600080fd5b5051600160a060020a031614613a0e57600080fd5b83600160a060020a031633600160a060020a031686600160a060020a0316600080516020613ea383398151915286866000806040518085600160a060020a0316600160a060020a03168152602001848152602001836004811115611a5f57fe5b6000805460408051600080516020613f4383398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b158015613ac257600080fd5b505af1158015613ad6573d6000803e3d6000fd5b505050506040513d6020811015613aec57600080fd5b50511515613af957600080fd5b82600160a060020a031684600160a060020a031686600160a060020a03167f68166bb2a567c21899b00209f52c286bf00ac613acc9f183da791ac5f5f47051856040518082815260200191505060405180910390a4506001949350505050565b600080600083600160a060020a03166327ce5b8c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015613b9c57600080fd5b505af1158015613bb0573d6000803e3d6000fd5b505050506040513d6020811015613bc657600080fd5b50519150600090505b81811015613c825760016003600086600160a060020a03166365957bf5856040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b158015613c2657600080fd5b505af1158015613c3a573d6000803e3d6000fd5b505050506040513d6020811015613c5057600080fd5b5051600160a060020a031681526020810191909152604001600020805460ff1916911515919091179055600101613bcf565b5050919050565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f556e697665727365466163746f727900000000000000000000000000000000006004820152905183928392600160a060020a039091169163f39ec1f79160248082019260209290919082900301818787803b158015613d1357600080fd5b505af1158015613d27573d6000803e3d6000fd5b505050506040513d6020811015613d3d57600080fd5b505160008054604080517f4837435f000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201528b83166024820152604481018b9052905193955090851692634837435f92606480840193602093929083900390910190829087803b158015613dbb57600080fd5b505af1158015613dcf573d6000803e3d6000fd5b505050506040513d6020811015613de557600080fd5b5051600160a060020a038082166000818152600160208181526040808420805460ff191690931790925581518a1515818301528281528b51928101929092528a519596509294938c16937f299eaafd0d27519eda3fe7195b73e5269e442b3d80928f19afa32b6db2f352b6938b938b93928392606084019287820192909102908190849084905b83811015613e84578181015183820152602001613e6c565b50505050905001935050505060405180910390a3969550505050505056003c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf9f7e1bf600000000000000000000000000000000000000000000000000000000262b80f2af08a1001d15a1df91dde9acb8441811543886659b3845a8c285748bc7c88d700000000000000000000000000000000000000000000000000000000075dd618f69c0f07adc97fe19ba435f3932ce6aa8cad287fb9bdfaf37639f703a3f08882f00000000000000000000000000000000000000000000000000000000a165627a7a723058201ba837d9fce758ad1fcc38e5c185bb5012ea00e74eb2c7883d7b370e4e70c3b90029`

// DeployAugur deploys a new Ethereum contract, binding an instance of Augur to it.
func DeployAugur(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Augur, error) {
	parsed, err := abi.JSON(strings.NewReader(AugurABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AugurBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Augur{AugurCaller: AugurCaller{contract: contract}, AugurTransactor: AugurTransactor{contract: contract}, AugurFilterer: AugurFilterer{contract: contract}}, nil
}

// Augur is an auto generated Go binding around an Ethereum contract.
type Augur struct {
	AugurCaller     // Read-only binding to the contract
	AugurTransactor // Write-only binding to the contract
	AugurFilterer   // Log filterer for contract events
}

// AugurCaller is an auto generated read-only Go binding around an Ethereum contract.
type AugurCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AugurTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AugurFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AugurSession struct {
	Contract     *Augur            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AugurCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AugurCallerSession struct {
	Contract *AugurCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AugurTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AugurTransactorSession struct {
	Contract     *AugurTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AugurRaw is an auto generated low-level Go binding around an Ethereum contract.
type AugurRaw struct {
	Contract *Augur // Generic contract binding to access the raw methods on
}

// AugurCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AugurCallerRaw struct {
	Contract *AugurCaller // Generic read-only contract binding to access the raw methods on
}

// AugurTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AugurTransactorRaw struct {
	Contract *AugurTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAugur creates a new instance of Augur, bound to a specific deployed contract.
func NewAugur(address common.Address, backend bind.ContractBackend) (*Augur, error) {
	contract, err := bindAugur(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Augur{AugurCaller: AugurCaller{contract: contract}, AugurTransactor: AugurTransactor{contract: contract}, AugurFilterer: AugurFilterer{contract: contract}}, nil
}

// NewAugurCaller creates a new read-only instance of Augur, bound to a specific deployed contract.
func NewAugurCaller(address common.Address, caller bind.ContractCaller) (*AugurCaller, error) {
	contract, err := bindAugur(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AugurCaller{contract: contract}, nil
}

// NewAugurTransactor creates a new write-only instance of Augur, bound to a specific deployed contract.
func NewAugurTransactor(address common.Address, transactor bind.ContractTransactor) (*AugurTransactor, error) {
	contract, err := bindAugur(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AugurTransactor{contract: contract}, nil
}

// NewAugurFilterer creates a new log filterer instance of Augur, bound to a specific deployed contract.
func NewAugurFilterer(address common.Address, filterer bind.ContractFilterer) (*AugurFilterer, error) {
	contract, err := bindAugur(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AugurFilterer{contract: contract}, nil
}

// bindAugur binds a generic wrapper to an already deployed contract.
func bindAugur(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AugurABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Augur *AugurRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Augur.Contract.AugurCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Augur *AugurRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Augur.Contract.AugurTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Augur *AugurRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Augur.Contract.AugurTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Augur *AugurCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Augur.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Augur *AugurTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Augur.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Augur *AugurTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Augur.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Augur *AugurCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Augur.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Augur *AugurSession) GetController() (common.Address, error) {
	return _Augur.Contract.GetController(&_Augur.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Augur *AugurCallerSession) GetController() (common.Address, error) {
	return _Augur.Contract.GetController(&_Augur.CallOpts)
}

// IsKnownCrowdsourcer is a free data retrieval call binding the contract method 0xb70da7dc.
//
// Solidity: function isKnownCrowdsourcer(_crowdsourcer address) constant returns(bool)
func (_Augur *AugurCaller) IsKnownCrowdsourcer(opts *bind.CallOpts, _crowdsourcer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Augur.contract.Call(opts, out, "isKnownCrowdsourcer", _crowdsourcer)
	return *ret0, err
}

// IsKnownCrowdsourcer is a free data retrieval call binding the contract method 0xb70da7dc.
//
// Solidity: function isKnownCrowdsourcer(_crowdsourcer address) constant returns(bool)
func (_Augur *AugurSession) IsKnownCrowdsourcer(_crowdsourcer common.Address) (bool, error) {
	return _Augur.Contract.IsKnownCrowdsourcer(&_Augur.CallOpts, _crowdsourcer)
}

// IsKnownCrowdsourcer is a free data retrieval call binding the contract method 0xb70da7dc.
//
// Solidity: function isKnownCrowdsourcer(_crowdsourcer address) constant returns(bool)
func (_Augur *AugurCallerSession) IsKnownCrowdsourcer(_crowdsourcer common.Address) (bool, error) {
	return _Augur.Contract.IsKnownCrowdsourcer(&_Augur.CallOpts, _crowdsourcer)
}

// IsKnownShareToken is a free data retrieval call binding the contract method 0x82e79da8.
//
// Solidity: function isKnownShareToken(_token address) constant returns(bool)
func (_Augur *AugurCaller) IsKnownShareToken(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Augur.contract.Call(opts, out, "isKnownShareToken", _token)
	return *ret0, err
}

// IsKnownShareToken is a free data retrieval call binding the contract method 0x82e79da8.
//
// Solidity: function isKnownShareToken(_token address) constant returns(bool)
func (_Augur *AugurSession) IsKnownShareToken(_token common.Address) (bool, error) {
	return _Augur.Contract.IsKnownShareToken(&_Augur.CallOpts, _token)
}

// IsKnownShareToken is a free data retrieval call binding the contract method 0x82e79da8.
//
// Solidity: function isKnownShareToken(_token address) constant returns(bool)
func (_Augur *AugurCallerSession) IsKnownShareToken(_token common.Address) (bool, error) {
	return _Augur.Contract.IsKnownShareToken(&_Augur.CallOpts, _token)
}

// IsKnownUniverse is a free data retrieval call binding the contract method 0x8cfb8f21.
//
// Solidity: function isKnownUniverse(_universe address) constant returns(bool)
func (_Augur *AugurCaller) IsKnownUniverse(opts *bind.CallOpts, _universe common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Augur.contract.Call(opts, out, "isKnownUniverse", _universe)
	return *ret0, err
}

// IsKnownUniverse is a free data retrieval call binding the contract method 0x8cfb8f21.
//
// Solidity: function isKnownUniverse(_universe address) constant returns(bool)
func (_Augur *AugurSession) IsKnownUniverse(_universe common.Address) (bool, error) {
	return _Augur.Contract.IsKnownUniverse(&_Augur.CallOpts, _universe)
}

// IsKnownUniverse is a free data retrieval call binding the contract method 0x8cfb8f21.
//
// Solidity: function isKnownUniverse(_universe address) constant returns(bool)
func (_Augur *AugurCallerSession) IsKnownUniverse(_universe common.Address) (bool, error) {
	return _Augur.Contract.IsKnownUniverse(&_Augur.CallOpts, _universe)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x8892bb73.
//
// Solidity: function createChildUniverse(_parentPayoutDistributionHash bytes32, _parentPayoutNumerators uint256[], _parentInvalid bool) returns(address)
func (_Augur *AugurTransactor) CreateChildUniverse(opts *bind.TransactOpts, _parentPayoutDistributionHash [32]byte, _parentPayoutNumerators []*big.Int, _parentInvalid bool) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "createChildUniverse", _parentPayoutDistributionHash, _parentPayoutNumerators, _parentInvalid)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x8892bb73.
//
// Solidity: function createChildUniverse(_parentPayoutDistributionHash bytes32, _parentPayoutNumerators uint256[], _parentInvalid bool) returns(address)
func (_Augur *AugurSession) CreateChildUniverse(_parentPayoutDistributionHash [32]byte, _parentPayoutNumerators []*big.Int, _parentInvalid bool) (*types.Transaction, error) {
	return _Augur.Contract.CreateChildUniverse(&_Augur.TransactOpts, _parentPayoutDistributionHash, _parentPayoutNumerators, _parentInvalid)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x8892bb73.
//
// Solidity: function createChildUniverse(_parentPayoutDistributionHash bytes32, _parentPayoutNumerators uint256[], _parentInvalid bool) returns(address)
func (_Augur *AugurTransactorSession) CreateChildUniverse(_parentPayoutDistributionHash [32]byte, _parentPayoutNumerators []*big.Int, _parentInvalid bool) (*types.Transaction, error) {
	return _Augur.Contract.CreateChildUniverse(&_Augur.TransactOpts, _parentPayoutDistributionHash, _parentPayoutNumerators, _parentInvalid)
}

// CreateGenesisUniverse is a paid mutator transaction binding the contract method 0x9684da1a.
//
// Solidity: function createGenesisUniverse() returns(address)
func (_Augur *AugurTransactor) CreateGenesisUniverse(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "createGenesisUniverse")
}

// CreateGenesisUniverse is a paid mutator transaction binding the contract method 0x9684da1a.
//
// Solidity: function createGenesisUniverse() returns(address)
func (_Augur *AugurSession) CreateGenesisUniverse() (*types.Transaction, error) {
	return _Augur.Contract.CreateGenesisUniverse(&_Augur.TransactOpts)
}

// CreateGenesisUniverse is a paid mutator transaction binding the contract method 0x9684da1a.
//
// Solidity: function createGenesisUniverse() returns(address)
func (_Augur *AugurTransactorSession) CreateGenesisUniverse() (*types.Transaction, error) {
	return _Augur.Contract.CreateGenesisUniverse(&_Augur.TransactOpts)
}

// DisputeCrowdsourcerCreated is a paid mutator transaction binding the contract method 0x7d4c7806.
//
// Solidity: function disputeCrowdsourcerCreated(_universe address, _market address, _disputeCrowdsourcer address, _payoutNumerators uint256[], _size uint256, _invalid bool) returns(bool)
func (_Augur *AugurTransactor) DisputeCrowdsourcerCreated(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _size *big.Int, _invalid bool) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "disputeCrowdsourcerCreated", _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _size, _invalid)
}

// DisputeCrowdsourcerCreated is a paid mutator transaction binding the contract method 0x7d4c7806.
//
// Solidity: function disputeCrowdsourcerCreated(_universe address, _market address, _disputeCrowdsourcer address, _payoutNumerators uint256[], _size uint256, _invalid bool) returns(bool)
func (_Augur *AugurSession) DisputeCrowdsourcerCreated(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _size *big.Int, _invalid bool) (*types.Transaction, error) {
	return _Augur.Contract.DisputeCrowdsourcerCreated(&_Augur.TransactOpts, _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _size, _invalid)
}

// DisputeCrowdsourcerCreated is a paid mutator transaction binding the contract method 0x7d4c7806.
//
// Solidity: function disputeCrowdsourcerCreated(_universe address, _market address, _disputeCrowdsourcer address, _payoutNumerators uint256[], _size uint256, _invalid bool) returns(bool)
func (_Augur *AugurTransactorSession) DisputeCrowdsourcerCreated(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _size *big.Int, _invalid bool) (*types.Transaction, error) {
	return _Augur.Contract.DisputeCrowdsourcerCreated(&_Augur.TransactOpts, _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _size, _invalid)
}

// LogCompleteSetsPurchased is a paid mutator transaction binding the contract method 0xc509d0b2.
//
// Solidity: function logCompleteSetsPurchased(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_Augur *AugurTransactor) LogCompleteSetsPurchased(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logCompleteSetsPurchased", _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsPurchased is a paid mutator transaction binding the contract method 0xc509d0b2.
//
// Solidity: function logCompleteSetsPurchased(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_Augur *AugurSession) LogCompleteSetsPurchased(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogCompleteSetsPurchased(&_Augur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsPurchased is a paid mutator transaction binding the contract method 0xc509d0b2.
//
// Solidity: function logCompleteSetsPurchased(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogCompleteSetsPurchased(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogCompleteSetsPurchased(&_Augur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsSold is a paid mutator transaction binding the contract method 0xed654fd7.
//
// Solidity: function logCompleteSetsSold(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_Augur *AugurTransactor) LogCompleteSetsSold(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logCompleteSetsSold", _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsSold is a paid mutator transaction binding the contract method 0xed654fd7.
//
// Solidity: function logCompleteSetsSold(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_Augur *AugurSession) LogCompleteSetsSold(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogCompleteSetsSold(&_Augur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsSold is a paid mutator transaction binding the contract method 0xed654fd7.
//
// Solidity: function logCompleteSetsSold(_universe address, _market address, _account address, _numCompleteSets uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogCompleteSetsSold(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogCompleteSetsSold(&_Augur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogDisputeCrowdsourcerCompleted is a paid mutator transaction binding the contract method 0x4c0019c3.
//
// Solidity: function logDisputeCrowdsourcerCompleted(_universe address, _market address, _disputeCrowdsourcer address) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerCompleted(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerCompleted", _universe, _market, _disputeCrowdsourcer)
}

// LogDisputeCrowdsourcerCompleted is a paid mutator transaction binding the contract method 0x4c0019c3.
//
// Solidity: function logDisputeCrowdsourcerCompleted(_universe address, _market address, _disputeCrowdsourcer address) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerCompleted(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerCompleted(&_Augur.TransactOpts, _universe, _market, _disputeCrowdsourcer)
}

// LogDisputeCrowdsourcerCompleted is a paid mutator transaction binding the contract method 0x4c0019c3.
//
// Solidity: function logDisputeCrowdsourcerCompleted(_universe address, _market address, _disputeCrowdsourcer address) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerCompleted(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerCompleted(&_Augur.TransactOpts, _universe, _market, _disputeCrowdsourcer)
}

// LogDisputeCrowdsourcerContribution is a paid mutator transaction binding the contract method 0x5f872147.
//
// Solidity: function logDisputeCrowdsourcerContribution(_universe address, _reporter address, _market address, _disputeCrowdsourcer address, _amountStaked uint256, _description string) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerContribution(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _amountStaked *big.Int, _description string) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerContribution", _universe, _reporter, _market, _disputeCrowdsourcer, _amountStaked, _description)
}

// LogDisputeCrowdsourcerContribution is a paid mutator transaction binding the contract method 0x5f872147.
//
// Solidity: function logDisputeCrowdsourcerContribution(_universe address, _reporter address, _market address, _disputeCrowdsourcer address, _amountStaked uint256, _description string) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerContribution(_universe common.Address, _reporter common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _amountStaked *big.Int, _description string) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerContribution(&_Augur.TransactOpts, _universe, _reporter, _market, _disputeCrowdsourcer, _amountStaked, _description)
}

// LogDisputeCrowdsourcerContribution is a paid mutator transaction binding the contract method 0x5f872147.
//
// Solidity: function logDisputeCrowdsourcerContribution(_universe address, _reporter address, _market address, _disputeCrowdsourcer address, _amountStaked uint256, _description string) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerContribution(_universe common.Address, _reporter common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _amountStaked *big.Int, _description string) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerContribution(&_Augur.TransactOpts, _universe, _reporter, _market, _disputeCrowdsourcer, _amountStaked, _description)
}

// LogDisputeCrowdsourcerRedeemed is a paid mutator transaction binding the contract method 0x06ba8e42.
//
// Solidity: function logDisputeCrowdsourcerRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerRedeemed(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerRedeemed", _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogDisputeCrowdsourcerRedeemed is a paid mutator transaction binding the contract method 0x06ba8e42.
//
// Solidity: function logDisputeCrowdsourcerRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerRedeemed(&_Augur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogDisputeCrowdsourcerRedeemed is a paid mutator transaction binding the contract method 0x06ba8e42.
//
// Solidity: function logDisputeCrowdsourcerRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerRedeemed(&_Augur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogDisputeCrowdsourcerTokensBurned is a paid mutator transaction binding the contract method 0x10561361.
//
// Solidity: function logDisputeCrowdsourcerTokensBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerTokensBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerTokensBurned", _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensBurned is a paid mutator transaction binding the contract method 0x10561361.
//
// Solidity: function logDisputeCrowdsourcerTokensBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensBurned is a paid mutator transaction binding the contract method 0x10561361.
//
// Solidity: function logDisputeCrowdsourcerTokensBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensMinted is a paid mutator transaction binding the contract method 0xb14823e1.
//
// Solidity: function logDisputeCrowdsourcerTokensMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerTokensMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerTokensMinted", _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensMinted is a paid mutator transaction binding the contract method 0xb14823e1.
//
// Solidity: function logDisputeCrowdsourcerTokensMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensMinted is a paid mutator transaction binding the contract method 0xb14823e1.
//
// Solidity: function logDisputeCrowdsourcerTokensMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogDisputeCrowdsourcerTokensTransferred is a paid mutator transaction binding the contract method 0x50bd5cb9.
//
// Solidity: function logDisputeCrowdsourcerTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerTokensTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerTokensTransferred", _universe, _from, _to, _value)
}

// LogDisputeCrowdsourcerTokensTransferred is a paid mutator transaction binding the contract method 0x50bd5cb9.
//
// Solidity: function logDisputeCrowdsourcerTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogDisputeCrowdsourcerTokensTransferred is a paid mutator transaction binding the contract method 0x50bd5cb9.
//
// Solidity: function logDisputeCrowdsourcerTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogFeeTokenBurned is a paid mutator transaction binding the contract method 0x979141ea.
//
// Solidity: function logFeeTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogFeeTokenBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logFeeTokenBurned", _universe, _target, _amount)
}

// LogFeeTokenBurned is a paid mutator transaction binding the contract method 0x979141ea.
//
// Solidity: function logFeeTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogFeeTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeTokenBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogFeeTokenBurned is a paid mutator transaction binding the contract method 0x979141ea.
//
// Solidity: function logFeeTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogFeeTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeTokenBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogFeeTokenMinted is a paid mutator transaction binding the contract method 0x2698eec9.
//
// Solidity: function logFeeTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogFeeTokenMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logFeeTokenMinted", _universe, _target, _amount)
}

// LogFeeTokenMinted is a paid mutator transaction binding the contract method 0x2698eec9.
//
// Solidity: function logFeeTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogFeeTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeTokenMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogFeeTokenMinted is a paid mutator transaction binding the contract method 0x2698eec9.
//
// Solidity: function logFeeTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogFeeTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeTokenMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogFeeTokenTransferred is a paid mutator transaction binding the contract method 0x37227c07.
//
// Solidity: function logFeeTokenTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactor) LogFeeTokenTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logFeeTokenTransferred", _universe, _from, _to, _value)
}

// LogFeeTokenTransferred is a paid mutator transaction binding the contract method 0x37227c07.
//
// Solidity: function logFeeTokenTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurSession) LogFeeTokenTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeTokenTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogFeeTokenTransferred is a paid mutator transaction binding the contract method 0x37227c07.
//
// Solidity: function logFeeTokenTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogFeeTokenTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeTokenTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogFeeWindowBurned is a paid mutator transaction binding the contract method 0x542e9b18.
//
// Solidity: function logFeeWindowBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogFeeWindowBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logFeeWindowBurned", _universe, _target, _amount)
}

// LogFeeWindowBurned is a paid mutator transaction binding the contract method 0x542e9b18.
//
// Solidity: function logFeeWindowBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogFeeWindowBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogFeeWindowBurned is a paid mutator transaction binding the contract method 0x542e9b18.
//
// Solidity: function logFeeWindowBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogFeeWindowBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogFeeWindowCreated is a paid mutator transaction binding the contract method 0x8d1b2afd.
//
// Solidity: function logFeeWindowCreated(_feeWindow address, _id uint256) returns(bool)
func (_Augur *AugurTransactor) LogFeeWindowCreated(opts *bind.TransactOpts, _feeWindow common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logFeeWindowCreated", _feeWindow, _id)
}

// LogFeeWindowCreated is a paid mutator transaction binding the contract method 0x8d1b2afd.
//
// Solidity: function logFeeWindowCreated(_feeWindow address, _id uint256) returns(bool)
func (_Augur *AugurSession) LogFeeWindowCreated(_feeWindow common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowCreated(&_Augur.TransactOpts, _feeWindow, _id)
}

// LogFeeWindowCreated is a paid mutator transaction binding the contract method 0x8d1b2afd.
//
// Solidity: function logFeeWindowCreated(_feeWindow address, _id uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogFeeWindowCreated(_feeWindow common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowCreated(&_Augur.TransactOpts, _feeWindow, _id)
}

// LogFeeWindowMinted is a paid mutator transaction binding the contract method 0x60fe103e.
//
// Solidity: function logFeeWindowMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogFeeWindowMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logFeeWindowMinted", _universe, _target, _amount)
}

// LogFeeWindowMinted is a paid mutator transaction binding the contract method 0x60fe103e.
//
// Solidity: function logFeeWindowMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogFeeWindowMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogFeeWindowMinted is a paid mutator transaction binding the contract method 0x60fe103e.
//
// Solidity: function logFeeWindowMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogFeeWindowMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogFeeWindowRedeemed is a paid mutator transaction binding the contract method 0x3b186bfc.
//
// Solidity: function logFeeWindowRedeemed(_universe address, _reporter address, _amountRedeemed uint256, _reportingFeesReceived uint256) returns(bool)
func (_Augur *AugurTransactor) LogFeeWindowRedeemed(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _amountRedeemed *big.Int, _reportingFeesReceived *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logFeeWindowRedeemed", _universe, _reporter, _amountRedeemed, _reportingFeesReceived)
}

// LogFeeWindowRedeemed is a paid mutator transaction binding the contract method 0x3b186bfc.
//
// Solidity: function logFeeWindowRedeemed(_universe address, _reporter address, _amountRedeemed uint256, _reportingFeesReceived uint256) returns(bool)
func (_Augur *AugurSession) LogFeeWindowRedeemed(_universe common.Address, _reporter common.Address, _amountRedeemed *big.Int, _reportingFeesReceived *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowRedeemed(&_Augur.TransactOpts, _universe, _reporter, _amountRedeemed, _reportingFeesReceived)
}

// LogFeeWindowRedeemed is a paid mutator transaction binding the contract method 0x3b186bfc.
//
// Solidity: function logFeeWindowRedeemed(_universe address, _reporter address, _amountRedeemed uint256, _reportingFeesReceived uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogFeeWindowRedeemed(_universe common.Address, _reporter common.Address, _amountRedeemed *big.Int, _reportingFeesReceived *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowRedeemed(&_Augur.TransactOpts, _universe, _reporter, _amountRedeemed, _reportingFeesReceived)
}

// LogFeeWindowTransferred is a paid mutator transaction binding the contract method 0x788873ea.
//
// Solidity: function logFeeWindowTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactor) LogFeeWindowTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logFeeWindowTransferred", _universe, _from, _to, _value)
}

// LogFeeWindowTransferred is a paid mutator transaction binding the contract method 0x788873ea.
//
// Solidity: function logFeeWindowTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurSession) LogFeeWindowTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogFeeWindowTransferred is a paid mutator transaction binding the contract method 0x788873ea.
//
// Solidity: function logFeeWindowTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogFeeWindowTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogFeeWindowTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogInitialReportSubmitted is a paid mutator transaction binding the contract method 0x9a0d59aa.
//
// Solidity: function logInitialReportSubmitted(_universe address, _reporter address, _market address, _amountStaked uint256, _isDesignatedReporter bool, _payoutNumerators uint256[], _invalid bool, _description string) returns(bool)
func (_Augur *AugurTransactor) LogInitialReportSubmitted(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _amountStaked *big.Int, _isDesignatedReporter bool, _payoutNumerators []*big.Int, _invalid bool, _description string) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logInitialReportSubmitted", _universe, _reporter, _market, _amountStaked, _isDesignatedReporter, _payoutNumerators, _invalid, _description)
}

// LogInitialReportSubmitted is a paid mutator transaction binding the contract method 0x9a0d59aa.
//
// Solidity: function logInitialReportSubmitted(_universe address, _reporter address, _market address, _amountStaked uint256, _isDesignatedReporter bool, _payoutNumerators uint256[], _invalid bool, _description string) returns(bool)
func (_Augur *AugurSession) LogInitialReportSubmitted(_universe common.Address, _reporter common.Address, _market common.Address, _amountStaked *big.Int, _isDesignatedReporter bool, _payoutNumerators []*big.Int, _invalid bool, _description string) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReportSubmitted(&_Augur.TransactOpts, _universe, _reporter, _market, _amountStaked, _isDesignatedReporter, _payoutNumerators, _invalid, _description)
}

// LogInitialReportSubmitted is a paid mutator transaction binding the contract method 0x9a0d59aa.
//
// Solidity: function logInitialReportSubmitted(_universe address, _reporter address, _market address, _amountStaked uint256, _isDesignatedReporter bool, _payoutNumerators uint256[], _invalid bool, _description string) returns(bool)
func (_Augur *AugurTransactorSession) LogInitialReportSubmitted(_universe common.Address, _reporter common.Address, _market common.Address, _amountStaked *big.Int, _isDesignatedReporter bool, _payoutNumerators []*big.Int, _invalid bool, _description string) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReportSubmitted(&_Augur.TransactOpts, _universe, _reporter, _market, _amountStaked, _isDesignatedReporter, _payoutNumerators, _invalid, _description)
}

// LogInitialReporterRedeemed is a paid mutator transaction binding the contract method 0xec18e2f1.
//
// Solidity: function logInitialReporterRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_Augur *AugurTransactor) LogInitialReporterRedeemed(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logInitialReporterRedeemed", _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogInitialReporterRedeemed is a paid mutator transaction binding the contract method 0xec18e2f1.
//
// Solidity: function logInitialReporterRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_Augur *AugurSession) LogInitialReporterRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReporterRedeemed(&_Augur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogInitialReporterRedeemed is a paid mutator transaction binding the contract method 0xec18e2f1.
//
// Solidity: function logInitialReporterRedeemed(_universe address, _reporter address, _market address, _amountRedeemed uint256, _repReceived uint256, _reportingFeesReceived uint256, _payoutNumerators uint256[]) returns(bool)
func (_Augur *AugurTransactorSession) LogInitialReporterRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _reportingFeesReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReporterRedeemed(&_Augur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _reportingFeesReceived, _payoutNumerators)
}

// LogInitialReporterTransferred is a paid mutator transaction binding the contract method 0xe3142e90.
//
// Solidity: function logInitialReporterTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_Augur *AugurTransactor) LogInitialReporterTransferred(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logInitialReporterTransferred", _universe, _market, _from, _to)
}

// LogInitialReporterTransferred is a paid mutator transaction binding the contract method 0xe3142e90.
//
// Solidity: function logInitialReporterTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_Augur *AugurSession) LogInitialReporterTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReporterTransferred(&_Augur.TransactOpts, _universe, _market, _from, _to)
}

// LogInitialReporterTransferred is a paid mutator transaction binding the contract method 0xe3142e90.
//
// Solidity: function logInitialReporterTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_Augur *AugurTransactorSession) LogInitialReporterTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReporterTransferred(&_Augur.TransactOpts, _universe, _market, _from, _to)
}

// LogMarketCreated is a paid mutator transaction binding the contract method 0xbc339f41.
//
// Solidity: function logMarketCreated(_topic bytes32, _description string, _extraInfo string, _universe address, _market address, _marketCreator address, _minPrice int256, _maxPrice int256, _marketType uint8) returns(bool)
func (_Augur *AugurTransactor) LogMarketCreated(opts *bind.TransactOpts, _topic [32]byte, _description string, _extraInfo string, _universe common.Address, _market common.Address, _marketCreator common.Address, _minPrice *big.Int, _maxPrice *big.Int, _marketType uint8) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketCreated", _topic, _description, _extraInfo, _universe, _market, _marketCreator, _minPrice, _maxPrice, _marketType)
}

// LogMarketCreated is a paid mutator transaction binding the contract method 0xbc339f41.
//
// Solidity: function logMarketCreated(_topic bytes32, _description string, _extraInfo string, _universe address, _market address, _marketCreator address, _minPrice int256, _maxPrice int256, _marketType uint8) returns(bool)
func (_Augur *AugurSession) LogMarketCreated(_topic [32]byte, _description string, _extraInfo string, _universe common.Address, _market common.Address, _marketCreator common.Address, _minPrice *big.Int, _maxPrice *big.Int, _marketType uint8) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketCreated(&_Augur.TransactOpts, _topic, _description, _extraInfo, _universe, _market, _marketCreator, _minPrice, _maxPrice, _marketType)
}

// LogMarketCreated is a paid mutator transaction binding the contract method 0xbc339f41.
//
// Solidity: function logMarketCreated(_topic bytes32, _description string, _extraInfo string, _universe address, _market address, _marketCreator address, _minPrice int256, _maxPrice int256, _marketType uint8) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketCreated(_topic [32]byte, _description string, _extraInfo string, _universe common.Address, _market common.Address, _marketCreator common.Address, _minPrice *big.Int, _maxPrice *big.Int, _marketType uint8) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketCreated(&_Augur.TransactOpts, _topic, _description, _extraInfo, _universe, _market, _marketCreator, _minPrice, _maxPrice, _marketType)
}

// LogMarketFinalized is a paid mutator transaction binding the contract method 0x27816ffc.
//
// Solidity: function logMarketFinalized(_universe address) returns(bool)
func (_Augur *AugurTransactor) LogMarketFinalized(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketFinalized", _universe)
}

// LogMarketFinalized is a paid mutator transaction binding the contract method 0x27816ffc.
//
// Solidity: function logMarketFinalized(_universe address) returns(bool)
func (_Augur *AugurSession) LogMarketFinalized(_universe common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketFinalized(&_Augur.TransactOpts, _universe)
}

// LogMarketFinalized is a paid mutator transaction binding the contract method 0x27816ffc.
//
// Solidity: function logMarketFinalized(_universe address) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketFinalized(_universe common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketFinalized(&_Augur.TransactOpts, _universe)
}

// LogMarketMailboxTransferred is a paid mutator transaction binding the contract method 0x339594f9.
//
// Solidity: function logMarketMailboxTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_Augur *AugurTransactor) LogMarketMailboxTransferred(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketMailboxTransferred", _universe, _market, _from, _to)
}

// LogMarketMailboxTransferred is a paid mutator transaction binding the contract method 0x339594f9.
//
// Solidity: function logMarketMailboxTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_Augur *AugurSession) LogMarketMailboxTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketMailboxTransferred(&_Augur.TransactOpts, _universe, _market, _from, _to)
}

// LogMarketMailboxTransferred is a paid mutator transaction binding the contract method 0x339594f9.
//
// Solidity: function logMarketMailboxTransferred(_universe address, _market address, _from address, _to address) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketMailboxTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketMailboxTransferred(&_Augur.TransactOpts, _universe, _market, _from, _to)
}

// LogMarketMigrated is a paid mutator transaction binding the contract method 0x17674e4d.
//
// Solidity: function logMarketMigrated(_market address, _originalUniverse address) returns(bool)
func (_Augur *AugurTransactor) LogMarketMigrated(opts *bind.TransactOpts, _market common.Address, _originalUniverse common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketMigrated", _market, _originalUniverse)
}

// LogMarketMigrated is a paid mutator transaction binding the contract method 0x17674e4d.
//
// Solidity: function logMarketMigrated(_market address, _originalUniverse address) returns(bool)
func (_Augur *AugurSession) LogMarketMigrated(_market common.Address, _originalUniverse common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketMigrated(&_Augur.TransactOpts, _market, _originalUniverse)
}

// LogMarketMigrated is a paid mutator transaction binding the contract method 0x17674e4d.
//
// Solidity: function logMarketMigrated(_market address, _originalUniverse address) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketMigrated(_market common.Address, _originalUniverse common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketMigrated(&_Augur.TransactOpts, _market, _originalUniverse)
}

// LogMarketParticipantsDisavowed is a paid mutator transaction binding the contract method 0xc67af5cc.
//
// Solidity: function logMarketParticipantsDisavowed(_universe address) returns(bool)
func (_Augur *AugurTransactor) LogMarketParticipantsDisavowed(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketParticipantsDisavowed", _universe)
}

// LogMarketParticipantsDisavowed is a paid mutator transaction binding the contract method 0xc67af5cc.
//
// Solidity: function logMarketParticipantsDisavowed(_universe address) returns(bool)
func (_Augur *AugurSession) LogMarketParticipantsDisavowed(_universe common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketParticipantsDisavowed(&_Augur.TransactOpts, _universe)
}

// LogMarketParticipantsDisavowed is a paid mutator transaction binding the contract method 0xc67af5cc.
//
// Solidity: function logMarketParticipantsDisavowed(_universe address) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketParticipantsDisavowed(_universe common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketParticipantsDisavowed(&_Augur.TransactOpts, _universe)
}

// LogMarketTransferred is a paid mutator transaction binding the contract method 0x23290737.
//
// Solidity: function logMarketTransferred(_universe address, _from address, _to address) returns(bool)
func (_Augur *AugurTransactor) LogMarketTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketTransferred", _universe, _from, _to)
}

// LogMarketTransferred is a paid mutator transaction binding the contract method 0x23290737.
//
// Solidity: function logMarketTransferred(_universe address, _from address, _to address) returns(bool)
func (_Augur *AugurSession) LogMarketTransferred(_universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketTransferred(&_Augur.TransactOpts, _universe, _from, _to)
}

// LogMarketTransferred is a paid mutator transaction binding the contract method 0x23290737.
//
// Solidity: function logMarketTransferred(_universe address, _from address, _to address) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketTransferred(_universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketTransferred(&_Augur.TransactOpts, _universe, _from, _to)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0x6e1636bb.
//
// Solidity: function logOrderCanceled(_universe address, _shareToken address, _sender address, _orderId bytes32, _orderType uint8, _tokenRefund uint256, _sharesRefund uint256) returns(bool)
func (_Augur *AugurTransactor) LogOrderCanceled(opts *bind.TransactOpts, _universe common.Address, _shareToken common.Address, _sender common.Address, _orderId [32]byte, _orderType uint8, _tokenRefund *big.Int, _sharesRefund *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logOrderCanceled", _universe, _shareToken, _sender, _orderId, _orderType, _tokenRefund, _sharesRefund)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0x6e1636bb.
//
// Solidity: function logOrderCanceled(_universe address, _shareToken address, _sender address, _orderId bytes32, _orderType uint8, _tokenRefund uint256, _sharesRefund uint256) returns(bool)
func (_Augur *AugurSession) LogOrderCanceled(_universe common.Address, _shareToken common.Address, _sender common.Address, _orderId [32]byte, _orderType uint8, _tokenRefund *big.Int, _sharesRefund *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogOrderCanceled(&_Augur.TransactOpts, _universe, _shareToken, _sender, _orderId, _orderType, _tokenRefund, _sharesRefund)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0x6e1636bb.
//
// Solidity: function logOrderCanceled(_universe address, _shareToken address, _sender address, _orderId bytes32, _orderType uint8, _tokenRefund uint256, _sharesRefund uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogOrderCanceled(_universe common.Address, _shareToken common.Address, _sender common.Address, _orderId [32]byte, _orderType uint8, _tokenRefund *big.Int, _sharesRefund *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogOrderCanceled(&_Augur.TransactOpts, _universe, _shareToken, _sender, _orderId, _orderType, _tokenRefund, _sharesRefund)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x0ae41574.
//
// Solidity: function logOrderCreated(_orderType uint8, _amount uint256, _price uint256, _creator address, _moneyEscrowed uint256, _sharesEscrowed uint256, _tradeGroupId bytes32, _orderId bytes32, _universe address, _shareToken address) returns(bool)
func (_Augur *AugurTransactor) LogOrderCreated(opts *bind.TransactOpts, _orderType uint8, _amount *big.Int, _price *big.Int, _creator common.Address, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _tradeGroupId [32]byte, _orderId [32]byte, _universe common.Address, _shareToken common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logOrderCreated", _orderType, _amount, _price, _creator, _moneyEscrowed, _sharesEscrowed, _tradeGroupId, _orderId, _universe, _shareToken)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x0ae41574.
//
// Solidity: function logOrderCreated(_orderType uint8, _amount uint256, _price uint256, _creator address, _moneyEscrowed uint256, _sharesEscrowed uint256, _tradeGroupId bytes32, _orderId bytes32, _universe address, _shareToken address) returns(bool)
func (_Augur *AugurSession) LogOrderCreated(_orderType uint8, _amount *big.Int, _price *big.Int, _creator common.Address, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _tradeGroupId [32]byte, _orderId [32]byte, _universe common.Address, _shareToken common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogOrderCreated(&_Augur.TransactOpts, _orderType, _amount, _price, _creator, _moneyEscrowed, _sharesEscrowed, _tradeGroupId, _orderId, _universe, _shareToken)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x0ae41574.
//
// Solidity: function logOrderCreated(_orderType uint8, _amount uint256, _price uint256, _creator address, _moneyEscrowed uint256, _sharesEscrowed uint256, _tradeGroupId bytes32, _orderId bytes32, _universe address, _shareToken address) returns(bool)
func (_Augur *AugurTransactorSession) LogOrderCreated(_orderType uint8, _amount *big.Int, _price *big.Int, _creator common.Address, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _tradeGroupId [32]byte, _orderId [32]byte, _universe common.Address, _shareToken common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogOrderCreated(&_Augur.TransactOpts, _orderType, _amount, _price, _creator, _moneyEscrowed, _sharesEscrowed, _tradeGroupId, _orderId, _universe, _shareToken)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x80d5398e.
//
// Solidity: function logOrderFilled(_universe address, _shareToken address, _filler address, _orderId bytes32, _numCreatorShares uint256, _numCreatorTokens uint256, _numFillerShares uint256, _numFillerTokens uint256, _marketCreatorFees uint256, _reporterFees uint256, _amountFilled uint256, _tradeGroupId bytes32) returns(bool)
func (_Augur *AugurTransactor) LogOrderFilled(opts *bind.TransactOpts, _universe common.Address, _shareToken common.Address, _filler common.Address, _orderId [32]byte, _numCreatorShares *big.Int, _numCreatorTokens *big.Int, _numFillerShares *big.Int, _numFillerTokens *big.Int, _marketCreatorFees *big.Int, _reporterFees *big.Int, _amountFilled *big.Int, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logOrderFilled", _universe, _shareToken, _filler, _orderId, _numCreatorShares, _numCreatorTokens, _numFillerShares, _numFillerTokens, _marketCreatorFees, _reporterFees, _amountFilled, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x80d5398e.
//
// Solidity: function logOrderFilled(_universe address, _shareToken address, _filler address, _orderId bytes32, _numCreatorShares uint256, _numCreatorTokens uint256, _numFillerShares uint256, _numFillerTokens uint256, _marketCreatorFees uint256, _reporterFees uint256, _amountFilled uint256, _tradeGroupId bytes32) returns(bool)
func (_Augur *AugurSession) LogOrderFilled(_universe common.Address, _shareToken common.Address, _filler common.Address, _orderId [32]byte, _numCreatorShares *big.Int, _numCreatorTokens *big.Int, _numFillerShares *big.Int, _numFillerTokens *big.Int, _marketCreatorFees *big.Int, _reporterFees *big.Int, _amountFilled *big.Int, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Augur.Contract.LogOrderFilled(&_Augur.TransactOpts, _universe, _shareToken, _filler, _orderId, _numCreatorShares, _numCreatorTokens, _numFillerShares, _numFillerTokens, _marketCreatorFees, _reporterFees, _amountFilled, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x80d5398e.
//
// Solidity: function logOrderFilled(_universe address, _shareToken address, _filler address, _orderId bytes32, _numCreatorShares uint256, _numCreatorTokens uint256, _numFillerShares uint256, _numFillerTokens uint256, _marketCreatorFees uint256, _reporterFees uint256, _amountFilled uint256, _tradeGroupId bytes32) returns(bool)
func (_Augur *AugurTransactorSession) LogOrderFilled(_universe common.Address, _shareToken common.Address, _filler common.Address, _orderId [32]byte, _numCreatorShares *big.Int, _numCreatorTokens *big.Int, _numFillerShares *big.Int, _numFillerTokens *big.Int, _marketCreatorFees *big.Int, _reporterFees *big.Int, _amountFilled *big.Int, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Augur.Contract.LogOrderFilled(&_Augur.TransactOpts, _universe, _shareToken, _filler, _orderId, _numCreatorShares, _numCreatorTokens, _numFillerShares, _numFillerTokens, _marketCreatorFees, _reporterFees, _amountFilled, _tradeGroupId)
}

// LogReportingParticipantDisavowed is a paid mutator transaction binding the contract method 0x17570e80.
//
// Solidity: function logReportingParticipantDisavowed(_universe address, _market address) returns(bool)
func (_Augur *AugurTransactor) LogReportingParticipantDisavowed(opts *bind.TransactOpts, _universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logReportingParticipantDisavowed", _universe, _market)
}

// LogReportingParticipantDisavowed is a paid mutator transaction binding the contract method 0x17570e80.
//
// Solidity: function logReportingParticipantDisavowed(_universe address, _market address) returns(bool)
func (_Augur *AugurSession) LogReportingParticipantDisavowed(_universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogReportingParticipantDisavowed(&_Augur.TransactOpts, _universe, _market)
}

// LogReportingParticipantDisavowed is a paid mutator transaction binding the contract method 0x17570e80.
//
// Solidity: function logReportingParticipantDisavowed(_universe address, _market address) returns(bool)
func (_Augur *AugurTransactorSession) LogReportingParticipantDisavowed(_universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogReportingParticipantDisavowed(&_Augur.TransactOpts, _universe, _market)
}

// LogReputationTokenBurned is a paid mutator transaction binding the contract method 0x4405a339.
//
// Solidity: function logReputationTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogReputationTokenBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logReputationTokenBurned", _universe, _target, _amount)
}

// LogReputationTokenBurned is a paid mutator transaction binding the contract method 0x4405a339.
//
// Solidity: function logReputationTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogReputationTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokenBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogReputationTokenBurned is a paid mutator transaction binding the contract method 0x4405a339.
//
// Solidity: function logReputationTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogReputationTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokenBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogReputationTokenMinted is a paid mutator transaction binding the contract method 0x79fff7a9.
//
// Solidity: function logReputationTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogReputationTokenMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logReputationTokenMinted", _universe, _target, _amount)
}

// LogReputationTokenMinted is a paid mutator transaction binding the contract method 0x79fff7a9.
//
// Solidity: function logReputationTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogReputationTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokenMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogReputationTokenMinted is a paid mutator transaction binding the contract method 0x79fff7a9.
//
// Solidity: function logReputationTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogReputationTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokenMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogReputationTokensTransferred is a paid mutator transaction binding the contract method 0xec37a6e4.
//
// Solidity: function logReputationTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactor) LogReputationTokensTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logReputationTokensTransferred", _universe, _from, _to, _value)
}

// LogReputationTokensTransferred is a paid mutator transaction binding the contract method 0xec37a6e4.
//
// Solidity: function logReputationTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurSession) LogReputationTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogReputationTokensTransferred is a paid mutator transaction binding the contract method 0xec37a6e4.
//
// Solidity: function logReputationTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogReputationTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogShareTokenBurned is a paid mutator transaction binding the contract method 0xa1b7887f.
//
// Solidity: function logShareTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogShareTokenBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logShareTokenBurned", _universe, _target, _amount)
}

// LogShareTokenBurned is a paid mutator transaction binding the contract method 0xa1b7887f.
//
// Solidity: function logShareTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogShareTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogShareTokenBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogShareTokenBurned is a paid mutator transaction binding the contract method 0xa1b7887f.
//
// Solidity: function logShareTokenBurned(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogShareTokenBurned(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogShareTokenBurned(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogShareTokenMinted is a paid mutator transaction binding the contract method 0xa1dfe545.
//
// Solidity: function logShareTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) LogShareTokenMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logShareTokenMinted", _universe, _target, _amount)
}

// LogShareTokenMinted is a paid mutator transaction binding the contract method 0xa1dfe545.
//
// Solidity: function logShareTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurSession) LogShareTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogShareTokenMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogShareTokenMinted is a paid mutator transaction binding the contract method 0xa1dfe545.
//
// Solidity: function logShareTokenMinted(_universe address, _target address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogShareTokenMinted(_universe common.Address, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogShareTokenMinted(&_Augur.TransactOpts, _universe, _target, _amount)
}

// LogShareTokensTransferred is a paid mutator transaction binding the contract method 0x86b9a1f4.
//
// Solidity: function logShareTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactor) LogShareTokensTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logShareTokensTransferred", _universe, _from, _to, _value)
}

// LogShareTokensTransferred is a paid mutator transaction binding the contract method 0x86b9a1f4.
//
// Solidity: function logShareTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurSession) LogShareTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogShareTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogShareTokensTransferred is a paid mutator transaction binding the contract method 0x86b9a1f4.
//
// Solidity: function logShareTokensTransferred(_universe address, _from address, _to address, _value uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogShareTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogShareTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value)
}

// LogTimestampSet is a paid mutator transaction binding the contract method 0xc8e6b2a8.
//
// Solidity: function logTimestampSet(_newTimestamp uint256) returns(bool)
func (_Augur *AugurTransactor) LogTimestampSet(opts *bind.TransactOpts, _newTimestamp *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logTimestampSet", _newTimestamp)
}

// LogTimestampSet is a paid mutator transaction binding the contract method 0xc8e6b2a8.
//
// Solidity: function logTimestampSet(_newTimestamp uint256) returns(bool)
func (_Augur *AugurSession) LogTimestampSet(_newTimestamp *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogTimestampSet(&_Augur.TransactOpts, _newTimestamp)
}

// LogTimestampSet is a paid mutator transaction binding the contract method 0xc8e6b2a8.
//
// Solidity: function logTimestampSet(_newTimestamp uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogTimestampSet(_newTimestamp *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogTimestampSet(&_Augur.TransactOpts, _newTimestamp)
}

// LogTradingProceedsClaimed is a paid mutator transaction binding the contract method 0x6051fa2c.
//
// Solidity: function logTradingProceedsClaimed(_universe address, _shareToken address, _sender address, _market address, _numShares uint256, _numPayoutTokens uint256, _finalTokenBalance uint256) returns(bool)
func (_Augur *AugurTransactor) LogTradingProceedsClaimed(opts *bind.TransactOpts, _universe common.Address, _shareToken common.Address, _sender common.Address, _market common.Address, _numShares *big.Int, _numPayoutTokens *big.Int, _finalTokenBalance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logTradingProceedsClaimed", _universe, _shareToken, _sender, _market, _numShares, _numPayoutTokens, _finalTokenBalance)
}

// LogTradingProceedsClaimed is a paid mutator transaction binding the contract method 0x6051fa2c.
//
// Solidity: function logTradingProceedsClaimed(_universe address, _shareToken address, _sender address, _market address, _numShares uint256, _numPayoutTokens uint256, _finalTokenBalance uint256) returns(bool)
func (_Augur *AugurSession) LogTradingProceedsClaimed(_universe common.Address, _shareToken common.Address, _sender common.Address, _market common.Address, _numShares *big.Int, _numPayoutTokens *big.Int, _finalTokenBalance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogTradingProceedsClaimed(&_Augur.TransactOpts, _universe, _shareToken, _sender, _market, _numShares, _numPayoutTokens, _finalTokenBalance)
}

// LogTradingProceedsClaimed is a paid mutator transaction binding the contract method 0x6051fa2c.
//
// Solidity: function logTradingProceedsClaimed(_universe address, _shareToken address, _sender address, _market address, _numShares uint256, _numPayoutTokens uint256, _finalTokenBalance uint256) returns(bool)
func (_Augur *AugurTransactorSession) LogTradingProceedsClaimed(_universe common.Address, _shareToken common.Address, _sender common.Address, _market common.Address, _numShares *big.Int, _numPayoutTokens *big.Int, _finalTokenBalance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogTradingProceedsClaimed(&_Augur.TransactOpts, _universe, _shareToken, _sender, _market, _numShares, _numPayoutTokens, _finalTokenBalance)
}

// LogUniverseForked is a paid mutator transaction binding the contract method 0x4a812023.
//
// Solidity: function logUniverseForked() returns(bool)
func (_Augur *AugurTransactor) LogUniverseForked(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logUniverseForked")
}

// LogUniverseForked is a paid mutator transaction binding the contract method 0x4a812023.
//
// Solidity: function logUniverseForked() returns(bool)
func (_Augur *AugurSession) LogUniverseForked() (*types.Transaction, error) {
	return _Augur.Contract.LogUniverseForked(&_Augur.TransactOpts)
}

// LogUniverseForked is a paid mutator transaction binding the contract method 0x4a812023.
//
// Solidity: function logUniverseForked() returns(bool)
func (_Augur *AugurTransactorSession) LogUniverseForked() (*types.Transaction, error) {
	return _Augur.Contract.LogUniverseForked(&_Augur.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Augur *AugurTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Augur *AugurSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Augur.Contract.SetController(&_Augur.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Augur *AugurTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Augur.Contract.SetController(&_Augur.TransactOpts, _controller)
}

// TrustedTransfer is a paid mutator transaction binding the contract method 0xec238994.
//
// Solidity: function trustedTransfer(_token address, _from address, _to address, _amount uint256) returns(bool)
func (_Augur *AugurTransactor) TrustedTransfer(opts *bind.TransactOpts, _token common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "trustedTransfer", _token, _from, _to, _amount)
}

// TrustedTransfer is a paid mutator transaction binding the contract method 0xec238994.
//
// Solidity: function trustedTransfer(_token address, _from address, _to address, _amount uint256) returns(bool)
func (_Augur *AugurSession) TrustedTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.TrustedTransfer(&_Augur.TransactOpts, _token, _from, _to, _amount)
}

// TrustedTransfer is a paid mutator transaction binding the contract method 0xec238994.
//
// Solidity: function trustedTransfer(_token address, _from address, _to address, _amount uint256) returns(bool)
func (_Augur *AugurTransactorSession) TrustedTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.TrustedTransfer(&_Augur.TransactOpts, _token, _from, _to, _amount)
}

// AugurCompleteSetsPurchasedIterator is returned from FilterCompleteSetsPurchased and is used to iterate over the raw logs and unpacked data for CompleteSetsPurchased events raised by the Augur contract.
type AugurCompleteSetsPurchasedIterator struct {
	Event *AugurCompleteSetsPurchased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurCompleteSetsPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurCompleteSetsPurchased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurCompleteSetsPurchased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurCompleteSetsPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurCompleteSetsPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurCompleteSetsPurchased represents a CompleteSetsPurchased event raised by the Augur contract.
type AugurCompleteSetsPurchased struct {
	Universe        common.Address
	Market          common.Address
	Account         common.Address
	NumCompleteSets *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCompleteSetsPurchased is a free log retrieval operation binding the contract event 0x349ab20f76ba930a00da1936627d07400af6bb7cd2e2b4c68bcab93ca8aff418.
//
// Solidity: e CompleteSetsPurchased(universe indexed address, market indexed address, account indexed address, numCompleteSets uint256)
func (_Augur *AugurFilterer) FilterCompleteSetsPurchased(opts *bind.FilterOpts, universe []common.Address, market []common.Address, account []common.Address) (*AugurCompleteSetsPurchasedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "CompleteSetsPurchased", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AugurCompleteSetsPurchasedIterator{contract: _Augur.contract, event: "CompleteSetsPurchased", logs: logs, sub: sub}, nil
}

// WatchCompleteSetsPurchased is a free log subscription operation binding the contract event 0x349ab20f76ba930a00da1936627d07400af6bb7cd2e2b4c68bcab93ca8aff418.
//
// Solidity: e CompleteSetsPurchased(universe indexed address, market indexed address, account indexed address, numCompleteSets uint256)
func (_Augur *AugurFilterer) WatchCompleteSetsPurchased(opts *bind.WatchOpts, sink chan<- *AugurCompleteSetsPurchased, universe []common.Address, market []common.Address, account []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "CompleteSetsPurchased", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurCompleteSetsPurchased)
				if err := _Augur.contract.UnpackLog(event, "CompleteSetsPurchased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurCompleteSetsSoldIterator is returned from FilterCompleteSetsSold and is used to iterate over the raw logs and unpacked data for CompleteSetsSold events raised by the Augur contract.
type AugurCompleteSetsSoldIterator struct {
	Event *AugurCompleteSetsSold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurCompleteSetsSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurCompleteSetsSold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurCompleteSetsSold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurCompleteSetsSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurCompleteSetsSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurCompleteSetsSold represents a CompleteSetsSold event raised by the Augur contract.
type AugurCompleteSetsSold struct {
	Universe        common.Address
	Market          common.Address
	Account         common.Address
	NumCompleteSets *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCompleteSetsSold is a free log retrieval operation binding the contract event 0x68166bb2a567c21899b00209f52c286bf00ac613acc9f183da791ac5f5f47051.
//
// Solidity: e CompleteSetsSold(universe indexed address, market indexed address, account indexed address, numCompleteSets uint256)
func (_Augur *AugurFilterer) FilterCompleteSetsSold(opts *bind.FilterOpts, universe []common.Address, market []common.Address, account []common.Address) (*AugurCompleteSetsSoldIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "CompleteSetsSold", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AugurCompleteSetsSoldIterator{contract: _Augur.contract, event: "CompleteSetsSold", logs: logs, sub: sub}, nil
}

// WatchCompleteSetsSold is a free log subscription operation binding the contract event 0x68166bb2a567c21899b00209f52c286bf00ac613acc9f183da791ac5f5f47051.
//
// Solidity: e CompleteSetsSold(universe indexed address, market indexed address, account indexed address, numCompleteSets uint256)
func (_Augur *AugurFilterer) WatchCompleteSetsSold(opts *bind.WatchOpts, sink chan<- *AugurCompleteSetsSold, universe []common.Address, market []common.Address, account []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "CompleteSetsSold", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurCompleteSetsSold)
				if err := _Augur.contract.UnpackLog(event, "CompleteSetsSold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurDisputeCrowdsourcerCompletedIterator is returned from FilterDisputeCrowdsourcerCompleted and is used to iterate over the raw logs and unpacked data for DisputeCrowdsourcerCompleted events raised by the Augur contract.
type AugurDisputeCrowdsourcerCompletedIterator struct {
	Event *AugurDisputeCrowdsourcerCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDisputeCrowdsourcerCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDisputeCrowdsourcerCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDisputeCrowdsourcerCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDisputeCrowdsourcerCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDisputeCrowdsourcerCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDisputeCrowdsourcerCompleted represents a DisputeCrowdsourcerCompleted event raised by the Augur contract.
type AugurDisputeCrowdsourcerCompleted struct {
	Universe            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDisputeCrowdsourcerCompleted is a free log retrieval operation binding the contract event 0xec05f094139821aeb3220a0837f5d14eb02aa619179aadf3b316ed95b3648abb.
//
// Solidity: e DisputeCrowdsourcerCompleted(universe indexed address, market indexed address, disputeCrowdsourcer address)
func (_Augur *AugurFilterer) FilterDisputeCrowdsourcerCompleted(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurDisputeCrowdsourcerCompletedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DisputeCrowdsourcerCompleted", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurDisputeCrowdsourcerCompletedIterator{contract: _Augur.contract, event: "DisputeCrowdsourcerCompleted", logs: logs, sub: sub}, nil
}

// WatchDisputeCrowdsourcerCompleted is a free log subscription operation binding the contract event 0xec05f094139821aeb3220a0837f5d14eb02aa619179aadf3b316ed95b3648abb.
//
// Solidity: e DisputeCrowdsourcerCompleted(universe indexed address, market indexed address, disputeCrowdsourcer address)
func (_Augur *AugurFilterer) WatchDisputeCrowdsourcerCompleted(opts *bind.WatchOpts, sink chan<- *AugurDisputeCrowdsourcerCompleted, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DisputeCrowdsourcerCompleted", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDisputeCrowdsourcerCompleted)
				if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurDisputeCrowdsourcerContributionIterator is returned from FilterDisputeCrowdsourcerContribution and is used to iterate over the raw logs and unpacked data for DisputeCrowdsourcerContribution events raised by the Augur contract.
type AugurDisputeCrowdsourcerContributionIterator struct {
	Event *AugurDisputeCrowdsourcerContribution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDisputeCrowdsourcerContributionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDisputeCrowdsourcerContribution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDisputeCrowdsourcerContribution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDisputeCrowdsourcerContributionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDisputeCrowdsourcerContributionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDisputeCrowdsourcerContribution represents a DisputeCrowdsourcerContribution event raised by the Augur contract.
type AugurDisputeCrowdsourcerContribution struct {
	Universe            common.Address
	Reporter            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	AmountStaked        *big.Int
	Description         string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDisputeCrowdsourcerContribution is a free log retrieval operation binding the contract event 0x1ba97c2894f2b4eb21d849bdb2c4b2007b3562407a13d5581e8cc603ccfc70aa.
//
// Solidity: e DisputeCrowdsourcerContribution(universe indexed address, reporter indexed address, market indexed address, disputeCrowdsourcer address, amountStaked uint256, description string)
func (_Augur *AugurFilterer) FilterDisputeCrowdsourcerContribution(opts *bind.FilterOpts, universe []common.Address, reporter []common.Address, market []common.Address) (*AugurDisputeCrowdsourcerContributionIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DisputeCrowdsourcerContribution", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurDisputeCrowdsourcerContributionIterator{contract: _Augur.contract, event: "DisputeCrowdsourcerContribution", logs: logs, sub: sub}, nil
}

// WatchDisputeCrowdsourcerContribution is a free log subscription operation binding the contract event 0x1ba97c2894f2b4eb21d849bdb2c4b2007b3562407a13d5581e8cc603ccfc70aa.
//
// Solidity: e DisputeCrowdsourcerContribution(universe indexed address, reporter indexed address, market indexed address, disputeCrowdsourcer address, amountStaked uint256, description string)
func (_Augur *AugurFilterer) WatchDisputeCrowdsourcerContribution(opts *bind.WatchOpts, sink chan<- *AugurDisputeCrowdsourcerContribution, universe []common.Address, reporter []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DisputeCrowdsourcerContribution", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDisputeCrowdsourcerContribution)
				if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerContribution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurDisputeCrowdsourcerCreatedIterator is returned from FilterDisputeCrowdsourcerCreated and is used to iterate over the raw logs and unpacked data for DisputeCrowdsourcerCreated events raised by the Augur contract.
type AugurDisputeCrowdsourcerCreatedIterator struct {
	Event *AugurDisputeCrowdsourcerCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDisputeCrowdsourcerCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDisputeCrowdsourcerCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDisputeCrowdsourcerCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDisputeCrowdsourcerCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDisputeCrowdsourcerCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDisputeCrowdsourcerCreated represents a DisputeCrowdsourcerCreated event raised by the Augur contract.
type AugurDisputeCrowdsourcerCreated struct {
	Universe            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	PayoutNumerators    []*big.Int
	Size                *big.Int
	Invalid             bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDisputeCrowdsourcerCreated is a free log retrieval operation binding the contract event 0xccc07058358a9411a6acb3cd58bf6d0b398c3ff1f0b2c8e97a6dbdbbe74eae41.
//
// Solidity: e DisputeCrowdsourcerCreated(universe indexed address, market indexed address, disputeCrowdsourcer address, payoutNumerators uint256[], size uint256, invalid bool)
func (_Augur *AugurFilterer) FilterDisputeCrowdsourcerCreated(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurDisputeCrowdsourcerCreatedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DisputeCrowdsourcerCreated", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurDisputeCrowdsourcerCreatedIterator{contract: _Augur.contract, event: "DisputeCrowdsourcerCreated", logs: logs, sub: sub}, nil
}

// WatchDisputeCrowdsourcerCreated is a free log subscription operation binding the contract event 0xccc07058358a9411a6acb3cd58bf6d0b398c3ff1f0b2c8e97a6dbdbbe74eae41.
//
// Solidity: e DisputeCrowdsourcerCreated(universe indexed address, market indexed address, disputeCrowdsourcer address, payoutNumerators uint256[], size uint256, invalid bool)
func (_Augur *AugurFilterer) WatchDisputeCrowdsourcerCreated(opts *bind.WatchOpts, sink chan<- *AugurDisputeCrowdsourcerCreated, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DisputeCrowdsourcerCreated", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDisputeCrowdsourcerCreated)
				if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurDisputeCrowdsourcerRedeemedIterator is returned from FilterDisputeCrowdsourcerRedeemed and is used to iterate over the raw logs and unpacked data for DisputeCrowdsourcerRedeemed events raised by the Augur contract.
type AugurDisputeCrowdsourcerRedeemedIterator struct {
	Event *AugurDisputeCrowdsourcerRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDisputeCrowdsourcerRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDisputeCrowdsourcerRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDisputeCrowdsourcerRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDisputeCrowdsourcerRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDisputeCrowdsourcerRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDisputeCrowdsourcerRedeemed represents a DisputeCrowdsourcerRedeemed event raised by the Augur contract.
type AugurDisputeCrowdsourcerRedeemed struct {
	Universe              common.Address
	Reporter              common.Address
	Market                common.Address
	DisputeCrowdsourcer   common.Address
	AmountRedeemed        *big.Int
	RepReceived           *big.Int
	ReportingFeesReceived *big.Int
	PayoutNumerators      []*big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDisputeCrowdsourcerRedeemed is a free log retrieval operation binding the contract event 0x450bd662d3b1e236c8f344457690d257aeae5dca1add336752839ac206613cc0.
//
// Solidity: e DisputeCrowdsourcerRedeemed(universe indexed address, reporter indexed address, market indexed address, disputeCrowdsourcer address, amountRedeemed uint256, repReceived uint256, reportingFeesReceived uint256, payoutNumerators uint256[])
func (_Augur *AugurFilterer) FilterDisputeCrowdsourcerRedeemed(opts *bind.FilterOpts, universe []common.Address, reporter []common.Address, market []common.Address) (*AugurDisputeCrowdsourcerRedeemedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DisputeCrowdsourcerRedeemed", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurDisputeCrowdsourcerRedeemedIterator{contract: _Augur.contract, event: "DisputeCrowdsourcerRedeemed", logs: logs, sub: sub}, nil
}

// WatchDisputeCrowdsourcerRedeemed is a free log subscription operation binding the contract event 0x450bd662d3b1e236c8f344457690d257aeae5dca1add336752839ac206613cc0.
//
// Solidity: e DisputeCrowdsourcerRedeemed(universe indexed address, reporter indexed address, market indexed address, disputeCrowdsourcer address, amountRedeemed uint256, repReceived uint256, reportingFeesReceived uint256, payoutNumerators uint256[])
func (_Augur *AugurFilterer) WatchDisputeCrowdsourcerRedeemed(opts *bind.WatchOpts, sink chan<- *AugurDisputeCrowdsourcerRedeemed, universe []common.Address, reporter []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DisputeCrowdsourcerRedeemed", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDisputeCrowdsourcerRedeemed)
				if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurEscapeHatchChangedIterator is returned from FilterEscapeHatchChanged and is used to iterate over the raw logs and unpacked data for EscapeHatchChanged events raised by the Augur contract.
type AugurEscapeHatchChangedIterator struct {
	Event *AugurEscapeHatchChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurEscapeHatchChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurEscapeHatchChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurEscapeHatchChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurEscapeHatchChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurEscapeHatchChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurEscapeHatchChanged represents a EscapeHatchChanged event raised by the Augur contract.
type AugurEscapeHatchChanged struct {
	IsOn bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEscapeHatchChanged is a free log retrieval operation binding the contract event 0x4b6202764c3d77dc2a0c06f5b94ed9051fca8b678f707f3e514479a2bc91eb66.
//
// Solidity: e EscapeHatchChanged(isOn bool)
func (_Augur *AugurFilterer) FilterEscapeHatchChanged(opts *bind.FilterOpts) (*AugurEscapeHatchChangedIterator, error) {

	logs, sub, err := _Augur.contract.FilterLogs(opts, "EscapeHatchChanged")
	if err != nil {
		return nil, err
	}
	return &AugurEscapeHatchChangedIterator{contract: _Augur.contract, event: "EscapeHatchChanged", logs: logs, sub: sub}, nil
}

// WatchEscapeHatchChanged is a free log subscription operation binding the contract event 0x4b6202764c3d77dc2a0c06f5b94ed9051fca8b678f707f3e514479a2bc91eb66.
//
// Solidity: e EscapeHatchChanged(isOn bool)
func (_Augur *AugurFilterer) WatchEscapeHatchChanged(opts *bind.WatchOpts, sink chan<- *AugurEscapeHatchChanged) (event.Subscription, error) {

	logs, sub, err := _Augur.contract.WatchLogs(opts, "EscapeHatchChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurEscapeHatchChanged)
				if err := _Augur.contract.UnpackLog(event, "EscapeHatchChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurFeeWindowCreatedIterator is returned from FilterFeeWindowCreated and is used to iterate over the raw logs and unpacked data for FeeWindowCreated events raised by the Augur contract.
type AugurFeeWindowCreatedIterator struct {
	Event *AugurFeeWindowCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurFeeWindowCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurFeeWindowCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurFeeWindowCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurFeeWindowCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurFeeWindowCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurFeeWindowCreated represents a FeeWindowCreated event raised by the Augur contract.
type AugurFeeWindowCreated struct {
	Universe  common.Address
	FeeWindow common.Address
	StartTime *big.Int
	EndTime   *big.Int
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeWindowCreated is a free log retrieval operation binding the contract event 0xbaba17e31bb9fbfbc0b794111d2b1236ed4e36067a5e0d7c3c3433ad66c99f9d.
//
// Solidity: e FeeWindowCreated(universe indexed address, feeWindow address, startTime uint256, endTime uint256, id uint256)
func (_Augur *AugurFilterer) FilterFeeWindowCreated(opts *bind.FilterOpts, universe []common.Address) (*AugurFeeWindowCreatedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "FeeWindowCreated", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurFeeWindowCreatedIterator{contract: _Augur.contract, event: "FeeWindowCreated", logs: logs, sub: sub}, nil
}

// WatchFeeWindowCreated is a free log subscription operation binding the contract event 0xbaba17e31bb9fbfbc0b794111d2b1236ed4e36067a5e0d7c3c3433ad66c99f9d.
//
// Solidity: e FeeWindowCreated(universe indexed address, feeWindow address, startTime uint256, endTime uint256, id uint256)
func (_Augur *AugurFilterer) WatchFeeWindowCreated(opts *bind.WatchOpts, sink chan<- *AugurFeeWindowCreated, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "FeeWindowCreated", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurFeeWindowCreated)
				if err := _Augur.contract.UnpackLog(event, "FeeWindowCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurFeeWindowRedeemedIterator is returned from FilterFeeWindowRedeemed and is used to iterate over the raw logs and unpacked data for FeeWindowRedeemed events raised by the Augur contract.
type AugurFeeWindowRedeemedIterator struct {
	Event *AugurFeeWindowRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurFeeWindowRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurFeeWindowRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurFeeWindowRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurFeeWindowRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurFeeWindowRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurFeeWindowRedeemed represents a FeeWindowRedeemed event raised by the Augur contract.
type AugurFeeWindowRedeemed struct {
	Universe              common.Address
	Reporter              common.Address
	FeeWindow             common.Address
	AmountRedeemed        *big.Int
	ReportingFeesReceived *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterFeeWindowRedeemed is a free log retrieval operation binding the contract event 0xc62cff53848fe243adb6130140cfe557ce16e8006861abd50adfe425150ba6c5.
//
// Solidity: e FeeWindowRedeemed(universe indexed address, reporter indexed address, feeWindow indexed address, amountRedeemed uint256, reportingFeesReceived uint256)
func (_Augur *AugurFilterer) FilterFeeWindowRedeemed(opts *bind.FilterOpts, universe []common.Address, reporter []common.Address, feeWindow []common.Address) (*AugurFeeWindowRedeemedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var feeWindowRule []interface{}
	for _, feeWindowItem := range feeWindow {
		feeWindowRule = append(feeWindowRule, feeWindowItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "FeeWindowRedeemed", universeRule, reporterRule, feeWindowRule)
	if err != nil {
		return nil, err
	}
	return &AugurFeeWindowRedeemedIterator{contract: _Augur.contract, event: "FeeWindowRedeemed", logs: logs, sub: sub}, nil
}

// WatchFeeWindowRedeemed is a free log subscription operation binding the contract event 0xc62cff53848fe243adb6130140cfe557ce16e8006861abd50adfe425150ba6c5.
//
// Solidity: e FeeWindowRedeemed(universe indexed address, reporter indexed address, feeWindow indexed address, amountRedeemed uint256, reportingFeesReceived uint256)
func (_Augur *AugurFilterer) WatchFeeWindowRedeemed(opts *bind.WatchOpts, sink chan<- *AugurFeeWindowRedeemed, universe []common.Address, reporter []common.Address, feeWindow []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var feeWindowRule []interface{}
	for _, feeWindowItem := range feeWindow {
		feeWindowRule = append(feeWindowRule, feeWindowItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "FeeWindowRedeemed", universeRule, reporterRule, feeWindowRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurFeeWindowRedeemed)
				if err := _Augur.contract.UnpackLog(event, "FeeWindowRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurInitialReportSubmittedIterator is returned from FilterInitialReportSubmitted and is used to iterate over the raw logs and unpacked data for InitialReportSubmitted events raised by the Augur contract.
type AugurInitialReportSubmittedIterator struct {
	Event *AugurInitialReportSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurInitialReportSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurInitialReportSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurInitialReportSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurInitialReportSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurInitialReportSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurInitialReportSubmitted represents a InitialReportSubmitted event raised by the Augur contract.
type AugurInitialReportSubmitted struct {
	Universe             common.Address
	Reporter             common.Address
	Market               common.Address
	AmountStaked         *big.Int
	IsDesignatedReporter bool
	PayoutNumerators     []*big.Int
	Invalid              bool
	Description          string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterInitialReportSubmitted is a free log retrieval operation binding the contract event 0x5f28d90f610b97bb029fe74dfab936367e76c0fe576309e0b65cf72a975beac4.
//
// Solidity: e InitialReportSubmitted(universe indexed address, reporter indexed address, market indexed address, amountStaked uint256, isDesignatedReporter bool, payoutNumerators uint256[], invalid bool, description string)
func (_Augur *AugurFilterer) FilterInitialReportSubmitted(opts *bind.FilterOpts, universe []common.Address, reporter []common.Address, market []common.Address) (*AugurInitialReportSubmittedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "InitialReportSubmitted", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurInitialReportSubmittedIterator{contract: _Augur.contract, event: "InitialReportSubmitted", logs: logs, sub: sub}, nil
}

// WatchInitialReportSubmitted is a free log subscription operation binding the contract event 0x5f28d90f610b97bb029fe74dfab936367e76c0fe576309e0b65cf72a975beac4.
//
// Solidity: e InitialReportSubmitted(universe indexed address, reporter indexed address, market indexed address, amountStaked uint256, isDesignatedReporter bool, payoutNumerators uint256[], invalid bool, description string)
func (_Augur *AugurFilterer) WatchInitialReportSubmitted(opts *bind.WatchOpts, sink chan<- *AugurInitialReportSubmitted, universe []common.Address, reporter []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "InitialReportSubmitted", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurInitialReportSubmitted)
				if err := _Augur.contract.UnpackLog(event, "InitialReportSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurInitialReporterRedeemedIterator is returned from FilterInitialReporterRedeemed and is used to iterate over the raw logs and unpacked data for InitialReporterRedeemed events raised by the Augur contract.
type AugurInitialReporterRedeemedIterator struct {
	Event *AugurInitialReporterRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurInitialReporterRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurInitialReporterRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurInitialReporterRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurInitialReporterRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurInitialReporterRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurInitialReporterRedeemed represents a InitialReporterRedeemed event raised by the Augur contract.
type AugurInitialReporterRedeemed struct {
	Universe              common.Address
	Reporter              common.Address
	Market                common.Address
	AmountRedeemed        *big.Int
	RepReceived           *big.Int
	ReportingFeesReceived *big.Int
	PayoutNumerators      []*big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterInitialReporterRedeemed is a free log retrieval operation binding the contract event 0xdd0dca2d338dc86ba5431017bdf6f3ad45247d608b0a38d866e3131a876be2cc.
//
// Solidity: e InitialReporterRedeemed(universe indexed address, reporter indexed address, market indexed address, amountRedeemed uint256, repReceived uint256, reportingFeesReceived uint256, payoutNumerators uint256[])
func (_Augur *AugurFilterer) FilterInitialReporterRedeemed(opts *bind.FilterOpts, universe []common.Address, reporter []common.Address, market []common.Address) (*AugurInitialReporterRedeemedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "InitialReporterRedeemed", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurInitialReporterRedeemedIterator{contract: _Augur.contract, event: "InitialReporterRedeemed", logs: logs, sub: sub}, nil
}

// WatchInitialReporterRedeemed is a free log subscription operation binding the contract event 0xdd0dca2d338dc86ba5431017bdf6f3ad45247d608b0a38d866e3131a876be2cc.
//
// Solidity: e InitialReporterRedeemed(universe indexed address, reporter indexed address, market indexed address, amountRedeemed uint256, repReceived uint256, reportingFeesReceived uint256, payoutNumerators uint256[])
func (_Augur *AugurFilterer) WatchInitialReporterRedeemed(opts *bind.WatchOpts, sink chan<- *AugurInitialReporterRedeemed, universe []common.Address, reporter []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "InitialReporterRedeemed", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurInitialReporterRedeemed)
				if err := _Augur.contract.UnpackLog(event, "InitialReporterRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurInitialReporterTransferredIterator is returned from FilterInitialReporterTransferred and is used to iterate over the raw logs and unpacked data for InitialReporterTransferred events raised by the Augur contract.
type AugurInitialReporterTransferredIterator struct {
	Event *AugurInitialReporterTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurInitialReporterTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurInitialReporterTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurInitialReporterTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurInitialReporterTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurInitialReporterTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurInitialReporterTransferred represents a InitialReporterTransferred event raised by the Augur contract.
type AugurInitialReporterTransferred struct {
	Universe common.Address
	Market   common.Address
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitialReporterTransferred is a free log retrieval operation binding the contract event 0xee62c58e2603b92f96a002e012f4f3bd5748102cfa3b711f6d778c6237fcaa96.
//
// Solidity: e InitialReporterTransferred(universe indexed address, market indexed address, from address, to address)
func (_Augur *AugurFilterer) FilterInitialReporterTransferred(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurInitialReporterTransferredIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "InitialReporterTransferred", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurInitialReporterTransferredIterator{contract: _Augur.contract, event: "InitialReporterTransferred", logs: logs, sub: sub}, nil
}

// WatchInitialReporterTransferred is a free log subscription operation binding the contract event 0xee62c58e2603b92f96a002e012f4f3bd5748102cfa3b711f6d778c6237fcaa96.
//
// Solidity: e InitialReporterTransferred(universe indexed address, market indexed address, from address, to address)
func (_Augur *AugurFilterer) WatchInitialReporterTransferred(opts *bind.WatchOpts, sink chan<- *AugurInitialReporterTransferred, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "InitialReporterTransferred", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurInitialReporterTransferred)
				if err := _Augur.contract.UnpackLog(event, "InitialReporterTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the Augur contract.
type AugurMarketCreatedIterator struct {
	Event *AugurMarketCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketCreated represents a MarketCreated event raised by the Augur contract.
type AugurMarketCreated struct {
	Topic             [32]byte
	Description       string
	ExtraInfo         string
	Universe          common.Address
	Market            common.Address
	MarketCreator     common.Address
	Outcomes          [][32]byte
	MarketCreationFee *big.Int
	MinPrice          *big.Int
	MaxPrice          *big.Int
	MarketType        uint8
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0xb2e65de73007eef46316e4f18ab1f301b4d0e31aa56733387b469612f90894df.
//
// Solidity: e MarketCreated(topic indexed bytes32, description string, extraInfo string, universe indexed address, market address, marketCreator indexed address, outcomes bytes32[], marketCreationFee uint256, minPrice int256, maxPrice int256, marketType uint8)
func (_Augur *AugurFilterer) FilterMarketCreated(opts *bind.FilterOpts, topic [][32]byte, universe []common.Address, marketCreator []common.Address) (*AugurMarketCreatedIterator, error) {

	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	var marketCreatorRule []interface{}
	for _, marketCreatorItem := range marketCreator {
		marketCreatorRule = append(marketCreatorRule, marketCreatorItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketCreated", topicRule, universeRule, marketCreatorRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketCreatedIterator{contract: _Augur.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0xb2e65de73007eef46316e4f18ab1f301b4d0e31aa56733387b469612f90894df.
//
// Solidity: e MarketCreated(topic indexed bytes32, description string, extraInfo string, universe indexed address, market address, marketCreator indexed address, outcomes bytes32[], marketCreationFee uint256, minPrice int256, maxPrice int256, marketType uint8)
func (_Augur *AugurFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *AugurMarketCreated, topic [][32]byte, universe []common.Address, marketCreator []common.Address) (event.Subscription, error) {

	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	var marketCreatorRule []interface{}
	for _, marketCreatorItem := range marketCreator {
		marketCreatorRule = append(marketCreatorRule, marketCreatorItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketCreated", topicRule, universeRule, marketCreatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketCreated)
				if err := _Augur.contract.UnpackLog(event, "MarketCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurMarketFinalizedIterator is returned from FilterMarketFinalized and is used to iterate over the raw logs and unpacked data for MarketFinalized events raised by the Augur contract.
type AugurMarketFinalizedIterator struct {
	Event *AugurMarketFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketFinalized represents a MarketFinalized event raised by the Augur contract.
type AugurMarketFinalized struct {
	Universe common.Address
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketFinalized is a free log retrieval operation binding the contract event 0x014ce4e12965529d7d31e11411d7a23b1778d448ab763ffc4d55830cbb4919d7.
//
// Solidity: e MarketFinalized(universe indexed address, market indexed address)
func (_Augur *AugurFilterer) FilterMarketFinalized(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurMarketFinalizedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketFinalized", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketFinalizedIterator{contract: _Augur.contract, event: "MarketFinalized", logs: logs, sub: sub}, nil
}

// WatchMarketFinalized is a free log subscription operation binding the contract event 0x014ce4e12965529d7d31e11411d7a23b1778d448ab763ffc4d55830cbb4919d7.
//
// Solidity: e MarketFinalized(universe indexed address, market indexed address)
func (_Augur *AugurFilterer) WatchMarketFinalized(opts *bind.WatchOpts, sink chan<- *AugurMarketFinalized, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketFinalized", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketFinalized)
				if err := _Augur.contract.UnpackLog(event, "MarketFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurMarketMailboxTransferredIterator is returned from FilterMarketMailboxTransferred and is used to iterate over the raw logs and unpacked data for MarketMailboxTransferred events raised by the Augur contract.
type AugurMarketMailboxTransferredIterator struct {
	Event *AugurMarketMailboxTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketMailboxTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketMailboxTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketMailboxTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketMailboxTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketMailboxTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketMailboxTransferred represents a MarketMailboxTransferred event raised by the Augur contract.
type AugurMarketMailboxTransferred struct {
	Universe common.Address
	Market   common.Address
	Mailbox  common.Address
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketMailboxTransferred is a free log retrieval operation binding the contract event 0x8a34ec183bf620d74d0b52e71165bb4255b0591d1c8e9d07c707a7f1d763158d.
//
// Solidity: e MarketMailboxTransferred(universe indexed address, market indexed address, mailbox indexed address, from address, to address)
func (_Augur *AugurFilterer) FilterMarketMailboxTransferred(opts *bind.FilterOpts, universe []common.Address, market []common.Address, mailbox []common.Address) (*AugurMarketMailboxTransferredIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var mailboxRule []interface{}
	for _, mailboxItem := range mailbox {
		mailboxRule = append(mailboxRule, mailboxItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketMailboxTransferred", universeRule, marketRule, mailboxRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketMailboxTransferredIterator{contract: _Augur.contract, event: "MarketMailboxTransferred", logs: logs, sub: sub}, nil
}

// WatchMarketMailboxTransferred is a free log subscription operation binding the contract event 0x8a34ec183bf620d74d0b52e71165bb4255b0591d1c8e9d07c707a7f1d763158d.
//
// Solidity: e MarketMailboxTransferred(universe indexed address, market indexed address, mailbox indexed address, from address, to address)
func (_Augur *AugurFilterer) WatchMarketMailboxTransferred(opts *bind.WatchOpts, sink chan<- *AugurMarketMailboxTransferred, universe []common.Address, market []common.Address, mailbox []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var mailboxRule []interface{}
	for _, mailboxItem := range mailbox {
		mailboxRule = append(mailboxRule, mailboxItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketMailboxTransferred", universeRule, marketRule, mailboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketMailboxTransferred)
				if err := _Augur.contract.UnpackLog(event, "MarketMailboxTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurMarketMigratedIterator is returned from FilterMarketMigrated and is used to iterate over the raw logs and unpacked data for MarketMigrated events raised by the Augur contract.
type AugurMarketMigratedIterator struct {
	Event *AugurMarketMigrated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketMigrated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketMigrated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketMigrated represents a MarketMigrated event raised by the Augur contract.
type AugurMarketMigrated struct {
	Market           common.Address
	OriginalUniverse common.Address
	NewUniverse      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMarketMigrated is a free log retrieval operation binding the contract event 0xc3cf07f8fa0fafc25a9dd0bad2cd6b961c55dad41b42c8ef8f931bc40e41e08c.
//
// Solidity: e MarketMigrated(market indexed address, originalUniverse indexed address, newUniverse indexed address)
func (_Augur *AugurFilterer) FilterMarketMigrated(opts *bind.FilterOpts, market []common.Address, originalUniverse []common.Address, newUniverse []common.Address) (*AugurMarketMigratedIterator, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var originalUniverseRule []interface{}
	for _, originalUniverseItem := range originalUniverse {
		originalUniverseRule = append(originalUniverseRule, originalUniverseItem)
	}
	var newUniverseRule []interface{}
	for _, newUniverseItem := range newUniverse {
		newUniverseRule = append(newUniverseRule, newUniverseItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketMigrated", marketRule, originalUniverseRule, newUniverseRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketMigratedIterator{contract: _Augur.contract, event: "MarketMigrated", logs: logs, sub: sub}, nil
}

// WatchMarketMigrated is a free log subscription operation binding the contract event 0xc3cf07f8fa0fafc25a9dd0bad2cd6b961c55dad41b42c8ef8f931bc40e41e08c.
//
// Solidity: e MarketMigrated(market indexed address, originalUniverse indexed address, newUniverse indexed address)
func (_Augur *AugurFilterer) WatchMarketMigrated(opts *bind.WatchOpts, sink chan<- *AugurMarketMigrated, market []common.Address, originalUniverse []common.Address, newUniverse []common.Address) (event.Subscription, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var originalUniverseRule []interface{}
	for _, originalUniverseItem := range originalUniverse {
		originalUniverseRule = append(originalUniverseRule, originalUniverseItem)
	}
	var newUniverseRule []interface{}
	for _, newUniverseItem := range newUniverse {
		newUniverseRule = append(newUniverseRule, newUniverseItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketMigrated", marketRule, originalUniverseRule, newUniverseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketMigrated)
				if err := _Augur.contract.UnpackLog(event, "MarketMigrated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurMarketParticipantsDisavowedIterator is returned from FilterMarketParticipantsDisavowed and is used to iterate over the raw logs and unpacked data for MarketParticipantsDisavowed events raised by the Augur contract.
type AugurMarketParticipantsDisavowedIterator struct {
	Event *AugurMarketParticipantsDisavowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketParticipantsDisavowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketParticipantsDisavowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketParticipantsDisavowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketParticipantsDisavowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketParticipantsDisavowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketParticipantsDisavowed represents a MarketParticipantsDisavowed event raised by the Augur contract.
type AugurMarketParticipantsDisavowed struct {
	Universe common.Address
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketParticipantsDisavowed is a free log retrieval operation binding the contract event 0x3b4f3db017516414df2695e5b0052661779d7163a6cd4368fd74313be73fa0b8.
//
// Solidity: e MarketParticipantsDisavowed(universe indexed address, market indexed address)
func (_Augur *AugurFilterer) FilterMarketParticipantsDisavowed(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurMarketParticipantsDisavowedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketParticipantsDisavowed", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketParticipantsDisavowedIterator{contract: _Augur.contract, event: "MarketParticipantsDisavowed", logs: logs, sub: sub}, nil
}

// WatchMarketParticipantsDisavowed is a free log subscription operation binding the contract event 0x3b4f3db017516414df2695e5b0052661779d7163a6cd4368fd74313be73fa0b8.
//
// Solidity: e MarketParticipantsDisavowed(universe indexed address, market indexed address)
func (_Augur *AugurFilterer) WatchMarketParticipantsDisavowed(opts *bind.WatchOpts, sink chan<- *AugurMarketParticipantsDisavowed, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketParticipantsDisavowed", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketParticipantsDisavowed)
				if err := _Augur.contract.UnpackLog(event, "MarketParticipantsDisavowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurMarketTransferredIterator is returned from FilterMarketTransferred and is used to iterate over the raw logs and unpacked data for MarketTransferred events raised by the Augur contract.
type AugurMarketTransferredIterator struct {
	Event *AugurMarketTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketTransferred represents a MarketTransferred event raised by the Augur contract.
type AugurMarketTransferred struct {
	Universe common.Address
	Market   common.Address
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketTransferred is a free log retrieval operation binding the contract event 0x55f2a7bfa32e835c3f3c3cff653a3d11c077ce1b00c5a41c6aaf09eedc1ac3b2.
//
// Solidity: e MarketTransferred(universe indexed address, market indexed address, from address, to address)
func (_Augur *AugurFilterer) FilterMarketTransferred(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurMarketTransferredIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketTransferred", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketTransferredIterator{contract: _Augur.contract, event: "MarketTransferred", logs: logs, sub: sub}, nil
}

// WatchMarketTransferred is a free log subscription operation binding the contract event 0x55f2a7bfa32e835c3f3c3cff653a3d11c077ce1b00c5a41c6aaf09eedc1ac3b2.
//
// Solidity: e MarketTransferred(universe indexed address, market indexed address, from address, to address)
func (_Augur *AugurFilterer) WatchMarketTransferred(opts *bind.WatchOpts, sink chan<- *AugurMarketTransferred, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketTransferred", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketTransferred)
				if err := _Augur.contract.UnpackLog(event, "MarketTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurOrderCanceledIterator is returned from FilterOrderCanceled and is used to iterate over the raw logs and unpacked data for OrderCanceled events raised by the Augur contract.
type AugurOrderCanceledIterator struct {
	Event *AugurOrderCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurOrderCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurOrderCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurOrderCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurOrderCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurOrderCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurOrderCanceled represents a OrderCanceled event raised by the Augur contract.
type AugurOrderCanceled struct {
	Universe     common.Address
	ShareToken   common.Address
	Sender       common.Address
	OrderId      [32]byte
	OrderType    uint8
	TokenRefund  *big.Int
	SharesRefund *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderCanceled is a free log retrieval operation binding the contract event 0x513d029ff62330c16d8d4b36b28fab53f09d10bb51b56fe121ab710ca2d1af80.
//
// Solidity: e OrderCanceled(universe indexed address, shareToken indexed address, sender indexed address, orderId bytes32, orderType uint8, tokenRefund uint256, sharesRefund uint256)
func (_Augur *AugurFilterer) FilterOrderCanceled(opts *bind.FilterOpts, universe []common.Address, shareToken []common.Address, sender []common.Address) (*AugurOrderCanceledIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var shareTokenRule []interface{}
	for _, shareTokenItem := range shareToken {
		shareTokenRule = append(shareTokenRule, shareTokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "OrderCanceled", universeRule, shareTokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AugurOrderCanceledIterator{contract: _Augur.contract, event: "OrderCanceled", logs: logs, sub: sub}, nil
}

// WatchOrderCanceled is a free log subscription operation binding the contract event 0x513d029ff62330c16d8d4b36b28fab53f09d10bb51b56fe121ab710ca2d1af80.
//
// Solidity: e OrderCanceled(universe indexed address, shareToken indexed address, sender indexed address, orderId bytes32, orderType uint8, tokenRefund uint256, sharesRefund uint256)
func (_Augur *AugurFilterer) WatchOrderCanceled(opts *bind.WatchOpts, sink chan<- *AugurOrderCanceled, universe []common.Address, shareToken []common.Address, sender []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var shareTokenRule []interface{}
	for _, shareTokenItem := range shareToken {
		shareTokenRule = append(shareTokenRule, shareTokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "OrderCanceled", universeRule, shareTokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurOrderCanceled)
				if err := _Augur.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the Augur contract.
type AugurOrderCreatedIterator struct {
	Event *AugurOrderCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurOrderCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurOrderCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurOrderCreated represents a OrderCreated event raised by the Augur contract.
type AugurOrderCreated struct {
	OrderType      uint8
	Amount         *big.Int
	Price          *big.Int
	Creator        common.Address
	MoneyEscrowed  *big.Int
	SharesEscrowed *big.Int
	TradeGroupId   [32]byte
	OrderId        [32]byte
	Universe       common.Address
	ShareToken     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x32d554e498d0c7f2a5c7fd8b6b234bfc4e1dfb5290466d998af09a813db32f31.
//
// Solidity: e OrderCreated(orderType uint8, amount uint256, price uint256, creator indexed address, moneyEscrowed uint256, sharesEscrowed uint256, tradeGroupId bytes32, orderId bytes32, universe indexed address, shareToken indexed address)
func (_Augur *AugurFilterer) FilterOrderCreated(opts *bind.FilterOpts, creator []common.Address, universe []common.Address, shareToken []common.Address) (*AugurOrderCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var shareTokenRule []interface{}
	for _, shareTokenItem := range shareToken {
		shareTokenRule = append(shareTokenRule, shareTokenItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "OrderCreated", creatorRule, universeRule, shareTokenRule)
	if err != nil {
		return nil, err
	}
	return &AugurOrderCreatedIterator{contract: _Augur.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x32d554e498d0c7f2a5c7fd8b6b234bfc4e1dfb5290466d998af09a813db32f31.
//
// Solidity: e OrderCreated(orderType uint8, amount uint256, price uint256, creator indexed address, moneyEscrowed uint256, sharesEscrowed uint256, tradeGroupId bytes32, orderId bytes32, universe indexed address, shareToken indexed address)
func (_Augur *AugurFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *AugurOrderCreated, creator []common.Address, universe []common.Address, shareToken []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var shareTokenRule []interface{}
	for _, shareTokenItem := range shareToken {
		shareTokenRule = append(shareTokenRule, shareTokenItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "OrderCreated", creatorRule, universeRule, shareTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurOrderCreated)
				if err := _Augur.contract.UnpackLog(event, "OrderCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the Augur contract.
type AugurOrderFilledIterator struct {
	Event *AugurOrderFilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurOrderFilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurOrderFilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurOrderFilled represents a OrderFilled event raised by the Augur contract.
type AugurOrderFilled struct {
	Universe          common.Address
	ShareToken        common.Address
	Filler            common.Address
	OrderId           [32]byte
	NumCreatorShares  *big.Int
	NumCreatorTokens  *big.Int
	NumFillerShares   *big.Int
	NumFillerTokens   *big.Int
	MarketCreatorFees *big.Int
	ReporterFees      *big.Int
	AmountFilled      *big.Int
	TradeGroupId      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xabb970462c1f0de9e237d127ad47c01c4e69caa179fd850d076ae9bfc529176e.
//
// Solidity: e OrderFilled(universe indexed address, shareToken indexed address, filler address, orderId bytes32, numCreatorShares uint256, numCreatorTokens uint256, numFillerShares uint256, numFillerTokens uint256, marketCreatorFees uint256, reporterFees uint256, amountFilled uint256, tradeGroupId bytes32)
func (_Augur *AugurFilterer) FilterOrderFilled(opts *bind.FilterOpts, universe []common.Address, shareToken []common.Address) (*AugurOrderFilledIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var shareTokenRule []interface{}
	for _, shareTokenItem := range shareToken {
		shareTokenRule = append(shareTokenRule, shareTokenItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "OrderFilled", universeRule, shareTokenRule)
	if err != nil {
		return nil, err
	}
	return &AugurOrderFilledIterator{contract: _Augur.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xabb970462c1f0de9e237d127ad47c01c4e69caa179fd850d076ae9bfc529176e.
//
// Solidity: e OrderFilled(universe indexed address, shareToken indexed address, filler address, orderId bytes32, numCreatorShares uint256, numCreatorTokens uint256, numFillerShares uint256, numFillerTokens uint256, marketCreatorFees uint256, reporterFees uint256, amountFilled uint256, tradeGroupId bytes32)
func (_Augur *AugurFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *AugurOrderFilled, universe []common.Address, shareToken []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var shareTokenRule []interface{}
	for _, shareTokenItem := range shareToken {
		shareTokenRule = append(shareTokenRule, shareTokenItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "OrderFilled", universeRule, shareTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurOrderFilled)
				if err := _Augur.contract.UnpackLog(event, "OrderFilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurReportingParticipantDisavowedIterator is returned from FilterReportingParticipantDisavowed and is used to iterate over the raw logs and unpacked data for ReportingParticipantDisavowed events raised by the Augur contract.
type AugurReportingParticipantDisavowedIterator struct {
	Event *AugurReportingParticipantDisavowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurReportingParticipantDisavowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurReportingParticipantDisavowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurReportingParticipantDisavowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurReportingParticipantDisavowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurReportingParticipantDisavowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurReportingParticipantDisavowed represents a ReportingParticipantDisavowed event raised by the Augur contract.
type AugurReportingParticipantDisavowed struct {
	Universe             common.Address
	Market               common.Address
	ReportingParticipant common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterReportingParticipantDisavowed is a free log retrieval operation binding the contract event 0xb20adf682c8f82b94a135452f54ac4483c9ee8c9b2324e946120696ab1d034b4.
//
// Solidity: e ReportingParticipantDisavowed(universe indexed address, market indexed address, reportingParticipant address)
func (_Augur *AugurFilterer) FilterReportingParticipantDisavowed(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurReportingParticipantDisavowedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "ReportingParticipantDisavowed", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurReportingParticipantDisavowedIterator{contract: _Augur.contract, event: "ReportingParticipantDisavowed", logs: logs, sub: sub}, nil
}

// WatchReportingParticipantDisavowed is a free log subscription operation binding the contract event 0xb20adf682c8f82b94a135452f54ac4483c9ee8c9b2324e946120696ab1d034b4.
//
// Solidity: e ReportingParticipantDisavowed(universe indexed address, market indexed address, reportingParticipant address)
func (_Augur *AugurFilterer) WatchReportingParticipantDisavowed(opts *bind.WatchOpts, sink chan<- *AugurReportingParticipantDisavowed, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "ReportingParticipantDisavowed", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurReportingParticipantDisavowed)
				if err := _Augur.contract.UnpackLog(event, "ReportingParticipantDisavowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurTimestampSetIterator is returned from FilterTimestampSet and is used to iterate over the raw logs and unpacked data for TimestampSet events raised by the Augur contract.
type AugurTimestampSetIterator struct {
	Event *AugurTimestampSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTimestampSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTimestampSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTimestampSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTimestampSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTimestampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTimestampSet represents a TimestampSet event raised by the Augur contract.
type AugurTimestampSet struct {
	NewTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTimestampSet is a free log retrieval operation binding the contract event 0x11dda748f0bd3af85a073da0088a0acb827d9584a4fdb825c81f1232a5309538.
//
// Solidity: e TimestampSet(newTimestamp uint256)
func (_Augur *AugurFilterer) FilterTimestampSet(opts *bind.FilterOpts) (*AugurTimestampSetIterator, error) {

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TimestampSet")
	if err != nil {
		return nil, err
	}
	return &AugurTimestampSetIterator{contract: _Augur.contract, event: "TimestampSet", logs: logs, sub: sub}, nil
}

// WatchTimestampSet is a free log subscription operation binding the contract event 0x11dda748f0bd3af85a073da0088a0acb827d9584a4fdb825c81f1232a5309538.
//
// Solidity: e TimestampSet(newTimestamp uint256)
func (_Augur *AugurFilterer) WatchTimestampSet(opts *bind.WatchOpts, sink chan<- *AugurTimestampSet) (event.Subscription, error) {

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TimestampSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTimestampSet)
				if err := _Augur.contract.UnpackLog(event, "TimestampSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurTokensBurnedIterator is returned from FilterTokensBurned and is used to iterate over the raw logs and unpacked data for TokensBurned events raised by the Augur contract.
type AugurTokensBurnedIterator struct {
	Event *AugurTokensBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTokensBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTokensBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTokensBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTokensBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTokensBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTokensBurned represents a TokensBurned event raised by the Augur contract.
type AugurTokensBurned struct {
	Universe  common.Address
	Token     common.Address
	Target    common.Address
	Amount    *big.Int
	TokenType uint8
	Market    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensBurned is a free log retrieval operation binding the contract event 0x262b80f2af08a1001d15a1df91dde9acb8441811543886659b3845a8c285748b.
//
// Solidity: e TokensBurned(universe indexed address, token indexed address, target indexed address, amount uint256, tokenType uint8, market address)
func (_Augur *AugurFilterer) FilterTokensBurned(opts *bind.FilterOpts, universe []common.Address, token []common.Address, target []common.Address) (*AugurTokensBurnedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TokensBurned", universeRule, tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AugurTokensBurnedIterator{contract: _Augur.contract, event: "TokensBurned", logs: logs, sub: sub}, nil
}

// WatchTokensBurned is a free log subscription operation binding the contract event 0x262b80f2af08a1001d15a1df91dde9acb8441811543886659b3845a8c285748b.
//
// Solidity: e TokensBurned(universe indexed address, token indexed address, target indexed address, amount uint256, tokenType uint8, market address)
func (_Augur *AugurFilterer) WatchTokensBurned(opts *bind.WatchOpts, sink chan<- *AugurTokensBurned, universe []common.Address, token []common.Address, target []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TokensBurned", universeRule, tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTokensBurned)
				if err := _Augur.contract.UnpackLog(event, "TokensBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurTokensMintedIterator is returned from FilterTokensMinted and is used to iterate over the raw logs and unpacked data for TokensMinted events raised by the Augur contract.
type AugurTokensMintedIterator struct {
	Event *AugurTokensMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTokensMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTokensMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTokensMinted represents a TokensMinted event raised by the Augur contract.
type AugurTokensMinted struct {
	Universe  common.Address
	Token     common.Address
	Target    common.Address
	Amount    *big.Int
	TokenType uint8
	Market    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensMinted is a free log retrieval operation binding the contract event 0x75dd618f69c0f07adc97fe19ba435f3932ce6aa8cad287fb9bdfaf37639f703a.
//
// Solidity: e TokensMinted(universe indexed address, token indexed address, target indexed address, amount uint256, tokenType uint8, market address)
func (_Augur *AugurFilterer) FilterTokensMinted(opts *bind.FilterOpts, universe []common.Address, token []common.Address, target []common.Address) (*AugurTokensMintedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TokensMinted", universeRule, tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AugurTokensMintedIterator{contract: _Augur.contract, event: "TokensMinted", logs: logs, sub: sub}, nil
}

// WatchTokensMinted is a free log subscription operation binding the contract event 0x75dd618f69c0f07adc97fe19ba435f3932ce6aa8cad287fb9bdfaf37639f703a.
//
// Solidity: e TokensMinted(universe indexed address, token indexed address, target indexed address, amount uint256, tokenType uint8, market address)
func (_Augur *AugurFilterer) WatchTokensMinted(opts *bind.WatchOpts, sink chan<- *AugurTokensMinted, universe []common.Address, token []common.Address, target []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TokensMinted", universeRule, tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTokensMinted)
				if err := _Augur.contract.UnpackLog(event, "TokensMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurTokensTransferredIterator is returned from FilterTokensTransferred and is used to iterate over the raw logs and unpacked data for TokensTransferred events raised by the Augur contract.
type AugurTokensTransferredIterator struct {
	Event *AugurTokensTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTokensTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTokensTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTokensTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTokensTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTokensTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTokensTransferred represents a TokensTransferred event raised by the Augur contract.
type AugurTokensTransferred struct {
	Universe  common.Address
	Token     common.Address
	From      common.Address
	To        common.Address
	Value     *big.Int
	TokenType uint8
	Market    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensTransferred is a free log retrieval operation binding the contract event 0x3c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf.
//
// Solidity: e TokensTransferred(universe indexed address, token indexed address, from indexed address, to address, value uint256, tokenType uint8, market address)
func (_Augur *AugurFilterer) FilterTokensTransferred(opts *bind.FilterOpts, universe []common.Address, token []common.Address, from []common.Address) (*AugurTokensTransferredIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TokensTransferred", universeRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &AugurTokensTransferredIterator{contract: _Augur.contract, event: "TokensTransferred", logs: logs, sub: sub}, nil
}

// WatchTokensTransferred is a free log subscription operation binding the contract event 0x3c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf.
//
// Solidity: e TokensTransferred(universe indexed address, token indexed address, from indexed address, to address, value uint256, tokenType uint8, market address)
func (_Augur *AugurFilterer) WatchTokensTransferred(opts *bind.WatchOpts, sink chan<- *AugurTokensTransferred, universe []common.Address, token []common.Address, from []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TokensTransferred", universeRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTokensTransferred)
				if err := _Augur.contract.UnpackLog(event, "TokensTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurTradingProceedsClaimedIterator is returned from FilterTradingProceedsClaimed and is used to iterate over the raw logs and unpacked data for TradingProceedsClaimed events raised by the Augur contract.
type AugurTradingProceedsClaimedIterator struct {
	Event *AugurTradingProceedsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTradingProceedsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTradingProceedsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTradingProceedsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTradingProceedsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTradingProceedsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTradingProceedsClaimed represents a TradingProceedsClaimed event raised by the Augur contract.
type AugurTradingProceedsClaimed struct {
	Universe          common.Address
	ShareToken        common.Address
	Sender            common.Address
	Market            common.Address
	NumShares         *big.Int
	NumPayoutTokens   *big.Int
	FinalTokenBalance *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTradingProceedsClaimed is a free log retrieval operation binding the contract event 0xa7e9373569caad2b7871ecb4d498619fc1c42840a6c0dbeb8dff20b131721e50.
//
// Solidity: e TradingProceedsClaimed(universe indexed address, shareToken indexed address, sender indexed address, market address, numShares uint256, numPayoutTokens uint256, finalTokenBalance uint256)
func (_Augur *AugurFilterer) FilterTradingProceedsClaimed(opts *bind.FilterOpts, universe []common.Address, shareToken []common.Address, sender []common.Address) (*AugurTradingProceedsClaimedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var shareTokenRule []interface{}
	for _, shareTokenItem := range shareToken {
		shareTokenRule = append(shareTokenRule, shareTokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TradingProceedsClaimed", universeRule, shareTokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AugurTradingProceedsClaimedIterator{contract: _Augur.contract, event: "TradingProceedsClaimed", logs: logs, sub: sub}, nil
}

// WatchTradingProceedsClaimed is a free log subscription operation binding the contract event 0xa7e9373569caad2b7871ecb4d498619fc1c42840a6c0dbeb8dff20b131721e50.
//
// Solidity: e TradingProceedsClaimed(universe indexed address, shareToken indexed address, sender indexed address, market address, numShares uint256, numPayoutTokens uint256, finalTokenBalance uint256)
func (_Augur *AugurFilterer) WatchTradingProceedsClaimed(opts *bind.WatchOpts, sink chan<- *AugurTradingProceedsClaimed, universe []common.Address, shareToken []common.Address, sender []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var shareTokenRule []interface{}
	for _, shareTokenItem := range shareToken {
		shareTokenRule = append(shareTokenRule, shareTokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TradingProceedsClaimed", universeRule, shareTokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTradingProceedsClaimed)
				if err := _Augur.contract.UnpackLog(event, "TradingProceedsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurUniverseCreatedIterator is returned from FilterUniverseCreated and is used to iterate over the raw logs and unpacked data for UniverseCreated events raised by the Augur contract.
type AugurUniverseCreatedIterator struct {
	Event *AugurUniverseCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurUniverseCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurUniverseCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurUniverseCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurUniverseCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurUniverseCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurUniverseCreated represents a UniverseCreated event raised by the Augur contract.
type AugurUniverseCreated struct {
	ParentUniverse   common.Address
	ChildUniverse    common.Address
	PayoutNumerators []*big.Int
	Invalid          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUniverseCreated is a free log retrieval operation binding the contract event 0x299eaafd0d27519eda3fe7195b73e5269e442b3d80928f19afa32b6db2f352b6.
//
// Solidity: e UniverseCreated(parentUniverse indexed address, childUniverse indexed address, payoutNumerators uint256[], invalid bool)
func (_Augur *AugurFilterer) FilterUniverseCreated(opts *bind.FilterOpts, parentUniverse []common.Address, childUniverse []common.Address) (*AugurUniverseCreatedIterator, error) {

	var parentUniverseRule []interface{}
	for _, parentUniverseItem := range parentUniverse {
		parentUniverseRule = append(parentUniverseRule, parentUniverseItem)
	}
	var childUniverseRule []interface{}
	for _, childUniverseItem := range childUniverse {
		childUniverseRule = append(childUniverseRule, childUniverseItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "UniverseCreated", parentUniverseRule, childUniverseRule)
	if err != nil {
		return nil, err
	}
	return &AugurUniverseCreatedIterator{contract: _Augur.contract, event: "UniverseCreated", logs: logs, sub: sub}, nil
}

// WatchUniverseCreated is a free log subscription operation binding the contract event 0x299eaafd0d27519eda3fe7195b73e5269e442b3d80928f19afa32b6db2f352b6.
//
// Solidity: e UniverseCreated(parentUniverse indexed address, childUniverse indexed address, payoutNumerators uint256[], invalid bool)
func (_Augur *AugurFilterer) WatchUniverseCreated(opts *bind.WatchOpts, sink chan<- *AugurUniverseCreated, parentUniverse []common.Address, childUniverse []common.Address) (event.Subscription, error) {

	var parentUniverseRule []interface{}
	for _, parentUniverseItem := range parentUniverse {
		parentUniverseRule = append(parentUniverseRule, parentUniverseItem)
	}
	var childUniverseRule []interface{}
	for _, childUniverseItem := range childUniverse {
		childUniverseRule = append(childUniverseRule, childUniverseItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "UniverseCreated", parentUniverseRule, childUniverseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurUniverseCreated)
				if err := _Augur.contract.UnpackLog(event, "UniverseCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AugurUniverseForkedIterator is returned from FilterUniverseForked and is used to iterate over the raw logs and unpacked data for UniverseForked events raised by the Augur contract.
type AugurUniverseForkedIterator struct {
	Event *AugurUniverseForked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurUniverseForkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurUniverseForked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurUniverseForked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurUniverseForkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurUniverseForkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurUniverseForked represents a UniverseForked event raised by the Augur contract.
type AugurUniverseForked struct {
	Universe common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUniverseForked is a free log retrieval operation binding the contract event 0xd4d990bbdf9b9a4383a394341465060ccb75513432ceee3d5fcd8788ab1a507f.
//
// Solidity: e UniverseForked(universe indexed address)
func (_Augur *AugurFilterer) FilterUniverseForked(opts *bind.FilterOpts, universe []common.Address) (*AugurUniverseForkedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "UniverseForked", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurUniverseForkedIterator{contract: _Augur.contract, event: "UniverseForked", logs: logs, sub: sub}, nil
}

// WatchUniverseForked is a free log subscription operation binding the contract event 0xd4d990bbdf9b9a4383a394341465060ccb75513432ceee3d5fcd8788ab1a507f.
//
// Solidity: e UniverseForked(universe indexed address)
func (_Augur *AugurFilterer) WatchUniverseForked(opts *bind.WatchOpts, sink chan<- *AugurUniverseForked, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "UniverseForked", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurUniverseForked)
				if err := _Augur.contract.UnpackLog(event, "UniverseForked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
