package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IRepPriceOracleABI is the input ABI used to generate the binding from.
const IRepPriceOracleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getRepPriceInAttoEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repPriceInAttoEth\",\"type\":\"uint256\"}],\"name\":\"setRepPriceInAttoEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IRepPriceOracleBin is the compiled bytecode used for deploying new contracts.
const IRepPriceOracleBin = `0x`

// DeployIRepPriceOracle deploys a new Ethereum contract, binding an instance of IRepPriceOracle to it.
func DeployIRepPriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IRepPriceOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(IRepPriceOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IRepPriceOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IRepPriceOracle{IRepPriceOracleCaller: IRepPriceOracleCaller{contract: contract}, IRepPriceOracleTransactor: IRepPriceOracleTransactor{contract: contract}, IRepPriceOracleFilterer: IRepPriceOracleFilterer{contract: contract}}, nil
}

// IRepPriceOracle is an auto generated Go binding around an Ethereum contract.
type IRepPriceOracle struct {
	IRepPriceOracleCaller     // Read-only binding to the contract
	IRepPriceOracleTransactor // Write-only binding to the contract
	IRepPriceOracleFilterer   // Log filterer for contract events
}

// IRepPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRepPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRepPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRepPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRepPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRepPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRepPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRepPriceOracleSession struct {
	Contract     *IRepPriceOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRepPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRepPriceOracleCallerSession struct {
	Contract *IRepPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IRepPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRepPriceOracleTransactorSession struct {
	Contract     *IRepPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IRepPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRepPriceOracleRaw struct {
	Contract *IRepPriceOracle // Generic contract binding to access the raw methods on
}

// IRepPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRepPriceOracleCallerRaw struct {
	Contract *IRepPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IRepPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRepPriceOracleTransactorRaw struct {
	Contract *IRepPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRepPriceOracle creates a new instance of IRepPriceOracle, bound to a specific deployed contract.
func NewIRepPriceOracle(address common.Address, backend bind.ContractBackend) (*IRepPriceOracle, error) {
	contract, err := bindIRepPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRepPriceOracle{IRepPriceOracleCaller: IRepPriceOracleCaller{contract: contract}, IRepPriceOracleTransactor: IRepPriceOracleTransactor{contract: contract}, IRepPriceOracleFilterer: IRepPriceOracleFilterer{contract: contract}}, nil
}

// NewIRepPriceOracleCaller creates a new read-only instance of IRepPriceOracle, bound to a specific deployed contract.
func NewIRepPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*IRepPriceOracleCaller, error) {
	contract, err := bindIRepPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRepPriceOracleCaller{contract: contract}, nil
}

// NewIRepPriceOracleTransactor creates a new write-only instance of IRepPriceOracle, bound to a specific deployed contract.
func NewIRepPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IRepPriceOracleTransactor, error) {
	contract, err := bindIRepPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRepPriceOracleTransactor{contract: contract}, nil
}

// NewIRepPriceOracleFilterer creates a new log filterer instance of IRepPriceOracle, bound to a specific deployed contract.
func NewIRepPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IRepPriceOracleFilterer, error) {
	contract, err := bindIRepPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRepPriceOracleFilterer{contract: contract}, nil
}

// bindIRepPriceOracle binds a generic wrapper to an already deployed contract.
func bindIRepPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRepPriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRepPriceOracle *IRepPriceOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRepPriceOracle.Contract.IRepPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRepPriceOracle *IRepPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRepPriceOracle.Contract.IRepPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRepPriceOracle *IRepPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRepPriceOracle.Contract.IRepPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRepPriceOracle *IRepPriceOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRepPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRepPriceOracle *IRepPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRepPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRepPriceOracle *IRepPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRepPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_IRepPriceOracle *IRepPriceOracleCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IRepPriceOracle.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_IRepPriceOracle *IRepPriceOracleSession) GetOwner() (common.Address, error) {
	return _IRepPriceOracle.Contract.GetOwner(&_IRepPriceOracle.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_IRepPriceOracle *IRepPriceOracleCallerSession) GetOwner() (common.Address, error) {
	return _IRepPriceOracle.Contract.GetOwner(&_IRepPriceOracle.CallOpts)
}

// GetRepPriceInAttoEth is a free data retrieval call binding the contract method 0x0f8fd363.
//
// Solidity: function getRepPriceInAttoEth() constant returns(uint256)
func (_IRepPriceOracle *IRepPriceOracleCaller) GetRepPriceInAttoEth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRepPriceOracle.contract.Call(opts, out, "getRepPriceInAttoEth")
	return *ret0, err
}

