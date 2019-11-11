// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeBank

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeBankABI is the input ABI used to generate the binding from.
const BridgeBankABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"viewCosmosDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeTokenCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"bridgeTokenWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createNewBridgeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"bytes\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cosmosBridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cosmosSender\",\"type\":\"bytes\"},{\"name\":\"_intendedRecipient\",\"type\":\"address\"},{\"name\":\"_bridgeTokenAddress\",\"type\":\"address\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintBridgeTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getCosmosDepositStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operatorAddress\",\"type\":\"address\"},{\"name\":\"_oracleAddress\",\"type\":\"address\"},{\"name\":\"_cosmosBridgeAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"LogLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"LogUnlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"LogNewBridgeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"LogBridgeTokenMint\",\"type\":\"event\"}]"

// BridgeBank is an auto generated Go binding around an Ethereum contract.
type BridgeBank struct {
	BridgeBankCaller     // Read-only binding to the contract
	BridgeBankTransactor // Write-only binding to the contract
	BridgeBankFilterer   // Log filterer for contract events
}

// BridgeBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBankSession struct {
	Contract     *BridgeBank       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBankCallerSession struct {
	Contract *BridgeBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBankTransactorSession struct {
	Contract     *BridgeBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBankRaw struct {
	Contract *BridgeBank // Generic contract binding to access the raw methods on
}

// BridgeBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBankCallerRaw struct {
	Contract *BridgeBankCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBankTransactorRaw struct {
	Contract *BridgeBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBank creates a new instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBank(address common.Address, backend bind.ContractBackend) (*BridgeBank, error) {
	contract, err := bindBridgeBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBank{BridgeBankCaller: BridgeBankCaller{contract: contract}, BridgeBankTransactor: BridgeBankTransactor{contract: contract}, BridgeBankFilterer: BridgeBankFilterer{contract: contract}}, nil
}

// NewBridgeBankCaller creates a new read-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankCaller(address common.Address, caller bind.ContractCaller) (*BridgeBankCaller, error) {
	contract, err := bindBridgeBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankCaller{contract: contract}, nil
}

// NewBridgeBankTransactor creates a new write-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBankTransactor, error) {
	contract, err := bindBridgeBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankTransactor{contract: contract}, nil
}

// NewBridgeBankFilterer creates a new log filterer instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBankFilterer, error) {
	contract, err := bindBridgeBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBankFilterer{contract: contract}, nil
}

// bindBridgeBank binds a generic wrapper to an already deployed contract.
func bindBridgeBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.BridgeBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transact(opts, method, params...)
}

// BridgeTokenCount is a free data retrieval call binding the contract method 0x328470ab.
//
// Solidity: function bridgeTokenCount() constant returns(uint256)
func (_BridgeBank *BridgeBankCaller) BridgeTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "bridgeTokenCount")
	return *ret0, err
}

// BridgeTokenCount is a free data retrieval call binding the contract method 0x328470ab.
//
// Solidity: function bridgeTokenCount() constant returns(uint256)
func (_BridgeBank *BridgeBankSession) BridgeTokenCount() (*big.Int, error) {
	return _BridgeBank.Contract.BridgeTokenCount(&_BridgeBank.CallOpts)
}

// BridgeTokenCount is a free data retrieval call binding the contract method 0x328470ab.
//
// Solidity: function bridgeTokenCount() constant returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) BridgeTokenCount() (*big.Int, error) {
	return _BridgeBank.Contract.BridgeTokenCount(&_BridgeBank.CallOpts)
}

// BridgeTokenWhitelist is a free data retrieval call binding the contract method 0x3f4d5681.
//
// Solidity: function bridgeTokenWhitelist(address ) constant returns(bool)
func (_BridgeBank *BridgeBankCaller) BridgeTokenWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "bridgeTokenWhitelist", arg0)
	return *ret0, err
}

// BridgeTokenWhitelist is a free data retrieval call binding the contract method 0x3f4d5681.
//
// Solidity: function bridgeTokenWhitelist(address ) constant returns(bool)
func (_BridgeBank *BridgeBankSession) BridgeTokenWhitelist(arg0 common.Address) (bool, error) {
	return _BridgeBank.Contract.BridgeTokenWhitelist(&_BridgeBank.CallOpts, arg0)
}

