package types

// This is what the database returned when it gets voters
type Voter struct {
	User               // Voters are also users
	PhoneNo     string `json:"phone_no"`
	Email       string `json:"email"`
	Fingerprint string `json:"fingerprint"`
	Location    string `json:"location"`
}
