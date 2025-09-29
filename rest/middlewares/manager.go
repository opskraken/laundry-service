package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type MiddlewareManager struct {
	globalMiddlewares []Middleware
}

func NewMiddlewareManager() *MiddlewareManager {
	return &MiddlewareManager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mng *MiddlewareManager) Use(middlewares ...Middleware) {
	mng.globalMiddlewares = append(mng.globalMiddlewares, middlewares...)
}

func (mng *MiddlewareManager) Apply(handler http.Handler) http.Handler {
	for i := len(mng.globalMiddlewares) - 1; i >= 0; i-- {
		handler = mng.globalMiddlewares[i](handler)
	}
	return handler
}

func (mng *MiddlewareManager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
