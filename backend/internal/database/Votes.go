package database

import (
	"database/sql"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
)

func (client *Client) StoreVote(vote types.Vote) error {
	_, err := client.db.Exec(`INSERT INTO votes (username, election_id, candidate_id)
    VALUES ($1, $2, $3);`, vote.Username, vote.ElectionID, vote.CandidateID)

	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}

func (client *Client) GetVotes(electionID int) ([]types.Vote, error) {
	log.Printf("Getting all votes")

	rows, err := client.db.Query(`SELECT username, election_id, candidate_id FROM votes WHERE election_id=$1;`, electionID)
	if err != nil {
		return []types.Vote{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	votes := []types.Vote{}

	for rows.Next() {
		vote := types.Vote{}

		if err := rows.Scan(&vote.Username, &vote.ElectionID, &vote.CandidateID); err != nil {
			return []types.Vote{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
		}
		votes = append(votes, vote)
	}

	if err = rows.Err(); err != nil {
		return []types.Vote{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return votes, nil
}

func (client *Client) DeleteVotes(tx *sql.Tx, electionID int) error {
	_, err := tx.Exec(`DELETE FROM votes WHERE election_id=$1;`, electionID)
	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}
