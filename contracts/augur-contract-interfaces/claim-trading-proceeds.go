package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ClaimTradingProceedsABI is the input ABI used to generate the binding from.
const ClaimTradingProceedsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"calculateCreatorFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"calculateReportingFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_numberOfShares\",\"type\":\"uint256\"}],\"name\":\"divideUpWinnings\",\"outputs\":[{\"name\":\"_proceeds\",\"type\":\"uint256\"},{\"name\":\"_shareHolderShare\",\"type\":\"uint256\"},{\"name\":\"_creatorShare\",\"type\":\"uint256\"},{\"name\":\"_reporterShare\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_shareHolder\",\"type\":\"address\"}],\"name\":\"claimTradingProceeds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_numberOfShares\",\"type\":\"uint256\"}],\"name\":\"calculateProceeds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ClaimTradingProceedsBin is the compiled bytecode used for deploying new contracts.
const ClaimTradingProceedsBin = `0x608060405260008054600160a860020a0319163317905561113a806100256000396000f3006080604052600436106100695763ffffffff60e060020a60003504166306065ced811461006e5780633018205f146100a457806381894407146100d557806392eefe9b146100f95780639f66cddf1461012e578063b53709af1461017b578063f2dc8266146101a2575b600080fd5b34801561007a57600080fd5b50610092600160a060020a03600435166024356101c9565b60408051918252519081900360200190f35b3480156100b057600080fd5b506100b9610247565b60408051600160a060020a039092168252519081900360200190f35b3480156100e157600080fd5b50610092600160a060020a0360043516602435610256565b34801561010557600080fd5b5061011a600160a060020a0360043516610366565b604080519115158252519081900360200190f35b34801561013a57600080fd5b50610155600160a060020a03600435166024356044356103b0565b604080519485526020850193909352838301919091526060830152519081900360800190f35b34801561018757600080fd5b5061011a600160a060020a0360043581169060243516610406565b3480156101ae57600080fd5b50610092600160a060020a0360043516602435604435610e31565b600082600160a060020a031663f8c52125836040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561021457600080fd5b505af1158015610228573d6000803e3d6000fd5b505050506040513d602081101561023e57600080fd5b50519392505050565b600054600160a060020a031690565b60008083600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561029757600080fd5b505af11580156102ab573d6000803e3d6000fd5b505050506040513d60208110156102c157600080fd5b5051604080517f8f93bffe0000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921691638f93bffe916004808201926020929091908290030181600087803b15801561032057600080fd5b505af1158015610334573d6000803e3d6000fd5b505050506040513d602081101561034a57600080fd5b5051905061035e838263ffffffff610ec416565b949350505050565b60008054600160a060020a0316331461037e57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b6000806000806103c1878787610e31565b93506103cd87856101c9565b91506103d98785610256565b90506103fb816103ef868563ffffffff610edb16565b9063ffffffff610edb16565b925093509350935093565b60008060008060008060008060008a600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561045557600080fd5b505af1158015610469573d6000803e3d6000fd5b505050506040513d602081101561047f57600080fd5b505160008054604080517f4e94c8290000000000000000000000000000000000000000000000000000000081529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b1580156104e657600080fd5b505af11580156104fa573d6000803e3d6000fd5b505050506040513d602081101561051057600080fd5b5051604080517f8cfb8f21000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b15801561057757600080fd5b505af115801561058b573d6000803e3d6000fd5b505050506040513d60208110156105a157600080fd5b505115156105ae57600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561060957600080fd5b505af115801561061d573d6000803e3d6000fd5b505050506040513d602081101561063357600080fd5b5051151561064057600080fd5b60005474010000000000000000000000000000000000000000900460ff161561066857600080fd5b6001600060146101000a81548160ff0219169083151502179055508c600160a060020a0316638d4e40836040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156106c157600080fd5b505af11580156106d5573d6000803e3d6000fd5b505050506040513d60208110156106eb57600080fd5b505115156106f857600080fd5b8c600160a060020a031663df2a29da6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561073657600080fd5b505af115801561074a573d6000803e3d6000fd5b505050506040513d602081101561076057600080fd5b50519950600098505b8c600160a060020a03166327ce5b8c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156107a757600080fd5b505af11580156107bb573d6000803e3d6000fd5b505050506040513d60208110156107d157600080fd5b5051891015610d96578c600160a060020a03166365957bf58a6040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561082357600080fd5b505af1158015610837573d6000803e3d6000fd5b505050506040513d602081101561084d57600080fd5b5051604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038f811660048301529151929a50908a16916370a08231916024808201926020929091908290030181600087803b1580156108b757600080fd5b505af11580156108cb573d6000803e3d6000fd5b505050506040513d60208110156108e157600080fd5b505196506108f08d8a896103b0565b9298509096509450925060008711156109a25787600160a060020a031663d333d7cf8d896040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561096657600080fd5b505af115801561097a573d6000803e3d6000fd5b505050506040513d602081101561099057600080fd5b506109a090508d898e8a89610ef0565b505b6000851115610ae457604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a038f81166004830152306024830152604482018890529151918c16916323b872dd916064808201926020929091908290030181600087803b158015610a1d57600080fd5b505af1158015610a31573d6000803e3d6000fd5b505050506040513d6020811015610a4757600080fd5b50511515610a5457600080fd5b89600160a060020a0316631baffe388d876040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015610ab757600080fd5b505af1158015610acb573d6000803e3d6000fd5b505050506040513d6020811015610ae157600080fd5b50505b6000841115610bf35789600160a060020a03166323b872dd8e8f600160a060020a031663ed23378b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610b3b57600080fd5b505af1158015610b4f573d6000803e3d6000fd5b505050506040513d6020811015610b6557600080fd5b50516040805160e060020a63ffffffff8616028152600160a060020a039384166004820152929091166024830152604482018890525160648083019260209291908290030181600087803b158015610bbc57600080fd5b505af1158015610bd0573d6000803e3d6000fd5b505050506040513d6020811015610be657600080fd5b50511515610bf357600080fd5b6000831115610d8b5789600160a060020a03166323b872dd8e8f600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610c4a57600080fd5b505af1158015610c5e573d6000803e3d6000fd5b505050506040513d6020811015610c7457600080fd5b5051604080517f0cc8c9af0000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921691630cc8c9af916004808201926020929091908290030181600087803b158015610cd357600080fd5b505af1158015610ce7573d6000803e3d6000fd5b505050506040513d6020811015610cfd57600080fd5b50516040805160e060020a63ffffffff8616028152600160a060020a039384166004820152929091166024830152604482018790525160648083019260209291908290030181600087803b158015610d5457600080fd5b505af1158015610d68573d6000803e3d6000fd5b505050506040513d6020811015610d7e57600080fd5b50511515610d8b57600080fd5b886001019850610769565b8c600160a060020a031663a0695f246040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610dd457600080fd5b505af1158015610de8573d6000803e3d6000fd5b505050506040513d6020811015610dfe57600080fd5b50506000805474ff0000000000000000000000000000000000000000191690555060019c9b505050505050505050505050565b60008084600160a060020a0316633c264820856040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b158015610e7d57600080fd5b505af1158015610e91573d6000803e3d6000fd5b505050506040513d6020811015610ea757600080fd5b50519050610ebb838263ffffffff6110ce16565b95945050505050565b6000808284811515610ed257fe5b04949350505050565b600082821115610eea57600080fd5b50900390565b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610f4457600080fd5b505af1158015610f58573d6000803e3d6000fd5b505050506040513d6020811015610f6e57600080fd5b5051604080517f870c426d0000000000000000000000000000000000000000000000000000000081529051600160a060020a0392831692636051fa2c92908a169163870c426d916004808201926020929091908290030181600087803b158015610fd757600080fd5b505af1158015610feb573d6000803e3d6000fd5b505050506040513d602081101561100157600080fd5b505187878a8888611022600160a060020a038516318263ffffffff6110fc16565b6040805160e060020a63ffffffff8b16028152600160a060020a03988916600482015296881660248801529487166044870152929095166064850152608484015260a483019390935260c4820192909252905160e48083019260209291908290030181600087803b15801561109657600080fd5b505af11580156110aa573d6000803e3d6000fd5b505050506040513d60208110156110c057600080fd5b506001979650505050505050565b60008282028315806110ea57508284828115156110e757fe5b04145b15156110f557600080fd5b9392505050565b6000828201838110156110f557600080fd00a165627a7a72305820adadebd6e2370e834fddf460a02c5bd4349d7688a7770f063443fbc92fcace480029`

