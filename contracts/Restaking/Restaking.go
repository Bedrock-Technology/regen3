// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Restaking

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

// BeaconChainProofsStateRootProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsStateRootProof struct {
	BeaconStateRoot [32]byte
	Proof           []byte
}

// RestakingMetaData contains all meta data concerning the Restaking contract.
var RestakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Pending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"podId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callDelegationManager\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createPod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedWithdrawalRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPod\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPodManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingWithdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getPod\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initializeV3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"podOwners\",\"outputs\":[{\"internalType\":\"contractIPodOwner\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"podId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"revertIfNoBalance\",\"type\":\"bool\"}],\"name\":\"startCheckPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"upgradeBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"podId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"oracleTimestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"name\":\"stateRootProof\",\"type\":\"tuple\"},{\"internalType\":\"uint40[]\",\"name\":\"validatorIndices\",\"type\":\"uint40[]\"},{\"internalType\":\"bytes[]\",\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"validatorFields\",\"type\":\"bytes32[][]\"}],\"name\":\"verifyWithdrawalCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RestakingABI is the input ABI used to generate the binding from.
// Deprecated: Use RestakingMetaData.ABI instead.
var RestakingABI = RestakingMetaData.ABI

// Restaking is an auto generated Go binding around an Ethereum contract.
type Restaking struct {
	RestakingCaller     // Read-only binding to the contract
	RestakingTransactor // Write-only binding to the contract
	RestakingFilterer   // Log filterer for contract events
}

// RestakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestakingSession struct {
	Contract     *Restaking        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RestakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestakingCallerSession struct {
	Contract *RestakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RestakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestakingTransactorSession struct {
	Contract     *RestakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RestakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestakingRaw struct {
	Contract *Restaking // Generic contract binding to access the raw methods on
}

// RestakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestakingCallerRaw struct {
	Contract *RestakingCaller // Generic read-only contract binding to access the raw methods on
}

// RestakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestakingTransactorRaw struct {
	Contract *RestakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestaking creates a new instance of Restaking, bound to a specific deployed contract.
func NewRestaking(address common.Address, backend bind.ContractBackend) (*Restaking, error) {
	contract, err := bindRestaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Restaking{RestakingCaller: RestakingCaller{contract: contract}, RestakingTransactor: RestakingTransactor{contract: contract}, RestakingFilterer: RestakingFilterer{contract: contract}}, nil
}

// NewRestakingCaller creates a new read-only instance of Restaking, bound to a specific deployed contract.
func NewRestakingCaller(address common.Address, caller bind.ContractCaller) (*RestakingCaller, error) {
	contract, err := bindRestaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestakingCaller{contract: contract}, nil
}

// NewRestakingTransactor creates a new write-only instance of Restaking, bound to a specific deployed contract.
func NewRestakingTransactor(address common.Address, transactor bind.ContractTransactor) (*RestakingTransactor, error) {
	contract, err := bindRestaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestakingTransactor{contract: contract}, nil
}

// NewRestakingFilterer creates a new log filterer instance of Restaking, bound to a specific deployed contract.
func NewRestakingFilterer(address common.Address, filterer bind.ContractFilterer) (*RestakingFilterer, error) {
	contract, err := bindRestaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestakingFilterer{contract: contract}, nil
}

// bindRestaking binds a generic wrapper to an already deployed contract.
func bindRestaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RestakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Restaking *RestakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Restaking.Contract.RestakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Restaking *RestakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restaking.Contract.RestakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Restaking *RestakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Restaking.Contract.RestakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Restaking *RestakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Restaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Restaking *RestakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Restaking *RestakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Restaking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Restaking *RestakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Restaking *RestakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Restaking.Contract.DEFAULTADMINROLE(&_Restaking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Restaking *RestakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Restaking.Contract.DEFAULTADMINROLE(&_Restaking.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Restaking *RestakingCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Restaking *RestakingSession) OPERATORROLE() ([32]byte, error) {
	return _Restaking.Contract.OPERATORROLE(&_Restaking.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Restaking *RestakingCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Restaking.Contract.OPERATORROLE(&_Restaking.CallOpts)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_Restaking *RestakingCaller) Beacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "beacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_Restaking *RestakingSession) Beacon() (common.Address, error) {
	return _Restaking.Contract.Beacon(&_Restaking.CallOpts)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_Restaking *RestakingCallerSession) Beacon() (common.Address, error) {
	return _Restaking.Contract.Beacon(&_Restaking.CallOpts)
}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_Restaking *RestakingCaller) DelayedWithdrawalRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "delayedWithdrawalRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_Restaking *RestakingSession) DelayedWithdrawalRouter() (common.Address, error) {
	return _Restaking.Contract.DelayedWithdrawalRouter(&_Restaking.CallOpts)
}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_Restaking *RestakingCallerSession) DelayedWithdrawalRouter() (common.Address, error) {
	return _Restaking.Contract.DelayedWithdrawalRouter(&_Restaking.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_Restaking *RestakingCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_Restaking *RestakingSession) DelegationManager() (common.Address, error) {
	return _Restaking.Contract.DelegationManager(&_Restaking.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_Restaking *RestakingCallerSession) DelegationManager() (common.Address, error) {
	return _Restaking.Contract.DelegationManager(&_Restaking.CallOpts)
}

// EigenPod is a free data retrieval call binding the contract method 0xa3aae136.
//
// Solidity: function eigenPod() view returns(address)
func (_Restaking *RestakingCaller) EigenPod(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "eigenPod")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPod is a free data retrieval call binding the contract method 0xa3aae136.
//
// Solidity: function eigenPod() view returns(address)
func (_Restaking *RestakingSession) EigenPod() (common.Address, error) {
	return _Restaking.Contract.EigenPod(&_Restaking.CallOpts)
}

// EigenPod is a free data retrieval call binding the contract method 0xa3aae136.
//
// Solidity: function eigenPod() view returns(address)
func (_Restaking *RestakingCallerSession) EigenPod() (common.Address, error) {
	return _Restaking.Contract.EigenPod(&_Restaking.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_Restaking *RestakingCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_Restaking *RestakingSession) EigenPodManager() (common.Address, error) {
	return _Restaking.Contract.EigenPodManager(&_Restaking.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_Restaking *RestakingCallerSession) EigenPodManager() (common.Address, error) {
	return _Restaking.Contract.EigenPodManager(&_Restaking.CallOpts)
}

// GetPendingWithdrawalAmount is a free data retrieval call binding the contract method 0x99cba074.
//
// Solidity: function getPendingWithdrawalAmount() view returns(uint256)
func (_Restaking *RestakingCaller) GetPendingWithdrawalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "getPendingWithdrawalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingWithdrawalAmount is a free data retrieval call binding the contract method 0x99cba074.
//
// Solidity: function getPendingWithdrawalAmount() view returns(uint256)
func (_Restaking *RestakingSession) GetPendingWithdrawalAmount() (*big.Int, error) {
	return _Restaking.Contract.GetPendingWithdrawalAmount(&_Restaking.CallOpts)
}

// GetPendingWithdrawalAmount is a free data retrieval call binding the contract method 0x99cba074.
//
// Solidity: function getPendingWithdrawalAmount() view returns(uint256)
func (_Restaking *RestakingCallerSession) GetPendingWithdrawalAmount() (*big.Int, error) {
	return _Restaking.Contract.GetPendingWithdrawalAmount(&_Restaking.CallOpts)
}

// GetPod is a free data retrieval call binding the contract method 0x6a93ee0f.
//
// Solidity: function getPod(uint256 i) view returns(address)
func (_Restaking *RestakingCaller) GetPod(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "getPod", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPod is a free data retrieval call binding the contract method 0x6a93ee0f.
//
// Solidity: function getPod(uint256 i) view returns(address)
func (_Restaking *RestakingSession) GetPod(i *big.Int) (common.Address, error) {
	return _Restaking.Contract.GetPod(&_Restaking.CallOpts, i)
}

// GetPod is a free data retrieval call binding the contract method 0x6a93ee0f.
//
// Solidity: function getPod(uint256 i) view returns(address)
func (_Restaking *RestakingCallerSession) GetPod(i *big.Int) (common.Address, error) {
	return _Restaking.Contract.GetPod(&_Restaking.CallOpts, i)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Restaking *RestakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Restaking *RestakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Restaking.Contract.GetRoleAdmin(&_Restaking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Restaking *RestakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Restaking.Contract.GetRoleAdmin(&_Restaking.CallOpts, role)
}

// GetTotalPods is a free data retrieval call binding the contract method 0xd2197bd3.
//
// Solidity: function getTotalPods() view returns(uint256)
func (_Restaking *RestakingCaller) GetTotalPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "getTotalPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPods is a free data retrieval call binding the contract method 0xd2197bd3.
//
// Solidity: function getTotalPods() view returns(uint256)
func (_Restaking *RestakingSession) GetTotalPods() (*big.Int, error) {
	return _Restaking.Contract.GetTotalPods(&_Restaking.CallOpts)
}

// GetTotalPods is a free data retrieval call binding the contract method 0xd2197bd3.
//
// Solidity: function getTotalPods() view returns(uint256)
func (_Restaking *RestakingCallerSession) GetTotalPods() (*big.Int, error) {
	return _Restaking.Contract.GetTotalPods(&_Restaking.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Restaking *RestakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Restaking *RestakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Restaking.Contract.HasRole(&_Restaking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Restaking *RestakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Restaking.Contract.HasRole(&_Restaking.CallOpts, role, account)
}

// PodOwners is a free data retrieval call binding the contract method 0xa1dccd11.
//
// Solidity: function podOwners(uint256 ) view returns(address)
func (_Restaking *RestakingCaller) PodOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "podOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PodOwners is a free data retrieval call binding the contract method 0xa1dccd11.
//
// Solidity: function podOwners(uint256 ) view returns(address)
func (_Restaking *RestakingSession) PodOwners(arg0 *big.Int) (common.Address, error) {
	return _Restaking.Contract.PodOwners(&_Restaking.CallOpts, arg0)
}

// PodOwners is a free data retrieval call binding the contract method 0xa1dccd11.
//
// Solidity: function podOwners(uint256 ) view returns(address)
func (_Restaking *RestakingCallerSession) PodOwners(arg0 *big.Int) (common.Address, error) {
	return _Restaking.Contract.PodOwners(&_Restaking.CallOpts, arg0)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_Restaking *RestakingCaller) StakingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "stakingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_Restaking *RestakingSession) StakingAddress() (common.Address, error) {
	return _Restaking.Contract.StakingAddress(&_Restaking.CallOpts)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_Restaking *RestakingCallerSession) StakingAddress() (common.Address, error) {
	return _Restaking.Contract.StakingAddress(&_Restaking.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Restaking *RestakingCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Restaking *RestakingSession) StrategyManager() (common.Address, error) {
	return _Restaking.Contract.StrategyManager(&_Restaking.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Restaking *RestakingCallerSession) StrategyManager() (common.Address, error) {
	return _Restaking.Contract.StrategyManager(&_Restaking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Restaking *RestakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Restaking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Restaking *RestakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Restaking.Contract.SupportsInterface(&_Restaking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Restaking *RestakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Restaking.Contract.SupportsInterface(&_Restaking.CallOpts, interfaceId)
}

// CallDelegationManager is a paid mutator transaction binding the contract method 0x37a26eef.
//
// Solidity: function callDelegationManager(uint256 podId, bytes data) returns(bytes)
func (_Restaking *RestakingTransactor) CallDelegationManager(opts *bind.TransactOpts, podId *big.Int, data []byte) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "callDelegationManager", podId, data)
}

// CallDelegationManager is a paid mutator transaction binding the contract method 0x37a26eef.
//
// Solidity: function callDelegationManager(uint256 podId, bytes data) returns(bytes)
func (_Restaking *RestakingSession) CallDelegationManager(podId *big.Int, data []byte) (*types.Transaction, error) {
	return _Restaking.Contract.CallDelegationManager(&_Restaking.TransactOpts, podId, data)
}

// CallDelegationManager is a paid mutator transaction binding the contract method 0x37a26eef.
//
// Solidity: function callDelegationManager(uint256 podId, bytes data) returns(bytes)
func (_Restaking *RestakingTransactorSession) CallDelegationManager(podId *big.Int, data []byte) (*types.Transaction, error) {
	return _Restaking.Contract.CallDelegationManager(&_Restaking.TransactOpts, podId, data)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns()
func (_Restaking *RestakingTransactor) CreatePod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "createPod")
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns()
func (_Restaking *RestakingSession) CreatePod() (*types.Transaction, error) {
	return _Restaking.Contract.CreatePod(&_Restaking.TransactOpts)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns()
func (_Restaking *RestakingTransactorSession) CreatePod() (*types.Transaction, error) {
	return _Restaking.Contract.CreatePod(&_Restaking.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address target, bytes data) returns(bytes)
func (_Restaking *RestakingTransactor) Execute(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "execute", target, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address target, bytes data) returns(bytes)
func (_Restaking *RestakingSession) Execute(target common.Address, data []byte) (*types.Transaction, error) {
	return _Restaking.Contract.Execute(&_Restaking.TransactOpts, target, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address target, bytes data) returns(bytes)
func (_Restaking *RestakingTransactorSession) Execute(target common.Address, data []byte) (*types.Transaction, error) {
	return _Restaking.Contract.Execute(&_Restaking.TransactOpts, target, data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Restaking *RestakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Restaking *RestakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.GrantRole(&_Restaking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Restaking *RestakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.GrantRole(&_Restaking.TransactOpts, role, account)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address impl) returns()
func (_Restaking *RestakingTransactor) InitializeV3(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "initializeV3", impl)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address impl) returns()
func (_Restaking *RestakingSession) InitializeV3(impl common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.InitializeV3(&_Restaking.TransactOpts, impl)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address impl) returns()
func (_Restaking *RestakingTransactorSession) InitializeV3(impl common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.InitializeV3(&_Restaking.TransactOpts, impl)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Restaking *RestakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Restaking *RestakingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.RenounceRole(&_Restaking.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Restaking *RestakingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.RenounceRole(&_Restaking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Restaking *RestakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Restaking *RestakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.RevokeRole(&_Restaking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Restaking *RestakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.RevokeRole(&_Restaking.TransactOpts, role, account)
}

// StartCheckPoint is a paid mutator transaction binding the contract method 0xece10d61.
//
// Solidity: function startCheckPoint(uint256 podId, bool revertIfNoBalance) returns()
func (_Restaking *RestakingTransactor) StartCheckPoint(opts *bind.TransactOpts, podId *big.Int, revertIfNoBalance bool) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "startCheckPoint", podId, revertIfNoBalance)
}

// StartCheckPoint is a paid mutator transaction binding the contract method 0xece10d61.
//
// Solidity: function startCheckPoint(uint256 podId, bool revertIfNoBalance) returns()
func (_Restaking *RestakingSession) StartCheckPoint(podId *big.Int, revertIfNoBalance bool) (*types.Transaction, error) {
	return _Restaking.Contract.StartCheckPoint(&_Restaking.TransactOpts, podId, revertIfNoBalance)
}

// StartCheckPoint is a paid mutator transaction binding the contract method 0xece10d61.
//
// Solidity: function startCheckPoint(uint256 podId, bool revertIfNoBalance) returns()
func (_Restaking *RestakingTransactorSession) StartCheckPoint(podId *big.Int, revertIfNoBalance bool) (*types.Transaction, error) {
	return _Restaking.Contract.StartCheckPoint(&_Restaking.TransactOpts, podId, revertIfNoBalance)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address target, uint256 amount) returns()
func (_Restaking *RestakingTransactor) Transfer(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "transfer", target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address target, uint256 amount) returns()
func (_Restaking *RestakingSession) Transfer(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Restaking.Contract.Transfer(&_Restaking.TransactOpts, target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address target, uint256 amount) returns()
func (_Restaking *RestakingTransactorSession) Transfer(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Restaking.Contract.Transfer(&_Restaking.TransactOpts, target, amount)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_Restaking *RestakingTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_Restaking *RestakingSession) Update() (*types.Transaction, error) {
	return _Restaking.Contract.Update(&_Restaking.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_Restaking *RestakingTransactorSession) Update() (*types.Transaction, error) {
	return _Restaking.Contract.Update(&_Restaking.TransactOpts)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x1bce4583.
//
// Solidity: function upgradeBeacon(address impl) returns()
func (_Restaking *RestakingTransactor) UpgradeBeacon(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "upgradeBeacon", impl)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x1bce4583.
//
// Solidity: function upgradeBeacon(address impl) returns()
func (_Restaking *RestakingSession) UpgradeBeacon(impl common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.UpgradeBeacon(&_Restaking.TransactOpts, impl)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x1bce4583.
//
// Solidity: function upgradeBeacon(address impl) returns()
func (_Restaking *RestakingTransactorSession) UpgradeBeacon(impl common.Address) (*types.Transaction, error) {
	return _Restaking.Contract.UpgradeBeacon(&_Restaking.TransactOpts, impl)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0xc125d9d7.
//
// Solidity: function verifyWithdrawalCredentials(uint256 podId, uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_Restaking *RestakingTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, podId *big.Int, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Restaking.contract.Transact(opts, "verifyWithdrawalCredentials", podId, oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0xc125d9d7.
//
// Solidity: function verifyWithdrawalCredentials(uint256 podId, uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_Restaking *RestakingSession) VerifyWithdrawalCredentials(podId *big.Int, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Restaking.Contract.VerifyWithdrawalCredentials(&_Restaking.TransactOpts, podId, oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0xc125d9d7.
//
// Solidity: function verifyWithdrawalCredentials(uint256 podId, uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_Restaking *RestakingTransactorSession) VerifyWithdrawalCredentials(podId *big.Int, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Restaking.Contract.VerifyWithdrawalCredentials(&_Restaking.TransactOpts, podId, oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Restaking *RestakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restaking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Restaking *RestakingSession) Receive() (*types.Transaction, error) {
	return _Restaking.Contract.Receive(&_Restaking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Restaking *RestakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Restaking.Contract.Receive(&_Restaking.TransactOpts)
}

// RestakingClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Restaking contract.
type RestakingClaimedIterator struct {
	Event *RestakingClaimed // Event containing the contract specifics and raw log

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
func (it *RestakingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingClaimed)
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
		it.Event = new(RestakingClaimed)
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
func (it *RestakingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingClaimed represents a Claimed event raised by the Restaking contract.
type RestakingClaimed struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8.
//
// Solidity: event Claimed(uint256 amount)
func (_Restaking *RestakingFilterer) FilterClaimed(opts *bind.FilterOpts) (*RestakingClaimedIterator, error) {

	logs, sub, err := _Restaking.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &RestakingClaimedIterator{contract: _Restaking.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8.
//
// Solidity: event Claimed(uint256 amount)
func (_Restaking *RestakingFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *RestakingClaimed) (event.Subscription, error) {

	logs, sub, err := _Restaking.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingClaimed)
				if err := _Restaking.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8.
//
// Solidity: event Claimed(uint256 amount)
func (_Restaking *RestakingFilterer) ParseClaimed(log types.Log) (*RestakingClaimed, error) {
	event := new(RestakingClaimed)
	if err := _Restaking.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Restaking contract.
type RestakingInitializedIterator struct {
	Event *RestakingInitialized // Event containing the contract specifics and raw log

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
func (it *RestakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingInitialized)
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
		it.Event = new(RestakingInitialized)
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
func (it *RestakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingInitialized represents a Initialized event raised by the Restaking contract.
type RestakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Restaking *RestakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*RestakingInitializedIterator, error) {

	logs, sub, err := _Restaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RestakingInitializedIterator{contract: _Restaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Restaking *RestakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RestakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Restaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingInitialized)
				if err := _Restaking.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Restaking *RestakingFilterer) ParseInitialized(log types.Log) (*RestakingInitialized, error) {
	event := new(RestakingInitialized)
	if err := _Restaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingPendingIterator is returned from FilterPending and is used to iterate over the raw logs and unpacked data for Pending events raised by the Restaking contract.
type RestakingPendingIterator struct {
	Event *RestakingPending // Event containing the contract specifics and raw log

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
func (it *RestakingPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingPending)
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
		it.Event = new(RestakingPending)
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
func (it *RestakingPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingPending represents a Pending event raised by the Restaking contract.
type RestakingPending struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPending is a free log retrieval operation binding the contract event 0x4e2c5ab35c5a6c864cb37f0d4b556e39e83aa85e744306ad9fd3ad7cc933028b.
//
// Solidity: event Pending(uint256 amount)
func (_Restaking *RestakingFilterer) FilterPending(opts *bind.FilterOpts) (*RestakingPendingIterator, error) {

	logs, sub, err := _Restaking.contract.FilterLogs(opts, "Pending")
	if err != nil {
		return nil, err
	}
	return &RestakingPendingIterator{contract: _Restaking.contract, event: "Pending", logs: logs, sub: sub}, nil
}

// WatchPending is a free log subscription operation binding the contract event 0x4e2c5ab35c5a6c864cb37f0d4b556e39e83aa85e744306ad9fd3ad7cc933028b.
//
// Solidity: event Pending(uint256 amount)
func (_Restaking *RestakingFilterer) WatchPending(opts *bind.WatchOpts, sink chan<- *RestakingPending) (event.Subscription, error) {

	logs, sub, err := _Restaking.contract.WatchLogs(opts, "Pending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingPending)
				if err := _Restaking.contract.UnpackLog(event, "Pending", log); err != nil {
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

// ParsePending is a log parse operation binding the contract event 0x4e2c5ab35c5a6c864cb37f0d4b556e39e83aa85e744306ad9fd3ad7cc933028b.
//
// Solidity: event Pending(uint256 amount)
func (_Restaking *RestakingFilterer) ParsePending(log types.Log) (*RestakingPending, error) {
	event := new(RestakingPending)
	if err := _Restaking.contract.UnpackLog(event, "Pending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Restaking contract.
type RestakingRoleAdminChangedIterator struct {
	Event *RestakingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RestakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingRoleAdminChanged)
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
		it.Event = new(RestakingRoleAdminChanged)
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
func (it *RestakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingRoleAdminChanged represents a RoleAdminChanged event raised by the Restaking contract.
type RestakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Restaking *RestakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RestakingRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Restaking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RestakingRoleAdminChangedIterator{contract: _Restaking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Restaking *RestakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RestakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Restaking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingRoleAdminChanged)
				if err := _Restaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Restaking *RestakingFilterer) ParseRoleAdminChanged(log types.Log) (*RestakingRoleAdminChanged, error) {
	event := new(RestakingRoleAdminChanged)
	if err := _Restaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Restaking contract.
type RestakingRoleGrantedIterator struct {
	Event *RestakingRoleGranted // Event containing the contract specifics and raw log

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
func (it *RestakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingRoleGranted)
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
		it.Event = new(RestakingRoleGranted)
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
func (it *RestakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingRoleGranted represents a RoleGranted event raised by the Restaking contract.
type RestakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Restaking *RestakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RestakingRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Restaking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RestakingRoleGrantedIterator{contract: _Restaking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Restaking *RestakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RestakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Restaking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingRoleGranted)
				if err := _Restaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Restaking *RestakingFilterer) ParseRoleGranted(log types.Log) (*RestakingRoleGranted, error) {
	event := new(RestakingRoleGranted)
	if err := _Restaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Restaking contract.
type RestakingRoleRevokedIterator struct {
	Event *RestakingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RestakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingRoleRevoked)
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
		it.Event = new(RestakingRoleRevoked)
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
func (it *RestakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingRoleRevoked represents a RoleRevoked event raised by the Restaking contract.
type RestakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Restaking *RestakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RestakingRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Restaking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RestakingRoleRevokedIterator{contract: _Restaking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Restaking *RestakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RestakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Restaking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingRoleRevoked)
				if err := _Restaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Restaking *RestakingFilterer) ParseRoleRevoked(log types.Log) (*RestakingRoleRevoked, error) {
	event := new(RestakingRoleRevoked)
	if err := _Restaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
