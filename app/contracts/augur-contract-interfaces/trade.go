package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ITradeABI is the input ABI used to generate the binding from.
const ITradeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_direction\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_tradeGroupID\",\"type\":\"uint256\"}],\"name\":\"publicFillBestOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupID\",\"type\":\"uint256\"}],\"name\":\"publicSell\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupID\",\"type\":\"uint256\"}],\"name\":\"publicBuy\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_direction\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupID\",\"type\":\"uint256\"}],\"name\":\"publicTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITradeBin is the compiled bytecode used for deploying new contracts.
const ITradeBin = `0x`

// DeployITrade deploys a new Ethereum contract, binding an instance of ITrade to it.
func DeployITrade(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ITrade, error) {
	parsed, err := abi.JSON(strings.NewReader(ITradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ITradeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ITrade{ITradeCaller: ITradeCaller{contract: contract}, ITradeTransactor: ITradeTransactor{contract: contract}, ITradeFilterer: ITradeFilterer{contract: contract}}, nil
}

// ITrade is an auto generated Go binding around an Ethereum contract.
type ITrade struct {
	ITradeCaller     // Read-only binding to the contract
	ITradeTransactor // Write-only binding to the contract
	ITradeFilterer   // Log filterer for contract events
}

// ITradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITradeSession struct {
	Contract     *ITrade           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITradeCallerSession struct {
	Contract *ITradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITradeTransactorSession struct {
	Contract     *ITradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITradeRaw struct {
	Contract *ITrade // Generic contract binding to access the raw methods on
}

// ITradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITradeCallerRaw struct {
	Contract *ITradeCaller // Generic read-only contract binding to access the raw methods on
}

// ITradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITradeTransactorRaw struct {
	Contract *ITradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITrade creates a new instance of ITrade, bound to a specific deployed contract.
func NewITrade(address common.Address, backend bind.ContractBackend) (*ITrade, error) {
	contract, err := bindITrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITrade{ITradeCaller: ITradeCaller{contract: contract}, ITradeTransactor: ITradeTransactor{contract: contract}, ITradeFilterer: ITradeFilterer{contract: contract}}, nil
}

// NewITradeCaller creates a new read-only instance of ITrade, bound to a specific deployed contract.
func NewITradeCaller(address common.Address, caller bind.ContractCaller) (*ITradeCaller, error) {
	contract, err := bindITrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITradeCaller{contract: contract}, nil
}

// NewITradeTransactor creates a new write-only instance of ITrade, bound to a specific deployed contract.
func NewITradeTransactor(address common.Address, transactor bind.ContractTransactor) (*ITradeTransactor, error) {
	contract, err := bindITrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITradeTransactor{contract: contract}, nil
}

// NewITradeFilterer creates a new log filterer instance of ITrade, bound to a specific deployed contract.
func NewITradeFilterer(address common.Address, filterer bind.ContractFilterer) (*ITradeFilterer, error) {
	contract, err := bindITrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITradeFilterer{contract: contract}, nil
}

// bindITrade binds a generic wrapper to an already deployed contract.
func bindITrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITrade *ITradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITrade.Contract.ITradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITrade *ITradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITrade.Contract.ITradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITrade *ITradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITrade.Contract.ITradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITrade *ITradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITrade *ITradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITrade *ITradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITrade.Contract.contract.Transact(opts, method, params...)
}

// PublicBuy is a paid mutator transaction binding the contract method 0xcecc78d7.
//
// Solidity: function publicBuy(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupID uint256) returns(bytes32)
func (_ITrade *ITradeTransactor) PublicBuy(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.contract.Transact(opts, "publicBuy", _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupID)
}

// PublicBuy is a paid mutator transaction binding the contract method 0xcecc78d7.
//
// Solidity: function publicBuy(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupID uint256) returns(bytes32)
func (_ITrade *ITradeSession) PublicBuy(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.Contract.PublicBuy(&_ITrade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupID)
}

// PublicBuy is a paid mutator transaction binding the contract method 0xcecc78d7.
//
// Solidity: function publicBuy(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupID uint256) returns(bytes32)
func (_ITrade *ITradeTransactorSession) PublicBuy(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.Contract.PublicBuy(&_ITrade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupID)
}

// PublicFillBestOrder is a paid mutator transaction binding the contract method 0x9fba60e0.
//
// Solidity: function publicFillBestOrder(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _tradeGroupID uint256) returns(uint256)
func (_ITrade *ITradeTransactor) PublicFillBestOrder(opts *bind.TransactOpts, _direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.contract.Transact(opts, "publicFillBestOrder", _direction, _market, _outcome, _fxpAmount, _price, _tradeGroupID)
}

// PublicFillBestOrder is a paid mutator transaction binding the contract method 0x9fba60e0.
//
// Solidity: function publicFillBestOrder(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _tradeGroupID uint256) returns(uint256)
func (_ITrade *ITradeSession) PublicFillBestOrder(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.Contract.PublicFillBestOrder(&_ITrade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _tradeGroupID)
}

// PublicFillBestOrder is a paid mutator transaction binding the contract method 0x9fba60e0.
//
// Solidity: function publicFillBestOrder(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _tradeGroupID uint256) returns(uint256)
func (_ITrade *ITradeTransactorSession) PublicFillBestOrder(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.Contract.PublicFillBestOrder(&_ITrade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _tradeGroupID)
}

// PublicSell is a paid mutator transaction binding the contract method 0xcc725362.
//
// Solidity: function publicSell(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupID uint256) returns(bytes32)
func (_ITrade *ITradeTransactor) PublicSell(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.contract.Transact(opts, "publicSell", _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupID)
}

// PublicSell is a paid mutator transaction binding the contract method 0xcc725362.
//
// Solidity: function publicSell(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupID uint256) returns(bytes32)
func (_ITrade *ITradeSession) PublicSell(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.Contract.PublicSell(&_ITrade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupID)
}

// PublicSell is a paid mutator transaction binding the contract method 0xcc725362.
//
// Solidity: function publicSell(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupID uint256) returns(bytes32)
func (_ITrade *ITradeTransactorSession) PublicSell(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.Contract.PublicSell(&_ITrade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupID)
}

// PublicTrade is a paid mutator transaction binding the contract method 0xff99b07f.
//
// Solidity: function publicTrade(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupID uint256) returns(bytes32)
func (_ITrade *ITradeTransactor) PublicTrade(opts *bind.TransactOpts, _direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.contract.Transact(opts, "publicTrade", _direction, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupID)
}

// PublicTrade is a paid mutator transaction binding the contract method 0xff99b07f.
//
// Solidity: function publicTrade(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupID uint256) returns(bytes32)
func (_ITrade *ITradeSession) PublicTrade(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.Contract.PublicTrade(&_ITrade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupID)
}

// PublicTrade is a paid mutator transaction binding the contract method 0xff99b07f.
//
// Solidity: function publicTrade(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupID uint256) returns(bytes32)
func (_ITrade *ITradeTransactorSession) PublicTrade(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupID *big.Int) (*types.Transaction, error) {
	return _ITrade.Contract.PublicTrade(&_ITrade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupID)
}

// TradeABI is the input ABI used to generate the binding from.
const TradeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"name\":\"_loopLimit\",\"type\":\"uint256\"}],\"name\":\"publicSellWithLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_direction\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"publicTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_direction\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"name\":\"_loopLimit\",\"type\":\"uint256\"}],\"name\":\"publicTradeWithLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_direction\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"name\":\"_loopLimit\",\"type\":\"uint256\"}],\"name\":\"publicFillBestOrderWithLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"publicSell\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_direction\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"publicFillBestOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"name\":\"_loopLimit\",\"type\":\"uint256\"}],\"name\":\"publicBuyWithLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_fxpAmount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"publicBuy\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// TradeBin is the compiled bytecode used for deploying new contracts.
const TradeBin = `0x608060405260008054600160a860020a03191633179055612a65806100256000396000f30060806040526004361061008a5763ffffffff60e060020a6000350416630a9a877b811461008f57806315c64b61146100ca5780633018205f146100f6578063554de08c1461012757806392eefe9b14610157578063a127d74e1461018c578063b1b12ef6146101b5578063ef01e5a3146101db578063f069788714610201578063f2beaf591461022a575b600080fd5b6100b8600160a060020a036004351660243560443560643560843560a43560c43560e435610250565b60408051918252519081900360200190f35b6100b860ff60043516600160a060020a036024351660443560643560843560a43560c43560e435610500565b34801561010257600080fd5b5061010b610765565b60408051600160a060020a039092168252519081900360200190f35b6100b860ff60043516600160a060020a036024351660443560643560843560a43560c43560e43561010435610774565b34801561016357600080fd5b50610178600160a060020a0360043516610a24565b604080519115158252519081900360200190f35b6100b860ff60043516600160a060020a036024351660443560643560843560a43560c435610a6e565b6100b8600160a060020a036004351660243560443560643560843560a43560c435610d1a565b6100b860ff60043516600160a060020a036024351660443560643560843560a435610f80565b6100b8600160a060020a036004351660243560443560643560843560a43560c43560e43561122a565b6100b8600160a060020a036004351660243560443560643560843560a43560c435611451565b60008089600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561029457600080fd5b505af11580156102a8573d6000803e3d6000fd5b505050506040513d60208110156102be57600080fd5b505160008054604080516000805160206129fa83398151915281529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b15801561031357600080fd5b505af1158015610327573d6000803e3d6000fd5b505050506040513d602081101561033d57600080fd5b505160408051600080516020612a1a8339815191528152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b15801561039257600080fd5b505af11580156103a6573d6000803e3d6000fd5b505050506040513d60208110156103bc57600080fd5b505115156103c957600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561042457600080fd5b505af1158015610438573d6000803e3d6000fd5b505050506040513d602081101561044e57600080fd5b5051151561045b57600080fd5b610463611677565b506104773360018e8e8e8e8e8e8e8e6117b9565b92508b600160a060020a031663a0695f246040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156104b757600080fd5b505af11580156104cb573d6000803e3d6000fd5b505050506040513d60208110156104e157600080fd5b509293508392506104f061197d565b5050505098975050505050505050565b60008088600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561054457600080fd5b505af1158015610558573d6000803e3d6000fd5b505050506040513d602081101561056e57600080fd5b505160008054604080516000805160206129fa83398151915281529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b1580156105c357600080fd5b505af11580156105d7573d6000803e3d6000fd5b505050506040513d60208110156105ed57600080fd5b505160408051600080516020612a1a8339815191528152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b15801561064257600080fd5b505af1158015610656573d6000803e3d6000fd5b505050506040513d602081101561066c57600080fd5b5051151561067957600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156106d457600080fd5b505af11580156106e8573d6000803e3d6000fd5b505050506040513d60208110156106fe57600080fd5b5051151561070b57600080fd5b610713611677565b50610725338d8d8d8d8d8d8d8d611c83565b92508a600160a060020a031663a0695f246040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156104b757600080fd5b600054600160a060020a031690565b60008089600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156107b857600080fd5b505af11580156107cc573d6000803e3d6000fd5b505050506040513d60208110156107e257600080fd5b505160008054604080516000805160206129fa83398151915281529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b15801561083757600080fd5b505af115801561084b573d6000803e3d6000fd5b505050506040513d602081101561086157600080fd5b505160408051600080516020612a1a8339815191528152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b1580156108b657600080fd5b505af11580156108ca573d6000803e3d6000fd5b505050506040513d60208110156108e057600080fd5b505115156108ed57600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561094857600080fd5b505af115801561095c573d6000803e3d6000fd5b505050506040513d602081101561097257600080fd5b5051151561097f57600080fd5b610987611677565b5061099a338e8e8e8e8e8e8e8e8e6117b9565b92508b600160a060020a031663a0695f246040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156109da57600080fd5b505af11580156109ee573d6000803e3d6000fd5b505050506040513d6020811015610a0457600080fd5b50929350839250610a1361197d565b505050509998505050505050505050565b60008054600160a060020a03163314610a3c57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60008087600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610ab257600080fd5b505af1158015610ac6573d6000803e3d6000fd5b505050506040513d6020811015610adc57600080fd5b505160008054604080516000805160206129fa83398151915281529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b158015610b3157600080fd5b505af1158015610b45573d6000803e3d6000fd5b505050506040513d6020811015610b5b57600080fd5b505160408051600080516020612a1a8339815191528152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b158015610bb057600080fd5b505af1158015610bc4573d6000803e3d6000fd5b505050506040513d6020811015610bda57600080fd5b50511515610be757600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b158015610c4257600080fd5b505af1158015610c56573d6000803e3d6000fd5b505050506040513d6020811015610c6c57600080fd5b50511515610c7957600080fd5b610c81611677565b50610c92338c8c8c8c8c8c8c611e74565b925089600160a060020a031663a0695f246040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610cd257600080fd5b505af1158015610ce6573d6000803e3d6000fd5b505050506040513d6020811015610cfc57600080fd5b50929350839250610d0b61197d565b50505050979650505050505050565b60008088600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610d5e57600080fd5b505af1158015610d72573d6000803e3d6000fd5b505050506040513d6020811015610d8857600080fd5b505160008054604080516000805160206129fa83398151915281529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b158015610ddd57600080fd5b505af1158015610df1573d6000803e3d6000fd5b505050506040513d6020811015610e0757600080fd5b505160408051600080516020612a1a8339815191528152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b158015610e5c57600080fd5b505af1158015610e70573d6000803e3d6000fd5b505050506040513d6020811015610e8657600080fd5b50511515610e9357600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b158015610eee57600080fd5b505af1158015610f02573d6000803e3d6000fd5b505050506040513d6020811015610f1857600080fd5b50511515610f2557600080fd5b610f2d611677565b50610f403360018d8d8d8d8d8d8d611c83565b92508a600160a060020a031663a0695f246040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610cd257600080fd5b60008086600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610fc457600080fd5b505af1158015610fd8573d6000803e3d6000fd5b505050506040513d6020811015610fee57600080fd5b505160008054604080516000805160206129fa83398151915281529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b15801561104357600080fd5b505af1158015611057573d6000803e3d6000fd5b505050506040513d602081101561106d57600080fd5b505160408051600080516020612a1a8339815191528152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b1580156110c257600080fd5b505af11580156110d6573d6000803e3d6000fd5b505050506040513d60208110156110ec57600080fd5b505115156110f957600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561115457600080fd5b505af1158015611168573d6000803e3d6000fd5b505050506040513d602081101561117e57600080fd5b5051151561118b57600080fd5b611193611677565b506111a3338b8b8b8b8b8b6123fc565b925088600160a060020a031663a0695f246040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156111e357600080fd5b505af11580156111f7573d6000803e3d6000fd5b505050506040513d602081101561120d57600080fd5b5092935083925061121c61197d565b505050509695505050505050565b60008089600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561126e57600080fd5b505af1158015611282573d6000803e3d6000fd5b505050506040513d602081101561129857600080fd5b505160008054604080516000805160206129fa83398151915281529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b1580156112ed57600080fd5b505af1158015611301573d6000803e3d6000fd5b505050506040513d602081101561131757600080fd5b505160408051600080516020612a1a8339815191528152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b15801561136c57600080fd5b505af1158015611380573d6000803e3d6000fd5b505050506040513d602081101561139657600080fd5b505115156113a357600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156113fe57600080fd5b505af1158015611412573d6000803e3d6000fd5b505050506040513d602081101561142857600080fd5b5051151561143557600080fd5b61143d611677565b506104773360008e8e8e8e8e8e8e8e6117b9565b60008088600081600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561149557600080fd5b505af11580156114a9573d6000803e3d6000fd5b505050506040513d60208110156114bf57600080fd5b505160008054604080516000805160206129fa83398151915281529051939450600160a060020a0390911692634e94c82992600480840193602093929083900390910190829087803b15801561151457600080fd5b505af1158015611528573d6000803e3d6000fd5b505050506040513d602081101561153e57600080fd5b505160408051600080516020612a1a8339815191528152600160a060020a03848116600483015291519190921691638cfb8f219160248083019260209291908290030181600087803b15801561159357600080fd5b505af11580156115a7573d6000803e3d6000fd5b505050506040513d60208110156115bd57600080fd5b505115156115ca57600080fd5b80600160a060020a0316639f7e1bf6836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561162557600080fd5b505af1158015611639573d6000803e3d6000fd5b505050506040513d602081101561164f57600080fd5b5051151561165c57600080fd5b611664611677565b50610f403360008d8d8d8d8d8d8d611c83565b6000803411156117b35760008054604080516000805160206129da83398151915281527f436173680000000000000000000000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b1580156116f657600080fd5b505af115801561170a573d6000803e3d6000fd5b505050506040513d602081101561172057600080fd5b5051604080517f4faa8a260000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691634faa8a26913491602480830192602092919082900301818588803b15801561178557600080fd5b505af1158015611799573d6000803e3d6000fd5b50505050506040513d60208110156117b057600080fd5b50505b50600190565b6000806117cc8c8c8c8c8c8c8a8a611e74565b90508015156117de576001915061196e565b60008054604080516000805160206129da83398151915281527f4372656174654f7264657200000000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b15801561185357600080fd5b505af1158015611867573d6000803e3d6000fd5b505050506040513d602081101561187d57600080fd5b5051600160a060020a031663f79383288d6118978e612984565b848b8f8f8d8d8d6040518a63ffffffff1660e060020a028152600401808a600160a060020a0316600160a060020a031681526020018960018111156118d857fe5b60ff168152602080820199909952604080820198909852600160a060020a03909616606087015250608085019390935260a084019190915260c083015260e08201529051610100808301955092935091908290030181600087803b15801561193f57600080fd5b505af1158015611953573d6000803e3d6000fd5b505050506040513d602081101561196957600080fd5b505191505b509a9950505050505050505050565b60008054604080516000805160206129da83398151915281527f436173680000000000000000000000000000000000000000000000000000000060048201529051839283928392600160a060020a039092169163f39ec1f79160248082019260209290919082900301818787803b1580156119f757600080fd5b505af1158015611a0b573d6000803e3d6000fd5b505050506040513d6020811015611a2157600080fd5b5051604080517f70a082310000000000000000000000000000000000000000000000000000000081523360048201529051919450600160a060020a038516916370a08231916024808201926020929091908290030181600087803b158015611a8857600080fd5b505af1158015611a9c573d6000803e3d6000fd5b505050506040513d6020811015611ab257600080fd5b505191506000821115611c79576000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611b1157600080fd5b505af1158015611b25573d6000803e3d6000fd5b505050506040513d6020811015611b3b57600080fd5b5051604080517fec238994000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301523360248301523060448301526064820186905291519293509083169163ec238994916084808201926020929091908290030181600087803b158015611bb857600080fd5b505af1158015611bcc573d6000803e3d6000fd5b505050506040513d6020811015611be257600080fd5b5050604080517f1baffe38000000000000000000000000000000000000000000000000000000008152336004820152602481018490529051600160a060020a03851691631baffe389160448083019260209291908290030181600087803b158015611c4c57600080fd5b505af1158015611c60573d6000803e3d6000fd5b505050506040513d6020811015611c7657600080fd5b50505b6001935050505090565b6000806000611c978c8c8c8c8c8c8a6123fc565b9150811515611ca95760019250611e65565b611cb16129a8565b5a1015611cc15760019250611e65565b611cca8b612984565b60008054604080516000805160206129da83398151915281527f4372656174654f7264657200000000000000000000000000000000000000000060048201529051939450600160a060020a039091169263f39ec1f792602480840193602093929083900390910190829087803b158015611d4357600080fd5b505af1158015611d57573d6000803e3d6000fd5b505050506040513d6020811015611d6d57600080fd5b8101908080519060200190929190505050600160a060020a031663f79383288d83858b8f8f8d8d8d6040518a63ffffffff1660e060020a028152600401808a600160a060020a0316600160a060020a03168152602001896001811115611dcf57fe5b60ff168152602080820199909952604080820198909852600160a060020a03909616606087015250608085019390935260a084019190915260c083015260e08201529051610100808301955092935091908290030181600087803b158015611e3657600080fd5b505af1158015611e4a573d6000803e3d6000fd5b505050506040513d6020811015611e6057600080fd5b505192505b50509998505050505050505050565b600080600080600080600060149054906101000a900460ff16151515611e9957600080fd5b6000805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055611ed78d6129af565b60008054604080516000805160206129da83398151915281527f4f7264657273000000000000000000000000000000000000000000000000000060048201529051939850600160a060020a039091169263f39ec1f792602480840193602093929083900390910190829087803b158015611f5057600080fd5b505af1158015611f64573d6000803e3d6000fd5b505050506040513d6020811015611f7a57600080fd5b50516040517f7b6eaa65000000000000000000000000000000000000000000000000000000008152909450600160a060020a03851690637b6eaa659087908f908f9060040180846001811115611fcc57fe5b60ff16815260200183600160a060020a0316600160a060020a031681526020018281526020019350505050602060405180830381600087803b15801561201157600080fd5b505af1158015612025573d6000803e3d6000fd5b505050506040513d602081101561203b57600080fd5b50518a965092505b82158015906120525750600086115b801561205e5750600087115b156123c157604080517f31d98b3f000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a038616916331d98b3f9160248083019260209291908290030181600087803b1580156120c557600080fd5b505af11580156120d9573d6000803e3d6000fd5b505050506040513d60208110156120ef57600080fd5b50519150600085600181111561210157fe5b1461210f5788821115612114565b888210155b156123b157604080517f8925f9e9000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03861691638925f9e99160248083019260209291908290030181600087803b15801561217b57600080fd5b505af115801561218f573d6000803e3d6000fd5b505050506040513d60208110156121a557600080fd5b5051604080517f3011e16a000000000000000000000000000000000000000000000000000000008152600160a060020a038f81166004830152602482018f905260448201869052915192935090861691633011e16a916064808201926020929091908290030181600087803b15801561221d57600080fd5b505af1158015612231573d6000803e3d6000fd5b505050506040513d602081101561224757600080fd5b505060008054604080516000805160206129da83398151915281527f46696c6c4f72646572000000000000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b1580156122be57600080fd5b505af11580156122d2573d6000803e3d6000fd5b505050506040513d60208110156122e857600080fd5b8101908080519060200190929190505050600160a060020a0316631faf89958f85898c6040518563ffffffff1660e060020a0281526004018085600160a060020a0316600160a060020a0316815260200184600019166000191681526020018381526020018260001916600019168152602001945050505050602060405180830381600087803b15801561237b57600080fd5b505af115801561238f573d6000803e3d6000fd5b505050506040513d60208110156123a557600080fd5b505195509150816123b6565b600092505b600187039650612043565b82156123cc57600095505b50506000805474ff00000000000000000000000000000000000000001916905550919a9950505050505050505050565b600080600080600080600060149054906101000a900460ff1615151561242157600080fd5b6000805474ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000017905561245f8c6129af565b60008054604080516000805160206129da83398151915281527f4f7264657273000000000000000000000000000000000000000000000000000060048201529051939850600160a060020a039091169263f39ec1f792602480840193602093929083900390910190829087803b1580156124d857600080fd5b505af11580156124ec573d6000803e3d6000fd5b505050506040513d602081101561250257600080fd5b50516040517f7b6eaa65000000000000000000000000000000000000000000000000000000008152909450600160a060020a03851690637b6eaa659087908e908e906004018084600181111561255457fe5b60ff16815260200183600160a060020a0316600160a060020a031681526020018281526020019350505050602060405180830381600087803b15801561259957600080fd5b505af11580156125ad573d6000803e3d6000fd5b505050506040513d60208110156125c357600080fd5b505189965092505b82158015906125da5750600086115b80156125ed57506125e96129d2565b5a10155b1561294a57604080517f31d98b3f000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a038616916331d98b3f9160248083019260209291908290030181600087803b15801561265457600080fd5b505af1158015612668573d6000803e3d6000fd5b505050506040513d602081101561267e57600080fd5b50519150600085600181111561269057fe5b1461269e57878211156126a3565b878210155b1561294057604080517f8925f9e9000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03861691638925f9e99160248083019260209291908290030181600087803b15801561270a57600080fd5b505af115801561271e573d6000803e3d6000fd5b505050506040513d602081101561273457600080fd5b5051604080517f3011e16a000000000000000000000000000000000000000000000000000000008152600160a060020a038e81166004830152602482018e905260448201869052915192935090861691633011e16a916064808201926020929091908290030181600087803b1580156127ac57600080fd5b505af11580156127c0573d6000803e3d6000fd5b505050506040513d60208110156127d657600080fd5b505060008054604080516000805160206129da83398151915281527f46696c6c4f72646572000000000000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b15801561284d57600080fd5b505af1158015612861573d6000803e3d6000fd5b505050506040513d602081101561287757600080fd5b8101908080519060200190929190505050600160a060020a0316631faf89958e85898b6040518563ffffffff1660e060020a0281526004018085600160a060020a0316600160a060020a0316815260200184600019166000191681526020018381526020018260001916600019168152602001945050505050602060405180830381600087803b15801561290a57600080fd5b505af115801561291e573d6000803e3d6000fd5b505050506040513d602081101561293457600080fd5b50519550915081612945565b600092505b6125cb565b821561295557600095505b50506000805474ff00000000000000000000000000000000000000001916905550919998505050505050505050565b60008082600181111561299357fe5b1461299f5760016129a2565b60005b92915050565b620aae6090565b6000808260018111156129be57fe5b146129ca5760006129a2565b506001919050565b621e8480905600f39ec1f7000000000000000000000000000000000000000000000000000000004e94c829000000000000000000000000000000000000000000000000000000008cfb8f2100000000000000000000000000000000000000000000000000000000a165627a7a723058202dfd756afb4c9176d23cfeddbfd7de47fc613e715267cfd1c923d01e4f1a5f9b0029`

// DeployTrade deploys a new Ethereum contract, binding an instance of Trade to it.
func DeployTrade(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trade, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TradeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trade{TradeCaller: TradeCaller{contract: contract}, TradeTransactor: TradeTransactor{contract: contract}, TradeFilterer: TradeFilterer{contract: contract}}, nil
}

// Trade is an auto generated Go binding around an Ethereum contract.
type Trade struct {
	TradeCaller     // Read-only binding to the contract
	TradeTransactor // Write-only binding to the contract
	TradeFilterer   // Log filterer for contract events
}

// TradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeSession struct {
	Contract     *Trade            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeCallerSession struct {
	Contract *TradeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeTransactorSession struct {
	Contract     *TradeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeRaw struct {
	Contract *Trade // Generic contract binding to access the raw methods on
}

// TradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeCallerRaw struct {
	Contract *TradeCaller // Generic read-only contract binding to access the raw methods on
}

// TradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeTransactorRaw struct {
	Contract *TradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrade creates a new instance of Trade, bound to a specific deployed contract.
func NewTrade(address common.Address, backend bind.ContractBackend) (*Trade, error) {
	contract, err := bindTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trade{TradeCaller: TradeCaller{contract: contract}, TradeTransactor: TradeTransactor{contract: contract}, TradeFilterer: TradeFilterer{contract: contract}}, nil
}

// NewTradeCaller creates a new read-only instance of Trade, bound to a specific deployed contract.
func NewTradeCaller(address common.Address, caller bind.ContractCaller) (*TradeCaller, error) {
	contract, err := bindTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradeCaller{contract: contract}, nil
}

// NewTradeTransactor creates a new write-only instance of Trade, bound to a specific deployed contract.
func NewTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeTransactor, error) {
	contract, err := bindTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradeTransactor{contract: contract}, nil
}

// NewTradeFilterer creates a new log filterer instance of Trade, bound to a specific deployed contract.
func NewTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*TradeFilterer, error) {
	contract, err := bindTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradeFilterer{contract: contract}, nil
}

// bindTrade binds a generic wrapper to an already deployed contract.
func bindTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.TradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transact(opts, method, params...)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Trade *TradeCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Trade *TradeSession) GetController() (common.Address, error) {
	return _Trade.Contract.GetController(&_Trade.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Trade *TradeCallerSession) GetController() (common.Address, error) {
	return _Trade.Contract.GetController(&_Trade.CallOpts)
}

// PublicBuy is a paid mutator transaction binding the contract method 0xf2beaf59.
//
// Solidity: function publicBuy(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_Trade *TradeTransactor) PublicBuy(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "publicBuy", _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicBuy is a paid mutator transaction binding the contract method 0xf2beaf59.
//
// Solidity: function publicBuy(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_Trade *TradeSession) PublicBuy(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.Contract.PublicBuy(&_Trade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicBuy is a paid mutator transaction binding the contract method 0xf2beaf59.
//
// Solidity: function publicBuy(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_Trade *TradeTransactorSession) PublicBuy(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.Contract.PublicBuy(&_Trade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicBuyWithLimit is a paid mutator transaction binding the contract method 0xf0697887.
//
// Solidity: function publicBuyWithLimit(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32, _loopLimit uint256) returns(bytes32)
func (_Trade *TradeTransactor) PublicBuyWithLimit(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "publicBuyWithLimit", _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit)
}

