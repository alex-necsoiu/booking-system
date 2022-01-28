package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func balance() {

	// Convert string into Address
	address := common.HexToAddress(rinkebyAddress)
	// cAddress := common.HexToAddress(contractAddres)

	// Create Ethereum Client
	client, err := ethclient.DialContext(context.Background(), infuraUrl)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	defer client.Close()

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	fmt.Println("     Address:", address)
	fmt.Println("   Gas price:", gasPrice)
	fmt.Println("     ChainId:", chainID)
	fmt.Println("Block Number:", block.Number())

	weiBalance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fBalance := new(big.Float)
	fBalance.SetString(weiBalance.String())

	// Divide wei by 10^18  to get ETH Balance
	ethBalance := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("\n ETH Balance:", ethBalance, " ETH")
	// Balance Value in Wei not ETH! 1 ETH = 10^18 Wei
	fmt.Println(" WEY Balance:", weiBalance, "  WEI\n") // Have to convert from *big.Int > *big.Float

}
