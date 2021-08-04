package service

type Service interface {
	Current() string
	Register()
}

type service struct {
}

func New() service {
	return service{}
}
