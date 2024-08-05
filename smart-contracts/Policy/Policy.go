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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"}],\"name\":\"PolicyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"PolicyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"}],\"name\":\"PolicyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_endTime\",\"type\":\"string\"}],\"name\":\"createPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deletePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolicies\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structPolicyContract.Policy[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structPolicyContract.Policy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"getPolicyByZone\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structPolicyContract.Policy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"policies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_endTime\",\"type\":\"string\"}],\"name\":\"updatePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506118a18061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c8063882b413e11610059578063882b413e146101125780639d47b49c1461012e578063d3e894831461015e578063fed27c8f1461019257610086565b80632b07fce31461008a578063394d9578146100ba5780633b04f6f1146100d657806361b8ce8c146100f4575b5f80fd5b6100a4600480360381019061009f9190610d83565b6101ae565b6040516100b19190610ed9565b60405180910390f35b6100d460048036038101906100cf919061104f565b6103af565b005b6100de6104bc565b6040516100eb919061120c565b60405180910390f35b6100fc610773565b604051610109919061123b565b60405180910390f35b61012c60048036038101906101279190611254565b610779565b005b610148600480360381019061014391906112f0565b610887565b6040516101559190610ed9565b60405180910390f35b61017860048036038101906101739190610d83565b610acb565b604051610189959493929190611381565b60405180910390f35b6101ac60048036038101906101a79190610d83565b610c23565b005b6101b6610d11565b600154821080156101f257505f82815481106101d5576101d46113e0565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b610231576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022890611457565b60405180910390fd5b5f8281548110610244576102436113e0565b5b905f5260205f2090600502016040518060a00160405290815f82015481526020016001820154815260200160028201805461027e906114a2565b80601f01602080910402602001604051908101604052809291908181526020018280546102aa906114a2565b80156102f55780601f106102cc576101008083540402835291602001916102f5565b820191905f5260205f20905b8154815290600101906020018083116102d857829003601f168201915b5050505050815260200160038201805461030e906114a2565b80601f016020809104026020016040519081016040528092919081815260200182805461033a906114a2565b80156103855780601f1061035c57610100808354040283529160200191610385565b820191905f5260205f20905b81548152906001019060200180831161036857829003601f168201915b50505050508152602001600482015f9054906101000a900460ff1615151515815250509050919050565b5f6040518060a00160405280600154815260200185815260200184815260200183815260200160011515815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f0155602082015181600101556040820151816002019081610429919061166f565b50606082015181600301908161043f919061166f565b506080820151816004015f6101000a81548160ff02191690831515021790555050507fe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e600154848484604051610498949392919061173e565b60405180910390a160015f8154809291906104b2906117bc565b9190505550505050565b60605f805b5f8054905081101561051d575f81815481106104e0576104df6113e0565b5b905f5260205f2090600502016004015f9054906101000a900460ff161561051057818061050c906117bc565b9250505b80806001019150506104c1565b505f8167ffffffffffffffff81111561053957610538610f2b565b5b60405190808252806020026020018201604052801561057257816020015b61055f610d11565b8152602001906001900390816105575790505b5090505f805b5f80549050811015610769575f8181548110610597576105966113e0565b5b905f5260205f2090600502016004015f9054906101000a900460ff161561075c575f81815481106105cb576105ca6113e0565b5b905f5260205f2090600502016040518060a00160405290815f820154815260200160018201548152602001600282018054610605906114a2565b80601f0160208091040260200160405190810160405280929190818152602001828054610631906114a2565b801561067c5780601f106106535761010080835404028352916020019161067c565b820191905f5260205f20905b81548152906001019060200180831161065f57829003601f168201915b50505050508152602001600382018054610695906114a2565b80601f01602080910402602001604051908101604052809291908181526020018280546106c1906114a2565b801561070c5780601f106106e35761010080835404028352916020019161070c565b820191905f5260205f20905b8154815290600101906020018083116106ef57829003601f168201915b50505050508152602001600482015f9054906101000a900460ff161515151581525050838381518110610742576107416113e0565b5b60200260200101819052508180610758906117bc565b9250505b8080600101915050610578565b5081935050505090565b60015481565b600154841080156107b557505f8481548110610798576107976113e0565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b6107f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107eb90611457565b60405180910390fd5b5f808581548110610808576108076113e0565b5b905f5260205f209060050201905083816001018190555082816002019081610830919061166f565b5081816003019081610842919061166f565b507f89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f9185858585604051610878949392919061173e565b60405180910390a15050505050565b61088f610d11565b5f5b5f80549050811015610a8a57825f82815481106108b1576108b06113e0565b5b905f5260205f209060050201600101541480156108f957505f81815481106108dc576108db6113e0565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b15610a7d575f8181548110610911576109106113e0565b5b905f5260205f2090600502016040518060a00160405290815f82015481526020016001820154815260200160028201805461094b906114a2565b80601f0160208091040260200160405190810160405280929190818152602001828054610977906114a2565b80156109c25780601f10610999576101008083540402835291602001916109c2565b820191905f5260205f20905b8154815290600101906020018083116109a557829003601f168201915b505050505081526020016003820180546109db906114a2565b80601f0160208091040260200160405190810160405280929190818152602001828054610a07906114a2565b8015610a525780601f10610a2957610100808354040283529160200191610a52565b820191905f5260205f20905b815481529060010190602001808311610a3557829003601f168201915b50505050508152602001600482015f9054906101000a900460ff161515151581525050915050610ac6565b8080600101915050610891565b506040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abd9061184d565b60405180910390fd5b919050565b5f8181548110610ad9575f80fd5b905f5260205f2090600502015f91509050805f015490806001015490806002018054610b04906114a2565b80601f0160208091040260200160405190810160405280929190818152602001828054610b30906114a2565b8015610b7b5780601f10610b5257610100808354040283529160200191610b7b565b820191905f5260205f20905b815481529060010190602001808311610b5e57829003601f168201915b505050505090806003018054610b90906114a2565b80601f0160208091040260200160405190810160405280929190818152602001828054610bbc906114a2565b8015610c075780601f10610bde57610100808354040283529160200191610c07565b820191905f5260205f20905b815481529060010190602001808311610bea57829003601f168201915b505050505090806004015f9054906101000a900460ff16905085565b60015481108015610c5f57505f8181548110610c4257610c416113e0565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b610c9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9590611457565b60405180910390fd5b5f808281548110610cb257610cb16113e0565b5b905f5260205f2090600502016004015f6101000a81548160ff0219169083151502179055507f433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d81604051610d06919061123b565b60405180910390a150565b6040518060a001604052805f81526020015f815260200160608152602001606081526020015f151581525090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610d6281610d50565b8114610d6c575f80fd5b50565b5f81359050610d7d81610d59565b92915050565b5f60208284031215610d9857610d97610d48565b5b5f610da584828501610d6f565b91505092915050565b610db781610d50565b82525050565b5f819050919050565b610dcf81610dbd565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e1782610dd5565b610e218185610ddf565b9350610e31818560208601610def565b610e3a81610dfd565b840191505092915050565b5f8115159050919050565b610e5981610e45565b82525050565b5f60a083015f830151610e745f860182610dae565b506020830151610e876020860182610dc6565b5060408301518482036040860152610e9f8282610e0d565b91505060608301518482036060860152610eb98282610e0d565b9150506080830151610ece6080860182610e50565b508091505092915050565b5f6020820190508181035f830152610ef18184610e5f565b905092915050565b610f0281610dbd565b8114610f0c575f80fd5b50565b5f81359050610f1d81610ef9565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610f6182610dfd565b810181811067ffffffffffffffff82111715610f8057610f7f610f2b565b5b80604052505050565b5f610f92610d3f565b9050610f9e8282610f58565b919050565b5f67ffffffffffffffff821115610fbd57610fbc610f2b565b5b610fc682610dfd565b9050602081019050919050565b828183375f83830152505050565b5f610ff3610fee84610fa3565b610f89565b90508281526020810184848401111561100f5761100e610f27565b5b61101a848285610fd3565b509392505050565b5f82601f83011261103657611035610f23565b5b8135611046848260208601610fe1565b91505092915050565b5f805f6060848603121561106657611065610d48565b5b5f61107386828701610f0f565b935050602084013567ffffffffffffffff81111561109457611093610d4c565b5b6110a086828701611022565b925050604084013567ffffffffffffffff8111156110c1576110c0610d4c565b5b6110cd86828701611022565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f60a083015f8301516111155f860182610dae565b5060208301516111286020860182610dc6565b50604083015184820360408601526111408282610e0d565b9150506060830151848203606086015261115a8282610e0d565b915050608083015161116f6080860182610e50565b508091505092915050565b5f6111858383611100565b905092915050565b5f602082019050919050565b5f6111a3826110d7565b6111ad81856110e1565b9350836020820285016111bf856110f1565b805f5b858110156111fa57848403895281516111db858261117a565b94506111e68361118d565b925060208a019950506001810190506111c2565b50829750879550505050505092915050565b5f6020820190508181035f8301526112248184611199565b905092915050565b61123581610d50565b82525050565b5f60208201905061124e5f83018461122c565b92915050565b5f805f806080858703121561126c5761126b610d48565b5b5f61127987828801610d6f565b945050602061128a87828801610f0f565b935050604085013567ffffffffffffffff8111156112ab576112aa610d4c565b5b6112b787828801611022565b925050606085013567ffffffffffffffff8111156112d8576112d7610d4c565b5b6112e487828801611022565b91505092959194509250565b5f6020828403121561130557611304610d48565b5b5f61131284828501610f0f565b91505092915050565b61132481610dbd565b82525050565b5f82825260208201905092915050565b5f61134482610dd5565b61134e818561132a565b935061135e818560208601610def565b61136781610dfd565b840191505092915050565b61137b81610e45565b82525050565b5f60a0820190506113945f83018861122c565b6113a1602083018761131b565b81810360408301526113b3818661133a565b905081810360608301526113c7818561133a565b90506113d66080830184611372565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f506f6c69637920646f6573206e6f7420657869737400000000000000000000005f82015250565b5f61144160158361132a565b915061144c8261140d565b602082019050919050565b5f6020820190508181035f83015261146e81611435565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806114b957607f821691505b6020821081036114cc576114cb611475565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261152e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826114f3565b61153886836114f3565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61157361156e61156984610d50565b611550565b610d50565b9050919050565b5f819050919050565b61158c83611559565b6115a06115988261157a565b8484546114ff565b825550505050565b5f90565b6115b46115a8565b6115bf818484611583565b505050565b5b818110156115e2576115d75f826115ac565b6001810190506115c5565b5050565b601f821115611627576115f8816114d2565b611601846114e4565b81016020851015611610578190505b61162461161c856114e4565b8301826115c4565b50505b505050565b5f82821c905092915050565b5f6116475f198460080261162c565b1980831691505092915050565b5f61165f8383611638565b9150826002028217905092915050565b61167882610dd5565b67ffffffffffffffff81111561169157611690610f2b565b5b61169b82546114a2565b6116a68282856115e6565b5f60209050601f8311600181146116d7575f84156116c5578287015190505b6116cf8582611654565b865550611736565b601f1984166116e5866114d2565b5f5b8281101561170c578489015182556001820191506020850194506020810190506116e7565b868310156117295784890151611725601f891682611638565b8355505b6001600288020188555050505b505050505050565b5f6080820190506117515f83018761122c565b61175e602083018661131b565b8181036040830152611770818561133a565b90508181036060830152611784818461133a565b905095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6117c682610d50565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117f8576117f761178f565b5b600182019050919050565b7f506f6c696379206e6f7420666f756e64000000000000000000000000000000005f82015250565b5f61183760108361132a565b915061184282611803565b602082019050919050565b5f6020820190508181035f8301526118648161182b565b905091905056fea2646970667358221220ae2b55b6133895df1224717fcc49217efb2d7c68939cb73890b6d2e1367d2da064736f6c634300081a0033",
}

// PolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use PolicyMetaData.ABI instead.
var PolicyABI = PolicyMetaData.ABI

// PolicyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolicyMetaData.Bin instead.
var PolicyBin = PolicyMetaData.Bin

// DeployPolicy deploys a new Ethereum contract, binding an instance of Policy to it.
func DeployPolicy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Policy, error) {
	parsed, err := PolicyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolicyBin), backend)
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

// GetPolicyByZone is a free data retrieval call binding the contract method 0x9d47b49c.
//
// Solidity: function getPolicyByZone(int256 _zone) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicyCaller) GetPolicyByZone(opts *bind.CallOpts, _zone *big.Int) (PolicyContractPolicy, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "getPolicyByZone", _zone)

	if err != nil {
		return *new(PolicyContractPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(PolicyContractPolicy)).(*PolicyContractPolicy)

	return out0, err

}

// GetPolicyByZone is a free data retrieval call binding the contract method 0x9d47b49c.
//
// Solidity: function getPolicyByZone(int256 _zone) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicySession) GetPolicyByZone(_zone *big.Int) (PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicyByZone(&_Policy.CallOpts, _zone)
}

// GetPolicyByZone is a free data retrieval call binding the contract method 0x9d47b49c.
//
// Solidity: function getPolicyByZone(int256 _zone) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicyCallerSession) GetPolicyByZone(_zone *big.Int) (PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicyByZone(&_Policy.CallOpts, _zone)
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
// Solidity: event PolicyCreated(uint256 id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) FilterPolicyCreated(opts *bind.FilterOpts) (*PolicyPolicyCreatedIterator, error) {

	logs, sub, err := _Policy.contract.FilterLogs(opts, "PolicyCreated")
	if err != nil {
		return nil, err
	}
	return &PolicyPolicyCreatedIterator{contract: _Policy.contract, event: "PolicyCreated", logs: logs, sub: sub}, nil
}

