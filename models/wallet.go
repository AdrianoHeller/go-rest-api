package models

import (
	"api/helpers"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type TransStatus int

const (
	Processed = iota
	OnHold
	Cancelled
)

type Transaction struct {
	Id     uuid.UUID   `json:"id"`
	From   string      `json:"from"`
	To     string      `json:"to"`
	Amount float64     `json:"amount"`
	Status TransStatus `json:"st"`
}

type Wallet struct {
	Id           uuid.UUID     `json:"id"`
	Owner        string        `json:"owner"`
	CreatedAt    int           `json:"created_at"`
	Transactions []Transaction `json:"transactions"`
}

func (h *Wallet) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	r.Body = http.MaxBytesReader(w, r.Body, 1040000)

	var wlt Wallet

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(&wlt)

	if err != nil {
		helpers.HandleCustomErrors(w, err)
		return
	}

	data, err := helpers.ConvertToJson(wlt)

	w.WriteHeader(http.StatusOK)

	if _, e := w.Write(data); e != nil {
		msg := fmt.Sprintf("Error found: %s", e)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

var transactionTest Transaction = Transaction{
	Id:     uuid.New(),
	From:   "John Bonhan",
	To:     "Robert Plant",
	Amount: 298.72,
	Status: Processed,
}

var Tr []Transaction

var WalletRef = &Wallet{
	Id:           uuid.New(),
	Owner:        "",
	CreatedAt:    000000000000,
	Transactions: append(Tr, transactionTest),
}
