package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IInitialReporterABI is the input ABI used to generate the binding from.
const IInitialReporterABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isInvalid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"liquidateLosing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"designatedReporterShowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_designatedReporter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDesignatedReporter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isDisavowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"returnRepFromDisavow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"designatedReporterWasCorrect\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReportTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_redeemer\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_payoutDistributionHash\",\"type\":\"bytes32\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_invalid\",\"type\":\"bool\"},{\"name\":\"_initialReportStake\",\"type\":\"uint256\"}],\"name\":\"report\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getPayoutNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPayoutDistributionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_designatedReporter\",\"type\":\"address\"}],\"name\":\"migrateToNewUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IInitialReporterBin is the compiled bytecode used for deploying new contracts.
const IInitialReporterBin = `0x`

// DeployIInitialReporter deploys a new Ethereum contract, binding an instance of IInitialReporter to it.
func DeployIInitialReporter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IInitialReporter, error) {
	parsed, err := abi.JSON(strings.NewReader(IInitialReporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IInitialReporterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IInitialReporter{IInitialReporterCaller: IInitialReporterCaller{contract: contract}, IInitialReporterTransactor: IInitialReporterTransactor{contract: contract}, IInitialReporterFilterer: IInitialReporterFilterer{contract: contract}}, nil
}

// IInitialReporter is an auto generated Go binding around an Ethereum contract.
type IInitialReporter struct {
	IInitialReporterCaller     // Read-only binding to the contract
	IInitialReporterTransactor // Write-only binding to the contract
	IInitialReporterFilterer   // Log filterer for contract events
}

// IInitialReporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInitialReporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInitialReporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInitialReporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInitialReporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInitialReporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInitialReporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInitialReporterSession struct {
	Contract     *IInitialReporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInitialReporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInitialReporterCallerSession struct {
	Contract *IInitialReporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IInitialReporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInitialReporterTransactorSession struct {
	Contract     *IInitialReporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IInitialReporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInitialReporterRaw struct {
	Contract *IInitialReporter // Generic contract binding to access the raw methods on
}

// IInitialReporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInitialReporterCallerRaw struct {
	Contract *IInitialReporterCaller // Generic read-only contract binding to access the raw methods on
}

// IInitialReporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInitialReporterTransactorRaw struct {
	Contract *IInitialReporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInitialReporter creates a new instance of IInitialReporter, bound to a specific deployed contract.
func NewIInitialReporter(address common.Address, backend bind.ContractBackend) (*IInitialReporter, error) {
	contract, err := bindIInitialReporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInitialReporter{IInitialReporterCaller: IInitialReporterCaller{contract: contract}, IInitialReporterTransactor: IInitialReporterTransactor{contract: contract}, IInitialReporterFilterer: IInitialReporterFilterer{contract: contract}}, nil
}

// NewIInitialReporterCaller creates a new read-only instance of IInitialReporter, bound to a specific deployed contract.
func NewIInitialReporterCaller(address common.Address, caller bind.ContractCaller) (*IInitialReporterCaller, error) {
	contract, err := bindIInitialReporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInitialReporterCaller{contract: contract}, nil
}

// NewIInitialReporterTransactor creates a new write-only instance of IInitialReporter, bound to a specific deployed contract.
func NewIInitialReporterTransactor(address common.Address, transactor bind.ContractTransactor) (*IInitialReporterTransactor, error) {
	contract, err := bindIInitialReporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInitialReporterTransactor{contract: contract}, nil
}

// NewIInitialReporterFilterer creates a new log filterer instance of IInitialReporter, bound to a specific deployed contract.
func NewIInitialReporterFilterer(address common.Address, filterer bind.ContractFilterer) (*IInitialReporterFilterer, error) {
	contract, err := bindIInitialReporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInitialReporterFilterer{contract: contract}, nil
}

// bindIInitialReporter binds a generic wrapper to an already deployed contract.
func bindIInitialReporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInitialReporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInitialReporter *IInitialReporterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IInitialReporter.Contract.IInitialReporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInitialReporter *IInitialReporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.Contract.IInitialReporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInitialReporter *IInitialReporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInitialReporter.Contract.IInitialReporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInitialReporter *IInitialReporterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IInitialReporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInitialReporter *IInitialReporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInitialReporter *IInitialReporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInitialReporter.Contract.contract.Transact(opts, method, params...)
}

// DesignatedReporterShowed is a free data retrieval call binding the contract method 0x3bf8f34a.
//
// Solidity: function designatedReporterShowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterCaller) DesignatedReporterShowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "designatedReporterShowed")
	return *ret0, err
}

// DesignatedReporterShowed is a free data retrieval call binding the contract method 0x3bf8f34a.
//
// Solidity: function designatedReporterShowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterSession) DesignatedReporterShowed() (bool, error) {
	return _IInitialReporter.Contract.DesignatedReporterShowed(&_IInitialReporter.CallOpts)
}

// DesignatedReporterShowed is a free data retrieval call binding the contract method 0x3bf8f34a.
//
// Solidity: function designatedReporterShowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterCallerSession) DesignatedReporterShowed() (bool, error) {
	return _IInitialReporter.Contract.DesignatedReporterShowed(&_IInitialReporter.CallOpts)
}

