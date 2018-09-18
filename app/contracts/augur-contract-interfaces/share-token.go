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

// IShareTokenABI is the input ABI used to generate the binding from.
const IShareTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createShares\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedFillOrderTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedOrderTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedCancelOrderTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOutcome\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"destroyShares\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// IShareTokenBin is the compiled bytecode used for deploying new contracts.
const IShareTokenBin = `0x`

// DeployIShareToken deploys a new Ethereum contract, binding an instance of IShareToken to it.
func DeployIShareToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IShareToken, error) {
	parsed, err := abi.JSON(strings.NewReader(IShareTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IShareTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IShareToken{IShareTokenCaller: IShareTokenCaller{contract: contract}, IShareTokenTransactor: IShareTokenTransactor{contract: contract}, IShareTokenFilterer: IShareTokenFilterer{contract: contract}}, nil
}

// IShareToken is an auto generated Go binding around an Ethereum contract.
type IShareToken struct {
	IShareTokenCaller     // Read-only binding to the contract
	IShareTokenTransactor // Write-only binding to the contract
	IShareTokenFilterer   // Log filterer for contract events
}

// IShareTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IShareTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShareTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IShareTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShareTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IShareTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShareTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IShareTokenSession struct {
	Contract     *IShareToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IShareTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IShareTokenCallerSession struct {
	Contract *IShareTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IShareTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IShareTokenTransactorSession struct {
	Contract     *IShareTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IShareTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IShareTokenRaw struct {
	Contract *IShareToken // Generic contract binding to access the raw methods on
}

// IShareTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IShareTokenCallerRaw struct {
	Contract *IShareTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IShareTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IShareTokenTransactorRaw struct {
	Contract *IShareTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIShareToken creates a new instance of IShareToken, bound to a specific deployed contract.
func NewIShareToken(address common.Address, backend bind.ContractBackend) (*IShareToken, error) {
	contract, err := bindIShareToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IShareToken{IShareTokenCaller: IShareTokenCaller{contract: contract}, IShareTokenTransactor: IShareTokenTransactor{contract: contract}, IShareTokenFilterer: IShareTokenFilterer{contract: contract}}, nil
}

// NewIShareTokenCaller creates a new read-only instance of IShareToken, bound to a specific deployed contract.
func NewIShareTokenCaller(address common.Address, caller bind.ContractCaller) (*IShareTokenCaller, error) {
	contract, err := bindIShareToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IShareTokenCaller{contract: contract}, nil
}

// NewIShareTokenTransactor creates a new write-only instance of IShareToken, bound to a specific deployed contract.
func NewIShareTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IShareTokenTransactor, error) {
	contract, err := bindIShareToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IShareTokenTransactor{contract: contract}, nil
}

// NewIShareTokenFilterer creates a new log filterer instance of IShareToken, bound to a specific deployed contract.
func NewIShareTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IShareTokenFilterer, error) {
	contract, err := bindIShareToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IShareTokenFilterer{contract: contract}, nil
}

// bindIShareToken binds a generic wrapper to an already deployed contract.
func bindIShareToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IShareTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IShareToken *IShareTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IShareToken.Contract.IShareTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IShareToken *IShareTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IShareToken.Contract.IShareTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IShareToken *IShareTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IShareToken.Contract.IShareTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IShareToken *IShareTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IShareToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IShareToken *IShareTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IShareToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IShareToken *IShareTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IShareToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IShareToken *IShareTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IShareToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IShareToken *IShareTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IShareToken.Contract.Allowance(&_IShareToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IShareToken *IShareTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IShareToken.Contract.Allowance(&_IShareToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IShareToken *IShareTokenCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IShareToken.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IShareToken *IShareTokenSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _IShareToken.Contract.BalanceOf(&_IShareToken.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IShareToken *IShareTokenCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _IShareToken.Contract.BalanceOf(&_IShareToken.CallOpts, _who)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IShareToken *IShareTokenCaller) GetMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IShareToken.contract.Call(opts, out, "getMarket")
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IShareToken *IShareTokenSession) GetMarket() (common.Address, error) {
	return _IShareToken.Contract.GetMarket(&_IShareToken.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_IShareToken *IShareTokenCallerSession) GetMarket() (common.Address, error) {
	return _IShareToken.Contract.GetMarket(&_IShareToken.CallOpts)
}

// GetOutcome is a free data retrieval call binding the contract method 0x7e7e4b47.
//
// Solidity: function getOutcome() constant returns(uint256)
func (_IShareToken *IShareTokenCaller) GetOutcome(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IShareToken.contract.Call(opts, out, "getOutcome")
	return *ret0, err
}

// GetOutcome is a free data retrieval call binding the contract method 0x7e7e4b47.
//
// Solidity: function getOutcome() constant returns(uint256)
func (_IShareToken *IShareTokenSession) GetOutcome() (*big.Int, error) {
	return _IShareToken.Contract.GetOutcome(&_IShareToken.CallOpts)
}

// GetOutcome is a free data retrieval call binding the contract method 0x7e7e4b47.
//
// Solidity: function getOutcome() constant returns(uint256)
func (_IShareToken *IShareTokenCallerSession) GetOutcome() (*big.Int, error) {
	return _IShareToken.Contract.GetOutcome(&_IShareToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IShareToken *IShareTokenCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IShareToken.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IShareToken *IShareTokenSession) GetTypeName() ([32]byte, error) {
	return _IShareToken.Contract.GetTypeName(&_IShareToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IShareToken *IShareTokenCallerSession) GetTypeName() ([32]byte, error) {
	return _IShareToken.Contract.GetTypeName(&_IShareToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IShareToken *IShareTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IShareToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IShareToken *IShareTokenSession) TotalSupply() (*big.Int, error) {
	return _IShareToken.Contract.TotalSupply(&_IShareToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IShareToken *IShareTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IShareToken.Contract.TotalSupply(&_IShareToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IShareToken *IShareTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IShareToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IShareToken *IShareTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.Approve(&_IShareToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IShareToken *IShareTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.Approve(&_IShareToken.TransactOpts, _spender, _value)
}

// CreateShares is a paid mutator transaction binding the contract method 0x0f9e5bbd.
//
// Solidity: function createShares(_owner address, _amount uint256) returns(bool)
func (_IShareToken *IShareTokenTransactor) CreateShares(opts *bind.TransactOpts, _owner common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IShareToken.contract.Transact(opts, "createShares", _owner, _amount)
}

// CreateShares is a paid mutator transaction binding the contract method 0x0f9e5bbd.
//
// Solidity: function createShares(_owner address, _amount uint256) returns(bool)
func (_IShareToken *IShareTokenSession) CreateShares(_owner common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.CreateShares(&_IShareToken.TransactOpts, _owner, _amount)
}

// CreateShares is a paid mutator transaction binding the contract method 0x0f9e5bbd.
//
// Solidity: function createShares(_owner address, _amount uint256) returns(bool)
func (_IShareToken *IShareTokenTransactorSession) CreateShares(_owner common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.CreateShares(&_IShareToken.TransactOpts, _owner, _amount)
}

// DestroyShares is a paid mutator transaction binding the contract method 0xd333d7cf.
//
// Solidity: function destroyShares( address, balance uint256) returns(bool)
func (_IShareToken *IShareTokenTransactor) DestroyShares(opts *bind.TransactOpts, arg0 common.Address, balance *big.Int) (*types.Transaction, error) {
	return _IShareToken.contract.Transact(opts, "destroyShares", arg0, balance)
}

// DestroyShares is a paid mutator transaction binding the contract method 0xd333d7cf.
//
// Solidity: function destroyShares( address, balance uint256) returns(bool)
func (_IShareToken *IShareTokenSession) DestroyShares(arg0 common.Address, balance *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.DestroyShares(&_IShareToken.TransactOpts, arg0, balance)
}

// DestroyShares is a paid mutator transaction binding the contract method 0xd333d7cf.
//
// Solidity: function destroyShares( address, balance uint256) returns(bool)
func (_IShareToken *IShareTokenTransactorSession) DestroyShares(arg0 common.Address, balance *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.DestroyShares(&_IShareToken.TransactOpts, arg0, balance)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_market address, _outcome uint256) returns(bool)
func (_IShareToken *IShareTokenTransactor) Initialize(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int) (*types.Transaction, error) {
	return _IShareToken.contract.Transact(opts, "initialize", _market, _outcome)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_market address, _outcome uint256) returns(bool)
func (_IShareToken *IShareTokenSession) Initialize(_market common.Address, _outcome *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.Initialize(&_IShareToken.TransactOpts, _market, _outcome)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_market address, _outcome uint256) returns(bool)
func (_IShareToken *IShareTokenTransactorSession) Initialize(_market common.Address, _outcome *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.Initialize(&_IShareToken.TransactOpts, _market, _outcome)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IShareToken *IShareTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IShareToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IShareToken *IShareTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.Transfer(&_IShareToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IShareToken *IShareTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.Transfer(&_IShareToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IShareToken *IShareTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IShareToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IShareToken *IShareTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.TransferFrom(&_IShareToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IShareToken *IShareTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.TransferFrom(&_IShareToken.TransactOpts, _from, _to, _value)
}

// TrustedCancelOrderTransfer is a paid mutator transaction binding the contract method 0x7b30074d.
//
// Solidity: function trustedCancelOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IShareToken *IShareTokenTransactor) TrustedCancelOrderTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IShareToken.contract.Transact(opts, "trustedCancelOrderTransfer", _source, _destination, _attotokens)
}

// TrustedCancelOrderTransfer is a paid mutator transaction binding the contract method 0x7b30074d.
//
// Solidity: function trustedCancelOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IShareToken *IShareTokenSession) TrustedCancelOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.TrustedCancelOrderTransfer(&_IShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedCancelOrderTransfer is a paid mutator transaction binding the contract method 0x7b30074d.
//
// Solidity: function trustedCancelOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IShareToken *IShareTokenTransactorSession) TrustedCancelOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.TrustedCancelOrderTransfer(&_IShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedFillOrderTransfer is a paid mutator transaction binding the contract method 0x3d3c5c9f.
//
// Solidity: function trustedFillOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IShareToken *IShareTokenTransactor) TrustedFillOrderTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IShareToken.contract.Transact(opts, "trustedFillOrderTransfer", _source, _destination, _attotokens)
}

// TrustedFillOrderTransfer is a paid mutator transaction binding the contract method 0x3d3c5c9f.
//
// Solidity: function trustedFillOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IShareToken *IShareTokenSession) TrustedFillOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.TrustedFillOrderTransfer(&_IShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedFillOrderTransfer is a paid mutator transaction binding the contract method 0x3d3c5c9f.
//
// Solidity: function trustedFillOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IShareToken *IShareTokenTransactorSession) TrustedFillOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.TrustedFillOrderTransfer(&_IShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedOrderTransfer is a paid mutator transaction binding the contract method 0x764c92f2.
//
// Solidity: function trustedOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IShareToken *IShareTokenTransactor) TrustedOrderTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IShareToken.contract.Transact(opts, "trustedOrderTransfer", _source, _destination, _attotokens)
}

// TrustedOrderTransfer is a paid mutator transaction binding the contract method 0x764c92f2.
//
// Solidity: function trustedOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IShareToken *IShareTokenSession) TrustedOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.TrustedOrderTransfer(&_IShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedOrderTransfer is a paid mutator transaction binding the contract method 0x764c92f2.
//
// Solidity: function trustedOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_IShareToken *IShareTokenTransactorSession) TrustedOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IShareToken.Contract.TrustedOrderTransfer(&_IShareToken.TransactOpts, _source, _destination, _attotokens)
}

// IShareTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IShareToken contract.
type IShareTokenApprovalIterator struct {
	Event *IShareTokenApproval // Event containing the contract specifics and raw log

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
func (it *IShareTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IShareTokenApproval)
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
		it.Event = new(IShareTokenApproval)
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
func (it *IShareTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IShareTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IShareTokenApproval represents a Approval event raised by the IShareToken contract.
type IShareTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IShareToken *IShareTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IShareTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IShareToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IShareTokenApprovalIterator{contract: _IShareToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IShareToken *IShareTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IShareTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IShareToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IShareTokenApproval)
				if err := _IShareToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IShareTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IShareToken contract.
type IShareTokenTransferIterator struct {
	Event *IShareTokenTransfer // Event containing the contract specifics and raw log

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
func (it *IShareTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IShareTokenTransfer)
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
		it.Event = new(IShareTokenTransfer)
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
func (it *IShareTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IShareTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IShareTokenTransfer represents a Transfer event raised by the IShareToken contract.
type IShareTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IShareToken *IShareTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IShareTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IShareToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IShareTokenTransferIterator{contract: _IShareToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IShareToken *IShareTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IShareTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IShareToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IShareTokenTransfer)
				if err := _IShareToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ShareTokenABI is the input ABI used to generate the binding from.
const ShareTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_fxpValue\",\"type\":\"uint256\"}],\"name\":\"createShares\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedFillOrderTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETERNAL_APPROVAL_VALUE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedOrderTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedCancelOrderTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOutcome\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_fxpValue\",\"type\":\"uint256\"}],\"name\":\"destroyShares\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ShareTokenBin is the compiled bytecode used for deploying new contracts.
const ShareTokenBin = `0x60806040526002805460ff1916905560008054600160a060020a0319163317905561123f8061002f6000396000f30060806040526004361061012f5763ffffffff60e060020a60003504166306fdde038114610134578063095ea7b3146101be5780630f9e5bbd146101f657806318160ddd1461021a57806323b872dd146102415780633018205f1461026b578063313ce5671461029c5780633d3c5c9f146102c7578063634eaff1146102f1578063661884631461030657806370a082311461032a578063764c92f21461034b5780637b30074d146103755780637e7e4b471461039f57806392eefe9b146103b457806395d89b41146103d5578063a9059cbb146103ea578063bef72fa21461040e578063cd6dc68714610423578063d333d7cf14610447578063d73dd6231461046b578063db0a087c1461048f578063dd62ed3e146104a4578063ee89dab4146104cb578063f1be1679146104e0575b600080fd5b34801561014057600080fd5b506101496104f5565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018357818101518382015260200161016b565b50505050905090810190601f1680156101b05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ca57600080fd5b506101e2600160a060020a036004351660243561052c565b604080519115158252519081900360200190f35b34801561020257600080fd5b506101e2600160a060020a0360043516602435610543565b34801561022657600080fd5b5061022f6105ea565b60408051918252519081900360200190f35b34801561024d57600080fd5b506101e2600160a060020a03600435811690602435166044356105f0565b34801561027757600080fd5b50610280610669565b60408051600160a060020a039092168252519081900360200190f35b3480156102a857600080fd5b506102b1610678565b6040805160ff9092168252519081900360200190f35b3480156102d357600080fd5b506101e2600160a060020a036004358116906024351660443561067d565b3480156102fd57600080fd5b5061022f61076a565b34801561031257600080fd5b506101e2600160a060020a0360043516602435610770565b34801561033657600080fd5b5061022f600160a060020a03600435166107d3565b34801561035757600080fd5b506101e2600160a060020a03600435811690602435166044356107ee565b34801561038157600080fd5b506101e2600160a060020a0360043581169060243516604435610876565b3480156103ab57600080fd5b5061022f6108fe565b3480156103c057600080fd5b506101e2600160a060020a0360043516610904565b3480156103e157600080fd5b5061014961094e565b3480156103f657600080fd5b506101e2600160a060020a0360043516602435610985565b34801561041a57600080fd5b5061022f610999565b34801561042f57600080fd5b506101e2600160a060020a036004351660243561099f565b34801561045357600080fd5b506101e2600160a060020a03600435166024356109f2565b34801561047757600080fd5b506101e2600160a060020a0360043516602435610a99565b34801561049b57600080fd5b5061022f610ad5565b3480156104b057600080fd5b5061022f600160a060020a0360043581169060243516610af9565b3480156104d757600080fd5b506101e2610b24565b3480156104ec57600080fd5b50610280610b2d565b60408051808201909152600681527f5368617265730000000000000000000000000000000000000000000000000000602082015281565b6000610539338484610b3c565b5060019392505050565b60008054604080517f3f08882f0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b1580156105a957600080fd5b505af11580156105bd573d6000803e3d6000fd5b505050506040513d60208110156105d357600080fd5b505115156105e057600080fd5b6105398383610ba7565b60035490565b600160a060020a038316600090815260056020908152604080832033845290915281205460001981146106525761062d818463ffffffff610c4816565b600160a060020a03861660009081526005602090815260408083203384529091529020555b61065d858585610c5d565b50600195945050505050565b600054600160a060020a031690565b600081565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f46696c6c4f7264657200000000000000000000000000000000000000000000006004820181905291519192600160a060020a03169163f39ec1f79160248082019260209290919082900301818887803b15801561070557600080fd5b505af1158015610719573d6000803e3d6000fd5b505050506040513d602081101561072f57600080fd5b5051600160a060020a0316331461074557600080fd5b60025460ff16151561075657600080fd5b610761858585610c5d565b95945050505050565b60001981565b336000908152600560209081526040808320600160a060020a0386168452909152812054808311156107ae576107a833856000610b3c565b50610539565b6107c833856107c3848763ffffffff610c4816565b610b3c565b505060019392505050565b600160a060020a031660009081526004602052604090205490565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4372656174654f726465720000000000000000000000000000000000000000006004820181905291519192600160a060020a03169163f39ec1f79160248082019260209290919082900301818887803b15801561070557600080fd5b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f43616e63656c4f726465720000000000000000000000000000000000000000006004820181905291519192600160a060020a03169163f39ec1f79160248082019260209290919082900301818887803b15801561070557600080fd5b60075490565b60008054600160a060020a0316331461091c57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60408051808201909152600581527f5348415245000000000000000000000000000000000000000000000000000000602082015281565b6000610992338484610c5d565b9392505050565b60015481565b60025460009060ff16156109b257600080fd5b6109ba610d28565b50506006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039390931692909217909155600755600190565b60008054604080517f3f08882f0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b158015610a5857600080fd5b505af1158015610a6c573d6000803e3d6000fd5b505050506040513d6020811015610a8257600080fd5b50511515610a8f57600080fd5b6105398383610d4f565b336000818152600560209081526040808320600160a060020a038716845290915281205490916105399185906107c3908663ffffffff610df016565b7f5368617265546f6b656e0000000000000000000000000000000000000000000090565b600160a060020a03918216600090815260056020908152604080832093909416825291909152205490565b60025460ff1690565b600654600160a060020a031690565b600160a060020a03808416600081815260056020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a038216600090815260046020526040812054610bd0908363ffffffff610df016565b600160a060020a038416600090815260046020526040902055600354610bfc908363ffffffff610df016565b600355604080518381529051600160a060020a038516917f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885919081900360200190a26105398383610e02565b600082821115610c5757600080fd5b50900390565b600160a060020a038316600090815260046020526040812054610c86908363ffffffff610c4816565b600160a060020a038086166000908152600460205260408082209390935590851681522054610cbb908363ffffffff610df016565b600160a060020a0380851660008181526004602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3610d1d848484610f93565b506001949350505050565b60025460009060ff1615610d3b57600080fd5b506002805460ff1916600190811790915590565b600160a060020a038216600090815260046020526040812054610d78908363ffffffff610c4816565b600160a060020a038416600090815260046020526040902055600354610da4908363ffffffff610c4816565b600355604080518381529051600160a060020a038516917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2610539838361112c565b60008282018381101561099257600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610e5657600080fd5b505af1158015610e6a573d6000803e3d6000fd5b505050506040513d6020811015610e8057600080fd5b5051600654604080517f870c426d0000000000000000000000000000000000000000000000000000000081529051600160a060020a039384169363a1dfe54593169163870c426d9160048083019260209291908290030181600087803b158015610ee957600080fd5b505af1158015610efd573d6000803e3d6000fd5b505050506040513d6020811015610f1357600080fd5b50516040805160e060020a63ffffffff8516028152600160a060020a0392831660048201529187166024830152604482018690525160648083019260209291908290030181600087803b158015610f6957600080fd5b505af1158015610f7d573d6000803e3d6000fd5b505050506040513d60208110156107c857600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610fe757600080fd5b505af1158015610ffb573d6000803e3d6000fd5b505050506040513d602081101561101157600080fd5b5051600654604080517f870c426d0000000000000000000000000000000000000000000000000000000081529051600160a060020a03938416936386b9a1f493169163870c426d9160048083019260209291908290030181600087803b15801561107a57600080fd5b505af115801561108e573d6000803e3d6000fd5b505050506040513d60208110156110a457600080fd5b50516040805160e060020a63ffffffff8516028152600160a060020a03928316600482015282891660248201529187166044830152606482018690525160848083019260209291908290030181600087803b15801561110257600080fd5b505af1158015611116573d6000803e3d6000fd5b505050506040513d602081101561065d57600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561118057600080fd5b505af1158015611194573d6000803e3d6000fd5b505050506040513d60208110156111aa57600080fd5b5051600654604080517f870c426d0000000000000000000000000000000000000000000000000000000081529051600160a060020a039384169363a1b7887f93169163870c426d9160048083019260209291908290030181600087803b158015610ee957600080fd00a165627a7a723058208767a424426cb6b8082a0b7e52905afc167369cd865c99437bcbfbd3629251b90029`

// DeployShareToken deploys a new Ethereum contract, binding an instance of ShareToken to it.
func DeployShareToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ShareToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ShareTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ShareTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShareToken{ShareTokenCaller: ShareTokenCaller{contract: contract}, ShareTokenTransactor: ShareTokenTransactor{contract: contract}, ShareTokenFilterer: ShareTokenFilterer{contract: contract}}, nil
}

// ShareToken is an auto generated Go binding around an Ethereum contract.
type ShareToken struct {
	ShareTokenCaller     // Read-only binding to the contract
	ShareTokenTransactor // Write-only binding to the contract
	ShareTokenFilterer   // Log filterer for contract events
}

// ShareTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShareTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShareTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShareTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShareTokenSession struct {
	Contract     *ShareToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShareTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShareTokenCallerSession struct {
	Contract *ShareTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ShareTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShareTokenTransactorSession struct {
	Contract     *ShareTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ShareTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShareTokenRaw struct {
	Contract *ShareToken // Generic contract binding to access the raw methods on
}

// ShareTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShareTokenCallerRaw struct {
	Contract *ShareTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ShareTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShareTokenTransactorRaw struct {
	Contract *ShareTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShareToken creates a new instance of ShareToken, bound to a specific deployed contract.
func NewShareToken(address common.Address, backend bind.ContractBackend) (*ShareToken, error) {
	contract, err := bindShareToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShareToken{ShareTokenCaller: ShareTokenCaller{contract: contract}, ShareTokenTransactor: ShareTokenTransactor{contract: contract}, ShareTokenFilterer: ShareTokenFilterer{contract: contract}}, nil
}

// NewShareTokenCaller creates a new read-only instance of ShareToken, bound to a specific deployed contract.
func NewShareTokenCaller(address common.Address, caller bind.ContractCaller) (*ShareTokenCaller, error) {
	contract, err := bindShareToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShareTokenCaller{contract: contract}, nil
}

// NewShareTokenTransactor creates a new write-only instance of ShareToken, bound to a specific deployed contract.
func NewShareTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ShareTokenTransactor, error) {
	contract, err := bindShareToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShareTokenTransactor{contract: contract}, nil
}

// NewShareTokenFilterer creates a new log filterer instance of ShareToken, bound to a specific deployed contract.
func NewShareTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ShareTokenFilterer, error) {
	contract, err := bindShareToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShareTokenFilterer{contract: contract}, nil
}

// bindShareToken binds a generic wrapper to an already deployed contract.
func bindShareToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShareTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareToken *ShareTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShareToken.Contract.ShareTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareToken *ShareTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareToken.Contract.ShareTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareToken *ShareTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareToken.Contract.ShareTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareToken *ShareTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShareToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareToken *ShareTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareToken *ShareTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareToken.Contract.contract.Transact(opts, method, params...)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_ShareToken *ShareTokenCaller) ETERNALAPPROVALVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "ETERNAL_APPROVAL_VALUE")
	return *ret0, err
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_ShareToken *ShareTokenSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _ShareToken.Contract.ETERNALAPPROVALVALUE(&_ShareToken.CallOpts)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_ShareToken *ShareTokenCallerSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _ShareToken.Contract.ETERNALAPPROVALVALUE(&_ShareToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ShareToken *ShareTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ShareToken *ShareTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ShareToken.Contract.Allowance(&_ShareToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ShareToken *ShareTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ShareToken.Contract.Allowance(&_ShareToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ShareToken *ShareTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ShareToken *ShareTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ShareToken.Contract.BalanceOf(&_ShareToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ShareToken *ShareTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ShareToken.Contract.BalanceOf(&_ShareToken.CallOpts, _owner)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_ShareToken *ShareTokenCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_ShareToken *ShareTokenSession) ControllerLookupName() ([32]byte, error) {
	return _ShareToken.Contract.ControllerLookupName(&_ShareToken.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_ShareToken *ShareTokenCallerSession) ControllerLookupName() ([32]byte, error) {
	return _ShareToken.Contract.ControllerLookupName(&_ShareToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ShareToken *ShareTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ShareToken *ShareTokenSession) Decimals() (uint8, error) {
	return _ShareToken.Contract.Decimals(&_ShareToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ShareToken *ShareTokenCallerSession) Decimals() (uint8, error) {
	return _ShareToken.Contract.Decimals(&_ShareToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ShareToken *ShareTokenCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ShareToken *ShareTokenSession) GetController() (common.Address, error) {
	return _ShareToken.Contract.GetController(&_ShareToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ShareToken *ShareTokenCallerSession) GetController() (common.Address, error) {
	return _ShareToken.Contract.GetController(&_ShareToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_ShareToken *ShareTokenCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_ShareToken *ShareTokenSession) GetInitialized() (bool, error) {
	return _ShareToken.Contract.GetInitialized(&_ShareToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_ShareToken *ShareTokenCallerSession) GetInitialized() (bool, error) {
	return _ShareToken.Contract.GetInitialized(&_ShareToken.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_ShareToken *ShareTokenCaller) GetMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getMarket")
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_ShareToken *ShareTokenSession) GetMarket() (common.Address, error) {
	return _ShareToken.Contract.GetMarket(&_ShareToken.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_ShareToken *ShareTokenCallerSession) GetMarket() (common.Address, error) {
	return _ShareToken.Contract.GetMarket(&_ShareToken.CallOpts)
}

// GetOutcome is a free data retrieval call binding the contract method 0x7e7e4b47.
//
// Solidity: function getOutcome() constant returns(uint256)
func (_ShareToken *ShareTokenCaller) GetOutcome(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getOutcome")
	return *ret0, err
}

// GetOutcome is a free data retrieval call binding the contract method 0x7e7e4b47.
//
// Solidity: function getOutcome() constant returns(uint256)
func (_ShareToken *ShareTokenSession) GetOutcome() (*big.Int, error) {
	return _ShareToken.Contract.GetOutcome(&_ShareToken.CallOpts)
}

// GetOutcome is a free data retrieval call binding the contract method 0x7e7e4b47.
//
// Solidity: function getOutcome() constant returns(uint256)
func (_ShareToken *ShareTokenCallerSession) GetOutcome() (*big.Int, error) {
	return _ShareToken.Contract.GetOutcome(&_ShareToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ShareToken *ShareTokenCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ShareToken *ShareTokenSession) GetTypeName() ([32]byte, error) {
	return _ShareToken.Contract.GetTypeName(&_ShareToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ShareToken *ShareTokenCallerSession) GetTypeName() ([32]byte, error) {
	return _ShareToken.Contract.GetTypeName(&_ShareToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ShareToken *ShareTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ShareToken *ShareTokenSession) Name() (string, error) {
	return _ShareToken.Contract.Name(&_ShareToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ShareToken *ShareTokenCallerSession) Name() (string, error) {
	return _ShareToken.Contract.Name(&_ShareToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ShareToken *ShareTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ShareToken *ShareTokenSession) Symbol() (string, error) {
	return _ShareToken.Contract.Symbol(&_ShareToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ShareToken *ShareTokenCallerSession) Symbol() (string, error) {
	return _ShareToken.Contract.Symbol(&_ShareToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ShareToken *ShareTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ShareToken *ShareTokenSession) TotalSupply() (*big.Int, error) {
	return _ShareToken.Contract.TotalSupply(&_ShareToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ShareToken *ShareTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ShareToken.Contract.TotalSupply(&_ShareToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ShareToken *ShareTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.Approve(&_ShareToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.Approve(&_ShareToken.TransactOpts, _spender, _value)
}

// CreateShares is a paid mutator transaction binding the contract method 0x0f9e5bbd.
//
// Solidity: function createShares(_owner address, _fxpValue uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) CreateShares(opts *bind.TransactOpts, _owner common.Address, _fxpValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "createShares", _owner, _fxpValue)
}

// CreateShares is a paid mutator transaction binding the contract method 0x0f9e5bbd.
//
// Solidity: function createShares(_owner address, _fxpValue uint256) returns(bool)
func (_ShareToken *ShareTokenSession) CreateShares(_owner common.Address, _fxpValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.CreateShares(&_ShareToken.TransactOpts, _owner, _fxpValue)
}

// CreateShares is a paid mutator transaction binding the contract method 0x0f9e5bbd.
//
// Solidity: function createShares(_owner address, _fxpValue uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) CreateShares(_owner common.Address, _fxpValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.CreateShares(&_ShareToken.TransactOpts, _owner, _fxpValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_ShareToken *ShareTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.DecreaseApproval(&_ShareToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.DecreaseApproval(&_ShareToken.TransactOpts, _spender, _subtractedValue)
}

// DestroyShares is a paid mutator transaction binding the contract method 0xd333d7cf.
//
// Solidity: function destroyShares(_owner address, _fxpValue uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) DestroyShares(opts *bind.TransactOpts, _owner common.Address, _fxpValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "destroyShares", _owner, _fxpValue)
}

// DestroyShares is a paid mutator transaction binding the contract method 0xd333d7cf.
//
// Solidity: function destroyShares(_owner address, _fxpValue uint256) returns(bool)
func (_ShareToken *ShareTokenSession) DestroyShares(_owner common.Address, _fxpValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.DestroyShares(&_ShareToken.TransactOpts, _owner, _fxpValue)
}

// DestroyShares is a paid mutator transaction binding the contract method 0xd333d7cf.
//
// Solidity: function destroyShares(_owner address, _fxpValue uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) DestroyShares(_owner common.Address, _fxpValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.DestroyShares(&_ShareToken.TransactOpts, _owner, _fxpValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ShareToken *ShareTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.IncreaseApproval(&_ShareToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.IncreaseApproval(&_ShareToken.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_market address, _outcome uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) Initialize(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "initialize", _market, _outcome)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_market address, _outcome uint256) returns(bool)
func (_ShareToken *ShareTokenSession) Initialize(_market common.Address, _outcome *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.Initialize(&_ShareToken.TransactOpts, _market, _outcome)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_market address, _outcome uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) Initialize(_market common.Address, _outcome *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.Initialize(&_ShareToken.TransactOpts, _market, _outcome)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ShareToken *ShareTokenTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ShareToken *ShareTokenSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ShareToken.Contract.SetController(&_ShareToken.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ShareToken.Contract.SetController(&_ShareToken.TransactOpts, _controller)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ShareToken *ShareTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.Transfer(&_ShareToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.Transfer(&_ShareToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ShareToken *ShareTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.TransferFrom(&_ShareToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.TransferFrom(&_ShareToken.TransactOpts, _from, _to, _value)
}

// TrustedCancelOrderTransfer is a paid mutator transaction binding the contract method 0x7b30074d.
//
// Solidity: function trustedCancelOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) TrustedCancelOrderTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "trustedCancelOrderTransfer", _source, _destination, _attotokens)
}

// TrustedCancelOrderTransfer is a paid mutator transaction binding the contract method 0x7b30074d.
//
// Solidity: function trustedCancelOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ShareToken *ShareTokenSession) TrustedCancelOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.TrustedCancelOrderTransfer(&_ShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedCancelOrderTransfer is a paid mutator transaction binding the contract method 0x7b30074d.
//
// Solidity: function trustedCancelOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) TrustedCancelOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.TrustedCancelOrderTransfer(&_ShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedFillOrderTransfer is a paid mutator transaction binding the contract method 0x3d3c5c9f.
//
// Solidity: function trustedFillOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) TrustedFillOrderTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "trustedFillOrderTransfer", _source, _destination, _attotokens)
}

// TrustedFillOrderTransfer is a paid mutator transaction binding the contract method 0x3d3c5c9f.
//
// Solidity: function trustedFillOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ShareToken *ShareTokenSession) TrustedFillOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.TrustedFillOrderTransfer(&_ShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedFillOrderTransfer is a paid mutator transaction binding the contract method 0x3d3c5c9f.
//
// Solidity: function trustedFillOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) TrustedFillOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.TrustedFillOrderTransfer(&_ShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedOrderTransfer is a paid mutator transaction binding the contract method 0x764c92f2.
//
// Solidity: function trustedOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ShareToken *ShareTokenTransactor) TrustedOrderTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "trustedOrderTransfer", _source, _destination, _attotokens)
}

// TrustedOrderTransfer is a paid mutator transaction binding the contract method 0x764c92f2.
//
// Solidity: function trustedOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ShareToken *ShareTokenSession) TrustedOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.TrustedOrderTransfer(&_ShareToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedOrderTransfer is a paid mutator transaction binding the contract method 0x764c92f2.
//
// Solidity: function trustedOrderTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) TrustedOrderTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.TrustedOrderTransfer(&_ShareToken.TransactOpts, _source, _destination, _attotokens)
}

// ShareTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ShareToken contract.
type ShareTokenApprovalIterator struct {
	Event *ShareTokenApproval // Event containing the contract specifics and raw log

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
func (it *ShareTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareTokenApproval)
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
		it.Event = new(ShareTokenApproval)
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
func (it *ShareTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareTokenApproval represents a Approval event raised by the ShareToken contract.
type ShareTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ShareToken *ShareTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ShareTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShareToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ShareTokenApprovalIterator{contract: _ShareToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ShareToken *ShareTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ShareTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShareToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareTokenApproval)
				if err := _ShareToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ShareTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ShareToken contract.
type ShareTokenBurnIterator struct {
	Event *ShareTokenBurn // Event containing the contract specifics and raw log

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
func (it *ShareTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareTokenBurn)
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
		it.Event = new(ShareTokenBurn)
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
func (it *ShareTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareTokenBurn represents a Burn event raised by the ShareToken contract.
type ShareTokenBurn struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_ShareToken *ShareTokenFilterer) FilterBurn(opts *bind.FilterOpts, target []common.Address) (*ShareTokenBurnIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ShareToken.contract.FilterLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return &ShareTokenBurnIterator{contract: _ShareToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_ShareToken *ShareTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ShareTokenBurn, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ShareToken.contract.WatchLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareTokenBurn)
				if err := _ShareToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ShareTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ShareToken contract.
type ShareTokenMintIterator struct {
	Event *ShareTokenMint // Event containing the contract specifics and raw log

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
func (it *ShareTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareTokenMint)
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
		it.Event = new(ShareTokenMint)
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
func (it *ShareTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareTokenMint represents a Mint event raised by the ShareToken contract.
type ShareTokenMint struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_ShareToken *ShareTokenFilterer) FilterMint(opts *bind.FilterOpts, target []common.Address) (*ShareTokenMintIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ShareToken.contract.FilterLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return &ShareTokenMintIterator{contract: _ShareToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_ShareToken *ShareTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ShareTokenMint, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ShareToken.contract.WatchLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareTokenMint)
				if err := _ShareToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ShareTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ShareToken contract.
type ShareTokenTransferIterator struct {
	Event *ShareTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ShareTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareTokenTransfer)
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
		it.Event = new(ShareTokenTransfer)
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
func (it *ShareTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareTokenTransfer represents a Transfer event raised by the ShareToken contract.
type ShareTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ShareToken *ShareTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShareTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShareToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShareTokenTransferIterator{contract: _ShareToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ShareToken *ShareTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ShareTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShareToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareTokenTransfer)
				if err := _ShareToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ShareTokenFactoryABI is the input ABI used to generate the binding from.
const ShareTokenFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"createShareToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ShareTokenFactoryBin is the compiled bytecode used for deploying new contracts.
const ShareTokenFactoryBin = `0x608060405234801561001057600080fd5b50610535806100206000396000f3006080604052600436106100405763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416635f97709a8114610045575b600080fd5b34801561005157600080fd5b5061007c73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356100a5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000806000856100b36101f1565b73ffffffffffffffffffffffffffffffffffffffff90911681527f5368617265546f6b656e000000000000000000000000000000000000000000006020820152604080519182900301906000f080158015610112573d6000803e3d6000fd5b5091508190508073ffffffffffffffffffffffffffffffffffffffff1663cd6dc68786866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156101bb57600080fd5b505af11580156101cf573d6000803e3d6000fd5b505050506040513d60208110156101e557600080fd5b50909695505050505050565b60405161030880610202833901905600608060405234801561001057600080fd5b506040516040806103088339810160405280516020909101516000805433600160a060020a03199182161716600160a060020a039093169290921782556001556102a890819061006090396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461014f57806392eefe9b1461018d578063bef72fa2146101cf575b60015460009015156100675761014c565b60008054600154604080517ff39ec1f700000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff9092169263f39ec1f7926024808401936020939083900390910190829087803b1580156100df57600080fd5b505af11580156100f3573d6000803e3d6000fd5b505050506040513d602081101561010957600080fd5b50516040805136601f8101601f191680830190935292935091600083376000803684865af4808015610147576040513d81016040523d6000823e3d81f35b600080fd5b50005b34801561015b57600080fd5b506101646101f6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019957600080fd5b506101bb73ffffffffffffffffffffffffffffffffffffffff60043516610212565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610276565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461023757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a7230582017c3a6b57ac8bbc7fd5320cba2543fe371eb8726332cd58700252790e09cf2140029a165627a7a723058202a0cd2dd7f11fec6dfbe359fc932a4c5675bda3d30fae0ff6f500b05e5fadf210029`

// DeployShareTokenFactory deploys a new Ethereum contract, binding an instance of ShareTokenFactory to it.
func DeployShareTokenFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ShareTokenFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ShareTokenFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ShareTokenFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShareTokenFactory{ShareTokenFactoryCaller: ShareTokenFactoryCaller{contract: contract}, ShareTokenFactoryTransactor: ShareTokenFactoryTransactor{contract: contract}, ShareTokenFactoryFilterer: ShareTokenFactoryFilterer{contract: contract}}, nil
}

// ShareTokenFactory is an auto generated Go binding around an Ethereum contract.
type ShareTokenFactory struct {
	ShareTokenFactoryCaller     // Read-only binding to the contract
	ShareTokenFactoryTransactor // Write-only binding to the contract
	ShareTokenFactoryFilterer   // Log filterer for contract events
}

// ShareTokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShareTokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareTokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShareTokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareTokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShareTokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareTokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShareTokenFactorySession struct {
	Contract     *ShareTokenFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ShareTokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShareTokenFactoryCallerSession struct {
	Contract *ShareTokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ShareTokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShareTokenFactoryTransactorSession struct {
	Contract     *ShareTokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ShareTokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShareTokenFactoryRaw struct {
	Contract *ShareTokenFactory // Generic contract binding to access the raw methods on
}

// ShareTokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShareTokenFactoryCallerRaw struct {
	Contract *ShareTokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ShareTokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShareTokenFactoryTransactorRaw struct {
	Contract *ShareTokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShareTokenFactory creates a new instance of ShareTokenFactory, bound to a specific deployed contract.
func NewShareTokenFactory(address common.Address, backend bind.ContractBackend) (*ShareTokenFactory, error) {
	contract, err := bindShareTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShareTokenFactory{ShareTokenFactoryCaller: ShareTokenFactoryCaller{contract: contract}, ShareTokenFactoryTransactor: ShareTokenFactoryTransactor{contract: contract}, ShareTokenFactoryFilterer: ShareTokenFactoryFilterer{contract: contract}}, nil
}

// NewShareTokenFactoryCaller creates a new read-only instance of ShareTokenFactory, bound to a specific deployed contract.
func NewShareTokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*ShareTokenFactoryCaller, error) {
	contract, err := bindShareTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShareTokenFactoryCaller{contract: contract}, nil
}

// NewShareTokenFactoryTransactor creates a new write-only instance of ShareTokenFactory, bound to a specific deployed contract.
func NewShareTokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ShareTokenFactoryTransactor, error) {
	contract, err := bindShareTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShareTokenFactoryTransactor{contract: contract}, nil
}

// NewShareTokenFactoryFilterer creates a new log filterer instance of ShareTokenFactory, bound to a specific deployed contract.
func NewShareTokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ShareTokenFactoryFilterer, error) {
	contract, err := bindShareTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShareTokenFactoryFilterer{contract: contract}, nil
}

// bindShareTokenFactory binds a generic wrapper to an already deployed contract.
func bindShareTokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShareTokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareTokenFactory *ShareTokenFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShareTokenFactory.Contract.ShareTokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareTokenFactory *ShareTokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareTokenFactory.Contract.ShareTokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareTokenFactory *ShareTokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareTokenFactory.Contract.ShareTokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareTokenFactory *ShareTokenFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShareTokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareTokenFactory *ShareTokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareTokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareTokenFactory *ShareTokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareTokenFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateShareToken is a paid mutator transaction binding the contract method 0x5f97709a.
//
// Solidity: function createShareToken(_controller address, _market address, _outcome uint256) returns(address)
func (_ShareTokenFactory *ShareTokenFactoryTransactor) CreateShareToken(opts *bind.TransactOpts, _controller common.Address, _market common.Address, _outcome *big.Int) (*types.Transaction, error) {
	return _ShareTokenFactory.contract.Transact(opts, "createShareToken", _controller, _market, _outcome)
}

// CreateShareToken is a paid mutator transaction binding the contract method 0x5f97709a.
//
// Solidity: function createShareToken(_controller address, _market address, _outcome uint256) returns(address)
func (_ShareTokenFactory *ShareTokenFactorySession) CreateShareToken(_controller common.Address, _market common.Address, _outcome *big.Int) (*types.Transaction, error) {
	return _ShareTokenFactory.Contract.CreateShareToken(&_ShareTokenFactory.TransactOpts, _controller, _market, _outcome)
}

// CreateShareToken is a paid mutator transaction binding the contract method 0x5f97709a.
//
// Solidity: function createShareToken(_controller address, _market address, _outcome uint256) returns(address)
func (_ShareTokenFactory *ShareTokenFactoryTransactorSession) CreateShareToken(_controller common.Address, _market common.Address, _outcome *big.Int) (*types.Transaction, error) {
	return _ShareTokenFactory.Contract.CreateShareToken(&_ShareTokenFactory.TransactOpts, _controller, _market, _outcome)
}
