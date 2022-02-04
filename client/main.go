package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	contract "github.com/alex-necsoiu/booking-system/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Private Key
	pKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		log.Fatal(err)
	}

	// Convert string into Address
	address := common.HexToAddress(rinkebyAddress)
	cAddress := common.HexToAddress(contractAddres)

	// date := big.NewInt(1514764800)
	client, err := ethClient()
	if err != nil {
		log.Fatal(err)
	}

	instance, err := contract.NewBookingsystem(cAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	instance2, err := contract.NewBookingsystemTransactor(cAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Call CountRoom
	_, err = CountRoom(instance)
	if err != nil {
		log.Fatal(err)
	}

	txOpts, err := TxOpts(context.Background(), client, pKey)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance2.CreateRoom(txOpts, "Alex Room2", address)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * 5000000)

	// Call CountRoom
	_, err = CountRoom(instance)
	if err != nil {
		log.Fatal(err)
	}
	// Call GetAllRooms
	_, err = GetAllRooms(instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Room cerated Tx:", tx.Hash())
	fmt.Println("Room cerated Tx:", tx.Value())

	// // Instantiate the contract and display its name
	// token, err := NewToken(common.HexToAddress("0x21e6fc92f93c8a1bb41e2be64b4e1f88a54d3576"), conn)
	// if err != nil {
	// 	log.Fatalf("Failed to instantiate a Token contract: %v", err)
	// }
	// name, err := token.Name(nil)
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve token name: %v", err)
	// }
	// fmt.Println("Token name:", name)

}

// TxOpts take a client and private key and returns *bind.TransactOpts
func TxOpts(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	add := common.HexToAddress(rinkebyAddress)

	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	// auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// auth.GasPrice = gasPrice
	// auth.GasLimit = uint64(6000000)
	// auth.Nonce = big.NewInt(int64(nonce))

	tx := &bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: auth.GasLimit,
		Value:    big.NewInt(10),
		Nonce:    auth.Nonce,
		Context:  ctx,
		GasPrice: auth.GasPrice,
	}

	return tx, nil

}

// GetAllRooms returns an array of []Room
func GetAllRooms(instance *contract.Bookingsystem) ([]contract.BookRoomRoom, error) {
	rooms, err := instance.AllRoomsByDate(nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("\nRoom", rooms)
	return rooms, nil
}

// // GetRoom returns an Room by index
// func GetRoom(instance *contract.Bookingsystem, idx int64) (*struct{Name:string bigInt}, error) {
// 	room, err := instance.GetRoom(nil, big.NewInt(idx))
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("\nRoom:", room)

// 	return room, nil
// }

// CountRoom returns an Room
func CountRoom(instance *contract.Bookingsystem) (uint8, error) {
	count, err := instance.RoomsCount(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nTotal Rooms:", count)
	return count, nil
}
