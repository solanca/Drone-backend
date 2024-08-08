// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Policy

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PolicyContractPolicy is an auto generated low-level Go binding around an user-defined struct.
type PolicyContractPolicy struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Exists    bool
}

// PolicyMetaData contains all meta data concerning the Policy contract.
var PolicyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"loggerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"}],\"name\":\"PolicyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"PolicyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"}],\"name\":\"PolicyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_endTime\",\"type\":\"string\"}],\"name\":\"createPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deletePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolicies\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structPolicyContract.Policy[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structPolicyContract.Policy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"getPolicyByZone\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structPolicyContract.Policy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logger\",\"outputs\":[{\"internalType\":\"contractLoggingContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"policies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_endTime\",\"type\":\"string\"}],\"name\":\"updatePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161237f38038061237f833981810160405281019061003191906100d5565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610100565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a48261007b565b9050919050565b6100b48161009a565b81146100be575f80fd5b50565b5f815190506100cf816100ab565b92915050565b5f602082840312156100ea576100e9610077565b5b5f6100f7848285016100c1565b91505092915050565b6122728061010d5f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c8063882b413e11610064578063882b413e1461011d5780639d47b49c14610139578063d3e8948314610169578063f24ccbfe1461019d578063fed27c8f146101bb57610091565b80632b07fce314610095578063394d9578146100c55780633b04f6f1146100e157806361b8ce8c146100ff575b5f80fd5b6100af60048036038101906100aa91906112b9565b6101d7565b6040516100bc919061140f565b60405180910390f35b6100df60048036038101906100da9190611585565b6103d8565b005b6100e96104e4565b6040516100f69190611742565b60405180910390f35b61010761079b565b6040516101149190611771565b60405180910390f35b6101376004803603810190610132919061178a565b6107a1565b005b610153600480360381019061014e9190611826565b6108ae565b604051610160919061140f565b60405180910390f35b610183600480360381019061017e91906112b9565b610c38565b6040516101949594939291906118b7565b60405180910390f35b6101a5610d90565b6040516101b29190611990565b60405180910390f35b6101d560048036038101906101d091906112b9565b610db5565b005b6101df611247565b6001548210801561021b57505f82815481106101fe576101fd6119a9565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b61025a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025190611a20565b60405180910390fd5b5f828154811061026d5761026c6119a9565b5b905f5260205f2090600502016040518060a00160405290815f8201548152602001600182015481526020016002820180546102a790611a6b565b80601f01602080910402602001604051908101604052809291908181526020018280546102d390611a6b565b801561031e5780601f106102f55761010080835404028352916020019161031e565b820191905f5260205f20905b81548152906001019060200180831161030157829003601f168201915b5050505050815260200160038201805461033790611a6b565b80601f016020809104026020016040519081016040528092919081815260200182805461036390611a6b565b80156103ae5780601f10610385576101008083540402835291602001916103ae565b820191905f5260205f20905b81548152906001019060200180831161039157829003601f168201915b50505050508152602001600482015f9054906101000a900460ff1615151515815250509050919050565b5f6040518060a00160405280600154815260200185815260200184815260200183815260200160011515815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f01556020820151816001015560408201518160020190816104529190611c2f565b5060608201518160030190816104689190611c2f565b506080820151816004015f6101000a81548160ff02191690831515021790555050506001547fe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e8484846040516104c093929190611cfe565b60405180910390a260015f8154809291906104da90611d6e565b9190505550505050565b60605f805b5f80549050811015610545575f8181548110610508576105076119a9565b5b905f5260205f2090600502016004015f9054906101000a900460ff161561053857818061053490611d6e565b9250505b80806001019150506104e9565b505f8167ffffffffffffffff81111561056157610560611461565b5b60405190808252806020026020018201604052801561059a57816020015b610587611247565b81526020019060019003908161057f5790505b5090505f805b5f80549050811015610791575f81815481106105bf576105be6119a9565b5b905f5260205f2090600502016004015f9054906101000a900460ff1615610784575f81815481106105f3576105f26119a9565b5b905f5260205f2090600502016040518060a00160405290815f82015481526020016001820154815260200160028201805461062d90611a6b565b80601f016020809104026020016040519081016040528092919081815260200182805461065990611a6b565b80156106a45780601f1061067b576101008083540402835291602001916106a4565b820191905f5260205f20905b81548152906001019060200180831161068757829003601f168201915b505050505081526020016003820180546106bd90611a6b565b80601f01602080910402602001604051908101604052809291908181526020018280546106e990611a6b565b80156107345780601f1061070b57610100808354040283529160200191610734565b820191905f5260205f20905b81548152906001019060200180831161071757829003601f168201915b50505050508152602001600482015f9054906101000a900460ff16151515158152505083838151811061076a576107696119a9565b5b6020026020010181905250818061078090611d6e565b9250505b80806001019150506105a0565b5081935050505090565b60015481565b600154841080156107dd57505f84815481106107c0576107bf6119a9565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b61081c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081390611a20565b60405180910390fd5b5f8085815481106108305761082f6119a9565b5b905f5260205f2090600502019050838160010181905550828160020190816108589190611c2f565b508181600301908161086a9190611c2f565b50847f89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f9185858560405161089f93929190611cfe565b60405180910390a25050505050565b6108b6611247565b5f5b5f80549050811015610bf757825f82815481106108d8576108d76119a9565b5b905f5260205f2090600502016001015414801561092057505f8181548110610903576109026119a9565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b15610bea575f6109505f838154811061093c5761093b6119a9565b5b905f5260205f2090600502015f0154610e99565b61097b5f8481548110610966576109656119a9565b5b905f5260205f20906005020160010154611017565b5f848154811061098e5761098d6119a9565b5b905f5260205f2090600502016002015f85815481106109b0576109af6119a9565b5b905f5260205f2090600502016003016040516020016109d29493929190611f97565b604051602081830303815290604052905060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166397f8abd8826040518263ffffffff1660e01b8152600401610a3d919061204a565b5f604051808303815f87803b158015610a54575f80fd5b505af1158015610a66573d5f803e3d5ffd5b505050505f8281548110610a7d57610a7c6119a9565b5b905f5260205f2090600502016040518060a00160405290815f820154815260200160018201548152602001600282018054610ab790611a6b565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae390611a6b565b8015610b2e5780601f10610b0557610100808354040283529160200191610b2e565b820191905f5260205f20905b815481529060010190602001808311610b1157829003601f168201915b50505050508152602001600382018054610b4790611a6b565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7390611a6b565b8015610bbe5780601f10610b9557610100808354040283529160200191610bbe565b820191905f5260205f20905b815481529060010190602001808311610ba157829003601f168201915b50505050508152602001600482015f9054906101000a900460ff16151515158152505092505050610c33565b80806001019150506108b8565b506040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2a906120c7565b60405180910390fd5b919050565b5f8181548110610c46575f80fd5b905f5260205f2090600502015f91509050805f015490806001015490806002018054610c7190611a6b565b80601f0160208091040260200160405190810160405280929190818152602001828054610c9d90611a6b565b8015610ce85780601f10610cbf57610100808354040283529160200191610ce8565b820191905f5260205f20905b815481529060010190602001808311610ccb57829003601f168201915b505050505090806003018054610cfd90611a6b565b80601f0160208091040260200160405190810160405280929190818152602001828054610d2990611a6b565b8015610d745780601f10610d4b57610100808354040283529160200191610d74565b820191905f5260205f20905b815481529060010190602001808311610d5757829003601f168201915b505050505090806004015f9054906101000a900460ff16905085565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481108015610df157505f8181548110610dd457610dd36119a9565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b610e30576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2790611a20565b60405180910390fd5b5f808281548110610e4457610e436119a9565b5b905f5260205f2090600502016004015f6101000a81548160ff021916908315150217905550807f433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d60405160405180910390a250565b60605f8203610edf576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611012565b5f8290505f5b5f8214610f0e578080610ef790611d6e565b915050600a82610f079190612112565b9150610ee5565b5f8167ffffffffffffffff811115610f2957610f28611461565b5b6040519080825280601f01601f191660200182016040528015610f5b5781602001600182028036833780820191505090505b5090505f8290505b5f861461100a57600181610f779190612142565b90505f600a8088610f889190612112565b610f929190612175565b87610f9d9190612142565b6030610fa991906121c2565b90505f8160f81b905080848481518110610fc657610fc56119a9565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a886110019190612112565b97505050610f63565b819450505050505b919050565b60605f820361105d576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611242565b5f80831290505f8161106f578361107a565b83611079906121f6565b5b90505f5b5f82146110a757808061109090611d6e565b915050600a826110a09190612112565b915061107e565b82156110bc5780806110b890611d6e565b9150505b5f8167ffffffffffffffff8111156110d7576110d6611461565b5b6040519080825280601f01601f1916602001820160405280156111095781602001600182028036833780820191505090505b5090505f8290508461111b5786611126565b86611125906121f6565b5b93505b5f84146111d05760018161113d9190612142565b90505f600a808661114e9190612112565b6111589190612175565b856111639190612142565b603061116f91906121c2565b90505f8160f81b90508084848151811061118c5761118b6119a9565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a866111c79190612112565b95505050611129565b8415611239577f2d00000000000000000000000000000000000000000000000000000000000000825f8151811061120a576112096119a9565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505b81955050505050505b919050565b6040518060a001604052805f81526020015f815260200160608152602001606081526020015f151581525090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61129881611286565b81146112a2575f80fd5b50565b5f813590506112b38161128f565b92915050565b5f602082840312156112ce576112cd61127e565b5b5f6112db848285016112a5565b91505092915050565b6112ed81611286565b82525050565b5f819050919050565b611305816112f3565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61134d8261130b565b6113578185611315565b9350611367818560208601611325565b61137081611333565b840191505092915050565b5f8115159050919050565b61138f8161137b565b82525050565b5f60a083015f8301516113aa5f8601826112e4565b5060208301516113bd60208601826112fc565b50604083015184820360408601526113d58282611343565b915050606083015184820360608601526113ef8282611343565b91505060808301516114046080860182611386565b508091505092915050565b5f6020820190508181035f8301526114278184611395565b905092915050565b611438816112f3565b8114611442575f80fd5b50565b5f813590506114538161142f565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61149782611333565b810181811067ffffffffffffffff821117156114b6576114b5611461565b5b80604052505050565b5f6114c8611275565b90506114d4828261148e565b919050565b5f67ffffffffffffffff8211156114f3576114f2611461565b5b6114fc82611333565b9050602081019050919050565b828183375f83830152505050565b5f611529611524846114d9565b6114bf565b9050828152602081018484840111156115455761154461145d565b5b611550848285611509565b509392505050565b5f82601f83011261156c5761156b611459565b5b813561157c848260208601611517565b91505092915050565b5f805f6060848603121561159c5761159b61127e565b5b5f6115a986828701611445565b935050602084013567ffffffffffffffff8111156115ca576115c9611282565b5b6115d686828701611558565b925050604084013567ffffffffffffffff8111156115f7576115f6611282565b5b61160386828701611558565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f60a083015f83015161164b5f8601826112e4565b50602083015161165e60208601826112fc565b50604083015184820360408601526116768282611343565b915050606083015184820360608601526116908282611343565b91505060808301516116a56080860182611386565b508091505092915050565b5f6116bb8383611636565b905092915050565b5f602082019050919050565b5f6116d98261160d565b6116e38185611617565b9350836020820285016116f585611627565b805f5b85811015611730578484038952815161171185826116b0565b945061171c836116c3565b925060208a019950506001810190506116f8565b50829750879550505050505092915050565b5f6020820190508181035f83015261175a81846116cf565b905092915050565b61176b81611286565b82525050565b5f6020820190506117845f830184611762565b92915050565b5f805f80608085870312156117a2576117a161127e565b5b5f6117af878288016112a5565b94505060206117c087828801611445565b935050604085013567ffffffffffffffff8111156117e1576117e0611282565b5b6117ed87828801611558565b925050606085013567ffffffffffffffff81111561180e5761180d611282565b5b61181a87828801611558565b91505092959194509250565b5f6020828403121561183b5761183a61127e565b5b5f61184884828501611445565b91505092915050565b61185a816112f3565b82525050565b5f82825260208201905092915050565b5f61187a8261130b565b6118848185611860565b9350611894818560208601611325565b61189d81611333565b840191505092915050565b6118b18161137b565b82525050565b5f60a0820190506118ca5f830188611762565b6118d76020830187611851565b81810360408301526118e98186611870565b905081810360608301526118fd8185611870565b905061190c60808301846118a8565b9695505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61195861195361194e84611916565b611935565b611916565b9050919050565b5f6119698261193e565b9050919050565b5f61197a8261195f565b9050919050565b61198a81611970565b82525050565b5f6020820190506119a35f830184611981565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f506f6c69637920646f6573206e6f7420657869737400000000000000000000005f82015250565b5f611a0a601583611860565b9150611a15826119d6565b602082019050919050565b5f6020820190508181035f830152611a37816119fe565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611a8257607f821691505b602082108103611a9557611a94611a3e565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611af77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611abc565b611b018683611abc565b95508019841693508086168417925050509392505050565b5f611b33611b2e611b2984611286565b611935565b611286565b9050919050565b5f819050919050565b611b4c83611b19565b611b60611b5882611b3a565b848454611ac8565b825550505050565b5f90565b611b74611b68565b611b7f818484611b43565b505050565b5b81811015611ba257611b975f82611b6c565b600181019050611b85565b5050565b601f821115611be757611bb881611a9b565b611bc184611aad565b81016020851015611bd0578190505b611be4611bdc85611aad565b830182611b84565b50505b505050565b5f82821c905092915050565b5f611c075f1984600802611bec565b1980831691505092915050565b5f611c1f8383611bf8565b9150826002028217905092915050565b611c388261130b565b67ffffffffffffffff811115611c5157611c50611461565b5b611c5b8254611a6b565b611c66828285611ba6565b5f60209050601f831160018114611c97575f8415611c85578287015190505b611c8f8582611c14565b865550611cf6565b601f198416611ca586611a9b565b5f5b82811015611ccc57848901518255600182019150602085019450602081019050611ca7565b86831015611ce95784890151611ce5601f891682611bf8565b8355505b6001600288020188555050505b505050505050565b5f606082019050611d115f830186611851565b8181036020830152611d238185611870565b90508181036040830152611d378184611870565b9050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611d7882611286565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611daa57611da9611d41565b5b600182019050919050565b5f81905092915050565b7f49443a20000000000000000000000000000000000000000000000000000000005f82015250565b5f611df3600483611db5565b9150611dfe82611dbf565b600482019050919050565b5f611e138261130b565b611e1d8185611db5565b9350611e2d818560208601611325565b80840191505092915050565b7f2c205a6f6e653a200000000000000000000000000000000000000000000000005f82015250565b5f611e6d600883611db5565b9150611e7882611e39565b600882019050919050565b7f2c20537461727454696d653a20000000000000000000000000000000000000005f82015250565b5f611eb7600d83611db5565b9150611ec282611e83565b600d82019050919050565b5f8154611ed981611a6b565b611ee38186611db5565b9450600182165f8114611efd5760018114611f1257611f44565b60ff1983168652811515820286019350611f44565b611f1b85611a9b565b5f5b83811015611f3c57815481890152600182019150602081019050611f1d565b838801955050505b50505092915050565b7f2c20456e6454696d653a200000000000000000000000000000000000000000005f82015250565b5f611f81600b83611db5565b9150611f8c82611f4d565b600b82019050919050565b5f611fa182611de7565b9150611fad8287611e09565b9150611fb882611e61565b9150611fc48286611e09565b9150611fcf82611eab565b9150611fdb8285611ecd565b9150611fe682611f75565b9150611ff28284611ecd565b915081905095945050505050565b7f50525000000000000000000000000000000000000000000000000000000000005f82015250565b5f612034600383611860565b915061203f82612000565b602082019050919050565b5f6040820190508181035f83015261206181612028565b905081810360208301526120758184611870565b905092915050565b7f506f6c696379206e6f7420666f756e64000000000000000000000000000000005f82015250565b5f6120b1601083611860565b91506120bc8261207d565b602082019050919050565b5f6020820190508181035f8301526120de816120a5565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61211c82611286565b915061212783611286565b925082612137576121366120e5565b5b828204905092915050565b5f61214c82611286565b915061215783611286565b925082820390508181111561216f5761216e611d41565b5b92915050565b5f61217f82611286565b915061218a83611286565b925082820261219881611286565b915082820484148315176121af576121ae611d41565b5b5092915050565b5f60ff82169050919050565b5f6121cc826121b6565b91506121d7836121b6565b9250828201905060ff8111156121f0576121ef611d41565b5b92915050565b5f612200826112f3565b91507f8000000000000000000000000000000000000000000000000000000000000000820361223257612231611d41565b5b815f03905091905056fea2646970667358221220091ac6b29fbc785c4a44a51a7f190e8a711899e03c916d88efd73b674d7dca7a64736f6c634300081a0033",
}

// PolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use PolicyMetaData.ABI instead.
var PolicyABI = PolicyMetaData.ABI

// PolicyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolicyMetaData.Bin instead.
var PolicyBin = PolicyMetaData.Bin

// DeployPolicy deploys a new Ethereum contract, binding an instance of Policy to it.
func DeployPolicy(auth *bind.TransactOpts, backend bind.ContractBackend, loggerAddress common.Address) (common.Address, *types.Transaction, *Policy, error) {
	parsed, err := PolicyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolicyBin), backend, loggerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Policy{PolicyCaller: PolicyCaller{contract: contract}, PolicyTransactor: PolicyTransactor{contract: contract}, PolicyFilterer: PolicyFilterer{contract: contract}}, nil
}

// Policy is an auto generated Go binding around an Ethereum contract.
type Policy struct {
	PolicyCaller     // Read-only binding to the contract
	PolicyTransactor // Write-only binding to the contract
	PolicyFilterer   // Log filterer for contract events
}

// PolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolicySession struct {
	Contract     *Policy           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolicyCallerSession struct {
	Contract *PolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolicyTransactorSession struct {
	Contract     *PolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolicyRaw struct {
	Contract *Policy // Generic contract binding to access the raw methods on
}

// PolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolicyCallerRaw struct {
	Contract *PolicyCaller // Generic read-only contract binding to access the raw methods on
}

// PolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolicyTransactorRaw struct {
	Contract *PolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolicy creates a new instance of Policy, bound to a specific deployed contract.
func NewPolicy(address common.Address, backend bind.ContractBackend) (*Policy, error) {
	contract, err := bindPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Policy{PolicyCaller: PolicyCaller{contract: contract}, PolicyTransactor: PolicyTransactor{contract: contract}, PolicyFilterer: PolicyFilterer{contract: contract}}, nil
}

// NewPolicyCaller creates a new read-only instance of Policy, bound to a specific deployed contract.
func NewPolicyCaller(address common.Address, caller bind.ContractCaller) (*PolicyCaller, error) {
	contract, err := bindPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyCaller{contract: contract}, nil
}

// NewPolicyTransactor creates a new write-only instance of Policy, bound to a specific deployed contract.
func NewPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*PolicyTransactor, error) {
	contract, err := bindPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyTransactor{contract: contract}, nil
}

