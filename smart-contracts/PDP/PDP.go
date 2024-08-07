// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PDP

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

// PDPMetaData contains all meta data concerning the PDP contract.
var PDPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_entityId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"}],\"name\":\"evaluateAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_loggerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"logger\",\"outputs\":[{\"internalType\":\"contractLoggingContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516114e23803806114e2833981810160405281019061003191906100d4565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f80fd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b6113d68061010c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063a9a11d1514610038578063f24ccbfe14610068575b5f80fd5b610052600480360381019061004d9190610bb8565b610086565b60405161005f9190610c6e565b60405180910390f35b610070610283565b60405161007d9190610d01565b60405180910390f35b5f805f90505f4290505f620151808261009f9190610d47565b90505f6100ab836102a6565b90506100b882888861034f565b156100c257600193505b5f84610103576040518060400160405280600d81526020017f4163636573732044656e6965640000000000000000000000000000000000000081525061013a565b6040518060400160405280600e81526020017f416363657373204772616e7465640000000000000000000000000000000000008152505b90505f6101468b6103bd565b61014f8b61053b565b8488610190576040518060400160405280600681526020017f44656e69656400000000000000000000000000000000000000000000000000008152506101c7565b6040518060400160405280600a81526020017f50726f67726573736564000000000000000000000000000000000000000000008152505b6040516020016101da9493929190610ef1565b60405160208183030381529060405290505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166397f8abd883836040518363ffffffff1660e01b8152600401610245929190610fa2565b5f604051808303815f87803b15801561025c575f80fd5b505af115801561026e573d5f803e3d5ffd5b50505050859650505050505050949350505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60605f6018603c80856102b99190611004565b6102c39190611004565b6102cd9190610d47565b90505f603c80856102de9190611004565b6102e89190610d47565b90505f603c856102f89190610d47565b90505f6103048461076b565b90505f6103108461076b565b90505f61031c8461076b565b90508282826040516020016103339392919061107e565b6040516020818303038152906040529650505050505050919050565b5f8061035a846107b6565b90505f610366846107b6565b90508082116103885781861015801561037f5750808611155b925050506103b6565b818611801561039a5750620151808611155b806103b157505f86101580156103b05750808611155b5b925050505b9392505050565b60605f8203610403576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610536565b5f8290505f5b5f821461043257808061041b906110c4565b915050600a8261042b9190611004565b9150610409565b5f8167ffffffffffffffff81111561044d5761044c610a94565b5b6040519080825280601f01601f19166020018201604052801561047f5781602001600182028036833780820191505090505b5090505f8290505b5f861461052e5760018161049b919061110b565b90505f600a80886104ac9190611004565b6104b6919061113e565b876104c1919061110b565b60306104cd919061118b565b90505f8160f81b9050808484815181106104ea576104e96111bf565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a886105259190611004565b97505050610487565b819450505050505b919050565b60605f8203610581576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610766565b5f80831290505f81610593578361059e565b8361059d906111ec565b5b90505f5b5f82146105cb5780806105b4906110c4565b915050600a826105c49190611004565b91506105a2565b82156105e05780806105dc906110c4565b9150505b5f8167ffffffffffffffff8111156105fb576105fa610a94565b5b6040519080825280601f01601f19166020018201604052801561062d5781602001600182028036833780820191505090505b5090505f8290508461063f578661064a565b86610649906111ec565b5b93505b5f84146106f457600181610661919061110b565b90505f600a80866106729190611004565b61067c919061113e565b85610687919061110b565b6030610693919061118b565b90505f8160f81b9050808484815181106106b0576106af6111bf565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a866106eb9190611004565b9550505061064d565b841561075d577f2d00000000000000000000000000000000000000000000000000000000000000825f8151811061072e5761072d6111bf565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505b81955050505050505b919050565b6060600a8210156107a55761077f826103bd565b60405160200161078f919061127c565b60405160208183030381529060405290506107b1565b6107ae826103bd565b90505b919050565b5f8082905060088151146107ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f6906112e7565b60405180910390fd5b5f603082600181518110610816576108156111bf565b5b602001015160f81c60f81b60f81c60ff16610831919061110b565b600a6030845f81518110610848576108476111bf565b5b602001015160f81c60f81b60f81c60ff16610863919061110b565b61086d919061113e565b6108779190611305565b90505f6030836004815181106108905761088f6111bf565b5b602001015160f81c60f81b60f81c60ff166108ab919061110b565b600a6030856003815181106108c3576108c26111bf565b5b602001015160f81c60f81b60f81c60ff166108de919061110b565b6108e8919061113e565b6108f29190611305565b90505f60308460078151811061090b5761090a6111bf565b5b602001015160f81c60f81b60f81c60ff16610926919061110b565b600a60308660068151811061093e5761093d6111bf565b5b602001015160f81c60f81b60f81c60ff16610959919061110b565b610963919061113e565b61096d9190611305565b905060188310801561097f5750603c82105b801561098b5750603c81105b6109ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c190611382565b60405180910390fd5b80603c836109d8919061113e565b610e10856109e6919061113e565b6109f09190611305565b6109fa9190611305565b945050505050919050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610a2881610a16565b8114610a32575f80fd5b50565b5f81359050610a4381610a1f565b92915050565b5f819050919050565b610a5b81610a49565b8114610a65575f80fd5b50565b5f81359050610a7681610a52565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610aca82610a84565b810181811067ffffffffffffffff82111715610ae957610ae8610a94565b5b80604052505050565b5f610afb610a05565b9050610b078282610ac1565b919050565b5f67ffffffffffffffff821115610b2657610b25610a94565b5b610b2f82610a84565b9050602081019050919050565b828183375f83830152505050565b5f610b5c610b5784610b0c565b610af2565b905082815260208101848484011115610b7857610b77610a80565b5b610b83848285610b3c565b509392505050565b5f82601f830112610b9f57610b9e610a7c565b5b8135610baf848260208601610b4a565b91505092915050565b5f805f8060808587031215610bd057610bcf610a0e565b5b5f610bdd87828801610a35565b9450506020610bee87828801610a68565b935050604085013567ffffffffffffffff811115610c0f57610c0e610a12565b5b610c1b87828801610b8b565b925050606085013567ffffffffffffffff811115610c3c57610c3b610a12565b5b610c4887828801610b8b565b91505092959194509250565b5f8115159050919050565b610c6881610c54565b82525050565b5f602082019050610c815f830184610c5f565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610cc9610cc4610cbf84610c87565b610ca6565b610c87565b9050919050565b5f610cda82610caf565b9050919050565b5f610ceb82610cd0565b9050919050565b610cfb81610ce1565b82525050565b5f602082019050610d145f830184610cf2565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610d5182610a16565b9150610d5c83610a16565b925082610d6c57610d6b610d1a565b5b828206905092915050565b5f81905092915050565b7f44726f6e652049443a20000000000000000000000000000000000000000000005f82015250565b5f610db5600a83610d77565b9150610dc082610d81565b600a82019050919050565b5f81519050919050565b8281835e5f83830152505050565b5f610ded82610dcb565b610df78185610d77565b9350610e07818560208601610dd5565b80840191505092915050565b7f2c205a6f6e653a200000000000000000000000000000000000000000000000005f82015250565b5f610e47600883610d77565b9150610e5282610e13565b600882019050919050565b7f2c2043757272656e7454696d653a2000000000000000000000000000000000005f82015250565b5f610e91600f83610d77565b9150610e9c82610e5d565b600f82019050919050565b7f2c204772616e7465643a200000000000000000000000000000000000000000005f82015250565b5f610edb600b83610d77565b9150610ee682610ea7565b600b82019050919050565b5f610efb82610da9565b9150610f078287610de3565b9150610f1282610e3b565b9150610f1e8286610de3565b9150610f2982610e85565b9150610f358285610de3565b9150610f4082610ecf565b9150610f4c8284610de3565b915081905095945050505050565b5f82825260208201905092915050565b5f610f7482610dcb565b610f7e8185610f5a565b9350610f8e818560208601610dd5565b610f9781610a84565b840191505092915050565b5f6040820190508181035f830152610fba8185610f6a565b90508181036020830152610fce8184610f6a565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61100e82610a16565b915061101983610a16565b92508261102957611028610d1a565b5b828204905092915050565b7f3a000000000000000000000000000000000000000000000000000000000000005f82015250565b5f611068600183610d77565b915061107382611034565b600182019050919050565b5f6110898286610de3565b91506110948261105c565b91506110a08285610de3565b91506110ab8261105c565b91506110b78284610de3565b9150819050949350505050565b5f6110ce82610a16565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611100576110ff610fd7565b5b600182019050919050565b5f61111582610a16565b915061112083610a16565b925082820390508181111561113857611137610fd7565b5b92915050565b5f61114882610a16565b915061115383610a16565b925082820261116181610a16565b9150828204841483151761117857611177610fd7565b5b5092915050565b5f60ff82169050919050565b5f6111958261117f565b91506111a08361117f565b9250828201905060ff8111156111b9576111b8610fd7565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6111f682610a49565b91507f8000000000000000000000000000000000000000000000000000000000000000820361122857611227610fd7565b5b815f039050919050565b7f30000000000000000000000000000000000000000000000000000000000000005f82015250565b5f611266600183610d77565b915061127182611232565b600182019050919050565b5f6112868261125a565b91506112928284610de3565b915081905092915050565b7f696e76616c69642074696d6520666f726d6174000000000000000000000000005f82015250565b5f6112d1601383610f5a565b91506112dc8261129d565b602082019050919050565b5f6020820190508181035f8301526112fe816112c5565b9050919050565b5f61130f82610a16565b915061131a83610a16565b925082820190508082111561133257611331610fd7565b5b92915050565b7f496e76616c69642074696d652076616c756500000000000000000000000000005f82015250565b5f61136c601283610f5a565b915061137782611338565b602082019050919050565b5f6020820190508181035f83015261139981611360565b905091905056fea2646970667358221220c9d1f4ee866df2861dbe60c70cf420b2e2c7439ea5d63dafb5699e1d8607479464736f6c634300081a0033",
}

