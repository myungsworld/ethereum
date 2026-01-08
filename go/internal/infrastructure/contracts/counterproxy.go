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

// CounterProxyMetaData contains all meta data concerning the CounterProxy contract.
var CounterProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"logic_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldLogic\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newLogic\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLogic\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161051c38038061051c83398101604081905261002e916100ee565b6001600160a01b0381166100945760405162461bcd60e51b815260206004820152602360248201527f50726f78793a206c6f67696320616464726573732063616e6e6f74206265207a60448201526265726f60e81b606482015260840160405180910390fd5b600280546001600160a01b0383166001600160a01b0319918216811790925560038054909116331790556040515f907f5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7908290a35061011b565b5f602082840312156100fe575f5ffd5b81516001600160a01b0381168114610114575f5ffd5b9392505050565b6103f4806101285f395ff3fe608060405260043610610042575f3560e01c80633659cfe6146100735780635c60da1b146100925780638f283970146100c7578063f851a440146100e65761005e565b3661005e5760025461005c906001600160a01b0316610103565b005b60025461005c906001600160a01b0316610103565b34801561007e575f5ffd5b5061005c61008d366004610391565b610121565b34801561009d575f5ffd5b506002546001600160a01b03165b6040516001600160a01b03909116815260200160405180910390f35b3480156100d2575f5ffd5b5061005c6100e1366004610391565b610285565b3480156100f1575f5ffd5b506003546001600160a01b03166100ab565b365f5f375f5f365f845af43d5f5f3e80801561011d573d5ff35b3d5ffd5b6003546001600160a01b031633146101805760405162461bcd60e51b815260206004820152601d60248201527f50726f78793a206f6e6c792061646d696e2063616e207570677261646500000060448201526064015b60405180910390fd5b6001600160a01b0381166101d65760405162461bcd60e51b815260206004820152601f60248201527f50726f78793a206e6577206c6f6769632063616e6e6f74206265207a65726f006044820152606401610177565b6002546001600160a01b03908116908216036102345760405162461bcd60e51b815260206004820152601960248201527f50726f78793a2073616d65206c6f6769632061646472657373000000000000006044820152606401610177565b600280546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7905f90a35050565b6003546001600160a01b031633146102ea5760405162461bcd60e51b815260206004820152602260248201527f50726f78793a206f6e6c792061646d696e2063616e206368616e67652061646d60448201526134b760f11b6064820152608401610177565b6001600160a01b0381166103405760405162461bcd60e51b815260206004820152601f60248201527f50726f78793a206e65772061646d696e2063616e6e6f74206265207a65726f006044820152606401610177565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f905f90a35050565b5f602082840312156103a1575f5ffd5b81356001600160a01b03811681146103b7575f5ffd5b939250505056fea26469706673582212200b80371c5aac30b7382933ba7f9aa4ca039a9458a9afd5c1b1ba8ea3009043bc64736f6c63430008210033",
}

// CounterProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterProxyMetaData.ABI instead.
var CounterProxyABI = CounterProxyMetaData.ABI

// CounterProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterProxyMetaData.Bin instead.
var CounterProxyBin = CounterProxyMetaData.Bin

// DeployCounterProxy deploys a new Ethereum contract, binding an instance of CounterProxy to it.
func DeployCounterProxy(auth *bind.TransactOpts, backend bind.ContractBackend, logic_ common.Address) (common.Address, *types.Transaction, *CounterProxy, error) {
	parsed, err := CounterProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterProxyBin), backend, logic_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CounterProxy{CounterProxyCaller: CounterProxyCaller{contract: contract}, CounterProxyTransactor: CounterProxyTransactor{contract: contract}, CounterProxyFilterer: CounterProxyFilterer{contract: contract}}, nil
}

// CounterProxy is an auto generated Go binding around an Ethereum contract.
type CounterProxy struct {
	CounterProxyCaller     // Read-only binding to the contract
	CounterProxyTransactor // Write-only binding to the contract
	CounterProxyFilterer   // Log filterer for contract events
}

// CounterProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterProxySession struct {
	Contract     *CounterProxy     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterProxyCallerSession struct {
	Contract *CounterProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CounterProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterProxyTransactorSession struct {
	Contract     *CounterProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CounterProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterProxyRaw struct {
	Contract *CounterProxy // Generic contract binding to access the raw methods on
}

// CounterProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterProxyCallerRaw struct {
	Contract *CounterProxyCaller // Generic read-only contract binding to access the raw methods on
}

// CounterProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterProxyTransactorRaw struct {
	Contract *CounterProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounterProxy creates a new instance of CounterProxy, bound to a specific deployed contract.
func NewCounterProxy(address common.Address, backend bind.ContractBackend) (*CounterProxy, error) {
	contract, err := bindCounterProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CounterProxy{CounterProxyCaller: CounterProxyCaller{contract: contract}, CounterProxyTransactor: CounterProxyTransactor{contract: contract}, CounterProxyFilterer: CounterProxyFilterer{contract: contract}}, nil
}

// NewCounterProxyCaller creates a new read-only instance of CounterProxy, bound to a specific deployed contract.
func NewCounterProxyCaller(address common.Address, caller bind.ContractCaller) (*CounterProxyCaller, error) {
	contract, err := bindCounterProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterProxyCaller{contract: contract}, nil
}

// NewCounterProxyTransactor creates a new write-only instance of CounterProxy, bound to a specific deployed contract.
func NewCounterProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterProxyTransactor, error) {
	contract, err := bindCounterProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterProxyTransactor{contract: contract}, nil
}

// NewCounterProxyFilterer creates a new log filterer instance of CounterProxy, bound to a specific deployed contract.
func NewCounterProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterProxyFilterer, error) {
	contract, err := bindCounterProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterProxyFilterer{contract: contract}, nil
}

