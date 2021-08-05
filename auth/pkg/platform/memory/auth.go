package memory

import (
	"github.com/mattfan00/gomite/utl/entity"
)

type Memory interface {
	Register(string, string) (entity.User, error)
}

type memory struct {
	users []entity.User
}

func New() *memory {
	return &memory{
		users: []entity.User{},
	}
}

func (m *memory) Register(username string, password string) (entity.User, error) {
	newUser := entity.User{
		Username: username,
		Password: password,
	}

	m.users = append(m.users, newUser)
	return newUser, nil
}
