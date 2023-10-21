package handlers

import (
	"encoding/json"
	"github.com/devmarciosieto/Client-Server-API/client/internal/dto"
	"github.com/devmarciosieto/Client-Server-API/client/service"
	"net/http"
)

func BuscaCotacaoUSDBRL(w http.ResponseWriter, r *http.Request) {
	request, err := service.BuscaCotacao()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var req dto.USDBRLRequest
	req.Bid = request.Bid

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(req)

}
