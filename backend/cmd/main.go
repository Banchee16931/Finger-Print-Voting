package main

import (
	"finger-print-voting-backend/internal/api"
	"finger-print-voting-backend/internal/config"
	"finger-print-voting-backend/internal/database"
	"log"
)

func main() {
	log.Println("Loading Config")
	cfg := config.Load()

	log.Println("Connecting to Database")
	db, err := database.NewDatabase(cfg.DB)
	if err != nil {
		panic(err)
	}

	if err := db.EnsureValidSchema(); err != nil {
		panic(err)
	}

	defer db.Close()

	log.Println("Setting up Server")
	srv := api.NewServer().WithDBClient(db).WithPasswordSecret(cfg.PasswordSecret)

	err = srv.Start(":8080")
	if err != nil {
		panic(err.Error())
	}
}
