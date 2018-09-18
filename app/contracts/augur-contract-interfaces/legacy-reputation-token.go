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

// LegacyReputationTokenABI is the input ABI used to generate the binding from.
const LegacyReputationTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalTheoreticalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_invalid\",\"type\":\"bool\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOutByPayout\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"faucet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETERNAL_APPROVAL_VALUE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOut\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLegacyRepToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIsMigratingFromLegacy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_holders\",\"type\":\"address[]\"}],\"name\":\"migrateBalancesFromLegacyRep\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedFeeWindowTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateIn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateParentTotalTheoreticalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedReportingParticipantTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"updateSiblingMigrationTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountMigrated\",\"type\":\"uint256\"}],\"name\":\"mintForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_spenders\",\"type\":\"address[]\"}],\"name\":\"migrateAllowancesFromLegacyRep\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedMarketTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedUniverseTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_universe\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_repBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"FundedAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// LegacyReputationTokenBin is the compiled bytecode used for deploying new contracts.
const LegacyReputationTokenBin = `0x60806040526002805460ff191690553480156200001b57600080fd5b5060008054600160a060020a031916331790556200005b731985365e9f78359a9b6ad760e32412f4a445e8626401000000006200250c6200006c82021704565b156200006657600080fd5b62000074565b6000903b1190565b61254080620000846000396000f3006080604052600436106101b35763ffffffff60e060020a60003504166306fdde0381146101b8578063095ea7b31461024257806318160ddd1461027a578063238d3590146102a157806323b872dd146102b65780633018205f146102e0578063313ce56714610311578063398c1a89146103265780635791589714610382578063634eaff11461039a57806366188463146103af5780636e7ce591146103d357806370a08231146103f757806377469275146104185780637cf99c331461042d5780637f68625914610442578063870c426d1461049757806390fbf84e146104ac57806391d76bbb146104d657806392eefe9b146104eb57806395d89b411461050c578063a0c1ca3414610521578063a819515d14610545578063a9059cbb1461055a578063b873e9a71461057e578063bef72fa2146105a8578063c4d66de8146105bd578063d73dd623146105de578063d9d3e07d14610602578063db05413414610623578063db0a087c1461063b578063dd62ed3e14610650578063de4c057414610677578063dea6aec714610705578063ee89dab41461071a578063f22b258a1461072f578063fe98184d14610759575b600080fd5b3480156101c457600080fd5b506101cd610783565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102075781810151838201526020016101ef565b50505050905090810190601f1680156102345780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024e57600080fd5b50610266600160a060020a03600435166024356107ba565b604080519115158252519081900360200190f35b34801561028657600080fd5b5061028f6107d1565b60408051918252519081900360200190f35b3480156102ad57600080fd5b5061028f6107d7565b3480156102c257600080fd5b50610266600160a060020a03600435811690602435166044356107dd565b3480156102ec57600080fd5b506102f5610803565b60408051600160a060020a039092168252519081900360200190f35b34801561031d57600080fd5b5061028f610812565b34801561033257600080fd5b506040805160206004803580820135838102808601850190965280855261026695369593946024949385019291829185019084908082843750949750505050823515159350505060200135610817565b34801561038e57600080fd5b50610266600435610a59565b3480156103a657600080fd5b5061028f610ac9565b3480156103bb57600080fd5b50610266600160a060020a0360043516602435610acf565b3480156103df57600080fd5b50610266600160a060020a0360043516602435610b35565b34801561040357600080fd5b5061028f600160a060020a0360043516610c0d565b34801561042457600080fd5b506102f5610c28565b34801561043957600080fd5b50610266610cdf565b34801561044e57600080fd5b506040805160206004803580820135838102808601850190965280855261026695369593946024949385019291829185019084908082843750949750610ce89650505050505050565b3480156104a357600080fd5b506102f5610d53565b3480156104b857600080fd5b50610266600160a060020a0360043581169060243516604435610d62565b3480156104e257600080fd5b5061028f610e2e565b3480156104f757600080fd5b50610266600160a060020a0360043516610e34565b34801561051857600080fd5b506101cd610e7e565b34801561052d57600080fd5b50610266600160a060020a0360043516602435610eb5565b34801561055157600080fd5b5061026661134b565b34801561056657600080fd5b50610266600160a060020a0360043516602435611518565b34801561058a57600080fd5b50610266600160a060020a036004358116906024351660443561153c565b3480156105b457600080fd5b5061028f6115c6565b3480156105c957600080fd5b50610266600160a060020a03600435166115cc565b3480156105ea57600080fd5b50610266600160a060020a036004351660243561175a565b34801561060e57600080fd5b50610266600160a060020a0360043516611796565b34801561062f57600080fd5b50610266600435611b1f565b34801561064757600080fd5b5061028f611c9b565b34801561065c57600080fd5b5061028f600160a060020a0360043581169060243516611cbf565b34801561068357600080fd5b506040805160206004803580820135838102808601850190965280855261026695369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750611cea9650505050505050565b34801561071157600080fd5b5061028f611e13565b34801561072657600080fd5b50610266611e19565b34801561073b57600080fd5b50610266600160a060020a0360043581169060243516604435611e22565b34801561076557600080fd5b50610266600160a060020a0360043581169060243516604435611eac565b60408051808201909152600a81527f52657075746174696f6e00000000000000000000000000000000000000000000602082015281565b60006107c7338484611ee7565b5060019392505050565b60035490565b600a5490565b600b5460009060ff16156107f057600080fd5b6107fb848484611f52565b949350505050565b600054600160a060020a031690565b601281565b600b546000908190819060ff161561082e57600080fd5b60025460ff16151561083f57600080fd5b6000841161084c57600080fd5b600654604080517fdf428e3b000000000000000000000000000000000000000000000000000000008152871515602482015260048101918252885160448201528851600160a060020a039093169263df428e3b928a928a928291606401906020808701910280838360005b838110156108cf5781810151838201526020016108b7565b505050509050019350505050602060405180830381600087803b1580156108f557600080fd5b505af1158015610909573d6000803e3d6000fd5b505050506040513d602081101561091f57600080fd5b5051604080517fb80907f20000000000000000000000000000000000000000000000000000000081529051919350600160a060020a0384169163b80907f2916004808201926020929091908290030181600087803b15801561098057600080fd5b505af1158015610994573d6000803e3d6000fd5b505050506040513d60208110156109aa57600080fd5b505190506109b83385611fcb565b50604080517fa0c1ca34000000000000000000000000000000000000000000000000000000008152336004820152602481018690529051600160a060020a0383169163a0c1ca349160448083019260209291908290030181600087803b158015610a2157600080fd5b505af1158015610a35573d6000803e3d6000fd5b505050506040513d6020811015610a4b57600080fd5b506001979650505050505050565b60007001000000000000000000000000000000008210610a7857600080fd5b610a82338361206c565b50604080518381524260208201528151339230927fbf88b5a3158512dc8ad44dadf221f46eb1df762a2e7fd1de86aab34b0af455cf929081900390910190a3506001919050565b60001981565b336000908152600560209081526040808320600160a060020a038616845290915281205480831115610b0d57610b0733856000611ee7565b50610b29565b610b273385610b22848763ffffffff61210d16565b611ee7565b505b600191505b5092915050565b600b5460009060ff1615610b4857600080fd5b60025460ff161515610b5957600080fd5b60008211610b6657600080fd5b610b6f83612122565b50610b7a3383611fcb565b50604080517fa0c1ca34000000000000000000000000000000000000000000000000000000008152336004820152602481018490529051600160a060020a0385169163a0c1ca349160448083019260209291908290030181600087803b158015610be357600080fd5b505af1158015610bf7573d6000803e3d6000fd5b505050506040513d6020811015610b2757600080fd5b600160a060020a031660009081526004602052604090205490565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4c656761637952657075746174696f6e546f6b656e000000000000000000000060048201529051600160a060020a039092169163f39ec1f79160248082019260209290919082900301818787803b158015610cae57600080fd5b505af1158015610cc2573d6000803e3d6000fd5b505050506040513d6020811015610cd857600080fd5b5051905090565b600b5460ff1690565b600b546000908190819060ff161515610d0057600080fd5b60025460ff161515610d1157600080fd5b610d19610c28565b9150600090505b83518110156107c757610d4a8482815181101515610d3a57fe5b90602001906020020151836122bc565b50600101610d20565b600654600160a060020a031690565b600b5460009060ff1615610d7557600080fd5b60025460ff161515610d8657600080fd5b600654604080517fc7c88d700000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163c7c88d70916024808201926020929091908290030181600087803b158015610dec57600080fd5b505af1158015610e00573d6000803e3d6000fd5b505050506040513d6020811015610e1657600080fd5b50511515610e2357600080fd5b6107fb8484846123cf565b60075490565b60008054600160a060020a03163314610e4c57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60408051808201909152600381527f5245500000000000000000000000000000000000000000000000000000000000602082015281565b600b546000908190819060ff1615610ecc57600080fd5b60025460ff161515610edd57600080fd5b600660009054906101000a9004600160a060020a0316600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610f3057600080fd5b505af1158015610f44573d6000803e3d6000fd5b505050506040513d6020811015610f5a57600080fd5b5051604080517fb80907f20000000000000000000000000000000000000000000000000000000081529051919350600160a060020a0384169163b80907f2916004808201926020929091908290030181600087803b158015610fbb57600080fd5b505af1158015610fcf573d6000803e3d6000fd5b505050506040513d6020811015610fe557600080fd5b5051600160a060020a03163314610ffb57600080fd5b611005858561206c565b506007805485019055604080517f77e71ee50000000000000000000000000000000000000000000000000000000081529051600160a060020a038416916377e71ee59160048083019260209291908290030181600087803b15801561106957600080fd5b505af115801561107d573d6000803e3d6000fd5b505050506040513d602081101561109357600080fd5b505160008054604080517f188ec3560000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169263188ec356926004808401936020939083900390910190829087803b1580156110f657600080fd5b505af115801561110a573d6000803e3d6000fd5b505050506040513d602081101561112057600080fd5b5051101561114f5761113984601463ffffffff61248f16565b9050611145858261206c565b50600a8054820190555b81600160a060020a031663cb1d84186040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561118d57600080fd5b505af11580156111a1573d6000803e3d6000fd5b505050506040513d60208110156111b757600080fd5b5051604080517f8d4e40830000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921691638d4e4083916004808201926020929091908290030181600087803b15801561121657600080fd5b505af115801561122a573d6000803e3d6000fd5b505050506040513d602081101561124057600080fd5b505115156113405781600160a060020a031663f7095d9d600660009054906101000a9004600160a060020a0316600160a060020a031663c38c0fa76040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156112aa57600080fd5b505af11580156112be573d6000803e3d6000fd5b505050506040513d60208110156112d457600080fd5b50516040805160e060020a63ffffffff851602815260048101929092525160248083019260209291908290030181600087803b15801561131357600080fd5b505af1158015611327573d6000803e3d6000fd5b505050506040513d602081101561133d57600080fd5b50505b506001949350505050565b600b54600090819060ff161561136057600080fd5b600660009054906101000a9004600160a060020a0316600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156113b357600080fd5b505af11580156113c7573d6000803e3d6000fd5b505050506040513d60208110156113dd57600080fd5b5051600954600a80549190910390559050600160a060020a038116151561140e576114066124a6565b600955611505565b80600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561144c57600080fd5b505af1158015611460573d6000803e3d6000fd5b505050506040513d602081101561147657600080fd5b5051604080517f238d35900000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169163238d3590916004808201926020929091908290030181600087803b1580156114d557600080fd5b505af11580156114e9573d6000803e3d6000fd5b505050506040513d60208110156114ff57600080fd5b50516009555b5050600954600a80549091019055600190565b600b5460009060ff161561152b57600080fd5b61153583836124b5565b9392505050565b600b5460009060ff161561154f57600080fd5b60025460ff16151561156057600080fd5b600654604080517ff76514c70000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163f76514c7916024808201926020929091908290030181600087803b158015610dec57600080fd5b60015481565b600254600090819060ff16156115e157600080fd5b6115e96124c2565b50600160a060020a03831615156115ff57600080fd5b6006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851617905561162f61134b565b50611638610c28565b90506000600160a060020a031683600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561168357600080fd5b505af1158015611697573d6000803e3d6000fd5b505050506040513d60208110156116ad57600080fd5b5051600b805460ff1916600160a060020a039283169390931492909217909155604080517f18160ddd0000000000000000000000000000000000000000000000000000000081529051918316916318160ddd916004808201926020929091908290030181600087803b15801561172257600080fd5b505af1158015611736573d6000803e3d6000fd5b505050506040513d602081101561174c57600080fd5b5051600c5550600192915050565b336000818152600560209081526040808320600160a060020a038716845290915281205490916107c7918590610b22908663ffffffff6124e916565b600b54600090819060ff16156117ab57600080fd5b600160a060020a0383163014156117c157600080fd5b82600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156117ff57600080fd5b505af1158015611813573d6000803e3d6000fd5b505050506040513d602081101561182957600080fd5b5051600654604080517fa63f13500000000000000000000000000000000000000000000000000000000081529051929350600160a060020a039091169163a63f1350916004808201926020929091908290030181600087803b15801561188e57600080fd5b505af11580156118a2573d6000803e3d6000fd5b505050506040513d60208110156118b857600080fd5b5051604080517fc38c0fa70000000000000000000000000000000000000000000000000000000081529051600160a060020a039283169263eceba876929085169163c38c0fa7916004808201926020929091908290030181600087803b15801561192157600080fd5b505af1158015611935573d6000803e3d6000fd5b505050506040513d602081101561194b57600080fd5b50516040805160e060020a63ffffffff851602815260048101929092525160248083019260209291908290030181600087803b15801561198a57600080fd5b505af115801561199e573d6000803e3d6000fd5b505050506040513d60208110156119b457600080fd5b5051604080517fb80907f20000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169163b80907f2916004808201926020929091908290030181600087803b158015611a1357600080fd5b505af1158015611a27573d6000803e3d6000fd5b505050506040513d6020811015611a3d57600080fd5b5051600160a060020a03848116911614611a5657600080fd5b600160a060020a038316600081815260086020908152604080832054600a8054909101905580517f91d76bbb00000000000000000000000000000000000000000000000000000000815290516391d76bbb93600480840194938390030190829087803b158015611ac557600080fd5b505af1158015611ad9573d6000803e3d6000fd5b505050506040513d6020811015611aef57600080fd5b5051600160a060020a039390931660009081526008602052604090208390555050600a8054919091039055600190565b600b5460009081908190819060ff1615611b3857600080fd5b60025460ff161515611b4957600080fd5b600660009054906101000a9004600160a060020a0316600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611b9c57600080fd5b505af1158015611bb0573d6000803e3d6000fd5b505050506040513d6020811015611bc657600080fd5b5051604080517ff76514c700000000000000000000000000000000000000000000000000000000815233600482018190529151929550909350600160a060020a0385169163f76514c7916024808201926020929091908290030181600087803b158015611c3257600080fd5b505af1158015611c46573d6000803e3d6000fd5b505050506040513d6020811015611c5c57600080fd5b50511515611c6957600080fd5b611c7a85600263ffffffff61248f16565b9050611c86828261206c565b50600a80548201905560019350505050919050565b7f52657075746174696f6e546f6b656e000000000000000000000000000000000090565b600160a060020a03918216600090815260056020908152604080832093909416825291909152205490565b600b546000908190819081908190819060ff161515611d0857600080fd5b60025460ff161515611d1957600080fd5b611d21610c28565b9450600093505b8751841015610a4b578784815181101515611d3f57fe5b9060200190602002015192508684815181101515611d5957fe5b6020908102909101810151604080517fdd62ed3e000000000000000000000000000000000000000000000000000000008152600160a060020a038781166004830152808416602483015291519295509088169263dd62ed3e926044808401938290030181600087803b158015611dce57600080fd5b505af1158015611de2573d6000803e3d6000fd5b505050506040513d6020811015611df857600080fd5b50519050611e07838383611ee7565b50600190930192611d28565b600c5490565b60025460ff1690565b600b5460009060ff1615611e3557600080fd5b60025460ff161515611e4657600080fd5b600654604080517f9f7e1bf60000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691639f7e1bf6916024808201926020929091908290030181600087803b158015610dec57600080fd5b600b5460009060ff1615611ebf57600080fd5b60025460ff161515611ed057600080fd5b600654600160a060020a03163314610e2357600080fd5b600160a060020a03808416600081815260056020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a03831660009081526005602090815260408083203384529091528120546000198114611fb457611f8f818463ffffffff61210d16565b600160a060020a03861660009081526005602090815260408083203384529091529020555b611fbf8585856123cf565b50600195945050505050565b600160a060020a038216600090815260046020526040812054611ff4908363ffffffff61210d16565b600160a060020a038416600090815260046020526040902055600354612020908363ffffffff61210d16565b600355604080518381529051600160a060020a038516917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a26107c783836124fb565b600160a060020a038216600090815260046020526040812054612095908363ffffffff6124e916565b600160a060020a0384166000908152600460205260409020556003546120c1908363ffffffff6124e916565b600355604080518381529051600160a060020a038516917f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885919081900360200190a26107c783836124fb565b60008282111561211c57600080fd5b50900390565b600080600083600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561216557600080fd5b505af1158015612179573d6000803e3d6000fd5b505050506040513d602081101561218f57600080fd5b5051600654604080517f9517317c000000000000000000000000000000000000000000000000000000008152600160a060020a0380851660048301529151939550911691639517317c916024808201926020929091908290030181600087803b1580156121fb57600080fd5b505af115801561220f573d6000803e3d6000fd5b505050506040513d602081101561222557600080fd5b5051151561223257600080fd5b81905083600160a060020a031681600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561227d57600080fd5b505af1158015612291573d6000803e3d6000fd5b505050506040513d60208110156122a757600080fd5b5051600160a060020a0316146107c757600080fd5b600b54600090819060ff1615156122d257600080fd5b60025460ff1615156122e357600080fd5b600160a060020a038416600090815260046020526040812054111561230b5760009150610b2e565b82600160a060020a03166370a08231856040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561236657600080fd5b505af115801561237a573d6000803e3d6000fd5b505050506040513d602081101561239057600080fd5b505190508015156123a45760009150610b2e565b6123ae848261206c565b50600354600c541415610b29575050600b805460ff19169055506001919050565b600160a060020a0383166000908152600460205260408120546123f8908363ffffffff61210d16565b600160a060020a03808616600090815260046020526040808220939093559085168152205461242d908363ffffffff6124e916565b600160a060020a0380851660008181526004602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3611340848484612503565b600080828481151561249d57fe5b04949350505050565b6a09195731e2ce35eb00000090565b60006115353384846123cf565b60025460009060ff16156124d557600080fd5b506002805460ff1916600190811790915590565b60008282018381101561153557600080fd5b600192915050565b60019392505050565b6000903b11905600a165627a7a72305820e58ed6a2e21e227a5a7b5301e84f18050e1782b6972607bc7bb4475fab17d63e0029`

