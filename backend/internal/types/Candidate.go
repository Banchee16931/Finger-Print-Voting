package types

import "fmt"

// This is what the database will output when getting a list of candidate
type Candidate struct {
	CandidateID int    `json:"candidate_id"` // PK
	ElectionID  int    `json:"election_id"`  // FK
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
	PartyColour string `json:"party_colour"`
	Photo       string `json:"photo"`
}

// This is what the API will input into the database to generate a Candidate
type CandidateRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
	PartyColour string `json:"party_colour"`
	Photo       string `json:"photo"`
}

// Ensures that the CandidateRequest has not blank values
func (req CandidateRequest) Validate() error {
	if req.FirstName == "" {
		return fmt.Errorf("first name is empty")
	}

	if req.LastName == "" {
		return fmt.Errorf("last name is empty")
	}

	if req.Party == "" {
		return fmt.Errorf("party is empty")
	}

	if req.PartyColour == "" {
		return fmt.Errorf("party colour is empty")
	}

	if req.Photo == "" {
		return fmt.Errorf("no photo given")
	}

	return nil
}
