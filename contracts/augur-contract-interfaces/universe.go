package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)


// IUniverseABI is the input ABI used to generate the binding from.
const IUniverseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeThresholdForDisputePacing\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementOpenInterestFromMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCreateNextFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"decrementOpenInterestFromMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRepMarketCapInAttoeth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyTarget\",\"type\":\"address\"}],\"name\":\"isContainerForFeeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeRoundDurationInSeconds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fork\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetRepMarketCapInAttoeth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyTarget\",\"type\":\"address\"}],\"name\":\"isContainerForShareToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_badMarkets\",\"type\":\"uint256\"},{\"name\":\"_totalMarkets\",\"type\":\"uint256\"},{\"name\":\"_targetDivisor\",\"type\":\"uint256\"},{\"name\":\"_previousValue\",\"type\":\"uint256\"},{\"name\":\"_defaultValue\",\"type\":\"uint256\"},{\"name\":\"_floor\",\"type\":\"uint256\"}],\"name\":\"calculateFloatingValue\",\"outputs\":[{\"name\":\"_newValue\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialReportMinValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCreateCurrentFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOpenInterestInAttoEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWinningChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeMarketFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkEndTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkReputationGoal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheReportingFeeDivisor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyChild\",\"type\":\"address\"}],\"name\":\"isParentOf\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyTarget\",\"type\":\"address\"}],\"name\":\"isContainerForMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParentUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheValidityBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"decrementOpenInterest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getOrCreateFeeWindowByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReputationToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentUniverse\",\"type\":\"address\"},{\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isForking\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParentPayoutDistributionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyTarget\",\"type\":\"address\"}],\"name\":\"isContainerForFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkingMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementOpenInterest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentPayoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_invalid\",\"type\":\"bool\"}],\"name\":\"createChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_feeWindow\",\"type\":\"address\"}],\"name\":\"getOrCreateFeeWindowBefore\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheDesignatedReportStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheMarketCreationCost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"getChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"updateTentativeWinningChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reportingParticipant\",\"type\":\"address\"}],\"name\":\"isContainerForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeThresholdForFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheDesignatedReportNoShowBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addMarketTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IUniverseBin is the compiled bytecode used for deploying new contracts.
const IUniverseBin = `0x`

// DeployIUniverse deploys a new Ethereum contract, binding an instance of IUniverse to it.
func DeployIUniverse(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IUniverse, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniverseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IUniverseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IUniverse{IUniverseCaller: IUniverseCaller{contract: contract}, IUniverseTransactor: IUniverseTransactor{contract: contract}, IUniverseFilterer: IUniverseFilterer{contract: contract}}, nil
}

// IUniverse is an auto generated Go binding around an Ethereum contract.
type IUniverse struct {
	IUniverseCaller     // Read-only binding to the contract
	IUniverseTransactor // Write-only binding to the contract
	IUniverseFilterer   // Log filterer for contract events
}

// IUniverseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniverseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniverseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniverseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniverseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniverseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniverseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniverseSession struct {
	Contract     *IUniverse        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUniverseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniverseCallerSession struct {
	Contract *IUniverseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IUniverseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniverseTransactorSession struct {
	Contract     *IUniverseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IUniverseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniverseRaw struct {
	Contract *IUniverse // Generic contract binding to access the raw methods on
}

// IUniverseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniverseCallerRaw struct {
	Contract *IUniverseCaller // Generic read-only contract binding to access the raw methods on
}

// IUniverseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniverseTransactorRaw struct {
	Contract *IUniverseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniverse creates a new instance of IUniverse, bound to a specific deployed contract.
func NewIUniverse(address common.Address, backend bind.ContractBackend) (*IUniverse, error) {
	contract, err := bindIUniverse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniverse{IUniverseCaller: IUniverseCaller{contract: contract}, IUniverseTransactor: IUniverseTransactor{contract: contract}, IUniverseFilterer: IUniverseFilterer{contract: contract}}, nil
}

// NewIUniverseCaller creates a new read-only instance of IUniverse, bound to a specific deployed contract.
func NewIUniverseCaller(address common.Address, caller bind.ContractCaller) (*IUniverseCaller, error) {
	contract, err := bindIUniverse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniverseCaller{contract: contract}, nil
}

// NewIUniverseTransactor creates a new write-only instance of IUniverse, bound to a specific deployed contract.
func NewIUniverseTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniverseTransactor, error) {
	contract, err := bindIUniverse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniverseTransactor{contract: contract}, nil
}

// NewIUniverseFilterer creates a new log filterer instance of IUniverse, bound to a specific deployed contract.
func NewIUniverseFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniverseFilterer, error) {
	contract, err := bindIUniverse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniverseFilterer{contract: contract}, nil
}

// bindIUniverse binds a generic wrapper to an already deployed contract.
func bindIUniverse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniverseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniverse *IUniverseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IUniverse.Contract.IUniverseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniverse *IUniverseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.Contract.IUniverseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniverse *IUniverseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniverse.Contract.IUniverseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniverse *IUniverseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IUniverse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniverse *IUniverseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniverse *IUniverseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniverse.Contract.contract.Transact(opts, method, params...)
}

// CalculateFloatingValue is a free data retrieval call binding the contract method 0x5bafecf5.
//
// Solidity: function calculateFloatingValue(_badMarkets uint256, _totalMarkets uint256, _targetDivisor uint256, _previousValue uint256, _defaultValue uint256, _floor uint256) constant returns(_newValue uint256)
func (_IUniverse *IUniverseCaller) CalculateFloatingValue(opts *bind.CallOpts, _badMarkets *big.Int, _totalMarkets *big.Int, _targetDivisor *big.Int, _previousValue *big.Int, _defaultValue *big.Int, _floor *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "calculateFloatingValue", _badMarkets, _totalMarkets, _targetDivisor, _previousValue, _defaultValue, _floor)
	return *ret0, err
}

// CalculateFloatingValue is a free data retrieval call binding the contract method 0x5bafecf5.
//
// Solidity: function calculateFloatingValue(_badMarkets uint256, _totalMarkets uint256, _targetDivisor uint256, _previousValue uint256, _defaultValue uint256, _floor uint256) constant returns(_newValue uint256)
func (_IUniverse *IUniverseSession) CalculateFloatingValue(_badMarkets *big.Int, _totalMarkets *big.Int, _targetDivisor *big.Int, _previousValue *big.Int, _defaultValue *big.Int, _floor *big.Int) (*big.Int, error) {
	return _IUniverse.Contract.CalculateFloatingValue(&_IUniverse.CallOpts, _badMarkets, _totalMarkets, _targetDivisor, _previousValue, _defaultValue, _floor)
}

