package models

import (
	"fmt"
	"log"
	"net/http"
)

type RunningEnv int

const (
	Develompment = iota
	Staging
	Production
)

type Server struct {
	Port string     `json:"port"`
	Env  RunningEnv `json:"env"`
}

func (s *Server) ServerRoute(mux *http.ServeMux) {

	mux.Handle("/", PingRef)
	mux.Handle("/users/all", UserRef)
	mux.Handle("/users/wallet", WalletRef)

	err := http.ListenAndServe(fmt.Sprintf(":%s", s.Port), mux)
	if err != nil {
		msg := fmt.Sprintf("Error found: %s", err)
		log.Fatal(msg)
	}
}
