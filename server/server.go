package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gopetbot/messenger/middleware"
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
func NewClt(handler *middleware.Router, host string, port string) *Simple {
	return &Simple{
		svr: &http.Server{
			Addr:              fmt.Sprintf("%v:%v", host, port),
			Handler:           handler.Mux,
			ReadTimeout:       600 * time.Second,
			ReadHeaderTimeout: 600 * time.Second,
		},
	}

}