// CalculateFloatingValue is a free data retrieval call binding the contract method 0x5bafecf5.
//
// Solidity: function calculateFloatingValue(_badMarkets uint256, _totalMarkets uint256, _targetDivisor uint256, _previousValue uint256, _defaultValue uint256, _floor uint256) constant returns(_newValue uint256)
func (_IUniverse *IUniverseCallerSession) CalculateFloatingValue(_badMarkets *big.Int, _totalMarkets *big.Int, _targetDivisor *big.Int, _previousValue *big.Int, _defaultValue *big.Int, _floor *big.Int) (*big.Int, error) {
	return _IUniverse.Contract.CalculateFloatingValue(&_IUniverse.CallOpts, _badMarkets, _totalMarkets, _targetDivisor, _previousValue, _defaultValue, _floor)
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(_parentPayoutDistributionHash bytes32) constant returns(address)
func (_IUniverse *IUniverseCaller) GetChildUniverse(opts *bind.CallOpts, _parentPayoutDistributionHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getChildUniverse", _parentPayoutDistributionHash)
	return *ret0, err
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(_parentPayoutDistributionHash bytes32) constant returns(address)
func (_IUniverse *IUniverseSession) GetChildUniverse(_parentPayoutDistributionHash [32]byte) (common.Address, error) {
	return _IUniverse.Contract.GetChildUniverse(&_IUniverse.CallOpts, _parentPayoutDistributionHash)
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(_parentPayoutDistributionHash bytes32) constant returns(address)
func (_IUniverse *IUniverseCallerSession) GetChildUniverse(_parentPayoutDistributionHash [32]byte) (common.Address, error) {
	return _IUniverse.Contract.GetChildUniverse(&_IUniverse.CallOpts, _parentPayoutDistributionHash)
}

// GetCurrentFeeWindow is a free data retrieval call binding the contract method 0x79f96600.
//
// Solidity: function getCurrentFeeWindow() constant returns(address)
func (_IUniverse *IUniverseCaller) GetCurrentFeeWindow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getCurrentFeeWindow")
	return *ret0, err
}

// GetCurrentFeeWindow is a free data retrieval call binding the contract method 0x79f96600.
//
// Solidity: function getCurrentFeeWindow() constant returns(address)
func (_IUniverse *IUniverseSession) GetCurrentFeeWindow() (common.Address, error) {
	return _IUniverse.Contract.GetCurrentFeeWindow(&_IUniverse.CallOpts)
}

// GetCurrentFeeWindow is a free data retrieval call binding the contract method 0x79f96600.
//
// Solidity: function getCurrentFeeWindow() constant returns(address)
func (_IUniverse *IUniverseCallerSession) GetCurrentFeeWindow() (common.Address, error) {
	return _IUniverse.Contract.GetCurrentFeeWindow(&_IUniverse.CallOpts)
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x331172f3.
//
// Solidity: function getDisputeRoundDurationInSeconds() constant returns(uint256)
func (_IUniverse *IUniverseCaller) GetDisputeRoundDurationInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getDisputeRoundDurationInSeconds")
	return *ret0, err
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x331172f3.
//
// Solidity: function getDisputeRoundDurationInSeconds() constant returns(uint256)
func (_IUniverse *IUniverseSession) GetDisputeRoundDurationInSeconds() (*big.Int, error) {
	return _IUniverse.Contract.GetDisputeRoundDurationInSeconds(&_IUniverse.CallOpts)
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x331172f3.
//
// Solidity: function getDisputeRoundDurationInSeconds() constant returns(uint256)
func (_IUniverse *IUniverseCallerSession) GetDisputeRoundDurationInSeconds() (*big.Int, error) {
	return _IUniverse.Contract.GetDisputeRoundDurationInSeconds(&_IUniverse.CallOpts)
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() constant returns(uint256)
func (_IUniverse *IUniverseCaller) GetDisputeThresholdForDisputePacing(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getDisputeThresholdForDisputePacing")
	return *ret0, err
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() constant returns(uint256)
func (_IUniverse *IUniverseSession) GetDisputeThresholdForDisputePacing() (*big.Int, error) {
	return _IUniverse.Contract.GetDisputeThresholdForDisputePacing(&_IUniverse.CallOpts)
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() constant returns(uint256)
func (_IUniverse *IUniverseCallerSession) GetDisputeThresholdForDisputePacing() (*big.Int, error) {
	return _IUniverse.Contract.GetDisputeThresholdForDisputePacing(&_IUniverse.CallOpts)
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() constant returns(uint256)
func (_IUniverse *IUniverseCaller) GetDisputeThresholdForFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getDisputeThresholdForFork")
	return *ret0, err
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() constant returns(uint256)
func (_IUniverse *IUniverseSession) GetDisputeThresholdForFork() (*big.Int, error) {
	return _IUniverse.Contract.GetDisputeThresholdForFork(&_IUniverse.CallOpts)
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() constant returns(uint256)
func (_IUniverse *IUniverseCallerSession) GetDisputeThresholdForFork() (*big.Int, error) {
	return _IUniverse.Contract.GetDisputeThresholdForFork(&_IUniverse.CallOpts)
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() constant returns(uint256)
func (_IUniverse *IUniverseCaller) GetForkEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getForkEndTime")
	return *ret0, err
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() constant returns(uint256)
func (_IUniverse *IUniverseSession) GetForkEndTime() (*big.Int, error) {
	return _IUniverse.Contract.GetForkEndTime(&_IUniverse.CallOpts)
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() constant returns(uint256)
func (_IUniverse *IUniverseCallerSession) GetForkEndTime() (*big.Int, error) {
	return _IUniverse.Contract.GetForkEndTime(&_IUniverse.CallOpts)
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() constant returns(uint256)
func (_IUniverse *IUniverseCaller) GetForkReputationGoal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getForkReputationGoal")
	return *ret0, err
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() constant returns(uint256)
func (_IUniverse *IUniverseSession) GetForkReputationGoal() (*big.Int, error) {
	return _IUniverse.Contract.GetForkReputationGoal(&_IUniverse.CallOpts)
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() constant returns(uint256)
func (_IUniverse *IUniverseCallerSession) GetForkReputationGoal() (*big.Int, error) {
	return _IUniverse.Contract.GetForkReputationGoal(&_IUniverse.CallOpts)
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() constant returns(address)
func (_IUniverse *IUniverseCaller) GetForkingMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getForkingMarket")
	return *ret0, err
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() constant returns(address)
func (_IUniverse *IUniverseSession) GetForkingMarket() (common.Address, error) {
	return _IUniverse.Contract.GetForkingMarket(&_IUniverse.CallOpts)
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() constant returns(address)
func (_IUniverse *IUniverseCallerSession) GetForkingMarket() (common.Address, error) {
	return _IUniverse.Contract.GetForkingMarket(&_IUniverse.CallOpts)
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() constant returns(uint256)
func (_IUniverse *IUniverseCaller) GetInitialReportMinValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getInitialReportMinValue")
	return *ret0, err
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() constant returns(uint256)
func (_IUniverse *IUniverseSession) GetInitialReportMinValue() (*big.Int, error) {
	return _IUniverse.Contract.GetInitialReportMinValue(&_IUniverse.CallOpts)
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() constant returns(uint256)
func (_IUniverse *IUniverseCallerSession) GetInitialReportMinValue() (*big.Int, error) {
	return _IUniverse.Contract.GetInitialReportMinValue(&_IUniverse.CallOpts)
}

// GetOpenInterestInAttoEth is a free data retrieval call binding the contract method 0x6a9d7629.
//
// Solidity: function getOpenInterestInAttoEth() constant returns(uint256)
func (_IUniverse *IUniverseCaller) GetOpenInterestInAttoEth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getOpenInterestInAttoEth")
	return *ret0, err
}

// GetOpenInterestInAttoEth is a free data retrieval call binding the contract method 0x6a9d7629.
//
// Solidity: function getOpenInterestInAttoEth() constant returns(uint256)
func (_IUniverse *IUniverseSession) GetOpenInterestInAttoEth() (*big.Int, error) {
	return _IUniverse.Contract.GetOpenInterestInAttoEth(&_IUniverse.CallOpts)
}

// GetOpenInterestInAttoEth is a free data retrieval call binding the contract method 0x6a9d7629.
//
// Solidity: function getOpenInterestInAttoEth() constant returns(uint256)
func (_IUniverse *IUniverseCallerSession) GetOpenInterestInAttoEth() (*big.Int, error) {
	return _IUniverse.Contract.GetOpenInterestInAttoEth(&_IUniverse.CallOpts)
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() constant returns(bytes32)
func (_IUniverse *IUniverseCaller) GetParentPayoutDistributionHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getParentPayoutDistributionHash")
	return *ret0, err
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() constant returns(bytes32)
func (_IUniverse *IUniverseSession) GetParentPayoutDistributionHash() ([32]byte, error) {
	return _IUniverse.Contract.GetParentPayoutDistributionHash(&_IUniverse.CallOpts)
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() constant returns(bytes32)
func (_IUniverse *IUniverseCallerSession) GetParentPayoutDistributionHash() ([32]byte, error) {
	return _IUniverse.Contract.GetParentPayoutDistributionHash(&_IUniverse.CallOpts)
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() constant returns(address)
func (_IUniverse *IUniverseCaller) GetParentUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getParentUniverse")
	return *ret0, err
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() constant returns(address)
func (_IUniverse *IUniverseSession) GetParentUniverse() (common.Address, error) {
	return _IUniverse.Contract.GetParentUniverse(&_IUniverse.CallOpts)
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() constant returns(address)
func (_IUniverse *IUniverseCallerSession) GetParentUniverse() (common.Address, error) {
	return _IUniverse.Contract.GetParentUniverse(&_IUniverse.CallOpts)
}

// GetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x13ad9ce1.
//
// Solidity: function getRepMarketCapInAttoeth() constant returns(uint256)
func (_IUniverse *IUniverseCaller) GetRepMarketCapInAttoeth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getRepMarketCapInAttoeth")
	return *ret0, err
}

// GetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x13ad9ce1.
//
// Solidity: function getRepMarketCapInAttoeth() constant returns(uint256)
func (_IUniverse *IUniverseSession) GetRepMarketCapInAttoeth() (*big.Int, error) {
	return _IUniverse.Contract.GetRepMarketCapInAttoeth(&_IUniverse.CallOpts)
}

// GetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x13ad9ce1.
//
// Solidity: function getRepMarketCapInAttoeth() constant returns(uint256)
func (_IUniverse *IUniverseCallerSession) GetRepMarketCapInAttoeth() (*big.Int, error) {
	return _IUniverse.Contract.GetRepMarketCapInAttoeth(&_IUniverse.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_IUniverse *IUniverseCaller) GetReputationToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getReputationToken")
	return *ret0, err
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_IUniverse *IUniverseSession) GetReputationToken() (common.Address, error) {
	return _IUniverse.Contract.GetReputationToken(&_IUniverse.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_IUniverse *IUniverseCallerSession) GetReputationToken() (common.Address, error) {
	return _IUniverse.Contract.GetReputationToken(&_IUniverse.CallOpts)
}

// GetTargetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x4a719a27.
//
// Solidity: function getTargetRepMarketCapInAttoeth() constant returns(uint256)
func (_IUniverse *IUniverseCaller) GetTargetRepMarketCapInAttoeth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getTargetRepMarketCapInAttoeth")
	return *ret0, err
}

// GetTargetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x4a719a27.
//
// Solidity: function getTargetRepMarketCapInAttoeth() constant returns(uint256)
func (_IUniverse *IUniverseSession) GetTargetRepMarketCapInAttoeth() (*big.Int, error) {
	return _IUniverse.Contract.GetTargetRepMarketCapInAttoeth(&_IUniverse.CallOpts)
}

// GetTargetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x4a719a27.
//
// Solidity: function getTargetRepMarketCapInAttoeth() constant returns(uint256)
func (_IUniverse *IUniverseCallerSession) GetTargetRepMarketCapInAttoeth() (*big.Int, error) {
	return _IUniverse.Contract.GetTargetRepMarketCapInAttoeth(&_IUniverse.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IUniverse *IUniverseCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IUniverse *IUniverseSession) GetTypeName() ([32]byte, error) {
	return _IUniverse.Contract.GetTypeName(&_IUniverse.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IUniverse *IUniverseCallerSession) GetTypeName() ([32]byte, error) {
	return _IUniverse.Contract.GetTypeName(&_IUniverse.CallOpts)
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() constant returns(address)
func (_IUniverse *IUniverseCaller) GetWinningChildUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "getWinningChildUniverse")
	return *ret0, err
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() constant returns(address)
func (_IUniverse *IUniverseSession) GetWinningChildUniverse() (common.Address, error) {
	return _IUniverse.Contract.GetWinningChildUniverse(&_IUniverse.CallOpts)
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() constant returns(address)
func (_IUniverse *IUniverseCallerSession) GetWinningChildUniverse() (common.Address, error) {
	return _IUniverse.Contract.GetWinningChildUniverse(&_IUniverse.CallOpts)
}

// IsContainerForFeeToken is a free data retrieval call binding the contract method 0x26d16bc9.
//
// Solidity: function isContainerForFeeToken(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseCaller) IsContainerForFeeToken(opts *bind.CallOpts, _shadyTarget common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "isContainerForFeeToken", _shadyTarget)
	return *ret0, err
}

// IsContainerForFeeToken is a free data retrieval call binding the contract method 0x26d16bc9.
//
// Solidity: function isContainerForFeeToken(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseSession) IsContainerForFeeToken(_shadyTarget common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForFeeToken(&_IUniverse.CallOpts, _shadyTarget)
}

// IsContainerForFeeToken is a free data retrieval call binding the contract method 0x26d16bc9.
//
// Solidity: function isContainerForFeeToken(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseCallerSession) IsContainerForFeeToken(_shadyTarget common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForFeeToken(&_IUniverse.CallOpts, _shadyTarget)
}

// IsContainerForFeeWindow is a free data retrieval call binding the contract method 0xc7c88d70.
//
// Solidity: function isContainerForFeeWindow(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseCaller) IsContainerForFeeWindow(opts *bind.CallOpts, _shadyTarget common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "isContainerForFeeWindow", _shadyTarget)
	return *ret0, err
}

// IsContainerForFeeWindow is a free data retrieval call binding the contract method 0xc7c88d70.
//
// Solidity: function isContainerForFeeWindow(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseSession) IsContainerForFeeWindow(_shadyTarget common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForFeeWindow(&_IUniverse.CallOpts, _shadyTarget)
}

// IsContainerForFeeWindow is a free data retrieval call binding the contract method 0xc7c88d70.
//
// Solidity: function isContainerForFeeWindow(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseCallerSession) IsContainerForFeeWindow(_shadyTarget common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForFeeWindow(&_IUniverse.CallOpts, _shadyTarget)
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseCaller) IsContainerForMarket(opts *bind.CallOpts, _shadyTarget common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "isContainerForMarket", _shadyTarget)
	return *ret0, err
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseSession) IsContainerForMarket(_shadyTarget common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForMarket(&_IUniverse.CallOpts, _shadyTarget)
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseCallerSession) IsContainerForMarket(_shadyTarget common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForMarket(&_IUniverse.CallOpts, _shadyTarget)
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(_reportingParticipant address) constant returns(bool)
func (_IUniverse *IUniverseCaller) IsContainerForReportingParticipant(opts *bind.CallOpts, _reportingParticipant common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "isContainerForReportingParticipant", _reportingParticipant)
	return *ret0, err
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(_reportingParticipant address) constant returns(bool)
func (_IUniverse *IUniverseSession) IsContainerForReportingParticipant(_reportingParticipant common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForReportingParticipant(&_IUniverse.CallOpts, _reportingParticipant)
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(_reportingParticipant address) constant returns(bool)
func (_IUniverse *IUniverseCallerSession) IsContainerForReportingParticipant(_reportingParticipant common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForReportingParticipant(&_IUniverse.CallOpts, _reportingParticipant)
}

// IsContainerForShareToken is a free data retrieval call binding the contract method 0x509a1061.
//
// Solidity: function isContainerForShareToken(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseCaller) IsContainerForShareToken(opts *bind.CallOpts, _shadyTarget common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "isContainerForShareToken", _shadyTarget)
	return *ret0, err
}

// IsContainerForShareToken is a free data retrieval call binding the contract method 0x509a1061.
//
// Solidity: function isContainerForShareToken(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseSession) IsContainerForShareToken(_shadyTarget common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForShareToken(&_IUniverse.CallOpts, _shadyTarget)
}

// IsContainerForShareToken is a free data retrieval call binding the contract method 0x509a1061.
//
// Solidity: function isContainerForShareToken(_shadyTarget address) constant returns(bool)
func (_IUniverse *IUniverseCallerSession) IsContainerForShareToken(_shadyTarget common.Address) (bool, error) {
	return _IUniverse.Contract.IsContainerForShareToken(&_IUniverse.CallOpts, _shadyTarget)
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() constant returns(bool)
func (_IUniverse *IUniverseCaller) IsForking(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "isForking")
	return *ret0, err
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() constant returns(bool)
func (_IUniverse *IUniverseSession) IsForking() (bool, error) {
	return _IUniverse.Contract.IsForking(&_IUniverse.CallOpts)
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() constant returns(bool)
func (_IUniverse *IUniverseCallerSession) IsForking() (bool, error) {
	return _IUniverse.Contract.IsForking(&_IUniverse.CallOpts)
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(_shadyChild address) constant returns(bool)
func (_IUniverse *IUniverseCaller) IsParentOf(opts *bind.CallOpts, _shadyChild common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUniverse.contract.Call(opts, out, "isParentOf", _shadyChild)
	return *ret0, err
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(_shadyChild address) constant returns(bool)
func (_IUniverse *IUniverseSession) IsParentOf(_shadyChild common.Address) (bool, error) {
	return _IUniverse.Contract.IsParentOf(&_IUniverse.CallOpts, _shadyChild)
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(_shadyChild address) constant returns(bool)
func (_IUniverse *IUniverseCallerSession) IsParentOf(_shadyChild common.Address) (bool, error) {
	return _IUniverse.Contract.IsParentOf(&_IUniverse.CallOpts, _shadyChild)
}

// AddMarketTo is a paid mutator transaction binding the contract method 0xfeeda367.
//
// Solidity: function addMarketTo() returns(bool)
func (_IUniverse *IUniverseTransactor) AddMarketTo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "addMarketTo")
}

// AddMarketTo is a paid mutator transaction binding the contract method 0xfeeda367.
//
// Solidity: function addMarketTo() returns(bool)
func (_IUniverse *IUniverseSession) AddMarketTo() (*types.Transaction, error) {
	return _IUniverse.Contract.AddMarketTo(&_IUniverse.TransactOpts)
}

// AddMarketTo is a paid mutator transaction binding the contract method 0xfeeda367.
//
// Solidity: function addMarketTo() returns(bool)
func (_IUniverse *IUniverseTransactorSession) AddMarketTo() (*types.Transaction, error) {
	return _IUniverse.Contract.AddMarketTo(&_IUniverse.TransactOpts)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0xdf428e3b.
//
// Solidity: function createChildUniverse(_parentPayoutNumerators uint256[], _invalid bool) returns(address)
func (_IUniverse *IUniverseTransactor) CreateChildUniverse(opts *bind.TransactOpts, _parentPayoutNumerators []*big.Int, _invalid bool) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "createChildUniverse", _parentPayoutNumerators, _invalid)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0xdf428e3b.
//
// Solidity: function createChildUniverse(_parentPayoutNumerators uint256[], _invalid bool) returns(address)
func (_IUniverse *IUniverseSession) CreateChildUniverse(_parentPayoutNumerators []*big.Int, _invalid bool) (*types.Transaction, error) {
	return _IUniverse.Contract.CreateChildUniverse(&_IUniverse.TransactOpts, _parentPayoutNumerators, _invalid)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0xdf428e3b.
//
// Solidity: function createChildUniverse(_parentPayoutNumerators uint256[], _invalid bool) returns(address)
func (_IUniverse *IUniverseTransactorSession) CreateChildUniverse(_parentPayoutNumerators []*big.Int, _invalid bool) (*types.Transaction, error) {
	return _IUniverse.Contract.CreateChildUniverse(&_IUniverse.TransactOpts, _parentPayoutNumerators, _invalid)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(_amount uint256) returns(bool)
func (_IUniverse *IUniverseTransactor) DecrementOpenInterest(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "decrementOpenInterest", _amount)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(_amount uint256) returns(bool)
func (_IUniverse *IUniverseSession) DecrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.DecrementOpenInterest(&_IUniverse.TransactOpts, _amount)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(_amount uint256) returns(bool)
func (_IUniverse *IUniverseTransactorSession) DecrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.DecrementOpenInterest(&_IUniverse.TransactOpts, _amount)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x0db3be6a.
//
// Solidity: function decrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_IUniverse *IUniverseTransactor) DecrementOpenInterestFromMarket(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "decrementOpenInterestFromMarket", _amount)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x0db3be6a.
//
// Solidity: function decrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_IUniverse *IUniverseSession) DecrementOpenInterestFromMarket(_amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.DecrementOpenInterestFromMarket(&_IUniverse.TransactOpts, _amount)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x0db3be6a.
//
// Solidity: function decrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_IUniverse *IUniverseTransactorSession) DecrementOpenInterestFromMarket(_amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.DecrementOpenInterestFromMarket(&_IUniverse.TransactOpts, _amount)
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_IUniverse *IUniverseTransactor) Fork(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "fork")
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_IUniverse *IUniverseSession) Fork() (*types.Transaction, error) {
	return _IUniverse.Contract.Fork(&_IUniverse.TransactOpts)
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_IUniverse *IUniverseTransactorSession) Fork() (*types.Transaction, error) {
	return _IUniverse.Contract.Fork(&_IUniverse.TransactOpts)
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_IUniverse *IUniverseTransactor) GetOrCacheDesignatedReportNoShowBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "getOrCacheDesignatedReportNoShowBond")
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_IUniverse *IUniverseSession) GetOrCacheDesignatedReportNoShowBond() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheDesignatedReportNoShowBond(&_IUniverse.TransactOpts)
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_IUniverse *IUniverseTransactorSession) GetOrCacheDesignatedReportNoShowBond() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheDesignatedReportNoShowBond(&_IUniverse.TransactOpts)
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_IUniverse *IUniverseTransactor) GetOrCacheDesignatedReportStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "getOrCacheDesignatedReportStake")
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_IUniverse *IUniverseSession) GetOrCacheDesignatedReportStake() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheDesignatedReportStake(&_IUniverse.TransactOpts)
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_IUniverse *IUniverseTransactorSession) GetOrCacheDesignatedReportStake() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheDesignatedReportStake(&_IUniverse.TransactOpts)
}

// GetOrCacheMarketCreationCost is a paid mutator transaction binding the contract method 0xec86fdbd.
//
// Solidity: function getOrCacheMarketCreationCost() returns(uint256)
func (_IUniverse *IUniverseTransactor) GetOrCacheMarketCreationCost(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "getOrCacheMarketCreationCost")
}

// GetOrCacheMarketCreationCost is a paid mutator transaction binding the contract method 0xec86fdbd.
//
// Solidity: function getOrCacheMarketCreationCost() returns(uint256)
func (_IUniverse *IUniverseSession) GetOrCacheMarketCreationCost() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheMarketCreationCost(&_IUniverse.TransactOpts)
}

// GetOrCacheMarketCreationCost is a paid mutator transaction binding the contract method 0xec86fdbd.
//
// Solidity: function getOrCacheMarketCreationCost() returns(uint256)
func (_IUniverse *IUniverseTransactorSession) GetOrCacheMarketCreationCost() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheMarketCreationCost(&_IUniverse.TransactOpts)
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_IUniverse *IUniverseTransactor) GetOrCacheReportingFeeDivisor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "getOrCacheReportingFeeDivisor")
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_IUniverse *IUniverseSession) GetOrCacheReportingFeeDivisor() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheReportingFeeDivisor(&_IUniverse.TransactOpts)
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_IUniverse *IUniverseTransactorSession) GetOrCacheReportingFeeDivisor() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheReportingFeeDivisor(&_IUniverse.TransactOpts)
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_IUniverse *IUniverseTransactor) GetOrCacheValidityBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "getOrCacheValidityBond")
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_IUniverse *IUniverseSession) GetOrCacheValidityBond() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheValidityBond(&_IUniverse.TransactOpts)
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_IUniverse *IUniverseTransactorSession) GetOrCacheValidityBond() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCacheValidityBond(&_IUniverse.TransactOpts)
}

// GetOrCreateCurrentFeeWindow is a paid mutator transaction binding the contract method 0x68a7effc.
//
// Solidity: function getOrCreateCurrentFeeWindow() returns(address)
func (_IUniverse *IUniverseTransactor) GetOrCreateCurrentFeeWindow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "getOrCreateCurrentFeeWindow")
}

// GetOrCreateCurrentFeeWindow is a paid mutator transaction binding the contract method 0x68a7effc.
//
// Solidity: function getOrCreateCurrentFeeWindow() returns(address)
func (_IUniverse *IUniverseSession) GetOrCreateCurrentFeeWindow() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCreateCurrentFeeWindow(&_IUniverse.TransactOpts)
}

// GetOrCreateCurrentFeeWindow is a paid mutator transaction binding the contract method 0x68a7effc.
//
// Solidity: function getOrCreateCurrentFeeWindow() returns(address)
func (_IUniverse *IUniverseTransactorSession) GetOrCreateCurrentFeeWindow() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCreateCurrentFeeWindow(&_IUniverse.TransactOpts)
}

// GetOrCreateFeeWindowBefore is a paid mutator transaction binding the contract method 0xe487c7a5.
//
// Solidity: function getOrCreateFeeWindowBefore(_feeWindow address) returns(address)
func (_IUniverse *IUniverseTransactor) GetOrCreateFeeWindowBefore(opts *bind.TransactOpts, _feeWindow common.Address) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "getOrCreateFeeWindowBefore", _feeWindow)
}

// GetOrCreateFeeWindowBefore is a paid mutator transaction binding the contract method 0xe487c7a5.
//
// Solidity: function getOrCreateFeeWindowBefore(_feeWindow address) returns(address)
func (_IUniverse *IUniverseSession) GetOrCreateFeeWindowBefore(_feeWindow common.Address) (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCreateFeeWindowBefore(&_IUniverse.TransactOpts, _feeWindow)
}

// GetOrCreateFeeWindowBefore is a paid mutator transaction binding the contract method 0xe487c7a5.
//
// Solidity: function getOrCreateFeeWindowBefore(_feeWindow address) returns(address)
func (_IUniverse *IUniverseTransactorSession) GetOrCreateFeeWindowBefore(_feeWindow common.Address) (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCreateFeeWindowBefore(&_IUniverse.TransactOpts, _feeWindow)
}

// GetOrCreateFeeWindowByTimestamp is a paid mutator transaction binding the contract method 0xb66afea3.
//
// Solidity: function getOrCreateFeeWindowByTimestamp(_timestamp uint256) returns(address)
func (_IUniverse *IUniverseTransactor) GetOrCreateFeeWindowByTimestamp(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "getOrCreateFeeWindowByTimestamp", _timestamp)
}