// PublicBuyWithLimit is a paid mutator transaction binding the contract method 0xf0697887.
//
// Solidity: function publicBuyWithLimit(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32, _loopLimit uint256) returns(bytes32)
func (_Trade *TradeSession) PublicBuyWithLimit(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.PublicBuyWithLimit(&_Trade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit)
}

// PublicBuyWithLimit is a paid mutator transaction binding the contract method 0xf0697887.
//
// Solidity: function publicBuyWithLimit(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32, _loopLimit uint256) returns(bytes32)
func (_Trade *TradeTransactorSession) PublicBuyWithLimit(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.PublicBuyWithLimit(&_Trade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit)
}

// PublicFillBestOrder is a paid mutator transaction binding the contract method 0xef01e5a3.
//
// Solidity: function publicFillBestOrder(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _tradeGroupId bytes32) returns(uint256)
func (_Trade *TradeTransactor) PublicFillBestOrder(opts *bind.TransactOpts, _direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "publicFillBestOrder", _direction, _market, _outcome, _fxpAmount, _price, _tradeGroupId)
}

// PublicFillBestOrder is a paid mutator transaction binding the contract method 0xef01e5a3.
//
// Solidity: function publicFillBestOrder(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _tradeGroupId bytes32) returns(uint256)
func (_Trade *TradeSession) PublicFillBestOrder(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.Contract.PublicFillBestOrder(&_Trade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _tradeGroupId)
}