// WatchPolicyCreated is a free log subscription operation binding the contract event 0xe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e.
//
// Solidity: event PolicyCreated(uint256 id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) WatchPolicyCreated(opts *bind.WatchOpts, sink chan<- *PolicyPolicyCreated) (event.Subscription, error) {

	logs, sub, err := _Policy.contract.WatchLogs(opts, "PolicyCreated")
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
// Solidity: event PolicyCreated(uint256 id, int256 zone, string startTime, string endTime)
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
// Solidity: event PolicyDeleted(uint256 id)
func (_Policy *PolicyFilterer) FilterPolicyDeleted(opts *bind.FilterOpts) (*PolicyPolicyDeletedIterator, error) {

	logs, sub, err := _Policy.contract.FilterLogs(opts, "PolicyDeleted")
	if err != nil {
		return nil, err
	}
	return &PolicyPolicyDeletedIterator{contract: _Policy.contract, event: "PolicyDeleted", logs: logs, sub: sub}, nil
}

// WatchPolicyDeleted is a free log subscription operation binding the contract event 0x433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d.
//
// Solidity: event PolicyDeleted(uint256 id)
func (_Policy *PolicyFilterer) WatchPolicyDeleted(opts *bind.WatchOpts, sink chan<- *PolicyPolicyDeleted) (event.Subscription, error) {

	logs, sub, err := _Policy.contract.WatchLogs(opts, "PolicyDeleted")
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
// Solidity: event PolicyDeleted(uint256 id)
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
// Solidity: event PolicyUpdated(uint256 id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) FilterPolicyUpdated(opts *bind.FilterOpts) (*PolicyPolicyUpdatedIterator, error) {

	logs, sub, err := _Policy.contract.FilterLogs(opts, "PolicyUpdated")
	if err != nil {
		return nil, err
	}
	return &PolicyPolicyUpdatedIterator{contract: _Policy.contract, event: "PolicyUpdated", logs: logs, sub: sub}, nil
}

// WatchPolicyUpdated is a free log subscription operation binding the contract event 0x89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f91.
//
// Solidity: event PolicyUpdated(uint256 id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) WatchPolicyUpdated(opts *bind.WatchOpts, sink chan<- *PolicyPolicyUpdated) (event.Subscription, error) {

	logs, sub, err := _Policy.contract.WatchLogs(opts, "PolicyUpdated")
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
// Solidity: event PolicyUpdated(uint256 id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) ParsePolicyUpdated(log types.Log) (*PolicyPolicyUpdated, error) {
	event := new(PolicyPolicyUpdated)
	if err := _Policy.contract.UnpackLog(event, "PolicyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