// DeployClaimTradingProceeds deploys a new Ethereum contract, binding an instance of ClaimTradingProceeds to it.
func DeployClaimTradingProceeds(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ClaimTradingProceeds, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimTradingProceedsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClaimTradingProceedsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClaimTradingProceeds{ClaimTradingProceedsCaller: ClaimTradingProceedsCaller{contract: contract}, ClaimTradingProceedsTransactor: ClaimTradingProceedsTransactor{contract: contract}, ClaimTradingProceedsFilterer: ClaimTradingProceedsFilterer{contract: contract}}, nil
}

// ClaimTradingProceeds is an auto generated Go binding around an Ethereum contract.
type ClaimTradingProceeds struct {
	ClaimTradingProceedsCaller     // Read-only binding to the contract
	ClaimTradingProceedsTransactor // Write-only binding to the contract
	ClaimTradingProceedsFilterer   // Log filterer for contract events
}

// ClaimTradingProceedsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimTradingProceedsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTradingProceedsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimTradingProceedsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTradingProceedsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimTradingProceedsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTradingProceedsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimTradingProceedsSession struct {
	Contract     *ClaimTradingProceeds // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ClaimTradingProceedsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimTradingProceedsCallerSession struct {
	Contract *ClaimTradingProceedsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ClaimTradingProceedsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimTradingProceedsTransactorSession struct {
	Contract     *ClaimTradingProceedsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ClaimTradingProceedsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimTradingProceedsRaw struct {
	Contract *ClaimTradingProceeds // Generic contract binding to access the raw methods on
}

// ClaimTradingProceedsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimTradingProceedsCallerRaw struct {
	Contract *ClaimTradingProceedsCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimTradingProceedsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimTradingProceedsTransactorRaw struct {
	Contract *ClaimTradingProceedsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimTradingProceeds creates a new instance of ClaimTradingProceeds, bound to a specific deployed contract.
func NewClaimTradingProceeds(address common.Address, backend bind.ContractBackend) (*ClaimTradingProceeds, error) {
	contract, err := bindClaimTradingProceeds(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimTradingProceeds{ClaimTradingProceedsCaller: ClaimTradingProceedsCaller{contract: contract}, ClaimTradingProceedsTransactor: ClaimTradingProceedsTransactor{contract: contract}, ClaimTradingProceedsFilterer: ClaimTradingProceedsFilterer{contract: contract}}, nil
}

// NewClaimTradingProceedsCaller creates a new read-only instance of ClaimTradingProceeds, bound to a specific deployed contract.
func NewClaimTradingProceedsCaller(address common.Address, caller bind.ContractCaller) (*ClaimTradingProceedsCaller, error) {
	contract, err := bindClaimTradingProceeds(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimTradingProceedsCaller{contract: contract}, nil
}

// NewClaimTradingProceedsTransactor creates a new write-only instance of ClaimTradingProceeds, bound to a specific deployed contract.
func NewClaimTradingProceedsTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimTradingProceedsTransactor, error) {
	contract, err := bindClaimTradingProceeds(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimTradingProceedsTransactor{contract: contract}, nil
}

// NewClaimTradingProceedsFilterer creates a new log filterer instance of ClaimTradingProceeds, bound to a specific deployed contract.
func NewClaimTradingProceedsFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimTradingProceedsFilterer, error) {
	contract, err := bindClaimTradingProceeds(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimTradingProceedsFilterer{contract: contract}, nil
}

// bindClaimTradingProceeds binds a generic wrapper to an already deployed contract.
func bindClaimTradingProceeds(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimTradingProceedsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimTradingProceeds *ClaimTradingProceedsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimTradingProceeds.Contract.ClaimTradingProceedsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimTradingProceeds *ClaimTradingProceedsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.ClaimTradingProceedsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimTradingProceeds *ClaimTradingProceedsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.ClaimTradingProceedsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimTradingProceeds *ClaimTradingProceedsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimTradingProceeds.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.contract.Transact(opts, method, params...)
}

// CalculateCreatorFee is a free data retrieval call binding the contract method 0x06065ced.
//
// Solidity: function calculateCreatorFee(_market address, _amount uint256) constant returns(uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsCaller) CalculateCreatorFee(opts *bind.CallOpts, _market common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClaimTradingProceeds.contract.Call(opts, out, "calculateCreatorFee", _market, _amount)
	return *ret0, err
}

// CalculateCreatorFee is a free data retrieval call binding the contract method 0x06065ced.
//
// Solidity: function calculateCreatorFee(_market address, _amount uint256) constant returns(uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsSession) CalculateCreatorFee(_market common.Address, _amount *big.Int) (*big.Int, error) {
	return _ClaimTradingProceeds.Contract.CalculateCreatorFee(&_ClaimTradingProceeds.CallOpts, _market, _amount)
}

// CalculateCreatorFee is a free data retrieval call binding the contract method 0x06065ced.
//
// Solidity: function calculateCreatorFee(_market address, _amount uint256) constant returns(uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsCallerSession) CalculateCreatorFee(_market common.Address, _amount *big.Int) (*big.Int, error) {
	return _ClaimTradingProceeds.Contract.CalculateCreatorFee(&_ClaimTradingProceeds.CallOpts, _market, _amount)
}

// CalculateProceeds is a free data retrieval call binding the contract method 0xf2dc8266.
//
// Solidity: function calculateProceeds(_market address, _outcome uint256, _numberOfShares uint256) constant returns(uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsCaller) CalculateProceeds(opts *bind.CallOpts, _market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClaimTradingProceeds.contract.Call(opts, out, "calculateProceeds", _market, _outcome, _numberOfShares)
	return *ret0, err
}

