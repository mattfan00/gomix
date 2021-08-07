package service

import (
	"github.com/mattfan00/gomite/auth/pkg/platform/pg"
	"github.com/mattfan00/gomite/utl/entity"
	"github.com/mattfan00/gomite/utl/jwt"
)

type Service interface {
	Register(string, string) (entity.User, entity.AuthToken, error)
	Login(string, string) (string, error)
}

type service struct {
	pg  pg.Store
	atg jwt.TokenGenerator // access token generator
	rtg jwt.TokenGenerator // refresh token generator
}

func New(pg pg.Store, atg jwt.TokenGenerator, rtg jwt.TokenGenerator) service {
	return service{
		pg:  pg,
		atg: atg,
		rtg: rtg,
	}
}
