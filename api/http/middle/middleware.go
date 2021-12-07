package middle

type Middleware struct {
}

func NewMiddleware() (m *Middleware, err error) {
	return &Middleware{}, nil
}
