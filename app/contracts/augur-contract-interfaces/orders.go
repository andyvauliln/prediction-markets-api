package augur-contract-interfaces

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// OrdersABI is the input ABI used to generate the binding from.
const OrdersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_moneyEscrowed\",\"type\":\"uint256\"},{\"name\":\"_sharesEscrowed\",\"type\":\"uint256\"}],\"name\":\"getOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"}],\"name\":\"assertIsNotBetterPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementTotalEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"name\":\"_sharesFilled\",\"type\":\"uint256\"},{\"name\":\"_tokensFilled\",\"type\":\"uint256\"}],\"name\":\"recordFillOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"getTotalEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"}],\"name\":\"assertIsNotWorsePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOrderMoneyEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"isBetterPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOutcome\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"},{\"name\":\"_moneyEscrowed\",\"type\":\"uint256\"},{\"name\":\"_sharesEscrowed\",\"type\":\"uint256\"},{\"name\":\"_betterOrderId\",\"type\":\"bytes32\"},{\"name\":\"_worseOrderId\",\"type\":\"bytes32\"},{\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"saveOrder\",\"outputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getBestOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getLastOutcomePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getWorseOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getWorstOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getBetterOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_market\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"decrementTotalEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOrderType\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"isWorsePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOrderCreator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOrderSharesEscrowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"removeOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrdersBin is the compiled bytecode used for deploying new contracts.
const OrdersBin = `0x608060405260008054600160a060020a03191633179055611e7a806100256000396000f3006080604052600436106101505763ffffffff60e060020a6000350416630e0a0d7481146101555780631727583c1461017f57806319e54fb3146101c15780632c491293146101f65780632ed5ca291461021a5780633011e16a146102385780633018205f1461025f57806331d98b3f1461029057806337ec114d146102a857806345c92c38146102c95780634a1a342b146102ea5780635cf17bbb146103025780635d1a3b82146103235780636bc29efa1461033b5780637b6eaa6514610385578063804fb410146103af5780638925f9e9146103d35780638e12ebad146103eb57806392eefe9b1461041557806394d26cb51461043657806395ede0321461044e578063bef72fa214610472578063c3c95c7b14610487578063cf3573641461049f578063d9b3d9fe146104db578063e7d80c70146104fc578063ebead05f14610514578063fde996681461052c575b600080fd5b34801561016157600080fd5b5061016d600435610544565b60408051918252519081900360200190f35b34801561018b57600080fd5b5061016d60ff60043516600160a060020a0360243581169060443590606435906084351660a43560c43560e43561010435610559565b3480156101cd57600080fd5b506101e260ff600435166024356044356106c4565b604080519115158252519081900360200190f35b34801561020257600080fd5b506101e2600160a060020a03600435166024356106e6565b34801561022657600080fd5b506101e2600435602435604435610798565b34801561024457600080fd5b506101e2600160a060020a0360043516602435604435610bf9565b34801561026b57600080fd5b50610274610cb6565b60408051600160a060020a039092168252519081900360200190f35b34801561029c57600080fd5b5061016d600435610cc5565b3480156102b457600080fd5b5061016d600160a060020a0360043516610cda565b3480156102d557600080fd5b506101e260ff60043516602435604435610cf5565b3480156102f657600080fd5b5061016d600435610d02565b34801561030e57600080fd5b506101e260ff60043516602435604435610d17565b34801561032f57600080fd5b5061016d600435610d74565b34801561034757600080fd5b5061016d60ff60043516600160a060020a0360243581169060443590606435906084351660a43560c43560e435610104356101243561014435610d89565b34801561039157600080fd5b5061016d60ff60043516600160a060020a03602435166044356112d6565b3480156103bb57600080fd5b5061016d600160a060020a0360043516602435611300565b3480156103df57600080fd5b5061016d60043561132c565b3480156103f757600080fd5b5061016d60ff60043516600160a060020a0360243516604435611341565b34801561042157600080fd5b506101e2600160a060020a0360043516611352565b34801561044257600080fd5b5061016d60043561139c565b34801561045a57600080fd5b506101e2600160a060020a03600435166024356113b1565b34801561047e57600080fd5b5061016d611464565b34801561049357600080fd5b5061027460043561146a565b3480156104ab57600080fd5b506104b7600435611488565b604051808260018111156104c757fe5b60ff16815260200191505060405180910390f35b3480156104e757600080fd5b506101e260ff600435166024356044356114a0565b34801561050857600080fd5b506102746004356114eb565b34801561052057600080fd5b5061016d600435611509565b34801561053857600080fd5b506101e260043561151e565b60009081526002602052604090206007015490565b600060028a8a8a8a8a8a8a8a8a604051602001808a600181111561057957fe5b60ff167f010000000000000000000000000000000000000000000000000000000000000002815260010189600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140188815260200187815260200186600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140185815260200184815260200183815260200182815260200199505050505050505050506040516020818303038152906040526040518082805190602001908083835b6020831061065e5780518252601f19909201916020918201910161063f565b51815160209384036101000a600019018019909216911617905260405191909301945091925050808303816000865af115801561069f573d6000803e3d6000fd5b5050506040513d60208110156106b457600080fd5b50519a9950505050505050505050565b60006106d1848484610d17565b156106db57600080fd5b5060015b9392505050565b6000805460408051600080516020611e2f83398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b15801561073a57600080fd5b505af115801561074e573d6000803e3d6000fd5b505050506040513d602081101561076457600080fd5b5051151561077157600080fd5b50600160a060020a0382166000908152600360205260409020805482019055600192915050565b6000805460408051600080516020611e2f83398151915281523360048201529051839283928392600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b1580156107f257600080fd5b505af1158015610806573d6000803e3d6000fd5b505050506040513d602081101561081c57600080fd5b5051151561082957600080fd5b6000878152600260209081526040808320600181015482517f27ce5b8c0000000000000000000000000000000000000000000000000000000081529251919750600160a060020a0316936327ce5b8c93600480850194919392918390030190829087803b15801561089957600080fd5b505af11580156108ad573d6000803e3d6000fd5b505050506040513d60208110156108c357600080fd5b50516005840154106108d457600080fd5b8615156108e057600080fd5b60098301548611156108f157600080fd5b600a83015485111561090257600080fd5b8260010160009054906101000a9004600160a060020a0316600160a060020a031663bad84c9e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561095757600080fd5b505af115801561096b573d6000803e3d6000fd5b505050506040513d602081101561098157600080fd5b50516008840154111561099357600080fd5b6000915081600684015460ff1660018111156109ab57fe5b14156109df576109d86109cb84600801548761164790919063ffffffff16565b879063ffffffff61165e16565b9150610ab3565b6001600684015460ff1660018111156109f457fe5b1415610ab35760088301546001840154604080517fbad84c9e0000000000000000000000000000000000000000000000000000000081529051610a9b9392600160a060020a03169163bad84c9e9160048083019260209291908290030181600087803b158015610a6357600080fd5b505af1158015610a77573d6000803e3d6000fd5b505050506040513d6020811015610a8d57600080fd5b50519063ffffffff61167016565b9050610ab06109cb868363ffffffff61164716565b91505b6007830154821115610ac457600080fd5b6007830180548390039055600a83018054869003905582546001840154604080517f95ede032000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015260248101899052905191909216916395ede0329160448083019260209291908290030181600087803b158015610b4d57600080fd5b505af1158015610b61573d6000803e3d6000fd5b505050506040513d6020811015610b7757600080fd5b5050600983018054879003905560078301541515610bec57600a83015415610b9e57600080fd5b600983015415610bad57600080fd5b610bb687611685565b5060006008840181905560048401805473ffffffffffffffffffffffffffffffffffffffff19169055600b8401819055600c8401555b5060019695505050505050565b6000805460408051600080516020611e2f83398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b158015610c4d57600080fd5b505af1158015610c61573d6000803e3d6000fd5b505050506040513d6020811015610c7757600080fd5b50511515610c8457600080fd5b50600160a060020a03929092166000908152600360209081526040808320938352600193840190915290209190915590565b600054600160a060020a031690565b60009081526002602052604090206008015490565b600160a060020a031660009081526003602052604090205490565b60006106d18484846114a0565b6000908152600260205260409020600a015490565b600080846001811115610d2657fe5b1415610d46575060008181526002602052604090206008015482116106df565b6001846001811115610d5457fe5b14156106df575060008181526002602052604090206008015482106106df565b60009081526002602052604090206005015490565b6000805460408051600080516020611e2f833981519152815233600482015290518392600160a060020a031691633f08882f91602480830192602092919082900301818787803b158015610ddc57600080fd5b505af1158015610df0573d6000803e3d6000fd5b505050506040513d6020811015610e0657600080fd5b50511515610e1357600080fd5b8b600160a060020a03166327ce5b8c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610e5157600080fd5b505af1158015610e65573d6000803e3d6000fd5b505050506040513d6020811015610e7b57600080fd5b50518810610e8857600080fd5b610e998d8d8d8d8d438e8e8e610559565b91506002600083600019166000191681526020019081526020016000209050308160000160006101000a815481600160a060020a030219169083600160a060020a031602179055508b8160010160006101000a815481600160a060020a030219169083600160a060020a03160217905550818160030181600019169055508c8160060160006101000a81548160ff02191690836001811115610f3757fe5b02179055508781600501819055508981600801819055508a8160070181905550888160040160006101000a815481600160a060020a030219169083600160a060020a031602179055508681600a01819055508060000160009054906101000a9004600160a060020a0316600160a060020a0316632c4912938d896040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561100357600080fd5b505af1158015611017573d6000803e3d6000fd5b505050506040513d602081101561102d57600080fd5b5050600981018690556110418186866117ad565b506000809054906101000a9004600160a060020a0316600160a060020a0316634e94c8296040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561109457600080fd5b505af11580156110a8573d6000803e3d6000fd5b505050506040513d60208110156110be57600080fd5b8101908080519060200190929190505050600160a060020a0316630ae415748e8d8d8d8c8c8a8a8a60010160009054906101000a9004600160a060020a0316600160a060020a031663870c426d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561113a57600080fd5b505af115801561114e573d6000803e3d6000fd5b505050506040513d602081101561116457600080fd5b505160018c015460058d0154604080517f65957bf5000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a03909216916365957bf5916024808201926020929091908290030181600087803b1580156111d457600080fd5b505af11580156111e8573d6000803e3d6000fd5b505050506040513d60208110156111fe57600080fd5b505160405160e060020a63ffffffff8d16028152600401808b600181111561122257fe5b60ff16815260208082019b909b5260408082019a909a52600160a060020a03988916606082015260808101979097525060a086019490945260c085019290925260e084015283166101008301529091166101208201529051610140808301945090918290030181600087803b15801561129a57600080fd5b505af11580156112ae573d6000803e3d6000fd5b505050506040513d60208110156112c457600080fd5b5050509b9a5050505050505050505050565b6000600460006112e7858588611b11565b8152602081019190915260400160002054949350505050565b600160a060020a0391909116600090815260036020908152604080832093835260019093019052205490565b6000908152600260205260409020600c015490565b6000600560006112e7858588611b11565b60008054600160a060020a0316331461136a57600080fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b6000908152600260205260409020600b015490565b6000805460408051600080516020611e2f83398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b15801561140557600080fd5b505af1158015611419573d6000803e3d6000fd5b505050506040513d602081101561142f57600080fd5b5051151561143c57600080fd5b50600160a060020a038216600090815260036020526040902080548290039055600192915050565b60015481565b600090815260026020526040902060010154600160a060020a031690565b60009081526002602052604090206006015460ff1690565b6000808460018111156114af57fe5b14156114d057506000818152600260205260409020600801548211156106df565b506000818152600260205260409020600801548210156106df565b600090815260026020526040902060040154600160a060020a031690565b60009081526002602052604090206009015490565b6000805460408051600080516020611e2f83398151915281523360048201529051600160a060020a0390921691633f08882f9160248082019260209290919082900301818787803b15801561157257600080fd5b505af1158015611586573d6000803e3d6000fd5b505050506040513d602081101561159c57600080fd5b505115156115a957600080fd5b6115b282611685565b505060009081526002602081905260408220805473ffffffffffffffffffffffffffffffffffffffff19908116825560018083018054831690559282018054821690556003820184905560048201805490911690556005810183905560068101805460ff19169055600781018390556008810183905560098101839055600a8101839055600b8101839055600c019190915590565b600080828481151561165557fe5b04949350505050565b6000828201838110156106df57600080fd5b60008282111561167f57600080fd5b50900390565b6000818152600260205260408120600681015460018201546005830154600b840154600c9094015460ff90931693600160a060020a0390921692909190866004876116d187878a611b11565b8152602081019190915260400160002054141561170a5780600460006116f887878a611b11565b81526020810191909152604001600020555b866005600061171a87878a611b11565b8152602081019190915260400160002054141561175357816005600061174187878a611b11565b81526020810191909152604001600020555b811561176e576000828152600260205260409020600c018190555b8015611789576000818152600260205260409020600b018290555b5050506000938452505060026020525060408120600b8101829055600c0155600190565b60018301546005840154600685015460009283928392839260049284926117e192600160a060020a0316919060ff16611b11565b81526020810191909152604001600090812054600189015460058a81015460068c015493975090939261182192600160a060020a0316919060ff16611b11565b81526020808201929092526040908101600090812054815483517ff39ec1f70000000000000000000000000000000000000000000000000000000081527f4f7264657273466574636865720000000000000000000000000000000000000060048201529351919650600160a060020a03169363f39ec1f793602480820194929392918390030190829087803b1580156118b957600080fd5b505af11580156118cd573d6000803e3d6000fd5b505050506040513d60208110156118e357600080fd5b5051600688015460088901546040517f3b01cf3c000000000000000000000000000000000000000000000000000000008152929350600160a060020a03841692633b01cf3c9260ff169190879087908c908c906004018087600181111561194657fe5b60ff1681526020810196909652506040808601949094526060850192909252608084015260a0830152805160c0808401945091929091908290030181600087803b15801561199357600080fd5b505af11580156119a7573d6000803e3d6000fd5b505050506040513d60408110156119bd57600080fd5b50805160209091015190965094506000600688015460ff1660018111156119e057fe5b1415611a425760038701546001880154600889015460058a0154611a0f9392600160a060020a03169190611c25565b6003880154600189015460088a015460058b0154939650611a3b93600160a060020a0390921691611cbb565b9150611a9a565b60038701546001880154600889015460058a0154611a6b9392600160a060020a03169190611d37565b6003880154600189015460088a015460058b0154939650611a9793600160a060020a0390921691611db2565b91505b6003870154831415611aab57600095505b6003870154821415611abc57600094505b8515611ae15760038701546000878152600260205260409020600c0155600b87018690555b8415610bec575050505060038301546000828152600260205260409020600b0155600c9092019190915550600190565b600060028484846040516020018084600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401838152602001826001811115611b5857fe5b60ff167f010000000000000000000000000000000000000000000000000000000000000002815260010193505050506040516020818303038152906040526040518082805190602001908083835b60208310611bc55780518252601f199092019160209182019101611ba6565b51815160209384036101000a600019018019909216911617905260405191909301945091925050808303816000865af1158015611c06573d6000803e3d6000fd5b5050506040513d6020811015611c1b57600080fd5b5051949350505050565b60008060046000611c3887866000611b11565b81526020810191909152604001600020549050801580611c68575060008181526002602052604090206008015484115b15611c90578560046000611c7e88876000611b11565b81526020810191909152604001600020555b60046000611ca087866000611b11565b81526020810191909152604001600020549695505050505050565b60008060056000611cce87866000611b11565b81526020810191909152604001600020549050801580611cff57506000818152600260205260409020600801548411155b15611d27578560056000611d1588876000611b11565b81526020810191909152604001600020555b60056000611ca087866000611b11565b60008060046000611d4a87866001611b11565b81526020810191909152604001600020549050801580611d7a575060008181526002602052604090206008015484105b15611da2578560046000611d9088876001611b11565b81526020810191909152604001600020555b60046000611ca087866001611b11565b60008060056000611dc587866001611b11565b81526020810191909152604001600020549050801580611df657506000818152600260205260409020600801548410155b15611e1e578560056000611e0c88876001611b11565b81526020810191909152604001600020555b60056000611ca087866001611b1156003f08882f00000000000000000000000000000000000000000000000000000000a165627a7a72305820b33586a0f51481722feda5d2c3f953b7f2a6d38cfc03264df2943a8c1a018f260029`

