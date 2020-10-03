package main

import (
	"net/http"

	"github.com/gopetbot/messenger/middleware"
	"github.com/gopetbot/messenger/server"
)

func main() {

	mux := middleware.NewRouterInstance("/v1")
	mux.AddRoute(&middleware.Context{
		Name:    "Pet",
		Method:  http.MethodPost,
		Path:    "/pet/webhook",
		Handler: middleware.PetProjectHandler,
	})

	server.NewClt(mux, "0.0.0.0", "8080").Start()

}
