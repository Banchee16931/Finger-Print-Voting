package types

import "fmt"

type Vote struct {
	Username    string `json:"username"`     // PK
	ElectionID  int    `json:"election_id"`  // PK
	CandidateID int    `json:"candidate_id"` // PK
}

type VoteRequest struct {
	ElectionID  int    `json:"election_id"`
	Fingerprint string `json:"fingerprint"`
	CandidateID int    `json:"candidate_id"`
}

func (req VoteRequest) Validate() error {
	if req.Fingerprint == "" {
		return fmt.Errorf("fingerprint is empty")
	}

	return nil
}
