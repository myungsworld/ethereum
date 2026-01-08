// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CounterV1MetaData contains all meta data concerning the CounterV1 contract.
var CounterV1MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"CountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102d18061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80632baeceb71461006457806354fd4d501461006e5780638129fc1c1461009b5780638da5cb5b146100a3578063a87d942c146100be578063d09de08a146100ce575b5f5ffd5b61006c6100d6565b005b6040805180820182526002815261563160f01b602082015290516100929190610226565b60405180910390f35b61006c61017e565b6001546040516001600160a01b039091168152602001610092565b5f54604051908152602001610092565b61006c610215565b5f5f541161012b5760405162461bcd60e51b815260206004820152601d60248201527f436f756e7465723a2063616e6e6f7420676f2062656c6f77207a65726f00000060448201526064015b60405180910390fd5b60015f5f82825461013c919061026f565b90915550505f54604080519182523360208301527fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a8291015b60405180910390a1565b6001546001600160a01b0316156101cd5760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b6044820152606401610122565b600180546001600160a01b031916339081179091555f80556040519081527f908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e690602001610174565b60015f5f82825461013c9190610288565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b634e487b7160e01b5f52601160045260245ffd5b818103818111156102825761028261025b565b92915050565b808201808211156102825761028261025b56fea2646970667358221220d48188008a4fc14a50e368a7606f612f0143b460e368115eb45abfbe35615d6264736f6c63430008210033",
}

// CounterV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterV1MetaData.ABI instead.
var CounterV1ABI = CounterV1MetaData.ABI

// CounterV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterV1MetaData.Bin instead.
var CounterV1Bin = CounterV1MetaData.Bin

// DeployCounterV1 deploys a new Ethereum contract, binding an instance of CounterV1 to it.
func DeployCounterV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CounterV1, error) {
	parsed, err := CounterV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CounterV1{CounterV1Caller: CounterV1Caller{contract: contract}, CounterV1Transactor: CounterV1Transactor{contract: contract}, CounterV1Filterer: CounterV1Filterer{contract: contract}}, nil
}

// CounterV1 is an auto generated Go binding around an Ethereum contract.
type CounterV1 struct {
	CounterV1Caller     // Read-only binding to the contract
	CounterV1Transactor // Write-only binding to the contract
	CounterV1Filterer   // Log filterer for contract events
}