// DesignatedReporterWasCorrect is a free data retrieval call binding the contract method 0x8ed882c5.
//
// Solidity: function designatedReporterWasCorrect() constant returns(bool)
func (_IInitialReporter *IInitialReporterCaller) DesignatedReporterWasCorrect(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "designatedReporterWasCorrect")
	return *ret0, err
}

// DesignatedReporterWasCorrect is a free data retrieval call binding the contract method 0x8ed882c5.
//
// Solidity: function designatedReporterWasCorrect() constant returns(bool)
func (_IInitialReporter *IInitialReporterSession) DesignatedReporterWasCorrect() (bool, error) {
	return _IInitialReporter.Contract.DesignatedReporterWasCorrect(&_IInitialReporter.CallOpts)
}

// DesignatedReporterWasCorrect is a free data retrieval call binding the contract method 0x8ed882c5.
//
// Solidity: function designatedReporterWasCorrect() constant returns(bool)
func (_IInitialReporter *IInitialReporterCallerSession) DesignatedReporterWasCorrect() (bool, error) {
	return _IInitialReporter.Contract.DesignatedReporterWasCorrect(&_IInitialReporter.CallOpts)
}

// GetDesignatedReporter is a free data retrieval call binding the contract method 0x5b6e2492.
//
// Solidity: function getDesignatedReporter() constant returns(address)
func (_IInitialReporter *IInitialReporterCaller) GetDesignatedReporter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getDesignatedReporter")
	return *ret0, err
}

// GetDesignatedReporter is a free data retrieval call binding the contract method 0x5b6e2492.
//
// Solidity: function getDesignatedReporter() constant returns(address)
func (_IInitialReporter *IInitialReporterSession) GetDesignatedReporter() (common.Address, error) {
	return _IInitialReporter.Contract.GetDesignatedReporter(&_IInitialReporter.CallOpts)
}

// GetDesignatedReporter is a free data retrieval call binding the contract method 0x5b6e2492.
//
// Solidity: function getDesignatedReporter() constant returns(address)
func (_IInitialReporter *IInitialReporterCallerSession) GetDesignatedReporter() (common.Address, error) {
	return _IInitialReporter.Contract.GetDesignatedReporter(&_IInitialReporter.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IInitialReporter *IInitialReporterCaller) GetMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getMarket")
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IInitialReporter *IInitialReporterSession) GetMarket() (common.Address, error) {
	return _IInitialReporter.Contract.GetMarket(&_IInitialReporter.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IInitialReporter *IInitialReporterCallerSession) GetMarket() (common.Address, error) {
	return _IInitialReporter.Contract.GetMarket(&_IInitialReporter.CallOpts)
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_IInitialReporter *IInitialReporterCaller) GetPayoutDistributionHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getPayoutDistributionHash")
	return *ret0, err
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_IInitialReporter *IInitialReporterSession) GetPayoutDistributionHash() ([32]byte, error) {
	return _IInitialReporter.Contract.GetPayoutDistributionHash(&_IInitialReporter.CallOpts)
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_IInitialReporter *IInitialReporterCallerSession) GetPayoutDistributionHash() ([32]byte, error) {
	return _IInitialReporter.Contract.GetPayoutDistributionHash(&_IInitialReporter.CallOpts)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_IInitialReporter *IInitialReporterCaller) GetPayoutNumerator(opts *bind.CallOpts, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getPayoutNumerator", _outcome)
	return *ret0, err
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_IInitialReporter *IInitialReporterSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _IInitialReporter.Contract.GetPayoutNumerator(&_IInitialReporter.CallOpts, _outcome)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_IInitialReporter *IInitialReporterCallerSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _IInitialReporter.Contract.GetPayoutNumerator(&_IInitialReporter.CallOpts, _outcome)
}

// GetReportTimestamp is a free data retrieval call binding the contract method 0x94a771d7.
//
// Solidity: function getReportTimestamp() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCaller) GetReportTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getReportTimestamp")
	return *ret0, err
}

// GetReportTimestamp is a free data retrieval call binding the contract method 0x94a771d7.
//
// Solidity: function getReportTimestamp() constant returns(uint256)
func (_IInitialReporter *IInitialReporterSession) GetReportTimestamp() (*big.Int, error) {
	return _IInitialReporter.Contract.GetReportTimestamp(&_IInitialReporter.CallOpts)
}

