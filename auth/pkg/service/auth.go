package service

func (s service) Current() string {
	return "this is me"
}

func (s service) Register() {
}

func (s service) Login(username string, password string) (string, error) {
	if username != "matt" || password != "password" {
		return "", nil
	}

	return "logged in", nil
}
