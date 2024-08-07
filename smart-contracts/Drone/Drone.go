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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"loggerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"}],\"name\":\"DroneCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DroneDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"}],\"name\":\"DroneUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"createDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drones\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getDrone\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structDroneContract.Drone\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDrones\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structDroneContract.Drone[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"getDronesByZone\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structDroneContract.Drone[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logger\",\"outputs\":[{\"internalType\":\"contractLoggingContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_model_type\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"updateDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161200c38038061200c833981810160405281019061003191906100d5565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610100565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a48261007b565b9050919050565b6100b48161009a565b81146100be575f80fd5b50565b5f815190506100cf816100ab565b92915050565b5f602082840312156100ea576100e9610077565b5b5f6100f7848285016100c1565b91505092915050565b611eff8061010d5f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c8063d7c7a2ba11610064578063d7c7a2ba14610120578063dbd1be771461013c578063ed88eb6d1461016c578063f24ccbfe14610188578063f6e76098146101a657610091565b806307d5ebf514610095578063276f5038146100b357806361b8ce8c146100cf5780639341bc00146100ed575b5f80fd5b61009d6101d6565b6040516100aa919061125d565b60405180910390f35b6100cd60048036038101906100c891906113e4565b6103fd565b005b6100d76104ea565b6040516100e4919061144d565b60405180910390f35b61010760048036038101906101029190611490565b6104f0565b6040516101179493929190611521565b60405180910390f35b61013a6004803603810190610135919061156b565b6105bc565b005b61015660048036038101906101519190611490565b6106b4565b6040516101639190611637565b60405180910390f35b61018660048036038101906101819190611490565b610927565b005b610190610a0b565b60405161019d91906116d1565b60405180910390f35b6101c060048036038101906101bb91906116ea565b610a30565b6040516101cd919061125d565b60405180910390f35b60605f805b5f80549050811015610237575f81815481106101fa576101f9611715565b5b905f5260205f2090600402016003015f9054906101000a900460ff161561022a5781806102269061176f565b9250505b80806001019150506101db565b505f8167ffffffffffffffff81111561025357610252611296565b5b60405190808252806020026020018201604052801561028c57816020015b610279611061565b8152602001906001900390816102715790505b5090505f805b5f805490508110156103f3575f81815481106102b1576102b0611715565b5b905f5260205f2090600402016003015f9054906101000a900460ff16156103e6575f81815481106102e5576102e4611715565b5b905f5260205f2090600402016040518060800160405290815f8201548152602001600182018054610315906117e3565b80601f0160208091040260200160405190810160405280929190818152602001828054610341906117e3565b801561038c5780601f106103635761010080835404028352916020019161038c565b820191905f5260205f20905b81548152906001019060200180831161036f57829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff1615151515815250508383815181106103cc576103cb611715565b5b602002602001018190525081806103e29061176f565b9250505b8080600101915050610292565b5081935050505090565b5f6040518060800160405280600154815260200184815260200183815260200160011515815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f0155602082015181600101908161046791906119a7565b50604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050506001547f6cc4ed91f20ea27eda237a37a11e00cb1e9e8cc036186b136f63b5dd811b64b583836040516104c7929190611a76565b60405180910390a260015f8154809291906104e19061176f565b91905055505050565b60015481565b5f81815481106104fe575f80fd5b905f5260205f2090600402015f91509050805f015490806001018054610523906117e3565b80601f016020809104026020016040519081016040528092919081815260200182805461054f906117e3565b801561059a5780601f106105715761010080835404028352916020019161059a565b820191905f5260205f20905b81548152906001019060200180831161057d57829003601f168201915b505050505090806002015490806003015f9054906101000a900460ff16905084565b600154831080156105f857505f83815481106105db576105da611715565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b610637576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062e90611aee565b60405180910390fd5b5f80848154811061064b5761064a611715565b5b905f5260205f20906004020190508281600101908161066a91906119a7565b50818160020181905550837f84303538537fc260b1a75f003f2666ebee27ee671e3fde7c9f365c5557f0dfe984846040516106a6929190611a76565b60405180910390a250505050565b6106bc611061565b600154821080156106f857505f82815481106106db576106da611715565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b610737576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072e90611aee565b60405180910390fd5b5f61074183610cb3565b5f848154811061075457610753611715565b5b905f5260205f20906004020160010161078e5f868154811061077957610778611715565b5b905f5260205f20906004020160020154610e31565b6040516020016107a093929190611ca4565b604051602081830303815290604052905060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166397f8abd8826040518263ffffffff1660e01b815260040161080b9190611d3f565b5f604051808303815f87803b158015610822575f80fd5b505af1158015610834573d5f803e3d5ffd5b505050505f838154811061084b5761084a611715565b5b905f5260205f2090600402016040518060800160405290815f820154815260200160018201805461087b906117e3565b80601f01602080910402602001604051908101604052809291908181526020018280546108a7906117e3565b80156108f25780601f106108c9576101008083540402835291602001916108f2565b820191905f5260205f20905b8154815290600101906020018083116108d557829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff161515151581525050915050919050565b6001548110801561096357505f818154811061094657610945611715565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b6109a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099990611aee565b60405180910390fd5b5f8082815481106109b6576109b5611715565b5b905f5260205f2090600402016003015f6101000a81548160ff021916908315150217905550807fa398f31db5e7963635c8e05cb7bdd246a84574c0b1671840aa00e069a9794b1260405160405180910390a250565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60605f805b5f80549050811015610abe575f8181548110610a5457610a53611715565b5b905f5260205f2090600402016003015f9054906101000a900460ff168015610a9d5750835f8281548110610a8b57610a8a611715565b5b905f5260205f20906004020160020154145b15610ab1578180610aad9061176f565b9250505b8080600101915050610a35565b505f8167ffffffffffffffff811115610ada57610ad9611296565b5b604051908082528060200260200182016040528015610b1357816020015b610b00611061565b815260200190600190039081610af85790505b5090505f805b5f80549050811015610ca7575f8181548110610b3857610b37611715565b5b905f5260205f2090600402016003015f9054906101000a900460ff168015610b815750855f8281548110610b6f57610b6e611715565b5b905f5260205f20906004020160020154145b15610c9a575f8181548110610b9957610b98611715565b5b905f5260205f2090600402016040518060800160405290815f8201548152602001600182018054610bc9906117e3565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf5906117e3565b8015610c405780601f10610c1757610100808354040283529160200191610c40565b820191905f5260205f20905b815481529060010190602001808311610c2357829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff161515151581525050838381518110610c8057610c7f611715565b5b60200260200101819052508180610c969061176f565b9250505b8080600101915050610b19565b50819350505050919050565b60605f8203610cf9576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610e2c565b5f8290505f5b5f8214610d28578080610d119061176f565b915050600a82610d219190611d9f565b9150610cff565b5f8167ffffffffffffffff811115610d4357610d42611296565b5b6040519080825280601f01601f191660200182016040528015610d755781602001600182028036833780820191505090505b5090505f8290505b5f8614610e2457600181610d919190611dcf565b90505f600a8088610da29190611d9f565b610dac9190611e02565b87610db79190611dcf565b6030610dc39190611e4f565b90505f8160f81b905080848481518110610de057610ddf611715565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a88610e1b9190611d9f565b97505050610d7d565b819450505050505b919050565b60605f8203610e77576040518060400160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815250905061105c565b5f80831290505f81610e895783610e94565b83610e9390611e83565b5b90505f5b5f8214610ec1578080610eaa9061176f565b915050600a82610eba9190611d9f565b9150610e98565b8215610ed6578080610ed29061176f565b9150505b5f8167ffffffffffffffff811115610ef157610ef0611296565b5b6040519080825280601f01601f191660200182016040528015610f235781602001600182028036833780820191505090505b5090505f82905084610f355786610f40565b86610f3f90611e83565b5b93505b5f8414610fea57600181610f579190611dcf565b90505f600a8086610f689190611d9f565b610f729190611e02565b85610f7d9190611dcf565b6030610f899190611e4f565b90505f8160f81b905080848481518110610fa657610fa5611715565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a86610fe19190611d9f565b95505050610f43565b8415611053577f2d00000000000000000000000000000000000000000000000000000000000000825f8151811061102457611023611715565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505b81955050505050505b919050565b60405180608001604052805f8152602001606081526020015f81526020015f151581525090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b6110c3816110b1565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61110b826110c9565b61111581856110d3565b93506111258185602086016110e3565b61112e816110f1565b840191505092915050565b5f819050919050565b61114b81611139565b82525050565b5f8115159050919050565b61116581611151565b82525050565b5f608083015f8301516111805f8601826110ba565b50602083015184820360208601526111988282611101565b91505060408301516111ad6040860182611142565b5060608301516111c0606086018261115c565b508091505092915050565b5f6111d6838361116b565b905092915050565b5f602082019050919050565b5f6111f482611088565b6111fe8185611092565b935083602082028501611210856110a2565b805f5b8581101561124b578484038952815161122c85826111cb565b9450611237836111de565b925060208a01995050600181019050611213565b50829750879550505050505092915050565b5f6020820190508181035f83015261127581846111ea565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6112cc826110f1565b810181811067ffffffffffffffff821117156112eb576112ea611296565b5b80604052505050565b5f6112fd61127d565b905061130982826112c3565b919050565b5f67ffffffffffffffff82111561132857611327611296565b5b611331826110f1565b9050602081019050919050565b828183375f83830152505050565b5f61135e6113598461130e565b6112f4565b90508281526020810184848401111561137a57611379611292565b5b61138584828561133e565b509392505050565b5f82601f8301126113a1576113a061128e565b5b81356113b184826020860161134c565b91505092915050565b6113c381611139565b81146113cd575f80fd5b50565b5f813590506113de816113ba565b92915050565b5f80604083850312156113fa576113f9611286565b5b5f83013567ffffffffffffffff8111156114175761141661128a565b5b6114238582860161138d565b9250506020611434858286016113d0565b9150509250929050565b611447816110b1565b82525050565b5f6020820190506114605f83018461143e565b92915050565b61146f816110b1565b8114611479575f80fd5b50565b5f8135905061148a81611466565b92915050565b5f602082840312156114a5576114a4611286565b5b5f6114b28482850161147c565b91505092915050565b5f82825260208201905092915050565b5f6114d5826110c9565b6114df81856114bb565b93506114ef8185602086016110e3565b6114f8816110f1565b840191505092915050565b61150c81611139565b82525050565b61151b81611151565b82525050565b5f6080820190506115345f83018761143e565b818103602083015261154681866114cb565b90506115556040830185611503565b6115626060830184611512565b95945050505050565b5f805f6060848603121561158257611581611286565b5b5f61158f8682870161147c565b935050602084013567ffffffffffffffff8111156115b0576115af61128a565b5b6115bc8682870161138d565b92505060406115cd868287016113d0565b9150509250925092565b5f608083015f8301516115ec5f8601826110ba565b50602083015184820360208601526116048282611101565b91505060408301516116196040860182611142565b50606083015161162c606086018261115c565b508091505092915050565b5f6020820190508181035f83015261164f81846115d7565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61169961169461168f84611657565b611676565b611657565b9050919050565b5f6116aa8261167f565b9050919050565b5f6116bb826116a0565b9050919050565b6116cb816116b1565b82525050565b5f6020820190506116e45f8301846116c2565b92915050565b5f602082840312156116ff576116fe611286565b5b5f61170c848285016113d0565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611779826110b1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117ab576117aa611742565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806117fa57607f821691505b60208210810361180d5761180c6117b6565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261186f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611834565b6118798683611834565b95508019841693508086168417925050509392505050565b5f6118ab6118a66118a1846110b1565b611676565b6110b1565b9050919050565b5f819050919050565b6118c483611891565b6118d86118d0826118b2565b848454611840565b825550505050565b5f90565b6118ec6118e0565b6118f78184846118bb565b505050565b5b8181101561191a5761190f5f826118e4565b6001810190506118fd565b5050565b601f82111561195f5761193081611813565b61193984611825565b81016020851015611948578190505b61195c61195485611825565b8301826118fc565b50505b505050565b5f82821c905092915050565b5f61197f5f1984600802611964565b1980831691505092915050565b5f6119978383611970565b9150826002028217905092915050565b6119b0826110c9565b67ffffffffffffffff8111156119c9576119c8611296565b5b6119d382546117e3565b6119de82828561191e565b5f60209050601f831160018114611a0f575f84156119fd578287015190505b611a07858261198c565b865550611a6e565b601f198416611a1d86611813565b5f5b82811015611a4457848901518255600182019150602085019450602081019050611a1f565b86831015611a615784890151611a5d601f891682611970565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f830152611a8e81856114cb565b9050611a9d6020830184611503565b9392505050565b7f44726f6e6520646f6573206e6f742065786973740000000000000000000000005f82015250565b5f611ad86014836114bb565b9150611ae382611aa4565b602082019050919050565b5f6020820190508181035f830152611b0581611acc565b9050919050565b5f81905092915050565b7f49443a20000000000000000000000000000000000000000000000000000000005f82015250565b5f611b4a600483611b0c565b9150611b5582611b16565b600482019050919050565b5f611b6a826110c9565b611b748185611b0c565b9350611b848185602086016110e3565b80840191505092915050565b7f2c204d6f64656c5f547970653a200000000000000000000000000000000000005f82015250565b5f611bc4600e83611b0c565b9150611bcf82611b90565b600e82019050919050565b5f8154611be6816117e3565b611bf08186611b0c565b9450600182165f8114611c0a5760018114611c1f57611c51565b60ff1983168652811515820286019350611c51565b611c2885611813565b5f5b83811015611c4957815481890152600182019150602081019050611c2a565b838801955050505b50505092915050565b7f2c205a6f6e653a200000000000000000000000000000000000000000000000005f82015250565b5f611c8e600883611b0c565b9150611c9982611c5a565b600882019050919050565b5f611cae82611b3e565b9150611cba8286611b60565b9150611cc582611bb8565b9150611cd18285611bda565b9150611cdc82611c82565b9150611ce88284611b60565b9150819050949350505050565b7f50495000000000000000000000000000000000000000000000000000000000005f82015250565b5f611d296003836114bb565b9150611d3482611cf5565b602082019050919050565b5f6040820190508181035f830152611d5681611d1d565b90508181036020830152611d6a81846114cb565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611da9826110b1565b9150611db4836110b1565b925082611dc457611dc3611d72565b5b828204905092915050565b5f611dd9826110b1565b9150611de4836110b1565b9250828203905081811115611dfc57611dfb611742565b5b92915050565b5f611e0c826110b1565b9150611e17836110b1565b9250828202611e25816110b1565b91508282048414831517611e3c57611e3b611742565b5b5092915050565b5f60ff82169050919050565b5f611e5982611e43565b9150611e6483611e43565b9250828201905060ff811115611e7d57611e7c611742565b5b92915050565b5f611e8d82611139565b91507f80000000000000000000000000000000000000000000000000000000000000008203611ebf57611ebe611742565b5b815f03905091905056fea2646970667358221220cc4acb1446da5150ca5454fd6b127ab38df911d8688f6381d1c6a913487e9ef564736f6c634300081a0033",
}

