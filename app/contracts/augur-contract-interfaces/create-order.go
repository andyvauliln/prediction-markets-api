package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CreateOrderABI is the input ABI used to generate the binding from.
const CreateOrderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_attoshares\",\"type\":\"uint256\"},{\"name\":\"_displayPrice\",\"type\":\"uint256\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"publicCreateOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_creator\",\"type\":\"address\"},{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_attoshares\",\"type\":\"uint256\"},{\"name\":\"_displayPrice\",\"type\":\"uint256\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"createOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CreateOrderBin is the compiled bytecode used for deploying new contracts.
const CreateOrderBin = `0x608060405260008054600160a860020a03191633179055611b35806100256000396000f3006080604052600436106100485763ffffffff60e060020a6000350416633018205f811461004d5780633ae4ce0a1461007e57806392eefe9b146100bc578063f7938328146100f1575b600080fd5b34801561005957600080fd5b50610062610134565b60408051600160a060020a039092168252519081900360200190f35b6100aa60ff60043516602435604435600160a060020a036064351660843560a43560c43560e435610143565b60408051918252519081900360200190f35b3480156100c857600080fd5b506100dd600160a060020a0360043516610508565b604080519115158252519081900360200190f35b3480156100fd57600080fd5b506100aa600160a060020a0360043581169060ff602435169060443590606435906084351660a43560c43560e43561010435610551565b600054600160a060020a031690565b60008086600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561018757600080fd5b505af115801561019b573d6000803e3d6000fd5b505050506040513d60208110156101b157600080fd5b505160008054604080517f4e94c8290000000000000000000000000000000000000000000000000000000081529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b15801561021857600080fd5b505af115801561022c573d6000803e3d6000fd5b505050506040513d602081101561024257600080fd5b5051604080517f8cfb8f21000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b1580156102a957600080fd5b505af11580156102bd573d6000803e3d6000fd5b505050506040513d60208110156102d357600080fd5b505115156102e057600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561033b57600080fd5b505af115801561034f573d6000803e3d6000fd5b505050506040513d602081101561036557600080fd5b5051151561037257600080fd5b61037a610741565b5030600160a060020a031663f7938328338e8e8e8e8e8e8e8e6040518a63ffffffff1660e060020a028152600401808a600160a060020a0316600160a060020a031681526020018960018111156103cd57fe5b60ff168152602080820199909952604080820198909852600160a060020a03909616606087015250608085019390935260a084019190915260c083015260e08201529051610100808301955092935091908290030181600087803b15801561043457600080fd5b505af1158015610448573d6000803e3d6000fd5b505050506040513d602081101561045e57600080fd5b5051604080517fa0695f240000000000000000000000000000000000000000000000000000000081529051919450600160a060020a038b169163a0695f24916004808201926020929091908290030181600087803b1580156104bf57600080fd5b505af11580156104d3573d6000803e3d6000fd5b505050506040513d60208110156104e957600080fd5b509293508392506104f8610895565b5050505098975050505050505050565b60008054600160a060020a0316331461052057600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03831617905560015b919050565b600061055b611aa1565b60008054604080517f3f08882f0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921692633f08882f926024808401936020939083900390910190829087803b1580156105c257600080fd5b505af11580156105d6573d6000803e3d6000fd5b505050506040513d60208110156105ec57600080fd5b505115156105f957600080fd5b60005474010000000000000000000000000000000000000000900460ff161561062157600080fd5b6000805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179081905561067290600160a060020a03168c888d8d8d8d8c8c610bad565b905061067d81610e5e565b508051600160a060020a0316630e0a0d7461069783610ea9565b6040805160e060020a63ffffffff851602815260048101929092525160248083019260209291908290030181600087803b1580156106d457600080fd5b505af11580156106e8573d6000803e3d6000fd5b505050506040513d60208110156106fe57600080fd5b50511561070a57600080fd5b6107148184611077565b6000805474ff0000000000000000000000000000000000000000191690559b9a5050505050505050505050565b60008034111561088f5760008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f436173680000000000000000000000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b1580156107d257600080fd5b505af11580156107e6573d6000803e3d6000fd5b505050506040513d60208110156107fc57600080fd5b5051604080517f4faa8a260000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691634faa8a26913491602480830192602092919082900301818588803b15801561086157600080fd5b505af1158015610875573d6000803e3d6000fd5b50505050506040513d602081101561088c57600080fd5b50505b50600190565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f436173680000000000000000000000000000000000000000000000000000000060048201529051839283928392600160a060020a039092169163f39ec1f79160248082019260209290919082900301818787803b15801561092157600080fd5b505af1158015610935573d6000803e3d6000fd5b505050506040513d602081101561094b57600080fd5b5051604080517f70a082310000000000000000000000000000000000000000000000000000000081523360048201529051919450600160a060020a038516916370a08231916024808201926020929091908290030181600087803b1580156109b257600080fd5b505af11580156109c6573d6000803e3d6000fd5b505050506040513d60208110156109dc57600080fd5b505191506000821115610ba3576000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610a3b57600080fd5b505af1158015610a4f573d6000803e3d6000fd5b505050506040513d6020811015610a6557600080fd5b5051604080517fec238994000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301523360248301523060448301526064820186905291519293509083169163ec238994916084808201926020929091908290030181600087803b158015610ae257600080fd5b505af1158015610af6573d6000803e3d6000fd5b505050506040513d6020811015610b0c57600080fd5b5050604080517f1baffe38000000000000000000000000000000000000000000000000000000008152336004820152602481018490529051600160a060020a03851691631baffe389160448083019260209291908290030181600087803b158015610b7657600080fd5b505af1158015610b8a573d6000803e3d6000fd5b505050506040513d6020811015610ba057600080fd5b50505b6001935050505090565b610bb5611aa1565b60008085600160a060020a03166327ce5b8c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610bf657600080fd5b505af1158015610c0a573d6000803e3d6000fd5b505050506040513d6020811015610c2057600080fd5b50518a10610c2d57600080fd5b85600160a060020a031663bad84c9e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610c6b57600080fd5b505af1158015610c7f573d6000803e3d6000fd5b505050506040513d6020811015610c9557600080fd5b50518710610ca257600080fd5b60008811610caf57600080fd5b8b600160a060020a031663f39ec1f76040518163ffffffff1660e060020a02815260040180807f4f726465727300000000000000000000000000000000000000000000000000008152506020019050602060405180830381600087803b158015610d1857600080fd5b505af1158015610d2c573d6000803e3d6000fd5b505050506040513d6020811015610d4257600080fd5b5051604080517f4e94c8290000000000000000000000000000000000000000000000000000000081529051919350600160a060020a038e1691634e94c829916004808201926020929091908290030181600087803b158015610da357600080fd5b505af1158015610db7573d6000803e3d6000fd5b505050506040513d6020811015610dcd57600080fd5b5051604080516101a081018252600160a060020a03808616825289811660208301528084169282019290925260006060820152908d16608082015260a081018c905290915060c081018a6001811115610e2257fe5b8152602081018a905260408101899052600060608201819052608082015260a0810187905260c001859052925050509998505050505050505050565b600060018260c001516001811115610e7257fe5b1415610e8857610e818261119c565b905061054c565b60008260c001516001811115610e9a57fe5b141561054c57610e818261157e565b60608101516000908190151561106e578260000151600160a060020a0316631727583c8460c0015185602001518660e001518761010001518860800151438a60a001518b61014001518c61012001516040518a63ffffffff1660e060020a028152600401808a6001811115610f1a57fe5b60ff16815260200189600160a060020a0316600160a060020a0316815260200188815260200187815260200186600160a060020a0316600160a060020a031681526020018581526020018481526020018381526020018281526020019950505050505050505050602060405180830381600087803b158015610f9b57600080fd5b505af1158015610faf573d6000803e3d6000fd5b505050506040513d6020811015610fc557600080fd5b50518351604080517f0e0a0d74000000000000000000000000000000000000000000000000000000008152600481018490529051929350600160a060020a0390911691630e0a0d74916024808201926020929091908290030181600087803b15801561103057600080fd5b505af1158015611044573d6000803e3d6000fd5b505050506040513d602081101561105a57600080fd5b50511561106657600080fd5b606083018190525b50506060015190565b60008260000151600160a060020a0316636bc29efa8460c0015185602001518660e0015187610100015188608001518960a001518a61014001518b61012001518c61016001518d61018001518d6040518c63ffffffff1660e060020a028152600401808c60018111156110e657fe5b60ff168152600160a060020a039b8c1660208083019190915260408083019c909c52606082019a909a5297909a16608088015260a087019590955260c086019390935260e085019190915261010084015261012083015261014082015291516101608084019550919350918290030181600087803b15801561116757600080fd5b505af115801561117b573d6000803e3d6000fd5b505050506040513d602081101561119157600080fd5b505190505b92915050565b60008060008084610140015160001415156111b657600080fd5b610120850151156111c657600080fd5b8460200151600160a060020a03166365957bf58660a001516040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561121757600080fd5b505af115801561122b573d6000803e3d6000fd5b505050506040513d602081101561124157600080fd5b505160e08601516080870151604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015290519396509194508516916370a08231916024808201926020929091908290030181600087803b1580156112b757600080fd5b505af11580156112cb573d6000803e3d6000fd5b505050506040513d60208110156112e157600080fd5b5051905060008111156113ac576112f88183611a47565b61012086018190526080860151602080880151604080517f764c92f2000000000000000000000000000000000000000000000000000000008152600160a060020a0394851660048201529184166024830152604482018590525193909503949186169263764c92f29260648082019392918290030181600087803b15801561137f57600080fd5b505af1158015611393573d6000803e3d6000fd5b505050506040513d60208110156113a957600080fd5b50505b6000821115611573576114488261143c8761010001518860200151600160a060020a031663bad84c9e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561140457600080fd5b505af1158015611418573d6000803e3d6000fd5b505050506040513d602081101561142e57600080fd5b50519063ffffffff611a5e16565b9063ffffffff611a7316565b856101400181815250508460400151600160a060020a031663ec2389948660200151600160a060020a031663df2a29da6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156114a757600080fd5b505af11580156114bb573d6000803e3d6000fd5b505050506040513d60208110156114d157600080fd5b505160808801516020898101516101408b01516040805160e060020a63ffffffff8916028152600160a060020a039687166004820152948616602486015294909116604484015260648301529151608480830193928290030181600087803b15801561153c57600080fd5b505af1158015611550573d6000803e3d6000fd5b505050506040513d602081101561156657600080fd5b5051151561157357600080fd5b506001949350505050565b600080600080600080866101400151600014151561159b57600080fd5b610120870151156115ab57600080fd5b8660e0015194508660200151600160a060020a03166327ce5b8c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156115f457600080fd5b505af1158015611608573d6000803e3d6000fd5b505050506040513d602081101561161e57600080fd5b505193507f40000000000000000000000000000000000000000000000000000000000000009250600091505b838210156117865760a0870151821461177b578660200151600160a060020a03166365957bf5836040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b1580156116aa57600080fd5b505af11580156116be573d6000803e3d6000fd5b505050506040513d60208110156116d457600080fd5b50516080880151604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152905191909216916370a082319160248083019260209291908290030181600087803b15801561174057600080fd5b505af1158015611754573d6000803e3d6000fd5b505050506040513d602081101561176a57600080fd5b505190506117788184611a47565b92505b60019091019061164a565b60008311156118ef576117998386611a47565b610120880181905290940393600091505b838210156118ef5760a087015182146118e4578660200151600160a060020a03166365957bf5836040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561180a57600080fd5b505af115801561181e573d6000803e3d6000fd5b505050506040513d602081101561183457600080fd5b505160808801516020898101516101208b0151604080517f764c92f2000000000000000000000000000000000000000000000000000000008152600160a060020a0395861660048201529285166024840152604483019190915251929093169263764c92f292606480830193928290030181600087803b1580156118b757600080fd5b505af11580156118cb573d6000803e3d6000fd5b505050506040513d60208110156118e157600080fd5b50505b6001909101906117aa565b6000851115611a3a5761010087015161190f90869063ffffffff611a7316565b876101400181815250508660400151600160a060020a031663ec2389948860200151600160a060020a031663df2a29da6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561196e57600080fd5b505af1158015611982573d6000803e3d6000fd5b505050506040513d602081101561199857600080fd5b505160808a015160208b8101516101408d01516040805160e060020a63ffffffff8916028152600160a060020a039687166004820152948616602486015294909116604484015260648301529151608480830193928290030181600087803b158015611a0357600080fd5b505af1158015611a17573d6000803e3d6000fd5b505050506040513d6020811015611a2d57600080fd5b50511515611a3a57600080fd5b5060019695505050505050565b6000818311611a57575081611196565b5080611196565b600082821115611a6d57600080fd5b50900390565b6000828202831580611a8f5750828482811515611a8c57fe5b04145b1515611a9a57600080fd5b9392505050565b604080516101a081018252600080825260208201819052918101829052606081018290526080810182905260a081018290529060c0820190815260006020820181905260408201819052606082018190526080820181905260a0820181905260c090910152905600a165627a7a72305820e56c41af3106f13d22095d1fccd9ac1b857c4075870667b014a6a1682334814b0029`

