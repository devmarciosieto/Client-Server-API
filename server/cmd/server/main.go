package main

import (
	"fmt"
	"github.com/devmarciosieto/Client-Server-API/server/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/cotacao", handlers.BuscaCotacaoUSDBRL)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
