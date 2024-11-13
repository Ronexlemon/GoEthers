package routes

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)


func GetRoutes( router *mux.Router,client *ethclient.Client){
	router.HandleFunc("/{address}",).Method("GET")
}