package types

import "fmt"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (userReq *LoginRequest) Validate() error {
	if userReq.Username == "" {
		return fmt.Errorf("username is empty")
	}

	if userReq.Password == "" {
		return fmt.Errorf("password is empty")
	}

	return nil
}
