package types

type Result struct {
	ResultID    int    `json:"result_id"` // PK
	ElectionID  int    `json:"election_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
	PartyColour string `json:"party_colour"`
	Votes       int    `json:"votes"`
}

type ResultRequest struct {
	ElectionID  int    `json:"election_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
	PartyColour string `json:"party_colour"`
	Votes       int    `json:"votes"`
}
