package main

import (
	"bank-list-api/api"
	"bank-list-api/storage"
	"log"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
