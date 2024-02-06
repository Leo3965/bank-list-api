package scripts

import (
	"bank-list-api/service"
	"bank-list-api/storage"
	"bank-list-api/structs"
	"fmt"
	"log"
)

func seedAccount(store storage.Storage, fname, lname, pw string) *structs.Account {
	acc, err := service.NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	id, err := store.CreateAccount(acc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID of the new account", id)

	return acc
}

func SeedAccounts(store storage.Storage) {
	seedAccount(store, "Leonardo", "O. Freitas", "hunter")
}
