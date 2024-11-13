package routes

import (
	

	
	"github.com/gorilla/mux"
)


func GetRoutes( router *mux.Router){
	router.HandleFunc("/{address}",).Method("GET")
}