// DeployOrders deploys a new Ethereum contract, binding an instance of Orders to it.
func DeployOrders(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Orders, error) {
	parsed, err := abi.JSON(strings.NewReader(OrdersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrdersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Orders{OrdersCaller: OrdersCaller{contract: contract}, OrdersTransactor: OrdersTransactor{contract: contract}, OrdersFilterer: OrdersFilterer{contract: contract}}, nil
}

// Orders is an auto generated Go binding around an Ethereum contract.
type Orders struct {
	OrdersCaller     // Read-only binding to the contract
	OrdersTransactor // Write-only binding to the contract
	OrdersFilterer   // Log filterer for contract events
}

// OrdersCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrdersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrdersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrdersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrdersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrdersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrdersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrdersSession struct {
	Contract     *Orders           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrdersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrdersCallerSession struct {
	Contract *OrdersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OrdersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrdersTransactorSession struct {
	Contract     *OrdersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrdersRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrdersRaw struct {
	Contract *Orders // Generic contract binding to access the raw methods on
}

// OrdersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrdersCallerRaw struct {
	Contract *OrdersCaller // Generic read-only contract binding to access the raw methods on
}

// OrdersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrdersTransactorRaw struct {
	Contract *OrdersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrders creates a new instance of Orders, bound to a specific deployed contract.
func NewOrders(address common.Address, backend bind.ContractBackend) (*Orders, error) {
	contract, err := bindOrders(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orders{OrdersCaller: OrdersCaller{contract: contract}, OrdersTransactor: OrdersTransactor{contract: contract}, OrdersFilterer: OrdersFilterer{contract: contract}}, nil
}

// NewOrdersCaller creates a new read-only instance of Orders, bound to a specific deployed contract.
func NewOrdersCaller(address common.Address, caller bind.ContractCaller) (*OrdersCaller, error) {
	contract, err := bindOrders(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrdersCaller{contract: contract}, nil
}

// NewOrdersTransactor creates a new write-only instance of Orders, bound to a specific deployed contract.
func NewOrdersTransactor(address common.Address, transactor bind.ContractTransactor) (*OrdersTransactor, error) {
	contract, err := bindOrders(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrdersTransactor{contract: contract}, nil
}

// NewOrdersFilterer creates a new log filterer instance of Orders, bound to a specific deployed contract.
func NewOrdersFilterer(address common.Address, filterer bind.ContractFilterer) (*OrdersFilterer, error) {
	contract, err := bindOrders(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrdersFilterer{contract: contract}, nil
}

// bindOrders binds a generic wrapper to an already deployed contract.
func bindOrders(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrdersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orders *OrdersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orders.Contract.OrdersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orders *OrdersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orders.Contract.OrdersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orders *OrdersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orders.Contract.OrdersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orders *OrdersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orders.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orders *OrdersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orders.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orders *OrdersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orders.Contract.contract.Transact(opts, method, params...)
}

// AssertIsNotBetterPrice is a free data retrieval call binding the contract method 0x19e54fb3.
//
// Solidity: function assertIsNotBetterPrice(_type uint8, _price uint256, _betterOrderId bytes32) constant returns(bool)
func (_Orders *OrdersCaller) AssertIsNotBetterPrice(opts *bind.CallOpts, _type uint8, _price *big.Int, _betterOrderId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "assertIsNotBetterPrice", _type, _price, _betterOrderId)
	return *ret0, err
}

// AssertIsNotBetterPrice is a free data retrieval call binding the contract method 0x19e54fb3.
//
// Solidity: function assertIsNotBetterPrice(_type uint8, _price uint256, _betterOrderId bytes32) constant returns(bool)
func (_Orders *OrdersSession) AssertIsNotBetterPrice(_type uint8, _price *big.Int, _betterOrderId [32]byte) (bool, error) {
	return _Orders.Contract.AssertIsNotBetterPrice(&_Orders.CallOpts, _type, _price, _betterOrderId)
}

// AssertIsNotBetterPrice is a free data retrieval call binding the contract method 0x19e54fb3.
//
// Solidity: function assertIsNotBetterPrice(_type uint8, _price uint256, _betterOrderId bytes32) constant returns(bool)
func (_Orders *OrdersCallerSession) AssertIsNotBetterPrice(_type uint8, _price *big.Int, _betterOrderId [32]byte) (bool, error) {
	return _Orders.Contract.AssertIsNotBetterPrice(&_Orders.CallOpts, _type, _price, _betterOrderId)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Orders *OrdersCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Orders *OrdersSession) ControllerLookupName() ([32]byte, error) {
	return _Orders.Contract.ControllerLookupName(&_Orders.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() constant returns(bytes32)
func (_Orders *OrdersCallerSession) ControllerLookupName() ([32]byte, error) {
	return _Orders.Contract.ControllerLookupName(&_Orders.CallOpts)
}

// GetAmount is a free data retrieval call binding the contract method 0x0e0a0d74.
//
// Solidity: function getAmount(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCaller) GetAmount(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getAmount", _orderId)
	return *ret0, err
}

// GetAmount is a free data retrieval call binding the contract method 0x0e0a0d74.
//
// Solidity: function getAmount(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersSession) GetAmount(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetAmount(&_Orders.CallOpts, _orderId)
}

// GetAmount is a free data retrieval call binding the contract method 0x0e0a0d74.
//
// Solidity: function getAmount(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCallerSession) GetAmount(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetAmount(&_Orders.CallOpts, _orderId)
}

// GetBestOrderId is a free data retrieval call binding the contract method 0x7b6eaa65.
//
// Solidity: function getBestOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_Orders *OrdersCaller) GetBestOrderId(opts *bind.CallOpts, _type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getBestOrderId", _type, _market, _outcome)
	return *ret0, err
}

// GetBestOrderId is a free data retrieval call binding the contract method 0x7b6eaa65.
//
// Solidity: function getBestOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_Orders *OrdersSession) GetBestOrderId(_type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	return _Orders.Contract.GetBestOrderId(&_Orders.CallOpts, _type, _market, _outcome)
}

// GetBestOrderId is a free data retrieval call binding the contract method 0x7b6eaa65.
//
// Solidity: function getBestOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_Orders *OrdersCallerSession) GetBestOrderId(_type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	return _Orders.Contract.GetBestOrderId(&_Orders.CallOpts, _type, _market, _outcome)
}

// GetBetterOrderId is a free data retrieval call binding the contract method 0x94d26cb5.
//
// Solidity: function getBetterOrderId(_orderId bytes32) constant returns(bytes32)
func (_Orders *OrdersCaller) GetBetterOrderId(opts *bind.CallOpts, _orderId [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getBetterOrderId", _orderId)
	return *ret0, err
}

// GetBetterOrderId is a free data retrieval call binding the contract method 0x94d26cb5.
//
// Solidity: function getBetterOrderId(_orderId bytes32) constant returns(bytes32)
func (_Orders *OrdersSession) GetBetterOrderId(_orderId [32]byte) ([32]byte, error) {
	return _Orders.Contract.GetBetterOrderId(&_Orders.CallOpts, _orderId)
}

// GetBetterOrderId is a free data retrieval call binding the contract method 0x94d26cb5.
//
// Solidity: function getBetterOrderId(_orderId bytes32) constant returns(bytes32)
func (_Orders *OrdersCallerSession) GetBetterOrderId(_orderId [32]byte) ([32]byte, error) {
	return _Orders.Contract.GetBetterOrderId(&_Orders.CallOpts, _orderId)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Orders *OrdersCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Orders *OrdersSession) GetController() (common.Address, error) {
	return _Orders.Contract.GetController(&_Orders.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() constant returns(address)
func (_Orders *OrdersCallerSession) GetController() (common.Address, error) {
	return _Orders.Contract.GetController(&_Orders.CallOpts)
}

// GetLastOutcomePrice is a free data retrieval call binding the contract method 0x804fb410.
//
// Solidity: function getLastOutcomePrice(_market address, _outcome uint256) constant returns(uint256)
func (_Orders *OrdersCaller) GetLastOutcomePrice(opts *bind.CallOpts, _market common.Address, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getLastOutcomePrice", _market, _outcome)
	return *ret0, err
}

// GetLastOutcomePrice is a free data retrieval call binding the contract method 0x804fb410.
//
// Solidity: function getLastOutcomePrice(_market address, _outcome uint256) constant returns(uint256)
func (_Orders *OrdersSession) GetLastOutcomePrice(_market common.Address, _outcome *big.Int) (*big.Int, error) {
	return _Orders.Contract.GetLastOutcomePrice(&_Orders.CallOpts, _market, _outcome)
}

// GetLastOutcomePrice is a free data retrieval call binding the contract method 0x804fb410.
//
// Solidity: function getLastOutcomePrice(_market address, _outcome uint256) constant returns(uint256)
func (_Orders *OrdersCallerSession) GetLastOutcomePrice(_market common.Address, _outcome *big.Int) (*big.Int, error) {
	return _Orders.Contract.GetLastOutcomePrice(&_Orders.CallOpts, _market, _outcome)
}

// GetMarket is a free data retrieval call binding the contract method 0xc3c95c7b.
//
// Solidity: function getMarket(_orderId bytes32) constant returns(address)
func (_Orders *OrdersCaller) GetMarket(opts *bind.CallOpts, _orderId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getMarket", _orderId)
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xc3c95c7b.
//
// Solidity: function getMarket(_orderId bytes32) constant returns(address)
func (_Orders *OrdersSession) GetMarket(_orderId [32]byte) (common.Address, error) {
	return _Orders.Contract.GetMarket(&_Orders.CallOpts, _orderId)
}

// GetMarket is a free data retrieval call binding the contract method 0xc3c95c7b.
//
// Solidity: function getMarket(_orderId bytes32) constant returns(address)
func (_Orders *OrdersCallerSession) GetMarket(_orderId [32]byte) (common.Address, error) {
	return _Orders.Contract.GetMarket(&_Orders.CallOpts, _orderId)
}

// GetOrderCreator is a free data retrieval call binding the contract method 0xe7d80c70.
//
// Solidity: function getOrderCreator(_orderId bytes32) constant returns(address)
func (_Orders *OrdersCaller) GetOrderCreator(opts *bind.CallOpts, _orderId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getOrderCreator", _orderId)
	return *ret0, err
}

// GetOrderCreator is a free data retrieval call binding the contract method 0xe7d80c70.
//
// Solidity: function getOrderCreator(_orderId bytes32) constant returns(address)
func (_Orders *OrdersSession) GetOrderCreator(_orderId [32]byte) (common.Address, error) {
	return _Orders.Contract.GetOrderCreator(&_Orders.CallOpts, _orderId)
}

// GetOrderCreator is a free data retrieval call binding the contract method 0xe7d80c70.
//
// Solidity: function getOrderCreator(_orderId bytes32) constant returns(address)
func (_Orders *OrdersCallerSession) GetOrderCreator(_orderId [32]byte) (common.Address, error) {
	return _Orders.Contract.GetOrderCreator(&_Orders.CallOpts, _orderId)
}

// GetOrderId is a free data retrieval call binding the contract method 0x1727583c.
//
// Solidity: function getOrderId(_type uint8, _market address, _amount uint256, _price uint256, _sender address, _blockNumber uint256, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256) constant returns(bytes32)
func (_Orders *OrdersCaller) GetOrderId(opts *bind.CallOpts, _type uint8, _market common.Address, _amount *big.Int, _price *big.Int, _sender common.Address, _blockNumber *big.Int, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getOrderId", _type, _market, _amount, _price, _sender, _blockNumber, _outcome, _moneyEscrowed, _sharesEscrowed)
	return *ret0, err
}

// GetOrderId is a free data retrieval call binding the contract method 0x1727583c.
//
// Solidity: function getOrderId(_type uint8, _market address, _amount uint256, _price uint256, _sender address, _blockNumber uint256, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256) constant returns(bytes32)
func (_Orders *OrdersSession) GetOrderId(_type uint8, _market common.Address, _amount *big.Int, _price *big.Int, _sender common.Address, _blockNumber *big.Int, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int) ([32]byte, error) {
	return _Orders.Contract.GetOrderId(&_Orders.CallOpts, _type, _market, _amount, _price, _sender, _blockNumber, _outcome, _moneyEscrowed, _sharesEscrowed)
}

// GetOrderId is a free data retrieval call binding the contract method 0x1727583c.
//
// Solidity: function getOrderId(_type uint8, _market address, _amount uint256, _price uint256, _sender address, _blockNumber uint256, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256) constant returns(bytes32)
func (_Orders *OrdersCallerSession) GetOrderId(_type uint8, _market common.Address, _amount *big.Int, _price *big.Int, _sender common.Address, _blockNumber *big.Int, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int) ([32]byte, error) {
	return _Orders.Contract.GetOrderId(&_Orders.CallOpts, _type, _market, _amount, _price, _sender, _blockNumber, _outcome, _moneyEscrowed, _sharesEscrowed)
}

// GetOrderMoneyEscrowed is a free data retrieval call binding the contract method 0x4a1a342b.
//
// Solidity: function getOrderMoneyEscrowed(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCaller) GetOrderMoneyEscrowed(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getOrderMoneyEscrowed", _orderId)
	return *ret0, err
}

// GetOrderMoneyEscrowed is a free data retrieval call binding the contract method 0x4a1a342b.
//
// Solidity: function getOrderMoneyEscrowed(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersSession) GetOrderMoneyEscrowed(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetOrderMoneyEscrowed(&_Orders.CallOpts, _orderId)
}

// GetOrderMoneyEscrowed is a free data retrieval call binding the contract method 0x4a1a342b.
//
// Solidity: function getOrderMoneyEscrowed(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCallerSession) GetOrderMoneyEscrowed(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetOrderMoneyEscrowed(&_Orders.CallOpts, _orderId)
}

// GetOrderSharesEscrowed is a free data retrieval call binding the contract method 0xebead05f.
//
// Solidity: function getOrderSharesEscrowed(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCaller) GetOrderSharesEscrowed(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getOrderSharesEscrowed", _orderId)
	return *ret0, err
}

// GetOrderSharesEscrowed is a free data retrieval call binding the contract method 0xebead05f.
//
// Solidity: function getOrderSharesEscrowed(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersSession) GetOrderSharesEscrowed(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetOrderSharesEscrowed(&_Orders.CallOpts, _orderId)
}

// GetOrderSharesEscrowed is a free data retrieval call binding the contract method 0xebead05f.
//
// Solidity: function getOrderSharesEscrowed(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCallerSession) GetOrderSharesEscrowed(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetOrderSharesEscrowed(&_Orders.CallOpts, _orderId)
}

