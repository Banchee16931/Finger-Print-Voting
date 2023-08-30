package types

type UserDetails struct {
	Username  string `json:"username"`
	Admin     bool   `json:"admin"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (res UserDetails) FromUser(user User) UserDetails {
	res.Username = user.Username
	res.Admin = user.Admin
	res.FirstName = user.FirstName
	res.LastName = user.LastName
	return res
}

const UserContext Context = "user"
