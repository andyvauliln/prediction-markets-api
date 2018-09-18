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

// FeeWindowABI is the input ABI used to generate the binding from.
const FeeWindowABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumMarkets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEndTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalFeeStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETERNAL_APPROVAL_VALUE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedUniverseBuy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"onMarketFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumIncorrectDesignatedReportMarkets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintFeeTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOver\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReputationToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"redeemForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStartTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_feeWindowId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumInvalidMarkets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumDesignatedReportNoShows\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// FeeWindowBin is the compiled bytecode used for deploying new contracts.
const FeeWindowBin = `0x60806040526005805460ff1916905560008054600160a060020a031916331790556121e18061002f6000396000f3006080604052600436106101a85763ffffffff60e060020a60003504166306fdde0381146101ad578063095ea7b31461023757806318160ddd1461026f57806322f3e2d41461029657806323b872dd146102ab578063295c39a5146102d55780633018205f146102ea578063313ce5671461031b578063439f5ac2146103465780635c3cdec81461035b578063634eaff114610370578063661884631461038557806370a08231146103a95780637303ed18146103ca578063870c426d146103ee5780638dc6e2f11461040357806392eefe9b1461041857806395a2251f1461043957806395d89b411461045a578063a52c05121461046f578063a7eb685b14610484578063a9059cbb1461049c578063b4bd9e27146104c0578063b80907f2146104d5578063b97a6c12146104ea578063bef72fa2146104ff578063c828371e14610514578063ca709a2514610529578063cd6dc6871461053e578063cf3d384914610562578063d73dd62314610577578063d96a094a1461059b578063da0b0c36146105b3578063db0a087c146105c8578063dd62ed3e146105dd578063ee89dab414610604575b600080fd5b3480156101b957600080fd5b506101c2610619565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101fc5781810151838201526020016101e4565b50505050905090810190601f1680156102295780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024357600080fd5b5061025b600160a060020a0360043516602435610650565b604080519115158252519081900360200190f35b34801561027b57600080fd5b50610284610667565b60408051918252519081900360200190f35b3480156102a257600080fd5b5061025b61066e565b3480156102b757600080fd5b5061025b600160a060020a03600435811690602435166044356107ae565b3480156102e157600080fd5b50610284610827565b3480156102f657600080fd5b506102ff610842565b60408051600160a060020a039092168252519081900360200190f35b34801561032757600080fd5b50610330610851565b6040805160ff9092168252519081900360200190f35b34801561035257600080fd5b50610284610856565b34801561036757600080fd5b5061028461088e565b34801561037c57600080fd5b50610284610935565b34801561039157600080fd5b5061025b600160a060020a036004351660243561093b565b3480156103b557600080fd5b50610284600160a060020a036004351661099e565b3480156103d657600080fd5b5061025b600160a060020a03600435166024356109b9565b3480156103fa57600080fd5b506102ff6109f3565b34801561040f57600080fd5b5061025b610a1c565b34801561042457600080fd5b5061025b600160a060020a0360043516610c72565b34801561044557600080fd5b5061025b600160a060020a0360043516610cbc565b34801561046657600080fd5b506101c2610cd2565b34801561047b57600080fd5b50610284610d09565b34801561049057600080fd5b5061025b600435610d0f565b3480156104a857600080fd5b5061025b600160a060020a0360043516602435610e47565b3480156104cc57600080fd5b5061025b610e5b565b3480156104e157600080fd5b506102ff610efc565b3480156104f657600080fd5b5061025b610f94565b34801561050b57600080fd5b50610284610fa9565b34801561052057600080fd5b50610284610faf565b34801561053557600080fd5b506102ff610fca565b34801561054a57600080fd5b5061025b600160a060020a0360043516602435610fee565b34801561056e57600080fd5b50610284611249565b34801561058357600080fd5b5061025b600160a060020a0360043516602435611264565b3480156105a757600080fd5b5061025b6004356112a0565b3480156105bf57600080fd5b506102846112be565b3480156105d457600080fd5b506102846112c4565b3480156105e957600080fd5b50610284600160a060020a03600435811690602435166112fd565b34801561061057600080fd5b5061025b611328565b60408051808201909152601381527f50617274696369706174696f6e20546f6b656e00000000000000000000000000602082015281565b600061065d338484611331565b5060019392505050565b6002545b90565b60055460009060ff16151561068257600080fd5b61068a610faf565b6000809054906101000a9004600160a060020a0316600160a060020a031663188ec3566040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156106dc57600080fd5b505af11580156106f0573d6000803e3d6000fd5b505050506040513d602081101561070657600080fd5b5051116107155750600061066b565b61071d610856565b6000809054906101000a9004600160a060020a0316600160a060020a031663188ec3566040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561076f57600080fd5b505af1158015610783573d6000803e3d6000fd5b505050506040513d602081101561079957600080fd5b5051106107a85750600061066b565b50600190565b600160a060020a03831660009081526004602090815260408083203384529091528120546000198114610810576107eb818463ffffffff61139c16565b600160a060020a03861660009081526004602090815260408083203384529091529020555b61081b8585856113b1565b50600195945050505050565b60055460009060ff16151561083b57600080fd5b5060075490565b600054600160a060020a031690565b600081565b60055460009060ff16151561086a57600080fd5b61088961087561147c565b61087d610faf565b9063ffffffff61148316565b905090565b600080600061089b610667565b9150600d60009054906101000a9004600160a060020a0316600160a060020a03166318160ddd6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156108f057600080fd5b505af1158015610904573d6000803e3d6000fd5b505050506040513d602081101561091a57600080fd5b5051905061092e828263ffffffff61148316565b9250505090565b60001981565b336000908152600460209081526040808320600160a060020a0386168452909152812054808311156109795761097333856000611331565b5061065d565b610993338561098e848763ffffffff61139c16565b611331565b505060019392505050565b600160a060020a031660009081526003602052604090205490565b60055460009060ff1615156109cd57600080fd5b6005546101009004600160a060020a031633146109e957600080fd5b61065d8383611495565b60055460009060ff161515610a0757600080fd5b506005546101009004600160a060020a031690565b600554600090819060ff161515610a3257600080fd5b50600554604080517f9f7e1bf60000000000000000000000000000000000000000000000000000000081523360048201819052915191926101009004600160a060020a031691639f7e1bf6916024808201926020929091908290030181600087803b158015610aa057600080fd5b505af1158015610ab4573d6000803e3d6000fd5b505050506040513d6020811015610aca57600080fd5b50511515610ad757600080fd5b600780546001019055604080517f04be2f500000000000000000000000000000000000000000000000000000000081529051600160a060020a038316916304be2f509160048083019260209291908290030181600087803b158015610b3b57600080fd5b505af1158015610b4f573d6000803e3d6000fd5b505050506040513d6020811015610b6557600080fd5b505115610b76576008805460010190555b80600160a060020a0316638ed882c56040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610bb457600080fd5b505af1158015610bc8573d6000803e3d6000fd5b505050506040513d6020811015610bde57600080fd5b50511515610bf0576009805460010190555b80600160a060020a0316633bf8f34a6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610c2e57600080fd5b505af1158015610c42573d6000803e3d6000fd5b505050506040513d6020811015610c5857600080fd5b50511515610c6a57600a805460010190555b600191505090565b60008054600160a060020a03163314610c8a57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b6000610cc9826000611604565b50600192915050565b60408051808201909152600281527f5054000000000000000000000000000000000000000000000000000000000000602082015281565b60095490565b600554604080517ff76514c700000000000000000000000000000000000000000000000000000000815233600482015290516000926101009004600160a060020a03169163f76514c791602480830192602092919082900301818787803b158015610d7957600080fd5b505af1158015610d8d573d6000803e3d6000fd5b505050506040513d6020811015610da357600080fd5b50511515610db057600080fd5b600d54604080517fa00ce6a5000000000000000000000000000000000000000000000000000000008152336004820152602481018590529051600160a060020a039092169163a00ce6a5916044808201926020929091908290030181600087803b158015610e1d57600080fd5b505af1158015610e31573d6000803e3d6000fd5b505050506040513d602081101561065d57600080fd5b6000610e543384846113b1565b9392505050565b60055460009060ff161515610e6f57600080fd5b610e77610856565b6000809054906101000a9004600160a060020a0316600160a060020a031663188ec3566040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610ec957600080fd5b505af1158015610edd573d6000803e3d6000fd5b505050506040513d6020811015610ef357600080fd5b50511015905090565b60055460009060ff161515610f1057600080fd5b600560019054906101000a9004600160a060020a0316600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610f6357600080fd5b505af1158015610f77573d6000803e3d6000fd5b505050506040513d6020811015610f8d57600080fd5b5051905090565b6000610fa1336001611604565b506001905090565b60015481565b60055460009060ff161515610fc357600080fd5b5060065490565b60055460009060ff161515610fde57600080fd5b50600d54600160a060020a031690565b60055460009060ff161561100157600080fd5b611009611cbe565b5082600560016101000a815481600160a060020a030219169083600160a060020a031602179055506110c0600560019054906101000a9004600160a060020a0316600160a060020a031663331172f36040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561108757600080fd5b505af115801561109b573d6000803e3d6000fd5b505050506040513d60208110156110b157600080fd5b5051839063ffffffff611ce516565b60065560008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f466565546f6b656e466163746f7279000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b15801561114a57600080fd5b505af115801561115e573d6000803e3d6000fd5b505050506040513d602081101561117457600080fd5b505160008054604080517f1e9e9dd4000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015230602482015290519190931692631e9e9dd49260448083019360209390929083900390910190829087803b1580156111e857600080fd5b505af11580156111fc573d6000803e3d6000fd5b505050506040513d602081101561121257600080fd5b5051600d805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390921691909117905550600192915050565b60055460009060ff16151561125d57600080fd5b5060085490565b336000818152600460209081526040808320600160a060020a0387168452909152812054909161065d91859061098e908663ffffffff61148316565b60055460009060ff1615156112b457600080fd5b610cc93383611495565b600a5490565b60055460009060ff1615156112d857600080fd5b507f46656557696e646f77000000000000000000000000000000000000000000000090565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205490565b60055460ff1690565b600160a060020a03808416600081815260046020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6000828211156113ab57600080fd5b50900390565b600160a060020a0383166000908152600360205260408120546113da908363ffffffff61139c16565b600160a060020a03808616600090815260036020526040808220939093559085168152205461140f908363ffffffff61148316565b600160a060020a0380851660008181526003602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3611471848484611d0c565b506001949350505050565b62093a8090565b600082820183811015610e5457600080fd5b60055460009060ff1615156114a957600080fd5b600082116114b657600080fd5b6114be61066e565b15156114c957600080fd5b600560019054906101000a9004600160a060020a0316600160a060020a031663becb1f356040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561151c57600080fd5b505af1158015611530573d6000803e3d6000fd5b505050506040513d602081101561154657600080fd5b50511561155257600080fd5b61155a610efc565b604080517f90fbf84e000000000000000000000000000000000000000000000000000000008152600160a060020a03868116600483015230602483015260448201869052915192909116916390fbf84e916064808201926020929091908290030181600087803b1580156115cd57600080fd5b505af11580156115e1573d6000803e3d6000fd5b505050506040513d60208110156115f757600080fd5b5061065d90508383611e3a565b600080600080600080600080611618610e5b565b8061169e5750600560019054906101000a9004600160a060020a0316600160a060020a031663becb1f356040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561167157600080fd5b505af1158015611685573d6000803e3d6000fd5b505050506040513d602081101561169b57600080fd5b50515b15156116a957600080fd5b600160a060020a03808b16600081815260036020908152604080832054600d5482517f70a0823100000000000000000000000000000000000000000000000000000000815260048101969096529151909c509416936370a0823193602480820194918390030190829087803b15801561172157600080fd5b505af1158015611735573d6000803e3d6000fd5b505050506040513d602081101561174b57600080fd5b5051955061175f878763ffffffff61148316565b945061176961088e565b9350861561181d5761177b8a88611edb565b50611784610efc565b600160a060020a031663a9059cbb8b896040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156117e657600080fd5b505af11580156117fa573d6000803e3d6000fd5b505050506040513d602081101561181057600080fd5b5051151561181d57600080fd5b85156118bf57600d54604080517ff7862ec2000000000000000000000000000000000000000000000000000000008152600160a060020a038d81166004830152602482018a90529151919092169163f7862ec29160448083019260209291908290030181600087803b15801561189257600080fd5b505af11580156118a6573d6000803e3d6000fd5b505050506040513d60208110156118bc57600080fd5b50505b8315156118cf5760019750611cb1565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f436173680000000000000000000000000000000000000000000000000000000060048201529051600160a060020a039092169263f39ec1f7926024808401936020939083900390910190829087803b15801561195657600080fd5b505af115801561196a573d6000803e3d6000fd5b505050506040513d602081101561198057600080fd5b5051604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051919450600160a060020a038516916370a08231916024808201926020929091908290030181600087803b1580156119e757600080fd5b505af11580156119fb573d6000803e3d6000fd5b505050506040513d6020811015611a1157600080fd5b50519150611a3584611a29848863ffffffff611ce516565b9063ffffffff611f7c16565b90506000811115611b75578815611ae55782600160a060020a031663a9059cbb8b836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015611aa957600080fd5b505af1158015611abd573d6000803e3d6000fd5b505050506040513d6020811015611ad357600080fd5b50511515611ae057600080fd5b611b75565b82600160a060020a0316631baffe388b836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015611b4857600080fd5b505af1158015611b5c573d6000803e3d6000fd5b505050506040513d6020811015611b7257600080fd5b50505b6000871115611cac576000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611bd057600080fd5b505af1158015611be4573d6000803e3d6000fd5b505050506040513d6020811015611bfa57600080fd5b5051600554604080517f3b186bfc000000000000000000000000000000000000000000000000000000008152610100909204600160a060020a0390811660048401528d81166024840152604483018b9052606483018590529051921691633b186bfc916084808201926020929091908290030181600087803b158015611c7f57600080fd5b505af1158015611c93573d6000803e3d6000fd5b505050506040513d6020811015611ca957600080fd5b50505b600197505b5050505050505092915050565b60055460009060ff1615611cd157600080fd5b506005805460ff1916600190811790915590565b6000828202831580611d015750828482811515611cfe57fe5b04145b1515610e5457600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611d6057600080fd5b505af1158015611d74573d6000803e3d6000fd5b505050506040513d6020811015611d8a57600080fd5b5051600554604080517f788873ea000000000000000000000000000000000000000000000000000000008152610100909204600160a060020a0390811660048401528781166024840152868116604484015260648301869052905192169163788873ea916084808201926020929091908290030181600087803b158015611e1057600080fd5b505af1158015611e24573d6000803e3d6000fd5b505050506040513d602081101561081b57600080fd5b600160a060020a038216600090815260036020526040812054611e63908363ffffffff61148316565b600160a060020a038416600090815260036020526040902055600254611e8f908363ffffffff61148316565b600255604080518381529051600160a060020a038516917f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885919081900360200190a261065d8383611f93565b600160a060020a038216600090815260036020526040812054611f04908363ffffffff61139c16565b600160a060020a038416600090815260036020526040902055600254611f30908363ffffffff61139c16565b600255604080518381529051600160a060020a038516917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a261065d83836120b9565b6000808284811515611f8a57fe5b04949350505050565b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611fe757600080fd5b505af1158015611ffb573d6000803e3d6000fd5b505050506040513d602081101561201157600080fd5b5051600554604080517f60fe103e000000000000000000000000000000000000000000000000000000008152610100909204600160a060020a03908116600484015286811660248401526044830186905290519216916360fe103e916064808201926020929091908290030181600087803b15801561208f57600080fd5b505af11580156120a3573d6000803e3d6000fd5b505050506040513d602081101561099357600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561210d57600080fd5b505af1158015612121573d6000803e3d6000fd5b505050506040513d602081101561213757600080fd5b5051600554604080517f542e9b18000000000000000000000000000000000000000000000000000000008152610100909204600160a060020a039081166004840152868116602484015260448301869052905192169163542e9b18916064808201926020929091908290030181600087803b15801561208f57600080fd00a165627a7a72305820f094018b3ecf7613cec763c67356b798852401a08e502db39889429cb5716f8c0029`

