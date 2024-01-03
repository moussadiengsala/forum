package internals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"learn.zone01dakar.sn/forum-rest-api/lib"
)

// App represents the core application structure, managing routes, middleware, and handling HTTP requests.
type App struct {
	routes     map[string]map[string]lib.Handler
	handler    http.Handler
	middleware []lib.Middleware
}

// NewApp creates a new instance of the App with default configurations.
func NewApp() *App {
	return &App{
		routes: make(map[string]map[string]lib.Handler),
		handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			http.Error(w, "Handler not configured", http.StatusInternalServerError)
		}),
	}
}

// Use adds middleware functions to the application middleware stack.
func (a *App) Use(middleware ...lib.Middleware) {
	a.middleware = append(a.middleware, middleware...)
}

// AddRoute adds a route to the framework.
func (a *App) AddRoute(method, path string, handler lib.Handler) {
	prefixedPath := fmt.Sprintf("/api%s", path)

	if _, ok := a.routes[prefixedPath]; !ok {
		a.routes[prefixedPath] = make(map[string]lib.Handler)
	}
	a.routes[prefixedPath][method] = handler
}

// GET is a shorthand method to add a GET route to the framework.
func (a *App) GET(path string, handler lib.Handler) {
	a.AddRoute("GET", path, handler)
}

// POST is a shorthand method to add a POST route to the framework.
func (a *App) POST(path string, handler lib.Handler) {
	a.AddRoute("POST", path, handler)
}

// ServeHTTP implements the http.Handler interface to handle incoming HTTP requests.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	for routePath, routeHandler := range a.routes {
		handler, ok := routeHandler[method]
		if strings.HasPrefix(path, routePath) && ok {
			finalHandler := handler
			for _, mw := range a.middleware {
				finalHandler = mw(finalHandler)
			}

			finalHandler(w, r)
			return
		} else if strings.HasPrefix(path, routePath) {
			response := map[string]interface{}{
				"Code":    http.StatusMethodNotAllowed,
				"Message": "Method Not Allowed",
				"Data":    nil,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(response)
			return
		}

	}

	response := map[string]interface{}{
		"Code":    http.StatusNotFound,
		"Message": "Not Found",
		"Data":    nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}