// GetRepPriceInAttoEth is a free data retrieval call binding the contract method 0x0f8fd363.
//
// Solidity: function getRepPriceInAttoEth() constant returns(uint256)
func (_IRepPriceOracle *IRepPriceOracleSession) GetRepPriceInAttoEth() (*big.Int, error) {
	return _IRepPriceOracle.Contract.GetRepPriceInAttoEth(&_IRepPriceOracle.CallOpts)
}

// GetRepPriceInAttoEth is a free data retrieval call binding the contract method 0x0f8fd363.
//
// Solidity: function getRepPriceInAttoEth() constant returns(uint256)
func (_IRepPriceOracle *IRepPriceOracleCallerSession) GetRepPriceInAttoEth() (*big.Int, error) {
	return _IRepPriceOracle.Contract.GetRepPriceInAttoEth(&_IRepPriceOracle.CallOpts)
}

// SetRepPriceInAttoEth is a paid mutator transaction binding the contract method 0xac4d8a26.
//
// Solidity: function setRepPriceInAttoEth(_repPriceInAttoEth uint256) returns(bool)
func (_IRepPriceOracle *IRepPriceOracleTransactor) SetRepPriceInAttoEth(opts *bind.TransactOpts, _repPriceInAttoEth *big.Int) (*types.Transaction, error) {
	return _IRepPriceOracle.contract.Transact(opts, "setRepPriceInAttoEth", _repPriceInAttoEth)
}

// SetRepPriceInAttoEth is a paid mutator transaction binding the contract method 0xac4d8a26.
//
// Solidity: function setRepPriceInAttoEth(_repPriceInAttoEth uint256) returns(bool)
func (_IRepPriceOracle *IRepPriceOracleSession) SetRepPriceInAttoEth(_repPriceInAttoEth *big.Int) (*types.Transaction, error) {
	return _IRepPriceOracle.Contract.SetRepPriceInAttoEth(&_IRepPriceOracle.TransactOpts, _repPriceInAttoEth)
}