// GetReportTimestamp is a free data retrieval call binding the contract method 0x94a771d7.
//
// Solidity: function getReportTimestamp() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCallerSession) GetReportTimestamp() (*big.Int, error) {
	return _IInitialReporter.Contract.GetReportTimestamp(&_IInitialReporter.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_IInitialReporter *IInitialReporterSession) GetSize() (*big.Int, error) {
	return _IInitialReporter.Contract.GetSize(&_IInitialReporter.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCallerSession) GetSize() (*big.Int, error) {
	return _IInitialReporter.Contract.GetSize(&_IInitialReporter.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCaller) GetStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getStake")
	return *ret0, err
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_IInitialReporter *IInitialReporterSession) GetStake() (*big.Int, error) {
	return _IInitialReporter.Contract.GetStake(&_IInitialReporter.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCallerSession) GetStake() (*big.Int, error) {
	return _IInitialReporter.Contract.GetStake(&_IInitialReporter.CallOpts)
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterCaller) IsDisavowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "isDisavowed")
	return *ret0, err
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterSession) IsDisavowed() (bool, error) {
	return _IInitialReporter.Contract.IsDisavowed(&_IInitialReporter.CallOpts)
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterCallerSession) IsDisavowed() (bool, error) {
	return _IInitialReporter.Contract.IsDisavowed(&_IInitialReporter.CallOpts)
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_IInitialReporter *IInitialReporterCaller) IsInvalid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "isInvalid")
	return *ret0, err
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_IInitialReporter *IInitialReporterSession) IsInvalid() (bool, error) {
	return _IInitialReporter.Contract.IsInvalid(&_IInitialReporter.CallOpts)
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_IInitialReporter *IInitialReporterCallerSession) IsInvalid() (bool, error) {
	return _IInitialReporter.Contract.IsInvalid(&_IInitialReporter.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_market address, _designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) Initialize(opts *bind.TransactOpts, _market common.Address, _designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "initialize", _market, _designatedReporter)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_market address, _designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterSession) Initialize(_market common.Address, _designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Initialize(&_IInitialReporter.TransactOpts, _market, _designatedReporter)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_market address, _designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) Initialize(_market common.Address, _designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Initialize(&_IInitialReporter.TransactOpts, _market, _designatedReporter)
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) LiquidateLosing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "liquidateLosing")
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_IInitialReporter *IInitialReporterSession) LiquidateLosing() (*types.Transaction, error) {
	return _IInitialReporter.Contract.LiquidateLosing(&_IInitialReporter.TransactOpts)
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) LiquidateLosing() (*types.Transaction, error) {
	return _IInitialReporter.Contract.LiquidateLosing(&_IInitialReporter.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_IInitialReporter *IInitialReporterSession) Migrate() (*types.Transaction, error) {
	return _IInitialReporter.Contract.Migrate(&_IInitialReporter.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) Migrate() (*types.Transaction, error) {
	return _IInitialReporter.Contract.Migrate(&_IInitialReporter.TransactOpts)
}

// MigrateToNewUniverse is a paid mutator transaction binding the contract method 0xef0ce37a.
//
// Solidity: function migrateToNewUniverse(_designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) MigrateToNewUniverse(opts *bind.TransactOpts, _designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "migrateToNewUniverse", _designatedReporter)
}

// MigrateToNewUniverse is a paid mutator transaction binding the contract method 0xef0ce37a.
//
// Solidity: function migrateToNewUniverse(_designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterSession) MigrateToNewUniverse(_designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.MigrateToNewUniverse(&_IInitialReporter.TransactOpts, _designatedReporter)
}

// MigrateToNewUniverse is a paid mutator transaction binding the contract method 0xef0ce37a.
//
// Solidity: function migrateToNewUniverse(_designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) MigrateToNewUniverse(_designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.MigrateToNewUniverse(&_IInitialReporter.TransactOpts, _designatedReporter)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) Redeem(opts *bind.TransactOpts, _redeemer common.Address) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "redeem", _redeemer)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_IInitialReporter *IInitialReporterSession) Redeem(_redeemer common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Redeem(&_IInitialReporter.TransactOpts, _redeemer)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) Redeem(_redeemer common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Redeem(&_IInitialReporter.TransactOpts, _redeemer)
}

// Report is a paid mutator transaction binding the contract method 0xc6f344f5.
//
// Solidity: function report(_reporter address, _payoutDistributionHash bytes32, _payoutNumerators uint256[], _invalid bool, _initialReportStake uint256) returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) Report(opts *bind.TransactOpts, _reporter common.Address, _payoutDistributionHash [32]byte, _payoutNumerators []*big.Int, _invalid bool, _initialReportStake *big.Int) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "report", _reporter, _payoutDistributionHash, _payoutNumerators, _invalid, _initialReportStake)
}

// Report is a paid mutator transaction binding the contract method 0xc6f344f5.
//
// Solidity: function report(_reporter address, _payoutDistributionHash bytes32, _payoutNumerators uint256[], _invalid bool, _initialReportStake uint256) returns(bool)
func (_IInitialReporter *IInitialReporterSession) Report(_reporter common.Address, _payoutDistributionHash [32]byte, _payoutNumerators []*big.Int, _invalid bool, _initialReportStake *big.Int) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Report(&_IInitialReporter.TransactOpts, _reporter, _payoutDistributionHash, _payoutNumerators, _invalid, _initialReportStake)
}

// Report is a paid mutator transaction binding the contract method 0xc6f344f5.
//
// Solidity: function report(_reporter address, _payoutDistributionHash bytes32, _payoutNumerators uint256[], _invalid bool, _initialReportStake uint256) returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) Report(_reporter common.Address, _payoutDistributionHash [32]byte, _payoutNumerators []*big.Int, _invalid bool, _initialReportStake *big.Int) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Report(&_IInitialReporter.TransactOpts, _reporter, _payoutDistributionHash, _payoutNumerators, _invalid, _initialReportStake)
}

