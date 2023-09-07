package types

import "fmt"

// This is what the database returned when getting votes
type Vote struct {
	Username    string `json:"username"`     // PK
	ElectionID  int    `json:"election_id"`  // PK
	CandidateID int    `json:"candidate_id"` // PK
}

// This is what the API will recieve to register a vote
type VoteRequest struct {
	ElectionID  int    `json:"election_id"`
	Fingerprint string `json:"fingerprint"`
	CandidateID int    `json:"candidate_id"`
}

// Ensures that the VoteRequest has not blank values
func (req VoteRequest) Validate() error {
	if req.Fingerprint == "" {
		return fmt.Errorf("fingerprint is empty")
	}

	return nil
}
