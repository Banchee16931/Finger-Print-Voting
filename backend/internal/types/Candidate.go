package types

type Candidate struct {
	CandidateID int    `json:"candidate_id"` // PK
	ElectionID  int    `json:"election_id"`  // FK
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
	PartyColour string `json:"party_colour"`
	Photo       string `json:"photo"`
}

type CandidateRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
	PartyColour string `json:"party_colour"`
	Photo       string `json:"photo"`
}
