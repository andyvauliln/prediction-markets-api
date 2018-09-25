package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ControllerABI is the input ABI used to generate the binding from.
const ControllerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getContractDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes20\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_commitHash\",\"type\":\"bytes20\"},{\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"registerContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"assertIsWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAugur\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"registry\",\"outputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"commitHash\",\"type\":\"bytes20\"},{\"name\":\"bytecodeHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"lookup\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ControllerBin is the compiled bytecode used for deploying new contracts.
const ControllerBin = `0x60806040526003805460ff1916905534801561001a57600080fd5b5060008054600160a060020a0319163390811782558152600160208190526040909120805460ff191690911790556106b0806100576000396000f3006080604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663185e996981146100c9578063188ec356146101175780632d6d85231461013e5780633f08882f1461018b5780634e94c829146101ac57806375f12b21146101dd5780637ef50298146101f25780638ab1d681146102465780638da5cb5b146102675780639b19251a1461027c578063e43252d71461029d578063f2fde38b146102be578063f39ec1f7146102df575b600080fd5b3480156100d557600080fd5b506100e16004356102f7565b60408051600160a060020a0390941684526bffffffffffffffffffffffff19909216602084015282820152519081900360600190f35b34801561012357600080fd5b5061012c610338565b60408051918252519081900360200190f35b34801561014a57600080fd5b50610177600435600160a060020a03602435166bffffffffffffffffffffffff19604435166064356103ea565b604080519115158252519081900360200190f35b34801561019757600080fd5b50610177600160a060020a03600435166104d1565b3480156101b857600080fd5b506101c1610500565b60408051600160a060020a039092168252519081900360200190f35b3480156101e957600080fd5b50610177610530565b3480156101fe57600080fd5b5061020a600435610539565b60408051948552600160a060020a0390931660208501526bffffffffffffffffffffffff19909116838301526060830152519081900360800190f35b34801561025257600080fd5b50610177600160a060020a036004351661057a565b34801561027357600080fd5b506101c16105b8565b34801561028857600080fd5b50610177600160a060020a03600435166105c7565b3480156102a957600080fd5b50610177600160a060020a03600435166105dc565b3480156102ca57600080fd5b50610177600160a060020a036004351661061c565b3480156102eb57600080fd5b506101c1600435610666565b6000908152600260208190526040909120600181015491810154600390910154600160a060020a03909216926c010000000000000000000000009091029190565b60006103637f54696d6500000000000000000000000000000000000000000000000000000000610666565b600160a060020a031663188ec3566040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156103b957600080fd5b505af11580156103cd573d6000803e3d6000fd5b505050506040513d60208110156103e357600080fd5b5051905090565b60008054600160a060020a0316331461040257600080fd5b600085815260026020526040902060010154600160a060020a03161561042757600080fd5b5060408051608081018252858152600160a060020a0380861660208084019182526bffffffffffffffffffffffff1987168486019081526060850187815260008b815260029384905296909620945185559151600180860180549290951673ffffffffffffffffffffffffffffffffffffffff1992831617909455915190840180546c01000000000000000000000000909204919092161790559151600390910155949350505050565b600160a060020a03811660009081526001602052604081205460ff1615156104f857600080fd5b506001919050565b600061052b7f4175677572000000000000000000000000000000000000000000000000000000610666565b905090565b60035460ff1681565b600260208190526000918252604090912080546001820154928201546003909201549092600160a060020a0316916c01000000000000000000000000029084565b60008054600160a060020a0316331461059257600080fd5b50600160a060020a03166000908152600160208190526040909120805460ff1916905590565b600054600160a060020a031681565b60016020526000908152604090205460ff1681565b60008054600160a060020a031633146105f457600080fd5b50600160a060020a03166000908152600160208190526040909120805460ff19168217905590565b60008054600160a060020a0316331461063457600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600090815260026020526040902060010154600160a060020a0316905600a165627a7a72305820ae70d5e95cd435aa3771fcd1f5d0bb6591dbd4b509c0894604cb62befea9443c0029`

// DeployController deploys a new Ethereum contract, binding an instance of Controller to it.
func DeployController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Controller, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// AssertIsWhitelisted is a free data retrieval call binding the contract method 0x3f08882f.
//
// Solidity: function assertIsWhitelisted(_target address) constant returns(bool)
func (_Controller *ControllerCaller) AssertIsWhitelisted(opts *bind.CallOpts, _target common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "assertIsWhitelisted", _target)
	return *ret0, err
}

