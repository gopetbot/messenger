package main

import (
	"github.com/idasilva/messenger/handlers"
	"github.com/idasilva/messenger/server"
	"net/http"
)

func main() {

	mux := handlers.NewRouterInstance()
	mux.AddRoute(&handlers.Context{
		Name:    "Pet",
		Method:  http.MethodGet,
		Path:    "/hello/pet/project",
		Handler: handlers.PetProject,
	})

    server.NewClt(mux, "8080").Start()

}
