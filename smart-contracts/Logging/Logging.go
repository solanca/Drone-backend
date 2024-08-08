// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Logging

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

// LoggingContractLogEntry is an auto generated low-level Go binding around an user-defined struct.
type LoggingContractLogEntry struct {
	User      common.Address
	Action    string
	Timestamp string
	Data      string
}

// LoggingMetaData contains all meta data concerning the Logging contract.
var LoggingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"ActionLogged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structLoggingContract.LogEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"logAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"logs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506114128061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806397f8abd814610043578063e79899bd1461005f578063fe290ccb14610092575b5f80fd5b61005d60048036038101906100589190610a14565b6100b0565b005b61007960048036038101906100749190610abd565b61020b565b6040516100899493929190610b87565b60405180910390f35b61009a6103f6565b6040516100a79190610d5f565b60405180910390f35b5f612a30426100bf9190610dac565b90505f6100cb82610655565b90505f60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200186815260200183815260200185815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816101849190610fd9565b50604082015181600201908161019a9190610fd9565b5060608201518160030190816101b09190610fd9565b5050503373ffffffffffffffffffffffffffffffffffffffff167f48c86d04f2513c301a1df6bf62da971314800bb3d672395fb3fc2abc063d4e6c8583866040516101fd939291906110a8565b60405180910390a250505050565b5f8181548110610219575f80fd5b905f5260205f2090600402015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461025d90610e0c565b80601f016020809104026020016040519081016040528092919081815260200182805461028990610e0c565b80156102d45780601f106102ab576101008083540402835291602001916102d4565b820191905f5260205f20905b8154815290600101906020018083116102b757829003601f168201915b5050505050908060020180546102e990610e0c565b80601f016020809104026020016040519081016040528092919081815260200182805461031590610e0c565b80156103605780601f1061033757610100808354040283529160200191610360565b820191905f5260205f20905b81548152906001019060200180831161034357829003601f168201915b50505050509080600301805461037590610e0c565b80601f01602080910402602001604051908101604052809291908181526020018280546103a190610e0c565b80156103ec5780601f106103c3576101008083540402835291602001916103ec565b820191905f5260205f20905b8154815290600101906020018083116103cf57829003601f168201915b5050505050905084565b60605f805480602002602001604051908101604052809291908181526020015f905b8282101561064c578382905f5260205f2090600402016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461049d90610e0c565b80601f01602080910402602001604051908101604052809291908181526020018280546104c990610e0c565b80156105145780601f106104eb57610100808354040283529160200191610514565b820191905f5260205f20905b8154815290600101906020018083116104f757829003601f168201915b5050505050815260200160028201805461052d90610e0c565b80601f016020809104026020016040519081016040528092919081815260200182805461055990610e0c565b80156105a45780601f1061057b576101008083540402835291602001916105a4565b820191905f5260205f20905b81548152906001019060200180831161058757829003601f168201915b505050505081526020016003820180546105bd90610e0c565b80601f01602080910402602001604051908101604052809291908181526020018280546105e990610e0c565b80156106345780601f1061060b57610100808354040283529160200191610634565b820191905f5260205f20905b81548152906001019060200180831161061757829003601f168201915b50505050508152505081526020019060010190610418565b50505050905090565b60605f6018603c8085610668919061111f565b610672919061111f565b61067c919061114f565b90505f603c808561068d919061111f565b610697919061114f565b90505f603c856106a7919061114f565b90505f6106b3846106fe565b90505f6106bf846106fe565b90505f6106cb846106fe565b90508282826040516020016106e293929190611203565b6040516020818303038152906040529650505050505050919050565b6060600a8210156107385761071282610749565b6040516020016107229190611293565b6040516020818303038152906040529050610744565b61074182610749565b90505b919050565b60605f820361078f576040518060400160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525090506108c2565b5f8290505f5b5f82146107be5780806107a7906112b4565b915050600a826107b7919061111f565b9150610795565b5f8167ffffffffffffffff8111156107d9576107d86108f0565b5b6040519080825280601f01601f19166020018201604052801561080b5781602001600182028036833780820191505090505b5090505f8290505b5f86146108ba5760018161082791906112fb565b90505f600a8088610838919061111f565b610842919061132e565b8761084d91906112fb565b6030610859919061137b565b90505f8160f81b905080848481518110610876576108756113af565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a886108b1919061111f565b97505050610813565b819450505050505b919050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610926826108e0565b810181811067ffffffffffffffff82111715610945576109446108f0565b5b80604052505050565b5f6109576108c7565b9050610963828261091d565b919050565b5f67ffffffffffffffff821115610982576109816108f0565b5b61098b826108e0565b9050602081019050919050565b828183375f83830152505050565b5f6109b86109b384610968565b61094e565b9050828152602081018484840111156109d4576109d36108dc565b5b6109df848285610998565b509392505050565b5f82601f8301126109fb576109fa6108d8565b5b8135610a0b8482602086016109a6565b91505092915050565b5f8060408385031215610a2a57610a296108d0565b5b5f83013567ffffffffffffffff811115610a4757610a466108d4565b5b610a53858286016109e7565b925050602083013567ffffffffffffffff811115610a7457610a736108d4565b5b610a80858286016109e7565b9150509250929050565b5f819050919050565b610a9c81610a8a565b8114610aa6575f80fd5b50565b5f81359050610ab781610a93565b92915050565b5f60208284031215610ad257610ad16108d0565b5b5f610adf84828501610aa9565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b1182610ae8565b9050919050565b610b2181610b07565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610b5982610b27565b610b638185610b31565b9350610b73818560208601610b41565b610b7c816108e0565b840191505092915050565b5f608082019050610b9a5f830187610b18565b8181036020830152610bac8186610b4f565b90508181036040830152610bc08185610b4f565b90508181036060830152610bd48184610b4f565b905095945050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610c1181610b07565b82525050565b5f82825260208201905092915050565b5f610c3182610b27565b610c3b8185610c17565b9350610c4b818560208601610b41565b610c54816108e0565b840191505092915050565b5f608083015f830151610c745f860182610c08565b5060208301518482036020860152610c8c8282610c27565b91505060408301518482036040860152610ca68282610c27565b91505060608301518482036060860152610cc08282610c27565b9150508091505092915050565b5f610cd88383610c5f565b905092915050565b5f602082019050919050565b5f610cf682610bdf565b610d008185610be9565b935083602082028501610d1285610bf9565b805f5b85811015610d4d5784840389528151610d2e8582610ccd565b9450610d3983610ce0565b925060208a01995050600181019050610d15565b50829750879550505050505092915050565b5f6020820190508181035f830152610d778184610cec565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610db682610a8a565b9150610dc183610a8a565b9250828201905080821115610dd957610dd8610d7f565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610e2357607f821691505b602082108103610e3657610e35610ddf565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610e987fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610e5d565b610ea28683610e5d565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610edd610ed8610ed384610a8a565b610eba565b610a8a565b9050919050565b5f819050919050565b610ef683610ec3565b610f0a610f0282610ee4565b848454610e69565b825550505050565b5f90565b610f1e610f12565b610f29818484610eed565b505050565b5b81811015610f4c57610f415f82610f16565b600181019050610f2f565b5050565b601f821115610f9157610f6281610e3c565b610f6b84610e4e565b81016020851015610f7a578190505b610f8e610f8685610e4e565b830182610f2e565b50505b505050565b5f82821c905092915050565b5f610fb15f1984600802610f96565b1980831691505092915050565b5f610fc98383610fa2565b9150826002028217905092915050565b610fe282610b27565b67ffffffffffffffff811115610ffb57610ffa6108f0565b5b6110058254610e0c565b611010828285610f50565b5f60209050601f831160018114611041575f841561102f578287015190505b6110398582610fbe565b8655506110a0565b601f19841661104f86610e3c565b5f5b8281101561107657848901518255600182019150602085019450602081019050611051565b86831015611093578489015161108f601f891682610fa2565b8355505b6001600288020188555050505b505050505050565b5f6060820190508181035f8301526110c08186610b4f565b905081810360208301526110d48185610b4f565b905081810360408301526110e88184610b4f565b9050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61112982610a8a565b915061113483610a8a565b925082611144576111436110f2565b5b828204905092915050565b5f61115982610a8a565b915061116483610a8a565b925082611174576111736110f2565b5b828206905092915050565b5f81905092915050565b5f61119382610b27565b61119d818561117f565b93506111ad818560208601610b41565b80840191505092915050565b7f3a000000000000000000000000000000000000000000000000000000000000005f82015250565b5f6111ed60018361117f565b91506111f8826111b9565b600182019050919050565b5f61120e8286611189565b9150611219826111e1565b91506112258285611189565b9150611230826111e1565b915061123c8284611189565b9150819050949350505050565b7f30000000000000000000000000000000000000000000000000000000000000005f82015250565b5f61127d60018361117f565b915061128882611249565b600182019050919050565b5f61129d82611271565b91506112a98284611189565b915081905092915050565b5f6112be82610a8a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036112f0576112ef610d7f565b5b600182019050919050565b5f61130582610a8a565b915061131083610a8a565b925082820390508181111561132857611327610d7f565b5b92915050565b5f61133882610a8a565b915061134383610a8a565b925082820261135181610a8a565b9150828204841483151761136857611367610d7f565b5b5092915050565b5f60ff82169050919050565b5f6113858261136f565b91506113908361136f565b9250828201905060ff8111156113a9576113a8610d7f565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea26469706673582212202159001b028b38b23fd02cce42731e715743076455e601ad067ebeac4f3268e564736f6c634300081a0033",
}

