// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bookingsystem

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

// BookRoomRoom is an auto generated low-level Go binding around an user-defined struct.
type BookRoomRoom struct {
	Id           *big.Int
	Name         string
	Owner        common.Address
	Capacity     *big.Int
	WorkingSlots *big.Int
}

// BookingsystemMetaData contains all meta data concerning the Bookingsystem contract.
var BookingsystemMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"allRoomsByDate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workingSlots\",\"type\":\"uint256\"}],\"internalType\":\"structBookRoom.Room[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createRoom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getRoom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoomCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rooms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workingSlots\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roomsCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5033600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600190505b60058110156200011d57620000ce6200007e826200012460201b620005c81760201c565b60405160200162000090919062000593565b604051602081830303815290604052600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166200029e60201b60201c565b60008081819054906101000a900460ff1680929190620000ee90620005f5565b91906101000a81548160ff021916908360ff16021790555050808062000114906200062e565b9150506200005a565b5062000917565b606060008214156200016e576040518060400160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815250905062000299565b600082905060005b60008214620001a65780806200018c906200062e565b915050600a826200019e9190620006ab565b915062000176565b60008167ffffffffffffffff811115620001c557620001c4620006e3565b5b6040519080825280601f01601f191660200182016040528015620001f85781602001600182028036833780820191505090505b5090505b60008514620002925760018262000214919062000712565b9150600a856200022591906200074d565b603062000233919062000785565b60f81b8183815181106200024c576200024b620007e2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856200028a9190620006ab565b9450620001fc565b8093505050505b919050565b60016040518060a00160405280620002bc84620003db60201b60201c565b81526020018481526020018373ffffffffffffffffffffffffffffffffffffffff16815260200160148152602001600a81525090806001815401808255809150506001900390600052602060002090600502016000909190919091506000820151816000015560208201518160010190805190602001906200034092919062000410565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015560808201518160040155505060008081819054906101000a900460ff1680929190620003be90620005f5565b91906101000a81548160ff021916908360ff160217905550505050565b600081604051602001620003f0919062000895565b6040516020818303038152906040528051906020012060001c9050919050565b8280546200041e90620008e1565b90600052602060002090601f0160209004810192826200044257600085556200048e565b82601f106200045d57805160ff19168380011785556200048e565b828001600101855582156200048e579182015b828111156200048d57825182559160200191906001019062000470565b5b5090506200049d9190620004a1565b5090565b5b80821115620004bc576000816000905550600101620004a2565b5090565b600081905092915050565b7f726f6f6d00000000000000000000000000000000000000000000000000000000600082015250565b600062000503600483620004c0565b91506200051082620004cb565b600482019050919050565b600081519050919050565b60005b838110156200054657808201518184015260208101905062000529565b8381111562000556576000848401525b50505050565b600062000569826200051b565b620005758185620004c0565b93506200058781856020860162000526565b80840191505092915050565b6000620005a082620004f4565b9150620005ae82846200055c565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff82169050919050565b60006200060282620005e8565b915060ff821415620006195762000618620005b9565b5b600182019050919050565b6000819050919050565b60006200063b8262000624565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620006715762000670620005b9565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000620006b88262000624565b9150620006c58362000624565b925082620006d857620006d76200067c565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006200071f8262000624565b91506200072c8362000624565b925082821015620007425762000741620005b9565b5b828203905092915050565b60006200075a8262000624565b9150620007678362000624565b9250826200077a57620007796200067c565b5b828206905092915050565b6000620007928262000624565b91506200079f8362000624565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620007d757620007d6620005b9565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200083e8262000811565b9050919050565b60008160601b9050919050565b60006200085f8262000845565b9050919050565b6000620008738262000852565b9050919050565b6200088f620008898262000831565b62000866565b82525050565b6000620008a382846200087a565b60148201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620008fa57607f821691505b60208210811415620009115762000910620008b2565b5b50919050565b6110a080620009276000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063183f5f95146100675780631bae0ac81461008557806341b82f3f146100b957806348d68a2d146100d75780636d8a74cb146100f3578063eadddb3c14610124575b600080fd5b61006f610142565b60405161007c9190610a2a565b60405180910390f35b61009f600480360381019061009a9190610a8c565b6102a7565b6040516100b0959493929190610b21565b60405180910390f35b6100c1610395565b6040516100ce9190610b97565b60405180910390f35b6100f160048036038101906100ec9190610d13565b6103a6565b005b61010d60048036038101906101089190610a8c565b6104d7565b60405161011b929190610d6f565b60405180910390f35b61012c6105bb565b6040516101399190610d9f565b60405180910390f35b60606001805480602002602001604051908101604052809291908181526020016000905b8282101561029e57838290600052602060002090600502016040518060a0016040529081600082015481526020016001820180546101a390610de9565b80601f01602080910402602001604051908101604052809291908181526020018280546101cf90610de9565b801561021c5780601f106101f15761010080835404028352916020019161021c565b820191906000526020600020905b8154815290600101906020018083116101ff57829003601f168201915b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152505081526020019060010190610166565b50505050905090565b600181815481106102b757600080fd5b90600052602060002090600502016000915090508060000154908060010180546102e090610de9565b80601f016020809104026020016040519081016040528092919081815260200182805461030c90610de9565b80156103595780601f1061032e57610100808354040283529160200191610359565b820191906000526020600020905b81548152906001019060200180831161033c57829003601f168201915b5050505050908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154905085565b60008054906101000a900460ff1681565b60016040518060a001604052806103bc84610729565b81526020018481526020018373ffffffffffffffffffffffffffffffffffffffff16815260200160148152602001600a815250908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000155602082015181600101908051906020019061043e92919061075c565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015560808201518160040155505060008081819054906101000a900460ff16809291906104ba90610e4a565b91906101000a81548160ff021916908360ff160217905550505050565b60606000600183815481106104ef576104ee610e74565b5b90600052602060002090600502016001016001848154811061051457610513610e74565b5b90600052602060002090600502016000015481805461053290610de9565b80601f016020809104026020016040519081016040528092919081815260200182805461055e90610de9565b80156105ab5780601f10610580576101008083540402835291602001916105ab565b820191906000526020600020905b81548152906001019060200180831161058e57829003601f168201915b5050505050915091509150915091565b6000600180549050905090565b60606000821415610610576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610724565b600082905060005b6000821461064257808061062b90610ea3565b915050600a8261063b9190610f1b565b9150610618565b60008167ffffffffffffffff81111561065e5761065d610bbc565b5b6040519080825280601f01601f1916602001820160405280156106905781602001600182028036833780820191505090505b5090505b6000851461071d576001826106a99190610f4c565b9150600a856106b89190610f80565b60306106c49190610fb1565b60f81b8183815181106106da576106d9610e74565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856107169190610f1b565b9450610694565b8093505050505b919050565b60008160405160200161073c919061104f565b6040516020818303038152906040528051906020012060001c9050919050565b82805461076890610de9565b90600052602060002090601f01602090048101928261078a57600085556107d1565b82601f106107a357805160ff19168380011785556107d1565b828001600101855582156107d1579182015b828111156107d05782518255916020019190600101906107b5565b5b5090506107de91906107e2565b5090565b5b808211156107fb5760008160009055506001016107e3565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b61083e8161082b565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561087e578082015181840152602081019050610863565b8381111561088d576000848401525b50505050565b6000601f19601f8301169050919050565b60006108af82610844565b6108b9818561084f565b93506108c9818560208601610860565b6108d281610893565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610908826108dd565b9050919050565b610918816108fd565b82525050565b600060a0830160008301516109366000860182610835565b506020830151848203602086015261094e82826108a4565b9150506040830151610963604086018261090f565b5060608301516109766060860182610835565b5060808301516109896080860182610835565b508091505092915050565b60006109a0838361091e565b905092915050565b6000602082019050919050565b60006109c0826107ff565b6109ca818561080a565b9350836020820285016109dc8561081b565b8060005b85811015610a1857848403895281516109f98582610994565b9450610a04836109a8565b925060208a019950506001810190506109e0565b50829750879550505050505092915050565b60006020820190508181036000830152610a4481846109b5565b905092915050565b6000604051905090565b600080fd5b600080fd5b610a698161082b565b8114610a7457600080fd5b50565b600081359050610a8681610a60565b92915050565b600060208284031215610aa257610aa1610a56565b5b6000610ab084828501610a77565b91505092915050565b610ac28161082b565b82525050565b600082825260208201905092915050565b6000610ae482610844565b610aee8185610ac8565b9350610afe818560208601610860565b610b0781610893565b840191505092915050565b610b1b816108fd565b82525050565b600060a082019050610b366000830188610ab9565b8181036020830152610b488187610ad9565b9050610b576040830186610b12565b610b646060830185610ab9565b610b716080830184610ab9565b9695505050505050565b600060ff82169050919050565b610b9181610b7b565b82525050565b6000602082019050610bac6000830184610b88565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610bf482610893565b810181811067ffffffffffffffff82111715610c1357610c12610bbc565b5b80604052505050565b6000610c26610a4c565b9050610c328282610beb565b919050565b600067ffffffffffffffff821115610c5257610c51610bbc565b5b610c5b82610893565b9050602081019050919050565b82818337600083830152505050565b6000610c8a610c8584610c37565b610c1c565b905082815260208101848484011115610ca657610ca5610bb7565b5b610cb1848285610c68565b509392505050565b600082601f830112610cce57610ccd610bb2565b5b8135610cde848260208601610c77565b91505092915050565b610cf0816108fd565b8114610cfb57600080fd5b50565b600081359050610d0d81610ce7565b92915050565b60008060408385031215610d2a57610d29610a56565b5b600083013567ffffffffffffffff811115610d4857610d47610a5b565b5b610d5485828601610cb9565b9250506020610d6585828601610cfe565b9150509250929050565b60006040820190508181036000830152610d898185610ad9565b9050610d986020830184610ab9565b9392505050565b6000602082019050610db46000830184610ab9565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e0157607f821691505b60208210811415610e1557610e14610dba565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e5582610b7b565b915060ff821415610e6957610e68610e1b565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610eae8261082b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610ee157610ee0610e1b565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610f268261082b565b9150610f318361082b565b925082610f4157610f40610eec565b5b828204905092915050565b6000610f578261082b565b9150610f628361082b565b925082821015610f7557610f74610e1b565b5b828203905092915050565b6000610f8b8261082b565b9150610f968361082b565b925082610fa657610fa5610eec565b5b828206905092915050565b6000610fbc8261082b565b9150610fc78361082b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610ffc57610ffb610e1b565b5b828201905092915050565b60008160601b9050919050565b600061101f82611007565b9050919050565b600061103182611014565b9050919050565b611049611044826108fd565b611026565b82525050565b600061105b8284611038565b6014820191508190509291505056fea2646970667358221220134f4bd3ccada08a8b1081a5323cbf5187ae910e658096898ee278025606276f64736f6c634300080b0033",
}

