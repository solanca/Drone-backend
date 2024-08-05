// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Attribute

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

// AttributeContractAttribute is an auto generated low-level Go binding around an user-defined struct.
type AttributeContractAttribute struct {
	Id    *big.Int
	Name  string
	Value []*big.Int
}

// AttributeMetaData contains all meta data concerning the Attribute contract.
var AttributeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"name\":\"AttributeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AttributeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"name\":\"AttributeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"_value\",\"type\":\"uint256[]\"}],\"name\":\"createAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"_value\",\"type\":\"uint256[]\"}],\"name\":\"updateAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"attributes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getAttribute\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structAttributeContract.Attribute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAttributes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structAttributeContract.Attribute[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAttributesByName\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structAttributeContract.Attribute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5061177f8061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c8063493fe80f11610059578063493fe80f1461011057806361b8ce8c146101405780637238b41f1461015e578063d05dcc6a1461017a57610086565b806306516a4a1461008a578063152583de146100ba5780631736a62b146100d85780632822800c146100f4575b5f80fd5b6100a4600480360381019061009f9190610c51565b6101ab565b6040516100b19190610e0c565b60405180910390f35b6100c26103e7565b6040516100cf9190610f3b565b60405180910390f35b6100f260048036038101906100ed9190611049565b610531565b005b61010e600480360381019061010991906110d1565b610603565b005b61012a60048036038101906101259190611147565b6106dc565b6040516101379190610e0c565b60405180910390f35b61014861084d565b6040516101559190611181565b60405180910390f35b61017860048036038101906101739190611147565b610853565b005b610194600480360381019061018f9190611147565b610917565b6040516101a29291906111e2565b60405180910390f35b6101b3610a23565b5f5b5f805490508110156103a6576102735f82815481106101d7576101d6611210565b5b905f5260205f20906003020160010180546101f19061126a565b80601f016020809104026020016040519081016040528092919081815260200182805461021d9061126a565b80156102685780601f1061023f57610100808354040283529160200191610268565b820191905f5260205f20905b81548152906001019060200180831161024b57829003601f168201915b5050505050846109cb565b15610399575f818154811061028b5761028a611210565b5b905f5260205f2090600302016040518060600160405290815f82015481526020016001820180546102bb9061126a565b80601f01602080910402602001604051908101604052809291908181526020018280546102e79061126a565b80156103325780601f1061030957610100808354040283529160200191610332565b820191905f5260205f20905b81548152906001019060200180831161031557829003601f168201915b505050505081526020016002820180548060200260200160405190810160405280929190818152602001828054801561038857602002820191905f5260205f20905b815481526020019060010190808311610374575b5050505050815250509150506103e2565b80806001019150506101b5565b506040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d9906112e4565b60405180910390fd5b919050565b60605f805480602002602001604051908101604052809291908181526020015f905b82821015610528578382905f5260205f2090600302016040518060600160405290815f82015481526020016001820180546104439061126a565b80601f016020809104026020016040519081016040528092919081815260200182805461046f9061126a565b80156104ba5780601f10610491576101008083540402835291602001916104ba565b820191905f5260205f20905b81548152906001019060200180831161049d57829003601f168201915b505050505081526020016002820180548060200260200160405190810160405280929190818152602001828054801561051057602002820191905f5260205f20905b8154815260200190600101908083116104fc575b50505050508152505081526020019060010190610409565b50505050905090565b6001548310610575576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056c9061134c565b60405180910390fd5b5f80848154811061058957610588611210565b5b905f5260205f2090600302019050828160010190816105a89190611507565b50818160020190805190602001906105c1929190610a43565b507f8f26e9a323213025e76e7c58dbf6338a9c36708b0f933b7e5babc20acca4e8d18484846040516105f593929190611642565b60405180910390a150505050565b5f6040518060600160405280600154815260200184815260200183815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f015560208201518160010190816106649190611507565b506040820151816002019080519060200190610681929190610a43565b5050507f92ff6eabb2f26553137551f9eaac5f9e399114f11d71898144a4e7e71f8689f160015483836040516106b993929190611642565b60405180910390a160015f8154809291906106d3906116b2565b91905055505050565b6106e4610a23565b6001548210610728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071f9061134c565b60405180910390fd5b5f80838154811061073c5761073b611210565b5b905f5260205f2090600302016040518060600160405290815f820154815260200160018201805461076c9061126a565b80601f01602080910402602001604051908101604052809291908181526020018280546107989061126a565b80156107e35780601f106107ba576101008083540402835291602001916107e3565b820191905f5260205f20905b8154815290600101906020018083116107c657829003601f168201915b505050505081526020016002820180548060200260200160405190810160405280929190818152602001828054801561083957602002820191905f5260205f20905b815481526020019060010190808311610825575b505050505081525050905080915050919050565b60015481565b6001548110610897576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088e9061134c565b60405180910390fd5b5f81815481106108aa576108a9611210565b5b905f5260205f2090600302015f8082015f9055600182015f6108cc9190610a8e565b600282015f6108db9190610acb565b50507ffe96137d7877697f141204379bad0e6617ff7023d7608f9db3be956d7ceef7118160405161090c9190611181565b60405180910390a150565b5f8181548110610925575f80fd5b905f5260205f2090600302015f91509050805f01549080600101805461094a9061126a565b80601f01602080910402602001604051908101604052809291908181526020018280546109769061126a565b80156109c15780601f10610998576101008083540402835291602001916109c1565b820191905f5260205f20905b8154815290600101906020018083116109a457829003601f168201915b5050505050905082565b5f816040516020016109dd9190611733565b6040516020818303038152906040528051906020012083604051602001610a049190611733565b6040516020818303038152906040528051906020012014905092915050565b60405180606001604052805f815260200160608152602001606081525090565b828054828255905f5260205f20908101928215610a7d579160200282015b82811115610a7c578251825591602001919060010190610a61565b5b509050610a8a9190610ae9565b5090565b508054610a9a9061126a565b5f825580601f10610aab5750610ac8565b601f0160209004905f5260205f2090810190610ac79190610ae9565b5b50565b5080545f8255905f5260205f2090810190610ae69190610ae9565b50565b5b80821115610b00575f815f905550600101610aea565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610b6382610b1d565b810181811067ffffffffffffffff82111715610b8257610b81610b2d565b5b80604052505050565b5f610b94610b04565b9050610ba08282610b5a565b919050565b5f67ffffffffffffffff821115610bbf57610bbe610b2d565b5b610bc882610b1d565b9050602081019050919050565b828183375f83830152505050565b5f610bf5610bf084610ba5565b610b8b565b905082815260208101848484011115610c1157610c10610b19565b5b610c1c848285610bd5565b509392505050565b5f82601f830112610c3857610c37610b15565b5b8135610c48848260208601610be3565b91505092915050565b5f60208284031215610c6657610c65610b0d565b5b5f82013567ffffffffffffffff811115610c8357610c82610b11565b5b610c8f84828501610c24565b91505092915050565b5f819050919050565b610caa81610c98565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610ce282610cb0565b610cec8185610cba565b9350610cfc818560208601610cca565b610d0581610b1d565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f610d448383610ca1565b60208301905092915050565b5f602082019050919050565b5f610d6682610d10565b610d708185610d1a565b9350610d7b83610d2a565b805f5b83811015610dab578151610d928882610d39565b9750610d9d83610d50565b925050600181019050610d7e565b5085935050505092915050565b5f606083015f830151610dcd5f860182610ca1565b5060208301518482036020860152610de58282610cd8565b91505060408301518482036040860152610dff8282610d5c565b9150508091505092915050565b5f6020820190508181035f830152610e248184610db8565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f606083015f830151610e6a5f860182610ca1565b5060208301518482036020860152610e828282610cd8565b91505060408301518482036040860152610e9c8282610d5c565b9150508091505092915050565b5f610eb48383610e55565b905092915050565b5f602082019050919050565b5f610ed282610e2c565b610edc8185610e36565b935083602082028501610eee85610e46565b805f5b85811015610f295784840389528151610f0a8582610ea9565b9450610f1583610ebc565b925060208a01995050600181019050610ef1565b50829750879550505050505092915050565b5f6020820190508181035f830152610f538184610ec8565b905092915050565b610f6481610c98565b8114610f6e575f80fd5b50565b5f81359050610f7f81610f5b565b92915050565b5f67ffffffffffffffff821115610f9f57610f9e610b2d565b5b602082029050602081019050919050565b5f80fd5b5f610fc6610fc184610f85565b610b8b565b90508083825260208201905060208402830185811115610fe957610fe8610fb0565b5b835b818110156110125780610ffe8882610f71565b845260208401935050602081019050610feb565b5050509392505050565b5f82601f8301126110305761102f610b15565b5b8135611040848260208601610fb4565b91505092915050565b5f805f606084860312156110605761105f610b0d565b5b5f61106d86828701610f71565b935050602084013567ffffffffffffffff81111561108e5761108d610b11565b5b61109a86828701610c24565b925050604084013567ffffffffffffffff8111156110bb576110ba610b11565b5b6110c78682870161101c565b9150509250925092565b5f80604083850312156110e7576110e6610b0d565b5b5f83013567ffffffffffffffff81111561110457611103610b11565b5b61111085828601610c24565b925050602083013567ffffffffffffffff81111561113157611130610b11565b5b61113d8582860161101c565b9150509250929050565b5f6020828403121561115c5761115b610b0d565b5b5f61116984828501610f71565b91505092915050565b61117b81610c98565b82525050565b5f6020820190506111945f830184611172565b92915050565b5f82825260208201905092915050565b5f6111b482610cb0565b6111be818561119a565b93506111ce818560208601610cca565b6111d781610b1d565b840191505092915050565b5f6040820190506111f55f830185611172565b818103602083015261120781846111aa565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061128157607f821691505b6020821081036112945761129361123d565b5b50919050565b7f417474726962757465206e6f7420666f756e64000000000000000000000000005f82015250565b5f6112ce60138361119a565b91506112d98261129a565b602082019050919050565b5f6020820190508181035f8301526112fb816112c2565b9050919050565b7f41747472696275746520646f6573206e6f7420657869737400000000000000005f82015250565b5f61133660188361119a565b915061134182611302565b602082019050919050565b5f6020820190508181035f8301526113638161132a565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026113c67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261138b565b6113d0868361138b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61140b61140661140184610c98565b6113e8565b610c98565b9050919050565b5f819050919050565b611424836113f1565b61143861143082611412565b848454611397565b825550505050565b5f90565b61144c611440565b61145781848461141b565b505050565b5b8181101561147a5761146f5f82611444565b60018101905061145d565b5050565b601f8211156114bf576114908161136a565b6114998461137c565b810160208510156114a8578190505b6114bc6114b48561137c565b83018261145c565b50505b505050565b5f82821c905092915050565b5f6114df5f19846008026114c4565b1980831691505092915050565b5f6114f783836114d0565b9150826002028217905092915050565b61151082610cb0565b67ffffffffffffffff81111561152957611528610b2d565b5b611533825461126a565b61153e82828561147e565b5f60209050601f83116001811461156f575f841561155d578287015190505b61156785826114ec565b8655506115ce565b601f19841661157d8661136a565b5f5b828110156115a45784890151825560018201915060208501945060208101905061157f565b868310156115c157848901516115bd601f8916826114d0565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f6115f082610d10565b6115fa81856115d6565b935061160583610d2a565b805f5b8381101561163557815161161c8882610d39565b975061162783610d50565b925050600181019050611608565b5085935050505092915050565b5f6060820190506116555f830186611172565b818103602083015261166781856111aa565b9050818103604083015261167b81846115e6565b9050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6116bc82610c98565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116ee576116ed611685565b5b600182019050919050565b5f81905092915050565b5f61170d82610cb0565b61171781856116f9565b9350611727818560208601610cca565b80840191505092915050565b5f61173e8284611703565b91508190509291505056fea26469706673582212200f6404c9dd962fe7c9b6fff819bf51f627c4eb2dc72e86e4461bbd51794b20db64736f6c634300081a0033",
}

