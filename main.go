package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RonexLemon/Goether/config"
	"github.com/RonexLemon/Goether/router"
	"github.com/RonexLemon/Goether/routes"
)



var RPC_URL = "https://alfajores-forno.celo-testnet.org"
var add = ":9090"
func main(){
	// Create a new client
	clientInstance := &config.Client{
		RPC_URL: RPC_URL,
	}
	client := clientInstance.NewClient()
	//router
	router:= router.NewRouter()
	routes.GetRoutes(router,client)
	routes.CreateWalletRoutes(router)
	routes.CheckAddressType(router,client)

	
	defer  client.Close()
	fmt.Println("conected to CLient RPC:", RPC_URL)
	fmt.Println("listening on port:", add)
	log.Fatal(http.ListenAndServe(add,router))
	
}