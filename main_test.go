package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopetbot/messenger/pkg"
)

func TestGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello/pet/project", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pkg.PetProject)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"hello": "world"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
