package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IReportingParticipantABI is the input ABI used to generate the binding from.
const IReportingParticipantABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isInvalid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"liquidateLosing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isDisavowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_redeemer\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getPayoutNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPayoutDistributionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IReportingParticipantBin is the compiled bytecode used for deploying new contracts.
const IReportingParticipantBin = `0x`

// DeployIReportingParticipant deploys a new Ethereum contract, binding an instance of IReportingParticipant to it.
func DeployIReportingParticipant(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IReportingParticipant, error) {
	parsed, err := abi.JSON(strings.NewReader(IReportingParticipantABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IReportingParticipantBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IReportingParticipant{IReportingParticipantCaller: IReportingParticipantCaller{contract: contract}, IReportingParticipantTransactor: IReportingParticipantTransactor{contract: contract}, IReportingParticipantFilterer: IReportingParticipantFilterer{contract: contract}}, nil
}

// IReportingParticipant is an auto generated Go binding around an Ethereum contract.
type IReportingParticipant struct {
	IReportingParticipantCaller     // Read-only binding to the contract
	IReportingParticipantTransactor // Write-only binding to the contract
	IReportingParticipantFilterer   // Log filterer for contract events
}

// IReportingParticipantCaller is an auto generated read-only Go binding around an Ethereum contract.
type IReportingParticipantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReportingParticipantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IReportingParticipantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReportingParticipantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IReportingParticipantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReportingParticipantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IReportingParticipantSession struct {
	Contract     *IReportingParticipant // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IReportingParticipantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IReportingParticipantCallerSession struct {
	Contract *IReportingParticipantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IReportingParticipantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IReportingParticipantTransactorSession struct {
	Contract     *IReportingParticipantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IReportingParticipantRaw is an auto generated low-level Go binding around an Ethereum contract.
type IReportingParticipantRaw struct {
	Contract *IReportingParticipant // Generic contract binding to access the raw methods on
}

// IReportingParticipantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IReportingParticipantCallerRaw struct {
	Contract *IReportingParticipantCaller // Generic read-only contract binding to access the raw methods on
}

// IReportingParticipantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IReportingParticipantTransactorRaw struct {
	Contract *IReportingParticipantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReportingParticipant creates a new instance of IReportingParticipant, bound to a specific deployed contract.
func NewIReportingParticipant(address common.Address, backend bind.ContractBackend) (*IReportingParticipant, error) {
	contract, err := bindIReportingParticipant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReportingParticipant{IReportingParticipantCaller: IReportingParticipantCaller{contract: contract}, IReportingParticipantTransactor: IReportingParticipantTransactor{contract: contract}, IReportingParticipantFilterer: IReportingParticipantFilterer{contract: contract}}, nil
}

// NewIReportingParticipantCaller creates a new read-only instance of IReportingParticipant, bound to a specific deployed contract.
func NewIReportingParticipantCaller(address common.Address, caller bind.ContractCaller) (*IReportingParticipantCaller, error) {
	contract, err := bindIReportingParticipant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IReportingParticipantCaller{contract: contract}, nil
}

// NewIReportingParticipantTransactor creates a new write-only instance of IReportingParticipant, bound to a specific deployed contract.
func NewIReportingParticipantTransactor(address common.Address, transactor bind.ContractTransactor) (*IReportingParticipantTransactor, error) {
	contract, err := bindIReportingParticipant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IReportingParticipantTransactor{contract: contract}, nil
}

// NewIReportingParticipantFilterer creates a new log filterer instance of IReportingParticipant, bound to a specific deployed contract.
func NewIReportingParticipantFilterer(address common.Address, filterer bind.ContractFilterer) (*IReportingParticipantFilterer, error) {
	contract, err := bindIReportingParticipant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IReportingParticipantFilterer{contract: contract}, nil
}

// bindIReportingParticipant binds a generic wrapper to an already deployed contract.
func bindIReportingParticipant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IReportingParticipantABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReportingParticipant *IReportingParticipantRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IReportingParticipant.Contract.IReportingParticipantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReportingParticipant *IReportingParticipantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReportingParticipant.Contract.IReportingParticipantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReportingParticipant *IReportingParticipantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReportingParticipant.Contract.IReportingParticipantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReportingParticipant *IReportingParticipantCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IReportingParticipant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReportingParticipant *IReportingParticipantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReportingParticipant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReportingParticipant *IReportingParticipantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReportingParticipant.Contract.contract.Transact(opts, method, params...)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IReportingParticipant *IReportingParticipantCaller) GetMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IReportingParticipant.contract.Call(opts, out, "getMarket")
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IReportingParticipant *IReportingParticipantSession) GetMarket() (common.Address, error) {
	return _IReportingParticipant.Contract.GetMarket(&_IReportingParticipant.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IReportingParticipant *IReportingParticipantCallerSession) GetMarket() (common.Address, error) {
	return _IReportingParticipant.Contract.GetMarket(&_IReportingParticipant.CallOpts)
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_IReportingParticipant *IReportingParticipantCaller) GetPayoutDistributionHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IReportingParticipant.contract.Call(opts, out, "getPayoutDistributionHash")
	return *ret0, err
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_IReportingParticipant *IReportingParticipantSession) GetPayoutDistributionHash() ([32]byte, error) {
	return _IReportingParticipant.Contract.GetPayoutDistributionHash(&_IReportingParticipant.CallOpts)
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_IReportingParticipant *IReportingParticipantCallerSession) GetPayoutDistributionHash() ([32]byte, error) {
	return _IReportingParticipant.Contract.GetPayoutDistributionHash(&_IReportingParticipant.CallOpts)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_IReportingParticipant *IReportingParticipantCaller) GetPayoutNumerator(opts *bind.CallOpts, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IReportingParticipant.contract.Call(opts, out, "getPayoutNumerator", _outcome)
	return *ret0, err
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_IReportingParticipant *IReportingParticipantSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _IReportingParticipant.Contract.GetPayoutNumerator(&_IReportingParticipant.CallOpts, _outcome)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_IReportingParticipant *IReportingParticipantCallerSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _IReportingParticipant.Contract.GetPayoutNumerator(&_IReportingParticipant.CallOpts, _outcome)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_IReportingParticipant *IReportingParticipantCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IReportingParticipant.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_IReportingParticipant *IReportingParticipantSession) GetSize() (*big.Int, error) {
	return _IReportingParticipant.Contract.GetSize(&_IReportingParticipant.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_IReportingParticipant *IReportingParticipantCallerSession) GetSize() (*big.Int, error) {
	return _IReportingParticipant.Contract.GetSize(&_IReportingParticipant.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_IReportingParticipant *IReportingParticipantCaller) GetStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IReportingParticipant.contract.Call(opts, out, "getStake")
	return *ret0, err
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_IReportingParticipant *IReportingParticipantSession) GetStake() (*big.Int, error) {
	return _IReportingParticipant.Contract.GetStake(&_IReportingParticipant.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_IReportingParticipant *IReportingParticipantCallerSession) GetStake() (*big.Int, error) {
	return _IReportingParticipant.Contract.GetStake(&_IReportingParticipant.CallOpts)
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_IReportingParticipant *IReportingParticipantCaller) IsDisavowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IReportingParticipant.contract.Call(opts, out, "isDisavowed")
	return *ret0, err
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_IReportingParticipant *IReportingParticipantSession) IsDisavowed() (bool, error) {
	return _IReportingParticipant.Contract.IsDisavowed(&_IReportingParticipant.CallOpts)
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_IReportingParticipant *IReportingParticipantCallerSession) IsDisavowed() (bool, error) {
	return _IReportingParticipant.Contract.IsDisavowed(&_IReportingParticipant.CallOpts)
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_IReportingParticipant *IReportingParticipantCaller) IsInvalid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IReportingParticipant.contract.Call(opts, out, "isInvalid")
	return *ret0, err
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_IReportingParticipant *IReportingParticipantSession) IsInvalid() (bool, error) {
	return _IReportingParticipant.Contract.IsInvalid(&_IReportingParticipant.CallOpts)
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_IReportingParticipant *IReportingParticipantCallerSession) IsInvalid() (bool, error) {
	return _IReportingParticipant.Contract.IsInvalid(&_IReportingParticipant.CallOpts)
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_IReportingParticipant *IReportingParticipantTransactor) LiquidateLosing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReportingParticipant.contract.Transact(opts, "liquidateLosing")
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_IReportingParticipant *IReportingParticipantSession) LiquidateLosing() (*types.Transaction, error) {
	return _IReportingParticipant.Contract.LiquidateLosing(&_IReportingParticipant.TransactOpts)
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_IReportingParticipant *IReportingParticipantTransactorSession) LiquidateLosing() (*types.Transaction, error) {
	return _IReportingParticipant.Contract.LiquidateLosing(&_IReportingParticipant.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_IReportingParticipant *IReportingParticipantTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReportingParticipant.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_IReportingParticipant *IReportingParticipantSession) Migrate() (*types.Transaction, error) {
	return _IReportingParticipant.Contract.Migrate(&_IReportingParticipant.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_IReportingParticipant *IReportingParticipantTransactorSession) Migrate() (*types.Transaction, error) {
	return _IReportingParticipant.Contract.Migrate(&_IReportingParticipant.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_IReportingParticipant *IReportingParticipantTransactor) Redeem(opts *bind.TransactOpts, _redeemer common.Address) (*types.Transaction, error) {
	return _IReportingParticipant.contract.Transact(opts, "redeem", _redeemer)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_IReportingParticipant *IReportingParticipantSession) Redeem(_redeemer common.Address) (*types.Transaction, error) {
	return _IReportingParticipant.Contract.Redeem(&_IReportingParticipant.TransactOpts, _redeemer)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_IReportingParticipant *IReportingParticipantTransactorSession) Redeem(_redeemer common.Address) (*types.Transaction, error) {
	return _IReportingParticipant.Contract.Redeem(&_IReportingParticipant.TransactOpts, _redeemer)
}

// ReportingABI is the input ABI used to generate the binding from.
const ReportingABI = "[]"

// ReportingBin is the compiled bytecode used for deploying new contracts.
const ReportingBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206ce5eae0740986f1a929443c3dc2f4930d12332e0b18cbe286ea91dcc0a787370029`

// DeployReporting deploys a new Ethereum contract, binding an instance of Reporting to it.
func DeployReporting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reporting, error) {
	parsed, err := abi.JSON(strings.NewReader(ReportingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReportingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reporting{ReportingCaller: ReportingCaller{contract: contract}, ReportingTransactor: ReportingTransactor{contract: contract}, ReportingFilterer: ReportingFilterer{contract: contract}}, nil
}

// Reporting is an auto generated Go binding around an Ethereum contract.
type Reporting struct {
	ReportingCaller     // Read-only binding to the contract
	ReportingTransactor // Write-only binding to the contract
	ReportingFilterer   // Log filterer for contract events
}

// ReportingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReportingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReportingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReportingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReportingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReportingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReportingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReportingSession struct {
	Contract     *Reporting        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReportingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReportingCallerSession struct {
	Contract *ReportingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ReportingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReportingTransactorSession struct {
	Contract     *ReportingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ReportingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReportingRaw struct {
	Contract *Reporting // Generic contract binding to access the raw methods on
}

// ReportingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReportingCallerRaw struct {
	Contract *ReportingCaller // Generic read-only contract binding to access the raw methods on
}

// ReportingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReportingTransactorRaw struct {
	Contract *ReportingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReporting creates a new instance of Reporting, bound to a specific deployed contract.
func NewReporting(address common.Address, backend bind.ContractBackend) (*Reporting, error) {
	contract, err := bindReporting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reporting{ReportingCaller: ReportingCaller{contract: contract}, ReportingTransactor: ReportingTransactor{contract: contract}, ReportingFilterer: ReportingFilterer{contract: contract}}, nil
}

// NewReportingCaller creates a new read-only instance of Reporting, bound to a specific deployed contract.
func NewReportingCaller(address common.Address, caller bind.ContractCaller) (*ReportingCaller, error) {
	contract, err := bindReporting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReportingCaller{contract: contract}, nil
}

// NewReportingTransactor creates a new write-only instance of Reporting, bound to a specific deployed contract.
func NewReportingTransactor(address common.Address, transactor bind.ContractTransactor) (*ReportingTransactor, error) {
	contract, err := bindReporting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReportingTransactor{contract: contract}, nil
}

// NewReportingFilterer creates a new log filterer instance of Reporting, bound to a specific deployed contract.
func NewReportingFilterer(address common.Address, filterer bind.ContractFilterer) (*ReportingFilterer, error) {
	contract, err := bindReporting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReportingFilterer{contract: contract}, nil
}

// bindReporting binds a generic wrapper to an already deployed contract.
func bindReporting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReportingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reporting *ReportingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reporting.Contract.ReportingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reporting *ReportingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reporting.Contract.ReportingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reporting *ReportingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reporting.Contract.ReportingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reporting *ReportingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reporting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reporting *ReportingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reporting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reporting *ReportingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reporting.Contract.contract.Transact(opts, method, params...)
}
