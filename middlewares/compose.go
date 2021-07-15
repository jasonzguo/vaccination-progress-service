package middleware

import "github.com/julienschmidt/httprouter"

type Middleware func(httprouter.Handle) httprouter.Handle

type stack struct {
	middlewares []Middleware
}

func NewStack() *stack {
	return &stack{
		middlewares: []Middleware{},
	}
}

func (s *stack) Use(mw Middleware) {
	s.middlewares = append(s.middlewares, mw)
}

func (s *stack) Wrap(fn httprouter.Handle) httprouter.Handle {
	l := len(s.middlewares)
	if l == 0 {
		return fn
	}

	var result httprouter.Handle
	result = s.middlewares[l-1](fn)

	for i := l - 2; i >= 0; i-- {
		result = s.middlewares[i](result)
	}

	return result
}
