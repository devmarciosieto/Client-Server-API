package handlers

import (
	"encoding/json"
	"github.com/devmarciosieto/Client-Server-API/server/internal/dto"
	"github.com/devmarciosieto/Client-Server-API/server/service"
	"github.com/devmarciosieto/Client-Server-API/server/storage"
	"net/http"
)

func BuscaCotacaoUSDBRL(w http.ResponseWriter, r *http.Request) {

	response, err := service.BuscaCotacao()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	storage.CreateFile(response)

	var res dto.USDBRLResponse
	res.Bid = response.USDBRL.Bid

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
