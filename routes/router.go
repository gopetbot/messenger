package routes

import (
	"context"
	"github.com/gopetbot/messenger/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

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
	Mux    *mux.Router
	ctx    context.Context
	sample *server.Sample
}

func (r *Router) Server() error {

	ctx, cancel := context.WithCancel(r.ctx)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		log.Println("signal caught. shutting down...")
		cancel()
		r.sample.Shutdown(ctx)
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		r.sample.Start(r.Mux)
	}()

	wg.Wait()
	return nil

}

//AddRoute, references a new handler
func (s *Router) AddRoute(c *Context) {
	s.Mux.HandleFunc(c.Path, c.Handler)

}

//NewRouterInstance created a new instance of the router
func NewRouterInstance(apiVersion string, host string, port string) *Router {
	return &Router{
		Mux:    mux.NewRouter().StrictSlash(true).PathPrefix(apiVersion).Subrouter(),
		ctx:    context.Background(),
		sample: server.NewSample(host, port),
	}

}
