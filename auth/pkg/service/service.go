package service

import (
	"github.com/mattfan00/gomite/auth/pkg/platform/memory"
	"github.com/mattfan00/gomite/utl/entity"
	"github.com/mattfan00/gomite/utl/jwt"
)

type Service interface {
	Register(string, string) (entity.User, entity.AuthToken, error)
	Login(string, string) (string, error)
}

type service struct {
	mem memory.Memory
	atg jwt.TokenGenerator // access token generator
	rtg jwt.TokenGenerator // refresh token generator
}

func New(mem memory.Memory, atg jwt.TokenGenerator, rtg jwt.TokenGenerator) service {
	return service{
		mem: mem,
		atg: atg,
		rtg: rtg,
	}
}