// BookingsystemABI is the input ABI used to generate the binding from.
// Deprecated: Use BookingsystemMetaData.ABI instead.
var BookingsystemABI = BookingsystemMetaData.ABI

// BookingsystemBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BookingsystemMetaData.Bin instead.
var BookingsystemBin = BookingsystemMetaData.Bin

// DeployBookingsystem deploys a new Ethereum contract, binding an instance of Bookingsystem to it.
func DeployBookingsystem(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bookingsystem, error) {
	parsed, err := BookingsystemMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BookingsystemBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bookingsystem{BookingsystemCaller: BookingsystemCaller{contract: contract}, BookingsystemTransactor: BookingsystemTransactor{contract: contract}, BookingsystemFilterer: BookingsystemFilterer{contract: contract}}, nil
}

// Bookingsystem is an auto generated Go binding around an Ethereum contract.
type Bookingsystem struct {
	BookingsystemCaller     // Read-only binding to the contract
	BookingsystemTransactor // Write-only binding to the contract
	BookingsystemFilterer   // Log filterer for contract events
}

// BookingsystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type BookingsystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingsystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BookingsystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingsystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BookingsystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingsystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BookingsystemSession struct {
	Contract     *Bookingsystem    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BookingsystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BookingsystemCallerSession struct {
	Contract *BookingsystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BookingsystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BookingsystemTransactorSession struct {
	Contract     *BookingsystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BookingsystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type BookingsystemRaw struct {
	Contract *Bookingsystem // Generic contract binding to access the raw methods on
}

// BookingsystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BookingsystemCallerRaw struct {
	Contract *BookingsystemCaller // Generic read-only contract binding to access the raw methods on
}

// BookingsystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BookingsystemTransactorRaw struct {
	Contract *BookingsystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBookingsystem creates a new instance of Bookingsystem, bound to a specific deployed contract.
func NewBookingsystem(address common.Address, backend bind.ContractBackend) (*Bookingsystem, error) {
	contract, err := bindBookingsystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bookingsystem{BookingsystemCaller: BookingsystemCaller{contract: contract}, BookingsystemTransactor: BookingsystemTransactor{contract: contract}, BookingsystemFilterer: BookingsystemFilterer{contract: contract}}, nil
}

// NewBookingsystemCaller creates a new read-only instance of Bookingsystem, bound to a specific deployed contract.
func NewBookingsystemCaller(address common.Address, caller bind.ContractCaller) (*BookingsystemCaller, error) {
	contract, err := bindBookingsystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BookingsystemCaller{contract: contract}, nil
}

// NewBookingsystemTransactor creates a new write-only instance of Bookingsystem, bound to a specific deployed contract.
func NewBookingsystemTransactor(address common.Address, transactor bind.ContractTransactor) (*BookingsystemTransactor, error) {
	contract, err := bindBookingsystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BookingsystemTransactor{contract: contract}, nil
}

// NewBookingsystemFilterer creates a new log filterer instance of Bookingsystem, bound to a specific deployed contract.
func NewBookingsystemFilterer(address common.Address, filterer bind.ContractFilterer) (*BookingsystemFilterer, error) {
	contract, err := bindBookingsystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BookingsystemFilterer{contract: contract}, nil
}

// bindBookingsystem binds a generic wrapper to an already deployed contract.
func bindBookingsystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BookingsystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingsystem *BookingsystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingsystem.Contract.BookingsystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingsystem *BookingsystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingsystem.Contract.BookingsystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingsystem *BookingsystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingsystem.Contract.BookingsystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingsystem *BookingsystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingsystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingsystem *BookingsystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingsystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingsystem *BookingsystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingsystem.Contract.contract.Transact(opts, method, params...)
}

