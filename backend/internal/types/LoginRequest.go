package types

import "fmt"

// This is what the API receives to generate a JWT
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Ensures that the LoginRequest has not blank values
func (userReq *LoginRequest) Validate() error {
	if userReq.Username == "" {
		return fmt.Errorf("username is empty")
	}

	if userReq.Password == "" {
		return fmt.Errorf("password is empty")
	}

	return nil
}
