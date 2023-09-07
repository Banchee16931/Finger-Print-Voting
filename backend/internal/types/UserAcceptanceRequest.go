package types

import "fmt"

// This is what the API will received to accept or decline a registration request
type UserAcceptanceRequest struct {
	RegistrantID int    `json:"registrant_id"`
	Accepted     bool   `json:"accepted"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

// Ensures that the UserAcceptanceRequest has not blank values
func (req UserAcceptanceRequest) Validate() error {
	if req.Username == "" {
		return fmt.Errorf("username is empty")
	}

	if req.Password == "" {
		return fmt.Errorf("password is empty")
	}

	return nil
}
