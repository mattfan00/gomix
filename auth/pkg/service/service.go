package service

import (
	"github.com/mattfan00/gomite/auth/pkg/platform/memory"
	"github.com/mattfan00/gomite/utl/entity"
)

type Service interface {
	Current() string
	Register(string, string) (entity.User, string, error)
	Login(string, string) (string, error)
}

type service struct {
	mem memory.Memory
}

func New(mem memory.Memory) service {
	return service{
		mem: mem,
	}
}
