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
	Timestamp *big.Int
	Data      string
}

// LoggingMetaData contains all meta data concerning the Logging contract.
var LoggingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"ActionLogged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structLoggingContract.LogEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"logAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"logs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50610d288061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806397f8abd814610043578063e79899bd1461005f578063fe290ccb14610092575b5f80fd5b61005d6004803603810190610058919061066b565b6100b0565b005b61007960048036038101906100749190610714565b6101e0565b60405161008994939291906107ed565b60405180910390f35b61009a610345565b6040516100a791906109c6565b60405180910390f35b5f60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184815260200142815260200183815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816101679190610be0565b506040820151816002015560608201518160030190816101879190610be0565b5050503373ffffffffffffffffffffffffffffffffffffffff167fbadb834156528988a5030c219956a45300d7fbc59c36a0c421bb12fdf9a3ba178342846040516101d493929190610caf565b60405180910390a25050565b5f81815481106101ee575f80fd5b905f5260205f2090600402015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461023290610a13565b80601f016020809104026020016040519081016040528092919081815260200182805461025e90610a13565b80156102a95780601f10610280576101008083540402835291602001916102a9565b820191905f5260205f20905b81548152906001019060200180831161028c57829003601f168201915b5050505050908060020154908060030180546102c490610a13565b80601f01602080910402602001604051908101604052809291908181526020018280546102f090610a13565b801561033b5780601f106103125761010080835404028352916020019161033b565b820191905f5260205f20905b81548152906001019060200180831161031e57829003601f168201915b5050505050905084565b60605f805480602002602001604051908101604052809291908181526020015f905b82821015610515578382905f5260205f2090600402016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180546103ec90610a13565b80601f016020809104026020016040519081016040528092919081815260200182805461041890610a13565b80156104635780601f1061043a57610100808354040283529160200191610463565b820191905f5260205f20905b81548152906001019060200180831161044657829003601f168201915b505050505081526020016002820154815260200160038201805461048690610a13565b80601f01602080910402602001604051908101604052809291908181526020018280546104b290610a13565b80156104fd5780601f106104d4576101008083540402835291602001916104fd565b820191905f5260205f20905b8154815290600101906020018083116104e057829003601f168201915b50505050508152505081526020019060010190610367565b50505050905090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61057d82610537565b810181811067ffffffffffffffff8211171561059c5761059b610547565b5b80604052505050565b5f6105ae61051e565b90506105ba8282610574565b919050565b5f67ffffffffffffffff8211156105d9576105d8610547565b5b6105e282610537565b9050602081019050919050565b828183375f83830152505050565b5f61060f61060a846105bf565b6105a5565b90508281526020810184848401111561062b5761062a610533565b5b6106368482856105ef565b509392505050565b5f82601f8301126106525761065161052f565b5b81356106628482602086016105fd565b91505092915050565b5f806040838503121561068157610680610527565b5b5f83013567ffffffffffffffff81111561069e5761069d61052b565b5b6106aa8582860161063e565b925050602083013567ffffffffffffffff8111156106cb576106ca61052b565b5b6106d78582860161063e565b9150509250929050565b5f819050919050565b6106f3816106e1565b81146106fd575f80fd5b50565b5f8135905061070e816106ea565b92915050565b5f6020828403121561072957610728610527565b5b5f61073684828501610700565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6107688261073f565b9050919050565b6107788161075e565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6107b08261077e565b6107ba8185610788565b93506107ca818560208601610798565b6107d381610537565b840191505092915050565b6107e7816106e1565b82525050565b5f6080820190506108005f83018761076f565b818103602083015261081281866107a6565b905061082160408301856107de565b818103606083015261083381846107a6565b905095945050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6108708161075e565b82525050565b5f82825260208201905092915050565b5f6108908261077e565b61089a8185610876565b93506108aa818560208601610798565b6108b381610537565b840191505092915050565b6108c7816106e1565b82525050565b5f608083015f8301516108e25f860182610867565b50602083015184820360208601526108fa8282610886565b915050604083015161090f60408601826108be565b50606083015184820360608601526109278282610886565b9150508091505092915050565b5f61093f83836108cd565b905092915050565b5f602082019050919050565b5f61095d8261083e565b6109678185610848565b93508360208202850161097985610858565b805f5b858110156109b457848403895281516109958582610934565b94506109a083610947565b925060208a0199505060018101905061097c565b50829750879550505050505092915050565b5f6020820190508181035f8301526109de8184610953565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610a2a57607f821691505b602082108103610a3d57610a3c6109e6565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610a9f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a64565b610aa98683610a64565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610ae4610adf610ada846106e1565b610ac1565b6106e1565b9050919050565b5f819050919050565b610afd83610aca565b610b11610b0982610aeb565b848454610a70565b825550505050565b5f90565b610b25610b19565b610b30818484610af4565b505050565b5b81811015610b5357610b485f82610b1d565b600181019050610b36565b5050565b601f821115610b9857610b6981610a43565b610b7284610a55565b81016020851015610b81578190505b610b95610b8d85610a55565b830182610b35565b50505b505050565b5f82821c905092915050565b5f610bb85f1984600802610b9d565b1980831691505092915050565b5f610bd08383610ba9565b9150826002028217905092915050565b610be98261077e565b67ffffffffffffffff811115610c0257610c01610547565b5b610c0c8254610a13565b610c17828285610b57565b5f60209050601f831160018114610c48575f8415610c36578287015190505b610c408582610bc5565b865550610ca7565b601f198416610c5686610a43565b5f5b82811015610c7d57848901518255600182019150602085019450602081019050610c58565b86831015610c9a5784890151610c96601f891682610ba9565b8355505b6001600288020188555050505b505050505050565b5f6060820190508181035f830152610cc781866107a6565b9050610cd660208301856107de565b8181036040830152610ce881846107a6565b905094935050505056fea264697066735822122008ecb58d7c73ede2b316ff2be333d51c818854d7db7524a8f49b7a9da77a40d364736f6c634300081a0033",
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
// Solidity: function getLogs() view returns((address,string,uint256,string)[])
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
// Solidity: function getLogs() view returns((address,string,uint256,string)[])
func (_Logging *LoggingSession) GetLogs() ([]LoggingContractLogEntry, error) {
	return _Logging.Contract.GetLogs(&_Logging.CallOpts)
}