// PublicFillBestOrder is a paid mutator transaction binding the contract method 0xef01e5a3.
//
// Solidity: function publicFillBestOrder(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _tradeGroupId bytes32) returns(uint256)
func (_Trade *TradeTransactorSession) PublicFillBestOrder(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.Contract.PublicFillBestOrder(&_Trade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _tradeGroupId)
}

// PublicFillBestOrderWithLimit is a paid mutator transaction binding the contract method 0xa127d74e.
//
// Solidity: function publicFillBestOrderWithLimit(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _tradeGroupId bytes32, _loopLimit uint256) returns(uint256)
func (_Trade *TradeTransactor) PublicFillBestOrderWithLimit(opts *bind.TransactOpts, _direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "publicFillBestOrderWithLimit", _direction, _market, _outcome, _fxpAmount, _price, _tradeGroupId, _loopLimit)
}

// PublicFillBestOrderWithLimit is a paid mutator transaction binding the contract method 0xa127d74e.
//
// Solidity: function publicFillBestOrderWithLimit(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _tradeGroupId bytes32, _loopLimit uint256) returns(uint256)
func (_Trade *TradeSession) PublicFillBestOrderWithLimit(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.PublicFillBestOrderWithLimit(&_Trade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _tradeGroupId, _loopLimit)
}

