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

// BaseReportingParticipantABI is the input ABI used to generate the binding from.
const BaseReportingParticipantABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isInvalid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"liquidateLosing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isDisavowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_redeemer\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getPayoutNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPayoutDistributionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BaseReportingParticipantBin is the compiled bytecode used for deploying new contracts.
const BaseReportingParticipantBin = `0x`

// DeployBaseReportingParticipant deploys a new Ethereum contract, binding an instance of BaseReportingParticipant to it.
func DeployBaseReportingParticipant(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseReportingParticipant, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseReportingParticipantABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseReportingParticipantBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseReportingParticipant{BaseReportingParticipantCaller: BaseReportingParticipantCaller{contract: contract}, BaseReportingParticipantTransactor: BaseReportingParticipantTransactor{contract: contract}, BaseReportingParticipantFilterer: BaseReportingParticipantFilterer{contract: contract}}, nil
}

// BaseReportingParticipant is an auto generated Go binding around an Ethereum contract.
type BaseReportingParticipant struct {
	BaseReportingParticipantCaller     // Read-only binding to the contract
	BaseReportingParticipantTransactor // Write-only binding to the contract
	BaseReportingParticipantFilterer   // Log filterer for contract events
}

// BaseReportingParticipantCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseReportingParticipantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseReportingParticipantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseReportingParticipantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseReportingParticipantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseReportingParticipantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseReportingParticipantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseReportingParticipantSession struct {
	Contract     *BaseReportingParticipant // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BaseReportingParticipantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseReportingParticipantCallerSession struct {
	Contract *BaseReportingParticipantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// BaseReportingParticipantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseReportingParticipantTransactorSession struct {
	Contract     *BaseReportingParticipantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// BaseReportingParticipantRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseReportingParticipantRaw struct {
	Contract *BaseReportingParticipant // Generic contract binding to access the raw methods on
}

// BaseReportingParticipantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseReportingParticipantCallerRaw struct {
	Contract *BaseReportingParticipantCaller // Generic read-only contract binding to access the raw methods on
}

// BaseReportingParticipantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseReportingParticipantTransactorRaw struct {
	Contract *BaseReportingParticipantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseReportingParticipant creates a new instance of BaseReportingParticipant, bound to a specific deployed contract.
func NewBaseReportingParticipant(address common.Address, backend bind.ContractBackend) (*BaseReportingParticipant, error) {
	contract, err := bindBaseReportingParticipant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseReportingParticipant{BaseReportingParticipantCaller: BaseReportingParticipantCaller{contract: contract}, BaseReportingParticipantTransactor: BaseReportingParticipantTransactor{contract: contract}, BaseReportingParticipantFilterer: BaseReportingParticipantFilterer{contract: contract}}, nil
}

// NewBaseReportingParticipantCaller creates a new read-only instance of BaseReportingParticipant, bound to a specific deployed contract.
func NewBaseReportingParticipantCaller(address common.Address, caller bind.ContractCaller) (*BaseReportingParticipantCaller, error) {
	contract, err := bindBaseReportingParticipant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseReportingParticipantCaller{contract: contract}, nil
}

// NewBaseReportingParticipantTransactor creates a new write-only instance of BaseReportingParticipant, bound to a specific deployed contract.
func NewBaseReportingParticipantTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseReportingParticipantTransactor, error) {
	contract, err := bindBaseReportingParticipant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseReportingParticipantTransactor{contract: contract}, nil
}

// NewBaseReportingParticipantFilterer creates a new log filterer instance of BaseReportingParticipant, bound to a specific deployed contract.
func NewBaseReportingParticipantFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseReportingParticipantFilterer, error) {
	contract, err := bindBaseReportingParticipant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseReportingParticipantFilterer{contract: contract}, nil
}