// ReturnRepFromDisavow is a paid mutator transaction binding the contract method 0x87882cb8.
//
// Solidity: function returnRepFromDisavow() returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) ReturnRepFromDisavow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "returnRepFromDisavow")
}

// ReturnRepFromDisavow is a paid mutator transaction binding the contract method 0x87882cb8.
//
// Solidity: function returnRepFromDisavow() returns(bool)
func (_IInitialReporter *IInitialReporterSession) ReturnRepFromDisavow() (*types.Transaction, error) {
	return _IInitialReporter.Contract.ReturnRepFromDisavow(&_IInitialReporter.TransactOpts)
}

// ReturnRepFromDisavow is a paid mutator transaction binding the contract method 0x87882cb8.
//
// Solidity: function returnRepFromDisavow() returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) ReturnRepFromDisavow() (*types.Transaction, error) {
	return _IInitialReporter.Contract.ReturnRepFromDisavow(&_IInitialReporter.TransactOpts)
}

// InitialReporterFactoryABI is the input ABI used to generate the binding from.
const InitialReporterFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_designatedReporter\",\"type\":\"address\"}],\"name\":\"createInitialReporter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InitialReporterFactoryBin is the compiled bytecode used for deploying new contracts.
const InitialReporterFactoryBin = `0x608060405234801561001057600080fd5b50610510806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663b30949d98114610045575b600080fd5b34801561005157600080fd5b5061007f73ffffffffffffffffffffffffffffffffffffffff600435811690602435811690604435166100a8565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000806000856100b66101cc565b73ffffffffffffffffffffffffffffffffffffffff90911681527f496e697469616c5265706f7274657200000000000000000000000000000000006020820152604080519182900301906000f080158015610115573d6000803e3d6000fd5b50604080517f485cc95500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152878116602483015291519294508493509083169163485cc955916044808201926020929091908290030181600087803b15801561019657600080fd5b505af11580156101aa573d6000803e3d6000fd5b505050506040513d60208110156101c057600080fd5b50909695505050505050565b604051610308806101dd833901905600608060405234801561001057600080fd5b506040516040806103088339810160405280516020909101516000805433600160a060020a03199182161716600160a060020a039093169290921782556001556102a890819061006090396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461014f57806392eefe9b1461018d578063bef72fa2146101cf575b60015460009015156100675761014c565b60008054600154604080517ff39ec1f700000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff9092169263f39ec1f7926024808401936020939083900390910190829087803b1580156100df57600080fd5b505af11580156100f3573d6000803e3d6000fd5b505050506040513d602081101561010957600080fd5b50516040805136601f8101601f191680830190935292935091600083376000803684865af4808015610147576040513d81016040523d6000823e3d81f35b600080fd5b50005b34801561015b57600080fd5b506101646101f6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019957600080fd5b506101bb73ffffffffffffffffffffffffffffffffffffffff60043516610212565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610276565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461023757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a7230582017c3a6b57ac8bbc7fd5320cba2543fe371eb8726332cd58700252790e09cf2140029a165627a7a72305820d253873ca38b6b638f84a97baf0cbebb58b93b5d1167fedc10735bbf40096dc30029`

// DeployInitialReporterFactory deploys a new Ethereum contract, binding an instance of InitialReporterFactory to it.
func DeployInitialReporterFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InitialReporterFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(InitialReporterFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InitialReporterFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InitialReporterFactory{InitialReporterFactoryCaller: InitialReporterFactoryCaller{contract: contract}, InitialReporterFactoryTransactor: InitialReporterFactoryTransactor{contract: contract}, InitialReporterFactoryFilterer: InitialReporterFactoryFilterer{contract: contract}}, nil
}

// InitialReporterFactory is an auto generated Go binding around an Ethereum contract.
type InitialReporterFactory struct {
	InitialReporterFactoryCaller     // Read-only binding to the contract
	InitialReporterFactoryTransactor // Write-only binding to the contract
	InitialReporterFactoryFilterer   // Log filterer for contract events
}

// InitialReporterFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitialReporterFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitialReporterFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitialReporterFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitialReporterFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitialReporterFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitialReporterFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitialReporterFactorySession struct {
	Contract     *InitialReporterFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InitialReporterFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitialReporterFactoryCallerSession struct {
	Contract *InitialReporterFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// InitialReporterFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitialReporterFactoryTransactorSession struct {
	Contract     *InitialReporterFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// InitialReporterFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitialReporterFactoryRaw struct {
	Contract *InitialReporterFactory // Generic contract binding to access the raw methods on
}

// InitialReporterFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitialReporterFactoryCallerRaw struct {
	Contract *InitialReporterFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// InitialReporterFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitialReporterFactoryTransactorRaw struct {
	Contract *InitialReporterFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitialReporterFactory creates a new instance of InitialReporterFactory, bound to a specific deployed contract.
