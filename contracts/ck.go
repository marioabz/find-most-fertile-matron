// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"promoCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_STARTING_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSiringAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_AUCTION_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSaleAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSiringAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createPromoKitty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSaleAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_v2Address\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAuctionBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"}],\"name\":\"createGen0Auction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROMO_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getKitty\",\"outputs\":[{\"name\":\"isGestating\",\"type\":\"bool\"},{\"name\":\"isReady\",\"type\":\"bool\"},{\"name\":\"cooldownIndex\",\"type\":\"uint256\"},{\"name\":\"nextActionAt\",\"type\":\"uint256\"},{\"name\":\"siringWithId\",\"type\":\"uint256\"},{\"name\":\"birthTime\",\"type\":\"uint256\"},{\"name\":\"matronId\",\"type\":\"uint256\"},{\"name\":\"sireId\",\"type\":\"uint256\"},{\"name\":\"generation\",\"type\":\"uint256\"},{\"name\":\"genes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sireId\",\"type\":\"uint256\"},{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"bidOnSiringAuction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0CreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() view returns(uint256)
func (_Token *TokenCaller) GEN0AUCTIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "GEN0_AUCTION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() view returns(uint256)
func (_Token *TokenSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _Token.Contract.GEN0AUCTIONDURATION(&_Token.CallOpts)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() view returns(uint256)
func (_Token *TokenCallerSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _Token.Contract.GEN0AUCTIONDURATION(&_Token.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() view returns(uint256)
func (_Token *TokenCaller) GEN0CREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "GEN0_CREATION_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() view returns(uint256)
func (_Token *TokenSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _Token.Contract.GEN0CREATIONLIMIT(&_Token.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() view returns(uint256)
func (_Token *TokenCallerSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _Token.Contract.GEN0CREATIONLIMIT(&_Token.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() view returns(uint256)
func (_Token *TokenCaller) GEN0STARTINGPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "GEN0_STARTING_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() view returns(uint256)
func (_Token *TokenSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _Token.Contract.GEN0STARTINGPRICE(&_Token.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() view returns(uint256)
func (_Token *TokenCallerSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _Token.Contract.GEN0STARTINGPRICE(&_Token.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() view returns(uint256)
func (_Token *TokenCaller) PROMOCREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "PROMO_CREATION_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() view returns(uint256)
func (_Token *TokenSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _Token.Contract.PROMOCREATIONLIMIT(&_Token.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() view returns(uint256)
func (_Token *TokenCallerSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _Token.Contract.PROMOCREATIONLIMIT(&_Token.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_Token *TokenCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "autoBirthFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_Token *TokenSession) AutoBirthFee() (*big.Int, error) {
	return _Token.Contract.AutoBirthFee(&_Token.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_Token *TokenCallerSession) AutoBirthFee() (*big.Int, error) {
	return _Token.Contract.AutoBirthFee(&_Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_Token *TokenCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "canBreedWith", _matronId, _sireId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_Token *TokenSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _Token.Contract.CanBreedWith(&_Token.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_Token *TokenCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _Token.Contract.CanBreedWith(&_Token.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_Token *TokenCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "ceoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_Token *TokenSession) CeoAddress() (common.Address, error) {
	return _Token.Contract.CeoAddress(&_Token.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_Token *TokenCallerSession) CeoAddress() (common.Address, error) {
	return _Token.Contract.CeoAddress(&_Token.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_Token *TokenCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "cfoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_Token *TokenSession) CfoAddress() (common.Address, error) {
	return _Token.Contract.CfoAddress(&_Token.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_Token *TokenCallerSession) CfoAddress() (common.Address, error) {
	return _Token.Contract.CfoAddress(&_Token.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_Token *TokenCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "cooAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_Token *TokenSession) CooAddress() (common.Address, error) {
	return _Token.Contract.CooAddress(&_Token.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_Token *TokenCallerSession) CooAddress() (common.Address, error) {
	return _Token.Contract.CooAddress(&_Token.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_Token *TokenCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "cooldowns", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_Token *TokenSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _Token.Contract.Cooldowns(&_Token.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_Token *TokenCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _Token.Contract.Cooldowns(&_Token.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_Token *TokenCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "erc721Metadata")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_Token *TokenSession) Erc721Metadata() (common.Address, error) {
	return _Token.Contract.Erc721Metadata(&_Token.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_Token *TokenCallerSession) Erc721Metadata() (common.Address, error) {
	return _Token.Contract.Erc721Metadata(&_Token.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() view returns(uint256)
func (_Token *TokenCaller) Gen0CreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "gen0CreatedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() view returns(uint256)
func (_Token *TokenSession) Gen0CreatedCount() (*big.Int, error) {
	return _Token.Contract.Gen0CreatedCount(&_Token.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() view returns(uint256)
func (_Token *TokenCallerSession) Gen0CreatedCount() (*big.Int, error) {
	return _Token.Contract.Gen0CreatedCount(&_Token.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_Token *TokenCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "geneScience")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_Token *TokenSession) GeneScience() (common.Address, error) {
	return _Token.Contract.GeneScience(&_Token.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_Token *TokenCallerSession) GeneScience() (common.Address, error) {
	return _Token.Contract.GeneScience(&_Token.CallOpts)
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(uint256 _id) view returns(bool isGestating, bool isReady, uint256 cooldownIndex, uint256 nextActionAt, uint256 siringWithId, uint256 birthTime, uint256 matronId, uint256 sireId, uint256 generation, uint256 genes)
func (_Token *TokenCaller) GetKitty(opts *bind.CallOpts, _id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getKitty", _id)

	outstruct := new(struct {
		IsGestating   bool
		IsReady       bool
		CooldownIndex *big.Int
		NextActionAt  *big.Int
		SiringWithId  *big.Int
		BirthTime     *big.Int
		MatronId      *big.Int
		SireId        *big.Int
		Generation    *big.Int
		Genes         *big.Int
	})

	outstruct.IsGestating = out[0].(bool)
	outstruct.IsReady = out[1].(bool)
	outstruct.CooldownIndex = out[2].(*big.Int)
	outstruct.NextActionAt = out[3].(*big.Int)
	outstruct.SiringWithId = out[4].(*big.Int)
	outstruct.BirthTime = out[5].(*big.Int)
	outstruct.MatronId = out[6].(*big.Int)
	outstruct.SireId = out[7].(*big.Int)
	outstruct.Generation = out[8].(*big.Int)
	outstruct.Genes = out[9].(*big.Int)

	return *outstruct, err

}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(uint256 _id) view returns(bool isGestating, bool isReady, uint256 cooldownIndex, uint256 nextActionAt, uint256 siringWithId, uint256 birthTime, uint256 matronId, uint256 sireId, uint256 generation, uint256 genes)
func (_Token *TokenSession) GetKitty(_id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	return _Token.Contract.GetKitty(&_Token.CallOpts, _id)
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(uint256 _id) view returns(bool isGestating, bool isReady, uint256 cooldownIndex, uint256 nextActionAt, uint256 siringWithId, uint256 birthTime, uint256 matronId, uint256 sireId, uint256 generation, uint256 genes)
func (_Token *TokenCallerSession) GetKitty(_id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	return _Token.Contract.GetKitty(&_Token.CallOpts, _id)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_Token *TokenCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isPregnant", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_Token *TokenSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _Token.Contract.IsPregnant(&_Token.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_Token *TokenCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _Token.Contract.IsPregnant(&_Token.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_Token *TokenCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isReadyToBreed", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_Token *TokenSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _Token.Contract.IsReadyToBreed(&_Token.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_Token *TokenCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _Token.Contract.IsReadyToBreed(&_Token.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_Token *TokenCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "kittyIndexToApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_Token *TokenSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.KittyIndexToApproved(&_Token.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_Token *TokenCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.KittyIndexToApproved(&_Token.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_Token *TokenCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "kittyIndexToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_Token *TokenSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.KittyIndexToOwner(&_Token.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_Token *TokenCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.KittyIndexToOwner(&_Token.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() view returns(address)
func (_Token *TokenCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "newContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() view returns(address)
func (_Token *TokenSession) NewContractAddress() (common.Address, error) {
	return _Token.Contract.NewContractAddress(&_Token.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() view returns(address)
func (_Token *TokenCallerSession) NewContractAddress() (common.Address, error) {
	return _Token.Contract.NewContractAddress(&_Token.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_Token *TokenCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_Token *TokenSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.OwnerOf(&_Token.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_Token *TokenCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.OwnerOf(&_Token.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Token *TokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Token *TokenSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Token *TokenCallerSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_Token *TokenCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "pregnantKitties")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_Token *TokenSession) PregnantKitties() (*big.Int, error) {
	return _Token.Contract.PregnantKitties(&_Token.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_Token *TokenCallerSession) PregnantKitties() (*big.Int, error) {
	return _Token.Contract.PregnantKitties(&_Token.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() view returns(uint256)
func (_Token *TokenCaller) PromoCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "promoCreatedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() view returns(uint256)
func (_Token *TokenSession) PromoCreatedCount() (*big.Int, error) {
	return _Token.Contract.PromoCreatedCount(&_Token.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() view returns(uint256)
func (_Token *TokenCallerSession) PromoCreatedCount() (*big.Int, error) {
	return _Token.Contract.PromoCreatedCount(&_Token.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_Token *TokenCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "saleAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_Token *TokenSession) SaleAuction() (common.Address, error) {
	return _Token.Contract.SaleAuction(&_Token.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_Token *TokenCallerSession) SaleAuction() (common.Address, error) {
	return _Token.Contract.SaleAuction(&_Token.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_Token *TokenCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "secondsPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_Token *TokenSession) SecondsPerBlock() (*big.Int, error) {
	return _Token.Contract.SecondsPerBlock(&_Token.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_Token *TokenCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _Token.Contract.SecondsPerBlock(&_Token.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_Token *TokenCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "sireAllowedToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_Token *TokenSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.SireAllowedToAddress(&_Token.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_Token *TokenCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.SireAllowedToAddress(&_Token.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_Token *TokenCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "siringAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_Token *TokenSession) SiringAuction() (common.Address, error) {
	return _Token.Contract.SiringAuction(&_Token.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_Token *TokenCallerSession) SiringAuction() (common.Address, error) {
	return _Token.Contract.SiringAuction(&_Token.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_Token *TokenCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_Token *TokenSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _Token.Contract.SupportsInterface(&_Token.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_Token *TokenCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _Token.Contract.SupportsInterface(&_Token.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_Token *TokenCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tokenMetadata", _tokenId, _preferredTransport)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_Token *TokenSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _Token.Contract.TokenMetadata(&_Token.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_Token *TokenCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _Token.Contract.TokenMetadata(&_Token.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_Token *TokenCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_Token *TokenSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Token.Contract.TokensOfOwner(&_Token.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_Token *TokenCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Token.Contract.TokensOfOwner(&_Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Token *TokenSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Token *TokenTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_Token *TokenTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_Token *TokenSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ApproveSiring(&_Token.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_Token *TokenTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ApproveSiring(&_Token.TransactOpts, _addr, _sireId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_Token *TokenTransactor) BidOnSiringAuction(opts *bind.TransactOpts, _sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "bidOnSiringAuction", _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_Token *TokenSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BidOnSiringAuction(&_Token.TransactOpts, _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_Token *TokenTransactorSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BidOnSiringAuction(&_Token.TransactOpts, _sireId, _matronId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_Token *TokenTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_Token *TokenSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BreedWithAuto(&_Token.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_Token *TokenTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BreedWithAuto(&_Token.TransactOpts, _matronId, _sireId)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(uint256 _genes) returns()
func (_Token *TokenTransactor) CreateGen0Auction(opts *bind.TransactOpts, _genes *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createGen0Auction", _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(uint256 _genes) returns()
func (_Token *TokenSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateGen0Auction(&_Token.TransactOpts, _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(uint256 _genes) returns()
func (_Token *TokenTransactorSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateGen0Auction(&_Token.TransactOpts, _genes)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(uint256 _genes, address _owner) returns()
func (_Token *TokenTransactor) CreatePromoKitty(opts *bind.TransactOpts, _genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createPromoKitty", _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(uint256 _genes, address _owner) returns()
func (_Token *TokenSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Token.Contract.CreatePromoKitty(&_Token.TransactOpts, _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(uint256 _genes, address _owner) returns()
func (_Token *TokenTransactorSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Token.Contract.CreatePromoKitty(&_Token.TransactOpts, _genes, _owner)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_Token *TokenTransactor) CreateSaleAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createSaleAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_Token *TokenSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateSaleAuction(&_Token.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_Token *TokenTransactorSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateSaleAuction(&_Token.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_Token *TokenTransactor) CreateSiringAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createSiringAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_Token *TokenSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateSiringAuction(&_Token.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_Token *TokenTransactorSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateSiringAuction(&_Token.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_Token *TokenTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_Token *TokenSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.GiveBirth(&_Token.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_Token *TokenTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.GiveBirth(&_Token.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactorSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_Token *TokenTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_Token *TokenSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetAutoBirthFee(&_Token.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_Token *TokenTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetAutoBirthFee(&_Token.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_Token *TokenTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_Token *TokenSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetCEO(&_Token.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_Token *TokenTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetCEO(&_Token.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_Token *TokenTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_Token *TokenSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetCFO(&_Token.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_Token *TokenTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetCFO(&_Token.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_Token *TokenTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_Token *TokenSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetCOO(&_Token.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_Token *TokenTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetCOO(&_Token.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_Token *TokenTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_Token *TokenSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetGeneScienceAddress(&_Token.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_Token *TokenTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetGeneScienceAddress(&_Token.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_Token *TokenTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_Token *TokenSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetMetadataAddress(&_Token.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_Token *TokenTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetMetadataAddress(&_Token.TransactOpts, _contractAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(address _v2Address) returns()
func (_Token *TokenTransactor) SetNewAddress(opts *bind.TransactOpts, _v2Address common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setNewAddress", _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(address _v2Address) returns()
func (_Token *TokenSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetNewAddress(&_Token.TransactOpts, _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(address _v2Address) returns()
func (_Token *TokenTransactorSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetNewAddress(&_Token.TransactOpts, _v2Address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_Token *TokenTransactor) SetSaleAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSaleAuctionAddress", _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_Token *TokenSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetSaleAuctionAddress(&_Token.TransactOpts, _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_Token *TokenTransactorSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetSaleAuctionAddress(&_Token.TransactOpts, _address)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_Token *TokenTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_Token *TokenSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSecondsPerBlock(&_Token.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_Token *TokenTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSecondsPerBlock(&_Token.TransactOpts, secs)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_Token *TokenTransactor) SetSiringAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSiringAuctionAddress", _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_Token *TokenSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetSiringAuctionAddress(&_Token.TransactOpts, _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_Token *TokenTransactorSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetSiringAuctionAddress(&_Token.TransactOpts, _address)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_Token *TokenSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_Token *TokenTransactor) WithdrawAuctionBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawAuctionBalances")
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_Token *TokenSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _Token.Contract.WithdrawAuctionBalances(&_Token.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_Token *TokenTransactorSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _Token.Contract.WithdrawAuctionBalances(&_Token.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_Token *TokenTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_Token *TokenSession) WithdrawBalance() (*types.Transaction, error) {
	return _Token.Contract.WithdrawBalance(&_Token.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_Token *TokenTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _Token.Contract.WithdrawBalance(&_Token.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Token.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.Fallback(&_Token.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.Fallback(&_Token.TransactOpts, calldata)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts) (*TokenApprovalIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the Token contract.
type TokenBirthIterator struct {
	Event *TokenBirth // Event containing the contract specifics and raw log

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
func (it *TokenBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBirth)
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
		it.Event = new(TokenBirth)
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
func (it *TokenBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBirth represents a Birth event raised by the Token contract.
type TokenBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_Token *TokenFilterer) FilterBirth(opts *bind.FilterOpts) (*TokenBirthIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &TokenBirthIterator{contract: _Token.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_Token *TokenFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *TokenBirth) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBirth)
				if err := _Token.contract.UnpackLog(event, "Birth", log); err != nil {
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

// ParseBirth is a log parse operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_Token *TokenFilterer) ParseBirth(log types.Log) (*TokenBirth, error) {
	event := new(TokenBirth)
	if err := _Token.contract.UnpackLog(event, "Birth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the Token contract.
type TokenContractUpgradeIterator struct {
	Event *TokenContractUpgrade // Event containing the contract specifics and raw log

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
func (it *TokenContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContractUpgrade)
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
		it.Event = new(TokenContractUpgrade)
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
func (it *TokenContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContractUpgrade represents a ContractUpgrade event raised by the Token contract.
type TokenContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_Token *TokenFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*TokenContractUpgradeIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &TokenContractUpgradeIterator{contract: _Token.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_Token *TokenFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *TokenContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContractUpgrade)
				if err := _Token.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// ParseContractUpgrade is a log parse operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_Token *TokenFilterer) ParseContractUpgrade(log types.Log) (*TokenContractUpgrade, error) {
	event := new(TokenContractUpgrade)
	if err := _Token.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the Token contract.
type TokenPregnantIterator struct {
	Event *TokenPregnant // Event containing the contract specifics and raw log

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
func (it *TokenPregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPregnant)
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
		it.Event = new(TokenPregnant)
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
func (it *TokenPregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPregnant represents a Pregnant event raised by the Token contract.
type TokenPregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_Token *TokenFilterer) FilterPregnant(opts *bind.FilterOpts) (*TokenPregnantIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &TokenPregnantIterator{contract: _Token.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_Token *TokenFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *TokenPregnant) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPregnant)
				if err := _Token.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// ParsePregnant is a log parse operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_Token *TokenFilterer) ParsePregnant(log types.Log) (*TokenPregnant, error) {
	event := new(TokenPregnant)
	if err := _Token.contract.UnpackLog(event, "Pregnant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts) (*TokenTransferIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