// LoggingABI is the input ABI used to generate the binding from.
// Deprecated: Use LoggingMetaData.ABI instead.
var LoggingABI = LoggingMetaData.ABI

// LoggingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LoggingMetaData.Bin instead.
var LoggingBin = LoggingMetaData.Bin

// DeployLogging deploys a new Ethereum contract, binding an instance of Logging to it.
func DeployLogging(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Logging, error) {
	parsed, err := LoggingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LoggingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Logging{LoggingCaller: LoggingCaller{contract: contract}, LoggingTransactor: LoggingTransactor{contract: contract}, LoggingFilterer: LoggingFilterer{contract: contract}}, nil
}

// Logging is an auto generated Go binding around an Ethereum contract.
type Logging struct {
	LoggingCaller     // Read-only binding to the contract
	LoggingTransactor // Write-only binding to the contract
	LoggingFilterer   // Log filterer for contract events
}

// LoggingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoggingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoggingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoggingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoggingSession struct {
	Contract     *Logging          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoggingCallerSession struct {
	Contract *LoggingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LoggingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoggingTransactorSession struct {
	Contract     *LoggingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LoggingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoggingRaw struct {
	Contract *Logging // Generic contract binding to access the raw methods on
}

// LoggingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoggingCallerRaw struct {
	Contract *LoggingCaller // Generic read-only contract binding to access the raw methods on
}

// LoggingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoggingTransactorRaw struct {
	Contract *LoggingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogging creates a new instance of Logging, bound to a specific deployed contract.
func NewLogging(address common.Address, backend bind.ContractBackend) (*Logging, error) {
	contract, err := bindLogging(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Logging{LoggingCaller: LoggingCaller{contract: contract}, LoggingTransactor: LoggingTransactor{contract: contract}, LoggingFilterer: LoggingFilterer{contract: contract}}, nil
}

// NewLoggingCaller creates a new read-only instance of Logging, bound to a specific deployed contract.
func NewLoggingCaller(address common.Address, caller bind.ContractCaller) (*LoggingCaller, error) {
	contract, err := bindLogging(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoggingCaller{contract: contract}, nil
}

// NewLoggingTransactor creates a new write-only instance of Logging, bound to a specific deployed contract.
func NewLoggingTransactor(address common.Address, transactor bind.ContractTransactor) (*LoggingTransactor, error) {
	contract, err := bindLogging(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoggingTransactor{contract: contract}, nil
}

// NewLoggingFilterer creates a new log filterer instance of Logging, bound to a specific deployed contract.
func NewLoggingFilterer(address common.Address, filterer bind.ContractFilterer) (*LoggingFilterer, error) {
	contract, err := bindLogging(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoggingFilterer{contract: contract}, nil
}

// bindLogging binds a generic wrapper to an already deployed contract.
func bindLogging(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoggingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logging *LoggingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logging.Contract.LoggingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logging *LoggingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logging.Contract.LoggingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logging *LoggingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logging.Contract.LoggingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logging *LoggingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logging.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logging *LoggingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logging.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logging *LoggingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logging.Contract.contract.Transact(opts, method, params...)
}

// GetLogs is a free data retrieval call binding the contract method 0xfe290ccb.
//
// Solidity: function getLogs() view returns((address,string,string,string)[])
func (_Logging *LoggingCaller) GetLogs(opts *bind.CallOpts) ([]LoggingContractLogEntry, error) {
	var out []interface{}
	err := _Logging.contract.Call(opts, &out, "getLogs")

	if err != nil {
		return *new([]LoggingContractLogEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]LoggingContractLogEntry)).(*[]LoggingContractLogEntry)

	return out0, err

}

// GetLogs is a free data retrieval call binding the contract method 0xfe290ccb.
//
// Solidity: function getLogs() view returns((address,string,string,string)[])
func (_Logging *LoggingSession) GetLogs() ([]LoggingContractLogEntry, error) {
	return _Logging.Contract.GetLogs(&_Logging.CallOpts)
}

// GetLogs is a free data retrieval call binding the contract method 0xfe290ccb.
//
// Solidity: function getLogs() view returns((address,string,string,string)[])
func (_Logging *LoggingCallerSession) GetLogs() ([]LoggingContractLogEntry, error) {
	return _Logging.Contract.GetLogs(&_Logging.CallOpts)
}

// Logs is a free data retrieval call binding the contract method 0xe79899bd.
//
// Solidity: function logs(uint256 ) view returns(address user, string action, string timestamp, string data)
func (_Logging *LoggingCaller) Logs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User      common.Address
	Action    string
	Timestamp string
	Data      string
}, error) {
	var out []interface{}
	err := _Logging.contract.Call(opts, &out, "logs", arg0)

	outstruct := new(struct {
		User      common.Address
		Action    string
		Timestamp string
		Data      string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Action = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Data = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Logs is a free data retrieval call binding the contract method 0xe79899bd.
//
// Solidity: function logs(uint256 ) view returns(address user, string action, string timestamp, string data)
func (_Logging *LoggingSession) Logs(arg0 *big.Int) (struct {
	User      common.Address
	Action    string
	Timestamp string
	Data      string
}, error) {
	return _Logging.Contract.Logs(&_Logging.CallOpts, arg0)
}

// Logs is a free data retrieval call binding the contract method 0xe79899bd.
//
// Solidity: function logs(uint256 ) view returns(address user, string action, string timestamp, string data)
func (_Logging *LoggingCallerSession) Logs(arg0 *big.Int) (struct {
	User      common.Address
	Action    string
	Timestamp string
	Data      string
}, error) {
	return _Logging.Contract.Logs(&_Logging.CallOpts, arg0)
}

// LogAction is a paid mutator transaction binding the contract method 0x97f8abd8.
//
// Solidity: function logAction(string action, string data) returns()
func (_Logging *LoggingTransactor) LogAction(opts *bind.TransactOpts, action string, data string) (*types.Transaction, error) {
	return _Logging.contract.Transact(opts, "logAction", action, data)
}

// LogAction is a paid mutator transaction binding the contract method 0x97f8abd8.
//
// Solidity: function logAction(string action, string data) returns()
func (_Logging *LoggingSession) LogAction(action string, data string) (*types.Transaction, error) {
	return _Logging.Contract.LogAction(&_Logging.TransactOpts, action, data)
}

// LogAction is a paid mutator transaction binding the contract method 0x97f8abd8.
//
// Solidity: function logAction(string action, string data) returns()
func (_Logging *LoggingTransactorSession) LogAction(action string, data string) (*types.Transaction, error) {
	return _Logging.Contract.LogAction(&_Logging.TransactOpts, action, data)
}

// LoggingActionLoggedIterator is returned from FilterActionLogged and is used to iterate over the raw logs and unpacked data for ActionLogged events raised by the Logging contract.
type LoggingActionLoggedIterator struct {
	Event *LoggingActionLogged // Event containing the contract specifics and raw log

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
func (it *LoggingActionLoggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggingActionLogged)
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
		it.Event = new(LoggingActionLogged)
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
func (it *LoggingActionLoggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggingActionLoggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggingActionLogged represents a ActionLogged event raised by the Logging contract.
type LoggingActionLogged struct {
	User      common.Address
	Action    string
	Timestamp string
	Data      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterActionLogged is a free log retrieval operation binding the contract event 0x48c86d04f2513c301a1df6bf62da971314800bb3d672395fb3fc2abc063d4e6c.
//
// Solidity: event ActionLogged(address indexed user, string action, string timestamp, string data)
func (_Logging *LoggingFilterer) FilterActionLogged(opts *bind.FilterOpts, user []common.Address) (*LoggingActionLoggedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Logging.contract.FilterLogs(opts, "ActionLogged", userRule)
	if err != nil {
		return nil, err
	}
	return &LoggingActionLoggedIterator{contract: _Logging.contract, event: "ActionLogged", logs: logs, sub: sub}, nil
}

// WatchActionLogged is a free log subscription operation binding the contract event 0x48c86d04f2513c301a1df6bf62da971314800bb3d672395fb3fc2abc063d4e6c.
//
// Solidity: event ActionLogged(address indexed user, string action, string timestamp, string data)
func (_Logging *LoggingFilterer) WatchActionLogged(opts *bind.WatchOpts, sink chan<- *LoggingActionLogged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Logging.contract.WatchLogs(opts, "ActionLogged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggingActionLogged)
				if err := _Logging.contract.UnpackLog(event, "ActionLogged", log); err != nil {
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

// ParseActionLogged is a log parse operation binding the contract event 0x48c86d04f2513c301a1df6bf62da971314800bb3d672395fb3fc2abc063d4e6c.
//
// Solidity: event ActionLogged(address indexed user, string action, string timestamp, string data)
func (_Logging *LoggingFilterer) ParseActionLogged(log types.Log) (*LoggingActionLogged, error) {
	event := new(LoggingActionLogged)
	if err := _Logging.contract.UnpackLog(event, "ActionLogged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
