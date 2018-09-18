package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CashAutoConverterABI is the input ABI used to generate the binding from.
const CashAutoConverterABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CashAutoConverterBin is the compiled bytecode used for deploying new contracts.
const CashAutoConverterBin = `0x608060405260008054600160a060020a0319163317905561017c806100256000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461005057806392eefe9b1461008e575b600080fd5b34801561005c57600080fd5b506100656100d0565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b506100bc73ffffffffffffffffffffffffffffffffffffffff600435166100ec565b604080519115158252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461011157600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff1990911617905560019190505600a165627a7a7230582081017bf7e2ca953107e6d62bf89c801f3898568c06d87dafe6894058e61793590029`

// DeployCashAutoConverter deploys a new Ethereum contract, binding an instance of CashAutoConverter to it.
func DeployCashAutoConverter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CashAutoConverter, error) {
	parsed, err := abi.JSON(strings.NewReader(CashAutoConverterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CashAutoConverterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CashAutoConverter{CashAutoConverterCaller: CashAutoConverterCaller{contract: contract}, CashAutoConverterTransactor: CashAutoConverterTransactor{contract: contract}, CashAutoConverterFilterer: CashAutoConverterFilterer{contract: contract}}, nil
}

// CashAutoConverter is an auto generated Go binding around an Ethereum contract.
type CashAutoConverter struct {
	CashAutoConverterCaller     // Read-only binding to the contract
	CashAutoConverterTransactor // Write-only binding to the contract
	CashAutoConverterFilterer   // Log filterer for contract events
}

// CashAutoConverterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CashAutoConverterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashAutoConverterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CashAutoConverterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashAutoConverterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CashAutoConverterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashAutoConverterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CashAutoConverterSession struct {
	Contract     *CashAutoConverter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CashAutoConverterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CashAutoConverterCallerSession struct {
	Contract *CashAutoConverterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CashAutoConverterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CashAutoConverterTransactorSession struct {
	Contract     *CashAutoConverterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CashAutoConverterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CashAutoConverterRaw struct {
	Contract *CashAutoConverter // Generic contract binding to access the raw methods on
}

// CashAutoConverterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CashAutoConverterCallerRaw struct {
	Contract *CashAutoConverterCaller // Generic read-only contract binding to access the raw methods on
}

// CashAutoConverterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CashAutoConverterTransactorRaw struct {
	Contract *CashAutoConverterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCashAutoConverter creates a new instance of CashAutoConverter, bound to a specific deployed contract.
func NewCashAutoConverter(address common.Address, backend bind.ContractBackend) (*CashAutoConverter, error) {
	contract, err := bindCashAutoConverter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CashAutoConverter{CashAutoConverterCaller: CashAutoConverterCaller{contract: contract}, CashAutoConverterTransactor: CashAutoConverterTransactor{contract: contract}, CashAutoConverterFilterer: CashAutoConverterFilterer{contract: contract}}, nil
}

// NewCashAutoConverterCaller creates a new read-only instance of CashAutoConverter, bound to a specific deployed contract.
func NewCashAutoConverterCaller(address common.Address, caller bind.ContractCaller) (*CashAutoConverterCaller, error) {
	contract, err := bindCashAutoConverter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CashAutoConverterCaller{contract: contract}, nil
}

// NewCashAutoConverterTransactor creates a new write-only instance of CashAutoConverter, bound to a specific deployed contract.
func NewCashAutoConverterTransactor(address common.Address, transactor bind.ContractTransactor) (*CashAutoConverterTransactor, error) {
	contract, err := bindCashAutoConverter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CashAutoConverterTransactor{contract: contract}, nil
}

// NewCashAutoConverterFilterer creates a new log filterer instance of CashAutoConverter, bound to a specific deployed contract.
func NewCashAutoConverterFilterer(address common.Address, filterer bind.ContractFilterer) (*CashAutoConverterFilterer, error) {
	contract, err := bindCashAutoConverter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CashAutoConverterFilterer{contract: contract}, nil
}

// bindCashAutoConverter binds a generic wrapper to an already deployed contract.
func bindCashAutoConverter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CashAutoConverterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CashAutoConverter *CashAutoConverterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CashAutoConverter.Contract.CashAutoConverterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CashAutoConverter *CashAutoConverterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CashAutoConverter.Contract.CashAutoConverterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CashAutoConverter *CashAutoConverterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CashAutoConverter.Contract.CashAutoConverterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CashAutoConverter *CashAutoConverterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CashAutoConverter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CashAutoConverter *CashAutoConverterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CashAutoConverter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CashAutoConverter *CashAutoConverterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CashAutoConverter.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_CashAutoConverter *CashAutoConverterCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CashAutoConverter.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_CashAutoConverter *CashAutoConverterSession) GetController() (common.Address, error) {
	return _CashAutoConverter.Contract.GetController(&_CashAutoConverter.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_CashAutoConverter *CashAutoConverterCallerSession) GetController() (common.Address, error) {
	return _CashAutoConverter.Contract.GetController(&_CashAutoConverter.CallOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_CashAutoConverter *CashAutoConverterTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _CashAutoConverter.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_CashAutoConverter *CashAutoConverterSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _CashAutoConverter.Contract.SetController(&_CashAutoConverter.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_CashAutoConverter *CashAutoConverterTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _CashAutoConverter.Contract.SetController(&_CashAutoConverter.TransactOpts, _controller)
}
