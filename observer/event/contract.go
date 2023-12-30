// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package event

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
)

// TestEventMetaData contains all meta data concerning the TestEvent contract.
var TestEventMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"number\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"toadd\",\"type\":\"uint8\"}],\"name\":\"Launch\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"number\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"toadd\",\"type\":\"uint8\"}],\"name\":\"Trigger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TestEventABI is the input ABI used to generate the binding from.
// Deprecated: Use TestEventMetaData.ABI instead.
var TestEventABI = TestEventMetaData.ABI

// TestEvent is an auto generated Go binding around an Ethereum contract.
type TestEvent struct {
	TestEventCaller     // Read-only binding to the contract
	TestEventTransactor // Write-only binding to the contract
	TestEventFilterer   // Log filterer for contract events
}

// TestEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestEventSession struct {
	Contract     *TestEvent        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestEventCallerSession struct {
	Contract *TestEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestEventTransactorSession struct {
	Contract     *TestEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestEventRaw struct {
	Contract *TestEvent // Generic contract binding to access the raw methods on
}

// TestEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestEventCallerRaw struct {
	Contract *TestEventCaller // Generic read-only contract binding to access the raw methods on
}

// TestEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestEventTransactorRaw struct {
	Contract *TestEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestEvent creates a new instance of TestEvent, bound to a specific deployed contract.
func NewTestEvent(address common.Address, backend bind.ContractBackend) (*TestEvent, error) {
	contract, err := bindTestEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestEvent{TestEventCaller: TestEventCaller{contract: contract}, TestEventTransactor: TestEventTransactor{contract: contract}, TestEventFilterer: TestEventFilterer{contract: contract}}, nil
}

// NewTestEventCaller creates a new read-only instance of TestEvent, bound to a specific deployed contract.
func NewTestEventCaller(address common.Address, caller bind.ContractCaller) (*TestEventCaller, error) {
	contract, err := bindTestEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestEventCaller{contract: contract}, nil
}

// NewTestEventTransactor creates a new write-only instance of TestEvent, bound to a specific deployed contract.
func NewTestEventTransactor(address common.Address, transactor bind.ContractTransactor) (*TestEventTransactor, error) {
	contract, err := bindTestEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestEventTransactor{contract: contract}, nil
}

// NewTestEventFilterer creates a new log filterer instance of TestEvent, bound to a specific deployed contract.
func NewTestEventFilterer(address common.Address, filterer bind.ContractFilterer) (*TestEventFilterer, error) {
	contract, err := bindTestEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestEventFilterer{contract: contract}, nil
}

// bindTestEvent binds a generic wrapper to an already deployed contract.
func bindTestEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestEventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEvent *TestEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestEvent.Contract.TestEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEvent *TestEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEvent.Contract.TestEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEvent *TestEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEvent.Contract.TestEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEvent *TestEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEvent *TestEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEvent *TestEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEvent.Contract.contract.Transact(opts, method, params...)
}

// Trigger is a paid mutator transaction binding the contract method 0x80455628.
//
// Solidity: function Trigger(uint8 number, uint8 toadd) returns()
func (_TestEvent *TestEventTransactor) Trigger(opts *bind.TransactOpts, number uint8, toadd uint8) (*types.Transaction, error) {
	return _TestEvent.contract.Transact(opts, "Trigger", number, toadd)
}

// Trigger is a paid mutator transaction binding the contract method 0x80455628.
//
// Solidity: function Trigger(uint8 number, uint8 toadd) returns()
func (_TestEvent *TestEventSession) Trigger(number uint8, toadd uint8) (*types.Transaction, error) {
	return _TestEvent.Contract.Trigger(&_TestEvent.TransactOpts, number, toadd)
}

// Trigger is a paid mutator transaction binding the contract method 0x80455628.
//
// Solidity: function Trigger(uint8 number, uint8 toadd) returns()
func (_TestEvent *TestEventTransactorSession) Trigger(number uint8, toadd uint8) (*types.Transaction, error) {
	return _TestEvent.Contract.Trigger(&_TestEvent.TransactOpts, number, toadd)
}

// TestEventLaunchIterator is returned from FilterLaunch and is used to iterate over the raw logs and unpacked data for Launch events raised by the TestEvent contract.
type TestEventLaunchIterator struct {
	Event *TestEventLaunch // Event containing the contract specifics and raw log

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
func (it *TestEventLaunchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEventLaunch)
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
		it.Event = new(TestEventLaunch)
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
func (it *TestEventLaunchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEventLaunchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEventLaunch represents a Launch event raised by the TestEvent contract.
type TestEventLaunch struct {
	Number uint8
	Toadd  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLaunch is a free log retrieval operation binding the contract event 0xd80d05dcc56133f514e98b094d950c262b8051c00409d936a6e371d3388093ce.
//
// Solidity: event Launch(uint8 indexed number, uint8 indexed toadd)
func (_TestEvent *TestEventFilterer) FilterLaunch(opts *bind.FilterOpts, number []uint8, toadd []uint8) (*TestEventLaunchIterator, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var toaddRule []interface{}
	for _, toaddItem := range toadd {
		toaddRule = append(toaddRule, toaddItem)
	}

	logs, sub, err := _TestEvent.contract.FilterLogs(opts, "Launch", numberRule, toaddRule)
	if err != nil {
		return nil, err
	}
	return &TestEventLaunchIterator{contract: _TestEvent.contract, event: "Launch", logs: logs, sub: sub}, nil
}

// WatchLaunch is a free log subscription operation binding the contract event 0xd80d05dcc56133f514e98b094d950c262b8051c00409d936a6e371d3388093ce.
//
// Solidity: event Launch(uint8 indexed number, uint8 indexed toadd)
func (_TestEvent *TestEventFilterer) WatchLaunch(opts *bind.WatchOpts, sink chan<- *TestEventLaunch, number []uint8, toadd []uint8) (event.Subscription, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var toaddRule []interface{}
	for _, toaddItem := range toadd {
		toaddRule = append(toaddRule, toaddItem)
	}

	logs, sub, err := _TestEvent.contract.WatchLogs(opts, "Launch", numberRule, toaddRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEventLaunch)
				if err := _TestEvent.contract.UnpackLog(event, "Launch", log); err != nil {
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

// ParseLaunch is a log parse operation binding the contract event 0xd80d05dcc56133f514e98b094d950c262b8051c00409d936a6e371d3388093ce.
//
// Solidity: event Launch(uint8 indexed number, uint8 indexed toadd)
func (_TestEvent *TestEventFilterer) ParseLaunch(log types.Log) (*TestEventLaunch, error) {
	event := new(TestEventLaunch)
	if err := _TestEvent.contract.UnpackLog(event, "Launch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