func NewInitialReporterFactory(address common.Address, backend bind.ContractBackend) (*InitialReporterFactory, error) {
	contract, err := bindInitialReporterFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InitialReporterFactory{InitialReporterFactoryCaller: InitialReporterFactoryCaller{contract: contract}, InitialReporterFactoryTransactor: InitialReporterFactoryTransactor{contract: contract}, InitialReporterFactoryFilterer: InitialReporterFactoryFilterer{contract: contract}}, nil
}

// NewInitialReporterFactoryCaller creates a new read-only instance of InitialReporterFactory, bound to a specific deployed contract.
func NewInitialReporterFactoryCaller(address common.Address, caller bind.ContractCaller) (*InitialReporterFactoryCaller, error) {
	contract, err := bindInitialReporterFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitialReporterFactoryCaller{contract: contract}, nil
}

// NewInitialReporterFactoryTransactor creates a new write-only instance of InitialReporterFactory, bound to a specific deployed contract.
func NewInitialReporterFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*InitialReporterFactoryTransactor, error) {
	contract, err := bindInitialReporterFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitialReporterFactoryTransactor{contract: contract}, nil
}

// NewInitialReporterFactoryFilterer creates a new log filterer instance of InitialReporterFactory, bound to a specific deployed contract.
func NewInitialReporterFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*InitialReporterFactoryFilterer, error) {
	contract, err := bindInitialReporterFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitialReporterFactoryFilterer{contract: contract}, nil
}

// bindInitialReporterFactory binds a generic wrapper to an already deployed contract.
func bindInitialReporterFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitialReporterFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitialReporterFactory *InitialReporterFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InitialReporterFactory.Contract.InitialReporterFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitialReporterFactory *InitialReporterFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitialReporterFactory.Contract.InitialReporterFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitialReporterFactory *InitialReporterFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitialReporterFactory.Contract.InitialReporterFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitialReporterFactory *InitialReporterFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InitialReporterFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitialReporterFactory *InitialReporterFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitialReporterFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitialReporterFactory *InitialReporterFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitialReporterFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateInitialReporter is a paid mutator transaction binding the contract method 0xb30949d9.
//
// Solidity: function createInitialReporter(_controller address, _market address, _designatedReporter address) returns(address)
func (_InitialReporterFactory *InitialReporterFactoryTransactor) CreateInitialReporter(opts *bind.TransactOpts, _controller common.Address, _market common.Address, _designatedReporter common.Address) (*types.Transaction, error) {
	return _InitialReporterFactory.contract.Transact(opts, "createInitialReporter", _controller, _market, _designatedReporter)
}

// CreateInitialReporter is a paid mutator transaction binding the contract method 0xb30949d9.
//
// Solidity: function createInitialReporter(_controller address, _market address, _designatedReporter address) returns(address)
func (_InitialReporterFactory *InitialReporterFactorySession) CreateInitialReporter(_controller common.Address, _market common.Address, _designatedReporter common.Address) (*types.Transaction, error) {
	return _InitialReporterFactory.Contract.CreateInitialReporter(&_InitialReporterFactory.TransactOpts, _controller, _market, _designatedReporter)
}

// CreateInitialReporter is a paid mutator transaction binding the contract method 0xb30949d9.
//
// Solidity: function createInitialReporter(_controller address, _market address, _designatedReporter address) returns(address)
func (_InitialReporterFactory *InitialReporterFactoryTransactorSession) CreateInitialReporter(_controller common.Address, _market common.Address, _designatedReporter common.Address) (*types.Transaction, error) {
	return _InitialReporterFactory.Contract.CreateInitialReporter(&_InitialReporterFactory.TransactOpts, _controller, _market, _designatedReporter)
}

// IInitialReporterABI is the input ABI used to generate the binding from.
const IInitialReporterABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isInvalid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"liquidateLosing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"designatedReporterShowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_designatedReporter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDesignatedReporter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isDisavowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"returnRepFromDisavow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"designatedReporterWasCorrect\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReportTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_redeemer\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_payoutDistributionHash\",\"type\":\"bytes32\"},{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_invalid\",\"type\":\"bool\"},{\"name\":\"_initialReportStake\",\"type\":\"uint256\"}],\"name\":\"report\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getPayoutNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPayoutDistributionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_designatedReporter\",\"type\":\"address\"}],\"name\":\"migrateToNewUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IInitialReporterBin is the compiled bytecode used for deploying new contracts.
const IInitialReporterBin = `0x`

