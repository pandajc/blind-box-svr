// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nftmarket

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

// NFTMarketOrder is an auto generated low-level Go binding around an user-defined struct.
type NFTMarketOrder struct {
	Token   common.Address
	TokenId *big.Int
	Owner   common.Address
	Price   *big.Int
	Amount  *big.Int
}

// NftmarketMetaData contains all meta data concerning the Nftmarket contract.
var NftmarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenIdOrderExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"List\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Purchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"UpdateAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"UpdatePrice\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structNFTMarket.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"updateAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NftmarketABI is the input ABI used to generate the binding from.
// Deprecated: Use NftmarketMetaData.ABI instead.
var NftmarketABI = NftmarketMetaData.ABI

// Nftmarket is an auto generated Go binding around an Ethereum contract.
type Nftmarket struct {
	NftmarketCaller     // Read-only binding to the contract
	NftmarketTransactor // Write-only binding to the contract
	NftmarketFilterer   // Log filterer for contract events
}

// NftmarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftmarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftmarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftmarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftmarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftmarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftmarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftmarketSession struct {
	Contract     *Nftmarket        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftmarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftmarketCallerSession struct {
	Contract *NftmarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NftmarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftmarketTransactorSession struct {
	Contract     *NftmarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NftmarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftmarketRaw struct {
	Contract *Nftmarket // Generic contract binding to access the raw methods on
}

// NftmarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftmarketCallerRaw struct {
	Contract *NftmarketCaller // Generic read-only contract binding to access the raw methods on
}

// NftmarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftmarketTransactorRaw struct {
	Contract *NftmarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftmarket creates a new instance of Nftmarket, bound to a specific deployed contract.
func NewNftmarket(address common.Address, backend bind.ContractBackend) (*Nftmarket, error) {
	contract, err := bindNftmarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nftmarket{NftmarketCaller: NftmarketCaller{contract: contract}, NftmarketTransactor: NftmarketTransactor{contract: contract}, NftmarketFilterer: NftmarketFilterer{contract: contract}}, nil
}

// NewNftmarketCaller creates a new read-only instance of Nftmarket, bound to a specific deployed contract.
func NewNftmarketCaller(address common.Address, caller bind.ContractCaller) (*NftmarketCaller, error) {
	contract, err := bindNftmarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftmarketCaller{contract: contract}, nil
}

// NewNftmarketTransactor creates a new write-only instance of Nftmarket, bound to a specific deployed contract.
func NewNftmarketTransactor(address common.Address, transactor bind.ContractTransactor) (*NftmarketTransactor, error) {
	contract, err := bindNftmarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftmarketTransactor{contract: contract}, nil
}

// NewNftmarketFilterer creates a new log filterer instance of Nftmarket, bound to a specific deployed contract.
func NewNftmarketFilterer(address common.Address, filterer bind.ContractFilterer) (*NftmarketFilterer, error) {
	contract, err := bindNftmarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftmarketFilterer{contract: contract}, nil
}

// bindNftmarket binds a generic wrapper to an already deployed contract.
func bindNftmarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftmarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nftmarket *NftmarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nftmarket.Contract.NftmarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nftmarket *NftmarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nftmarket.Contract.NftmarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nftmarket *NftmarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nftmarket.Contract.NftmarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nftmarket *NftmarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nftmarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nftmarket *NftmarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nftmarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nftmarket *NftmarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nftmarket.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _tokenId) view returns((address,uint256,address,uint256,uint256))
func (_Nftmarket *NftmarketCaller) Get(opts *bind.CallOpts, _tokenId *big.Int) (NFTMarketOrder, error) {
	var out []interface{}
	err := _Nftmarket.contract.Call(opts, &out, "get", _tokenId)

	if err != nil {
		return *new(NFTMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(NFTMarketOrder)).(*NFTMarketOrder)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _tokenId) view returns((address,uint256,address,uint256,uint256))
func (_Nftmarket *NftmarketSession) Get(_tokenId *big.Int) (NFTMarketOrder, error) {
	return _Nftmarket.Contract.Get(&_Nftmarket.CallOpts, _tokenId)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _tokenId) view returns((address,uint256,address,uint256,uint256))
func (_Nftmarket *NftmarketCallerSession) Get(_tokenId *big.Int) (NFTMarketOrder, error) {
	return _Nftmarket.Contract.Get(&_Nftmarket.CallOpts, _tokenId)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address token, uint256 tokenId, address owner, uint256 price, uint256 amount)
func (_Nftmarket *NftmarketCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token   common.Address
	TokenId *big.Int
	Owner   common.Address
	Price   *big.Int
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _Nftmarket.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Token   common.Address
		TokenId *big.Int
		Owner   common.Address
		Price   *big.Int
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address token, uint256 tokenId, address owner, uint256 price, uint256 amount)
func (_Nftmarket *NftmarketSession) Orders(arg0 *big.Int) (struct {
	Token   common.Address
	TokenId *big.Int
	Owner   common.Address
	Price   *big.Int
	Amount  *big.Int
}, error) {
	return _Nftmarket.Contract.Orders(&_Nftmarket.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address token, uint256 tokenId, address owner, uint256 price, uint256 amount)
func (_Nftmarket *NftmarketCallerSession) Orders(arg0 *big.Int) (struct {
	Token   common.Address
	TokenId *big.Int
	Owner   common.Address
	Price   *big.Int
	Amount  *big.Int
}, error) {
	return _Nftmarket.Contract.Orders(&_Nftmarket.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nftmarket *NftmarketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Nftmarket.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nftmarket *NftmarketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nftmarket.Contract.SupportsInterface(&_Nftmarket.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nftmarket *NftmarketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nftmarket.Contract.SupportsInterface(&_Nftmarket.CallOpts, interfaceId)
}

// List is a paid mutator transaction binding the contract method 0xbb74a1c2.
//
// Solidity: function list(address _nftAddr, uint256 _tokenId, uint256 _price, uint256 _amount) returns(uint256)
func (_Nftmarket *NftmarketTransactor) List(opts *bind.TransactOpts, _nftAddr common.Address, _tokenId *big.Int, _price *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Nftmarket.contract.Transact(opts, "list", _nftAddr, _tokenId, _price, _amount)
}

// List is a paid mutator transaction binding the contract method 0xbb74a1c2.
//
// Solidity: function list(address _nftAddr, uint256 _tokenId, uint256 _price, uint256 _amount) returns(uint256)
func (_Nftmarket *NftmarketSession) List(_nftAddr common.Address, _tokenId *big.Int, _price *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.List(&_Nftmarket.TransactOpts, _nftAddr, _tokenId, _price, _amount)
}

// List is a paid mutator transaction binding the contract method 0xbb74a1c2.
//
// Solidity: function list(address _nftAddr, uint256 _tokenId, uint256 _price, uint256 _amount) returns(uint256)
func (_Nftmarket *NftmarketTransactorSession) List(_nftAddr common.Address, _tokenId *big.Int, _price *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.List(&_Nftmarket.TransactOpts, _nftAddr, _tokenId, _price, _amount)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Nftmarket *NftmarketTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Nftmarket.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Nftmarket *NftmarketSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Nftmarket.Contract.OnERC1155BatchReceived(&_Nftmarket.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Nftmarket *NftmarketTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Nftmarket.Contract.OnERC1155BatchReceived(&_Nftmarket.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Nftmarket *NftmarketTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Nftmarket.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Nftmarket *NftmarketSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Nftmarket.Contract.OnERC1155Received(&_Nftmarket.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Nftmarket *NftmarketTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Nftmarket.Contract.OnERC1155Received(&_Nftmarket.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// Purchase is a paid mutator transaction binding the contract method 0x70876c98.
//
// Solidity: function purchase(uint256 _orderId, uint256 _amount) payable returns()
func (_Nftmarket *NftmarketTransactor) Purchase(opts *bind.TransactOpts, _orderId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Nftmarket.contract.Transact(opts, "purchase", _orderId, _amount)
}

// Purchase is a paid mutator transaction binding the contract method 0x70876c98.
//
// Solidity: function purchase(uint256 _orderId, uint256 _amount) payable returns()
func (_Nftmarket *NftmarketSession) Purchase(_orderId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.Purchase(&_Nftmarket.TransactOpts, _orderId, _amount)
}

// Purchase is a paid mutator transaction binding the contract method 0x70876c98.
//
// Solidity: function purchase(uint256 _orderId, uint256 _amount) payable returns()
func (_Nftmarket *NftmarketTransactorSession) Purchase(_orderId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.Purchase(&_Nftmarket.TransactOpts, _orderId, _amount)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 _orderId) returns()
func (_Nftmarket *NftmarketTransactor) Revoke(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _Nftmarket.contract.Transact(opts, "revoke", _orderId)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 _orderId) returns()
func (_Nftmarket *NftmarketSession) Revoke(_orderId *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.Revoke(&_Nftmarket.TransactOpts, _orderId)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 _orderId) returns()
func (_Nftmarket *NftmarketTransactorSession) Revoke(_orderId *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.Revoke(&_Nftmarket.TransactOpts, _orderId)
}

// UpdateAmount is a paid mutator transaction binding the contract method 0x0179cc52.
//
// Solidity: function updateAmount(uint256 _orderId, uint256 _amount) returns()
func (_Nftmarket *NftmarketTransactor) UpdateAmount(opts *bind.TransactOpts, _orderId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Nftmarket.contract.Transact(opts, "updateAmount", _orderId, _amount)
}

// UpdateAmount is a paid mutator transaction binding the contract method 0x0179cc52.
//
// Solidity: function updateAmount(uint256 _orderId, uint256 _amount) returns()
func (_Nftmarket *NftmarketSession) UpdateAmount(_orderId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.UpdateAmount(&_Nftmarket.TransactOpts, _orderId, _amount)
}

// UpdateAmount is a paid mutator transaction binding the contract method 0x0179cc52.
//
// Solidity: function updateAmount(uint256 _orderId, uint256 _amount) returns()
func (_Nftmarket *NftmarketTransactorSession) UpdateAmount(_orderId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.UpdateAmount(&_Nftmarket.TransactOpts, _orderId, _amount)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x82367b2d.
//
// Solidity: function updatePrice(uint256 _orderId, uint256 _newPrice) returns()
func (_Nftmarket *NftmarketTransactor) UpdatePrice(opts *bind.TransactOpts, _orderId *big.Int, _newPrice *big.Int) (*types.Transaction, error) {
	return _Nftmarket.contract.Transact(opts, "updatePrice", _orderId, _newPrice)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x82367b2d.
//
// Solidity: function updatePrice(uint256 _orderId, uint256 _newPrice) returns()
func (_Nftmarket *NftmarketSession) UpdatePrice(_orderId *big.Int, _newPrice *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.UpdatePrice(&_Nftmarket.TransactOpts, _orderId, _newPrice)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x82367b2d.
//
// Solidity: function updatePrice(uint256 _orderId, uint256 _newPrice) returns()
func (_Nftmarket *NftmarketTransactorSession) UpdatePrice(_orderId *big.Int, _newPrice *big.Int) (*types.Transaction, error) {
	return _Nftmarket.Contract.UpdatePrice(&_Nftmarket.TransactOpts, _orderId, _newPrice)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nftmarket *NftmarketTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Nftmarket.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nftmarket *NftmarketSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Nftmarket.Contract.Fallback(&_Nftmarket.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nftmarket *NftmarketTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Nftmarket.Contract.Fallback(&_Nftmarket.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Nftmarket *NftmarketTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nftmarket.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Nftmarket *NftmarketSession) Receive() (*types.Transaction, error) {
	return _Nftmarket.Contract.Receive(&_Nftmarket.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Nftmarket *NftmarketTransactorSession) Receive() (*types.Transaction, error) {
	return _Nftmarket.Contract.Receive(&_Nftmarket.TransactOpts)
}

// NftmarketListIterator is returned from FilterList and is used to iterate over the raw logs and unpacked data for List events raised by the Nftmarket contract.
type NftmarketListIterator struct {
	Event *NftmarketList // Event containing the contract specifics and raw log

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
func (it *NftmarketListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftmarketList)
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
		it.Event = new(NftmarketList)
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
func (it *NftmarketListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftmarketListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftmarketList represents a List event raised by the Nftmarket contract.
type NftmarketList struct {
	Seller  common.Address
	NftAddr common.Address
	OrderId *big.Int
	Price   *big.Int
	Amount  *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterList is a free log retrieval operation binding the contract event 0xb9f26d3c47b1e6cc5b8e24d075d969f3a79727a118e23cf23cdec7d26f74c682.
//
// Solidity: event List(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 price, uint256 amount, uint256 tokenId)
func (_Nftmarket *NftmarketFilterer) FilterList(opts *bind.FilterOpts, seller []common.Address, nftAddr []common.Address, orderId []*big.Int) (*NftmarketListIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.FilterLogs(opts, "List", sellerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return &NftmarketListIterator{contract: _Nftmarket.contract, event: "List", logs: logs, sub: sub}, nil
}

// WatchList is a free log subscription operation binding the contract event 0xb9f26d3c47b1e6cc5b8e24d075d969f3a79727a118e23cf23cdec7d26f74c682.
//
// Solidity: event List(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 price, uint256 amount, uint256 tokenId)
func (_Nftmarket *NftmarketFilterer) WatchList(opts *bind.WatchOpts, sink chan<- *NftmarketList, seller []common.Address, nftAddr []common.Address, orderId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.WatchLogs(opts, "List", sellerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftmarketList)
				if err := _Nftmarket.contract.UnpackLog(event, "List", log); err != nil {
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

// ParseList is a log parse operation binding the contract event 0xb9f26d3c47b1e6cc5b8e24d075d969f3a79727a118e23cf23cdec7d26f74c682.
//
// Solidity: event List(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 price, uint256 amount, uint256 tokenId)
func (_Nftmarket *NftmarketFilterer) ParseList(log types.Log) (*NftmarketList, error) {
	event := new(NftmarketList)
	if err := _Nftmarket.contract.UnpackLog(event, "List", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftmarketPurchaseIterator is returned from FilterPurchase and is used to iterate over the raw logs and unpacked data for Purchase events raised by the Nftmarket contract.
type NftmarketPurchaseIterator struct {
	Event *NftmarketPurchase // Event containing the contract specifics and raw log

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
func (it *NftmarketPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftmarketPurchase)
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
		it.Event = new(NftmarketPurchase)
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
func (it *NftmarketPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftmarketPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftmarketPurchase represents a Purchase event raised by the Nftmarket contract.
type NftmarketPurchase struct {
	Buyer   common.Address
	NftAddr common.Address
	OrderId *big.Int
	Price   *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPurchase is a free log retrieval operation binding the contract event 0x8f7a55179307cea51948432d653dbd53a23fedc388bcb3e04e311f8220d87864.
//
// Solidity: event Purchase(address indexed buyer, address indexed nftAddr, uint256 indexed orderId, uint256 price, uint256 amount)
func (_Nftmarket *NftmarketFilterer) FilterPurchase(opts *bind.FilterOpts, buyer []common.Address, nftAddr []common.Address, orderId []*big.Int) (*NftmarketPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.FilterLogs(opts, "Purchase", buyerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return &NftmarketPurchaseIterator{contract: _Nftmarket.contract, event: "Purchase", logs: logs, sub: sub}, nil
}

// WatchPurchase is a free log subscription operation binding the contract event 0x8f7a55179307cea51948432d653dbd53a23fedc388bcb3e04e311f8220d87864.
//
// Solidity: event Purchase(address indexed buyer, address indexed nftAddr, uint256 indexed orderId, uint256 price, uint256 amount)
func (_Nftmarket *NftmarketFilterer) WatchPurchase(opts *bind.WatchOpts, sink chan<- *NftmarketPurchase, buyer []common.Address, nftAddr []common.Address, orderId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.WatchLogs(opts, "Purchase", buyerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftmarketPurchase)
				if err := _Nftmarket.contract.UnpackLog(event, "Purchase", log); err != nil {
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

// ParsePurchase is a log parse operation binding the contract event 0x8f7a55179307cea51948432d653dbd53a23fedc388bcb3e04e311f8220d87864.
//
// Solidity: event Purchase(address indexed buyer, address indexed nftAddr, uint256 indexed orderId, uint256 price, uint256 amount)
func (_Nftmarket *NftmarketFilterer) ParsePurchase(log types.Log) (*NftmarketPurchase, error) {
	event := new(NftmarketPurchase)
	if err := _Nftmarket.contract.UnpackLog(event, "Purchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftmarketRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the Nftmarket contract.
type NftmarketRevokeIterator struct {
	Event *NftmarketRevoke // Event containing the contract specifics and raw log

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
func (it *NftmarketRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftmarketRevoke)
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
		it.Event = new(NftmarketRevoke)
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
func (it *NftmarketRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftmarketRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftmarketRevoke represents a Revoke event raised by the Nftmarket contract.
type NftmarketRevoke struct {
	Seller  common.Address
	NftAddr common.Address
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0xb698e31a2abee5824d0d7bcfd2339aead7f9e9ae413fba50bf554ff3fa470b7b.
//
// Solidity: event Revoke(address indexed seller, address indexed nftAddr, uint256 indexed orderId)
func (_Nftmarket *NftmarketFilterer) FilterRevoke(opts *bind.FilterOpts, seller []common.Address, nftAddr []common.Address, orderId []*big.Int) (*NftmarketRevokeIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.FilterLogs(opts, "Revoke", sellerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return &NftmarketRevokeIterator{contract: _Nftmarket.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0xb698e31a2abee5824d0d7bcfd2339aead7f9e9ae413fba50bf554ff3fa470b7b.
//
// Solidity: event Revoke(address indexed seller, address indexed nftAddr, uint256 indexed orderId)
func (_Nftmarket *NftmarketFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *NftmarketRevoke, seller []common.Address, nftAddr []common.Address, orderId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.WatchLogs(opts, "Revoke", sellerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftmarketRevoke)
				if err := _Nftmarket.contract.UnpackLog(event, "Revoke", log); err != nil {
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

// ParseRevoke is a log parse operation binding the contract event 0xb698e31a2abee5824d0d7bcfd2339aead7f9e9ae413fba50bf554ff3fa470b7b.
//
// Solidity: event Revoke(address indexed seller, address indexed nftAddr, uint256 indexed orderId)
func (_Nftmarket *NftmarketFilterer) ParseRevoke(log types.Log) (*NftmarketRevoke, error) {
	event := new(NftmarketRevoke)
	if err := _Nftmarket.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftmarketUpdateAmountIterator is returned from FilterUpdateAmount and is used to iterate over the raw logs and unpacked data for UpdateAmount events raised by the Nftmarket contract.
type NftmarketUpdateAmountIterator struct {
	Event *NftmarketUpdateAmount // Event containing the contract specifics and raw log

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
func (it *NftmarketUpdateAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftmarketUpdateAmount)
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
		it.Event = new(NftmarketUpdateAmount)
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
func (it *NftmarketUpdateAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftmarketUpdateAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftmarketUpdateAmount represents a UpdateAmount event raised by the Nftmarket contract.
type NftmarketUpdateAmount struct {
	Seller    common.Address
	NftAddr   common.Address
	OrderId   *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateAmount is a free log retrieval operation binding the contract event 0x1ce11c0a5c07a2f5318c58c72462d38c3bf9d28fd1d22b42bb2c61f409f97b1e.
//
// Solidity: event UpdateAmount(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 newAmount)
func (_Nftmarket *NftmarketFilterer) FilterUpdateAmount(opts *bind.FilterOpts, seller []common.Address, nftAddr []common.Address, orderId []*big.Int) (*NftmarketUpdateAmountIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.FilterLogs(opts, "UpdateAmount", sellerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return &NftmarketUpdateAmountIterator{contract: _Nftmarket.contract, event: "UpdateAmount", logs: logs, sub: sub}, nil
}

// WatchUpdateAmount is a free log subscription operation binding the contract event 0x1ce11c0a5c07a2f5318c58c72462d38c3bf9d28fd1d22b42bb2c61f409f97b1e.
//
// Solidity: event UpdateAmount(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 newAmount)
func (_Nftmarket *NftmarketFilterer) WatchUpdateAmount(opts *bind.WatchOpts, sink chan<- *NftmarketUpdateAmount, seller []common.Address, nftAddr []common.Address, orderId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.WatchLogs(opts, "UpdateAmount", sellerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftmarketUpdateAmount)
				if err := _Nftmarket.contract.UnpackLog(event, "UpdateAmount", log); err != nil {
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

// ParseUpdateAmount is a log parse operation binding the contract event 0x1ce11c0a5c07a2f5318c58c72462d38c3bf9d28fd1d22b42bb2c61f409f97b1e.
//
// Solidity: event UpdateAmount(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 newAmount)
func (_Nftmarket *NftmarketFilterer) ParseUpdateAmount(log types.Log) (*NftmarketUpdateAmount, error) {
	event := new(NftmarketUpdateAmount)
	if err := _Nftmarket.contract.UnpackLog(event, "UpdateAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftmarketUpdatePriceIterator is returned from FilterUpdatePrice and is used to iterate over the raw logs and unpacked data for UpdatePrice events raised by the Nftmarket contract.
type NftmarketUpdatePriceIterator struct {
	Event *NftmarketUpdatePrice // Event containing the contract specifics and raw log

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
func (it *NftmarketUpdatePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftmarketUpdatePrice)
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
		it.Event = new(NftmarketUpdatePrice)
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
func (it *NftmarketUpdatePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftmarketUpdatePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftmarketUpdatePrice represents a UpdatePrice event raised by the Nftmarket contract.
type NftmarketUpdatePrice struct {
	Seller   common.Address
	NftAddr  common.Address
	OrderId  *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatePrice is a free log retrieval operation binding the contract event 0xa1d79c5222548888b4d922f69eccfbc97181419ea8d09ccd06b6aa2756682542.
//
// Solidity: event UpdatePrice(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 newPrice)
func (_Nftmarket *NftmarketFilterer) FilterUpdatePrice(opts *bind.FilterOpts, seller []common.Address, nftAddr []common.Address, orderId []*big.Int) (*NftmarketUpdatePriceIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.FilterLogs(opts, "UpdatePrice", sellerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return &NftmarketUpdatePriceIterator{contract: _Nftmarket.contract, event: "UpdatePrice", logs: logs, sub: sub}, nil
}

// WatchUpdatePrice is a free log subscription operation binding the contract event 0xa1d79c5222548888b4d922f69eccfbc97181419ea8d09ccd06b6aa2756682542.
//
// Solidity: event UpdatePrice(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 newPrice)
func (_Nftmarket *NftmarketFilterer) WatchUpdatePrice(opts *bind.WatchOpts, sink chan<- *NftmarketUpdatePrice, seller []common.Address, nftAddr []common.Address, orderId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Nftmarket.contract.WatchLogs(opts, "UpdatePrice", sellerRule, nftAddrRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftmarketUpdatePrice)
				if err := _Nftmarket.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
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

// ParseUpdatePrice is a log parse operation binding the contract event 0xa1d79c5222548888b4d922f69eccfbc97181419ea8d09ccd06b6aa2756682542.
//
// Solidity: event UpdatePrice(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 newPrice)
func (_Nftmarket *NftmarketFilterer) ParseUpdatePrice(log types.Log) (*NftmarketUpdatePrice, error) {
	event := new(NftmarketUpdatePrice)
	if err := _Nftmarket.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