// PDPABI is the input ABI used to generate the binding from.
// Deprecated: Use PDPMetaData.ABI instead.
var PDPABI = PDPMetaData.ABI

// PDPBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PDPMetaData.Bin instead.
var PDPBin = PDPMetaData.Bin

// DeployPDP deploys a new Ethereum contract, binding an instance of PDP to it.
func DeployPDP(auth *bind.TransactOpts, backend bind.ContractBackend, _loggerAddress common.Address) (common.Address, *types.Transaction, *PDP, error) {
	parsed, err := PDPMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PDPBin), backend, _loggerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PDP{PDPCaller: PDPCaller{contract: contract}, PDPTransactor: PDPTransactor{contract: contract}, PDPFilterer: PDPFilterer{contract: contract}}, nil
}

// PDP is an auto generated Go binding around an Ethereum contract.
type PDP struct {
	PDPCaller     // Read-only binding to the contract
	PDPTransactor // Write-only binding to the contract
	PDPFilterer   // Log filterer for contract events
}

// PDPCaller is an auto generated read-only Go binding around an Ethereum contract.
type PDPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PDPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PDPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PDPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PDPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PDPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PDPSession struct {
	Contract     *PDP              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PDPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PDPCallerSession struct {
	Contract *PDPCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PDPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PDPTransactorSession struct {
	Contract     *PDPTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PDPRaw is an auto generated low-level Go binding around an Ethereum contract.
type PDPRaw struct {
	Contract *PDP // Generic contract binding to access the raw methods on
}

// PDPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PDPCallerRaw struct {
	Contract *PDPCaller // Generic read-only contract binding to access the raw methods on
}

// PDPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PDPTransactorRaw struct {
	Contract *PDPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPDP creates a new instance of PDP, bound to a specific deployed contract.
func NewPDP(address common.Address, backend bind.ContractBackend) (*PDP, error) {
	contract, err := bindPDP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PDP{PDPCaller: PDPCaller{contract: contract}, PDPTransactor: PDPTransactor{contract: contract}, PDPFilterer: PDPFilterer{contract: contract}}, nil
}

// NewPDPCaller creates a new read-only instance of PDP, bound to a specific deployed contract.
func NewPDPCaller(address common.Address, caller bind.ContractCaller) (*PDPCaller, error) {
	contract, err := bindPDP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PDPCaller{contract: contract}, nil
}

// NewPDPTransactor creates a new write-only instance of PDP, bound to a specific deployed contract.
func NewPDPTransactor(address common.Address, transactor bind.ContractTransactor) (*PDPTransactor, error) {
	contract, err := bindPDP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PDPTransactor{contract: contract}, nil
}

// NewPDPFilterer creates a new log filterer instance of PDP, bound to a specific deployed contract.
func NewPDPFilterer(address common.Address, filterer bind.ContractFilterer) (*PDPFilterer, error) {
	contract, err := bindPDP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PDPFilterer{contract: contract}, nil
}

// bindPDP binds a generic wrapper to an already deployed contract.
func bindPDP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PDPMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PDP *PDPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PDP.Contract.PDPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PDP *PDPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PDP.Contract.PDPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PDP *PDPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PDP.Contract.PDPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PDP *PDPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PDP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PDP *PDPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PDP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PDP *PDPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PDP.Contract.contract.Transact(opts, method, params...)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_PDP *PDPCaller) Logger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PDP.contract.Call(opts, &out, "logger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_PDP *PDPSession) Logger() (common.Address, error) {
	return _PDP.Contract.Logger(&_PDP.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_PDP *PDPCallerSession) Logger() (common.Address, error) {
	return _PDP.Contract.Logger(&_PDP.CallOpts)
}

// EvaluateAccess is a paid mutator transaction binding the contract method 0xa9a11d15.
//
// Solidity: function evaluateAccess(uint256 _entityId, int256 zone, string startTime, string endTime) returns(bool)
func (_PDP *PDPTransactor) EvaluateAccess(opts *bind.TransactOpts, _entityId *big.Int, zone *big.Int, startTime string, endTime string) (*types.Transaction, error) {
	return _PDP.contract.Transact(opts, "evaluateAccess", _entityId, zone, startTime, endTime)
}

// EvaluateAccess is a paid mutator transaction binding the contract method 0xa9a11d15.
//
// Solidity: function evaluateAccess(uint256 _entityId, int256 zone, string startTime, string endTime) returns(bool)
func (_PDP *PDPSession) EvaluateAccess(_entityId *big.Int, zone *big.Int, startTime string, endTime string) (*types.Transaction, error) {
	return _PDP.Contract.EvaluateAccess(&_PDP.TransactOpts, _entityId, zone, startTime, endTime)
}

// EvaluateAccess is a paid mutator transaction binding the contract method 0xa9a11d15.
//
// Solidity: function evaluateAccess(uint256 _entityId, int256 zone, string startTime, string endTime) returns(bool)
func (_PDP *PDPTransactorSession) EvaluateAccess(_entityId *big.Int, zone *big.Int, startTime string, endTime string) (*types.Transaction, error) {
	return _PDP.Contract.EvaluateAccess(&_PDP.TransactOpts, _entityId, zone, startTime, endTime)
}
