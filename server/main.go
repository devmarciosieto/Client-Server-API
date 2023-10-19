package main

import (
	"encoding/json"
	"fmt"
	"github.com/devmarciosieto/Client-Server-API/server/internal/dto"
	"github.com/devmarciosieto/Client-Server-API/server/storage"
	"io"
	"net/http"
)

func BuscaCotacaoUSDBRL(w http.ResponseWriter, r *http.Request) {

	response, err := BuscaCotacao()
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

func BuscaCotacao() (*dto.ApiResponse, error) {
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var response dto.ApiResponse
	err = json.Unmarshal(res, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func main() {
	http.HandleFunc("/cotacao", BuscaCotacaoUSDBRL)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
