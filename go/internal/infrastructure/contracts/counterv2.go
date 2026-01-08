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

// CounterV2MetaData contains all meta data concerning the CounterV2 contract.
var CounterV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"CountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resetBy\",\"type\":\"address\"}],\"name\":\"CountReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"incrementBy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"setCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105248061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610090575f3560e01c80638da5cb5b116100635780638da5cb5b146100e6578063a87d942c14610101578063d09de08a14610111578063d14e62b814610119578063d826f88f1461012c575f5ffd5b806303df179c146100945780632baeceb7146100a957806354fd4d50146100b15780638129fc1c146100de575b5f5ffd5b6100a76100a2366004610462565b610134565b005b6100a76101ed565b60408051808201825260028152612b1960f11b602082015290516100d59190610479565b60405180910390f35b6100a7610290565b6001546040516001600160a01b0390911681526020016100d5565b5f546040519081526020016100d5565b6100a7610327565b6100a7610127366004610462565b610338565b6100a76103d5565b5f811161019a5760405162461bcd60e51b815260206004820152602960248201527f436f756e7465723a20616d6f756e74206d7573742062652067726561746572206044820152687468616e207a65726f60b81b60648201526084015b60405180910390fd5b805f5f8282546101aa91906104c2565b90915550505f54604080519182523360208301527fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a8291015b60405180910390a150565b5f5f541161023d5760405162461bcd60e51b815260206004820152601d60248201527f436f756e7465723a2063616e6e6f7420676f2062656c6f77207a65726f0000006044820152606401610191565b60015f5f82825461024e91906104db565b90915550505f54604080519182523360208301527fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a8291015b60405180910390a1565b6001546001600160a01b0316156102df5760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b6044820152606401610191565b600180546001600160a01b031916339081179091555f80556040519081527f908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e690602001610286565b60015f5f82825461024e91906104c2565b6001546001600160a01b0316331461039c5760405162461bcd60e51b815260206004820152602160248201527f436f756e7465723a206f6e6c79206f776e65722063616e2073657420636f756e6044820152601d60fa1b6064820152608401610191565b5f819055604080518281523360208201527fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a8291016101e2565b6001546001600160a01b0316331461042f5760405162461bcd60e51b815260206004820152601d60248201527f436f756e7465723a206f6e6c79206f776e65722063616e2072657365740000006044820152606401610191565b5f80556040513381527fa5ee6258204973c56c5a39c4ac31e61723f410d84f9e8117ba52b76b7cea990c90602001610286565b5f60208284031215610472575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156104d5576104d56104ae565b92915050565b818103818111156104d5576104d56104ae56fea2646970667358221220d9bce519e2118963d48b0d21f6e8ecfa28470a0948749955b781fa563c56f7af64736f6c63430008210033",
}

// CounterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterV2MetaData.ABI instead.
var CounterV2ABI = CounterV2MetaData.ABI

// CounterV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterV2MetaData.Bin instead.
var CounterV2Bin = CounterV2MetaData.Bin

// DeployCounterV2 deploys a new Ethereum contract, binding an instance of CounterV2 to it.
func DeployCounterV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CounterV2, error) {
	parsed, err := CounterV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CounterV2{CounterV2Caller: CounterV2Caller{contract: contract}, CounterV2Transactor: CounterV2Transactor{contract: contract}, CounterV2Filterer: CounterV2Filterer{contract: contract}}, nil
}

// CounterV2 is an auto generated Go binding around an Ethereum contract.
type CounterV2 struct {
	CounterV2Caller     // Read-only binding to the contract
	CounterV2Transactor // Write-only binding to the contract
	CounterV2Filterer   // Log filterer for contract events
}

// CounterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CounterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterV2Session struct {
	Contract     *CounterV2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterV2CallerSession struct {
	Contract *CounterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CounterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterV2TransactorSession struct {
	Contract     *CounterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CounterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CounterV2Raw struct {
	Contract *CounterV2 // Generic contract binding to access the raw methods on
}

// CounterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterV2CallerRaw struct {
	Contract *CounterV2Caller // Generic read-only contract binding to access the raw methods on
}

// CounterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterV2TransactorRaw struct {
	Contract *CounterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCounterV2 creates a new instance of CounterV2, bound to a specific deployed contract.
func NewCounterV2(address common.Address, backend bind.ContractBackend) (*CounterV2, error) {
	contract, err := bindCounterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CounterV2{CounterV2Caller: CounterV2Caller{contract: contract}, CounterV2Transactor: CounterV2Transactor{contract: contract}, CounterV2Filterer: CounterV2Filterer{contract: contract}}, nil
}

// NewCounterV2Caller creates a new read-only instance of CounterV2, bound to a specific deployed contract.
func NewCounterV2Caller(address common.Address, caller bind.ContractCaller) (*CounterV2Caller, error) {
	contract, err := bindCounterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterV2Caller{contract: contract}, nil
}

// NewCounterV2Transactor creates a new write-only instance of CounterV2, bound to a specific deployed contract.
func NewCounterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*CounterV2Transactor, error) {
	contract, err := bindCounterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterV2Transactor{contract: contract}, nil
}

