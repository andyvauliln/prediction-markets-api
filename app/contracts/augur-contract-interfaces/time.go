package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ITimeABI is the input ABI used to generate the binding from.
const ITimeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ITimeBin is the compiled bytecode used for deploying new contracts.
const ITimeBin = `0x`

// DeployITime deploys a new Ethereum contract, binding an instance of ITime to it.
func DeployITime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ITime, error) {
	parsed, err := abi.JSON(strings.NewReader(ITimeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ITimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ITime{ITimeCaller: ITimeCaller{contract: contract}, ITimeTransactor: ITimeTransactor{contract: contract}, ITimeFilterer: ITimeFilterer{contract: contract}}, nil
}

// ITime is an auto generated Go binding around an Ethereum contract.
type ITime struct {
	ITimeCaller     // Read-only binding to the contract
	ITimeTransactor // Write-only binding to the contract
	ITimeFilterer   // Log filterer for contract events
}

// ITimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITimeSession struct {
	Contract     *ITime            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITimeCallerSession struct {
	Contract *ITimeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITimeTransactorSession struct {
	Contract     *ITimeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITimeRaw struct {
	Contract *ITime // Generic contract binding to access the raw methods on
}

// ITimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITimeCallerRaw struct {
	Contract *ITimeCaller // Generic read-only contract binding to access the raw methods on
}

// ITimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITimeTransactorRaw struct {
	Contract *ITimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITime creates a new instance of ITime, bound to a specific deployed contract.
func NewITime(address common.Address, backend bind.ContractBackend) (*ITime, error) {
	contract, err := bindITime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITime{ITimeCaller: ITimeCaller{contract: contract}, ITimeTransactor: ITimeTransactor{contract: contract}, ITimeFilterer: ITimeFilterer{contract: contract}}, nil
}

// NewITimeCaller creates a new read-only instance of ITime, bound to a specific deployed contract.
func NewITimeCaller(address common.Address, caller bind.ContractCaller) (*ITimeCaller, error) {
	contract, err := bindITime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITimeCaller{contract: contract}, nil
}

// NewITimeTransactor creates a new write-only instance of ITime, bound to a specific deployed contract.
func NewITimeTransactor(address common.Address, transactor bind.ContractTransactor) (*ITimeTransactor, error) {
	contract, err := bindITime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITimeTransactor{contract: contract}, nil
}

// NewITimeFilterer creates a new log filterer instance of ITime, bound to a specific deployed contract.
func NewITimeFilterer(address common.Address, filterer bind.ContractFilterer) (*ITimeFilterer, error) {
	contract, err := bindITime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITimeFilterer{contract: contract}, nil
}

// bindITime binds a generic wrapper to an already deployed contract.
func bindITime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITime *ITimeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITime.Contract.ITimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITime *ITimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITime.Contract.ITimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITime *ITimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITime.Contract.ITimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITime *ITimeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITime *ITimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITime *ITimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITime.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ITime *ITimeCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ITime.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ITime *ITimeSession) GetController() (common.Address, error) {
	return _ITime.Contract.GetController(&_ITime.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ITime *ITimeCallerSession) GetController() (common.Address, error) {
	return _ITime.Contract.GetController(&_ITime.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_ITime *ITimeCaller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ITime.contract.Call(opts, out, "getTimestamp")
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_ITime *ITimeSession) GetTimestamp() (*big.Int, error) {
	return _ITime.Contract.GetTimestamp(&_ITime.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_ITime *ITimeCallerSession) GetTimestamp() (*big.Int, error) {
	return _ITime.Contract.GetTimestamp(&_ITime.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ITime *ITimeCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ITime.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ITime *ITimeSession) GetTypeName() ([32]byte, error) {
	return _ITime.Contract.GetTypeName(&_ITime.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_ITime *ITimeCallerSession) GetTypeName() ([32]byte, error) {
	return _ITime.Contract.GetTypeName(&_ITime.CallOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ITime *ITimeTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _ITime.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ITime *ITimeSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ITime.Contract.SetController(&_ITime.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ITime *ITimeTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ITime.Contract.SetController(&_ITime.TransactOpts, _controller)
}

// TimeABI is the input ABI used to generate the binding from.
const TimeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TimeBin is the compiled bytecode used for deploying new contracts.
const TimeBin = `0x608060405260008054600160a060020a031916331790556101f6806100256000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663188ec35681146100665780633018205f1461008d57806392eefe9b146100cb578063db0a087c1461010d575b600080fd5b34801561007257600080fd5b5061007b610122565b60408051918252519081900360200190f35b34801561009957600080fd5b506100a2610126565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100d757600080fd5b506100f973ffffffffffffffffffffffffffffffffffffffff60043516610142565b604080519115158252519081900360200190f35b34801561011957600080fd5b5061007b6101a6565b4290565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461016757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b7f54696d6500000000000000000000000000000000000000000000000000000000905600a165627a7a7230582039da3abdbbb3da20737b57f6b2b3ba92c203cf64d56b8a970f73d6c89ccf6a460029`

// DeployTime deploys a new Ethereum contract, binding an instance of Time to it.
func DeployTime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Time, error) {
	parsed, err := abi.JSON(strings.NewReader(TimeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Time{TimeCaller: TimeCaller{contract: contract}, TimeTransactor: TimeTransactor{contract: contract}, TimeFilterer: TimeFilterer{contract: contract}}, nil
}

// Time is an auto generated Go binding around an Ethereum contract.
type Time struct {
	TimeCaller     // Read-only binding to the contract
	TimeTransactor // Write-only binding to the contract
	TimeFilterer   // Log filterer for contract events
}

// TimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeSession struct {
	Contract     *Time             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeCallerSession struct {
	Contract *TimeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeTransactorSession struct {
	Contract     *TimeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeRaw struct {
	Contract *Time // Generic contract binding to access the raw methods on
}

// TimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeCallerRaw struct {
	Contract *TimeCaller // Generic read-only contract binding to access the raw methods on
}

// TimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeTransactorRaw struct {
	Contract *TimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTime creates a new instance of Time, bound to a specific deployed contract.
func NewTime(address common.Address, backend bind.ContractBackend) (*Time, error) {
	contract, err := bindTime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Time{TimeCaller: TimeCaller{contract: contract}, TimeTransactor: TimeTransactor{contract: contract}, TimeFilterer: TimeFilterer{contract: contract}}, nil
}

// NewTimeCaller creates a new read-only instance of Time, bound to a specific deployed contract.
func NewTimeCaller(address common.Address, caller bind.ContractCaller) (*TimeCaller, error) {
	contract, err := bindTime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeCaller{contract: contract}, nil
}

// NewTimeTransactor creates a new write-only instance of Time, bound to a specific deployed contract.
func NewTimeTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeTransactor, error) {
	contract, err := bindTime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeTransactor{contract: contract}, nil
}

// NewTimeFilterer creates a new log filterer instance of Time, bound to a specific deployed contract.
func NewTimeFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeFilterer, error) {
	contract, err := bindTime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeFilterer{contract: contract}, nil
}

// bindTime binds a generic wrapper to an already deployed contract.
func bindTime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Time *TimeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Time.Contract.TimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Time *TimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Time.Contract.TimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Time *TimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Time.Contract.TimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Time *TimeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Time.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Time *TimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Time.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Time *TimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Time.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Time *TimeCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Time.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Time *TimeSession) GetController() (common.Address, error) {
	return _Time.Contract.GetController(&_Time.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Time *TimeCallerSession) GetController() (common.Address, error) {
	return _Time.Contract.GetController(&_Time.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_Time *TimeCaller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Time.contract.Call(opts, out, "getTimestamp")
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_Time *TimeSession) GetTimestamp() (*big.Int, error) {
	return _Time.Contract.GetTimestamp(&_Time.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_Time *TimeCallerSession) GetTimestamp() (*big.Int, error) {
	return _Time.Contract.GetTimestamp(&_Time.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_Time *TimeCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Time.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_Time *TimeSession) GetTypeName() ([32]byte, error) {
	return _Time.Contract.GetTypeName(&_Time.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_Time *TimeCallerSession) GetTypeName() ([32]byte, error) {
	return _Time.Contract.GetTypeName(&_Time.CallOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Time *TimeTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Time.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Time *TimeSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Time.Contract.SetController(&_Time.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Time *TimeTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Time.Contract.SetController(&_Time.TransactOpts, _controller)
}

// TimeControlledABI is the input ABI used to generate the binding from.
const TimeControlledABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"setTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TimeControlledBin is the compiled bytecode used for deploying new contracts.
const TimeControlledBin = `0x6080604052600160025534801561001557600080fd5b506000805433600160a060020a03199182168117909255600180549091169091179055610062731985365e9f78359a9b6ad760e32412f4a445e862640100000000610075810261047b1704565b1561006c57600080fd5b4260025561007d565b6000903b1190565b6104af8061008c6000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663188ec35681146100925780633018205f146100b9578063893d20e8146100ea57806392eefe9b146100ff578063a0a2b57314610134578063b610c75e1461014c578063db0a087c14610164578063f2fde38b14610179575b600080fd5b34801561009e57600080fd5b506100a761019a565b60408051918252519081900360200190f35b3480156100c557600080fd5b506100ce6101a0565b60408051600160a060020a039092168252519081900360200190f35b3480156100f657600080fd5b506100ce6101af565b34801561010b57600080fd5b50610120600160a060020a03600435166101be565b604080519115158252519081900360200190f35b34801561014057600080fd5b50610120600435610208565b34801561015857600080fd5b5061012060043561035b565b34801561017057600080fd5b506100a76103de565b34801561018557600080fd5b50610120600160a060020a0360043516610402565b60025490565b600054600160a060020a031690565b600154600160a060020a031690565b60008054600160a060020a031633146101d657600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154600090600160a060020a0316331461022257600080fd5b816002819055506000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561029457600080fd5b505af11580156102a8573d6000803e3d6000fd5b505050506040513d60208110156102be57600080fd5b5051600254604080517fc8e6b2a8000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a039092169163c8e6b2a8916024808201926020929091908290030181600087803b15801561032757600080fd5b505af115801561033b573d6000803e3d6000fd5b505050506040513d602081101561035157600080fd5b5060019392505050565b600154600090600160a060020a0316331461037557600080fd5b600280548301905560008054604080517f4e94c8290000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921692634e94c829926004808401936020939083900390910190829087803b15801561029457600080fd5b7f54696d65436f6e74726f6c6c656400000000000000000000000000000000000090565b600154600090600160a060020a0316331461041c57600080fd5b600160a060020a0382161561046b5760015461044190600160a060020a031683610473565b506001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790555b506001919050565b600192915050565b6000903b11905600a165627a7a723058205b606499646a2282fefb6e13d2afc01be49d94f65468a269ffbedf5699a787730029`

// DeployTimeControlled deploys a new Ethereum contract, binding an instance of TimeControlled to it.
func DeployTimeControlled(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TimeControlled, error) {
	parsed, err := abi.JSON(strings.NewReader(TimeControlledABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TimeControlledBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TimeControlled{TimeControlledCaller: TimeControlledCaller{contract: contract}, TimeControlledTransactor: TimeControlledTransactor{contract: contract}, TimeControlledFilterer: TimeControlledFilterer{contract: contract}}, nil
}

// TimeControlled is an auto generated Go binding around an Ethereum contract.
type TimeControlled struct {
	TimeControlledCaller     // Read-only binding to the contract
	TimeControlledTransactor // Write-only binding to the contract
	TimeControlledFilterer   // Log filterer for contract events
}

// TimeControlledCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeControlledCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeControlledTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeControlledTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeControlledFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeControlledFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeControlledSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeControlledSession struct {
	Contract     *TimeControlled   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeControlledCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeControlledCallerSession struct {
	Contract *TimeControlledCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TimeControlledTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeControlledTransactorSession struct {
	Contract     *TimeControlledTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TimeControlledRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeControlledRaw struct {
	Contract *TimeControlled // Generic contract binding to access the raw methods on
}

// TimeControlledCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeControlledCallerRaw struct {
	Contract *TimeControlledCaller // Generic read-only contract binding to access the raw methods on
}

// TimeControlledTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeControlledTransactorRaw struct {
	Contract *TimeControlledTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimeControlled creates a new instance of TimeControlled, bound to a specific deployed contract.
func NewTimeControlled(address common.Address, backend bind.ContractBackend) (*TimeControlled, error) {
	contract, err := bindTimeControlled(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimeControlled{TimeControlledCaller: TimeControlledCaller{contract: contract}, TimeControlledTransactor: TimeControlledTransactor{contract: contract}, TimeControlledFilterer: TimeControlledFilterer{contract: contract}}, nil
}

// NewTimeControlledCaller creates a new read-only instance of TimeControlled, bound to a specific deployed contract.
func NewTimeControlledCaller(address common.Address, caller bind.ContractCaller) (*TimeControlledCaller, error) {
	contract, err := bindTimeControlled(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeControlledCaller{contract: contract}, nil
}

// NewTimeControlledTransactor creates a new write-only instance of TimeControlled, bound to a specific deployed contract.
func NewTimeControlledTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeControlledTransactor, error) {
	contract, err := bindTimeControlled(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeControlledTransactor{contract: contract}, nil
}

// NewTimeControlledFilterer creates a new log filterer instance of TimeControlled, bound to a specific deployed contract.
func NewTimeControlledFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeControlledFilterer, error) {
	contract, err := bindTimeControlled(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeControlledFilterer{contract: contract}, nil
}

// bindTimeControlled binds a generic wrapper to an already deployed contract.
func bindTimeControlled(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimeControlledABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeControlled *TimeControlledRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TimeControlled.Contract.TimeControlledCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeControlled *TimeControlledRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeControlled.Contract.TimeControlledTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeControlled *TimeControlledRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeControlled.Contract.TimeControlledTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeControlled *TimeControlledCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TimeControlled.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeControlled *TimeControlledTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeControlled.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeControlled *TimeControlledTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeControlled.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_TimeControlled *TimeControlledCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TimeControlled.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_TimeControlled *TimeControlledSession) GetController() (common.Address, error) {
	return _TimeControlled.Contract.GetController(&_TimeControlled.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_TimeControlled *TimeControlledCallerSession) GetController() (common.Address, error) {
	return _TimeControlled.Contract.GetController(&_TimeControlled.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_TimeControlled *TimeControlledCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TimeControlled.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_TimeControlled *TimeControlledSession) GetOwner() (common.Address, error) {
	return _TimeControlled.Contract.GetOwner(&_TimeControlled.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_TimeControlled *TimeControlledCallerSession) GetOwner() (common.Address, error) {
	return _TimeControlled.Contract.GetOwner(&_TimeControlled.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_TimeControlled *TimeControlledCaller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimeControlled.contract.Call(opts, out, "getTimestamp")
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_TimeControlled *TimeControlledSession) GetTimestamp() (*big.Int, error) {
	return _TimeControlled.Contract.GetTimestamp(&_TimeControlled.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_TimeControlled *TimeControlledCallerSession) GetTimestamp() (*big.Int, error) {
	return _TimeControlled.Contract.GetTimestamp(&_TimeControlled.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_TimeControlled *TimeControlledCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TimeControlled.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_TimeControlled *TimeControlledSession) GetTypeName() ([32]byte, error) {
	return _TimeControlled.Contract.GetTypeName(&_TimeControlled.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_TimeControlled *TimeControlledCallerSession) GetTypeName() ([32]byte, error) {
	return _TimeControlled.Contract.GetTypeName(&_TimeControlled.CallOpts)
}

// IncrementTimestamp is a paid mutator transaction binding the contract method 0xb610c75e.
//
// Solidity: function incrementTimestamp(_amount uint256) returns(bool)
func (_TimeControlled *TimeControlledTransactor) IncrementTimestamp(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _TimeControlled.contract.Transact(opts, "incrementTimestamp", _amount)
}

// IncrementTimestamp is a paid mutator transaction binding the contract method 0xb610c75e.
//
// Solidity: function incrementTimestamp(_amount uint256) returns(bool)
func (_TimeControlled *TimeControlledSession) IncrementTimestamp(_amount *big.Int) (*types.Transaction, error) {
	return _TimeControlled.Contract.IncrementTimestamp(&_TimeControlled.TransactOpts, _amount)
}

// IncrementTimestamp is a paid mutator transaction binding the contract method 0xb610c75e.
//
// Solidity: function incrementTimestamp(_amount uint256) returns(bool)
func (_TimeControlled *TimeControlledTransactorSession) IncrementTimestamp(_amount *big.Int) (*types.Transaction, error) {
	return _TimeControlled.Contract.IncrementTimestamp(&_TimeControlled.TransactOpts, _amount)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_TimeControlled *TimeControlledTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _TimeControlled.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_TimeControlled *TimeControlledSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _TimeControlled.Contract.SetController(&_TimeControlled.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_TimeControlled *TimeControlledTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _TimeControlled.Contract.SetController(&_TimeControlled.TransactOpts, _controller)
}

// SetTimestamp is a paid mutator transaction binding the contract method 0xa0a2b573.
//
// Solidity: function setTimestamp(_timestamp uint256) returns(bool)
func (_TimeControlled *TimeControlledTransactor) SetTimestamp(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _TimeControlled.contract.Transact(opts, "setTimestamp", _timestamp)
}

// SetTimestamp is a paid mutator transaction binding the contract method 0xa0a2b573.
//
// Solidity: function setTimestamp(_timestamp uint256) returns(bool)
func (_TimeControlled *TimeControlledSession) SetTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _TimeControlled.Contract.SetTimestamp(&_TimeControlled.TransactOpts, _timestamp)
}

// SetTimestamp is a paid mutator transaction binding the contract method 0xa0a2b573.
//
// Solidity: function setTimestamp(_timestamp uint256) returns(bool)
func (_TimeControlled *TimeControlledTransactorSession) SetTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _TimeControlled.Contract.SetTimestamp(&_TimeControlled.TransactOpts, _timestamp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_TimeControlled *TimeControlledTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _TimeControlled.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_TimeControlled *TimeControlledSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _TimeControlled.Contract.TransferOwnership(&_TimeControlled.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_TimeControlled *TimeControlledTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _TimeControlled.Contract.TransferOwnership(&_TimeControlled.TransactOpts, _newOwner)
}
