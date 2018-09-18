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

// OldLegacyReputationTokenABI is the input ABI used to generate the binding from.
const OldLegacyReputationTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalTheoreticalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"name\":\"_invalid\",\"type\":\"bool\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOutByPayout\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETERNAL_APPROVAL_VALUE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOut\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLegacyRepToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIsMigratingFromLegacy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_holders\",\"type\":\"address[]\"}],\"name\":\"migrateBalancesFromLegacyRep\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUniverse\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedFeeWindowTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reporter\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateIn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateParentTotalTheoreticalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedReportingParticipantTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"updateSiblingMigrationTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountMigrated\",\"type\":\"uint256\"}],\"name\":\"mintForReportingParticipant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_spenders\",\"type\":\"address[]\"}],\"name\":\"migrateAllowancesFromLegacyRep\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedMarketTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedUniverseTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// OldLegacyReputationTokenBin is the compiled bytecode used for deploying new contracts.
const OldLegacyReputationTokenBin = `0x60806040526002805460ff1916905560008054600160a060020a031916331790556127c18061002f6000396000f3006080604052600436106101a85763ffffffff60e060020a60003504166306fdde0381146101ad578063095ea7b31461023757806318160ddd1461026f578063238d35901461029657806323b872dd146102ab5780633018205f146102d5578063313ce56714610306578063398c1a8914610331578063634eaff11461038d57806366188463146103a25780636e7ce591146103c657806370a08231146103ea578063774692751461040b5780637cf99c33146104205780637f68625914610435578063870c426d1461048a57806390fbf84e1461049f57806391d76bbb146104c957806392eefe9b146104de57806395d89b41146104ff578063a0c1ca3414610514578063a819515d14610538578063a9059cbb1461054d578063b873e9a714610571578063bef72fa21461059b578063c4d66de8146105b0578063d73dd623146105d1578063d9d3e07d146105f5578063db05413414610616578063db0a087c1461062e578063dd62ed3e14610643578063de4c05741461066a578063dea6aec7146106f8578063ee89dab41461070d578063f22b258a14610722578063fe98184d1461074c575b600080fd5b3480156101b957600080fd5b506101c2610776565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101fc5781810151838201526020016101e4565b50505050905090810190601f1680156102295780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024357600080fd5b5061025b600160a060020a03600435166024356107ad565b604080519115158252519081900360200190f35b34801561027b57600080fd5b506102846107c4565b60408051918252519081900360200190f35b3480156102a257600080fd5b506102846107ca565b3480156102b757600080fd5b5061025b600160a060020a03600435811690602435166044356107d0565b3480156102e157600080fd5b506102ea6107f6565b60408051600160a060020a039092168252519081900360200190f35b34801561031257600080fd5b5061031b610805565b6040805160ff9092168252519081900360200190f35b34801561033d57600080fd5b506040805160206004803580820135838102808601850190965280855261025b9536959394602494938501929182918501908490808284375094975050505082351515935050506020013561080a565b34801561039957600080fd5b50610284610a4c565b3480156103ae57600080fd5b5061025b600160a060020a0360043516602435610a52565b3480156103d257600080fd5b5061025b600160a060020a0360043516602435610ab8565b3480156103f657600080fd5b50610284600160a060020a0360043516610b90565b34801561041757600080fd5b506102ea610bab565b34801561042c57600080fd5b5061025b610c62565b34801561044157600080fd5b506040805160206004803580820135838102808601850190965280855261025b95369593946024949385019291829185019084908082843750949750610c6b9650505050505050565b34801561049657600080fd5b506102ea610cd6565b3480156104ab57600080fd5b5061025b600160a060020a0360043581169060243516604435610ce5565b3480156104d557600080fd5b50610284610db1565b3480156104ea57600080fd5b5061025b600160a060020a0360043516610db7565b34801561050b57600080fd5b506101c2610e01565b34801561052057600080fd5b5061025b600160a060020a0360043516602435610e38565b34801561054457600080fd5b5061025b6112ce565b34801561055957600080fd5b5061025b600160a060020a036004351660243561149b565b34801561057d57600080fd5b5061025b600160a060020a03600435811690602435166044356114bf565b3480156105a757600080fd5b50610284611549565b3480156105bc57600080fd5b5061025b600160a060020a036004351661154f565b3480156105dd57600080fd5b5061025b600160a060020a03600435166024356116dd565b34801561060157600080fd5b5061025b600160a060020a0360043516611719565b34801561062257600080fd5b5061025b600435611aa2565b34801561063a57600080fd5b50610284611c1e565b34801561064f57600080fd5b50610284600160a060020a0360043581169060243516611c42565b34801561067657600080fd5b506040805160206004803580820135838102808601850190965280855261025b95369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750611c6d9650505050505050565b34801561070457600080fd5b50610284611d96565b34801561071957600080fd5b5061025b611d9c565b34801561072e57600080fd5b5061025b600160a060020a0360043581169060243516604435611da5565b34801561075857600080fd5b5061025b600160a060020a0360043581169060243516604435611e2f565b60408051808201909152600a81527f52657075746174696f6e00000000000000000000000000000000000000000000602082015281565b60006107ba338484611e6a565b5060019392505050565b60035490565b600a5490565b600b5460009060ff16156107e357600080fd5b6107ee848484611ed5565b949350505050565b600054600160a060020a031690565b601281565b600b546000908190819060ff161561082157600080fd5b60025460ff16151561083257600080fd5b6000841161083f57600080fd5b600654604080517fdf428e3b000000000000000000000000000000000000000000000000000000008152871515602482015260048101918252885160448201528851600160a060020a039093169263df428e3b928a928a928291606401906020808701910280838360005b838110156108c25781810151838201526020016108aa565b505050509050019350505050602060405180830381600087803b1580156108e857600080fd5b505af11580156108fc573d6000803e3d6000fd5b505050506040513d602081101561091257600080fd5b5051604080517fb80907f20000000000000000000000000000000000000000000000000000000081529051919350600160a060020a0384169163b80907f2916004808201926020929091908290030181600087803b15801561097357600080fd5b505af1158015610987573d6000803e3d6000fd5b505050506040513d602081101561099d57600080fd5b505190506109ab3385611f4e565b50604080517fa0c1ca34000000000000000000000000000000000000000000000000000000008152336004820152602481018690529051600160a060020a0383169163a0c1ca349160448083019260209291908290030181600087803b158015610a1457600080fd5b505af1158015610a28573d6000803e3d6000fd5b505050506040513d6020811015610a3e57600080fd5b506001979650505050505050565b60001981565b336000908152600560209081526040808320600160a060020a038616845290915281205480831115610a9057610a8a33856000611e6a565b50610aac565b610aaa3385610aa5848763ffffffff611fef16565b611e6a565b505b600191505b5092915050565b600b5460009060ff1615610acb57600080fd5b60025460ff161515610adc57600080fd5b60008211610ae957600080fd5b610af283612004565b50610afd3383611f4e565b50604080517fa0c1ca34000000000000000000000000000000000000000000000000000000008152336004820152602481018490529051600160a060020a0385169163a0c1ca349160448083019260209291908290030181600087803b158015610b6657600080fd5b505af1158015610b7a573d6000803e3d6000fd5b505050506040513d6020811015610aaa57600080fd5b600160a060020a031660009081526004602052604090205490565b60008054604080517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4c656761637952657075746174696f6e546f6b656e000000000000000000000060048201529051600160a060020a039092169163f39ec1f79160248082019260209290919082900301818787803b158015610c3157600080fd5b505af1158015610c45573d6000803e3d6000fd5b505050506040513d6020811015610c5b57600080fd5b5051905090565b600b5460ff1690565b600b546000908190819060ff161515610c8357600080fd5b60025460ff161515610c9457600080fd5b610c9c610bab565b9150600090505b83518110156107ba57610ccd8482815181101515610cbd57fe5b906020019060200201518361219e565b50600101610ca3565b600654600160a060020a031690565b600b5460009060ff1615610cf857600080fd5b60025460ff161515610d0957600080fd5b600654604080517fc7c88d700000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163c7c88d70916024808201926020929091908290030181600087803b158015610d6f57600080fd5b505af1158015610d83573d6000803e3d6000fd5b505050506040513d6020811015610d9957600080fd5b50511515610da657600080fd5b6107ee8484846122b1565b60075490565b60008054600160a060020a03163314610dcf57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60408051808201909152600381527f5245500000000000000000000000000000000000000000000000000000000000602082015281565b600b546000908190819060ff1615610e4f57600080fd5b60025460ff161515610e6057600080fd5b600660009054906101000a9004600160a060020a0316600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610eb357600080fd5b505af1158015610ec7573d6000803e3d6000fd5b505050506040513d6020811015610edd57600080fd5b5051604080517fb80907f20000000000000000000000000000000000000000000000000000000081529051919350600160a060020a0384169163b80907f2916004808201926020929091908290030181600087803b158015610f3e57600080fd5b505af1158015610f52573d6000803e3d6000fd5b505050506040513d6020811015610f6857600080fd5b5051600160a060020a03163314610f7e57600080fd5b610f888585612371565b506007805485019055604080517f77e71ee50000000000000000000000000000000000000000000000000000000081529051600160a060020a038416916377e71ee59160048083019260209291908290030181600087803b158015610fec57600080fd5b505af1158015611000573d6000803e3d6000fd5b505050506040513d602081101561101657600080fd5b505160008054604080517f188ec3560000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169263188ec356926004808401936020939083900390910190829087803b15801561107957600080fd5b505af115801561108d573d6000803e3d6000fd5b505050506040513d60208110156110a357600080fd5b505110156110d2576110bc84601463ffffffff61241216565b90506110c88582612371565b50600a8054820190555b81600160a060020a031663cb1d84186040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561111057600080fd5b505af1158015611124573d6000803e3d6000fd5b505050506040513d602081101561113a57600080fd5b5051604080517f8d4e40830000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921691638d4e4083916004808201926020929091908290030181600087803b15801561119957600080fd5b505af11580156111ad573d6000803e3d6000fd5b505050506040513d60208110156111c357600080fd5b505115156112c35781600160a060020a031663f7095d9d600660009054906101000a9004600160a060020a0316600160a060020a031663c38c0fa76040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561122d57600080fd5b505af1158015611241573d6000803e3d6000fd5b505050506040513d602081101561125757600080fd5b50516040805160e060020a63ffffffff851602815260048101929092525160248083019260209291908290030181600087803b15801561129657600080fd5b505af11580156112aa573d6000803e3d6000fd5b505050506040513d60208110156112c057600080fd5b50505b506001949350505050565b600b54600090819060ff16156112e357600080fd5b600660009054906101000a9004600160a060020a0316600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561133657600080fd5b505af115801561134a573d6000803e3d6000fd5b505050506040513d602081101561136057600080fd5b5051600954600a80549190910390559050600160a060020a038116151561139157611389612429565b600955611488565b80600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156113cf57600080fd5b505af11580156113e3573d6000803e3d6000fd5b505050506040513d60208110156113f957600080fd5b5051604080517f238d35900000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169163238d3590916004808201926020929091908290030181600087803b15801561145857600080fd5b505af115801561146c573d6000803e3d6000fd5b505050506040513d602081101561148257600080fd5b50516009555b5050600954600a80549091019055600190565b600b5460009060ff16156114ae57600080fd5b6114b88383612438565b9392505050565b600b5460009060ff16156114d257600080fd5b60025460ff1615156114e357600080fd5b600654604080517ff76514c70000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163f76514c7916024808201926020929091908290030181600087803b158015610d6f57600080fd5b60015481565b600254600090819060ff161561156457600080fd5b61156c612445565b50600160a060020a038316151561158257600080fd5b6006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385161790556115b26112ce565b506115bb610bab565b90506000600160a060020a031683600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561160657600080fd5b505af115801561161a573d6000803e3d6000fd5b505050506040513d602081101561163057600080fd5b5051600b805460ff1916600160a060020a039283169390931492909217909155604080517f18160ddd0000000000000000000000000000000000000000000000000000000081529051918316916318160ddd916004808201926020929091908290030181600087803b1580156116a557600080fd5b505af11580156116b9573d6000803e3d6000fd5b505050506040513d60208110156116cf57600080fd5b5051600c5550600192915050565b336000818152600560209081526040808320600160a060020a038716845290915281205490916107ba918590610aa5908663ffffffff61246c16565b600b54600090819060ff161561172e57600080fd5b600160a060020a03831630141561174457600080fd5b82600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561178257600080fd5b505af1158015611796573d6000803e3d6000fd5b505050506040513d60208110156117ac57600080fd5b5051600654604080517fa63f13500000000000000000000000000000000000000000000000000000000081529051929350600160a060020a039091169163a63f1350916004808201926020929091908290030181600087803b15801561181157600080fd5b505af1158015611825573d6000803e3d6000fd5b505050506040513d602081101561183b57600080fd5b5051604080517fc38c0fa70000000000000000000000000000000000000000000000000000000081529051600160a060020a039283169263eceba876929085169163c38c0fa7916004808201926020929091908290030181600087803b1580156118a457600080fd5b505af11580156118b8573d6000803e3d6000fd5b505050506040513d60208110156118ce57600080fd5b50516040805160e060020a63ffffffff851602815260048101929092525160248083019260209291908290030181600087803b15801561190d57600080fd5b505af1158015611921573d6000803e3d6000fd5b505050506040513d602081101561193757600080fd5b5051604080517fb80907f20000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169163b80907f2916004808201926020929091908290030181600087803b15801561199657600080fd5b505af11580156119aa573d6000803e3d6000fd5b505050506040513d60208110156119c057600080fd5b5051600160a060020a038481169116146119d957600080fd5b600160a060020a038316600081815260086020908152604080832054600a8054909101905580517f91d76bbb00000000000000000000000000000000000000000000000000000000815290516391d76bbb93600480840194938390030190829087803b158015611a4857600080fd5b505af1158015611a5c573d6000803e3d6000fd5b505050506040513d6020811015611a7257600080fd5b5051600160a060020a039390931660009081526008602052604090208390555050600a8054919091039055600190565b600b5460009081908190819060ff1615611abb57600080fd5b60025460ff161515611acc57600080fd5b600660009054906101000a9004600160a060020a0316600160a060020a031663a63f13506040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611b1f57600080fd5b505af1158015611b33573d6000803e3d6000fd5b505050506040513d6020811015611b4957600080fd5b5051604080517ff76514c700000000000000000000000000000000000000000000000000000000815233600482018190529151929550909350600160a060020a0385169163f76514c7916024808201926020929091908290030181600087803b158015611bb557600080fd5b505af1158015611bc9573d6000803e3d6000fd5b505050506040513d6020811015611bdf57600080fd5b50511515611bec57600080fd5b611bfd85600263ffffffff61241216565b9050611c098282612371565b50600a80548201905560019350505050919050565b7f52657075746174696f6e546f6b656e000000000000000000000000000000000090565b600160a060020a03918216600090815260056020908152604080832093909416825291909152205490565b600b546000908190819081908190819060ff161515611c8b57600080fd5b60025460ff161515611c9c57600080fd5b611ca4610bab565b9450600093505b8751841015610a3e578784815181101515611cc257fe5b9060200190602002015192508684815181101515611cdc57fe5b6020908102909101810151604080517fdd62ed3e000000000000000000000000000000000000000000000000000000008152600160a060020a038781166004830152808416602483015291519295509088169263dd62ed3e926044808401938290030181600087803b158015611d5157600080fd5b505af1158015611d65573d6000803e3d6000fd5b505050506040513d6020811015611d7b57600080fd5b50519050611d8a838383611e6a565b50600190930192611cab565b600c5490565b60025460ff1690565b600b5460009060ff1615611db857600080fd5b60025460ff161515611dc957600080fd5b600654604080517f9f7e1bf60000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a0390921691639f7e1bf6916024808201926020929091908290030181600087803b158015610d6f57600080fd5b600b5460009060ff1615611e4257600080fd5b60025460ff161515611e5357600080fd5b600654600160a060020a03163314610da657600080fd5b600160a060020a03808416600081815260056020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a03831660009081526005602090815260408083203384529091528120546000198114611f3757611f12818463ffffffff611fef16565b600160a060020a03861660009081526005602090815260408083203384529091529020555b611f428585856122b1565b50600195945050505050565b600160a060020a038216600090815260046020526040812054611f77908363ffffffff611fef16565b600160a060020a038416600090815260046020526040902055600354611fa3908363ffffffff611fef16565b600355604080518381529051600160a060020a038516917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a26107ba838361247e565b600082821115611ffe57600080fd5b50900390565b600080600083600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561204757600080fd5b505af115801561205b573d6000803e3d6000fd5b505050506040513d602081101561207157600080fd5b5051600654604080517f9517317c000000000000000000000000000000000000000000000000000000008152600160a060020a0380851660048301529151939550911691639517317c916024808201926020929091908290030181600087803b1580156120dd57600080fd5b505af11580156120f1573d6000803e3d6000fd5b505050506040513d602081101561210757600080fd5b5051151561211457600080fd5b81905083600160a060020a031681600160a060020a031663b80907f26040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561215f57600080fd5b505af1158015612173573d6000803e3d6000fd5b505050506040513d602081101561218957600080fd5b5051600160a060020a0316146107ba57600080fd5b600b54600090819060ff1615156121b457600080fd5b60025460ff1615156121c557600080fd5b600160a060020a03841660009081526004602052604081205411156121ed5760009150610ab1565b82600160a060020a03166370a08231856040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561224857600080fd5b505af115801561225c573d6000803e3d6000fd5b505050506040513d602081101561227257600080fd5b505190508015156122865760009150610ab1565b6122908482612371565b50600354600c541415610aac575050600b805460ff19169055506001919050565b600160a060020a0383166000908152600460205260408120546122da908363ffffffff611fef16565b600160a060020a03808616600090815260046020526040808220939093559085168152205461230f908363ffffffff61246c16565b600160a060020a0380851660008181526004602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a36112c3848484612575565b600160a060020a03821660009081526004602052604081205461239a908363ffffffff61246c16565b600160a060020a0384166000908152600460205260409020556003546123c6908363ffffffff61246c16565b600355604080518381529051600160a060020a038516917f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885919081900360200190a26107ba838361269e565b600080828481151561242057fe5b04949350505050565b6a09195731e2ce35eb00000090565b60006114b83384846122b1565b60025460009060ff161561245857600080fd5b506002805460ff1916600190811790915590565b6000828201838110156114b857600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156124d257600080fd5b505af11580156124e6573d6000803e3d6000fd5b505050506040513d60208110156124fc57600080fd5b5051600654604080517f4405a339000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015286831660248201526044810186905290519190921691634405a3399160648083019260209291908290030181600087803b158015610b6657600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156125c957600080fd5b505af11580156125dd573d6000803e3d6000fd5b505050506040513d60208110156125f357600080fd5b5051600654604080517fec37a6e4000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015287831660248201528683166044820152606481018690529051919092169163ec37a6e49160848083019260209291908290030181600087803b15801561267457600080fd5b505af1158015612688573d6000803e3d6000fd5b505050506040513d6020811015611f4257600080fd5b60008060009054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156126f257600080fd5b505af1158015612706573d6000803e3d6000fd5b505050506040513d602081101561271c57600080fd5b5051600654604080517f79fff7a9000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152868316602482015260448101869052905191909216916379fff7a99160648083019260209291908290030181600087803b158015610b6657600080fd00a165627a7a7230582073454304f923443e62bd77fa2d7cf38c07f89b281c1698b3fbf4785145069bd20029`

