package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"strconv"

	"github.com/RonexLemon/Goether/responses"
	"github.com/RonexLemon/Goether/services"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)


func IsContractOrAddress (client *ethclient.Client)http.HandlerFunc{
	return(func(w http.ResponseWriter, r *http.Request){
		params:=mux.Vars(r)
		address:=params["address"]
		ctx,cancel:= context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()
		isCont_add,result,err:=services.IsContractOrAddress(ctx,client,address)
		if err!=nil{
			response:= responses.MessageResponse{Status:http.StatusInternalServerError, Message:"error",Data:map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		response:= responses.MessageResponse{Status:http.StatusOK, Message:"Success",Data:map[string]interface{}{"Is":map[string]string{
			"Contract":strconv.FormatBool(isCont_add),
			"Address":result,
		}}}
		json.NewEncoder(w).Encode(response)

		


	})
}