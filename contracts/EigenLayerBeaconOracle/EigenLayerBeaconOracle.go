// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EigenLayerBeaconOracle

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

// EigenLayerBeaconOracleMetaData contains all meta data concerning the EigenLayerBeaconOracle contract.
var EigenLayerBeaconOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_genesisBlockTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidBlockTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampOutOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockRoot\",\"type\":\"bytes32\"}],\"name\":\"EigenLayerBeaconOracleUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GENESIS_BLOCK_TIMESTAMP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetTimestamp\",\"type\":\"uint256\"}],\"name\":\"addTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_targetSlot\",\"type\":\"uint64\"}],\"name\":\"findBlockRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timestampToBlockRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EigenLayerBeaconOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenLayerBeaconOracleMetaData.ABI instead.
var EigenLayerBeaconOracleABI = EigenLayerBeaconOracleMetaData.ABI

// EigenLayerBeaconOracle is an auto generated Go binding around an Ethereum contract.
type EigenLayerBeaconOracle struct {
	EigenLayerBeaconOracleCaller     // Read-only binding to the contract
	EigenLayerBeaconOracleTransactor // Write-only binding to the contract
	EigenLayerBeaconOracleFilterer   // Log filterer for contract events
}

// EigenLayerBeaconOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type EigenLayerBeaconOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenLayerBeaconOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EigenLayerBeaconOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenLayerBeaconOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EigenLayerBeaconOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenLayerBeaconOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EigenLayerBeaconOracleSession struct {
	Contract     *EigenLayerBeaconOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EigenLayerBeaconOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EigenLayerBeaconOracleCallerSession struct {
	Contract *EigenLayerBeaconOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// EigenLayerBeaconOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EigenLayerBeaconOracleTransactorSession struct {
	Contract     *EigenLayerBeaconOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// EigenLayerBeaconOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type EigenLayerBeaconOracleRaw struct {
	Contract *EigenLayerBeaconOracle // Generic contract binding to access the raw methods on
}

// EigenLayerBeaconOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EigenLayerBeaconOracleCallerRaw struct {
	Contract *EigenLayerBeaconOracleCaller // Generic read-only contract binding to access the raw methods on
}

// EigenLayerBeaconOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EigenLayerBeaconOracleTransactorRaw struct {
	Contract *EigenLayerBeaconOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEigenLayerBeaconOracle creates a new instance of EigenLayerBeaconOracle, bound to a specific deployed contract.
func NewEigenLayerBeaconOracle(address common.Address, backend bind.ContractBackend) (*EigenLayerBeaconOracle, error) {
	contract, err := bindEigenLayerBeaconOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EigenLayerBeaconOracle{EigenLayerBeaconOracleCaller: EigenLayerBeaconOracleCaller{contract: contract}, EigenLayerBeaconOracleTransactor: EigenLayerBeaconOracleTransactor{contract: contract}, EigenLayerBeaconOracleFilterer: EigenLayerBeaconOracleFilterer{contract: contract}}, nil
}

// NewEigenLayerBeaconOracleCaller creates a new read-only instance of EigenLayerBeaconOracle, bound to a specific deployed contract.
func NewEigenLayerBeaconOracleCaller(address common.Address, caller bind.ContractCaller) (*EigenLayerBeaconOracleCaller, error) {
	contract, err := bindEigenLayerBeaconOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EigenLayerBeaconOracleCaller{contract: contract}, nil
}

// NewEigenLayerBeaconOracleTransactor creates a new write-only instance of EigenLayerBeaconOracle, bound to a specific deployed contract.
func NewEigenLayerBeaconOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*EigenLayerBeaconOracleTransactor, error) {
	contract, err := bindEigenLayerBeaconOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EigenLayerBeaconOracleTransactor{contract: contract}, nil
}

// NewEigenLayerBeaconOracleFilterer creates a new log filterer instance of EigenLayerBeaconOracle, bound to a specific deployed contract.
func NewEigenLayerBeaconOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*EigenLayerBeaconOracleFilterer, error) {
	contract, err := bindEigenLayerBeaconOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EigenLayerBeaconOracleFilterer{contract: contract}, nil
}

// bindEigenLayerBeaconOracle binds a generic wrapper to an already deployed contract.
func bindEigenLayerBeaconOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EigenLayerBeaconOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenLayerBeaconOracle.Contract.EigenLayerBeaconOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenLayerBeaconOracle.Contract.EigenLayerBeaconOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenLayerBeaconOracle.Contract.EigenLayerBeaconOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenLayerBeaconOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenLayerBeaconOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenLayerBeaconOracle.Contract.contract.Transact(opts, method, params...)
}