// bindCounterProxy binds a generic wrapper to an already deployed contract.
func bindCounterProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterProxy *CounterProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterProxy.Contract.CounterProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterProxy *CounterProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterProxy.Contract.CounterProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterProxy *CounterProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterProxy.Contract.CounterProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterProxy *CounterProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterProxy *CounterProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterProxy *CounterProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CounterProxy *CounterProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CounterProxy.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CounterProxy *CounterProxySession) Admin() (common.Address, error) {
	return _CounterProxy.Contract.Admin(&_CounterProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CounterProxy *CounterProxyCallerSession) Admin() (common.Address, error) {
	return _CounterProxy.Contract.Admin(&_CounterProxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_CounterProxy *CounterProxyCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CounterProxy.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_CounterProxy *CounterProxySession) Implementation() (common.Address, error) {
	return _CounterProxy.Contract.Implementation(&_CounterProxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_CounterProxy *CounterProxyCallerSession) Implementation() (common.Address, error) {
	return _CounterProxy.Contract.Implementation(&_CounterProxy.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_CounterProxy *CounterProxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _CounterProxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_CounterProxy *CounterProxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _CounterProxy.Contract.ChangeAdmin(&_CounterProxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_CounterProxy *CounterProxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _CounterProxy.Contract.ChangeAdmin(&_CounterProxy.TransactOpts, newAdmin)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newLogic) returns()
func (_CounterProxy *CounterProxyTransactor) UpgradeTo(opts *bind.TransactOpts, newLogic common.Address) (*types.Transaction, error) {
	return _CounterProxy.contract.Transact(opts, "upgradeTo", newLogic)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newLogic) returns()
func (_CounterProxy *CounterProxySession) UpgradeTo(newLogic common.Address) (*types.Transaction, error) {
	return _CounterProxy.Contract.UpgradeTo(&_CounterProxy.TransactOpts, newLogic)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newLogic) returns()
func (_CounterProxy *CounterProxyTransactorSession) UpgradeTo(newLogic common.Address) (*types.Transaction, error) {
	return _CounterProxy.Contract.UpgradeTo(&_CounterProxy.TransactOpts, newLogic)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CounterProxy *CounterProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CounterProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CounterProxy *CounterProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CounterProxy.Contract.Fallback(&_CounterProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CounterProxy *CounterProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CounterProxy.Contract.Fallback(&_CounterProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CounterProxy *CounterProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CounterProxy *CounterProxySession) Receive() (*types.Transaction, error) {
	return _CounterProxy.Contract.Receive(&_CounterProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CounterProxy *CounterProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _CounterProxy.Contract.Receive(&_CounterProxy.TransactOpts)
}

// CounterProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the CounterProxy contract.
type CounterProxyAdminChangedIterator struct {
	Event *CounterProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *CounterProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterProxyAdminChanged)
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
		it.Event = new(CounterProxyAdminChanged)
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
func (it *CounterProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterProxyAdminChanged represents a AdminChanged event raised by the CounterProxy contract.
type CounterProxyAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed oldAdmin, address indexed newAdmin)
func (_CounterProxy *CounterProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*CounterProxyAdminChangedIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CounterProxy.contract.FilterLogs(opts, "AdminChanged", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &CounterProxyAdminChangedIterator{contract: _CounterProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed oldAdmin, address indexed newAdmin)
func (_CounterProxy *CounterProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *CounterProxyAdminChanged, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CounterProxy.contract.WatchLogs(opts, "AdminChanged", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterProxyAdminChanged)
				if err := _CounterProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed oldAdmin, address indexed newAdmin)
func (_CounterProxy *CounterProxyFilterer) ParseAdminChanged(log types.Log) (*CounterProxyAdminChanged, error) {
	event := new(CounterProxyAdminChanged)
	if err := _CounterProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CounterProxy contract.
type CounterProxyUpgradedIterator struct {
	Event *CounterProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *CounterProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterProxyUpgraded)
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
		it.Event = new(CounterProxyUpgraded)
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
func (it *CounterProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterProxyUpgraded represents a Upgraded event raised by the CounterProxy contract.
type CounterProxyUpgraded struct {
	OldLogic common.Address
	NewLogic common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed oldLogic, address indexed newLogic)
func (_CounterProxy *CounterProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, oldLogic []common.Address, newLogic []common.Address) (*CounterProxyUpgradedIterator, error) {

	var oldLogicRule []interface{}
	for _, oldLogicItem := range oldLogic {
		oldLogicRule = append(oldLogicRule, oldLogicItem)
	}
	var newLogicRule []interface{}
	for _, newLogicItem := range newLogic {
		newLogicRule = append(newLogicRule, newLogicItem)
	}

	logs, sub, err := _CounterProxy.contract.FilterLogs(opts, "Upgraded", oldLogicRule, newLogicRule)
	if err != nil {
		return nil, err
	}
	return &CounterProxyUpgradedIterator{contract: _CounterProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed oldLogic, address indexed newLogic)
func (_CounterProxy *CounterProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CounterProxyUpgraded, oldLogic []common.Address, newLogic []common.Address) (event.Subscription, error) {

	var oldLogicRule []interface{}
	for _, oldLogicItem := range oldLogic {
		oldLogicRule = append(oldLogicRule, oldLogicItem)
	}
	var newLogicRule []interface{}
	for _, newLogicItem := range newLogic {
		newLogicRule = append(newLogicRule, newLogicItem)
	}

	logs, sub, err := _CounterProxy.contract.WatchLogs(opts, "Upgraded", oldLogicRule, newLogicRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterProxyUpgraded)
				if err := _CounterProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed oldLogic, address indexed newLogic)
func (_CounterProxy *CounterProxyFilterer) ParseUpgraded(log types.Log) (*CounterProxyUpgraded, error) {
	event := new(CounterProxyUpgraded)
	if err := _CounterProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
