package models

import (
	"api/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
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
	Id           string        `json:"id"`
	Owner        string        `json:"owner"`
	CreatedAt    time.Time     `json:"created_at"`
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

	data, err := helpers.ConvertToJson(r.Body)

	w.WriteHeader(http.StatusOK)
	_, e := w.Write(data)

	if e != nil {
		msg := fmt.Sprintf("Error found: %s", e)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}