// AllRoomsByDate is a free data retrieval call binding the contract method 0x183f5f95.
//
// Solidity: function allRoomsByDate() view returns((uint256,string,address,uint256,uint256)[])
func (_Bookingsystem *BookingsystemCaller) AllRoomsByDate(opts *bind.CallOpts) ([]BookRoomRoom, error) {
	var out []interface{}
	err := _Bookingsystem.contract.Call(opts, &out, "allRoomsByDate")

	if err != nil {
		return *new([]BookRoomRoom), err
	}

	out0 := *abi.ConvertType(out[0], new([]BookRoomRoom)).(*[]BookRoomRoom)

	return out0, err

}

// AllRoomsByDate is a free data retrieval call binding the contract method 0x183f5f95.
//
// Solidity: function allRoomsByDate() view returns((uint256,string,address,uint256,uint256)[])
func (_Bookingsystem *BookingsystemSession) AllRoomsByDate() ([]BookRoomRoom, error) {
	return _Bookingsystem.Contract.AllRoomsByDate(&_Bookingsystem.CallOpts)
}

// AllRoomsByDate is a free data retrieval call binding the contract method 0x183f5f95.
//
// Solidity: function allRoomsByDate() view returns((uint256,string,address,uint256,uint256)[])
func (_Bookingsystem *BookingsystemCallerSession) AllRoomsByDate() ([]BookRoomRoom, error) {
	return _Bookingsystem.Contract.AllRoomsByDate(&_Bookingsystem.CallOpts)
}