// DeployCreateOrder deploys a new Ethereum contract, binding an instance of CreateOrder to it.
func DeployCreateOrder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CreateOrder, error) {
	parsed, err := abi.JSON(strings.NewReader(CreateOrderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CreateOrderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CreateOrder{CreateOrderCaller: CreateOrderCaller{contract: contract}, CreateOrderTransactor: CreateOrderTransactor{contract: contract}, CreateOrderFilterer: CreateOrderFilterer{contract: contract}}, nil
}

// CreateOrder is an auto generated Go binding around an Ethereum contract.
type CreateOrder struct {
	CreateOrderCaller     // Read-only binding to the contract
	CreateOrderTransactor // Write-only binding to the contract
	CreateOrderFilterer   // Log filterer for contract events
}

// CreateOrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreateOrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateOrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreateOrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateOrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreateOrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateOrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreateOrderSession struct {
	Contract     *CreateOrder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreateOrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreateOrderCallerSession struct {
	Contract *CreateOrderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CreateOrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreateOrderTransactorSession struct {
	Contract     *CreateOrderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CreateOrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreateOrderRaw struct {
	Contract *CreateOrder // Generic contract binding to access the raw methods on
}

// CreateOrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreateOrderCallerRaw struct {
	Contract *CreateOrderCaller // Generic read-only contract binding to access the raw methods on
}

// CreateOrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreateOrderTransactorRaw struct {
	Contract *CreateOrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreateOrder creates a new instance of CreateOrder, bound to a specific deployed contract.
func NewCreateOrder(address common.Address, backend bind.ContractBackend) (*CreateOrder, error) {
	contract, err := bindCreateOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreateOrder{CreateOrderCaller: CreateOrderCaller{contract: contract}, CreateOrderTransactor: CreateOrderTransactor{contract: contract}, CreateOrderFilterer: CreateOrderFilterer{contract: contract}}, nil
}

// NewCreateOrderCaller creates a new read-only instance of CreateOrder, bound to a specific deployed contract.
func NewCreateOrderCaller(address common.Address, caller bind.ContractCaller) (*CreateOrderCaller, error) {
	contract, err := bindCreateOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreateOrderCaller{contract: contract}, nil
}

// NewCreateOrderTransactor creates a new write-only instance of CreateOrder, bound to a specific deployed contract.
func NewCreateOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*CreateOrderTransactor, error) {
	contract, err := bindCreateOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreateOrderTransactor{contract: contract}, nil
}

// NewCreateOrderFilterer creates a new log filterer instance of CreateOrder, bound to a specific deployed contract.
func NewCreateOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*CreateOrderFilterer, error) {
	contract, err := bindCreateOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreateOrderFilterer{contract: contract}, nil
}

// bindCreateOrder binds a generic wrapper to an already deployed contract.
func bindCreateOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreateOrderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreateOrder *CreateOrderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CreateOrder.Contract.CreateOrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreateOrder *CreateOrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreateOrder.Contract.CreateOrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreateOrder *CreateOrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreateOrder.Contract.CreateOrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreateOrder *CreateOrderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CreateOrder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreateOrder *CreateOrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreateOrder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreateOrder *CreateOrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreateOrder.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_CreateOrder *CreateOrderCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CreateOrder.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_CreateOrder *CreateOrderSession) GetController() (common.Address, error) {
	return _CreateOrder.Contract.GetController(&_CreateOrder.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_CreateOrder *CreateOrderCallerSession) GetController() (common.Address, error) {
	return _CreateOrder.Contract.GetController(&_CreateOrder.CallOpts)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf7938328.
//
// Solidity: function createOrder(_creator address, _type uint8, _attoshares uint256, _displayPrice uint256, _market address, _outcome uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_CreateOrder *CreateOrderTransactor) CreateOrder(opts *bind.TransactOpts, _creator common.Address, _type uint8, _attoshares *big.Int, _displayPrice *big.Int, _market common.Address, _outcome *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _CreateOrder.contract.Transact(opts, "createOrder", _creator, _type, _attoshares, _displayPrice, _market, _outcome, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf7938328.
//
// Solidity: function createOrder(_creator address, _type uint8, _attoshares uint256, _displayPrice uint256, _market address, _outcome uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_CreateOrder *CreateOrderSession) CreateOrder(_creator common.Address, _type uint8, _attoshares *big.Int, _displayPrice *big.Int, _market common.Address, _outcome *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _CreateOrder.Contract.CreateOrder(&_CreateOrder.TransactOpts, _creator, _type, _attoshares, _displayPrice, _market, _outcome, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf7938328.
//
// Solidity: function createOrder(_creator address, _type uint8, _attoshares uint256, _displayPrice uint256, _market address, _outcome uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_CreateOrder *CreateOrderTransactorSession) CreateOrder(_creator common.Address, _type uint8, _attoshares *big.Int, _displayPrice *big.Int, _market common.Address, _outcome *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _CreateOrder.Contract.CreateOrder(&_CreateOrder.TransactOpts, _creator, _type, _attoshares, _displayPrice, _market, _outcome, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicCreateOrder is a paid mutator transaction binding the contract method 0x3ae4ce0a.
//
// Solidity: function publicCreateOrder(_type uint8, _attoshares uint256, _displayPrice uint256, _market address, _outcome uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_CreateOrder *CreateOrderTransactor) PublicCreateOrder(opts *bind.TransactOpts, _type uint8, _attoshares *big.Int, _displayPrice *big.Int, _market common.Address, _outcome *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _CreateOrder.contract.Transact(opts, "publicCreateOrder", _type, _attoshares, _displayPrice, _market, _outcome, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicCreateOrder is a paid mutator transaction binding the contract method 0x3ae4ce0a.
//
// Solidity: function publicCreateOrder(_type uint8, _attoshares uint256, _displayPrice uint256, _market address, _outcome uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_CreateOrder *CreateOrderSession) PublicCreateOrder(_type uint8, _attoshares *big.Int, _displayPrice *big.Int, _market common.Address, _outcome *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _CreateOrder.Contract.PublicCreateOrder(&_CreateOrder.TransactOpts, _type, _attoshares, _displayPrice, _market, _outcome, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicCreateOrder is a paid mutator transaction binding the contract method 0x3ae4ce0a.
//
// Solidity: function publicCreateOrder(_type uint8, _attoshares uint256, _displayPrice uint256, _market address, _outcome uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_CreateOrder *CreateOrderTransactorSession) PublicCreateOrder(_type uint8, _attoshares *big.Int, _displayPrice *big.Int, _market common.Address, _outcome *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _CreateOrder.Contract.PublicCreateOrder(&_CreateOrder.TransactOpts, _type, _attoshares, _displayPrice, _market, _outcome, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_CreateOrder *CreateOrderTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _CreateOrder.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_CreateOrder *CreateOrderSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _CreateOrder.Contract.SetController(&_CreateOrder.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_CreateOrder *CreateOrderTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _CreateOrder.Contract.SetController(&_CreateOrder.TransactOpts, _controller)
}

// ICreateOrderABI is the input ABI used to generate the binding from.
const ICreateOrderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"publicCreateOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"createOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ICreateOrderBin is the compiled bytecode used for deploying new contracts.
const ICreateOrderBin = `0x`

// DeployICreateOrder deploys a new Ethereum contract, binding an instance of ICreateOrder to it.
func DeployICreateOrder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ICreateOrder, error) {
	parsed, err := abi.JSON(strings.NewReader(ICreateOrderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ICreateOrderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ICreateOrder{ICreateOrderCaller: ICreateOrderCaller{contract: contract}, ICreateOrderTransactor: ICreateOrderTransactor{contract: contract}, ICreateOrderFilterer: ICreateOrderFilterer{contract: contract}}, nil
}

// ICreateOrder is an auto generated Go binding around an Ethereum contract.
type ICreateOrder struct {
	ICreateOrderCaller     // Read-only binding to the contract
	ICreateOrderTransactor // Write-only binding to the contract
	ICreateOrderFilterer   // Log filterer for contract events
}

// ICreateOrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICreateOrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreateOrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICreateOrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreateOrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICreateOrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreateOrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICreateOrderSession struct {
	Contract     *ICreateOrder     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICreateOrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICreateOrderCallerSession struct {
	Contract *ICreateOrderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ICreateOrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICreateOrderTransactorSession struct {
	Contract     *ICreateOrderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ICreateOrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICreateOrderRaw struct {
	Contract *ICreateOrder // Generic contract binding to access the raw methods on
}

// ICreateOrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICreateOrderCallerRaw struct {
	Contract *ICreateOrderCaller // Generic read-only contract binding to access the raw methods on
}

// ICreateOrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICreateOrderTransactorRaw struct {
	Contract *ICreateOrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICreateOrder creates a new instance of ICreateOrder, bound to a specific deployed contract.
func NewICreateOrder(address common.Address, backend bind.ContractBackend) (*ICreateOrder, error) {
	contract, err := bindICreateOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICreateOrder{ICreateOrderCaller: ICreateOrderCaller{contract: contract}, ICreateOrderTransactor: ICreateOrderTransactor{contract: contract}, ICreateOrderFilterer: ICreateOrderFilterer{contract: contract}}, nil
}

// NewICreateOrderCaller creates a new read-only instance of ICreateOrder, bound to a specific deployed contract.
func NewICreateOrderCaller(address common.Address, caller bind.ContractCaller) (*ICreateOrderCaller, error) {
	contract, err := bindICreateOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICreateOrderCaller{contract: contract}, nil
}

// NewICreateOrderTransactor creates a new write-only instance of ICreateOrder, bound to a specific deployed contract.
func NewICreateOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*ICreateOrderTransactor, error) {
	contract, err := bindICreateOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICreateOrderTransactor{contract: contract}, nil
}

// NewICreateOrderFilterer creates a new log filterer instance of ICreateOrder, bound to a specific deployed contract.
func NewICreateOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*ICreateOrderFilterer, error) {
	contract, err := bindICreateOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICreateOrderFilterer{contract: contract}, nil
}

// bindICreateOrder binds a generic wrapper to an already deployed contract.
func bindICreateOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICreateOrderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreateOrder *ICreateOrderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICreateOrder.Contract.ICreateOrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreateOrder *ICreateOrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreateOrder.Contract.ICreateOrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreateOrder *ICreateOrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreateOrder.Contract.ICreateOrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreateOrder *ICreateOrderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICreateOrder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreateOrder *ICreateOrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreateOrder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreateOrder *ICreateOrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreateOrder.Contract.contract.Transact(opts, method, params...)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf7938328.
//
// Solidity: function createOrder( address,  uint8,  uint256,  uint256,  address,  uint256,  bytes32,  bytes32,  bytes32) returns(bytes32)
func (_ICreateOrder *ICreateOrderTransactor) CreateOrder(opts *bind.TransactOpts, arg0 common.Address, arg1 uint8, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 [32]byte, arg7 [32]byte, arg8 [32]byte) (*types.Transaction, error) {
	return _ICreateOrder.contract.Transact(opts, "createOrder", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf7938328.
//
// Solidity: function createOrder( address,  uint8,  uint256,  uint256,  address,  uint256,  bytes32,  bytes32,  bytes32) returns(bytes32)
func (_ICreateOrder *ICreateOrderSession) CreateOrder(arg0 common.Address, arg1 uint8, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 [32]byte, arg7 [32]byte, arg8 [32]byte) (*types.Transaction, error) {
	return _ICreateOrder.Contract.CreateOrder(&_ICreateOrder.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xf7938328.
//
// Solidity: function createOrder( address,  uint8,  uint256,  uint256,  address,  uint256,  bytes32,  bytes32,  bytes32) returns(bytes32)
func (_ICreateOrder *ICreateOrderTransactorSession) CreateOrder(arg0 common.Address, arg1 uint8, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 [32]byte, arg7 [32]byte, arg8 [32]byte) (*types.Transaction, error) {
	return _ICreateOrder.Contract.CreateOrder(&_ICreateOrder.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// PublicCreateOrder is a paid mutator transaction binding the contract method 0x3ae4ce0a.
//
// Solidity: function publicCreateOrder( uint8,  uint256,  uint256,  address,  uint256,  bytes32,  bytes32,  bytes32) returns(bytes32)
func (_ICreateOrder *ICreateOrderTransactor) PublicCreateOrder(opts *bind.TransactOpts, arg0 uint8, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 [32]byte, arg6 [32]byte, arg7 [32]byte) (*types.Transaction, error) {
	return _ICreateOrder.contract.Transact(opts, "publicCreateOrder", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// PublicCreateOrder is a paid mutator transaction binding the contract method 0x3ae4ce0a.
//
// Solidity: function publicCreateOrder( uint8,  uint256,  uint256,  address,  uint256,  bytes32,  bytes32,  bytes32) returns(bytes32)
func (_ICreateOrder *ICreateOrderSession) PublicCreateOrder(arg0 uint8, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 [32]byte, arg6 [32]byte, arg7 [32]byte) (*types.Transaction, error) {
	return _ICreateOrder.Contract.PublicCreateOrder(&_ICreateOrder.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// PublicCreateOrder is a paid mutator transaction binding the contract method 0x3ae4ce0a.
//
// Solidity: function publicCreateOrder( uint8,  uint256,  uint256,  address,  uint256,  bytes32,  bytes32,  bytes32) returns(bytes32)
func (_ICreateOrder *ICreateOrderTransactorSession) PublicCreateOrder(arg0 uint8, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 [32]byte, arg6 [32]byte, arg7 [32]byte) (*types.Transaction, error) {
	return _ICreateOrder.Contract.PublicCreateOrder(&_ICreateOrder.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}