// BridgeTokenWhitelist is a free data retrieval call binding the contract method 0x3f4d5681.
//
// Solidity: function bridgeTokenWhitelist(address ) constant returns(bool)
func (_BridgeBank *BridgeBankCallerSession) BridgeTokenWhitelist(arg0 common.Address) (bool, error) {
	return _BridgeBank.Contract.BridgeTokenWhitelist(&_BridgeBank.CallOpts, arg0)
}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() constant returns(address)
func (_BridgeBank *BridgeBankCaller) CosmosBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "cosmosBridge")
	return *ret0, err
}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() constant returns(address)
func (_BridgeBank *BridgeBankSession) CosmosBridge() (common.Address, error) {
	return _BridgeBank.Contract.CosmosBridge(&_BridgeBank.CallOpts)
}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() constant returns(address)
func (_BridgeBank *BridgeBankCallerSession) CosmosBridge() (common.Address, error) {
	return _BridgeBank.Contract.CosmosBridge(&_BridgeBank.CallOpts)
}

// GetCosmosDepositStatus is a free data retrieval call binding the contract method 0xe06064f4.
//
// Solidity: function getCosmosDepositStatus(bytes32 _id) constant returns(bool)
func (_BridgeBank *BridgeBankCaller) GetCosmosDepositStatus(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "getCosmosDepositStatus", _id)
	return *ret0, err
}

// GetCosmosDepositStatus is a free data retrieval call binding the contract method 0xe06064f4.
//
// Solidity: function getCosmosDepositStatus(bytes32 _id) constant returns(bool)
func (_BridgeBank *BridgeBankSession) GetCosmosDepositStatus(_id [32]byte) (bool, error) {
	return _BridgeBank.Contract.GetCosmosDepositStatus(&_BridgeBank.CallOpts, _id)
}

// GetCosmosDepositStatus is a free data retrieval call binding the contract method 0xe06064f4.
//
// Solidity: function getCosmosDepositStatus(bytes32 _id) constant returns(bool)
func (_BridgeBank *BridgeBankCallerSession) GetCosmosDepositStatus(_id [32]byte) (bool, error) {
	return _BridgeBank.Contract.GetCosmosDepositStatus(&_BridgeBank.CallOpts, _id)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() constant returns(uint256)
func (_BridgeBank *BridgeBankCaller) LockNonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "lockNonce")
	return *ret0, err
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() constant returns(uint256)
func (_BridgeBank *BridgeBankSession) LockNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockNonce(&_BridgeBank.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() constant returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) LockNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockNonce(&_BridgeBank.CallOpts)
}

// LockedFunds is a free data retrieval call binding the contract method 0xb86247d7.
//
// Solidity: function lockedFunds(address ) constant returns(uint256)
func (_BridgeBank *BridgeBankCaller) LockedFunds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "lockedFunds", arg0)
	return *ret0, err
}

// LockedFunds is a free data retrieval call binding the contract method 0xb86247d7.
//
// Solidity: function lockedFunds(address ) constant returns(uint256)
func (_BridgeBank *BridgeBankSession) LockedFunds(arg0 common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.LockedFunds(&_BridgeBank.CallOpts, arg0)
}

// LockedFunds is a free data retrieval call binding the contract method 0xb86247d7.
//
// Solidity: function lockedFunds(address ) constant returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) LockedFunds(arg0 common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.LockedFunds(&_BridgeBank.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_BridgeBank *BridgeBankCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_BridgeBank *BridgeBankSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_BridgeBank *BridgeBankCallerSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_BridgeBank *BridgeBankCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_BridgeBank *BridgeBankSession) Oracle() (common.Address, error) {
	return _BridgeBank.Contract.Oracle(&_BridgeBank.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_BridgeBank *BridgeBankCallerSession) Oracle() (common.Address, error) {
	return _BridgeBank.Contract.Oracle(&_BridgeBank.CallOpts)
}

// ViewCosmosDeposit is a free data retrieval call binding the contract method 0x2852a414.
//
// Solidity: function viewCosmosDeposit(bytes32 _id) constant returns(bytes, address, address, uint256)
func (_BridgeBank *BridgeBankCaller) ViewCosmosDeposit(opts *bind.CallOpts, _id [32]byte) ([]byte, common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new([]byte)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _BridgeBank.contract.Call(opts, out, "viewCosmosDeposit", _id)
	return *ret0, *ret1, *ret2, *ret3, err
}

