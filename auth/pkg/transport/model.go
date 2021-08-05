package transport

import (
	"github.com/mattfan00/gomite/utl/entity"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Message string `json:"message"`
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerResponse struct {
	User  entity.User `json:"user"`
	Token string      `json:"token"`
}