// AttributeABI is the input ABI used to generate the binding from.
// Deprecated: Use AttributeMetaData.ABI instead.
var AttributeABI = AttributeMetaData.ABI

// AttributeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttributeMetaData.Bin instead.
var AttributeBin = AttributeMetaData.Bin

// DeployAttribute deploys a new Ethereum contract, binding an instance of Attribute to it.
func DeployAttribute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Attribute, error) {
	parsed, err := AttributeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttributeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Attribute{AttributeCaller: AttributeCaller{contract: contract}, AttributeTransactor: AttributeTransactor{contract: contract}, AttributeFilterer: AttributeFilterer{contract: contract}}, nil
}

// Attribute is an auto generated Go binding around an Ethereum contract.
type Attribute struct {
	AttributeCaller     // Read-only binding to the contract
	AttributeTransactor // Write-only binding to the contract
	AttributeFilterer   // Log filterer for contract events
}

// AttributeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttributeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttributeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttributeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttributeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttributeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttributeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttributeSession struct {
	Contract     *Attribute        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttributeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttributeCallerSession struct {
	Contract *AttributeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AttributeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttributeTransactorSession struct {
	Contract     *AttributeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AttributeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttributeRaw struct {
	Contract *Attribute // Generic contract binding to access the raw methods on
}

// AttributeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttributeCallerRaw struct {
	Contract *AttributeCaller // Generic read-only contract binding to access the raw methods on
}

// AttributeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttributeTransactorRaw struct {
	Contract *AttributeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttribute creates a new instance of Attribute, bound to a specific deployed contract.
func NewAttribute(address common.Address, backend bind.ContractBackend) (*Attribute, error) {
	contract, err := bindAttribute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attribute{AttributeCaller: AttributeCaller{contract: contract}, AttributeTransactor: AttributeTransactor{contract: contract}, AttributeFilterer: AttributeFilterer{contract: contract}}, nil
}

// NewAttributeCaller creates a new read-only instance of Attribute, bound to a specific deployed contract.
func NewAttributeCaller(address common.Address, caller bind.ContractCaller) (*AttributeCaller, error) {
	contract, err := bindAttribute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttributeCaller{contract: contract}, nil
}

// NewAttributeTransactor creates a new write-only instance of Attribute, bound to a specific deployed contract.
func NewAttributeTransactor(address common.Address, transactor bind.ContractTransactor) (*AttributeTransactor, error) {
	contract, err := bindAttribute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttributeTransactor{contract: contract}, nil
}

// NewAttributeFilterer creates a new log filterer instance of Attribute, bound to a specific deployed contract.
func NewAttributeFilterer(address common.Address, filterer bind.ContractFilterer) (*AttributeFilterer, error) {
	contract, err := bindAttribute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttributeFilterer{contract: contract}, nil
}

// bindAttribute binds a generic wrapper to an already deployed contract.
func bindAttribute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AttributeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attribute *AttributeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attribute.Contract.AttributeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attribute *AttributeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attribute.Contract.AttributeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attribute *AttributeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attribute.Contract.AttributeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attribute *AttributeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attribute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attribute *AttributeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attribute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attribute *AttributeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attribute.Contract.contract.Transact(opts, method, params...)
}