// GetOrderType is a free data retrieval call binding the contract method 0xcf357364.
//
// Solidity: function getOrderType(_orderId bytes32) constant returns(uint8)
func (_Orders *OrdersCaller) GetOrderType(opts *bind.CallOpts, _orderId [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getOrderType", _orderId)
	return *ret0, err
}

// GetOrderType is a free data retrieval call binding the contract method 0xcf357364.
//
// Solidity: function getOrderType(_orderId bytes32) constant returns(uint8)
func (_Orders *OrdersSession) GetOrderType(_orderId [32]byte) (uint8, error) {
	return _Orders.Contract.GetOrderType(&_Orders.CallOpts, _orderId)
}

// GetOrderType is a free data retrieval call binding the contract method 0xcf357364.
//
// Solidity: function getOrderType(_orderId bytes32) constant returns(uint8)
func (_Orders *OrdersCallerSession) GetOrderType(_orderId [32]byte) (uint8, error) {
	return _Orders.Contract.GetOrderType(&_Orders.CallOpts, _orderId)
}

// GetOutcome is a free data retrieval call binding the contract method 0x5d1a3b82.
//
// Solidity: function getOutcome(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCaller) GetOutcome(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getOutcome", _orderId)
	return *ret0, err
}

// GetOutcome is a free data retrieval call binding the contract method 0x5d1a3b82.
//
// Solidity: function getOutcome(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersSession) GetOutcome(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetOutcome(&_Orders.CallOpts, _orderId)
}

