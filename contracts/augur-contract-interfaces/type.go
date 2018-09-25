package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ITypedABI is the input ABI used to generate the binding from.
const ITypedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ITypedBin is the compiled bytecode used for deploying new contracts.
const ITypedBin = `0x`

// DeployITyped deploys a new Ethereum contract, binding an instance of ITyped to it.
func DeployITyped(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ITyped, error) {
	parsed, err := abi.JSON(strings.NewReader(ITypedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ITypedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ITyped{ITypedCaller: ITypedCaller{contract: contract}, ITypedTransactor: ITypedTransactor{contract: contract}, ITypedFilterer: ITypedFilterer{contract: contract}}, nil
}

// ITyped is an auto generated Go binding around an Ethereum contract.
type ITyped struct {
	ITypedCaller     // Read-only binding to the contract
	ITypedTransactor // Write-only binding to the contract
	ITypedFilterer   // Log filterer for contract events
}

// ITypedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITypedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITypedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITypedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITypedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITypedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITypedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITypedSession struct {
	Contract     *ITyped           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITypedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITypedCallerSession struct {
	Contract *ITypedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITypedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITypedTransactorSession struct {
	Contract     *ITypedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITypedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITypedRaw struct {
	Contract *ITyped // Generic contract binding to access the raw methods on
}

// ITypedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITypedCallerRaw struct {
	Contract *ITypedCaller // Generic read-only contract binding to access the raw methods on
}

// ITypedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITypedTransactorRaw struct {
	Contract *ITypedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITyped creates a new instance of ITyped, bound to a specific deployed contract.
func NewITyped(address common.Address, backend bind.ContractBackend) (*ITyped, error) {
	contract, err := bindITyped(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITyped{ITypedCaller: ITypedCaller{contract: contract}, ITypedTransactor: ITypedTransactor{contract: contract}, ITypedFilterer: ITypedFilterer{contract: contract}}, nil
}

// NewITypedCaller creates a new read-only instance of ITyped, bound to a specific deployed contract.
func NewITypedCaller(address common.Address, caller bind.ContractCaller) (*ITypedCaller, error) {
	contract, err := bindITyped(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITypedCaller{contract: contract}, nil
}

// NewITypedTransactor creates a new write-only instance of ITyped, bound to a specific deployed contract.
func NewITypedTransactor(address common.Address, transactor bind.ContractTransactor) (*ITypedTransactor, error) {
	contract, err := bindITyped(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITypedTransactor{contract: contract}, nil
}

// NewITypedFilterer creates a new log filterer instance of ITyped, bound to a specific deployed contract.
func NewITypedFilterer(address common.Address, filterer bind.ContractFilterer) (*ITypedFilterer, error) {
	contract, err := bindITyped(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITypedFilterer{contract: contract}, nil
}

// bindITyped binds a generic wrapper to an already deployed contract.
func bindITyped(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITypedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITyped *ITypedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITyped.Contract.ITypedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITyped *ITypedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITyped.Contract.ITypedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITyped *ITypedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITyped.Contract.ITypedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITyped *ITypedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITyped.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITyped *ITypedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITyped.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITyped *ITypedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITyped.Contract.contract.Transact(opts, method, params...)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ITyped *ITypedCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ITyped.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ITyped *ITypedSession) GetTypeName() ([32]byte, error) {
	return _ITyped.Contract.GetTypeName(&_ITyped.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ITyped *ITypedCallerSession) GetTypeName() ([32]byte, error) {
	return _ITyped.Contract.GetTypeName(&_ITyped.CallOpts)
}
