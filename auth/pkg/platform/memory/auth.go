package memory

type Memory interface {
	Register(string, string) ([]string, error)
}

type memory struct {
	users []string
}

func New() *memory {
	return &memory{
		users: []string{},
	}
}

func (m *memory) Register(username string, password string) ([]string, error) {
	m.users = append(m.users, username)
	return m.users, nil
}
