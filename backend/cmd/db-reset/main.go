package main

import (
	"finger-print-voting-backend/internal/config"
	"finger-print-voting-backend/internal/database"
	"finger-print-voting-backend/internal/types"
	"finger-print-voting-backend/internal/users"
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

	if err := db.DropDBTables(); err != nil {
		panic(err)
	}

	codebaseLoc, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	schemaLoc := fmt.Sprintf("%s\\internal\\database\\schemas", codebaseLoc)

	if err := db.SetupSchema(schemaLoc); err != nil {
		panic(err)
	}

	if err := users.NewUser(db, types.User{
		Username:  "admin",
		Password:  "firstplease",
		Admin:     true,
		FirstName: "Admin",
		LastName:  "Account",
	}); err != nil {
		panic(err)
	}

	log.Println("Created admin account")

	defer db.Close()
}
