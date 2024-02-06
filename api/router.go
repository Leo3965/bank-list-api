package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandleFunc(s.handleAccountID), s.store))

	router.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer))

	log.Println("API Server running on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