// Attributes is a free data retrieval call binding the contract method 0xd05dcc6a.
//
// Solidity: function attributes(uint256 ) view returns(uint256 id, string name)
func (_Attribute *AttributeCaller) Attributes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id   *big.Int
	Name string
}, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "attributes", arg0)

	outstruct := new(struct {
		Id   *big.Int
		Name string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Attributes is a free data retrieval call binding the contract method 0xd05dcc6a.
//
// Solidity: function attributes(uint256 ) view returns(uint256 id, string name)
func (_Attribute *AttributeSession) Attributes(arg0 *big.Int) (struct {
	Id   *big.Int
	Name string
}, error) {
	return _Attribute.Contract.Attributes(&_Attribute.CallOpts, arg0)
}

// Attributes is a free data retrieval call binding the contract method 0xd05dcc6a.
//
// Solidity: function attributes(uint256 ) view returns(uint256 id, string name)
func (_Attribute *AttributeCallerSession) Attributes(arg0 *big.Int) (struct {
	Id   *big.Int
	Name string
}, error) {
	return _Attribute.Contract.Attributes(&_Attribute.CallOpts, arg0)
}

// GetAttribute is a free data retrieval call binding the contract method 0x493fe80f.
//
// Solidity: function getAttribute(uint256 _id) view returns((uint256,string,uint256[]))
func (_Attribute *AttributeCaller) GetAttribute(opts *bind.CallOpts, _id *big.Int) (AttributeContractAttribute, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "getAttribute", _id)

	if err != nil {
		return *new(AttributeContractAttribute), err
	}

	out0 := *abi.ConvertType(out[0], new(AttributeContractAttribute)).(*AttributeContractAttribute)

	return out0, err

}

// GetAttribute is a free data retrieval call binding the contract method 0x493fe80f.
//
// Solidity: function getAttribute(uint256 _id) view returns((uint256,string,uint256[]))
func (_Attribute *AttributeSession) GetAttribute(_id *big.Int) (AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttribute(&_Attribute.CallOpts, _id)
}

// GetAttribute is a free data retrieval call binding the contract method 0x493fe80f.
//
// Solidity: function getAttribute(uint256 _id) view returns((uint256,string,uint256[]))
func (_Attribute *AttributeCallerSession) GetAttribute(_id *big.Int) (AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttribute(&_Attribute.CallOpts, _id)
}

// GetAttributes is a free data retrieval call binding the contract method 0x152583de.
//
// Solidity: function getAttributes() view returns((uint256,string,uint256[])[])
func (_Attribute *AttributeCaller) GetAttributes(opts *bind.CallOpts) ([]AttributeContractAttribute, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "getAttributes")

	if err != nil {
		return *new([]AttributeContractAttribute), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttributeContractAttribute)).(*[]AttributeContractAttribute)

	return out0, err

}

