package users

import (
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/database"
	"finger-print-voting-backend/internal/types"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func NewUser(db database.Database, user types.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.Password = string(encryptedPassword)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = db.StoreUser(tx, user)
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
