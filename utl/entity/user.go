package entity

import (
	"github.com/satori/go.uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" pg:",pk,type:uuid,default:uuid_generate_v4()"`
	Username string    `json:"username" pg:",unique"`
	Password string    `json:"-"`
}

type AuthToken struct {
	Access  string `json:"at"`
	Refresh string `json:"rt"`
}
