// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Drone

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

// DroneContractDrone is an auto generated low-level Go binding around an user-defined struct.
type DroneContractDrone struct {
	Id     *big.Int
	Model  string
	Zone   *big.Int
	Exists bool
}

// DroneMetaData contains all meta data concerning the Drone contract.
var DroneMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_model\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"createDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"model\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"}],\"name\":\"DroneCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DroneDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"model\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"}],\"name\":\"DroneUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_model\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"updateDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drones\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getDrone\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structDroneContract.Drone\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDrones\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structDroneContract.Drone[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506112f08061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061007b575f3560e01c80639341bc00116100595780639341bc00146100d7578063d7c7a2ba1461010a578063dbd1be7714610126578063ed88eb6d146101565761007b565b806307d5ebf51461007f578063276f50381461009d57806361b8ce8c146100b9575b5f80fd5b610087610172565b6040516100949190610ab2565b60405180910390f35b6100b760048036038101906100b29190610c39565b610399565b005b6100c1610487565b6040516100ce9190610ca2565b60405180910390f35b6100f160048036038101906100ec9190610ce5565b61048d565b6040516101019493929190610d76565b60405180910390f35b610124600480360381019061011f9190610dc0565b610559565b005b610140600480360381019061013b9190610ce5565b610652565b60405161014d9190610e8c565b60405180910390f35b610170600480360381019061016b9190610ce5565b6107c8565b005b60605f805b5f805490508110156101d3575f818154811061019657610195610eac565b5b905f5260205f2090600402016003015f9054906101000a900460ff16156101c65781806101c290610f06565b9250505b8080600101915050610177565b505f8167ffffffffffffffff8111156101ef576101ee610aeb565b5b60405190808252806020026020018201604052801561022857816020015b6102156108b6565b81526020019060019003908161020d5790505b5090505f805b5f8054905081101561038f575f818154811061024d5761024c610eac565b5b905f5260205f2090600402016003015f9054906101000a900460ff1615610382575f818154811061028157610280610eac565b5b905f5260205f2090600402016040518060800160405290815f82015481526020016001820180546102b190610f7a565b80601f01602080910402602001604051908101604052809291908181526020018280546102dd90610f7a565b80156103285780601f106102ff57610100808354040283529160200191610328565b820191905f5260205f20905b81548152906001019060200180831161030b57829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff16151515158152505083838151811061036857610367610eac565b5b6020026020010181905250818061037e90610f06565b9250505b808060010191505061022e565b5081935050505090565b5f6040518060800160405280600154815260200184815260200183815260200160011515815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015560208201518160010190816104039190611147565b50604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050507f6cc4ed91f20ea27eda237a37a11e00cb1e9e8cc036186b136f63b5dd811b64b5600154838360405161046493929190611216565b60405180910390a160015f81548092919061047e90610f06565b91905055505050565b60015481565b5f818154811061049b575f80fd5b905f5260205f2090600402015f91509050805f0154908060010180546104c090610f7a565b80601f01602080910402602001604051908101604052809291908181526020018280546104ec90610f7a565b80156105375780601f1061050e57610100808354040283529160200191610537565b820191905f5260205f20905b81548152906001019060200180831161051a57829003601f168201915b505050505090806002015490806003015f9054906101000a900460ff16905084565b6001548310801561059557505f838154811061057857610577610eac565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b6105d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105cb9061129c565b60405180910390fd5b5f8084815481106105e8576105e7610eac565b5b905f5260205f2090600402019050828160010190816106079190611147565b508181600201819055507f84303538537fc260b1a75f003f2666ebee27ee671e3fde7c9f365c5557f0dfe984848460405161064493929190611216565b60405180910390a150505050565b61065a6108b6565b6001548210801561069657505f828154811061067957610678610eac565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b6106d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106cc9061129c565b60405180910390fd5b5f8083815481106106e9576106e8610eac565b5b905f5260205f2090600402016040518060800160405290815f820154815260200160018201805461071990610f7a565b80601f016020809104026020016040519081016040528092919081815260200182805461074590610f7a565b80156107905780601f1061076757610100808354040283529160200191610790565b820191905f5260205f20905b81548152906001019060200180831161077357829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff161515151581525050905080915050919050565b6001548110801561080457505f81815481106107e7576107e6610eac565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b610843576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083a9061129c565b60405180910390fd5b5f80828154811061085757610856610eac565b5b905f5260205f2090600402016003015f6101000a81548160ff0219169083151502179055507fa398f31db5e7963635c8e05cb7bdd246a84574c0b1671840aa00e069a9794b12816040516108ab9190610ca2565b60405180910390a150565b60405180608001604052805f8152602001606081526020015f81526020015f151581525090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61091881610906565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6109608261091e565b61096a8185610928565b935061097a818560208601610938565b61098381610946565b840191505092915050565b5f819050919050565b6109a08161098e565b82525050565b5f8115159050919050565b6109ba816109a6565b82525050565b5f608083015f8301516109d55f86018261090f565b50602083015184820360208601526109ed8282610956565b9150506040830151610a026040860182610997565b506060830151610a1560608601826109b1565b508091505092915050565b5f610a2b83836109c0565b905092915050565b5f602082019050919050565b5f610a49826108dd565b610a5381856108e7565b935083602082028501610a65856108f7565b805f5b85811015610aa05784840389528151610a818582610a20565b9450610a8c83610a33565b925060208a01995050600181019050610a68565b50829750879550505050505092915050565b5f6020820190508181035f830152610aca8184610a3f565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610b2182610946565b810181811067ffffffffffffffff82111715610b4057610b3f610aeb565b5b80604052505050565b5f610b52610ad2565b9050610b5e8282610b18565b919050565b5f67ffffffffffffffff821115610b7d57610b7c610aeb565b5b610b8682610946565b9050602081019050919050565b828183375f83830152505050565b5f610bb3610bae84610b63565b610b49565b905082815260208101848484011115610bcf57610bce610ae7565b5b610bda848285610b93565b509392505050565b5f82601f830112610bf657610bf5610ae3565b5b8135610c06848260208601610ba1565b91505092915050565b610c188161098e565b8114610c22575f80fd5b50565b5f81359050610c3381610c0f565b92915050565b5f8060408385031215610c4f57610c4e610adb565b5b5f83013567ffffffffffffffff811115610c6c57610c6b610adf565b5b610c7885828601610be2565b9250506020610c8985828601610c25565b9150509250929050565b610c9c81610906565b82525050565b5f602082019050610cb55f830184610c93565b92915050565b610cc481610906565b8114610cce575f80fd5b50565b5f81359050610cdf81610cbb565b92915050565b5f60208284031215610cfa57610cf9610adb565b5b5f610d0784828501610cd1565b91505092915050565b5f82825260208201905092915050565b5f610d2a8261091e565b610d348185610d10565b9350610d44818560208601610938565b610d4d81610946565b840191505092915050565b610d618161098e565b82525050565b610d70816109a6565b82525050565b5f608082019050610d895f830187610c93565b8181036020830152610d9b8186610d20565b9050610daa6040830185610d58565b610db76060830184610d67565b95945050505050565b5f805f60608486031215610dd757610dd6610adb565b5b5f610de486828701610cd1565b935050602084013567ffffffffffffffff811115610e0557610e04610adf565b5b610e1186828701610be2565b9250506040610e2286828701610c25565b9150509250925092565b5f608083015f830151610e415f86018261090f565b5060208301518482036020860152610e598282610956565b9150506040830151610e6e6040860182610997565b506060830151610e8160608601826109b1565b508091505092915050565b5f6020820190508181035f830152610ea48184610e2c565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f1082610906565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f4257610f41610ed9565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610f9157607f821691505b602082108103610fa457610fa3610f4d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026110067fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610fcb565b6110108683610fcb565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61104b61104661104184610906565b611028565b610906565b9050919050565b5f819050919050565b61106483611031565b61107861107082611052565b848454610fd7565b825550505050565b5f90565b61108c611080565b61109781848461105b565b505050565b5b818110156110ba576110af5f82611084565b60018101905061109d565b5050565b601f8211156110ff576110d081610faa565b6110d984610fbc565b810160208510156110e8578190505b6110fc6110f485610fbc565b83018261109c565b50505b505050565b5f82821c905092915050565b5f61111f5f1984600802611104565b1980831691505092915050565b5f6111378383611110565b9150826002028217905092915050565b6111508261091e565b67ffffffffffffffff81111561116957611168610aeb565b5b6111738254610f7a565b61117e8282856110be565b5f60209050601f8311600181146111af575f841561119d578287015190505b6111a7858261112c565b86555061120e565b601f1984166111bd86610faa565b5f5b828110156111e4578489015182556001820191506020850194506020810190506111bf565b8683101561120157848901516111fd601f891682611110565b8355505b6001600288020188555050505b505050505050565b5f6060820190506112295f830186610c93565b818103602083015261123b8185610d20565b905061124a6040830184610d58565b949350505050565b7f44726f6e6520646f6573206e6f742065786973740000000000000000000000005f82015250565b5f611286601483610d10565b915061129182611252565b602082019050919050565b5f6020820190508181035f8301526112b38161127a565b905091905056fea264697066735822122014d076b15d4fd2ac4ef9c41078bbe63eb7505d616965c71bdc64976efc9e46b764736f6c634300081a0033",
}

