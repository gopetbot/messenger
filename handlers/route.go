package handlers

import (
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type Context struct {
	Name    string
	Method  string
	Path    string
	Handler Handler
}

//PetProject contains a new handler
func PetProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Salvado e abençoando")
	w.WriteHeader(http.StatusOK)
}