// SetRepPriceInAttoEth is a paid mutator transaction binding the contract method 0xac4d8a26.
//
// Solidity: function setRepPriceInAttoEth(_repPriceInAttoEth uint256) returns(bool)
func (_IRepPriceOracle *IRepPriceOracleTransactorSession) SetRepPriceInAttoEth(_repPriceInAttoEth *big.Int) (*types.Transaction, error) {
	return _IRepPriceOracle.Contract.SetRepPriceInAttoEth(&_IRepPriceOracle.TransactOpts, _repPriceInAttoEth)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_IRepPriceOracle *IRepPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _IRepPriceOracle.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_IRepPriceOracle *IRepPriceOracleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _IRepPriceOracle.Contract.TransferOwnership(&_IRepPriceOracle.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_IRepPriceOracle *IRepPriceOracleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _IRepPriceOracle.Contract.TransferOwnership(&_IRepPriceOracle.TransactOpts, _newOwner)
}

// RepPriceOracleABI is the input ABI used to generate the binding from.
const RepPriceOracleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getRepPriceInAttoEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repPriceInAttoEth\",\"type\":\"uint256\"}],\"name\":\"setRepPriceInAttoEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RepPriceOracleBin is the compiled bytecode used for deploying new contracts.
const RepPriceOracleBin = `0x608060405266d529ae9e86000060015560008054600160a060020a031916331790556101e4806100306000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630f8fd3638114610066578063893d20e81461008d578063ac4d8a26146100be578063f2fde38b146100ea575b600080fd5b34801561007257600080fd5b5061007b61010b565b60408051918252519081900360200190f35b34801561009957600080fd5b506100a2610111565b60408051600160a060020a039092168252519081900360200190f35b3480156100ca57600080fd5b506100d6600435610120565b604080519115158252519081900360200190f35b3480156100f657600080fd5b506100d6600160a060020a0360043516610141565b60015490565b600054600160a060020a031690565b60008054600160a060020a0316331461013857600080fd5b50600190815590565b60008054600160a060020a0316331461015957600080fd5b600160a060020a038216156101a85760005461017e90600160a060020a0316836101b0565b506000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790555b506001919050565b6001929150505600a165627a7a723058202dbd09fbd8c7ff3b00bee39807fb2a2c7ef66ecc73969e2617633a0a4fc0a8a30029`

// DeployRepPriceOracle deploys a new Ethereum contract, binding an instance of RepPriceOracle to it.
func DeployRepPriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RepPriceOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(RepPriceOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RepPriceOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RepPriceOracle{RepPriceOracleCaller: RepPriceOracleCaller{contract: contract}, RepPriceOracleTransactor: RepPriceOracleTransactor{contract: contract}, RepPriceOracleFilterer: RepPriceOracleFilterer{contract: contract}}, nil
}

// RepPriceOracle is an auto generated Go binding around an Ethereum contract.
type RepPriceOracle struct {
	RepPriceOracleCaller     // Read-only binding to the contract
	RepPriceOracleTransactor // Write-only binding to the contract
	RepPriceOracleFilterer   // Log filterer for contract events
}

// RepPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RepPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RepPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RepPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RepPriceOracleSession struct {
	Contract     *RepPriceOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RepPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RepPriceOracleCallerSession struct {
	Contract *RepPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RepPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RepPriceOracleTransactorSession struct {
	Contract     *RepPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RepPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RepPriceOracleRaw struct {
	Contract *RepPriceOracle // Generic contract binding to access the raw methods on
}

// RepPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RepPriceOracleCallerRaw struct {
	Contract *RepPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// RepPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RepPriceOracleTransactorRaw struct {
	Contract *RepPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRepPriceOracle creates a new instance of RepPriceOracle, bound to a specific deployed contract.
func NewRepPriceOracle(address common.Address, backend bind.ContractBackend) (*RepPriceOracle, error) {
	contract, err := bindRepPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RepPriceOracle{RepPriceOracleCaller: RepPriceOracleCaller{contract: contract}, RepPriceOracleTransactor: RepPriceOracleTransactor{contract: contract}, RepPriceOracleFilterer: RepPriceOracleFilterer{contract: contract}}, nil
}

// NewRepPriceOracleCaller creates a new read-only instance of RepPriceOracle, bound to a specific deployed contract.
func NewRepPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*RepPriceOracleCaller, error) {
	contract, err := bindRepPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RepPriceOracleCaller{contract: contract}, nil
}

// NewRepPriceOracleTransactor creates a new write-only instance of RepPriceOracle, bound to a specific deployed contract.
func NewRepPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*RepPriceOracleTransactor, error) {
	contract, err := bindRepPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RepPriceOracleTransactor{contract: contract}, nil
}

// NewRepPriceOracleFilterer creates a new log filterer instance of RepPriceOracle, bound to a specific deployed contract.
func NewRepPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*RepPriceOracleFilterer, error) {
	contract, err := bindRepPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RepPriceOracleFilterer{contract: contract}, nil
}