// DeployFeeWindow deploys a new Ethereum contract, binding an instance of FeeWindow to it.
func DeployFeeWindow(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FeeWindow, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeWindowABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeeWindowBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeWindow{FeeWindowCaller: FeeWindowCaller{contract: contract}, FeeWindowTransactor: FeeWindowTransactor{contract: contract}, FeeWindowFilterer: FeeWindowFilterer{contract: contract}}, nil
}

// FeeWindow is an auto generated Go binding around an Ethereum contract.
type FeeWindow struct {
	FeeWindowCaller     // Read-only binding to the contract
	FeeWindowTransactor // Write-only binding to the contract
	FeeWindowFilterer   // Log filterer for contract events
}

// FeeWindowCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeWindowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeWindowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeWindowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeWindowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeWindowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeWindowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeWindowSession struct {
	Contract     *FeeWindow        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeWindowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeWindowCallerSession struct {
	Contract *FeeWindowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FeeWindowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeWindowTransactorSession struct {
	Contract     *FeeWindowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FeeWindowRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeWindowRaw struct {
	Contract *FeeWindow // Generic contract binding to access the raw methods on
}

// FeeWindowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeWindowCallerRaw struct {
	Contract *FeeWindowCaller // Generic read-only contract binding to access the raw methods on
}

// FeeWindowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeWindowTransactorRaw struct {
	Contract *FeeWindowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeWindow creates a new instance of FeeWindow, bound to a specific deployed contract.
func NewFeeWindow(address common.Address, backend bind.ContractBackend) (*FeeWindow, error) {
	contract, err := bindFeeWindow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeWindow{FeeWindowCaller: FeeWindowCaller{contract: contract}, FeeWindowTransactor: FeeWindowTransactor{contract: contract}, FeeWindowFilterer: FeeWindowFilterer{contract: contract}}, nil
}

// NewFeeWindowCaller creates a new read-only instance of FeeWindow, bound to a specific deployed contract.
func NewFeeWindowCaller(address common.Address, caller bind.ContractCaller) (*FeeWindowCaller, error) {
	contract, err := bindFeeWindow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeWindowCaller{contract: contract}, nil
}

// NewFeeWindowTransactor creates a new write-only instance of FeeWindow, bound to a specific deployed contract.
func NewFeeWindowTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeWindowTransactor, error) {
	contract, err := bindFeeWindow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeWindowTransactor{contract: contract}, nil
}

