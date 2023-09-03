package main

import (
	"finger-print-voting-backend/internal/api"
	"finger-print-voting-backend/internal/background"
	"finger-print-voting-backend/internal/config"
	"finger-print-voting-backend/internal/database"
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Loading Config")
	cfg := config.Load()

	log.Println("Connecting to Database")
	db, err := database.NewDatabase(cfg.DB)
	if err != nil {
		panic(err)
	}

	codebaseLoc, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	schemaLoc := fmt.Sprintf("%s\\internal\\database\\schemas", codebaseLoc)

	if err := db.EnsureValidSchema(schemaLoc); err != nil {
		panic(err)
	}

	defer db.Close()

	go background.UpdateElections(db)

	log.Println("Setting up Server")
	srv := api.NewServer().WithDBClient(db).WithPasswordSecret(cfg.PasswordSecret)

	err = srv.Start(":8080")
	if err != nil {
		panic(err.Error())
	}
}
