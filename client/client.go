package main

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddres = "0x3607092459d8AF58be4CfdBae4059B90fBA6eE93"
	rinkebyAddress = "0xAa3777F59260b8bD003e850E321AdBc576115b06"
	ganacheUrl     = "http://localhost:7545"
	infuraUrl      = "https://rinkeby.infura.io/v3/7db961138c114b7882db2ff9788cded0"
	pKey           = "fd2179823adc4e772665848074ff54c3a0351a06accbb9e08498a51da1b519c5"
)

func ethClient() (*ethclient.Client, error) {

	// Convert string into Address
	// cAddress := common.HexToAddress(contractAddres)

	// Create Ethereum Client
	client, err := ethclient.DialContext(context.Background(), infuraUrl)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// instance, err := contract.NewBookingsystem(cAddress, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return client, nil
}
