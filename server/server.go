package server

type Server interface {
	Start() error
	Addr()string
}
