// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DelayedWithdrawalRouter

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

// IDelayedWithdrawalRouterDelayedWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelayedWithdrawalRouterDelayedWithdrawal struct {
	Amount       *big.Int
	BlockCreated uint32
}

// IDelayedWithdrawalRouterUserDelayedWithdrawals is an auto generated low-level Go binding around an user-defined struct.
type IDelayedWithdrawalRouterUserDelayedWithdrawals struct {
	DelayedWithdrawalsCompleted *big.Int
	DelayedWithdrawals          []IDelayedWithdrawalRouterDelayedWithdrawal
}

// DelayedWithdrawalRouterMetaData contains all meta data concerning the DelayedWithdrawalRouter contract.
var DelayedWithdrawalRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"_eigenPodManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DelayedWithdrawalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountClaimed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delayedWithdrawalsCompleted\",\"type\":\"uint256\"}],\"name\":\"DelayedWithdrawalsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"WithdrawalDelayBlocksSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAWAL_DELAY_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"canClaimDelayedWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxNumberOfDelayedWithdrawalsToClaim\",\"type\":\"uint256\"}],\"name\":\"claimDelayedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxNumberOfDelayedWithdrawalsToClaim\",\"type\":\"uint256\"}],\"name\":\"claimDelayedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"createDelayedWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPodManager\",\"outputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimableUserDelayedWithdrawals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"blockCreated\",\"type\":\"uint32\"}],\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserDelayedWithdrawals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"blockCreated\",\"type\":\"uint32\"}],\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initOwner\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initPausedStatus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalDelayBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"userDelayedWithdrawalByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"blockCreated\",\"type\":\"uint32\"}],\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userWithdrawals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayedWithdrawalsCompleted\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"blockCreated\",\"type\":\"uint32\"}],\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal[]\",\"name\":\"delayedWithdrawals\",\"type\":\"tuple[]\"}],\"internalType\":\"structIDelayedWithdrawalRouter.UserDelayedWithdrawals\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userWithdrawalsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DelayedWithdrawalRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use DelayedWithdrawalRouterMetaData.ABI instead.
var DelayedWithdrawalRouterABI = DelayedWithdrawalRouterMetaData.ABI

// DelayedWithdrawalRouter is an auto generated Go binding around an Ethereum contract.
type DelayedWithdrawalRouter struct {
	DelayedWithdrawalRouterCaller     // Read-only binding to the contract
	DelayedWithdrawalRouterTransactor // Write-only binding to the contract
	DelayedWithdrawalRouterFilterer   // Log filterer for contract events
}

// DelayedWithdrawalRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelayedWithdrawalRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedWithdrawalRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelayedWithdrawalRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedWithdrawalRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelayedWithdrawalRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedWithdrawalRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelayedWithdrawalRouterSession struct {
	Contract     *DelayedWithdrawalRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DelayedWithdrawalRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelayedWithdrawalRouterCallerSession struct {
	Contract *DelayedWithdrawalRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// DelayedWithdrawalRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelayedWithdrawalRouterTransactorSession struct {
	Contract     *DelayedWithdrawalRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// DelayedWithdrawalRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelayedWithdrawalRouterRaw struct {
	Contract *DelayedWithdrawalRouter // Generic contract binding to access the raw methods on
}

// DelayedWithdrawalRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelayedWithdrawalRouterCallerRaw struct {
	Contract *DelayedWithdrawalRouterCaller // Generic read-only contract binding to access the raw methods on
}

// DelayedWithdrawalRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelayedWithdrawalRouterTransactorRaw struct {
	Contract *DelayedWithdrawalRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelayedWithdrawalRouter creates a new instance of DelayedWithdrawalRouter, bound to a specific deployed contract.
func NewDelayedWithdrawalRouter(address common.Address, backend bind.ContractBackend) (*DelayedWithdrawalRouter, error) {
	contract, err := bindDelayedWithdrawalRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouter{DelayedWithdrawalRouterCaller: DelayedWithdrawalRouterCaller{contract: contract}, DelayedWithdrawalRouterTransactor: DelayedWithdrawalRouterTransactor{contract: contract}, DelayedWithdrawalRouterFilterer: DelayedWithdrawalRouterFilterer{contract: contract}}, nil
}