// bindBaseReportingParticipant binds a generic wrapper to an already deployed contract.
func bindBaseReportingParticipant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseReportingParticipantABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseReportingParticipant *BaseReportingParticipantRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseReportingParticipant.Contract.BaseReportingParticipantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseReportingParticipant *BaseReportingParticipantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.BaseReportingParticipantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseReportingParticipant *BaseReportingParticipantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.BaseReportingParticipantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseReportingParticipant *BaseReportingParticipantCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseReportingParticipant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseReportingParticipant *BaseReportingParticipantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseReportingParticipant *BaseReportingParticipantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_BaseReportingParticipant *BaseReportingParticipantCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseReportingParticipant.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_BaseReportingParticipant *BaseReportingParticipantSession) GetController() (common.Address, error) {
	return _BaseReportingParticipant.Contract.GetController(&_BaseReportingParticipant.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_BaseReportingParticipant *BaseReportingParticipantCallerSession) GetController() (common.Address, error) {
	return _BaseReportingParticipant.Contract.GetController(&_BaseReportingParticipant.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_BaseReportingParticipant *BaseReportingParticipantCaller) GetMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseReportingParticipant.contract.Call(opts, out, "getMarket")
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_BaseReportingParticipant *BaseReportingParticipantSession) GetMarket() (common.Address, error) {
	return _BaseReportingParticipant.Contract.GetMarket(&_BaseReportingParticipant.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_BaseReportingParticipant *BaseReportingParticipantCallerSession) GetMarket() (common.Address, error) {
	return _BaseReportingParticipant.Contract.GetMarket(&_BaseReportingParticipant.CallOpts)
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_BaseReportingParticipant *BaseReportingParticipantCaller) GetPayoutDistributionHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseReportingParticipant.contract.Call(opts, out, "getPayoutDistributionHash")
	return *ret0, err
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_BaseReportingParticipant *BaseReportingParticipantSession) GetPayoutDistributionHash() ([32]byte, error) {
	return _BaseReportingParticipant.Contract.GetPayoutDistributionHash(&_BaseReportingParticipant.CallOpts)
}

// GetPayoutDistributionHash is a free data retrieval call binding the contract method 0xdd62f6f3.
//
// Solidity: function getPayoutDistributionHash() constant returns(bytes32)
func (_BaseReportingParticipant *BaseReportingParticipantCallerSession) GetPayoutDistributionHash() ([32]byte, error) {
	return _BaseReportingParticipant.Contract.GetPayoutDistributionHash(&_BaseReportingParticipant.CallOpts)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_BaseReportingParticipant *BaseReportingParticipantCaller) GetPayoutNumerator(opts *bind.CallOpts, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseReportingParticipant.contract.Call(opts, out, "getPayoutNumerator", _outcome)
	return *ret0, err
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_BaseReportingParticipant *BaseReportingParticipantSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _BaseReportingParticipant.Contract.GetPayoutNumerator(&_BaseReportingParticipant.CallOpts, _outcome)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(_outcome uint256) constant returns(uint256)
func (_BaseReportingParticipant *BaseReportingParticipantCallerSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _BaseReportingParticipant.Contract.GetPayoutNumerator(&_BaseReportingParticipant.CallOpts, _outcome)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_BaseReportingParticipant *BaseReportingParticipantCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseReportingParticipant.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_BaseReportingParticipant *BaseReportingParticipantSession) GetSize() (*big.Int, error) {
	return _BaseReportingParticipant.Contract.GetSize(&_BaseReportingParticipant.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_BaseReportingParticipant *BaseReportingParticipantCallerSession) GetSize() (*big.Int, error) {
	return _BaseReportingParticipant.Contract.GetSize(&_BaseReportingParticipant.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_BaseReportingParticipant *BaseReportingParticipantCaller) GetStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseReportingParticipant.contract.Call(opts, out, "getStake")
	return *ret0, err
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_BaseReportingParticipant *BaseReportingParticipantSession) GetStake() (*big.Int, error) {
	return _BaseReportingParticipant.Contract.GetStake(&_BaseReportingParticipant.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_BaseReportingParticipant *BaseReportingParticipantCallerSession) GetStake() (*big.Int, error) {
	return _BaseReportingParticipant.Contract.GetStake(&_BaseReportingParticipant.CallOpts)
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantCaller) IsDisavowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseReportingParticipant.contract.Call(opts, out, "isDisavowed")
	return *ret0, err
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantSession) IsDisavowed() (bool, error) {
	return _BaseReportingParticipant.Contract.IsDisavowed(&_BaseReportingParticipant.CallOpts)
}

// IsDisavowed is a free data retrieval call binding the contract method 0x64020842.
//
// Solidity: function isDisavowed() constant returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantCallerSession) IsDisavowed() (bool, error) {
	return _BaseReportingParticipant.Contract.IsDisavowed(&_BaseReportingParticipant.CallOpts)
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantCaller) IsInvalid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseReportingParticipant.contract.Call(opts, out, "isInvalid")
	return *ret0, err
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantSession) IsInvalid() (bool, error) {
	return _BaseReportingParticipant.Contract.IsInvalid(&_BaseReportingParticipant.CallOpts)
}

// IsInvalid is a free data retrieval call binding the contract method 0x04be2f50.
//
// Solidity: function isInvalid() constant returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantCallerSession) IsInvalid() (bool, error) {
	return _BaseReportingParticipant.Contract.IsInvalid(&_BaseReportingParticipant.CallOpts)
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantTransactor) LiquidateLosing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseReportingParticipant.contract.Transact(opts, "liquidateLosing")
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantSession) LiquidateLosing() (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.LiquidateLosing(&_BaseReportingParticipant.TransactOpts)
}

// LiquidateLosing is a paid mutator transaction binding the contract method 0x22b152a3.
//
// Solidity: function liquidateLosing() returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantTransactorSession) LiquidateLosing() (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.LiquidateLosing(&_BaseReportingParticipant.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseReportingParticipant.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantSession) Migrate() (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.Migrate(&_BaseReportingParticipant.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantTransactorSession) Migrate() (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.Migrate(&_BaseReportingParticipant.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantTransactor) Redeem(opts *bind.TransactOpts, _redeemer common.Address) (*types.Transaction, error) {
	return _BaseReportingParticipant.contract.Transact(opts, "redeem", _redeemer)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantSession) Redeem(_redeemer common.Address) (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.Redeem(&_BaseReportingParticipant.TransactOpts, _redeemer)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_redeemer address) returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantTransactorSession) Redeem(_redeemer common.Address) (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.Redeem(&_BaseReportingParticipant.TransactOpts, _redeemer)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _BaseReportingParticipant.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.SetController(&_BaseReportingParticipant.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_BaseReportingParticipant *BaseReportingParticipantTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _BaseReportingParticipant.Contract.SetController(&_BaseReportingParticipant.TransactOpts, _controller)
}

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
