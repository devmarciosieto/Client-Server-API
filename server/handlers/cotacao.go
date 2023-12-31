package handlers

import (
	"context"
	"encoding/json"
	"github.com/devmarciosieto/Client-Server-API/server/internal/dto"
	"github.com/devmarciosieto/Client-Server-API/server/service"
	"log"
	"net/http"
	"time"
)

func BuscaCotacaoUSDBRL(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	log.Println("Requeste iniciada")
	defer log.Println("Request finalizada")

	response, err := service.BuscaCotacao(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout de 200ms atingido ao buscar cotação")
			w.WriteHeader(http.StatusRequestTimeout)
			return
		}
	default:
	}

	if response == nil {
		log.Println("Resposta inesperada ao buscar cotação")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var res dto.USDBRLResponse
	res.Bid = response.USDBRL.Bid

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
