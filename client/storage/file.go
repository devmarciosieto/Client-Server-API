package storage

import (
	"fmt"
	"github.com/devmarciosieto/Client-Server-API/client/internal/dto"
	"os"
	"time"
)

func CreateFile(response *dto.USDBRLRequest) {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
	}
	defer file.Close()

	currentDateTimeUTC := time.Now().UTC()
	formatted := currentDateTimeUTC.Format("2006-01-02 15:04:05")

	_, err = file.WriteString(fmt.Sprintf("Data e Hora da contação em UTC %s Valor do Dólar do bid: $%s\n", formatted, response.Bid))
}
