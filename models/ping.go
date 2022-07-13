package models

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
)

type Ping struct {
	Message    string     `json:"message"`
	StatusCode int        `json:"status_code"`
	Port       string     `json:"port"`
	Env        RunningEnv `json:"env"`
}

func (h *Ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	json, err := json2.Marshal(h)

	if err != nil {
		msg := fmt.Sprintf("Error found: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.StatusCode)
	_, e := w.Write(json)
	if e != nil {
		msg := fmt.Sprintf("Error found: %s", e)
		http.Error(w, msg, http.StatusBadRequest)
	}
}

var Port string = "3000"

var PingRef = &Ping{
	Port:       Port,
	Env:        Staging,
	Message:    "OK",
	StatusCode: 200,
}
