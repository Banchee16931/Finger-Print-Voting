package database

import (
	"database/sql"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
)

func (client *Client) StoreElection(election types.ElectionRequest) error {
	tx, err := client.db.Begin()

	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	lastIndexID := -1

	err = tx.QueryRow(`INSERT INTO elections (election_start, election_end, authority_location)
    VALUES ($1, $2, $3) RETURNING election_id;`, election.Start, election.End, election.Location).Scan(&lastIndexID)

	if err != nil {
		rollErr := tx.Rollback()
		if rollErr != nil {
			return fmt.Errorf("%w: %s: %s", cerr.ErrDB, rollErr.Error(), err.Error())
		}
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	if lastIndexID == -1 {
		err = fmt.Errorf("last index ID not returned")
		rollErr := tx.Rollback()
		if rollErr != nil {
			return fmt.Errorf("%w: %s: %s", cerr.ErrDB, rollErr.Error(), err.Error())
		}
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	err = StoreCandidates(tx, lastIndexID, election.Candidates)

	if err != nil {
		rollErr := tx.Rollback()
		if rollErr != nil {
			return fmt.Errorf("%w: %s: %s", cerr.ErrDB, rollErr.Error(), err.Error())
		}
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	err = tx.Commit()

	if err != nil {
		rollErr := tx.Rollback()
		if rollErr != nil {
			return fmt.Errorf("%w: %s: %s", cerr.ErrDB, rollErr.Error(), err.Error())
		}
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}

func StoreCandidates(tx *sql.Tx, electionID int, candidates []types.CandidateRequest) error {
	log.Printf("Store candidate")

	log.Print(candidates)

	for _, candidate := range candidates {

		log.Printf(candidate.FirstName, candidate.LastName, candidate.Party, candidate.PartyColour, candidate.Photo)
		log.Println("here")
		log.Println(tx == nil)

		_, err := tx.Exec(`INSERT INTO candidates (election_id, first_name, last_name, party, party_colour, photo)
		VALUES ($1, $2, $3, $4, $5, $6);`, electionID, candidate.FirstName, candidate.LastName, candidate.Party, candidate.PartyColour, candidate.Photo)

		if err != nil {
			log.Println("err")
			return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
		}
	}

	log.Printf("Complete")

	return nil
}

func (client *Client) GetElection(electionID int) (types.Election, error) {
	log.Printf("Getting election: %d", electionID)
	election := types.Election{}
	err := client.db.QueryRow(`SELECT election_id, election_start, election_end, authority_location FROM elections WHERE election_id=$1;`,
		electionID).Scan(&election.ElectionID, &election.Start, &election.End, &election.Location)
	if err != nil {
		return types.Election{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return election, nil
}

func (client *Client) GetElections() ([]types.Election, error) {
	log.Printf("Getting all elections")

	rows, err := client.db.Query(`SELECT election_id, election_start, election_end, authority_location FROM elections;`)
	if err != nil {
		return []types.Election{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	elections := []types.Election{}

	for rows.Next() {
		election := types.Election{}

		if err := rows.Scan(&election.ElectionID, &election.Start, &election.End, &election.Location); err != nil {
			return []types.Election{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
		}

		elections = append(elections, election)
	}

	if err = rows.Err(); err != nil {
		return []types.Election{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return elections, nil
}

func (client *Client) GetElectionByLocation(location string) ([]types.Election, error) {
	log.Printf("Getting elections from %s", location)

	rows, err := client.db.Query(`SELECT election_id, election_start, election_end, authority_location FROM elections WHERE authority_location=$1;`,
		location)
	if err != nil {
		return []types.Election{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	elections := []types.Election{}

	for rows.Next() {
		election := types.Election{}

		if err := rows.Scan(&election.ElectionID, &election.Start, &election.End, &election.Location); err != nil {
			return []types.Election{}, err
		}

		if err != nil {
			return []types.Election{}, err
		}

		elections = append(elections, election)
	}

	if err = rows.Err(); err != nil {
		return []types.Election{}, err
	}

	return elections, nil
}

func (client *Client) GetCandidates(electionID int) ([]types.Candidate, error) {
	log.Printf("Getting candidates")

	rows, err := client.db.Query(`SELECT candidate_id, election_id, first_name, last_name, party, party_colour, photo FROM candidates;`)
	if err != nil {
		return []types.Candidate{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	candidates := []types.Candidate{}

	for rows.Next() {
		candidate := types.Candidate{}

		if err := rows.Scan(&candidate.CandidateID, &candidate.ElectionID, &candidate.FirstName, &candidate.LastName, &candidate.Party, &candidate.PartyColour, &candidate.Photo); err != nil {
			return []types.Candidate{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
		}
		candidates = append(candidates, candidate)
	}

	if err = rows.Err(); err != nil {
		return []types.Candidate{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return candidates, nil
}

func (client *Client) DeleteCandidates(electionID int) error {
	_, err := client.db.Exec(`DELETE FROM candidates WHERE election_id = $1;`, electionID)
	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}