// DeployOldLegacyReputationToken deploys a new Ethereum contract, binding an instance of OldLegacyReputationToken to it.
func DeployOldLegacyReputationToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OldLegacyReputationToken, error) {
	parsed, err := abi.JSON(strings.NewReader(OldLegacyReputationTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OldLegacyReputationTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OldLegacyReputationToken{OldLegacyReputationTokenCaller: OldLegacyReputationTokenCaller{contract: contract}, OldLegacyReputationTokenTransactor: OldLegacyReputationTokenTransactor{contract: contract}, OldLegacyReputationTokenFilterer: OldLegacyReputationTokenFilterer{contract: contract}}, nil
}

// OldLegacyReputationToken is an auto generated Go binding around an Ethereum contract.
type OldLegacyReputationToken struct {
	OldLegacyReputationTokenCaller     // Read-only binding to the contract
	OldLegacyReputationTokenTransactor // Write-only binding to the contract
	OldLegacyReputationTokenFilterer   // Log filterer for contract events
}

// OldLegacyReputationTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type OldLegacyReputationTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldLegacyReputationTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OldLegacyReputationTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldLegacyReputationTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OldLegacyReputationTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldLegacyReputationTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OldLegacyReputationTokenSession struct {
	Contract     *OldLegacyReputationToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OldLegacyReputationTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OldLegacyReputationTokenCallerSession struct {
	Contract *OldLegacyReputationTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// OldLegacyReputationTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OldLegacyReputationTokenTransactorSession struct {
	Contract     *OldLegacyReputationTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// OldLegacyReputationTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type OldLegacyReputationTokenRaw struct {
	Contract *OldLegacyReputationToken // Generic contract binding to access the raw methods on
}

// OldLegacyReputationTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OldLegacyReputationTokenCallerRaw struct {
	Contract *OldLegacyReputationTokenCaller // Generic read-only contract binding to access the raw methods on
}

// OldLegacyReputationTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OldLegacyReputationTokenTransactorRaw struct {
	Contract *OldLegacyReputationTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOldLegacyReputationToken creates a new instance of OldLegacyReputationToken, bound to a specific deployed contract.
func NewOldLegacyReputationToken(address common.Address, backend bind.ContractBackend) (*OldLegacyReputationToken, error) {
	contract, err := bindOldLegacyReputationToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OldLegacyReputationToken{OldLegacyReputationTokenCaller: OldLegacyReputationTokenCaller{contract: contract}, OldLegacyReputationTokenTransactor: OldLegacyReputationTokenTransactor{contract: contract}, OldLegacyReputationTokenFilterer: OldLegacyReputationTokenFilterer{contract: contract}}, nil
}

// NewOldLegacyReputationTokenCaller creates a new read-only instance of OldLegacyReputationToken, bound to a specific deployed contract.
func NewOldLegacyReputationTokenCaller(address common.Address, caller bind.ContractCaller) (*OldLegacyReputationTokenCaller, error) {
	contract, err := bindOldLegacyReputationToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OldLegacyReputationTokenCaller{contract: contract}, nil
}

// NewOldLegacyReputationTokenTransactor creates a new write-only instance of OldLegacyReputationToken, bound to a specific deployed contract.
func NewOldLegacyReputationTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*OldLegacyReputationTokenTransactor, error) {
	contract, err := bindOldLegacyReputationToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OldLegacyReputationTokenTransactor{contract: contract}, nil
}