// DeployIInitialReporter deploys a new Ethereum contract, binding an instance of IInitialReporter to it.
func DeployIInitialReporter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IInitialReporter, error) {
	parsed, err := abi.JSON(strings.NewReader(IInitialReporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IInitialReporterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IInitialReporter{IInitialReporterCaller: IInitialReporterCaller{contract: contract}, IInitialReporterTransactor: IInitialReporterTransactor{contract: contract}, IInitialReporterFilterer: IInitialReporterFilterer{contract: contract}}, nil
}

// IInitialReporter is an auto generated Go binding around an Ethereum contract.
type IInitialReporter struct {
	IInitialReporterCaller     // Read-only binding to the contract
	IInitialReporterTransactor // Write-only binding to the contract
	IInitialReporterFilterer   // Log filterer for contract events
}

// IInitialReporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInitialReporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInitialReporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInitialReporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInitialReporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInitialReporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInitialReporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInitialReporterSession struct {
	Contract     *IInitialReporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInitialReporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInitialReporterCallerSession struct {
	Contract *IInitialReporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IInitialReporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInitialReporterTransactorSession struct {
	Contract     *IInitialReporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IInitialReporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInitialReporterRaw struct {
	Contract *IInitialReporter // Generic contract binding to access the raw methods on
}

// IInitialReporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInitialReporterCallerRaw struct {
	Contract *IInitialReporterCaller // Generic read-only contract binding to access the raw methods on
}

// IInitialReporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInitialReporterTransactorRaw struct {
	Contract *IInitialReporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInitialReporter creates a new instance of IInitialReporter, bound to a specific deployed contract.
func NewIInitialReporter(address common.Address, backend bind.ContractBackend) (*IInitialReporter, error) {
	contract, err := bindIInitialReporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInitialReporter{IInitialReporterCaller: IInitialReporterCaller{contract: contract}, IInitialReporterTransactor: IInitialReporterTransactor{contract: contract}, IInitialReporterFilterer: IInitialReporterFilterer{contract: contract}}, nil
}

// NewIInitialReporterCaller creates a new read-only instance of IInitialReporter, bound to a specific deployed contract.
func NewIInitialReporterCaller(address common.Address, caller bind.ContractCaller) (*IInitialReporterCaller, error) {
	contract, err := bindIInitialReporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInitialReporterCaller{contract: contract}, nil
}

// NewIInitialReporterTransactor creates a new write-only instance of IInitialReporter, bound to a specific deployed contract.
func NewIInitialReporterTransactor(address common.Address, transactor bind.ContractTransactor) (*IInitialReporterTransactor, error) {
	contract, err := bindIInitialReporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInitialReporterTransactor{contract: contract}, nil
}

// NewIInitialReporterFilterer creates a new log filterer instance of IInitialReporter, bound to a specific deployed contract.
func NewIInitialReporterFilterer(address common.Address, filterer bind.ContractFilterer) (*IInitialReporterFilterer, error) {
	contract, err := bindIInitialReporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInitialReporterFilterer{contract: contract}, nil
}

// bindIInitialReporter binds a generic wrapper to an already deployed contract.
func bindIInitialReporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInitialReporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInitialReporter *IInitialReporterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IInitialReporter.Contract.IInitialReporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInitialReporter *IInitialReporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.Contract.IInitialReporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInitialReporter *IInitialReporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInitialReporter.Contract.IInitialReporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInitialReporter *IInitialReporterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IInitialReporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInitialReporter *IInitialReporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInitialReporter *IInitialReporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInitialReporter.Contract.contract.Transact(opts, method, params...)
}

// DesignatedReporterShowed is a free data retrieval call binding the contract method 0x3bf8f34a.
//
// Solidity: function designatedReporterShowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterCaller) DesignatedReporterShowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "designatedReporterShowed")
	return *ret0, err
}

// DesignatedReporterShowed is a free data retrieval call binding the contract method 0x3bf8f34a.
//
// Solidity: function designatedReporterShowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterSession) DesignatedReporterShowed() (bool, error) {
	return _IInitialReporter.Contract.DesignatedReporterShowed(&_IInitialReporter.CallOpts)
}

// DesignatedReporterShowed is a free data retrieval call binding the contract method 0x3bf8f34a.
//
// Solidity: function designatedReporterShowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterCallerSession) DesignatedReporterShowed() (bool, error) {
	return _IInitialReporter.Contract.DesignatedReporterShowed(&_IInitialReporter.CallOpts)
}

// DesignatedReporterWasCorrect is a free data retrieval call binding the contract method 0x8ed882c5.
//
// Solidity: function designatedReporterWasCorrect() constant returns(bool)
func (_IInitialReporter *IInitialReporterCaller) DesignatedReporterWasCorrect(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "designatedReporterWasCorrect")
	return *ret0, err
}

// DesignatedReporterWasCorrect is a free data retrieval call binding the contract method 0x8ed882c5.
//
// Solidity: function designatedReporterWasCorrect() constant returns(bool)
func (_IInitialReporter *IInitialReporterSession) DesignatedReporterWasCorrect() (bool, error) {
	return _IInitialReporter.Contract.DesignatedReporterWasCorrect(&_IInitialReporter.CallOpts)
}

// DesignatedReporterWasCorrect is a free data retrieval call binding the contract method 0x8ed882c5.
//
// Solidity: function designatedReporterWasCorrect() constant returns(bool)
func (_IInitialReporter *IInitialReporterCallerSession) DesignatedReporterWasCorrect() (bool, error) {
	return _IInitialReporter.Contract.DesignatedReporterWasCorrect(&_IInitialReporter.CallOpts)
}

// GetDesignatedReporter is a free data retrieval call binding the contract method 0x5b6e2492.
//
// Solidity: function getDesignatedReporter() constant returns(address)
func (_IInitialReporter *IInitialReporterCaller) GetDesignatedReporter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getDesignatedReporter")
	return *ret0, err
}

