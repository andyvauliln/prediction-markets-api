package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// FeeTokenABI is the input ABI used to generate the binding from.
const FeeTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETERNAL_APPROVAL_VALUE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_feeWindow\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"feeWindowBurn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// FeeTokenBin is the compiled bytecode used for deploying new contracts.
const FeeTokenBin = `0x60806040526005805460ff1916905560008054600160a060020a03191633179055610e858061002f6000396000f3006080604052600436106100f85763ffffffff60e060020a60003504166306fdde0381146100fd578063095ea7b31461018757806318160ddd146101bf57806323b872dd146101e65780633018205f14610210578063313ce56714610241578063634eaff11461026c578063661884631461028157806370a08231146102a557806392eefe9b146102c657806395d89b41146102e7578063a00ce6a5146102fc578063a9059cbb14610320578063bef72fa214610344578063c4d66de814610359578063d73dd6231461037a578063dd62ed3e1461039e578063ee89dab4146103c5578063f77f29b1146103da578063f7862ec2146103ef575b600080fd5b34801561010957600080fd5b50610112610413565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014c578181015183820152602001610134565b50505050905090810190601f1680156101795780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019357600080fd5b506101ab600160a060020a036004351660243561044a565b604080519115158252519081900360200190f35b3480156101cb57600080fd5b506101d4610461565b60408051918252519081900360200190f35b3480156101f257600080fd5b506101ab600160a060020a0360043581169060243516604435610467565b34801561021c57600080fd5b506102256104e0565b60408051600160a060020a039092168252519081900360200190f35b34801561024d57600080fd5b506102566104ef565b6040805160ff9092168252519081900360200190f35b34801561027857600080fd5b506101d46104f4565b34801561028d57600080fd5b506101ab600160a060020a03600435166024356104fa565b3480156102b157600080fd5b506101d4600160a060020a036004351661055d565b3480156102d257600080fd5b506101ab600160a060020a0360043516610578565b3480156102f357600080fd5b506101126105c2565b34801561030857600080fd5b506101ab600160a060020a03600435166024356105f9565b34801561032c57600080fd5b506101ab600160a060020a0360043516602435610633565b34801561035057600080fd5b506101d4610647565b34801561036557600080fd5b506101ab600160a060020a036004351661064d565b34801561038657600080fd5b506101ab600160a060020a03600435166024356106a0565b3480156103aa57600080fd5b506101d4600160a060020a03600435811690602435166106dc565b3480156103d157600080fd5b506101ab610707565b3480156103e657600080fd5b50610225610710565b3480156103fb57600080fd5b506101ab600160a060020a0360043516602435610739565b60408051808201909152600881527f466565546f6b656e000000000000000000000000000000000000000000000000602082015281565b6000610457338484610773565b5060019392505050565b60025490565b600160a060020a038316600090815260046020908152604080832033845290915281205460001981146104c9576104a4818463ffffffff6107de16565b600160a060020a03861660009081526004602090815260408083203384529091529020555b6104d48585856107f3565b50600195945050505050565b600054600160a060020a031690565b600081565b60001981565b336000908152600460209081526040808320600160a060020a0386168452909152812054808311156105385761053233856000610773565b50610457565b610552338561054d848763ffffffff6107de16565b610773565b505060019392505050565b600160a060020a031660009081526003602052604090205490565b60008054600160a060020a0316331461059057600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60408051808201909152600381527f4645450000000000000000000000000000000000000000000000000000000000602082015281565b60055460009060ff16151561060d57600080fd5b6005546101009004600160a060020a0316331461062957600080fd5b61045783836108be565b60006106403384846107f3565b9392505050565b60015481565b60055460009060ff161561066057600080fd5b61066861095f565b505060058054600160a060020a0383166101000274ffffffffffffffffffffffffffffffffffffffff00199091161790556001919050565b336000818152600460209081526040808320600160a060020a0387168452909152812054909161045791859061054d908663ffffffff61098616565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205490565b60055460ff1690565b60055460009060ff16151561072457600080fd5b506005546101009004600160a060020a031690565b60055460009060ff16151561074d57600080fd5b6005546101009004600160a060020a0316331461076957600080fd5b6104578383610998565b600160a060020a03808416600081815260046020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6000828211156107ed57600080fd5b50900390565b600160a060020a03831660009081526003602052604081205461081c908363ffffffff6107de16565b600160a060020a038086166000908152600360205260408082209390935590851681522054610851908363ffffffff61098616565b600160a060020a0380851660008181526003602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a36108b3848484610a39565b506001949350505050565b600160a060020a0382166000908152600360205260408120546108e7908363ffffffff61098616565b600160a060020a038416600090815260036020526040902055600254610913908363ffffffff61098616565b600255604080518381529051600160a060020a038516917f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885919081900360200190a26104578383610bd7565b60055460009060ff161561097257600080fd5b506005805460ff1916600190811790915590565b60008282018381101561064057600080fd5b600160a060020a0382166000908152600360205260408120546109c1908363ffffffff6107de16565b600160a060020a0384166000908152600360205260409020556002546109ed908363ffffffff6107de16565b600255604080518381529051600160a060020a038516917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a26104578383610d6d565b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610a8d57600080fd5b505af1158015610aa1573d6000803e3d6000fd5b505050506040513d6020811015610ab757600080fd5b5051600554604080517f870c426d0000000000000000000000000000000000000000000000000000000081529051600160a060020a03938416936337227c07936101009004169163870c426d9160048083019260209291908290030181600087803b158015610b2557600080fd5b505af1158015610b39573d6000803e3d6000fd5b505050506040513d6020811015610b4f57600080fd5b50516040805160e060020a63ffffffff8516028152600160a060020a03928316600482015282891660248201529187166044830152606482018690525160848083019260209291908290030181600087803b158015610bad57600080fd5b505af1158015610bc1573d6000803e3d6000fd5b505050506040513d60208110156104d457600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610c2b57600080fd5b505af1158015610c3f573d6000803e3d6000fd5b505050506040513d6020811015610c5557600080fd5b5051600554604080517f870c426d0000000000000000000000000000000000000000000000000000000081529051600160a060020a0393841693632698eec9936101009004169163870c426d9160048083019260209291908290030181600087803b158015610cc357600080fd5b505af1158015610cd7573d6000803e3d6000fd5b505050506040513d6020811015610ced57600080fd5b50516040805160e060020a63ffffffff8516028152600160a060020a0392831660048201529187166024830152604482018690525160648083019260209291908290030181600087803b158015610d4357600080fd5b505af1158015610d57573d6000803e3d6000fd5b505050506040513d602081101561055257600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610dc157600080fd5b505af1158015610dd5573d6000803e3d6000fd5b505050506040513d6020811015610deb57600080fd5b5051600554604080517f870c426d0000000000000000000000000000000000000000000000000000000081529051600160a060020a039384169363979141ea936101009004169163870c426d9160048083019260209291908290030181600087803b158015610cc357600080fd00a165627a7a72305820117326a820a64319e388911152f83e538aeeaa49a0dc64f9858df12040ef3f8c0029`

