package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// online test web: https://developer.metamask.io/key/bc339063a9b54795a3097d2e9b7d413f/active-endpoints
var httpUrl = "https://sepolia.infura.io/v3/bc339063a9b54795a3097d2e9b7d413f"
var localUrl = "http://127.0.0.1:8545" // commond line input: ganache-cli

func main() {
	fmt.Println("hello")
	client, err := ethclient.DialContext(context.Background(), localUrl)
	if err != nil {
		log.Fatalf("error to create a eth client:%v", err)
	}
	defer client.Close()
	fmt.Println("get client success!")
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("error to get block:%v", err)
	}
	fmt.Println(block.Number())
	// section 2: https://www.youtube.com/watch?v=nuivtRUaSw8&list=PLay9kDOVd_x7hbhssw4pTKZHzzc6OG0e_&index=2
}
