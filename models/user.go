package models

import (
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if !strings.Contains(r.Header.Get("Authorization"), "Bearer ") {
		msg := "you must provide a token"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	if r.Method != "POST" {
		msg := "method not allowed"
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	var u User

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	validId := u.Id != "" && len(u.Id) > 0 && len(u.Id) < 25

	validName := u.Name != "" && strings.Contains(u.Name, " ") && len(u.Name) < 50

	if !validId && !validName {
		msg := "invalid request fields"
		http.Error(w, msg, http.StatusNotAcceptable)
		return
	}

	jsonData, err := json.Marshal(u)

	if err != nil {
		msg := "could not parse incoming payload"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, e := w.Write(jsonData); e != nil {
		msg := "could not write bytes into response"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

}

var UserRef = &User{
	Id:   "",
	Name: "",
}