// NewPolicyFilterer creates a new log filterer instance of Policy, bound to a specific deployed contract.
func NewPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*PolicyFilterer, error) {
	contract, err := bindPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolicyFilterer{contract: contract}, nil
}

// bindPolicy binds a generic wrapper to an already deployed contract.
func bindPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolicyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.PolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transact(opts, method, params...)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns((uint256,int256,string,string,bool)[])
func (_Policy *PolicyCaller) GetPolicies(opts *bind.CallOpts) ([]PolicyContractPolicy, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "getPolicies")

	if err != nil {
		return *new([]PolicyContractPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new([]PolicyContractPolicy)).(*[]PolicyContractPolicy)

	return out0, err

}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns((uint256,int256,string,string,bool)[])
func (_Policy *PolicySession) GetPolicies() ([]PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicies(&_Policy.CallOpts)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns((uint256,int256,string,string,bool)[])
func (_Policy *PolicyCallerSession) GetPolicies() ([]PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicies(&_Policy.CallOpts)
}

// GetPolicy is a free data retrieval call binding the contract method 0x2b07fce3.
//
// Solidity: function getPolicy(uint256 _id) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicyCaller) GetPolicy(opts *bind.CallOpts, _id *big.Int) (PolicyContractPolicy, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "getPolicy", _id)

	if err != nil {
		return *new(PolicyContractPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(PolicyContractPolicy)).(*PolicyContractPolicy)

	return out0, err

}

// GetPolicy is a free data retrieval call binding the contract method 0x2b07fce3.
//
// Solidity: function getPolicy(uint256 _id) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicySession) GetPolicy(_id *big.Int) (PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicy(&_Policy.CallOpts, _id)
}

