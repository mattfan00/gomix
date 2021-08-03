package auth

type Service interface {
	Current() string
}

type Auth struct {
}

func New() Auth {
	return Auth{}
}
