package api

import (
	"bank-list-api/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {

	if r.Method == http.MethodPost {
		return s.handleTransferAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {
	transferReq := structs.TransferRequest{}
	if err := json.NewDecoder(r.Body).Decode(&transferReq); err != nil {
		return err
	}

	defer r.Body.Close()

	return WriteJSON(w, http.StatusOK, transferReq)
}