// NewOldLegacyReputationTokenFilterer creates a new log filterer instance of OldLegacyReputationToken, bound to a specific deployed contract.
func NewOldLegacyReputationTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*OldLegacyReputationTokenFilterer, error) {
	contract, err := bindOldLegacyReputationToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OldLegacyReputationTokenFilterer{contract: contract}, nil
}

// bindOldLegacyReputationToken binds a generic wrapper to an already deployed contract.
func bindOldLegacyReputationToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OldLegacyReputationTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldLegacyReputationToken *OldLegacyReputationTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OldLegacyReputationToken.Contract.OldLegacyReputationTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldLegacyReputationToken *OldLegacyReputationTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.OldLegacyReputationTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldLegacyReputationToken *OldLegacyReputationTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.OldLegacyReputationTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OldLegacyReputationToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.contract.Transact(opts, method, params...)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) ETERNALAPPROVALVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "ETERNAL_APPROVAL_VALUE")
	return *ret0, err
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.ETERNALAPPROVALVALUE(&_OldLegacyReputationToken.CallOpts)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.ETERNALAPPROVALVALUE(&_OldLegacyReputationToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.Allowance(&_OldLegacyReputationToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.Allowance(&_OldLegacyReputationToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.BalanceOf(&_OldLegacyReputationToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.BalanceOf(&_OldLegacyReputationToken.CallOpts, _owner)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) ControllerLookupName() ([32]byte, error) {
	return _OldLegacyReputationToken.Contract.ControllerLookupName(&_OldLegacyReputationToken.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) ControllerLookupName() ([32]byte, error) {
	return _OldLegacyReputationToken.Contract.ControllerLookupName(&_OldLegacyReputationToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) Decimals() (uint8, error) {
	return _OldLegacyReputationToken.Contract.Decimals(&_OldLegacyReputationToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) Decimals() (uint8, error) {
	return _OldLegacyReputationToken.Contract.Decimals(&_OldLegacyReputationToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) GetController() (common.Address, error) {
	return _OldLegacyReputationToken.Contract.GetController(&_OldLegacyReputationToken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) GetController() (common.Address, error) {
	return _OldLegacyReputationToken.Contract.GetController(&_OldLegacyReputationToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) GetInitialized() (bool, error) {
	return _OldLegacyReputationToken.Contract.GetInitialized(&_OldLegacyReputationToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() constant returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) GetInitialized() (bool, error) {
	return _OldLegacyReputationToken.Contract.GetInitialized(&_OldLegacyReputationToken.CallOpts)
}

// GetIsMigratingFromLegacy is a free data retrieval call binding the contract method 0x7cf99c33.
//
// Solidity: function getIsMigratingFromLegacy() constant returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) GetIsMigratingFromLegacy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "getIsMigratingFromLegacy")
	return *ret0, err
}

// GetIsMigratingFromLegacy is a free data retrieval call binding the contract method 0x7cf99c33.
//
// Solidity: function getIsMigratingFromLegacy() constant returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) GetIsMigratingFromLegacy() (bool, error) {
	return _OldLegacyReputationToken.Contract.GetIsMigratingFromLegacy(&_OldLegacyReputationToken.CallOpts)
}

// GetIsMigratingFromLegacy is a free data retrieval call binding the contract method 0x7cf99c33.
//
// Solidity: function getIsMigratingFromLegacy() constant returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) GetIsMigratingFromLegacy() (bool, error) {
	return _OldLegacyReputationToken.Contract.GetIsMigratingFromLegacy(&_OldLegacyReputationToken.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() constant returns(address)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) GetLegacyRepToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "getLegacyRepToken")
	return *ret0, err
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() constant returns(address)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) GetLegacyRepToken() (common.Address, error) {
	return _OldLegacyReputationToken.Contract.GetLegacyRepToken(&_OldLegacyReputationToken.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() constant returns(address)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) GetLegacyRepToken() (common.Address, error) {
	return _OldLegacyReputationToken.Contract.GetLegacyRepToken(&_OldLegacyReputationToken.CallOpts)
}

// GetTargetSupply is a free data retrieval call binding the contract method 0xdea6aec7.
//
// Solidity: function getTargetSupply() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) GetTargetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "getTargetSupply")
	return *ret0, err
}

// GetTargetSupply is a free data retrieval call binding the contract method 0xdea6aec7.
//
// Solidity: function getTargetSupply() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) GetTargetSupply() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.GetTargetSupply(&_OldLegacyReputationToken.CallOpts)
}

// GetTargetSupply is a free data retrieval call binding the contract method 0xdea6aec7.
//
// Solidity: function getTargetSupply() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) GetTargetSupply() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.GetTargetSupply(&_OldLegacyReputationToken.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) GetTotalMigrated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "getTotalMigrated")
	return *ret0, err
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) GetTotalMigrated() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.GetTotalMigrated(&_OldLegacyReputationToken.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) GetTotalMigrated() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.GetTotalMigrated(&_OldLegacyReputationToken.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) GetTotalTheoreticalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "getTotalTheoreticalSupply")
	return *ret0, err
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.GetTotalTheoreticalSupply(&_OldLegacyReputationToken.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.GetTotalTheoreticalSupply(&_OldLegacyReputationToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) GetTypeName() ([32]byte, error) {
	return _OldLegacyReputationToken.Contract.GetTypeName(&_OldLegacyReputationToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() constant returns(bytes32)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) GetTypeName() ([32]byte, error) {
	return _OldLegacyReputationToken.Contract.GetTypeName(&_OldLegacyReputationToken.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) GetUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "getUniverse")
	return *ret0, err
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) GetUniverse() (common.Address, error) {
	return _OldLegacyReputationToken.Contract.GetUniverse(&_OldLegacyReputationToken.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() constant returns(address)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) GetUniverse() (common.Address, error) {
	return _OldLegacyReputationToken.Contract.GetUniverse(&_OldLegacyReputationToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) Name() (string, error) {
	return _OldLegacyReputationToken.Contract.Name(&_OldLegacyReputationToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) Name() (string, error) {
	return _OldLegacyReputationToken.Contract.Name(&_OldLegacyReputationToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) Symbol() (string, error) {
	return _OldLegacyReputationToken.Contract.Symbol(&_OldLegacyReputationToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) Symbol() (string, error) {
	return _OldLegacyReputationToken.Contract.Symbol(&_OldLegacyReputationToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldLegacyReputationToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) TotalSupply() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.TotalSupply(&_OldLegacyReputationToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _OldLegacyReputationToken.Contract.TotalSupply(&_OldLegacyReputationToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.Approve(&_OldLegacyReputationToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.Approve(&_OldLegacyReputationToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.DecreaseApproval(&_OldLegacyReputationToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.DecreaseApproval(&_OldLegacyReputationToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.IncreaseApproval(&_OldLegacyReputationToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.IncreaseApproval(&_OldLegacyReputationToken.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) Initialize(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "initialize", _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.Initialize(&_OldLegacyReputationToken.TransactOpts, _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(_universe address) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.Initialize(&_OldLegacyReputationToken.TransactOpts, _universe)
}

// MigrateAllowancesFromLegacyRep is a paid mutator transaction binding the contract method 0xde4c0574.
//
// Solidity: function migrateAllowancesFromLegacyRep(_owners address[], _spenders address[]) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) MigrateAllowancesFromLegacyRep(opts *bind.TransactOpts, _owners []common.Address, _spenders []common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "migrateAllowancesFromLegacyRep", _owners, _spenders)
}

// MigrateAllowancesFromLegacyRep is a paid mutator transaction binding the contract method 0xde4c0574.
//
// Solidity: function migrateAllowancesFromLegacyRep(_owners address[], _spenders address[]) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) MigrateAllowancesFromLegacyRep(_owners []common.Address, _spenders []common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateAllowancesFromLegacyRep(&_OldLegacyReputationToken.TransactOpts, _owners, _spenders)
}

// MigrateAllowancesFromLegacyRep is a paid mutator transaction binding the contract method 0xde4c0574.
//
// Solidity: function migrateAllowancesFromLegacyRep(_owners address[], _spenders address[]) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) MigrateAllowancesFromLegacyRep(_owners []common.Address, _spenders []common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateAllowancesFromLegacyRep(&_OldLegacyReputationToken.TransactOpts, _owners, _spenders)
}

// MigrateBalancesFromLegacyRep is a paid mutator transaction binding the contract method 0x7f686259.
//
// Solidity: function migrateBalancesFromLegacyRep(_holders address[]) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) MigrateBalancesFromLegacyRep(opts *bind.TransactOpts, _holders []common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "migrateBalancesFromLegacyRep", _holders)
}

// MigrateBalancesFromLegacyRep is a paid mutator transaction binding the contract method 0x7f686259.
//
// Solidity: function migrateBalancesFromLegacyRep(_holders address[]) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) MigrateBalancesFromLegacyRep(_holders []common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateBalancesFromLegacyRep(&_OldLegacyReputationToken.TransactOpts, _holders)
}

// MigrateBalancesFromLegacyRep is a paid mutator transaction binding the contract method 0x7f686259.
//
// Solidity: function migrateBalancesFromLegacyRep(_holders address[]) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) MigrateBalancesFromLegacyRep(_holders []common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateBalancesFromLegacyRep(&_OldLegacyReputationToken.TransactOpts, _holders)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) MigrateIn(opts *bind.TransactOpts, _reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "migrateIn", _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateIn(&_OldLegacyReputationToken.TransactOpts, _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(_reporter address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateIn(&_OldLegacyReputationToken.TransactOpts, _reporter, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) MigrateOut(opts *bind.TransactOpts, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "migrateOut", _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateOut(&_OldLegacyReputationToken.TransactOpts, _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(_destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateOut(&_OldLegacyReputationToken.TransactOpts, _destination, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x398c1a89.
//
// Solidity: function migrateOutByPayout(_payoutNumerators uint256[], _invalid bool, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) MigrateOutByPayout(opts *bind.TransactOpts, _payoutNumerators []*big.Int, _invalid bool, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "migrateOutByPayout", _payoutNumerators, _invalid, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x398c1a89.
//
// Solidity: function migrateOutByPayout(_payoutNumerators uint256[], _invalid bool, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) MigrateOutByPayout(_payoutNumerators []*big.Int, _invalid bool, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateOutByPayout(&_OldLegacyReputationToken.TransactOpts, _payoutNumerators, _invalid, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x398c1a89.
//
// Solidity: function migrateOutByPayout(_payoutNumerators uint256[], _invalid bool, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) MigrateOutByPayout(_payoutNumerators []*big.Int, _invalid bool, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MigrateOutByPayout(&_OldLegacyReputationToken.TransactOpts, _payoutNumerators, _invalid, _attotokens)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) MintForReportingParticipant(opts *bind.TransactOpts, _amountMigrated *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "mintForReportingParticipant", _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MintForReportingParticipant(&_OldLegacyReputationToken.TransactOpts, _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(_amountMigrated uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.MintForReportingParticipant(&_OldLegacyReputationToken.TransactOpts, _amountMigrated)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.SetController(&_OldLegacyReputationToken.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.SetController(&_OldLegacyReputationToken.TransactOpts, _controller)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.Transfer(&_OldLegacyReputationToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.Transfer(&_OldLegacyReputationToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TransferFrom(&_OldLegacyReputationToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TransferFrom(&_OldLegacyReputationToken.TransactOpts, _from, _to, _value)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) TrustedFeeWindowTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "trustedFeeWindowTransfer", _source, _destination, _attotokens)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) TrustedFeeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TrustedFeeWindowTransfer(&_OldLegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedFeeWindowTransfer is a paid mutator transaction binding the contract method 0x90fbf84e.
//
// Solidity: function trustedFeeWindowTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) TrustedFeeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TrustedFeeWindowTransfer(&_OldLegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) TrustedMarketTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "trustedMarketTransfer", _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TrustedMarketTransfer(&_OldLegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TrustedMarketTransfer(&_OldLegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) TrustedReportingParticipantTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "trustedReportingParticipantTransfer", _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TrustedReportingParticipantTransfer(&_OldLegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TrustedReportingParticipantTransfer(&_OldLegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) TrustedUniverseTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "trustedUniverseTransfer", _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TrustedUniverseTransfer(&_OldLegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(_source address, _destination address, _attotokens uint256) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.TrustedUniverseTransfer(&_OldLegacyReputationToken.TransactOpts, _source, _destination, _attotokens)
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) UpdateParentTotalTheoreticalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "updateParentTotalTheoreticalSupply")
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) UpdateParentTotalTheoreticalSupply() (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.UpdateParentTotalTheoreticalSupply(&_OldLegacyReputationToken.TransactOpts)
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) UpdateParentTotalTheoreticalSupply() (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.UpdateParentTotalTheoreticalSupply(&_OldLegacyReputationToken.TransactOpts)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(_token address) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactor) UpdateSiblingMigrationTotal(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.contract.Transact(opts, "updateSiblingMigrationTotal", _token)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(_token address) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenSession) UpdateSiblingMigrationTotal(_token common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.UpdateSiblingMigrationTotal(&_OldLegacyReputationToken.TransactOpts, _token)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(_token address) returns(bool)
func (_OldLegacyReputationToken *OldLegacyReputationTokenTransactorSession) UpdateSiblingMigrationTotal(_token common.Address) (*types.Transaction, error) {
	return _OldLegacyReputationToken.Contract.UpdateSiblingMigrationTotal(&_OldLegacyReputationToken.TransactOpts, _token)
}

// OldLegacyReputationTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OldLegacyReputationToken contract.
type OldLegacyReputationTokenApprovalIterator struct {
	Event *OldLegacyReputationTokenApproval // Event containing the contract specifics and raw log

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
func (it *OldLegacyReputationTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldLegacyReputationTokenApproval)
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
		it.Event = new(OldLegacyReputationTokenApproval)
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
func (it *OldLegacyReputationTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldLegacyReputationTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldLegacyReputationTokenApproval represents a Approval event raised by the OldLegacyReputationToken contract.
type OldLegacyReputationTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OldLegacyReputationTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OldLegacyReputationToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OldLegacyReputationTokenApprovalIterator{contract: _OldLegacyReputationToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OldLegacyReputationTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OldLegacyReputationToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldLegacyReputationTokenApproval)
				if err := _OldLegacyReputationToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// OldLegacyReputationTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the OldLegacyReputationToken contract.
type OldLegacyReputationTokenBurnIterator struct {
	Event *OldLegacyReputationTokenBurn // Event containing the contract specifics and raw log

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
func (it *OldLegacyReputationTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldLegacyReputationTokenBurn)
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
		it.Event = new(OldLegacyReputationTokenBurn)
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
func (it *OldLegacyReputationTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldLegacyReputationTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldLegacyReputationTokenBurn represents a Burn event raised by the OldLegacyReputationToken contract.
type OldLegacyReputationTokenBurn struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenFilterer) FilterBurn(opts *bind.FilterOpts, target []common.Address) (*OldLegacyReputationTokenBurnIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _OldLegacyReputationToken.contract.FilterLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return &OldLegacyReputationTokenBurnIterator{contract: _OldLegacyReputationToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(target indexed address, value uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *OldLegacyReputationTokenBurn, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _OldLegacyReputationToken.contract.WatchLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldLegacyReputationTokenBurn)
				if err := _OldLegacyReputationToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// OldLegacyReputationTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the OldLegacyReputationToken contract.
type OldLegacyReputationTokenMintIterator struct {
	Event *OldLegacyReputationTokenMint // Event containing the contract specifics and raw log

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
func (it *OldLegacyReputationTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldLegacyReputationTokenMint)
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
		it.Event = new(OldLegacyReputationTokenMint)
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
func (it *OldLegacyReputationTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldLegacyReputationTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldLegacyReputationTokenMint represents a Mint event raised by the OldLegacyReputationToken contract.
type OldLegacyReputationTokenMint struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenFilterer) FilterMint(opts *bind.FilterOpts, target []common.Address) (*OldLegacyReputationTokenMintIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _OldLegacyReputationToken.contract.FilterLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return &OldLegacyReputationTokenMintIterator{contract: _OldLegacyReputationToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(target indexed address, value uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *OldLegacyReputationTokenMint, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _OldLegacyReputationToken.contract.WatchLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldLegacyReputationTokenMint)
				if err := _OldLegacyReputationToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// OldLegacyReputationTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OldLegacyReputationToken contract.
type OldLegacyReputationTokenTransferIterator struct {
	Event *OldLegacyReputationTokenTransfer // Event containing the contract specifics and raw log

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
func (it *OldLegacyReputationTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldLegacyReputationTokenTransfer)
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
		it.Event = new(OldLegacyReputationTokenTransfer)
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
func (it *OldLegacyReputationTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldLegacyReputationTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldLegacyReputationTokenTransfer represents a Transfer event raised by the OldLegacyReputationToken contract.
type OldLegacyReputationTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OldLegacyReputationTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OldLegacyReputationToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OldLegacyReputationTokenTransferIterator{contract: _OldLegacyReputationToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_OldLegacyReputationToken *OldLegacyReputationTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OldLegacyReputationTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OldLegacyReputationToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldLegacyReputationTokenTransfer)
				if err := _OldLegacyReputationToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
