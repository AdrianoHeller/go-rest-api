package main

import (
	"api/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

//var conn *pgx.Conn

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

	guid, err := uuidGenerator()

	if err != nil {
		fmt.Printf("Error found while creating uuid: %s", err.Error())
		os.Exit(1)
	}

	newUser := models.User{
		Id:   guid,
		Name: "John Dog",
	}

	if err := insertUser(conn, &newUser, "users"); err != nil {
		fmt.Printf("Error while listing users: %s", err.Error())
		os.Exit(1)
	}

	if err := listUsers(conn); err != nil {
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

func listUsers(conn *pgx.Conn) error {
	selectQuery := "select * from users"
	rows, _ := conn.Query(context.Background(), selectQuery)
	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		fmt.Printf("id: %s, name: %s", id, name)
	}
	return rows.Err()
}

func insertUser(conn *pgx.Conn, userData *models.User, tableName string) error {

	formattedQuery := fmt.Sprintf("insert into %s (id,name)values('%s','%s')", tableName, userData.Id, userData.Name)
	_, err := conn.Exec(context.Background(), formattedQuery)
	if err != nil {
		return err
	}
	return nil
}

func uuidGenerator() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	customUuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
	return customUuid, nil
}