// GetRoom is a free data retrieval call binding the contract method 0x6d8a74cb.
//
// Solidity: function getRoom(uint256 idx) view returns(string name, uint256 id)
func (_Bookingsystem *BookingsystemCaller) GetRoom(opts *bind.CallOpts, idx *big.Int) (struct {
	Name string
	Id   *big.Int
}, error) {
	var out []interface{}
	err := _Bookingsystem.contract.Call(opts, &out, "getRoom", idx)

	outstruct := new(struct {
		Name string
		Id   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoom is a free data retrieval call binding the contract method 0x6d8a74cb.
//
// Solidity: function getRoom(uint256 idx) view returns(string name, uint256 id)
func (_Bookingsystem *BookingsystemSession) GetRoom(idx *big.Int) (struct {
	Name string
	Id   *big.Int
}, error) {
	return _Bookingsystem.Contract.GetRoom(&_Bookingsystem.CallOpts, idx)
}

// GetRoom is a free data retrieval call binding the contract method 0x6d8a74cb.
//
// Solidity: function getRoom(uint256 idx) view returns(string name, uint256 id)
func (_Bookingsystem *BookingsystemCallerSession) GetRoom(idx *big.Int) (struct {
	Name string
	Id   *big.Int
}, error) {
	return _Bookingsystem.Contract.GetRoom(&_Bookingsystem.CallOpts, idx)
}

// GetRoomCount is a free data retrieval call binding the contract method 0xeadddb3c.
//
// Solidity: function getRoomCount() view returns(uint256)
func (_Bookingsystem *BookingsystemCaller) GetRoomCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bookingsystem.contract.Call(opts, &out, "getRoomCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoomCount is a free data retrieval call binding the contract method 0xeadddb3c.
//
// Solidity: function getRoomCount() view returns(uint256)
func (_Bookingsystem *BookingsystemSession) GetRoomCount() (*big.Int, error) {
	return _Bookingsystem.Contract.GetRoomCount(&_Bookingsystem.CallOpts)
}

// GetRoomCount is a free data retrieval call binding the contract method 0xeadddb3c.
//
// Solidity: function getRoomCount() view returns(uint256)
func (_Bookingsystem *BookingsystemCallerSession) GetRoomCount() (*big.Int, error) {
	return _Bookingsystem.Contract.GetRoomCount(&_Bookingsystem.CallOpts)
}

// Rooms is a free data retrieval call binding the contract method 0x1bae0ac8.
//
// Solidity: function rooms(uint256 ) view returns(uint256 id, string name, address owner, uint256 capacity, uint256 workingSlots)
func (_Bookingsystem *BookingsystemCaller) Rooms(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           *big.Int
	Name         string
	Owner        common.Address
	Capacity     *big.Int
	WorkingSlots *big.Int
}, error) {
	var out []interface{}
	err := _Bookingsystem.contract.Call(opts, &out, "rooms", arg0)

	outstruct := new(struct {
		Id           *big.Int
		Name         string
		Owner        common.Address
		Capacity     *big.Int
		WorkingSlots *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Capacity = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.WorkingSlots = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rooms is a free data retrieval call binding the contract method 0x1bae0ac8.
//
// Solidity: function rooms(uint256 ) view returns(uint256 id, string name, address owner, uint256 capacity, uint256 workingSlots)
func (_Bookingsystem *BookingsystemSession) Rooms(arg0 *big.Int) (struct {
	Id           *big.Int
	Name         string
	Owner        common.Address
	Capacity     *big.Int
	WorkingSlots *big.Int
}, error) {
	return _Bookingsystem.Contract.Rooms(&_Bookingsystem.CallOpts, arg0)
}

// Rooms is a free data retrieval call binding the contract method 0x1bae0ac8.
//
// Solidity: function rooms(uint256 ) view returns(uint256 id, string name, address owner, uint256 capacity, uint256 workingSlots)
func (_Bookingsystem *BookingsystemCallerSession) Rooms(arg0 *big.Int) (struct {
	Id           *big.Int
	Name         string
	Owner        common.Address
	Capacity     *big.Int
	WorkingSlots *big.Int
}, error) {
	return _Bookingsystem.Contract.Rooms(&_Bookingsystem.CallOpts, arg0)
}

// RoomsCount is a free data retrieval call binding the contract method 0x41b82f3f.
//
// Solidity: function roomsCount() view returns(uint8)
func (_Bookingsystem *BookingsystemCaller) RoomsCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bookingsystem.contract.Call(opts, &out, "roomsCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RoomsCount is a free data retrieval call binding the contract method 0x41b82f3f.
//
// Solidity: function roomsCount() view returns(uint8)
func (_Bookingsystem *BookingsystemSession) RoomsCount() (uint8, error) {
	return _Bookingsystem.Contract.RoomsCount(&_Bookingsystem.CallOpts)
}

// RoomsCount is a free data retrieval call binding the contract method 0x41b82f3f.
//
// Solidity: function roomsCount() view returns(uint8)
func (_Bookingsystem *BookingsystemCallerSession) RoomsCount() (uint8, error) {
	return _Bookingsystem.Contract.RoomsCount(&_Bookingsystem.CallOpts)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x48d68a2d.
//
// Solidity: function createRoom(string _name, address _owner) returns()
func (_Bookingsystem *BookingsystemTransactor) CreateRoom(opts *bind.TransactOpts, _name string, _owner common.Address) (*types.Transaction, error) {
	return _Bookingsystem.contract.Transact(opts, "createRoom", _name, _owner)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x48d68a2d.
//
// Solidity: function createRoom(string _name, address _owner) returns()
func (_Bookingsystem *BookingsystemSession) CreateRoom(_name string, _owner common.Address) (*types.Transaction, error) {
	return _Bookingsystem.Contract.CreateRoom(&_Bookingsystem.TransactOpts, _name, _owner)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x48d68a2d.
//
// Solidity: function createRoom(string _name, address _owner) returns()
func (_Bookingsystem *BookingsystemTransactorSession) CreateRoom(_name string, _owner common.Address) (*types.Transaction, error) {
	return _Bookingsystem.Contract.CreateRoom(&_Bookingsystem.TransactOpts, _name, _owner)
}
