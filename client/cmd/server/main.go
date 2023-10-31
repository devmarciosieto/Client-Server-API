package main

import (
	"fmt"
	"github.com/devmarciosieto/Client-Server-API/client/handlers"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("../../public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/cotacao", handlers.BuscaCotacaoUSDBRL)
	fmt.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
