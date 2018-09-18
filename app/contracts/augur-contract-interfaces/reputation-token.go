package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IReputationTokenABI is the input ABI used to generate the binding from.
const IReputationTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalTheoreticalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOut\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedFeeWindowTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateIn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedReportingParticipantTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountMigrated\",\"type\":\"uint256\"}],\"name\":\"mintForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedMarketTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedUniverseTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// IReputationTokenBin is the compiled bytecode used for deploying new contracts.
const IReputationTokenBin = `0x`

// DeployIReputationToken deploys a new Ethereum contract, binding an instance of IReputationToken to it.
func DeployIReputationToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IReputationToken, error) {
	parsed, err := abi.JSON(strings.NewReader(IReputationTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IReputationTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IReputationToken{IReputationTokenCaller: IReputationTokenCaller{contract: contract}, IReputationTokenTransactor: IReputationTokenTransactor{contract: contract}, IReputationTokenFilterer: IReputationTokenFilterer{contract: contract}}, nil
}

// IReputationToken is an auto generated Go binding around an Ethereum contract.
type IReputationToken struct {
	IReputationTokenCaller     // Read-only binding to the contract
	IReputationTokenTransactor // Write-only binding to the contract
	IReputationTokenFilterer   // Log filterer for contract events
}

// IReputationTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IReputationTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReputationTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IReputationTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReputationTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IReputationTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReputationTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IReputationTokenSession struct {
	Contract     *IReputationToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IReputationTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IReputationTokenCallerSession struct {
	Contract *IReputationTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IReputationTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IReputationTokenTransactorSession struct {
	Contract     *IReputationTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IReputationTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IReputationTokenRaw struct {
	Contract *IReputationToken // Generic contract binding to access the raw methods on
}

// IReputationTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IReputationTokenCallerRaw struct {
	Contract *IReputationTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IReputationTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IReputationTokenTransactorRaw struct {
	Contract *IReputationTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReputationToken creates a new instance of IReputationToken, bound to a specific deployed contract.
func NewIReputationToken(address common.Address, backend bind.ContractBackend) (*IReputationToken, error) {
	contract, err := bindIReputationToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReputationToken{IReputationTokenCaller: IReputationTokenCaller{contract: contract}, IReputationTokenTransactor: IReputationTokenTransactor{contract: contract}, IReputationTokenFilterer: IReputationTokenFilterer{contract: contract}}, nil
}

// NewIReputationTokenCaller creates a new read-only instance of IReputationToken, bound to a specific deployed contract.
func NewIReputationTokenCaller(address common.Address, caller bind.ContractCaller) (*IReputationTokenCaller, error) {
	contract, err := bindIReputationToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IReputationTokenCaller{contract: contract}, nil
}

// NewIReputationTokenTransactor creates a new write-only instance of IReputationToken, bound to a specific deployed contract.
func NewIReputationTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IReputationTokenTransactor, error) {
	contract, err := bindIReputationToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IReputationTokenTransactor{contract: contract}, nil
}

// NewIReputationTokenFilterer creates a new log filterer instance of IReputationToken, bound to a specific deployed contract.
func NewIReputationTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IReputationTokenFilterer, error) {
	contract, err := bindIReputationToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IReputationTokenFilterer{contract: contract}, nil
}

// bindIReputationToken binds a generic wrapper to an already deployed contract.
func bindIReputationToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IReputationTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReputationToken *IReputationTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IReputationToken.Contract.IReputationTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReputationToken *IReputationTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReputationToken.Contract.IReputationTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReputationToken *IReputationTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReputationToken.Contract.IReputationTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReputationToken *IReputationTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IReputationToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReputationToken *IReputationTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReputationToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReputationToken *IReputationTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReputationToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IReputationToken *IReputationTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IReputationToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IReputationToken *IReputationTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IReputationToken.Contract.Allowance(&_IReputationToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IReputationToken *IReputationTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IReputationToken.Contract.Allowance(&_IReputationToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IReputationToken *IReputationTokenCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IReputationToken.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IReputationToken *IReputationTokenSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _IReputationToken.Contract.BalanceOf(&_IReputationToken.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IReputationToken *IReputationTokenCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _IReputationToken.Contract.BalanceOf(&_IReputationToken.CallOpts, _who)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_IReputationToken *IReputationTokenCaller) GetTotalMigrated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IReputationToken.contract.Call(opts, out, "getTotalMigrated")
	return *ret0, err
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_IReputationToken *IReputationTokenSession) GetTotalMigrated() (*big.Int, error) {
	return _IReputationToken.Contract.GetTotalMigrated(&_IReputationToken.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_IReputationToken *IReputationTokenCallerSession) GetTotalMigrated() (*big.Int, error) {
	return _IReputationToken.Contract.GetTotalMigrated(&_IReputationToken.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_IReputationToken *IReputationTokenCaller) GetTotalTheoreticalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IReputationToken.contract.Call(opts, out, "getTotalTheoreticalSupply")
	return *ret0, err
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_IReputationToken *IReputationTokenSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _IReputationToken.Contract.GetTotalTheoreticalSupply(&_IReputationToken.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_IReputationToken *IReputationTokenCallerSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _IReputationToken.Contract.GetTotalTheoreticalSupply(&_IReputationToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IReputationToken *IReputationTokenCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IReputationToken.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IReputationToken *IReputationTokenSession) GetTypeName() ([32]byte, error) {
	return _IReputationToken.Contract.GetTypeName(&_IReputationToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IReputationToken *IReputationTokenCallerSession) GetTypeName() ([32]byte, error) {
	return _IReputationToken.Contract.GetTypeName(&_IReputationToken.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_IReputationToken *IReputationTokenCaller) GetUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IReputationToken.contract.Call(opts, out, "getUniverse")
	return *ret0, err
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_IReputationToken *IReputationTokenSession) GetUniverse() (common.Address, error) {
	return _IReputationToken.Contract.GetUniverse(&_IReputationToken.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_IReputationToken *IReputationTokenCallerSession) GetUniverse() (common.Address, error) {
	return _IReputationToken.Contract.GetUniverse(&_IReputationToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IReputationToken *IReputationTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IReputationToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IReputationToken *IReputationTokenSession) TotalSupply() (*big.Int, error) {
	return _IReputationToken.Contract.TotalSupply(&_IReputationToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IReputationToken *IReputationTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IReputationToken.Contract.TotalSupply(&_IReputationToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.Approve(&_IReputationToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.Approve(&_IReputationToken.TransactOpts, _spender, _value)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) Initialize(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "initialize", _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_IReputationToken *IReputationTokenSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _IReputationToken.Contract.Initialize(&_IReputationToken.TransactOpts, _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _IReputationToken.Contract.Initialize(&_IReputationToken.TransactOpts, _universe)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) MigrateIn(opts *bind.TransactOpts, _reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "migrateIn", _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.MigrateIn(&_IReputationToken.TransactOpts, _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.MigrateIn(&_IReputationToken.TransactOpts, _reporter, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) MigrateOut(opts *bind.TransactOpts, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "migrateOut", _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.MigrateOut(&_IReputationToken.TransactOpts, _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.MigrateOut(&_IReputationToken.TransactOpts, _destination, _attotokens)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) MintForReportingParticipant(opts *bind.TransactOpts, _amountMigrated *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "mintForReportingParticipant", _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.MintForReportingParticipant(&_IReputationToken.TransactOpts, _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.MintForReportingParticipant(&_IReputationToken.TransactOpts, _amountMigrated)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.Transfer(&_IReputationToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.Transfer(&_IReputationToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TransferFrom(&_IReputationToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TransferFrom(&_IReputationToken.TransactOpts, _from, _to, _value)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) TrustedFeeWindowTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "trustedFeeWindowTransfer", _source, _destination, _attotokens)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) TrustedFeeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TrustedFeeWindowTransfer(&_IReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) TrustedFeeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TrustedFeeWindowTransfer(&_IReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) TrustedMarketTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "trustedMarketTransfer", _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TrustedMarketTransfer(&_IReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TrustedMarketTransfer(&_IReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) TrustedReportingParticipantTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "trustedReportingParticipantTransfer", _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TrustedReportingParticipantTransfer(&_IReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TrustedReportingParticipantTransfer(&_IReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactor) TrustedUniverseTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.contract.Transact(opts, "trustedUniverseTransfer", _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TrustedUniverseTransfer(&_IReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IReputationToken *IReputationTokenTransactorSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IReputationToken.Contract.TrustedUniverseTransfer(&_IReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// IReputationTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IReputationToken contract.
type IReputationTokenApprovalIterator struct {
	Event *IReputationTokenApproval // Event containing the contract specifics and raw log

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
func (it *IReputationTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReputationTokenApproval)
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
		it.Event = new(IReputationTokenApproval)
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
func (it *IReputationTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReputationTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReputationTokenApproval represents a Approval event raised by the IReputationToken contract.
type IReputationTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IReputationToken *IReputationTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IReputationTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IReputationToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IReputationTokenApprovalIterator{contract: _IReputationToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IReputationToken *IReputationTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IReputationTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IReputationToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReputationTokenApproval)
				if err := _IReputationToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IReputationTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IReputationToken contract.
type IReputationTokenTransferIterator struct {
	Event *IReputationTokenTransfer // Event containing the contract specifics and raw log

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
func (it *IReputationTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReputationTokenTransfer)
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
		it.Event = new(IReputationTokenTransfer)
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
func (it *IReputationTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReputationTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReputationTokenTransfer represents a Transfer event raised by the IReputationToken contract.
type IReputationTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IReputationToken *IReputationTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IReputationTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IReputationToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IReputationTokenTransferIterator{contract: _IReputationToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IReputationToken *IReputationTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IReputationTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IReputationToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReputationTokenTransfer)
				if err := _IReputationToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ReputationTokenABI is the input ABI used to generate the binding from.
const ReputationTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalTheoreticalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_invalid\",\"type\":\"bool\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOutByPayout\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETERNAL_APPROVAL_VALUE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOut\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrateFromLegacyReputationToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLegacyRepToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedFeeWindowTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateIn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateParentTotalTheoreticalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedReportingParticipantTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"updateSiblingMigrationTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountMigrated\",\"type\":\"uint256\"}],\"name\":\"mintForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedMarketTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedUniverseTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ReputationTokenBin is the compiled bytecode used for deploying new contracts.
const ReputationTokenBin = `0x60806040526002805460ff1916905560008054600160a060020a031916331790556123d08061002f6000396000f3006080604052600436106101875763ffffffff60e060020a60003504166306fdde03811461018c578063095ea7b31461021657806318160ddd1461024e578063238d35901461027557806323b872dd1461028a5780633018205f146102b4578063313ce567146102e5578063398c1a8914610310578063634eaff11461036c57806366188463146103815780636e7ce591146103a557806370a08231146103c957806375d9aa1a146103ea57806377469275146103ff578063870c426d1461041457806390fbf84e1461042957806391d76bbb1461045357806392eefe9b1461046857806395d89b4114610489578063a0c1ca341461049e578063a819515d146104c2578063a9059cbb146104d7578063b873e9a7146104fb578063bef72fa214610525578063c4d66de81461053a578063d73dd6231461055b578063d9d3e07d1461057f578063db054134146105a0578063db0a087c146105b8578063dd62ed3e146105cd578063ee89dab4146105f4578063f22b258a14610609578063fe98184d14610633575b600080fd5b34801561019857600080fd5b506101a161065d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101db5781810151838201526020016101c3565b50505050905090810190601f1680156102085780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561022257600080fd5b5061023a600160a060020a0360043516602435610694565b604080519115158252519081900360200190f35b34801561025a57600080fd5b506102636106ab565b60408051918252519081900360200190f35b34801561028157600080fd5b506102636106b1565b34801561029657600080fd5b5061023a600160a060020a03600435811690602435166044356106b7565b3480156102c057600080fd5b506102c96106cc565b60408051600160a060020a039092168252519081900360200190f35b3480156102f157600080fd5b506102fa6106db565b6040805160ff9092168252519081900360200190f35b34801561031c57600080fd5b506040805160206004803580820135838102808601850190965280855261023a953695939460249493850192918291850190849080828437509497505050508235151593505050602001356106e0565b34801561037857600080fd5b50610263610912565b34801561038d57600080fd5b5061023a600160a060020a0360043516602435610918565b3480156103b157600080fd5b5061023a600160a060020a036004351660243561097b565b3480156103d557600080fd5b50610263600160a060020a0360043516610a43565b3480156103f657600080fd5b5061023a610a5e565b34801561040b57600080fd5b506102c9610c7b565b34801561042057600080fd5b506102c9610d32565b34801561043557600080fd5b5061023a600160a060020a0360043581169060243516604435610d41565b34801561045f57600080fd5b50610263610dfd565b34801561047457600080fd5b5061023a600160a060020a0360043516610e03565b34801561049557600080fd5b506101a1610e4d565b3480156104aa57600080fd5b5061023a600160a060020a0360043516602435610e84565b3480156104ce57600080fd5b5061023a6112e3565b3480156104e357600080fd5b5061023a600160a060020a036004351660243561149e565b34801561050757600080fd5b5061023a600160a060020a03600435811690602435166044356114b1565b34801561053157600080fd5b5061026361152b565b34801561054657600080fd5b5061023a600160a060020a0360043516611531565b34801561056757600080fd5b5061023a600160a060020a036004351660243561159b565b34801561058b57600080fd5b5061023a600160a060020a03600435166115d7565b3480156105ac57600080fd5b5061023a60043561194e565b3480156105c457600080fd5b50610263611aba565b3480156105d957600080fd5b50610263600160a060020a0360043581169060243516611ade565b34801561060057600080fd5b5061023a611b09565b34801561061557600080fd5b5061023a600160a060020a0360043581169060243516604435611b12565b34801561063f57600080fd5b5061023a600160a060020a0360043581169060243516604435611b8c565b60408051808201909152600a81527f52657075746174696f6e00000000000000000000000000000000000000000000602082015281565b60006106a1338484611bb7565b5060019392505050565b60035490565b600a5490565b60006106c4848484611c22565b949350505050565b600054600160a060020a031690565b601281565b6002546000908190819060ff1615156106f857600080fd5b6000841161070557600080fd5b600654604080517fdf428e3b000000000000000000000000000000000000000000000000000000008152871515602482015260048101918252885160448201528851600160a060020a039093169263df428e3b928a928a928291606401906020808701910280838360005b83811015610788578181015183820152602001610770565b505050509050019350505050602060405180830381600087803b1580156107ae57600080fd5b505af11580156107c2573d6000803e3d6000fd5b505050506040513d60208110156107d857600080fd5b5051604080517fb80907f20000000000000000000000000000000000000000000000000000000081529051919350600160a060020a0384169163b80907f2916004808201926020929091908290030181600087803b15801561083957600080fd5b505af115801561084d573d6000803e3d6000fd5b505050506040513d602081101561086357600080fd5b505190506108713385611c8f565b50604080517fa0c1ca34000000000000000000000000000000000000000000000000000000008152336004820152602481018690529051600160a060020a0383169163a0c1ca349160448083019260209291908290030181600087803b1580156108da57600080fd5b505af11580156108ee573d6000803e3d6000fd5b505050506040513d602081101561090457600080fd5b506001979650505050505050565b60001981565b336000908152600560209081526040808320600160a060020a0386168452909152812054808311156109565761095033856000611bb7565b506106a1565b610970338561096b848763ffffffff611d3016565b611bb7565b505060019392505050565b60025460009060ff16151561098f57600080fd5b6000821161099c57600080fd5b6109a583611d45565b506109b03383611c8f565b50604080517fa0c1ca34000000000000000000000000000000000000000000000000000000008152336004820152602481018490529051600160a060020a0385169163a0c1ca349160448083019260209291908290030181600087803b158015610a1957600080fd5b505af1158015610a2d573d6000803e3d6000fd5b505050506040513d602081101561097057600080fd5b600160a060020a031660009081526004602052604090205490565b6002546000908190819060ff161515610a7657600080fd5b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4c656761637952657075746174696f6e546f6b656e000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b158015610afd57600080fd5b505af1158015610b11573d6000803e3d6000fd5b505050506040513d6020811015610b2757600080fd5b5051604080517f70a082310000000000000000000000000000000000000000000000000000000081523360048201529051919350600160a060020a038416916370a08231916024808201926020929091908290030181600087803b158015610b8e57600080fd5b505af1158015610ba2573d6000803e3d6000fd5b505050506040513d6020811015610bb857600080fd5b5051604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152600060248201819052604482018490529151929350600160a060020a038516926323b872dd92606480840193602093929083900390910190829087803b158015610c3057600080fd5b505af1158015610c44573d6000803e3d6000fd5b505050506040513d6020811015610c5a57600080fd5b50511515610c6757600080fd5b610c713382611edf565b5060019250505090565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4c656761637952657075746174696f6e546f6b656e000000000000000000000060048201529051600160a060020a039092169163f39ec1f79160248082019260209290919082900301818787803b158015610d0157600080fd5b505af1158015610d15573d6000803e3d6000fd5b505050506040513d6020811015610d2b57600080fd5b5051905090565b600654600160a060020a031690565b60025460009060ff161515610d5557600080fd5b600654604080517fc7c88d700000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163c7c88d70916024808201926020929091908290030181600087803b158015610dbb57600080fd5b505af1158015610dcf573d6000803e3d6000fd5b505050506040513d6020811015610de557600080fd5b50511515610df257600080fd5b6106c4848484611f80565b60075490565b60008054600160a060020a03163314610e1b57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60408051808201909152600381527f5245500000000000000000000000000000000000000000000000000000000000602082015281565b600254600090819060ff161515610e9a57600080fd5b600660009054906101000a9004600160a060020a0316600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610eed57600080fd5b505af1158015610f01573d6000803e3d6000fd5b505050506040513d6020811015610f1757600080fd5b5051604080517fb80907f20000000000000000000000000000000000000000000000000000000081529051919250600160a060020a0383169163b80907f2916004808201926020929091908290030181600087803b158015610f7857600080fd5b505af1158015610f8c573d6000803e3d6000fd5b505050506040513d6020811015610fa257600080fd5b5051600160a060020a03163314610fb857600080fd5b80600160a060020a03166377e71ee56040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610ff657600080fd5b505af115801561100a573d6000803e3d6000fd5b505050506040513d602081101561102057600080fd5b505160008054604080517f188ec3560000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169263188ec356926004808401936020939083900390910190829087803b15801561108357600080fd5b505af1158015611097573d6000803e3d6000fd5b505050506040513d60208110156110ad57600080fd5b5051106110b957600080fd5b6110c38484611edf565b506007805484019055604080517fcb1d84180000000000000000000000000000000000000000000000000000000081529051600160a060020a0383169163cb1d84189160048083019260209291908290030181600087803b15801561112757600080fd5b505af115801561113b573d6000803e3d6000fd5b505050506040513d602081101561115157600080fd5b5051604080517f8d4e40830000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921691638d4e4083916004808201926020929091908290030181600087803b1580156111b057600080fd5b505af11580156111c4573d6000803e3d6000fd5b505050506040513d60208110156111da57600080fd5b505115156106a15780600160a060020a031663f7095d9d600660009054906101000a9004600160a060020a0316600160a060020a031663c38c0fa76040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561124457600080fd5b505af1158015611258573d6000803e3d6000fd5b505050506040513d602081101561126e57600080fd5b50516040805160e060020a63ffffffff851602815260048101929092525160248083019260209291908290030181600087803b1580156112ad57600080fd5b505af11580156112c1573d6000803e3d6000fd5b505050506040513d60208110156112d757600080fd5b50600195945050505050565b600080600660009054906101000a9004600160a060020a0316600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561133957600080fd5b505af115801561134d573d6000803e3d6000fd5b505050506040513d602081101561136357600080fd5b5051600954600a80549190910390559050600160a060020a03811615156113945761138c61204b565b60095561148b565b80600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156113d257600080fd5b505af11580156113e6573d6000803e3d6000fd5b505050506040513d60208110156113fc57600080fd5b5051604080517f238d35900000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169163238d3590916004808201926020929091908290030181600087803b15801561145b57600080fd5b505af115801561146f573d6000803e3d6000fd5b505050506040513d602081101561148557600080fd5b50516009555b5050600954600a80549091019055600190565b60006114aa838361205a565b9392505050565b60025460009060ff1615156114c557600080fd5b600654604080517ff76514c70000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163f76514c7916024808201926020929091908290030181600087803b158015610dbb57600080fd5b60015481565b60025460009060ff161561154457600080fd5b61154c612067565b50600160a060020a038216151561156257600080fd5b6006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790556115926112e3565b50600192915050565b336000818152600560209081526040808320600160a060020a038716845290915281205490916106a191859061096b908663ffffffff61208e16565b600080600160a060020a0383163014156115f057600080fd5b82600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561162e57600080fd5b505af1158015611642573d6000803e3d6000fd5b505050506040513d602081101561165857600080fd5b5051600654604080517fa63f13500000000000000000000000000000000000000000000000000000000081529051929350600160a060020a039091169163a63f1350916004808201926020929091908290030181600087803b1580156116bd57600080fd5b505af11580156116d1573d6000803e3d6000fd5b505050506040513d60208110156116e757600080fd5b5051604080517fc38c0fa70000000000000000000000000000000000000000000000000000000081529051600160a060020a039283169263eceba876929085169163c38c0fa7916004808201926020929091908290030181600087803b15801561175057600080fd5b505af1158015611764573d6000803e3d6000fd5b505050506040513d602081101561177a57600080fd5b50516040805160e060020a63ffffffff851602815260048101929092525160248083019260209291908290030181600087803b1580156117b957600080fd5b505af11580156117cd573d6000803e3d6000fd5b505050506040513d60208110156117e357600080fd5b5051604080517fb80907f20000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169163b80907f2916004808201926020929091908290030181600087803b15801561184257600080fd5b505af1158015611856573d6000803e3d6000fd5b505050506040513d602081101561186c57600080fd5b5051600160a060020a0384811691161461188557600080fd5b600160a060020a038316600081815260086020908152604080832054600a8054909101905580517f91d76bbb00000000000000000000000000000000000000000000000000000000815290516391d76bbb93600480840194938390030190829087803b1580156118f457600080fd5b505af1158015611908573d6000803e3d6000fd5b505050506040513d602081101561191e57600080fd5b5051600160a060020a039390931660009081526008602052604090208390555050600a8054919091039055600190565b60025460009081908190819060ff16151561196857600080fd5b600660009054906101000a9004600160a060020a0316600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156119bb57600080fd5b505af11580156119cf573d6000803e3d6000fd5b505050506040513d60208110156119e557600080fd5b5051604080517ff76514c700000000000000000000000000000000000000000000000000000000815233600482018190529151929550909350600160a060020a0385169163f76514c7916024808201926020929091908290030181600087803b158015611a5157600080fd5b505af1158015611a65573d6000803e3d6000fd5b505050506040513d6020811015611a7b57600080fd5b50511515611a8857600080fd5b611a9985600263ffffffff6120a016565b9050611aa58282611edf565b50600a80548201905560019350505050919050565b7f52657075746174696f6e546f6b656e000000000000000000000000000000000090565b600160a060020a03918216600090815260056020908152604080832093909416825291909152205490565b60025460ff1690565b60025460009060ff161515611b2657600080fd5b600654604080517f9f7e1bf60000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691639f7e1bf6916024808201926020929091908290030181600087803b158015610dbb57600080fd5b60025460009060ff161515611ba057600080fd5b600654600160a060020a03163314610df257600080fd5b600160a060020a03808416600081815260056020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a03831660009081526005602090815260408083203384529091528120546000198114611c8457611c5f818463ffffffff611d3016565b600160a060020a03861660009081526005602090815260408083203384529091529020555b6112d7858585611f80565b600160a060020a038216600090815260046020526040812054611cb8908363ffffffff611d3016565b600160a060020a038416600090815260046020526040902055600354611ce4908363ffffffff611d3016565b600355604080518381529051600160a060020a038516917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a26106a183836120b7565b600082821115611d3f57600080fd5b50900390565b600080600083600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611d8857600080fd5b505af1158015611d9c573d6000803e3d6000fd5b505050506040513d6020811015611db257600080fd5b5051600654604080517f9517317c000000000000000000000000000000000000000000000000000000008152600160a060020a0380851660048301529151939550911691639517317c916024808201926020929091908290030181600087803b158015611e1e57600080fd5b505af1158015611e32573d6000803e3d6000fd5b505050506040513d6020811015611e4857600080fd5b50511515611e5557600080fd5b81905083600160a060020a031681600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611ea057600080fd5b505af1158015611eb4573d6000803e3d6000fd5b505050506040513d6020811015611eca57600080fd5b5051600160a060020a0316146106a157600080fd5b600160a060020a038216600090815260046020526040812054611f08908363ffffffff61208e16565b600160a060020a038416600090815260046020526040902055600354611f34908363ffffffff61208e16565b600355604080518381529051600160a060020a038516917f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885919081900360200190a26106a183836121ae565b600160a060020a038316600090815260046020526040812054611fa9908363ffffffff611d3016565b600160a060020a038086166000908152600460205260408082209390935590851681522054611fde908363ffffffff61208e16565b600160a060020a0380851660008181526004602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a36120408484846122a5565b506001949350505050565b6a09195731e2ce35eb00000090565b60006114aa338484611f80565b60025460009060ff161561207a57600080fd5b506002805460ff1916600190811790915590565b6000828201838110156114aa57600080fd5b60008082848115156120ae57fe5b04949350505050565b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561210b57600080fd5b505af115801561211f573d6000803e3d6000fd5b505050506040513d602081101561213557600080fd5b5051600654604080517f4405a339000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015286831660248201526044810186905290519190921691634405a3399160648083019260209291908290030181600087803b158015610a1957600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561220257600080fd5b505af1158015612216573d6000803e3d6000fd5b505050506040513d602081101561222c57600080fd5b5051600654604080517f79fff7a9000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152868316602482015260448101869052905191909216916379fff7a99160648083019260209291908290030181600087803b158015610a1957600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156122f957600080fd5b505af115801561230d573d6000803e3d6000fd5b505050506040513d602081101561232357600080fd5b5051600654604080517fec37a6e4000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015287831660248201528683166044820152606481018690529051919092169163ec37a6e49160848083019260209291908290030181600087803b1580156112ad57600080fd00a165627a7a723058201d96aaa81ff21ed6d331b39a77c5a7837a3e2cf2ac04bc3036160efd332680e40029`

// DeployReputationToken deploys a new Ethereum contract, binding an instance of ReputationToken to it.
func DeployReputationToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReputationToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ReputationTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReputationTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReputationToken{ReputationTokenCaller: ReputationTokenCaller{contract: contract}, ReputationTokenTransactor: ReputationTokenTransactor{contract: contract}, ReputationTokenFilterer: ReputationTokenFilterer{contract: contract}}, nil
}

// ReputationToken is an auto generated Go binding around an Ethereum contract.
type ReputationToken struct {
	ReputationTokenCaller     // Read-only binding to the contract
	ReputationTokenTransactor // Write-only binding to the contract
	ReputationTokenFilterer   // Log filterer for contract events
}

// ReputationTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationTokenSession struct {
	Contract     *ReputationToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReputationTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationTokenCallerSession struct {
	Contract *ReputationTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReputationTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationTokenTransactorSession struct {
	Contract     *ReputationTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReputationTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationTokenRaw struct {
	Contract *ReputationToken // Generic contract binding to access the raw methods on
}

// ReputationTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationTokenCallerRaw struct {
	Contract *ReputationTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationTokenTransactorRaw struct {
	Contract *ReputationTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputationToken creates a new instance of ReputationToken, bound to a specific deployed contract.
func NewReputationToken(address common.Address, backend bind.ContractBackend) (*ReputationToken, error) {
	contract, err := bindReputationToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReputationToken{ReputationTokenCaller: ReputationTokenCaller{contract: contract}, ReputationTokenTransactor: ReputationTokenTransactor{contract: contract}, ReputationTokenFilterer: ReputationTokenFilterer{contract: contract}}, nil
}

// NewReputationTokenCaller creates a new read-only instance of ReputationToken, bound to a specific deployed contract.
func NewReputationTokenCaller(address common.Address, caller bind.ContractCaller) (*ReputationTokenCaller, error) {
	contract, err := bindReputationToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenCaller{contract: contract}, nil
}

// NewReputationTokenTransactor creates a new write-only instance of ReputationToken, bound to a specific deployed contract.
func NewReputationTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationTokenTransactor, error) {
	contract, err := bindReputationToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenTransactor{contract: contract}, nil
}

// NewReputationTokenFilterer creates a new log filterer instance of ReputationToken, bound to a specific deployed contract.
func NewReputationTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationTokenFilterer, error) {
	contract, err := bindReputationToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenFilterer{contract: contract}, nil
}

// bindReputationToken binds a generic wrapper to an already deployed contract.
func bindReputationToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReputationTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationToken *ReputationTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReputationToken.Contract.ReputationTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationToken *ReputationTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationToken.Contract.ReputationTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationToken *ReputationTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationToken.Contract.ReputationTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationToken *ReputationTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReputationToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationToken *ReputationTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationToken *ReputationTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationToken.Contract.contract.Transact(opts, method, params...)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_ReputationToken *ReputationTokenCaller) ETERNALAPPROVALVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "ETERNAL_APPROVAL_VALUE")
	return *ret0, err
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_ReputationToken *ReputationTokenSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _ReputationToken.Contract.ETERNALAPPROVALVALUE(&_ReputationToken.CallOpts)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_ReputationToken *ReputationTokenCallerSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _ReputationToken.Contract.ETERNALAPPROVALVALUE(&_ReputationToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReputationToken *ReputationTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReputationToken *ReputationTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ReputationToken.Contract.Allowance(&_ReputationToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReputationToken *ReputationTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ReputationToken.Contract.Allowance(&_ReputationToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ReputationToken *ReputationTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ReputationToken *ReputationTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ReputationToken.Contract.BalanceOf(&_ReputationToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ReputationToken *ReputationTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ReputationToken.Contract.BalanceOf(&_ReputationToken.CallOpts, _owner)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_ReputationToken *ReputationTokenCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_ReputationToken *ReputationTokenSession) ControllerLookupName() ([32]byte, error) {
	return _ReputationToken.Contract.ControllerLookupName(&_ReputationToken.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_ReputationToken *ReputationTokenCallerSession) ControllerLookupName() ([32]byte, error) {
	return _ReputationToken.Contract.ControllerLookupName(&_ReputationToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ReputationToken *ReputationTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ReputationToken *ReputationTokenSession) Decimals() (uint8, error) {
	return _ReputationToken.Contract.Decimals(&_ReputationToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ReputationToken *ReputationTokenCallerSession) Decimals() (uint8, error) {
	return _ReputationToken.Contract.Decimals(&_ReputationToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ReputationToken *ReputationTokenCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ReputationToken *ReputationTokenSession) GetController() (common.Address, error) {
	return _ReputationToken.Contract.GetController(&_ReputationToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ReputationToken *ReputationTokenCallerSession) GetController() (common.Address, error) {
	return _ReputationToken.Contract.GetController(&_ReputationToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_ReputationToken *ReputationTokenCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_ReputationToken *ReputationTokenSession) GetInitialized() (bool, error) {
	return _ReputationToken.Contract.GetInitialized(&_ReputationToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_ReputationToken *ReputationTokenCallerSession) GetInitialized() (bool, error) {
	return _ReputationToken.Contract.GetInitialized(&_ReputationToken.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() constant returns(address)
func (_ReputationToken *ReputationTokenCaller) GetLegacyRepToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "getLegacyRepToken")
	return *ret0, err
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() constant returns(address)
func (_ReputationToken *ReputationTokenSession) GetLegacyRepToken() (common.Address, error) {
	return _ReputationToken.Contract.GetLegacyRepToken(&_ReputationToken.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() constant returns(address)
func (_ReputationToken *ReputationTokenCallerSession) GetLegacyRepToken() (common.Address, error) {
	return _ReputationToken.Contract.GetLegacyRepToken(&_ReputationToken.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_ReputationToken *ReputationTokenCaller) GetTotalMigrated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "getTotalMigrated")
	return *ret0, err
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_ReputationToken *ReputationTokenSession) GetTotalMigrated() (*big.Int, error) {
	return _ReputationToken.Contract.GetTotalMigrated(&_ReputationToken.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_ReputationToken *ReputationTokenCallerSession) GetTotalMigrated() (*big.Int, error) {
	return _ReputationToken.Contract.GetTotalMigrated(&_ReputationToken.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_ReputationToken *ReputationTokenCaller) GetTotalTheoreticalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "getTotalTheoreticalSupply")
	return *ret0, err
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_ReputationToken *ReputationTokenSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _ReputationToken.Contract.GetTotalTheoreticalSupply(&_ReputationToken.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_ReputationToken *ReputationTokenCallerSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _ReputationToken.Contract.GetTotalTheoreticalSupply(&_ReputationToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ReputationToken *ReputationTokenCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ReputationToken *ReputationTokenSession) GetTypeName() ([32]byte, error) {
	return _ReputationToken.Contract.GetTypeName(&_ReputationToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ReputationToken *ReputationTokenCallerSession) GetTypeName() ([32]byte, error) {
	return _ReputationToken.Contract.GetTypeName(&_ReputationToken.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_ReputationToken *ReputationTokenCaller) GetUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "getUniverse")
	return *ret0, err
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_ReputationToken *ReputationTokenSession) GetUniverse() (common.Address, error) {
	return _ReputationToken.Contract.GetUniverse(&_ReputationToken.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_ReputationToken *ReputationTokenCallerSession) GetUniverse() (common.Address, error) {
	return _ReputationToken.Contract.GetUniverse(&_ReputationToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReputationToken *ReputationTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReputationToken *ReputationTokenSession) Name() (string, error) {
	return _ReputationToken.Contract.Name(&_ReputationToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReputationToken *ReputationTokenCallerSession) Name() (string, error) {
	return _ReputationToken.Contract.Name(&_ReputationToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ReputationToken *ReputationTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ReputationToken *ReputationTokenSession) Symbol() (string, error) {
	return _ReputationToken.Contract.Symbol(&_ReputationToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ReputationToken *ReputationTokenCallerSession) Symbol() (string, error) {
	return _ReputationToken.Contract.Symbol(&_ReputationToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReputationToken *ReputationTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReputationToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReputationToken *ReputationTokenSession) TotalSupply() (*big.Int, error) {
	return _ReputationToken.Contract.TotalSupply(&_ReputationToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReputationToken *ReputationTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ReputationToken.Contract.TotalSupply(&_ReputationToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.Approve(&_ReputationToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.Approve(&_ReputationToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.DecreaseApproval(&_ReputationToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.DecreaseApproval(&_ReputationToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.IncreaseApproval(&_ReputationToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.IncreaseApproval(&_ReputationToken.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) Initialize(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "initialize", _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_ReputationToken *ReputationTokenSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _ReputationToken.Contract.Initialize(&_ReputationToken.TransactOpts, _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _ReputationToken.Contract.Initialize(&_ReputationToken.TransactOpts, _universe)
}

// MigrateFromLegacyReputationToken is a paid mutator transaction binding the contract method 0x75d9aa1a.
//
// Solidity: function migrateFromLegacyReputationToken() returns(bool)
func (_ReputationToken *ReputationTokenTransactor) MigrateFromLegacyReputationToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "migrateFromLegacyReputationToken")
}

// MigrateFromLegacyReputationToken is a paid mutator transaction binding the contract method 0x75d9aa1a.
//
// Solidity: function migrateFromLegacyReputationToken() returns(bool)
func (_ReputationToken *ReputationTokenSession) MigrateFromLegacyReputationToken() (*types.Transaction, error) {
	return _ReputationToken.Contract.MigrateFromLegacyReputationToken(&_ReputationToken.TransactOpts)
}

// MigrateFromLegacyReputationToken is a paid mutator transaction binding the contract method 0x75d9aa1a.
//
// Solidity: function migrateFromLegacyReputationToken() returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) MigrateFromLegacyReputationToken() (*types.Transaction, error) {
	return _ReputationToken.Contract.MigrateFromLegacyReputationToken(&_ReputationToken.TransactOpts)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) MigrateIn(opts *bind.TransactOpts, _reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "migrateIn", _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.MigrateIn(&_ReputationToken.TransactOpts, _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.MigrateIn(&_ReputationToken.TransactOpts, _reporter, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) MigrateOut(opts *bind.TransactOpts, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "migrateOut", _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.MigrateOut(&_ReputationToken.TransactOpts, _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.MigrateOut(&_ReputationToken.TransactOpts, _destination, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x398c1a89.
//
// Solidity: function migrateOutByPayout(_payoutNumerators uint256[], _invalid bool, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) MigrateOutByPayout(opts *bind.TransactOpts, _payoutNumerators []*big.Int, _invalid bool, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "migrateOutByPayout", _payoutNumerators, _invalid, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x398c1a89.
//
// Solidity: function migrateOutByPayout(_payoutNumerators uint256[], _invalid bool, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) MigrateOutByPayout(_payoutNumerators []*big.Int, _invalid bool, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.MigrateOutByPayout(&_ReputationToken.TransactOpts, _payoutNumerators, _invalid, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x398c1a89.
//
// Solidity: function migrateOutByPayout(_payoutNumerators uint256[], _invalid bool, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) MigrateOutByPayout(_payoutNumerators []*big.Int, _invalid bool, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.MigrateOutByPayout(&_ReputationToken.TransactOpts, _payoutNumerators, _invalid, _attotokens)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) MintForReportingParticipant(opts *bind.TransactOpts, _amountMigrated *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "mintForReportingParticipant", _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.MintForReportingParticipant(&_ReputationToken.TransactOpts, _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.MintForReportingParticipant(&_ReputationToken.TransactOpts, _amountMigrated)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ReputationToken *ReputationTokenSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ReputationToken.Contract.SetController(&_ReputationToken.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ReputationToken.Contract.SetController(&_ReputationToken.TransactOpts, _controller)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.Transfer(&_ReputationToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.Transfer(&_ReputationToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TransferFrom(&_ReputationToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TransferFrom(&_ReputationToken.TransactOpts, _from, _to, _value)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) TrustedFeeWindowTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "trustedFeeWindowTransfer", _source, _destination, _attotokens)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) TrustedFeeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TrustedFeeWindowTransfer(&_ReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) TrustedFeeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TrustedFeeWindowTransfer(&_ReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) TrustedMarketTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "trustedMarketTransfer", _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TrustedMarketTransfer(&_ReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TrustedMarketTransfer(&_ReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) TrustedReportingParticipantTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "trustedReportingParticipantTransfer", _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TrustedReportingParticipantTransfer(&_ReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TrustedReportingParticipantTransfer(&_ReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) TrustedUniverseTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "trustedUniverseTransfer", _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TrustedUniverseTransfer(&_ReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ReputationToken.Contract.TrustedUniverseTransfer(&_ReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_ReputationToken *ReputationTokenTransactor) UpdateParentTotalTheoreticalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "updateParentTotalTheoreticalSupply")
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_ReputationToken *ReputationTokenSession) UpdateParentTotalTheoreticalSupply() (*types.Transaction, error) {
	return _ReputationToken.Contract.UpdateParentTotalTheoreticalSupply(&_ReputationToken.TransactOpts)
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) UpdateParentTotalTheoreticalSupply() (*types.Transaction, error) {
	return _ReputationToken.Contract.UpdateParentTotalTheoreticalSupply(&_ReputationToken.TransactOpts)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(_token address) returns(bool)
func (_ReputationToken *ReputationTokenTransactor) UpdateSiblingMigrationTotal(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ReputationToken.contract.Transact(opts, "updateSiblingMigrationTotal", _token)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(_token address) returns(bool)
func (_ReputationToken *ReputationTokenSession) UpdateSiblingMigrationTotal(_token common.Address) (*types.Transaction, error) {
	return _ReputationToken.Contract.UpdateSiblingMigrationTotal(&_ReputationToken.TransactOpts, _token)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(_token address) returns(bool)
func (_ReputationToken *ReputationTokenTransactorSession) UpdateSiblingMigrationTotal(_token common.Address) (*types.Transaction, error) {
	return _ReputationToken.Contract.UpdateSiblingMigrationTotal(&_ReputationToken.TransactOpts, _token)
}

// ReputationTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ReputationToken contract.
type ReputationTokenApprovalIterator struct {
	Event *ReputationTokenApproval // Event containing the contract specifics and raw log

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
func (it *ReputationTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationTokenApproval)
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
		it.Event = new(ReputationTokenApproval)
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
func (it *ReputationTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationTokenApproval represents a Approval event raised by the ReputationToken contract.
type ReputationTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ReputationToken *ReputationTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ReputationTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ReputationToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenApprovalIterator{contract: _ReputationToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ReputationToken *ReputationTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ReputationTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ReputationToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationTokenApproval)
				if err := _ReputationToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ReputationTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ReputationToken contract.
type ReputationTokenBurnIterator struct {
	Event *ReputationTokenBurn // Event containing the contract specifics and raw log

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
func (it *ReputationTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationTokenBurn)
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
		it.Event = new(ReputationTokenBurn)
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
func (it *ReputationTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationTokenBurn represents a Burn event raised by the ReputationToken contract.
type ReputationTokenBurn struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_ReputationToken *ReputationTokenFilterer) FilterBurn(opts *bind.FilterOpts, target []common.Address) (*ReputationTokenBurnIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ReputationToken.contract.FilterLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenBurnIterator{contract: _ReputationToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_ReputationToken *ReputationTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ReputationTokenBurn, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ReputationToken.contract.WatchLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationTokenBurn)
				if err := _ReputationToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ReputationTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ReputationToken contract.
type ReputationTokenMintIterator struct {
	Event *ReputationTokenMint // Event containing the contract specifics and raw log

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
func (it *ReputationTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationTokenMint)
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
		it.Event = new(ReputationTokenMint)
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
func (it *ReputationTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationTokenMint represents a Mint event raised by the ReputationToken contract.
type ReputationTokenMint struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_ReputationToken *ReputationTokenFilterer) FilterMint(opts *bind.FilterOpts, target []common.Address) (*ReputationTokenMintIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ReputationToken.contract.FilterLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenMintIterator{contract: _ReputationToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_ReputationToken *ReputationTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ReputationTokenMint, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ReputationToken.contract.WatchLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationTokenMint)
				if err := _ReputationToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ReputationTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ReputationToken contract.
type ReputationTokenTransferIterator struct {
	Event *ReputationTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ReputationTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationTokenTransfer)
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
		it.Event = new(ReputationTokenTransfer)
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
func (it *ReputationTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationTokenTransfer represents a Transfer event raised by the ReputationToken contract.
type ReputationTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ReputationToken *ReputationTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ReputationTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ReputationToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenTransferIterator{contract: _ReputationToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ReputationToken *ReputationTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ReputationTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ReputationToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationTokenTransfer)
				if err := _ReputationToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ReputationTokenFactoryABI is the input ABI used to generate the binding from.
const ReputationTokenFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"createReputationToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ReputationTokenFactoryBin is the compiled bytecode used for deploying new contracts.
const ReputationTokenFactoryBin = `0x608060405234801561001057600080fd5b50610529806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663812f82d98114610045575b600080fd5b34801561005157600080fd5b5061007973ffffffffffffffffffffffffffffffffffffffff600435811690602435166100a2565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000806000846100b06101e5565b73ffffffffffffffffffffffffffffffffffffffff90911681527f52657075746174696f6e546f6b656e00000000000000000000000000000000006020820152604080519182900301906000f08015801561010f573d6000803e3d6000fd5b5091508190508073ffffffffffffffffffffffffffffffffffffffff1663c4d66de8856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156101b057600080fd5b505af11580156101c4573d6000803e3d6000fd5b505050506040513d60208110156101da57600080fd5b509095945050505050565b604051610308806101f6833901905600608060405234801561001057600080fd5b506040516040806103088339810160405280516020909101516000805433600160a060020a03199182161716600160a060020a039093169290921782556001556102a890819061006090396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461014f57806392eefe9b1461018d578063bef72fa2146101cf575b60015460009015156100675761014c565b60008054600154604080517ff39ec1f700000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff9092169263f39ec1f7926024808401936020939083900390910190829087803b1580156100df57600080fd5b505af11580156100f3573d6000803e3d6000fd5b505050506040513d602081101561010957600080fd5b50516040805136601f8101601f191680830190935292935091600083376000803684865af4808015610147576040513d81016040523d6000823e3d81f35b600080fd5b50005b34801561015b57600080fd5b506101646101f6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019957600080fd5b506101bb73ffffffffffffffffffffffffffffffffffffffff60043516610212565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610276565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461023757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a7230582017c3a6b57ac8bbc7fd5320cba2543fe371eb8726332cd58700252790e09cf2140029a165627a7a72305820b71bbd72d851aed15eb8ca465f1397fc5cf20445facf1ac5429fdadee99440c00029`

// DeployReputationTokenFactory deploys a new Ethereum contract, binding an instance of ReputationTokenFactory to it.
func DeployReputationTokenFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReputationTokenFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ReputationTokenFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReputationTokenFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReputationTokenFactory{ReputationTokenFactoryCaller: ReputationTokenFactoryCaller{contract: contract}, ReputationTokenFactoryTransactor: ReputationTokenFactoryTransactor{contract: contract}, ReputationTokenFactoryFilterer: ReputationTokenFactoryFilterer{contract: contract}}, nil
}

// ReputationTokenFactory is an auto generated Go binding around an Ethereum contract.
type ReputationTokenFactory struct {
	ReputationTokenFactoryCaller     // Read-only binding to the contract
	ReputationTokenFactoryTransactor // Write-only binding to the contract
	ReputationTokenFactoryFilterer   // Log filterer for contract events
}

// ReputationTokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationTokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationTokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationTokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationTokenFactorySession struct {
	Contract     *ReputationTokenFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ReputationTokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationTokenFactoryCallerSession struct {
	Contract *ReputationTokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ReputationTokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationTokenFactoryTransactorSession struct {
	Contract     *ReputationTokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ReputationTokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationTokenFactoryRaw struct {
	Contract *ReputationTokenFactory // Generic contract binding to access the raw methods on
}

// ReputationTokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationTokenFactoryCallerRaw struct {
	Contract *ReputationTokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationTokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationTokenFactoryTransactorRaw struct {
	Contract *ReputationTokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputationTokenFactory creates a new instance of ReputationTokenFactory, bound to a specific deployed contract.
func NewReputationTokenFactory(address common.Address, backend bind.ContractBackend) (*ReputationTokenFactory, error) {
	contract, err := bindReputationTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenFactory{ReputationTokenFactoryCaller: ReputationTokenFactoryCaller{contract: contract}, ReputationTokenFactoryTransactor: ReputationTokenFactoryTransactor{contract: contract}, ReputationTokenFactoryFilterer: ReputationTokenFactoryFilterer{contract: contract}}, nil
}

// NewReputationTokenFactoryCaller creates a new read-only instance of ReputationTokenFactory, bound to a specific deployed contract.
func NewReputationTokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*ReputationTokenFactoryCaller, error) {
	contract, err := bindReputationTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenFactoryCaller{contract: contract}, nil
}

// NewReputationTokenFactoryTransactor creates a new write-only instance of ReputationTokenFactory, bound to a specific deployed contract.
func NewReputationTokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationTokenFactoryTransactor, error) {
	contract, err := bindReputationTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenFactoryTransactor{contract: contract}, nil
}

// NewReputationTokenFactoryFilterer creates a new log filterer instance of ReputationTokenFactory, bound to a specific deployed contract.
func NewReputationTokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationTokenFactoryFilterer, error) {
	contract, err := bindReputationTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationTokenFactoryFilterer{contract: contract}, nil
}

// bindReputationTokenFactory binds a generic wrapper to an already deployed contract.
func bindReputationTokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReputationTokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationTokenFactory *ReputationTokenFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReputationTokenFactory.Contract.ReputationTokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationTokenFactory *ReputationTokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationTokenFactory.Contract.ReputationTokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationTokenFactory *ReputationTokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationTokenFactory.Contract.ReputationTokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationTokenFactory *ReputationTokenFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReputationTokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationTokenFactory *ReputationTokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationTokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationTokenFactory *ReputationTokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationTokenFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateReputationToken is a paid mutator transaction binding the contract method 0x812f82d9.
//
// Solidity: function createReputationToken(_controller address, _universe address) returns(address)
func (_ReputationTokenFactory *ReputationTokenFactoryTransactor) CreateReputationToken(opts *bind.TransactOpts, _controller common.Address, _universe common.Address) (*types.Transaction, error) {
	return _ReputationTokenFactory.contract.Transact(opts, "createReputationToken", _controller, _universe)
}

// CreateReputationToken is a paid mutator transaction binding the contract method 0x812f82d9.
//
// Solidity: function createReputationToken(_controller address, _universe address) returns(address)
func (_ReputationTokenFactory *ReputationTokenFactorySession) CreateReputationToken(_controller common.Address, _universe common.Address) (*types.Transaction, error) {
	return _ReputationTokenFactory.Contract.CreateReputationToken(&_ReputationTokenFactory.TransactOpts, _controller, _universe)
}

// CreateReputationToken is a paid mutator transaction binding the contract method 0x812f82d9.
//
// Solidity: function createReputationToken(_controller address, _universe address) returns(address)
func (_ReputationTokenFactory *ReputationTokenFactoryTransactorSession) CreateReputationToken(_controller common.Address, _universe common.Address) (*types.Transaction, error) {
	return _ReputationTokenFactory.Contract.CreateReputationToken(&_ReputationTokenFactory.TransactOpts, _controller, _universe)
}