// GetOutcome is a free data retrieval call binding the contract method 0x5d1a3b82.
//
// Solidity: function getOutcome(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCallerSession) GetOutcome(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetOutcome(&_Orders.CallOpts, _orderId)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCaller) GetPrice(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getPrice", _orderId)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersSession) GetPrice(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetPrice(&_Orders.CallOpts, _orderId)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(_orderId bytes32) constant returns(uint256)
func (_Orders *OrdersCallerSession) GetPrice(_orderId [32]byte) (*big.Int, error) {
	return _Orders.Contract.GetPrice(&_Orders.CallOpts, _orderId)
}

// GetTotalEscrowed is a free data retrieval call binding the contract method 0x37ec114d.
//
// Solidity: function getTotalEscrowed(_market address) constant returns(uint256)
func (_Orders *OrdersCaller) GetTotalEscrowed(opts *bind.CallOpts, _market common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getTotalEscrowed", _market)
	return *ret0, err
}

// GetTotalEscrowed is a free data retrieval call binding the contract method 0x37ec114d.
//
// Solidity: function getTotalEscrowed(_market address) constant returns(uint256)
func (_Orders *OrdersSession) GetTotalEscrowed(_market common.Address) (*big.Int, error) {
	return _Orders.Contract.GetTotalEscrowed(&_Orders.CallOpts, _market)
}

