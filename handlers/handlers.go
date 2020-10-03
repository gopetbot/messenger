package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gopetbot/messenger/factory"
	"io/ioutil"
	"net/http"
)

//PetProjectHandler contains a new handler
func PetProjectHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	s := DialogFlowRequest{}
	jsonByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(jsonByte, &s)
	if err != nil {
		fmt.Println(err)
	}

	platform, err := factory.SamplePlatform(s.OriginalRequest.Source)
	if err == nil {

		message := fmt.Sprintf("Platform %v does not exist!", s.OriginalRequest.Source)
		handleResponse(w, message, http.StatusNotFound)
		return
	}
	message, err := platform.HandleMessage("I am the message!")
	if err != nil {
		handleResponse(w, message, http.StatusInternalServerError)
	}
	handleResponse(w, message, http.StatusOK)
}


func handleResponse(w http.ResponseWriter, message string, httpCode int) {
	response, _ := json.Marshal(Response{Message: message})
	w.WriteHeader(httpCode)
	w.Write([]byte(string(response)))
}