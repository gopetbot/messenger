package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopetbot/messenger/middleware"
)

func TestGetFacebookWebhook(t *testing.T) {

	requestPayload := &middleware.DialogFlowRequest{
		middleware.OriginalRequest{
			Source: "facebook",
		},
	}

	marshaledValue, _ := json.Marshal(requestPayload)
	req, err := http.NewRequest("POST", "/v1/pet/webhook", bytes.NewReader(marshaledValue))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware.PetProjectHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"message":"handled for facebook"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
