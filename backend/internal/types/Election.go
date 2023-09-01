package types

type ElectionRequest struct {
	Start      string             `json:"start"`
	End        string             `json:"end"`
	Location   string             `json:"location"`
	Candidates []CandidateRequest `json:"candidates"`
}

type Election struct {
	ElectionID int    `json:"election_id"` // PK
	Start      string `json:"start"`
	End        string `json:"end"`
	Location   string `json:"location"`
}