// GetTotalEscrowed is a free data retrieval call binding the contract method 0x37ec114d.
//
// Solidity: function getTotalEscrowed(_market address) constant returns(uint256)
func (_Orders *OrdersCallerSession) GetTotalEscrowed(_market common.Address) (*big.Int, error) {
	return _Orders.Contract.GetTotalEscrowed(&_Orders.CallOpts, _market)
}

// GetWorseOrderId is a free data retrieval call binding the contract method 0x8925f9e9.
//
// Solidity: function getWorseOrderId(_orderId bytes32) constant returns(bytes32)
func (_Orders *OrdersCaller) GetWorseOrderId(opts *bind.CallOpts, _orderId [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getWorseOrderId", _orderId)
	return *ret0, err
}

// GetWorseOrderId is a free data retrieval call binding the contract method 0x8925f9e9.
//
// Solidity: function getWorseOrderId(_orderId bytes32) constant returns(bytes32)
func (_Orders *OrdersSession) GetWorseOrderId(_orderId [32]byte) ([32]byte, error) {
	return _Orders.Contract.GetWorseOrderId(&_Orders.CallOpts, _orderId)
}

// GetWorseOrderId is a free data retrieval call binding the contract method 0x8925f9e9.
//
// Solidity: function getWorseOrderId(_orderId bytes32) constant returns(bytes32)
func (_Orders *OrdersCallerSession) GetWorseOrderId(_orderId [32]byte) ([32]byte, error) {
	return _Orders.Contract.GetWorseOrderId(&_Orders.CallOpts, _orderId)
}

// GetWorstOrderId is a free data retrieval call binding the contract method 0x8e12ebad.
//
// Solidity: function getWorstOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_Orders *OrdersCaller) GetWorstOrderId(opts *bind.CallOpts, _type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "getWorstOrderId", _type, _market, _outcome)
	return *ret0, err
}

