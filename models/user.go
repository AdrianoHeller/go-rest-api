package models

import (
	"api/helpers"
	"encoding/json"
	"net/http"
	"reflect"
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

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(&u)

	if err != nil {
		helpers.HandleCustomErrors(w, err)
		return
	}

	validId := reflect.TypeOf(u.Id) != nil

	validName := u.Name != "" && strings.Contains(u.Name, " ") && len(u.Name) < 50

	if !validId || !validName {
		msg := "invalid request fields"
		http.Error(w, msg, http.StatusNotAcceptable)
		return
	}

	jsonData, err := helpers.ConvertToJson(u)

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
