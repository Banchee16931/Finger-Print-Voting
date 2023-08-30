package users

import (
	"finger-print-voting-backend/internal/database"
	"finger-print-voting-backend/internal/types"

	"golang.org/x/crypto/bcrypt"
)

func NewUser(db database.Database, user types.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.Password = string(encryptedPassword)

	return db.StoreUser(user)
}