// GetWorstOrderId is a free data retrieval call binding the contract method 0x8e12ebad.
//
// Solidity: function getWorstOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_Orders *OrdersSession) GetWorstOrderId(_type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	return _Orders.Contract.GetWorstOrderId(&_Orders.CallOpts, _type, _market, _outcome)
}

// GetWorstOrderId is a free data retrieval call binding the contract method 0x8e12ebad.
//
// Solidity: function getWorstOrderId(_type uint8, _market address, _outcome uint256) constant returns(bytes32)
func (_Orders *OrdersCallerSession) GetWorstOrderId(_type uint8, _market common.Address, _outcome *big.Int) ([32]byte, error) {
	return _Orders.Contract.GetWorstOrderId(&_Orders.CallOpts, _type, _market, _outcome)
}

// IsBetterPrice is a free data retrieval call binding the contract method 0x5cf17bbb.
//
// Solidity: function isBetterPrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_Orders *OrdersCaller) IsBetterPrice(opts *bind.CallOpts, _type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "isBetterPrice", _type, _price, _orderId)
	return *ret0, err
}

// IsBetterPrice is a free data retrieval call binding the contract method 0x5cf17bbb.
//
// Solidity: function isBetterPrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_Orders *OrdersSession) IsBetterPrice(_type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	return _Orders.Contract.IsBetterPrice(&_Orders.CallOpts, _type, _price, _orderId)
}