// GENESISBLOCKTIMESTAMP is a free data retrieval call binding the contract method 0xaa4a95c7.
//
// Solidity: function GENESIS_BLOCK_TIMESTAMP() view returns(uint256)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleCaller) GENESISBLOCKTIMESTAMP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenLayerBeaconOracle.contract.Call(opts, &out, "GENESIS_BLOCK_TIMESTAMP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISBLOCKTIMESTAMP is a free data retrieval call binding the contract method 0xaa4a95c7.
//
// Solidity: function GENESIS_BLOCK_TIMESTAMP() view returns(uint256)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleSession) GENESISBLOCKTIMESTAMP() (*big.Int, error) {
	return _EigenLayerBeaconOracle.Contract.GENESISBLOCKTIMESTAMP(&_EigenLayerBeaconOracle.CallOpts)
}

// GENESISBLOCKTIMESTAMP is a free data retrieval call binding the contract method 0xaa4a95c7.
//
// Solidity: function GENESIS_BLOCK_TIMESTAMP() view returns(uint256)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleCallerSession) GENESISBLOCKTIMESTAMP() (*big.Int, error) {
	return _EigenLayerBeaconOracle.Contract.GENESISBLOCKTIMESTAMP(&_EigenLayerBeaconOracle.CallOpts)
}

// FindBlockRoot is a free data retrieval call binding the contract method 0x5b6ae8c0.
//
// Solidity: function findBlockRoot(uint64 _targetSlot) view returns(bytes32 blockRoot)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleCaller) FindBlockRoot(opts *bind.CallOpts, _targetSlot uint64) ([32]byte, error) {
	var out []interface{}
	err := _EigenLayerBeaconOracle.contract.Call(opts, &out, "findBlockRoot", _targetSlot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FindBlockRoot is a free data retrieval call binding the contract method 0x5b6ae8c0.
//
// Solidity: function findBlockRoot(uint64 _targetSlot) view returns(bytes32 blockRoot)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleSession) FindBlockRoot(_targetSlot uint64) ([32]byte, error) {
	return _EigenLayerBeaconOracle.Contract.FindBlockRoot(&_EigenLayerBeaconOracle.CallOpts, _targetSlot)
}

// FindBlockRoot is a free data retrieval call binding the contract method 0x5b6ae8c0.
//
// Solidity: function findBlockRoot(uint64 _targetSlot) view returns(bytes32 blockRoot)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleCallerSession) FindBlockRoot(_targetSlot uint64) ([32]byte, error) {
	return _EigenLayerBeaconOracle.Contract.FindBlockRoot(&_EigenLayerBeaconOracle.CallOpts, _targetSlot)
}

// TimestampToBlockRoot is a free data retrieval call binding the contract method 0x643599f2.
//
// Solidity: function timestampToBlockRoot(uint256 ) view returns(bytes32)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleCaller) TimestampToBlockRoot(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EigenLayerBeaconOracle.contract.Call(opts, &out, "timestampToBlockRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimestampToBlockRoot is a free data retrieval call binding the contract method 0x643599f2.
//
// Solidity: function timestampToBlockRoot(uint256 ) view returns(bytes32)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleSession) TimestampToBlockRoot(arg0 *big.Int) ([32]byte, error) {
	return _EigenLayerBeaconOracle.Contract.TimestampToBlockRoot(&_EigenLayerBeaconOracle.CallOpts, arg0)
}

// TimestampToBlockRoot is a free data retrieval call binding the contract method 0x643599f2.
//
// Solidity: function timestampToBlockRoot(uint256 ) view returns(bytes32)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleCallerSession) TimestampToBlockRoot(arg0 *big.Int) ([32]byte, error) {
	return _EigenLayerBeaconOracle.Contract.TimestampToBlockRoot(&_EigenLayerBeaconOracle.CallOpts, arg0)
}

// AddTimestamp is a paid mutator transaction binding the contract method 0x1bb90003.
//
// Solidity: function addTimestamp(uint256 _targetTimestamp) returns()
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleTransactor) AddTimestamp(opts *bind.TransactOpts, _targetTimestamp *big.Int) (*types.Transaction, error) {
	return _EigenLayerBeaconOracle.contract.Transact(opts, "addTimestamp", _targetTimestamp)
}

// AddTimestamp is a paid mutator transaction binding the contract method 0x1bb90003.
//
// Solidity: function addTimestamp(uint256 _targetTimestamp) returns()
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleSession) AddTimestamp(_targetTimestamp *big.Int) (*types.Transaction, error) {
	return _EigenLayerBeaconOracle.Contract.AddTimestamp(&_EigenLayerBeaconOracle.TransactOpts, _targetTimestamp)
}