// DroneABI is the input ABI used to generate the binding from.
// Deprecated: Use DroneMetaData.ABI instead.
var DroneABI = DroneMetaData.ABI

// DroneBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DroneMetaData.Bin instead.
var DroneBin = DroneMetaData.Bin

// DeployDrone deploys a new Ethereum contract, binding an instance of Drone to it.
func DeployDrone(auth *bind.TransactOpts, backend bind.ContractBackend, loggerAddress common.Address) (common.Address, *types.Transaction, *Drone, error) {
	parsed, err := DroneMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DroneBin), backend, loggerAddress)
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

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Drone *DroneCaller) Logger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Drone.contract.Call(opts, &out, "logger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Drone *DroneSession) Logger() (common.Address, error) {
	return _Drone.Contract.Logger(&_Drone.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Drone *DroneCallerSession) Logger() (common.Address, error) {
	return _Drone.Contract.Logger(&_Drone.CallOpts)
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

// GetDrone is a paid mutator transaction binding the contract method 0xdbd1be77.
//
// Solidity: function getDrone(uint256 _id) returns((uint256,string,int256,bool))
func (_Drone *DroneTransactor) GetDrone(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Drone.contract.Transact(opts, "getDrone", _id)
}

// GetDrone is a paid mutator transaction binding the contract method 0xdbd1be77.
//
// Solidity: function getDrone(uint256 _id) returns((uint256,string,int256,bool))
func (_Drone *DroneSession) GetDrone(_id *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.GetDrone(&_Drone.TransactOpts, _id)
}

// GetDrone is a paid mutator transaction binding the contract method 0xdbd1be77.
//
// Solidity: function getDrone(uint256 _id) returns((uint256,string,int256,bool))
func (_Drone *DroneTransactorSession) GetDrone(_id *big.Int) (*types.Transaction, error) {
	return _Drone.Contract.GetDrone(&_Drone.TransactOpts, _id)
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
