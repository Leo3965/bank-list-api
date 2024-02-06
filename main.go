package main

import (
	"bank-list-api/api"
	"bank-list-api/scripts"
	"bank-list-api/storage"
	"flag"
	"fmt"
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

	//seed stuff
	seed := flag.Bool("seed", false, "seeding a user to db")
	flag.Parse()

	if *seed {
		fmt.Println("seeding the database")
		scripts.SeedAccounts(store)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