// DeployFeeToken deploys a new Ethereum contract, binding an instance of FeeToken to it.
func DeployFeeToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FeeToken, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeeTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeToken{FeeTokenCaller: FeeTokenCaller{contract: contract}, FeeTokenTransactor: FeeTokenTransactor{contract: contract}, FeeTokenFilterer: FeeTokenFilterer{contract: contract}}, nil
}

// FeeToken is an auto generated Go binding around an Ethereum contract.
type FeeToken struct {
	FeeTokenCaller     // Read-only binding to the contract
	FeeTokenTransactor // Write-only binding to the contract
	FeeTokenFilterer   // Log filterer for contract events
}

// FeeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeTokenSession struct {
	Contract     *FeeToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeTokenCallerSession struct {
	Contract *FeeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FeeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeTokenTransactorSession struct {
	Contract     *FeeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FeeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeTokenRaw struct {
	Contract *FeeToken // Generic contract binding to access the raw methods on
}

// FeeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeTokenCallerRaw struct {
	Contract *FeeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FeeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeTokenTransactorRaw struct {
	Contract *FeeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeToken creates a new instance of FeeToken, bound to a specific deployed contract.
func NewFeeToken(address common.Address, backend bind.ContractBackend) (*FeeToken, error) {
	contract, err := bindFeeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeToken{FeeTokenCaller: FeeTokenCaller{contract: contract}, FeeTokenTransactor: FeeTokenTransactor{contract: contract}, FeeTokenFilterer: FeeTokenFilterer{contract: contract}}, nil
}

// NewFeeTokenCaller creates a new read-only instance of FeeToken, bound to a specific deployed contract.
func NewFeeTokenCaller(address common.Address, caller bind.ContractCaller) (*FeeTokenCaller, error) {
	contract, err := bindFeeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeTokenCaller{contract: contract}, nil
}

// NewFeeTokenTransactor creates a new write-only instance of FeeToken, bound to a specific deployed contract.
func NewFeeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeTokenTransactor, error) {
	contract, err := bindFeeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeTokenTransactor{contract: contract}, nil
}

// NewFeeTokenFilterer creates a new log filterer instance of FeeToken, bound to a specific deployed contract.
func NewFeeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeTokenFilterer, error) {
	contract, err := bindFeeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeTokenFilterer{contract: contract}, nil
}

// bindFeeToken binds a generic wrapper to an already deployed contract.
func bindFeeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeToken *FeeTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeToken.Contract.FeeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeToken *FeeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeToken.Contract.FeeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeToken *FeeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeToken.Contract.FeeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeToken *FeeTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeToken *FeeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeToken *FeeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeToken.Contract.contract.Transact(opts, method, params...)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_FeeToken *FeeTokenCaller) ETERNALAPPROVALVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "ETERNAL_APPROVAL_VALUE")
	return *ret0, err
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_FeeToken *FeeTokenSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _FeeToken.Contract.ETERNALAPPROVALVALUE(&_FeeToken.CallOpts)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_FeeToken *FeeTokenCallerSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _FeeToken.Contract.ETERNALAPPROVALVALUE(&_FeeToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_FeeToken *FeeTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_FeeToken *FeeTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _FeeToken.Contract.Allowance(&_FeeToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_FeeToken *FeeTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _FeeToken.Contract.Allowance(&_FeeToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FeeToken *FeeTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FeeToken *FeeTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _FeeToken.Contract.BalanceOf(&_FeeToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FeeToken *FeeTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _FeeToken.Contract.BalanceOf(&_FeeToken.CallOpts, _owner)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_FeeToken *FeeTokenCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_FeeToken *FeeTokenSession) ControllerLookupName() ([32]byte, error) {
	return _FeeToken.Contract.ControllerLookupName(&_FeeToken.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_FeeToken *FeeTokenCallerSession) ControllerLookupName() ([32]byte, error) {
	return _FeeToken.Contract.ControllerLookupName(&_FeeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FeeToken *FeeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FeeToken *FeeTokenSession) Decimals() (uint8, error) {
	return _FeeToken.Contract.Decimals(&_FeeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FeeToken *FeeTokenCallerSession) Decimals() (uint8, error) {
	return _FeeToken.Contract.Decimals(&_FeeToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_FeeToken *FeeTokenCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_FeeToken *FeeTokenSession) GetController() (common.Address, error) {
	return _FeeToken.Contract.GetController(&_FeeToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_FeeToken *FeeTokenCallerSession) GetController() (common.Address, error) {
	return _FeeToken.Contract.GetController(&_FeeToken.CallOpts)
}

// GetFeeWindow is a free data retrieval call binding the contract method 0xf77f29b1.
//
// Solidity: function getFeeWindow() constant returns(address)
func (_FeeToken *FeeTokenCaller) GetFeeWindow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "getFeeWindow")
	return *ret0, err
}

// GetFeeWindow is a free data retrieval call binding the contract method 0xf77f29b1.
//
// Solidity: function getFeeWindow() constant returns(address)
func (_FeeToken *FeeTokenSession) GetFeeWindow() (common.Address, error) {
	return _FeeToken.Contract.GetFeeWindow(&_FeeToken.CallOpts)
}

// GetFeeWindow is a free data retrieval call binding the contract method 0xf77f29b1.
//
// Solidity: function getFeeWindow() constant returns(address)
func (_FeeToken *FeeTokenCallerSession) GetFeeWindow() (common.Address, error) {
	return _FeeToken.Contract.GetFeeWindow(&_FeeToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_FeeToken *FeeTokenCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_FeeToken *FeeTokenSession) GetInitialized() (bool, error) {
	return _FeeToken.Contract.GetInitialized(&_FeeToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_FeeToken *FeeTokenCallerSession) GetInitialized() (bool, error) {
	return _FeeToken.Contract.GetInitialized(&_FeeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FeeToken *FeeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FeeToken *FeeTokenSession) Name() (string, error) {
	return _FeeToken.Contract.Name(&_FeeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FeeToken *FeeTokenCallerSession) Name() (string, error) {
	return _FeeToken.Contract.Name(&_FeeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FeeToken *FeeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FeeToken *FeeTokenSession) Symbol() (string, error) {
	return _FeeToken.Contract.Symbol(&_FeeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FeeToken *FeeTokenCallerSession) Symbol() (string, error) {
	return _FeeToken.Contract.Symbol(&_FeeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FeeToken *FeeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FeeToken *FeeTokenSession) TotalSupply() (*big.Int, error) {
	return _FeeToken.Contract.TotalSupply(&_FeeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FeeToken *FeeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _FeeToken.Contract.TotalSupply(&_FeeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_FeeToken *FeeTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_FeeToken *FeeTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.Approve(&_FeeToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_FeeToken *FeeTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.Approve(&_FeeToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_FeeToken *FeeTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _FeeToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_FeeToken *FeeTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.DecreaseApproval(&_FeeToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_FeeToken *FeeTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.DecreaseApproval(&_FeeToken.TransactOpts, _spender, _subtractedValue)
}

// FeeWindowBurn is a paid mutator transaction binding the contract method 0xf7862ec2.
//
// Solidity: function feeWindowBurn(_target address, _amount uint256) returns(bool)
func (_FeeToken *FeeTokenTransactor) FeeWindowBurn(opts *bind.TransactOpts, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FeeToken.contract.Transact(opts, "feeWindowBurn", _target, _amount)
}

// FeeWindowBurn is a paid mutator transaction binding the contract method 0xf7862ec2.
//
// Solidity: function feeWindowBurn(_target address, _amount uint256) returns(bool)
func (_FeeToken *FeeTokenSession) FeeWindowBurn(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.FeeWindowBurn(&_FeeToken.TransactOpts, _target, _amount)
}

// FeeWindowBurn is a paid mutator transaction binding the contract method 0xf7862ec2.
//
// Solidity: function feeWindowBurn(_target address, _amount uint256) returns(bool)
func (_FeeToken *FeeTokenTransactorSession) FeeWindowBurn(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.FeeWindowBurn(&_FeeToken.TransactOpts, _target, _amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_FeeToken *FeeTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _FeeToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_FeeToken *FeeTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.IncreaseApproval(&_FeeToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_FeeToken *FeeTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.IncreaseApproval(&_FeeToken.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_feeWindow address) returns(bool)
func (_FeeToken *FeeTokenTransactor) Initialize(opts *bind.TransactOpts, _feeWindow common.Address) (*types.Transaction, error) {
	return _FeeToken.contract.Transact(opts, "initialize", _feeWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_feeWindow address) returns(bool)
func (_FeeToken *FeeTokenSession) Initialize(_feeWindow common.Address) (*types.Transaction, error) {
	return _FeeToken.Contract.Initialize(&_FeeToken.TransactOpts, _feeWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_feeWindow address) returns(bool)
func (_FeeToken *FeeTokenTransactorSession) Initialize(_feeWindow common.Address) (*types.Transaction, error) {
	return _FeeToken.Contract.Initialize(&_FeeToken.TransactOpts, _feeWindow)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xa00ce6a5.
//
// Solidity: function mintForReportingParticipant(_target address, _amount uint256) returns(bool)
func (_FeeToken *FeeTokenTransactor) MintForReportingParticipant(opts *bind.TransactOpts, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FeeToken.contract.Transact(opts, "mintForReportingParticipant", _target, _amount)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xa00ce6a5.
//
// Solidity: function mintForReportingParticipant(_target address, _amount uint256) returns(bool)
func (_FeeToken *FeeTokenSession) MintForReportingParticipant(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.MintForReportingParticipant(&_FeeToken.TransactOpts, _target, _amount)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xa00ce6a5.
//
// Solidity: function mintForReportingParticipant(_target address, _amount uint256) returns(bool)
func (_FeeToken *FeeTokenTransactorSession) MintForReportingParticipant(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.MintForReportingParticipant(&_FeeToken.TransactOpts, _target, _amount)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_FeeToken *FeeTokenTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _FeeToken.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_FeeToken *FeeTokenSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _FeeToken.Contract.SetController(&_FeeToken.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_FeeToken *FeeTokenTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _FeeToken.Contract.SetController(&_FeeToken.TransactOpts, _controller)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_FeeToken *FeeTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_FeeToken *FeeTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.Transfer(&_FeeToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_FeeToken *FeeTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.Transfer(&_FeeToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_FeeToken *FeeTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_FeeToken *FeeTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.TransferFrom(&_FeeToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_FeeToken *FeeTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeToken.Contract.TransferFrom(&_FeeToken.TransactOpts, _from, _to, _value)
}

// FeeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FeeToken contract.
type FeeTokenApprovalIterator struct {
	Event *FeeTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTokenApproval represents a Approval event raised by the FeeToken contract.
type FeeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_FeeToken *FeeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FeeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FeeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FeeTokenApprovalIterator{contract: _FeeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_FeeToken *FeeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FeeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FeeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTokenApproval)
				if err := _FeeToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FeeTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the FeeToken contract.
type FeeTokenBurnIterator struct {
	Event *FeeTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTokenBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeTokenBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTokenBurn represents a Burn event raised by the FeeToken contract.
type FeeTokenBurn struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_FeeToken *FeeTokenFilterer) FilterBurn(opts *bind.FilterOpts, target []common.Address) (*FeeTokenBurnIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeToken.contract.FilterLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return &FeeTokenBurnIterator{contract: _FeeToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_FeeToken *FeeTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *FeeTokenBurn, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeToken.contract.WatchLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTokenBurn)
				if err := _FeeToken.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FeeTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the FeeToken contract.
type FeeTokenMintIterator struct {
	Event *FeeTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTokenMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeTokenMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTokenMint represents a Mint event raised by the FeeToken contract.
type FeeTokenMint struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_FeeToken *FeeTokenFilterer) FilterMint(opts *bind.FilterOpts, target []common.Address) (*FeeTokenMintIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeToken.contract.FilterLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return &FeeTokenMintIterator{contract: _FeeToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_FeeToken *FeeTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *FeeTokenMint, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeToken.contract.WatchLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTokenMint)
				if err := _FeeToken.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FeeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FeeToken contract.
type FeeTokenTransferIterator struct {
	Event *FeeTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeTokenTransfer represents a Transfer event raised by the FeeToken contract.
type FeeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_FeeToken *FeeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeTokenTransferIterator{contract: _FeeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_FeeToken *FeeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FeeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeTokenTransfer)
				if err := _FeeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FeeTokenFactoryABI is the input ABI used to generate the binding from.
const FeeTokenFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_feeWindow\",\"type\":\"address\"}],\"name\":\"createFeeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeeTokenFactoryBin is the compiled bytecode used for deploying new contracts.
const FeeTokenFactoryBin = `0x608060405234801561001057600080fd5b50610529806100206000396000f3006080604052600436106100405763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631e9e9dd48114610045575b600080fd5b34801561005157600080fd5b5061007973ffffffffffffffffffffffffffffffffffffffff600435811690602435166100a2565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000806000846100b06101e5565b73ffffffffffffffffffffffffffffffffffffffff90911681527f466565546f6b656e0000000000000000000000000000000000000000000000006020820152604080519182900301906000f08015801561010f573d6000803e3d6000fd5b5091508190508073ffffffffffffffffffffffffffffffffffffffff1663c4d66de8856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156101b057600080fd5b505af11580156101c4573d6000803e3d6000fd5b505050506040513d60208110156101da57600080fd5b509095945050505050565b604051610308806101f6833901905600608060405234801561001057600080fd5b506040516040806103088339810160405280516020909101516000805433600160a060020a03199182161716600160a060020a039093169290921782556001556102a890819061006090396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461014f57806392eefe9b1461018d578063bef72fa2146101cf575b60015460009015156100675761014c565b60008054600154604080517ff39ec1f700000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff9092169263f39ec1f7926024808401936020939083900390910190829087803b1580156100df57600080fd5b505af11580156100f3573d6000803e3d6000fd5b505050506040513d602081101561010957600080fd5b50516040805136601f8101601f191680830190935292935091600083376000803684865af4808015610147576040513d81016040523d6000823e3d81f35b600080fd5b50005b34801561015b57600080fd5b506101646101f6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019957600080fd5b506101bb73ffffffffffffffffffffffffffffffffffffffff60043516610212565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610276565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461023757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a7230582017c3a6b57ac8bbc7fd5320cba2543fe371eb8726332cd58700252790e09cf2140029a165627a7a72305820c3ba7283e647fba0be0ac7737f535e979717138bb181c5a280781fc53f52d88d0029`

// DeployFeeTokenFactory deploys a new Ethereum contract, binding an instance of FeeTokenFactory to it.
func DeployFeeTokenFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FeeTokenFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeTokenFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeeTokenFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeTokenFactory{FeeTokenFactoryCaller: FeeTokenFactoryCaller{contract: contract}, FeeTokenFactoryTransactor: FeeTokenFactoryTransactor{contract: contract}, FeeTokenFactoryFilterer: FeeTokenFactoryFilterer{contract: contract}}, nil
}

// FeeTokenFactory is an auto generated Go binding around an Ethereum contract.
type FeeTokenFactory struct {
	FeeTokenFactoryCaller     // Read-only binding to the contract
	FeeTokenFactoryTransactor // Write-only binding to the contract
	FeeTokenFactoryFilterer   // Log filterer for contract events
}

// FeeTokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeTokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeTokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeTokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeTokenFactorySession struct {
	Contract     *FeeTokenFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeTokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeTokenFactoryCallerSession struct {
	Contract *FeeTokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FeeTokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeTokenFactoryTransactorSession struct {
	Contract     *FeeTokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FeeTokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeTokenFactoryRaw struct {
	Contract *FeeTokenFactory // Generic contract binding to access the raw methods on
}

// FeeTokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeTokenFactoryCallerRaw struct {
	Contract *FeeTokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FeeTokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeTokenFactoryTransactorRaw struct {
	Contract *FeeTokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeTokenFactory creates a new instance of FeeTokenFactory, bound to a specific deployed contract.
func NewFeeTokenFactory(address common.Address, backend bind.ContractBackend) (*FeeTokenFactory, error) {
	contract, err := bindFeeTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeTokenFactory{FeeTokenFactoryCaller: FeeTokenFactoryCaller{contract: contract}, FeeTokenFactoryTransactor: FeeTokenFactoryTransactor{contract: contract}, FeeTokenFactoryFilterer: FeeTokenFactoryFilterer{contract: contract}}, nil
}

// NewFeeTokenFactoryCaller creates a new read-only instance of FeeTokenFactory, bound to a specific deployed contract.
func NewFeeTokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*FeeTokenFactoryCaller, error) {
	contract, err := bindFeeTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeTokenFactoryCaller{contract: contract}, nil
}

// NewFeeTokenFactoryTransactor creates a new write-only instance of FeeTokenFactory, bound to a specific deployed contract.
func NewFeeTokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeTokenFactoryTransactor, error) {
	contract, err := bindFeeTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeTokenFactoryTransactor{contract: contract}, nil
}

// NewFeeTokenFactoryFilterer creates a new log filterer instance of FeeTokenFactory, bound to a specific deployed contract.
func NewFeeTokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeTokenFactoryFilterer, error) {
	contract, err := bindFeeTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeTokenFactoryFilterer{contract: contract}, nil
}

// bindFeeTokenFactory binds a generic wrapper to an already deployed contract.
func bindFeeTokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeTokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeTokenFactory *FeeTokenFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeTokenFactory.Contract.FeeTokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeTokenFactory *FeeTokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeTokenFactory.Contract.FeeTokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeTokenFactory *FeeTokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeTokenFactory.Contract.FeeTokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeTokenFactory *FeeTokenFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeTokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeTokenFactory *FeeTokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeTokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeTokenFactory *FeeTokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeTokenFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateFeeToken is a paid mutator transaction binding the contract method 0x1e9e9dd4.
//
// Solidity: function createFeeToken(_controller address, _feeWindow address) returns(address)
func (_FeeTokenFactory *FeeTokenFactoryTransactor) CreateFeeToken(opts *bind.TransactOpts, _controller common.Address, _feeWindow common.Address) (*types.Transaction, error) {
	return _FeeTokenFactory.contract.Transact(opts, "createFeeToken", _controller, _feeWindow)
}

// CreateFeeToken is a paid mutator transaction binding the contract method 0x1e9e9dd4.
//
// Solidity: function createFeeToken(_controller address, _feeWindow address) returns(address)
func (_FeeTokenFactory *FeeTokenFactorySession) CreateFeeToken(_controller common.Address, _feeWindow common.Address) (*types.Transaction, error) {
	return _FeeTokenFactory.Contract.CreateFeeToken(&_FeeTokenFactory.TransactOpts, _controller, _feeWindow)
}

// CreateFeeToken is a paid mutator transaction binding the contract method 0x1e9e9dd4.
//
// Solidity: function createFeeToken(_controller address, _feeWindow address) returns(address)
func (_FeeTokenFactory *FeeTokenFactoryTransactorSession) CreateFeeToken(_controller common.Address, _feeWindow common.Address) (*types.Transaction, error) {
	return _FeeTokenFactory.Contract.CreateFeeToken(&_FeeTokenFactory.TransactOpts, _controller, _feeWindow)
}

// IFeeTokenABI is the input ABI used to generate the binding from.
const IFeeTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_feeWindow\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"feeWindowBurn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// IFeeTokenBin is the compiled bytecode used for deploying new contracts.
const IFeeTokenBin = `0x`

// DeployIFeeToken deploys a new Ethereum contract, binding an instance of IFeeToken to it.
func DeployIFeeToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IFeeToken, error) {
	parsed, err := abi.JSON(strings.NewReader(IFeeTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IFeeTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IFeeToken{IFeeTokenCaller: IFeeTokenCaller{contract: contract}, IFeeTokenTransactor: IFeeTokenTransactor{contract: contract}, IFeeTokenFilterer: IFeeTokenFilterer{contract: contract}}, nil
}

// IFeeToken is an auto generated Go binding around an Ethereum contract.
type IFeeToken struct {
	IFeeTokenCaller     // Read-only binding to the contract
	IFeeTokenTransactor // Write-only binding to the contract
	IFeeTokenFilterer   // Log filterer for contract events
}

// IFeeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFeeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFeeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFeeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFeeTokenSession struct {
	Contract     *IFeeToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFeeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFeeTokenCallerSession struct {
	Contract *IFeeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IFeeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFeeTokenTransactorSession struct {
	Contract     *IFeeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IFeeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFeeTokenRaw struct {
	Contract *IFeeToken // Generic contract binding to access the raw methods on
}

// IFeeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFeeTokenCallerRaw struct {
	Contract *IFeeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IFeeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFeeTokenTransactorRaw struct {
	Contract *IFeeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFeeToken creates a new instance of IFeeToken, bound to a specific deployed contract.
func NewIFeeToken(address common.Address, backend bind.ContractBackend) (*IFeeToken, error) {
	contract, err := bindIFeeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFeeToken{IFeeTokenCaller: IFeeTokenCaller{contract: contract}, IFeeTokenTransactor: IFeeTokenTransactor{contract: contract}, IFeeTokenFilterer: IFeeTokenFilterer{contract: contract}}, nil
}

// NewIFeeTokenCaller creates a new read-only instance of IFeeToken, bound to a specific deployed contract.
func NewIFeeTokenCaller(address common.Address, caller bind.ContractCaller) (*IFeeTokenCaller, error) {
	contract, err := bindIFeeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeTokenCaller{contract: contract}, nil
}

// NewIFeeTokenTransactor creates a new write-only instance of IFeeToken, bound to a specific deployed contract.
func NewIFeeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IFeeTokenTransactor, error) {
	contract, err := bindIFeeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeTokenTransactor{contract: contract}, nil
}

// NewIFeeTokenFilterer creates a new log filterer instance of IFeeToken, bound to a specific deployed contract.
func NewIFeeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IFeeTokenFilterer, error) {
	contract, err := bindIFeeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFeeTokenFilterer{contract: contract}, nil
}

// bindIFeeToken binds a generic wrapper to an already deployed contract.
func bindIFeeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFeeTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeToken *IFeeTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IFeeToken.Contract.IFeeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeToken *IFeeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeToken.Contract.IFeeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeToken *IFeeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeToken.Contract.IFeeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeToken *IFeeTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IFeeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeToken *IFeeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeToken *IFeeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IFeeToken *IFeeTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IFeeToken *IFeeTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IFeeToken.Contract.Allowance(&_IFeeToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IFeeToken *IFeeTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IFeeToken.Contract.Allowance(&_IFeeToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IFeeToken *IFeeTokenCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeToken.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IFeeToken *IFeeTokenSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _IFeeToken.Contract.BalanceOf(&_IFeeToken.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IFeeToken *IFeeTokenCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _IFeeToken.Contract.BalanceOf(&_IFeeToken.CallOpts, _who)
}

// GetFeeWindow is a free data retrieval call binding the contract method 0xf77f29b1.
//
// Solidity: function getFeeWindow() constant returns(address)
func (_IFeeToken *IFeeTokenCaller) GetFeeWindow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IFeeToken.contract.Call(opts, out, "getFeeWindow")
	return *ret0, err
}

// GetFeeWindow is a free data retrieval call binding the contract method 0xf77f29b1.
//
// Solidity: function getFeeWindow() constant returns(address)
func (_IFeeToken *IFeeTokenSession) GetFeeWindow() (common.Address, error) {
	return _IFeeToken.Contract.GetFeeWindow(&_IFeeToken.CallOpts)
}

// GetFeeWindow is a free data retrieval call binding the contract method 0xf77f29b1.
//
// Solidity: function getFeeWindow() constant returns(address)
func (_IFeeToken *IFeeTokenCallerSession) GetFeeWindow() (common.Address, error) {
	return _IFeeToken.Contract.GetFeeWindow(&_IFeeToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_IFeeToken *IFeeTokenCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IFeeToken.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_IFeeToken *IFeeTokenSession) GetInitialized() (bool, error) {
	return _IFeeToken.Contract.GetInitialized(&_IFeeToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_IFeeToken *IFeeTokenCallerSession) GetInitialized() (bool, error) {
	return _IFeeToken.Contract.GetInitialized(&_IFeeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IFeeToken *IFeeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IFeeToken *IFeeTokenSession) TotalSupply() (*big.Int, error) {
	return _IFeeToken.Contract.TotalSupply(&_IFeeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IFeeToken *IFeeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IFeeToken.Contract.TotalSupply(&_IFeeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IFeeToken *IFeeTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.Approve(&_IFeeToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.Approve(&_IFeeToken.TransactOpts, _spender, _value)
}

// FeeWindowBurn is a paid mutator transaction binding the contract method 0xf7862ec2.
//
// Solidity: function feeWindowBurn(_target address, _amount uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactor) FeeWindowBurn(opts *bind.TransactOpts, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IFeeToken.contract.Transact(opts, "feeWindowBurn", _target, _amount)
}

// FeeWindowBurn is a paid mutator transaction binding the contract method 0xf7862ec2.
//
// Solidity: function feeWindowBurn(_target address, _amount uint256) returns(bool)
func (_IFeeToken *IFeeTokenSession) FeeWindowBurn(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.FeeWindowBurn(&_IFeeToken.TransactOpts, _target, _amount)
}

// FeeWindowBurn is a paid mutator transaction binding the contract method 0xf7862ec2.
//
// Solidity: function feeWindowBurn(_target address, _amount uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactorSession) FeeWindowBurn(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.FeeWindowBurn(&_IFeeToken.TransactOpts, _target, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_feeWindow address) returns(bool)
func (_IFeeToken *IFeeTokenTransactor) Initialize(opts *bind.TransactOpts, _feeWindow common.Address) (*types.Transaction, error) {
	return _IFeeToken.contract.Transact(opts, "initialize", _feeWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_feeWindow address) returns(bool)
func (_IFeeToken *IFeeTokenSession) Initialize(_feeWindow common.Address) (*types.Transaction, error) {
	return _IFeeToken.Contract.Initialize(&_IFeeToken.TransactOpts, _feeWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_feeWindow address) returns(bool)
func (_IFeeToken *IFeeTokenTransactorSession) Initialize(_feeWindow common.Address) (*types.Transaction, error) {
	return _IFeeToken.Contract.Initialize(&_IFeeToken.TransactOpts, _feeWindow)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xa00ce6a5.
//
// Solidity: function mintForReportingParticipant(_target address, _amount uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactor) MintForReportingParticipant(opts *bind.TransactOpts, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IFeeToken.contract.Transact(opts, "mintForReportingParticipant", _target, _amount)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xa00ce6a5.
//
// Solidity: function mintForReportingParticipant(_target address, _amount uint256) returns(bool)
func (_IFeeToken *IFeeTokenSession) MintForReportingParticipant(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.MintForReportingParticipant(&_IFeeToken.TransactOpts, _target, _amount)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xa00ce6a5.
//
// Solidity: function mintForReportingParticipant(_target address, _amount uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactorSession) MintForReportingParticipant(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.MintForReportingParticipant(&_IFeeToken.TransactOpts, _target, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IFeeToken *IFeeTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.Transfer(&_IFeeToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.Transfer(&_IFeeToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IFeeToken *IFeeTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.TransferFrom(&_IFeeToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IFeeToken *IFeeTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeToken.Contract.TransferFrom(&_IFeeToken.TransactOpts, _from, _to, _value)
}

// IFeeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IFeeToken contract.
type IFeeTokenApprovalIterator struct {
	Event *IFeeTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IFeeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IFeeTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IFeeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeTokenApproval represents a Approval event raised by the IFeeToken contract.
type IFeeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IFeeToken *IFeeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IFeeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IFeeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IFeeTokenApprovalIterator{contract: _IFeeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IFeeToken *IFeeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IFeeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IFeeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeTokenApproval)
				if err := _IFeeToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IFeeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IFeeToken contract.
type IFeeTokenTransferIterator struct {
	Event *IFeeTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IFeeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IFeeTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IFeeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeTokenTransfer represents a Transfer event raised by the IFeeToken contract.
type IFeeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IFeeToken *IFeeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IFeeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFeeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFeeTokenTransferIterator{contract: _IFeeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IFeeToken *IFeeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IFeeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFeeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeTokenTransfer)
				if err := _IFeeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
