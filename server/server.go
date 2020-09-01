package server

import (
	"github.com/gopetbot/messenger/handlers"
	"log"
	"net/http"
	"time"
)

type Simple struct {
	svr *http.Server
}

func (s *Simple) Start() {
	log.Println("Create a simple server!!! Port: ", s.svr.Addr)
	if err := s.svr.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}

//New create a new instance of the server to http request and response
func NewClt(handler *handlers.Router, addr string) *Simple {
	return &Simple{
		svr: &http.Server{
			Addr:              addr,
			Handler:           handler.Mux,
			ReadTimeout:       600 * time.Second,
			ReadHeaderTimeout: 600 * time.Second,
		},
	}

}