// GetLogs is a free data retrieval call binding the contract method 0xfe290ccb.
//
// Solidity: function getLogs() view returns((address,string,uint256,string)[])
func (_Logging *LoggingCallerSession) GetLogs() ([]LoggingContractLogEntry, error) {
	return _Logging.Contract.GetLogs(&_Logging.CallOpts)
}

// Logs is a free data retrieval call binding the contract method 0xe79899bd.
//
// Solidity: function logs(uint256 ) view returns(address user, string action, uint256 timestamp, string data)
func (_Logging *LoggingCaller) Logs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User      common.Address
	Action    string
	Timestamp *big.Int
	Data      string
}, error) {
	var out []interface{}
	err := _Logging.contract.Call(opts, &out, "logs", arg0)

	outstruct := new(struct {
		User      common.Address
		Action    string
		Timestamp *big.Int
		Data      string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Action = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Logs is a free data retrieval call binding the contract method 0xe79899bd.
//
// Solidity: function logs(uint256 ) view returns(address user, string action, uint256 timestamp, string data)
func (_Logging *LoggingSession) Logs(arg0 *big.Int) (struct {
	User      common.Address
	Action    string
	Timestamp *big.Int
	Data      string
}, error) {
	return _Logging.Contract.Logs(&_Logging.CallOpts, arg0)
}

// Logs is a free data retrieval call binding the contract method 0xe79899bd.
//
// Solidity: function logs(uint256 ) view returns(address user, string action, uint256 timestamp, string data)
func (_Logging *LoggingCallerSession) Logs(arg0 *big.Int) (struct {
	User      common.Address
	Action    string
	Timestamp *big.Int
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
	Timestamp *big.Int
	Data      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterActionLogged is a free log retrieval operation binding the contract event 0xbadb834156528988a5030c219956a45300d7fbc59c36a0c421bb12fdf9a3ba17.
//
// Solidity: event ActionLogged(address indexed user, string action, uint256 timestamp, string data)
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

// WatchActionLogged is a free log subscription operation binding the contract event 0xbadb834156528988a5030c219956a45300d7fbc59c36a0c421bb12fdf9a3ba17.
//
// Solidity: event ActionLogged(address indexed user, string action, uint256 timestamp, string data)
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

// ParseActionLogged is a log parse operation binding the contract event 0xbadb834156528988a5030c219956a45300d7fbc59c36a0c421bb12fdf9a3ba17.
//
// Solidity: event ActionLogged(address indexed user, string action, uint256 timestamp, string data)
func (_Logging *LoggingFilterer) ParseActionLogged(log types.Log) (*LoggingActionLogged, error) {
	event := new(LoggingActionLogged)
	if err := _Logging.contract.UnpackLog(event, "ActionLogged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
