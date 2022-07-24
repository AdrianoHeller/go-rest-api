package services

import (
	"api/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"math/rand"
	"time"
)

func ListUsers(conn *pgx.Conn) error {
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
func ListSingleUser(conn *pgx.Conn, patternToMatch string, valueToMatch string) error {
	selectSingleQuery := fmt.Sprintf("select * from users where %s = '%s'", patternToMatch, valueToMatch)
	rows, _ := conn.Query(context.Background(), selectSingleQuery)
	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		fmt.Printf("username: %s, id: %s", name, id)
	}
	return rows.Err()
}

func InsertUser(conn *pgx.Conn, userData *models.User, tableName string) error {

	formattedQuery := fmt.Sprintf("insert into %s (id,name)values('%s','%s')", tableName, userData.Id, userData.Name)
	_, err := conn.Exec(context.Background(), formattedQuery)
	if err != nil {
		return err
	}
	return nil
}

func UuidGenerator() (string, error) {
	rand.Seed(time.Now().UnixNano())
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
