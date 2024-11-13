package main

import (
	"fmt"
	

	"github.com/RonexLemon/Goether/config"
)



var RPC_URL = "https://alfajores-forno.celo-testnet.org"
func main(){
	// Create a new client
	clientInstance := &config.Client{
		RPC_URL: RPC_URL,
	}
	client := clientInstance.NewClient()
	
	defer  client.Close()
	fmt.Println("conected to CLient RPC:", RPC_URL)
	
}