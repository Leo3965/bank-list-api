package service

import (
	"bank-list-api/structs"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func NewAccount(firstName, lastName, password string) (*structs.Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &structs.Account{
		FirstName:         firstName,
		LastName:          lastName,
		Number:            int64(rand.Intn(10000)),
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
