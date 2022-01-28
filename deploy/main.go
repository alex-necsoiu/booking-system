package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	contract "github.com/alex-necsoiu/booking-system/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ganacheUrl     = "http://localhost:7545"
	addr2          = "0xD890c3FC59FCBddf5Ce62aC9AFfC90DEbb7C88DE"
	rinkebyAddress = "0xAa3777F59260b8bD003e850E321AdBc576115b06"
	infuraUrl      = "https://rinkeby.infura.io/v3/7db961138c114b7882db2ff9788cded0"
	pKey           = "fd2179823adc4e772665848074ff54c3a0351a06accbb9e08498a51da1b519c5"
)

func main() {

	// Private Key
	privateKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		log.Fatal(err)
	}

	// Client
	client, err := ethclient.Dial(infuraUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	add := common.HexToAddress(rinkebyAddress)

	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(8000000)
	auth.Nonce = big.NewInt(int64(nonce))

	a, tx, _, err := contract.DeployBookingsystem(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------------------------")
	fmt.Println("CONTRACT ADDRESS:", a.Hex())
	fmt.Println("TRANSACTION HASH:", tx.Hash().Hex())
	fmt.Println("-----------------------------------")
}
