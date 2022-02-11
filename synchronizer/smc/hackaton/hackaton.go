// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hackaton

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

// HackatonMetaData contains all meta data concerning the Hackaton contract.
var HackatonMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"information\",\"type\":\"string\"}],\"name\":\"Data\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"sendEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"proofB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofC\",\"type\":\"uint256[2]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6109578061010d6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806327e0b8991461005c5780633d1aa90e14610078578063715018a6146100945780638da5cb5b1461009e578063f2fde38b146100bc575b600080fd5b61007660048036038101906100719190610528565b6100d8565b005b610092600480360381019061008d91906105f0565b610112565b005b61009c610159565b005b6100a66101e1565b6040516100b39190610699565b60405180910390f35b6100d660048036038101906100d191906106e0565b61020a565b005b7fe8e349d9d60379b7e0a717f216d3d9efaf00b5dce8b634a63f24a5d76251059f816040516101079190610795565b60405180910390a150565b6001610153576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014a90610803565b60405180910390fd5b50505050565b610161610302565b73ffffffffffffffffffffffffffffffffffffffff1661017f6101e1565b73ffffffffffffffffffffffffffffffffffffffff16146101d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101cc9061086f565b60405180910390fd5b6101df600061030a565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610212610302565b73ffffffffffffffffffffffffffffffffffffffff166102306101e1565b73ffffffffffffffffffffffffffffffffffffffff1614610286576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027d9061086f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156102f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ed90610901565b60405180910390fd5b6102ff8161030a565b50565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610435826103ec565b810181811067ffffffffffffffff82111715610454576104536103fd565b5b80604052505050565b60006104676103ce565b9050610473828261042c565b919050565b600067ffffffffffffffff821115610493576104926103fd565b5b61049c826103ec565b9050602081019050919050565b82818337600083830152505050565b60006104cb6104c684610478565b61045d565b9050828152602081018484840111156104e7576104e66103e7565b5b6104f28482856104a9565b509392505050565b600082601f83011261050f5761050e6103e2565b5b813561051f8482602086016104b8565b91505092915050565b60006020828403121561053e5761053d6103d8565b5b600082013567ffffffffffffffff81111561055c5761055b6103dd565b5b610568848285016104fa565b91505092915050565b6000819050919050565b61058481610571565b811461058f57600080fd5b50565b6000813590506105a18161057b565b92915050565b600080fd5b6000819050826020600202820111156105c8576105c76105a7565b5b92915050565b6000819050826040600202820111156105ea576105e96105a7565b5b92915050565b600080600080610120858703121561060b5761060a6103d8565b5b600061061987828801610592565b945050602061062a878288016105ac565b935050606061063b878288016105ce565b92505060e061064c878288016105ac565b91505092959194509250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061068382610658565b9050919050565b61069381610678565b82525050565b60006020820190506106ae600083018461068a565b92915050565b6106bd81610678565b81146106c857600080fd5b50565b6000813590506106da816106b4565b92915050565b6000602082840312156106f6576106f56103d8565b5b6000610704848285016106cb565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561074757808201518184015260208101905061072c565b83811115610756576000848401525b50505050565b60006107678261070d565b6107718185610718565b9350610781818560208601610729565b61078a816103ec565b840191505092915050565b600060208201905081810360008301526107af818461075c565b905092915050565b7f566572696669636174696f6e3a20494e56414c49445f50524f4f460000000000600082015250565b60006107ed601b83610718565b91506107f8826107b7565b602082019050919050565b6000602082019050818103600083015261081c816107e0565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610859602083610718565b915061086482610823565b602082019050919050565b600060208201905081810360008301526108888161084c565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006108eb602683610718565b91506108f68261088f565b604082019050919050565b6000602082019050818103600083015261091a816108de565b905091905056fea2646970667358221220b110a8a7b738e798a24b0942f6d3dcb9103c5235d9c813618e6bfecd20a8037764736f6c63430008090033",
}

// HackatonABI is the input ABI used to generate the binding from.
// Deprecated: Use HackatonMetaData.ABI instead.
var HackatonABI = HackatonMetaData.ABI

// HackatonBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HackatonMetaData.Bin instead.
var HackatonBin = HackatonMetaData.Bin