// GetPolicy is a free data retrieval call binding the contract method 0x2b07fce3.
//
// Solidity: function getPolicy(uint256 _id) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicyCallerSession) GetPolicy(_id *big.Int) (PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicy(&_Policy.CallOpts, _id)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Policy *PolicyCaller) Logger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "logger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Policy *PolicySession) Logger() (common.Address, error) {
	return _Policy.Contract.Logger(&_Policy.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Policy *PolicyCallerSession) Logger() (common.Address, error) {
	return _Policy.Contract.Logger(&_Policy.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Policy *PolicyCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Policy *PolicySession) NextId() (*big.Int, error) {
	return _Policy.Contract.NextId(&_Policy.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Policy *PolicyCallerSession) NextId() (*big.Int, error) {
	return _Policy.Contract.NextId(&_Policy.CallOpts)
}

// Policies is a free data retrieval call binding the contract method 0xd3e89483.
//
// Solidity: function policies(uint256 ) view returns(uint256 id, int256 zone, string startTime, string endTime, bool exists)
func (_Policy *PolicyCaller) Policies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Exists    bool
}, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "policies", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Zone      *big.Int
		StartTime string
		EndTime   string
		Exists    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Zone = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.EndTime = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Exists = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Policies is a free data retrieval call binding the contract method 0xd3e89483.
//
// Solidity: function policies(uint256 ) view returns(uint256 id, int256 zone, string startTime, string endTime, bool exists)
func (_Policy *PolicySession) Policies(arg0 *big.Int) (struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Exists    bool
}, error) {
	return _Policy.Contract.Policies(&_Policy.CallOpts, arg0)
}

