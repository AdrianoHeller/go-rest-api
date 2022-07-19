package main

import (
	"api/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"net/http"
	"os"
	"strings"
)

var conn *pgx.Conn

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

	if err := listUsers(); err != nil {
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

func listUsers() error {
	selectQuery := "select * from users"
	rows, _ := conn.Query(context.Background(), selectQuery)
	for rows.Next() {
		var id uuid.UUID
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		fmt.Printf("id: %d, name: %s", id, name)
	}
	return rows.Err()
}

func insertUser(userData *models.User, tableName string) error {
	formattedQuery := fmt.Sprintf(`insert into %s ("id","name") values("%d","%s")`, tableName, userData.Id, userData.Name)
	_, err := conn.Exec(context.Background(), formattedQuery)
	if err != nil {
		return err
	}
	return nil
}
