package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ControlledABI is the input ABI used to generate the binding from.
const ControlledABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ControlledBin is the compiled bytecode used for deploying new contracts.
const ControlledBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561017c806100326000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461005057806392eefe9b1461008e575b600080fd5b34801561005c57600080fd5b506100656100d0565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b506100bc73ffffffffffffffffffffffffffffffffffffffff600435166100ec565b604080519115158252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461011157600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff1990911617905560019190505600a165627a7a72305820e47dad7468b569dda417cf495cff6ae9c66bace21e8923c4f9e381227faaf93b0029`

// DeployControlled deploys a new Ethereum contract, binding an instance of Controlled to it.
func DeployControlled(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Controlled, error) {
	parsed, err := abi.JSON(strings.NewReader(ControlledABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ControlledBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Controlled{ControlledCaller: ControlledCaller{contract: contract}, ControlledTransactor: ControlledTransactor{contract: contract}, ControlledFilterer: ControlledFilterer{contract: contract}}, nil
}

// Controlled is an auto generated Go binding around an Ethereum contract.
type Controlled struct {
	ControlledCaller     // Read-only binding to the contract
	ControlledTransactor // Write-only binding to the contract
	ControlledFilterer   // Log filterer for contract events
}

// ControlledCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControlledCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlledTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControlledTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlledFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControlledFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlledSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControlledSession struct {
	Contract     *Controlled       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControlledCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControlledCallerSession struct {
	Contract *ControlledCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControlledTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControlledTransactorSession struct {
	Contract     *ControlledTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControlledRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControlledRaw struct {
	Contract *Controlled // Generic contract binding to access the raw methods on
}

// ControlledCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControlledCallerRaw struct {
	Contract *ControlledCaller // Generic read-only contract binding to access the raw methods on
}

// ControlledTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControlledTransactorRaw struct {
	Contract *ControlledTransactor // Generic write-only contract binding to access the raw methods on
}

// NewControlled creates a new instance of Controlled, bound to a specific deployed contract.
func NewControlled(address common.Address, backend bind.ContractBackend) (*Controlled, error) {
	contract, err := bindControlled(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controlled{ControlledCaller: ControlledCaller{contract: contract}, ControlledTransactor: ControlledTransactor{contract: contract}, ControlledFilterer: ControlledFilterer{contract: contract}}, nil
}

// NewControlledCaller creates a new read-only instance of Controlled, bound to a specific deployed contract.
func NewControlledCaller(address common.Address, caller bind.ContractCaller) (*ControlledCaller, error) {
	contract, err := bindControlled(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControlledCaller{contract: contract}, nil
}

// NewControlledTransactor creates a new write-only instance of Controlled, bound to a specific deployed contract.
func NewControlledTransactor(address common.Address, transactor bind.ContractTransactor) (*ControlledTransactor, error) {
	contract, err := bindControlled(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControlledTransactor{contract: contract}, nil
}

// NewControlledFilterer creates a new log filterer instance of Controlled, bound to a specific deployed contract.
func NewControlledFilterer(address common.Address, filterer bind.ContractFilterer) (*ControlledFilterer, error) {
	contract, err := bindControlled(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControlledFilterer{contract: contract}, nil
}

// bindControlled binds a generic wrapper to an already deployed contract.
func bindControlled(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControlledABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controlled *ControlledRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controlled.Contract.ControlledCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controlled *ControlledRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controlled.Contract.ControlledTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controlled *ControlledRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controlled.Contract.ControlledTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controlled *ControlledCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controlled.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controlled *ControlledTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controlled.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controlled *ControlledTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controlled.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Controlled *ControlledCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Controlled.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Controlled *ControlledSession) GetController() (common.Address, error) {
	return _Controlled.Contract.GetController(&_Controlled.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Controlled *ControlledCallerSession) GetController() (common.Address, error) {
	return _Controlled.Contract.GetController(&_Controlled.CallOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Controlled *ControlledTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Controlled.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Controlled *ControlledSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Controlled.Contract.SetController(&_Controlled.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Controlled *ControlledTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Controlled.Contract.SetController(&_Controlled.TransactOpts, _controller)
}

// IControlledABI is the input ABI used to generate the binding from.
const IControlledABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IControlledBin is the compiled bytecode used for deploying new contracts.
const IControlledBin = `0x`

// DeployIControlled deploys a new Ethereum contract, binding an instance of IControlled to it.
func DeployIControlled(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IControlled, error) {
	parsed, err := abi.JSON(strings.NewReader(IControlledABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IControlledBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IControlled{IControlledCaller: IControlledCaller{contract: contract}, IControlledTransactor: IControlledTransactor{contract: contract}, IControlledFilterer: IControlledFilterer{contract: contract}}, nil
}

// IControlled is an auto generated Go binding around an Ethereum contract.
type IControlled struct {
	IControlledCaller     // Read-only binding to the contract
	IControlledTransactor // Write-only binding to the contract
	IControlledFilterer   // Log filterer for contract events
}

// IControlledCaller is an auto generated read-only Go binding around an Ethereum contract.
type IControlledCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlledTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IControlledTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlledFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControlledFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlledSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControlledSession struct {
	Contract     *IControlled      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControlledCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControlledCallerSession struct {
	Contract *IControlledCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IControlledTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControlledTransactorSession struct {
	Contract     *IControlledTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IControlledRaw is an auto generated low-level Go binding around an Ethereum contract.
type IControlledRaw struct {
	Contract *IControlled // Generic contract binding to access the raw methods on
}

// IControlledCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControlledCallerRaw struct {
	Contract *IControlledCaller // Generic read-only contract binding to access the raw methods on
}

// IControlledTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControlledTransactorRaw struct {
	Contract *IControlledTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIControlled creates a new instance of IControlled, bound to a specific deployed contract.
func NewIControlled(address common.Address, backend bind.ContractBackend) (*IControlled, error) {
	contract, err := bindIControlled(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControlled{IControlledCaller: IControlledCaller{contract: contract}, IControlledTransactor: IControlledTransactor{contract: contract}, IControlledFilterer: IControlledFilterer{contract: contract}}, nil
}

// NewIControlledCaller creates a new read-only instance of IControlled, bound to a specific deployed contract.
func NewIControlledCaller(address common.Address, caller bind.ContractCaller) (*IControlledCaller, error) {
	contract, err := bindIControlled(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControlledCaller{contract: contract}, nil
}

// NewIControlledTransactor creates a new write-only instance of IControlled, bound to a specific deployed contract.
func NewIControlledTransactor(address common.Address, transactor bind.ContractTransactor) (*IControlledTransactor, error) {
	contract, err := bindIControlled(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControlledTransactor{contract: contract}, nil
}

// NewIControlledFilterer creates a new log filterer instance of IControlled, bound to a specific deployed contract.
func NewIControlledFilterer(address common.Address, filterer bind.ContractFilterer) (*IControlledFilterer, error) {
	contract, err := bindIControlled(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControlledFilterer{contract: contract}, nil
}

// bindIControlled binds a generic wrapper to an already deployed contract.
func bindIControlled(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControlledABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControlled *IControlledRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IControlled.Contract.IControlledCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControlled *IControlledRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControlled.Contract.IControlledTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControlled *IControlledRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControlled.Contract.IControlledTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControlled *IControlledCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IControlled.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControlled *IControlledTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControlled.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControlled *IControlledTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControlled.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_IControlled *IControlledCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IControlled.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_IControlled *IControlledSession) GetController() (common.Address, error) {
	return _IControlled.Contract.GetController(&_IControlled.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_IControlled *IControlledCallerSession) GetController() (common.Address, error) {
	return _IControlled.Contract.GetController(&_IControlled.CallOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_IControlled *IControlledTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _IControlled.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_IControlled *IControlledSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _IControlled.Contract.SetController(&_IControlled.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_IControlled *IControlledTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _IControlled.Contract.SetController(&_IControlled.TransactOpts, _controller)
}