// DroneABI is the input ABI used to generate the binding from.
// Deprecated: Use DroneMetaData.ABI instead.
var DroneABI = DroneMetaData.ABI

// DroneBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DroneMetaData.Bin instead.
var DroneBin = DroneMetaData.Bin

// DeployDrone deploys a new Ethereum contract, binding an instance of Drone to it.
func DeployDrone(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Drone, error) {
	parsed, err := DroneMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DroneBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Drone{DroneCaller: DroneCaller{contract: contract}, DroneTransactor: DroneTransactor{contract: contract}, DroneFilterer: DroneFilterer{contract: contract}}, nil
}

// Drone is an auto generated Go binding around an Ethereum contract.
type Drone struct {
	DroneCaller     // Read-only binding to the contract
	DroneTransactor // Write-only binding to the contract
	DroneFilterer   // Log filterer for contract events
}

// DroneCaller is an auto generated read-only Go binding around an Ethereum contract.
type DroneCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DroneTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DroneTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DroneFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DroneFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DroneSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DroneSession struct {
	Contract     *Drone            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DroneCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DroneCallerSession struct {
	Contract *DroneCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DroneTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DroneTransactorSession struct {
	Contract     *DroneTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DroneRaw is an auto generated low-level Go binding around an Ethereum contract.
type DroneRaw struct {
	Contract *Drone // Generic contract binding to access the raw methods on
}

// DroneCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DroneCallerRaw struct {
	Contract *DroneCaller // Generic read-only contract binding to access the raw methods on
}

// DroneTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DroneTransactorRaw struct {
	Contract *DroneTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDrone creates a new instance of Drone, bound to a specific deployed contract.
func NewDrone(address common.Address, backend bind.ContractBackend) (*Drone, error) {
	contract, err := bindDrone(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Drone{DroneCaller: DroneCaller{contract: contract}, DroneTransactor: DroneTransactor{contract: contract}, DroneFilterer: DroneFilterer{contract: contract}}, nil
}

// NewDroneCaller creates a new read-only instance of Drone, bound to a specific deployed contract.
func NewDroneCaller(address common.Address, caller bind.ContractCaller) (*DroneCaller, error) {
	contract, err := bindDrone(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DroneCaller{contract: contract}, nil
}

// NewDroneTransactor creates a new write-only instance of Drone, bound to a specific deployed contract.
func NewDroneTransactor(address common.Address, transactor bind.ContractTransactor) (*DroneTransactor, error) {
	contract, err := bindDrone(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DroneTransactor{contract: contract}, nil
}

// NewDroneFilterer creates a new log filterer instance of Drone, bound to a specific deployed contract.
func NewDroneFilterer(address common.Address, filterer bind.ContractFilterer) (*DroneFilterer, error) {
	contract, err := bindDrone(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DroneFilterer{contract: contract}, nil
}

// bindDrone binds a generic wrapper to an already deployed contract.
func bindDrone(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DroneMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Drone *DroneRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Drone.Contract.DroneCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Drone *DroneRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Drone.Contract.DroneTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Drone *DroneRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Drone.Contract.DroneTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Drone *DroneCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Drone.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Drone *DroneTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Drone.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Drone *DroneTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Drone.Contract.contract.Transact(opts, method, params...)
}

// Drones is a free data retrieval call binding the contract method 0x9341bc00.
//
// Solidity: function drones(uint256 ) view returns(uint256 id, string model, int256 zone, bool exists)
func (_Drone *DroneCaller) Drones(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Model  string
	Zone   *big.Int
	Exists bool
}, error) {
	var out []interface{}
	err := _Drone.contract.Call(opts, &out, "drones", arg0)

	outstruct := new(struct {
		Id     *big.Int
		Model  string
		Zone   *big.Int
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Model = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Zone = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Drones is a free data retrieval call binding the contract method 0x9341bc00.
//
// Solidity: function drones(uint256 ) view returns(uint256 id, string model, int256 zone, bool exists)
func (_Drone *DroneSession) Drones(arg0 *big.Int) (struct {
	Id     *big.Int
	Model  string
	Zone   *big.Int
	Exists bool
}, error) {
	return _Drone.Contract.Drones(&_Drone.CallOpts, arg0)
}

// Drones is a free data retrieval call binding the contract method 0x9341bc00.
//
// Solidity: function drones(uint256 ) view returns(uint256 id, string model, int256 zone, bool exists)
func (_Drone *DroneCallerSession) Drones(arg0 *big.Int) (struct {
	Id     *big.Int
	Model  string
	Zone   *big.Int
	Exists bool
}, error) {
	return _Drone.Contract.Drones(&_Drone.CallOpts, arg0)
}

// GetDrone is a free data retrieval call binding the contract method 0xdbd1be77.
//
// Solidity: function getDrone(uint256 _id) view returns((uint256,string,int256,bool))
func (_Drone *DroneCaller) GetDrone(opts *bind.CallOpts, _id *big.Int) (DroneContractDrone, error) {
	var out []interface{}
	err := _Drone.contract.Call(opts, &out, "getDrone", _id)

	if err != nil {
		return *new(DroneContractDrone), err
	}

	out0 := *abi.ConvertType(out[0], new(DroneContractDrone)).(*DroneContractDrone)

	return out0, err

}

// GetDrone is a free data retrieval call binding the contract method 0xdbd1be77.
//
// Solidity: function getDrone(uint256 _id) view returns((uint256,string,int256,bool))
func (_Drone *DroneSession) GetDrone(_id *big.Int) (DroneContractDrone, error) {
	return _Drone.Contract.GetDrone(&_Drone.CallOpts, _id)
}

// GetDrone is a free data retrieval call binding the contract method 0xdbd1be77.
//
// Solidity: function getDrone(uint256 _id) view returns((uint256,string,int256,bool))
func (_Drone *DroneCallerSession) GetDrone(_id *big.Int) (DroneContractDrone, error) {
	return _Drone.Contract.GetDrone(&_Drone.CallOpts, _id)
}

// GetDrones is a free data retrieval call binding the contract method 0x07d5ebf5.
//
// Solidity: function getDrones() view returns((uint256,string,int256,bool)[])
func (_Drone *DroneCaller) GetDrones(opts *bind.CallOpts) ([]DroneContractDrone, error) {
	var out []interface{}
	err := _Drone.contract.Call(opts, &out, "getDrones")

	if err != nil {
		return *new([]DroneContractDrone), err
	}

	out0 := *abi.ConvertType(out[0], new([]DroneContractDrone)).(*[]DroneContractDrone)

	return out0, err

}

// GetDrones is a free data retrieval call binding the contract method 0x07d5ebf5.
//
// Solidity: function getDrones() view returns((uint256,string,int256,bool)[])
func (_Drone *DroneSession) GetDrones() ([]DroneContractDrone, error) {
	return _Drone.Contract.GetDrones(&_Drone.CallOpts)
}

// GetDrones is a free data retrieval call binding the contract method 0x07d5ebf5.
//
// Solidity: function getDrones() view returns((uint256,string,int256,bool)[])
func (_Drone *DroneCallerSession) GetDrones() ([]DroneContractDrone, error) {
	return _Drone.Contract.GetDrones(&_Drone.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Drone *DroneCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Drone.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Drone *DroneSession) NextId() (*big.Int, error) {
	return _Drone.Contract.NextId(&_Drone.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Drone *DroneCallerSession) NextId() (*big.Int, error) {
	return _Drone.Contract.NextId(&_Drone.CallOpts)
}

// CreateDrone is a paid mutator transaction binding the contract method 0x276f5038.
//
// Solidity: function createDrone(string _model, int256 _zone) returns()
func (_Drone *DroneTransactor) CreateDrone(opts *bind.TransactOpts, _model string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.contract.Transact(opts, "createDrone", _model, _zone)
}

// CreateDrone is a paid mutator transaction binding the contract method 0x276f5038.
//
// Solidity: function createDrone(string _model, int256 _zone) returns()
func (_Drone *DroneSession) CreateDrone(_model string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.CreateDrone(&_Drone.TransactOpts, _model, _zone)
}

// CreateDrone is a paid mutator transaction binding the contract method 0x276f5038.
//
// Solidity: function createDrone(string _model, int256 _zone) returns()
func (_Drone *DroneTransactorSession) CreateDrone(_model string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.CreateDrone(&_Drone.TransactOpts, _model, _zone)
}

// DeleteDrone is a paid mutator transaction binding the contract method 0xed88eb6d.
//
// Solidity: function deleteDrone(uint256 _id) returns()
func (_Drone *DroneTransactor) DeleteDrone(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Drone.contract.Transact(opts, "deleteDrone", _id)
}

// DeleteDrone is a paid mutator transaction binding the contract method 0xed88eb6d.
//
// Solidity: function deleteDrone(uint256 _id) returns()
func (_Drone *DroneSession) DeleteDrone(_id *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.DeleteDrone(&_Drone.TransactOpts, _id)
}

// DeleteDrone is a paid mutator transaction binding the contract method 0xed88eb6d.
//
// Solidity: function deleteDrone(uint256 _id) returns()
func (_Drone *DroneTransactorSession) DeleteDrone(_id *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.DeleteDrone(&_Drone.TransactOpts, _id)
}

// UpdateDrone is a paid mutator transaction binding the contract method 0xd7c7a2ba.
//
// Solidity: function updateDrone(uint256 _id, string _model, int256 _zone) returns()
func (_Drone *DroneTransactor) UpdateDrone(opts *bind.TransactOpts, _id *big.Int, _model string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.contract.Transact(opts, "updateDrone", _id, _model, _zone)
}

// UpdateDrone is a paid mutator transaction binding the contract method 0xd7c7a2ba.
//
// Solidity: function updateDrone(uint256 _id, string _model, int256 _zone) returns()
func (_Drone *DroneSession) UpdateDrone(_id *big.Int, _model string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.UpdateDrone(&_Drone.TransactOpts, _id, _model, _zone)
}

// UpdateDrone is a paid mutator transaction binding the contract method 0xd7c7a2ba.
//
// Solidity: function updateDrone(uint256 _id, string _model, int256 _zone) returns()
func (_Drone *DroneTransactorSession) UpdateDrone(_id *big.Int, _model string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.UpdateDrone(&_Drone.TransactOpts, _id, _model, _zone)
}

// DroneDroneCreatedIterator is returned from FilterDroneCreated and is used to iterate over the raw logs and unpacked data for DroneCreated events raised by the Drone contract.
type DroneDroneCreatedIterator struct {
	Event *DroneDroneCreated // Event containing the contract specifics and raw log

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
func (it *DroneDroneCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DroneDroneCreated)
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
		it.Event = new(DroneDroneCreated)
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
func (it *DroneDroneCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DroneDroneCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DroneDroneCreated represents a DroneCreated event raised by the Drone contract.
type DroneDroneCreated struct {
	Id    *big.Int
	Model string
	Zone  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDroneCreated is a free log retrieval operation binding the contract event 0x6cc4ed91f20ea27eda237a37a11e00cb1e9e8cc036186b136f63b5dd811b64b5.
//
// Solidity: event DroneCreated(uint256 id, string model, int256 zone)
func (_Drone *DroneFilterer) FilterDroneCreated(opts *bind.FilterOpts) (*DroneDroneCreatedIterator, error) {

	logs, sub, err := _Drone.contract.FilterLogs(opts, "DroneCreated")
	if err != nil {
		return nil, err
	}
	return &DroneDroneCreatedIterator{contract: _Drone.contract, event: "DroneCreated", logs: logs, sub: sub}, nil
}

// WatchDroneCreated is a free log subscription operation binding the contract event 0x6cc4ed91f20ea27eda237a37a11e00cb1e9e8cc036186b136f63b5dd811b64b5.
//
// Solidity: event DroneCreated(uint256 id, string model, int256 zone)
func (_Drone *DroneFilterer) WatchDroneCreated(opts *bind.WatchOpts, sink chan<- *DroneDroneCreated) (event.Subscription, error) {

	logs, sub, err := _Drone.contract.WatchLogs(opts, "DroneCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DroneDroneCreated)
				if err := _Drone.contract.UnpackLog(event, "DroneCreated", log); err != nil {
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

// ParseDroneCreated is a log parse operation binding the contract event 0x6cc4ed91f20ea27eda237a37a11e00cb1e9e8cc036186b136f63b5dd811b64b5.
//
// Solidity: event DroneCreated(uint256 id, string model, int256 zone)
func (_Drone *DroneFilterer) ParseDroneCreated(log types.Log) (*DroneDroneCreated, error) {
	event := new(DroneDroneCreated)
	if err := _Drone.contract.UnpackLog(event, "DroneCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DroneDroneDeletedIterator is returned from FilterDroneDeleted and is used to iterate over the raw logs and unpacked data for DroneDeleted events raised by the Drone contract.
type DroneDroneDeletedIterator struct {
	Event *DroneDroneDeleted // Event containing the contract specifics and raw log

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
func (it *DroneDroneDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DroneDroneDeleted)
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
		it.Event = new(DroneDroneDeleted)
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
func (it *DroneDroneDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DroneDroneDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DroneDroneDeleted represents a DroneDeleted event raised by the Drone contract.
type DroneDroneDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDroneDeleted is a free log retrieval operation binding the contract event 0xa398f31db5e7963635c8e05cb7bdd246a84574c0b1671840aa00e069a9794b12.
//
// Solidity: event DroneDeleted(uint256 id)
func (_Drone *DroneFilterer) FilterDroneDeleted(opts *bind.FilterOpts) (*DroneDroneDeletedIterator, error) {

	logs, sub, err := _Drone.contract.FilterLogs(opts, "DroneDeleted")
	if err != nil {
		return nil, err
	}
	return &DroneDroneDeletedIterator{contract: _Drone.contract, event: "DroneDeleted", logs: logs, sub: sub}, nil
}

// WatchDroneDeleted is a free log subscription operation binding the contract event 0xa398f31db5e7963635c8e05cb7bdd246a84574c0b1671840aa00e069a9794b12.
//
// Solidity: event DroneDeleted(uint256 id)
func (_Drone *DroneFilterer) WatchDroneDeleted(opts *bind.WatchOpts, sink chan<- *DroneDroneDeleted) (event.Subscription, error) {

	logs, sub, err := _Drone.contract.WatchLogs(opts, "DroneDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DroneDroneDeleted)
				if err := _Drone.contract.UnpackLog(event, "DroneDeleted", log); err != nil {
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

// ParseDroneDeleted is a log parse operation binding the contract event 0xa398f31db5e7963635c8e05cb7bdd246a84574c0b1671840aa00e069a9794b12.
//
// Solidity: event DroneDeleted(uint256 id)
func (_Drone *DroneFilterer) ParseDroneDeleted(log types.Log) (*DroneDroneDeleted, error) {
	event := new(DroneDroneDeleted)
	if err := _Drone.contract.UnpackLog(event, "DroneDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DroneDroneUpdatedIterator is returned from FilterDroneUpdated and is used to iterate over the raw logs and unpacked data for DroneUpdated events raised by the Drone contract.
type DroneDroneUpdatedIterator struct {
	Event *DroneDroneUpdated // Event containing the contract specifics and raw log

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
func (it *DroneDroneUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DroneDroneUpdated)
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
		it.Event = new(DroneDroneUpdated)
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
func (it *DroneDroneUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DroneDroneUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DroneDroneUpdated represents a DroneUpdated event raised by the Drone contract.
type DroneDroneUpdated struct {
	Id    *big.Int
	Model string
	Zone  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDroneUpdated is a free log retrieval operation binding the contract event 0x84303538537fc260b1a75f003f2666ebee27ee671e3fde7c9f365c5557f0dfe9.
//
// Solidity: event DroneUpdated(uint256 id, string model, int256 zone)
func (_Drone *DroneFilterer) FilterDroneUpdated(opts *bind.FilterOpts) (*DroneDroneUpdatedIterator, error) {

	logs, sub, err := _Drone.contract.FilterLogs(opts, "DroneUpdated")
	if err != nil {
		return nil, err
	}
	return &DroneDroneUpdatedIterator{contract: _Drone.contract, event: "DroneUpdated", logs: logs, sub: sub}, nil
}

// WatchDroneUpdated is a free log subscription operation binding the contract event 0x84303538537fc260b1a75f003f2666ebee27ee671e3fde7c9f365c5557f0dfe9.
//
// Solidity: event DroneUpdated(uint256 id, string model, int256 zone)
func (_Drone *DroneFilterer) WatchDroneUpdated(opts *bind.WatchOpts, sink chan<- *DroneDroneUpdated) (event.Subscription, error) {

	logs, sub, err := _Drone.contract.WatchLogs(opts, "DroneUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DroneDroneUpdated)
				if err := _Drone.contract.UnpackLog(event, "DroneUpdated", log); err != nil {
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

// ParseDroneUpdated is a log parse operation binding the contract event 0x84303538537fc260b1a75f003f2666ebee27ee671e3fde7c9f365c5557f0dfe9.
//
// Solidity: event DroneUpdated(uint256 id, string model, int256 zone)
func (_Drone *DroneFilterer) ParseDroneUpdated(log types.Log) (*DroneDroneUpdated, error) {
	event := new(DroneDroneUpdated)
	if err := _Drone.contract.UnpackLog(event, "DroneUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
