package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SafeMathInt256ABI is the input ABI used to generate the binding from.
const SafeMathInt256ABI = "[]"

// SafeMathInt256Bin is the compiled bytecode used for deploying new contracts.
const SafeMathInt256Bin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820878705e15568687f69684760ba6ab6529b7a76a23ff611484211c1aed51a2e330029`

// DeploySafeMathInt256 deploys a new Ethereum contract, binding an instance of SafeMathInt256 to it.
func DeploySafeMathInt256(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathInt256, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathInt256ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathInt256Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMathInt256{SafeMathInt256Caller: SafeMathInt256Caller{contract: contract}, SafeMathInt256Transactor: SafeMathInt256Transactor{contract: contract}, SafeMathInt256Filterer: SafeMathInt256Filterer{contract: contract}}, nil
}

// SafeMathInt256 is an auto generated Go binding around an Ethereum contract.
type SafeMathInt256 struct {
	SafeMathInt256Caller     // Read-only binding to the contract
	SafeMathInt256Transactor // Write-only binding to the contract
	SafeMathInt256Filterer   // Log filterer for contract events
}

// SafeMathInt256Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathInt256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathInt256Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathInt256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathInt256Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathInt256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathInt256Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathInt256Session struct {
	Contract     *SafeMathInt256   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathInt256CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathInt256CallerSession struct {
	Contract *SafeMathInt256Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SafeMathInt256TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathInt256TransactorSession struct {
	Contract     *SafeMathInt256Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SafeMathInt256Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathInt256Raw struct {
	Contract *SafeMathInt256 // Generic contract binding to access the raw methods on
}

// SafeMathInt256CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathInt256CallerRaw struct {
	Contract *SafeMathInt256Caller // Generic read-only contract binding to access the raw methods on
}

// SafeMathInt256TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathInt256TransactorRaw struct {
	Contract *SafeMathInt256Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMathInt256 creates a new instance of SafeMathInt256, bound to a specific deployed contract.
func NewSafeMathInt256(address common.Address, backend bind.ContractBackend) (*SafeMathInt256, error) {
	contract, err := bindSafeMathInt256(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMathInt256{SafeMathInt256Caller: SafeMathInt256Caller{contract: contract}, SafeMathInt256Transactor: SafeMathInt256Transactor{contract: contract}, SafeMathInt256Filterer: SafeMathInt256Filterer{contract: contract}}, nil
}

// NewSafeMathInt256Caller creates a new read-only instance of SafeMathInt256, bound to a specific deployed contract.
func NewSafeMathInt256Caller(address common.Address, caller bind.ContractCaller) (*SafeMathInt256Caller, error) {
	contract, err := bindSafeMathInt256(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathInt256Caller{contract: contract}, nil
}

// NewSafeMathInt256Transactor creates a new write-only instance of SafeMathInt256, bound to a specific deployed contract.
func NewSafeMathInt256Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathInt256Transactor, error) {
	contract, err := bindSafeMathInt256(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathInt256Transactor{contract: contract}, nil
}

// NewSafeMathInt256Filterer creates a new log filterer instance of SafeMathInt256, bound to a specific deployed contract.
func NewSafeMathInt256Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathInt256Filterer, error) {
	contract, err := bindSafeMathInt256(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathInt256Filterer{contract: contract}, nil
}

// bindSafeMathInt256 binds a generic wrapper to an already deployed contract.
func bindSafeMathInt256(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathInt256ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathInt256 *SafeMathInt256Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMathInt256.Contract.SafeMathInt256Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathInt256 *SafeMathInt256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathInt256.Contract.SafeMathInt256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathInt256 *SafeMathInt256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathInt256.Contract.SafeMathInt256Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathInt256 *SafeMathInt256CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMathInt256.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathInt256 *SafeMathInt256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathInt256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathInt256 *SafeMathInt256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathInt256.Contract.contract.Transact(opts, method, params...)
}

// SafeMathUint256ABI is the input ABI used to generate the binding from.
const SafeMathUint256ABI = "[]"

// SafeMathUint256Bin is the compiled bytecode used for deploying new contracts.
const SafeMathUint256Bin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058203cbd8ef335ce044264783f07c720584ee9a3901642f83c5ccd7f0c1abfb164180029`