// ViewCosmosDeposit is a free data retrieval call binding the contract method 0x2852a414.
//
// Solidity: function viewCosmosDeposit(bytes32 _id) constant returns(bytes, address, address, uint256)
func (_BridgeBank *BridgeBankSession) ViewCosmosDeposit(_id [32]byte) ([]byte, common.Address, common.Address, *big.Int, error) {
	return _BridgeBank.Contract.ViewCosmosDeposit(&_BridgeBank.CallOpts, _id)
}

// ViewCosmosDeposit is a free data retrieval call binding the contract method 0x2852a414.
//
// Solidity: function viewCosmosDeposit(bytes32 _id) constant returns(bytes, address, address, uint256)
func (_BridgeBank *BridgeBankCallerSession) ViewCosmosDeposit(_id [32]byte) ([]byte, common.Address, common.Address, *big.Int, error) {
	return _BridgeBank.Contract.ViewCosmosDeposit(&_BridgeBank.CallOpts, _id)
}

// CreateNewBridgeToken is a paid mutator transaction binding the contract method 0x50b06e4d.
//
// Solidity: function createNewBridgeToken(string _symbol) returns(address)
func (_BridgeBank *BridgeBankTransactor) CreateNewBridgeToken(opts *bind.TransactOpts, _symbol string) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "createNewBridgeToken", _symbol)
}

// CreateNewBridgeToken is a paid mutator transaction binding the contract method 0x50b06e4d.
//
// Solidity: function createNewBridgeToken(string _symbol) returns(address)
func (_BridgeBank *BridgeBankSession) CreateNewBridgeToken(_symbol string) (*types.Transaction, error) {
	return _BridgeBank.Contract.CreateNewBridgeToken(&_BridgeBank.TransactOpts, _symbol)
}

// CreateNewBridgeToken is a paid mutator transaction binding the contract method 0x50b06e4d.
//
// Solidity: function createNewBridgeToken(string _symbol) returns(address)
func (_BridgeBank *BridgeBankTransactorSession) CreateNewBridgeToken(_symbol string) (*types.Transaction, error) {
	return _BridgeBank.Contract.CreateNewBridgeToken(&_BridgeBank.TransactOpts, _symbol)
}

// Lock is a paid mutator transaction binding the contract method 0x9df2a385.
//
// Solidity: function lock(bytes _recipient, address _token, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactor) Lock(opts *bind.TransactOpts, _recipient []byte, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "lock", _recipient, _token, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x9df2a385.
//
// Solidity: function lock(bytes _recipient, address _token, uint256 _amount) returns()
func (_BridgeBank *BridgeBankSession) Lock(_recipient []byte, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Lock(&_BridgeBank.TransactOpts, _recipient, _token, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x9df2a385.
//
// Solidity: function lock(bytes _recipient, address _token, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactorSession) Lock(_recipient []byte, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Lock(&_BridgeBank.TransactOpts, _recipient, _token, _amount)
}

// MintBridgeTokens is a paid mutator transaction binding the contract method 0xcdf68c41.
//
// Solidity: function mintBridgeTokens(bytes _cosmosSender, address _intendedRecipient, address _bridgeTokenAddress, string _symbol, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactor) MintBridgeTokens(opts *bind.TransactOpts, _cosmosSender []byte, _intendedRecipient common.Address, _bridgeTokenAddress common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "mintBridgeTokens", _cosmosSender, _intendedRecipient, _bridgeTokenAddress, _symbol, _amount)
}

// MintBridgeTokens is a paid mutator transaction binding the contract method 0xcdf68c41.
//
// Solidity: function mintBridgeTokens(bytes _cosmosSender, address _intendedRecipient, address _bridgeTokenAddress, string _symbol, uint256 _amount) returns()
func (_BridgeBank *BridgeBankSession) MintBridgeTokens(_cosmosSender []byte, _intendedRecipient common.Address, _bridgeTokenAddress common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.MintBridgeTokens(&_BridgeBank.TransactOpts, _cosmosSender, _intendedRecipient, _bridgeTokenAddress, _symbol, _amount)
}