// CalculateProceeds is a free data retrieval call binding the contract method 0xf2dc8266.
//
// Solidity: function calculateProceeds(_market address, _outcome uint256, _numberOfShares uint256) constant returns(uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsSession) CalculateProceeds(_market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*big.Int, error) {
	return _ClaimTradingProceeds.Contract.CalculateProceeds(&_ClaimTradingProceeds.CallOpts, _market, _outcome, _numberOfShares)
}

// CalculateProceeds is a free data retrieval call binding the contract method 0xf2dc8266.
//
// Solidity: function calculateProceeds(_market address, _outcome uint256, _numberOfShares uint256) constant returns(uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsCallerSession) CalculateProceeds(_market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*big.Int, error) {
	return _ClaimTradingProceeds.Contract.CalculateProceeds(&_ClaimTradingProceeds.CallOpts, _market, _outcome, _numberOfShares)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ClaimTradingProceeds *ClaimTradingProceedsCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClaimTradingProceeds.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ClaimTradingProceeds *ClaimTradingProceedsSession) GetController() (common.Address, error) {
	return _ClaimTradingProceeds.Contract.GetController(&_ClaimTradingProceeds.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_ClaimTradingProceeds *ClaimTradingProceedsCallerSession) GetController() (common.Address, error) {
	return _ClaimTradingProceeds.Contract.GetController(&_ClaimTradingProceeds.CallOpts)
}

// CalculateReportingFee is a paid mutator transaction binding the contract method 0x81894407.
//
// Solidity: function calculateReportingFee(_market address, _amount uint256) returns(uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactor) CalculateReportingFee(opts *bind.TransactOpts, _market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClaimTradingProceeds.contract.Transact(opts, "calculateReportingFee", _market, _amount)
}

