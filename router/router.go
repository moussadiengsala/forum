package router

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type NewRouter struct {
	Middlewares []Middleware
}

func (router *NewRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var handler = Handlers(w, r)
	for _, mw := range router.Middlewares {
		handler = mw(handler)
	}
	handler(w, r)
}
