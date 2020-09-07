package pkg

import (
	"log"
	"net/http"
)

//PetProject contains a new handler
func PetProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Salvado e abençoando...")
	w.WriteHeader(http.StatusOK)
}
