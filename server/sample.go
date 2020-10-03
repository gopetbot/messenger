package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Sample struct {
	*http.Server
}
func (s *Sample) Addr() string{
	 return s.Server.Addr
}

func (s *Sample) Start(handler http.Handler) error {
	s.Handler = handler
	log.Println("Create a simple server!!! Port: ", s.Addr())
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return  nil
}

//NewSample create a new instance of the server to http request and response
func NewSample(host string, port string) *Sample {
	return &Sample{
		 &http.Server{
		 	Addr: fmt.Sprintf("%v:%v", host, port),
			ReadTimeout:       600 * time.Second,
			ReadHeaderTimeout: 600 * time.Second,
		},
	}

}
