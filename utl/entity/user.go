package entity

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

type AuthToken struct {
	Access  string `json:"at"`
	Refresh string `json:"rt"`
}