// Policies is a free data retrieval call binding the contract method 0xd3e89483.
//
// Solidity: function policies(uint256 ) view returns(uint256 id, int256 zone, string startTime, string endTime, bool exists)
func (_Policy *PolicyCallerSession) Policies(arg0 *big.Int) (struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Exists    bool
}, error) {
	return _Policy.Contract.Policies(&_Policy.CallOpts, arg0)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x394d9578.
//
// Solidity: function createPolicy(int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicyTransactor) CreatePolicy(opts *bind.TransactOpts, _zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "createPolicy", _zone, _startTime, _endTime)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x394d9578.
//
// Solidity: function createPolicy(int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicySession) CreatePolicy(_zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.Contract.CreatePolicy(&_Policy.TransactOpts, _zone, _startTime, _endTime)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x394d9578.
//
// Solidity: function createPolicy(int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicyTransactorSession) CreatePolicy(_zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.Contract.CreatePolicy(&_Policy.TransactOpts, _zone, _startTime, _endTime)
}

// DeletePolicy is a paid mutator transaction binding the contract method 0xfed27c8f.
//
// Solidity: function deletePolicy(uint256 _id) returns()
func (_Policy *PolicyTransactor) DeletePolicy(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "deletePolicy", _id)
}

// DeletePolicy is a paid mutator transaction binding the contract method 0xfed27c8f.
//
// Solidity: function deletePolicy(uint256 _id) returns()
func (_Policy *PolicySession) DeletePolicy(_id *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.DeletePolicy(&_Policy.TransactOpts, _id)
}

