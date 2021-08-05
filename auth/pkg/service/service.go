package service

import (
	"github.com/mattfan00/gomite/auth/pkg/platform/memory"
)

type Service interface {
	Current() string
	Register(string, string) ([]string, error)
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
