package storage

import (
	"bank-list-api/structs"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(account *structs.Account) (int, error)
	DeleteAccount(int) error
	UpdateAccount(account *structs.Account) error
	GetAccountByID(int) (*structs.Account, error)
	GetAccounts() ([]*structs.Account, error)
}
