package bcrypt

import (
	"errors"

	bcryptGo "golang.org/x/crypto/bcrypt"
)

var (
	ErrHashingPasword = errors.New("Error hashing password")
)

type Service interface {
	Hash(string) (string, error)
	Compare(string, string) bool
}

type service struct {
	cost int
}

func New(cost int) service {
	return service{
		cost: cost,
	}
}

func (s service) Hash(password string) (string, error) {
	hashed, err := bcryptGo.GenerateFromPassword([]byte(password), s.cost)
	if err != nil {
		return "", ErrHashingPasword
	}

	return string(hashed), nil
}

func (s service) Compare(hashedPassword string, password string) bool {
	err := bcryptGo.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	// if there is an error, then they are not the same
	return err == nil
}