// CalculateReportingFee is a paid mutator transaction binding the contract method 0x81894407.
//
// Solidity: function calculateReportingFee(_market address, _amount uint256) returns(uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsSession) CalculateReportingFee(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.CalculateReportingFee(&_ClaimTradingProceeds.TransactOpts, _market, _amount)
}

// CalculateReportingFee is a paid mutator transaction binding the contract method 0x81894407.
//
// Solidity: function calculateReportingFee(_market address, _amount uint256) returns(uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactorSession) CalculateReportingFee(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.CalculateReportingFee(&_ClaimTradingProceeds.TransactOpts, _market, _amount)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xb53709af.
//
// Solidity: function claimTradingProceeds(_market address, _shareHolder address) returns(bool)
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactor) ClaimTradingProceeds(opts *bind.TransactOpts, _market common.Address, _shareHolder common.Address) (*types.Transaction, error) {
	return _ClaimTradingProceeds.contract.Transact(opts, "claimTradingProceeds", _market, _shareHolder)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xb53709af.
//
// Solidity: function claimTradingProceeds(_market address, _shareHolder address) returns(bool)
func (_ClaimTradingProceeds *ClaimTradingProceedsSession) ClaimTradingProceeds(_market common.Address, _shareHolder common.Address) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.ClaimTradingProceeds(&_ClaimTradingProceeds.TransactOpts, _market, _shareHolder)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xb53709af.
//
// Solidity: function claimTradingProceeds(_market address, _shareHolder address) returns(bool)
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactorSession) ClaimTradingProceeds(_market common.Address, _shareHolder common.Address) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.ClaimTradingProceeds(&_ClaimTradingProceeds.TransactOpts, _market, _shareHolder)
}

