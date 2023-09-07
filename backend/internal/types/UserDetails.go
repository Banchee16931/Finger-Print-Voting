package types

// A user without the password information
type UserDetails struct {
	Username  string `json:"username"`
	Admin     bool   `json:"admin"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Transfers the data from a User into the UserDetails
func (res UserDetails) FromUser(user User) UserDetails {
	res.Username = user.Username
	res.Admin = user.Admin
	res.FirstName = user.FirstName
	res.LastName = user.LastName
	return res
}

const UserContext Context = "user"