// IsBetterPrice is a free data retrieval call binding the contract method 0x5cf17bbb.
//
// Solidity: function isBetterPrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_Orders *OrdersCallerSession) IsBetterPrice(_type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	return _Orders.Contract.IsBetterPrice(&_Orders.CallOpts, _type, _price, _orderId)
}

// IsWorsePrice is a free data retrieval call binding the contract method 0xd9b3d9fe.
//
// Solidity: function isWorsePrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_Orders *OrdersCaller) IsWorsePrice(opts *bind.CallOpts, _type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Orders.contract.Call(opts, out, "isWorsePrice", _type, _price, _orderId)
	return *ret0, err
}

// IsWorsePrice is a free data retrieval call binding the contract method 0xd9b3d9fe.
//
// Solidity: function isWorsePrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_Orders *OrdersSession) IsWorsePrice(_type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	return _Orders.Contract.IsWorsePrice(&_Orders.CallOpts, _type, _price, _orderId)
}

// IsWorsePrice is a free data retrieval call binding the contract method 0xd9b3d9fe.
//
// Solidity: function isWorsePrice(_type uint8, _price uint256, _orderId bytes32) constant returns(bool)
func (_Orders *OrdersCallerSession) IsWorsePrice(_type uint8, _price *big.Int, _orderId [32]byte) (bool, error) {
	return _Orders.Contract.IsWorsePrice(&_Orders.CallOpts, _type, _price, _orderId)
}

// AssertIsNotWorsePrice is a paid mutator transaction binding the contract method 0x45c92c38.
//
// Solidity: function assertIsNotWorsePrice(_type uint8, _price uint256, _worseOrderId bytes32) returns(bool)
func (_Orders *OrdersTransactor) AssertIsNotWorsePrice(opts *bind.TransactOpts, _type uint8, _price *big.Int, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _Orders.contract.Transact(opts, "assertIsNotWorsePrice", _type, _price, _worseOrderId)
}

// AssertIsNotWorsePrice is a paid mutator transaction binding the contract method 0x45c92c38.
//
// Solidity: function assertIsNotWorsePrice(_type uint8, _price uint256, _worseOrderId bytes32) returns(bool)
func (_Orders *OrdersSession) AssertIsNotWorsePrice(_type uint8, _price *big.Int, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _Orders.Contract.AssertIsNotWorsePrice(&_Orders.TransactOpts, _type, _price, _worseOrderId)
}

// AssertIsNotWorsePrice is a paid mutator transaction binding the contract method 0x45c92c38.
//
// Solidity: function assertIsNotWorsePrice(_type uint8, _price uint256, _worseOrderId bytes32) returns(bool)
func (_Orders *OrdersTransactorSession) AssertIsNotWorsePrice(_type uint8, _price *big.Int, _worseOrderId [32]byte) (*types.Transaction, error) {
	return _Orders.Contract.AssertIsNotWorsePrice(&_Orders.TransactOpts, _type, _price, _worseOrderId)
}

// DecrementTotalEscrowed is a paid mutator transaction binding the contract method 0x95ede032.
//
// Solidity: function decrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_Orders *OrdersTransactor) DecrementTotalEscrowed(opts *bind.TransactOpts, _market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Orders.contract.Transact(opts, "decrementTotalEscrowed", _market, _amount)
}

// DecrementTotalEscrowed is a paid mutator transaction binding the contract method 0x95ede032.
//
// Solidity: function decrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_Orders *OrdersSession) DecrementTotalEscrowed(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Orders.Contract.DecrementTotalEscrowed(&_Orders.TransactOpts, _market, _amount)
}

// DecrementTotalEscrowed is a paid mutator transaction binding the contract method 0x95ede032.
//
// Solidity: function decrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_Orders *OrdersTransactorSession) DecrementTotalEscrowed(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Orders.Contract.DecrementTotalEscrowed(&_Orders.TransactOpts, _market, _amount)
}

// IncrementTotalEscrowed is a paid mutator transaction binding the contract method 0x2c491293.
//
// Solidity: function incrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_Orders *OrdersTransactor) IncrementTotalEscrowed(opts *bind.TransactOpts, _market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Orders.contract.Transact(opts, "incrementTotalEscrowed", _market, _amount)
}

// IncrementTotalEscrowed is a paid mutator transaction binding the contract method 0x2c491293.
//
// Solidity: function incrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_Orders *OrdersSession) IncrementTotalEscrowed(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Orders.Contract.IncrementTotalEscrowed(&_Orders.TransactOpts, _market, _amount)
}

// IncrementTotalEscrowed is a paid mutator transaction binding the contract method 0x2c491293.
//
// Solidity: function incrementTotalEscrowed(_market address, _amount uint256) returns(bool)
func (_Orders *OrdersTransactorSession) IncrementTotalEscrowed(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Orders.Contract.IncrementTotalEscrowed(&_Orders.TransactOpts, _market, _amount)
}

