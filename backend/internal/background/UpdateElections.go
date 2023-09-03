package background

import (
	"database/sql"
	"finger-print-voting-backend/internal/database"
	"finger-print-voting-backend/internal/types"
	"log"
	"time"
)

func UpdateElections(db database.Database) {
	for {
		time.Sleep(5 * time.Minute)
		log.Println("Updating Elections")

		tx, err := db.Begin()
		if err != nil {
			log.Printf("UpdateElections Begin Error: %s", err.Error())
			continue
		}

		if err := updateElection(tx, db); err != nil {
			tx.Rollback()
			log.Printf("UpdateElections Error: %s", err.Error())
			continue
		}

		if err := tx.Commit(); err != nil {
			tx.Rollback()
			log.Printf("UpdateElections Commit Error: %s", err.Error())
			continue
		}
	}
}

func updateElection(tx *sql.Tx, db database.Database) error {
	elections, _ := db.GetElections()

	updateElections := []types.Election{}
	now := time.Now()
	for _, election := range elections {
		endDate, err := types.StringToDate(election.End)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if now.Unix() > endDate.Unix() {
			updateElections = append(updateElections, election)
		}
	}

	for _, election := range updateElections {
		results, err := db.GetResults(election.ElectionID)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if len(results) != 0 {
			continue
		}

		candidates, err := db.GetCandidates(election.ElectionID)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		votes, err := db.GetVotes(election.ElectionID)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		voteMap := make(map[int]int)

		for _, vote := range votes {
			voteMap[vote.CandidateID] = voteMap[vote.CandidateID] + 1
		}

		newResults := []types.ResultRequest{}
		for _, candidate := range candidates {
			newResults = append(newResults, types.ResultRequest{
				ElectionID:  election.ElectionID,
				FirstName:   candidate.FirstName,
				LastName:    candidate.LastName,
				Party:       candidate.Party,
				PartyColour: candidate.PartyColour,
				Votes:       voteMap[candidate.CandidateID],
			})
		}

		for _, result := range newResults {
			db.StoreResult(tx, result)
		}

		db.DeleteCandidates(tx, election.ElectionID)
		db.DeleteVotes(tx, election.ElectionID)

		log.Printf("updated: %d", election.ElectionID)
	}

	return nil
}