// GetDesignatedReporter is a free data retrieval call binding the contract method 0x5b6e2492.
//
// Solidity: function getDesignatedReporter() constant returns(address)
func (_IInitialReporter *IInitialReporterSession) GetDesignatedReporter() (common.Address, error) {
	return _IInitialReporter.Contract.GetDesignatedReporter(&_IInitialReporter.CallOpts)
}

// GetDesignatedReporter is a free data retrieval call binding the contract method 0x5b6e2492.
//
// Solidity: function getDesignatedReporter() constant returns(address)
func (_IInitialReporter *IInitialReporterCallerSession) GetDesignatedReporter() (common.Address, error) {
	return _IInitialReporter.Contract.GetDesignatedReporter(&_IInitialReporter.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IInitialReporter *IInitialReporterCaller) GetMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getMarket")
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IInitialReporter *IInitialReporterSession) GetMarket() (common.Address, error) {
	return _IInitialReporter.Contract.GetMarket(&_IInitialReporter.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IInitialReporter *IInitialReporterCallerSession) GetMarket() (common.Address, error) {
	return _IInitialReporter.Contract.GetMarket(&_IInitialReporter.CallOpts)
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_IInitialReporter *IInitialReporterCaller) GetPayoutDistributionHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getPayoutDistributionHash")
	return *ret0, err
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_IInitialReporter *IInitialReporterSession) GetPayoutDistributionHash() ([32]byte, error) {
	return _IInitialReporter.Contract.GetPayoutDistributionHash(&_IInitialReporter.CallOpts)
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_IInitialReporter *IInitialReporterCallerSession) GetPayoutDistributionHash() ([32]byte, error) {
	return _IInitialReporter.Contract.GetPayoutDistributionHash(&_IInitialReporter.CallOpts)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_IInitialReporter *IInitialReporterCaller) GetPayoutNumerator(opts *bind.CallOpts, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getPayoutNumerator", _outcome)
	return *ret0, err
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_IInitialReporter *IInitialReporterSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _IInitialReporter.Contract.GetPayoutNumerator(&_IInitialReporter.CallOpts, _outcome)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_IInitialReporter *IInitialReporterCallerSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _IInitialReporter.Contract.GetPayoutNumerator(&_IInitialReporter.CallOpts, _outcome)
}

// GetReportTimestamp is a free data retrieval call binding the contract method 0x94a771d7.
//
// Solidity: function getReportTimestamp() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCaller) GetReportTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getReportTimestamp")
	return *ret0, err
}

// GetReportTimestamp is a free data retrieval call binding the contract method 0x94a771d7.
//
// Solidity: function getReportTimestamp() constant returns(uint256)
func (_IInitialReporter *IInitialReporterSession) GetReportTimestamp() (*big.Int, error) {
	return _IInitialReporter.Contract.GetReportTimestamp(&_IInitialReporter.CallOpts)
}

// GetReportTimestamp is a free data retrieval call binding the contract method 0x94a771d7.
//
// Solidity: function getReportTimestamp() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCallerSession) GetReportTimestamp() (*big.Int, error) {
	return _IInitialReporter.Contract.GetReportTimestamp(&_IInitialReporter.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_IInitialReporter *IInitialReporterSession) GetSize() (*big.Int, error) {
	return _IInitialReporter.Contract.GetSize(&_IInitialReporter.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCallerSession) GetSize() (*big.Int, error) {
	return _IInitialReporter.Contract.GetSize(&_IInitialReporter.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCaller) GetStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "getStake")
	return *ret0, err
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_IInitialReporter *IInitialReporterSession) GetStake() (*big.Int, error) {
	return _IInitialReporter.Contract.GetStake(&_IInitialReporter.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_IInitialReporter *IInitialReporterCallerSession) GetStake() (*big.Int, error) {
	return _IInitialReporter.Contract.GetStake(&_IInitialReporter.CallOpts)
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterCaller) IsDisavowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "isDisavowed")
	return *ret0, err
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterSession) IsDisavowed() (bool, error) {
	return _IInitialReporter.Contract.IsDisavowed(&_IInitialReporter.CallOpts)
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_IInitialReporter *IInitialReporterCallerSession) IsDisavowed() (bool, error) {
	return _IInitialReporter.Contract.IsDisavowed(&_IInitialReporter.CallOpts)
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_IInitialReporter *IInitialReporterCaller) IsInvalid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IInitialReporter.contract.Call(opts, out, "isInvalid")
	return *ret0, err
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_IInitialReporter *IInitialReporterSession) IsInvalid() (bool, error) {
	return _IInitialReporter.Contract.IsInvalid(&_IInitialReporter.CallOpts)
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_IInitialReporter *IInitialReporterCallerSession) IsInvalid() (bool, error) {
	return _IInitialReporter.Contract.IsInvalid(&_IInitialReporter.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_market address, _designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) Initialize(opts *bind.TransactOpts, _market common.Address, _designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "initialize", _market, _designatedReporter)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_market address, _designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterSession) Initialize(_market common.Address, _designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Initialize(&_IInitialReporter.TransactOpts, _market, _designatedReporter)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_market address, _designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) Initialize(_market common.Address, _designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Initialize(&_IInitialReporter.TransactOpts, _market, _designatedReporter)
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) LiquidateLosing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "liquidateLosing")
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_IInitialReporter *IInitialReporterSession) LiquidateLosing() (*types.Transaction, error) {
	return _IInitialReporter.Contract.LiquidateLosing(&_IInitialReporter.TransactOpts)
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) LiquidateLosing() (*types.Transaction, error) {
	return _IInitialReporter.Contract.LiquidateLosing(&_IInitialReporter.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_IInitialReporter *IInitialReporterSession) Migrate() (*types.Transaction, error) {
	return _IInitialReporter.Contract.Migrate(&_IInitialReporter.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) Migrate() (*types.Transaction, error) {
	return _IInitialReporter.Contract.Migrate(&_IInitialReporter.TransactOpts)
}

// MigrateToNewUniverse is a paid mutator transaction binding the contract method 0xef0ce37a.
//
// Solidity: function migrateToNewUniverse(_designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) MigrateToNewUniverse(opts *bind.TransactOpts, _designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "migrateToNewUniverse", _designatedReporter)
}

// MigrateToNewUniverse is a paid mutator transaction binding the contract method 0xef0ce37a.
//
// Solidity: function migrateToNewUniverse(_designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterSession) MigrateToNewUniverse(_designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.MigrateToNewUniverse(&_IInitialReporter.TransactOpts, _designatedReporter)
}

