package main

import (
	"github.com/devmarciosieto/Client-Server-API/client/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", handlers.BuscaCotacaoUSDBRL)
	log.Fatal(http.ListenAndServe(":8081", mux))
}
