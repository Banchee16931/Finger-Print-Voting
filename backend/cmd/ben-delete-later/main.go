package main

import (
	"finger-print-voting-backend/internal/config"
	"finger-print-voting-backend/internal/database"
	"finger-print-voting-backend/internal/types"
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

	defer db.Close()

	if err := db.DropDBTables(); err != nil {
		panic(err)
	}

	if err := db.SetupSchema(); err != nil {
		panic(err)
	}

	// write test code here

	registrant := types.RegistrationRequest{FirstName: "First", LastName: "Last", Email: "Email", PhoneNo: "num", Fingerprint: "print", ProofOfIdentity: "proof", Location: "local"}
	db.StoreRegistrant(registrant)
	registrants, _ := db.GetRegistrants()
	log.Println(registrants)

	election := types.ElectionRequest{Start: "2002-02-02", End: "2003-03-03", Location: "Here", Candidates: []types.CandidateRequest{{FirstName: "James", LastName: "Jamison", Party: "Whig", PartyColour: "Purple", Photo: "Cheese"}}}
	err = db.StoreElection(election)
	if err != nil {
		panic(err)
	}
	elections, _ := db.GetElections()
	log.Println(elections)
	candidate, _ := db.GetCandidates(elections[0].ElectionID)
	log.Println(candidate)

	user := types.User{Username: "bancheeban", Password: "pass1", Admin: false, FirstName: "Callum", LastName: "Stone"}
	voter := types.Voter{User: user, PhoneNo: "01438 872874", Email: "email@email.email", Fingerprint: "print", Location: "Here"}
	err = db.StoreVoter(voter)
	if err != nil {
		panic(err)
	}
	log.Println(db.GetVoter("bancheeban"))
}
