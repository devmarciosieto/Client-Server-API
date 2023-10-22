package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/devmarciosieto/Client-Server-API/client/internal/dto"
	"github.com/devmarciosieto/Client-Server-API/client/storage"
	"io"
	"net/http"
	"time"
)

func BuscaCotacao() (*dto.USDBRLRequest, error) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

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

	var request dto.USDBRLRequest
	err = json.Unmarshal(res, &request)
	if err != nil {
		return nil, err
	}

	storage.CreateFile(&request)

	return &request, nil
}