// NewDelayedWithdrawalRouterCaller creates a new read-only instance of DelayedWithdrawalRouter, bound to a specific deployed contract.
func NewDelayedWithdrawalRouterCaller(address common.Address, caller bind.ContractCaller) (*DelayedWithdrawalRouterCaller, error) {
	contract, err := bindDelayedWithdrawalRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterCaller{contract: contract}, nil
}

// NewDelayedWithdrawalRouterTransactor creates a new write-only instance of DelayedWithdrawalRouter, bound to a specific deployed contract.
func NewDelayedWithdrawalRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*DelayedWithdrawalRouterTransactor, error) {
	contract, err := bindDelayedWithdrawalRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterTransactor{contract: contract}, nil
}

// NewDelayedWithdrawalRouterFilterer creates a new log filterer instance of DelayedWithdrawalRouter, bound to a specific deployed contract.
func NewDelayedWithdrawalRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*DelayedWithdrawalRouterFilterer, error) {
	contract, err := bindDelayedWithdrawalRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterFilterer{contract: contract}, nil
}

// bindDelayedWithdrawalRouter binds a generic wrapper to an already deployed contract.
func bindDelayedWithdrawalRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelayedWithdrawalRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedWithdrawalRouter.Contract.DelayedWithdrawalRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.DelayedWithdrawalRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.DelayedWithdrawalRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedWithdrawalRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.contract.Transact(opts, method, params...)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) MAXWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.MAXWITHDRAWALDELAYBLOCKS(&_DelayedWithdrawalRouter.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.MAXWITHDRAWALDELAYBLOCKS(&_DelayedWithdrawalRouter.CallOpts)
}

// CanClaimDelayedWithdrawal is a free data retrieval call binding the contract method 0x75608896.
//
// Solidity: function canClaimDelayedWithdrawal(address user, uint256 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) CanClaimDelayedWithdrawal(opts *bind.CallOpts, user common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "canClaimDelayedWithdrawal", user, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimDelayedWithdrawal is a free data retrieval call binding the contract method 0x75608896.
//
// Solidity: function canClaimDelayedWithdrawal(address user, uint256 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) CanClaimDelayedWithdrawal(user common.Address, index *big.Int) (bool, error) {
	return _DelayedWithdrawalRouter.Contract.CanClaimDelayedWithdrawal(&_DelayedWithdrawalRouter.CallOpts, user, index)
}

// CanClaimDelayedWithdrawal is a free data retrieval call binding the contract method 0x75608896.
//
// Solidity: function canClaimDelayedWithdrawal(address user, uint256 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) CanClaimDelayedWithdrawal(user common.Address, index *big.Int) (bool, error) {
	return _DelayedWithdrawalRouter.Contract.CanClaimDelayedWithdrawal(&_DelayedWithdrawalRouter.CallOpts, user, index)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) EigenPodManager() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.EigenPodManager(&_DelayedWithdrawalRouter.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) EigenPodManager() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.EigenPodManager(&_DelayedWithdrawalRouter.CallOpts)
}

// GetClaimableUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x1f39d87f.
//
// Solidity: function getClaimableUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) GetClaimableUserDelayedWithdrawals(opts *bind.CallOpts, user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "getClaimableUserDelayedWithdrawals", user)

	if err != nil {
		return *new([]IDelayedWithdrawalRouterDelayedWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelayedWithdrawalRouterDelayedWithdrawal)).(*[]IDelayedWithdrawalRouterDelayedWithdrawal)

	return out0, err

}

// GetClaimableUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x1f39d87f.
//
// Solidity: function getClaimableUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) GetClaimableUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.GetClaimableUserDelayedWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// GetClaimableUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x1f39d87f.
//
// Solidity: function getClaimableUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) GetClaimableUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.GetClaimableUserDelayedWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// GetUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x3e1de008.
//
// Solidity: function getUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) GetUserDelayedWithdrawals(opts *bind.CallOpts, user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "getUserDelayedWithdrawals", user)

	if err != nil {
		return *new([]IDelayedWithdrawalRouterDelayedWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelayedWithdrawalRouterDelayedWithdrawal)).(*[]IDelayedWithdrawalRouterDelayedWithdrawal)

	return out0, err

}

// GetUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x3e1de008.
//
// Solidity: function getUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) GetUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.GetUserDelayedWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// GetUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x3e1de008.
//
// Solidity: function getUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) GetUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.GetUserDelayedWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Owner() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.Owner(&_DelayedWithdrawalRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) Owner() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.Owner(&_DelayedWithdrawalRouter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Paused(index uint8) (bool, error) {
	return _DelayedWithdrawalRouter.Contract.Paused(&_DelayedWithdrawalRouter.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) Paused(index uint8) (bool, error) {
	return _DelayedWithdrawalRouter.Contract.Paused(&_DelayedWithdrawalRouter.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Paused0() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.Paused0(&_DelayedWithdrawalRouter.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) Paused0() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.Paused0(&_DelayedWithdrawalRouter.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) PauserRegistry() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.PauserRegistry(&_DelayedWithdrawalRouter.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) PauserRegistry() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.PauserRegistry(&_DelayedWithdrawalRouter.CallOpts)
}

// UserDelayedWithdrawalByIndex is a free data retrieval call binding the contract method 0x85594e58.
//
// Solidity: function userDelayedWithdrawalByIndex(address user, uint256 index) view returns((uint224,uint32))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) UserDelayedWithdrawalByIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "userDelayedWithdrawalByIndex", user, index)

	if err != nil {
		return *new(IDelayedWithdrawalRouterDelayedWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelayedWithdrawalRouterDelayedWithdrawal)).(*IDelayedWithdrawalRouterDelayedWithdrawal)

	return out0, err

}

// UserDelayedWithdrawalByIndex is a free data retrieval call binding the contract method 0x85594e58.
//
// Solidity: function userDelayedWithdrawalByIndex(address user, uint256 index) view returns((uint224,uint32))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) UserDelayedWithdrawalByIndex(user common.Address, index *big.Int) (IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.UserDelayedWithdrawalByIndex(&_DelayedWithdrawalRouter.CallOpts, user, index)
}

// UserDelayedWithdrawalByIndex is a free data retrieval call binding the contract method 0x85594e58.
//
// Solidity: function userDelayedWithdrawalByIndex(address user, uint256 index) view returns((uint224,uint32))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) UserDelayedWithdrawalByIndex(user common.Address, index *big.Int) (IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.UserDelayedWithdrawalByIndex(&_DelayedWithdrawalRouter.CallOpts, user, index)
}

// UserWithdrawals is a free data retrieval call binding the contract method 0xecb7cb1b.
//
// Solidity: function userWithdrawals(address user) view returns((uint256,(uint224,uint32)[]))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) UserWithdrawals(opts *bind.CallOpts, user common.Address) (IDelayedWithdrawalRouterUserDelayedWithdrawals, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "userWithdrawals", user)

	if err != nil {
		return *new(IDelayedWithdrawalRouterUserDelayedWithdrawals), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelayedWithdrawalRouterUserDelayedWithdrawals)).(*IDelayedWithdrawalRouterUserDelayedWithdrawals)

	return out0, err

}

// UserWithdrawals is a free data retrieval call binding the contract method 0xecb7cb1b.
//
// Solidity: function userWithdrawals(address user) view returns((uint256,(uint224,uint32)[]))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) UserWithdrawals(user common.Address) (IDelayedWithdrawalRouterUserDelayedWithdrawals, error) {
	return _DelayedWithdrawalRouter.Contract.UserWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// UserWithdrawals is a free data retrieval call binding the contract method 0xecb7cb1b.
//
// Solidity: function userWithdrawals(address user) view returns((uint256,(uint224,uint32)[]))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) UserWithdrawals(user common.Address) (IDelayedWithdrawalRouterUserDelayedWithdrawals, error) {
	return _DelayedWithdrawalRouter.Contract.UserWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// UserWithdrawalsLength is a free data retrieval call binding the contract method 0xe4f4f887.
//
// Solidity: function userWithdrawalsLength(address user) view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) UserWithdrawalsLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "userWithdrawalsLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserWithdrawalsLength is a free data retrieval call binding the contract method 0xe4f4f887.
//
// Solidity: function userWithdrawalsLength(address user) view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) UserWithdrawalsLength(user common.Address) (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.UserWithdrawalsLength(&_DelayedWithdrawalRouter.CallOpts, user)
}