// MintBridgeTokens is a paid mutator transaction binding the contract method 0xcdf68c41.
//
// Solidity: function mintBridgeTokens(bytes _cosmosSender, address _intendedRecipient, address _bridgeTokenAddress, string _symbol, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactorSession) MintBridgeTokens(_cosmosSender []byte, _intendedRecipient common.Address, _bridgeTokenAddress common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.MintBridgeTokens(&_BridgeBank.TransactOpts, _cosmosSender, _intendedRecipient, _bridgeTokenAddress, _symbol, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f89c91c.
//
// Solidity: function unlock(address _recipient, address _token, string _symbol, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactor) Unlock(opts *bind.TransactOpts, _recipient common.Address, _token common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "unlock", _recipient, _token, _symbol, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f89c91c.
//
// Solidity: function unlock(address _recipient, address _token, string _symbol, uint256 _amount) returns()
func (_BridgeBank *BridgeBankSession) Unlock(_recipient common.Address, _token common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Unlock(&_BridgeBank.TransactOpts, _recipient, _token, _symbol, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f89c91c.
//
// Solidity: function unlock(address _recipient, address _token, string _symbol, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactorSession) Unlock(_recipient common.Address, _token common.Address, _symbol string, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Unlock(&_BridgeBank.TransactOpts, _recipient, _token, _symbol, _amount)
}

// BridgeBankLogBridgeTokenMintIterator is returned from FilterLogBridgeTokenMint and is used to iterate over the raw logs and unpacked data for LogBridgeTokenMint events raised by the BridgeBank contract.
type BridgeBankLogBridgeTokenMintIterator struct {
	Event *BridgeBankLogBridgeTokenMint // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogBridgeTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogBridgeTokenMint)
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
		it.Event = new(BridgeBankLogBridgeTokenMint)
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
func (it *BridgeBankLogBridgeTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogBridgeTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogBridgeTokenMint represents a LogBridgeTokenMint event raised by the BridgeBank contract.
type BridgeBankLogBridgeTokenMint struct {
	Token       common.Address
	Symbol      string
	Amount      *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogBridgeTokenMint is a free log retrieval operation binding the contract event 0x262f97360779b7c2bb05fd24ef49f22d51435f78d3abd1ab35c323b22064cd4d.
//
// Solidity: event LogBridgeTokenMint(address _token, string _symbol, uint256 _amount, address _beneficiary)
func (_BridgeBank *BridgeBankFilterer) FilterLogBridgeTokenMint(opts *bind.FilterOpts) (*BridgeBankLogBridgeTokenMintIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogBridgeTokenMint")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogBridgeTokenMintIterator{contract: _BridgeBank.contract, event: "LogBridgeTokenMint", logs: logs, sub: sub}, nil
}

// WatchLogBridgeTokenMint is a free log subscription operation binding the contract event 0x262f97360779b7c2bb05fd24ef49f22d51435f78d3abd1ab35c323b22064cd4d.
//
// Solidity: event LogBridgeTokenMint(address _token, string _symbol, uint256 _amount, address _beneficiary)
func (_BridgeBank *BridgeBankFilterer) WatchLogBridgeTokenMint(opts *bind.WatchOpts, sink chan<- *BridgeBankLogBridgeTokenMint) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogBridgeTokenMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogBridgeTokenMint)
				if err := _BridgeBank.contract.UnpackLog(event, "LogBridgeTokenMint", log); err != nil {
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

// BridgeBankLogLockIterator is returned from FilterLogLock and is used to iterate over the raw logs and unpacked data for LogLock events raised by the BridgeBank contract.
type BridgeBankLogLockIterator struct {
	Event *BridgeBankLogLock // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogLock)
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
		it.Event = new(BridgeBankLogLock)
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
func (it *BridgeBankLogLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogLock represents a LogLock event raised by the BridgeBank contract.
type BridgeBankLogLock struct {
	From   common.Address
	To     []byte
	Token  common.Address
	Symbol string
	Value  *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogLock is a free log retrieval operation binding the contract event 0x374449c83a37309524754bbdfc5b8306d3694b5d14609b8fbb1b50cc5c0319a7.
//
// Solidity: event LogLock(address _from, bytes _to, address _token, string _symbol, uint256 _value, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) FilterLogLock(opts *bind.FilterOpts) (*BridgeBankLogLockIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogLock")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogLockIterator{contract: _BridgeBank.contract, event: "LogLock", logs: logs, sub: sub}, nil
}

// WatchLogLock is a free log subscription operation binding the contract event 0x374449c83a37309524754bbdfc5b8306d3694b5d14609b8fbb1b50cc5c0319a7.
//
// Solidity: event LogLock(address _from, bytes _to, address _token, string _symbol, uint256 _value, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) WatchLogLock(opts *bind.WatchOpts, sink chan<- *BridgeBankLogLock) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogLock)
				if err := _BridgeBank.contract.UnpackLog(event, "LogLock", log); err != nil {
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

// BridgeBankLogNewBridgeTokenIterator is returned from FilterLogNewBridgeToken and is used to iterate over the raw logs and unpacked data for LogNewBridgeToken events raised by the BridgeBank contract.
type BridgeBankLogNewBridgeTokenIterator struct {
	Event *BridgeBankLogNewBridgeToken // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogNewBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogNewBridgeToken)
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
		it.Event = new(BridgeBankLogNewBridgeToken)
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
func (it *BridgeBankLogNewBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogNewBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogNewBridgeToken represents a LogNewBridgeToken event raised by the BridgeBank contract.
type BridgeBankLogNewBridgeToken struct {
	Token  common.Address
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogNewBridgeToken is a free log retrieval operation binding the contract event 0x0ec4ab372af15f8db6003eb14d91402a44b20dff79fbac33b4ee0df68fafe9c0.
//
// Solidity: event LogNewBridgeToken(address _token, string _symbol)
func (_BridgeBank *BridgeBankFilterer) FilterLogNewBridgeToken(opts *bind.FilterOpts) (*BridgeBankLogNewBridgeTokenIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogNewBridgeToken")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogNewBridgeTokenIterator{contract: _BridgeBank.contract, event: "LogNewBridgeToken", logs: logs, sub: sub}, nil
}

// WatchLogNewBridgeToken is a free log subscription operation binding the contract event 0x0ec4ab372af15f8db6003eb14d91402a44b20dff79fbac33b4ee0df68fafe9c0.
//
// Solidity: event LogNewBridgeToken(address _token, string _symbol)
func (_BridgeBank *BridgeBankFilterer) WatchLogNewBridgeToken(opts *bind.WatchOpts, sink chan<- *BridgeBankLogNewBridgeToken) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogNewBridgeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogNewBridgeToken)
				if err := _BridgeBank.contract.UnpackLog(event, "LogNewBridgeToken", log); err != nil {
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

// BridgeBankLogUnlockIterator is returned from FilterLogUnlock and is used to iterate over the raw logs and unpacked data for LogUnlock events raised by the BridgeBank contract.
type BridgeBankLogUnlockIterator struct {
	Event *BridgeBankLogUnlock // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogUnlock)
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
		it.Event = new(BridgeBankLogUnlock)
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
func (it *BridgeBankLogUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogUnlock represents a LogUnlock event raised by the BridgeBank contract.
type BridgeBankLogUnlock struct {
	To     common.Address
	Token  common.Address
	Symbol string
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogUnlock is a free log retrieval operation binding the contract event 0x802cd873de701272ec903860b690986bd460b5bcd57e30ac1fdfdeece10528ac.
//
// Solidity: event LogUnlock(address _to, address _token, string _symbol, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) FilterLogUnlock(opts *bind.FilterOpts) (*BridgeBankLogUnlockIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogUnlockIterator{contract: _BridgeBank.contract, event: "LogUnlock", logs: logs, sub: sub}, nil
}

// WatchLogUnlock is a free log subscription operation binding the contract event 0x802cd873de701272ec903860b690986bd460b5bcd57e30ac1fdfdeece10528ac.
//
// Solidity: event LogUnlock(address _to, address _token, string _symbol, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) WatchLogUnlock(opts *bind.WatchOpts, sink chan<- *BridgeBankLogUnlock) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogUnlock)
				if err := _BridgeBank.contract.UnpackLog(event, "LogUnlock", log); err != nil {
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
