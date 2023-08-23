package types

type Voter struct {
	User               // Voters are also users
	PhoneNo     string `json:"phone_no"`
	Email       string `json:"email"`
	Fingerprint string `json:"fingerprint"`
	Location    string `json:"location"`
}
