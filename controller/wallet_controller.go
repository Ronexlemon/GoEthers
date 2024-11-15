package controller

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/RonexLemon/Goether/responses"
	"github.com/RonexLemon/Goether/services"
)

// CreateWallet handles wallet creation
func CreateWallet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		privateKey, publicKey,address, err := services.GeneratePrivateKey()
		privateKeyHex := hex.EncodeToString(privateKey)
		publicKeyHex := hex.EncodeToString(publicKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.MessageResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error generating keys",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.MessageResponse{
			Status:  http.StatusCreated,
			Message: "Success",
			Data: map[string]interface{}{
				"keys": map[string]string{
					"privateKey": string(privateKeyHex),
					"publicKey":  string(publicKeyHex),
					"address":    address,
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}
}