// GetAttributes is a free data retrieval call binding the contract method 0x152583de.
//
// Solidity: function getAttributes() view returns((uint256,string,uint256[])[])
func (_Attribute *AttributeSession) GetAttributes() ([]AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttributes(&_Attribute.CallOpts)
}

// GetAttributes is a free data retrieval call binding the contract method 0x152583de.
//
// Solidity: function getAttributes() view returns((uint256,string,uint256[])[])
func (_Attribute *AttributeCallerSession) GetAttributes() ([]AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttributes(&_Attribute.CallOpts)
}

// GetAttributesByName is a free data retrieval call binding the contract method 0x06516a4a.
//
// Solidity: function getAttributesByName(string _name) view returns((uint256,string,uint256[]))
func (_Attribute *AttributeCaller) GetAttributesByName(opts *bind.CallOpts, _name string) (AttributeContractAttribute, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "getAttributesByName", _name)

	if err != nil {
		return *new(AttributeContractAttribute), err
	}

	out0 := *abi.ConvertType(out[0], new(AttributeContractAttribute)).(*AttributeContractAttribute)

	return out0, err

}

// GetAttributesByName is a free data retrieval call binding the contract method 0x06516a4a.
//
// Solidity: function getAttributesByName(string _name) view returns((uint256,string,uint256[]))
func (_Attribute *AttributeSession) GetAttributesByName(_name string) (AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttributesByName(&_Attribute.CallOpts, _name)
}

