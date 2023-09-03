package types

import "fmt"

type UserAcceptanceRequest struct {
	RegistrantID int    `json:"registrant_id"`
	Accepted     bool   `json:"accepted"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func (req UserAcceptanceRequest) Validate() error {
	if req.Username == "" {
		return fmt.Errorf("username is empty")
	}

	if req.Password == "" {
		return fmt.Errorf("password is empty")
	}

	return nil
}
