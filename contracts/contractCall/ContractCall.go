// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractCall

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

// ContractCallMetaData contains all meta data concerning the ContractCall contract.
var ContractCallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchSafeBurnCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"tokenIds\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"batchSafeMintCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"from\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchSecurityTransferCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minterAddress\",\"type\":\"address\"}],\"name\":\"setMinterCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610ed4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063a7cb998111610050578063a7cb9981146100e3578063d17598d514610113578063e75e0e0b1461014357610072565b8063211a01f1146100775780635c2e05b814610095578063688db184146100c5575b600080fd5b61007f610173565b60405161008c91906104f8565b60405180910390f35b6100af60048036038101906100aa919061086a565b6101e5565b6040516100bc91906104f8565b60405180910390f35b6100cd61026a565b6040516100da91906104f8565b60405180910390f35b6100fd60048036038101906100f89190610a0a565b6102dc565b60405161010a91906104f8565b60405180910390f35b61012d60048036038101906101289190610ab1565b610361565b60405161013a91906104f8565b60405180910390f35b61015d60048036038101906101589190610afa565b6103e0565b60405161016a91906104f8565b60405180910390f35b6060633f4ba83a60e01b604051602401604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905090565b606063df3720f660e01b84848460405160240161020493929190610cfc565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090509392505050565b6060638456cb5960e01b604051602401604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905090565b606063708f17db60e01b8484846040516024016102fb93929190610e06565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090509392505050565b606063cece0b9060e01b8260405160240161037c9190610e52565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050919050565b606063fca3b5aa60e01b826040516024016103fb9190610e83565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561049957808201518184015260208101905061047e565b838111156104a8576000848401525b50505050565b6000601f19601f8301169050919050565b60006104ca8261045f565b6104d4818561046a565b93506104e481856020860161047b565b6104ed816104ae565b840191505092915050565b6000602082019050818103600083015261051281846104bf565b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61056b826104ae565b810181811067ffffffffffffffff8211171561058a57610589610533565b5b80604052505050565b600061059d61051a565b90506105a98282610562565b919050565b600067ffffffffffffffff8211156105c9576105c8610533565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061060a826105df565b9050919050565b61061a816105ff565b811461062557600080fd5b50565b60008135905061063781610611565b92915050565b600061065061064b846105ae565b610593565b90508083825260208201905060208402830185811115610673576106726105da565b5b835b8181101561069c57806106888882610628565b845260208401935050602081019050610675565b5050509392505050565b600082601f8301126106bb576106ba61052e565b5b81356106cb84826020860161063d565b91505092915050565b600067ffffffffffffffff8211156106ef576106ee610533565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff8211156107205761071f610533565b5b610729826104ae565b9050602081019050919050565b82818337600083830152505050565b600061075861075384610705565b610593565b90508281526020810184848401111561077457610773610700565b5b61077f848285610736565b509392505050565b600082601f83011261079c5761079b61052e565b5b81356107ac848260208601610745565b91505092915050565b60006107c86107c3846106d4565b610593565b905080838252602082019050602084028301858111156107eb576107ea6105da565b5b835b8181101561083257803567ffffffffffffffff8111156108105761080f61052e565b5b80860161081d8982610787565b855260208501945050506020810190506107ed565b5050509392505050565b600082601f8301126108515761085061052e565b5b81356108618482602086016107b5565b91505092915050565b60008060006060848603121561088357610882610524565b5b600084013567ffffffffffffffff8111156108a1576108a0610529565b5b6108ad868287016106a6565b935050602084013567ffffffffffffffff8111156108ce576108cd610529565b5b6108da8682870161083c565b925050604084013567ffffffffffffffff8111156108fb576108fa610529565b5b6109078682870161083c565b9150509250925092565b600067ffffffffffffffff82111561092c5761092b610533565b5b602082029050602081019050919050565b6000819050919050565b6109508161093d565b811461095b57600080fd5b50565b60008135905061096d81610947565b92915050565b600061098661098184610911565b610593565b905080838252602082019050602084028301858111156109a9576109a86105da565b5b835b818110156109d257806109be888261095e565b8452602084019350506020810190506109ab565b5050509392505050565b600082601f8301126109f1576109f061052e565b5b8135610a01848260208601610973565b91505092915050565b600080600060608486031215610a2357610a22610524565b5b600084013567ffffffffffffffff811115610a4157610a40610529565b5b610a4d868287016106a6565b935050602084013567ffffffffffffffff811115610a6e57610a6d610529565b5b610a7a868287016106a6565b925050604084013567ffffffffffffffff811115610a9b57610a9a610529565b5b610aa7868287016109dc565b9150509250925092565b600060208284031215610ac757610ac6610524565b5b600082013567ffffffffffffffff811115610ae557610ae4610529565b5b610af1848285016109dc565b91505092915050565b600060208284031215610b1057610b0f610524565b5b6000610b1e84828501610628565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610b5c816105ff565b82525050565b6000610b6e8383610b53565b60208301905092915050565b6000602082019050919050565b6000610b9282610b27565b610b9c8185610b32565b9350610ba783610b43565b8060005b83811015610bd8578151610bbf8882610b62565b9750610bca83610b7a565b925050600181019050610bab565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b6000610c3882610c11565b610c428185610c1c565b9350610c5281856020860161047b565b610c5b816104ae565b840191505092915050565b6000610c728383610c2d565b905092915050565b6000602082019050919050565b6000610c9282610be5565b610c9c8185610bf0565b935083602082028501610cae85610c01565b8060005b85811015610cea5784840389528151610ccb8582610c66565b9450610cd683610c7a565b925060208a01995050600181019050610cb2565b50829750879550505050505092915050565b60006060820190508181036000830152610d168186610b87565b90508181036020830152610d2a8185610c87565b90508181036040830152610d3e8184610c87565b9050949350505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610d7d8161093d565b82525050565b6000610d8f8383610d74565b60208301905092915050565b6000602082019050919050565b6000610db382610d48565b610dbd8185610d53565b9350610dc883610d64565b8060005b83811015610df9578151610de08882610d83565b9750610deb83610d9b565b925050600181019050610dcc565b5085935050505092915050565b60006060820190508181036000830152610e208186610b87565b90508181036020830152610e348185610b87565b90508181036040830152610e488184610da8565b9050949350505050565b60006020820190508181036000830152610e6c8184610da8565b905092915050565b610e7d816105ff565b82525050565b6000602082019050610e986000830184610e74565b9291505056fea2646970667358221220c4c06bdafa06f92aeec05491bc67ff12706c5878e0cd6d5c8963f29d0665b8de64736f6c63430008090033",
}

// ContractCallABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractCallMetaData.ABI instead.
var ContractCallABI = ContractCallMetaData.ABI

// ContractCallBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractCallMetaData.Bin instead.
var ContractCallBin = ContractCallMetaData.Bin

// DeployContractCall deploys a new Ethereum contract, binding an instance of ContractCall to it.
func DeployContractCall(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractCall, error) {
	parsed, err := ContractCallMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractCallBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractCall{ContractCallCaller: ContractCallCaller{contract: contract}, ContractCallTransactor: ContractCallTransactor{contract: contract}, ContractCallFilterer: ContractCallFilterer{contract: contract}}, nil
}

// ContractCall is an auto generated Go binding around an Ethereum contract.
type ContractCall struct {
	ContractCallCaller     // Read-only binding to the contract
	ContractCallTransactor // Write-only binding to the contract
	ContractCallFilterer   // Log filterer for contract events
}

// ContractCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractCallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractCallSession struct {
	Contract     *ContractCall     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallCallerSession struct {
	Contract *ContractCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ContractCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractCallTransactorSession struct {
	Contract     *ContractCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractCallRaw struct {
	Contract *ContractCall // Generic contract binding to access the raw methods on
}

// ContractCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallCallerRaw struct {
	Contract *ContractCallCaller // Generic read-only contract binding to access the raw methods on
}

// ContractCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractCallTransactorRaw struct {
	Contract *ContractCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractCall creates a new instance of ContractCall, bound to a specific deployed contract.
func NewContractCall(address common.Address, backend bind.ContractBackend) (*ContractCall, error) {
	contract, err := bindContractCall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractCall{ContractCallCaller: ContractCallCaller{contract: contract}, ContractCallTransactor: ContractCallTransactor{contract: contract}, ContractCallFilterer: ContractCallFilterer{contract: contract}}, nil
}

// NewContractCallCaller creates a new read-only instance of ContractCall, bound to a specific deployed contract.
func NewContractCallCaller(address common.Address, caller bind.ContractCaller) (*ContractCallCaller, error) {
	contract, err := bindContractCall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCallCaller{contract: contract}, nil
}

// NewContractCallTransactor creates a new write-only instance of ContractCall, bound to a specific deployed contract.
func NewContractCallTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractCallTransactor, error) {
	contract, err := bindContractCall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCallTransactor{contract: contract}, nil
}

