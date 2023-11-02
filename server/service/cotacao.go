package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/devmarciosieto/Client-Server-API/server/internal/domain/entity"
	"github.com/devmarciosieto/Client-Server-API/server/internal/dto"
	"github.com/devmarciosieto/Client-Server-API/server/internal/infra/database"
	"io"
	"log"
	"net/http"
	"time"
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

	db := database.NewDb()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	//ctxWithTimeout, cancel := context.WithTimeout(ctx, 1*time.Nanosecond)
	defer cancel()

	if ctxWithTimeout.Err() == context.DeadlineExceeded {
		log.Println("database context canceled")

	}

	repo := database.NewUSDBRLRepository(ctxWithTimeout, db)
	err = repo.InsertUSDBRL(entity.NewUSDBRL(response.USDBRL.Code, response.USDBRL.Codein, response.USDBRL.Name, response.USDBRL.High, response.USDBRL.Low, response.USDBRL.VarBid, response.USDBRL.PctChange, response.USDBRL.Bid, response.USDBRL.Ask, response.USDBRL.Timestamp, response.USDBRL.Create_date))

	if err != nil {
		return nil, err
	}

	return &response, nil
}