// AddTimestamp is a paid mutator transaction binding the contract method 0x1bb90003.
//
// Solidity: function addTimestamp(uint256 _targetTimestamp) returns()
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleTransactorSession) AddTimestamp(_targetTimestamp *big.Int) (*types.Transaction, error) {
	return _EigenLayerBeaconOracle.Contract.AddTimestamp(&_EigenLayerBeaconOracle.TransactOpts, _targetTimestamp)
}

// EigenLayerBeaconOracleEigenLayerBeaconOracleUpdateIterator is returned from FilterEigenLayerBeaconOracleUpdate and is used to iterate over the raw logs and unpacked data for EigenLayerBeaconOracleUpdate events raised by the EigenLayerBeaconOracle contract.
type EigenLayerBeaconOracleEigenLayerBeaconOracleUpdateIterator struct {
	Event *EigenLayerBeaconOracleEigenLayerBeaconOracleUpdate // Event containing the contract specifics and raw log

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
func (it *EigenLayerBeaconOracleEigenLayerBeaconOracleUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenLayerBeaconOracleEigenLayerBeaconOracleUpdate)
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
		it.Event = new(EigenLayerBeaconOracleEigenLayerBeaconOracleUpdate)
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
func (it *EigenLayerBeaconOracleEigenLayerBeaconOracleUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenLayerBeaconOracleEigenLayerBeaconOracleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenLayerBeaconOracleEigenLayerBeaconOracleUpdate represents a EigenLayerBeaconOracleUpdate event raised by the EigenLayerBeaconOracle contract.
type EigenLayerBeaconOracleEigenLayerBeaconOracleUpdate struct {
	Slot      *big.Int
	Timestamp *big.Int
	BlockRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEigenLayerBeaconOracleUpdate is a free log retrieval operation binding the contract event 0xc8bd73742a63a8ef565902b94304e65f46c90f6974df076790e322b8f7cbeba0.
//
// Solidity: event EigenLayerBeaconOracleUpdate(uint256 slot, uint256 timestamp, bytes32 blockRoot)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleFilterer) FilterEigenLayerBeaconOracleUpdate(opts *bind.FilterOpts) (*EigenLayerBeaconOracleEigenLayerBeaconOracleUpdateIterator, error) {

	logs, sub, err := _EigenLayerBeaconOracle.contract.FilterLogs(opts, "EigenLayerBeaconOracleUpdate")
	if err != nil {
		return nil, err
	}
	return &EigenLayerBeaconOracleEigenLayerBeaconOracleUpdateIterator{contract: _EigenLayerBeaconOracle.contract, event: "EigenLayerBeaconOracleUpdate", logs: logs, sub: sub}, nil
}

// WatchEigenLayerBeaconOracleUpdate is a free log subscription operation binding the contract event 0xc8bd73742a63a8ef565902b94304e65f46c90f6974df076790e322b8f7cbeba0.
//
// Solidity: event EigenLayerBeaconOracleUpdate(uint256 slot, uint256 timestamp, bytes32 blockRoot)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleFilterer) WatchEigenLayerBeaconOracleUpdate(opts *bind.WatchOpts, sink chan<- *EigenLayerBeaconOracleEigenLayerBeaconOracleUpdate) (event.Subscription, error) {

	logs, sub, err := _EigenLayerBeaconOracle.contract.WatchLogs(opts, "EigenLayerBeaconOracleUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenLayerBeaconOracleEigenLayerBeaconOracleUpdate)
				if err := _EigenLayerBeaconOracle.contract.UnpackLog(event, "EigenLayerBeaconOracleUpdate", log); err != nil {
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

// ParseEigenLayerBeaconOracleUpdate is a log parse operation binding the contract event 0xc8bd73742a63a8ef565902b94304e65f46c90f6974df076790e322b8f7cbeba0.
//
// Solidity: event EigenLayerBeaconOracleUpdate(uint256 slot, uint256 timestamp, bytes32 blockRoot)
func (_EigenLayerBeaconOracle *EigenLayerBeaconOracleFilterer) ParseEigenLayerBeaconOracleUpdate(log types.Log) (*EigenLayerBeaconOracleEigenLayerBeaconOracleUpdate, error) {
	event := new(EigenLayerBeaconOracleEigenLayerBeaconOracleUpdate)
	if err := _EigenLayerBeaconOracle.contract.UnpackLog(event, "EigenLayerBeaconOracleUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