// UserWithdrawalsLength is a free data retrieval call binding the contract method 0xe4f4f887.
//
// Solidity: function userWithdrawalsLength(address user) view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) UserWithdrawalsLength(user common.Address) (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.UserWithdrawalsLength(&_DelayedWithdrawalRouter.CallOpts, user)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) WithdrawalDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "withdrawalDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.WithdrawalDelayBlocks(&_DelayedWithdrawalRouter.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.WithdrawalDelayBlocks(&_DelayedWithdrawalRouter.CallOpts)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) ClaimDelayedWithdrawals(opts *bind.TransactOpts, maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "claimDelayedWithdrawals", maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) ClaimDelayedWithdrawals(maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals(&_DelayedWithdrawalRouter.TransactOpts, maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) ClaimDelayedWithdrawals(maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals(&_DelayedWithdrawalRouter.TransactOpts, maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals0 is a paid mutator transaction binding the contract method 0xe5db06c0.
//
// Solidity: function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) ClaimDelayedWithdrawals0(opts *bind.TransactOpts, recipient common.Address, maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "claimDelayedWithdrawals0", recipient, maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals0 is a paid mutator transaction binding the contract method 0xe5db06c0.
//
// Solidity: function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) ClaimDelayedWithdrawals0(recipient common.Address, maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals0(&_DelayedWithdrawalRouter.TransactOpts, recipient, maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals0 is a paid mutator transaction binding the contract method 0xe5db06c0.
//
// Solidity: function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) ClaimDelayedWithdrawals0(recipient common.Address, maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals0(&_DelayedWithdrawalRouter.TransactOpts, recipient, maxNumberOfDelayedWithdrawalsToClaim)
}

// CreateDelayedWithdrawal is a paid mutator transaction binding the contract method 0xc0db354c.
//
// Solidity: function createDelayedWithdrawal(address podOwner, address recipient) payable returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) CreateDelayedWithdrawal(opts *bind.TransactOpts, podOwner common.Address, recipient common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "createDelayedWithdrawal", podOwner, recipient)
}

// CreateDelayedWithdrawal is a paid mutator transaction binding the contract method 0xc0db354c.
//
// Solidity: function createDelayedWithdrawal(address podOwner, address recipient) payable returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) CreateDelayedWithdrawal(podOwner common.Address, recipient common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.CreateDelayedWithdrawal(&_DelayedWithdrawalRouter.TransactOpts, podOwner, recipient)
}

// CreateDelayedWithdrawal is a paid mutator transaction binding the contract method 0xc0db354c.
//
// Solidity: function createDelayedWithdrawal(address podOwner, address recipient) payable returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) CreateDelayedWithdrawal(podOwner common.Address, recipient common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.CreateDelayedWithdrawal(&_DelayedWithdrawalRouter.TransactOpts, podOwner, recipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address initOwner, address _pauserRegistry, uint256 initPausedStatus, uint256 _withdrawalDelayBlocks) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) Initialize(opts *bind.TransactOpts, initOwner common.Address, _pauserRegistry common.Address, initPausedStatus *big.Int, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "initialize", initOwner, _pauserRegistry, initPausedStatus, _withdrawalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address initOwner, address _pauserRegistry, uint256 initPausedStatus, uint256 _withdrawalDelayBlocks) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Initialize(initOwner common.Address, _pauserRegistry common.Address, initPausedStatus *big.Int, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Initialize(&_DelayedWithdrawalRouter.TransactOpts, initOwner, _pauserRegistry, initPausedStatus, _withdrawalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address initOwner, address _pauserRegistry, uint256 initPausedStatus, uint256 _withdrawalDelayBlocks) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) Initialize(initOwner common.Address, _pauserRegistry common.Address, initPausedStatus *big.Int, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Initialize(&_DelayedWithdrawalRouter.TransactOpts, initOwner, _pauserRegistry, initPausedStatus, _withdrawalDelayBlocks)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Pause(&_DelayedWithdrawalRouter.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Pause(&_DelayedWithdrawalRouter.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) PauseAll() (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.PauseAll(&_DelayedWithdrawalRouter.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) PauseAll() (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.PauseAll(&_DelayedWithdrawalRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.RenounceOwnership(&_DelayedWithdrawalRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.RenounceOwnership(&_DelayedWithdrawalRouter.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.SetPauserRegistry(&_DelayedWithdrawalRouter.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.SetPauserRegistry(&_DelayedWithdrawalRouter.TransactOpts, newPauserRegistry)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newValue) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) SetWithdrawalDelayBlocks(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "setWithdrawalDelayBlocks", newValue)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newValue) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) SetWithdrawalDelayBlocks(newValue *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.SetWithdrawalDelayBlocks(&_DelayedWithdrawalRouter.TransactOpts, newValue)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newValue) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) SetWithdrawalDelayBlocks(newValue *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.SetWithdrawalDelayBlocks(&_DelayedWithdrawalRouter.TransactOpts, newValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.TransferOwnership(&_DelayedWithdrawalRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.TransferOwnership(&_DelayedWithdrawalRouter.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Unpause(&_DelayedWithdrawalRouter.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Unpause(&_DelayedWithdrawalRouter.TransactOpts, newPausedStatus)
}

// DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator is returned from FilterDelayedWithdrawalCreated and is used to iterate over the raw logs and unpacked data for DelayedWithdrawalCreated events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator struct {
	Event *DelayedWithdrawalRouterDelayedWithdrawalCreated // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterDelayedWithdrawalCreated)
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
		it.Event = new(DelayedWithdrawalRouterDelayedWithdrawalCreated)
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
func (it *DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterDelayedWithdrawalCreated represents a DelayedWithdrawalCreated event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterDelayedWithdrawalCreated struct {
	PodOwner  common.Address
	Recipient common.Address
	Amount    *big.Int
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayedWithdrawalCreated is a free log retrieval operation binding the contract event 0xb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f59.
//
// Solidity: event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterDelayedWithdrawalCreated(opts *bind.FilterOpts) (*DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "DelayedWithdrawalCreated")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator{contract: _DelayedWithdrawalRouter.contract, event: "DelayedWithdrawalCreated", logs: logs, sub: sub}, nil
}

// WatchDelayedWithdrawalCreated is a free log subscription operation binding the contract event 0xb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f59.
//
// Solidity: event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchDelayedWithdrawalCreated(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterDelayedWithdrawalCreated) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "DelayedWithdrawalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterDelayedWithdrawalCreated)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalCreated", log); err != nil {
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

// ParseDelayedWithdrawalCreated is a log parse operation binding the contract event 0xb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f59.
//
// Solidity: event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseDelayedWithdrawalCreated(log types.Log) (*DelayedWithdrawalRouterDelayedWithdrawalCreated, error) {
	event := new(DelayedWithdrawalRouterDelayedWithdrawalCreated)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator is returned from FilterDelayedWithdrawalsClaimed and is used to iterate over the raw logs and unpacked data for DelayedWithdrawalsClaimed events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator struct {
	Event *DelayedWithdrawalRouterDelayedWithdrawalsClaimed // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterDelayedWithdrawalsClaimed)
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
		it.Event = new(DelayedWithdrawalRouterDelayedWithdrawalsClaimed)
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
func (it *DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterDelayedWithdrawalsClaimed represents a DelayedWithdrawalsClaimed event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterDelayedWithdrawalsClaimed struct {
	Recipient                   common.Address
	AmountClaimed               *big.Int
	DelayedWithdrawalsCompleted *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterDelayedWithdrawalsClaimed is a free log retrieval operation binding the contract event 0x6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943.
//
// Solidity: event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterDelayedWithdrawalsClaimed(opts *bind.FilterOpts) (*DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "DelayedWithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator{contract: _DelayedWithdrawalRouter.contract, event: "DelayedWithdrawalsClaimed", logs: logs, sub: sub}, nil
}

// WatchDelayedWithdrawalsClaimed is a free log subscription operation binding the contract event 0x6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943.
//
// Solidity: event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchDelayedWithdrawalsClaimed(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterDelayedWithdrawalsClaimed) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "DelayedWithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterDelayedWithdrawalsClaimed)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalsClaimed", log); err != nil {
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

// ParseDelayedWithdrawalsClaimed is a log parse operation binding the contract event 0x6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943.
//
// Solidity: event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseDelayedWithdrawalsClaimed(log types.Log) (*DelayedWithdrawalRouterDelayedWithdrawalsClaimed, error) {
	event := new(DelayedWithdrawalRouterDelayedWithdrawalsClaimed)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterInitializedIterator struct {
	Event *DelayedWithdrawalRouterInitialized // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterInitialized)
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
		it.Event = new(DelayedWithdrawalRouterInitialized)
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
func (it *DelayedWithdrawalRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterInitialized represents a Initialized event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*DelayedWithdrawalRouterInitializedIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterInitializedIterator{contract: _DelayedWithdrawalRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterInitialized)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseInitialized(log types.Log) (*DelayedWithdrawalRouterInitialized, error) {
	event := new(DelayedWithdrawalRouterInitialized)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterOwnershipTransferredIterator struct {
	Event *DelayedWithdrawalRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterOwnershipTransferred)
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
		it.Event = new(DelayedWithdrawalRouterOwnershipTransferred)
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
func (it *DelayedWithdrawalRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterOwnershipTransferred represents a OwnershipTransferred event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DelayedWithdrawalRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterOwnershipTransferredIterator{contract: _DelayedWithdrawalRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterOwnershipTransferred)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseOwnershipTransferred(log types.Log) (*DelayedWithdrawalRouterOwnershipTransferred, error) {
	event := new(DelayedWithdrawalRouterOwnershipTransferred)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterPausedIterator struct {
	Event *DelayedWithdrawalRouterPaused // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterPaused)
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
		it.Event = new(DelayedWithdrawalRouterPaused)
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
func (it *DelayedWithdrawalRouterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterPaused represents a Paused event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*DelayedWithdrawalRouterPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterPausedIterator{contract: _DelayedWithdrawalRouter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterPaused)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParsePaused(log types.Log) (*DelayedWithdrawalRouterPaused, error) {
	event := new(DelayedWithdrawalRouterPaused)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterPauserRegistrySetIterator struct {
	Event *DelayedWithdrawalRouterPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterPauserRegistrySet)
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
		it.Event = new(DelayedWithdrawalRouterPauserRegistrySet)
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
func (it *DelayedWithdrawalRouterPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterPauserRegistrySet represents a PauserRegistrySet event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*DelayedWithdrawalRouterPauserRegistrySetIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterPauserRegistrySetIterator{contract: _DelayedWithdrawalRouter.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterPauserRegistrySet)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParsePauserRegistrySet(log types.Log) (*DelayedWithdrawalRouterPauserRegistrySet, error) {
	event := new(DelayedWithdrawalRouterPauserRegistrySet)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterUnpausedIterator struct {
	Event *DelayedWithdrawalRouterUnpaused // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterUnpaused)
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
		it.Event = new(DelayedWithdrawalRouterUnpaused)
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
func (it *DelayedWithdrawalRouterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterUnpaused represents a Unpaused event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*DelayedWithdrawalRouterUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterUnpausedIterator{contract: _DelayedWithdrawalRouter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterUnpaused)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseUnpaused(log types.Log) (*DelayedWithdrawalRouterUnpaused, error) {
	event := new(DelayedWithdrawalRouterUnpaused)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator is returned from FilterWithdrawalDelayBlocksSet and is used to iterate over the raw logs and unpacked data for WithdrawalDelayBlocksSet events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator struct {
	Event *DelayedWithdrawalRouterWithdrawalDelayBlocksSet // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterWithdrawalDelayBlocksSet)
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
		it.Event = new(DelayedWithdrawalRouterWithdrawalDelayBlocksSet)
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
func (it *DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterWithdrawalDelayBlocksSet represents a WithdrawalDelayBlocksSet event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterWithdrawalDelayBlocksSet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayBlocksSet is a free log retrieval operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterWithdrawalDelayBlocksSet(opts *bind.FilterOpts) (*DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator{contract: _DelayedWithdrawalRouter.contract, event: "WithdrawalDelayBlocksSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayBlocksSet is a free log subscription operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchWithdrawalDelayBlocksSet(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterWithdrawalDelayBlocksSet) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterWithdrawalDelayBlocksSet)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
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

// ParseWithdrawalDelayBlocksSet is a log parse operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseWithdrawalDelayBlocksSet(log types.Log) (*DelayedWithdrawalRouterWithdrawalDelayBlocksSet, error) {
	event := new(DelayedWithdrawalRouterWithdrawalDelayBlocksSet)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