// DeployLegacyReputationToken deploys a new Ethereum contract, binding an instance of LegacyReputationToken to it.
func DeployLegacyReputationToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LegacyReputationToken, error) {
	parsed, err := abi.JSON(strings.NewReader(LegacyReputationTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LegacyReputationTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LegacyReputationToken{LegacyReputationTokenCaller: LegacyReputationTokenCaller{contract: contract}, LegacyReputationTokenTransactor: LegacyReputationTokenTransactor{contract: contract}, LegacyReputationTokenFilterer: LegacyReputationTokenFilterer{contract: contract}}, nil
}

// LegacyReputationToken is an auto generated Go binding around an Ethereum contract.
type LegacyReputationToken struct {
	LegacyReputationTokenCaller     // Read-only binding to the contract
	LegacyReputationTokenTransactor // Write-only binding to the contract
	LegacyReputationTokenFilterer   // Log filterer for contract events
}

// LegacyReputationTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyReputationTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyReputationTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyReputationTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyReputationTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyReputationTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyReputationTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyReputationTokenSession struct {
	Contract     *LegacyReputationToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LegacyReputationTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyReputationTokenCallerSession struct {
	Contract *LegacyReputationTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// LegacyReputationTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyReputationTokenTransactorSession struct {
	Contract     *LegacyReputationTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// LegacyReputationTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyReputationTokenRaw struct {
	Contract *LegacyReputationToken // Generic contract binding to access the raw methods on
}

// LegacyReputationTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyReputationTokenCallerRaw struct {
	Contract *LegacyReputationTokenCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyReputationTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyReputationTokenTransactorRaw struct {
	Contract *LegacyReputationTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyReputationToken creates a new instance of LegacyReputationToken, bound to a specific deployed contract.
func NewLegacyReputationToken(address common.Address, backend bind.ContractBackend) (*LegacyReputationToken, error) {
	contract, err := bindLegacyReputationToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LegacyReputationToken{LegacyReputationTokenCaller: LegacyReputationTokenCaller{contract: contract}, LegacyReputationTokenTransactor: LegacyReputationTokenTransactor{contract: contract}, LegacyReputationTokenFilterer: LegacyReputationTokenFilterer{contract: contract}}, nil
}

// NewLegacyReputationTokenCaller creates a new read-only instance of LegacyReputationToken, bound to a specific deployed contract.
func NewLegacyReputationTokenCaller(address common.Address, caller bind.ContractCaller) (*LegacyReputationTokenCaller, error) {
	contract, err := bindLegacyReputationToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyReputationTokenCaller{contract: contract}, nil
}

// NewLegacyReputationTokenTransactor creates a new write-only instance of LegacyReputationToken, bound to a specific deployed contract.
func NewLegacyReputationTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyReputationTokenTransactor, error) {
	contract, err := bindLegacyReputationToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyReputationTokenTransactor{contract: contract}, nil
}

// NewLegacyReputationTokenFilterer creates a new log filterer instance of LegacyReputationToken, bound to a specific deployed contract.
func NewLegacyReputationTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyReputationTokenFilterer, error) {
	contract, err := bindLegacyReputationToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyReputationTokenFilterer{contract: contract}, nil
}

// bindLegacyReputationToken binds a generic wrapper to an already deployed contract.
func bindLegacyReputationToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LegacyReputationTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyReputationToken *LegacyReputationTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LegacyReputationToken.Contract.LegacyReputationTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyReputationToken *LegacyReputationTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.LegacyReputationTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyReputationToken *LegacyReputationTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.LegacyReputationTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyReputationToken *LegacyReputationTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LegacyReputationToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyReputationToken *LegacyReputationTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyReputationToken *LegacyReputationTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.contract.Transact(opts, method, params...)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCaller) ETERNALAPPROVALVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "ETERNAL_APPROVAL_VALUE")
	return *ret0, err
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _LegacyReputationToken.Contract.ETERNALAPPROVALVALUE(&_LegacyReputationToken.CallOpts)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _LegacyReputationToken.Contract.ETERNALAPPROVALVALUE(&_LegacyReputationToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_LegacyReputationToken *LegacyReputationTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_LegacyReputationToken *LegacyReputationTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LegacyReputationToken.Contract.Allowance(&_LegacyReputationToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LegacyReputationToken.Contract.Allowance(&_LegacyReputationToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LegacyReputationToken.Contract.BalanceOf(&_LegacyReputationToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LegacyReputationToken.Contract.BalanceOf(&_LegacyReputationToken.CallOpts, _owner)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_LegacyReputationToken *LegacyReputationTokenCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_LegacyReputationToken *LegacyReputationTokenSession) ControllerLookupName() ([32]byte, error) {
	return _LegacyReputationToken.Contract.ControllerLookupName(&_LegacyReputationToken.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) ControllerLookupName() ([32]byte, error) {
	return _LegacyReputationToken.Contract.ControllerLookupName(&_LegacyReputationToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenSession) Decimals() (*big.Int, error) {
	return _LegacyReputationToken.Contract.Decimals(&_LegacyReputationToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) Decimals() (*big.Int, error) {
	return _LegacyReputationToken.Contract.Decimals(&_LegacyReputationToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_LegacyReputationToken *LegacyReputationTokenCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_LegacyReputationToken *LegacyReputationTokenSession) GetController() (common.Address, error) {
	return _LegacyReputationToken.Contract.GetController(&_LegacyReputationToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) GetController() (common.Address, error) {
	return _LegacyReputationToken.Contract.GetController(&_LegacyReputationToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) GetInitialized() (bool, error) {
	return _LegacyReputationToken.Contract.GetInitialized(&_LegacyReputationToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) GetInitialized() (bool, error) {
	return _LegacyReputationToken.Contract.GetInitialized(&_LegacyReputationToken.CallOpts)
}

// GetIsMigratingFromLegacy is a free data retrieval call binding the contract method 0x7cf99c33.
//
// Solidity: function getIsMigratingFromLegacy() constant returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenCaller) GetIsMigratingFromLegacy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "getIsMigratingFromLegacy")
	return *ret0, err
}

// GetIsMigratingFromLegacy is a free data retrieval call binding the contract method 0x7cf99c33.
//
// Solidity: function getIsMigratingFromLegacy() constant returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) GetIsMigratingFromLegacy() (bool, error) {
	return _LegacyReputationToken.Contract.GetIsMigratingFromLegacy(&_LegacyReputationToken.CallOpts)
}

// GetIsMigratingFromLegacy is a free data retrieval call binding the contract method 0x7cf99c33.
//
// Solidity: function getIsMigratingFromLegacy() constant returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) GetIsMigratingFromLegacy() (bool, error) {
	return _LegacyReputationToken.Contract.GetIsMigratingFromLegacy(&_LegacyReputationToken.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() constant returns(address)
func (_LegacyReputationToken *LegacyReputationTokenCaller) GetLegacyRepToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "getLegacyRepToken")
	return *ret0, err
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() constant returns(address)
func (_LegacyReputationToken *LegacyReputationTokenSession) GetLegacyRepToken() (common.Address, error) {
	return _LegacyReputationToken.Contract.GetLegacyRepToken(&_LegacyReputationToken.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() constant returns(address)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) GetLegacyRepToken() (common.Address, error) {
	return _LegacyReputationToken.Contract.GetLegacyRepToken(&_LegacyReputationToken.CallOpts)
}

// GetTargetSupply is a free data retrieval call binding the contract method 0xdea6aec7.
//
// Solidity: function getTargetSupply() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCaller) GetTargetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "getTargetSupply")
	return *ret0, err
}

// GetTargetSupply is a free data retrieval call binding the contract method 0xdea6aec7.
//
// Solidity: function getTargetSupply() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenSession) GetTargetSupply() (*big.Int, error) {
	return _LegacyReputationToken.Contract.GetTargetSupply(&_LegacyReputationToken.CallOpts)
}

// GetTargetSupply is a free data retrieval call binding the contract method 0xdea6aec7.
//
// Solidity: function getTargetSupply() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) GetTargetSupply() (*big.Int, error) {
	return _LegacyReputationToken.Contract.GetTargetSupply(&_LegacyReputationToken.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCaller) GetTotalMigrated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "getTotalMigrated")
	return *ret0, err
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenSession) GetTotalMigrated() (*big.Int, error) {
	return _LegacyReputationToken.Contract.GetTotalMigrated(&_LegacyReputationToken.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) GetTotalMigrated() (*big.Int, error) {
	return _LegacyReputationToken.Contract.GetTotalMigrated(&_LegacyReputationToken.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCaller) GetTotalTheoreticalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "getTotalTheoreticalSupply")
	return *ret0, err
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _LegacyReputationToken.Contract.GetTotalTheoreticalSupply(&_LegacyReputationToken.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _LegacyReputationToken.Contract.GetTotalTheoreticalSupply(&_LegacyReputationToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_LegacyReputationToken *LegacyReputationTokenCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_LegacyReputationToken *LegacyReputationTokenSession) GetTypeName() ([32]byte, error) {
	return _LegacyReputationToken.Contract.GetTypeName(&_LegacyReputationToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) GetTypeName() ([32]byte, error) {
	return _LegacyReputationToken.Contract.GetTypeName(&_LegacyReputationToken.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_LegacyReputationToken *LegacyReputationTokenCaller) GetUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "getUniverse")
	return *ret0, err
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_LegacyReputationToken *LegacyReputationTokenSession) GetUniverse() (common.Address, error) {
	return _LegacyReputationToken.Contract.GetUniverse(&_LegacyReputationToken.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) GetUniverse() (common.Address, error) {
	return _LegacyReputationToken.Contract.GetUniverse(&_LegacyReputationToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_LegacyReputationToken *LegacyReputationTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_LegacyReputationToken *LegacyReputationTokenSession) Name() (string, error) {
	return _LegacyReputationToken.Contract.Name(&_LegacyReputationToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) Name() (string, error) {
	return _LegacyReputationToken.Contract.Name(&_LegacyReputationToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_LegacyReputationToken *LegacyReputationTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_LegacyReputationToken *LegacyReputationTokenSession) Symbol() (string, error) {
	return _LegacyReputationToken.Contract.Symbol(&_LegacyReputationToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) Symbol() (string, error) {
	return _LegacyReputationToken.Contract.Symbol(&_LegacyReputationToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LegacyReputationToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenSession) TotalSupply() (*big.Int, error) {
	return _LegacyReputationToken.Contract.TotalSupply(&_LegacyReputationToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_LegacyReputationToken *LegacyReputationTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _LegacyReputationToken.Contract.TotalSupply(&_LegacyReputationToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.Approve(&_LegacyReputationToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.Approve(&_LegacyReputationToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.DecreaseApproval(&_LegacyReputationToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.DecreaseApproval(&_LegacyReputationToken.TransactOpts, _spender, _subtractedValue)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(_amount uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) Faucet(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "faucet", _amount)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(_amount uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) Faucet(_amount *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.Faucet(&_LegacyReputationToken.TransactOpts, _amount)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(_amount uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) Faucet(_amount *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.Faucet(&_LegacyReputationToken.TransactOpts, _amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.IncreaseApproval(&_LegacyReputationToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.IncreaseApproval(&_LegacyReputationToken.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) Initialize(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "initialize", _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.Initialize(&_LegacyReputationToken.TransactOpts, _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.Initialize(&_LegacyReputationToken.TransactOpts, _universe)
}

// MigrateAllowancesFromLegacyRep is a paid mutator transaction binding the contract method 0xde4c0574.
//
// Solidity: function migrateAllowancesFromLegacyRep(_owners address[], _spenders address[]) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) MigrateAllowancesFromLegacyRep(opts *bind.TransactOpts, _owners []common.Address, _spenders []common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "migrateAllowancesFromLegacyRep", _owners, _spenders)
}

// MigrateAllowancesFromLegacyRep is a paid mutator transaction binding the contract method 0xde4c0574.
//
// Solidity: function migrateAllowancesFromLegacyRep(_owners address[], _spenders address[]) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) MigrateAllowancesFromLegacyRep(_owners []common.Address, _spenders []common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateAllowancesFromLegacyRep(&_LegacyReputationToken.TransactOpts, _owners, _spenders)
}

// MigrateAllowancesFromLegacyRep is a paid mutator transaction binding the contract method 0xde4c0574.
//
// Solidity: function migrateAllowancesFromLegacyRep(_owners address[], _spenders address[]) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) MigrateAllowancesFromLegacyRep(_owners []common.Address, _spenders []common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateAllowancesFromLegacyRep(&_LegacyReputationToken.TransactOpts, _owners, _spenders)
}

// MigrateBalancesFromLegacyRep is a paid mutator transaction binding the contract method 0x7f686259.
//
// Solidity: function migrateBalancesFromLegacyRep(_holders address[]) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) MigrateBalancesFromLegacyRep(opts *bind.TransactOpts, _holders []common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "migrateBalancesFromLegacyRep", _holders)
}

// MigrateBalancesFromLegacyRep is a paid mutator transaction binding the contract method 0x7f686259.
//
// Solidity: function migrateBalancesFromLegacyRep(_holders address[]) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) MigrateBalancesFromLegacyRep(_holders []common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateBalancesFromLegacyRep(&_LegacyReputationToken.TransactOpts, _holders)
}

// MigrateBalancesFromLegacyRep is a paid mutator transaction binding the contract method 0x7f686259.
//
// Solidity: function migrateBalancesFromLegacyRep(_holders address[]) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) MigrateBalancesFromLegacyRep(_holders []common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateBalancesFromLegacyRep(&_LegacyReputationToken.TransactOpts, _holders)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) MigrateIn(opts *bind.TransactOpts, _reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "migrateIn", _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateIn(&_LegacyReputationToken.TransactOpts, _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateIn(&_LegacyReputationToken.TransactOpts, _reporter, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) MigrateOut(opts *bind.TransactOpts, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "migrateOut", _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateOut(&_LegacyReputationToken.TransactOpts, _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateOut(&_LegacyReputationToken.TransactOpts, _destination, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x398c1a89.
//
// Solidity: function migrateOutByPayout(_payoutNumerators uint256[], _invalid bool, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) MigrateOutByPayout(opts *bind.TransactOpts, _payoutNumerators []*big.Int, _invalid bool, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "migrateOutByPayout", _payoutNumerators, _invalid, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x398c1a89.
//
// Solidity: function migrateOutByPayout(_payoutNumerators uint256[], _invalid bool, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) MigrateOutByPayout(_payoutNumerators []*big.Int, _invalid bool, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateOutByPayout(&_LegacyReputationToken.TransactOpts, _payoutNumerators, _invalid, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x398c1a89.
//
// Solidity: function migrateOutByPayout(_payoutNumerators uint256[], _invalid bool, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) MigrateOutByPayout(_payoutNumerators []*big.Int, _invalid bool, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MigrateOutByPayout(&_LegacyReputationToken.TransactOpts, _payoutNumerators, _invalid, _attotokens)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) MintForReportingParticipant(opts *bind.TransactOpts, _amountMigrated *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "mintForReportingParticipant", _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MintForReportingParticipant(&_LegacyReputationToken.TransactOpts, _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.MintForReportingParticipant(&_LegacyReputationToken.TransactOpts, _amountMigrated)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.SetController(&_LegacyReputationToken.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.SetController(&_LegacyReputationToken.TransactOpts, _controller)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.Transfer(&_LegacyReputationToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.Transfer(&_LegacyReputationToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TransferFrom(&_LegacyReputationToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TransferFrom(&_LegacyReputationToken.TransactOpts, _from, _to, _value)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) TrustedFeeWindowTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "trustedFeeWindowTransfer", _source, _destination, _attotokens)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) TrustedFeeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TrustedFeeWindowTransfer(&_LegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) TrustedFeeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TrustedFeeWindowTransfer(&_LegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) TrustedMarketTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "trustedMarketTransfer", _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TrustedMarketTransfer(&_LegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TrustedMarketTransfer(&_LegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) TrustedReportingParticipantTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "trustedReportingParticipantTransfer", _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TrustedReportingParticipantTransfer(&_LegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TrustedReportingParticipantTransfer(&_LegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) TrustedUniverseTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "trustedUniverseTransfer", _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TrustedUniverseTransfer(&_LegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.TrustedUniverseTransfer(&_LegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) UpdateParentTotalTheoreticalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "updateParentTotalTheoreticalSupply")
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) UpdateParentTotalTheoreticalSupply() (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.UpdateParentTotalTheoreticalSupply(&_LegacyReputationToken.TransactOpts)
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) UpdateParentTotalTheoreticalSupply() (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.UpdateParentTotalTheoreticalSupply(&_LegacyReputationToken.TransactOpts)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(_token address) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactor) UpdateSiblingMigrationTotal(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.contract.Transact(opts, "updateSiblingMigrationTotal", _token)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(_token address) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenSession) UpdateSiblingMigrationTotal(_token common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.UpdateSiblingMigrationTotal(&_LegacyReputationToken.TransactOpts, _token)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(_token address) returns(bool)
func (_LegacyReputationToken *LegacyReputationTokenTransactorSession) UpdateSiblingMigrationTotal(_token common.Address) (*types.Transaction, error) {
	return _LegacyReputationToken.Contract.UpdateSiblingMigrationTotal(&_LegacyReputationToken.TransactOpts, _token)
}

// LegacyReputationTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LegacyReputationToken contract.
type LegacyReputationTokenApprovalIterator struct {
	Event *LegacyReputationTokenApproval // Event containing the contract specifics and raw log

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
func (it *LegacyReputationTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyReputationTokenApproval)
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
		it.Event = new(LegacyReputationTokenApproval)
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
func (it *LegacyReputationTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyReputationTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyReputationTokenApproval represents a Approval event raised by the LegacyReputationToken contract.
type LegacyReputationTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LegacyReputationTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LegacyReputationTokenApprovalIterator{contract: _LegacyReputationToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LegacyReputationTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyReputationTokenApproval)
				if err := _LegacyReputationToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// LegacyReputationTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the LegacyReputationToken contract.
type LegacyReputationTokenBurnIterator struct {
	Event *LegacyReputationTokenBurn // Event containing the contract specifics and raw log

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
func (it *LegacyReputationTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyReputationTokenBurn)
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
		it.Event = new(LegacyReputationTokenBurn)
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
func (it *LegacyReputationTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyReputationTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyReputationTokenBurn represents a Burn event raised by the LegacyReputationToken contract.
type LegacyReputationTokenBurn struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) FilterBurn(opts *bind.FilterOpts, target []common.Address) (*LegacyReputationTokenBurnIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.FilterLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return &LegacyReputationTokenBurnIterator{contract: _LegacyReputationToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *LegacyReputationTokenBurn, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.WatchLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyReputationTokenBurn)
				if err := _LegacyReputationToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// LegacyReputationTokenFundedAccountIterator is returned from FilterFundedAccount and is used to iterate over the raw logs and unpacked data for FundedAccount events raised by the LegacyReputationToken contract.
type LegacyReputationTokenFundedAccountIterator struct {
	Event *LegacyReputationTokenFundedAccount // Event containing the contract specifics and raw log

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
func (it *LegacyReputationTokenFundedAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyReputationTokenFundedAccount)
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
		it.Event = new(LegacyReputationTokenFundedAccount)
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
func (it *LegacyReputationTokenFundedAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyReputationTokenFundedAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyReputationTokenFundedAccount represents a FundedAccount event raised by the LegacyReputationToken contract.
type LegacyReputationTokenFundedAccount struct {
	Universe   common.Address
	Sender     common.Address
	RepBalance *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFundedAccount is a free log retrieval operation binding the contract event 0xbf88b5a3158512dc8ad44dadf221f46eb1df762a2e7fd1de86aab34b0af455cf.
//
// Solidity: e FundedAccount(_universe indexed address, _sender indexed address, _repBalance uint256, _timestamp uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) FilterFundedAccount(opts *bind.FilterOpts, _universe []common.Address, _sender []common.Address) (*LegacyReputationTokenFundedAccountIterator, error) {

	var _universeRule []interface{}
	for _, _universeItem := range _universe {
		_universeRule = append(_universeRule, _universeItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.FilterLogs(opts, "FundedAccount", _universeRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return &LegacyReputationTokenFundedAccountIterator{contract: _LegacyReputationToken.contract, event: "FundedAccount", logs: logs, sub: sub}, nil
}

// WatchFundedAccount is a free log subscription operation binding the contract event 0xbf88b5a3158512dc8ad44dadf221f46eb1df762a2e7fd1de86aab34b0af455cf.
//
// Solidity: e FundedAccount(_universe indexed address, _sender indexed address, _repBalance uint256, _timestamp uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) WatchFundedAccount(opts *bind.WatchOpts, sink chan<- *LegacyReputationTokenFundedAccount, _universe []common.Address, _sender []common.Address) (event.Subscription, error) {

	var _universeRule []interface{}
	for _, _universeItem := range _universe {
		_universeRule = append(_universeRule, _universeItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.WatchLogs(opts, "FundedAccount", _universeRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyReputationTokenFundedAccount)
				if err := _LegacyReputationToken.contract.UnpackLog(event, "FundedAccount", log); err != nil {
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

// LegacyReputationTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the LegacyReputationToken contract.
type LegacyReputationTokenMintIterator struct {
	Event *LegacyReputationTokenMint // Event containing the contract specifics and raw log

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
func (it *LegacyReputationTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyReputationTokenMint)
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
		it.Event = new(LegacyReputationTokenMint)
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
func (it *LegacyReputationTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyReputationTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyReputationTokenMint represents a Mint event raised by the LegacyReputationToken contract.
type LegacyReputationTokenMint struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) FilterMint(opts *bind.FilterOpts, target []common.Address) (*LegacyReputationTokenMintIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.FilterLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return &LegacyReputationTokenMintIterator{contract: _LegacyReputationToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *LegacyReputationTokenMint, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.WatchLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyReputationTokenMint)
				if err := _LegacyReputationToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// LegacyReputationTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LegacyReputationToken contract.
type LegacyReputationTokenTransferIterator struct {
	Event *LegacyReputationTokenTransfer // Event containing the contract specifics and raw log

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
func (it *LegacyReputationTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyReputationTokenTransfer)
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
		it.Event = new(LegacyReputationTokenTransfer)
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
func (it *LegacyReputationTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyReputationTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyReputationTokenTransfer represents a Transfer event raised by the LegacyReputationToken contract.
type LegacyReputationTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LegacyReputationTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LegacyReputationTokenTransferIterator{contract: _LegacyReputationToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_LegacyReputationToken *LegacyReputationTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LegacyReputationTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyReputationToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyReputationTokenTransfer)
				if err := _LegacyReputationToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
