package services

import (
	"api/helpers"
	"api/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
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

func CreateWallet(conn *pgx.Conn, owner string) error {
	uuid, err := helpers.UuidGenerator()
	if err != nil {
		errMsg := fmt.Sprintf("Error found: %s", err)
		return errors.New(errMsg)
	}
	newWallet := models.Wallet{
		Id:        uuid,
		Owner:     owner,
		CreatedAt: int(time.Now().Unix()),
	}
	insertNewQuery := fmt.Sprintf("insert into wallets (id,owner,created_at) values('%s','%s','%d')", newWallet.Id, newWallet.Owner, newWallet.CreatedAt)
	if _, err := conn.Exec(context.Background(), insertNewQuery); err != nil {
		errMsg := fmt.Sprintf("Error found: %s", err)
		return errors.New(errMsg)
	}
	return nil
}

func CreateTransaction(conn *pgx.Conn, from string, to string, amount float64) error {
	uuid, err := helpers.UuidGenerator()
	if err != nil {
		errMsg := fmt.Sprintf("Error found: %s", err)
		return errors.New(errMsg)
	}
	newTransaction := models.Transaction{
		Id:     uuid,
		From:   from,
		To:     to,
		Amount: amount,
		Status: models.Processed,
	}
	insertNewQuery := fmt.Sprintf("insert into transactions (id,from,to,amount,status) values('%s','%s','%s','%f','%d')", newTransaction.Id, newTransaction.From, newTransaction.To, newTransaction.Amount, newTransaction.Status)
	if _, err := conn.Exec(context.Background(), insertNewQuery); err != nil {
		errMsg := fmt.Sprintf("Error found: %s", err)
		return errors.New(errMsg)
	}
	return nil
}