// bindRepPriceOracle binds a generic wrapper to an already deployed contract.
func bindRepPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RepPriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepPriceOracle *RepPriceOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepPriceOracle.Contract.RepPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepPriceOracle *RepPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepPriceOracle.Contract.RepPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepPriceOracle *RepPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepPriceOracle.Contract.RepPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepPriceOracle *RepPriceOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepPriceOracle *RepPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepPriceOracle *RepPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_RepPriceOracle *RepPriceOracleCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepPriceOracle.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_RepPriceOracle *RepPriceOracleSession) GetOwner() (common.Address, error) {
	return _RepPriceOracle.Contract.GetOwner(&_RepPriceOracle.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_RepPriceOracle *RepPriceOracleCallerSession) GetOwner() (common.Address, error) {
	return _RepPriceOracle.Contract.GetOwner(&_RepPriceOracle.CallOpts)
}

// GetRepPriceInAttoEth is a free data retrieval call binding the contract method 0x0f8fd363.
//
// Solidity: function getRepPriceInAttoEth() constant returns(uint256)
func (_RepPriceOracle *RepPriceOracleCaller) GetRepPriceInAttoEth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepPriceOracle.contract.Call(opts, out, "getRepPriceInAttoEth")
	return *ret0, err
}

// GetRepPriceInAttoEth is a free data retrieval call binding the contract method 0x0f8fd363.
//
// Solidity: function getRepPriceInAttoEth() constant returns(uint256)
func (_RepPriceOracle *RepPriceOracleSession) GetRepPriceInAttoEth() (*big.Int, error) {
	return _RepPriceOracle.Contract.GetRepPriceInAttoEth(&_RepPriceOracle.CallOpts)
}

// GetRepPriceInAttoEth is a free data retrieval call binding the contract method 0x0f8fd363.
//
// Solidity: function getRepPriceInAttoEth() constant returns(uint256)
func (_RepPriceOracle *RepPriceOracleCallerSession) GetRepPriceInAttoEth() (*big.Int, error) {
	return _RepPriceOracle.Contract.GetRepPriceInAttoEth(&_RepPriceOracle.CallOpts)
}

// SetRepPriceInAttoEth is a paid mutator transaction binding the contract method 0xac4d8a26.
//
// Solidity: function setRepPriceInAttoEth(_repPriceInAttoEth uint256) returns(bool)
func (_RepPriceOracle *RepPriceOracleTransactor) SetRepPriceInAttoEth(opts *bind.TransactOpts, _repPriceInAttoEth *big.Int) (*types.Transaction, error) {
	return _RepPriceOracle.contract.Transact(opts, "setRepPriceInAttoEth", _repPriceInAttoEth)
}

// SetRepPriceInAttoEth is a paid mutator transaction binding the contract method 0xac4d8a26.
//
// Solidity: function setRepPriceInAttoEth(_repPriceInAttoEth uint256) returns(bool)
func (_RepPriceOracle *RepPriceOracleSession) SetRepPriceInAttoEth(_repPriceInAttoEth *big.Int) (*types.Transaction, error) {
	return _RepPriceOracle.Contract.SetRepPriceInAttoEth(&_RepPriceOracle.TransactOpts, _repPriceInAttoEth)
}

// SetRepPriceInAttoEth is a paid mutator transaction binding the contract method 0xac4d8a26.
//
// Solidity: function setRepPriceInAttoEth(_repPriceInAttoEth uint256) returns(bool)
func (_RepPriceOracle *RepPriceOracleTransactorSession) SetRepPriceInAttoEth(_repPriceInAttoEth *big.Int) (*types.Transaction, error) {
	return _RepPriceOracle.Contract.SetRepPriceInAttoEth(&_RepPriceOracle.TransactOpts, _repPriceInAttoEth)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RepPriceOracle *RepPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RepPriceOracle.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RepPriceOracle *RepPriceOracleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RepPriceOracle.Contract.TransferOwnership(&_RepPriceOracle.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RepPriceOracle *RepPriceOracleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RepPriceOracle.Contract.TransferOwnership(&_RepPriceOracle.TransactOpts, _newOwner)
}