// DeletePolicy is a paid mutator transaction binding the contract method 0xfed27c8f.
//
// Solidity: function deletePolicy(uint256 _id) returns()
func (_Policy *PolicyTransactorSession) DeletePolicy(_id *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.DeletePolicy(&_Policy.TransactOpts, _id)
}

// GetPolicyByZone is a paid mutator transaction binding the contract method 0x9d47b49c.
//
// Solidity: function getPolicyByZone(int256 _zone) returns((uint256,int256,string,string,bool))
func (_Policy *PolicyTransactor) GetPolicyByZone(opts *bind.TransactOpts, _zone *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "getPolicyByZone", _zone)
}

// GetPolicyByZone is a paid mutator transaction binding the contract method 0x9d47b49c.
//
// Solidity: function getPolicyByZone(int256 _zone) returns((uint256,int256,string,string,bool))
func (_Policy *PolicySession) GetPolicyByZone(_zone *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.GetPolicyByZone(&_Policy.TransactOpts, _zone)
}

// GetPolicyByZone is a paid mutator transaction binding the contract method 0x9d47b49c.
//
// Solidity: function getPolicyByZone(int256 _zone) returns((uint256,int256,string,string,bool))
func (_Policy *PolicyTransactorSession) GetPolicyByZone(_zone *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.GetPolicyByZone(&_Policy.TransactOpts, _zone)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0x882b413e.
//
// Solidity: function updatePolicy(uint256 _id, int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicyTransactor) UpdatePolicy(opts *bind.TransactOpts, _id *big.Int, _zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "updatePolicy", _id, _zone, _startTime, _endTime)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0x882b413e.
//
// Solidity: function updatePolicy(uint256 _id, int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicySession) UpdatePolicy(_id *big.Int, _zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.Contract.UpdatePolicy(&_Policy.TransactOpts, _id, _zone, _startTime, _endTime)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0x882b413e.
//
// Solidity: function updatePolicy(uint256 _id, int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicyTransactorSession) UpdatePolicy(_id *big.Int, _zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.Contract.UpdatePolicy(&_Policy.TransactOpts, _id, _zone, _startTime, _endTime)
}

