package types

type Registrant struct {
	RegistrantID    int    `json:"registrant_id"` // PK
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	PhoneNo         string `json:"phone_no"`
	Fingerprint     string `json:"fingerprint"`
	ProofOfIdentity string `json:"proof_of_identity"`
}
