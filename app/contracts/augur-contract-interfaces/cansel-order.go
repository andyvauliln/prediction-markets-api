package augur-contract-interfaces

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CancelOrderABI is the input ABI used to generate the binding from.
const CancelOrderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CancelOrderBin is the compiled bytecode used for deploying new contracts.
const CancelOrderBin = `0x608060405260008054600160a860020a031916331790556112ad806100256000396000f30060806040526004361061003d5763ffffffff60e060020a6000350416633018205f81146100425780637489ec231461007357806392eefe9b1461009f575b600080fd5b34801561004e57600080fd5b506100576100c0565b60408051600160a060020a039092168252519081900360200190f35b34801561007f57600080fd5b5061008b6004356100cf565b604080519115158252519081900360200190f35b3480156100ab57600080fd5b5061008b600160a060020a03600435166109c1565b600054600160a060020a031690565b60008060008060008060008060149054906101000a900460ff161515156100f557600080fd5b6000805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055610132610a0b565b5087151561013f57600080fd5b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4f7264657273000000000000000000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b1580156101c657600080fd5b505af11580156101da573d6000803e3d6000fd5b505050506040513d60208110156101f057600080fd5b5051604080517f4a1a342b000000000000000000000000000000000000000000000000000000008152600481018b90529051919750600160a060020a03881691634a1a342b916024808201926020929091908290030181600087803b15801561025857600080fd5b505af115801561026c573d6000803e3d6000fd5b505050506040513d602081101561028257600080fd5b5051604080517febead05f000000000000000000000000000000000000000000000000000000008152600481018b90529051919650600160a060020a0388169163ebead05f916024808201926020929091908290030181600087803b1580156102ea57600080fd5b505af11580156102fe573d6000803e3d6000fd5b505050506040513d602081101561031457600080fd5b5051604080517fcf357364000000000000000000000000000000000000000000000000000000008152600481018b90529051919550600160a060020a0388169163cf357364916024808201926020929091908290030181600087803b15801561037c57600080fd5b505af1158015610390573d6000803e3d6000fd5b505050506040513d60208110156103a657600080fd5b5051604080517fc3c95c7b000000000000000000000000000000000000000000000000000000008152600481018b90529051919450600160a060020a0388169163c3c95c7b916024808201926020929091908290030181600087803b15801561040e57600080fd5b505af1158015610422573d6000803e3d6000fd5b505050506040513d602081101561043857600080fd5b5051604080517f5d1a3b82000000000000000000000000000000000000000000000000000000008152600481018b90529051919350600160a060020a03881691635d1a3b82916024808201926020929091908290030181600087803b1580156104a057600080fd5b505af11580156104b4573d6000803e3d6000fd5b505050506040513d60208110156104ca57600080fd5b5051604080517fe7d80c70000000000000000000000000000000000000000000000000000000008152600481018b90529051919250600160a060020a0388169163e7d80c70916024808201926020929091908290030181600087803b15801561053257600080fd5b505af1158015610546573d6000803e3d6000fd5b505050506040513d602081101561055c57600080fd5b5051600160a060020a0316331461057257600080fd5b604080517ffde99668000000000000000000000000000000000000000000000000000000008152600481018a90529051600160a060020a0388169163fde996689160248083019260209291908290030181600087803b1580156105d457600080fd5b505af11580156105e8573d6000803e3d6000fd5b505050506040513d60208110156105fe57600080fd5b5061060f9050338486888686610b5f565b5085600160a060020a03166395ede03283876040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561067357600080fd5b505af1158015610687573d6000803e3d6000fd5b505050506040513d602081101561069d57600080fd5b5050604080517fa0695f240000000000000000000000000000000000000000000000000000000081529051600160a060020a0384169163a0695f249160048083019260209291908290030181600087803b1580156106fa57600080fd5b505af115801561070e573d6000803e3d6000fd5b505050506040513d602081101561072457600080fd5b505060008054604080517f4e94c8290000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921692634e94c829926004808401936020939083900390910190829087803b15801561078757600080fd5b505af115801561079b573d6000803e3d6000fd5b505050506040513d60208110156107b157600080fd5b5051604080517f870c426d0000000000000000000000000000000000000000000000000000000081529051600160a060020a0392831692636e1636bb929086169163870c426d916004808201926020929091908290030181600087803b15801561081a57600080fd5b505af115801561082e573d6000803e3d6000fd5b505050506040513d602081101561084457600080fd5b5051604080517f65957bf5000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a038716916365957bf59160248083019260209291908290030181600087803b1580156108a857600080fd5b505af11580156108bc573d6000803e3d6000fd5b505050506040513d60208110156108d257600080fd5b505160405160e060020a63ffffffff8516028152600160a060020a038084166004830190815290831660248301523360448301819052606483018f9052918e918a918d918d9160840184600181111561092757fe5b60ff168152602001838152602001828152602001975050505050505050602060405180830381600087803b15801561095e57600080fd5b505af1158015610972573d6000803e3d6000fd5b505050506040513d602081101561098857600080fd5b506001975050610996610f69565b50506000805474ff000000000000000000000000000000000000000019169055509395945050505050565b60008054600160a060020a031633146109d957600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600080341115610b595760008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f436173680000000000000000000000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b158015610a9c57600080fd5b505af1158015610ab0573d6000803e3d6000fd5b505050506040513d6020811015610ac657600080fd5b5051604080517f4faa8a260000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691634faa8a26913491602480830192602092919082900301818588803b158015610b2b57600080fd5b505af1158015610b3f573d6000803e3d6000fd5b50505050506040513d6020811015610b5657600080fd5b50505b50600190565b600080600080871115610e39576000886001811115610b7a57fe5b1415610d2357600091505b84600160a060020a03166327ce5b8c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610bc357600080fd5b505af1158015610bd7573d6000803e3d6000fd5b505050506040513d6020811015610bed57600080fd5b5051821015610d1e57818414610d135784600160a060020a03166365957bf5836040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b158015610c4657600080fd5b505af1158015610c5a573d6000803e3d6000fd5b505050506040513d6020811015610c7057600080fd5b5051604080517f7b30074d000000000000000000000000000000000000000000000000000000008152600160a060020a0388811660048301528c81166024830152604482018b905291519190921691637b30074d9160648083019260209291908290030181600087803b158015610ce657600080fd5b505af1158015610cfa573d6000803e3d6000fd5b505050506040513d6020811015610d1057600080fd5b50505b816001019150610b85565b610e39565b84600160a060020a03166365957bf5856040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b158015610d6c57600080fd5b505af1158015610d80573d6000803e3d6000fd5b505050506040513d6020811015610d9657600080fd5b5051604080517f7b30074d000000000000000000000000000000000000000000000000000000008152600160a060020a0388811660048301528c81166024830152604482018b905291519190921691637b30074d9160648083019260209291908290030181600087803b158015610e0c57600080fd5b505af1158015610e20573d6000803e3d6000fd5b505050506040513d6020811015610e3657600080fd5b50505b6000861115610f5a5784600160a060020a031663df2a29da6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610e8057600080fd5b505af1158015610e94573d6000803e3d6000fd5b505050506040513d6020811015610eaa57600080fd5b5051604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0388811660048301528c81166024830152604482018a90529151929350908316916323b872dd916064808201926020929091908290030181600087803b158015610f2357600080fd5b505af1158015610f37573d6000803e3d6000fd5b505050506040513d6020811015610f4d57600080fd5b50511515610f5a57600080fd5b50600198975050505050505050565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f436173680000000000000000000000000000000000000000000000000000000060048201529051839283928392600160a060020a039092169163f39ec1f79160248082019260209290919082900301818787803b158015610ff557600080fd5b505af1158015611009573d6000803e3d6000fd5b505050506040513d602081101561101f57600080fd5b5051604080517f70a082310000000000000000000000000000000000000000000000000000000081523360048201529051919450600160a060020a038516916370a08231916024808201926020929091908290030181600087803b15801561108657600080fd5b505af115801561109a573d6000803e3d6000fd5b505050506040513d60208110156110b057600080fd5b505191506000821115611277576000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561110f57600080fd5b505af1158015611123573d6000803e3d6000fd5b505050506040513d602081101561113957600080fd5b5051604080517fec238994000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301523360248301523060448301526064820186905291519293509083169163ec238994916084808201926020929091908290030181600087803b1580156111b657600080fd5b505af11580156111ca573d6000803e3d6000fd5b505050506040513d60208110156111e057600080fd5b5050604080517f1baffe38000000000000000000000000000000000000000000000000000000008152336004820152602481018490529051600160a060020a03851691631baffe389160448083019260209291908290030181600087803b15801561124a57600080fd5b505af115801561125e573d6000803e3d6000fd5b505050506040513d602081101561127457600080fd5b50505b60019350505050905600a165627a7a72305820006b24bb0c4a8a70cbd56e7c7f6ab6fa7113801258827c307a0435ab8d89cdd50029`

