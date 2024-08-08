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
	Bin: "0x608060405234801561000f575f80fd5b506040516114ae3803806114ae833981810160405281019061003191906100d4565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f80fd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b6113a28061010c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063a9a11d1514610038578063f24ccbfe14610068575b5f80fd5b610052600480360381019061004d9190610b84565b610086565b60405161005f9190610c3a565b60405180910390f35b61007061024f565b60405161007d9190610ccd565b60405180910390f35b5f805f90505f612a304261009a9190610d13565b90505f62015180826100ac9190610d73565b90505f6100b883610272565b90506100c582888861031b565b156100cf57600193505b5f6040518060400160405280600e81526020017f416363657373205265717565737400000000000000000000000000000000000081525090505f6101128b610389565b61011b8b610507565b848861015c576040518060400160405280600681526020017f44656e6965640000000000000000000000000000000000000000000000000000815250610193565b6040518060400160405280600a81526020017f50726f67726573736564000000000000000000000000000000000000000000008152505b6040516020016101a69493929190610f1d565b60405160208183030381529060405290505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166397f8abd883836040518363ffffffff1660e01b8152600401610211929190610fce565b5f604051808303815f87803b158015610228575f80fd5b505af115801561023a573d5f803e3d5ffd5b50505050859650505050505050949350505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60605f6018603c80856102859190611003565b61028f9190611003565b6102999190610d73565b90505f603c80856102aa9190611003565b6102b49190610d73565b90505f603c856102c49190610d73565b90505f6102d084610737565b90505f6102dc84610737565b90505f6102e884610737565b90508282826040516020016102ff9392919061107d565b6040516020818303038152906040529650505050505050919050565b5f8061032684610782565b90505f61033284610782565b90508082116103545781861015801561034b5750808611155b92505050610382565b81861180156103665750620151808611155b8061037d57505f861015801561037c5750808611155b5b925050505b9392505050565b60605f82036103cf576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610502565b5f8290505f5b5f82146103fe5780806103e7906110c3565b915050600a826103f79190611003565b91506103d5565b5f8167ffffffffffffffff81111561041957610418610a60565b5b6040519080825280601f01601f19166020018201604052801561044b5781602001600182028036833780820191505090505b5090505f8290505b5f86146104fa57600181610467919061110a565b90505f600a80886104789190611003565b610482919061113d565b8761048d919061110a565b6030610499919061118a565b90505f8160f81b9050808484815181106104b6576104b56111be565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a886104f19190611003565b97505050610453565b819450505050505b919050565b60605f820361054d576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610732565b5f80831290505f8161055f578361056a565b83610569906111eb565b5b90505f5b5f8214610597578080610580906110c3565b915050600a826105909190611003565b915061056e565b82156105ac5780806105a8906110c3565b9150505b5f8167ffffffffffffffff8111156105c7576105c6610a60565b5b6040519080825280601f01601f1916602001820160405280156105f95781602001600182028036833780820191505090505b5090505f8290508461060b5786610616565b86610615906111eb565b5b93505b5f84146106c05760018161062d919061110a565b90505f600a808661063e9190611003565b610648919061113d565b85610653919061110a565b603061065f919061118a565b90505f8160f81b90508084848151811061067c5761067b6111be565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a866106b79190611003565b95505050610619565b8415610729577f2d00000000000000000000000000000000000000000000000000000000000000825f815181106106fa576106f96111be565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505b81955050505050505b919050565b6060600a8210156107715761074b82610389565b60405160200161075b919061127b565b604051602081830303815290604052905061077d565b61077a82610389565b90505b919050565b5f8082905060088151146107cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c2906112e6565b60405180910390fd5b5f6030826001815181106107e2576107e16111be565b5b602001015160f81c60f81b60f81c60ff166107fd919061110a565b600a6030845f81518110610814576108136111be565b5b602001015160f81c60f81b60f81c60ff1661082f919061110a565b610839919061113d565b6108439190610d13565b90505f60308360048151811061085c5761085b6111be565b5b602001015160f81c60f81b60f81c60ff16610877919061110a565b600a60308560038151811061088f5761088e6111be565b5b602001015160f81c60f81b60f81c60ff166108aa919061110a565b6108b4919061113d565b6108be9190610d13565b90505f6030846007815181106108d7576108d66111be565b5b602001015160f81c60f81b60f81c60ff166108f2919061110a565b600a60308660068151811061090a576109096111be565b5b602001015160f81c60f81b60f81c60ff16610925919061110a565b61092f919061113d565b6109399190610d13565b905060188310801561094b5750603c82105b80156109575750603c81105b610996576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098d9061134e565b60405180910390fd5b80603c836109a4919061113d565b610e10856109b2919061113d565b6109bc9190610d13565b6109c69190610d13565b945050505050919050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6109f4816109e2565b81146109fe575f80fd5b50565b5f81359050610a0f816109eb565b92915050565b5f819050919050565b610a2781610a15565b8114610a31575f80fd5b50565b5f81359050610a4281610a1e565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610a9682610a50565b810181811067ffffffffffffffff82111715610ab557610ab4610a60565b5b80604052505050565b5f610ac76109d1565b9050610ad38282610a8d565b919050565b5f67ffffffffffffffff821115610af257610af1610a60565b5b610afb82610a50565b9050602081019050919050565b828183375f83830152505050565b5f610b28610b2384610ad8565b610abe565b905082815260208101848484011115610b4457610b43610a4c565b5b610b4f848285610b08565b509392505050565b5f82601f830112610b6b57610b6a610a48565b5b8135610b7b848260208601610b16565b91505092915050565b5f805f8060808587031215610b9c57610b9b6109da565b5b5f610ba987828801610a01565b9450506020610bba87828801610a34565b935050604085013567ffffffffffffffff811115610bdb57610bda6109de565b5b610be787828801610b57565b925050606085013567ffffffffffffffff811115610c0857610c076109de565b5b610c1487828801610b57565b91505092959194509250565b5f8115159050919050565b610c3481610c20565b82525050565b5f602082019050610c4d5f830184610c2b565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610c95610c90610c8b84610c53565b610c72565b610c53565b9050919050565b5f610ca682610c7b565b9050919050565b5f610cb782610c9c565b9050919050565b610cc781610cad565b82525050565b5f602082019050610ce05f830184610cbe565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610d1d826109e2565b9150610d28836109e2565b9250828201905080821115610d4057610d3f610ce6565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610d7d826109e2565b9150610d88836109e2565b925082610d9857610d97610d46565b5b828206905092915050565b5f81905092915050565b7f44726f6e652049443a20000000000000000000000000000000000000000000005f82015250565b5f610de1600a83610da3565b9150610dec82610dad565b600a82019050919050565b5f81519050919050565b8281835e5f83830152505050565b5f610e1982610df7565b610e238185610da3565b9350610e33818560208601610e01565b80840191505092915050565b7f2c205a6f6e653a200000000000000000000000000000000000000000000000005f82015250565b5f610e73600883610da3565b9150610e7e82610e3f565b600882019050919050565b7f2c2043757272656e7454696d653a2000000000000000000000000000000000005f82015250565b5f610ebd600f83610da3565b9150610ec882610e89565b600f82019050919050565b7f2c204772616e7465643a200000000000000000000000000000000000000000005f82015250565b5f610f07600b83610da3565b9150610f1282610ed3565b600b82019050919050565b5f610f2782610dd5565b9150610f338287610e0f565b9150610f3e82610e67565b9150610f4a8286610e0f565b9150610f5582610eb1565b9150610f618285610e0f565b9150610f6c82610efb565b9150610f788284610e0f565b915081905095945050505050565b5f82825260208201905092915050565b5f610fa082610df7565b610faa8185610f86565b9350610fba818560208601610e01565b610fc381610a50565b840191505092915050565b5f6040820190508181035f830152610fe68185610f96565b90508181036020830152610ffa8184610f96565b90509392505050565b5f61100d826109e2565b9150611018836109e2565b92508261102857611027610d46565b5b828204905092915050565b7f3a000000000000000000000000000000000000000000000000000000000000005f82015250565b5f611067600183610da3565b915061107282611033565b600182019050919050565b5f6110888286610e0f565b91506110938261105b565b915061109f8285610e0f565b91506110aa8261105b565b91506110b68284610e0f565b9150819050949350505050565b5f6110cd826109e2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110ff576110fe610ce6565b5b600182019050919050565b5f611114826109e2565b915061111f836109e2565b925082820390508181111561113757611136610ce6565b5b92915050565b5f611147826109e2565b9150611152836109e2565b9250828202611160816109e2565b9150828204841483151761117757611176610ce6565b5b5092915050565b5f60ff82169050919050565b5f6111948261117e565b915061119f8361117e565b9250828201905060ff8111156111b8576111b7610ce6565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6111f582610a15565b91507f8000000000000000000000000000000000000000000000000000000000000000820361122757611226610ce6565b5b815f039050919050565b7f30000000000000000000000000000000000000000000000000000000000000005f82015250565b5f611265600183610da3565b915061127082611231565b600182019050919050565b5f61128582611259565b91506112918284610e0f565b915081905092915050565b7f696e76616c69642074696d6520666f726d6174000000000000000000000000005f82015250565b5f6112d0601383610f86565b91506112db8261129c565b602082019050919050565b5f6020820190508181035f8301526112fd816112c4565b9050919050565b7f496e76616c69642074696d652076616c756500000000000000000000000000005f82015250565b5f611338601283610f86565b915061134382611304565b602082019050919050565b5f6020820190508181035f8301526113658161132c565b905091905056fea264697066735822122071c9f9236b529cc4c0a3cf015f802fb2d8e5e23231d482a3f6c7817fdf508e1e64736f6c634300081a0033",
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
