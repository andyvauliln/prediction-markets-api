package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IMailboxABI is the input ABI used to generate the binding from.
const IMailboxABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IMailboxBin is the compiled bytecode used for deploying new contracts.
const IMailboxBin = `0x`

// DeployIMailbox deploys a new Ethereum contract, binding an instance of IMailbox to it.
func DeployIMailbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IMailbox, error) {
	parsed, err := abi.JSON(strings.NewReader(IMailboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IMailboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IMailbox{IMailboxCaller: IMailboxCaller{contract: contract}, IMailboxTransactor: IMailboxTransactor{contract: contract}, IMailboxFilterer: IMailboxFilterer{contract: contract}}, nil
}

// IMailbox is an auto generated Go binding around an Ethereum contract.
type IMailbox struct {
	IMailboxCaller     // Read-only binding to the contract
	IMailboxTransactor // Write-only binding to the contract
	IMailboxFilterer   // Log filterer for contract events
}

// IMailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMailboxSession struct {
	Contract     *IMailbox         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMailboxCallerSession struct {
	Contract *IMailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IMailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMailboxTransactorSession struct {
	Contract     *IMailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IMailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMailboxRaw struct {
	Contract *IMailbox // Generic contract binding to access the raw methods on
}

// IMailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMailboxCallerRaw struct {
	Contract *IMailboxCaller // Generic read-only contract binding to access the raw methods on
}

// IMailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMailboxTransactorRaw struct {
	Contract *IMailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMailbox creates a new instance of IMailbox, bound to a specific deployed contract.
func NewIMailbox(address common.Address, backend bind.ContractBackend) (*IMailbox, error) {
	contract, err := bindIMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMailbox{IMailboxCaller: IMailboxCaller{contract: contract}, IMailboxTransactor: IMailboxTransactor{contract: contract}, IMailboxFilterer: IMailboxFilterer{contract: contract}}, nil
}

// NewIMailboxCaller creates a new read-only instance of IMailbox, bound to a specific deployed contract.
func NewIMailboxCaller(address common.Address, caller bind.ContractCaller) (*IMailboxCaller, error) {
	contract, err := bindIMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMailboxCaller{contract: contract}, nil
}

// NewIMailboxTransactor creates a new write-only instance of IMailbox, bound to a specific deployed contract.
func NewIMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IMailboxTransactor, error) {
	contract, err := bindIMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMailboxTransactor{contract: contract}, nil
}

// NewIMailboxFilterer creates a new log filterer instance of IMailbox, bound to a specific deployed contract.
func NewIMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IMailboxFilterer, error) {
	contract, err := bindIMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMailboxFilterer{contract: contract}, nil
}

// bindIMailbox binds a generic wrapper to an already deployed contract.
func bindIMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMailboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMailbox *IMailboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMailbox.Contract.IMailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMailbox *IMailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMailbox.Contract.IMailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMailbox *IMailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMailbox.Contract.IMailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMailbox *IMailboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMailbox *IMailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMailbox *IMailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMailbox.Contract.contract.Transact(opts, method, params...)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns(bool)
func (_IMailbox *IMailboxTransactor) DepositEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMailbox.contract.Transact(opts, "depositEther")
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns(bool)
func (_IMailbox *IMailboxSession) DepositEther() (*types.Transaction, error) {
	return _IMailbox.Contract.DepositEther(&_IMailbox.TransactOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns(bool)
func (_IMailbox *IMailboxTransactorSession) DepositEther() (*types.Transaction, error) {
	return _IMailbox.Contract.DepositEther(&_IMailbox.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_owner address, _market address) returns(bool)
func (_IMailbox *IMailboxTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _market common.Address) (*types.Transaction, error) {
	return _IMailbox.contract.Transact(opts, "initialize", _owner, _market)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_owner address, _market address) returns(bool)
func (_IMailbox *IMailboxSession) Initialize(_owner common.Address, _market common.Address) (*types.Transaction, error) {
	return _IMailbox.Contract.Initialize(&_IMailbox.TransactOpts, _owner, _market)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_owner address, _market address) returns(bool)
func (_IMailbox *IMailboxTransactorSession) Initialize(_owner common.Address, _market common.Address) (*types.Transaction, error) {
	return _IMailbox.Contract.Initialize(&_IMailbox.TransactOpts, _owner, _market)
}

// MailboxABI is the input ABI used to generate the binding from.
const MailboxABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MailboxBin is the compiled bytecode used for deploying new contracts.
const MailboxBin = `0x60806040526002805460008054600160a060020a03191633908117909155600160a860020a03199091161790556109328061003b6000396000f30060806040526004361061008a5763ffffffff60e060020a6000350416633018205f811461008f578063485cc955146100c057806349df728c146100fb5780637362377b1461011c578063893d20e81461013157806392eefe9b1461014657806398ea5fca14610167578063bef72fa21461016f578063ee89dab414610196578063f2fde38b146101ab575b600080fd5b34801561009b57600080fd5b506100a46101cc565b60408051600160a060020a039092168252519081900360200190f35b3480156100cc57600080fd5b506100e7600160a060020a03600435811690602435166101db565b604080519115158252519081900360200190f35b34801561010757600080fd5b506100e7600160a060020a036004351661024e565b34801561012857600080fd5b506100e76103a9565b34801561013d57600080fd5b506100a4610601565b34801561015257600080fd5b506100e7600160a060020a0360043516610610565b6100e761065a565b34801561017b57600080fd5b5061018461065f565b60408051918252519081900360200190f35b3480156101a257600080fd5b506100e7610665565b3480156101b757600080fd5b506100e7600160a060020a0360043516610686565b600054600160a060020a031690565b60025460009074010000000000000000000000000000000000000000900460ff161561020657600080fd5b61020e6106f7565b505060028054600160a060020a0393841673ffffffffffffffffffffffffffffffffffffffff199182161790915560038054929093169116179055600190565b6002546000908190600160a060020a0316331461026a57600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038516916370a082319160248083019260209291908290030181600087803b1580156102cb57600080fd5b505af11580156102df573d6000803e3d6000fd5b505050506040513d60208110156102f557600080fd5b5051600254604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201526024810184905290519293509085169163a9059cbb916044808201926020929091908290030181600087803b15801561036957600080fd5b505af115801561037d573d6000803e3d6000fd5b505050506040513d602081101561039357600080fd5b505115156103a057600080fd5b50600192915050565b60025460009081908190600160a060020a031633146103c757600080fd5b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f436173680000000000000000000000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b15801561044e57600080fd5b505af1158015610462573d6000803e3d6000fd5b505050506040513d602081101561047857600080fd5b5051604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051919350600160a060020a038416916370a08231916024808201926020929091908290030181600087803b1580156104df57600080fd5b505af11580156104f3573d6000803e3d6000fd5b505050506040513d602081101561050957600080fd5b5051905060008111156105b257600254604080517f1baffe38000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015260248101849052905191841691631baffe38916044808201926020929091908290030181600087803b15801561058557600080fd5b505af1158015610599573d6000803e3d6000fd5b505050506040513d60208110156105af57600080fd5b50505b6000303111156105f857600254604051600160a060020a0390911690303180156108fc02916000818181858888f193505050501580156105f6573d6000803e3d6000fd5b505b60019250505090565b600254600160a060020a031690565b60008054600160a060020a0316331461062857600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600190565b60015481565b60025474010000000000000000000000000000000000000000900460ff1690565b600254600090600160a060020a031633146106a057600080fd5b600160a060020a038216156106ef576002546106c590600160a060020a03168361075d565b506002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790555b506001919050565b60025460009074010000000000000000000000000000000000000000900460ff161561072257600080fd5b506002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055600190565b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156107b157600080fd5b505af11580156107c5573d6000803e3d6000fd5b505050506040513d60208110156107db57600080fd5b5051600354604080517f870c426d0000000000000000000000000000000000000000000000000000000081529051600160a060020a039384169363339594f993169163870c426d9160048083019260209291908290030181600087803b15801561084457600080fd5b505af1158015610858573d6000803e3d6000fd5b505050506040513d602081101561086e57600080fd5b50516003546040805160e060020a63ffffffff8616028152600160a060020a039384166004820152918316602483015287831660448301529186166064820152905160848083019260209291908290030181600087803b1580156108d157600080fd5b505af11580156108e5573d6000803e3d6000fd5b505050506040513d60208110156108fb57600080fd5b5060019493505050505600a165627a7a723058206ed302f76d846971cb9893c1f05eea83bd12fece0365881626806f0575e1ba5a0029`

// DeployMailbox deploys a new Ethereum contract, binding an instance of Mailbox to it.
func DeployMailbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mailbox, error) {
	parsed, err := abi.JSON(strings.NewReader(MailboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MailboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mailbox{MailboxCaller: MailboxCaller{contract: contract}, MailboxTransactor: MailboxTransactor{contract: contract}, MailboxFilterer: MailboxFilterer{contract: contract}}, nil
}

// Mailbox is an auto generated Go binding around an Ethereum contract.
type Mailbox struct {
	MailboxCaller     // Read-only binding to the contract
	MailboxTransactor // Write-only binding to the contract
	MailboxFilterer   // Log filterer for contract events
}

// MailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MailboxSession struct {
	Contract     *Mailbox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MailboxCallerSession struct {
	Contract *MailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MailboxTransactorSession struct {
	Contract     *MailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MailboxRaw struct {
	Contract *Mailbox // Generic contract binding to access the raw methods on
}

// MailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MailboxCallerRaw struct {
	Contract *MailboxCaller // Generic read-only contract binding to access the raw methods on
}

// MailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MailboxTransactorRaw struct {
	Contract *MailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMailbox creates a new instance of Mailbox, bound to a specific deployed contract.
func NewMailbox(address common.Address, backend bind.ContractBackend) (*Mailbox, error) {
	contract, err := bindMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mailbox{MailboxCaller: MailboxCaller{contract: contract}, MailboxTransactor: MailboxTransactor{contract: contract}, MailboxFilterer: MailboxFilterer{contract: contract}}, nil
}

// NewMailboxCaller creates a new read-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxCaller(address common.Address, caller bind.ContractCaller) (*MailboxCaller, error) {
	contract, err := bindMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxCaller{contract: contract}, nil
}

// NewMailboxTransactor creates a new write-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*MailboxTransactor, error) {
	contract, err := bindMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxTransactor{contract: contract}, nil
}

// NewMailboxFilterer creates a new log filterer instance of Mailbox, bound to a specific deployed contract.
func NewMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*MailboxFilterer, error) {
	contract, err := bindMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MailboxFilterer{contract: contract}, nil
}

// bindMailbox binds a generic wrapper to an already deployed contract.
func bindMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MailboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.MailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transact(opts, method, params...)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Mailbox *MailboxCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mailbox.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Mailbox *MailboxSession) ControllerLookupName() ([32]byte, error) {
	return _Mailbox.Contract.ControllerLookupName(&_Mailbox.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Mailbox *MailboxCallerSession) ControllerLookupName() ([32]byte, error) {
	return _Mailbox.Contract.ControllerLookupName(&_Mailbox.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Mailbox *MailboxCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mailbox.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Mailbox *MailboxSession) GetController() (common.Address, error) {
	return _Mailbox.Contract.GetController(&_Mailbox.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Mailbox *MailboxCallerSession) GetController() (common.Address, error) {
	return _Mailbox.Contract.GetController(&_Mailbox.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_Mailbox *MailboxCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mailbox.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_Mailbox *MailboxSession) GetInitialized() (bool, error) {
	return _Mailbox.Contract.GetInitialized(&_Mailbox.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_Mailbox *MailboxCallerSession) GetInitialized() (bool, error) {
	return _Mailbox.Contract.GetInitialized(&_Mailbox.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Mailbox *MailboxCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mailbox.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Mailbox *MailboxSession) GetOwner() (common.Address, error) {
	return _Mailbox.Contract.GetOwner(&_Mailbox.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Mailbox *MailboxCallerSession) GetOwner() (common.Address, error) {
	return _Mailbox.Contract.GetOwner(&_Mailbox.CallOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns(bool)
func (_Mailbox *MailboxTransactor) DepositEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "depositEther")
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns(bool)
func (_Mailbox *MailboxSession) DepositEther() (*types.Transaction, error) {
	return _Mailbox.Contract.DepositEther(&_Mailbox.TransactOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns(bool)
func (_Mailbox *MailboxTransactorSession) DepositEther() (*types.Transaction, error) {
	return _Mailbox.Contract.DepositEther(&_Mailbox.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_owner address, _market address) returns(bool)
func (_Mailbox *MailboxTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _market common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "initialize", _owner, _market)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_owner address, _market address) returns(bool)
func (_Mailbox *MailboxSession) Initialize(_owner common.Address, _market common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Initialize(&_Mailbox.TransactOpts, _owner, _market)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(_owner address, _market address) returns(bool)
func (_Mailbox *MailboxTransactorSession) Initialize(_owner common.Address, _market common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Initialize(&_Mailbox.TransactOpts, _owner, _market)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Mailbox *MailboxTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Mailbox *MailboxSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetController(&_Mailbox.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Mailbox *MailboxTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetController(&_Mailbox.TransactOpts, _controller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Mailbox *MailboxTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Mailbox *MailboxSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.TransferOwnership(&_Mailbox.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Mailbox *MailboxTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.TransferOwnership(&_Mailbox.TransactOpts, _newOwner)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns(bool)
func (_Mailbox *MailboxTransactor) WithdrawEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "withdrawEther")
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns(bool)
func (_Mailbox *MailboxSession) WithdrawEther() (*types.Transaction, error) {
	return _Mailbox.Contract.WithdrawEther(&_Mailbox.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns(bool)
func (_Mailbox *MailboxTransactorSession) WithdrawEther() (*types.Transaction, error) {
	return _Mailbox.Contract.WithdrawEther(&_Mailbox.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(_token address) returns(bool)
func (_Mailbox *MailboxTransactor) WithdrawTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "withdrawTokens", _token)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(_token address) returns(bool)
func (_Mailbox *MailboxSession) WithdrawTokens(_token common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.WithdrawTokens(&_Mailbox.TransactOpts, _token)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(_token address) returns(bool)
func (_Mailbox *MailboxTransactorSession) WithdrawTokens(_token common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.WithdrawTokens(&_Mailbox.TransactOpts, _token)
}

// MailboxFactoryABI is the input ABI used to generate the binding from.
const MailboxFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"createMailbox\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MailboxFactoryBin is the compiled bytecode used for deploying new contracts.
const MailboxFactoryBin = `0x608060405234801561001057600080fd5b50610510806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663a73b24b48114610045575b600080fd5b34801561005157600080fd5b5061007f73ffffffffffffffffffffffffffffffffffffffff600435811690602435811690604435166100a8565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000806000856100b66101cc565b73ffffffffffffffffffffffffffffffffffffffff90911681527f4d61696c626f78000000000000000000000000000000000000000000000000006020820152604080519182900301906000f080158015610115573d6000803e3d6000fd5b50604080517f485cc95500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152878116602483015291519294508493509083169163485cc955916044808201926020929091908290030181600087803b15801561019657600080fd5b505af11580156101aa573d6000803e3d6000fd5b505050506040513d60208110156101c057600080fd5b50909695505050505050565b604051610308806101dd833901905600608060405234801561001057600080fd5b506040516040806103088339810160405280516020909101516000805433600160a060020a03199182161716600160a060020a039093169290921782556001556102a890819061006090396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461014f57806392eefe9b1461018d578063bef72fa2146101cf575b60015460009015156100675761014c565b60008054600154604080517ff39ec1f700000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff9092169263f39ec1f7926024808401936020939083900390910190829087803b1580156100df57600080fd5b505af11580156100f3573d6000803e3d6000fd5b505050506040513d602081101561010957600080fd5b50516040805136601f8101601f191680830190935292935091600083376000803684865af4808015610147576040513d81016040523d6000823e3d81f35b600080fd5b50005b34801561015b57600080fd5b506101646101f6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019957600080fd5b506101bb73ffffffffffffffffffffffffffffffffffffffff60043516610212565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610276565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461023757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a7230582017c3a6b57ac8bbc7fd5320cba2543fe371eb8726332cd58700252790e09cf2140029a165627a7a72305820ee49130120a34c2e71c6fd314ce7b347b2676c7fa0bb4f5d3f83033eaf2c7c480029`

// DeployMailboxFactory deploys a new Ethereum contract, binding an instance of MailboxFactory to it.
func DeployMailboxFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MailboxFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(MailboxFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MailboxFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MailboxFactory{MailboxFactoryCaller: MailboxFactoryCaller{contract: contract}, MailboxFactoryTransactor: MailboxFactoryTransactor{contract: contract}, MailboxFactoryFilterer: MailboxFactoryFilterer{contract: contract}}, nil
}

// MailboxFactory is an auto generated Go binding around an Ethereum contract.
type MailboxFactory struct {
	MailboxFactoryCaller     // Read-only binding to the contract
	MailboxFactoryTransactor // Write-only binding to the contract
	MailboxFactoryFilterer   // Log filterer for contract events
}

// MailboxFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MailboxFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MailboxFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MailboxFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MailboxFactorySession struct {
	Contract     *MailboxFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailboxFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MailboxFactoryCallerSession struct {
	Contract *MailboxFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MailboxFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MailboxFactoryTransactorSession struct {
	Contract     *MailboxFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MailboxFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MailboxFactoryRaw struct {
	Contract *MailboxFactory // Generic contract binding to access the raw methods on
}

// MailboxFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MailboxFactoryCallerRaw struct {
	Contract *MailboxFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MailboxFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MailboxFactoryTransactorRaw struct {
	Contract *MailboxFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMailboxFactory creates a new instance of MailboxFactory, bound to a specific deployed contract.
func NewMailboxFactory(address common.Address, backend bind.ContractBackend) (*MailboxFactory, error) {
	contract, err := bindMailboxFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MailboxFactory{MailboxFactoryCaller: MailboxFactoryCaller{contract: contract}, MailboxFactoryTransactor: MailboxFactoryTransactor{contract: contract}, MailboxFactoryFilterer: MailboxFactoryFilterer{contract: contract}}, nil
}

// NewMailboxFactoryCaller creates a new read-only instance of MailboxFactory, bound to a specific deployed contract.
func NewMailboxFactoryCaller(address common.Address, caller bind.ContractCaller) (*MailboxFactoryCaller, error) {
	contract, err := bindMailboxFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxFactoryCaller{contract: contract}, nil
}

// NewMailboxFactoryTransactor creates a new write-only instance of MailboxFactory, bound to a specific deployed contract.
func NewMailboxFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MailboxFactoryTransactor, error) {
	contract, err := bindMailboxFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxFactoryTransactor{contract: contract}, nil
}

// NewMailboxFactoryFilterer creates a new log filterer instance of MailboxFactory, bound to a specific deployed contract.
func NewMailboxFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MailboxFactoryFilterer, error) {
	contract, err := bindMailboxFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MailboxFactoryFilterer{contract: contract}, nil
}

// bindMailboxFactory binds a generic wrapper to an already deployed contract.
func bindMailboxFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MailboxFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MailboxFactory *MailboxFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MailboxFactory.Contract.MailboxFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MailboxFactory *MailboxFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MailboxFactory.Contract.MailboxFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MailboxFactory *MailboxFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MailboxFactory.Contract.MailboxFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MailboxFactory *MailboxFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MailboxFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MailboxFactory *MailboxFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MailboxFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MailboxFactory *MailboxFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MailboxFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateMailbox is a paid mutator transaction binding the contract method 0xa73b24b4.
//
// Solidity: function createMailbox(_controller address, _owner address, _market address) returns(address)
func (_MailboxFactory *MailboxFactoryTransactor) CreateMailbox(opts *bind.TransactOpts, _controller common.Address, _owner common.Address, _market common.Address) (*types.Transaction, error) {
	return _MailboxFactory.contract.Transact(opts, "createMailbox", _controller, _owner, _market)
}

// CreateMailbox is a paid mutator transaction binding the contract method 0xa73b24b4.
//
// Solidity: function createMailbox(_controller address, _owner address, _market address) returns(address)
func (_MailboxFactory *MailboxFactorySession) CreateMailbox(_controller common.Address, _owner common.Address, _market common.Address) (*types.Transaction, error) {
	return _MailboxFactory.Contract.CreateMailbox(&_MailboxFactory.TransactOpts, _controller, _owner, _market)
}

// CreateMailbox is a paid mutator transaction binding the contract method 0xa73b24b4.
//
// Solidity: function createMailbox(_controller address, _owner address, _market address) returns(address)
func (_MailboxFactory *MailboxFactoryTransactorSession) CreateMailbox(_controller common.Address, _owner common.Address, _market common.Address) (*types.Transaction, error) {
	return _MailboxFactory.Contract.CreateMailbox(&_MailboxFactory.TransactOpts, _controller, _owner, _market)
}
