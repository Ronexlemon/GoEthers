package main

import ("github.com/ethereum/go-ethereum/ethclient")

var RPC_URL = "https://alfajores-forno.celo-testnet.org"
func main(){
	// Create a new client
	client, err := ethclient.Dial(RPC_URL)
	
}