// GetAttributesByName is a free data retrieval call binding the contract method 0x06516a4a.
//
// Solidity: function getAttributesByName(string _name) view returns((uint256,string,uint256[]))
func (_Attribute *AttributeCallerSession) GetAttributesByName(_name string) (AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttributesByName(&_Attribute.CallOpts, _name)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Attribute *AttributeCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Attribute *AttributeSession) NextId() (*big.Int, error) {
	return _Attribute.Contract.NextId(&_Attribute.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Attribute *AttributeCallerSession) NextId() (*big.Int, error) {
	return _Attribute.Contract.NextId(&_Attribute.CallOpts)
}

// CreateAttribute is a paid mutator transaction binding the contract method 0x2822800c.
//
// Solidity: function createAttribute(string _name, uint256[] _value) returns()
func (_Attribute *AttributeTransactor) CreateAttribute(opts *bind.TransactOpts, _name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.contract.Transact(opts, "createAttribute", _name, _value)
}

// CreateAttribute is a paid mutator transaction binding the contract method 0x2822800c.
//
// Solidity: function createAttribute(string _name, uint256[] _value) returns()
func (_Attribute *AttributeSession) CreateAttribute(_name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.CreateAttribute(&_Attribute.TransactOpts, _name, _value)
}

// CreateAttribute is a paid mutator transaction binding the contract method 0x2822800c.
//
// Solidity: function createAttribute(string _name, uint256[] _value) returns()
func (_Attribute *AttributeTransactorSession) CreateAttribute(_name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.CreateAttribute(&_Attribute.TransactOpts, _name, _value)
}

// DeleteAttribute is a paid mutator transaction binding the contract method 0x7238b41f.
//
// Solidity: function deleteAttribute(uint256 _id) returns()
func (_Attribute *AttributeTransactor) DeleteAttribute(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Attribute.contract.Transact(opts, "deleteAttribute", _id)
}

// DeleteAttribute is a paid mutator transaction binding the contract method 0x7238b41f.
//
// Solidity: function deleteAttribute(uint256 _id) returns()
func (_Attribute *AttributeSession) DeleteAttribute(_id *big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.DeleteAttribute(&_Attribute.TransactOpts, _id)
}

// DeleteAttribute is a paid mutator transaction binding the contract method 0x7238b41f.
//
// Solidity: function deleteAttribute(uint256 _id) returns()
func (_Attribute *AttributeTransactorSession) DeleteAttribute(_id *big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.DeleteAttribute(&_Attribute.TransactOpts, _id)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x1736a62b.
//
// Solidity: function updateAttribute(uint256 _id, string _name, uint256[] _value) returns()
func (_Attribute *AttributeTransactor) UpdateAttribute(opts *bind.TransactOpts, _id *big.Int, _name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.contract.Transact(opts, "updateAttribute", _id, _name, _value)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x1736a62b.
//
// Solidity: function updateAttribute(uint256 _id, string _name, uint256[] _value) returns()
func (_Attribute *AttributeSession) UpdateAttribute(_id *big.Int, _name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.UpdateAttribute(&_Attribute.TransactOpts, _id, _name, _value)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x1736a62b.
//
// Solidity: function updateAttribute(uint256 _id, string _name, uint256[] _value) returns()
func (_Attribute *AttributeTransactorSession) UpdateAttribute(_id *big.Int, _name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.UpdateAttribute(&_Attribute.TransactOpts, _id, _name, _value)
}

// AttributeAttributeCreatedIterator is returned from FilterAttributeCreated and is used to iterate over the raw logs and unpacked data for AttributeCreated events raised by the Attribute contract.
type AttributeAttributeCreatedIterator struct {
	Event *AttributeAttributeCreated // Event containing the contract specifics and raw log

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
func (it *AttributeAttributeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttributeAttributeCreated)
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
		it.Event = new(AttributeAttributeCreated)
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
func (it *AttributeAttributeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttributeAttributeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttributeAttributeCreated represents a AttributeCreated event raised by the Attribute contract.
type AttributeAttributeCreated struct {
	Id    *big.Int
	Name  string
	Value []*big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAttributeCreated is a free log retrieval operation binding the contract event 0x92ff6eabb2f26553137551f9eaac5f9e399114f11d71898144a4e7e71f8689f1.
//
// Solidity: event AttributeCreated(uint256 id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) FilterAttributeCreated(opts *bind.FilterOpts) (*AttributeAttributeCreatedIterator, error) {

	logs, sub, err := _Attribute.contract.FilterLogs(opts, "AttributeCreated")
	if err != nil {
		return nil, err
	}
	return &AttributeAttributeCreatedIterator{contract: _Attribute.contract, event: "AttributeCreated", logs: logs, sub: sub}, nil
}

// WatchAttributeCreated is a free log subscription operation binding the contract event 0x92ff6eabb2f26553137551f9eaac5f9e399114f11d71898144a4e7e71f8689f1.
//
// Solidity: event AttributeCreated(uint256 id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) WatchAttributeCreated(opts *bind.WatchOpts, sink chan<- *AttributeAttributeCreated) (event.Subscription, error) {

	logs, sub, err := _Attribute.contract.WatchLogs(opts, "AttributeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttributeAttributeCreated)
				if err := _Attribute.contract.UnpackLog(event, "AttributeCreated", log); err != nil {
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

// ParseAttributeCreated is a log parse operation binding the contract event 0x92ff6eabb2f26553137551f9eaac5f9e399114f11d71898144a4e7e71f8689f1.
//
// Solidity: event AttributeCreated(uint256 id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) ParseAttributeCreated(log types.Log) (*AttributeAttributeCreated, error) {
	event := new(AttributeAttributeCreated)
	if err := _Attribute.contract.UnpackLog(event, "AttributeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttributeAttributeDeletedIterator is returned from FilterAttributeDeleted and is used to iterate over the raw logs and unpacked data for AttributeDeleted events raised by the Attribute contract.
type AttributeAttributeDeletedIterator struct {
	Event *AttributeAttributeDeleted // Event containing the contract specifics and raw log

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
func (it *AttributeAttributeDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttributeAttributeDeleted)
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
		it.Event = new(AttributeAttributeDeleted)
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
func (it *AttributeAttributeDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttributeAttributeDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttributeAttributeDeleted represents a AttributeDeleted event raised by the Attribute contract.
type AttributeAttributeDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAttributeDeleted is a free log retrieval operation binding the contract event 0xfe96137d7877697f141204379bad0e6617ff7023d7608f9db3be956d7ceef711.
//
// Solidity: event AttributeDeleted(uint256 id)
func (_Attribute *AttributeFilterer) FilterAttributeDeleted(opts *bind.FilterOpts) (*AttributeAttributeDeletedIterator, error) {

	logs, sub, err := _Attribute.contract.FilterLogs(opts, "AttributeDeleted")
	if err != nil {
		return nil, err
	}
	return &AttributeAttributeDeletedIterator{contract: _Attribute.contract, event: "AttributeDeleted", logs: logs, sub: sub}, nil
}

// WatchAttributeDeleted is a free log subscription operation binding the contract event 0xfe96137d7877697f141204379bad0e6617ff7023d7608f9db3be956d7ceef711.
//
// Solidity: event AttributeDeleted(uint256 id)
func (_Attribute *AttributeFilterer) WatchAttributeDeleted(opts *bind.WatchOpts, sink chan<- *AttributeAttributeDeleted) (event.Subscription, error) {

	logs, sub, err := _Attribute.contract.WatchLogs(opts, "AttributeDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttributeAttributeDeleted)
				if err := _Attribute.contract.UnpackLog(event, "AttributeDeleted", log); err != nil {
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

// ParseAttributeDeleted is a log parse operation binding the contract event 0xfe96137d7877697f141204379bad0e6617ff7023d7608f9db3be956d7ceef711.
//
// Solidity: event AttributeDeleted(uint256 id)
func (_Attribute *AttributeFilterer) ParseAttributeDeleted(log types.Log) (*AttributeAttributeDeleted, error) {
	event := new(AttributeAttributeDeleted)
	if err := _Attribute.contract.UnpackLog(event, "AttributeDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttributeAttributeUpdatedIterator is returned from FilterAttributeUpdated and is used to iterate over the raw logs and unpacked data for AttributeUpdated events raised by the Attribute contract.
type AttributeAttributeUpdatedIterator struct {
	Event *AttributeAttributeUpdated // Event containing the contract specifics and raw log

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
func (it *AttributeAttributeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttributeAttributeUpdated)
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
		it.Event = new(AttributeAttributeUpdated)
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
func (it *AttributeAttributeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttributeAttributeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttributeAttributeUpdated represents a AttributeUpdated event raised by the Attribute contract.
type AttributeAttributeUpdated struct {
	Id    *big.Int
	Name  string
	Value []*big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAttributeUpdated is a free log retrieval operation binding the contract event 0x8f26e9a323213025e76e7c58dbf6338a9c36708b0f933b7e5babc20acca4e8d1.
//
// Solidity: event AttributeUpdated(uint256 id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) FilterAttributeUpdated(opts *bind.FilterOpts) (*AttributeAttributeUpdatedIterator, error) {

	logs, sub, err := _Attribute.contract.FilterLogs(opts, "AttributeUpdated")
	if err != nil {
		return nil, err
	}
	return &AttributeAttributeUpdatedIterator{contract: _Attribute.contract, event: "AttributeUpdated", logs: logs, sub: sub}, nil
}

// WatchAttributeUpdated is a free log subscription operation binding the contract event 0x8f26e9a323213025e76e7c58dbf6338a9c36708b0f933b7e5babc20acca4e8d1.
//
// Solidity: event AttributeUpdated(uint256 id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) WatchAttributeUpdated(opts *bind.WatchOpts, sink chan<- *AttributeAttributeUpdated) (event.Subscription, error) {

	logs, sub, err := _Attribute.contract.WatchLogs(opts, "AttributeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttributeAttributeUpdated)
				if err := _Attribute.contract.UnpackLog(event, "AttributeUpdated", log); err != nil {
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

// ParseAttributeUpdated is a log parse operation binding the contract event 0x8f26e9a323213025e76e7c58dbf6338a9c36708b0f933b7e5babc20acca4e8d1.
//
// Solidity: event AttributeUpdated(uint256 id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) ParseAttributeUpdated(log types.Log) (*AttributeAttributeUpdated, error) {
	event := new(AttributeAttributeUpdated)
	if err := _Attribute.contract.UnpackLog(event, "AttributeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