// DivideUpWinnings is a paid mutator transaction binding the contract method 0x9f66cddf.
//
// Solidity: function divideUpWinnings(_market address, _outcome uint256, _numberOfShares uint256) returns(_proceeds uint256, _shareHolderShare uint256, _creatorShare uint256, _reporterShare uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactor) DivideUpWinnings(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*types.Transaction, error) {
	return _ClaimTradingProceeds.contract.Transact(opts, "divideUpWinnings", _market, _outcome, _numberOfShares)
}

// DivideUpWinnings is a paid mutator transaction binding the contract method 0x9f66cddf.
//
// Solidity: function divideUpWinnings(_market address, _outcome uint256, _numberOfShares uint256) returns(_proceeds uint256, _shareHolderShare uint256, _creatorShare uint256, _reporterShare uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsSession) DivideUpWinnings(_market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.DivideUpWinnings(&_ClaimTradingProceeds.TransactOpts, _market, _outcome, _numberOfShares)
}

// DivideUpWinnings is a paid mutator transaction binding the contract method 0x9f66cddf.
//
// Solidity: function divideUpWinnings(_market address, _outcome uint256, _numberOfShares uint256) returns(_proceeds uint256, _shareHolderShare uint256, _creatorShare uint256, _reporterShare uint256)
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactorSession) DivideUpWinnings(_market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.DivideUpWinnings(&_ClaimTradingProceeds.TransactOpts, _market, _outcome, _numberOfShares)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _ClaimTradingProceeds.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ClaimTradingProceeds *ClaimTradingProceedsSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.SetController(&_ClaimTradingProceeds.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_ClaimTradingProceeds *ClaimTradingProceedsTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ClaimTradingProceeds.Contract.SetController(&_ClaimTradingProceeds.TransactOpts, _controller)
}

// IClaimTradingProceedsABI is the input ABI used to generate the binding from.
const IClaimTradingProceedsABI = "[]"

// IClaimTradingProceedsBin is the compiled bytecode used for deploying new contracts.
const IClaimTradingProceedsBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a7230582087fbf0c1540e89774ac5ed156b0238aa5ec20cb40ddff291447351978c0fd7330029`

// DeployIClaimTradingProceeds deploys a new Ethereum contract, binding an instance of IClaimTradingProceeds to it.
func DeployIClaimTradingProceeds(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IClaimTradingProceeds, error) {
	parsed, err := abi.JSON(strings.NewReader(IClaimTradingProceedsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IClaimTradingProceedsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IClaimTradingProceeds{IClaimTradingProceedsCaller: IClaimTradingProceedsCaller{contract: contract}, IClaimTradingProceedsTransactor: IClaimTradingProceedsTransactor{contract: contract}, IClaimTradingProceedsFilterer: IClaimTradingProceedsFilterer{contract: contract}}, nil
}

// IClaimTradingProceeds is an auto generated Go binding around an Ethereum contract.
type IClaimTradingProceeds struct {
	IClaimTradingProceedsCaller     // Read-only binding to the contract
	IClaimTradingProceedsTransactor // Write-only binding to the contract
	IClaimTradingProceedsFilterer   // Log filterer for contract events
}

// IClaimTradingProceedsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IClaimTradingProceedsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IClaimTradingProceedsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IClaimTradingProceedsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IClaimTradingProceedsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IClaimTradingProceedsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IClaimTradingProceedsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IClaimTradingProceedsSession struct {
	Contract     *IClaimTradingProceeds // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IClaimTradingProceedsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IClaimTradingProceedsCallerSession struct {
	Contract *IClaimTradingProceedsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IClaimTradingProceedsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IClaimTradingProceedsTransactorSession struct {
	Contract     *IClaimTradingProceedsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IClaimTradingProceedsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IClaimTradingProceedsRaw struct {
	Contract *IClaimTradingProceeds // Generic contract binding to access the raw methods on
}

// IClaimTradingProceedsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IClaimTradingProceedsCallerRaw struct {
	Contract *IClaimTradingProceedsCaller // Generic read-only contract binding to access the raw methods on
}

// IClaimTradingProceedsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IClaimTradingProceedsTransactorRaw struct {
	Contract *IClaimTradingProceedsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIClaimTradingProceeds creates a new instance of IClaimTradingProceeds, bound to a specific deployed contract.
func NewIClaimTradingProceeds(address common.Address, backend bind.ContractBackend) (*IClaimTradingProceeds, error) {
	contract, err := bindIClaimTradingProceeds(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IClaimTradingProceeds{IClaimTradingProceedsCaller: IClaimTradingProceedsCaller{contract: contract}, IClaimTradingProceedsTransactor: IClaimTradingProceedsTransactor{contract: contract}, IClaimTradingProceedsFilterer: IClaimTradingProceedsFilterer{contract: contract}}, nil
}

// NewIClaimTradingProceedsCaller creates a new read-only instance of IClaimTradingProceeds, bound to a specific deployed contract.
func NewIClaimTradingProceedsCaller(address common.Address, caller bind.ContractCaller) (*IClaimTradingProceedsCaller, error) {
	contract, err := bindIClaimTradingProceeds(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IClaimTradingProceedsCaller{contract: contract}, nil
}

// NewIClaimTradingProceedsTransactor creates a new write-only instance of IClaimTradingProceeds, bound to a specific deployed contract.
func NewIClaimTradingProceedsTransactor(address common.Address, transactor bind.ContractTransactor) (*IClaimTradingProceedsTransactor, error) {
	contract, err := bindIClaimTradingProceeds(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IClaimTradingProceedsTransactor{contract: contract}, nil
}

// NewIClaimTradingProceedsFilterer creates a new log filterer instance of IClaimTradingProceeds, bound to a specific deployed contract.
func NewIClaimTradingProceedsFilterer(address common.Address, filterer bind.ContractFilterer) (*IClaimTradingProceedsFilterer, error) {
	contract, err := bindIClaimTradingProceeds(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IClaimTradingProceedsFilterer{contract: contract}, nil
}

// bindIClaimTradingProceeds binds a generic wrapper to an already deployed contract.
func bindIClaimTradingProceeds(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IClaimTradingProceedsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IClaimTradingProceeds *IClaimTradingProceedsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IClaimTradingProceeds.Contract.IClaimTradingProceedsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IClaimTradingProceeds *IClaimTradingProceedsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IClaimTradingProceeds.Contract.IClaimTradingProceedsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IClaimTradingProceeds *IClaimTradingProceedsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IClaimTradingProceeds.Contract.IClaimTradingProceedsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IClaimTradingProceeds *IClaimTradingProceedsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IClaimTradingProceeds.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IClaimTradingProceeds *IClaimTradingProceedsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IClaimTradingProceeds.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IClaimTradingProceeds *IClaimTradingProceedsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IClaimTradingProceeds.Contract.contract.Transact(opts, method, params...)
}
