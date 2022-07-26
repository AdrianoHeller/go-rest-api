package main

import (
	"api/helpers"
	"api/models"
	"api/services"
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

	fmt.Println("Connected to Postgres")

	guid, err := helpers.UuidGenerator()

	if err != nil {
		fmt.Printf("Error found while creating uuid: %s", err.Error())
		os.Exit(1)
	}

	newUser := models.User{
		Id:   guid,
		Name: "John Dog",
	}

	if err := services.InsertUser(conn, &newUser, "users"); err != nil {
		fmt.Printf("Error while listing users: %s", err.Error())
		os.Exit(1)
	}

	if err := services.ListUsers(conn); err != nil {
		fmt.Printf("Error while listing users: %s", err.Error())
		os.Exit(1)
	}

	if err := services.ListSingleUser(conn, "id", "ed723ce6-6b24-aaa6-d53f-cda68b1da87e"); err != nil {
		fmt.Printf("Error while listing users: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close(ctx)

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
