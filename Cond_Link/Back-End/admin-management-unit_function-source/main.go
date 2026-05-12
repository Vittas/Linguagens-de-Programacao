package main

import (
	"log"
	"net/http"

	"functions.com/controlview-management/function"
)

func main() {
	http.HandleFunc("/", function.BootUp)

	log.Println("Servidor rodando em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
