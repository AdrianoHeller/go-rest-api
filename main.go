package main

import (
	"api/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
	"os"
	"strings"
)

func main() {

	ctx := context.Background()

	postgresConnectionString := os.Getenv("PSQL_CONN")

	Port := os.Getenv("SERVER_PORT")

	if !strings.Contains(postgresConnectionString, "postgresql://") {
		panic("you must provide a connection string")
	}

	conn, err := pgx.Connect(ctx, postgresConnectionString)

	if err != nil {
		fmt.Printf("Error found while connecting: %s", err.Error())
		os.Exit(1)
	}

	if err := conn.Close(ctx); err != nil {
		fmt.Printf("Error found while cclosing connection: %s", err.Error())
		os.Exit(1)
	}

	if Port == "" {
		Port = "5000"
	}

	mux := http.NewServeMux()

	server := models.Server{
		Port: Port,
		Env:  models.Staging,
	}

	server.ServerRoute(mux)
}
