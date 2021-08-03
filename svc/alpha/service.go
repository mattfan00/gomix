package alpha

type Service interface {
	Hello() string
}

type Alpha struct {
}

func New() Alpha {
	return Alpha{}
}
