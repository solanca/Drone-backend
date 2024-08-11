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
	Id        *big.Int
	ModelType string
	Zone      *big.Int
	Exists    bool
}

// DroneMetaData contains all meta data concerning the Drone contract.
var DroneMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"createDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"}],\"name\":\"DroneCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DroneDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"}],\"name\":\"DroneUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"updateDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drones\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getDrone\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structDroneContract.Drone\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDrones\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structDroneContract.Drone[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"getDronesByZone\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structDroneContract.Drone[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506115ba8061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c8063d7c7a2ba11610059578063d7c7a2ba14610115578063dbd1be7714610131578063ed88eb6d14610161578063f6e760981461017d57610086565b806307d5ebf51461008a578063276f5038146100a857806361b8ce8c146100c45780639341bc00146100e2575b5f80fd5b6100926101ad565b60405161009f9190610d5f565b60405180910390f35b6100c260048036038101906100bd9190610ee6565b6103d4565b005b6100cc6104c1565b6040516100d99190610f4f565b60405180910390f35b6100fc60048036038101906100f79190610f92565b6104c7565b60405161010c9493929190611023565b60405180910390f35b61012f600480360381019061012a919061106d565b610593565b005b61014b60048036038101906101469190610f92565b61068b565b6040516101589190611139565b60405180910390f35b61017b60048036038101906101769190610f92565b6107fc565b005b61019760048036038101906101929190611159565b6108e0565b6040516101a49190610d5f565b60405180910390f35b60605f805b5f8054905081101561020e575f81815481106101d1576101d0611184565b5b905f5260205f2090600402016003015f9054906101000a900460ff16156102015781806101fd906111de565b9250505b80806001019150506101b2565b505f8167ffffffffffffffff81111561022a57610229610d98565b5b60405190808252806020026020018201604052801561026357816020015b610250610b63565b8152602001906001900390816102485790505b5090505f805b5f805490508110156103ca575f818154811061028857610287611184565b5b905f5260205f2090600402016003015f9054906101000a900460ff16156103bd575f81815481106102bc576102bb611184565b5b905f5260205f2090600402016040518060800160405290815f82015481526020016001820180546102ec90611252565b80601f016020809104026020016040519081016040528092919081815260200182805461031890611252565b80156103635780601f1061033a57610100808354040283529160200191610363565b820191905f5260205f20905b81548152906001019060200180831161034657829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff1615151515815250508383815181106103a3576103a2611184565b5b602002602001018190525081806103b9906111de565b9250505b8080600101915050610269565b5081935050505090565b5f6040518060800160405280600154815260200184815260200183815260200160011515815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f0155602082015181600101908161043e919061141f565b50604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050506001547f6cc4ed91f20ea27eda237a37a11e00cb1e9e8cc036186b136f63b5dd811b64b5838360405161049e9291906114ee565b60405180910390a260015f8154809291906104b8906111de565b91905055505050565b60015481565b5f81815481106104d5575f80fd5b905f5260205f2090600402015f91509050805f0154908060010180546104fa90611252565b80601f016020809104026020016040519081016040528092919081815260200182805461052690611252565b80156105715780601f1061054857610100808354040283529160200191610571565b820191905f5260205f20905b81548152906001019060200180831161055457829003601f168201915b505050505090806002015490806003015f9054906101000a900460ff16905084565b600154831080156105cf57505f83815481106105b2576105b1611184565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b61060e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060590611566565b60405180910390fd5b5f80848154811061062257610621611184565b5b905f5260205f209060040201905082816001019081610641919061141f565b50818160020181905550837f84303538537fc260b1a75f003f2666ebee27ee671e3fde7c9f365c5557f0dfe9848460405161067d9291906114ee565b60405180910390a250505050565b610693610b63565b600154821080156106cf57505f82815481106106b2576106b1611184565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b61070e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070590611566565b60405180910390fd5b5f828154811061072157610720611184565b5b905f5260205f2090600402016040518060800160405290815f820154815260200160018201805461075190611252565b80601f016020809104026020016040519081016040528092919081815260200182805461077d90611252565b80156107c85780601f1061079f576101008083540402835291602001916107c8565b820191905f5260205f20905b8154815290600101906020018083116107ab57829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff1615151515815250509050919050565b6001548110801561083857505f818154811061081b5761081a611184565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b610877576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086e90611566565b60405180910390fd5b5f80828154811061088b5761088a611184565b5b905f5260205f2090600402016003015f6101000a81548160ff021916908315150217905550807fa398f31db5e7963635c8e05cb7bdd246a84574c0b1671840aa00e069a9794b1260405160405180910390a250565b60605f805b5f8054905081101561096e575f818154811061090457610903611184565b5b905f5260205f2090600402016003015f9054906101000a900460ff16801561094d5750835f828154811061093b5761093a611184565b5b905f5260205f20906004020160020154145b1561096157818061095d906111de565b9250505b80806001019150506108e5565b505f8167ffffffffffffffff81111561098a57610989610d98565b5b6040519080825280602002602001820160405280156109c357816020015b6109b0610b63565b8152602001906001900390816109a85790505b5090505f805b5f80549050811015610b57575f81815481106109e8576109e7611184565b5b905f5260205f2090600402016003015f9054906101000a900460ff168015610a315750855f8281548110610a1f57610a1e611184565b5b905f5260205f20906004020160020154145b15610b4a575f8181548110610a4957610a48611184565b5b905f5260205f2090600402016040518060800160405290815f8201548152602001600182018054610a7990611252565b80601f0160208091040260200160405190810160405280929190818152602001828054610aa590611252565b8015610af05780601f10610ac757610100808354040283529160200191610af0565b820191905f5260205f20905b815481529060010190602001808311610ad357829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff161515151581525050838381518110610b3057610b2f611184565b5b60200260200101819052508180610b46906111de565b9250505b80806001019150506109c9565b50819350505050919050565b60405180608001604052805f8152602001606081526020015f81526020015f151581525090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b610bc581610bb3565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610c0d82610bcb565b610c178185610bd5565b9350610c27818560208601610be5565b610c3081610bf3565b840191505092915050565b5f819050919050565b610c4d81610c3b565b82525050565b5f8115159050919050565b610c6781610c53565b82525050565b5f608083015f830151610c825f860182610bbc565b5060208301518482036020860152610c9a8282610c03565b9150506040830151610caf6040860182610c44565b506060830151610cc26060860182610c5e565b508091505092915050565b5f610cd88383610c6d565b905092915050565b5f602082019050919050565b5f610cf682610b8a565b610d008185610b94565b935083602082028501610d1285610ba4565b805f5b85811015610d4d5784840389528151610d2e8582610ccd565b9450610d3983610ce0565b925060208a01995050600181019050610d15565b50829750879550505050505092915050565b5f6020820190508181035f830152610d778184610cec565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610dce82610bf3565b810181811067ffffffffffffffff82111715610ded57610dec610d98565b5b80604052505050565b5f610dff610d7f565b9050610e0b8282610dc5565b919050565b5f67ffffffffffffffff821115610e2a57610e29610d98565b5b610e3382610bf3565b9050602081019050919050565b828183375f83830152505050565b5f610e60610e5b84610e10565b610df6565b905082815260208101848484011115610e7c57610e7b610d94565b5b610e87848285610e40565b509392505050565b5f82601f830112610ea357610ea2610d90565b5b8135610eb3848260208601610e4e565b91505092915050565b610ec581610c3b565b8114610ecf575f80fd5b50565b5f81359050610ee081610ebc565b92915050565b5f8060408385031215610efc57610efb610d88565b5b5f83013567ffffffffffffffff811115610f1957610f18610d8c565b5b610f2585828601610e8f565b9250506020610f3685828601610ed2565b9150509250929050565b610f4981610bb3565b82525050565b5f602082019050610f625f830184610f40565b92915050565b610f7181610bb3565b8114610f7b575f80fd5b50565b5f81359050610f8c81610f68565b92915050565b5f60208284031215610fa757610fa6610d88565b5b5f610fb484828501610f7e565b91505092915050565b5f82825260208201905092915050565b5f610fd782610bcb565b610fe18185610fbd565b9350610ff1818560208601610be5565b610ffa81610bf3565b840191505092915050565b61100e81610c3b565b82525050565b61101d81610c53565b82525050565b5f6080820190506110365f830187610f40565b81810360208301526110488186610fcd565b90506110576040830185611005565b6110646060830184611014565b95945050505050565b5f805f6060848603121561108457611083610d88565b5b5f61109186828701610f7e565b935050602084013567ffffffffffffffff8111156110b2576110b1610d8c565b5b6110be86828701610e8f565b92505060406110cf86828701610ed2565b9150509250925092565b5f608083015f8301516110ee5f860182610bbc565b50602083015184820360208601526111068282610c03565b915050604083015161111b6040860182610c44565b50606083015161112e6060860182610c5e565b508091505092915050565b5f6020820190508181035f83015261115181846110d9565b905092915050565b5f6020828403121561116e5761116d610d88565b5b5f61117b84828501610ed2565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111e882610bb3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361121a576112196111b1565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061126957607f821691505b60208210810361127c5761127b611225565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026112de7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826112a3565b6112e886836112a3565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61132361131e61131984610bb3565b611300565b610bb3565b9050919050565b5f819050919050565b61133c83611309565b6113506113488261132a565b8484546112af565b825550505050565b5f90565b611364611358565b61136f818484611333565b505050565b5b81811015611392576113875f8261135c565b600181019050611375565b5050565b601f8211156113d7576113a881611282565b6113b184611294565b810160208510156113c0578190505b6113d46113cc85611294565b830182611374565b50505b505050565b5f82821c905092915050565b5f6113f75f19846008026113dc565b1980831691505092915050565b5f61140f83836113e8565b9150826002028217905092915050565b61142882610bcb565b67ffffffffffffffff81111561144157611440610d98565b5b61144b8254611252565b611456828285611396565b5f60209050601f831160018114611487575f8415611475578287015190505b61147f8582611404565b8655506114e6565b601f19841661149586611282565b5f5b828110156114bc57848901518255600182019150602085019450602081019050611497565b868310156114d957848901516114d5601f8916826113e8565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f8301526115068185610fcd565b90506115156020830184611005565b9392505050565b7f44726f6e6520646f6573206e6f742065786973740000000000000000000000005f82015250565b5f611550601483610fbd565b915061155b8261151c565b602082019050919050565b5f6020820190508181035f83015261157d81611544565b905091905056fea264697066735822122054b0dce1db9bd05b653c26f6096ea113a8b1b4fc9f0a3dee4b9dd7d0065cec8164736f6c634300081a0033",
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
// Solidity: function drones(uint256 ) view returns(uint256 id, string model_type, int256 zone, bool exists)
func (_Drone *DroneCaller) Drones(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	ModelType string
	Zone      *big.Int
	Exists    bool
}, error) {
	var out []interface{}
	err := _Drone.contract.Call(opts, &out, "drones", arg0)

	outstruct := new(struct {
		Id        *big.Int
		ModelType string
		Zone      *big.Int
		Exists    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ModelType = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Zone = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Drones is a free data retrieval call binding the contract method 0x9341bc00.
//
// Solidity: function drones(uint256 ) view returns(uint256 id, string model_type, int256 zone, bool exists)
func (_Drone *DroneSession) Drones(arg0 *big.Int) (struct {
	Id        *big.Int
	ModelType string
	Zone      *big.Int
	Exists    bool
}, error) {
	return _Drone.Contract.Drones(&_Drone.CallOpts, arg0)
}

// Drones is a free data retrieval call binding the contract method 0x9341bc00.
//
// Solidity: function drones(uint256 ) view returns(uint256 id, string model_type, int256 zone, bool exists)
func (_Drone *DroneCallerSession) Drones(arg0 *big.Int) (struct {
	Id        *big.Int
	ModelType string
	Zone      *big.Int
	Exists    bool
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

// GetDronesByZone is a free data retrieval call binding the contract method 0xf6e76098.
//
// Solidity: function getDronesByZone(int256 _zone) view returns((uint256,string,int256,bool)[])
func (_Drone *DroneCaller) GetDronesByZone(opts *bind.CallOpts, _zone *big.Int) ([]DroneContractDrone, error) {
	var out []interface{}
	err := _Drone.contract.Call(opts, &out, "getDronesByZone", _zone)

	if err != nil {
		return *new([]DroneContractDrone), err
	}

	out0 := *abi.ConvertType(out[0], new([]DroneContractDrone)).(*[]DroneContractDrone)

	return out0, err

}

// GetDronesByZone is a free data retrieval call binding the contract method 0xf6e76098.
//
// Solidity: function getDronesByZone(int256 _zone) view returns((uint256,string,int256,bool)[])
func (_Drone *DroneSession) GetDronesByZone(_zone *big.Int) ([]DroneContractDrone, error) {
	return _Drone.Contract.GetDronesByZone(&_Drone.CallOpts, _zone)
}

// GetDronesByZone is a free data retrieval call binding the contract method 0xf6e76098.
//
// Solidity: function getDronesByZone(int256 _zone) view returns((uint256,string,int256,bool)[])
func (_Drone *DroneCallerSession) GetDronesByZone(_zone *big.Int) ([]DroneContractDrone, error) {
	return _Drone.Contract.GetDronesByZone(&_Drone.CallOpts, _zone)
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
// Solidity: function createDrone(string _model_type, int256 _zone) returns()
func (_Drone *DroneTransactor) CreateDrone(opts *bind.TransactOpts, _model_type string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.contract.Transact(opts, "createDrone", _model_type, _zone)
}

// CreateDrone is a paid mutator transaction binding the contract method 0x276f5038.
//
// Solidity: function createDrone(string _model_type, int256 _zone) returns()
func (_Drone *DroneSession) CreateDrone(_model_type string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.CreateDrone(&_Drone.TransactOpts, _model_type, _zone)
}

// CreateDrone is a paid mutator transaction binding the contract method 0x276f5038.
//
// Solidity: function createDrone(string _model_type, int256 _zone) returns()
func (_Drone *DroneTransactorSession) CreateDrone(_model_type string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.CreateDrone(&_Drone.TransactOpts, _model_type, _zone)
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
// Solidity: function updateDrone(uint256 _id, string _model_type, int256 _zone) returns()
func (_Drone *DroneTransactor) UpdateDrone(opts *bind.TransactOpts, _id *big.Int, _model_type string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.contract.Transact(opts, "updateDrone", _id, _model_type, _zone)
}

// UpdateDrone is a paid mutator transaction binding the contract method 0xd7c7a2ba.
//
// Solidity: function updateDrone(uint256 _id, string _model_type, int256 _zone) returns()
func (_Drone *DroneSession) UpdateDrone(_id *big.Int, _model_type string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.UpdateDrone(&_Drone.TransactOpts, _id, _model_type, _zone)
}

// UpdateDrone is a paid mutator transaction binding the contract method 0xd7c7a2ba.
//
// Solidity: function updateDrone(uint256 _id, string _model_type, int256 _zone) returns()
func (_Drone *DroneTransactorSession) UpdateDrone(_id *big.Int, _model_type string, _zone *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.UpdateDrone(&_Drone.TransactOpts, _id, _model_type, _zone)
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
	Id        *big.Int
	ModelType string
	Zone      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDroneCreated is a free log retrieval operation binding the contract event 0x6cc4ed91f20ea27eda237a37a11e00cb1e9e8cc036186b136f63b5dd811b64b5.
//
// Solidity: event DroneCreated(uint256 indexed id, string model_type, int256 zone)
func (_Drone *DroneFilterer) FilterDroneCreated(opts *bind.FilterOpts, id []*big.Int) (*DroneDroneCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Drone.contract.FilterLogs(opts, "DroneCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &DroneDroneCreatedIterator{contract: _Drone.contract, event: "DroneCreated", logs: logs, sub: sub}, nil
}

// WatchDroneCreated is a free log subscription operation binding the contract event 0x6cc4ed91f20ea27eda237a37a11e00cb1e9e8cc036186b136f63b5dd811b64b5.
//
// Solidity: event DroneCreated(uint256 indexed id, string model_type, int256 zone)
func (_Drone *DroneFilterer) WatchDroneCreated(opts *bind.WatchOpts, sink chan<- *DroneDroneCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Drone.contract.WatchLogs(opts, "DroneCreated", idRule)
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
// Solidity: event DroneCreated(uint256 indexed id, string model_type, int256 zone)
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
// Solidity: event DroneDeleted(uint256 indexed id)
func (_Drone *DroneFilterer) FilterDroneDeleted(opts *bind.FilterOpts, id []*big.Int) (*DroneDroneDeletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Drone.contract.FilterLogs(opts, "DroneDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return &DroneDroneDeletedIterator{contract: _Drone.contract, event: "DroneDeleted", logs: logs, sub: sub}, nil
}

// WatchDroneDeleted is a free log subscription operation binding the contract event 0xa398f31db5e7963635c8e05cb7bdd246a84574c0b1671840aa00e069a9794b12.
//
// Solidity: event DroneDeleted(uint256 indexed id)
func (_Drone *DroneFilterer) WatchDroneDeleted(opts *bind.WatchOpts, sink chan<- *DroneDroneDeleted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Drone.contract.WatchLogs(opts, "DroneDeleted", idRule)
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
// Solidity: event DroneDeleted(uint256 indexed id)
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
	Id        *big.Int
	ModelType string
	Zone      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDroneUpdated is a free log retrieval operation binding the contract event 0x84303538537fc260b1a75f003f2666ebee27ee671e3fde7c9f365c5557f0dfe9.
//
// Solidity: event DroneUpdated(uint256 indexed id, string model_type, int256 zone)
func (_Drone *DroneFilterer) FilterDroneUpdated(opts *bind.FilterOpts, id []*big.Int) (*DroneDroneUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Drone.contract.FilterLogs(opts, "DroneUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &DroneDroneUpdatedIterator{contract: _Drone.contract, event: "DroneUpdated", logs: logs, sub: sub}, nil
}

// WatchDroneUpdated is a free log subscription operation binding the contract event 0x84303538537fc260b1a75f003f2666ebee27ee671e3fde7c9f365c5557f0dfe9.
//
// Solidity: event DroneUpdated(uint256 indexed id, string model_type, int256 zone)
func (_Drone *DroneFilterer) WatchDroneUpdated(opts *bind.WatchOpts, sink chan<- *DroneDroneUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Drone.contract.WatchLogs(opts, "DroneUpdated", idRule)
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
// Solidity: event DroneUpdated(uint256 indexed id, string model_type, int256 zone)
func (_Drone *DroneFilterer) ParseDroneUpdated(log types.Log) (*DroneDroneUpdated, error) {
	event := new(DroneDroneUpdated)
	if err := _Drone.contract.UnpackLog(event, "DroneUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