// DeployHackaton deploys a new Ethereum contract, binding an instance of Hackaton to it.
func DeployHackaton(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hackaton, error) {
	parsed, err := HackatonMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HackatonBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hackaton{HackatonCaller: HackatonCaller{contract: contract}, HackatonTransactor: HackatonTransactor{contract: contract}, HackatonFilterer: HackatonFilterer{contract: contract}}, nil
}

// Hackaton is an auto generated Go binding around an Ethereum contract.
type Hackaton struct {
	HackatonCaller     // Read-only binding to the contract
	HackatonTransactor // Write-only binding to the contract
	HackatonFilterer   // Log filterer for contract events
}

// HackatonCaller is an auto generated read-only Go binding around an Ethereum contract.
type HackatonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HackatonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HackatonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HackatonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HackatonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HackatonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HackatonSession struct {
	Contract     *Hackaton         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HackatonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HackatonCallerSession struct {
	Contract *HackatonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HackatonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HackatonTransactorSession struct {
	Contract     *HackatonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HackatonRaw is an auto generated low-level Go binding around an Ethereum contract.
type HackatonRaw struct {
	Contract *Hackaton // Generic contract binding to access the raw methods on
}

// HackatonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HackatonCallerRaw struct {
	Contract *HackatonCaller // Generic read-only contract binding to access the raw methods on
}

// HackatonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HackatonTransactorRaw struct {
	Contract *HackatonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHackaton creates a new instance of Hackaton, bound to a specific deployed contract.
func NewHackaton(address common.Address, backend bind.ContractBackend) (*Hackaton, error) {
	contract, err := bindHackaton(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hackaton{HackatonCaller: HackatonCaller{contract: contract}, HackatonTransactor: HackatonTransactor{contract: contract}, HackatonFilterer: HackatonFilterer{contract: contract}}, nil
}

// NewHackatonCaller creates a new read-only instance of Hackaton, bound to a specific deployed contract.
func NewHackatonCaller(address common.Address, caller bind.ContractCaller) (*HackatonCaller, error) {
	contract, err := bindHackaton(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HackatonCaller{contract: contract}, nil
}

// NewHackatonTransactor creates a new write-only instance of Hackaton, bound to a specific deployed contract.
func NewHackatonTransactor(address common.Address, transactor bind.ContractTransactor) (*HackatonTransactor, error) {
	contract, err := bindHackaton(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HackatonTransactor{contract: contract}, nil
}

// NewHackatonFilterer creates a new log filterer instance of Hackaton, bound to a specific deployed contract.
func NewHackatonFilterer(address common.Address, filterer bind.ContractFilterer) (*HackatonFilterer, error) {
	contract, err := bindHackaton(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HackatonFilterer{contract: contract}, nil
}

// bindHackaton binds a generic wrapper to an already deployed contract.
func bindHackaton(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HackatonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hackaton *HackatonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hackaton.Contract.HackatonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hackaton *HackatonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hackaton.Contract.HackatonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hackaton *HackatonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hackaton.Contract.HackatonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hackaton *HackatonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hackaton.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hackaton *HackatonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hackaton.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hackaton *HackatonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hackaton.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hackaton *HackatonCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hackaton.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hackaton *HackatonSession) Owner() (common.Address, error) {
	return _Hackaton.Contract.Owner(&_Hackaton.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hackaton *HackatonCallerSession) Owner() (common.Address, error) {
	return _Hackaton.Contract.Owner(&_Hackaton.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hackaton *HackatonTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hackaton.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hackaton *HackatonSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hackaton.Contract.RenounceOwnership(&_Hackaton.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hackaton *HackatonTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hackaton.Contract.RenounceOwnership(&_Hackaton.TransactOpts)
}

// SendEvent is a paid mutator transaction binding the contract method 0x27e0b899.
//
// Solidity: function sendEvent(string data) returns()
func (_Hackaton *HackatonTransactor) SendEvent(opts *bind.TransactOpts, data string) (*types.Transaction, error) {
	return _Hackaton.contract.Transact(opts, "sendEvent", data)
}

// SendEvent is a paid mutator transaction binding the contract method 0x27e0b899.
//
// Solidity: function sendEvent(string data) returns()
func (_Hackaton *HackatonSession) SendEvent(data string) (*types.Transaction, error) {
	return _Hackaton.Contract.SendEvent(&_Hackaton.TransactOpts, data)
}

// SendEvent is a paid mutator transaction binding the contract method 0x27e0b899.
//
// Solidity: function sendEvent(string data) returns()
func (_Hackaton *HackatonTransactorSession) SendEvent(data string) (*types.Transaction, error) {
	return _Hackaton.Contract.SendEvent(&_Hackaton.TransactOpts, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hackaton *HackatonTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hackaton.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hackaton *HackatonSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hackaton.Contract.TransferOwnership(&_Hackaton.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hackaton *HackatonTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hackaton.Contract.TransferOwnership(&_Hackaton.TransactOpts, newOwner)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x3d1aa90e.
//
// Solidity: function verifyProof(bytes32 hash, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Hackaton *HackatonTransactor) VerifyProof(opts *bind.TransactOpts, hash [32]byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Hackaton.contract.Transact(opts, "verifyProof", hash, proofA, proofB, proofC)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x3d1aa90e.
//
// Solidity: function verifyProof(bytes32 hash, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Hackaton *HackatonSession) VerifyProof(hash [32]byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Hackaton.Contract.VerifyProof(&_Hackaton.TransactOpts, hash, proofA, proofB, proofC)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x3d1aa90e.
//
// Solidity: function verifyProof(bytes32 hash, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Hackaton *HackatonTransactorSession) VerifyProof(hash [32]byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Hackaton.Contract.VerifyProof(&_Hackaton.TransactOpts, hash, proofA, proofB, proofC)
}

// HackatonDataIterator is returned from FilterData and is used to iterate over the raw logs and unpacked data for Data events raised by the Hackaton contract.
type HackatonDataIterator struct {
	Event *HackatonData // Event containing the contract specifics and raw log

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
func (it *HackatonDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HackatonData)
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
		it.Event = new(HackatonData)
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
func (it *HackatonDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HackatonDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HackatonData represents a Data event raised by the Hackaton contract.
type HackatonData struct {
	Information string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterData is a free log retrieval operation binding the contract event 0xe8e349d9d60379b7e0a717f216d3d9efaf00b5dce8b634a63f24a5d76251059f.
//
// Solidity: event Data(string information)
func (_Hackaton *HackatonFilterer) FilterData(opts *bind.FilterOpts) (*HackatonDataIterator, error) {

	logs, sub, err := _Hackaton.contract.FilterLogs(opts, "Data")
	if err != nil {
		return nil, err
	}
	return &HackatonDataIterator{contract: _Hackaton.contract, event: "Data", logs: logs, sub: sub}, nil
}

// WatchData is a free log subscription operation binding the contract event 0xe8e349d9d60379b7e0a717f216d3d9efaf00b5dce8b634a63f24a5d76251059f.
//
// Solidity: event Data(string information)
func (_Hackaton *HackatonFilterer) WatchData(opts *bind.WatchOpts, sink chan<- *HackatonData) (event.Subscription, error) {

	logs, sub, err := _Hackaton.contract.WatchLogs(opts, "Data")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HackatonData)
				if err := _Hackaton.contract.UnpackLog(event, "Data", log); err != nil {
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

// ParseData is a log parse operation binding the contract event 0xe8e349d9d60379b7e0a717f216d3d9efaf00b5dce8b634a63f24a5d76251059f.
//
// Solidity: event Data(string information)
func (_Hackaton *HackatonFilterer) ParseData(log types.Log) (*HackatonData, error) {
	event := new(HackatonData)
	if err := _Hackaton.contract.UnpackLog(event, "Data", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HackatonOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hackaton contract.
type HackatonOwnershipTransferredIterator struct {
	Event *HackatonOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HackatonOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HackatonOwnershipTransferred)
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
		it.Event = new(HackatonOwnershipTransferred)
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
func (it *HackatonOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HackatonOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HackatonOwnershipTransferred represents a OwnershipTransferred event raised by the Hackaton contract.
type HackatonOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hackaton *HackatonFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HackatonOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hackaton.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HackatonOwnershipTransferredIterator{contract: _Hackaton.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hackaton *HackatonFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HackatonOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hackaton.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HackatonOwnershipTransferred)
				if err := _Hackaton.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Hackaton *HackatonFilterer) ParseOwnershipTransferred(log types.Log) (*HackatonOwnershipTransferred, error) {
	event := new(HackatonOwnershipTransferred)
	if err := _Hackaton.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