// AssertIsWhitelisted is a free data retrieval call binding the contract method 0x3f08882f.
//
// Solidity: function assertIsWhitelisted(_target address) constant returns(bool)
func (_Controller *ControllerSession) AssertIsWhitelisted(_target common.Address) (bool, error) {
	return _Controller.Contract.AssertIsWhitelisted(&_Controller.CallOpts, _target)
}

// AssertIsWhitelisted is a free data retrieval call binding the contract method 0x3f08882f.
//
// Solidity: function assertIsWhitelisted(_target address) constant returns(bool)
func (_Controller *ControllerCallerSession) AssertIsWhitelisted(_target common.Address) (bool, error) {
	return _Controller.Contract.AssertIsWhitelisted(&_Controller.CallOpts, _target)
}

// GetAugur is a free data retrieval call binding the contract method 0x4e94c829.
//
// Solidity: function getAugur() constant returns(address)
func (_Controller *ControllerCaller) GetAugur(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "getAugur")
	return *ret0, err
}

// GetAugur is a free data retrieval call binding the contract method 0x4e94c829.
//
// Solidity: function getAugur() constant returns(address)
func (_Controller *ControllerSession) GetAugur() (common.Address, error) {
	return _Controller.Contract.GetAugur(&_Controller.CallOpts)
}

// GetAugur is a free data retrieval call binding the contract method 0x4e94c829.
//
// Solidity: function getAugur() constant returns(address)
func (_Controller *ControllerCallerSession) GetAugur() (common.Address, error) {
	return _Controller.Contract.GetAugur(&_Controller.CallOpts)
}

