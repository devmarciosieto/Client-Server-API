package storage

import (
	"fmt"
	"github.com/devmarciosieto/Client-Server-API/server/internal/dto"
	"os"
)

func CreateFile(response *dto.ApiResponse) {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", response.USDBRL.Bid))
}
