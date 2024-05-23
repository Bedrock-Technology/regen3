// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Staking

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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"BalanceSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creditor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountEther\",\"type\":\"uint256\"}],\"name\":\"DebtQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DepositContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerAccountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"milli\",\"type\":\"uint256\"}],\"name\":\"ManagerFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ManagerFeeWithdrawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ManagerRevenueCompounded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RestakingAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RevenueAccounted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UserRevenueCompounded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextValidatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stoppedCount\",\"type\":\"uint256\"}],\"name\":\"ValidatorSlashedStopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stoppedCount\",\"type\":\"uint256\"}],\"name\":\"ValidatorStopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"WhiteListToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawCredential\",\"type\":\"bytes32\"}],\"name\":\"WithdrawCredentialSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SAFE_PUSH_REWARDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"checkDebt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDepositContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountedBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountedManagerRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountedUserRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDebtQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingEthers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecentReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecentStopped\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idx_to\",\"type\":\"uint256\"}],\"name\":\"getRegisteredValidators\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bool[]\",\"name\":\"stopped\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idx_to\",\"type\":\"uint256\"}],\"name\":\"getRegisteredValidatorsV2\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bool[]\",\"name\":\"stopped\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"restaking\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportedValidatorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportedValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStoppedValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVectorClock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"instantSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"previewInstantSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEthersToSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokensToBurn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pushBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"clock\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxRewards\",\"type\":\"uint256\"}],\"name\":\"pushBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethersToRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxToBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"redeemFromValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"burned\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"uint8[]\",\"name\":\"podIds\",\"type\":\"uint8[]\"}],\"name\":\"registerRestakingValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"registerRestakingValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"registerValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"oldpubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"restaking\",\"type\":\"bool\"},{\"internalType\":\"uint8[]\",\"name\":\"podIds\",\"type\":\"uint8[]\"}],\"name\":\"replaceValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethDepositContract\",\"type\":\"address\"}],\"name\":\"setETHDepositContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"milli\",\"type\":\"uint256\"}],\"name\":\"setManagerFeeShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"withdrawalCredentials_\",\"type\":\"bytes32\"}],\"name\":\"setWithdrawCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorRegistry\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"stopped\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"restaking\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"eigenpod\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_stoppedPubKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"clock\",\"type\":\"bytes32\"}],\"name\":\"validatorStopped\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawManagerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xETHAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Staking.Contract.DEFAULTADMINROLE(&_Staking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Staking.Contract.DEFAULTADMINROLE(&_Staking.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_Staking *StakingCaller) DEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_Staking *StakingSession) DEPOSITSIZE() (*big.Int, error) {
	return _Staking.Contract.DEPOSITSIZE(&_Staking.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_Staking *StakingCallerSession) DEPOSITSIZE() (*big.Int, error) {
	return _Staking.Contract.DEPOSITSIZE(&_Staking.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Staking *StakingSession) MANAGERROLE() ([32]byte, error) {
	return _Staking.Contract.MANAGERROLE(&_Staking.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) MANAGERROLE() ([32]byte, error) {
	return _Staking.Contract.MANAGERROLE(&_Staking.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) ORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_Staking *StakingSession) ORACLEROLE() ([32]byte, error) {
	return _Staking.Contract.ORACLEROLE(&_Staking.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) ORACLEROLE() ([32]byte, error) {
	return _Staking.Contract.ORACLEROLE(&_Staking.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Staking *StakingSession) PAUSERROLE() ([32]byte, error) {
	return _Staking.Contract.PAUSERROLE(&_Staking.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Staking.Contract.PAUSERROLE(&_Staking.CallOpts)
}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) REGISTRYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "REGISTRY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_Staking *StakingSession) REGISTRYROLE() ([32]byte, error) {
	return _Staking.Contract.REGISTRYROLE(&_Staking.CallOpts)
}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) REGISTRYROLE() ([32]byte, error) {
	return _Staking.Contract.REGISTRYROLE(&_Staking.CallOpts)
}

// SAFEPUSHREWARDS is a free data retrieval call binding the contract method 0x506a7bec.
//
// Solidity: function SAFE_PUSH_REWARDS() view returns(uint256)
func (_Staking *StakingCaller) SAFEPUSHREWARDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "SAFE_PUSH_REWARDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SAFEPUSHREWARDS is a free data retrieval call binding the contract method 0x506a7bec.
//
// Solidity: function SAFE_PUSH_REWARDS() view returns(uint256)
func (_Staking *StakingSession) SAFEPUSHREWARDS() (*big.Int, error) {
	return _Staking.Contract.SAFEPUSHREWARDS(&_Staking.CallOpts)
}

// SAFEPUSHREWARDS is a free data retrieval call binding the contract method 0x506a7bec.
//
// Solidity: function SAFE_PUSH_REWARDS() view returns(uint256)
func (_Staking *StakingCallerSession) SAFEPUSHREWARDS() (*big.Int, error) {
	return _Staking.Contract.SAFEPUSHREWARDS(&_Staking.CallOpts)
}

// CheckDebt is a free data retrieval call binding the contract method 0xc8c3df4a.
//
// Solidity: function checkDebt(uint256 index) view returns(address account, uint256 amount)
func (_Staking *StakingCaller) CheckDebt(opts *bind.CallOpts, index *big.Int) (struct {
	Account common.Address
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "checkDebt", index)

	outstruct := new(struct {
		Account common.Address
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CheckDebt is a free data retrieval call binding the contract method 0xc8c3df4a.
//
// Solidity: function checkDebt(uint256 index) view returns(address account, uint256 amount)
func (_Staking *StakingSession) CheckDebt(index *big.Int) (struct {
	Account common.Address
	Amount  *big.Int
}, error) {
	return _Staking.Contract.CheckDebt(&_Staking.CallOpts, index)
}

// CheckDebt is a free data retrieval call binding the contract method 0xc8c3df4a.
//
// Solidity: function checkDebt(uint256 index) view returns(address account, uint256 amount)
func (_Staking *StakingCallerSession) CheckDebt(index *big.Int) (struct {
	Account common.Address
	Amount  *big.Int
}, error) {
	return _Staking.Contract.CheckDebt(&_Staking.CallOpts, index)
}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_Staking *StakingCaller) CurrentReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "currentReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_Staking *StakingSession) CurrentReserve() (*big.Int, error) {
	return _Staking.Contract.CurrentReserve(&_Staking.CallOpts)
}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_Staking *StakingCallerSession) CurrentReserve() (*big.Int, error) {
	return _Staking.Contract.CurrentReserve(&_Staking.CallOpts)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_Staking *StakingCaller) DebtOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "debtOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_Staking *StakingSession) DebtOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.DebtOf(&_Staking.CallOpts, account)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_Staking *StakingCallerSession) DebtOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.DebtOf(&_Staking.CallOpts, account)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_Staking *StakingCaller) EthDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "ethDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_Staking *StakingSession) EthDepositContract() (common.Address, error) {
	return _Staking.Contract.EthDepositContract(&_Staking.CallOpts)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_Staking *StakingCallerSession) EthDepositContract() (common.Address, error) {
	return _Staking.Contract.EthDepositContract(&_Staking.CallOpts)
}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256)
func (_Staking *StakingCaller) ExchangeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "exchangeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256)
func (_Staking *StakingSession) ExchangeRatio() (*big.Int, error) {
	return _Staking.Contract.ExchangeRatio(&_Staking.CallOpts)
}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256)
func (_Staking *StakingCallerSession) ExchangeRatio() (*big.Int, error) {
	return _Staking.Contract.ExchangeRatio(&_Staking.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(int256)
func (_Staking *StakingCaller) GetAccountedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getAccountedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(int256)
func (_Staking *StakingSession) GetAccountedBalance() (*big.Int, error) {
	return _Staking.Contract.GetAccountedBalance(&_Staking.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(int256)
func (_Staking *StakingCallerSession) GetAccountedBalance() (*big.Int, error) {
	return _Staking.Contract.GetAccountedBalance(&_Staking.CallOpts)
}

// GetAccountedManagerRevenue is a free data retrieval call binding the contract method 0x08b84c0c.
//
// Solidity: function getAccountedManagerRevenue() view returns(uint256)
func (_Staking *StakingCaller) GetAccountedManagerRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getAccountedManagerRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedManagerRevenue is a free data retrieval call binding the contract method 0x08b84c0c.
//
// Solidity: function getAccountedManagerRevenue() view returns(uint256)
func (_Staking *StakingSession) GetAccountedManagerRevenue() (*big.Int, error) {
	return _Staking.Contract.GetAccountedManagerRevenue(&_Staking.CallOpts)
}

// GetAccountedManagerRevenue is a free data retrieval call binding the contract method 0x08b84c0c.
//
// Solidity: function getAccountedManagerRevenue() view returns(uint256)
func (_Staking *StakingCallerSession) GetAccountedManagerRevenue() (*big.Int, error) {
	return _Staking.Contract.GetAccountedManagerRevenue(&_Staking.CallOpts)
}

// GetAccountedUserRevenue is a free data retrieval call binding the contract method 0x61c993c5.
//
// Solidity: function getAccountedUserRevenue() view returns(uint256)
func (_Staking *StakingCaller) GetAccountedUserRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getAccountedUserRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedUserRevenue is a free data retrieval call binding the contract method 0x61c993c5.
//
// Solidity: function getAccountedUserRevenue() view returns(uint256)
func (_Staking *StakingSession) GetAccountedUserRevenue() (*big.Int, error) {
	return _Staking.Contract.GetAccountedUserRevenue(&_Staking.CallOpts)
}

// GetAccountedUserRevenue is a free data retrieval call binding the contract method 0x61c993c5.
//
// Solidity: function getAccountedUserRevenue() view returns(uint256)
func (_Staking *StakingCallerSession) GetAccountedUserRevenue() (*big.Int, error) {
	return _Staking.Contract.GetAccountedUserRevenue(&_Staking.CallOpts)
}

// GetCurrentDebts is a free data retrieval call binding the contract method 0x8b0bfd35.
//
// Solidity: function getCurrentDebts() view returns(uint256)
func (_Staking *StakingCaller) GetCurrentDebts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getCurrentDebts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentDebts is a free data retrieval call binding the contract method 0x8b0bfd35.
//
// Solidity: function getCurrentDebts() view returns(uint256)
func (_Staking *StakingSession) GetCurrentDebts() (*big.Int, error) {
	return _Staking.Contract.GetCurrentDebts(&_Staking.CallOpts)
}

// GetCurrentDebts is a free data retrieval call binding the contract method 0x8b0bfd35.
//
// Solidity: function getCurrentDebts() view returns(uint256)
func (_Staking *StakingCallerSession) GetCurrentDebts() (*big.Int, error) {
	return _Staking.Contract.GetCurrentDebts(&_Staking.CallOpts)
}

// GetDebtQueue is a free data retrieval call binding the contract method 0xdc3fc3b2.
//
// Solidity: function getDebtQueue() view returns(uint256 first, uint256 last)
func (_Staking *StakingCaller) GetDebtQueue(opts *bind.CallOpts) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDebtQueue")

	outstruct := new(struct {
		First *big.Int
		Last  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.First = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Last = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDebtQueue is a free data retrieval call binding the contract method 0xdc3fc3b2.
//
// Solidity: function getDebtQueue() view returns(uint256 first, uint256 last)
func (_Staking *StakingSession) GetDebtQueue() (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _Staking.Contract.GetDebtQueue(&_Staking.CallOpts)
}

// GetDebtQueue is a free data retrieval call binding the contract method 0xdc3fc3b2.
//
// Solidity: function getDebtQueue() view returns(uint256 first, uint256 last)
func (_Staking *StakingCallerSession) GetDebtQueue() (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _Staking.Contract.GetDebtQueue(&_Staking.CallOpts)
}

// GetNextValidatorId is a free data retrieval call binding the contract method 0xda863b3b.
//
// Solidity: function getNextValidatorId() view returns(uint256)
func (_Staking *StakingCaller) GetNextValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getNextValidatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextValidatorId is a free data retrieval call binding the contract method 0xda863b3b.
//
// Solidity: function getNextValidatorId() view returns(uint256)
func (_Staking *StakingSession) GetNextValidatorId() (*big.Int, error) {
	return _Staking.Contract.GetNextValidatorId(&_Staking.CallOpts)
}

// GetNextValidatorId is a free data retrieval call binding the contract method 0xda863b3b.
//
// Solidity: function getNextValidatorId() view returns(uint256)
func (_Staking *StakingCallerSession) GetNextValidatorId() (*big.Int, error) {
	return _Staking.Contract.GetNextValidatorId(&_Staking.CallOpts)
}

// GetPendingEthers is a free data retrieval call binding the contract method 0x64363f2b.
//
// Solidity: function getPendingEthers() view returns(uint256)
func (_Staking *StakingCaller) GetPendingEthers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getPendingEthers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingEthers is a free data retrieval call binding the contract method 0x64363f2b.
//
// Solidity: function getPendingEthers() view returns(uint256)
func (_Staking *StakingSession) GetPendingEthers() (*big.Int, error) {
	return _Staking.Contract.GetPendingEthers(&_Staking.CallOpts)
}

// GetPendingEthers is a free data retrieval call binding the contract method 0x64363f2b.
//
// Solidity: function getPendingEthers() view returns(uint256)
func (_Staking *StakingCallerSession) GetPendingEthers() (*big.Int, error) {
	return _Staking.Contract.GetPendingEthers(&_Staking.CallOpts)
}

// GetRecentReceived is a free data retrieval call binding the contract method 0xe08f2d89.
//
// Solidity: function getRecentReceived() view returns(uint256)
func (_Staking *StakingCaller) GetRecentReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRecentReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRecentReceived is a free data retrieval call binding the contract method 0xe08f2d89.
//
// Solidity: function getRecentReceived() view returns(uint256)
func (_Staking *StakingSession) GetRecentReceived() (*big.Int, error) {
	return _Staking.Contract.GetRecentReceived(&_Staking.CallOpts)
}

// GetRecentReceived is a free data retrieval call binding the contract method 0xe08f2d89.
//
// Solidity: function getRecentReceived() view returns(uint256)
func (_Staking *StakingCallerSession) GetRecentReceived() (*big.Int, error) {
	return _Staking.Contract.GetRecentReceived(&_Staking.CallOpts)
}

// GetRecentStopped is a free data retrieval call binding the contract method 0x2ae45fa1.
//
// Solidity: function getRecentStopped() view returns(uint256)
func (_Staking *StakingCaller) GetRecentStopped(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRecentStopped")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRecentStopped is a free data retrieval call binding the contract method 0x2ae45fa1.
//
// Solidity: function getRecentStopped() view returns(uint256)
func (_Staking *StakingSession) GetRecentStopped() (*big.Int, error) {
	return _Staking.Contract.GetRecentStopped(&_Staking.CallOpts)
}

// GetRecentStopped is a free data retrieval call binding the contract method 0x2ae45fa1.
//
// Solidity: function getRecentStopped() view returns(uint256)
func (_Staking *StakingCallerSession) GetRecentStopped() (*big.Int, error) {
	return _Staking.Contract.GetRecentStopped(&_Staking.CallOpts)
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xf22abf37.
//
// Solidity: function getRegisteredValidators(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped)
func (_Staking *StakingCaller) GetRegisteredValidators(opts *bind.CallOpts, idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRegisteredValidators", idx_from, idx_to)

	outstruct := new(struct {
		Pubkeys    [][]byte
		Signatures [][]byte
		Stopped    []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkeys = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.Signatures = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)
	outstruct.Stopped = *abi.ConvertType(out[2], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xf22abf37.
//
// Solidity: function getRegisteredValidators(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped)
func (_Staking *StakingSession) GetRegisteredValidators(idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
}, error) {
	return _Staking.Contract.GetRegisteredValidators(&_Staking.CallOpts, idx_from, idx_to)
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xf22abf37.
//
// Solidity: function getRegisteredValidators(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped)
func (_Staking *StakingCallerSession) GetRegisteredValidators(idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
}, error) {
	return _Staking.Contract.GetRegisteredValidators(&_Staking.CallOpts, idx_from, idx_to)
}

// GetRegisteredValidatorsCount is a free data retrieval call binding the contract method 0x30b12c8d.
//
// Solidity: function getRegisteredValidatorsCount() view returns(uint256)
func (_Staking *StakingCaller) GetRegisteredValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRegisteredValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRegisteredValidatorsCount is a free data retrieval call binding the contract method 0x30b12c8d.
//
// Solidity: function getRegisteredValidatorsCount() view returns(uint256)
func (_Staking *StakingSession) GetRegisteredValidatorsCount() (*big.Int, error) {
	return _Staking.Contract.GetRegisteredValidatorsCount(&_Staking.CallOpts)
}

// GetRegisteredValidatorsCount is a free data retrieval call binding the contract method 0x30b12c8d.
//
// Solidity: function getRegisteredValidatorsCount() view returns(uint256)
func (_Staking *StakingCallerSession) GetRegisteredValidatorsCount() (*big.Int, error) {
	return _Staking.Contract.GetRegisteredValidatorsCount(&_Staking.CallOpts)
}

// GetRegisteredValidatorsV2 is a free data retrieval call binding the contract method 0x3bf39dca.
//
// Solidity: function getRegisteredValidatorsV2(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped, bool[] restaking)
func (_Staking *StakingCaller) GetRegisteredValidatorsV2(opts *bind.CallOpts, idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
	Restaking  []bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRegisteredValidatorsV2", idx_from, idx_to)

	outstruct := new(struct {
		Pubkeys    [][]byte
		Signatures [][]byte
		Stopped    []bool
		Restaking  []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkeys = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.Signatures = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)
	outstruct.Stopped = *abi.ConvertType(out[2], new([]bool)).(*[]bool)
	outstruct.Restaking = *abi.ConvertType(out[3], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetRegisteredValidatorsV2 is a free data retrieval call binding the contract method 0x3bf39dca.
//
// Solidity: function getRegisteredValidatorsV2(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped, bool[] restaking)
func (_Staking *StakingSession) GetRegisteredValidatorsV2(idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
	Restaking  []bool
}, error) {
	return _Staking.Contract.GetRegisteredValidatorsV2(&_Staking.CallOpts, idx_from, idx_to)
}

// GetRegisteredValidatorsV2 is a free data retrieval call binding the contract method 0x3bf39dca.
//
// Solidity: function getRegisteredValidatorsV2(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped, bool[] restaking)
func (_Staking *StakingCallerSession) GetRegisteredValidatorsV2(idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
	Restaking  []bool
}, error) {
	return _Staking.Contract.GetRegisteredValidatorsV2(&_Staking.CallOpts, idx_from, idx_to)
}

// GetReportedValidatorBalance is a free data retrieval call binding the contract method 0xcdb54a1b.
//
// Solidity: function getReportedValidatorBalance() view returns(uint256)
func (_Staking *StakingCaller) GetReportedValidatorBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getReportedValidatorBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportedValidatorBalance is a free data retrieval call binding the contract method 0xcdb54a1b.
//
// Solidity: function getReportedValidatorBalance() view returns(uint256)
func (_Staking *StakingSession) GetReportedValidatorBalance() (*big.Int, error) {
	return _Staking.Contract.GetReportedValidatorBalance(&_Staking.CallOpts)
}

// GetReportedValidatorBalance is a free data retrieval call binding the contract method 0xcdb54a1b.
//
// Solidity: function getReportedValidatorBalance() view returns(uint256)
func (_Staking *StakingCallerSession) GetReportedValidatorBalance() (*big.Int, error) {
	return _Staking.Contract.GetReportedValidatorBalance(&_Staking.CallOpts)
}

// GetReportedValidators is a free data retrieval call binding the contract method 0x6a42602c.
//
// Solidity: function getReportedValidators() view returns(uint256)
func (_Staking *StakingCaller) GetReportedValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getReportedValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportedValidators is a free data retrieval call binding the contract method 0x6a42602c.
//
// Solidity: function getReportedValidators() view returns(uint256)
func (_Staking *StakingSession) GetReportedValidators() (*big.Int, error) {
	return _Staking.Contract.GetReportedValidators(&_Staking.CallOpts)
}

// GetReportedValidators is a free data retrieval call binding the contract method 0x6a42602c.
//
// Solidity: function getReportedValidators() view returns(uint256)
func (_Staking *StakingCallerSession) GetReportedValidators() (*big.Int, error) {
	return _Staking.Contract.GetReportedValidators(&_Staking.CallOpts)
}

// GetRewardDebts is a free data retrieval call binding the contract method 0xa065913f.
//
// Solidity: function getRewardDebts() view returns(uint256)
func (_Staking *StakingCaller) GetRewardDebts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRewardDebts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardDebts is a free data retrieval call binding the contract method 0xa065913f.
//
// Solidity: function getRewardDebts() view returns(uint256)
func (_Staking *StakingSession) GetRewardDebts() (*big.Int, error) {
	return _Staking.Contract.GetRewardDebts(&_Staking.CallOpts)
}

// GetRewardDebts is a free data retrieval call binding the contract method 0xa065913f.
//
// Solidity: function getRewardDebts() view returns(uint256)
func (_Staking *StakingCallerSession) GetRewardDebts() (*big.Int, error) {
	return _Staking.Contract.GetRewardDebts(&_Staking.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetRoleAdmin(&_Staking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetRoleAdmin(&_Staking.CallOpts, role)
}

// GetStoppedValidatorsCount is a free data retrieval call binding the contract method 0x49557df9.
//
// Solidity: function getStoppedValidatorsCount() view returns(uint256)
func (_Staking *StakingCaller) GetStoppedValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStoppedValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStoppedValidatorsCount is a free data retrieval call binding the contract method 0x49557df9.
//
// Solidity: function getStoppedValidatorsCount() view returns(uint256)
func (_Staking *StakingSession) GetStoppedValidatorsCount() (*big.Int, error) {
	return _Staking.Contract.GetStoppedValidatorsCount(&_Staking.CallOpts)
}

// GetStoppedValidatorsCount is a free data retrieval call binding the contract method 0x49557df9.
//
// Solidity: function getStoppedValidatorsCount() view returns(uint256)
func (_Staking *StakingCallerSession) GetStoppedValidatorsCount() (*big.Int, error) {
	return _Staking.Contract.GetStoppedValidatorsCount(&_Staking.CallOpts)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x0917e776.
//
// Solidity: function getTotalStaked() view returns(uint256)
func (_Staking *StakingCaller) GetTotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getTotalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStaked is a free data retrieval call binding the contract method 0x0917e776.
//
// Solidity: function getTotalStaked() view returns(uint256)
func (_Staking *StakingSession) GetTotalStaked() (*big.Int, error) {
	return _Staking.Contract.GetTotalStaked(&_Staking.CallOpts)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x0917e776.
//
// Solidity: function getTotalStaked() view returns(uint256)
func (_Staking *StakingCallerSession) GetTotalStaked() (*big.Int, error) {
	return _Staking.Contract.GetTotalStaked(&_Staking.CallOpts)
}

// GetVectorClock is a free data retrieval call binding the contract method 0xecacf56d.
//
// Solidity: function getVectorClock() view returns(bytes32)
func (_Staking *StakingCaller) GetVectorClock(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getVectorClock")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVectorClock is a free data retrieval call binding the contract method 0xecacf56d.
//
// Solidity: function getVectorClock() view returns(bytes32)
func (_Staking *StakingSession) GetVectorClock() ([32]byte, error) {
	return _Staking.Contract.GetVectorClock(&_Staking.CallOpts)
}

// GetVectorClock is a free data retrieval call binding the contract method 0xecacf56d.
//
// Solidity: function getVectorClock() view returns(bytes32)
func (_Staking *StakingCallerSession) GetVectorClock() ([32]byte, error) {
	return _Staking.Contract.GetVectorClock(&_Staking.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Staking.Contract.HasRole(&_Staking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Staking.Contract.HasRole(&_Staking.CallOpts, role, account)
}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_Staking *StakingCaller) ManagerFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "managerFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_Staking *StakingSession) ManagerFeeShare() (*big.Int, error) {
	return _Staking.Contract.ManagerFeeShare(&_Staking.CallOpts)
}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_Staking *StakingCallerSession) ManagerFeeShare() (*big.Int, error) {
	return _Staking.Contract.ManagerFeeShare(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCallerSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// PreviewInstantSwap is a free data retrieval call binding the contract method 0xfa695018.
//
// Solidity: function previewInstantSwap(uint256 tokenAmount) view returns(uint256 maxEthersToSwap, uint256 maxTokensToBurn)
func (_Staking *StakingCaller) PreviewInstantSwap(opts *bind.CallOpts, tokenAmount *big.Int) (struct {
	MaxEthersToSwap *big.Int
	MaxTokensToBurn *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "previewInstantSwap", tokenAmount)

	outstruct := new(struct {
		MaxEthersToSwap *big.Int
		MaxTokensToBurn *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxEthersToSwap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxTokensToBurn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PreviewInstantSwap is a free data retrieval call binding the contract method 0xfa695018.
//
// Solidity: function previewInstantSwap(uint256 tokenAmount) view returns(uint256 maxEthersToSwap, uint256 maxTokensToBurn)
func (_Staking *StakingSession) PreviewInstantSwap(tokenAmount *big.Int) (struct {
	MaxEthersToSwap *big.Int
	MaxTokensToBurn *big.Int
}, error) {
	return _Staking.Contract.PreviewInstantSwap(&_Staking.CallOpts, tokenAmount)
}

// PreviewInstantSwap is a free data retrieval call binding the contract method 0xfa695018.
//
// Solidity: function previewInstantSwap(uint256 tokenAmount) view returns(uint256 maxEthersToSwap, uint256 maxTokensToBurn)
func (_Staking *StakingCallerSession) PreviewInstantSwap(tokenAmount *big.Int) (struct {
	MaxEthersToSwap *big.Int
	MaxTokensToBurn *big.Int
}, error) {
	return _Staking.Contract.PreviewInstantSwap(&_Staking.CallOpts, tokenAmount)
}

// RedeemContract is a free data retrieval call binding the contract method 0x7a4473e1.
//
// Solidity: function redeemContract() view returns(address)
func (_Staking *StakingCaller) RedeemContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "redeemContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedeemContract is a free data retrieval call binding the contract method 0x7a4473e1.
//
// Solidity: function redeemContract() view returns(address)
func (_Staking *StakingSession) RedeemContract() (common.Address, error) {
	return _Staking.Contract.RedeemContract(&_Staking.CallOpts)
}

// RedeemContract is a free data retrieval call binding the contract method 0x7a4473e1.
//
// Solidity: function redeemContract() view returns(address)
func (_Staking *StakingCallerSession) RedeemContract() (common.Address, error) {
	return _Staking.Contract.RedeemContract(&_Staking.CallOpts)
}

// RestakingContract is a free data retrieval call binding the contract method 0xf1f3b3e7.
//
// Solidity: function restakingContract() view returns(address)
func (_Staking *StakingCaller) RestakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "restakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RestakingContract is a free data retrieval call binding the contract method 0xf1f3b3e7.
//
// Solidity: function restakingContract() view returns(address)
func (_Staking *StakingSession) RestakingContract() (common.Address, error) {
	return _Staking.Contract.RestakingContract(&_Staking.CallOpts)
}

// RestakingContract is a free data retrieval call binding the contract method 0xf1f3b3e7.
//
// Solidity: function restakingContract() view returns(address)
func (_Staking *StakingCallerSession) RestakingContract() (common.Address, error) {
	return _Staking.Contract.RestakingContract(&_Staking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Staking.Contract.SupportsInterface(&_Staking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Staking.Contract.SupportsInterface(&_Staking.CallOpts, interfaceId)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(bytes pubkey, bytes signature, bool stopped, bool restaking, uint8 eigenpod)
func (_Staking *StakingCaller) ValidatorRegistry(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Pubkey    []byte
	Signature []byte
	Stopped   bool
	Restaking bool
	Eigenpod  uint8
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validatorRegistry", arg0)

	outstruct := new(struct {
		Pubkey    []byte
		Signature []byte
		Stopped   bool
		Restaking bool
		Eigenpod  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkey = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Signature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Stopped = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Restaking = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Eigenpod = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(bytes pubkey, bytes signature, bool stopped, bool restaking, uint8 eigenpod)
func (_Staking *StakingSession) ValidatorRegistry(arg0 *big.Int) (struct {
	Pubkey    []byte
	Signature []byte
	Stopped   bool
	Restaking bool
	Eigenpod  uint8
}, error) {
	return _Staking.Contract.ValidatorRegistry(&_Staking.CallOpts, arg0)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(bytes pubkey, bytes signature, bool stopped, bool restaking, uint8 eigenpod)
func (_Staking *StakingCallerSession) ValidatorRegistry(arg0 *big.Int) (struct {
	Pubkey    []byte
	Signature []byte
	Stopped   bool
	Restaking bool
	Eigenpod  uint8
}, error) {
	return _Staking.Contract.ValidatorRegistry(&_Staking.CallOpts, arg0)
}

// WithdrawalCredentials is a free data retrieval call binding the contract method 0x4cd79e0a.
//
// Solidity: function withdrawalCredentials() view returns(bytes32)
func (_Staking *StakingCaller) WithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawalCredentials is a free data retrieval call binding the contract method 0x4cd79e0a.
//
// Solidity: function withdrawalCredentials() view returns(bytes32)
func (_Staking *StakingSession) WithdrawalCredentials() ([32]byte, error) {
	return _Staking.Contract.WithdrawalCredentials(&_Staking.CallOpts)
}

// WithdrawalCredentials is a free data retrieval call binding the contract method 0x4cd79e0a.
//
// Solidity: function withdrawalCredentials() view returns(bytes32)
func (_Staking *StakingCallerSession) WithdrawalCredentials() ([32]byte, error) {
	return _Staking.Contract.WithdrawalCredentials(&_Staking.CallOpts)
}

// XETHAddress is a free data retrieval call binding the contract method 0xb181033a.
//
// Solidity: function xETHAddress() view returns(address)
func (_Staking *StakingCaller) XETHAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "xETHAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XETHAddress is a free data retrieval call binding the contract method 0xb181033a.
//
// Solidity: function xETHAddress() view returns(address)
func (_Staking *StakingSession) XETHAddress() (common.Address, error) {
	return _Staking.Contract.XETHAddress(&_Staking.CallOpts)
}

// XETHAddress is a free data retrieval call binding the contract method 0xb181033a.
//
// Solidity: function xETHAddress() view returns(address)
func (_Staking *StakingCallerSession) XETHAddress() (common.Address, error) {
	return _Staking.Contract.XETHAddress(&_Staking.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GrantRole(&_Staking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GrantRole(&_Staking.TransactOpts, role, account)
}

// InstantSwap is a paid mutator transaction binding the contract method 0xfadb574d.
//
// Solidity: function instantSwap(uint256 tokenAmount) returns()
func (_Staking *StakingTransactor) InstantSwap(opts *bind.TransactOpts, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "instantSwap", tokenAmount)
}

// InstantSwap is a paid mutator transaction binding the contract method 0xfadb574d.
//
// Solidity: function instantSwap(uint256 tokenAmount) returns()
func (_Staking *StakingSession) InstantSwap(tokenAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.InstantSwap(&_Staking.TransactOpts, tokenAmount)
}

// InstantSwap is a paid mutator transaction binding the contract method 0xfadb574d.
//
// Solidity: function instantSwap(uint256 tokenAmount) returns()
func (_Staking *StakingTransactorSession) InstantSwap(tokenAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.InstantSwap(&_Staking.TransactOpts, tokenAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 minToMint, uint256 deadline) payable returns(uint256 minted)
func (_Staking *StakingTransactor) Mint(opts *bind.TransactOpts, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "mint", minToMint, deadline)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 minToMint, uint256 deadline) payable returns(uint256 minted)
func (_Staking *StakingSession) Mint(minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Mint(&_Staking.TransactOpts, minToMint, deadline)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 minToMint, uint256 deadline) payable returns(uint256 minted)
func (_Staking *StakingTransactorSession) Mint(minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Mint(&_Staking.TransactOpts, minToMint, deadline)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactorSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// PushBeacon is a paid mutator transaction binding the contract method 0x849138af.
//
// Solidity: function pushBeacon() returns()
func (_Staking *StakingTransactor) PushBeacon(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "pushBeacon")
}

// PushBeacon is a paid mutator transaction binding the contract method 0x849138af.
//
// Solidity: function pushBeacon() returns()
func (_Staking *StakingSession) PushBeacon() (*types.Transaction, error) {
	return _Staking.Contract.PushBeacon(&_Staking.TransactOpts)
}

// PushBeacon is a paid mutator transaction binding the contract method 0x849138af.
//
// Solidity: function pushBeacon() returns()
func (_Staking *StakingTransactorSession) PushBeacon() (*types.Transaction, error) {
	return _Staking.Contract.PushBeacon(&_Staking.TransactOpts)
}

// PushBeacon0 is a paid mutator transaction binding the contract method 0xa13cb538.
//
// Solidity: function pushBeacon(bytes32 clock, uint256 maxRewards) returns()
func (_Staking *StakingTransactor) PushBeacon0(opts *bind.TransactOpts, clock [32]byte, maxRewards *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "pushBeacon0", clock, maxRewards)
}

// PushBeacon0 is a paid mutator transaction binding the contract method 0xa13cb538.
//
// Solidity: function pushBeacon(bytes32 clock, uint256 maxRewards) returns()
func (_Staking *StakingSession) PushBeacon0(clock [32]byte, maxRewards *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.PushBeacon0(&_Staking.TransactOpts, clock, maxRewards)
}

// PushBeacon0 is a paid mutator transaction binding the contract method 0xa13cb538.
//
// Solidity: function pushBeacon(bytes32 clock, uint256 maxRewards) returns()
func (_Staking *StakingTransactorSession) PushBeacon0(clock [32]byte, maxRewards *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.PushBeacon0(&_Staking.TransactOpts, clock, maxRewards)
}

// RedeemFromValidators is a paid mutator transaction binding the contract method 0xf5404d60.
//
// Solidity: function redeemFromValidators(uint256 ethersToRedeem, uint256 maxToBurn, uint256 deadline) returns(uint256 burned)
func (_Staking *StakingTransactor) RedeemFromValidators(opts *bind.TransactOpts, ethersToRedeem *big.Int, maxToBurn *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "redeemFromValidators", ethersToRedeem, maxToBurn, deadline)
}

// RedeemFromValidators is a paid mutator transaction binding the contract method 0xf5404d60.
//
// Solidity: function redeemFromValidators(uint256 ethersToRedeem, uint256 maxToBurn, uint256 deadline) returns(uint256 burned)
func (_Staking *StakingSession) RedeemFromValidators(ethersToRedeem *big.Int, maxToBurn *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.RedeemFromValidators(&_Staking.TransactOpts, ethersToRedeem, maxToBurn, deadline)
}

// RedeemFromValidators is a paid mutator transaction binding the contract method 0xf5404d60.
//
// Solidity: function redeemFromValidators(uint256 ethersToRedeem, uint256 maxToBurn, uint256 deadline) returns(uint256 burned)
func (_Staking *StakingTransactorSession) RedeemFromValidators(ethersToRedeem *big.Int, maxToBurn *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.RedeemFromValidators(&_Staking.TransactOpts, ethersToRedeem, maxToBurn, deadline)
}

// RegisterRestakingValidators is a paid mutator transaction binding the contract method 0x0a1521fc.
//
// Solidity: function registerRestakingValidators(bytes[] pubkeys, bytes[] signatures, uint8[] podIds) returns()
func (_Staking *StakingTransactor) RegisterRestakingValidators(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte, podIds []uint8) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "registerRestakingValidators", pubkeys, signatures, podIds)
}

// RegisterRestakingValidators is a paid mutator transaction binding the contract method 0x0a1521fc.
//
// Solidity: function registerRestakingValidators(bytes[] pubkeys, bytes[] signatures, uint8[] podIds) returns()
func (_Staking *StakingSession) RegisterRestakingValidators(pubkeys [][]byte, signatures [][]byte, podIds []uint8) (*types.Transaction, error) {
	return _Staking.Contract.RegisterRestakingValidators(&_Staking.TransactOpts, pubkeys, signatures, podIds)
}

// RegisterRestakingValidators is a paid mutator transaction binding the contract method 0x0a1521fc.
//
// Solidity: function registerRestakingValidators(bytes[] pubkeys, bytes[] signatures, uint8[] podIds) returns()
func (_Staking *StakingTransactorSession) RegisterRestakingValidators(pubkeys [][]byte, signatures [][]byte, podIds []uint8) (*types.Transaction, error) {
	return _Staking.Contract.RegisterRestakingValidators(&_Staking.TransactOpts, pubkeys, signatures, podIds)
}

// RegisterRestakingValidators0 is a paid mutator transaction binding the contract method 0x4e8ab1a1.
//
// Solidity: function registerRestakingValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_Staking *StakingTransactor) RegisterRestakingValidators0(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "registerRestakingValidators0", pubkeys, signatures)
}

// RegisterRestakingValidators0 is a paid mutator transaction binding the contract method 0x4e8ab1a1.
//
// Solidity: function registerRestakingValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_Staking *StakingSession) RegisterRestakingValidators0(pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.RegisterRestakingValidators0(&_Staking.TransactOpts, pubkeys, signatures)
}

// RegisterRestakingValidators0 is a paid mutator transaction binding the contract method 0x4e8ab1a1.
//
// Solidity: function registerRestakingValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_Staking *StakingTransactorSession) RegisterRestakingValidators0(pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.RegisterRestakingValidators0(&_Staking.TransactOpts, pubkeys, signatures)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xe7efb747.
//
// Solidity: function registerValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_Staking *StakingTransactor) RegisterValidators(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "registerValidators", pubkeys, signatures)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xe7efb747.
//
// Solidity: function registerValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_Staking *StakingSession) RegisterValidators(pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.RegisterValidators(&_Staking.TransactOpts, pubkeys, signatures)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xe7efb747.
//
// Solidity: function registerValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_Staking *StakingTransactorSession) RegisterValidators(pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.RegisterValidators(&_Staking.TransactOpts, pubkeys, signatures)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RenounceRole(&_Staking.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RenounceRole(&_Staking.TransactOpts, role, account)
}

// ReplaceValidators is a paid mutator transaction binding the contract method 0x4f13b2e3.
//
// Solidity: function replaceValidators(bytes[] oldpubkeys, bytes[] pubkeys, bytes[] signatures, bool restaking, uint8[] podIds) returns()
func (_Staking *StakingTransactor) ReplaceValidators(opts *bind.TransactOpts, oldpubkeys [][]byte, pubkeys [][]byte, signatures [][]byte, restaking bool, podIds []uint8) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "replaceValidators", oldpubkeys, pubkeys, signatures, restaking, podIds)
}

// ReplaceValidators is a paid mutator transaction binding the contract method 0x4f13b2e3.
//
// Solidity: function replaceValidators(bytes[] oldpubkeys, bytes[] pubkeys, bytes[] signatures, bool restaking, uint8[] podIds) returns()
func (_Staking *StakingSession) ReplaceValidators(oldpubkeys [][]byte, pubkeys [][]byte, signatures [][]byte, restaking bool, podIds []uint8) (*types.Transaction, error) {
	return _Staking.Contract.ReplaceValidators(&_Staking.TransactOpts, oldpubkeys, pubkeys, signatures, restaking, podIds)
}

// ReplaceValidators is a paid mutator transaction binding the contract method 0x4f13b2e3.
//
// Solidity: function replaceValidators(bytes[] oldpubkeys, bytes[] pubkeys, bytes[] signatures, bool restaking, uint8[] podIds) returns()
func (_Staking *StakingTransactorSession) ReplaceValidators(oldpubkeys [][]byte, pubkeys [][]byte, signatures [][]byte, restaking bool, podIds []uint8) (*types.Transaction, error) {
	return _Staking.Contract.ReplaceValidators(&_Staking.TransactOpts, oldpubkeys, pubkeys, signatures, restaking, podIds)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RevokeRole(&_Staking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RevokeRole(&_Staking.TransactOpts, role, account)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_Staking *StakingTransactor) SetETHDepositContract(opts *bind.TransactOpts, _ethDepositContract common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setETHDepositContract", _ethDepositContract)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_Staking *StakingSession) SetETHDepositContract(_ethDepositContract common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetETHDepositContract(&_Staking.TransactOpts, _ethDepositContract)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_Staking *StakingTransactorSession) SetETHDepositContract(_ethDepositContract common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetETHDepositContract(&_Staking.TransactOpts, _ethDepositContract)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_Staking *StakingTransactor) SetManagerFeeShare(opts *bind.TransactOpts, milli *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setManagerFeeShare", milli)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_Staking *StakingSession) SetManagerFeeShare(milli *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetManagerFeeShare(&_Staking.TransactOpts, milli)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_Staking *StakingTransactorSession) SetManagerFeeShare(milli *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetManagerFeeShare(&_Staking.TransactOpts, milli)
}

// SetWithdrawCredential is a paid mutator transaction binding the contract method 0xd7d25461.
//
// Solidity: function setWithdrawCredential(bytes32 withdrawalCredentials_) returns()
func (_Staking *StakingTransactor) SetWithdrawCredential(opts *bind.TransactOpts, withdrawalCredentials_ [32]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setWithdrawCredential", withdrawalCredentials_)
}

// SetWithdrawCredential is a paid mutator transaction binding the contract method 0xd7d25461.
//
// Solidity: function setWithdrawCredential(bytes32 withdrawalCredentials_) returns()
func (_Staking *StakingSession) SetWithdrawCredential(withdrawalCredentials_ [32]byte) (*types.Transaction, error) {
	return _Staking.Contract.SetWithdrawCredential(&_Staking.TransactOpts, withdrawalCredentials_)
}

// SetWithdrawCredential is a paid mutator transaction binding the contract method 0xd7d25461.
//
// Solidity: function setWithdrawCredential(bytes32 withdrawalCredentials_) returns()
func (_Staking *StakingTransactorSession) SetWithdrawCredential(withdrawalCredentials_ [32]byte) (*types.Transaction, error) {
	return _Staking.Contract.SetWithdrawCredential(&_Staking.TransactOpts, withdrawalCredentials_)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Staking *StakingSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Staking *StakingTransactorSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// SyncBalance is a paid mutator transaction binding the contract method 0xfd9c652b.
//
// Solidity: function syncBalance() returns()
func (_Staking *StakingTransactor) SyncBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "syncBalance")
}

// SyncBalance is a paid mutator transaction binding the contract method 0xfd9c652b.
//
// Solidity: function syncBalance() returns()
func (_Staking *StakingSession) SyncBalance() (*types.Transaction, error) {
	return _Staking.Contract.SyncBalance(&_Staking.TransactOpts)
}

// SyncBalance is a paid mutator transaction binding the contract method 0xfd9c652b.
//
// Solidity: function syncBalance() returns()
func (_Staking *StakingTransactorSession) SyncBalance() (*types.Transaction, error) {
	return _Staking.Contract.SyncBalance(&_Staking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// ValidatorStopped is a paid mutator transaction binding the contract method 0x99629f58.
//
// Solidity: function validatorStopped(bytes[] _stoppedPubKeys, bytes32 clock) returns()
func (_Staking *StakingTransactor) ValidatorStopped(opts *bind.TransactOpts, _stoppedPubKeys [][]byte, clock [32]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "validatorStopped", _stoppedPubKeys, clock)
}

// ValidatorStopped is a paid mutator transaction binding the contract method 0x99629f58.
//
// Solidity: function validatorStopped(bytes[] _stoppedPubKeys, bytes32 clock) returns()
func (_Staking *StakingSession) ValidatorStopped(_stoppedPubKeys [][]byte, clock [32]byte) (*types.Transaction, error) {
	return _Staking.Contract.ValidatorStopped(&_Staking.TransactOpts, _stoppedPubKeys, clock)
}

// ValidatorStopped is a paid mutator transaction binding the contract method 0x99629f58.
//
// Solidity: function validatorStopped(bytes[] _stoppedPubKeys, bytes32 clock) returns()
func (_Staking *StakingTransactorSession) ValidatorStopped(_stoppedPubKeys [][]byte, clock [32]byte) (*types.Transaction, error) {
	return _Staking.Contract.ValidatorStopped(&_Staking.TransactOpts, _stoppedPubKeys, clock)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0x8d40152c.
//
// Solidity: function withdrawManagerFee(address to) returns()
func (_Staking *StakingTransactor) WithdrawManagerFee(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdrawManagerFee", to)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0x8d40152c.
//
// Solidity: function withdrawManagerFee(address to) returns()
func (_Staking *StakingSession) WithdrawManagerFee(to common.Address) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawManagerFee(&_Staking.TransactOpts, to)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0x8d40152c.
//
// Solidity: function withdrawManagerFee(address to) returns()
func (_Staking *StakingTransactorSession) WithdrawManagerFee(to common.Address) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawManagerFee(&_Staking.TransactOpts, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// StakingBalanceSyncedIterator is returned from FilterBalanceSynced and is used to iterate over the raw logs and unpacked data for BalanceSynced events raised by the Staking contract.
type StakingBalanceSyncedIterator struct {
	Event *StakingBalanceSynced // Event containing the contract specifics and raw log

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
func (it *StakingBalanceSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBalanceSynced)
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
		it.Event = new(StakingBalanceSynced)
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
func (it *StakingBalanceSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBalanceSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBalanceSynced represents a BalanceSynced event raised by the Staking contract.
type StakingBalanceSynced struct {
	Diff *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBalanceSynced is a free log retrieval operation binding the contract event 0xe7948c33eb604391785037114655100edf93283c25b69884e9238ae197f07817.
//
// Solidity: event BalanceSynced(uint256 diff)
func (_Staking *StakingFilterer) FilterBalanceSynced(opts *bind.FilterOpts) (*StakingBalanceSyncedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "BalanceSynced")
	if err != nil {
		return nil, err
	}
	return &StakingBalanceSyncedIterator{contract: _Staking.contract, event: "BalanceSynced", logs: logs, sub: sub}, nil
}

// WatchBalanceSynced is a free log subscription operation binding the contract event 0xe7948c33eb604391785037114655100edf93283c25b69884e9238ae197f07817.
//
// Solidity: event BalanceSynced(uint256 diff)
func (_Staking *StakingFilterer) WatchBalanceSynced(opts *bind.WatchOpts, sink chan<- *StakingBalanceSynced) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "BalanceSynced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBalanceSynced)
				if err := _Staking.contract.UnpackLog(event, "BalanceSynced", log); err != nil {
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

// ParseBalanceSynced is a log parse operation binding the contract event 0xe7948c33eb604391785037114655100edf93283c25b69884e9238ae197f07817.
//
// Solidity: event BalanceSynced(uint256 diff)
func (_Staking *StakingFilterer) ParseBalanceSynced(log types.Log) (*StakingBalanceSynced, error) {
	event := new(StakingBalanceSynced)
	if err := _Staking.contract.UnpackLog(event, "BalanceSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDebtQueuedIterator is returned from FilterDebtQueued and is used to iterate over the raw logs and unpacked data for DebtQueued events raised by the Staking contract.
type StakingDebtQueuedIterator struct {
	Event *StakingDebtQueued // Event containing the contract specifics and raw log

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
func (it *StakingDebtQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDebtQueued)
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
		it.Event = new(StakingDebtQueued)
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
func (it *StakingDebtQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDebtQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDebtQueued represents a DebtQueued event raised by the Staking contract.
type StakingDebtQueued struct {
	Creditor    common.Address
	AmountEther *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDebtQueued is a free log retrieval operation binding the contract event 0x889f3b08db1ce169841b4f7e2aadfe4298088b1bce57b31fecae3260dd435829.
//
// Solidity: event DebtQueued(address creditor, uint256 amountEther)
func (_Staking *StakingFilterer) FilterDebtQueued(opts *bind.FilterOpts) (*StakingDebtQueuedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DebtQueued")
	if err != nil {
		return nil, err
	}
	return &StakingDebtQueuedIterator{contract: _Staking.contract, event: "DebtQueued", logs: logs, sub: sub}, nil
}

// WatchDebtQueued is a free log subscription operation binding the contract event 0x889f3b08db1ce169841b4f7e2aadfe4298088b1bce57b31fecae3260dd435829.
//
// Solidity: event DebtQueued(address creditor, uint256 amountEther)
func (_Staking *StakingFilterer) WatchDebtQueued(opts *bind.WatchOpts, sink chan<- *StakingDebtQueued) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DebtQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDebtQueued)
				if err := _Staking.contract.UnpackLog(event, "DebtQueued", log); err != nil {
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

// ParseDebtQueued is a log parse operation binding the contract event 0x889f3b08db1ce169841b4f7e2aadfe4298088b1bce57b31fecae3260dd435829.
//
// Solidity: event DebtQueued(address creditor, uint256 amountEther)
func (_Staking *StakingFilterer) ParseDebtQueued(log types.Log) (*StakingDebtQueued, error) {
	event := new(StakingDebtQueued)
	if err := _Staking.contract.UnpackLog(event, "DebtQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDepositContractSetIterator is returned from FilterDepositContractSet and is used to iterate over the raw logs and unpacked data for DepositContractSet events raised by the Staking contract.
type StakingDepositContractSetIterator struct {
	Event *StakingDepositContractSet // Event containing the contract specifics and raw log

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
func (it *StakingDepositContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDepositContractSet)
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
		it.Event = new(StakingDepositContractSet)
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
func (it *StakingDepositContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDepositContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDepositContractSet represents a DepositContractSet event raised by the Staking contract.
type StakingDepositContractSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDepositContractSet is a free log retrieval operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_Staking *StakingFilterer) FilterDepositContractSet(opts *bind.FilterOpts) (*StakingDepositContractSetIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DepositContractSet")
	if err != nil {
		return nil, err
	}
	return &StakingDepositContractSetIterator{contract: _Staking.contract, event: "DepositContractSet", logs: logs, sub: sub}, nil
}

// WatchDepositContractSet is a free log subscription operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_Staking *StakingFilterer) WatchDepositContractSet(opts *bind.WatchOpts, sink chan<- *StakingDepositContractSet) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DepositContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDepositContractSet)
				if err := _Staking.contract.UnpackLog(event, "DepositContractSet", log); err != nil {
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

// ParseDepositContractSet is a log parse operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_Staking *StakingFilterer) ParseDepositContractSet(log types.Log) (*StakingDepositContractSet, error) {
	event := new(StakingDepositContractSet)
	if err := _Staking.contract.UnpackLog(event, "DepositContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerAccountSetIterator is returned from FilterManagerAccountSet and is used to iterate over the raw logs and unpacked data for ManagerAccountSet events raised by the Staking contract.
type StakingManagerAccountSetIterator struct {
	Event *StakingManagerAccountSet // Event containing the contract specifics and raw log

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
func (it *StakingManagerAccountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerAccountSet)
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
		it.Event = new(StakingManagerAccountSet)
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
func (it *StakingManagerAccountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerAccountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerAccountSet represents a ManagerAccountSet event raised by the Staking contract.
type StakingManagerAccountSet struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAccountSet is a free log retrieval operation binding the contract event 0x3aa9ccc8285bfd4a1ad6b4cbde39a3b91da051237323010b14c1fa4130be9b7f.
//
// Solidity: event ManagerAccountSet(address account)
func (_Staking *StakingFilterer) FilterManagerAccountSet(opts *bind.FilterOpts) (*StakingManagerAccountSetIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ManagerAccountSet")
	if err != nil {
		return nil, err
	}
	return &StakingManagerAccountSetIterator{contract: _Staking.contract, event: "ManagerAccountSet", logs: logs, sub: sub}, nil
}

// WatchManagerAccountSet is a free log subscription operation binding the contract event 0x3aa9ccc8285bfd4a1ad6b4cbde39a3b91da051237323010b14c1fa4130be9b7f.
//
// Solidity: event ManagerAccountSet(address account)
func (_Staking *StakingFilterer) WatchManagerAccountSet(opts *bind.WatchOpts, sink chan<- *StakingManagerAccountSet) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ManagerAccountSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerAccountSet)
				if err := _Staking.contract.UnpackLog(event, "ManagerAccountSet", log); err != nil {
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

// ParseManagerAccountSet is a log parse operation binding the contract event 0x3aa9ccc8285bfd4a1ad6b4cbde39a3b91da051237323010b14c1fa4130be9b7f.
//
// Solidity: event ManagerAccountSet(address account)
func (_Staking *StakingFilterer) ParseManagerAccountSet(log types.Log) (*StakingManagerAccountSet, error) {
	event := new(StakingManagerAccountSet)
	if err := _Staking.contract.UnpackLog(event, "ManagerAccountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerFeeSetIterator is returned from FilterManagerFeeSet and is used to iterate over the raw logs and unpacked data for ManagerFeeSet events raised by the Staking contract.
type StakingManagerFeeSetIterator struct {
	Event *StakingManagerFeeSet // Event containing the contract specifics and raw log

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
func (it *StakingManagerFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerFeeSet)
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
		it.Event = new(StakingManagerFeeSet)
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
func (it *StakingManagerFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerFeeSet represents a ManagerFeeSet event raised by the Staking contract.
type StakingManagerFeeSet struct {
	Milli *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeSet is a free log retrieval operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_Staking *StakingFilterer) FilterManagerFeeSet(opts *bind.FilterOpts) (*StakingManagerFeeSetIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ManagerFeeSet")
	if err != nil {
		return nil, err
	}
	return &StakingManagerFeeSetIterator{contract: _Staking.contract, event: "ManagerFeeSet", logs: logs, sub: sub}, nil
}

// WatchManagerFeeSet is a free log subscription operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_Staking *StakingFilterer) WatchManagerFeeSet(opts *bind.WatchOpts, sink chan<- *StakingManagerFeeSet) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ManagerFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerFeeSet)
				if err := _Staking.contract.UnpackLog(event, "ManagerFeeSet", log); err != nil {
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

// ParseManagerFeeSet is a log parse operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_Staking *StakingFilterer) ParseManagerFeeSet(log types.Log) (*StakingManagerFeeSet, error) {
	event := new(StakingManagerFeeSet)
	if err := _Staking.contract.UnpackLog(event, "ManagerFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerFeeWithdrawedIterator is returned from FilterManagerFeeWithdrawed and is used to iterate over the raw logs and unpacked data for ManagerFeeWithdrawed events raised by the Staking contract.
type StakingManagerFeeWithdrawedIterator struct {
	Event *StakingManagerFeeWithdrawed // Event containing the contract specifics and raw log

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
func (it *StakingManagerFeeWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerFeeWithdrawed)
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
		it.Event = new(StakingManagerFeeWithdrawed)
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
func (it *StakingManagerFeeWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerFeeWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerFeeWithdrawed represents a ManagerFeeWithdrawed event raised by the Staking contract.
type StakingManagerFeeWithdrawed struct {
	Amount *big.Int
	Arg1   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeWithdrawed is a free log retrieval operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address arg1)
func (_Staking *StakingFilterer) FilterManagerFeeWithdrawed(opts *bind.FilterOpts) (*StakingManagerFeeWithdrawedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return &StakingManagerFeeWithdrawedIterator{contract: _Staking.contract, event: "ManagerFeeWithdrawed", logs: logs, sub: sub}, nil
}

// WatchManagerFeeWithdrawed is a free log subscription operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address arg1)
func (_Staking *StakingFilterer) WatchManagerFeeWithdrawed(opts *bind.WatchOpts, sink chan<- *StakingManagerFeeWithdrawed) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerFeeWithdrawed)
				if err := _Staking.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
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

// ParseManagerFeeWithdrawed is a log parse operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address arg1)
func (_Staking *StakingFilterer) ParseManagerFeeWithdrawed(log types.Log) (*StakingManagerFeeWithdrawed, error) {
	event := new(StakingManagerFeeWithdrawed)
	if err := _Staking.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRevenueCompoundedIterator is returned from FilterManagerRevenueCompounded and is used to iterate over the raw logs and unpacked data for ManagerRevenueCompounded events raised by the Staking contract.
type StakingManagerRevenueCompoundedIterator struct {
	Event *StakingManagerRevenueCompounded // Event containing the contract specifics and raw log

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
func (it *StakingManagerRevenueCompoundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRevenueCompounded)
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
		it.Event = new(StakingManagerRevenueCompounded)
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
func (it *StakingManagerRevenueCompoundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRevenueCompoundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRevenueCompounded represents a ManagerRevenueCompounded event raised by the Staking contract.
type StakingManagerRevenueCompounded struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterManagerRevenueCompounded is a free log retrieval operation binding the contract event 0x20af1949d4837141c257165539bba8e24981f6d6c45872735b01fbf85848e5db.
//
// Solidity: event ManagerRevenueCompounded(uint256 amount)
func (_Staking *StakingFilterer) FilterManagerRevenueCompounded(opts *bind.FilterOpts) (*StakingManagerRevenueCompoundedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ManagerRevenueCompounded")
	if err != nil {
		return nil, err
	}
	return &StakingManagerRevenueCompoundedIterator{contract: _Staking.contract, event: "ManagerRevenueCompounded", logs: logs, sub: sub}, nil
}

// WatchManagerRevenueCompounded is a free log subscription operation binding the contract event 0x20af1949d4837141c257165539bba8e24981f6d6c45872735b01fbf85848e5db.
//
// Solidity: event ManagerRevenueCompounded(uint256 amount)
func (_Staking *StakingFilterer) WatchManagerRevenueCompounded(opts *bind.WatchOpts, sink chan<- *StakingManagerRevenueCompounded) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ManagerRevenueCompounded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRevenueCompounded)
				if err := _Staking.contract.UnpackLog(event, "ManagerRevenueCompounded", log); err != nil {
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

// ParseManagerRevenueCompounded is a log parse operation binding the contract event 0x20af1949d4837141c257165539bba8e24981f6d6c45872735b01fbf85848e5db.
//
// Solidity: event ManagerRevenueCompounded(uint256 amount)
func (_Staking *StakingFilterer) ParseManagerRevenueCompounded(log types.Log) (*StakingManagerRevenueCompounded, error) {
	event := new(StakingManagerRevenueCompounded)
	if err := _Staking.contract.UnpackLog(event, "ManagerRevenueCompounded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Staking contract.
type StakingPausedIterator struct {
	Event *StakingPaused // Event containing the contract specifics and raw log

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
func (it *StakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPaused)
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
		it.Event = new(StakingPaused)
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
func (it *StakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPaused represents a Paused event raised by the Staking contract.
type StakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingPausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingPausedIterator{contract: _Staking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingPaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPaused)
				if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) ParsePaused(log types.Log) (*StakingPaused, error) {
	event := new(StakingPaused)
	if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRestakingAddressSetIterator is returned from FilterRestakingAddressSet and is used to iterate over the raw logs and unpacked data for RestakingAddressSet events raised by the Staking contract.
type StakingRestakingAddressSetIterator struct {
	Event *StakingRestakingAddressSet // Event containing the contract specifics and raw log

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
func (it *StakingRestakingAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRestakingAddressSet)
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
		it.Event = new(StakingRestakingAddressSet)
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
func (it *StakingRestakingAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRestakingAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRestakingAddressSet represents a RestakingAddressSet event raised by the Staking contract.
type StakingRestakingAddressSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRestakingAddressSet is a free log retrieval operation binding the contract event 0x834a16436cbc9545643638861484439875a76defa59e448db456c67d36e4e417.
//
// Solidity: event RestakingAddressSet(address addr)
func (_Staking *StakingFilterer) FilterRestakingAddressSet(opts *bind.FilterOpts) (*StakingRestakingAddressSetIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RestakingAddressSet")
	if err != nil {
		return nil, err
	}
	return &StakingRestakingAddressSetIterator{contract: _Staking.contract, event: "RestakingAddressSet", logs: logs, sub: sub}, nil
}

// WatchRestakingAddressSet is a free log subscription operation binding the contract event 0x834a16436cbc9545643638861484439875a76defa59e448db456c67d36e4e417.
//
// Solidity: event RestakingAddressSet(address addr)
func (_Staking *StakingFilterer) WatchRestakingAddressSet(opts *bind.WatchOpts, sink chan<- *StakingRestakingAddressSet) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RestakingAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRestakingAddressSet)
				if err := _Staking.contract.UnpackLog(event, "RestakingAddressSet", log); err != nil {
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

// ParseRestakingAddressSet is a log parse operation binding the contract event 0x834a16436cbc9545643638861484439875a76defa59e448db456c67d36e4e417.
//
// Solidity: event RestakingAddressSet(address addr)
func (_Staking *StakingFilterer) ParseRestakingAddressSet(log types.Log) (*StakingRestakingAddressSet, error) {
	event := new(StakingRestakingAddressSet)
	if err := _Staking.contract.UnpackLog(event, "RestakingAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRevenueAccountedIterator is returned from FilterRevenueAccounted and is used to iterate over the raw logs and unpacked data for RevenueAccounted events raised by the Staking contract.
type StakingRevenueAccountedIterator struct {
	Event *StakingRevenueAccounted // Event containing the contract specifics and raw log

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
func (it *StakingRevenueAccountedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRevenueAccounted)
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
		it.Event = new(StakingRevenueAccounted)
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
func (it *StakingRevenueAccountedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRevenueAccountedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRevenueAccounted represents a RevenueAccounted event raised by the Staking contract.
type StakingRevenueAccounted struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevenueAccounted is a free log retrieval operation binding the contract event 0x82f24840c0f58d92529afdd441950ddc6e8f2d60138d4458a8d74ba367540cda.
//
// Solidity: event RevenueAccounted(uint256 amount)
func (_Staking *StakingFilterer) FilterRevenueAccounted(opts *bind.FilterOpts) (*StakingRevenueAccountedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RevenueAccounted")
	if err != nil {
		return nil, err
	}
	return &StakingRevenueAccountedIterator{contract: _Staking.contract, event: "RevenueAccounted", logs: logs, sub: sub}, nil
}

// WatchRevenueAccounted is a free log subscription operation binding the contract event 0x82f24840c0f58d92529afdd441950ddc6e8f2d60138d4458a8d74ba367540cda.
//
// Solidity: event RevenueAccounted(uint256 amount)
func (_Staking *StakingFilterer) WatchRevenueAccounted(opts *bind.WatchOpts, sink chan<- *StakingRevenueAccounted) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RevenueAccounted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRevenueAccounted)
				if err := _Staking.contract.UnpackLog(event, "RevenueAccounted", log); err != nil {
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

// ParseRevenueAccounted is a log parse operation binding the contract event 0x82f24840c0f58d92529afdd441950ddc6e8f2d60138d4458a8d74ba367540cda.
//
// Solidity: event RevenueAccounted(uint256 amount)
func (_Staking *StakingFilterer) ParseRevenueAccounted(log types.Log) (*StakingRevenueAccounted, error) {
	event := new(StakingRevenueAccounted)
	if err := _Staking.contract.UnpackLog(event, "RevenueAccounted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Staking contract.
type StakingRoleAdminChangedIterator struct {
	Event *StakingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleAdminChanged)
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
		it.Event = new(StakingRoleAdminChanged)
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
func (it *StakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleAdminChanged represents a RoleAdminChanged event raised by the Staking contract.
type StakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Staking *StakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleAdminChangedIterator{contract: _Staking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Staking *StakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleAdminChanged)
				if err := _Staking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleAdminChanged(log types.Log) (*StakingRoleAdminChanged, error) {
	event := new(StakingRoleAdminChanged)
	if err := _Staking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Staking contract.
type StakingRoleGrantedIterator struct {
	Event *StakingRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleGranted)
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
		it.Event = new(StakingRoleGranted)
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
func (it *StakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleGranted represents a RoleGranted event raised by the Staking contract.
type StakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingRoleGrantedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleGrantedIterator{contract: _Staking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleGranted)
				if err := _Staking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleGranted(log types.Log) (*StakingRoleGranted, error) {
	event := new(StakingRoleGranted)
	if err := _Staking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Staking contract.
type StakingRoleRevokedIterator struct {
	Event *StakingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleRevoked)
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
		it.Event = new(StakingRoleRevoked)
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
func (it *StakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleRevoked represents a RoleRevoked event raised by the Staking contract.
type StakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingRoleRevokedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleRevokedIterator{contract: _Staking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleRevoked)
				if err := _Staking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleRevoked(log types.Log) (*StakingRoleRevoked, error) {
	event := new(StakingRoleRevoked)
	if err := _Staking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Staking contract.
type StakingUnpausedIterator struct {
	Event *StakingUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnpaused)
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
		it.Event = new(StakingUnpaused)
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
func (it *StakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnpaused represents a Unpaused event raised by the Staking contract.
type StakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingUnpausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingUnpausedIterator{contract: _Staking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnpaused)
				if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) ParseUnpaused(log types.Log) (*StakingUnpaused, error) {
	event := new(StakingUnpaused)
	if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUserRevenueCompoundedIterator is returned from FilterUserRevenueCompounded and is used to iterate over the raw logs and unpacked data for UserRevenueCompounded events raised by the Staking contract.
type StakingUserRevenueCompoundedIterator struct {
	Event *StakingUserRevenueCompounded // Event containing the contract specifics and raw log

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
func (it *StakingUserRevenueCompoundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUserRevenueCompounded)
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
		it.Event = new(StakingUserRevenueCompounded)
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
func (it *StakingUserRevenueCompoundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUserRevenueCompoundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUserRevenueCompounded represents a UserRevenueCompounded event raised by the Staking contract.
type StakingUserRevenueCompounded struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUserRevenueCompounded is a free log retrieval operation binding the contract event 0x0605699e1cc0752b1c63cb9863a6f07f7ff8f525c69391ab42f8af6b14d9a2bf.
//
// Solidity: event UserRevenueCompounded(uint256 amount)
func (_Staking *StakingFilterer) FilterUserRevenueCompounded(opts *bind.FilterOpts) (*StakingUserRevenueCompoundedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "UserRevenueCompounded")
	if err != nil {
		return nil, err
	}
	return &StakingUserRevenueCompoundedIterator{contract: _Staking.contract, event: "UserRevenueCompounded", logs: logs, sub: sub}, nil
}

// WatchUserRevenueCompounded is a free log subscription operation binding the contract event 0x0605699e1cc0752b1c63cb9863a6f07f7ff8f525c69391ab42f8af6b14d9a2bf.
//
// Solidity: event UserRevenueCompounded(uint256 amount)
func (_Staking *StakingFilterer) WatchUserRevenueCompounded(opts *bind.WatchOpts, sink chan<- *StakingUserRevenueCompounded) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "UserRevenueCompounded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUserRevenueCompounded)
				if err := _Staking.contract.UnpackLog(event, "UserRevenueCompounded", log); err != nil {
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

// ParseUserRevenueCompounded is a log parse operation binding the contract event 0x0605699e1cc0752b1c63cb9863a6f07f7ff8f525c69391ab42f8af6b14d9a2bf.
//
// Solidity: event UserRevenueCompounded(uint256 amount)
func (_Staking *StakingFilterer) ParseUserRevenueCompounded(log types.Log) (*StakingUserRevenueCompounded, error) {
	event := new(StakingUserRevenueCompounded)
	if err := _Staking.contract.UnpackLog(event, "UserRevenueCompounded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorActivatedIterator is returned from FilterValidatorActivated and is used to iterate over the raw logs and unpacked data for ValidatorActivated events raised by the Staking contract.
type StakingValidatorActivatedIterator struct {
	Event *StakingValidatorActivated // Event containing the contract specifics and raw log

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
func (it *StakingValidatorActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorActivated)
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
		it.Event = new(StakingValidatorActivated)
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
func (it *StakingValidatorActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorActivated represents a ValidatorActivated event raised by the Staking contract.
type StakingValidatorActivated struct {
	NextValidatorId *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorActivated is a free log retrieval operation binding the contract event 0xe2a191ee805447bcf5adabadd39cb816b1b46de1364263aef69980bdafd8370f.
//
// Solidity: event ValidatorActivated(uint256 nextValidatorId)
func (_Staking *StakingFilterer) FilterValidatorActivated(opts *bind.FilterOpts) (*StakingValidatorActivatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorActivated")
	if err != nil {
		return nil, err
	}
	return &StakingValidatorActivatedIterator{contract: _Staking.contract, event: "ValidatorActivated", logs: logs, sub: sub}, nil
}

// WatchValidatorActivated is a free log subscription operation binding the contract event 0xe2a191ee805447bcf5adabadd39cb816b1b46de1364263aef69980bdafd8370f.
//
// Solidity: event ValidatorActivated(uint256 nextValidatorId)
func (_Staking *StakingFilterer) WatchValidatorActivated(opts *bind.WatchOpts, sink chan<- *StakingValidatorActivated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorActivated)
				if err := _Staking.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
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

// ParseValidatorActivated is a log parse operation binding the contract event 0xe2a191ee805447bcf5adabadd39cb816b1b46de1364263aef69980bdafd8370f.
//
// Solidity: event ValidatorActivated(uint256 nextValidatorId)
func (_Staking *StakingFilterer) ParseValidatorActivated(log types.Log) (*StakingValidatorActivated, error) {
	event := new(StakingValidatorActivated)
	if err := _Staking.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorSlashedStoppedIterator is returned from FilterValidatorSlashedStopped and is used to iterate over the raw logs and unpacked data for ValidatorSlashedStopped events raised by the Staking contract.
type StakingValidatorSlashedStoppedIterator struct {
	Event *StakingValidatorSlashedStopped // Event containing the contract specifics and raw log

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
func (it *StakingValidatorSlashedStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorSlashedStopped)
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
		it.Event = new(StakingValidatorSlashedStopped)
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
func (it *StakingValidatorSlashedStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorSlashedStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorSlashedStopped represents a ValidatorSlashedStopped event raised by the Staking contract.
type StakingValidatorSlashedStopped struct {
	StoppedCount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashedStopped is a free log retrieval operation binding the contract event 0x0c5c941f6f7b3b7b1624e8a6738cdc96705c25510de1f11ff37671927a5c47c0.
//
// Solidity: event ValidatorSlashedStopped(uint256 stoppedCount)
func (_Staking *StakingFilterer) FilterValidatorSlashedStopped(opts *bind.FilterOpts) (*StakingValidatorSlashedStoppedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorSlashedStopped")
	if err != nil {
		return nil, err
	}
	return &StakingValidatorSlashedStoppedIterator{contract: _Staking.contract, event: "ValidatorSlashedStopped", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashedStopped is a free log subscription operation binding the contract event 0x0c5c941f6f7b3b7b1624e8a6738cdc96705c25510de1f11ff37671927a5c47c0.
//
// Solidity: event ValidatorSlashedStopped(uint256 stoppedCount)
func (_Staking *StakingFilterer) WatchValidatorSlashedStopped(opts *bind.WatchOpts, sink chan<- *StakingValidatorSlashedStopped) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorSlashedStopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorSlashedStopped)
				if err := _Staking.contract.UnpackLog(event, "ValidatorSlashedStopped", log); err != nil {
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

// ParseValidatorSlashedStopped is a log parse operation binding the contract event 0x0c5c941f6f7b3b7b1624e8a6738cdc96705c25510de1f11ff37671927a5c47c0.
//
// Solidity: event ValidatorSlashedStopped(uint256 stoppedCount)
func (_Staking *StakingFilterer) ParseValidatorSlashedStopped(log types.Log) (*StakingValidatorSlashedStopped, error) {
	event := new(StakingValidatorSlashedStopped)
	if err := _Staking.contract.UnpackLog(event, "ValidatorSlashedStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorStoppedIterator is returned from FilterValidatorStopped and is used to iterate over the raw logs and unpacked data for ValidatorStopped events raised by the Staking contract.
type StakingValidatorStoppedIterator struct {
	Event *StakingValidatorStopped // Event containing the contract specifics and raw log

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
func (it *StakingValidatorStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorStopped)
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
		it.Event = new(StakingValidatorStopped)
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
func (it *StakingValidatorStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorStopped represents a ValidatorStopped event raised by the Staking contract.
type StakingValidatorStopped struct {
	StoppedCount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorStopped is a free log retrieval operation binding the contract event 0xf25558665a382a9abb684f20b20021df5923b51485bbf2829ff0089b5b271410.
//
// Solidity: event ValidatorStopped(uint256 stoppedCount)
func (_Staking *StakingFilterer) FilterValidatorStopped(opts *bind.FilterOpts) (*StakingValidatorStoppedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorStopped")
	if err != nil {
		return nil, err
	}
	return &StakingValidatorStoppedIterator{contract: _Staking.contract, event: "ValidatorStopped", logs: logs, sub: sub}, nil
}

// WatchValidatorStopped is a free log subscription operation binding the contract event 0xf25558665a382a9abb684f20b20021df5923b51485bbf2829ff0089b5b271410.
//
// Solidity: event ValidatorStopped(uint256 stoppedCount)
func (_Staking *StakingFilterer) WatchValidatorStopped(opts *bind.WatchOpts, sink chan<- *StakingValidatorStopped) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorStopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorStopped)
				if err := _Staking.contract.UnpackLog(event, "ValidatorStopped", log); err != nil {
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

// ParseValidatorStopped is a log parse operation binding the contract event 0xf25558665a382a9abb684f20b20021df5923b51485bbf2829ff0089b5b271410.
//
// Solidity: event ValidatorStopped(uint256 stoppedCount)
func (_Staking *StakingFilterer) ParseValidatorStopped(log types.Log) (*StakingValidatorStopped, error) {
	event := new(StakingValidatorStopped)
	if err := _Staking.contract.UnpackLog(event, "ValidatorStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWhiteListToggleIterator is returned from FilterWhiteListToggle and is used to iterate over the raw logs and unpacked data for WhiteListToggle events raised by the Staking contract.
type StakingWhiteListToggleIterator struct {
	Event *StakingWhiteListToggle // Event containing the contract specifics and raw log

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
func (it *StakingWhiteListToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWhiteListToggle)
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
		it.Event = new(StakingWhiteListToggle)
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
func (it *StakingWhiteListToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWhiteListToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWhiteListToggle represents a WhiteListToggle event raised by the Staking contract.
type StakingWhiteListToggle struct {
	Account common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhiteListToggle is a free log retrieval operation binding the contract event 0x96993f732e425c7615ed977e16e6e83d84bb45203ca23885018335b66c2e85be.
//
// Solidity: event WhiteListToggle(address account, bool enabled)
func (_Staking *StakingFilterer) FilterWhiteListToggle(opts *bind.FilterOpts) (*StakingWhiteListToggleIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WhiteListToggle")
	if err != nil {
		return nil, err
	}
	return &StakingWhiteListToggleIterator{contract: _Staking.contract, event: "WhiteListToggle", logs: logs, sub: sub}, nil
}

// WatchWhiteListToggle is a free log subscription operation binding the contract event 0x96993f732e425c7615ed977e16e6e83d84bb45203ca23885018335b66c2e85be.
//
// Solidity: event WhiteListToggle(address account, bool enabled)
func (_Staking *StakingFilterer) WatchWhiteListToggle(opts *bind.WatchOpts, sink chan<- *StakingWhiteListToggle) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WhiteListToggle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWhiteListToggle)
				if err := _Staking.contract.UnpackLog(event, "WhiteListToggle", log); err != nil {
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

// ParseWhiteListToggle is a log parse operation binding the contract event 0x96993f732e425c7615ed977e16e6e83d84bb45203ca23885018335b66c2e85be.
//
// Solidity: event WhiteListToggle(address account, bool enabled)
func (_Staking *StakingFilterer) ParseWhiteListToggle(log types.Log) (*StakingWhiteListToggle, error) {
	event := new(StakingWhiteListToggle)
	if err := _Staking.contract.UnpackLog(event, "WhiteListToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWithdrawCredentialSetIterator is returned from FilterWithdrawCredentialSet and is used to iterate over the raw logs and unpacked data for WithdrawCredentialSet events raised by the Staking contract.
type StakingWithdrawCredentialSetIterator struct {
	Event *StakingWithdrawCredentialSet // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawCredentialSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdrawCredentialSet)
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
		it.Event = new(StakingWithdrawCredentialSet)
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
func (it *StakingWithdrawCredentialSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawCredentialSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdrawCredentialSet represents a WithdrawCredentialSet event raised by the Staking contract.
type StakingWithdrawCredentialSet struct {
	WithdrawCredential [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCredentialSet is a free log retrieval operation binding the contract event 0x690facbfacb53c9319489117a6ac422718b5cb059a6ffade4871ff10f6f9aee9.
//
// Solidity: event WithdrawCredentialSet(bytes32 withdrawCredential)
func (_Staking *StakingFilterer) FilterWithdrawCredentialSet(opts *bind.FilterOpts) (*StakingWithdrawCredentialSetIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WithdrawCredentialSet")
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawCredentialSetIterator{contract: _Staking.contract, event: "WithdrawCredentialSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawCredentialSet is a free log subscription operation binding the contract event 0x690facbfacb53c9319489117a6ac422718b5cb059a6ffade4871ff10f6f9aee9.
//
// Solidity: event WithdrawCredentialSet(bytes32 withdrawCredential)
func (_Staking *StakingFilterer) WatchWithdrawCredentialSet(opts *bind.WatchOpts, sink chan<- *StakingWithdrawCredentialSet) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WithdrawCredentialSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdrawCredentialSet)
				if err := _Staking.contract.UnpackLog(event, "WithdrawCredentialSet", log); err != nil {
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

// ParseWithdrawCredentialSet is a log parse operation binding the contract event 0x690facbfacb53c9319489117a6ac422718b5cb059a6ffade4871ff10f6f9aee9.
//
// Solidity: event WithdrawCredentialSet(bytes32 withdrawCredential)
func (_Staking *StakingFilterer) ParseWithdrawCredentialSet(log types.Log) (*StakingWithdrawCredentialSet, error) {
	event := new(StakingWithdrawCredentialSet)
	if err := _Staking.contract.UnpackLog(event, "WithdrawCredentialSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
