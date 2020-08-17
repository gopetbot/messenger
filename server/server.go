package server

import (
	"fmt"
	"github.com/idasilva/messenger/handlers"
	"net/http"
	"time"
)

type Simple struct {
	svr *http.Server
}

func (s *Simple) Start() {
	fmt.Println("Create a simple server!!! Port: ",s.svr.Addr)
	if err := s.svr.ListenAndServe(); err != nil{
		fmt.Println(err.Error())
	}
}

//New create a new instance of the server to http request and response
func NewClt(handler *handlers.Router, addr string) *Simple {
	return &Simple{
		svr: &http.Server{
			Addr:              addr,
			Handler:           handler.Mux,
			ReadTimeout:       600* time.Second,
			ReadHeaderTimeout: 600 * time.Second,
		},
	}

}
