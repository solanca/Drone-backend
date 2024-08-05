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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pipAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prpAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entityId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"granted\",\"type\":\"bool\"}],\"name\":\"AccessDecision\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_entityId\",\"type\":\"uint256\"}],\"name\":\"evaluateAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pip\",\"outputs\":[{\"internalType\":\"contractDroneContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prp\",\"outputs\":[{\"internalType\":\"contractPolicyContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610e6b380380610e6b83398181016040528101906100319190610115565b815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610153565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100e4826100bb565b9050919050565b6100f4816100da565b81146100fe575f80fd5b50565b5f8151905061010f816100eb565b92915050565b5f806040838503121561012b5761012a6100b7565b5b5f61013885828601610101565b925050602061014985828601610101565b9150509250929050565b610d0b806101605f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80638b8245cc14610043578063d0345cd914610073578063d741e2f914610091575b5f80fd5b61005d600480360381019061005891906105b3565b6100af565b60405161006a91906105f8565b60405180910390f35b61007b61026a565b604051610088919061068b565b60405180910390f35b61009961028f565b6040516100a691906106c4565b60405180910390f35b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dbd1be77846040518263ffffffff1660e01b815260040161010991906106ec565b5f60405180830381865afa158015610123573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061014b919061090d565b90505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639d47b49c83604001516040518263ffffffff1660e01b81526004016101ac9190610963565b5f60405180830381865afa1580156101c6573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906101ee9190610a29565b90505f804290505f62015180826102059190610a9d565b905061021a81856040015186606001516102b2565b1561022457600192505b7f91e57896bf675ec8aea77f6172d04b52c5ffb7c87d49e67a2342a35b750da4e68784604051610255929190610acd565b60405180910390a18295505050505050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f806102bd84610320565b90505f6102c984610320565b90508082116102eb578186101580156102e25750808611155b92505050610319565b81861180156102fd5750620151808611155b8061031457505f86101580156103135750808611155b5b925050505b9392505050565b5f808290506008815114610369576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036090610b4e565b60405180910390fd5b5f6030826001815181106103805761037f610b6c565b5b602001015160f81c60f81b60f81c60ff1661039b9190610bc6565b600a6030845f815181106103b2576103b1610b6c565b5b602001015160f81c60f81b60f81c60ff166103cd9190610bc6565b6103d79190610bf9565b6103e19190610c3a565b90505f6030836004815181106103fa576103f9610b6c565b5b602001015160f81c60f81b60f81c60ff166104159190610bc6565b600a60308560038151811061042d5761042c610b6c565b5b602001015160f81c60f81b60f81c60ff166104489190610bc6565b6104529190610bf9565b61045c9190610c3a565b90505f60308460078151811061047557610474610b6c565b5b602001015160f81c60f81b60f81c60ff166104909190610bc6565b600a6030866006815181106104a8576104a7610b6c565b5b602001015160f81c60f81b60f81c60ff166104c39190610bc6565b6104cd9190610bf9565b6104d79190610c3a565b90506018831080156104e95750603c82105b80156104f55750603c81105b610534576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052b90610cb7565b60405180910390fd5b80603c836105429190610bf9565b610e10856105509190610bf9565b61055a9190610c3a565b6105649190610c3a565b945050505050919050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61059281610580565b811461059c575f80fd5b50565b5f813590506105ad81610589565b92915050565b5f602082840312156105c8576105c7610578565b5b5f6105d58482850161059f565b91505092915050565b5f8115159050919050565b6105f2816105de565b82525050565b5f60208201905061060b5f8301846105e9565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61065361064e61064984610611565b610630565b610611565b9050919050565b5f61066482610639565b9050919050565b5f6106758261065a565b9050919050565b6106858161066b565b82525050565b5f60208201905061069e5f83018461067c565b92915050565b5f6106ae8261065a565b9050919050565b6106be816106a4565b82525050565b5f6020820190506106d75f8301846106b5565b92915050565b6106e681610580565b82525050565b5f6020820190506106ff5f8301846106dd565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61074f82610709565b810181811067ffffffffffffffff8211171561076e5761076d610719565b5b80604052505050565b5f61078061056f565b905061078c8282610746565b919050565b5f80fd5b5f815190506107a381610589565b92915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff8211156107cb576107ca610719565b5b6107d482610709565b9050602081019050919050565b8281835e5f83830152505050565b5f6108016107fc846107b1565b610777565b90508281526020810184848401111561081d5761081c6107ad565b5b6108288482856107e1565b509392505050565b5f82601f830112610844576108436107a9565b5b81516108548482602086016107ef565b91505092915050565b5f819050919050565b61086f8161085d565b8114610879575f80fd5b50565b5f8151905061088a81610866565b92915050565b5f606082840312156108a5576108a4610705565b5b6108af6060610777565b90505f6108be84828501610795565b5f83015250602082015167ffffffffffffffff8111156108e1576108e0610791565b5b6108ed84828501610830565b60208301525060406109018482850161087c565b60408301525092915050565b5f6020828403121561092257610921610578565b5b5f82015167ffffffffffffffff81111561093f5761093e61057c565b5b61094b84828501610890565b91505092915050565b61095d8161085d565b82525050565b5f6020820190506109765f830184610954565b92915050565b5f6080828403121561099157610990610705565b5b61099b6080610777565b90505f6109aa84828501610795565b5f8301525060206109bd8482850161087c565b602083015250604082015167ffffffffffffffff8111156109e1576109e0610791565b5b6109ed84828501610830565b604083015250606082015167ffffffffffffffff811115610a1157610a10610791565b5b610a1d84828501610830565b60608301525092915050565b5f60208284031215610a3e57610a3d610578565b5b5f82015167ffffffffffffffff811115610a5b57610a5a61057c565b5b610a678482850161097c565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610aa782610580565b9150610ab283610580565b925082610ac257610ac1610a70565b5b828206905092915050565b5f604082019050610ae05f8301856106dd565b610aed60208301846105e9565b9392505050565b5f82825260208201905092915050565b7f696e76616c69642074696d6520666f726d6174000000000000000000000000005f82015250565b5f610b38601383610af4565b9150610b4382610b04565b602082019050919050565b5f6020820190508181035f830152610b6581610b2c565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610bd082610580565b9150610bdb83610580565b9250828203905081811115610bf357610bf2610b99565b5b92915050565b5f610c0382610580565b9150610c0e83610580565b9250828202610c1c81610580565b91508282048414831517610c3357610c32610b99565b5b5092915050565b5f610c4482610580565b9150610c4f83610580565b9250828201905080821115610c6757610c66610b99565b5b92915050565b7f496e76616c69642074696d652076616c756500000000000000000000000000005f82015250565b5f610ca1601283610af4565b9150610cac82610c6d565b602082019050919050565b5f6020820190508181035f830152610cce81610c95565b905091905056fea2646970667358221220164919763ef07056311b327f64206bbd973ccd4927215bc0b2cd0a1b1a5d3dd364736f6c634300081a0033",
}

