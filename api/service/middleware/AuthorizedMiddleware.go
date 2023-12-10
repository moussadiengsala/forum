package middleware

import (
	"fmt"
	"net/http"
)

func AuthorizedMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("headerrrrrrrrrrrrrrrrrrrrrrrrrrrr %T\n", r.Header)
		handler(w, r)
	}
}
