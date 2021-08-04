package service

type Service interface {
	Current() string
	Register()
	Login(string, string) (string, error)
}

type service struct {
}

func New() service {
	return service{}
}
