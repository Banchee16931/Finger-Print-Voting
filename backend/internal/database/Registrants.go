package database

import (
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
)

func (client *Client) StoreRegistrant(registrant types.RegistrationRequest) error {
	_, err := client.db.Exec(`INSERT INTO registrants (first_name, last_name, email, phone_no, fingerprint, proof, authority_location)
    VALUES ($1, $2, $3, $4, $5, $6, $7);`, registrant.FirstName, registrant.LastName, registrant.Email,
		registrant.PhoneNo, registrant.Fingerprint, registrant.ProofOfIdentity, registrant.Location)

	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}

func (client *Client) GetRegistrants() ([]types.Registrant, error) {
	log.Printf("Getting all registrants")

	rows, err := client.db.Query(`SELECT registrant_id, first_name, last_name, email, phone_no, fingerprint, proof, authority_location FROM registrants;`)
	if err != nil {
		return []types.Registrant{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	registrants := []types.Registrant{}

	for rows.Next() {
		registrant := types.Registrant{}

		if err := rows.Scan(&registrant.RegistrantID, &registrant.FirstName, &registrant.LastName, &registrant.Email, &registrant.PhoneNo,
			&registrant.Fingerprint, &registrant.ProofOfIdentity, &registrant.Location); err != nil {
			return []types.Registrant{}, err
		}
		registrants = append(registrants, registrant)
	}

	if err = rows.Err(); err != nil {
		return []types.Registrant{}, err
	}

	return registrants, nil
}

func (client *Client) DeleteRegistrant(registrantID int) error {
	_, err := client.db.Exec(`DELETE FROM registrants WHERE registrant_id = $1;`, registrantID)
	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}
