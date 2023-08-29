package database

import (
	"database/sql"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
)

func (client *Client) StoreUser(user types.User) error {
	_, err := client.db.Exec(`INSERT INTO users (username, encrypted_password, is_admin, first_name, last_name)
    VALUES ($1, $2, $3, $4, $5);`, user.Username, user.Password, user.Admin, user.FirstName, user.LastName)

	if err != nil {
		return fmt.Errorf("%w: %s", cerr.ErrDB, err.Error())
	}

	return nil
}

func (client *Client) StoreVoter(voter types.Voter) error {
	return cerr.ErrUnimplemented
}

func (client *Client) GetVoter(username string) (types.Voter, error) {
	return types.Voter{}, cerr.ErrUnimplemented
}

func (client *Client) GetUser(username string) (types.User, error) {
	log.Printf("Getting details of username: %s", username)
	row := client.db.QueryRow(`SELECT username, encrypted_password, is_admin, first_name, last_name FROM users WHERE username=$1;`, username)

	user := types.User{}

	if err := row.Scan(&user.Username, &user.Password, &user.Admin, &user.FirstName, &user.LastName); err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, cerr.ErrNotFound
		}

		return types.User{}, err
	}

	return user, nil
}

func (client *Client) DeleteVoter(voter types.Voter) error {
	return cerr.ErrUnimplemented
}
