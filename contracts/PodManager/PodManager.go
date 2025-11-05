// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PodManager

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// PodManagerMetaData contains all meta data concerning the PodManager contract.
var PodManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIETHPOSDeposit\",\"name\":\"_ethPOS\",\"type\":\"address\"},{\"internalType\":\"contractIBeacon\",\"name\":\"_eigenPodBeacon\",\"type\":\"address\"},{\"internalType\":\"contractIDelegationManager\",\"name\":\"_delegationManager\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenPodAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStrategy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LegacyWithdrawalsNotCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDelegationManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEigenPod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProofTimestampSetter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SharesNegative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SharesNotMultipleOfGwei\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BeaconChainETHDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"BeaconChainETHWithdrawalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"prevBeaconChainSlashingFactor\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newBeaconChainSlashingFactor\",\"type\":\"uint64\"}],\"name\":\"BeaconChainSlashingFactorDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"BurnableETHSharesIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"newTotalShares\",\"type\":\"int256\"}],\"name\":\"NewTotalShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPectraForkTimestamp\",\"type\":\"uint64\"}],\"name\":\"PectraForkTimestampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eigenPod\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"PodDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"sharesDelta\",\"type\":\"int256\"}],\"name\":\"PodSharesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProofTimestampSetter\",\"type\":\"address\"}],\"name\":\"ProofTimestampSetterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"addShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconChainETHStrategy\",\"outputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"beaconChainSlashingFactor\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnableETHShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createPod\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManager\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPodBeacon\",\"outputs\":[{\"internalType\":\"contractIBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethPOS\",\"outputs\":[{\"internalType\":\"contractIETHPOSDeposit\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"getPod\",\"outputs\":[{\"internalType\":\"contractIEigenPod\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"hasPod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedSharesToBurn\",\"type\":\"uint256\"}],\"name\":\"increaseBurnOrRedistributableShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initPausedStatus\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numPods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"ownerToPod\",\"outputs\":[{\"internalType\":\"contractIEigenPod\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pectraForkTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"podOwnerDepositShares\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"shares\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofTimestampSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevRestakedBalanceWei\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"balanceDeltaWei\",\"type\":\"int256\"}],\"name\":\"recordBeaconChainETHBalanceUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"depositSharesToRemove\",\"type\":\"uint256\"}],\"name\":\"removeDepositShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"setPectraForkTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProofTimestampSetter\",\"type\":\"address\"}],\"name\":\"setProofTimestampSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"depositDataRoot\",\"type\":\"bytes32\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"stakerDepositShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"withdrawSharesAsTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PodManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PodManagerMetaData.ABI instead.
var PodManagerABI = PodManagerMetaData.ABI

// PodManager is an auto generated Go binding around an Ethereum contract.
type PodManager struct {
	PodManagerCaller     // Read-only binding to the contract
	PodManagerTransactor // Write-only binding to the contract
	PodManagerFilterer   // Log filterer for contract events
}

// PodManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PodManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PodManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PodManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PodManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PodManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PodManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PodManagerSession struct {
	Contract     *PodManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PodManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PodManagerCallerSession struct {
	Contract *PodManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PodManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PodManagerTransactorSession struct {
	Contract     *PodManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PodManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PodManagerRaw struct {
	Contract *PodManager // Generic contract binding to access the raw methods on
}

// PodManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PodManagerCallerRaw struct {
	Contract *PodManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PodManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PodManagerTransactorRaw struct {
	Contract *PodManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPodManager creates a new instance of PodManager, bound to a specific deployed contract.
func NewPodManager(address common.Address, backend bind.ContractBackend) (*PodManager, error) {
	contract, err := bindPodManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PodManager{PodManagerCaller: PodManagerCaller{contract: contract}, PodManagerTransactor: PodManagerTransactor{contract: contract}, PodManagerFilterer: PodManagerFilterer{contract: contract}}, nil
}

// NewPodManagerCaller creates a new read-only instance of PodManager, bound to a specific deployed contract.
func NewPodManagerCaller(address common.Address, caller bind.ContractCaller) (*PodManagerCaller, error) {
	contract, err := bindPodManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PodManagerCaller{contract: contract}, nil
}

// NewPodManagerTransactor creates a new write-only instance of PodManager, bound to a specific deployed contract.
func NewPodManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PodManagerTransactor, error) {
	contract, err := bindPodManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PodManagerTransactor{contract: contract}, nil
}

// NewPodManagerFilterer creates a new log filterer instance of PodManager, bound to a specific deployed contract.
func NewPodManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PodManagerFilterer, error) {
	contract, err := bindPodManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PodManagerFilterer{contract: contract}, nil
}

// bindPodManager binds a generic wrapper to an already deployed contract.
func bindPodManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PodManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PodManager *PodManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PodManager.Contract.PodManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PodManager *PodManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.Contract.PodManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PodManager *PodManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PodManager.Contract.PodManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PodManager *PodManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PodManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PodManager *PodManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PodManager *PodManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PodManager.Contract.contract.Transact(opts, method, params...)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_PodManager *PodManagerCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_PodManager *PodManagerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _PodManager.Contract.BeaconChainETHStrategy(&_PodManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_PodManager *PodManagerCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _PodManager.Contract.BeaconChainETHStrategy(&_PodManager.CallOpts)
}

// BeaconChainSlashingFactor is a free data retrieval call binding the contract method 0xa3d75e09.
//
// Solidity: function beaconChainSlashingFactor(address podOwner) view returns(uint64)
func (_PodManager *PodManagerCaller) BeaconChainSlashingFactor(opts *bind.CallOpts, podOwner common.Address) (uint64, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "beaconChainSlashingFactor", podOwner)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BeaconChainSlashingFactor is a free data retrieval call binding the contract method 0xa3d75e09.
//
// Solidity: function beaconChainSlashingFactor(address podOwner) view returns(uint64)
func (_PodManager *PodManagerSession) BeaconChainSlashingFactor(podOwner common.Address) (uint64, error) {
	return _PodManager.Contract.BeaconChainSlashingFactor(&_PodManager.CallOpts, podOwner)
}

// BeaconChainSlashingFactor is a free data retrieval call binding the contract method 0xa3d75e09.
//
// Solidity: function beaconChainSlashingFactor(address podOwner) view returns(uint64)
func (_PodManager *PodManagerCallerSession) BeaconChainSlashingFactor(podOwner common.Address) (uint64, error) {
	return _PodManager.Contract.BeaconChainSlashingFactor(&_PodManager.CallOpts, podOwner)
}

// BurnableETHShares is a free data retrieval call binding the contract method 0xf5d4fed3.
//
// Solidity: function burnableETHShares() view returns(uint256)
func (_PodManager *PodManagerCaller) BurnableETHShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "burnableETHShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnableETHShares is a free data retrieval call binding the contract method 0xf5d4fed3.
//
// Solidity: function burnableETHShares() view returns(uint256)
func (_PodManager *PodManagerSession) BurnableETHShares() (*big.Int, error) {
	return _PodManager.Contract.BurnableETHShares(&_PodManager.CallOpts)
}

// BurnableETHShares is a free data retrieval call binding the contract method 0xf5d4fed3.
//
// Solidity: function burnableETHShares() view returns(uint256)
func (_PodManager *PodManagerCallerSession) BurnableETHShares() (*big.Int, error) {
	return _PodManager.Contract.BurnableETHShares(&_PodManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_PodManager *PodManagerCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_PodManager *PodManagerSession) DelegationManager() (common.Address, error) {
	return _PodManager.Contract.DelegationManager(&_PodManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_PodManager *PodManagerCallerSession) DelegationManager() (common.Address, error) {
	return _PodManager.Contract.DelegationManager(&_PodManager.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_PodManager *PodManagerCaller) EigenPodBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "eigenPodBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_PodManager *PodManagerSession) EigenPodBeacon() (common.Address, error) {
	return _PodManager.Contract.EigenPodBeacon(&_PodManager.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_PodManager *PodManagerCallerSession) EigenPodBeacon() (common.Address, error) {
	return _PodManager.Contract.EigenPodBeacon(&_PodManager.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_PodManager *PodManagerCaller) EthPOS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "ethPOS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_PodManager *PodManagerSession) EthPOS() (common.Address, error) {
	return _PodManager.Contract.EthPOS(&_PodManager.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_PodManager *PodManagerCallerSession) EthPOS() (common.Address, error) {
	return _PodManager.Contract.EthPOS(&_PodManager.CallOpts)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_PodManager *PodManagerCaller) GetPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_PodManager *PodManagerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _PodManager.Contract.GetPod(&_PodManager.CallOpts, podOwner)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_PodManager *PodManagerCallerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _PodManager.Contract.GetPod(&_PodManager.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_PodManager *PodManagerCaller) HasPod(opts *bind.CallOpts, podOwner common.Address) (bool, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "hasPod", podOwner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_PodManager *PodManagerSession) HasPod(podOwner common.Address) (bool, error) {
	return _PodManager.Contract.HasPod(&_PodManager.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_PodManager *PodManagerCallerSession) HasPod(podOwner common.Address) (bool, error) {
	return _PodManager.Contract.HasPod(&_PodManager.CallOpts, podOwner)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_PodManager *PodManagerCaller) NumPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "numPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_PodManager *PodManagerSession) NumPods() (*big.Int, error) {
	return _PodManager.Contract.NumPods(&_PodManager.CallOpts)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_PodManager *PodManagerCallerSession) NumPods() (*big.Int, error) {
	return _PodManager.Contract.NumPods(&_PodManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PodManager *PodManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PodManager *PodManagerSession) Owner() (common.Address, error) {
	return _PodManager.Contract.Owner(&_PodManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PodManager *PodManagerCallerSession) Owner() (common.Address, error) {
	return _PodManager.Contract.Owner(&_PodManager.CallOpts)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address podOwner) view returns(address)
func (_PodManager *PodManagerCaller) OwnerToPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "ownerToPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address podOwner) view returns(address)
func (_PodManager *PodManagerSession) OwnerToPod(podOwner common.Address) (common.Address, error) {
	return _PodManager.Contract.OwnerToPod(&_PodManager.CallOpts, podOwner)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address podOwner) view returns(address)
func (_PodManager *PodManagerCallerSession) OwnerToPod(podOwner common.Address) (common.Address, error) {
	return _PodManager.Contract.OwnerToPod(&_PodManager.CallOpts, podOwner)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PodManager *PodManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PodManager *PodManagerSession) Paused(index uint8) (bool, error) {
	return _PodManager.Contract.Paused(&_PodManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PodManager *PodManagerCallerSession) Paused(index uint8) (bool, error) {
	return _PodManager.Contract.Paused(&_PodManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PodManager *PodManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PodManager *PodManagerSession) Paused0() (*big.Int, error) {
	return _PodManager.Contract.Paused0(&_PodManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PodManager *PodManagerCallerSession) Paused0() (*big.Int, error) {
	return _PodManager.Contract.Paused0(&_PodManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PodManager *PodManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PodManager *PodManagerSession) PauserRegistry() (common.Address, error) {
	return _PodManager.Contract.PauserRegistry(&_PodManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PodManager *PodManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _PodManager.Contract.PauserRegistry(&_PodManager.CallOpts)
}

// PectraForkTimestamp is a free data retrieval call binding the contract method 0x2704351a.
//
// Solidity: function pectraForkTimestamp() view returns(uint64)
func (_PodManager *PodManagerCaller) PectraForkTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "pectraForkTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PectraForkTimestamp is a free data retrieval call binding the contract method 0x2704351a.
//
// Solidity: function pectraForkTimestamp() view returns(uint64)
func (_PodManager *PodManagerSession) PectraForkTimestamp() (uint64, error) {
	return _PodManager.Contract.PectraForkTimestamp(&_PodManager.CallOpts)
}

// PectraForkTimestamp is a free data retrieval call binding the contract method 0x2704351a.
//
// Solidity: function pectraForkTimestamp() view returns(uint64)
func (_PodManager *PodManagerCallerSession) PectraForkTimestamp() (uint64, error) {
	return _PodManager.Contract.PectraForkTimestamp(&_PodManager.CallOpts)
}

// PodOwnerDepositShares is a free data retrieval call binding the contract method 0xd48e8894.
//
// Solidity: function podOwnerDepositShares(address podOwner) view returns(int256 shares)
func (_PodManager *PodManagerCaller) PodOwnerDepositShares(opts *bind.CallOpts, podOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "podOwnerDepositShares", podOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PodOwnerDepositShares is a free data retrieval call binding the contract method 0xd48e8894.
//
// Solidity: function podOwnerDepositShares(address podOwner) view returns(int256 shares)
func (_PodManager *PodManagerSession) PodOwnerDepositShares(podOwner common.Address) (*big.Int, error) {
	return _PodManager.Contract.PodOwnerDepositShares(&_PodManager.CallOpts, podOwner)
}

// PodOwnerDepositShares is a free data retrieval call binding the contract method 0xd48e8894.
//
// Solidity: function podOwnerDepositShares(address podOwner) view returns(int256 shares)
func (_PodManager *PodManagerCallerSession) PodOwnerDepositShares(podOwner common.Address) (*big.Int, error) {
	return _PodManager.Contract.PodOwnerDepositShares(&_PodManager.CallOpts, podOwner)
}

// ProofTimestampSetter is a free data retrieval call binding the contract method 0x595edbcb.
//
// Solidity: function proofTimestampSetter() view returns(address)
func (_PodManager *PodManagerCaller) ProofTimestampSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "proofTimestampSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProofTimestampSetter is a free data retrieval call binding the contract method 0x595edbcb.
//
// Solidity: function proofTimestampSetter() view returns(address)
func (_PodManager *PodManagerSession) ProofTimestampSetter() (common.Address, error) {
	return _PodManager.Contract.ProofTimestampSetter(&_PodManager.CallOpts)
}

// ProofTimestampSetter is a free data retrieval call binding the contract method 0x595edbcb.
//
// Solidity: function proofTimestampSetter() view returns(address)
func (_PodManager *PodManagerCallerSession) ProofTimestampSetter() (common.Address, error) {
	return _PodManager.Contract.ProofTimestampSetter(&_PodManager.CallOpts)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 depositShares)
func (_PodManager *PodManagerCaller) StakerDepositShares(opts *bind.CallOpts, user common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "stakerDepositShares", user, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 depositShares)
func (_PodManager *PodManagerSession) StakerDepositShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _PodManager.Contract.StakerDepositShares(&_PodManager.CallOpts, user, strategy)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 depositShares)
func (_PodManager *PodManagerCallerSession) StakerDepositShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _PodManager.Contract.StakerDepositShares(&_PodManager.CallOpts, user, strategy)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_PodManager *PodManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_PodManager *PodManagerSession) Version() (string, error) {
	return _PodManager.Contract.Version(&_PodManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_PodManager *PodManagerCallerSession) Version() (string, error) {
	return _PodManager.Contract.Version(&_PodManager.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_PodManager *PodManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "addShares", staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_PodManager *PodManagerSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.AddShares(&_PodManager.TransactOpts, staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_PodManager *PodManagerTransactorSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.AddShares(&_PodManager.TransactOpts, staker, strategy, shares)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_PodManager *PodManagerTransactor) CreatePod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "createPod")
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_PodManager *PodManagerSession) CreatePod() (*types.Transaction, error) {
	return _PodManager.Contract.CreatePod(&_PodManager.TransactOpts)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_PodManager *PodManagerTransactorSession) CreatePod() (*types.Transaction, error) {
	return _PodManager.Contract.CreatePod(&_PodManager.TransactOpts)
}

// IncreaseBurnOrRedistributableShares is a paid mutator transaction binding the contract method 0x3fb99ca5.
//
// Solidity: function increaseBurnOrRedistributableShares((address,uint32) , uint256 , address , uint256 addedSharesToBurn) returns()
func (_PodManager *PodManagerTransactor) IncreaseBurnOrRedistributableShares(opts *bind.TransactOpts, arg0 OperatorSet, arg1 *big.Int, arg2 common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "increaseBurnOrRedistributableShares", arg0, arg1, arg2, addedSharesToBurn)
}

// IncreaseBurnOrRedistributableShares is a paid mutator transaction binding the contract method 0x3fb99ca5.
//
// Solidity: function increaseBurnOrRedistributableShares((address,uint32) , uint256 , address , uint256 addedSharesToBurn) returns()
func (_PodManager *PodManagerSession) IncreaseBurnOrRedistributableShares(arg0 OperatorSet, arg1 *big.Int, arg2 common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.IncreaseBurnOrRedistributableShares(&_PodManager.TransactOpts, arg0, arg1, arg2, addedSharesToBurn)
}

// IncreaseBurnOrRedistributableShares is a paid mutator transaction binding the contract method 0x3fb99ca5.
//
// Solidity: function increaseBurnOrRedistributableShares((address,uint32) , uint256 , address , uint256 addedSharesToBurn) returns()
func (_PodManager *PodManagerTransactorSession) IncreaseBurnOrRedistributableShares(arg0 OperatorSet, arg1 *big.Int, arg2 common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.IncreaseBurnOrRedistributableShares(&_PodManager.TransactOpts, arg0, arg1, arg2, addedSharesToBurn)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 _initPausedStatus) returns()
func (_PodManager *PodManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _initPausedStatus *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "initialize", initialOwner, _initPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 _initPausedStatus) returns()
func (_PodManager *PodManagerSession) Initialize(initialOwner common.Address, _initPausedStatus *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.Initialize(&_PodManager.TransactOpts, initialOwner, _initPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 _initPausedStatus) returns()
func (_PodManager *PodManagerTransactorSession) Initialize(initialOwner common.Address, _initPausedStatus *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.Initialize(&_PodManager.TransactOpts, initialOwner, _initPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PodManager *PodManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PodManager *PodManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.Pause(&_PodManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PodManager *PodManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.Pause(&_PodManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PodManager *PodManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PodManager *PodManagerSession) PauseAll() (*types.Transaction, error) {
	return _PodManager.Contract.PauseAll(&_PodManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PodManager *PodManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _PodManager.Contract.PauseAll(&_PodManager.TransactOpts)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xa1ca780b.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 prevRestakedBalanceWei, int256 balanceDeltaWei) returns()
func (_PodManager *PodManagerTransactor) RecordBeaconChainETHBalanceUpdate(opts *bind.TransactOpts, podOwner common.Address, prevRestakedBalanceWei *big.Int, balanceDeltaWei *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "recordBeaconChainETHBalanceUpdate", podOwner, prevRestakedBalanceWei, balanceDeltaWei)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xa1ca780b.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 prevRestakedBalanceWei, int256 balanceDeltaWei) returns()
func (_PodManager *PodManagerSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, prevRestakedBalanceWei *big.Int, balanceDeltaWei *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_PodManager.TransactOpts, podOwner, prevRestakedBalanceWei, balanceDeltaWei)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xa1ca780b.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 prevRestakedBalanceWei, int256 balanceDeltaWei) returns()
func (_PodManager *PodManagerTransactorSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, prevRestakedBalanceWei *big.Int, balanceDeltaWei *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_PodManager.TransactOpts, podOwner, prevRestakedBalanceWei, balanceDeltaWei)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_PodManager *PodManagerTransactor) RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "removeDepositShares", staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_PodManager *PodManagerSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.RemoveDepositShares(&_PodManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_PodManager *PodManagerTransactorSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.RemoveDepositShares(&_PodManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PodManager *PodManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PodManager *PodManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PodManager.Contract.RenounceOwnership(&_PodManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PodManager *PodManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PodManager.Contract.RenounceOwnership(&_PodManager.TransactOpts)
}

// SetPectraForkTimestamp is a paid mutator transaction binding the contract method 0x5a26fbf4.
//
// Solidity: function setPectraForkTimestamp(uint64 timestamp) returns()
func (_PodManager *PodManagerTransactor) SetPectraForkTimestamp(opts *bind.TransactOpts, timestamp uint64) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "setPectraForkTimestamp", timestamp)
}

// SetPectraForkTimestamp is a paid mutator transaction binding the contract method 0x5a26fbf4.
//
// Solidity: function setPectraForkTimestamp(uint64 timestamp) returns()
func (_PodManager *PodManagerSession) SetPectraForkTimestamp(timestamp uint64) (*types.Transaction, error) {
	return _PodManager.Contract.SetPectraForkTimestamp(&_PodManager.TransactOpts, timestamp)
}

// SetPectraForkTimestamp is a paid mutator transaction binding the contract method 0x5a26fbf4.
//
// Solidity: function setPectraForkTimestamp(uint64 timestamp) returns()
func (_PodManager *PodManagerTransactorSession) SetPectraForkTimestamp(timestamp uint64) (*types.Transaction, error) {
	return _PodManager.Contract.SetPectraForkTimestamp(&_PodManager.TransactOpts, timestamp)
}

// SetProofTimestampSetter is a paid mutator transaction binding the contract method 0x0d1e9de1.
//
// Solidity: function setProofTimestampSetter(address newProofTimestampSetter) returns()
func (_PodManager *PodManagerTransactor) SetProofTimestampSetter(opts *bind.TransactOpts, newProofTimestampSetter common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "setProofTimestampSetter", newProofTimestampSetter)
}

// SetProofTimestampSetter is a paid mutator transaction binding the contract method 0x0d1e9de1.
//
// Solidity: function setProofTimestampSetter(address newProofTimestampSetter) returns()
func (_PodManager *PodManagerSession) SetProofTimestampSetter(newProofTimestampSetter common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.SetProofTimestampSetter(&_PodManager.TransactOpts, newProofTimestampSetter)
}

// SetProofTimestampSetter is a paid mutator transaction binding the contract method 0x0d1e9de1.
//
// Solidity: function setProofTimestampSetter(address newProofTimestampSetter) returns()
func (_PodManager *PodManagerTransactorSession) SetProofTimestampSetter(newProofTimestampSetter common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.SetProofTimestampSetter(&_PodManager.TransactOpts, newProofTimestampSetter)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_PodManager *PodManagerTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_PodManager *PodManagerSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _PodManager.Contract.Stake(&_PodManager.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_PodManager *PodManagerTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _PodManager.Contract.Stake(&_PodManager.TransactOpts, pubkey, signature, depositDataRoot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PodManager *PodManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PodManager *PodManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.TransferOwnership(&_PodManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PodManager *PodManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.TransferOwnership(&_PodManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PodManager *PodManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PodManager *PodManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.Unpause(&_PodManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PodManager *PodManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.Unpause(&_PodManager.TransactOpts, newPausedStatus)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address , uint256 shares) returns()
func (_PodManager *PodManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "withdrawSharesAsTokens", staker, strategy, arg2, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address , uint256 shares) returns()
func (_PodManager *PodManagerSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawSharesAsTokens(&_PodManager.TransactOpts, staker, strategy, arg2, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address , uint256 shares) returns()
func (_PodManager *PodManagerTransactorSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawSharesAsTokens(&_PodManager.TransactOpts, staker, strategy, arg2, shares)
}

// PodManagerBeaconChainETHDepositedIterator is returned from FilterBeaconChainETHDeposited and is used to iterate over the raw logs and unpacked data for BeaconChainETHDeposited events raised by the PodManager contract.
type PodManagerBeaconChainETHDepositedIterator struct {
	Event *PodManagerBeaconChainETHDeposited // Event containing the contract specifics and raw log

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
func (it *PodManagerBeaconChainETHDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBeaconChainETHDeposited)
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
		it.Event = new(PodManagerBeaconChainETHDeposited)
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
func (it *PodManagerBeaconChainETHDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBeaconChainETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBeaconChainETHDeposited represents a BeaconChainETHDeposited event raised by the PodManager contract.
type PodManagerBeaconChainETHDeposited struct {
	PodOwner common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainETHDeposited is a free log retrieval operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_PodManager *PodManagerFilterer) FilterBeaconChainETHDeposited(opts *bind.FilterOpts, podOwner []common.Address) (*PodManagerBeaconChainETHDepositedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerBeaconChainETHDepositedIterator{contract: _PodManager.contract, event: "BeaconChainETHDeposited", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHDeposited is a free log subscription operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_PodManager *PodManagerFilterer) WatchBeaconChainETHDeposited(opts *bind.WatchOpts, sink chan<- *PodManagerBeaconChainETHDeposited, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBeaconChainETHDeposited)
				if err := _PodManager.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
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

// ParseBeaconChainETHDeposited is a log parse operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_PodManager *PodManagerFilterer) ParseBeaconChainETHDeposited(log types.Log) (*PodManagerBeaconChainETHDeposited, error) {
	event := new(PodManagerBeaconChainETHDeposited)
	if err := _PodManager.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerBeaconChainETHWithdrawalCompletedIterator is returned from FilterBeaconChainETHWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for BeaconChainETHWithdrawalCompleted events raised by the PodManager contract.
type PodManagerBeaconChainETHWithdrawalCompletedIterator struct {
	Event *PodManagerBeaconChainETHWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *PodManagerBeaconChainETHWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBeaconChainETHWithdrawalCompleted)
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
		it.Event = new(PodManagerBeaconChainETHWithdrawalCompleted)
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
func (it *PodManagerBeaconChainETHWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBeaconChainETHWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBeaconChainETHWithdrawalCompleted represents a BeaconChainETHWithdrawalCompleted event raised by the PodManager contract.
type PodManagerBeaconChainETHWithdrawalCompleted struct {
	PodOwner         common.Address
	Shares           *big.Int
	Nonce            *big.Int
	DelegatedAddress common.Address
	Withdrawer       common.Address
	WithdrawalRoot   [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainETHWithdrawalCompleted is a free log retrieval operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_PodManager *PodManagerFilterer) FilterBeaconChainETHWithdrawalCompleted(opts *bind.FilterOpts, podOwner []common.Address) (*PodManagerBeaconChainETHWithdrawalCompletedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerBeaconChainETHWithdrawalCompletedIterator{contract: _PodManager.contract, event: "BeaconChainETHWithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHWithdrawalCompleted is a free log subscription operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_PodManager *PodManagerFilterer) WatchBeaconChainETHWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *PodManagerBeaconChainETHWithdrawalCompleted, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBeaconChainETHWithdrawalCompleted)
				if err := _PodManager.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
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

// ParseBeaconChainETHWithdrawalCompleted is a log parse operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_PodManager *PodManagerFilterer) ParseBeaconChainETHWithdrawalCompleted(log types.Log) (*PodManagerBeaconChainETHWithdrawalCompleted, error) {
	event := new(PodManagerBeaconChainETHWithdrawalCompleted)
	if err := _PodManager.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerBeaconChainSlashingFactorDecreasedIterator is returned from FilterBeaconChainSlashingFactorDecreased and is used to iterate over the raw logs and unpacked data for BeaconChainSlashingFactorDecreased events raised by the PodManager contract.
type PodManagerBeaconChainSlashingFactorDecreasedIterator struct {
	Event *PodManagerBeaconChainSlashingFactorDecreased // Event containing the contract specifics and raw log

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
func (it *PodManagerBeaconChainSlashingFactorDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBeaconChainSlashingFactorDecreased)
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
		it.Event = new(PodManagerBeaconChainSlashingFactorDecreased)
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
func (it *PodManagerBeaconChainSlashingFactorDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBeaconChainSlashingFactorDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBeaconChainSlashingFactorDecreased represents a BeaconChainSlashingFactorDecreased event raised by the PodManager contract.
type PodManagerBeaconChainSlashingFactorDecreased struct {
	Staker                        common.Address
	PrevBeaconChainSlashingFactor uint64
	NewBeaconChainSlashingFactor  uint64
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainSlashingFactorDecreased is a free log retrieval operation binding the contract event 0xb160ab8589bf47dc04ea11b50d46678d21590cea2ed3e454e7bd3e41510f98cf.
//
// Solidity: event BeaconChainSlashingFactorDecreased(address staker, uint64 prevBeaconChainSlashingFactor, uint64 newBeaconChainSlashingFactor)
func (_PodManager *PodManagerFilterer) FilterBeaconChainSlashingFactorDecreased(opts *bind.FilterOpts) (*PodManagerBeaconChainSlashingFactorDecreasedIterator, error) {

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BeaconChainSlashingFactorDecreased")
	if err != nil {
		return nil, err
	}
	return &PodManagerBeaconChainSlashingFactorDecreasedIterator{contract: _PodManager.contract, event: "BeaconChainSlashingFactorDecreased", logs: logs, sub: sub}, nil
}

// WatchBeaconChainSlashingFactorDecreased is a free log subscription operation binding the contract event 0xb160ab8589bf47dc04ea11b50d46678d21590cea2ed3e454e7bd3e41510f98cf.
//
// Solidity: event BeaconChainSlashingFactorDecreased(address staker, uint64 prevBeaconChainSlashingFactor, uint64 newBeaconChainSlashingFactor)
func (_PodManager *PodManagerFilterer) WatchBeaconChainSlashingFactorDecreased(opts *bind.WatchOpts, sink chan<- *PodManagerBeaconChainSlashingFactorDecreased) (event.Subscription, error) {

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BeaconChainSlashingFactorDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBeaconChainSlashingFactorDecreased)
				if err := _PodManager.contract.UnpackLog(event, "BeaconChainSlashingFactorDecreased", log); err != nil {
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

// ParseBeaconChainSlashingFactorDecreased is a log parse operation binding the contract event 0xb160ab8589bf47dc04ea11b50d46678d21590cea2ed3e454e7bd3e41510f98cf.
//
// Solidity: event BeaconChainSlashingFactorDecreased(address staker, uint64 prevBeaconChainSlashingFactor, uint64 newBeaconChainSlashingFactor)
func (_PodManager *PodManagerFilterer) ParseBeaconChainSlashingFactorDecreased(log types.Log) (*PodManagerBeaconChainSlashingFactorDecreased, error) {
	event := new(PodManagerBeaconChainSlashingFactorDecreased)
	if err := _PodManager.contract.UnpackLog(event, "BeaconChainSlashingFactorDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerBurnableETHSharesIncreasedIterator is returned from FilterBurnableETHSharesIncreased and is used to iterate over the raw logs and unpacked data for BurnableETHSharesIncreased events raised by the PodManager contract.
type PodManagerBurnableETHSharesIncreasedIterator struct {
	Event *PodManagerBurnableETHSharesIncreased // Event containing the contract specifics and raw log

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
func (it *PodManagerBurnableETHSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBurnableETHSharesIncreased)
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
		it.Event = new(PodManagerBurnableETHSharesIncreased)
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
func (it *PodManagerBurnableETHSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBurnableETHSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBurnableETHSharesIncreased represents a BurnableETHSharesIncreased event raised by the PodManager contract.
type PodManagerBurnableETHSharesIncreased struct {
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnableETHSharesIncreased is a free log retrieval operation binding the contract event 0x1ed04b7fd262c0d9e50fa02957f32a81a151f03baaa367faeedc7521b001c4a4.
//
// Solidity: event BurnableETHSharesIncreased(uint256 shares)
func (_PodManager *PodManagerFilterer) FilterBurnableETHSharesIncreased(opts *bind.FilterOpts) (*PodManagerBurnableETHSharesIncreasedIterator, error) {

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BurnableETHSharesIncreased")
	if err != nil {
		return nil, err
	}
	return &PodManagerBurnableETHSharesIncreasedIterator{contract: _PodManager.contract, event: "BurnableETHSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchBurnableETHSharesIncreased is a free log subscription operation binding the contract event 0x1ed04b7fd262c0d9e50fa02957f32a81a151f03baaa367faeedc7521b001c4a4.
//
// Solidity: event BurnableETHSharesIncreased(uint256 shares)
func (_PodManager *PodManagerFilterer) WatchBurnableETHSharesIncreased(opts *bind.WatchOpts, sink chan<- *PodManagerBurnableETHSharesIncreased) (event.Subscription, error) {

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BurnableETHSharesIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBurnableETHSharesIncreased)
				if err := _PodManager.contract.UnpackLog(event, "BurnableETHSharesIncreased", log); err != nil {
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

// ParseBurnableETHSharesIncreased is a log parse operation binding the contract event 0x1ed04b7fd262c0d9e50fa02957f32a81a151f03baaa367faeedc7521b001c4a4.
//
// Solidity: event BurnableETHSharesIncreased(uint256 shares)
func (_PodManager *PodManagerFilterer) ParseBurnableETHSharesIncreased(log types.Log) (*PodManagerBurnableETHSharesIncreased, error) {
	event := new(PodManagerBurnableETHSharesIncreased)
	if err := _PodManager.contract.UnpackLog(event, "BurnableETHSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PodManager contract.
type PodManagerInitializedIterator struct {
	Event *PodManagerInitialized // Event containing the contract specifics and raw log

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
func (it *PodManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerInitialized)
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
		it.Event = new(PodManagerInitialized)
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
func (it *PodManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerInitialized represents a Initialized event raised by the PodManager contract.
type PodManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PodManager *PodManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PodManagerInitializedIterator, error) {

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PodManagerInitializedIterator{contract: _PodManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PodManager *PodManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PodManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerInitialized)
				if err := _PodManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PodManager *PodManagerFilterer) ParseInitialized(log types.Log) (*PodManagerInitialized, error) {
	event := new(PodManagerInitialized)
	if err := _PodManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerNewTotalSharesIterator is returned from FilterNewTotalShares and is used to iterate over the raw logs and unpacked data for NewTotalShares events raised by the PodManager contract.
type PodManagerNewTotalSharesIterator struct {
	Event *PodManagerNewTotalShares // Event containing the contract specifics and raw log

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
func (it *PodManagerNewTotalSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerNewTotalShares)
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
		it.Event = new(PodManagerNewTotalShares)
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
func (it *PodManagerNewTotalSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerNewTotalSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerNewTotalShares represents a NewTotalShares event raised by the PodManager contract.
type PodManagerNewTotalShares struct {
	PodOwner       common.Address
	NewTotalShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewTotalShares is a free log retrieval operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_PodManager *PodManagerFilterer) FilterNewTotalShares(opts *bind.FilterOpts, podOwner []common.Address) (*PodManagerNewTotalSharesIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "NewTotalShares", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerNewTotalSharesIterator{contract: _PodManager.contract, event: "NewTotalShares", logs: logs, sub: sub}, nil
}

// WatchNewTotalShares is a free log subscription operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_PodManager *PodManagerFilterer) WatchNewTotalShares(opts *bind.WatchOpts, sink chan<- *PodManagerNewTotalShares, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "NewTotalShares", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerNewTotalShares)
				if err := _PodManager.contract.UnpackLog(event, "NewTotalShares", log); err != nil {
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

// ParseNewTotalShares is a log parse operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_PodManager *PodManagerFilterer) ParseNewTotalShares(log types.Log) (*PodManagerNewTotalShares, error) {
	event := new(PodManagerNewTotalShares)
	if err := _PodManager.contract.UnpackLog(event, "NewTotalShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PodManager contract.
type PodManagerOwnershipTransferredIterator struct {
	Event *PodManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PodManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerOwnershipTransferred)
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
		it.Event = new(PodManagerOwnershipTransferred)
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
func (it *PodManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PodManager contract.
type PodManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PodManager *PodManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PodManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerOwnershipTransferredIterator{contract: _PodManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PodManager *PodManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PodManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerOwnershipTransferred)
				if err := _PodManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PodManager *PodManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PodManagerOwnershipTransferred, error) {
	event := new(PodManagerOwnershipTransferred)
	if err := _PodManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PodManager contract.
type PodManagerPausedIterator struct {
	Event *PodManagerPaused // Event containing the contract specifics and raw log

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
func (it *PodManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerPaused)
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
		it.Event = new(PodManagerPaused)
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
func (it *PodManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerPaused represents a Paused event raised by the PodManager contract.
type PodManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_PodManager *PodManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*PodManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerPausedIterator{contract: _PodManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_PodManager *PodManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PodManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerPaused)
				if err := _PodManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PodManager *PodManagerFilterer) ParsePaused(log types.Log) (*PodManagerPaused, error) {
	event := new(PodManagerPaused)
	if err := _PodManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerPectraForkTimestampSetIterator is returned from FilterPectraForkTimestampSet and is used to iterate over the raw logs and unpacked data for PectraForkTimestampSet events raised by the PodManager contract.
type PodManagerPectraForkTimestampSetIterator struct {
	Event *PodManagerPectraForkTimestampSet // Event containing the contract specifics and raw log

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
func (it *PodManagerPectraForkTimestampSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerPectraForkTimestampSet)
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
		it.Event = new(PodManagerPectraForkTimestampSet)
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
func (it *PodManagerPectraForkTimestampSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerPectraForkTimestampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerPectraForkTimestampSet represents a PectraForkTimestampSet event raised by the PodManager contract.
type PodManagerPectraForkTimestampSet struct {
	NewPectraForkTimestamp uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPectraForkTimestampSet is a free log retrieval operation binding the contract event 0x1bc8f042a52db3a437620dea4548f2031fb2a16dd8d3b0b854295528dd2cdd33.
//
// Solidity: event PectraForkTimestampSet(uint64 newPectraForkTimestamp)
func (_PodManager *PodManagerFilterer) FilterPectraForkTimestampSet(opts *bind.FilterOpts) (*PodManagerPectraForkTimestampSetIterator, error) {

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "PectraForkTimestampSet")
	if err != nil {
		return nil, err
	}
	return &PodManagerPectraForkTimestampSetIterator{contract: _PodManager.contract, event: "PectraForkTimestampSet", logs: logs, sub: sub}, nil
}

// WatchPectraForkTimestampSet is a free log subscription operation binding the contract event 0x1bc8f042a52db3a437620dea4548f2031fb2a16dd8d3b0b854295528dd2cdd33.
//
// Solidity: event PectraForkTimestampSet(uint64 newPectraForkTimestamp)
func (_PodManager *PodManagerFilterer) WatchPectraForkTimestampSet(opts *bind.WatchOpts, sink chan<- *PodManagerPectraForkTimestampSet) (event.Subscription, error) {

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "PectraForkTimestampSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerPectraForkTimestampSet)
				if err := _PodManager.contract.UnpackLog(event, "PectraForkTimestampSet", log); err != nil {
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

// ParsePectraForkTimestampSet is a log parse operation binding the contract event 0x1bc8f042a52db3a437620dea4548f2031fb2a16dd8d3b0b854295528dd2cdd33.
//
// Solidity: event PectraForkTimestampSet(uint64 newPectraForkTimestamp)
func (_PodManager *PodManagerFilterer) ParsePectraForkTimestampSet(log types.Log) (*PodManagerPectraForkTimestampSet, error) {
	event := new(PodManagerPectraForkTimestampSet)
	if err := _PodManager.contract.UnpackLog(event, "PectraForkTimestampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerPodDeployedIterator is returned from FilterPodDeployed and is used to iterate over the raw logs and unpacked data for PodDeployed events raised by the PodManager contract.
type PodManagerPodDeployedIterator struct {
	Event *PodManagerPodDeployed // Event containing the contract specifics and raw log

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
func (it *PodManagerPodDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerPodDeployed)
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
		it.Event = new(PodManagerPodDeployed)
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
func (it *PodManagerPodDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerPodDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerPodDeployed represents a PodDeployed event raised by the PodManager contract.
type PodManagerPodDeployed struct {
	EigenPod common.Address
	PodOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPodDeployed is a free log retrieval operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_PodManager *PodManagerFilterer) FilterPodDeployed(opts *bind.FilterOpts, eigenPod []common.Address, podOwner []common.Address) (*PodManagerPodDeployedIterator, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerPodDeployedIterator{contract: _PodManager.contract, event: "PodDeployed", logs: logs, sub: sub}, nil
}

// WatchPodDeployed is a free log subscription operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_PodManager *PodManagerFilterer) WatchPodDeployed(opts *bind.WatchOpts, sink chan<- *PodManagerPodDeployed, eigenPod []common.Address, podOwner []common.Address) (event.Subscription, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerPodDeployed)
				if err := _PodManager.contract.UnpackLog(event, "PodDeployed", log); err != nil {
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

// ParsePodDeployed is a log parse operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_PodManager *PodManagerFilterer) ParsePodDeployed(log types.Log) (*PodManagerPodDeployed, error) {
	event := new(PodManagerPodDeployed)
	if err := _PodManager.contract.UnpackLog(event, "PodDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerPodSharesUpdatedIterator is returned from FilterPodSharesUpdated and is used to iterate over the raw logs and unpacked data for PodSharesUpdated events raised by the PodManager contract.
type PodManagerPodSharesUpdatedIterator struct {
	Event *PodManagerPodSharesUpdated // Event containing the contract specifics and raw log

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
func (it *PodManagerPodSharesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerPodSharesUpdated)
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
		it.Event = new(PodManagerPodSharesUpdated)
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
func (it *PodManagerPodSharesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerPodSharesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerPodSharesUpdated represents a PodSharesUpdated event raised by the PodManager contract.
type PodManagerPodSharesUpdated struct {
	PodOwner    common.Address
	SharesDelta *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPodSharesUpdated is a free log retrieval operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_PodManager *PodManagerFilterer) FilterPodSharesUpdated(opts *bind.FilterOpts, podOwner []common.Address) (*PodManagerPodSharesUpdatedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerPodSharesUpdatedIterator{contract: _PodManager.contract, event: "PodSharesUpdated", logs: logs, sub: sub}, nil
}

// WatchPodSharesUpdated is a free log subscription operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_PodManager *PodManagerFilterer) WatchPodSharesUpdated(opts *bind.WatchOpts, sink chan<- *PodManagerPodSharesUpdated, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerPodSharesUpdated)
				if err := _PodManager.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
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

// ParsePodSharesUpdated is a log parse operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_PodManager *PodManagerFilterer) ParsePodSharesUpdated(log types.Log) (*PodManagerPodSharesUpdated, error) {
	event := new(PodManagerPodSharesUpdated)
	if err := _PodManager.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerProofTimestampSetterSetIterator is returned from FilterProofTimestampSetterSet and is used to iterate over the raw logs and unpacked data for ProofTimestampSetterSet events raised by the PodManager contract.
type PodManagerProofTimestampSetterSetIterator struct {
	Event *PodManagerProofTimestampSetterSet // Event containing the contract specifics and raw log

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
func (it *PodManagerProofTimestampSetterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerProofTimestampSetterSet)
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
		it.Event = new(PodManagerProofTimestampSetterSet)
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
func (it *PodManagerProofTimestampSetterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerProofTimestampSetterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerProofTimestampSetterSet represents a ProofTimestampSetterSet event raised by the PodManager contract.
type PodManagerProofTimestampSetterSet struct {
	NewProofTimestampSetter common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterProofTimestampSetterSet is a free log retrieval operation binding the contract event 0x7025c71a9fe60d709e71b377dc5f7c72c3e1d8539f8022574254e736ceca01e5.
//
// Solidity: event ProofTimestampSetterSet(address newProofTimestampSetter)
func (_PodManager *PodManagerFilterer) FilterProofTimestampSetterSet(opts *bind.FilterOpts) (*PodManagerProofTimestampSetterSetIterator, error) {

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "ProofTimestampSetterSet")
	if err != nil {
		return nil, err
	}
	return &PodManagerProofTimestampSetterSetIterator{contract: _PodManager.contract, event: "ProofTimestampSetterSet", logs: logs, sub: sub}, nil
}

// WatchProofTimestampSetterSet is a free log subscription operation binding the contract event 0x7025c71a9fe60d709e71b377dc5f7c72c3e1d8539f8022574254e736ceca01e5.
//
// Solidity: event ProofTimestampSetterSet(address newProofTimestampSetter)
func (_PodManager *PodManagerFilterer) WatchProofTimestampSetterSet(opts *bind.WatchOpts, sink chan<- *PodManagerProofTimestampSetterSet) (event.Subscription, error) {

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "ProofTimestampSetterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerProofTimestampSetterSet)
				if err := _PodManager.contract.UnpackLog(event, "ProofTimestampSetterSet", log); err != nil {
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

// ParseProofTimestampSetterSet is a log parse operation binding the contract event 0x7025c71a9fe60d709e71b377dc5f7c72c3e1d8539f8022574254e736ceca01e5.
//
// Solidity: event ProofTimestampSetterSet(address newProofTimestampSetter)
func (_PodManager *PodManagerFilterer) ParseProofTimestampSetterSet(log types.Log) (*PodManagerProofTimestampSetterSet, error) {
	event := new(PodManagerProofTimestampSetterSet)
	if err := _PodManager.contract.UnpackLog(event, "ProofTimestampSetterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PodManager contract.
type PodManagerUnpausedIterator struct {
	Event *PodManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *PodManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerUnpaused)
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
		it.Event = new(PodManagerUnpaused)
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
func (it *PodManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerUnpaused represents a Unpaused event raised by the PodManager contract.
type PodManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_PodManager *PodManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*PodManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerUnpausedIterator{contract: _PodManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_PodManager *PodManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PodManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerUnpaused)
				if err := _PodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PodManager *PodManagerFilterer) ParseUnpaused(log types.Log) (*PodManagerUnpaused, error) {
	event := new(PodManagerUnpaused)
	if err := _PodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
