package main

import (
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(account *Account) error
	DeleteAccount(int) error
	UpdateAccount(account *Account) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}
