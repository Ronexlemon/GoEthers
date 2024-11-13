package routes

import (
	"github.com/RonexLemon/Goether/controller"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)


func GetRoutes( router *mux.Router,client *ethclient.Client){
	router.HandleFunc("/balance/{address}",controller.GetUserBalance(client)).Methods("GET")
	router.HandleFunc("/blocknumber",controller.GetBlockNumber(client)).Methods("GET")
}