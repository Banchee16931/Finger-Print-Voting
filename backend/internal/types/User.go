package types

type User struct {
	Username  string `json:"username"` // PK
	Password  string `json:"password"`
	Admin     bool   `json:"admin"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
