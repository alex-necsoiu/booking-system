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

// var (
// 	// Dial is an alias for github.com/ethereum/go-ethereum/ethclient.Dial
// 	Dial = ethclient.Dial
// )

// // Client is go-ethereum.ethclient re-exported.
// type Client = ethclient.Client

// // MustDial calls Dial with an URL and returns an ethereum client.
// // On error, the error is logged an the software exits with fatal error message.
// func MustDial(url string) *ethclient.Client {
// 	client, err := ethclient.Dial(url)
// 	if err != nil {
// 		log.Fatal(
// 			"Ethereum client", errors.Wrap(err, "Unable to connect to node"),
// 			"provider", "blockchain",
// 			"chain", "ethereum",
// 		)
// 	}
// 	return client
// }
