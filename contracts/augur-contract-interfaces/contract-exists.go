package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ContractExistsABI is the input ABI used to generate the binding from.
const ContractExistsABI = "[]"

// ContractExistsBin is the compiled bytecode used for deploying new contracts.
const ContractExistsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e618fa31d9d6b9407d08ca50e12847b025fdeddcc3508155857b1620687d95b30029`

// DeployContractExists deploys a new Ethereum contract, binding an instance of ContractExists to it.
func DeployContractExists(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractExists, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractExistsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractExistsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractExists{ContractExistsCaller: ContractExistsCaller{contract: contract}, ContractExistsTransactor: ContractExistsTransactor{contract: contract}, ContractExistsFilterer: ContractExistsFilterer{contract: contract}}, nil
}

// ContractExists is an auto generated Go binding around an Ethereum contract.
type ContractExists struct {
	ContractExistsCaller     // Read-only binding to the contract
	ContractExistsTransactor // Write-only binding to the contract
	ContractExistsFilterer   // Log filterer for contract events
}

// ContractExistsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractExistsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractExistsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractExistsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractExistsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractExistsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractExistsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractExistsSession struct {
	Contract     *ContractExists   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractExistsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractExistsCallerSession struct {
	Contract *ContractExistsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ContractExistsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractExistsTransactorSession struct {
	Contract     *ContractExistsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractExistsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractExistsRaw struct {
	Contract *ContractExists // Generic contract binding to access the raw methods on
}

// ContractExistsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractExistsCallerRaw struct {
	Contract *ContractExistsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractExistsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractExistsTransactorRaw struct {
	Contract *ContractExistsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractExists creates a new instance of ContractExists, bound to a specific deployed contract.
func NewContractExists(address common.Address, backend bind.ContractBackend) (*ContractExists, error) {
	contract, err := bindContractExists(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractExists{ContractExistsCaller: ContractExistsCaller{contract: contract}, ContractExistsTransactor: ContractExistsTransactor{contract: contract}, ContractExistsFilterer: ContractExistsFilterer{contract: contract}}, nil
}

// NewContractExistsCaller creates a new read-only instance of ContractExists, bound to a specific deployed contract.
func NewContractExistsCaller(address common.Address, caller bind.ContractCaller) (*ContractExistsCaller, error) {
	contract, err := bindContractExists(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractExistsCaller{contract: contract}, nil
}

// NewContractExistsTransactor creates a new write-only instance of ContractExists, bound to a specific deployed contract.
func NewContractExistsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractExistsTransactor, error) {
	contract, err := bindContractExists(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractExistsTransactor{contract: contract}, nil
}

// NewContractExistsFilterer creates a new log filterer instance of ContractExists, bound to a specific deployed contract.
func NewContractExistsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractExistsFilterer, error) {
	contract, err := bindContractExists(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractExistsFilterer{contract: contract}, nil
}

// bindContractExists binds a generic wrapper to an already deployed contract.
func bindContractExists(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractExistsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractExists *ContractExistsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractExists.Contract.ContractExistsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractExists *ContractExistsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractExists.Contract.ContractExistsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractExists *ContractExistsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractExists.Contract.ContractExistsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractExists *ContractExistsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractExists.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractExists *ContractExistsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractExists.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractExists *ContractExistsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractExists.Contract.contract.Transact(opts, method, params...)
}