// PublicFillBestOrderWithLimit is a paid mutator transaction binding the contract method 0xa127d74e.
//
// Solidity: function publicFillBestOrderWithLimit(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _tradeGroupId bytes32, _loopLimit uint256) returns(uint256)
func (_Trade *TradeTransactorSession) PublicFillBestOrderWithLimit(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.PublicFillBestOrderWithLimit(&_Trade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _tradeGroupId, _loopLimit)
}

// PublicSell is a paid mutator transaction binding the contract method 0xb1b12ef6.
//
// Solidity: function publicSell(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_Trade *TradeTransactor) PublicSell(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "publicSell", _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicSell is a paid mutator transaction binding the contract method 0xb1b12ef6.
//
// Solidity: function publicSell(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_Trade *TradeSession) PublicSell(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.Contract.PublicSell(&_Trade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicSell is a paid mutator transaction binding the contract method 0xb1b12ef6.
//
// Solidity: function publicSell(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_Trade *TradeTransactorSession) PublicSell(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.Contract.PublicSell(&_Trade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicSellWithLimit is a paid mutator transaction binding the contract method 0x0a9a877b.
//
// Solidity: function publicSellWithLimit(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32, _loopLimit uint256) returns(bytes32)
func (_Trade *TradeTransactor) PublicSellWithLimit(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "publicSellWithLimit", _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit)
}