// DeploySafeMathUint256 deploys a new Ethereum contract, binding an instance of SafeMathUint256 to it.
func DeploySafeMathUint256(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathUint256, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathUint256ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathUint256Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMathUint256{SafeMathUint256Caller: SafeMathUint256Caller{contract: contract}, SafeMathUint256Transactor: SafeMathUint256Transactor{contract: contract}, SafeMathUint256Filterer: SafeMathUint256Filterer{contract: contract}}, nil
}

// SafeMathUint256 is an auto generated Go binding around an Ethereum contract.
type SafeMathUint256 struct {
	SafeMathUint256Caller     // Read-only binding to the contract
	SafeMathUint256Transactor // Write-only binding to the contract
	SafeMathUint256Filterer   // Log filterer for contract events
}

// SafeMathUint256Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathUint256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathUint256Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathUint256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathUint256Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathUint256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathUint256Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathUint256Session struct {
	Contract     *SafeMathUint256  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathUint256CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathUint256CallerSession struct {
	Contract *SafeMathUint256Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SafeMathUint256TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathUint256TransactorSession struct {
	Contract     *SafeMathUint256Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SafeMathUint256Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathUint256Raw struct {
	Contract *SafeMathUint256 // Generic contract binding to access the raw methods on
}

// SafeMathUint256CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathUint256CallerRaw struct {
	Contract *SafeMathUint256Caller // Generic read-only contract binding to access the raw methods on
}

// SafeMathUint256TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathUint256TransactorRaw struct {
	Contract *SafeMathUint256Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMathUint256 creates a new instance of SafeMathUint256, bound to a specific deployed contract.
func NewSafeMathUint256(address common.Address, backend bind.ContractBackend) (*SafeMathUint256, error) {
	contract, err := bindSafeMathUint256(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMathUint256{SafeMathUint256Caller: SafeMathUint256Caller{contract: contract}, SafeMathUint256Transactor: SafeMathUint256Transactor{contract: contract}, SafeMathUint256Filterer: SafeMathUint256Filterer{contract: contract}}, nil
}

// NewSafeMathUint256Caller creates a new read-only instance of SafeMathUint256, bound to a specific deployed contract.
func NewSafeMathUint256Caller(address common.Address, caller bind.ContractCaller) (*SafeMathUint256Caller, error) {
	contract, err := bindSafeMathUint256(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathUint256Caller{contract: contract}, nil
}

// NewSafeMathUint256Transactor creates a new write-only instance of SafeMathUint256, bound to a specific deployed contract.
func NewSafeMathUint256Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathUint256Transactor, error) {
	contract, err := bindSafeMathUint256(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathUint256Transactor{contract: contract}, nil
}

// NewSafeMathUint256Filterer creates a new log filterer instance of SafeMathUint256, bound to a specific deployed contract.
func NewSafeMathUint256Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathUint256Filterer, error) {
	contract, err := bindSafeMathUint256(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathUint256Filterer{contract: contract}, nil
}

// bindSafeMathUint256 binds a generic wrapper to an already deployed contract.
func bindSafeMathUint256(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathUint256ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathUint256 *SafeMathUint256Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMathUint256.Contract.SafeMathUint256Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathUint256 *SafeMathUint256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathUint256.Contract.SafeMathUint256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathUint256 *SafeMathUint256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathUint256.Contract.SafeMathUint256Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathUint256 *SafeMathUint256CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMathUint256.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathUint256 *SafeMathUint256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathUint256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathUint256 *SafeMathUint256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathUint256.Contract.contract.Transact(opts, method, params...)
}
