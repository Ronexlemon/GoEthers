package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/RonexLemon/Goether/responses"
	"github.com/RonexLemon/Goether/services"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func GetUserBalance(client *ethclient.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		params := mux.Vars(r)
		address := params["address"]

		balance, err := services.GetBalance(ctx, client, address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.MessageResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := responses.MessageResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"balance": balance}}
		json.NewEncoder(w).Encode(response)



		

	}
}

//get blockNumber
func GetBlockNumber(client *ethclient.Client)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		blocknumber,err := services.GetBlockNumber(ctx,client)
		if err !=nil{
			w.WriteHeader(http.StatusInternalServerError)
			response:= responses.MessageResponse{Status:http.StatusInternalServerError,Message:"error",Data:map[string]interface{}{"blocknumber":err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		response:= responses.MessageResponse{Status:http.StatusOK,Message:"Success",Data:map[string]interface{}{"blocknumber":blocknumber}}
		json.NewEncoder(w).Encode(response)

	}
}

func GetBlockDetails(client *ethclient.Client)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		block,err:= services.GetLatestBlockDetails(ctx,client);
		if err!=nil{
			response:=responses.MessageResponse{Status:http.StatusInternalServerError,Message:"error",Data:map[string]interface{}{"data":err.Error()}}
			json.NewEncoder(w).Encode(response)
			return}
			response:=responses.MessageResponse{Status:http.StatusOK,Message:"Sucess",Data:map[string]interface{}{"data":block}}
			json.NewEncoder(w).Encode(response)


	}
}

func TransactionByHash(client *ethclient.Client)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		params := mux.Vars(r)
		hash := params["hash"]
		transaction,isPending,err := services.GetTransactionByHash(ctx,client,hash)
		if err!=nil{
			response := responses.MessageResponse{Status:http.StatusInternalServerError,Message:"error",Data:map[string]interface{}{"Error":err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
			}
			if isPending{
				response := responses.MessageResponse{Status:http.StatusPartialContent,Message:"Pending",Data:map[string]interface{}{"transaction":isPending}}
			json.NewEncoder(w).Encode(response)
			return

			}
			response := responses.MessageResponse{Status:http.StatusOK,Message:"Success",Data:map[string]interface{}{"transaction":transaction}}
			json.NewEncoder(w).Encode(response)

	}
}