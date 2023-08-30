package database

import (
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
	"fmt"
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
	return []types.Registrant{}, cerr.ErrUnimplemented
}

func (client *Client) DeleteRegistrant(registrantID int) error {
	return cerr.ErrUnimplemented
}
