package api

import (
	"bank-list-api/structs"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == http.MethodGet {
		return s.handleListAccount(w, r)
	}

	if r.Method == http.MethodPost {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == http.MethodDelete {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleAccountID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetAccount(w, r)
	}

	if r.Method == http.MethodDelete {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getIDFromRequest(r)
	if err != nil {
		return err
	}

	acc, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, acc)
}

func (s *APIServer) handleListAccount(w http.ResponseWriter, _ *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	accReq := new(structs.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(accReq); err != nil {
		return err
	}

	account := structs.NewAccount(accReq.FirstName, accReq.LastName)

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getIDFromRequest(r)
	if err != nil {
		return err
	}

	err = s.store.DeleteAccount(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func getIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("can't convert %s into a integer", idStr)
	}
	return id, nil
}
