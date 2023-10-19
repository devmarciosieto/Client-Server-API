package service

import (
	"encoding/json"
	"github.com/devmarciosieto/Client-Server-API/server/internal/dto"
	"io"
	"net/http"
)

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