// DeployCancelOrder deploys a new Ethereum contract, binding an instance of CancelOrder to it.
func DeployCancelOrder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CancelOrder, error) {
	parsed, err := abi.JSON(strings.NewReader(CancelOrderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CancelOrderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CancelOrder{CancelOrderCaller: CancelOrderCaller{contract: contract}, CancelOrderTransactor: CancelOrderTransactor{contract: contract}, CancelOrderFilterer: CancelOrderFilterer{contract: contract}}, nil
}

// CancelOrder is an auto generated Go binding around an Ethereum contract.
type CancelOrder struct {
	CancelOrderCaller     // Read-only binding to the contract
	CancelOrderTransactor // Write-only binding to the contract
	CancelOrderFilterer   // Log filterer for contract events
}

// CancelOrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type CancelOrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancelOrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CancelOrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancelOrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CancelOrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CancelOrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CancelOrderSession struct {
	Contract     *CancelOrder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CancelOrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CancelOrderCallerSession struct {
	Contract *CancelOrderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CancelOrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CancelOrderTransactorSession struct {
	Contract     *CancelOrderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CancelOrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type CancelOrderRaw struct {
	Contract *CancelOrder // Generic contract binding to access the raw methods on
}

// CancelOrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CancelOrderCallerRaw struct {
	Contract *CancelOrderCaller // Generic read-only contract binding to access the raw methods on
}

// CancelOrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CancelOrderTransactorRaw struct {
	Contract *CancelOrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCancelOrder creates a new instance of CancelOrder, bound to a specific deployed contract.
func NewCancelOrder(address common.Address, backend bind.ContractBackend) (*CancelOrder, error) {
	contract, err := bindCancelOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CancelOrder{CancelOrderCaller: CancelOrderCaller{contract: contract}, CancelOrderTransactor: CancelOrderTransactor{contract: contract}, CancelOrderFilterer: CancelOrderFilterer{contract: contract}}, nil
}

// NewCancelOrderCaller creates a new read-only instance of CancelOrder, bound to a specific deployed contract.
func NewCancelOrderCaller(address common.Address, caller bind.ContractCaller) (*CancelOrderCaller, error) {
	contract, err := bindCancelOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CancelOrderCaller{contract: contract}, nil
}

// NewCancelOrderTransactor creates a new write-only instance of CancelOrder, bound to a specific deployed contract.
func NewCancelOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*CancelOrderTransactor, error) {
	contract, err := bindCancelOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CancelOrderTransactor{contract: contract}, nil
}

// NewCancelOrderFilterer creates a new log filterer instance of CancelOrder, bound to a specific deployed contract.
func NewCancelOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*CancelOrderFilterer, error) {
	contract, err := bindCancelOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CancelOrderFilterer{contract: contract}, nil
}

// bindCancelOrder binds a generic wrapper to an already deployed contract.
func bindCancelOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CancelOrderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CancelOrder *CancelOrderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CancelOrder.Contract.CancelOrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CancelOrder *CancelOrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancelOrder.Contract.CancelOrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CancelOrder *CancelOrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CancelOrder.Contract.CancelOrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CancelOrder *CancelOrderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CancelOrder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CancelOrder *CancelOrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CancelOrder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CancelOrder *CancelOrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CancelOrder.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_CancelOrder *CancelOrderCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CancelOrder.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_CancelOrder *CancelOrderSession) GetController() (common.Address, error) {
	return _CancelOrder.Contract.GetController(&_CancelOrder.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_CancelOrder *CancelOrderCallerSession) GetController() (common.Address, error) {
	return _CancelOrder.Contract.GetController(&_CancelOrder.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(_orderId bytes32) returns(bool)
func (_CancelOrder *CancelOrderTransactor) CancelOrder(opts *bind.TransactOpts, _orderId [32]byte) (*types.Transaction, error) {
	return _CancelOrder.contract.Transact(opts, "cancelOrder", _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(_orderId bytes32) returns(bool)
func (_CancelOrder *CancelOrderSession) CancelOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _CancelOrder.Contract.CancelOrder(&_CancelOrder.TransactOpts, _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(_orderId bytes32) returns(bool)
func (_CancelOrder *CancelOrderTransactorSession) CancelOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _CancelOrder.Contract.CancelOrder(&_CancelOrder.TransactOpts, _orderId)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_CancelOrder *CancelOrderTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _CancelOrder.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_CancelOrder *CancelOrderSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _CancelOrder.Contract.SetController(&_CancelOrder.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_CancelOrder *CancelOrderTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _CancelOrder.Contract.SetController(&_CancelOrder.TransactOpts, _controller)
}

// ICancelOrderABI is the input ABI used to generate the binding from.
const ICancelOrderABI = "[]"

// ICancelOrderBin is the compiled bytecode used for deploying new contracts.
const ICancelOrderBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a723058200a124e4405a38b45d23b2e888b8bc38f8a1cd3b139fb7dd3f66ab8194d35a2110029`

// DeployICancelOrder deploys a new Ethereum contract, binding an instance of ICancelOrder to it.
func DeployICancelOrder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ICancelOrder, error) {
	parsed, err := abi.JSON(strings.NewReader(ICancelOrderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ICancelOrderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ICancelOrder{ICancelOrderCaller: ICancelOrderCaller{contract: contract}, ICancelOrderTransactor: ICancelOrderTransactor{contract: contract}, ICancelOrderFilterer: ICancelOrderFilterer{contract: contract}}, nil
}

// ICancelOrder is an auto generated Go binding around an Ethereum contract.
type ICancelOrder struct {
	ICancelOrderCaller     // Read-only binding to the contract
	ICancelOrderTransactor // Write-only binding to the contract
	ICancelOrderFilterer   // Log filterer for contract events
}

// ICancelOrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICancelOrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICancelOrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICancelOrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICancelOrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICancelOrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICancelOrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICancelOrderSession struct {
	Contract     *ICancelOrder     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICancelOrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICancelOrderCallerSession struct {
	Contract *ICancelOrderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ICancelOrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICancelOrderTransactorSession struct {
	Contract     *ICancelOrderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ICancelOrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICancelOrderRaw struct {
	Contract *ICancelOrder // Generic contract binding to access the raw methods on
}

// ICancelOrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICancelOrderCallerRaw struct {
	Contract *ICancelOrderCaller // Generic read-only contract binding to access the raw methods on
}

// ICancelOrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICancelOrderTransactorRaw struct {
	Contract *ICancelOrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICancelOrder creates a new instance of ICancelOrder, bound to a specific deployed contract.
func NewICancelOrder(address common.Address, backend bind.ContractBackend) (*ICancelOrder, error) {
	contract, err := bindICancelOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICancelOrder{ICancelOrderCaller: ICancelOrderCaller{contract: contract}, ICancelOrderTransactor: ICancelOrderTransactor{contract: contract}, ICancelOrderFilterer: ICancelOrderFilterer{contract: contract}}, nil
}

// NewICancelOrderCaller creates a new read-only instance of ICancelOrder, bound to a specific deployed contract.
func NewICancelOrderCaller(address common.Address, caller bind.ContractCaller) (*ICancelOrderCaller, error) {
	contract, err := bindICancelOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICancelOrderCaller{contract: contract}, nil
}

// NewICancelOrderTransactor creates a new write-only instance of ICancelOrder, bound to a specific deployed contract.
func NewICancelOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*ICancelOrderTransactor, error) {
	contract, err := bindICancelOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICancelOrderTransactor{contract: contract}, nil
}

// NewICancelOrderFilterer creates a new log filterer instance of ICancelOrder, bound to a specific deployed contract.
func NewICancelOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*ICancelOrderFilterer, error) {
	contract, err := bindICancelOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICancelOrderFilterer{contract: contract}, nil
}

// bindICancelOrder binds a generic wrapper to an already deployed contract.
func bindICancelOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICancelOrderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICancelOrder *ICancelOrderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICancelOrder.Contract.ICancelOrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICancelOrder *ICancelOrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICancelOrder.Contract.ICancelOrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICancelOrder *ICancelOrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICancelOrder.Contract.ICancelOrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICancelOrder *ICancelOrderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICancelOrder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICancelOrder *ICancelOrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICancelOrder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICancelOrder *ICancelOrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICancelOrder.Contract.contract.Transact(opts, method, params...)
}
