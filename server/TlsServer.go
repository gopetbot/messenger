package server

import (
	"log"
	"net/http"
)

type TLsSample struct {
	*http.Server
	TlsKey string
	TlsPem string
}

func (s *TLsSample) Addr() string{
	return s.Server.Addr
}


func (t *TLsSample) Start() error {
	log.Println("Create a tls server!!! Port: ", t.Addr())
	if err := t.ListenAndServeTLS(t.TlsPem, t.TlsKey); err != nil {
		return err
	}
	return nil

}

//NewTls create a new instance of the server to https request and response
func NewTls(host string, port string) *Sample {
	return &Sample{}

}