// PublicSellWithLimit is a paid mutator transaction binding the contract method 0x0a9a877b.
//
// Solidity: function publicSellWithLimit(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32, _loopLimit uint256) returns(bytes32)
func (_Trade *TradeSession) PublicSellWithLimit(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.PublicSellWithLimit(&_Trade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit)
}

// PublicSellWithLimit is a paid mutator transaction binding the contract method 0x0a9a877b.
//
// Solidity: function publicSellWithLimit(_market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32, _loopLimit uint256) returns(bytes32)
func (_Trade *TradeTransactorSession) PublicSellWithLimit(_market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.PublicSellWithLimit(&_Trade.TransactOpts, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit)
}

// PublicTrade is a paid mutator transaction binding the contract method 0x15c64b61.
//
// Solidity: function publicTrade(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_Trade *TradeTransactor) PublicTrade(opts *bind.TransactOpts, _direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "publicTrade", _direction, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicTrade is a paid mutator transaction binding the contract method 0x15c64b61.
//
// Solidity: function publicTrade(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_Trade *TradeSession) PublicTrade(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.Contract.PublicTrade(&_Trade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicTrade is a paid mutator transaction binding the contract method 0x15c64b61.
//
// Solidity: function publicTrade(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(bytes32)
func (_Trade *TradeTransactorSession) PublicTrade(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Trade.Contract.PublicTrade(&_Trade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// PublicTradeWithLimit is a paid mutator transaction binding the contract method 0x554de08c.
//
// Solidity: function publicTradeWithLimit(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32, _loopLimit uint256) returns(bytes32)
func (_Trade *TradeTransactor) PublicTradeWithLimit(opts *bind.TransactOpts, _direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "publicTradeWithLimit", _direction, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit)
}

// PublicTradeWithLimit is a paid mutator transaction binding the contract method 0x554de08c.
//
// Solidity: function publicTradeWithLimit(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32, _loopLimit uint256) returns(bytes32)
func (_Trade *TradeSession) PublicTradeWithLimit(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.PublicTradeWithLimit(&_Trade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit)
}

// PublicTradeWithLimit is a paid mutator transaction binding the contract method 0x554de08c.
//
// Solidity: function publicTradeWithLimit(_direction uint8, _market address, _outcome uint256, _fxpAmount uint256, _price uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32, _loopLimit uint256) returns(bytes32)
func (_Trade *TradeTransactorSession) PublicTradeWithLimit(_direction uint8, _market common.Address, _outcome *big.Int, _fxpAmount *big.Int, _price *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte, _loopLimit *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.PublicTradeWithLimit(&_Trade.TransactOpts, _direction, _market, _outcome, _fxpAmount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Trade *TradeTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Trade *TradeSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetController(&_Trade.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Trade *TradeTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetController(&_Trade.TransactOpts, _controller)
}

// Trade1ABI is the input ABI used to generate the binding from.
const Trade1ABI = "[]"

// Trade1Bin is the compiled bytecode used for deploying new contracts.
const Trade1Bin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582063e77be9f1fbbb5ffdd1d13c26c209b03b2e7af1dd83b822e86bde4c3a07bcb30029`

// DeployTrade1 deploys a new Ethereum contract, binding an instance of Trade1 to it.
func DeployTrade1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trade1, error) {
	parsed, err := abi.JSON(strings.NewReader(Trade1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Trade1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trade1{Trade1Caller: Trade1Caller{contract: contract}, Trade1Transactor: Trade1Transactor{contract: contract}, Trade1Filterer: Trade1Filterer{contract: contract}}, nil
}

// Trade1 is an auto generated Go binding around an Ethereum contract.
type Trade1 struct {
	Trade1Caller     // Read-only binding to the contract
	Trade1Transactor // Write-only binding to the contract
	Trade1Filterer   // Log filterer for contract events
}

// Trade1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Trade1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Trade1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Trade1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Trade1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Trade1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Trade1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Trade1Session struct {
	Contract     *Trade1           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Trade1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Trade1CallerSession struct {
	Contract *Trade1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Trade1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Trade1TransactorSession struct {
	Contract     *Trade1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Trade1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Trade1Raw struct {
	Contract *Trade1 // Generic contract binding to access the raw methods on
}

// Trade1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Trade1CallerRaw struct {
	Contract *Trade1Caller // Generic read-only contract binding to access the raw methods on
}

// Trade1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Trade1TransactorRaw struct {
	Contract *Trade1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTrade1 creates a new instance of Trade1, bound to a specific deployed contract.
func NewTrade1(address common.Address, backend bind.ContractBackend) (*Trade1, error) {
	contract, err := bindTrade1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trade1{Trade1Caller: Trade1Caller{contract: contract}, Trade1Transactor: Trade1Transactor{contract: contract}, Trade1Filterer: Trade1Filterer{contract: contract}}, nil
}

// NewTrade1Caller creates a new read-only instance of Trade1, bound to a specific deployed contract.
func NewTrade1Caller(address common.Address, caller bind.ContractCaller) (*Trade1Caller, error) {
	contract, err := bindTrade1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Trade1Caller{contract: contract}, nil
}

// NewTrade1Transactor creates a new write-only instance of Trade1, bound to a specific deployed contract.
func NewTrade1Transactor(address common.Address, transactor bind.ContractTransactor) (*Trade1Transactor, error) {
	contract, err := bindTrade1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Trade1Transactor{contract: contract}, nil
}

// NewTrade1Filterer creates a new log filterer instance of Trade1, bound to a specific deployed contract.
func NewTrade1Filterer(address common.Address, filterer bind.ContractFilterer) (*Trade1Filterer, error) {
	contract, err := bindTrade1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Trade1Filterer{contract: contract}, nil
}

// bindTrade1 binds a generic wrapper to an already deployed contract.
func bindTrade1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Trade1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade1 *Trade1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trade1.Contract.Trade1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade1 *Trade1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade1.Contract.Trade1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade1 *Trade1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade1.Contract.Trade1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade1 *Trade1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trade1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade1 *Trade1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade1 *Trade1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade1.Contract.contract.Transact(opts, method, params...)
}
