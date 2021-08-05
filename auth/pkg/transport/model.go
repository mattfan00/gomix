package transport

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
	Users []string `json:"users"`
}
