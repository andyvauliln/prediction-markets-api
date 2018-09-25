package augur-contract-interfaces

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)
// MapABI is the input ABI used to generate the binding from.
const MapABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"contains\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getAsAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getValueOrZero\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getAsAddressOrZero\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MapBin is the compiled bytecode used for deploying new contracts.
const MapBin = `0x60806040526002805460008054600160a060020a03191633908117909155600160a860020a03199091161790556105d38061003b6000396000f3006080604052600436106100e55763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631d1a696d81146100ea5780633018205f1461011657806361641bdc1461014757806379f57e681461016b5780638196b8c714610183578063893d20e8146101ad5780638eaa6ac0146101c257806392eefe9b146101da57806395bc2673146101fb578063a87d942c14610213578063bef72fa214610228578063c4d66de81461023d578063d1de592a1461025e578063d8e5e64e14610279578063ee89dab414610291578063f2fde38b146102a6575b600080fd5b3480156100f657600080fd5b506101026004356102c7565b604080519115158252519081900360200190f35b34801561012257600080fd5b5061012b6102de565b60408051600160a060020a039092168252519081900360200190f35b34801561015357600080fd5b50610102600435600160a060020a03602435166102ed565b34801561017757600080fd5b5061012b600435610323565b34801561018f57600080fd5b5061019b60043561032e565b60408051918252519081900360200190f35b3480156101b957600080fd5b5061012b610340565b3480156101ce57600080fd5b5061019b60043561034f565b3480156101e657600080fd5b50610102600160a060020a036004351661036a565b34801561020757600080fd5b506101026004356103b4565b34801561021f57600080fd5b5061019b610404565b34801561023457600080fd5b5061019b61040a565b34801561024957600080fd5b50610102600160a060020a0360043516610410565b34801561026a57600080fd5b50610102600435602435610476565b34801561028557600080fd5b5061012b6004356104c6565b34801561029d57600080fd5b506101026104d1565b3480156102b257600080fd5b50610102600160a060020a03600435166104f2565b60008181526003602052604090205415155b919050565b600054600160a060020a031690565b600254600090600160a060020a0316331461030757600080fd5b61031a83600160a060020a038416610476565b90505b92915050565b600061031d8261034f565b60009081526003602052604090205490565b600254600160a060020a031690565b60008181526003602052604081205480151561031d57600080fd5b60008054600160a060020a0316331461038257600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600254600090600160a060020a031633146103ce57600080fd5b6103d7826102c7565b15156103e5575060006102d9565b5060009081526003602052604081205560048054600019019055600190565b60045490565b60015481565b60025460009074010000000000000000000000000000000000000000900460ff161561043b57600080fd5b610443610539565b505060028054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600254600090600160a060020a0316331461049057600080fd5b610499836102c7565b156104a65750600061031d565b506000918252600360205260409091205560048054600190810190915590565b600061031d8261032e565b60025474010000000000000000000000000000000000000000900460ff1690565b600254600090600160a060020a0316331461050c57600080fd5b600160a060020a038216156105315760025461044390600160a060020a03168361059f565b506001919050565b60025460009074010000000000000000000000000000000000000000900460ff161561056457600080fd5b506002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055600190565b6001929150505600a165627a7a72305820b1f713d81b9a3360edb3384acd4fd644d4c298bfbf1bbd55a51d185a51c0337c0029`

