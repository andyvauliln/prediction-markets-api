package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// OrderABI is the input ABI used to generate the binding from.
const OrderABI = "[]"

// OrderBin is the compiled bytecode used for deploying new contracts.
const OrderBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206db376a38e0a65dba8e9b8f7fa862806d61943963a58d931ed6fe37037a08f850029`

// DeployOrder deploys a new Ethereum contract, binding an instance of Order to it.
func DeployOrder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Order, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Order{OrderCaller: OrderCaller{contract: contract}, OrderTransactor: OrderTransactor{contract: contract}, OrderFilterer: OrderFilterer{contract: contract}}, nil
}

// Order is an auto generated Go binding around an Ethereum contract.
type Order struct {
	OrderCaller     // Read-only binding to the contract
	OrderTransactor // Write-only binding to the contract
	OrderFilterer   // Log filterer for contract events
}

// OrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderSession struct {
	Contract     *Order            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderCallerSession struct {
	Contract *OrderCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderTransactorSession struct {
	Contract     *OrderTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderRaw struct {
	Contract *Order // Generic contract binding to access the raw methods on
}

// OrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderCallerRaw struct {
	Contract *OrderCaller // Generic read-only contract binding to access the raw methods on
}

// OrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderTransactorRaw struct {
	Contract *OrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrder creates a new instance of Order, bound to a specific deployed contract.
func NewOrder(address common.Address, backend bind.ContractBackend) (*Order, error) {
	contract, err := bindOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Order{OrderCaller: OrderCaller{contract: contract}, OrderTransactor: OrderTransactor{contract: contract}, OrderFilterer: OrderFilterer{contract: contract}}, nil
}

// NewOrderCaller creates a new read-only instance of Order, bound to a specific deployed contract.
func NewOrderCaller(address common.Address, caller bind.ContractCaller) (*OrderCaller, error) {
	contract, err := bindOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderCaller{contract: contract}, nil
}

// NewOrderTransactor creates a new write-only instance of Order, bound to a specific deployed contract.
func NewOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderTransactor, error) {
	contract, err := bindOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderTransactor{contract: contract}, nil
}

// NewOrderFilterer creates a new log filterer instance of Order, bound to a specific deployed contract.
func NewOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderFilterer, error) {
	contract, err := bindOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderFilterer{contract: contract}, nil
}

// bindOrder binds a generic wrapper to an already deployed contract.
func bindOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Order *OrderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Order.Contract.OrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Order *OrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.Contract.OrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Order *OrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Order.Contract.OrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Order *OrderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Order.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Order *OrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Order *OrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Order.Contract.contract.Transact(opts, method, params...)
}



// OrdersFetcherABI is the input ABI used to generate the binding from.
const OrdersFetcherABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_highestOrderId\",\"type\":\"bytes32\"}],\"name\":\"descendOrderList\",\"outputs\":[{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_bestOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worstOrderId\",\"type\":\"bytes32\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"}],\"name\":\"findBoundingOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_lowestOrderId\",\"type\":\"bytes32\"}],\"name\":\"ascendOrderList\",\"outputs\":[{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OrdersFetcherBin is the compiled bytecode used for deploying new contracts.
const OrdersFetcherBin = `0x608060405260008054600160a060020a031916331790556114ca806100256000396000f3006080604052600436106100535763ffffffff60e060020a6000350416630e1424b581146100585780633018205f146100925780633b01cf3c146100c357806392eefe9b146100ed578063f8266a7a14610122575b600080fd5b34801561006457600080fd5b5061007960ff60043516602435604435610143565b6040805192835260208301919091528051918290030190f35b34801561009e57600080fd5b506100a76107a4565b60408051600160a060020a039092168252519081900360200190f35b3480156100cf57600080fd5b5061007960ff6004351660243560443560643560843560a4356107b3565b3480156100f957600080fd5b5061010e600160a060020a0360043516610d12565b604080519115158252519081900360200190f35b34801561012e57600080fd5b5061007960ff60043516602435604435610d5c565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4f7264657273000000000000000000000000000000000000000000000000000060048201529051849392839283928392600160a060020a03169163f39ec1f791602480830192602092919082900301818787803b1580156101cf57600080fd5b505af11580156101e3573d6000803e3d6000fd5b505050506040513d60208110156101f957600080fd5b50519150600088600181111561020b57fe5b1415610296576040805160008051602061145f8339815191528152600481018790529051600160a060020a038416916331d98b3f9160248083019260209291908290030181600087803b15801561026157600080fd5b505af1158015610275573d6000803e3d6000fd5b505050506040513d602081101561028b57600080fd5b50518711925061032b565b60018860018111156102a457fe5b141561032b576040805160008051602061145f8339815191528152600481018790529051600160a060020a038416916331d98b3f9160248083019260209291908290030181600087803b1580156102fa57600080fd5b505af115801561030e573d6000803e3d6000fd5b505050506040513d602081101561032457600080fd5b5051871092505b821561033b576000949350610799565b81600160a060020a031663d9b3d9fe8989886040518463ffffffff1660e060020a0281526004018084600181111561036f57fe5b60ff16815260200183815260200182600019166000191681526020019350505050602060405180830381600087803b1580156103aa57600080fd5b505af11580156103be573d6000803e3d6000fd5b505050506040513d60208110156103d457600080fd5b505190505b80801561046057506040805160008051602061147f8339815191528152600481018790529051600160a060020a03841691638925f9e99160248083019260209291908290030181600087803b15801561043157600080fd5b505af1158015610445573d6000803e3d6000fd5b505050506040513d602081101561045b57600080fd5b505115155b1561071a576040805160008051602061147f8339815191528152600481018790529051600160a060020a03841691638925f9e99160248083019260209291908290030181600087803b1580156104b557600080fd5b505af11580156104c9573d6000803e3d6000fd5b505050506040513d60208110156104df57600080fd5b50516040517fd9b3d9fe000000000000000000000000000000000000000000000000000000008152909450600160a060020a0383169063d9b3d9fe908a908a9088906004018084600181111561053157fe5b60ff16815260200183815260200182600019166000191681526020019350505050602060405180830381600087803b15801561056c57600080fd5b505af1158015610580573d6000803e3d6000fd5b505050506040513d602081101561059657600080fd5b50519050808061069157506040805160008051602061147f8339815191528152600481018790529051600160a060020a038416916331d98b3f918391638925f9e99160248083019260209291908290030181600087803b1580156105f957600080fd5b505af115801561060d573d6000803e3d6000fd5b505050506040513d602081101561062357600080fd5b50516040805160e060020a63ffffffff851602815260048101929092525160248083019260209291908290030181600087803b15801561066257600080fd5b505af1158015610676573d6000803e3d6000fd5b505050506040513d602081101561068c57600080fd5b505187145b15610715576040805160008051602061147f8339815191528152600481018790529051600160a060020a03841691638925f9e99160248083019260209291908290030181600087803b1580156106e657600080fd5b505af11580156106fa573d6000803e3d6000fd5b505050506040513d602081101561071057600080fd5b505194505b6103d9565b6040805160008051602061147f8339815191528152600481018790529051600160a060020a03841691638925f9e99160248083019260209291908290030181600087803b15801561076a57600080fd5b505af115801561077e573d6000803e3d6000fd5b505050506040513d602081101561079457600080fd5b505193505b505050935093915050565b600054600160a060020a031690565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4f726465727300000000000000000000000000000000000000000000000000006004820152905183928392600160a060020a039091169163f39ec1f79160248082019260209290919082900301818787803b15801561083d57600080fd5b505af1158015610851573d6000803e3d6000fd5b505050506040513d602081101561086757600080fd5b505190508686141561093e578615156108865760009250829150610d06565b80600160a060020a0316635cf17bbb8a8a8a6040518463ffffffff1660e060020a028152600401808460018111156108ba57fe5b60ff16815260200183815260200182600019166000191681526020019350505050602060405180830381600087803b1580156108f557600080fd5b505af1158015610909573d6000803e3d6000fd5b505050506040513d602081101561091f57600080fd5b5051156109325760009250869150610d06565b86925060009150610d06565b8415610a6b576040805160008051602061145f8339815191528152600481018790529051600160a060020a038316916331d98b3f9160248083019260209291908290030181600087803b15801561099457600080fd5b505af11580156109a8573d6000803e3d6000fd5b505050506040513d60208110156109be57600080fd5b505115156109cf5760009450610a6b565b80600160a060020a03166319e54fb38a8a886040518463ffffffff1660e060020a02815260040180846001811115610a0357fe5b60ff16815260200183815260200182600019166000191681526020019350505050602060405180830381600087803b158015610a3e57600080fd5b505af1158015610a52573d6000803e3d6000fd5b505050506040513d6020811015610a6857600080fd5b50505b8315610b98576040805160008051602061145f8339815191528152600481018690529051600160a060020a038316916331d98b3f9160248083019260209291908290030181600087803b158015610ac157600080fd5b505af1158015610ad5573d6000803e3d6000fd5b505050506040513d6020811015610aeb57600080fd5b50511515610afc5760009350610b98565b80600160a060020a03166345c92c388a8a876040518463ffffffff1660e060020a02815260040180846001811115610b3057fe5b60ff16815260200183815260200182600019166000191681526020019350505050602060405180830381600087803b158015610b6b57600080fd5b505af1158015610b7f573d6000803e3d6000fd5b505050506040513d6020811015610b9557600080fd5b50505b84158015610ba4575083155b15610bbd57610bb4898989610143565b92509250610d06565b841515610bcf57610bb4898986610d5c565b831515610be157610bb4898987610143565b6040805160008051602061147f83398151915281526004810187905290518591600160a060020a03841691638925f9e9916024808201926020929091908290030181600087803b158015610c3457600080fd5b505af1158015610c48573d6000803e3d6000fd5b505050506040513d6020811015610c5e57600080fd5b505114610c7057610bb4898987610143565b6040805160008051602061143f83398151915281526004810186905290518691600160a060020a038416916394d26cb5916024808201926020929091908290030181600087803b158015610cc357600080fd5b505af1158015610cd7573d6000803e3d6000fd5b505050506040513d6020811015610ced57600080fd5b505114610cff57610bb4898986610d5c565b8484925092505b50965096945050505050565b60008054600160a060020a03163314610d2a57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4f72646572730000000000000000000000000000000000000000000000000000600482015290518492849283928392600160a060020a03169163f39ec1f791602480830192602092919082900301818787803b158015610de757600080fd5b505af1158015610dfb573d6000803e3d6000fd5b505050506040513d6020811015610e1157600080fd5b505191506000886001811115610e2357fe5b1415610eaf576040805160008051602061145f8339815191528152600481018690529051600160a060020a038416916331d98b3f9160248083019260209291908290030181600087803b158015610e7957600080fd5b505af1158015610e8d573d6000803e3d6000fd5b505050506040513d6020811015610ea357600080fd5b50518711159250610f45565b6001886001811115610ebd57fe5b1415610f45576040805160008051602061145f8339815191528152600481018690529051600160a060020a038416916331d98b3f9160248083019260209291908290030181600087803b158015610f1357600080fd5b505af1158015610f27573d6000803e3d6000fd5b505050506040513d6020811015610f3d57600080fd5b505187101592505b8215610fd4576040805160008051602061147f83398151915281526004810186905290518591600160a060020a03851691638925f9e9916024808201926020929091908290030181600087803b158015610f9e57600080fd5b505af1158015610fb2573d6000803e3d6000fd5b505050506040513d6020811015610fc857600080fd5b50519095509350610799565b81600160a060020a0316635cf17bbb8989876040518463ffffffff1660e060020a0281526004018084600181111561100857fe5b60ff16815260200183815260200182600019166000191681526020019350505050602060405180830381600087803b15801561104357600080fd5b505af1158015611057573d6000803e3d6000fd5b505050506040513d602081101561106d57600080fd5b505190505b8080156110f957506040805160008051602061143f8339815191528152600481018690529051600160a060020a038416916394d26cb59160248083019260209291908290030181600087803b1580156110ca57600080fd5b505af11580156110de573d6000803e3d6000fd5b505050506040513d60208110156110f457600080fd5b505115155b80156111f157506040805160008051602061143f8339815191528152600481018690529051600160a060020a038416916331d98b3f9183916394d26cb59160248083019260209291908290030181600087803b15801561115857600080fd5b505af115801561116c573d6000803e3d6000fd5b505050506040513d602081101561118257600080fd5b50516040805160e060020a63ffffffff851602815260048101929092525160248083019260209291908290030181600087803b1580156111c157600080fd5b505af11580156111d5573d6000803e3d6000fd5b505050506040513d60208110156111eb57600080fd5b50518714155b156113b5576040805160008051602061143f8339815191528152600481018690529051600160a060020a038416916394d26cb59160248083019260209291908290030181600087803b15801561124657600080fd5b505af115801561125a573d6000803e3d6000fd5b505050506040513d602081101561127057600080fd5b50516040517f5cf17bbb000000000000000000000000000000000000000000000000000000008152909550600160a060020a03831690635cf17bbb908a908a908990600401808460018111156112c257fe5b60ff16815260200183815260200182600019166000191681526020019350505050602060405180830381600087803b1580156112fd57600080fd5b505af1158015611311573d6000803e3d6000fd5b505050506040513d602081101561132757600080fd5b5051905080156113b0576040805160008051602061143f8339815191528152600481018690529051600160a060020a038416916394d26cb59160248083019260209291908290030181600087803b15801561138157600080fd5b505af1158015611395573d6000803e3d6000fd5b505050506040513d60208110156113ab57600080fd5b505193505b611072565b6040805160008051602061143f8339815191528152600481018690529051600160a060020a038416916394d26cb59160248083019260209291908290030181600087803b15801561140557600080fd5b505af1158015611419573d6000803e3d6000fd5b505050506040513d602081101561142f57600080fd5b50519450505050935093915050560094d26cb50000000000000000000000000000000000000000000000000000000031d98b3f000000000000000000000000000000000000000000000000000000008925f9e900000000000000000000000000000000000000000000000000000000a165627a7a72305820829216ef082a0eecdf894c21932aa2377f844676390f1f3ad7d49eb8a4bbf29a0029`

// DeployOrdersFetcher deploys a new Ethereum contract, binding an instance of OrdersFetcher to it.
func DeployOrdersFetcher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrdersFetcher, error) {
	parsed, err := abi.JSON(strings.NewReader(OrdersFetcherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrdersFetcherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrdersFetcher{OrdersFetcherCaller: OrdersFetcherCaller{contract: contract}, OrdersFetcherTransactor: OrdersFetcherTransactor{contract: contract}, OrdersFetcherFilterer: OrdersFetcherFilterer{contract: contract}}, nil
}

// OrdersFetcher is an auto generated Go binding around an Ethereum contract.
type OrdersFetcher struct {
	OrdersFetcherCaller     // Read-only binding to the contract
	OrdersFetcherTransactor // Write-only binding to the contract
	OrdersFetcherFilterer   // Log filterer for contract events
}

// OrdersFetcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrdersFetcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrdersFetcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrdersFetcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrdersFetcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrdersFetcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrdersFetcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrdersFetcherSession struct {
	Contract     *OrdersFetcher    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrdersFetcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrdersFetcherCallerSession struct {
	Contract *OrdersFetcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OrdersFetcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrdersFetcherTransactorSession struct {
	Contract     *OrdersFetcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OrdersFetcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrdersFetcherRaw struct {
	Contract *OrdersFetcher // Generic contract binding to access the raw methods on
}

// OrdersFetcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrdersFetcherCallerRaw struct {
	Contract *OrdersFetcherCaller // Generic read-only contract binding to access the raw methods on
}

// OrdersFetcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrdersFetcherTransactorRaw struct {
	Contract *OrdersFetcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrdersFetcher creates a new instance of OrdersFetcher, bound to a specific deployed contract.
func NewOrdersFetcher(address common.Address, backend bind.ContractBackend) (*OrdersFetcher, error) {
	contract, err := bindOrdersFetcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrdersFetcher{OrdersFetcherCaller: OrdersFetcherCaller{contract: contract}, OrdersFetcherTransactor: OrdersFetcherTransactor{contract: contract}, OrdersFetcherFilterer: OrdersFetcherFilterer{contract: contract}}, nil
}

// NewOrdersFetcherCaller creates a new read-only instance of OrdersFetcher, bound to a specific deployed contract.
func NewOrdersFetcherCaller(address common.Address, caller bind.ContractCaller) (*OrdersFetcherCaller, error) {
	contract, err := bindOrdersFetcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrdersFetcherCaller{contract: contract}, nil
}

// NewOrdersFetcherTransactor creates a new write-only instance of OrdersFetcher, bound to a specific deployed contract.
func NewOrdersFetcherTransactor(address common.Address, transactor bind.ContractTransactor) (*OrdersFetcherTransactor, error) {
	contract, err := bindOrdersFetcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrdersFetcherTransactor{contract: contract}, nil
}

// NewOrdersFetcherFilterer creates a new log filterer instance of OrdersFetcher, bound to a specific deployed contract.
func NewOrdersFetcherFilterer(address common.Address, filterer bind.ContractFilterer) (*OrdersFetcherFilterer, error) {
	contract, err := bindOrdersFetcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrdersFetcherFilterer{contract: contract}, nil
}

// bindOrdersFetcher binds a generic wrapper to an already deployed contract.
func bindOrdersFetcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrdersFetcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrdersFetcher *OrdersFetcherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrdersFetcher.Contract.OrdersFetcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrdersFetcher *OrdersFetcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrdersFetcher.Contract.OrdersFetcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrdersFetcher *OrdersFetcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrdersFetcher.Contract.OrdersFetcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrdersFetcher *OrdersFetcherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrdersFetcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrdersFetcher *OrdersFetcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrdersFetcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrdersFetcher *OrdersFetcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrdersFetcher.Contract.contract.Transact(opts, method, params...)
}

// AscendOrderList is a free data retrieval call binding the contract method 0xf8266a7a.
//
// Solidity: function ascendOrderList(_type uint8, _price uint256, _lowestOrderId bytes32) constant returns(_betterOrderId bytes32, _worseOrderId bytes32)
func (_OrdersFetcher *OrdersFetcherCaller) AscendOrderList(opts *bind.CallOpts, _type uint8, _price *big.Int, _lowestOrderId [32]byte) (struct {
	BetterOrderId [32]byte
	WorseOrderId  [32]byte
}, error) {
	ret := new(struct {
		BetterOrderId [32]byte
		WorseOrderId  [32]byte
	})
	out := ret
	err := _OrdersFetcher.contract.Call(opts, out, "ascendOrderList", _type, _price, _lowestOrderId)
	return *ret, err
}

// AscendOrderList is a free data retrieval call binding the contract method 0xf8266a7a.
//
// Solidity: function ascendOrderList(_type uint8, _price uint256, _lowestOrderId bytes32) constant returns(_betterOrderId bytes32, _worseOrderId bytes32)
func (_OrdersFetcher *OrdersFetcherSession) AscendOrderList(_type uint8, _price *big.Int, _lowestOrderId [32]byte) (struct {
	BetterOrderId [32]byte
	WorseOrderId  [32]byte
}, error) {
	return _OrdersFetcher.Contract.AscendOrderList(&_OrdersFetcher.CallOpts, _type, _price, _lowestOrderId)
}

// AscendOrderList is a free data retrieval call binding the contract method 0xf8266a7a.
//
// Solidity: function ascendOrderList(_type uint8, _price uint256, _lowestOrderId bytes32) constant returns(_betterOrderId bytes32, _worseOrderId bytes32)
func (_OrdersFetcher *OrdersFetcherCallerSession) AscendOrderList(_type uint8, _price *big.Int, _lowestOrderId [32]byte) (struct {
	BetterOrderId [32]byte
	WorseOrderId  [32]byte
}, error) {
	return _OrdersFetcher.Contract.AscendOrderList(&_OrdersFetcher.CallOpts, _type, _price, _lowestOrderId)
}

// DescendOrderList is a free data retrieval call binding the contract method 0x0e1424b5.
//
// Solidity: function descendOrderList(_type uint8, _price uint256, _highestOrderId bytes32) constant returns(_betterOrderId bytes32, _worseOrderId bytes32)
func (_OrdersFetcher *OrdersFetcherCaller) DescendOrderList(opts *bind.CallOpts, _type uint8, _price *big.Int, _highestOrderId [32]byte) (struct {
	BetterOrderId [32]byte
	WorseOrderId  [32]byte
}, error) {
	ret := new(struct {
		BetterOrderId [32]byte
		WorseOrderId  [32]byte
	})
	out := ret
	err := _OrdersFetcher.contract.Call(opts, out, "descendOrderList", _type, _price, _highestOrderId)
	return *ret, err
}

// DescendOrderList is a free data retrieval call binding the contract method 0x0e1424b5.
//
// Solidity: function descendOrderList(_type uint8, _price uint256, _highestOrderId bytes32) constant returns(_betterOrderId bytes32, _worseOrderId bytes32)
func (_OrdersFetcher *OrdersFetcherSession) DescendOrderList(_type uint8, _price *big.Int, _highestOrderId [32]byte) (struct {
	BetterOrderId [32]byte
	WorseOrderId  [32]byte
}, error) {
	return _OrdersFetcher.Contract.DescendOrderList(&_OrdersFetcher.CallOpts, _type, _price, _highestOrderId)
}

// DescendOrderList is a free data retrieval call binding the contract method 0x0e1424b5.
//
// Solidity: function descendOrderList(_type uint8, _price uint256, _highestOrderId bytes32) constant returns(_betterOrderId bytes32, _worseOrderId bytes32)
func (_OrdersFetcher *OrdersFetcherCallerSession) DescendOrderList(_type uint8, _price *big.Int, _highestOrderId [32]byte) (struct {
	BetterOrderId [32]byte
	WorseOrderId  [32]byte
}, error) {
	return _OrdersFetcher.Contract.DescendOrderList(&_OrdersFetcher.CallOpts, _type, _price, _highestOrderId)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_OrdersFetcher *OrdersFetcherCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OrdersFetcher.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_OrdersFetcher *OrdersFetcherSession) GetController() (common.Address, error) {
	return _OrdersFetcher.Contract.GetController(&_OrdersFetcher.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_OrdersFetcher *OrdersFetcherCallerSession) GetController() (common.Address, error) {
	return _OrdersFetcher.Contract.GetController(&_OrdersFetcher.CallOpts)
}

// FindBoundingOrders is a paid mutator transaction binding the contract method 0x3b01cf3c.
//
// Solidity: function findBoundingOrders(_type uint8, _price uint256, _bestOrderId bytes32, _worstOrderId bytes32, _betterOrderId bytes32, _worseOrderId bytes32) returns(bytes32, bytes32)
func (_OrdersFetcher *OrdersFetcherTransactor) FindBoundingOrders(opts *bind.TransactOpts, _type uint8, _price *big.Int, _bestOrderId [32]byte, _worstOrderId [32]byte, _betterOrderId [32]byte, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _OrdersFetcher.contract.Transact(opts, "findBoundingOrders", _type, _price, _bestOrderId, _worstOrderId, _betterOrderId, _worseOrderId)
}

// FindBoundingOrders is a paid mutator transaction binding the contract method 0x3b01cf3c.
//
// Solidity: function findBoundingOrders(_type uint8, _price uint256, _bestOrderId bytes32, _worstOrderId bytes32, _betterOrderId bytes32, _worseOrderId bytes32) returns(bytes32, bytes32)
func (_OrdersFetcher *OrdersFetcherSession) FindBoundingOrders(_type uint8, _price *big.Int, _bestOrderId [32]byte, _worstOrderId [32]byte, _betterOrderId [32]byte, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _OrdersFetcher.Contract.FindBoundingOrders(&_OrdersFetcher.TransactOpts, _type, _price, _bestOrderId, _worstOrderId, _betterOrderId, _worseOrderId)
}

// FindBoundingOrders is a paid mutator transaction binding the contract method 0x3b01cf3c.
//
// Solidity: function findBoundingOrders(_type uint8, _price uint256, _bestOrderId bytes32, _worstOrderId bytes32, _betterOrderId bytes32, _worseOrderId bytes32) returns(bytes32, bytes32)
func (_OrdersFetcher *OrdersFetcherTransactorSession) FindBoundingOrders(_type uint8, _price *big.Int, _bestOrderId [32]byte, _worstOrderId [32]byte, _betterOrderId [32]byte, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _OrdersFetcher.Contract.FindBoundingOrders(&_OrdersFetcher.TransactOpts, _type, _price, _bestOrderId, _worstOrderId, _betterOrderId, _worseOrderId)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_OrdersFetcher *OrdersFetcherTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _OrdersFetcher.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_OrdersFetcher *OrdersFetcherSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _OrdersFetcher.Contract.SetController(&_OrdersFetcher.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_OrdersFetcher *OrdersFetcherTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _OrdersFetcher.Contract.SetController(&_OrdersFetcher.TransactOpts, _controller)
}


// IOrdersABI is the input ABI used to generate the binding from.
const IOrdersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_moneyEscrowed\",\"type\":\"uint256\"},{\"name\":\"_sharesEscrowed\",\"type\":\"uint256\"}],\"name\":\"getOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"}],\"name\":\"assertIsNotBetterPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementTotalEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"name\":\"_sharesFilled\",\"type\":\"uint256\"},{\"name\":\"_tokensFilled\",\"type\":\"uint256\"}],\"name\":\"recordFillOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"getTotalEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"}],\"name\":\"assertIsNotWorsePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOrderMoneyEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"isBetterPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOutcome\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_moneyEscrowed\",\"type\":\"uint256\"},{\"name\":\"_sharesEscrowed\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"saveOrder\",\"outputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getBestOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getLastOutcomePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getWorseOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getWorstOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getBetterOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"decrementTotalEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOrderType\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"isWorsePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOrderCreator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOrderSharesEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"removeOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IOrdersBin is the compiled bytecode used for deploying new contracts.
const IOrdersBin = `0x`

// DeployIOrders deploys a new Ethereum contract, binding an instance of IOrders to it.
func DeployIOrders(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IOrders, error) {
	parsed, err := abi.JSON(strings.NewReader(IOrdersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IOrdersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IOrders{IOrdersCaller: IOrdersCaller{contract: contract}, IOrdersTransactor: IOrdersTransactor{contract: contract}, IOrdersFilterer: IOrdersFilterer{contract: contract}}, nil
}

// IOrders is an auto generated Go binding around an Ethereum contract.
type IOrders struct {
	IOrdersCaller     // Read-only binding to the contract
	IOrdersTransactor // Write-only binding to the contract
	IOrdersFilterer   // Log filterer for contract events
}

// IOrdersCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOrdersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrdersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOrdersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrdersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOrdersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrdersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOrdersSession struct {
	Contract     *IOrders          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOrdersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOrdersCallerSession struct {
	Contract *IOrdersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IOrdersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOrdersTransactorSession struct {
	Contract     *IOrdersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IOrdersRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOrdersRaw struct {
	Contract *IOrders // Generic contract binding to access the raw methods on
}

// IOrdersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOrdersCallerRaw struct {
	Contract *IOrdersCaller // Generic read-only contract binding to access the raw methods on
}

// IOrdersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOrdersTransactorRaw struct {
	Contract *IOrdersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOrders creates a new instance of IOrders, bound to a specific deployed contract.
func NewIOrders(address common.Address, backend bind.ContractBackend) (*IOrders, error) {
	contract, err := bindIOrders(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOrders{IOrdersCaller: IOrdersCaller{contract: contract}, IOrdersTransactor: IOrdersTransactor{contract: contract}, IOrdersFilterer: IOrdersFilterer{contract: contract}}, nil
}

// NewIOrdersCaller creates a new read-only instance of IOrders, bound to a specific deployed contract.
func NewIOrdersCaller(address common.Address, caller bind.ContractCaller) (*IOrdersCaller, error) {
	contract, err := bindIOrders(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOrdersCaller{contract: contract}, nil
}

// NewIOrdersTransactor creates a new write-only instance of IOrders, bound to a specific deployed contract.
func NewIOrdersTransactor(address common.Address, transactor bind.ContractTransactor) (*IOrdersTransactor, error) {
	contract, err := bindIOrders(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOrdersTransactor{contract: contract}, nil
}

// NewIOrdersFilterer creates a new log filterer instance of IOrders, bound to a specific deployed contract.
func NewIOrdersFilterer(address common.Address, filterer bind.ContractFilterer) (*IOrdersFilterer, error) {
	contract, err := bindIOrders(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOrdersFilterer{contract: contract}, nil
}

// bindIOrders binds a generic wrapper to an already deployed contract.
func bindIOrders(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOrdersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrders *IOrdersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IOrders.Contract.IOrdersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrders *IOrdersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrders.Contract.IOrdersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrders *IOrdersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrders.Contract.IOrdersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrders *IOrdersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IOrders.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrders *IOrdersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrders.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrders *IOrdersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrders.Contract.contract.Transact(opts, method, params...)
}

// AssertIsNotBetterPrice is a free data retrieval call binding the contract method 0x19e54fb3.
//
// Solidity: function assertIsNotBetterPrice(_type uint8, _price uint256, _betterOrderId bytes32) constant returns(bool)
func (_IOrders *IOrdersCaller) AssertIsNotBetterPrice(opts *bind.CallOpts, _type uint8, _price *big.Int, _betterOrderId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "assertIsNotBetterPrice", _type, _price, _betterOrderId)
	return *ret0, err
}

// AssertIsNotBetterPrice is a free data retrieval call binding the contract method 0x19e54fb3.
//
// Solidity: function assertIsNotBetterPrice(_type uint8, _price uint256, _betterOrderId bytes32) constant returns(bool)
func (_IOrders *IOrdersSession) AssertIsNotBetterPrice(_type uint8, _price *big.Int, _betterOrderId [32]byte) (bool, error) {
	return _IOrders.Contract.AssertIsNotBetterPrice(&_IOrders.CallOpts, _type, _price, _betterOrderId)
}

// AssertIsNotBetterPrice is a free data retrieval call binding the contract method 0x19e54fb3.
//
// Solidity: function assertIsNotBetterPrice(_type uint8, _price uint256, _betterOrderId bytes32) constant returns(bool)
func (_IOrders *IOrdersCallerSession) AssertIsNotBetterPrice(_type uint8, _price *big.Int, _betterOrderId [32]byte) (bool, error) {
	return _IOrders.Contract.AssertIsNotBetterPrice(&_IOrders.CallOpts, _type, _price, _betterOrderId)
}

// GetAmount is a free data retrieval call binding the contract method 0x0e0a0d74.
//
// Solidity: function getAmount(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCaller) GetAmount(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getAmount", _orderId)
	return *ret0, err
}

// GetAmount is a free data retrieval call binding the contract method 0x0e0a0d74.
//
// Solidity: function getAmount(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersSession) GetAmount(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetAmount(&_IOrders.CallOpts, _orderId)
}

// GetAmount is a free data retrieval call binding the contract method 0x0e0a0d74.
//
// Solidity: function getAmount(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCallerSession) GetAmount(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetAmount(&_IOrders.CallOpts, _orderId)
}

// GetBestOrderId is a free data retrieval call binding the contract method 0x7b6eaa65.
//
// Solidity: function getBestOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_IOrders *IOrdersCaller) GetBestOrderId(opts *bind.CallOpts, _type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getBestOrderId", _type, _market, _outcome)
	return *ret0, err
}

// GetBestOrderId is a free data retrieval call binding the contract method 0x7b6eaa65.
//
// Solidity: function getBestOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_IOrders *IOrdersSession) GetBestOrderId(_type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	return _IOrders.Contract.GetBestOrderId(&_IOrders.CallOpts, _type, _market, _outcome)
}

// GetBestOrderId is a free data retrieval call binding the contract method 0x7b6eaa65.
//
// Solidity: function getBestOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_IOrders *IOrdersCallerSession) GetBestOrderId(_type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	return _IOrders.Contract.GetBestOrderId(&_IOrders.CallOpts, _type, _market, _outcome)
}

// GetBetterOrderId is a free data retrieval call binding the contract method 0x94d26cb5.
//
// Solidity: function getBetterOrderId(_orderId bytes32) constant returns(bytes32)
func (_IOrders *IOrdersCaller) GetBetterOrderId(opts *bind.CallOpts, _orderId [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getBetterOrderId", _orderId)
	return *ret0, err
}

// GetBetterOrderId is a free data retrieval call binding the contract method 0x94d26cb5.
//
// Solidity: function getBetterOrderId(_orderId bytes32) constant returns(bytes32)
func (_IOrders *IOrdersSession) GetBetterOrderId(_orderId [32]byte) ([32]byte, error) {
	return _IOrders.Contract.GetBetterOrderId(&_IOrders.CallOpts, _orderId)
}

// GetBetterOrderId is a free data retrieval call binding the contract method 0x94d26cb5.
//
// Solidity: function getBetterOrderId(_orderId bytes32) constant returns(bytes32)
func (_IOrders *IOrdersCallerSession) GetBetterOrderId(_orderId [32]byte) ([32]byte, error) {
	return _IOrders.Contract.GetBetterOrderId(&_IOrders.CallOpts, _orderId)
}

// GetLastOutcomePrice is a free data retrieval call binding the contract method 0x804fb410.
//
// Solidity: function getLastOutcomePrice(_market address, _outcome uint256) constant returns(uint256)
func (_IOrders *IOrdersCaller) GetLastOutcomePrice(opts *bind.CallOpts, _market common.Address, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getLastOutcomePrice", _market, _outcome)
	return *ret0, err
}

// GetLastOutcomePrice is a free data retrieval call binding the contract method 0x804fb410.
//
// Solidity: function getLastOutcomePrice(_market address, _outcome uint256) constant returns(uint256)
func (_IOrders *IOrdersSession) GetLastOutcomePrice(_market common.Address, _outcome *big.Int) (*big.Int, error) {
	return _IOrders.Contract.GetLastOutcomePrice(&_IOrders.CallOpts, _market, _outcome)
}

// GetLastOutcomePrice is a free data retrieval call binding the contract method 0x804fb410.
//
// Solidity: function getLastOutcomePrice(_market address, _outcome uint256) constant returns(uint256)
func (_IOrders *IOrdersCallerSession) GetLastOutcomePrice(_market common.Address, _outcome *big.Int) (*big.Int, error) {
	return _IOrders.Contract.GetLastOutcomePrice(&_IOrders.CallOpts, _market, _outcome)
}

// GetMarket is a free data retrieval call binding the contract method 0xc3c95c7b.
//
// Solidity: function getMarket(_orderId bytes32) constant returns(address)
func (_IOrders *IOrdersCaller) GetMarket(opts *bind.CallOpts, _orderId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getMarket", _orderId)
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xc3c95c7b.
//
// Solidity: function getMarket(_orderId bytes32) constant returns(address)
func (_IOrders *IOrdersSession) GetMarket(_orderId [32]byte) (common.Address, error) {
	return _IOrders.Contract.GetMarket(&_IOrders.CallOpts, _orderId)
}

// GetMarket is a free data retrieval call binding the contract method 0xc3c95c7b.
//
// Solidity: function getMarket(_orderId bytes32) constant returns(address)
func (_IOrders *IOrdersCallerSession) GetMarket(_orderId [32]byte) (common.Address, error) {
	return _IOrders.Contract.GetMarket(&_IOrders.CallOpts, _orderId)
}

// GetOrderCreator is a free data retrieval call binding the contract method 0xe7d80c70.
//
// Solidity: function getOrderCreator(_orderId bytes32) constant returns(address)
func (_IOrders *IOrdersCaller) GetOrderCreator(opts *bind.CallOpts, _orderId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getOrderCreator", _orderId)
	return *ret0, err
}

// GetOrderCreator is a free data retrieval call binding the contract method 0xe7d80c70.
//
// Solidity: function getOrderCreator(_orderId bytes32) constant returns(address)
func (_IOrders *IOrdersSession) GetOrderCreator(_orderId [32]byte) (common.Address, error) {
	return _IOrders.Contract.GetOrderCreator(&_IOrders.CallOpts, _orderId)
}

// GetOrderCreator is a free data retrieval call binding the contract method 0xe7d80c70.
//
// Solidity: function getOrderCreator(_orderId bytes32) constant returns(address)
func (_IOrders *IOrdersCallerSession) GetOrderCreator(_orderId [32]byte) (common.Address, error) {
	return _IOrders.Contract.GetOrderCreator(&_IOrders.CallOpts, _orderId)
}

// GetOrderId is a free data retrieval call binding the contract method 0x1727583c.
//
// Solidity: function getOrderId(_type uint8, _market address, _fxpAmount uint256, _price uint256, _sender address, _blockNumber uint256, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256) constant returns(bytes32)
func (_IOrders *IOrdersCaller) GetOrderId(opts *bind.CallOpts, _type uint8, _market common.Address, _fxpAmount *big.Int, _price *big.Int, _sender common.Address, _blockNumber *big.Int, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getOrderId", _type, _market, _fxpAmount, _price, _sender, _blockNumber, _outcome, _moneyEscrowed, _sharesEscrowed)
	return *ret0, err
}

// GetOrderId is a free data retrieval call binding the contract method 0x1727583c.
//
// Solidity: function getOrderId(_type uint8, _market address, _fxpAmount uint256, _price uint256, _sender address, _blockNumber uint256, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256) constant returns(bytes32)
func (_IOrders *IOrdersSession) GetOrderId(_type uint8, _market common.Address, _fxpAmount *big.Int, _price *big.Int, _sender common.Address, _blockNumber *big.Int, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int) ([32]byte, error) {
	return _IOrders.Contract.GetOrderId(&_IOrders.CallOpts, _type, _market, _fxpAmount, _price, _sender, _blockNumber, _outcome, _moneyEscrowed, _sharesEscrowed)
}

// GetOrderId is a free data retrieval call binding the contract method 0x1727583c.
//
// Solidity: function getOrderId(_type uint8, _market address, _fxpAmount uint256, _price uint256, _sender address, _blockNumber uint256, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256) constant returns(bytes32)
func (_IOrders *IOrdersCallerSession) GetOrderId(_type uint8, _market common.Address, _fxpAmount *big.Int, _price *big.Int, _sender common.Address, _blockNumber *big.Int, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int) ([32]byte, error) {
	return _IOrders.Contract.GetOrderId(&_IOrders.CallOpts, _type, _market, _fxpAmount, _price, _sender, _blockNumber, _outcome, _moneyEscrowed, _sharesEscrowed)
}

// GetOrderMoneyEscrowed is a free data retrieval call binding the contract method 0x4a1a342b.
//
// Solidity: function getOrderMoneyEscrowed(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCaller) GetOrderMoneyEscrowed(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getOrderMoneyEscrowed", _orderId)
	return *ret0, err
}

// GetOrderMoneyEscrowed is a free data retrieval call binding the contract method 0x4a1a342b.
//
// Solidity: function getOrderMoneyEscrowed(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersSession) GetOrderMoneyEscrowed(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetOrderMoneyEscrowed(&_IOrders.CallOpts, _orderId)
}

// GetOrderMoneyEscrowed is a free data retrieval call binding the contract method 0x4a1a342b.
//
// Solidity: function getOrderMoneyEscrowed(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCallerSession) GetOrderMoneyEscrowed(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetOrderMoneyEscrowed(&_IOrders.CallOpts, _orderId)
}

// GetOrderSharesEscrowed is a free data retrieval call binding the contract method 0xebead05f.
//
// Solidity: function getOrderSharesEscrowed(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCaller) GetOrderSharesEscrowed(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getOrderSharesEscrowed", _orderId)
	return *ret0, err
}

// GetOrderSharesEscrowed is a free data retrieval call binding the contract method 0xebead05f.
//
// Solidity: function getOrderSharesEscrowed(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersSession) GetOrderSharesEscrowed(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetOrderSharesEscrowed(&_IOrders.CallOpts, _orderId)
}

// GetOrderSharesEscrowed is a free data retrieval call binding the contract method 0xebead05f.
//
// Solidity: function getOrderSharesEscrowed(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCallerSession) GetOrderSharesEscrowed(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetOrderSharesEscrowed(&_IOrders.CallOpts, _orderId)
}

// GetOrderType is a free data retrieval call binding the contract method 0xcf357364.
//
// Solidity: function getOrderType(_orderId bytes32) constant returns(uint8)
func (_IOrders *IOrdersCaller) GetOrderType(opts *bind.CallOpts, _orderId [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getOrderType", _orderId)
	return *ret0, err
}

// GetOrderType is a free data retrieval call binding the contract method 0xcf357364.
//
// Solidity: function getOrderType(_orderId bytes32) constant returns(uint8)
func (_IOrders *IOrdersSession) GetOrderType(_orderId [32]byte) (uint8, error) {
	return _IOrders.Contract.GetOrderType(&_IOrders.CallOpts, _orderId)
}

// GetOrderType is a free data retrieval call binding the contract method 0xcf357364.
//
// Solidity: function getOrderType(_orderId bytes32) constant returns(uint8)
func (_IOrders *IOrdersCallerSession) GetOrderType(_orderId [32]byte) (uint8, error) {
	return _IOrders.Contract.GetOrderType(&_IOrders.CallOpts, _orderId)
}

// GetOutcome is a free data retrieval call binding the contract method 0x5d1a3b82.
//
// Solidity: function getOutcome(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCaller) GetOutcome(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getOutcome", _orderId)
	return *ret0, err
}

// GetOutcome is a free data retrieval call binding the contract method 0x5d1a3b82.
//
// Solidity: function getOutcome(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersSession) GetOutcome(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetOutcome(&_IOrders.CallOpts, _orderId)
}

// GetOutcome is a free data retrieval call binding the contract method 0x5d1a3b82.
//
// Solidity: function getOutcome(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCallerSession) GetOutcome(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetOutcome(&_IOrders.CallOpts, _orderId)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCaller) GetPrice(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getPrice", _orderId)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersSession) GetPrice(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetPrice(&_IOrders.CallOpts, _orderId)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(_orderId bytes32) constant returns(uint256)
func (_IOrders *IOrdersCallerSession) GetPrice(_orderId [32]byte) (*big.Int, error) {
	return _IOrders.Contract.GetPrice(&_IOrders.CallOpts, _orderId)
}

// GetTotalEscrowed is a free data retrieval call binding the contract method 0x37ec114d.
//
// Solidity: function getTotalEscrowed(_market address) constant returns(uint256)
func (_IOrders *IOrdersCaller) GetTotalEscrowed(opts *bind.CallOpts, _market common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getTotalEscrowed", _market)
	return *ret0, err
}

// GetTotalEscrowed is a free data retrieval call binding the contract method 0x37ec114d.
//
// Solidity: function getTotalEscrowed(_market address) constant returns(uint256)
func (_IOrders *IOrdersSession) GetTotalEscrowed(_market common.Address) (*big.Int, error) {
	return _IOrders.Contract.GetTotalEscrowed(&_IOrders.CallOpts, _market)
}

// GetTotalEscrowed is a free data retrieval call binding the contract method 0x37ec114d.
//
// Solidity: function getTotalEscrowed(_market address) constant returns(uint256)
func (_IOrders *IOrdersCallerSession) GetTotalEscrowed(_market common.Address) (*big.Int, error) {
	return _IOrders.Contract.GetTotalEscrowed(&_IOrders.CallOpts, _market)
}

// GetWorseOrderId is a free data retrieval call binding the contract method 0x8925f9e9.
//
// Solidity: function getWorseOrderId(_orderId bytes32) constant returns(bytes32)
func (_IOrders *IOrdersCaller) GetWorseOrderId(opts *bind.CallOpts, _orderId [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getWorseOrderId", _orderId)
	return *ret0, err
}

// GetWorseOrderId is a free data retrieval call binding the contract method 0x8925f9e9.
//
// Solidity: function getWorseOrderId(_orderId bytes32) constant returns(bytes32)
func (_IOrders *IOrdersSession) GetWorseOrderId(_orderId [32]byte) ([32]byte, error) {
	return _IOrders.Contract.GetWorseOrderId(&_IOrders.CallOpts, _orderId)
}

// GetWorseOrderId is a free data retrieval call binding the contract method 0x8925f9e9.
//
// Solidity: function getWorseOrderId(_orderId bytes32) constant returns(bytes32)
func (_IOrders *IOrdersCallerSession) GetWorseOrderId(_orderId [32]byte) ([32]byte, error) {
	return _IOrders.Contract.GetWorseOrderId(&_IOrders.CallOpts, _orderId)
}

// GetWorstOrderId is a free data retrieval call binding the contract method 0x8e12ebad.
//
// Solidity: function getWorstOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_IOrders *IOrdersCaller) GetWorstOrderId(opts *bind.CallOpts, _type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "getWorstOrderId", _type, _market, _outcome)
	return *ret0, err
}

// GetWorstOrderId is a free data retrieval call binding the contract method 0x8e12ebad.
//
// Solidity: function getWorstOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_IOrders *IOrdersSession) GetWorstOrderId(_type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	return _IOrders.Contract.GetWorstOrderId(&_IOrders.CallOpts, _type, _market, _outcome)
}

// GetWorstOrderId is a free data retrieval call binding the contract method 0x8e12ebad.
//
// Solidity: function getWorstOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_IOrders *IOrdersCallerSession) GetWorstOrderId(_type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	return _IOrders.Contract.GetWorstOrderId(&_IOrders.CallOpts, _type, _market, _outcome)
}

// IsBetterPrice is a free data retrieval call binding the contract method 0x5cf17bbb.
//
// Solidity: function isBetterPrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_IOrders *IOrdersCaller) IsBetterPrice(opts *bind.CallOpts, _type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "isBetterPrice", _type, _price, _orderId)
	return *ret0, err
}

// IsBetterPrice is a free data retrieval call binding the contract method 0x5cf17bbb.
//
// Solidity: function isBetterPrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_IOrders *IOrdersSession) IsBetterPrice(_type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	return _IOrders.Contract.IsBetterPrice(&_IOrders.CallOpts, _type, _price, _orderId)
}

// IsBetterPrice is a free data retrieval call binding the contract method 0x5cf17bbb.
//
// Solidity: function isBetterPrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_IOrders *IOrdersCallerSession) IsBetterPrice(_type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	return _IOrders.Contract.IsBetterPrice(&_IOrders.CallOpts, _type, _price, _orderId)
}

// IsWorsePrice is a free data retrieval call binding the contract method 0xd9b3d9fe.
//
// Solidity: function isWorsePrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_IOrders *IOrdersCaller) IsWorsePrice(opts *bind.CallOpts, _type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IOrders.contract.Call(opts, out, "isWorsePrice", _type, _price, _orderId)
	return *ret0, err
}

// IsWorsePrice is a free data retrieval call binding the contract method 0xd9b3d9fe.
//
// Solidity: function isWorsePrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_IOrders *IOrdersSession) IsWorsePrice(_type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	return _IOrders.Contract.IsWorsePrice(&_IOrders.CallOpts, _type, _price, _orderId)
}

// IsWorsePrice is a free data retrieval call binding the contract method 0xd9b3d9fe.
//
// Solidity: function isWorsePrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_IOrders *IOrdersCallerSession) IsWorsePrice(_type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	return _IOrders.Contract.IsWorsePrice(&_IOrders.CallOpts, _type, _price, _orderId)
}

// AssertIsNotWorsePrice is a paid mutator transaction binding the contract method 0x45c92c38.
//
// Solidity: function assertIsNotWorsePrice(_type uint8, _price uint256, _worseOrderId bytes32) returns(bool)
func (_IOrders *IOrdersTransactor) AssertIsNotWorsePrice(opts *bind.TransactOpts, _type uint8, _price *big.Int, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _IOrders.contract.Transact(opts, "assertIsNotWorsePrice", _type, _price, _worseOrderId)
}

// AssertIsNotWorsePrice is a paid mutator transaction binding the contract method 0x45c92c38.
//
// Solidity: function assertIsNotWorsePrice(_type uint8, _price uint256, _worseOrderId bytes32) returns(bool)
func (_IOrders *IOrdersSession) AssertIsNotWorsePrice(_type uint8, _price *big.Int, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _IOrders.Contract.AssertIsNotWorsePrice(&_IOrders.TransactOpts, _type, _price, _worseOrderId)
}

// AssertIsNotWorsePrice is a paid mutator transaction binding the contract method 0x45c92c38.
//
// Solidity: function assertIsNotWorsePrice(_type uint8, _price uint256, _worseOrderId bytes32) returns(bool)
func (_IOrders *IOrdersTransactorSession) AssertIsNotWorsePrice(_type uint8, _price *big.Int, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _IOrders.Contract.AssertIsNotWorsePrice(&_IOrders.TransactOpts, _type, _price, _worseOrderId)
}

// DecrementTotalEscrowed is a paid mutator transaction binding the contract method 0x95ede032.
//
// Solidity: function decrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_IOrders *IOrdersTransactor) DecrementTotalEscrowed(opts *bind.TransactOpts, _market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IOrders.contract.Transact(opts, "decrementTotalEscrowed", _market, _amount)
}

// DecrementTotalEscrowed is a paid mutator transaction binding the contract method 0x95ede032.
//
// Solidity: function decrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_IOrders *IOrdersSession) DecrementTotalEscrowed(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IOrders.Contract.DecrementTotalEscrowed(&_IOrders.TransactOpts, _market, _amount)
}

// DecrementTotalEscrowed is a paid mutator transaction binding the contract method 0x95ede032.
//
// Solidity: function decrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_IOrders *IOrdersTransactorSession) DecrementTotalEscrowed(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IOrders.Contract.DecrementTotalEscrowed(&_IOrders.TransactOpts, _market, _amount)
}

// IncrementTotalEscrowed is a paid mutator transaction binding the contract method 0x2c491293.
//
// Solidity: function incrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_IOrders *IOrdersTransactor) IncrementTotalEscrowed(opts *bind.TransactOpts, _market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IOrders.contract.Transact(opts, "incrementTotalEscrowed", _market, _amount)
}

// IncrementTotalEscrowed is a paid mutator transaction binding the contract method 0x2c491293.
//
// Solidity: function incrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_IOrders *IOrdersSession) IncrementTotalEscrowed(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IOrders.Contract.IncrementTotalEscrowed(&_IOrders.TransactOpts, _market, _amount)
}

// IncrementTotalEscrowed is a paid mutator transaction binding the contract method 0x2c491293.
//
// Solidity: function incrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_IOrders *IOrdersTransactorSession) IncrementTotalEscrowed(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IOrders.Contract.IncrementTotalEscrowed(&_IOrders.TransactOpts, _market, _amount)
}

// RecordFillOrder is a paid mutator transaction binding the contract method 0x2ed5ca29.
//
// Solidity: function recordFillOrder(_orderId bytes32, _sharesFilled uint256, _tokensFilled uint256) returns(bool)
func (_IOrders *IOrdersTransactor) RecordFillOrder(opts *bind.TransactOpts, _orderId [32]byte, _sharesFilled *big.Int, _tokensFilled *big.Int) (*types.Transaction, error) {
	return _IOrders.contract.Transact(opts, "recordFillOrder", _orderId, _sharesFilled, _tokensFilled)
}

// RecordFillOrder is a paid mutator transaction binding the contract method 0x2ed5ca29.
//
// Solidity: function recordFillOrder(_orderId bytes32, _sharesFilled uint256, _tokensFilled uint256) returns(bool)
func (_IOrders *IOrdersSession) RecordFillOrder(_orderId [32]byte, _sharesFilled *big.Int, _tokensFilled *big.Int) (*types.Transaction, error) {
	return _IOrders.Contract.RecordFillOrder(&_IOrders.TransactOpts, _orderId, _sharesFilled, _tokensFilled)
}

// RecordFillOrder is a paid mutator transaction binding the contract method 0x2ed5ca29.
//
// Solidity: function recordFillOrder(_orderId bytes32, _sharesFilled uint256, _tokensFilled uint256) returns(bool)
func (_IOrders *IOrdersTransactorSession) RecordFillOrder(_orderId [32]byte, _sharesFilled *big.Int, _tokensFilled *big.Int) (*types.Transaction, error) {
	return _IOrders.Contract.RecordFillOrder(&_IOrders.TransactOpts, _orderId, _sharesFilled, _tokensFilled)
}

// RemoveOrder is a paid mutator transaction binding the contract method 0xfde99668.
//
// Solidity: function removeOrder(_orderId bytes32) returns(bool)
func (_IOrders *IOrdersTransactor) RemoveOrder(opts *bind.TransactOpts, _orderId [32]byte) (*types.Transaction, error) {
	return _IOrders.contract.Transact(opts, "removeOrder", _orderId)
}

// RemoveOrder is a paid mutator transaction binding the contract method 0xfde99668.
//
// Solidity: function removeOrder(_orderId bytes32) returns(bool)
func (_IOrders *IOrdersSession) RemoveOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _IOrders.Contract.RemoveOrder(&_IOrders.TransactOpts, _orderId)
}

// RemoveOrder is a paid mutator transaction binding the contract method 0xfde99668.
//
// Solidity: function removeOrder(_orderId bytes32) returns(bool)
func (_IOrders *IOrdersTransactorSession) RemoveOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _IOrders.Contract.RemoveOrder(&_IOrders.TransactOpts, _orderId)
}

// SaveOrder is a paid mutator transaction binding the contract method 0x6bc29efa.
//
// Solidity: function saveOrder(_type uint8, _market address, _fxpAmount uint256, _price uint256, _sender address, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(_orderId bytes32)
func (_IOrders *IOrdersTransactor) SaveOrder(opts *bind.TransactOpts, _type uint8, _market common.Address, _fxpAmount *big.Int, _price *big.Int, _sender common.Address, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _IOrders.contract.Transact(opts, "saveOrder", _type, _market, _fxpAmount, _price, _sender, _outcome, _moneyEscrowed, _sharesEscrowed, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// SaveOrder is a paid mutator transaction binding the contract method 0x6bc29efa.
//
// Solidity: function saveOrder(_type uint8, _market address, _fxpAmount uint256, _price uint256, _sender address, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(_orderId bytes32)
func (_IOrders *IOrdersSession) SaveOrder(_type uint8, _market common.Address, _fxpAmount *big.Int, _price *big.Int, _sender common.Address, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _IOrders.Contract.SaveOrder(&_IOrders.TransactOpts, _type, _market, _fxpAmount, _price, _sender, _outcome, _moneyEscrowed, _sharesEscrowed, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// SaveOrder is a paid mutator transaction binding the contract method 0x6bc29efa.
//
// Solidity: function saveOrder(_type uint8, _market address, _fxpAmount uint256, _price uint256, _sender address, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(_orderId bytes32)
func (_IOrders *IOrdersTransactorSession) SaveOrder(_type uint8, _market common.Address, _fxpAmount *big.Int, _price *big.Int, _sender common.Address, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _IOrders.Contract.SaveOrder(&_IOrders.TransactOpts, _type, _market, _fxpAmount, _price, _sender, _outcome, _moneyEscrowed, _sharesEscrowed, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// SetPrice is a paid mutator transaction binding the contract method 0x3011e16a.
//
// Solidity: function setPrice(_market address, _outcome uint256, _price uint256) returns(bool)
func (_IOrders *IOrdersTransactor) SetPrice(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _IOrders.contract.Transact(opts, "setPrice", _market, _outcome, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x3011e16a.
//
// Solidity: function setPrice(_market address, _outcome uint256, _price uint256) returns(bool)
func (_IOrders *IOrdersSession) SetPrice(_market common.Address, _outcome *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _IOrders.Contract.SetPrice(&_IOrders.TransactOpts, _market, _outcome, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x3011e16a.
//
// Solidity: function setPrice(_market address, _outcome uint256, _price uint256) returns(bool)
func (_IOrders *IOrdersTransactorSession) SetPrice(_market common.Address, _outcome *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _IOrders.Contract.SetPrice(&_IOrders.TransactOpts, _market, _outcome, _price)
}



// IOrdersFetcherABI is the input ABI used to generate the binding from.
const IOrdersFetcherABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_bestOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worstOrderId\",\"type\":\"bytes32\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"}],\"name\":\"findBoundingOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IOrdersFetcherBin is the compiled bytecode used for deploying new contracts.
const IOrdersFetcherBin = `0x`

// DeployIOrdersFetcher deploys a new Ethereum contract, binding an instance of IOrdersFetcher to it.
func DeployIOrdersFetcher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IOrdersFetcher, error) {
	parsed, err := abi.JSON(strings.NewReader(IOrdersFetcherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IOrdersFetcherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IOrdersFetcher{IOrdersFetcherCaller: IOrdersFetcherCaller{contract: contract}, IOrdersFetcherTransactor: IOrdersFetcherTransactor{contract: contract}, IOrdersFetcherFilterer: IOrdersFetcherFilterer{contract: contract}}, nil
}

// IOrdersFetcher is an auto generated Go binding around an Ethereum contract.
type IOrdersFetcher struct {
	IOrdersFetcherCaller     // Read-only binding to the contract
	IOrdersFetcherTransactor // Write-only binding to the contract
	IOrdersFetcherFilterer   // Log filterer for contract events
}

// IOrdersFetcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOrdersFetcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrdersFetcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOrdersFetcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrdersFetcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOrdersFetcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrdersFetcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOrdersFetcherSession struct {
	Contract     *IOrdersFetcher   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOrdersFetcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOrdersFetcherCallerSession struct {
	Contract *IOrdersFetcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IOrdersFetcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOrdersFetcherTransactorSession struct {
	Contract     *IOrdersFetcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IOrdersFetcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOrdersFetcherRaw struct {
	Contract *IOrdersFetcher // Generic contract binding to access the raw methods on
}

// IOrdersFetcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOrdersFetcherCallerRaw struct {
	Contract *IOrdersFetcherCaller // Generic read-only contract binding to access the raw methods on
}

// IOrdersFetcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOrdersFetcherTransactorRaw struct {
	Contract *IOrdersFetcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOrdersFetcher creates a new instance of IOrdersFetcher, bound to a specific deployed contract.
func NewIOrdersFetcher(address common.Address, backend bind.ContractBackend) (*IOrdersFetcher, error) {
	contract, err := bindIOrdersFetcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOrdersFetcher{IOrdersFetcherCaller: IOrdersFetcherCaller{contract: contract}, IOrdersFetcherTransactor: IOrdersFetcherTransactor{contract: contract}, IOrdersFetcherFilterer: IOrdersFetcherFilterer{contract: contract}}, nil
}

// NewIOrdersFetcherCaller creates a new read-only instance of IOrdersFetcher, bound to a specific deployed contract.
func NewIOrdersFetcherCaller(address common.Address, caller bind.ContractCaller) (*IOrdersFetcherCaller, error) {
	contract, err := bindIOrdersFetcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOrdersFetcherCaller{contract: contract}, nil
}

// NewIOrdersFetcherTransactor creates a new write-only instance of IOrdersFetcher, bound to a specific deployed contract.
func NewIOrdersFetcherTransactor(address common.Address, transactor bind.ContractTransactor) (*IOrdersFetcherTransactor, error) {
	contract, err := bindIOrdersFetcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOrdersFetcherTransactor{contract: contract}, nil
}

// NewIOrdersFetcherFilterer creates a new log filterer instance of IOrdersFetcher, bound to a specific deployed contract.
func NewIOrdersFetcherFilterer(address common.Address, filterer bind.ContractFilterer) (*IOrdersFetcherFilterer, error) {
	contract, err := bindIOrdersFetcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOrdersFetcherFilterer{contract: contract}, nil
}

// bindIOrdersFetcher binds a generic wrapper to an already deployed contract.
func bindIOrdersFetcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOrdersFetcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrdersFetcher *IOrdersFetcherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IOrdersFetcher.Contract.IOrdersFetcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrdersFetcher *IOrdersFetcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrdersFetcher.Contract.IOrdersFetcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrdersFetcher *IOrdersFetcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrdersFetcher.Contract.IOrdersFetcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrdersFetcher *IOrdersFetcherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IOrdersFetcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrdersFetcher *IOrdersFetcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrdersFetcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrdersFetcher *IOrdersFetcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrdersFetcher.Contract.contract.Transact(opts, method, params...)
}

// FindBoundingOrders is a paid mutator transaction binding the contract method 0x3b01cf3c.
//
// Solidity: function findBoundingOrders(_type uint8, _price uint256, _bestOrderId bytes32, _worstOrderId bytes32, _betterOrderId bytes32, _worseOrderId bytes32) returns(bytes32, bytes32)
func (_IOrdersFetcher *IOrdersFetcherTransactor) FindBoundingOrders(opts *bind.TransactOpts, _type uint8, _price *big.Int, _bestOrderId [32]byte, _worstOrderId [32]byte, _betterOrderId [32]byte, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _IOrdersFetcher.contract.Transact(opts, "findBoundingOrders", _type, _price, _bestOrderId, _worstOrderId, _betterOrderId, _worseOrderId)
}

// FindBoundingOrders is a paid mutator transaction binding the contract method 0x3b01cf3c.
//
// Solidity: function findBoundingOrders(_type uint8, _price uint256, _bestOrderId bytes32, _worstOrderId bytes32, _betterOrderId bytes32, _worseOrderId bytes32) returns(bytes32, bytes32)
func (_IOrdersFetcher *IOrdersFetcherSession) FindBoundingOrders(_type uint8, _price *big.Int, _bestOrderId [32]byte, _worstOrderId [32]byte, _betterOrderId [32]byte, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _IOrdersFetcher.Contract.FindBoundingOrders(&_IOrdersFetcher.TransactOpts, _type, _price, _bestOrderId, _worstOrderId, _betterOrderId, _worseOrderId)
}

// FindBoundingOrders is a paid mutator transaction binding the contract method 0x3b01cf3c.
//
// Solidity: function findBoundingOrders(_type uint8, _price uint256, _bestOrderId bytes32, _worstOrderId bytes32, _betterOrderId bytes32, _worseOrderId bytes32) returns(bytes32, bytes32)
func (_IOrdersFetcher *IOrdersFetcherTransactorSession) FindBoundingOrders(_type uint8, _price *big.Int, _bestOrderId [32]byte, _worstOrderId [32]byte, _betterOrderId [32]byte, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _IOrdersFetcher.Contract.FindBoundingOrders(&_IOrdersFetcher.TransactOpts, _type, _price, _bestOrderId, _worstOrderId, _betterOrderId, _worseOrderId)
}
