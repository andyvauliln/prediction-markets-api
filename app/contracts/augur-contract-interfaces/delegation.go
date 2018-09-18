package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DelegationTargetABI is the input ABI used to generate the binding from.
const DelegationTargetABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DelegationTargetBin is the compiled bytecode used for deploying new contracts.
const DelegationTargetBin = `0x608060405260008054600160a060020a031916331790556101b4806100256000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461005b57806392eefe9b14610099578063bef72fa2146100db575b600080fd5b34801561006757600080fd5b50610070610102565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100a557600080fd5b506100c773ffffffffffffffffffffffffffffffffffffffff6004351661011e565b604080519115158252519081900360200190f35b3480156100e757600080fd5b506100f0610182565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461014357600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a723058202750fc39d056f0618c5cf2e66de80bd05710071e099269165c938e596db930850029`

// DeployDelegationTarget deploys a new Ethereum contract, binding an instance of DelegationTarget to it.
func DeployDelegationTarget(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DelegationTarget, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegationTargetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DelegationTargetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelegationTarget{DelegationTargetCaller: DelegationTargetCaller{contract: contract}, DelegationTargetTransactor: DelegationTargetTransactor{contract: contract}, DelegationTargetFilterer: DelegationTargetFilterer{contract: contract}}, nil
}

// DelegationTarget is an auto generated Go binding around an Ethereum contract.
type DelegationTarget struct {
	DelegationTargetCaller     // Read-only binding to the contract
	DelegationTargetTransactor // Write-only binding to the contract
	DelegationTargetFilterer   // Log filterer for contract events
}

// DelegationTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegationTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegationTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegationTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegationTargetSession struct {
	Contract     *DelegationTarget // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegationTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegationTargetCallerSession struct {
	Contract *DelegationTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DelegationTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegationTargetTransactorSession struct {
	Contract     *DelegationTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DelegationTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegationTargetRaw struct {
	Contract *DelegationTarget // Generic contract binding to access the raw methods on
}

// DelegationTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegationTargetCallerRaw struct {
	Contract *DelegationTargetCaller // Generic read-only contract binding to access the raw methods on
}

// DelegationTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegationTargetTransactorRaw struct {
	Contract *DelegationTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegationTarget creates a new instance of DelegationTarget, bound to a specific deployed contract.
func NewDelegationTarget(address common.Address, backend bind.ContractBackend) (*DelegationTarget, error) {
	contract, err := bindDelegationTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegationTarget{DelegationTargetCaller: DelegationTargetCaller{contract: contract}, DelegationTargetTransactor: DelegationTargetTransactor{contract: contract}, DelegationTargetFilterer: DelegationTargetFilterer{contract: contract}}, nil
}

// NewDelegationTargetCaller creates a new read-only instance of DelegationTarget, bound to a specific deployed contract.
func NewDelegationTargetCaller(address common.Address, caller bind.ContractCaller) (*DelegationTargetCaller, error) {
	contract, err := bindDelegationTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegationTargetCaller{contract: contract}, nil
}

// NewDelegationTargetTransactor creates a new write-only instance of DelegationTarget, bound to a specific deployed contract.
func NewDelegationTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegationTargetTransactor, error) {
	contract, err := bindDelegationTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegationTargetTransactor{contract: contract}, nil
}

// NewDelegationTargetFilterer creates a new log filterer instance of DelegationTarget, bound to a specific deployed contract.
func NewDelegationTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegationTargetFilterer, error) {
	contract, err := bindDelegationTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegationTargetFilterer{contract: contract}, nil
}

