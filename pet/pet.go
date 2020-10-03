package pet

import (
	"github.com/gopetbot/messenger/handlers"
	"github.com/gopetbot/messenger/routes"
	"net/http"
)

func Run() *routes.Router {
	r := routes.NewRouterInstance("/v1", "0.0.0.0", "8080")
	r.AddRoute(&routes.Context{
		Name:    "Pet",
		Path:    "/hello/pet/project",
		Method:  http.MethodPost,
		Handler: handlers.PetProjectHandler,
	})

	return r
}