// RecordFillOrder is a paid mutator transaction binding the contract method 0x2ed5ca29.
//
// Solidity: function recordFillOrder(_orderId bytes32, _sharesFilled uint256, _tokensFilled uint256) returns(bool)
func (_Orders *OrdersTransactor) RecordFillOrder(opts *bind.TransactOpts, _orderId [32]byte, _sharesFilled *big.Int, _tokensFilled *big.Int) (*types.Transaction, error) {
	return _Orders.contract.Transact(opts, "recordFillOrder", _orderId, _sharesFilled, _tokensFilled)
}

// RecordFillOrder is a paid mutator transaction binding the contract method 0x2ed5ca29.
//
// Solidity: function recordFillOrder(_orderId bytes32, _sharesFilled uint256, _tokensFilled uint256) returns(bool)
func (_Orders *OrdersSession) RecordFillOrder(_orderId [32]byte, _sharesFilled *big.Int, _tokensFilled *big.Int) (*types.Transaction, error) {
	return _Orders.Contract.RecordFillOrder(&_Orders.TransactOpts, _orderId, _sharesFilled, _tokensFilled)
}

// RecordFillOrder is a paid mutator transaction binding the contract method 0x2ed5ca29.
//
// Solidity: function recordFillOrder(_orderId bytes32, _sharesFilled uint256, _tokensFilled uint256) returns(bool)
func (_Orders *OrdersTransactorSession) RecordFillOrder(_orderId [32]byte, _sharesFilled *big.Int, _tokensFilled *big.Int) (*types.Transaction, error) {
	return _Orders.Contract.RecordFillOrder(&_Orders.TransactOpts, _orderId, _sharesFilled, _tokensFilled)
}

// RemoveOrder is a paid mutator transaction binding the contract method 0xfde99668.
//
// Solidity: function removeOrder(_orderId bytes32) returns(bool)
func (_Orders *OrdersTransactor) RemoveOrder(opts *bind.TransactOpts, _orderId [32]byte) (*types.Transaction, error) {
	return _Orders.contract.Transact(opts, "removeOrder", _orderId)
}

// RemoveOrder is a paid mutator transaction binding the contract method 0xfde99668.
//
// Solidity: function removeOrder(_orderId bytes32) returns(bool)
func (_Orders *OrdersSession) RemoveOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _Orders.Contract.RemoveOrder(&_Orders.TransactOpts, _orderId)
}

// RemoveOrder is a paid mutator transaction binding the contract method 0xfde99668.
//
// Solidity: function removeOrder(_orderId bytes32) returns(bool)
func (_Orders *OrdersTransactorSession) RemoveOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _Orders.Contract.RemoveOrder(&_Orders.TransactOpts, _orderId)
}

// SaveOrder is a paid mutator transaction binding the contract method 0x6bc29efa.
//
// Solidity: function saveOrder(_type uint8, _market address, _amount uint256, _price uint256, _sender address, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(_orderId bytes32)
func (_Orders *OrdersTransactor) SaveOrder(opts *bind.TransactOpts, _type uint8, _market common.Address, _amount *big.Int, _price *big.Int, _sender common.Address, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Orders.contract.Transact(opts, "saveOrder", _type, _market, _amount, _price, _sender, _outcome, _moneyEscrowed, _sharesEscrowed, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// SaveOrder is a paid mutator transaction binding the contract method 0x6bc29efa.
//
// Solidity: function saveOrder(_type uint8, _market address, _amount uint256, _price uint256, _sender address, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(_orderId bytes32)
func (_Orders *OrdersSession) SaveOrder(_type uint8, _market common.Address, _amount *big.Int, _price *big.Int, _sender common.Address, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Orders.Contract.SaveOrder(&_Orders.TransactOpts, _type, _market, _amount, _price, _sender, _outcome, _moneyEscrowed, _sharesEscrowed, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// SaveOrder is a paid mutator transaction binding the contract method 0x6bc29efa.
//
// Solidity: function saveOrder(_type uint8, _market address, _amount uint256, _price uint256, _sender address, _outcome uint256, _moneyEscrowed uint256, _sharesEscrowed uint256, _betterOrderId bytes32, _worseOrderId bytes32, _tradeGroupId bytes32) returns(_orderId bytes32)
func (_Orders *OrdersTransactorSession) SaveOrder(_type uint8, _market common.Address, _amount *big.Int, _price *big.Int, _sender common.Address, _outcome *big.Int, _moneyEscrowed *big.Int, _sharesEscrowed *big.Int, _betterOrderId [32]byte, _worseOrderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Orders.Contract.SaveOrder(&_Orders.TransactOpts, _type, _market, _amount, _price, _sender, _outcome, _moneyEscrowed, _sharesEscrowed, _betterOrderId, _worseOrderId, _tradeGroupId)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Orders *OrdersTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Orders.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Orders *OrdersSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Orders.Contract.SetController(&_Orders.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_Orders *OrdersTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Orders.Contract.SetController(&_Orders.TransactOpts, _controller)
}

// SetPrice is a paid mutator transaction binding the contract method 0x3011e16a.
//
// Solidity: function setPrice(_market address, _outcome uint256, _price uint256) returns(bool)
func (_Orders *OrdersTransactor) SetPrice(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Orders.contract.Transact(opts, "setPrice", _market, _outcome, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x3011e16a.
//
// Solidity: function setPrice(_market address, _outcome uint256, _price uint256) returns(bool)
func (_Orders *OrdersSession) SetPrice(_market common.Address, _outcome *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Orders.Contract.SetPrice(&_Orders.TransactOpts, _market, _outcome, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x3011e16a.
//
// Solidity: function setPrice(_market address, _outcome uint256, _price uint256) returns(bool)
func (_Orders *OrdersTransactorSession) SetPrice(_market common.Address, _outcome *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Orders.Contract.SetPrice(&_Orders.TransactOpts, _market, _outcome, _price)
}