// CounterV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type CounterV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterV1Session struct {
	Contract     *CounterV1        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterV1CallerSession struct {
	Contract *CounterV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CounterV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterV1TransactorSession struct {
	Contract     *CounterV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CounterV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type CounterV1Raw struct {
	Contract *CounterV1 // Generic contract binding to access the raw methods on
}

// CounterV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterV1CallerRaw struct {
	Contract *CounterV1Caller // Generic read-only contract binding to access the raw methods on
}

// CounterV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterV1TransactorRaw struct {
	Contract *CounterV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCounterV1 creates a new instance of CounterV1, bound to a specific deployed contract.
func NewCounterV1(address common.Address, backend bind.ContractBackend) (*CounterV1, error) {
	contract, err := bindCounterV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CounterV1{CounterV1Caller: CounterV1Caller{contract: contract}, CounterV1Transactor: CounterV1Transactor{contract: contract}, CounterV1Filterer: CounterV1Filterer{contract: contract}}, nil
}

// NewCounterV1Caller creates a new read-only instance of CounterV1, bound to a specific deployed contract.
func NewCounterV1Caller(address common.Address, caller bind.ContractCaller) (*CounterV1Caller, error) {
	contract, err := bindCounterV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterV1Caller{contract: contract}, nil
}

// NewCounterV1Transactor creates a new write-only instance of CounterV1, bound to a specific deployed contract.
func NewCounterV1Transactor(address common.Address, transactor bind.ContractTransactor) (*CounterV1Transactor, error) {
	contract, err := bindCounterV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterV1Transactor{contract: contract}, nil
}

// NewCounterV1Filterer creates a new log filterer instance of CounterV1, bound to a specific deployed contract.
func NewCounterV1Filterer(address common.Address, filterer bind.ContractFilterer) (*CounterV1Filterer, error) {
	contract, err := bindCounterV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterV1Filterer{contract: contract}, nil
}

// bindCounterV1 binds a generic wrapper to an already deployed contract.
func bindCounterV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterV1 *CounterV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterV1.Contract.CounterV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterV1 *CounterV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV1.Contract.CounterV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterV1 *CounterV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterV1.Contract.CounterV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterV1 *CounterV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterV1 *CounterV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterV1 *CounterV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterV1.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_CounterV1 *CounterV1Caller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CounterV1.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_CounterV1 *CounterV1Session) GetCount() (*big.Int, error) {
	return _CounterV1.Contract.GetCount(&_CounterV1.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_CounterV1 *CounterV1CallerSession) GetCount() (*big.Int, error) {
	return _CounterV1.Contract.GetCount(&_CounterV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CounterV1 *CounterV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CounterV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CounterV1 *CounterV1Session) Owner() (common.Address, error) {
	return _CounterV1.Contract.Owner(&_CounterV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CounterV1 *CounterV1CallerSession) Owner() (common.Address, error) {
	return _CounterV1.Contract.Owner(&_CounterV1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CounterV1 *CounterV1Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CounterV1.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CounterV1 *CounterV1Session) Version() (string, error) {
	return _CounterV1.Contract.Version(&_CounterV1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CounterV1 *CounterV1CallerSession) Version() (string, error) {
	return _CounterV1.Contract.Version(&_CounterV1.CallOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_CounterV1 *CounterV1Transactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV1.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_CounterV1 *CounterV1Session) Decrement() (*types.Transaction, error) {
	return _CounterV1.Contract.Decrement(&_CounterV1.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_CounterV1 *CounterV1TransactorSession) Decrement() (*types.Transaction, error) {
	return _CounterV1.Contract.Decrement(&_CounterV1.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterV1 *CounterV1Transactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV1.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterV1 *CounterV1Session) Increment() (*types.Transaction, error) {
	return _CounterV1.Contract.Increment(&_CounterV1.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterV1 *CounterV1TransactorSession) Increment() (*types.Transaction, error) {
	return _CounterV1.Contract.Increment(&_CounterV1.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CounterV1 *CounterV1Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV1.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CounterV1 *CounterV1Session) Initialize() (*types.Transaction, error) {
	return _CounterV1.Contract.Initialize(&_CounterV1.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CounterV1 *CounterV1TransactorSession) Initialize() (*types.Transaction, error) {
	return _CounterV1.Contract.Initialize(&_CounterV1.TransactOpts)
}

// CounterV1CountChangedIterator is returned from FilterCountChanged and is used to iterate over the raw logs and unpacked data for CountChanged events raised by the CounterV1 contract.
type CounterV1CountChangedIterator struct {
	Event *CounterV1CountChanged // Event containing the contract specifics and raw log

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
func (it *CounterV1CountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterV1CountChanged)
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
		it.Event = new(CounterV1CountChanged)
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
func (it *CounterV1CountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterV1CountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterV1CountChanged represents a CountChanged event raised by the CounterV1 contract.
type CounterV1CountChanged struct {
	NewCount  *big.Int
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCountChanged is a free log retrieval operation binding the contract event 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82.
//
// Solidity: event CountChanged(uint256 newCount, address changedBy)
func (_CounterV1 *CounterV1Filterer) FilterCountChanged(opts *bind.FilterOpts) (*CounterV1CountChangedIterator, error) {

	logs, sub, err := _CounterV1.contract.FilterLogs(opts, "CountChanged")
	if err != nil {
		return nil, err
	}
	return &CounterV1CountChangedIterator{contract: _CounterV1.contract, event: "CountChanged", logs: logs, sub: sub}, nil
}

// WatchCountChanged is a free log subscription operation binding the contract event 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82.
//
// Solidity: event CountChanged(uint256 newCount, address changedBy)
func (_CounterV1 *CounterV1Filterer) WatchCountChanged(opts *bind.WatchOpts, sink chan<- *CounterV1CountChanged) (event.Subscription, error) {

	logs, sub, err := _CounterV1.contract.WatchLogs(opts, "CountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterV1CountChanged)
				if err := _CounterV1.contract.UnpackLog(event, "CountChanged", log); err != nil {
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

// ParseCountChanged is a log parse operation binding the contract event 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82.
//
// Solidity: event CountChanged(uint256 newCount, address changedBy)
func (_CounterV1 *CounterV1Filterer) ParseCountChanged(log types.Log) (*CounterV1CountChanged, error) {
	event := new(CounterV1CountChanged)
	if err := _CounterV1.contract.UnpackLog(event, "CountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CounterV1 contract.
type CounterV1InitializedIterator struct {
	Event *CounterV1Initialized // Event containing the contract specifics and raw log

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
func (it *CounterV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterV1Initialized)
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
		it.Event = new(CounterV1Initialized)
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
func (it *CounterV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterV1Initialized represents a Initialized event raised by the CounterV1 contract.
type CounterV1Initialized struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e6.
//
// Solidity: event Initialized(address owner)
func (_CounterV1 *CounterV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*CounterV1InitializedIterator, error) {

	logs, sub, err := _CounterV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CounterV1InitializedIterator{contract: _CounterV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e6.
//
// Solidity: event Initialized(address owner)
func (_CounterV1 *CounterV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CounterV1Initialized) (event.Subscription, error) {

	logs, sub, err := _CounterV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterV1Initialized)
				if err := _CounterV1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e6.
//
// Solidity: event Initialized(address owner)
func (_CounterV1 *CounterV1Filterer) ParseInitialized(log types.Log) (*CounterV1Initialized, error) {
	event := new(CounterV1Initialized)
	if err := _CounterV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
