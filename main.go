package main

import (
	"fmt"
	"log"
	"net/http"
	"open-banking-expenses/config"
	"open-banking-expenses/routes"
)

func main() {
	config.ConnectDatabase()

	r := routes.SetupRouter()

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