// NewCounterV2Filterer creates a new log filterer instance of CounterV2, bound to a specific deployed contract.
func NewCounterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*CounterV2Filterer, error) {
	contract, err := bindCounterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterV2Filterer{contract: contract}, nil
}

// bindCounterV2 binds a generic wrapper to an already deployed contract.
func bindCounterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterV2 *CounterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterV2.Contract.CounterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterV2 *CounterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV2.Contract.CounterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterV2 *CounterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterV2.Contract.CounterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterV2 *CounterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterV2 *CounterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterV2 *CounterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterV2.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_CounterV2 *CounterV2Caller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CounterV2.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_CounterV2 *CounterV2Session) GetCount() (*big.Int, error) {
	return _CounterV2.Contract.GetCount(&_CounterV2.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_CounterV2 *CounterV2CallerSession) GetCount() (*big.Int, error) {
	return _CounterV2.Contract.GetCount(&_CounterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CounterV2 *CounterV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CounterV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CounterV2 *CounterV2Session) Owner() (common.Address, error) {
	return _CounterV2.Contract.Owner(&_CounterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CounterV2 *CounterV2CallerSession) Owner() (common.Address, error) {
	return _CounterV2.Contract.Owner(&_CounterV2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CounterV2 *CounterV2Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CounterV2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CounterV2 *CounterV2Session) Version() (string, error) {
	return _CounterV2.Contract.Version(&_CounterV2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CounterV2 *CounterV2CallerSession) Version() (string, error) {
	return _CounterV2.Contract.Version(&_CounterV2.CallOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_CounterV2 *CounterV2Transactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV2.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_CounterV2 *CounterV2Session) Decrement() (*types.Transaction, error) {
	return _CounterV2.Contract.Decrement(&_CounterV2.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_CounterV2 *CounterV2TransactorSession) Decrement() (*types.Transaction, error) {
	return _CounterV2.Contract.Decrement(&_CounterV2.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterV2 *CounterV2Transactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV2.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterV2 *CounterV2Session) Increment() (*types.Transaction, error) {
	return _CounterV2.Contract.Increment(&_CounterV2.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterV2 *CounterV2TransactorSession) Increment() (*types.Transaction, error) {
	return _CounterV2.Contract.Increment(&_CounterV2.TransactOpts)
}

// IncrementBy is a paid mutator transaction binding the contract method 0x03df179c.
//
// Solidity: function incrementBy(uint256 amount) returns()
func (_CounterV2 *CounterV2Transactor) IncrementBy(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CounterV2.contract.Transact(opts, "incrementBy", amount)
}

// IncrementBy is a paid mutator transaction binding the contract method 0x03df179c.
//
// Solidity: function incrementBy(uint256 amount) returns()
func (_CounterV2 *CounterV2Session) IncrementBy(amount *big.Int) (*types.Transaction, error) {
	return _CounterV2.Contract.IncrementBy(&_CounterV2.TransactOpts, amount)
}

// IncrementBy is a paid mutator transaction binding the contract method 0x03df179c.
//
// Solidity: function incrementBy(uint256 amount) returns()
func (_CounterV2 *CounterV2TransactorSession) IncrementBy(amount *big.Int) (*types.Transaction, error) {
	return _CounterV2.Contract.IncrementBy(&_CounterV2.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CounterV2 *CounterV2Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV2.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CounterV2 *CounterV2Session) Initialize() (*types.Transaction, error) {
	return _CounterV2.Contract.Initialize(&_CounterV2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CounterV2 *CounterV2TransactorSession) Initialize() (*types.Transaction, error) {
	return _CounterV2.Contract.Initialize(&_CounterV2.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_CounterV2 *CounterV2Transactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterV2.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_CounterV2 *CounterV2Session) Reset() (*types.Transaction, error) {
	return _CounterV2.Contract.Reset(&_CounterV2.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_CounterV2 *CounterV2TransactorSession) Reset() (*types.Transaction, error) {
	return _CounterV2.Contract.Reset(&_CounterV2.TransactOpts)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 newCount) returns()
func (_CounterV2 *CounterV2Transactor) SetCount(opts *bind.TransactOpts, newCount *big.Int) (*types.Transaction, error) {
	return _CounterV2.contract.Transact(opts, "setCount", newCount)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 newCount) returns()
func (_CounterV2 *CounterV2Session) SetCount(newCount *big.Int) (*types.Transaction, error) {
	return _CounterV2.Contract.SetCount(&_CounterV2.TransactOpts, newCount)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 newCount) returns()
func (_CounterV2 *CounterV2TransactorSession) SetCount(newCount *big.Int) (*types.Transaction, error) {
	return _CounterV2.Contract.SetCount(&_CounterV2.TransactOpts, newCount)
}

// CounterV2CountChangedIterator is returned from FilterCountChanged and is used to iterate over the raw logs and unpacked data for CountChanged events raised by the CounterV2 contract.
type CounterV2CountChangedIterator struct {
	Event *CounterV2CountChanged // Event containing the contract specifics and raw log

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
func (it *CounterV2CountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterV2CountChanged)
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
		it.Event = new(CounterV2CountChanged)
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
func (it *CounterV2CountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterV2CountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterV2CountChanged represents a CountChanged event raised by the CounterV2 contract.
type CounterV2CountChanged struct {
	NewCount  *big.Int
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCountChanged is a free log retrieval operation binding the contract event 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82.
//
// Solidity: event CountChanged(uint256 newCount, address changedBy)
func (_CounterV2 *CounterV2Filterer) FilterCountChanged(opts *bind.FilterOpts) (*CounterV2CountChangedIterator, error) {

	logs, sub, err := _CounterV2.contract.FilterLogs(opts, "CountChanged")
	if err != nil {
		return nil, err
	}
	return &CounterV2CountChangedIterator{contract: _CounterV2.contract, event: "CountChanged", logs: logs, sub: sub}, nil
}

// WatchCountChanged is a free log subscription operation binding the contract event 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82.
//
// Solidity: event CountChanged(uint256 newCount, address changedBy)
func (_CounterV2 *CounterV2Filterer) WatchCountChanged(opts *bind.WatchOpts, sink chan<- *CounterV2CountChanged) (event.Subscription, error) {

	logs, sub, err := _CounterV2.contract.WatchLogs(opts, "CountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterV2CountChanged)
				if err := _CounterV2.contract.UnpackLog(event, "CountChanged", log); err != nil {
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
func (_CounterV2 *CounterV2Filterer) ParseCountChanged(log types.Log) (*CounterV2CountChanged, error) {
	event := new(CounterV2CountChanged)
	if err := _CounterV2.contract.UnpackLog(event, "CountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterV2CountResetIterator is returned from FilterCountReset and is used to iterate over the raw logs and unpacked data for CountReset events raised by the CounterV2 contract.
type CounterV2CountResetIterator struct {
	Event *CounterV2CountReset // Event containing the contract specifics and raw log

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
func (it *CounterV2CountResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterV2CountReset)
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
		it.Event = new(CounterV2CountReset)
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
func (it *CounterV2CountResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterV2CountResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterV2CountReset represents a CountReset event raised by the CounterV2 contract.
type CounterV2CountReset struct {
	ResetBy common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCountReset is a free log retrieval operation binding the contract event 0xa5ee6258204973c56c5a39c4ac31e61723f410d84f9e8117ba52b76b7cea990c.
//
// Solidity: event CountReset(address resetBy)
func (_CounterV2 *CounterV2Filterer) FilterCountReset(opts *bind.FilterOpts) (*CounterV2CountResetIterator, error) {

	logs, sub, err := _CounterV2.contract.FilterLogs(opts, "CountReset")
	if err != nil {
		return nil, err
	}
	return &CounterV2CountResetIterator{contract: _CounterV2.contract, event: "CountReset", logs: logs, sub: sub}, nil
}

// WatchCountReset is a free log subscription operation binding the contract event 0xa5ee6258204973c56c5a39c4ac31e61723f410d84f9e8117ba52b76b7cea990c.
//
// Solidity: event CountReset(address resetBy)
func (_CounterV2 *CounterV2Filterer) WatchCountReset(opts *bind.WatchOpts, sink chan<- *CounterV2CountReset) (event.Subscription, error) {

	logs, sub, err := _CounterV2.contract.WatchLogs(opts, "CountReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterV2CountReset)
				if err := _CounterV2.contract.UnpackLog(event, "CountReset", log); err != nil {
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

// ParseCountReset is a log parse operation binding the contract event 0xa5ee6258204973c56c5a39c4ac31e61723f410d84f9e8117ba52b76b7cea990c.
//
// Solidity: event CountReset(address resetBy)
func (_CounterV2 *CounterV2Filterer) ParseCountReset(log types.Log) (*CounterV2CountReset, error) {
	event := new(CounterV2CountReset)
	if err := _CounterV2.contract.UnpackLog(event, "CountReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CounterV2 contract.
type CounterV2InitializedIterator struct {
	Event *CounterV2Initialized // Event containing the contract specifics and raw log

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
func (it *CounterV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterV2Initialized)
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
		it.Event = new(CounterV2Initialized)
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
func (it *CounterV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterV2Initialized represents a Initialized event raised by the CounterV2 contract.
type CounterV2Initialized struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e6.
//
// Solidity: event Initialized(address owner)
func (_CounterV2 *CounterV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*CounterV2InitializedIterator, error) {

	logs, sub, err := _CounterV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CounterV2InitializedIterator{contract: _CounterV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x908408e307fc569b417f6cbec5d5a06f44a0a505ac0479b47d421a4b2fd6a1e6.
//
// Solidity: event Initialized(address owner)
func (_CounterV2 *CounterV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CounterV2Initialized) (event.Subscription, error) {

	logs, sub, err := _CounterV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterV2Initialized)
				if err := _CounterV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CounterV2 *CounterV2Filterer) ParseInitialized(log types.Log) (*CounterV2Initialized, error) {
	event := new(CounterV2Initialized)
	if err := _CounterV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
