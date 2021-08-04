package auth

type Service interface {
	Current() string
	Register()
}

type Auth struct {
}

func New() Auth {
	return Auth{}
}
