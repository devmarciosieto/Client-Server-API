package main

import (
	"encoding/json"
	"fmt"
	"github.com/devmarciosieto/Client-Server-API/server/internal/dto"
	"io"
	"net/http"
	"os"
)

func BuscaCotacaoUSDBRL(w http.ResponseWriter, r *http.Request) {

	response, err := BuscaCotacao()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	createFile(err, response)
	var res dto.USDBRLResponse
	res.Bid = response.USDBRL.Bid

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func createFile(err error, response *dto.ApiResponse) {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", response.USDBRL.Bid))
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