// DeployMap deploys a new Ethereum contract, binding an instance of Map to it.
func DeployMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Map, error) {
	parsed, err := abi.JSON(strings.NewReader(MapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Map{MapCaller: MapCaller{contract: contract}, MapTransactor: MapTransactor{contract: contract}, MapFilterer: MapFilterer{contract: contract}}, nil
}

// Map is an auto generated Go binding around an Ethereum contract.
type Map struct {
	MapCaller     // Read-only binding to the contract
	MapTransactor // Write-only binding to the contract
	MapFilterer   // Log filterer for contract events
}

// MapCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapSession struct {
	Contract     *Map              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapCallerSession struct {
	Contract *MapCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapTransactorSession struct {
	Contract     *MapTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapRaw struct {
	Contract *Map // Generic contract binding to access the raw methods on
}

// MapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapCallerRaw struct {
	Contract *MapCaller // Generic read-only contract binding to access the raw methods on
}

// MapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapTransactorRaw struct {
	Contract *MapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMap creates a new instance of Map, bound to a specific deployed contract.
func NewMap(address common.Address, backend bind.ContractBackend) (*Map, error) {
	contract, err := bindMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Map{MapCaller: MapCaller{contract: contract}, MapTransactor: MapTransactor{contract: contract}, MapFilterer: MapFilterer{contract: contract}}, nil
}

// NewMapCaller creates a new read-only instance of Map, bound to a specific deployed contract.
func NewMapCaller(address common.Address, caller bind.ContractCaller) (*MapCaller, error) {
	contract, err := bindMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapCaller{contract: contract}, nil
}

// NewMapTransactor creates a new write-only instance of Map, bound to a specific deployed contract.
func NewMapTransactor(address common.Address, transactor bind.ContractTransactor) (*MapTransactor, error) {
	contract, err := bindMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapTransactor{contract: contract}, nil
}

// NewMapFilterer creates a new log filterer instance of Map, bound to a specific deployed contract.
func NewMapFilterer(address common.Address, filterer bind.ContractFilterer) (*MapFilterer, error) {
	contract, err := bindMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapFilterer{contract: contract}, nil
}

// bindMap binds a generic wrapper to an already deployed contract.
func bindMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Map *MapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Map.Contract.MapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Map *MapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Map.Contract.MapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Map *MapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Map.Contract.MapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Map *MapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Map.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Map *MapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Map.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Map *MapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Map.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0x1d1a696d.
//
// Solidity: function contains(_key bytes32) constant returns(bool)
func (_Map *MapCaller) Contains(opts *bind.CallOpts, _key [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "contains", _key)
	return *ret0, err
}

// Contains is a free data retrieval call binding the contract method 0x1d1a696d.
//
// Solidity: function contains(_key bytes32) constant returns(bool)
func (_Map *MapSession) Contains(_key [32]byte) (bool, error) {
	return _Map.Contract.Contains(&_Map.CallOpts, _key)
}

// Contains is a free data retrieval call binding the contract method 0x1d1a696d.
//
// Solidity: function contains(_key bytes32) constant returns(bool)
func (_Map *MapCallerSession) Contains(_key [32]byte) (bool, error) {
	return _Map.Contract.Contains(&_Map.CallOpts, _key)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Map *MapCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Map *MapSession) ControllerLookupName() ([32]byte, error) {
	return _Map.Contract.ControllerLookupName(&_Map.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Map *MapCallerSession) ControllerLookupName() ([32]byte, error) {
	return _Map.Contract.ControllerLookupName(&_Map.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(_key bytes32) constant returns(bytes32)
func (_Map *MapCaller) Get(opts *bind.CallOpts, _key [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "get", _key)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(_key bytes32) constant returns(bytes32)
func (_Map *MapSession) Get(_key [32]byte) ([32]byte, error) {
	return _Map.Contract.Get(&_Map.CallOpts, _key)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(_key bytes32) constant returns(bytes32)
func (_Map *MapCallerSession) Get(_key [32]byte) ([32]byte, error) {
	return _Map.Contract.Get(&_Map.CallOpts, _key)
}

// GetAsAddress is a free data retrieval call binding the contract method 0x79f57e68.
//
// Solidity: function getAsAddress(_key bytes32) constant returns(address)
func (_Map *MapCaller) GetAsAddress(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "getAsAddress", _key)
	return *ret0, err
}

// GetAsAddress is a free data retrieval call binding the contract method 0x79f57e68.
//
// Solidity: function getAsAddress(_key bytes32) constant returns(address)
func (_Map *MapSession) GetAsAddress(_key [32]byte) (common.Address, error) {
	return _Map.Contract.GetAsAddress(&_Map.CallOpts, _key)
}

// GetAsAddress is a free data retrieval call binding the contract method 0x79f57e68.
//
// Solidity: function getAsAddress(_key bytes32) constant returns(address)
func (_Map *MapCallerSession) GetAsAddress(_key [32]byte) (common.Address, error) {
	return _Map.Contract.GetAsAddress(&_Map.CallOpts, _key)
}

// GetAsAddressOrZero is a free data retrieval call binding the contract method 0xd8e5e64e.
//
// Solidity: function getAsAddressOrZero(_key bytes32) constant returns(address)
func (_Map *MapCaller) GetAsAddressOrZero(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "getAsAddressOrZero", _key)
	return *ret0, err
}

// GetAsAddressOrZero is a free data retrieval call binding the contract method 0xd8e5e64e.
//
// Solidity: function getAsAddressOrZero(_key bytes32) constant returns(address)
func (_Map *MapSession) GetAsAddressOrZero(_key [32]byte) (common.Address, error) {
	return _Map.Contract.GetAsAddressOrZero(&_Map.CallOpts, _key)
}

// GetAsAddressOrZero is a free data retrieval call binding the contract method 0xd8e5e64e.
//
// Solidity: function getAsAddressOrZero(_key bytes32) constant returns(address)
func (_Map *MapCallerSession) GetAsAddressOrZero(_key [32]byte) (common.Address, error) {
	return _Map.Contract.GetAsAddressOrZero(&_Map.CallOpts, _key)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Map *MapCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Map *MapSession) GetController() (common.Address, error) {
	return _Map.Contract.GetController(&_Map.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Map *MapCallerSession) GetController() (common.Address, error) {
	return _Map.Contract.GetController(&_Map.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() constant returns(uint256)
func (_Map *MapCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "getCount")
	return *ret0, err
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() constant returns(uint256)
func (_Map *MapSession) GetCount() (*big.Int, error) {
	return _Map.Contract.GetCount(&_Map.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() constant returns(uint256)
func (_Map *MapCallerSession) GetCount() (*big.Int, error) {
	return _Map.Contract.GetCount(&_Map.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_Map *MapCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_Map *MapSession) GetInitialized() (bool, error) {
	return _Map.Contract.GetInitialized(&_Map.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_Map *MapCallerSession) GetInitialized() (bool, error) {
	return _Map.Contract.GetInitialized(&_Map.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Map *MapCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Map *MapSession) GetOwner() (common.Address, error) {
	return _Map.Contract.GetOwner(&_Map.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Map *MapCallerSession) GetOwner() (common.Address, error) {
	return _Map.Contract.GetOwner(&_Map.CallOpts)
}

// GetValueOrZero is a free data retrieval call binding the contract method 0x8196b8c7.
//
// Solidity: function getValueOrZero(_key bytes32) constant returns(bytes32)
func (_Map *MapCaller) GetValueOrZero(opts *bind.CallOpts, _key [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Map.contract.Call(opts, out, "getValueOrZero", _key)
	return *ret0, err
}

// GetValueOrZero is a free data retrieval call binding the contract method 0x8196b8c7.
//
// Solidity: function getValueOrZero(_key bytes32) constant returns(bytes32)
func (_Map *MapSession) GetValueOrZero(_key [32]byte) ([32]byte, error) {
	return _Map.Contract.GetValueOrZero(&_Map.CallOpts, _key)
}

// GetValueOrZero is a free data retrieval call binding the contract method 0x8196b8c7.
//
// Solidity: function getValueOrZero(_key bytes32) constant returns(bytes32)
func (_Map *MapCallerSession) GetValueOrZero(_key [32]byte) ([32]byte, error) {
	return _Map.Contract.GetValueOrZero(&_Map.CallOpts, _key)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(_key bytes32, _value bytes32) returns(bool)
func (_Map *MapTransactor) Add(opts *bind.TransactOpts, _key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Map.contract.Transact(opts, "add", _key, _value)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(_key bytes32, _value bytes32) returns(bool)
func (_Map *MapSession) Add(_key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Map.Contract.Add(&_Map.TransactOpts, _key, _value)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(_key bytes32, _value bytes32) returns(bool)
func (_Map *MapTransactorSession) Add(_key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Map.Contract.Add(&_Map.TransactOpts, _key, _value)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_owner address) returns(bool)
func (_Map *MapTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Map.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_owner address) returns(bool)
func (_Map *MapSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Map.Contract.Initialize(&_Map.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_owner address) returns(bool)
func (_Map *MapTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Map.Contract.Initialize(&_Map.TransactOpts, _owner)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(_key bytes32) returns(bool)
func (_Map *MapTransactor) Remove(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _Map.contract.Transact(opts, "remove", _key)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(_key bytes32) returns(bool)
func (_Map *MapSession) Remove(_key [32]byte) (*types.Transaction, error) {
	return _Map.Contract.Remove(&_Map.TransactOpts, _key)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(_key bytes32) returns(bool)
func (_Map *MapTransactorSession) Remove(_key [32]byte) (*types.Transaction, error) {
	return _Map.Contract.Remove(&_Map.TransactOpts, _key)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Map *MapTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Map.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Map *MapSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Map.Contract.SetController(&_Map.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Map *MapTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Map.Contract.SetController(&_Map.TransactOpts, _controller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Map *MapTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Map.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Map *MapSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Map.Contract.TransferOwnership(&_Map.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Map *MapTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Map.Contract.TransferOwnership(&_Map.TransactOpts, _newOwner)
}

// MapFactoryABI is the input ABI used to generate the binding from.
const MapFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MapFactoryBin is the compiled bytecode used for deploying new contracts.
const MapFactoryBin = `0x608060405234801561001057600080fd5b50610529806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663b782f6138114610045575b600080fd5b34801561005157600080fd5b5061007973ffffffffffffffffffffffffffffffffffffffff600435811690602435166100a2565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000806000846100b06101e5565b73ffffffffffffffffffffffffffffffffffffffff90911681527f4d617000000000000000000000000000000000000000000000000000000000006020820152604080519182900301906000f08015801561010f573d6000803e3d6000fd5b5091508190508073ffffffffffffffffffffffffffffffffffffffff1663c4d66de8856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156101b057600080fd5b505af11580156101c4573d6000803e3d6000fd5b505050506040513d60208110156101da57600080fd5b509095945050505050565b604051610308806101f6833901905600608060405234801561001057600080fd5b506040516040806103088339810160405280516020909101516000805433600160a060020a03199182161716600160a060020a039093169290921782556001556102a890819061006090396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461014f57806392eefe9b1461018d578063bef72fa2146101cf575b60015460009015156100675761014c565b60008054600154604080517ff39ec1f700000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff9092169263f39ec1f7926024808401936020939083900390910190829087803b1580156100df57600080fd5b505af11580156100f3573d6000803e3d6000fd5b505050506040513d602081101561010957600080fd5b50516040805136601f8101601f191680830190935292935091600083376000803684865af4808015610147576040513d81016040523d6000823e3d81f35b600080fd5b50005b34801561015b57600080fd5b506101646101f6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019957600080fd5b506101bb73ffffffffffffffffffffffffffffffffffffffff60043516610212565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610276565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461023757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a7230582017c3a6b57ac8bbc7fd5320cba2543fe371eb8726332cd58700252790e09cf2140029a165627a7a723058203c11567a28352985143e09fa8123e3a282cf9773ffa150197f019151b77e0b510029`

// DeployMapFactory deploys a new Ethereum contract, binding an instance of MapFactory to it.
func DeployMapFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MapFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(MapFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MapFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MapFactory{MapFactoryCaller: MapFactoryCaller{contract: contract}, MapFactoryTransactor: MapFactoryTransactor{contract: contract}, MapFactoryFilterer: MapFactoryFilterer{contract: contract}}, nil
}

// MapFactory is an auto generated Go binding around an Ethereum contract.
type MapFactory struct {
	MapFactoryCaller     // Read-only binding to the contract
	MapFactoryTransactor // Write-only binding to the contract
	MapFactoryFilterer   // Log filterer for contract events
}

// MapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapFactorySession struct {
	Contract     *MapFactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapFactoryCallerSession struct {
	Contract *MapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapFactoryTransactorSession struct {
	Contract     *MapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapFactoryRaw struct {
	Contract *MapFactory // Generic contract binding to access the raw methods on
}

// MapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapFactoryCallerRaw struct {
	Contract *MapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapFactoryTransactorRaw struct {
	Contract *MapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMapFactory creates a new instance of MapFactory, bound to a specific deployed contract.
func NewMapFactory(address common.Address, backend bind.ContractBackend) (*MapFactory, error) {
	contract, err := bindMapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MapFactory{MapFactoryCaller: MapFactoryCaller{contract: contract}, MapFactoryTransactor: MapFactoryTransactor{contract: contract}, MapFactoryFilterer: MapFactoryFilterer{contract: contract}}, nil
}

// NewMapFactoryCaller creates a new read-only instance of MapFactory, bound to a specific deployed contract.
func NewMapFactoryCaller(address common.Address, caller bind.ContractCaller) (*MapFactoryCaller, error) {
	contract, err := bindMapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapFactoryCaller{contract: contract}, nil
}

// NewMapFactoryTransactor creates a new write-only instance of MapFactory, bound to a specific deployed contract.
func NewMapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MapFactoryTransactor, error) {
	contract, err := bindMapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapFactoryTransactor{contract: contract}, nil
}

// NewMapFactoryFilterer creates a new log filterer instance of MapFactory, bound to a specific deployed contract.
func NewMapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MapFactoryFilterer, error) {
	contract, err := bindMapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapFactoryFilterer{contract: contract}, nil
}

// bindMapFactory binds a generic wrapper to an already deployed contract.
func bindMapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MapFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapFactory *MapFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MapFactory.Contract.MapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapFactory *MapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapFactory.Contract.MapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapFactory *MapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapFactory.Contract.MapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapFactory *MapFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapFactory *MapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapFactory *MapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateMap is a paid mutator transaction binding the contract method 0xb782f613.
//
// Solidity: function createMap(_controller address, _owner address) returns(address)
func (_MapFactory *MapFactoryTransactor) CreateMap(opts *bind.TransactOpts, _controller common.Address, _owner common.Address) (*types.Transaction, error) {
	return _MapFactory.contract.Transact(opts, "createMap", _controller, _owner)
}

// CreateMap is a paid mutator transaction binding the contract method 0xb782f613.
//
// Solidity: function createMap(_controller address, _owner address) returns(address)
func (_MapFactory *MapFactorySession) CreateMap(_controller common.Address, _owner common.Address) (*types.Transaction, error) {
	return _MapFactory.Contract.CreateMap(&_MapFactory.TransactOpts, _controller, _owner)
}

// CreateMap is a paid mutator transaction binding the contract method 0xb782f613.
//
// Solidity: function createMap(_controller address, _owner address) returns(address)
func (_MapFactory *MapFactoryTransactorSession) CreateMap(_controller common.Address, _owner common.Address) (*types.Transaction, error) {
	return _MapFactory.Contract.CreateMap(&_MapFactory.TransactOpts, _controller, _owner)
}