// NewFeeWindowFilterer creates a new log filterer instance of FeeWindow, bound to a specific deployed contract.
func NewFeeWindowFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeWindowFilterer, error) {
	contract, err := bindFeeWindow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeWindowFilterer{contract: contract}, nil
}

// bindFeeWindow binds a generic wrapper to an already deployed contract.
func bindFeeWindow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeWindowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeWindow *FeeWindowRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeWindow.Contract.FeeWindowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeWindow *FeeWindowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeWindow.Contract.FeeWindowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeWindow *FeeWindowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeWindow.Contract.FeeWindowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeWindow *FeeWindowCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeWindow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeWindow *FeeWindowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeWindow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeWindow *FeeWindowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeWindow.Contract.contract.Transact(opts, method, params...)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) ETERNALAPPROVALVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "ETERNAL_APPROVAL_VALUE")
	return *ret0, err
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_FeeWindow *FeeWindowSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _FeeWindow.Contract.ETERNALAPPROVALVALUE(&_FeeWindow.CallOpts)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _FeeWindow.Contract.ETERNALAPPROVALVALUE(&_FeeWindow.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_FeeWindow *FeeWindowCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_FeeWindow *FeeWindowSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _FeeWindow.Contract.Allowance(&_FeeWindow.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_FeeWindow *FeeWindowCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _FeeWindow.Contract.Allowance(&_FeeWindow.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FeeWindow *FeeWindowSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _FeeWindow.Contract.BalanceOf(&_FeeWindow.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _FeeWindow.Contract.BalanceOf(&_FeeWindow.CallOpts, _owner)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_FeeWindow *FeeWindowCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_FeeWindow *FeeWindowSession) ControllerLookupName() ([32]byte, error) {
	return _FeeWindow.Contract.ControllerLookupName(&_FeeWindow.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_FeeWindow *FeeWindowCallerSession) ControllerLookupName() ([32]byte, error) {
	return _FeeWindow.Contract.ControllerLookupName(&_FeeWindow.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FeeWindow *FeeWindowCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FeeWindow *FeeWindowSession) Decimals() (uint8, error) {
	return _FeeWindow.Contract.Decimals(&_FeeWindow.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FeeWindow *FeeWindowCallerSession) Decimals() (uint8, error) {
	return _FeeWindow.Contract.Decimals(&_FeeWindow.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_FeeWindow *FeeWindowCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_FeeWindow *FeeWindowSession) GetController() (common.Address, error) {
	return _FeeWindow.Contract.GetController(&_FeeWindow.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_FeeWindow *FeeWindowCallerSession) GetController() (common.Address, error) {
	return _FeeWindow.Contract.GetController(&_FeeWindow.CallOpts)
}

// GetEndTime is a free data retrieval call binding the contract method 0x439f5ac2.
//
// Solidity: function getEndTime() constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) GetEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getEndTime")
	return *ret0, err
}

// GetEndTime is a free data retrieval call binding the contract method 0x439f5ac2.
//
// Solidity: function getEndTime() constant returns(uint256)
func (_FeeWindow *FeeWindowSession) GetEndTime() (*big.Int, error) {
	return _FeeWindow.Contract.GetEndTime(&_FeeWindow.CallOpts)
}

// GetEndTime is a free data retrieval call binding the contract method 0x439f5ac2.
//
// Solidity: function getEndTime() constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) GetEndTime() (*big.Int, error) {
	return _FeeWindow.Contract.GetEndTime(&_FeeWindow.CallOpts)
}

// GetFeeToken is a free data retrieval call binding the contract method 0xca709a25.
//
// Solidity: function getFeeToken() constant returns(address)
func (_FeeWindow *FeeWindowCaller) GetFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getFeeToken")
	return *ret0, err
}

// GetFeeToken is a free data retrieval call binding the contract method 0xca709a25.
//
// Solidity: function getFeeToken() constant returns(address)
func (_FeeWindow *FeeWindowSession) GetFeeToken() (common.Address, error) {
	return _FeeWindow.Contract.GetFeeToken(&_FeeWindow.CallOpts)
}

// GetFeeToken is a free data retrieval call binding the contract method 0xca709a25.
//
// Solidity: function getFeeToken() constant returns(address)
func (_FeeWindow *FeeWindowCallerSession) GetFeeToken() (common.Address, error) {
	return _FeeWindow.Contract.GetFeeToken(&_FeeWindow.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_FeeWindow *FeeWindowCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_FeeWindow *FeeWindowSession) GetInitialized() (bool, error) {
	return _FeeWindow.Contract.GetInitialized(&_FeeWindow.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_FeeWindow *FeeWindowCallerSession) GetInitialized() (bool, error) {
	return _FeeWindow.Contract.GetInitialized(&_FeeWindow.CallOpts)
}

// GetNumDesignatedReportNoShows is a free data retrieval call binding the contract method 0xda0b0c36.
//
// Solidity: function getNumDesignatedReportNoShows() constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) GetNumDesignatedReportNoShows(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getNumDesignatedReportNoShows")
	return *ret0, err
}

// GetNumDesignatedReportNoShows is a free data retrieval call binding the contract method 0xda0b0c36.
//
// Solidity: function getNumDesignatedReportNoShows() constant returns(uint256)
func (_FeeWindow *FeeWindowSession) GetNumDesignatedReportNoShows() (*big.Int, error) {
	return _FeeWindow.Contract.GetNumDesignatedReportNoShows(&_FeeWindow.CallOpts)
}

// GetNumDesignatedReportNoShows is a free data retrieval call binding the contract method 0xda0b0c36.
//
// Solidity: function getNumDesignatedReportNoShows() constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) GetNumDesignatedReportNoShows() (*big.Int, error) {
	return _FeeWindow.Contract.GetNumDesignatedReportNoShows(&_FeeWindow.CallOpts)
}

// GetNumIncorrectDesignatedReportMarkets is a free data retrieval call binding the contract method 0xa52c0512.
//
// Solidity: function getNumIncorrectDesignatedReportMarkets() constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) GetNumIncorrectDesignatedReportMarkets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getNumIncorrectDesignatedReportMarkets")
	return *ret0, err
}

// GetNumIncorrectDesignatedReportMarkets is a free data retrieval call binding the contract method 0xa52c0512.
//
// Solidity: function getNumIncorrectDesignatedReportMarkets() constant returns(uint256)
func (_FeeWindow *FeeWindowSession) GetNumIncorrectDesignatedReportMarkets() (*big.Int, error) {
	return _FeeWindow.Contract.GetNumIncorrectDesignatedReportMarkets(&_FeeWindow.CallOpts)
}

// GetNumIncorrectDesignatedReportMarkets is a free data retrieval call binding the contract method 0xa52c0512.
//
// Solidity: function getNumIncorrectDesignatedReportMarkets() constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) GetNumIncorrectDesignatedReportMarkets() (*big.Int, error) {
	return _FeeWindow.Contract.GetNumIncorrectDesignatedReportMarkets(&_FeeWindow.CallOpts)
}

// GetNumInvalidMarkets is a free data retrieval call binding the contract method 0xcf3d3849.
//
// Solidity: function getNumInvalidMarkets() constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) GetNumInvalidMarkets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getNumInvalidMarkets")
	return *ret0, err
}

// GetNumInvalidMarkets is a free data retrieval call binding the contract method 0xcf3d3849.
//
// Solidity: function getNumInvalidMarkets() constant returns(uint256)
func (_FeeWindow *FeeWindowSession) GetNumInvalidMarkets() (*big.Int, error) {
	return _FeeWindow.Contract.GetNumInvalidMarkets(&_FeeWindow.CallOpts)
}

// GetNumInvalidMarkets is a free data retrieval call binding the contract method 0xcf3d3849.
//
// Solidity: function getNumInvalidMarkets() constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) GetNumInvalidMarkets() (*big.Int, error) {
	return _FeeWindow.Contract.GetNumInvalidMarkets(&_FeeWindow.CallOpts)
}

// GetNumMarkets is a free data retrieval call binding the contract method 0x295c39a5.
//
// Solidity: function getNumMarkets() constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) GetNumMarkets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getNumMarkets")
	return *ret0, err
}

