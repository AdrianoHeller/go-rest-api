package main

import (
	"api/models"
	"net/http"
	"os"
	"strings"
)

func main() {

	//ctx := context.Background()

	Port := os.Getenv("SERVER_PORT")

	postrgesConnectionString := os.Getenv("PSQL_CONN")

	if Port == "" {
		Port = "5000"
	}

	if !strings.Contains(postrgesConnectionString, "postgresql://") {
		panic("you must provide a connection string")
	}

	mux := http.NewServeMux()

	server := models.Server{
		Port: Port,
		Env:  models.Staging,
	}

	server.ServerRoute(mux)
}
