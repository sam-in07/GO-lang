package middleware

import (
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	mngr := Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
	return &mngr
}

func (mngr *Manager) Use(middleware ...Middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
	return mngr
}

func (mngr *Manager) With(next http.Handler, middleware ...Middleware) http.Handler {
	n := next

	// Apply middlewares in order
	// If middleware = [hudie, logger]
	// 1. n = hudie(n)
	// 2. n = logger(n)
	for _, middleware := range middleware {
		n = middleware(n)
	}
  
	for _, globalMiddlewares := range mngr.globalMiddlewares{
		n = globalMiddlewares(n)
	}
        

	// Now n is the final http.Handler, which matches the return type
	
	
	
	return n
}