// GetNumMarkets is a free data retrieval call binding the contract method 0x295c39a5.
//
// Solidity: function getNumMarkets() constant returns(uint256)
func (_FeeWindow *FeeWindowSession) GetNumMarkets() (*big.Int, error) {
	return _FeeWindow.Contract.GetNumMarkets(&_FeeWindow.CallOpts)
}

// GetNumMarkets is a free data retrieval call binding the contract method 0x295c39a5.
//
// Solidity: function getNumMarkets() constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) GetNumMarkets() (*big.Int, error) {
	return _FeeWindow.Contract.GetNumMarkets(&_FeeWindow.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_FeeWindow *FeeWindowCaller) GetReputationToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getReputationToken")
	return *ret0, err
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_FeeWindow *FeeWindowSession) GetReputationToken() (common.Address, error) {
	return _FeeWindow.Contract.GetReputationToken(&_FeeWindow.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_FeeWindow *FeeWindowCallerSession) GetReputationToken() (common.Address, error) {
	return _FeeWindow.Contract.GetReputationToken(&_FeeWindow.CallOpts)
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) GetStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getStartTime")
	return *ret0, err
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() constant returns(uint256)
func (_FeeWindow *FeeWindowSession) GetStartTime() (*big.Int, error) {
	return _FeeWindow.Contract.GetStartTime(&_FeeWindow.CallOpts)
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) GetStartTime() (*big.Int, error) {
	return _FeeWindow.Contract.GetStartTime(&_FeeWindow.CallOpts)
}

// GetTotalFeeStake is a free data retrieval call binding the contract method 0x5c3cdec8.
//
// Solidity: function getTotalFeeStake() constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) GetTotalFeeStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getTotalFeeStake")
	return *ret0, err
}

// GetTotalFeeStake is a free data retrieval call binding the contract method 0x5c3cdec8.
//
// Solidity: function getTotalFeeStake() constant returns(uint256)
func (_FeeWindow *FeeWindowSession) GetTotalFeeStake() (*big.Int, error) {
	return _FeeWindow.Contract.GetTotalFeeStake(&_FeeWindow.CallOpts)
}