// bindDelegationTarget binds a generic wrapper to an already deployed contract.
func bindDelegationTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegationTargetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegationTarget *DelegationTargetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DelegationTarget.Contract.DelegationTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegationTarget *DelegationTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationTarget.Contract.DelegationTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegationTarget *DelegationTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegationTarget.Contract.DelegationTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegationTarget *DelegationTargetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DelegationTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegationTarget *DelegationTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegationTarget *DelegationTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegationTarget.Contract.contract.Transact(opts, method, params...)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_DelegationTarget *DelegationTargetCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DelegationTarget.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_DelegationTarget *DelegationTargetSession) ControllerLookupName() ([32]byte, error) {
	return _DelegationTarget.Contract.ControllerLookupName(&_DelegationTarget.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_DelegationTarget *DelegationTargetCallerSession) ControllerLookupName() ([32]byte, error) {
	return _DelegationTarget.Contract.ControllerLookupName(&_DelegationTarget.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_DelegationTarget *DelegationTargetCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DelegationTarget.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_DelegationTarget *DelegationTargetSession) GetController() (common.Address, error) {
	return _DelegationTarget.Contract.GetController(&_DelegationTarget.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_DelegationTarget *DelegationTargetCallerSession) GetController() (common.Address, error) {
	return _DelegationTarget.Contract.GetController(&_DelegationTarget.CallOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_DelegationTarget *DelegationTargetTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _DelegationTarget.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_DelegationTarget *DelegationTargetSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _DelegationTarget.Contract.SetController(&_DelegationTarget.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_DelegationTarget *DelegationTargetTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _DelegationTarget.Contract.SetController(&_DelegationTarget.TransactOpts, _controller)
}

// DelegatorABI is the input ABI used to generate the binding from.
const DelegatorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_controllerLookupName\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// DelegatorBin is the compiled bytecode used for deploying new contracts.
const DelegatorBin = `0x608060405234801561001057600080fd5b506040516040806103088339810160405280516020909101516000805433600160a060020a03199182161716600160a060020a039093169290921782556001556102a890819061006090396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461014f57806392eefe9b1461018d578063bef72fa2146101cf575b60015460009015156100675761014c565b60008054600154604080517ff39ec1f700000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff9092169263f39ec1f7926024808401936020939083900390910190829087803b1580156100df57600080fd5b505af11580156100f3573d6000803e3d6000fd5b505050506040513d602081101561010957600080fd5b50516040805136601f8101601f191680830190935292935091600083376000803684865af4808015610147576040513d81016040523d6000823e3d81f35b600080fd5b50005b34801561015b57600080fd5b506101646101f6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019957600080fd5b506101bb73ffffffffffffffffffffffffffffffffffffffff60043516610212565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610276565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461023757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a7230582017c3a6b57ac8bbc7fd5320cba2543fe371eb8726332cd58700252790e09cf2140029`

// DeployDelegator deploys a new Ethereum contract, binding an instance of Delegator to it.
func DeployDelegator(auth *bind.TransactOpts, backend bind.ContractBackend, _controller common.Address, _controllerLookupName [32]byte) (common.Address, *types.Transaction, *Delegator, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DelegatorBin), backend, _controller, _controllerLookupName)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Delegator{DelegatorCaller: DelegatorCaller{contract: contract}, DelegatorTransactor: DelegatorTransactor{contract: contract}, DelegatorFilterer: DelegatorFilterer{contract: contract}}, nil
}

// Delegator is an auto generated Go binding around an Ethereum contract.
type Delegator struct {
	DelegatorCaller     // Read-only binding to the contract
	DelegatorTransactor // Write-only binding to the contract
	DelegatorFilterer   // Log filterer for contract events
}

// DelegatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegatorSession struct {
	Contract     *Delegator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegatorCallerSession struct {
	Contract *DelegatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DelegatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegatorTransactorSession struct {
	Contract     *DelegatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DelegatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegatorRaw struct {
	Contract *Delegator // Generic contract binding to access the raw methods on
}

// DelegatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegatorCallerRaw struct {
	Contract *DelegatorCaller // Generic read-only contract binding to access the raw methods on
}

// DelegatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegatorTransactorRaw struct {
	Contract *DelegatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegator creates a new instance of Delegator, bound to a specific deployed contract.
func NewDelegator(address common.Address, backend bind.ContractBackend) (*Delegator, error) {
	contract, err := bindDelegator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Delegator{DelegatorCaller: DelegatorCaller{contract: contract}, DelegatorTransactor: DelegatorTransactor{contract: contract}, DelegatorFilterer: DelegatorFilterer{contract: contract}}, nil
}

// NewDelegatorCaller creates a new read-only instance of Delegator, bound to a specific deployed contract.
func NewDelegatorCaller(address common.Address, caller bind.ContractCaller) (*DelegatorCaller, error) {
	contract, err := bindDelegator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorCaller{contract: contract}, nil
}

// NewDelegatorTransactor creates a new write-only instance of Delegator, bound to a specific deployed contract.
func NewDelegatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegatorTransactor, error) {
	contract, err := bindDelegator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorTransactor{contract: contract}, nil
}

// NewDelegatorFilterer creates a new log filterer instance of Delegator, bound to a specific deployed contract.
func NewDelegatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegatorFilterer, error) {
	contract, err := bindDelegator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegatorFilterer{contract: contract}, nil
}

// bindDelegator binds a generic wrapper to an already deployed contract.
func bindDelegator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegator *DelegatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Delegator.Contract.DelegatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegator *DelegatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.Contract.DelegatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegator *DelegatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegator.Contract.DelegatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegator *DelegatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Delegator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegator *DelegatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegator *DelegatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegator.Contract.contract.Transact(opts, method, params...)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Delegator *DelegatorCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Delegator.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Delegator *DelegatorSession) ControllerLookupName() ([32]byte, error) {
	return _Delegator.Contract.ControllerLookupName(&_Delegator.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Delegator *DelegatorCallerSession) ControllerLookupName() ([32]byte, error) {
	return _Delegator.Contract.ControllerLookupName(&_Delegator.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Delegator *DelegatorCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Delegator.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Delegator *DelegatorSession) GetController() (common.Address, error) {
	return _Delegator.Contract.GetController(&_Delegator.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Delegator *DelegatorCallerSession) GetController() (common.Address, error) {
	return _Delegator.Contract.GetController(&_Delegator.CallOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Delegator *DelegatorTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Delegator *DelegatorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.SetController(&_Delegator.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Delegator *DelegatorTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.SetController(&_Delegator.TransactOpts, _controller)
}
