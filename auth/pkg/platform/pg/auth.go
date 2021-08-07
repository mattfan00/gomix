package pg

import (
	"github.com/mattfan00/gomix/utl/entity"

	"github.com/go-pg/pg/v10"
)

type Store interface {
	Register(entity.User) (entity.User, error)
}

type store struct {
	db *pg.DB
}

func New(db *pg.DB) store {
	return store{
		db: db,
	}
}

func (s store) Register(newUser entity.User) (entity.User, error) {
	_, err := s.db.Model(&newUser).Insert()

	return newUser, err
}