// PDPABI is the input ABI used to generate the binding from.
// Deprecated: Use PDPMetaData.ABI instead.
var PDPABI = PDPMetaData.ABI

// PDPBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PDPMetaData.Bin instead.
var PDPBin = PDPMetaData.Bin

// DeployPDP deploys a new Ethereum contract, binding an instance of PDP to it.
func DeployPDP(auth *bind.TransactOpts, backend bind.ContractBackend, _pipAddress common.Address, _prpAddress common.Address) (common.Address, *types.Transaction, *PDP, error) {
	parsed, err := PDPMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PDPBin), backend, _pipAddress, _prpAddress)
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

// Pip is a free data retrieval call binding the contract method 0xd741e2f9.
//
// Solidity: function pip() view returns(address)
func (_PDP *PDPCaller) Pip(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PDP.contract.Call(opts, &out, "pip")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pip is a free data retrieval call binding the contract method 0xd741e2f9.
//
// Solidity: function pip() view returns(address)
func (_PDP *PDPSession) Pip() (common.Address, error) {
	return _PDP.Contract.Pip(&_PDP.CallOpts)
}

// Pip is a free data retrieval call binding the contract method 0xd741e2f9.
//
// Solidity: function pip() view returns(address)
func (_PDP *PDPCallerSession) Pip() (common.Address, error) {
	return _PDP.Contract.Pip(&_PDP.CallOpts)
}

// Prp is a free data retrieval call binding the contract method 0xd0345cd9.
//
// Solidity: function prp() view returns(address)
func (_PDP *PDPCaller) Prp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PDP.contract.Call(opts, &out, "prp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prp is a free data retrieval call binding the contract method 0xd0345cd9.
//
// Solidity: function prp() view returns(address)
func (_PDP *PDPSession) Prp() (common.Address, error) {
	return _PDP.Contract.Prp(&_PDP.CallOpts)
}

// Prp is a free data retrieval call binding the contract method 0xd0345cd9.
//
// Solidity: function prp() view returns(address)
func (_PDP *PDPCallerSession) Prp() (common.Address, error) {
	return _PDP.Contract.Prp(&_PDP.CallOpts)
}

// EvaluateAccess is a paid mutator transaction binding the contract method 0x8b8245cc.
//
// Solidity: function evaluateAccess(uint256 _entityId) returns(bool)
func (_PDP *PDPTransactor) EvaluateAccess(opts *bind.TransactOpts, _entityId *big.Int) (*types.Transaction, error) {
	return _PDP.contract.Transact(opts, "evaluateAccess", _entityId)
}

// EvaluateAccess is a paid mutator transaction binding the contract method 0x8b8245cc.
//
// Solidity: function evaluateAccess(uint256 _entityId) returns(bool)
func (_PDP *PDPSession) EvaluateAccess(_entityId *big.Int) (*types.Transaction, error) {
	return _PDP.Contract.EvaluateAccess(&_PDP.TransactOpts, _entityId)
}

// EvaluateAccess is a paid mutator transaction binding the contract method 0x8b8245cc.
//
// Solidity: function evaluateAccess(uint256 _entityId) returns(bool)
func (_PDP *PDPTransactorSession) EvaluateAccess(_entityId *big.Int) (*types.Transaction, error) {
	return _PDP.Contract.EvaluateAccess(&_PDP.TransactOpts, _entityId)
}

// PDPAccessDecisionIterator is returned from FilterAccessDecision and is used to iterate over the raw logs and unpacked data for AccessDecision events raised by the PDP contract.
type PDPAccessDecisionIterator struct {
	Event *PDPAccessDecision // Event containing the contract specifics and raw log

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
func (it *PDPAccessDecisionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PDPAccessDecision)
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
		it.Event = new(PDPAccessDecision)
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
func (it *PDPAccessDecisionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PDPAccessDecisionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PDPAccessDecision represents a AccessDecision event raised by the PDP contract.
type PDPAccessDecision struct {
	EntityId *big.Int
	Granted  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccessDecision is a free log retrieval operation binding the contract event 0x91e57896bf675ec8aea77f6172d04b52c5ffb7c87d49e67a2342a35b750da4e6.
//
// Solidity: event AccessDecision(uint256 entityId, bool granted)
func (_PDP *PDPFilterer) FilterAccessDecision(opts *bind.FilterOpts) (*PDPAccessDecisionIterator, error) {

	logs, sub, err := _PDP.contract.FilterLogs(opts, "AccessDecision")
	if err != nil {
		return nil, err
	}
	return &PDPAccessDecisionIterator{contract: _PDP.contract, event: "AccessDecision", logs: logs, sub: sub}, nil
}

// WatchAccessDecision is a free log subscription operation binding the contract event 0x91e57896bf675ec8aea77f6172d04b52c5ffb7c87d49e67a2342a35b750da4e6.
//
// Solidity: event AccessDecision(uint256 entityId, bool granted)
func (_PDP *PDPFilterer) WatchAccessDecision(opts *bind.WatchOpts, sink chan<- *PDPAccessDecision) (event.Subscription, error) {

	logs, sub, err := _PDP.contract.WatchLogs(opts, "AccessDecision")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PDPAccessDecision)
				if err := _PDP.contract.UnpackLog(event, "AccessDecision", log); err != nil {
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

// ParseAccessDecision is a log parse operation binding the contract event 0x91e57896bf675ec8aea77f6172d04b52c5ffb7c87d49e67a2342a35b750da4e6.
//
// Solidity: event AccessDecision(uint256 entityId, bool granted)
func (_PDP *PDPFilterer) ParseAccessDecision(log types.Log) (*PDPAccessDecision, error) {
	event := new(PDPAccessDecision)
	if err := _PDP.contract.UnpackLog(event, "AccessDecision", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
