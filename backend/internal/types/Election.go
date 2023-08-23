package types

import "time"

type Election struct {
	ElectionID int         `json:"election_id"` // PK
	Start      time.Time   `json:"start"`
	End        time.Time   `json:"end"`
	Location   string      `json:"location"`
	Candidates []Candidate `json:"candidates"`
}
