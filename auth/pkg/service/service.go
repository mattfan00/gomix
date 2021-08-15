package service

import (
	"github.com/mattfan00/gomix/auth/pkg/platform/pg"
	"github.com/mattfan00/gomix/utl/bcrypt"
	"github.com/mattfan00/gomix/utl/entity"
	"github.com/mattfan00/gomix/utl/jwt"
)

type Service interface {
	Register(string, string) (entity.User, entity.AuthToken, error)
	Login(string, string) (entity.User, entity.AuthToken, error)
}

type service struct {
	pg  pg.Store
	atg jwt.TokenGenerator // access token generator
	rtg jwt.TokenGenerator // refresh token generator
	bc  bcrypt.Service
}

func New(
	pg pg.Store,
	atg jwt.TokenGenerator,
	rtg jwt.TokenGenerator,
	bc bcrypt.Service,
) service {
	return service{
		pg:  pg,
		atg: atg,
		rtg: rtg,
		bc:  bc,
	}
}