// GetTotalFeeStake is a free data retrieval call binding the contract method 0x5c3cdec8.
//
// Solidity: function getTotalFeeStake() constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) GetTotalFeeStake() (*big.Int, error) {
	return _FeeWindow.Contract.GetTotalFeeStake(&_FeeWindow.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_FeeWindow *FeeWindowCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_FeeWindow *FeeWindowSession) GetTypeName() ([32]byte, error) {
	return _FeeWindow.Contract.GetTypeName(&_FeeWindow.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_FeeWindow *FeeWindowCallerSession) GetTypeName() ([32]byte, error) {
	return _FeeWindow.Contract.GetTypeName(&_FeeWindow.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_FeeWindow *FeeWindowCaller) GetUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "getUniverse")
	return *ret0, err
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_FeeWindow *FeeWindowSession) GetUniverse() (common.Address, error) {
	return _FeeWindow.Contract.GetUniverse(&_FeeWindow.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_FeeWindow *FeeWindowCallerSession) GetUniverse() (common.Address, error) {
	return _FeeWindow.Contract.GetUniverse(&_FeeWindow.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_FeeWindow *FeeWindowCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "isActive")
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_FeeWindow *FeeWindowSession) IsActive() (bool, error) {
	return _FeeWindow.Contract.IsActive(&_FeeWindow.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_FeeWindow *FeeWindowCallerSession) IsActive() (bool, error) {
	return _FeeWindow.Contract.IsActive(&_FeeWindow.CallOpts)
}

// IsOver is a free data retrieval call binding the contract method 0xb4bd9e27.
//
// Solidity: function isOver() constant returns(bool)
func (_FeeWindow *FeeWindowCaller) IsOver(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "isOver")
	return *ret0, err
}

// IsOver is a free data retrieval call binding the contract method 0xb4bd9e27.
//
// Solidity: function isOver() constant returns(bool)
func (_FeeWindow *FeeWindowSession) IsOver() (bool, error) {
	return _FeeWindow.Contract.IsOver(&_FeeWindow.CallOpts)
}

// IsOver is a free data retrieval call binding the contract method 0xb4bd9e27.
//
// Solidity: function isOver() constant returns(bool)
func (_FeeWindow *FeeWindowCallerSession) IsOver() (bool, error) {
	return _FeeWindow.Contract.IsOver(&_FeeWindow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FeeWindow *FeeWindowCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FeeWindow *FeeWindowSession) Name() (string, error) {
	return _FeeWindow.Contract.Name(&_FeeWindow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FeeWindow *FeeWindowCallerSession) Name() (string, error) {
	return _FeeWindow.Contract.Name(&_FeeWindow.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FeeWindow *FeeWindowCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FeeWindow *FeeWindowSession) Symbol() (string, error) {
	return _FeeWindow.Contract.Symbol(&_FeeWindow.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FeeWindow *FeeWindowCallerSession) Symbol() (string, error) {
	return _FeeWindow.Contract.Symbol(&_FeeWindow.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FeeWindow *FeeWindowCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeWindow.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FeeWindow *FeeWindowSession) TotalSupply() (*big.Int, error) {
	return _FeeWindow.Contract.TotalSupply(&_FeeWindow.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FeeWindow *FeeWindowCallerSession) TotalSupply() (*big.Int, error) {
	return _FeeWindow.Contract.TotalSupply(&_FeeWindow.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_FeeWindow *FeeWindowSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.Approve(&_FeeWindow.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.Approve(&_FeeWindow.TransactOpts, _spender, _value)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_attotokens uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactor) Buy(opts *bind.TransactOpts, _attotokens *big.Int) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "buy", _attotokens)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_attotokens uint256) returns(bool)
func (_FeeWindow *FeeWindowSession) Buy(_attotokens *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.Buy(&_FeeWindow.TransactOpts, _attotokens)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_attotokens uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) Buy(_attotokens *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.Buy(&_FeeWindow.TransactOpts, _attotokens)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_FeeWindow *FeeWindowSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.DecreaseApproval(&_FeeWindow.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.DecreaseApproval(&_FeeWindow.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_FeeWindow *FeeWindowSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.IncreaseApproval(&_FeeWindow.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.IncreaseApproval(&_FeeWindow.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_universe address, _feeWindowId uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactor) Initialize(opts *bind.TransactOpts, _universe common.Address, _feeWindowId *big.Int) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "initialize", _universe, _feeWindowId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_universe address, _feeWindowId uint256) returns(bool)
func (_FeeWindow *FeeWindowSession) Initialize(_universe common.Address, _feeWindowId *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.Initialize(&_FeeWindow.TransactOpts, _universe, _feeWindowId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_universe address, _feeWindowId uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) Initialize(_universe common.Address, _feeWindowId *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.Initialize(&_FeeWindow.TransactOpts, _universe, _feeWindowId)
}

// MintFeeTokens is a paid mutator transaction binding the contract method 0xa7eb685b.
//
// Solidity: function mintFeeTokens(_amount uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactor) MintFeeTokens(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "mintFeeTokens", _amount)
}

// MintFeeTokens is a paid mutator transaction binding the contract method 0xa7eb685b.
//
// Solidity: function mintFeeTokens(_amount uint256) returns(bool)
func (_FeeWindow *FeeWindowSession) MintFeeTokens(_amount *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.MintFeeTokens(&_FeeWindow.TransactOpts, _amount)
}

// MintFeeTokens is a paid mutator transaction binding the contract method 0xa7eb685b.
//
// Solidity: function mintFeeTokens(_amount uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) MintFeeTokens(_amount *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.MintFeeTokens(&_FeeWindow.TransactOpts, _amount)
}

// OnMarketFinalized is a paid mutator transaction binding the contract method 0x8dc6e2f1.
//
// Solidity: function onMarketFinalized() returns(bool)
func (_FeeWindow *FeeWindowTransactor) OnMarketFinalized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "onMarketFinalized")
}

// OnMarketFinalized is a paid mutator transaction binding the contract method 0x8dc6e2f1.
//
// Solidity: function onMarketFinalized() returns(bool)
func (_FeeWindow *FeeWindowSession) OnMarketFinalized() (*types.Transaction, error) {
	return _FeeWindow.Contract.OnMarketFinalized(&_FeeWindow.TransactOpts)
}

// OnMarketFinalized is a paid mutator transaction binding the contract method 0x8dc6e2f1.
//
// Solidity: function onMarketFinalized() returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) OnMarketFinalized() (*types.Transaction, error) {
	return _FeeWindow.Contract.OnMarketFinalized(&_FeeWindow.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_sender address) returns(bool)
func (_FeeWindow *FeeWindowTransactor) Redeem(opts *bind.TransactOpts, _sender common.Address) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "redeem", _sender)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_sender address) returns(bool)
func (_FeeWindow *FeeWindowSession) Redeem(_sender common.Address) (*types.Transaction, error) {
	return _FeeWindow.Contract.Redeem(&_FeeWindow.TransactOpts, _sender)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_sender address) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) Redeem(_sender common.Address) (*types.Transaction, error) {
	return _FeeWindow.Contract.Redeem(&_FeeWindow.TransactOpts, _sender)
}

// RedeemForReportingParticipant is a paid mutator transaction binding the contract method 0xb97a6c12.
//
// Solidity: function redeemForReportingParticipant() returns(bool)
func (_FeeWindow *FeeWindowTransactor) RedeemForReportingParticipant(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "redeemForReportingParticipant")
}

// RedeemForReportingParticipant is a paid mutator transaction binding the contract method 0xb97a6c12.
//
// Solidity: function redeemForReportingParticipant() returns(bool)
func (_FeeWindow *FeeWindowSession) RedeemForReportingParticipant() (*types.Transaction, error) {
	return _FeeWindow.Contract.RedeemForReportingParticipant(&_FeeWindow.TransactOpts)
}

// RedeemForReportingParticipant is a paid mutator transaction binding the contract method 0xb97a6c12.
//
// Solidity: function redeemForReportingParticipant() returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) RedeemForReportingParticipant() (*types.Transaction, error) {
	return _FeeWindow.Contract.RedeemForReportingParticipant(&_FeeWindow.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_FeeWindow *FeeWindowTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_FeeWindow *FeeWindowSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _FeeWindow.Contract.SetController(&_FeeWindow.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _FeeWindow.Contract.SetController(&_FeeWindow.TransactOpts, _controller)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_FeeWindow *FeeWindowSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.Transfer(&_FeeWindow.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.Transfer(&_FeeWindow.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_FeeWindow *FeeWindowSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.TransferFrom(&_FeeWindow.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.TransferFrom(&_FeeWindow.TransactOpts, _from, _to, _value)
}

// TrustedUniverseBuy is a paid mutator transaction binding the contract method 0x7303ed18.
//
// Solidity: function trustedUniverseBuy(_buyer address, _attotokens uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactor) TrustedUniverseBuy(opts *bind.TransactOpts, _buyer common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _FeeWindow.contract.Transact(opts, "trustedUniverseBuy", _buyer, _attotokens)
}

// TrustedUniverseBuy is a paid mutator transaction binding the contract method 0x7303ed18.
//
// Solidity: function trustedUniverseBuy(_buyer address, _attotokens uint256) returns(bool)
func (_FeeWindow *FeeWindowSession) TrustedUniverseBuy(_buyer common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.TrustedUniverseBuy(&_FeeWindow.TransactOpts, _buyer, _attotokens)
}

// TrustedUniverseBuy is a paid mutator transaction binding the contract method 0x7303ed18.
//
// Solidity: function trustedUniverseBuy(_buyer address, _attotokens uint256) returns(bool)
func (_FeeWindow *FeeWindowTransactorSession) TrustedUniverseBuy(_buyer common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _FeeWindow.Contract.TrustedUniverseBuy(&_FeeWindow.TransactOpts, _buyer, _attotokens)
}

// FeeWindowApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FeeWindow contract.
type FeeWindowApprovalIterator struct {
	Event *FeeWindowApproval // Event containing the contract specifics and raw log

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
func (it *FeeWindowApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeWindowApproval)
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
		it.Event = new(FeeWindowApproval)
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
func (it *FeeWindowApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeWindowApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeWindowApproval represents a Approval event raised by the FeeWindow contract.
type FeeWindowApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_FeeWindow *FeeWindowFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FeeWindowApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FeeWindow.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FeeWindowApprovalIterator{contract: _FeeWindow.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_FeeWindow *FeeWindowFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FeeWindowApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FeeWindow.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeWindowApproval)
				if err := _FeeWindow.contract.UnpackLog(event, "Approval", log); err != nil {
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

// FeeWindowBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the FeeWindow contract.
type FeeWindowBurnIterator struct {
	Event *FeeWindowBurn // Event containing the contract specifics and raw log

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
func (it *FeeWindowBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeWindowBurn)
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
		it.Event = new(FeeWindowBurn)
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
func (it *FeeWindowBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeWindowBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeWindowBurn represents a Burn event raised by the FeeWindow contract.
type FeeWindowBurn struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_FeeWindow *FeeWindowFilterer) FilterBurn(opts *bind.FilterOpts, target []common.Address) (*FeeWindowBurnIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeWindow.contract.FilterLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return &FeeWindowBurnIterator{contract: _FeeWindow.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_FeeWindow *FeeWindowFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *FeeWindowBurn, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeWindow.contract.WatchLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeWindowBurn)
				if err := _FeeWindow.contract.UnpackLog(event, "Burn", log); err != nil {
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

// FeeWindowMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the FeeWindow contract.
type FeeWindowMintIterator struct {
	Event *FeeWindowMint // Event containing the contract specifics and raw log

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
func (it *FeeWindowMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeWindowMint)
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
		it.Event = new(FeeWindowMint)
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
func (it *FeeWindowMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeWindowMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeWindowMint represents a Mint event raised by the FeeWindow contract.
type FeeWindowMint struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_FeeWindow *FeeWindowFilterer) FilterMint(opts *bind.FilterOpts, target []common.Address) (*FeeWindowMintIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeWindow.contract.FilterLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return &FeeWindowMintIterator{contract: _FeeWindow.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_FeeWindow *FeeWindowFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *FeeWindowMint, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeWindow.contract.WatchLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeWindowMint)
				if err := _FeeWindow.contract.UnpackLog(event, "Mint", log); err != nil {
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

// FeeWindowTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FeeWindow contract.
type FeeWindowTransferIterator struct {
	Event *FeeWindowTransfer // Event containing the contract specifics and raw log

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
func (it *FeeWindowTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeWindowTransfer)
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
		it.Event = new(FeeWindowTransfer)
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
func (it *FeeWindowTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeWindowTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeWindowTransfer represents a Transfer event raised by the FeeWindow contract.
type FeeWindowTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_FeeWindow *FeeWindowFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeWindowTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeWindow.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeWindowTransferIterator{contract: _FeeWindow.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_FeeWindow *FeeWindowFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FeeWindowTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeWindow.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeWindowTransfer)
				if err := _FeeWindow.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// FeeWindowFactoryABI is the input ABI used to generate the binding from.
const FeeWindowFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"},{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_feeWindowId\",\"type\":\"uint256\"}],\"name\":\"createFeeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeeWindowFactoryBin is the compiled bytecode used for deploying new contracts.
const FeeWindowFactoryBin = `0x608060405234801561001057600080fd5b50610535806100206000396000f3006080604052600436106100405763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327237b1d8114610045575b600080fd5b34801561005157600080fd5b5061007c73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356100a5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000806000856100b36101f1565b73ffffffffffffffffffffffffffffffffffffffff90911681527f46656557696e646f7700000000000000000000000000000000000000000000006020820152604080519182900301906000f080158015610112573d6000803e3d6000fd5b5091508190508073ffffffffffffffffffffffffffffffffffffffff1663cd6dc68786866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156101bb57600080fd5b505af11580156101cf573d6000803e3d6000fd5b505050506040513d60208110156101e557600080fd5b50909695505050505050565b60405161030880610202833901905600608060405234801561001057600080fd5b506040516040806103088339810160405280516020909101516000805433600160a060020a03199182161716600160a060020a039093169290921782556001556102a890819061006090396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633018205f811461014f57806392eefe9b1461018d578063bef72fa2146101cf575b60015460009015156100675761014c565b60008054600154604080517ff39ec1f700000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff9092169263f39ec1f7926024808401936020939083900390910190829087803b1580156100df57600080fd5b505af11580156100f3573d6000803e3d6000fd5b505050506040513d602081101561010957600080fd5b50516040805136601f8101601f191680830190935292935091600083376000803684865af4808015610147576040513d81016040523d6000823e3d81f35b600080fd5b50005b34801561015b57600080fd5b506101646101f6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019957600080fd5b506101bb73ffffffffffffffffffffffffffffffffffffffff60043516610212565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610276565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461023757600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600154815600a165627a7a7230582017c3a6b57ac8bbc7fd5320cba2543fe371eb8726332cd58700252790e09cf2140029a165627a7a72305820f5b03e69d5106ff0df1caf9061a7b3afc27a3239f6538d985674f2e4184750a70029`

// DeployFeeWindowFactory deploys a new Ethereum contract, binding an instance of FeeWindowFactory to it.
func DeployFeeWindowFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FeeWindowFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeWindowFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeeWindowFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeWindowFactory{FeeWindowFactoryCaller: FeeWindowFactoryCaller{contract: contract}, FeeWindowFactoryTransactor: FeeWindowFactoryTransactor{contract: contract}, FeeWindowFactoryFilterer: FeeWindowFactoryFilterer{contract: contract}}, nil
}

// FeeWindowFactory is an auto generated Go binding around an Ethereum contract.
type FeeWindowFactory struct {
	FeeWindowFactoryCaller     // Read-only binding to the contract
	FeeWindowFactoryTransactor // Write-only binding to the contract
	FeeWindowFactoryFilterer   // Log filterer for contract events
}

// FeeWindowFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeWindowFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeWindowFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeWindowFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeWindowFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeWindowFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeWindowFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeWindowFactorySession struct {
	Contract     *FeeWindowFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeWindowFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeWindowFactoryCallerSession struct {
	Contract *FeeWindowFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FeeWindowFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeWindowFactoryTransactorSession struct {
	Contract     *FeeWindowFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FeeWindowFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeWindowFactoryRaw struct {
	Contract *FeeWindowFactory // Generic contract binding to access the raw methods on
}

// FeeWindowFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeWindowFactoryCallerRaw struct {
	Contract *FeeWindowFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FeeWindowFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeWindowFactoryTransactorRaw struct {
	Contract *FeeWindowFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeWindowFactory creates a new instance of FeeWindowFactory, bound to a specific deployed contract.
func NewFeeWindowFactory(address common.Address, backend bind.ContractBackend) (*FeeWindowFactory, error) {
	contract, err := bindFeeWindowFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeWindowFactory{FeeWindowFactoryCaller: FeeWindowFactoryCaller{contract: contract}, FeeWindowFactoryTransactor: FeeWindowFactoryTransactor{contract: contract}, FeeWindowFactoryFilterer: FeeWindowFactoryFilterer{contract: contract}}, nil
}

// NewFeeWindowFactoryCaller creates a new read-only instance of FeeWindowFactory, bound to a specific deployed contract.
func NewFeeWindowFactoryCaller(address common.Address, caller bind.ContractCaller) (*FeeWindowFactoryCaller, error) {
	contract, err := bindFeeWindowFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeWindowFactoryCaller{contract: contract}, nil
}

// NewFeeWindowFactoryTransactor creates a new write-only instance of FeeWindowFactory, bound to a specific deployed contract.
func NewFeeWindowFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeWindowFactoryTransactor, error) {
	contract, err := bindFeeWindowFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeWindowFactoryTransactor{contract: contract}, nil
}

// NewFeeWindowFactoryFilterer creates a new log filterer instance of FeeWindowFactory, bound to a specific deployed contract.
func NewFeeWindowFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeWindowFactoryFilterer, error) {
	contract, err := bindFeeWindowFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeWindowFactoryFilterer{contract: contract}, nil
}

// bindFeeWindowFactory binds a generic wrapper to an already deployed contract.
func bindFeeWindowFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeWindowFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeWindowFactory *FeeWindowFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeWindowFactory.Contract.FeeWindowFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeWindowFactory *FeeWindowFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeWindowFactory.Contract.FeeWindowFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeWindowFactory *FeeWindowFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeWindowFactory.Contract.FeeWindowFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeWindowFactory *FeeWindowFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeWindowFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeWindowFactory *FeeWindowFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeWindowFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeWindowFactory *FeeWindowFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeWindowFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateFeeWindow is a paid mutator transaction binding the contract method 0x27237b1d.
//
// Solidity: function createFeeWindow(_controller address, _universe address, _feeWindowId uint256) returns(address)
func (_FeeWindowFactory *FeeWindowFactoryTransactor) CreateFeeWindow(opts *bind.TransactOpts, _controller common.Address, _universe common.Address, _feeWindowId *big.Int) (*types.Transaction, error) {
	return _FeeWindowFactory.contract.Transact(opts, "createFeeWindow", _controller, _universe, _feeWindowId)
}

// CreateFeeWindow is a paid mutator transaction binding the contract method 0x27237b1d.
//
// Solidity: function createFeeWindow(_controller address, _universe address, _feeWindowId uint256) returns(address)
func (_FeeWindowFactory *FeeWindowFactorySession) CreateFeeWindow(_controller common.Address, _universe common.Address, _feeWindowId *big.Int) (*types.Transaction, error) {
	return _FeeWindowFactory.Contract.CreateFeeWindow(&_FeeWindowFactory.TransactOpts, _controller, _universe, _feeWindowId)
}

// CreateFeeWindow is a paid mutator transaction binding the contract method 0x27237b1d.
//
// Solidity: function createFeeWindow(_controller address, _universe address, _feeWindowId uint256) returns(address)
func (_FeeWindowFactory *FeeWindowFactoryTransactorSession) CreateFeeWindow(_controller common.Address, _universe common.Address, _feeWindowId *big.Int) (*types.Transaction, error) {
	return _FeeWindowFactory.Contract.CreateFeeWindow(&_FeeWindowFactory.TransactOpts, _controller, _universe, _feeWindowId)
}

/ IFeeWindowABI is the input ABI used to generate the binding from.
const IFeeWindowABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumMarkets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEndTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedUniverseBuy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"onMarketFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumIncorrectDesignatedReportMarkets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintFeeTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOver\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReputationToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"redeemForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStartTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"},{\"name\":\"_feeWindowId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumInvalidMarkets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumDesignatedReportNoShows\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// IFeeWindowBin is the compiled bytecode used for deploying new contracts.
const IFeeWindowBin = `0x`

// DeployIFeeWindow deploys a new Ethereum contract, binding an instance of IFeeWindow to it.
func DeployIFeeWindow(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IFeeWindow, error) {
	parsed, err := abi.JSON(strings.NewReader(IFeeWindowABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IFeeWindowBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IFeeWindow{IFeeWindowCaller: IFeeWindowCaller{contract: contract}, IFeeWindowTransactor: IFeeWindowTransactor{contract: contract}, IFeeWindowFilterer: IFeeWindowFilterer{contract: contract}}, nil
}

// IFeeWindow is an auto generated Go binding around an Ethereum contract.
type IFeeWindow struct {
	IFeeWindowCaller     // Read-only binding to the contract
	IFeeWindowTransactor // Write-only binding to the contract
	IFeeWindowFilterer   // Log filterer for contract events
}

// IFeeWindowCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFeeWindowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeWindowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFeeWindowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeWindowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFeeWindowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeWindowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFeeWindowSession struct {
	Contract     *IFeeWindow       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFeeWindowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFeeWindowCallerSession struct {
	Contract *IFeeWindowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IFeeWindowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFeeWindowTransactorSession struct {
	Contract     *IFeeWindowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IFeeWindowRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFeeWindowRaw struct {
	Contract *IFeeWindow // Generic contract binding to access the raw methods on
}

// IFeeWindowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFeeWindowCallerRaw struct {
	Contract *IFeeWindowCaller // Generic read-only contract binding to access the raw methods on
}

// IFeeWindowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFeeWindowTransactorRaw struct {
	Contract *IFeeWindowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFeeWindow creates a new instance of IFeeWindow, bound to a specific deployed contract.
func NewIFeeWindow(address common.Address, backend bind.ContractBackend) (*IFeeWindow, error) {
	contract, err := bindIFeeWindow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFeeWindow{IFeeWindowCaller: IFeeWindowCaller{contract: contract}, IFeeWindowTransactor: IFeeWindowTransactor{contract: contract}, IFeeWindowFilterer: IFeeWindowFilterer{contract: contract}}, nil
}

// NewIFeeWindowCaller creates a new read-only instance of IFeeWindow, bound to a specific deployed contract.
func NewIFeeWindowCaller(address common.Address, caller bind.ContractCaller) (*IFeeWindowCaller, error) {
	contract, err := bindIFeeWindow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeWindowCaller{contract: contract}, nil
}

// NewIFeeWindowTransactor creates a new write-only instance of IFeeWindow, bound to a specific deployed contract.
func NewIFeeWindowTransactor(address common.Address, transactor bind.ContractTransactor) (*IFeeWindowTransactor, error) {
	contract, err := bindIFeeWindow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeWindowTransactor{contract: contract}, nil
}

// NewIFeeWindowFilterer creates a new log filterer instance of IFeeWindow, bound to a specific deployed contract.
func NewIFeeWindowFilterer(address common.Address, filterer bind.ContractFilterer) (*IFeeWindowFilterer, error) {
	contract, err := bindIFeeWindow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFeeWindowFilterer{contract: contract}, nil
}

// bindIFeeWindow binds a generic wrapper to an already deployed contract.
func bindIFeeWindow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFeeWindowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeWindow *IFeeWindowRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IFeeWindow.Contract.IFeeWindowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeWindow *IFeeWindowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeWindow.Contract.IFeeWindowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeWindow *IFeeWindowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeWindow.Contract.IFeeWindowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeWindow *IFeeWindowCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IFeeWindow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeWindow *IFeeWindowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeWindow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeWindow *IFeeWindowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeWindow.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IFeeWindow *IFeeWindowCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IFeeWindow *IFeeWindowSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IFeeWindow.Contract.Allowance(&_IFeeWindow.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IFeeWindow *IFeeWindowCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IFeeWindow.Contract.Allowance(&_IFeeWindow.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IFeeWindow *IFeeWindowCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IFeeWindow *IFeeWindowSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _IFeeWindow.Contract.BalanceOf(&_IFeeWindow.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_IFeeWindow *IFeeWindowCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _IFeeWindow.Contract.BalanceOf(&_IFeeWindow.CallOpts, _who)
}

// GetEndTime is a free data retrieval call binding the contract method 0x439f5ac2.
//
// Solidity: function getEndTime() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCaller) GetEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getEndTime")
	return *ret0, err
}

// GetEndTime is a free data retrieval call binding the contract method 0x439f5ac2.
//
// Solidity: function getEndTime() constant returns(uint256)
func (_IFeeWindow *IFeeWindowSession) GetEndTime() (*big.Int, error) {
	return _IFeeWindow.Contract.GetEndTime(&_IFeeWindow.CallOpts)
}

// GetEndTime is a free data retrieval call binding the contract method 0x439f5ac2.
//
// Solidity: function getEndTime() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCallerSession) GetEndTime() (*big.Int, error) {
	return _IFeeWindow.Contract.GetEndTime(&_IFeeWindow.CallOpts)
}

// GetFeeToken is a free data retrieval call binding the contract method 0xca709a25.
//
// Solidity: function getFeeToken() constant returns(address)
func (_IFeeWindow *IFeeWindowCaller) GetFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getFeeToken")
	return *ret0, err
}

// GetFeeToken is a free data retrieval call binding the contract method 0xca709a25.
//
// Solidity: function getFeeToken() constant returns(address)
func (_IFeeWindow *IFeeWindowSession) GetFeeToken() (common.Address, error) {
	return _IFeeWindow.Contract.GetFeeToken(&_IFeeWindow.CallOpts)
}

// GetFeeToken is a free data retrieval call binding the contract method 0xca709a25.
//
// Solidity: function getFeeToken() constant returns(address)
func (_IFeeWindow *IFeeWindowCallerSession) GetFeeToken() (common.Address, error) {
	return _IFeeWindow.Contract.GetFeeToken(&_IFeeWindow.CallOpts)
}

// GetNumDesignatedReportNoShows is a free data retrieval call binding the contract method 0xda0b0c36.
//
// Solidity: function getNumDesignatedReportNoShows() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCaller) GetNumDesignatedReportNoShows(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getNumDesignatedReportNoShows")
	return *ret0, err
}

// GetNumDesignatedReportNoShows is a free data retrieval call binding the contract method 0xda0b0c36.
//
// Solidity: function getNumDesignatedReportNoShows() constant returns(uint256)
func (_IFeeWindow *IFeeWindowSession) GetNumDesignatedReportNoShows() (*big.Int, error) {
	return _IFeeWindow.Contract.GetNumDesignatedReportNoShows(&_IFeeWindow.CallOpts)
}

// GetNumDesignatedReportNoShows is a free data retrieval call binding the contract method 0xda0b0c36.
//
// Solidity: function getNumDesignatedReportNoShows() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCallerSession) GetNumDesignatedReportNoShows() (*big.Int, error) {
	return _IFeeWindow.Contract.GetNumDesignatedReportNoShows(&_IFeeWindow.CallOpts)
}

// GetNumIncorrectDesignatedReportMarkets is a free data retrieval call binding the contract method 0xa52c0512.
//
// Solidity: function getNumIncorrectDesignatedReportMarkets() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCaller) GetNumIncorrectDesignatedReportMarkets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getNumIncorrectDesignatedReportMarkets")
	return *ret0, err
}

// GetNumIncorrectDesignatedReportMarkets is a free data retrieval call binding the contract method 0xa52c0512.
//
// Solidity: function getNumIncorrectDesignatedReportMarkets() constant returns(uint256)
func (_IFeeWindow *IFeeWindowSession) GetNumIncorrectDesignatedReportMarkets() (*big.Int, error) {
	return _IFeeWindow.Contract.GetNumIncorrectDesignatedReportMarkets(&_IFeeWindow.CallOpts)
}

// GetNumIncorrectDesignatedReportMarkets is a free data retrieval call binding the contract method 0xa52c0512.
//
// Solidity: function getNumIncorrectDesignatedReportMarkets() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCallerSession) GetNumIncorrectDesignatedReportMarkets() (*big.Int, error) {
	return _IFeeWindow.Contract.GetNumIncorrectDesignatedReportMarkets(&_IFeeWindow.CallOpts)
}

// GetNumInvalidMarkets is a free data retrieval call binding the contract method 0xcf3d3849.
//
// Solidity: function getNumInvalidMarkets() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCaller) GetNumInvalidMarkets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getNumInvalidMarkets")
	return *ret0, err
}

// GetNumInvalidMarkets is a free data retrieval call binding the contract method 0xcf3d3849.
//
// Solidity: function getNumInvalidMarkets() constant returns(uint256)
func (_IFeeWindow *IFeeWindowSession) GetNumInvalidMarkets() (*big.Int, error) {
	return _IFeeWindow.Contract.GetNumInvalidMarkets(&_IFeeWindow.CallOpts)
}

// GetNumInvalidMarkets is a free data retrieval call binding the contract method 0xcf3d3849.
//
// Solidity: function getNumInvalidMarkets() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCallerSession) GetNumInvalidMarkets() (*big.Int, error) {
	return _IFeeWindow.Contract.GetNumInvalidMarkets(&_IFeeWindow.CallOpts)
}

// GetNumMarkets is a free data retrieval call binding the contract method 0x295c39a5.
//
// Solidity: function getNumMarkets() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCaller) GetNumMarkets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getNumMarkets")
	return *ret0, err
}

// GetNumMarkets is a free data retrieval call binding the contract method 0x295c39a5.
//
// Solidity: function getNumMarkets() constant returns(uint256)
func (_IFeeWindow *IFeeWindowSession) GetNumMarkets() (*big.Int, error) {
	return _IFeeWindow.Contract.GetNumMarkets(&_IFeeWindow.CallOpts)
}

// GetNumMarkets is a free data retrieval call binding the contract method 0x295c39a5.
//
// Solidity: function getNumMarkets() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCallerSession) GetNumMarkets() (*big.Int, error) {
	return _IFeeWindow.Contract.GetNumMarkets(&_IFeeWindow.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_IFeeWindow *IFeeWindowCaller) GetReputationToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getReputationToken")
	return *ret0, err
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_IFeeWindow *IFeeWindowSession) GetReputationToken() (common.Address, error) {
	return _IFeeWindow.Contract.GetReputationToken(&_IFeeWindow.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() constant returns(address)
func (_IFeeWindow *IFeeWindowCallerSession) GetReputationToken() (common.Address, error) {
	return _IFeeWindow.Contract.GetReputationToken(&_IFeeWindow.CallOpts)
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCaller) GetStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getStartTime")
	return *ret0, err
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() constant returns(uint256)
func (_IFeeWindow *IFeeWindowSession) GetStartTime() (*big.Int, error) {
	return _IFeeWindow.Contract.GetStartTime(&_IFeeWindow.CallOpts)
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCallerSession) GetStartTime() (*big.Int, error) {
	return _IFeeWindow.Contract.GetStartTime(&_IFeeWindow.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IFeeWindow *IFeeWindowCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IFeeWindow *IFeeWindowSession) GetTypeName() ([32]byte, error) {
	return _IFeeWindow.Contract.GetTypeName(&_IFeeWindow.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_IFeeWindow *IFeeWindowCallerSession) GetTypeName() ([32]byte, error) {
	return _IFeeWindow.Contract.GetTypeName(&_IFeeWindow.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_IFeeWindow *IFeeWindowCaller) GetUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "getUniverse")
	return *ret0, err
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_IFeeWindow *IFeeWindowSession) GetUniverse() (common.Address, error) {
	return _IFeeWindow.Contract.GetUniverse(&_IFeeWindow.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_IFeeWindow *IFeeWindowCallerSession) GetUniverse() (common.Address, error) {
	return _IFeeWindow.Contract.GetUniverse(&_IFeeWindow.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_IFeeWindow *IFeeWindowCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "isActive")
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_IFeeWindow *IFeeWindowSession) IsActive() (bool, error) {
	return _IFeeWindow.Contract.IsActive(&_IFeeWindow.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_IFeeWindow *IFeeWindowCallerSession) IsActive() (bool, error) {
	return _IFeeWindow.Contract.IsActive(&_IFeeWindow.CallOpts)
}

// IsOver is a free data retrieval call binding the contract method 0xb4bd9e27.
//
// Solidity: function isOver() constant returns(bool)
func (_IFeeWindow *IFeeWindowCaller) IsOver(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "isOver")
	return *ret0, err
}

// IsOver is a free data retrieval call binding the contract method 0xb4bd9e27.
//
// Solidity: function isOver() constant returns(bool)
func (_IFeeWindow *IFeeWindowSession) IsOver() (bool, error) {
	return _IFeeWindow.Contract.IsOver(&_IFeeWindow.CallOpts)
}

// IsOver is a free data retrieval call binding the contract method 0xb4bd9e27.
//
// Solidity: function isOver() constant returns(bool)
func (_IFeeWindow *IFeeWindowCallerSession) IsOver() (bool, error) {
	return _IFeeWindow.Contract.IsOver(&_IFeeWindow.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFeeWindow.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IFeeWindow *IFeeWindowSession) TotalSupply() (*big.Int, error) {
	return _IFeeWindow.Contract.TotalSupply(&_IFeeWindow.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IFeeWindow *IFeeWindowCallerSession) TotalSupply() (*big.Int, error) {
	return _IFeeWindow.Contract.TotalSupply(&_IFeeWindow.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IFeeWindow *IFeeWindowSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Approve(&_IFeeWindow.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Approve(&_IFeeWindow.TransactOpts, _spender, _value)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_attotokens uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) Buy(opts *bind.TransactOpts, _attotokens *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "buy", _attotokens)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_attotokens uint256) returns(bool)
func (_IFeeWindow *IFeeWindowSession) Buy(_attotokens *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Buy(&_IFeeWindow.TransactOpts, _attotokens)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(_attotokens uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) Buy(_attotokens *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Buy(&_IFeeWindow.TransactOpts, _attotokens)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_universe address, _feeWindowId uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) Initialize(opts *bind.TransactOpts, _universe common.Address, _feeWindowId *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "initialize", _universe, _feeWindowId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_universe address, _feeWindowId uint256) returns(bool)
func (_IFeeWindow *IFeeWindowSession) Initialize(_universe common.Address, _feeWindowId *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Initialize(&_IFeeWindow.TransactOpts, _universe, _feeWindowId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(_universe address, _feeWindowId uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) Initialize(_universe common.Address, _feeWindowId *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Initialize(&_IFeeWindow.TransactOpts, _universe, _feeWindowId)
}

// MintFeeTokens is a paid mutator transaction binding the contract method 0xa7eb685b.
//
// Solidity: function mintFeeTokens(_amount uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) MintFeeTokens(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "mintFeeTokens", _amount)
}

// MintFeeTokens is a paid mutator transaction binding the contract method 0xa7eb685b.
//
// Solidity: function mintFeeTokens(_amount uint256) returns(bool)
func (_IFeeWindow *IFeeWindowSession) MintFeeTokens(_amount *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.MintFeeTokens(&_IFeeWindow.TransactOpts, _amount)
}

// MintFeeTokens is a paid mutator transaction binding the contract method 0xa7eb685b.
//
// Solidity: function mintFeeTokens(_amount uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) MintFeeTokens(_amount *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.MintFeeTokens(&_IFeeWindow.TransactOpts, _amount)
}

// OnMarketFinalized is a paid mutator transaction binding the contract method 0x8dc6e2f1.
//
// Solidity: function onMarketFinalized() returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) OnMarketFinalized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "onMarketFinalized")
}

// OnMarketFinalized is a paid mutator transaction binding the contract method 0x8dc6e2f1.
//
// Solidity: function onMarketFinalized() returns(bool)
func (_IFeeWindow *IFeeWindowSession) OnMarketFinalized() (*types.Transaction, error) {
	return _IFeeWindow.Contract.OnMarketFinalized(&_IFeeWindow.TransactOpts)
}

// OnMarketFinalized is a paid mutator transaction binding the contract method 0x8dc6e2f1.
//
// Solidity: function onMarketFinalized() returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) OnMarketFinalized() (*types.Transaction, error) {
	return _IFeeWindow.Contract.OnMarketFinalized(&_IFeeWindow.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_sender address) returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) Redeem(opts *bind.TransactOpts, _sender common.Address) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "redeem", _sender)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_sender address) returns(bool)
func (_IFeeWindow *IFeeWindowSession) Redeem(_sender common.Address) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Redeem(&_IFeeWindow.TransactOpts, _sender)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(_sender address) returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) Redeem(_sender common.Address) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Redeem(&_IFeeWindow.TransactOpts, _sender)
}

// RedeemForReportingParticipant is a paid mutator transaction binding the contract method 0xb97a6c12.
//
// Solidity: function redeemForReportingParticipant() returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) RedeemForReportingParticipant(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "redeemForReportingParticipant")
}

// RedeemForReportingParticipant is a paid mutator transaction binding the contract method 0xb97a6c12.
//
// Solidity: function redeemForReportingParticipant() returns(bool)
func (_IFeeWindow *IFeeWindowSession) RedeemForReportingParticipant() (*types.Transaction, error) {
	return _IFeeWindow.Contract.RedeemForReportingParticipant(&_IFeeWindow.TransactOpts)
}

// RedeemForReportingParticipant is a paid mutator transaction binding the contract method 0xb97a6c12.
//
// Solidity: function redeemForReportingParticipant() returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) RedeemForReportingParticipant() (*types.Transaction, error) {
	return _IFeeWindow.Contract.RedeemForReportingParticipant(&_IFeeWindow.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IFeeWindow *IFeeWindowSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Transfer(&_IFeeWindow.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.Transfer(&_IFeeWindow.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IFeeWindow *IFeeWindowSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.TransferFrom(&_IFeeWindow.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.TransferFrom(&_IFeeWindow.TransactOpts, _from, _to, _value)
}

// TrustedUniverseBuy is a paid mutator transaction binding the contract method 0x7303ed18.
//
// Solidity: function trustedUniverseBuy(_buyer address, _attotokens uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactor) TrustedUniverseBuy(opts *bind.TransactOpts, _buyer common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.contract.Transact(opts, "trustedUniverseBuy", _buyer, _attotokens)
}

// TrustedUniverseBuy is a paid mutator transaction binding the contract method 0x7303ed18.
//
// Solidity: function trustedUniverseBuy(_buyer address, _attotokens uint256) returns(bool)
func (_IFeeWindow *IFeeWindowSession) TrustedUniverseBuy(_buyer common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.TrustedUniverseBuy(&_IFeeWindow.TransactOpts, _buyer, _attotokens)
}

// TrustedUniverseBuy is a paid mutator transaction binding the contract method 0x7303ed18.
//
// Solidity: function trustedUniverseBuy(_buyer address, _attotokens uint256) returns(bool)
func (_IFeeWindow *IFeeWindowTransactorSession) TrustedUniverseBuy(_buyer common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _IFeeWindow.Contract.TrustedUniverseBuy(&_IFeeWindow.TransactOpts, _buyer, _attotokens)
}

// IFeeWindowApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IFeeWindow contract.
type IFeeWindowApprovalIterator struct {
	Event *IFeeWindowApproval // Event containing the contract specifics and raw log

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
func (it *IFeeWindowApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeWindowApproval)
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
		it.Event = new(IFeeWindowApproval)
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
func (it *IFeeWindowApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeWindowApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeWindowApproval represents a Approval event raised by the IFeeWindow contract.
type IFeeWindowApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IFeeWindow *IFeeWindowFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IFeeWindowApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IFeeWindow.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IFeeWindowApprovalIterator{contract: _IFeeWindow.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IFeeWindow *IFeeWindowFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IFeeWindowApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IFeeWindow.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeWindowApproval)
				if err := _IFeeWindow.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IFeeWindowTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IFeeWindow contract.
type IFeeWindowTransferIterator struct {
	Event *IFeeWindowTransfer // Event containing the contract specifics and raw log

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
func (it *IFeeWindowTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeWindowTransfer)
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
		it.Event = new(IFeeWindowTransfer)
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
func (it *IFeeWindowTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeWindowTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeWindowTransfer represents a Transfer event raised by the IFeeWindow contract.
type IFeeWindowTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IFeeWindow *IFeeWindowFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IFeeWindowTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFeeWindow.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFeeWindowTransferIterator{contract: _IFeeWindow.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IFeeWindow *IFeeWindowFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IFeeWindowTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFeeWindow.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeWindowTransfer)
				if err := _IFeeWindow.contract.UnpackLog(event, "Transfer", log); err != nil {
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