// GetOrCreateFeeWindowByTimestamp is a paid mutator transaction binding the contract method 0xb66afea3.
//
// Solidity: function getOrCreateFeeWindowByTimestamp(_timestamp uint256) returns(address)
func (_IUniverse *IUniverseSession) GetOrCreateFeeWindowByTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCreateFeeWindowByTimestamp(&_IUniverse.TransactOpts, _timestamp)
}

// GetOrCreateFeeWindowByTimestamp is a paid mutator transaction binding the contract method 0xb66afea3.
//
// Solidity: function getOrCreateFeeWindowByTimestamp(_timestamp uint256) returns(address)
func (_IUniverse *IUniverseTransactorSession) GetOrCreateFeeWindowByTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCreateFeeWindowByTimestamp(&_IUniverse.TransactOpts, _timestamp)
}

// GetOrCreateNextFeeWindow is a paid mutator transaction binding the contract method 0x0cc8c9af.
//
// Solidity: function getOrCreateNextFeeWindow() returns(address)
func (_IUniverse *IUniverseTransactor) GetOrCreateNextFeeWindow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "getOrCreateNextFeeWindow")
}

// GetOrCreateNextFeeWindow is a paid mutator transaction binding the contract method 0x0cc8c9af.
//
// Solidity: function getOrCreateNextFeeWindow() returns(address)
func (_IUniverse *IUniverseSession) GetOrCreateNextFeeWindow() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCreateNextFeeWindow(&_IUniverse.TransactOpts)
}

// GetOrCreateNextFeeWindow is a paid mutator transaction binding the contract method 0x0cc8c9af.
//
// Solidity: function getOrCreateNextFeeWindow() returns(address)
func (_IUniverse *IUniverseTransactorSession) GetOrCreateNextFeeWindow() (*types.Transaction, error) {
	return _IUniverse.Contract.GetOrCreateNextFeeWindow(&_IUniverse.TransactOpts)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(_amount uint256) returns(bool)
func (_IUniverse *IUniverseTransactor) IncrementOpenInterest(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "incrementOpenInterest", _amount)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(_amount uint256) returns(bool)
func (_IUniverse *IUniverseSession) IncrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.IncrementOpenInterest(&_IUniverse.TransactOpts, _amount)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(_amount uint256) returns(bool)
func (_IUniverse *IUniverseTransactorSession) IncrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.IncrementOpenInterest(&_IUniverse.TransactOpts, _amount)
}

// IncrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x061777ed.
//
// Solidity: function incrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_IUniverse *IUniverseTransactor) IncrementOpenInterestFromMarket(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "incrementOpenInterestFromMarket", _amount)
}

// IncrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x061777ed.
//
// Solidity: function incrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_IUniverse *IUniverseSession) IncrementOpenInterestFromMarket(_amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.IncrementOpenInterestFromMarket(&_IUniverse.TransactOpts, _amount)
}

// IncrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x061777ed.
//
// Solidity: function incrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_IUniverse *IUniverseTransactorSession) IncrementOpenInterestFromMarket(_amount *big.Int) (*types.Transaction, error) {
	return _IUniverse.Contract.IncrementOpenInterestFromMarket(&_IUniverse.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(_parentUniverse address, _parentPayoutDistributionHash bytes32) returns(bool)
func (_IUniverse *IUniverseTransactor) Initialize(opts *bind.TransactOpts, _parentUniverse common.Address, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "initialize", _parentUniverse, _parentPayoutDistributionHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(_parentUniverse address, _parentPayoutDistributionHash bytes32) returns(bool)
func (_IUniverse *IUniverseSession) Initialize(_parentUniverse common.Address, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _IUniverse.Contract.Initialize(&_IUniverse.TransactOpts, _parentUniverse, _parentPayoutDistributionHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(_parentUniverse address, _parentPayoutDistributionHash bytes32) returns(bool)
func (_IUniverse *IUniverseTransactorSession) Initialize(_parentUniverse common.Address, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _IUniverse.Contract.Initialize(&_IUniverse.TransactOpts, _parentUniverse, _parentPayoutDistributionHash)
}

// RemoveMarketFrom is a paid mutator transaction binding the contract method 0x77a3a0a2.
//
// Solidity: function removeMarketFrom() returns(bool)
func (_IUniverse *IUniverseTransactor) RemoveMarketFrom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "removeMarketFrom")
}

// RemoveMarketFrom is a paid mutator transaction binding the contract method 0x77a3a0a2.
//
// Solidity: function removeMarketFrom() returns(bool)
func (_IUniverse *IUniverseSession) RemoveMarketFrom() (*types.Transaction, error) {
	return _IUniverse.Contract.RemoveMarketFrom(&_IUniverse.TransactOpts)
}

// RemoveMarketFrom is a paid mutator transaction binding the contract method 0x77a3a0a2.
//
// Solidity: function removeMarketFrom() returns(bool)
func (_IUniverse *IUniverseTransactorSession) RemoveMarketFrom() (*types.Transaction, error) {
	return _IUniverse.Contract.RemoveMarketFrom(&_IUniverse.TransactOpts)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(_parentPayoutDistributionHash bytes32) returns(bool)
func (_IUniverse *IUniverseTransactor) UpdateTentativeWinningChildUniverse(opts *bind.TransactOpts, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _IUniverse.contract.Transact(opts, "updateTentativeWinningChildUniverse", _parentPayoutDistributionHash)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(_parentPayoutDistributionHash bytes32) returns(bool)
func (_IUniverse *IUniverseSession) UpdateTentativeWinningChildUniverse(_parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _IUniverse.Contract.UpdateTentativeWinningChildUniverse(&_IUniverse.TransactOpts, _parentPayoutDistributionHash)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(_parentPayoutDistributionHash bytes32) returns(bool)
func (_IUniverse *IUniverseTransactorSession) UpdateTentativeWinningChildUniverse(_parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _IUniverse.Contract.UpdateTentativeWinningChildUniverse(&_IUniverse.TransactOpts, _parentPayoutDistributionHash)
}



// UniverseABI is the input ABI used to generate the binding from.
const UniverseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeThresholdForDisputePacing\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementOpenInterestFromMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCreateNextFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"decrementOpenInterestFromMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRepMarketCapInAttoeth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCreatePreviousFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyFeeToken\",\"type\":\"address\"}],\"name\":\"isContainerForFeeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeRoundDurationInSeconds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_feeWindowId\",\"type\":\"uint256\"}],\"name\":\"getFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fork\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetRepMarketCapInAttoeth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCreatePreviousPreviousFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyShareToken\",\"type\":\"address\"}],\"name\":\"isContainerForShareToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"buyParticipationTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_badMarkets\",\"type\":\"uint256\"},{\"name\":\"_totalMarkets\",\"type\":\"uint256\"},{\"name\":\"_targetDivisor\",\"type\":\"uint256\"},{\"name\":\"_previousValue\",\"type\":\"uint256\"},{\"name\":\"_defaultValue\",\"type\":\"uint256\"},{\"name\":\"_floor\",\"type\":\"uint256\"}],\"name\":\"calculateFloatingValue\",\"outputs\":[{\"name\":\"_newValue\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialReportMinValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCreateCurrentFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOpenInterestInAttoEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPreviousFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWinningChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_endTime\",\"type\":\"uint256\"},{\"name\":\"_feePerEthInWei\",\"type\":\"uint256\"},{\"name\":\"_denominationToken\",\"type\":\"address\"},{\"name\":\"_designatedReporterAddress\",\"type\":\"address\"},{\"name\":\"_minPrice\",\"type\":\"int256\"},{\"name\":\"_maxPrice\",\"type\":\"int256\"},{\"name\":\"_numTicks\",\"type\":\"uint256\"},{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_extraInfo\",\"type\":\"string\"}],\"name\":\"createScalarMarket\",\"outputs\":[{\"name\":\"_newMarket\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeMarketFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkEndTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkReputationGoal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reportingParticipants\",\"type\":\"address[]\"},{\"name\":\"_feeWindows\",\"type\":\"address[]\"}],\"name\":\"redeemStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheReportingFeeDivisor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyChild\",\"type\":\"address\"}],\"name\":\"isParentOf\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateForkValues\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyMarket\",\"type\":\"address\"}],\"name\":\"isContainerForMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParentUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getFeeWindowByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheValidityBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getInitialReportStakeSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"decrementOpenInterest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getOrCreateFeeWindowByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReputationToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentUniverse\",\"type\":\"address\"},{\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isForking\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParentPayoutDistributionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyFeeWindow\",\"type\":\"address\"}],\"name\":\"isContainerForFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_endTime\",\"type\":\"uint256\"},{\"name\":\"_feePerEthInWei\",\"type\":\"uint256\"},{\"name\":\"_denominationToken\",\"type\":\"address\"},{\"name\":\"_designatedReporterAddress\",\"type\":\"address\"},{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_extraInfo\",\"type\":\"string\"}],\"name\":\"createYesNoMarket\",\"outputs\":[{\"name\":\"_newMarket\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkingMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementOpenInterest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_endTime\",\"type\":\"uint256\"},{\"name\":\"_feePerEthInWei\",\"type\":\"uint256\"},{\"name\":\"_denominationToken\",\"type\":\"address\"},{\"name\":\"_designatedReporterAddress\",\"type\":\"address\"},{\"name\":\"_outcomes\",\"type\":\"bytes32[]\"},{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_extraInfo\",\"type\":\"string\"}],\"name\":\"createCategoricalMarket\",\"outputs\":[{\"name\":\"_newMarket\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getFeeWindowId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentPayoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_parentInvalid\",\"type\":\"bool\"}],\"name\":\"createChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_feeWindow\",\"type\":\"address\"}],\"name\":\"getOrCreateFeeWindowBefore\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheDesignatedReportStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheMarketCreationCost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"getChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"updateTentativeWinningChildUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shadyReportingParticipant\",\"type\":\"address\"}],\"name\":\"isContainerForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeThresholdForFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheDesignatedReportNoShowBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addMarketTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UniverseBin is the compiled bytecode used for deploying new contracts.
const UniverseBin = `0x60806040526002805460ff1916905560008054600160a060020a03191633179055613d678061002f6000396000f3006080604052600436106102c65763ffffffff60e060020a600035041663047825c781146102cb578063061777ed146102f25780630cc8c9af1461031e5780630db3be6a1461034f57806313ad9ce114610367578063225640ab1461037c57806326d16bc9146103915780633018205f146103b2578063331172f3146103c757806340b73897146103dc5780634591c060146103f45780634a719a27146104095780634f8035161461041e578063509a10611461043357806358ab010d146104545780635bafecf51461046c5780635f723b501461049357806368a7effc146104a85780636a9d7629146104bd5780636def1efc146104d25780636f70b9cb146104e757806372d61dfa146104fc57806377a3a0a2146105b557806377e71ee5146105ca57806379f96600146105df5780637c377d74146105f45780638d8cc634146106095780638f93bffe1461069757806392eefe9b146106ac5780639517317c146106cd5780639ab448d9146106ee5780639f7e1bf614610703578063a63f135014610724578063aa2bebb714610739578063af4cd45714610751578063af6b361514610766578063b62418a11461077b578063b66afea314610793578063b80907f2146107ab578063be13f47c146107c0578063becb1f35146107e4578063bef72fa2146107f9578063c38c0fa71461080e578063c7c88d7014610823578063cabc6f1414610844578063cb1d8418146108ef578063ce483e8814610904578063d28a16851461091c578063dae8af7714610a06578063db0a087c14610a1e578063df428e3b14610a33578063e487c7a514610a8c578063e79609e214610aad578063ec86fdbd14610ac2578063eceba87614610ad7578063edd1d02e14610aef578063ee89dab414610b04578063f7095d9d14610b19578063f76514c714610b31578063fb03eaea14610b52578063fd1e5e7a14610b67578063feeda36714610b7c575b600080fd5b3480156102d757600080fd5b506102e0610b91565b60408051918252519081900360200190f35b3480156102fe57600080fd5b5061030a600435610b97565b604080519115158252519081900360200190f35b34801561032a57600080fd5b50610333610bcb565b60408051600160a060020a039092168252519081900360200190f35b34801561035b57600080fd5b5061030a600435610c6f565b34801561037357600080fd5b506102e0610c98565b34801561038857600080fd5b50610333610e67565b34801561039d57600080fd5b5061030a600160a060020a0360043516610f01565b3480156103be57600080fd5b50610333611011565b3480156103d357600080fd5b506102e0611020565b3480156103e857600080fd5b5061033360043561102a565b34801561040057600080fd5b5061030a611045565b34801561041557600080fd5b506102e06111c9565b34801561042a57600080fd5b506103336111f5565b34801561043f57600080fd5b5061030a600160a060020a036004351661120a565b34801561046057600080fd5b5061030a600435611346565b34801561047857600080fd5b506102e060043560243560443560643560843560a4356113f2565b34801561049f57600080fd5b506102e06114e7565b3480156104b457600080fd5b506103336114ed565b3480156104c957600080fd5b506102e061157f565b3480156104de57600080fd5b50610333611585565b3480156104f357600080fd5b5061033361159a565b60408051602060046101043581810135601f8101849004840285018401909552848452610333948235946024803595604435600160a060020a0390811696606435909116956084359560a4359560c4359560e435953695946101249492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506117639650505050505050565b3480156105c157600080fd5b5061030a611a1c565b3480156105d657600080fd5b506102e0611a4f565b3480156105eb57600080fd5b50610333611a55565b34801561060057600080fd5b506102e0611ae7565b34801561061557600080fd5b506040805160206004803580820135838102808601850190965280855261030a95369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750611aed9650505050505050565b3480156106a357600080fd5b506102e0611c5f565b3480156106b857600080fd5b5061030a600160a060020a0360043516611d6d565b3480156106d957600080fd5b5061030a600160a060020a0360043516611db7565b3480156106fa57600080fd5b5061030a611e4a565b34801561070f57600080fd5b5061030a600160a060020a0360043516611f4f565b34801561073057600080fd5b50610333611f6d565b34801561074557600080fd5b50610333600435611f81565b34801561075d57600080fd5b506102e0611fac565b34801561077257600080fd5b506102e0612153565b34801561078757600080fd5b5061030a600435612174565b34801561079f57600080fd5b50610333600435612211565b3480156107b757600080fd5b506103336124f7565b3480156107cc57600080fd5b5061030a600160a060020a0360043516602435612506565b3480156107f057600080fd5b5061030a6126eb565b34801561080557600080fd5b506102e06126fc565b34801561081a57600080fd5b506102e0612702565b34801561082f57600080fd5b5061030a600160a060020a0360043516612708565b604080516020600460a43581810135601f8101849004840285018401909552848452610333948235946024803595604435600160a060020a039081169660643590911695608435953695929460c494920191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506127c29650505050505050565b3480156108fb57600080fd5b50610333612a2e565b34801561091057600080fd5b5061030a600435612a3d565b6040805160206004608435818101358381028086018501909652808552610333958335956024803596604435600160a060020a0390811697606435909116963696919560a49594909101928291908501908490808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b8a359b909a90999401975091955091820193509150819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750612ada9650505050505050565b348015610a1257600080fd5b506102e0600435612db6565b348015610a2a57600080fd5b506102e0612dd0565b348015610a3f57600080fd5b506040805160206004803580820135838102808601850190965280855261033395369593946024949385019291829185019084908082843750949750505050913515159250612df4915050565b348015610a9857600080fd5b50610333600160a060020a0360043516613088565b348015610ab957600080fd5b506102e06130d0565b348015610ace57600080fd5b506102e061326d565b348015610ae357600080fd5b50610333600435613277565b348015610afb57600080fd5b50610333613292565b348015610b1057600080fd5b5061030a6132a2565b348015610b2557600080fd5b5061030a6004356132ab565b348015610b3d57600080fd5b5061030a600160a060020a036004351661356c565b348015610b5e57600080fd5b506102e0613674565b348015610b7357600080fd5b506102e061367a565b348015610b8857600080fd5b5061030a61380b565b600a5490565b6000610ba233611f4f565b1515610bad57600080fd5b600f54610bc0908363ffffffff6139c416565b600f55506001919050565b6000610c6a610c65610bdb611020565b6000809054906101000a9004600160a060020a0316600160a060020a031663188ec3566040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610c2d57600080fd5b505af1158015610c41573d6000803e3d6000fd5b505050506040513d6020811015610c5757600080fd5b50519063ffffffff6139c416565b612211565b905090565b6000610c7a33611f4f565b1515610c8557600080fd5b600f54610bc0908363ffffffff6139d616565b6000805460408051600080516020613d1c83398151915281527f52657050726963654f7261636c650000000000000000000000000000000000006004820152905183928392600160a060020a039091169163f39ec1f79160248082019260209290919082900301818787803b158015610d1057600080fd5b505af1158015610d24573d6000803e3d6000fd5b505050506040513d6020811015610d3a57600080fd5b5051604080517f0f8fd3630000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921691630f8fd363916004808201926020929091908290030181600087803b158015610d9957600080fd5b505af1158015610dad573d6000803e3d6000fd5b505050506040513d6020811015610dc357600080fd5b50519150610e60670de0b6b3a7640000610e5484610ddf6124f7565b600160a060020a03166318160ddd6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610e1c57600080fd5b505af1158015610e30573d6000803e3d6000fd5b505050506040513d6020811015610e4657600080fd5b50519063ffffffff6139eb16565b9063ffffffff613a1216565b9392505050565b6000610c6a610c65610e77611020565b6000809054906101000a9004600160a060020a0316600160a060020a031663188ec3566040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610ec957600080fd5b505af1158015610edd573d6000803e3d6000fd5b505050506040513d6020811015610ef357600080fd5b50519063ffffffff6139d616565b600080600083600160a060020a031663f77f29b16040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610f4457600080fd5b505af1158015610f58573d6000803e3d6000fd5b505050506040513d6020811015610f6e57600080fd5b50519150610f7b82612708565b1515610f8657600080fd5b81905083600160a060020a031681600160a060020a031663ca709a256040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610fd157600080fd5b505af1158015610fe5573d6000803e3d6000fd5b505050506040513d6020811015610ffb57600080fd5b5051600160a060020a03161492505b5050919050565b600054600160a060020a031690565b6000610c6a613a29565b6000908152600c6020526040902054600160a060020a031690565b60025460009060ff16151561105957600080fd5b6110616126eb565b1561106b57600080fd5b61107433611f4f565b151561107f57600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916331790556110a9610bdb613a30565b60075560008054604080517f4e94c8290000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921692634e94c829926004808401936020939083900390910190829087803b15801561110d57600080fd5b505af1158015611121573d6000803e3d6000fd5b505050506040513d602081101561113757600080fd5b5051604080517f4a8120230000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921691634a812023916004808201926020929091908290030181600087803b15801561119657600080fd5b505af11580156111aa573d6000803e3d6000fd5b505050506040513d60208110156111c057600080fd5b50600191505090565b6000610c6a6111d6613a37565b610e546111e1613a3c565b6111e961157f565b9063ffffffff6139eb16565b6000610c6a610c65610e7760026111e9611020565b600080600083600160a060020a031663f1be16796040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561124d57600080fd5b505af1158015611261573d6000803e3d6000fd5b505050506040513d602081101561127757600080fd5b50519150600160a060020a0382161515611294576000925061100a565b61129d82611f4f565b15156112ac576000925061100a565b50604080517f509a1061000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301529151839283169163509a10619160248083019260209291908290030181600087803b15801561131257600080fd5b505af1158015611326573d6000803e3d6000fd5b505050506040513d602081101561133c57600080fd5b5051949350505050565b6000806113516114ed565b604080517f7303ed18000000000000000000000000000000000000000000000000000000008152336004820152602481018690529051919250600160a060020a03831691637303ed18916044808201926020929091908290030181600087803b1580156113bd57600080fd5b505af11580156113d1573d6000803e3d6000fd5b505050506040513d60208110156113e757600080fd5b506001949350505050565b60008515156114025750816114dd565b83151561140d578293505b848681151561141857fe5b04871161146557611433856111e9898763ffffffff6139eb16565b9050858181151561144057fe5b04905060028104905061145e6002858391900463ffffffff6139c416565b90506114ca565b6114a4611497868681151561147657fe5b0461148b89610e54898d63ffffffff6139eb16565b9063ffffffff6139d616565b869063ffffffff6139eb16565b905060018503818115156114b457fe5b0490506114c7818563ffffffff6139c416565b90505b6114da818363ffffffff613a4116565b90505b9695505050505050565b600b5490565b60008054604080517f188ec3560000000000000000000000000000000000000000000000000000000081529051610c6a92600160a060020a03169163188ec35691600480830192602092919082900301818887803b15801561154e57600080fd5b505af1158015611562573d6000803e3d6000fd5b505050506040513d602081101561157857600080fd5b5051612211565b600f5490565b6000610c6a611595610e77611020565b611f81565b60008060006115a76126eb565b15156115b257600080fd5b60065415156115c057600080fd5b6115cb600654613277565b915081600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561160b57600080fd5b505af115801561161f573d6000803e3d6000fd5b505050506040513d602081101561163557600080fd5b5051604080517f91d76bbb0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216916391d76bbb916004808201926020929091908290030181600087803b15801561169457600080fd5b505af11580156116a8573d6000803e3d6000fd5b505050506040513d60208110156116be57600080fd5b50516008549091508110158061175257506007546000809054906101000a9004600160a060020a0316600160a060020a031663188ec3566040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561172457600080fd5b505af1158015611738573d6000803e3d6000fd5b505050506040513d602081101561174e57600080fd5b5051115b151561175d57600080fd5b50919050565b60025460009060ff16151561177757600080fd5b825160001061178557600080fd5b85871261179157600080fd5b6117a285600263ffffffff613a5816565b15156117ad57600080fd5b6117bd8b8b8b8b3360028b613a6e565b90506000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561181157600080fd5b505af1158015611825573d6000803e3d6000fd5b505050506040513d602081101561183b57600080fd5b8101908080519060200190929190505050600160a060020a031663bc339f418585853086338e8e60026040518a63ffffffff1660e060020a028152600401808a60001916600019168152602001806020018060200189600160a060020a0316600160a060020a0316815260200188600160a060020a0316600160a060020a0316815260200187600160a060020a0316600160a060020a031681526020018681526020018581526020018460028111156118f057fe5b60ff16815260200183810383528b818151815260200191508051906020019080838360005b8381101561192d578181015183820152602001611915565b50505050905090810190601f16801561195a5780820380516001836020036101000a031916815260200191505b5083810382528a5181528a516020918201918c019080838360005b8381101561198d578181015183820152602001611975565b50505050905090810190601f1680156119ba5780820380516001836020036101000a031916815260200191505b509b505050505050505050505050602060405180830381600087803b1580156119e257600080fd5b505af11580156119f6573d6000803e3d6000fd5b505050506040513d6020811015611a0c57600080fd5b50509a9950505050505050505050565b6000611a2733611f4f565b1515611a3257600080fd5b50336000908152600d60205260409020805460ff19169055600190565b60075490565b60008054604080517f188ec3560000000000000000000000000000000000000000000000000000000081529051610c6a92600160a060020a03169163188ec35691600480830192602092919082900301818887803b158015611ab657600080fd5b505af1158015611aca573d6000803e3d6000fd5b505050506040513d6020811015611ae057600080fd5b5051611f81565b60085490565b600080805b8451821015611ba8578482815181101515611b0957fe5b6020908102909101810151604080517f95a2251f0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a03909216926395a2251f926024808401938290030181600087803b158015611b7157600080fd5b505af1158015611b85573d6000803e3d6000fd5b505050506040513d6020811015611b9b57600080fd5b5050600190910190611af2565b5060005b83518110156113e7578381815181101515611bc357fe5b6020908102909101810151604080517f95a2251f0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a03909216926395a2251f926024808401938290030181600087803b158015611c2b57600080fd5b505af1158015611c3f573d6000803e3d6000fd5b505050506040513d6020811015611c5557600080fd5b5050600101611bac565b6000806000806000806000611c726114ed565b9550611c7c610e67565b600160a060020a03871660009081526014602052604090205490955093508315611ca857839650611d64565b611cb0610c98565b9250611cba6111c9565b600160a060020a0386166000908152601460205260409020549092509050801515611cee57611ce7613cc9565b9350611d14565b811515611cfd57611ce7613cc9565b611d1182610e54838663ffffffff6139eb16565b93505b611d43611d1f613cce565b611d37611d2a613cd4565b879063ffffffff613a4116565b9063ffffffff613cd916565b600160a060020a038716600090815260146020526040902081905596508693505b50505050505090565b60008054600160a060020a03163314611d8557600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60008082600160a060020a031663c38c0fa76040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611df857600080fd5b505af1158015611e0c573d6000803e3d6000fd5b505050506040513d6020811015611e2257600080fd5b50519050600160a060020a038316611e3982613277565b600160a060020a0316149392505050565b60048054604080517f238d359000000000000000000000000000000000000000000000000000000000815290516000938493600160a060020a03169263238d3590928183019260209282900301818787803b158015611ea857600080fd5b505af1158015611ebc573d6000803e3d6000fd5b505050506040513d6020811015611ed257600080fd5b50519050611ee781600263ffffffff613a1216565b600855611efb81602863ffffffff613a1216565b6009819055611f2d90600190611f21906204000090610e5490600363ffffffff613a1216565b9063ffffffff6139c416565b600b55600954611f4490602063ffffffff613a1216565b600a55600191505090565b600160a060020a03166000908152600d602052604090205460ff1690565b6002546101009004600160a060020a031690565b600080611f8d83612db6565b6000908152600c6020526040902054600160a060020a03169392505050565b6000806000806000806000611fbf6114ed565b9550611fc96111f5565b600160a060020a03871660009081526010602052604090205490955093508315611ff557839650611d64565b84600160a060020a031663295c39a56040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561203357600080fd5b505af1158015612047573d6000803e3d6000fd5b505050506040513d602081101561205d57600080fd5b5051604080517fcf3d38490000000000000000000000000000000000000000000000000000000081529051919450600160a060020a0387169163cf3d3849916004808201926020929091908290030181600087803b1580156120be57600080fd5b505af11580156120d2573d6000803e3d6000fd5b505050506040513d60208110156120e857600080fd5b5051600160a060020a038616600090815260106020526040902054909250905061212b8284612115613cc9565b8461211e613ce9565b612126613ce9565b6113f2565b600160a060020a03969096166000908152601060205260409020869055509395945050505050565b6000610c6a6121606130d0565b61216861367a565b9063ffffffff613a4116565b60008054604080517f3f08882f0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b1580156121da57600080fd5b505af11580156121ee573d6000803e3d6000fd5b505050506040513d602081101561220457600080fd5b50511515610c8557600080fd5b600080600061221f84612db6565b6000818152600c6020526040902054909250600160a060020a031615156124d8576000805460408051600080516020613d1c83398151915281527f46656557696e646f77466163746f72790000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b1580156122b557600080fd5b505af11580156122c9573d6000803e3d6000fd5b505050506040513d60208110156122df57600080fd5b505160008054604080517f27237b1d000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015230602482015260448101879052905191909316926327237b1d9260648083019360209390929083900390910190829087803b15801561235a57600080fd5b505af115801561236e573d6000803e3d6000fd5b505050506040513d602081101561238457600080fd5b50516000838152600c60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0380871691909117909155835482517f4e94c82900000000000000000000000000000000000000000000000000000000815292519596501693634e94c82993600480840194938390030190829087803b15801561241357600080fd5b505af1158015612427573d6000803e3d6000fd5b505050506040513d602081101561243d57600080fd5b5051604080517f8d1b2afd000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301526024820186905291519190921691638d1b2afd9160448083019260209291908290030181600087803b1580156124ab57600080fd5b505af11580156124bf573d6000803e3d6000fd5b505050506040513d60208110156124d557600080fd5b50505b506000908152600c6020526040902054600160a060020a031692915050565b600454600160a060020a031690565b60025460009060ff161561251957600080fd5b612521613cf4565b5060028054600160a060020a038086166101000274ffffffffffffffffffffffffffffffffffffffff00199092169190911790915560038390556000805460408051600080516020613d1c83398151915281527f52657075746174696f6e546f6b656e466163746f72790000000000000000000060048201529051919093169263f39ec1f79260248083019360209390929083900390910190829087803b1580156125cb57600080fd5b505af11580156125df573d6000803e3d6000fd5b505050506040513d60208110156125f557600080fd5b505160008054604080517f812f82d9000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201523060248201529051919093169263812f82d99260448083019360209390929083900390910190829087803b15801561266957600080fd5b505af115801561267d573d6000803e3d6000fd5b505050506040513d602081101561269357600080fd5b50516004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790556126c9611e4a565b50600454600160a060020a031615156126e157600080fd5b5060015b92915050565b600554600160a060020a0316151590565b60015481565b60035490565b60008060008084600160a060020a031663c828371e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561274c57600080fd5b505af1158015612760573d6000803e3d6000fd5b505050506040513d602081101561277657600080fd5b5051925082151561278a57600093506127ba565b61279383612db6565b6000818152600c6020526040902054600160a060020a038781169116908114955090925090505b505050919050565b60025460009060ff1615156127d657600080fd5b82516000106127e457600080fd5b6127f688888888336002612710613a6e565b90506000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561284a57600080fd5b505af115801561285e573d6000803e3d6000fd5b505050506040513d602081101561287457600080fd5b50516040517fbc339f41000000000000000000000000000000000000000000000000000000008152600481018681523060648301819052600160a060020a0385811660848501523360a48501819052600060c48601819052670de0b6b3a764000060e487018190526101048701829052610120602488019081528b516101248901528b51949098169763bc339f41978d978d978d9790968d969095909490938593926044820191610144019060208d0190808383895b8381101561294257818101518382015260200161292a565b50505050905090810190601f16801561296f5780820380516001836020036101000a031916815260200191505b5083810382528a5181528a516020918201918c019080838360005b838110156129a257818101518382015260200161298a565b50505050905090810190601f1680156129cf5780820380516001836020036101000a031916815260200191505b509b505050505050505050505050602060405180830381600087803b1580156129f757600080fd5b505af1158015612a0b573d6000803e3d6000fd5b505050506040513d6020811015612a2157600080fd5b5050979650505050505050565b600554600160a060020a031690565b60008054604080517f3f08882f0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b158015612aa357600080fd5b505af1158015612ab7573d6000803e3d6000fd5b505050506040513d6020811015612acd57600080fd5b50511515610bad57600080fd5b60025460009060ff161515612aee57600080fd5b8251600010612afc57600080fd5b612b0e89898989338a51612710613a6e565b90506000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612b6257600080fd5b505af1158015612b76573d6000803e3d6000fd5b505050506040513d6020811015612b8c57600080fd5b8101908080519060200190929190505050600160a060020a03166307c1880a8585853086338c6000670de0b6b3a764000060016040518b63ffffffff1660e060020a028152600401808b6000191660001916815260200180602001806020018a600160a060020a0316600160a060020a0316815260200189600160a060020a0316600160a060020a0316815260200188600160a060020a0316600160a060020a0316815260200180602001878152602001868152602001856002811115612c4f57fe5b60ff16815260200184810384528d818151815260200191508051906020019080838360005b83811015612c8c578181015183820152602001612c74565b50505050905090810190601f168015612cb95780820380516001836020036101000a031916815260200191505b5084810383528c5181528c516020918201918e019080838360005b83811015612cec578181015183820152602001612cd4565b50505050905090810190601f168015612d195780820380516001836020036101000a031916815260200191505b508481038252885181528851602091820191808b01910280838360005b83811015612d4e578181015183820152602001612d36565b505050509050019d5050505050505050505050505050602060405180830381600087803b158015612d7e57600080fd5b505af1158015612d92573d6000803e3d6000fd5b505050506040513d6020811015612da857600080fd5b505098975050505050505050565b60006126e5612dc3611020565b839063ffffffff613a1216565b7f556e69766572736500000000000000000000000000000000000000000000000090565b600554604080517f2443f0ae000000000000000000000000000000000000000000000000000000008152831515602482015260048101918252845160448201528451600093849384938493600160a060020a0390931692632443f0ae928a928a92829160640190602080870191028083838c5b83811015612e7f578181015183820152602001612e67565b505050509050019350505050602060405180830381600087803b158015612ea557600080fd5b505af1158015612eb9573d6000803e3d6000fd5b505050506040513d6020811015612ecf57600080fd5b50519250612edc83613277565b91506000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612f3057600080fd5b505af1158015612f44573d6000803e3d6000fd5b505050506040513d6020811015612f5a57600080fd5b50519050600160a060020a038216151561307f576040517f8892bb73000000000000000000000000000000000000000000000000000000008152600481018481528615156044830152606060248301908152885160648401528851600160a060020a03851693638892bb739388938c938c93608401906020808701910280838360005b83811015612ff5578181015183820152602001612fdd565b50505050905001945050505050602060405180830381600087803b15801561301c57600080fd5b505af1158015613030573d6000803e3d6000fd5b505050506040513d602081101561304657600080fd5b50516000848152600e60205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03831617905591505b50949350505050565b60006126e5610c65600284600160a060020a031663c828371e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610ec957600080fd5b60008060008060008060006130e36114ed565b95506130ed6111f5565b600160a060020a0387166000908152601260205260409020549095509350831561311957839650611d64565b84600160a060020a031663295c39a56040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561315757600080fd5b505af115801561316b573d6000803e3d6000fd5b505050506040513d602081101561318157600080fd5b5051604080517fa52c05120000000000000000000000000000000000000000000000000000000081529051919450600160a060020a0387169163a52c0512916004808201926020929091908290030181600087803b1580156131e257600080fd5b505af11580156131f6573d6000803e3d6000fd5b505050506040513d602081101561320c57600080fd5b5051600160a060020a03861660009081526012602052604090205490925090506132458284613239613cc9565b84600b54600b546113f2565b600160a060020a03969096166000908152601260205260409020869055509395945050505050565b6000610c6a611fac565b6000908152600e6020526040902054600160a060020a031690565b6000610c6a611595610bdb611020565b60025460ff1690565b60008060008060006132be600654613277565b93506132c986613277565b925060009150600160a060020a038416156133d45783600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561331c57600080fd5b505af1158015613330573d6000803e3d6000fd5b505050506040513d602081101561334657600080fd5b5051604080517f91d76bbb0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216916391d76bbb916004808201926020929091908290030181600087803b1580156133a557600080fd5b505af11580156133b9573d6000803e3d6000fd5b505050506040513d60208110156133cf57600080fd5b505191505b82600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561341257600080fd5b505af1158015613426573d6000803e3d6000fd5b505050506040513d602081101561343c57600080fd5b5051604080517f91d76bbb0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216916391d76bbb916004808201926020929091908290030181600087803b15801561349b57600080fd5b505af11580156134af573d6000803e3d6000fd5b505050506040513d60208110156134c557600080fd5b50519050818111156134d75760068690555b600854811061356057600560009054906101000a9004600160a060020a0316600160a060020a031663c3fc47876040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561353357600080fd5b505af1158015613547573d6000803e3d6000fd5b505050506040513d602081101561355d57600080fd5b50505b50600195945050505050565b600080600083600160a060020a031663f1be16796040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156135af57600080fd5b505af11580156135c3573d6000803e3d6000fd5b505050506040513d60208110156135d957600080fd5b50519150600160a060020a03821615156135f6576000925061100a565b6135ff82611f4f565b151561360e576000925061100a565b50604080517ff76514c7000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301529151839283169163f76514c79160248083019260209291908290030181600087803b15801561131257600080fd5b60095490565b600080600080600080600061368d6114ed565b95506136976111f5565b600160a060020a038716600090815260136020526040902054909550935083156136c357839650611d64565b84600160a060020a031663295c39a56040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561370157600080fd5b505af1158015613715573d6000803e3d6000fd5b505050506040513d602081101561372b57600080fd5b5051604080517fda0b0c360000000000000000000000000000000000000000000000000000000081529051919450600160a060020a0387169163da0b0c36916004808201926020929091908290030181600087803b15801561378c57600080fd5b505af11580156137a0573d6000803e3d6000fd5b505050506040513d60208110156137b657600080fd5b5051600160a060020a03861660009081526013602052604090205490925090506137e38284613239613cc9565b600160a060020a03969096166000908152601360205260409020869055509395945050505050565b600254604080517f9f7e1bf600000000000000000000000000000000000000000000000000000000815233600482015290516000926101009004600160a060020a031691639f7e1bf691602480830192602092919082900301818787803b15801561387557600080fd5b505af1158015613889573d6000803e3d6000fd5b505050506040513d602081101561389f57600080fd5b505115156138ac57600080fd5b336000908152600d60209081526040808320805460ff19166001179055825481517f4e94c8290000000000000000000000000000000000000000000000000000000081529151600160a060020a039190911693634e94c82993600480850194919392918390030190829087803b15801561392557600080fd5b505af1158015613939573d6000803e3d6000fd5b505050506040513d602081101561394f57600080fd5b5051600254604080517f17674e4d000000000000000000000000000000000000000000000000000000008152336004820152610100909204600160a060020a03908116602484015290519216916317674e4d916044808201926020929091908290030181600087803b15801561119657600080fd5b600082820183811015610e6057600080fd5b6000828211156139e557600080fd5b50900390565b6000828202831580613a075750828482811515613a0457fe5b04145b1515610e6057600080fd5b6000808284811515613a2057fe5b04949350505050565b62093a8090565b624f1a0090565b600290565b600f90565b6000818310613a515750816126e5565b50806126e5565b60008183811515613a6557fe5b06159392505050565b600254600090819060ff161515613a8457600080fd5b6000805460408051600080516020613d1c83398151915281527f4d61726b6574466163746f72790000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b158015613af957600080fd5b505af1158015613b0d573d6000803e3d6000fd5b505050506040513d6020811015613b2357600080fd5b50519050613b2f6124f7565b600160a060020a031663fe98184d8683613b4761367a565b6040805160e060020a63ffffffff8716028152600160a060020a0394851660048201529290931660248301526044820152905160648083019260209291908290030181600087803b158015613b9b57600080fd5b505af1158015613baf573d6000803e3d6000fd5b505050506040513d6020811015613bc557600080fd5b5050600054604080517f1d48329b000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152306024820152604481018c9052606481018b9052898316608482015288831660a482015287831660c482015260e481018790526101048101869052905191831691631d48329b91349161012480830192602092919082900301818588803b158015613c6b57600080fd5b505af1158015613c7f573d6000803e3d6000fd5b50505050506040513d6020811015613c9657600080fd5b5051600160a060020a0381166000908152600d60205260409020805460ff19166001179055915050979650505050505050565b606490565b61271090565b600390565b6000818311613a515750816126e5565b662386f26fc1000090565b60025460009060ff1615613d0757600080fd5b506002805460ff19166001908117909155905600f39ec1f700000000000000000000000000000000000000000000000000000000a165627a7a72305820eae6db2807d07bd7bbcb987c0398eb61e2d3b36938c1f6a37d1218692fc7c1f40029`

// DeployUniverse deploys a new Ethereum contract, binding an instance of Universe to it.
func DeployUniverse(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Universe, error) {
	parsed, err := abi.JSON(strings.NewReader(UniverseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UniverseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Universe{UniverseCaller: UniverseCaller{contract: contract}, UniverseTransactor: UniverseTransactor{contract: contract}, UniverseFilterer: UniverseFilterer{contract: contract}}, nil
}

// Universe is an auto generated Go binding around an Ethereum contract.
type Universe struct {
	UniverseCaller     // Read-only binding to the contract
	UniverseTransactor // Write-only binding to the contract
	UniverseFilterer   // Log filterer for contract events
}

// UniverseCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniverseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniverseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniverseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniverseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniverseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniverseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniverseSession struct {
	Contract     *Universe         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniverseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniverseCallerSession struct {
	Contract *UniverseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UniverseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniverseTransactorSession struct {
	Contract     *UniverseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UniverseRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniverseRaw struct {
	Contract *Universe // Generic contract binding to access the raw methods on
}

// UniverseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniverseCallerRaw struct {
	Contract *UniverseCaller // Generic read-only contract binding to access the raw methods on
}

// UniverseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniverseTransactorRaw struct {
	Contract *UniverseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniverse creates a new instance of Universe, bound to a specific deployed contract.
func NewUniverse(address common.Address, backend bind.ContractBackend) (*Universe, error) {
	contract, err := bindUniverse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Universe{UniverseCaller: UniverseCaller{contract: contract}, UniverseTransactor: UniverseTransactor{contract: contract}, UniverseFilterer: UniverseFilterer{contract: contract}}, nil
}

// NewUniverseCaller creates a new read-only instance of Universe, bound to a specific deployed contract.
func NewUniverseCaller(address common.Address, caller bind.ContractCaller) (*UniverseCaller, error) {
	contract, err := bindUniverse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniverseCaller{contract: contract}, nil
}

// NewUniverseTransactor creates a new write-only instance of Universe, bound to a specific deployed contract.
func NewUniverseTransactor(address common.Address, transactor bind.ContractTransactor) (*UniverseTransactor, error) {
	contract, err := bindUniverse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniverseTransactor{contract: contract}, nil
}

// NewUniverseFilterer creates a new log filterer instance of Universe, bound to a specific deployed contract.
func NewUniverseFilterer(address common.Address, filterer bind.ContractFilterer) (*UniverseFilterer, error) {
	contract, err := bindUniverse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniverseFilterer{contract: contract}, nil
}

// bindUniverse binds a generic wrapper to an already deployed contract.
func bindUniverse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniverseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universe *UniverseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Universe.Contract.UniverseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universe *UniverseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.Contract.UniverseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universe *UniverseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Universe.Contract.UniverseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universe *UniverseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Universe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universe *UniverseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universe *UniverseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Universe.Contract.contract.Transact(opts, method, params...)
}

// CalculateFloatingValue is a free data retrieval call binding the contract method 0x5bafecf5.
//
// Solidity: function calculateFloatingValue(_badMarkets uint256, _totalMarkets uint256, _targetDivisor uint256, _previousValue uint256, _defaultValue uint256, _floor uint256) constant returns(_newValue uint256)
func (_Universe *UniverseCaller) CalculateFloatingValue(opts *bind.CallOpts, _badMarkets *big.Int, _totalMarkets *big.Int, _targetDivisor *big.Int, _previousValue *big.Int, _defaultValue *big.Int, _floor *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "calculateFloatingValue", _badMarkets, _totalMarkets, _targetDivisor, _previousValue, _defaultValue, _floor)
	return *ret0, err
}

// CalculateFloatingValue is a free data retrieval call binding the contract method 0x5bafecf5.
//
// Solidity: function calculateFloatingValue(_badMarkets uint256, _totalMarkets uint256, _targetDivisor uint256, _previousValue uint256, _defaultValue uint256, _floor uint256) constant returns(_newValue uint256)
func (_Universe *UniverseSession) CalculateFloatingValue(_badMarkets *big.Int, _totalMarkets *big.Int, _targetDivisor *big.Int, _previousValue *big.Int, _defaultValue *big.Int, _floor *big.Int) (*big.Int, error) {
	return _Universe.Contract.CalculateFloatingValue(&_Universe.CallOpts, _badMarkets, _totalMarkets, _targetDivisor, _previousValue, _defaultValue, _floor)
}

// CalculateFloatingValue is a free data retrieval call binding the contract method 0x5bafecf5.
//
// Solidity: function calculateFloatingValue(_badMarkets uint256, _totalMarkets uint256, _targetDivisor uint256, _previousValue uint256, _defaultValue uint256, _floor uint256) constant returns(_newValue uint256)
func (_Universe *UniverseCallerSession) CalculateFloatingValue(_badMarkets *big.Int, _totalMarkets *big.Int, _targetDivisor *big.Int, _previousValue *big.Int, _defaultValue *big.Int, _floor *big.Int) (*big.Int, error) {
	return _Universe.Contract.CalculateFloatingValue(&_Universe.CallOpts, _badMarkets, _totalMarkets, _targetDivisor, _previousValue, _defaultValue, _floor)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Universe *UniverseCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Universe *UniverseSession) ControllerLookupName() ([32]byte, error) {
	return _Universe.Contract.ControllerLookupName(&_Universe.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Universe *UniverseCallerSession) ControllerLookupName() ([32]byte, error) {
	return _Universe.Contract.ControllerLookupName(&_Universe.CallOpts)
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(_parentPayoutDistributionHash bytes32) constant returns(address)
func (_Universe *UniverseCaller) GetChildUniverse(opts *bind.CallOpts, _parentPayoutDistributionHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getChildUniverse", _parentPayoutDistributionHash)
	return *ret0, err
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(_parentPayoutDistributionHash bytes32) constant returns(address)
func (_Universe *UniverseSession) GetChildUniverse(_parentPayoutDistributionHash [32]byte) (common.Address, error) {
	return _Universe.Contract.GetChildUniverse(&_Universe.CallOpts, _parentPayoutDistributionHash)
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(_parentPayoutDistributionHash bytes32) constant returns(address)
func (_Universe *UniverseCallerSession) GetChildUniverse(_parentPayoutDistributionHash [32]byte) (common.Address, error) {
	return _Universe.Contract.GetChildUniverse(&_Universe.CallOpts, _parentPayoutDistributionHash)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Universe *UniverseCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Universe *UniverseSession) GetController() (common.Address, error) {
	return _Universe.Contract.GetController(&_Universe.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Universe *UniverseCallerSession) GetController() (common.Address, error) {
	return _Universe.Contract.GetController(&_Universe.CallOpts)
}

// GetCurrentFeeWindow is a free data retrieval call binding the contract method 0x79f96600.
//
// Solidity: function getCurrentFeeWindow() constant returns(address)
func (_Universe *UniverseCaller) GetCurrentFeeWindow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getCurrentFeeWindow")
	return *ret0, err
}

// GetCurrentFeeWindow is a free data retrieval call binding the contract method 0x79f96600.
//
// Solidity: function getCurrentFeeWindow() constant returns(address)
func (_Universe *UniverseSession) GetCurrentFeeWindow() (common.Address, error) {
	return _Universe.Contract.GetCurrentFeeWindow(&_Universe.CallOpts)
}

// GetCurrentFeeWindow is a free data retrieval call binding the contract method 0x79f96600.
//
// Solidity: function getCurrentFeeWindow() constant returns(address)
func (_Universe *UniverseCallerSession) GetCurrentFeeWindow() (common.Address, error) {
	return _Universe.Contract.GetCurrentFeeWindow(&_Universe.CallOpts)
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x331172f3.
//
// Solidity: function getDisputeRoundDurationInSeconds() constant returns(uint256)
func (_Universe *UniverseCaller) GetDisputeRoundDurationInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getDisputeRoundDurationInSeconds")
	return *ret0, err
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x331172f3.
//
// Solidity: function getDisputeRoundDurationInSeconds() constant returns(uint256)
func (_Universe *UniverseSession) GetDisputeRoundDurationInSeconds() (*big.Int, error) {
	return _Universe.Contract.GetDisputeRoundDurationInSeconds(&_Universe.CallOpts)
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x331172f3.
//
// Solidity: function getDisputeRoundDurationInSeconds() constant returns(uint256)
func (_Universe *UniverseCallerSession) GetDisputeRoundDurationInSeconds() (*big.Int, error) {
	return _Universe.Contract.GetDisputeRoundDurationInSeconds(&_Universe.CallOpts)
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() constant returns(uint256)
func (_Universe *UniverseCaller) GetDisputeThresholdForDisputePacing(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getDisputeThresholdForDisputePacing")
	return *ret0, err
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() constant returns(uint256)
func (_Universe *UniverseSession) GetDisputeThresholdForDisputePacing() (*big.Int, error) {
	return _Universe.Contract.GetDisputeThresholdForDisputePacing(&_Universe.CallOpts)
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() constant returns(uint256)
func (_Universe *UniverseCallerSession) GetDisputeThresholdForDisputePacing() (*big.Int, error) {
	return _Universe.Contract.GetDisputeThresholdForDisputePacing(&_Universe.CallOpts)
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() constant returns(uint256)
func (_Universe *UniverseCaller) GetDisputeThresholdForFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getDisputeThresholdForFork")
	return *ret0, err
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() constant returns(uint256)
func (_Universe *UniverseSession) GetDisputeThresholdForFork() (*big.Int, error) {
	return _Universe.Contract.GetDisputeThresholdForFork(&_Universe.CallOpts)
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() constant returns(uint256)
func (_Universe *UniverseCallerSession) GetDisputeThresholdForFork() (*big.Int, error) {
	return _Universe.Contract.GetDisputeThresholdForFork(&_Universe.CallOpts)
}

// GetFeeWindow is a free data retrieval call binding the contract method 0x40b73897.
//
// Solidity: function getFeeWindow(_feeWindowId uint256) constant returns(address)
func (_Universe *UniverseCaller) GetFeeWindow(opts *bind.CallOpts, _feeWindowId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getFeeWindow", _feeWindowId)
	return *ret0, err
}

// GetFeeWindow is a free data retrieval call binding the contract method 0x40b73897.
//
// Solidity: function getFeeWindow(_feeWindowId uint256) constant returns(address)
func (_Universe *UniverseSession) GetFeeWindow(_feeWindowId *big.Int) (common.Address, error) {
	return _Universe.Contract.GetFeeWindow(&_Universe.CallOpts, _feeWindowId)
}

// GetFeeWindow is a free data retrieval call binding the contract method 0x40b73897.
//
// Solidity: function getFeeWindow(_feeWindowId uint256) constant returns(address)
func (_Universe *UniverseCallerSession) GetFeeWindow(_feeWindowId *big.Int) (common.Address, error) {
	return _Universe.Contract.GetFeeWindow(&_Universe.CallOpts, _feeWindowId)
}

// GetFeeWindowByTimestamp is a free data retrieval call binding the contract method 0xaa2bebb7.
//
// Solidity: function getFeeWindowByTimestamp(_timestamp uint256) constant returns(address)
func (_Universe *UniverseCaller) GetFeeWindowByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getFeeWindowByTimestamp", _timestamp)
	return *ret0, err
}

// GetFeeWindowByTimestamp is a free data retrieval call binding the contract method 0xaa2bebb7.
//
// Solidity: function getFeeWindowByTimestamp(_timestamp uint256) constant returns(address)
func (_Universe *UniverseSession) GetFeeWindowByTimestamp(_timestamp *big.Int) (common.Address, error) {
	return _Universe.Contract.GetFeeWindowByTimestamp(&_Universe.CallOpts, _timestamp)
}

// GetFeeWindowByTimestamp is a free data retrieval call binding the contract method 0xaa2bebb7.
//
// Solidity: function getFeeWindowByTimestamp(_timestamp uint256) constant returns(address)
func (_Universe *UniverseCallerSession) GetFeeWindowByTimestamp(_timestamp *big.Int) (common.Address, error) {
	return _Universe.Contract.GetFeeWindowByTimestamp(&_Universe.CallOpts, _timestamp)
}

// GetFeeWindowId is a free data retrieval call binding the contract method 0xdae8af77.
//
// Solidity: function getFeeWindowId(_timestamp uint256) constant returns(uint256)
func (_Universe *UniverseCaller) GetFeeWindowId(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getFeeWindowId", _timestamp)
	return *ret0, err
}

// GetFeeWindowId is a free data retrieval call binding the contract method 0xdae8af77.
//
// Solidity: function getFeeWindowId(_timestamp uint256) constant returns(uint256)
func (_Universe *UniverseSession) GetFeeWindowId(_timestamp *big.Int) (*big.Int, error) {
	return _Universe.Contract.GetFeeWindowId(&_Universe.CallOpts, _timestamp)
}

// GetFeeWindowId is a free data retrieval call binding the contract method 0xdae8af77.
//
// Solidity: function getFeeWindowId(_timestamp uint256) constant returns(uint256)
func (_Universe *UniverseCallerSession) GetFeeWindowId(_timestamp *big.Int) (*big.Int, error) {
	return _Universe.Contract.GetFeeWindowId(&_Universe.CallOpts, _timestamp)
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() constant returns(uint256)
func (_Universe *UniverseCaller) GetForkEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getForkEndTime")
	return *ret0, err
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() constant returns(uint256)
func (_Universe *UniverseSession) GetForkEndTime() (*big.Int, error) {
	return _Universe.Contract.GetForkEndTime(&_Universe.CallOpts)
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() constant returns(uint256)
func (_Universe *UniverseCallerSession) GetForkEndTime() (*big.Int, error) {
	return _Universe.Contract.GetForkEndTime(&_Universe.CallOpts)
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() constant returns(uint256)
func (_Universe *UniverseCaller) GetForkReputationGoal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getForkReputationGoal")
	return *ret0, err
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() constant returns(uint256)
func (_Universe *UniverseSession) GetForkReputationGoal() (*big.Int, error) {
	return _Universe.Contract.GetForkReputationGoal(&_Universe.CallOpts)
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() constant returns(uint256)
func (_Universe *UniverseCallerSession) GetForkReputationGoal() (*big.Int, error) {
	return _Universe.Contract.GetForkReputationGoal(&_Universe.CallOpts)
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() constant returns(address)
func (_Universe *UniverseCaller) GetForkingMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getForkingMarket")
	return *ret0, err
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() constant returns(address)
func (_Universe *UniverseSession) GetForkingMarket() (common.Address, error) {
	return _Universe.Contract.GetForkingMarket(&_Universe.CallOpts)
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() constant returns(address)
func (_Universe *UniverseCallerSession) GetForkingMarket() (common.Address, error) {
	return _Universe.Contract.GetForkingMarket(&_Universe.CallOpts)
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() constant returns(uint256)
func (_Universe *UniverseCaller) GetInitialReportMinValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getInitialReportMinValue")
	return *ret0, err
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() constant returns(uint256)
func (_Universe *UniverseSession) GetInitialReportMinValue() (*big.Int, error) {
	return _Universe.Contract.GetInitialReportMinValue(&_Universe.CallOpts)
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() constant returns(uint256)
func (_Universe *UniverseCallerSession) GetInitialReportMinValue() (*big.Int, error) {
	return _Universe.Contract.GetInitialReportMinValue(&_Universe.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_Universe *UniverseCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_Universe *UniverseSession) GetInitialized() (bool, error) {
	return _Universe.Contract.GetInitialized(&_Universe.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_Universe *UniverseCallerSession) GetInitialized() (bool, error) {
	return _Universe.Contract.GetInitialized(&_Universe.CallOpts)
}

// GetNextFeeWindow is a free data retrieval call binding the contract method 0xedd1d02e.
//
// Solidity: function getNextFeeWindow() constant returns(address)
func (_Universe *UniverseCaller) GetNextFeeWindow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getNextFeeWindow")
	return *ret0, err
}

// GetNextFeeWindow is a free data retrieval call binding the contract method 0xedd1d02e.
//
// Solidity: function getNextFeeWindow() constant returns(address)
func (_Universe *UniverseSession) GetNextFeeWindow() (common.Address, error) {
	return _Universe.Contract.GetNextFeeWindow(&_Universe.CallOpts)
}

// GetNextFeeWindow is a free data retrieval call binding the contract method 0xedd1d02e.
//
// Solidity: function getNextFeeWindow() constant returns(address)
func (_Universe *UniverseCallerSession) GetNextFeeWindow() (common.Address, error) {
	return _Universe.Contract.GetNextFeeWindow(&_Universe.CallOpts)
}

// GetOpenInterestInAttoEth is a free data retrieval call binding the contract method 0x6a9d7629.
//
// Solidity: function getOpenInterestInAttoEth() constant returns(uint256)
func (_Universe *UniverseCaller) GetOpenInterestInAttoEth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getOpenInterestInAttoEth")
	return *ret0, err
}

// GetOpenInterestInAttoEth is a free data retrieval call binding the contract method 0x6a9d7629.
//
// Solidity: function getOpenInterestInAttoEth() constant returns(uint256)
func (_Universe *UniverseSession) GetOpenInterestInAttoEth() (*big.Int, error) {
	return _Universe.Contract.GetOpenInterestInAttoEth(&_Universe.CallOpts)
}

// GetOpenInterestInAttoEth is a free data retrieval call binding the contract method 0x6a9d7629.
//
// Solidity: function getOpenInterestInAttoEth() constant returns(uint256)
func (_Universe *UniverseCallerSession) GetOpenInterestInAttoEth() (*big.Int, error) {
	return _Universe.Contract.GetOpenInterestInAttoEth(&_Universe.CallOpts)
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() constant returns(bytes32)
func (_Universe *UniverseCaller) GetParentPayoutDistributionHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getParentPayoutDistributionHash")
	return *ret0, err
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() constant returns(bytes32)
func (_Universe *UniverseSession) GetParentPayoutDistributionHash() ([32]byte, error) {
	return _Universe.Contract.GetParentPayoutDistributionHash(&_Universe.CallOpts)
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() constant returns(bytes32)
func (_Universe *UniverseCallerSession) GetParentPayoutDistributionHash() ([32]byte, error) {
	return _Universe.Contract.GetParentPayoutDistributionHash(&_Universe.CallOpts)
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() constant returns(address)
func (_Universe *UniverseCaller) GetParentUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getParentUniverse")
	return *ret0, err
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() constant returns(address)
func (_Universe *UniverseSession) GetParentUniverse() (common.Address, error) {
	return _Universe.Contract.GetParentUniverse(&_Universe.CallOpts)
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() constant returns(address)
func (_Universe *UniverseCallerSession) GetParentUniverse() (common.Address, error) {
	return _Universe.Contract.GetParentUniverse(&_Universe.CallOpts)
}

// GetPreviousFeeWindow is a free data retrieval call binding the contract method 0x6def1efc.
//
// Solidity: function getPreviousFeeWindow() constant returns(address)
func (_Universe *UniverseCaller) GetPreviousFeeWindow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getPreviousFeeWindow")
	return *ret0, err
}

// GetPreviousFeeWindow is a free data retrieval call binding the contract method 0x6def1efc.
//
// Solidity: function getPreviousFeeWindow() constant returns(address)
func (_Universe *UniverseSession) GetPreviousFeeWindow() (common.Address, error) {
	return _Universe.Contract.GetPreviousFeeWindow(&_Universe.CallOpts)
}

// GetPreviousFeeWindow is a free data retrieval call binding the contract method 0x6def1efc.
//
// Solidity: function getPreviousFeeWindow() constant returns(address)
func (_Universe *UniverseCallerSession) GetPreviousFeeWindow() (common.Address, error) {
	return _Universe.Contract.GetPreviousFeeWindow(&_Universe.CallOpts)
}

// GetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x13ad9ce1.
//
// Solidity: function getRepMarketCapInAttoeth() constant returns(uint256)
func (_Universe *UniverseCaller) GetRepMarketCapInAttoeth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getRepMarketCapInAttoeth")
	return *ret0, err
}

// GetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x13ad9ce1.
//
// Solidity: function getRepMarketCapInAttoeth() constant returns(uint256)
func (_Universe *UniverseSession) GetRepMarketCapInAttoeth() (*big.Int, error) {
	return _Universe.Contract.GetRepMarketCapInAttoeth(&_Universe.CallOpts)
}

// GetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x13ad9ce1.
//
// Solidity: function getRepMarketCapInAttoeth() constant returns(uint256)
func (_Universe *UniverseCallerSession) GetRepMarketCapInAttoeth() (*big.Int, error) {
	return _Universe.Contract.GetRepMarketCapInAttoeth(&_Universe.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_Universe *UniverseCaller) GetReputationToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getReputationToken")
	return *ret0, err
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_Universe *UniverseSession) GetReputationToken() (common.Address, error) {
	return _Universe.Contract.GetReputationToken(&_Universe.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_Universe *UniverseCallerSession) GetReputationToken() (common.Address, error) {
	return _Universe.Contract.GetReputationToken(&_Universe.CallOpts)
}

// GetTargetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x4a719a27.
//
// Solidity: function getTargetRepMarketCapInAttoeth() constant returns(uint256)
func (_Universe *UniverseCaller) GetTargetRepMarketCapInAttoeth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getTargetRepMarketCapInAttoeth")
	return *ret0, err
}

// GetTargetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x4a719a27.
//
// Solidity: function getTargetRepMarketCapInAttoeth() constant returns(uint256)
func (_Universe *UniverseSession) GetTargetRepMarketCapInAttoeth() (*big.Int, error) {
	return _Universe.Contract.GetTargetRepMarketCapInAttoeth(&_Universe.CallOpts)
}

// GetTargetRepMarketCapInAttoeth is a free data retrieval call binding the contract method 0x4a719a27.
//
// Solidity: function getTargetRepMarketCapInAttoeth() constant returns(uint256)
func (_Universe *UniverseCallerSession) GetTargetRepMarketCapInAttoeth() (*big.Int, error) {
	return _Universe.Contract.GetTargetRepMarketCapInAttoeth(&_Universe.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_Universe *UniverseCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_Universe *UniverseSession) GetTypeName() ([32]byte, error) {
	return _Universe.Contract.GetTypeName(&_Universe.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_Universe *UniverseCallerSession) GetTypeName() ([32]byte, error) {
	return _Universe.Contract.GetTypeName(&_Universe.CallOpts)
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() constant returns(address)
func (_Universe *UniverseCaller) GetWinningChildUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "getWinningChildUniverse")
	return *ret0, err
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() constant returns(address)
func (_Universe *UniverseSession) GetWinningChildUniverse() (common.Address, error) {
	return _Universe.Contract.GetWinningChildUniverse(&_Universe.CallOpts)
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() constant returns(address)
func (_Universe *UniverseCallerSession) GetWinningChildUniverse() (common.Address, error) {
	return _Universe.Contract.GetWinningChildUniverse(&_Universe.CallOpts)
}

// IsContainerForFeeToken is a free data retrieval call binding the contract method 0x26d16bc9.
//
// Solidity: function isContainerForFeeToken(_shadyFeeToken address) constant returns(bool)
func (_Universe *UniverseCaller) IsContainerForFeeToken(opts *bind.CallOpts, _shadyFeeToken common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "isContainerForFeeToken", _shadyFeeToken)
	return *ret0, err
}

// IsContainerForFeeToken is a free data retrieval call binding the contract method 0x26d16bc9.
//
// Solidity: function isContainerForFeeToken(_shadyFeeToken address) constant returns(bool)
func (_Universe *UniverseSession) IsContainerForFeeToken(_shadyFeeToken common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForFeeToken(&_Universe.CallOpts, _shadyFeeToken)
}

// IsContainerForFeeToken is a free data retrieval call binding the contract method 0x26d16bc9.
//
// Solidity: function isContainerForFeeToken(_shadyFeeToken address) constant returns(bool)
func (_Universe *UniverseCallerSession) IsContainerForFeeToken(_shadyFeeToken common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForFeeToken(&_Universe.CallOpts, _shadyFeeToken)
}

// IsContainerForFeeWindow is a free data retrieval call binding the contract method 0xc7c88d70.
//
// Solidity: function isContainerForFeeWindow(_shadyFeeWindow address) constant returns(bool)
func (_Universe *UniverseCaller) IsContainerForFeeWindow(opts *bind.CallOpts, _shadyFeeWindow common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "isContainerForFeeWindow", _shadyFeeWindow)
	return *ret0, err
}

// IsContainerForFeeWindow is a free data retrieval call binding the contract method 0xc7c88d70.
//
// Solidity: function isContainerForFeeWindow(_shadyFeeWindow address) constant returns(bool)
func (_Universe *UniverseSession) IsContainerForFeeWindow(_shadyFeeWindow common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForFeeWindow(&_Universe.CallOpts, _shadyFeeWindow)
}

// IsContainerForFeeWindow is a free data retrieval call binding the contract method 0xc7c88d70.
//
// Solidity: function isContainerForFeeWindow(_shadyFeeWindow address) constant returns(bool)
func (_Universe *UniverseCallerSession) IsContainerForFeeWindow(_shadyFeeWindow common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForFeeWindow(&_Universe.CallOpts, _shadyFeeWindow)
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(_shadyMarket address) constant returns(bool)
func (_Universe *UniverseCaller) IsContainerForMarket(opts *bind.CallOpts, _shadyMarket common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "isContainerForMarket", _shadyMarket)
	return *ret0, err
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(_shadyMarket address) constant returns(bool)
func (_Universe *UniverseSession) IsContainerForMarket(_shadyMarket common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForMarket(&_Universe.CallOpts, _shadyMarket)
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(_shadyMarket address) constant returns(bool)
func (_Universe *UniverseCallerSession) IsContainerForMarket(_shadyMarket common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForMarket(&_Universe.CallOpts, _shadyMarket)
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(_shadyReportingParticipant address) constant returns(bool)
func (_Universe *UniverseCaller) IsContainerForReportingParticipant(opts *bind.CallOpts, _shadyReportingParticipant common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "isContainerForReportingParticipant", _shadyReportingParticipant)
	return *ret0, err
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(_shadyReportingParticipant address) constant returns(bool)
func (_Universe *UniverseSession) IsContainerForReportingParticipant(_shadyReportingParticipant common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForReportingParticipant(&_Universe.CallOpts, _shadyReportingParticipant)
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(_shadyReportingParticipant address) constant returns(bool)
func (_Universe *UniverseCallerSession) IsContainerForReportingParticipant(_shadyReportingParticipant common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForReportingParticipant(&_Universe.CallOpts, _shadyReportingParticipant)
}

// IsContainerForShareToken is a free data retrieval call binding the contract method 0x509a1061.
//
// Solidity: function isContainerForShareToken(_shadyShareToken address) constant returns(bool)
func (_Universe *UniverseCaller) IsContainerForShareToken(opts *bind.CallOpts, _shadyShareToken common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "isContainerForShareToken", _shadyShareToken)
	return *ret0, err
}

// IsContainerForShareToken is a free data retrieval call binding the contract method 0x509a1061.
//
// Solidity: function isContainerForShareToken(_shadyShareToken address) constant returns(bool)
func (_Universe *UniverseSession) IsContainerForShareToken(_shadyShareToken common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForShareToken(&_Universe.CallOpts, _shadyShareToken)
}

// IsContainerForShareToken is a free data retrieval call binding the contract method 0x509a1061.
//
// Solidity: function isContainerForShareToken(_shadyShareToken address) constant returns(bool)
func (_Universe *UniverseCallerSession) IsContainerForShareToken(_shadyShareToken common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForShareToken(&_Universe.CallOpts, _shadyShareToken)
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() constant returns(bool)
func (_Universe *UniverseCaller) IsForking(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "isForking")
	return *ret0, err
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() constant returns(bool)
func (_Universe *UniverseSession) IsForking() (bool, error) {
	return _Universe.Contract.IsForking(&_Universe.CallOpts)
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() constant returns(bool)
func (_Universe *UniverseCallerSession) IsForking() (bool, error) {
	return _Universe.Contract.IsForking(&_Universe.CallOpts)
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(_shadyChild address) constant returns(bool)
func (_Universe *UniverseCaller) IsParentOf(opts *bind.CallOpts, _shadyChild common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Universe.contract.Call(opts, out, "isParentOf", _shadyChild)
	return *ret0, err
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(_shadyChild address) constant returns(bool)
func (_Universe *UniverseSession) IsParentOf(_shadyChild common.Address) (bool, error) {
	return _Universe.Contract.IsParentOf(&_Universe.CallOpts, _shadyChild)
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(_shadyChild address) constant returns(bool)
func (_Universe *UniverseCallerSession) IsParentOf(_shadyChild common.Address) (bool, error) {
	return _Universe.Contract.IsParentOf(&_Universe.CallOpts, _shadyChild)
}

// AddMarketTo is a paid mutator transaction binding the contract method 0xfeeda367.
//
// Solidity: function addMarketTo() returns(bool)
func (_Universe *UniverseTransactor) AddMarketTo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "addMarketTo")
}

// AddMarketTo is a paid mutator transaction binding the contract method 0xfeeda367.
//
// Solidity: function addMarketTo() returns(bool)
func (_Universe *UniverseSession) AddMarketTo() (*types.Transaction, error) {
	return _Universe.Contract.AddMarketTo(&_Universe.TransactOpts)
}

// AddMarketTo is a paid mutator transaction binding the contract method 0xfeeda367.
//
// Solidity: function addMarketTo() returns(bool)
func (_Universe *UniverseTransactorSession) AddMarketTo() (*types.Transaction, error) {
	return _Universe.Contract.AddMarketTo(&_Universe.TransactOpts)
}

// BuyParticipationTokens is a paid mutator transaction binding the contract method 0x58ab010d.
//
// Solidity: function buyParticipationTokens(_attotokens uint256) returns(bool)
func (_Universe *UniverseTransactor) BuyParticipationTokens(opts *bind.TransactOpts, _attotokens *big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "buyParticipationTokens", _attotokens)
}

// BuyParticipationTokens is a paid mutator transaction binding the contract method 0x58ab010d.
//
// Solidity: function buyParticipationTokens(_attotokens uint256) returns(bool)
func (_Universe *UniverseSession) BuyParticipationTokens(_attotokens *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.BuyParticipationTokens(&_Universe.TransactOpts, _attotokens)
}

// BuyParticipationTokens is a paid mutator transaction binding the contract method 0x58ab010d.
//
// Solidity: function buyParticipationTokens(_attotokens uint256) returns(bool)
func (_Universe *UniverseTransactorSession) BuyParticipationTokens(_attotokens *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.BuyParticipationTokens(&_Universe.TransactOpts, _attotokens)
}

// CreateCategoricalMarket is a paid mutator transaction binding the contract method 0xd28a1685.
//
// Solidity: function createCategoricalMarket(_endTime uint256, _feePerEthInWei uint256, _denominationToken address, _designatedReporterAddress address, _outcomes bytes32[], _topic bytes32, _description string, _extraInfo string) returns(_newMarket address)
func (_Universe *UniverseTransactor) CreateCategoricalMarket(opts *bind.TransactOpts, _endTime *big.Int, _feePerEthInWei *big.Int, _denominationToken common.Address, _designatedReporterAddress common.Address, _outcomes [][32]byte, _topic [32]byte, _description string, _extraInfo string) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "createCategoricalMarket", _endTime, _feePerEthInWei, _denominationToken, _designatedReporterAddress, _outcomes, _topic, _description, _extraInfo)
}

// CreateCategoricalMarket is a paid mutator transaction binding the contract method 0xd28a1685.
//
// Solidity: function createCategoricalMarket(_endTime uint256, _feePerEthInWei uint256, _denominationToken address, _designatedReporterAddress address, _outcomes bytes32[], _topic bytes32, _description string, _extraInfo string) returns(_newMarket address)
func (_Universe *UniverseSession) CreateCategoricalMarket(_endTime *big.Int, _feePerEthInWei *big.Int, _denominationToken common.Address, _designatedReporterAddress common.Address, _outcomes [][32]byte, _topic [32]byte, _description string, _extraInfo string) (*types.Transaction, error) {
	return _Universe.Contract.CreateCategoricalMarket(&_Universe.TransactOpts, _endTime, _feePerEthInWei, _denominationToken, _designatedReporterAddress, _outcomes, _topic, _description, _extraInfo)
}

// CreateCategoricalMarket is a paid mutator transaction binding the contract method 0xd28a1685.
//
// Solidity: function createCategoricalMarket(_endTime uint256, _feePerEthInWei uint256, _denominationToken address, _designatedReporterAddress address, _outcomes bytes32[], _topic bytes32, _description string, _extraInfo string) returns(_newMarket address)
func (_Universe *UniverseTransactorSession) CreateCategoricalMarket(_endTime *big.Int, _feePerEthInWei *big.Int, _denominationToken common.Address, _designatedReporterAddress common.Address, _outcomes [][32]byte, _topic [32]byte, _description string, _extraInfo string) (*types.Transaction, error) {
	return _Universe.Contract.CreateCategoricalMarket(&_Universe.TransactOpts, _endTime, _feePerEthInWei, _denominationToken, _designatedReporterAddress, _outcomes, _topic, _description, _extraInfo)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0xdf428e3b.
//
// Solidity: function createChildUniverse(_parentPayoutNumerators uint256[], _parentInvalid bool) returns(address)
func (_Universe *UniverseTransactor) CreateChildUniverse(opts *bind.TransactOpts, _parentPayoutNumerators []*big.Int, _parentInvalid bool) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "createChildUniverse", _parentPayoutNumerators, _parentInvalid)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0xdf428e3b.
//
// Solidity: function createChildUniverse(_parentPayoutNumerators uint256[], _parentInvalid bool) returns(address)
func (_Universe *UniverseSession) CreateChildUniverse(_parentPayoutNumerators []*big.Int, _parentInvalid bool) (*types.Transaction, error) {
	return _Universe.Contract.CreateChildUniverse(&_Universe.TransactOpts, _parentPayoutNumerators, _parentInvalid)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0xdf428e3b.
//
// Solidity: function createChildUniverse(_parentPayoutNumerators uint256[], _parentInvalid bool) returns(address)
func (_Universe *UniverseTransactorSession) CreateChildUniverse(_parentPayoutNumerators []*big.Int, _parentInvalid bool) (*types.Transaction, error) {
	return _Universe.Contract.CreateChildUniverse(&_Universe.TransactOpts, _parentPayoutNumerators, _parentInvalid)
}

// CreateScalarMarket is a paid mutator transaction binding the contract method 0x72d61dfa.
//
// Solidity: function createScalarMarket(_endTime uint256, _feePerEthInWei uint256, _denominationToken address, _designatedReporterAddress address, _minPrice int256, _maxPrice int256, _numTicks uint256, _topic bytes32, _description string, _extraInfo string) returns(_newMarket address)
func (_Universe *UniverseTransactor) CreateScalarMarket(opts *bind.TransactOpts, _endTime *big.Int, _feePerEthInWei *big.Int, _denominationToken common.Address, _designatedReporterAddress common.Address, _minPrice *big.Int, _maxPrice *big.Int, _numTicks *big.Int, _topic [32]byte, _description string, _extraInfo string) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "createScalarMarket", _endTime, _feePerEthInWei, _denominationToken, _designatedReporterAddress, _minPrice, _maxPrice, _numTicks, _topic, _description, _extraInfo)
}

// CreateScalarMarket is a paid mutator transaction binding the contract method 0x72d61dfa.
//
// Solidity: function createScalarMarket(_endTime uint256, _feePerEthInWei uint256, _denominationToken address, _designatedReporterAddress address, _minPrice int256, _maxPrice int256, _numTicks uint256, _topic bytes32, _description string, _extraInfo string) returns(_newMarket address)
func (_Universe *UniverseSession) CreateScalarMarket(_endTime *big.Int, _feePerEthInWei *big.Int, _denominationToken common.Address, _designatedReporterAddress common.Address, _minPrice *big.Int, _maxPrice *big.Int, _numTicks *big.Int, _topic [32]byte, _description string, _extraInfo string) (*types.Transaction, error) {
	return _Universe.Contract.CreateScalarMarket(&_Universe.TransactOpts, _endTime, _feePerEthInWei, _denominationToken, _designatedReporterAddress, _minPrice, _maxPrice, _numTicks, _topic, _description, _extraInfo)
}

// CreateScalarMarket is a paid mutator transaction binding the contract method 0x72d61dfa.
//
// Solidity: function createScalarMarket(_endTime uint256, _feePerEthInWei uint256, _denominationToken address, _designatedReporterAddress address, _minPrice int256, _maxPrice int256, _numTicks uint256, _topic bytes32, _description string, _extraInfo string) returns(_newMarket address)
func (_Universe *UniverseTransactorSession) CreateScalarMarket(_endTime *big.Int, _feePerEthInWei *big.Int, _denominationToken common.Address, _designatedReporterAddress common.Address, _minPrice *big.Int, _maxPrice *big.Int, _numTicks *big.Int, _topic [32]byte, _description string, _extraInfo string) (*types.Transaction, error) {
	return _Universe.Contract.CreateScalarMarket(&_Universe.TransactOpts, _endTime, _feePerEthInWei, _denominationToken, _designatedReporterAddress, _minPrice, _maxPrice, _numTicks, _topic, _description, _extraInfo)
}

// CreateYesNoMarket is a paid mutator transaction binding the contract method 0xcabc6f14.
//
// Solidity: function createYesNoMarket(_endTime uint256, _feePerEthInWei uint256, _denominationToken address, _designatedReporterAddress address, _topic bytes32, _description string, _extraInfo string) returns(_newMarket address)
func (_Universe *UniverseTransactor) CreateYesNoMarket(opts *bind.TransactOpts, _endTime *big.Int, _feePerEthInWei *big.Int, _denominationToken common.Address, _designatedReporterAddress common.Address, _topic [32]byte, _description string, _extraInfo string) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "createYesNoMarket", _endTime, _feePerEthInWei, _denominationToken, _designatedReporterAddress, _topic, _description, _extraInfo)
}

// CreateYesNoMarket is a paid mutator transaction binding the contract method 0xcabc6f14.
//
// Solidity: function createYesNoMarket(_endTime uint256, _feePerEthInWei uint256, _denominationToken address, _designatedReporterAddress address, _topic bytes32, _description string, _extraInfo string) returns(_newMarket address)
func (_Universe *UniverseSession) CreateYesNoMarket(_endTime *big.Int, _feePerEthInWei *big.Int, _denominationToken common.Address, _designatedReporterAddress common.Address, _topic [32]byte, _description string, _extraInfo string) (*types.Transaction, error) {
	return _Universe.Contract.CreateYesNoMarket(&_Universe.TransactOpts, _endTime, _feePerEthInWei, _denominationToken, _designatedReporterAddress, _topic, _description, _extraInfo)
}

// CreateYesNoMarket is a paid mutator transaction binding the contract method 0xcabc6f14.
//
// Solidity: function createYesNoMarket(_endTime uint256, _feePerEthInWei uint256, _denominationToken address, _designatedReporterAddress address, _topic bytes32, _description string, _extraInfo string) returns(_newMarket address)
func (_Universe *UniverseTransactorSession) CreateYesNoMarket(_endTime *big.Int, _feePerEthInWei *big.Int, _denominationToken common.Address, _designatedReporterAddress common.Address, _topic [32]byte, _description string, _extraInfo string) (*types.Transaction, error) {
	return _Universe.Contract.CreateYesNoMarket(&_Universe.TransactOpts, _endTime, _feePerEthInWei, _denominationToken, _designatedReporterAddress, _topic, _description, _extraInfo)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(_amount uint256) returns(bool)
func (_Universe *UniverseTransactor) DecrementOpenInterest(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "decrementOpenInterest", _amount)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(_amount uint256) returns(bool)
func (_Universe *UniverseSession) DecrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.DecrementOpenInterest(&_Universe.TransactOpts, _amount)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(_amount uint256) returns(bool)
func (_Universe *UniverseTransactorSession) DecrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.DecrementOpenInterest(&_Universe.TransactOpts, _amount)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x0db3be6a.
//
// Solidity: function decrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_Universe *UniverseTransactor) DecrementOpenInterestFromMarket(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "decrementOpenInterestFromMarket", _amount)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x0db3be6a.
//
// Solidity: function decrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_Universe *UniverseSession) DecrementOpenInterestFromMarket(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.DecrementOpenInterestFromMarket(&_Universe.TransactOpts, _amount)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x0db3be6a.
//
// Solidity: function decrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_Universe *UniverseTransactorSession) DecrementOpenInterestFromMarket(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.DecrementOpenInterestFromMarket(&_Universe.TransactOpts, _amount)
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_Universe *UniverseTransactor) Fork(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "fork")
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_Universe *UniverseSession) Fork() (*types.Transaction, error) {
	return _Universe.Contract.Fork(&_Universe.TransactOpts)
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_Universe *UniverseTransactorSession) Fork() (*types.Transaction, error) {
	return _Universe.Contract.Fork(&_Universe.TransactOpts)
}

// GetInitialReportStakeSize is a paid mutator transaction binding the contract method 0xaf6b3615.
//
// Solidity: function getInitialReportStakeSize() returns(uint256)
func (_Universe *UniverseTransactor) GetInitialReportStakeSize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getInitialReportStakeSize")
}

// GetInitialReportStakeSize is a paid mutator transaction binding the contract method 0xaf6b3615.
//
// Solidity: function getInitialReportStakeSize() returns(uint256)
func (_Universe *UniverseSession) GetInitialReportStakeSize() (*types.Transaction, error) {
	return _Universe.Contract.GetInitialReportStakeSize(&_Universe.TransactOpts)
}

// GetInitialReportStakeSize is a paid mutator transaction binding the contract method 0xaf6b3615.
//
// Solidity: function getInitialReportStakeSize() returns(uint256)
func (_Universe *UniverseTransactorSession) GetInitialReportStakeSize() (*types.Transaction, error) {
	return _Universe.Contract.GetInitialReportStakeSize(&_Universe.TransactOpts)
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheDesignatedReportNoShowBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheDesignatedReportNoShowBond")
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheDesignatedReportNoShowBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheDesignatedReportNoShowBond(&_Universe.TransactOpts)
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheDesignatedReportNoShowBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheDesignatedReportNoShowBond(&_Universe.TransactOpts)
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheDesignatedReportStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheDesignatedReportStake")
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheDesignatedReportStake() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheDesignatedReportStake(&_Universe.TransactOpts)
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheDesignatedReportStake() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheDesignatedReportStake(&_Universe.TransactOpts)
}

// GetOrCacheMarketCreationCost is a paid mutator transaction binding the contract method 0xec86fdbd.
//
// Solidity: function getOrCacheMarketCreationCost() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheMarketCreationCost(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheMarketCreationCost")
}

// GetOrCacheMarketCreationCost is a paid mutator transaction binding the contract method 0xec86fdbd.
//
// Solidity: function getOrCacheMarketCreationCost() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheMarketCreationCost() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheMarketCreationCost(&_Universe.TransactOpts)
}

// GetOrCacheMarketCreationCost is a paid mutator transaction binding the contract method 0xec86fdbd.
//
// Solidity: function getOrCacheMarketCreationCost() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheMarketCreationCost() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheMarketCreationCost(&_Universe.TransactOpts)
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheReportingFeeDivisor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheReportingFeeDivisor")
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheReportingFeeDivisor() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheReportingFeeDivisor(&_Universe.TransactOpts)
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheReportingFeeDivisor() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheReportingFeeDivisor(&_Universe.TransactOpts)
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheValidityBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheValidityBond")
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheValidityBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheValidityBond(&_Universe.TransactOpts)
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheValidityBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheValidityBond(&_Universe.TransactOpts)
}

// GetOrCreateCurrentFeeWindow is a paid mutator transaction binding the contract method 0x68a7effc.
//
// Solidity: function getOrCreateCurrentFeeWindow() returns(address)
func (_Universe *UniverseTransactor) GetOrCreateCurrentFeeWindow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreateCurrentFeeWindow")
}

// GetOrCreateCurrentFeeWindow is a paid mutator transaction binding the contract method 0x68a7effc.
//
// Solidity: function getOrCreateCurrentFeeWindow() returns(address)
func (_Universe *UniverseSession) GetOrCreateCurrentFeeWindow() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateCurrentFeeWindow(&_Universe.TransactOpts)
}

// GetOrCreateCurrentFeeWindow is a paid mutator transaction binding the contract method 0x68a7effc.
//
// Solidity: function getOrCreateCurrentFeeWindow() returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreateCurrentFeeWindow() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateCurrentFeeWindow(&_Universe.TransactOpts)
}

// GetOrCreateFeeWindowBefore is a paid mutator transaction binding the contract method 0xe487c7a5.
//
// Solidity: function getOrCreateFeeWindowBefore(_feeWindow address) returns(address)
func (_Universe *UniverseTransactor) GetOrCreateFeeWindowBefore(opts *bind.TransactOpts, _feeWindow common.Address) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreateFeeWindowBefore", _feeWindow)
}

// GetOrCreateFeeWindowBefore is a paid mutator transaction binding the contract method 0xe487c7a5.
//
// Solidity: function getOrCreateFeeWindowBefore(_feeWindow address) returns(address)
func (_Universe *UniverseSession) GetOrCreateFeeWindowBefore(_feeWindow common.Address) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateFeeWindowBefore(&_Universe.TransactOpts, _feeWindow)
}

// GetOrCreateFeeWindowBefore is a paid mutator transaction binding the contract method 0xe487c7a5.
//
// Solidity: function getOrCreateFeeWindowBefore(_feeWindow address) returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreateFeeWindowBefore(_feeWindow common.Address) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateFeeWindowBefore(&_Universe.TransactOpts, _feeWindow)
}

// GetOrCreateFeeWindowByTimestamp is a paid mutator transaction binding the contract method 0xb66afea3.
//
// Solidity: function getOrCreateFeeWindowByTimestamp(_timestamp uint256) returns(address)
func (_Universe *UniverseTransactor) GetOrCreateFeeWindowByTimestamp(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreateFeeWindowByTimestamp", _timestamp)
}

// GetOrCreateFeeWindowByTimestamp is a paid mutator transaction binding the contract method 0xb66afea3.
//
// Solidity: function getOrCreateFeeWindowByTimestamp(_timestamp uint256) returns(address)
func (_Universe *UniverseSession) GetOrCreateFeeWindowByTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateFeeWindowByTimestamp(&_Universe.TransactOpts, _timestamp)
}

// GetOrCreateFeeWindowByTimestamp is a paid mutator transaction binding the contract method 0xb66afea3.
//
// Solidity: function getOrCreateFeeWindowByTimestamp(_timestamp uint256) returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreateFeeWindowByTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateFeeWindowByTimestamp(&_Universe.TransactOpts, _timestamp)
}

// GetOrCreateNextFeeWindow is a paid mutator transaction binding the contract method 0x0cc8c9af.
//
// Solidity: function getOrCreateNextFeeWindow() returns(address)
func (_Universe *UniverseTransactor) GetOrCreateNextFeeWindow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreateNextFeeWindow")
}

// GetOrCreateNextFeeWindow is a paid mutator transaction binding the contract method 0x0cc8c9af.
//
// Solidity: function getOrCreateNextFeeWindow() returns(address)
func (_Universe *UniverseSession) GetOrCreateNextFeeWindow() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateNextFeeWindow(&_Universe.TransactOpts)
}

// GetOrCreateNextFeeWindow is a paid mutator transaction binding the contract method 0x0cc8c9af.
//
// Solidity: function getOrCreateNextFeeWindow() returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreateNextFeeWindow() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateNextFeeWindow(&_Universe.TransactOpts)
}

// GetOrCreatePreviousFeeWindow is a paid mutator transaction binding the contract method 0x225640ab.
//
// Solidity: function getOrCreatePreviousFeeWindow() returns(address)
func (_Universe *UniverseTransactor) GetOrCreatePreviousFeeWindow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreatePreviousFeeWindow")
}

// GetOrCreatePreviousFeeWindow is a paid mutator transaction binding the contract method 0x225640ab.
//
// Solidity: function getOrCreatePreviousFeeWindow() returns(address)
func (_Universe *UniverseSession) GetOrCreatePreviousFeeWindow() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreatePreviousFeeWindow(&_Universe.TransactOpts)
}

// GetOrCreatePreviousFeeWindow is a paid mutator transaction binding the contract method 0x225640ab.
//
// Solidity: function getOrCreatePreviousFeeWindow() returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreatePreviousFeeWindow() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreatePreviousFeeWindow(&_Universe.TransactOpts)
}

// GetOrCreatePreviousPreviousFeeWindow is a paid mutator transaction binding the contract method 0x4f803516.
//
// Solidity: function getOrCreatePreviousPreviousFeeWindow() returns(address)
func (_Universe *UniverseTransactor) GetOrCreatePreviousPreviousFeeWindow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreatePreviousPreviousFeeWindow")
}

// GetOrCreatePreviousPreviousFeeWindow is a paid mutator transaction binding the contract method 0x4f803516.
//
// Solidity: function getOrCreatePreviousPreviousFeeWindow() returns(address)
func (_Universe *UniverseSession) GetOrCreatePreviousPreviousFeeWindow() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreatePreviousPreviousFeeWindow(&_Universe.TransactOpts)
}

// GetOrCreatePreviousPreviousFeeWindow is a paid mutator transaction binding the contract method 0x4f803516.
//
// Solidity: function getOrCreatePreviousPreviousFeeWindow() returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreatePreviousPreviousFeeWindow() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreatePreviousPreviousFeeWindow(&_Universe.TransactOpts)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(_amount uint256) returns(bool)
func (_Universe *UniverseTransactor) IncrementOpenInterest(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "incrementOpenInterest", _amount)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(_amount uint256) returns(bool)
func (_Universe *UniverseSession) IncrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.IncrementOpenInterest(&_Universe.TransactOpts, _amount)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(_amount uint256) returns(bool)
func (_Universe *UniverseTransactorSession) IncrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.IncrementOpenInterest(&_Universe.TransactOpts, _amount)
}

// IncrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x061777ed.
//
// Solidity: function incrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_Universe *UniverseTransactor) IncrementOpenInterestFromMarket(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "incrementOpenInterestFromMarket", _amount)
}

// IncrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x061777ed.
//
// Solidity: function incrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_Universe *UniverseSession) IncrementOpenInterestFromMarket(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.IncrementOpenInterestFromMarket(&_Universe.TransactOpts, _amount)
}

// IncrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x061777ed.
//
// Solidity: function incrementOpenInterestFromMarket(_amount uint256) returns(bool)
func (_Universe *UniverseTransactorSession) IncrementOpenInterestFromMarket(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.IncrementOpenInterestFromMarket(&_Universe.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(_parentUniverse address, _parentPayoutDistributionHash bytes32) returns(bool)
func (_Universe *UniverseTransactor) Initialize(opts *bind.TransactOpts, _parentUniverse common.Address, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "initialize", _parentUniverse, _parentPayoutDistributionHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(_parentUniverse address, _parentPayoutDistributionHash bytes32) returns(bool)
func (_Universe *UniverseSession) Initialize(_parentUniverse common.Address, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Universe.Contract.Initialize(&_Universe.TransactOpts, _parentUniverse, _parentPayoutDistributionHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(_parentUniverse address, _parentPayoutDistributionHash bytes32) returns(bool)
func (_Universe *UniverseTransactorSession) Initialize(_parentUniverse common.Address, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Universe.Contract.Initialize(&_Universe.TransactOpts, _parentUniverse, _parentPayoutDistributionHash)
}

// RedeemStake is a paid mutator transaction binding the contract method 0x8d8cc634.
//
// Solidity: function redeemStake(_reportingParticipants address[], _feeWindows address[]) returns(bool)
func (_Universe *UniverseTransactor) RedeemStake(opts *bind.TransactOpts, _reportingParticipants []common.Address, _feeWindows []common.Address) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "redeemStake", _reportingParticipants, _feeWindows)
}

// RedeemStake is a paid mutator transaction binding the contract method 0x8d8cc634.
//
// Solidity: function redeemStake(_reportingParticipants address[], _feeWindows address[]) returns(bool)
func (_Universe *UniverseSession) RedeemStake(_reportingParticipants []common.Address, _feeWindows []common.Address) (*types.Transaction, error) {
	return _Universe.Contract.RedeemStake(&_Universe.TransactOpts, _reportingParticipants, _feeWindows)
}

// RedeemStake is a paid mutator transaction binding the contract method 0x8d8cc634.
//
// Solidity: function redeemStake(_reportingParticipants address[], _feeWindows address[]) returns(bool)
func (_Universe *UniverseTransactorSession) RedeemStake(_reportingParticipants []common.Address, _feeWindows []common.Address) (*types.Transaction, error) {
	return _Universe.Contract.RedeemStake(&_Universe.TransactOpts, _reportingParticipants, _feeWindows)
}

// RemoveMarketFrom is a paid mutator transaction binding the contract method 0x77a3a0a2.
//
// Solidity: function removeMarketFrom() returns(bool)
func (_Universe *UniverseTransactor) RemoveMarketFrom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "removeMarketFrom")
}

// RemoveMarketFrom is a paid mutator transaction binding the contract method 0x77a3a0a2.
//
// Solidity: function removeMarketFrom() returns(bool)
func (_Universe *UniverseSession) RemoveMarketFrom() (*types.Transaction, error) {
	return _Universe.Contract.RemoveMarketFrom(&_Universe.TransactOpts)
}

// RemoveMarketFrom is a paid mutator transaction binding the contract method 0x77a3a0a2.
//
// Solidity: function removeMarketFrom() returns(bool)
func (_Universe *UniverseTransactorSession) RemoveMarketFrom() (*types.Transaction, error) {
	return _Universe.Contract.RemoveMarketFrom(&_Universe.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Universe *UniverseTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Universe *UniverseSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Universe.Contract.SetController(&_Universe.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Universe *UniverseTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Universe.Contract.SetController(&_Universe.TransactOpts, _controller)
}

// UpdateForkValues is a paid mutator transaction binding the contract method 0x9ab448d9.
//
// Solidity: function updateForkValues() returns(bool)
func (_Universe *UniverseTransactor) UpdateForkValues(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "updateForkValues")
}

// UpdateForkValues is a paid mutator transaction binding the contract method 0x9ab448d9.
//
// Solidity: function updateForkValues() returns(bool)
func (_Universe *UniverseSession) UpdateForkValues() (*types.Transaction, error) {
	return _Universe.Contract.UpdateForkValues(&_Universe.TransactOpts)
}

// UpdateForkValues is a paid mutator transaction binding the contract method 0x9ab448d9.
//
// Solidity: function updateForkValues() returns(bool)
func (_Universe *UniverseTransactorSession) UpdateForkValues() (*types.Transaction, error) {
	return _Universe.Contract.UpdateForkValues(&_Universe.TransactOpts)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(_parentPayoutDistributionHash bytes32) returns(bool)
func (_Universe *UniverseTransactor) UpdateTentativeWinningChildUniverse(opts *bind.TransactOpts, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "updateTentativeWinningChildUniverse", _parentPayoutDistributionHash)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(_parentPayoutDistributionHash bytes32) returns(bool)
func (_Universe *UniverseSession) UpdateTentativeWinningChildUniverse(_parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Universe.Contract.UpdateTentativeWinningChildUniverse(&_Universe.TransactOpts, _parentPayoutDistributionHash)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(_parentPayoutDistributionHash bytes32) returns(bool)
func (_Universe *UniverseTransactorSession) UpdateTentativeWinningChildUniverse(_parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Universe.Contract.UpdateTentativeWinningChildUniverse(&_Universe.TransactOpts, _parentPayoutDistributionHash)
}

// UniverseFactoryABI is the input ABI used to generate the binding from.
const UniverseFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_parentUniverse\",\"type\":\"address\"},{\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"createUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UniverseFactoryBin is the compiled bytecode used for deploying new contracts.
const UniverseFactoryBin = `0x608060405234801561001057600080fd5b5061050c806100206000396000f3006080604052600436106100405763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634837435f8114610045575b600080fd5b34801561005157600080fd5b5061007c73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356100a5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000806000856100b36101c8565b73ffffffffffffffffffffffffffffffffffffffff90911681527f556e6976657273650000000000000000000000000000000000000000000000006020820152604080519182900301906000f080158015610112573d6000803e3d6000fd5b50604080517fbe13f47c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff88811660048301526024820188905291519294508493509083169163be13f47c916044808201926020929091908290030181600087803b15801561019257600080fd5b505af11580156101a6573d6000803e3d6000fd5b505050506040513d60208110156101bc57600080fd5b50909695505050505050565b604051610308806101d9833901905600608060405234801561001057600080fd5b506040516040806103088339810160405280516020909101516000805433600160a060020a03199182161716600160a060020a039093169290921782556001556102a890819061006090396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461014f57806392eefe9b1461018d578063bef72fa2146101cf575b60015460009015156100675761014c565b60008054600154604080517ff39ec1f700000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff9092169263f39ec1f7926024808401936020939083900390910190829087803b1580156100df57600080fd5b505af11580156100f3573d6000803e3d6000fd5b505050506040513d602081101561010957600080fd5b50516040805136601f8101601f191680830190935292935091600083376000803684865af4808015610147576040513d81016040523d6000823e3d81f35b600080fd5b50005b34801561015b57600080fd5b506101646101f6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019957600080fd5b506101bb73ffffffffffffffffffffffffffffffffffffffff60043516610212565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610276565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461023757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a7230582017c3a6b57ac8bbc7fd5320cba2543fe371eb8726332cd58700252790e09cf2140029a165627a7a723058203bb9860d0c4317eb360a22ec936677eafef29be3761a4630c143b2bf3d57b9190029`

// DeployUniverseFactory deploys a new Ethereum contract, binding an instance of UniverseFactory to it.
func DeployUniverseFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniverseFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(UniverseFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UniverseFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniverseFactory{UniverseFactoryCaller: UniverseFactoryCaller{contract: contract}, UniverseFactoryTransactor: UniverseFactoryTransactor{contract: contract}, UniverseFactoryFilterer: UniverseFactoryFilterer{contract: contract}}, nil
}

// UniverseFactory is an auto generated Go binding around an Ethereum contract.
type UniverseFactory struct {
	UniverseFactoryCaller     // Read-only binding to the contract
	UniverseFactoryTransactor // Write-only binding to the contract
	UniverseFactoryFilterer   // Log filterer for contract events
}

// UniverseFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniverseFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniverseFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniverseFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniverseFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniverseFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniverseFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniverseFactorySession struct {
	Contract     *UniverseFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniverseFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniverseFactoryCallerSession struct {
	Contract *UniverseFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniverseFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniverseFactoryTransactorSession struct {
	Contract     *UniverseFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniverseFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniverseFactoryRaw struct {
	Contract *UniverseFactory // Generic contract binding to access the raw methods on
}

// UniverseFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniverseFactoryCallerRaw struct {
	Contract *UniverseFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UniverseFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniverseFactoryTransactorRaw struct {
	Contract *UniverseFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniverseFactory creates a new instance of UniverseFactory, bound to a specific deployed contract.
func NewUniverseFactory(address common.Address, backend bind.ContractBackend) (*UniverseFactory, error) {
	contract, err := bindUniverseFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniverseFactory{UniverseFactoryCaller: UniverseFactoryCaller{contract: contract}, UniverseFactoryTransactor: UniverseFactoryTransactor{contract: contract}, UniverseFactoryFilterer: UniverseFactoryFilterer{contract: contract}}, nil
}

// NewUniverseFactoryCaller creates a new read-only instance of UniverseFactory, bound to a specific deployed contract.
func NewUniverseFactoryCaller(address common.Address, caller bind.ContractCaller) (*UniverseFactoryCaller, error) {
	contract, err := bindUniverseFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniverseFactoryCaller{contract: contract}, nil
}

// NewUniverseFactoryTransactor creates a new write-only instance of UniverseFactory, bound to a specific deployed contract.
func NewUniverseFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniverseFactoryTransactor, error) {
	contract, err := bindUniverseFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniverseFactoryTransactor{contract: contract}, nil
}

// NewUniverseFactoryFilterer creates a new log filterer instance of UniverseFactory, bound to a specific deployed contract.
func NewUniverseFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniverseFactoryFilterer, error) {
	contract, err := bindUniverseFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniverseFactoryFilterer{contract: contract}, nil
}

// bindUniverseFactory binds a generic wrapper to an already deployed contract.
func bindUniverseFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniverseFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniverseFactory *UniverseFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UniverseFactory.Contract.UniverseFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniverseFactory *UniverseFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniverseFactory.Contract.UniverseFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniverseFactory *UniverseFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniverseFactory.Contract.UniverseFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniverseFactory *UniverseFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UniverseFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniverseFactory *UniverseFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniverseFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniverseFactory *UniverseFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniverseFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateUniverse is a paid mutator transaction binding the contract method 0x4837435f.
//
// Solidity: function createUniverse(_controller address, _parentUniverse address, _parentPayoutDistributionHash bytes32) returns(address)
func (_UniverseFactory *UniverseFactoryTransactor) CreateUniverse(opts *bind.TransactOpts, _controller common.Address, _parentUniverse common.Address, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _UniverseFactory.contract.Transact(opts, "createUniverse", _controller, _parentUniverse, _parentPayoutDistributionHash)
}

// CreateUniverse is a paid mutator transaction binding the contract method 0x4837435f.
//
// Solidity: function createUniverse(_controller address, _parentUniverse address, _parentPayoutDistributionHash bytes32) returns(address)
func (_UniverseFactory *UniverseFactorySession) CreateUniverse(_controller common.Address, _parentUniverse common.Address, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _UniverseFactory.Contract.CreateUniverse(&_UniverseFactory.TransactOpts, _controller, _parentUniverse, _parentPayoutDistributionHash)
}

// CreateUniverse is a paid mutator transaction binding the contract method 0x4837435f.
//
// Solidity: function createUniverse(_controller address, _parentUniverse address, _parentPayoutDistributionHash bytes32) returns(address)
func (_UniverseFactory *UniverseFactoryTransactorSession) CreateUniverse(_controller common.Address, _parentUniverse common.Address, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _UniverseFactory.Contract.CreateUniverse(&_UniverseFactory.TransactOpts, _controller, _parentUniverse, _parentPayoutDistributionHash)
}