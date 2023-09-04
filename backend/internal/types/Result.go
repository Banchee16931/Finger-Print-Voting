package types

// This is what the database returned when getting results
type Result struct {
	ResultID    int    `json:"result_id"` // PK
	ElectionID  int    `json:"election_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
	PartyColour string `json:"party_colour"`
	Votes       int    `json:"votes"`
}

// This is what the API will input into the database to generate a Result
type ResultRequest struct {
	ElectionID  int    `json:"election_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
	PartyColour string `json:"party_colour"`
	Votes       int    `json:"votes"`
}
