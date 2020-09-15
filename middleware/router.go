package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type Context struct {
	Name    string
	Method  string
	Path    string
	Handler Handler
}

type Router struct {
	Mux *mux.Router
}

//AddRoute, references a new handler
func (s *Router) AddRoute(c *Context) {
	s.Mux.HandleFunc(c.Path, c.Handler)

}

//NewRouterInstance created a new instance of the router
func NewRouterInstance(apiVersion string) *Router {
	return &Router{
		Mux: mux.NewRouter().StrictSlash(true).PathPrefix(apiVersion).Subrouter(),
	}

}
