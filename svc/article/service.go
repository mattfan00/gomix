package article

type Service interface {
	Get() string
}

type Article struct {
}

func New() Article {
	return Article{}
}
