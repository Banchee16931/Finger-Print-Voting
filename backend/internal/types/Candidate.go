package types

type Candidate struct {
	CandidateID int    `json:"candidate_id"` // PK
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Photo       string `json:"photo"`
}