// PolicyPolicyCreatedIterator is returned from FilterPolicyCreated and is used to iterate over the raw logs and unpacked data for PolicyCreated events raised by the Policy contract.
type PolicyPolicyCreatedIterator struct {
	Event *PolicyPolicyCreated // Event containing the contract specifics and raw log

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
func (it *PolicyPolicyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyPolicyCreated)
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
		it.Event = new(PolicyPolicyCreated)
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
func (it *PolicyPolicyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyPolicyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyPolicyCreated represents a PolicyCreated event raised by the Policy contract.
type PolicyPolicyCreated struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolicyCreated is a free log retrieval operation binding the contract event 0xe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e.
//
// Solidity: event PolicyCreated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) FilterPolicyCreated(opts *bind.FilterOpts, id []*big.Int) (*PolicyPolicyCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "PolicyCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &PolicyPolicyCreatedIterator{contract: _Policy.contract, event: "PolicyCreated", logs: logs, sub: sub}, nil
}

// WatchPolicyCreated is a free log subscription operation binding the contract event 0xe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e.
//
// Solidity: event PolicyCreated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) WatchPolicyCreated(opts *bind.WatchOpts, sink chan<- *PolicyPolicyCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "PolicyCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyPolicyCreated)
				if err := _Policy.contract.UnpackLog(event, "PolicyCreated", log); err != nil {
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

// ParsePolicyCreated is a log parse operation binding the contract event 0xe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e.
//
// Solidity: event PolicyCreated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) ParsePolicyCreated(log types.Log) (*PolicyPolicyCreated, error) {
	event := new(PolicyPolicyCreated)
	if err := _Policy.contract.UnpackLog(event, "PolicyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicyPolicyDeletedIterator is returned from FilterPolicyDeleted and is used to iterate over the raw logs and unpacked data for PolicyDeleted events raised by the Policy contract.
type PolicyPolicyDeletedIterator struct {
	Event *PolicyPolicyDeleted // Event containing the contract specifics and raw log

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
func (it *PolicyPolicyDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyPolicyDeleted)
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
		it.Event = new(PolicyPolicyDeleted)
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
func (it *PolicyPolicyDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyPolicyDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyPolicyDeleted represents a PolicyDeleted event raised by the Policy contract.
type PolicyPolicyDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPolicyDeleted is a free log retrieval operation binding the contract event 0x433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d.
//
// Solidity: event PolicyDeleted(uint256 indexed id)
func (_Policy *PolicyFilterer) FilterPolicyDeleted(opts *bind.FilterOpts, id []*big.Int) (*PolicyPolicyDeletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "PolicyDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return &PolicyPolicyDeletedIterator{contract: _Policy.contract, event: "PolicyDeleted", logs: logs, sub: sub}, nil
}

// WatchPolicyDeleted is a free log subscription operation binding the contract event 0x433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d.
//
// Solidity: event PolicyDeleted(uint256 indexed id)
func (_Policy *PolicyFilterer) WatchPolicyDeleted(opts *bind.WatchOpts, sink chan<- *PolicyPolicyDeleted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "PolicyDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyPolicyDeleted)
				if err := _Policy.contract.UnpackLog(event, "PolicyDeleted", log); err != nil {
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

// ParsePolicyDeleted is a log parse operation binding the contract event 0x433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d.
//
// Solidity: event PolicyDeleted(uint256 indexed id)
func (_Policy *PolicyFilterer) ParsePolicyDeleted(log types.Log) (*PolicyPolicyDeleted, error) {
	event := new(PolicyPolicyDeleted)
	if err := _Policy.contract.UnpackLog(event, "PolicyDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicyPolicyUpdatedIterator is returned from FilterPolicyUpdated and is used to iterate over the raw logs and unpacked data for PolicyUpdated events raised by the Policy contract.
type PolicyPolicyUpdatedIterator struct {
	Event *PolicyPolicyUpdated // Event containing the contract specifics and raw log

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
func (it *PolicyPolicyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyPolicyUpdated)
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
		it.Event = new(PolicyPolicyUpdated)
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
func (it *PolicyPolicyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyPolicyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyPolicyUpdated represents a PolicyUpdated event raised by the Policy contract.
type PolicyPolicyUpdated struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolicyUpdated is a free log retrieval operation binding the contract event 0x89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f91.
//
// Solidity: event PolicyUpdated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) FilterPolicyUpdated(opts *bind.FilterOpts, id []*big.Int) (*PolicyPolicyUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "PolicyUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &PolicyPolicyUpdatedIterator{contract: _Policy.contract, event: "PolicyUpdated", logs: logs, sub: sub}, nil
}

// WatchPolicyUpdated is a free log subscription operation binding the contract event 0x89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f91.
//
// Solidity: event PolicyUpdated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) WatchPolicyUpdated(opts *bind.WatchOpts, sink chan<- *PolicyPolicyUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "PolicyUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyPolicyUpdated)
				if err := _Policy.contract.UnpackLog(event, "PolicyUpdated", log); err != nil {
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

// ParsePolicyUpdated is a log parse operation binding the contract event 0x89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f91.
//
// Solidity: event PolicyUpdated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) ParsePolicyUpdated(log types.Log) (*PolicyPolicyUpdated, error) {
	event := new(PolicyPolicyUpdated)
	if err := _Policy.contract.UnpackLog(event, "PolicyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
