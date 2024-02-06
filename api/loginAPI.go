package api

import (
	"bank-list-api/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	var req structs.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	acc, err := s.store.GetAccountByID(req.Identification)
	if err != nil {
		return err
	}

	if !acc.ValidatePassword(req.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := createJWTToken(acc)
	if err != nil {
		return err
	}

	resp := structs.LoginResponse{
		ID:    acc.ID,
		Token: token,
	}

	return WriteJSON(w, http.StatusOK, resp)
}