// MigrateToNewUniverse is a paid mutator transaction binding the contract method 0xef0ce37a.
//
// Solidity: function migrateToNewUniverse(_designatedReporter address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) MigrateToNewUniverse(_designatedReporter common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.MigrateToNewUniverse(&_IInitialReporter.TransactOpts, _designatedReporter)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) Redeem(opts *bind.TransactOpts, _redeemer common.Address) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "redeem", _redeemer)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_IInitialReporter *IInitialReporterSession) Redeem(_redeemer common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Redeem(&_IInitialReporter.TransactOpts, _redeemer)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) Redeem(_redeemer common.Address) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Redeem(&_IInitialReporter.TransactOpts, _redeemer)
}

// Report is a paid mutator transaction binding the contract method 0xc6f344f5.
//
// Solidity: function report(_reporter address, _payoutDistributionHash bytes32, _payoutNumerators uint256[], _invalid bool, _initialReportStake uint256) returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) Report(opts *bind.TransactOpts, _reporter common.Address, _payoutDistributionHash [32]byte, _payoutNumerators []*big.Int, _invalid bool, _initialReportStake *big.Int) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "report", _reporter, _payoutDistributionHash, _payoutNumerators, _invalid, _initialReportStake)
}

// Report is a paid mutator transaction binding the contract method 0xc6f344f5.
//
// Solidity: function report(_reporter address, _payoutDistributionHash bytes32, _payoutNumerators uint256[], _invalid bool, _initialReportStake uint256) returns(bool)
func (_IInitialReporter *IInitialReporterSession) Report(_reporter common.Address, _payoutDistributionHash [32]byte, _payoutNumerators []*big.Int, _invalid bool, _initialReportStake *big.Int) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Report(&_IInitialReporter.TransactOpts, _reporter, _payoutDistributionHash, _payoutNumerators, _invalid, _initialReportStake)
}

// Report is a paid mutator transaction binding the contract method 0xc6f344f5.
//
// Solidity: function report(_reporter address, _payoutDistributionHash bytes32, _payoutNumerators uint256[], _invalid bool, _initialReportStake uint256) returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) Report(_reporter common.Address, _payoutDistributionHash [32]byte, _payoutNumerators []*big.Int, _invalid bool, _initialReportStake *big.Int) (*types.Transaction, error) {
	return _IInitialReporter.Contract.Report(&_IInitialReporter.TransactOpts, _reporter, _payoutDistributionHash, _payoutNumerators, _invalid, _initialReportStake)
}

// ReturnRepFromDisavow is a paid mutator transaction binding the contract method 0x87882cb8.
//
// Solidity: function returnRepFromDisavow() returns(bool)
func (_IInitialReporter *IInitialReporterTransactor) ReturnRepFromDisavow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInitialReporter.contract.Transact(opts, "returnRepFromDisavow")
}

// ReturnRepFromDisavow is a paid mutator transaction binding the contract method 0x87882cb8.
//
// Solidity: function returnRepFromDisavow() returns(bool)
func (_IInitialReporter *IInitialReporterSession) ReturnRepFromDisavow() (*types.Transaction, error) {
	return _IInitialReporter.Contract.ReturnRepFromDisavow(&_IInitialReporter.TransactOpts)
}

// ReturnRepFromDisavow is a paid mutator transaction binding the contract method 0x87882cb8.
//
// Solidity: function returnRepFromDisavow() returns(bool)
func (_IInitialReporter *IInitialReporterTransactorSession) ReturnRepFromDisavow() (*types.Transaction, error) {
	return _IInitialReporter.Contract.ReturnRepFromDisavow(&_IInitialReporter.TransactOpts)
}
