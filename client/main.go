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

const addr = "0xAa3777F59260b8bD003e850E321AdBc576115b06"
const addr2 = "0xD890c3FC59FCBddf5Ce62aC9AFfC90DEbb7C88DE"
const infuraUrl = "https://ropsten.infura.io/v3/7db961138c114b7882db2ff9788cded0"
const ganacheUrl = "http://localhost:7545"

func main() {

	// Create Ethereum Client
	client, err := ethclient.DialContext(context.Background(), ganacheUrl)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	fmt.Println("Block Number:", block.Number())

	// Convert string into Address
	address := common.HexToAddress(addr2)

	weiBalance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// Balance Value in Wei not ETH! 1 ETH = 10^18 Wei
	fmt.Println("BAlance:", weiBalance, " WEI") // Have to convert from *big.Int > *big.Float
	fBalance := new(big.Float)
	fBalance.SetString(weiBalance.String())

	// Divide wei by 10^18  to get ETH Balance
	ethBalance := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("BAlance:", ethBalance, " ETH")

}
