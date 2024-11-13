package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)



var RPC_URL = "https://alfajores-forno.celo-testnet.org"
func main(){
	// Create a new client
	client, err := ethclient.Dial(RPC_URL)
	if err !=nil{
		log.Fatal(err)
	}
	defer  client.Close()
	fmt.Println("conected to CLient RPC:", RPC_URL)
	
}