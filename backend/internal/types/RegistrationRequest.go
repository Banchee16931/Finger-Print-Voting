package types

import "fmt"

type RegistrationRequest struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	PhoneNo         string `json:"phone_no"`
	Fingerprint     string `json:"fingerprint"`
	ProofOfIdentity string `json:"proof_of_identity"`
	Location        string `json:"location"`
}

func (req RegistrationRequest) Validate() error {
	if req.FirstName == "" {
		return fmt.Errorf("first name is empty")
	}

	if req.LastName == "" {
		return fmt.Errorf("last name is empty")
	}

	if req.Email == "" {
		return fmt.Errorf("email is empty")
	}

	if req.PhoneNo == "" {
		return fmt.Errorf("phone number is empty")
	}

	if req.Fingerprint == "" {
		return fmt.Errorf("fingerprint is empty")
	}

	if req.ProofOfIdentity == "" {
		return fmt.Errorf("proof of identity is empty")
	}

	if req.Location == "" {
		return fmt.Errorf("location is empty")
	}

	return nil
}