// GetContractDetails is a free data retrieval call binding the contract method 0x185e9969.
//
// Solidity: function getContractDetails(_key bytes32) constant returns(address, bytes20, bytes32)
func (_Controller *ControllerCaller) GetContractDetails(opts *bind.CallOpts, _key [32]byte) (common.Address, [20]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new([20]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Controller.contract.Call(opts, out, "getContractDetails", _key)
	return *ret0, *ret1, *ret2, err
}

// GetContractDetails is a free data retrieval call binding the contract method 0x185e9969.
//
// Solidity: function getContractDetails(_key bytes32) constant returns(address, bytes20, bytes32)
func (_Controller *ControllerSession) GetContractDetails(_key [32]byte) (common.Address, [20]byte, [32]byte, error) {
	return _Controller.Contract.GetContractDetails(&_Controller.CallOpts, _key)
}

// GetContractDetails is a free data retrieval call binding the contract method 0x185e9969.
//
// Solidity: function getContractDetails(_key bytes32) constant returns(address, bytes20, bytes32)
func (_Controller *ControllerCallerSession) GetContractDetails(_key [32]byte) (common.Address, [20]byte, [32]byte, error) {
	return _Controller.Contract.GetContractDetails(&_Controller.CallOpts, _key)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_Controller *ControllerCaller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "getTimestamp")
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_Controller *ControllerSession) GetTimestamp() (*big.Int, error) {
	return _Controller.Contract.GetTimestamp(&_Controller.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_Controller *ControllerCallerSession) GetTimestamp() (*big.Int, error) {
	return _Controller.Contract.GetTimestamp(&_Controller.CallOpts)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(_key bytes32) constant returns(address)
func (_Controller *ControllerCaller) Lookup(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "lookup", _key)
	return *ret0, err
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(_key bytes32) constant returns(address)
func (_Controller *ControllerSession) Lookup(_key [32]byte) (common.Address, error) {
	return _Controller.Contract.Lookup(&_Controller.CallOpts, _key)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(_key bytes32) constant returns(address)
func (_Controller *ControllerCallerSession) Lookup(_key [32]byte) (common.Address, error) {
	return _Controller.Contract.Lookup(&_Controller.CallOpts, _key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7ef50298.
//
// Solidity: function registry( bytes32) constant returns(name bytes32, contractAddress address, commitHash bytes20, bytecodeHash bytes32)
func (_Controller *ControllerCaller) Registry(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Name            [32]byte
	ContractAddress common.Address
	CommitHash      [20]byte
	BytecodeHash    [32]byte
}, error) {
	ret := new(struct {
		Name            [32]byte
		ContractAddress common.Address
		CommitHash      [20]byte
		BytecodeHash    [32]byte
	})
	out := ret
	err := _Controller.contract.Call(opts, out, "registry", arg0)
	return *ret, err
}

// Registry is a free data retrieval call binding the contract method 0x7ef50298.
//
// Solidity: function registry( bytes32) constant returns(name bytes32, contractAddress address, commitHash bytes20, bytecodeHash bytes32)
func (_Controller *ControllerSession) Registry(arg0 [32]byte) (struct {
	Name            [32]byte
	ContractAddress common.Address
	CommitHash      [20]byte
	BytecodeHash    [32]byte
}, error) {
	return _Controller.Contract.Registry(&_Controller.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x7ef50298.
//
// Solidity: function registry( bytes32) constant returns(name bytes32, contractAddress address, commitHash bytes20, bytecodeHash bytes32)
func (_Controller *ControllerCallerSession) Registry(arg0 [32]byte) (struct {
	Name            [32]byte
	ContractAddress common.Address
	CommitHash      [20]byte
	BytecodeHash    [32]byte
}, error) {
	return _Controller.Contract.Registry(&_Controller.CallOpts, arg0)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Controller *ControllerCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Controller *ControllerSession) Stopped() (bool, error) {
	return _Controller.Contract.Stopped(&_Controller.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Controller *ControllerCallerSession) Stopped() (bool, error) {
	return _Controller.Contract.Stopped(&_Controller.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Controller *ControllerCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Controller.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Controller *ControllerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Controller.Contract.Whitelist(&_Controller.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Controller *ControllerCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Controller.Contract.Whitelist(&_Controller.CallOpts, arg0)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_target address) returns(bool)
func (_Controller *ControllerTransactor) AddToWhitelist(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addToWhitelist", _target)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_target address) returns(bool)
func (_Controller *ControllerSession) AddToWhitelist(_target common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddToWhitelist(&_Controller.TransactOpts, _target)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_target address) returns(bool)
func (_Controller *ControllerTransactorSession) AddToWhitelist(_target common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddToWhitelist(&_Controller.TransactOpts, _target)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x2d6d8523.
//
// Solidity: function registerContract(_key bytes32, _address address, _commitHash bytes20, _bytecodeHash bytes32) returns(bool)
func (_Controller *ControllerTransactor) RegisterContract(opts *bind.TransactOpts, _key [32]byte, _address common.Address, _commitHash [20]byte, _bytecodeHash [32]byte) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "registerContract", _key, _address, _commitHash, _bytecodeHash)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x2d6d8523.
//
// Solidity: function registerContract(_key bytes32, _address address, _commitHash bytes20, _bytecodeHash bytes32) returns(bool)
func (_Controller *ControllerSession) RegisterContract(_key [32]byte, _address common.Address, _commitHash [20]byte, _bytecodeHash [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.RegisterContract(&_Controller.TransactOpts, _key, _address, _commitHash, _bytecodeHash)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x2d6d8523.
//
// Solidity: function registerContract(_key bytes32, _address address, _commitHash bytes20, _bytecodeHash bytes32) returns(bool)
func (_Controller *ControllerTransactorSession) RegisterContract(_key [32]byte, _address common.Address, _commitHash [20]byte, _bytecodeHash [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.RegisterContract(&_Controller.TransactOpts, _key, _address, _commitHash, _bytecodeHash)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_target address) returns(bool)
func (_Controller *ControllerTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "removeFromWhitelist", _target)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_target address) returns(bool)
func (_Controller *ControllerSession) RemoveFromWhitelist(_target common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveFromWhitelist(&_Controller.TransactOpts, _target)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_target address) returns(bool)
func (_Controller *ControllerTransactorSession) RemoveFromWhitelist(_target common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveFromWhitelist(&_Controller.TransactOpts, _target)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Controller *ControllerTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Controller *ControllerSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Controller *ControllerTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, _newOwner)
}

// IControllerABI is the input ABI used to generate the binding from.
const IControllerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"assertIsWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAugur\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"lookup\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IControllerBin is the compiled bytecode used for deploying new contracts.
const IControllerBin = `0x`

// DeployIController deploys a new Ethereum contract, binding an instance of IController to it.
func DeployIController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IController, error) {
	parsed, err := abi.JSON(strings.NewReader(IControllerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IController{IControllerCaller: IControllerCaller{contract: contract}, IControllerTransactor: IControllerTransactor{contract: contract}, IControllerFilterer: IControllerFilterer{contract: contract}}, nil
}

// IController is an auto generated Go binding around an Ethereum contract.
type IController struct {
	IControllerCaller     // Read-only binding to the contract
	IControllerTransactor // Write-only binding to the contract
	IControllerFilterer   // Log filterer for contract events
}

// IControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControllerSession struct {
	Contract     *IController      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControllerCallerSession struct {
	Contract *IControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControllerTransactorSession struct {
	Contract     *IControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IControllerRaw struct {
	Contract *IController // Generic contract binding to access the raw methods on
}

// IControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControllerCallerRaw struct {
	Contract *IControllerCaller // Generic read-only contract binding to access the raw methods on
}

// IControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControllerTransactorRaw struct {
	Contract *IControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIController creates a new instance of IController, bound to a specific deployed contract.
func NewIController(address common.Address, backend bind.ContractBackend) (*IController, error) {
	contract, err := bindIController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IController{IControllerCaller: IControllerCaller{contract: contract}, IControllerTransactor: IControllerTransactor{contract: contract}, IControllerFilterer: IControllerFilterer{contract: contract}}, nil
}

// NewIControllerCaller creates a new read-only instance of IController, bound to a specific deployed contract.
func NewIControllerCaller(address common.Address, caller bind.ContractCaller) (*IControllerCaller, error) {
	contract, err := bindIController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControllerCaller{contract: contract}, nil
}

// NewIControllerTransactor creates a new write-only instance of IController, bound to a specific deployed contract.
func NewIControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*IControllerTransactor, error) {
	contract, err := bindIController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControllerTransactor{contract: contract}, nil
}

// NewIControllerFilterer creates a new log filterer instance of IController, bound to a specific deployed contract.
func NewIControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*IControllerFilterer, error) {
	contract, err := bindIController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControllerFilterer{contract: contract}, nil
}

// bindIController binds a generic wrapper to an already deployed contract.
func bindIController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IController *IControllerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IController.Contract.IControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IController *IControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IController.Contract.IControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IController *IControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IController.Contract.IControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IController *IControllerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IController *IControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IController *IControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IController.Contract.contract.Transact(opts, method, params...)
}

// AssertIsWhitelisted is a free data retrieval call binding the contract method 0x3f08882f.
//
// Solidity: function assertIsWhitelisted(_target address) constant returns(bool)
func (_IController *IControllerCaller) AssertIsWhitelisted(opts *bind.CallOpts, _target common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IController.contract.Call(opts, out, "assertIsWhitelisted", _target)
	return *ret0, err
}

// AssertIsWhitelisted is a free data retrieval call binding the contract method 0x3f08882f.
//
// Solidity: function assertIsWhitelisted(_target address) constant returns(bool)
func (_IController *IControllerSession) AssertIsWhitelisted(_target common.Address) (bool, error) {
	return _IController.Contract.AssertIsWhitelisted(&_IController.CallOpts, _target)
}

// AssertIsWhitelisted is a free data retrieval call binding the contract method 0x3f08882f.
//
// Solidity: function assertIsWhitelisted(_target address) constant returns(bool)
func (_IController *IControllerCallerSession) AssertIsWhitelisted(_target common.Address) (bool, error) {
	return _IController.Contract.AssertIsWhitelisted(&_IController.CallOpts, _target)
}

// GetAugur is a free data retrieval call binding the contract method 0x4e94c829.
//
// Solidity: function getAugur() constant returns(address)
func (_IController *IControllerCaller) GetAugur(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IController.contract.Call(opts, out, "getAugur")
	return *ret0, err
}

// GetAugur is a free data retrieval call binding the contract method 0x4e94c829.
//
// Solidity: function getAugur() constant returns(address)
func (_IController *IControllerSession) GetAugur() (common.Address, error) {
	return _IController.Contract.GetAugur(&_IController.CallOpts)
}

// GetAugur is a free data retrieval call binding the contract method 0x4e94c829.
//
// Solidity: function getAugur() constant returns(address)
func (_IController *IControllerCallerSession) GetAugur() (common.Address, error) {
	return _IController.Contract.GetAugur(&_IController.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_IController *IControllerCaller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IController.contract.Call(opts, out, "getTimestamp")
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_IController *IControllerSession) GetTimestamp() (*big.Int, error) {
	return _IController.Contract.GetTimestamp(&_IController.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() constant returns(uint256)
func (_IController *IControllerCallerSession) GetTimestamp() (*big.Int, error) {
	return _IController.Contract.GetTimestamp(&_IController.CallOpts)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(_key bytes32) constant returns(address)
func (_IController *IControllerCaller) Lookup(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IController.contract.Call(opts, out, "lookup", _key)
	return *ret0, err
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(_key bytes32) constant returns(address)
func (_IController *IControllerSession) Lookup(_key [32]byte) (common.Address, error) {
	return _IController.Contract.Lookup(&_IController.CallOpts, _key)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(_key bytes32) constant returns(address)
func (_IController *IControllerCallerSession) Lookup(_key [32]byte) (common.Address, error) {
	return _IController.Contract.Lookup(&_IController.CallOpts, _key)
}
