package database

import (
	"database/sql"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
)

func (client *Client) StoreUser(tx *sql.Tx, user types.User) error {
	log.Println("details: ", user.Username, user.Password, user.Admin, user.FirstName, user.LastName)
	_, err := client.db.Exec(`INSERT INTO users (username, encrypted_password, is_admin, first_name, last_name)
    VALUES ($1, $2, $3, $4, $5);`, user.Username, user.Password, user.Admin, user.FirstName, user.LastName)

	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}

func (client *Client) StoreVoter(tx *sql.Tx, voter types.Voter) error {
	log.Println("details: ", voter.Username, voter.PhoneNo, voter.Email)
	_, err := tx.Exec(`INSERT INTO voter_details (username, phone_no, email, fingerprint, authority_location)
	VALUES ($1, $2, $3, $4, $5);`, voter.User.Username, voter.PhoneNo, voter.Email, voter.Fingerprint, voter.Location)

	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	err = client.StoreUser(tx, voter.User)

	if err != nil {
		return fmt.Errorf("stored user: %w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}

func (client *Client) GetVoter(username string) (types.Voter, error) {

	row := client.db.QueryRow(`SELECT phone_no, email, fingerprint, authority_location FROM voter_details WHERE username=$1;`, username)

	voter := types.Voter{}

	if err := row.Scan(&voter.PhoneNo, &voter.Email, &voter.Fingerprint, &voter.Location); err != nil {
		if err == sql.ErrNoRows {
			return types.Voter{}, cerr.ErrNotFound
		}

		return types.Voter{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	user, err := client.GetUser(username)

	if err != nil {
		return types.Voter{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	voter.User = user

	return voter, nil
}

func (client *Client) GetUser(username string) (types.User, error) {
	log.Printf("Getting details of username: %s", username)
	row := client.db.QueryRow(`SELECT username, encrypted_password, is_admin, first_name, last_name FROM users WHERE username=$1;`, username)

	user := types.User{}

	if err := row.Scan(&user.Username, &user.Password, &user.Admin, &user.FirstName, &user.LastName); err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, cerr.ErrNotFound
		}

		return types.User{}, fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return user, nil
}

func (client *Client) DeleteVoter(username string) error {
	_, err := client.db.Exec(`DELETE FROM voter_details WHERE username = $1;`, username)
	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	_, err = client.db.Exec(`DELETE FROM users WHERE username = $1;`, username)
	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}
