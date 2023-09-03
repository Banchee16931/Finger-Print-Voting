package database

import (
	"database/sql"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
)

func (client *Client) StoreResult(tx *sql.Tx, result types.ResultRequest) error {
	_, err := tx.Exec(`INSERT INTO result (election_id, first_name, last_name, party, votes)
    VALUES ($1, $2, $3, $4, $5);`, result.ElectionID, result.FirstName, result.LastName, result.Party, result.Votes)

	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}

func (client *Client) GetResults(electionID int) ([]types.Result, error) {
	log.Printf("Getting all results")

	rows, err := client.db.Query(`SELECT result_id, election_id, first_name, last_name, party, votes FROM result WHERE election_id=$1;`, electionID)
	if err != nil {
		return []types.Result{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	results := []types.Result{}

	for rows.Next() {
		result := types.Result{}

		if err := rows.Scan(&result.ResultID, &result.ElectionID, &result.FirstName, &result.LastName, &result.Party, &result.Votes); err != nil {
			return []types.Result{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return []types.Result{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return results, nil
}
