package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/devmarciosieto/Client-Server-API/server/internal/dto"
	"io"
	"net/http"
)

func BuscaCotacao(ctx context.Context) (*dto.ApiResponse, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	res, err := io.ReadAll(resp.Body)
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
