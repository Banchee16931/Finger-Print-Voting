package types

type Vote struct {
	Username    string `json:"username"`     // PK
	ElectionID  int    `json:"election_id"`  // PK
	CandidateID int    `json:"candidate_id"` // PK
}
