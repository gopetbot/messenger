package handlers

import (
	mux "github.com/gorilla/mux"
)
type Router struct {
	Mux *mux.Router

}

//AddRoute, references a new handler
func(s *Router) AddRoute(c *Context){
	s.Mux.HandleFunc(c.Path,c.Handler)

}
//NewRouterInstance created a new instance of the router
func NewRouterInstance()*Router{
	return &Router{
		Mux: mux.NewRouter().StrictSlash(true),
	}

}