// NewContractCallFilterer creates a new log filterer instance of ContractCall, bound to a specific deployed contract.
func NewContractCallFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractCallFilterer, error) {
	contract, err := bindContractCall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractCallFilterer{contract: contract}, nil
}

// bindContractCall binds a generic wrapper to an already deployed contract.
func bindContractCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractCallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractCall *ContractCallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractCall.Contract.ContractCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractCall *ContractCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCall.Contract.ContractCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractCall *ContractCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractCall.Contract.ContractCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractCall *ContractCallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractCall *ContractCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractCall *ContractCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractCall.Contract.contract.Transact(opts, method, params...)
}

// BatchSafeBurnCall is a free data retrieval call binding the contract method 0xd17598d5.
//
// Solidity: function batchSafeBurnCall(uint256[] tokenIds) pure returns(bytes data)
func (_ContractCall *ContractCallCaller) BatchSafeBurnCall(opts *bind.CallOpts, tokenIds []*big.Int) ([]byte, error) {
	var out []interface{}
	err := _ContractCall.contract.Call(opts, &out, "batchSafeBurnCall", tokenIds)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BatchSafeBurnCall is a free data retrieval call binding the contract method 0xd17598d5.
//
// Solidity: function batchSafeBurnCall(uint256[] tokenIds) pure returns(bytes data)
func (_ContractCall *ContractCallSession) BatchSafeBurnCall(tokenIds []*big.Int) ([]byte, error) {
	return _ContractCall.Contract.BatchSafeBurnCall(&_ContractCall.CallOpts, tokenIds)
}

// BatchSafeBurnCall is a free data retrieval call binding the contract method 0xd17598d5.
//
// Solidity: function batchSafeBurnCall(uint256[] tokenIds) pure returns(bytes data)
func (_ContractCall *ContractCallCallerSession) BatchSafeBurnCall(tokenIds []*big.Int) ([]byte, error) {
	return _ContractCall.Contract.BatchSafeBurnCall(&_ContractCall.CallOpts, tokenIds)
}

// BatchSafeMintCall is a free data retrieval call binding the contract method 0x5c2e05b8.
//
// Solidity: function batchSafeMintCall(address[] tos, string[] tokenIds, string[] uris) pure returns(bytes data)
func (_ContractCall *ContractCallCaller) BatchSafeMintCall(opts *bind.CallOpts, tos []common.Address, tokenIds []string, uris []string) ([]byte, error) {
	var out []interface{}
	err := _ContractCall.contract.Call(opts, &out, "batchSafeMintCall", tos, tokenIds, uris)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BatchSafeMintCall is a free data retrieval call binding the contract method 0x5c2e05b8.
//
// Solidity: function batchSafeMintCall(address[] tos, string[] tokenIds, string[] uris) pure returns(bytes data)
func (_ContractCall *ContractCallSession) BatchSafeMintCall(tos []common.Address, tokenIds []string, uris []string) ([]byte, error) {
	return _ContractCall.Contract.BatchSafeMintCall(&_ContractCall.CallOpts, tos, tokenIds, uris)
}

// BatchSafeMintCall is a free data retrieval call binding the contract method 0x5c2e05b8.
//
// Solidity: function batchSafeMintCall(address[] tos, string[] tokenIds, string[] uris) pure returns(bytes data)
func (_ContractCall *ContractCallCallerSession) BatchSafeMintCall(tos []common.Address, tokenIds []string, uris []string) ([]byte, error) {
	return _ContractCall.Contract.BatchSafeMintCall(&_ContractCall.CallOpts, tos, tokenIds, uris)
}

// BatchSecurityTransferCall is a free data retrieval call binding the contract method 0xa7cb9981.
//
// Solidity: function batchSecurityTransferCall(address[] from, address[] to, uint256[] tokenIds) pure returns(bytes data)
func (_ContractCall *ContractCallCaller) BatchSecurityTransferCall(opts *bind.CallOpts, from []common.Address, to []common.Address, tokenIds []*big.Int) ([]byte, error) {
	var out []interface{}
	err := _ContractCall.contract.Call(opts, &out, "batchSecurityTransferCall", from, to, tokenIds)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BatchSecurityTransferCall is a free data retrieval call binding the contract method 0xa7cb9981.
//
// Solidity: function batchSecurityTransferCall(address[] from, address[] to, uint256[] tokenIds) pure returns(bytes data)
func (_ContractCall *ContractCallSession) BatchSecurityTransferCall(from []common.Address, to []common.Address, tokenIds []*big.Int) ([]byte, error) {
	return _ContractCall.Contract.BatchSecurityTransferCall(&_ContractCall.CallOpts, from, to, tokenIds)
}

// BatchSecurityTransferCall is a free data retrieval call binding the contract method 0xa7cb9981.
//
// Solidity: function batchSecurityTransferCall(address[] from, address[] to, uint256[] tokenIds) pure returns(bytes data)
func (_ContractCall *ContractCallCallerSession) BatchSecurityTransferCall(from []common.Address, to []common.Address, tokenIds []*big.Int) ([]byte, error) {
	return _ContractCall.Contract.BatchSecurityTransferCall(&_ContractCall.CallOpts, from, to, tokenIds)
}

// PauseCall is a free data retrieval call binding the contract method 0x688db184.
//
// Solidity: function pauseCall() pure returns(bytes data)
func (_ContractCall *ContractCallCaller) PauseCall(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ContractCall.contract.Call(opts, &out, "pauseCall")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PauseCall is a free data retrieval call binding the contract method 0x688db184.
//
// Solidity: function pauseCall() pure returns(bytes data)
func (_ContractCall *ContractCallSession) PauseCall() ([]byte, error) {
	return _ContractCall.Contract.PauseCall(&_ContractCall.CallOpts)
}

// PauseCall is a free data retrieval call binding the contract method 0x688db184.
//
// Solidity: function pauseCall() pure returns(bytes data)
func (_ContractCall *ContractCallCallerSession) PauseCall() ([]byte, error) {
	return _ContractCall.Contract.PauseCall(&_ContractCall.CallOpts)
}

// SetMinterCall is a free data retrieval call binding the contract method 0xe75e0e0b.
//
// Solidity: function setMinterCall(address minterAddress) pure returns(bytes data)
func (_ContractCall *ContractCallCaller) SetMinterCall(opts *bind.CallOpts, minterAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _ContractCall.contract.Call(opts, &out, "setMinterCall", minterAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SetMinterCall is a free data retrieval call binding the contract method 0xe75e0e0b.
//
// Solidity: function setMinterCall(address minterAddress) pure returns(bytes data)
func (_ContractCall *ContractCallSession) SetMinterCall(minterAddress common.Address) ([]byte, error) {
	return _ContractCall.Contract.SetMinterCall(&_ContractCall.CallOpts, minterAddress)
}

// SetMinterCall is a free data retrieval call binding the contract method 0xe75e0e0b.
//
// Solidity: function setMinterCall(address minterAddress) pure returns(bytes data)
func (_ContractCall *ContractCallCallerSession) SetMinterCall(minterAddress common.Address) ([]byte, error) {
	return _ContractCall.Contract.SetMinterCall(&_ContractCall.CallOpts, minterAddress)
}

// UnpauseCall is a free data retrieval call binding the contract method 0x211a01f1.
//
// Solidity: function unpauseCall() pure returns(bytes data)
func (_ContractCall *ContractCallCaller) UnpauseCall(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ContractCall.contract.Call(opts, &out, "unpauseCall")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// UnpauseCall is a free data retrieval call binding the contract method 0x211a01f1.
//
// Solidity: function unpauseCall() pure returns(bytes data)
func (_ContractCall *ContractCallSession) UnpauseCall() ([]byte, error) {
	return _ContractCall.Contract.UnpauseCall(&_ContractCall.CallOpts)
}

// UnpauseCall is a free data retrieval call binding the contract method 0x211a01f1.
//
// Solidity: function unpauseCall() pure returns(bytes data)
func (_ContractCall *ContractCallCallerSession) UnpauseCall() ([]byte, error) {
	return _ContractCall.Contract.UnpauseCall(&_ContractCall.CallOpts)
}
