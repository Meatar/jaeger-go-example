package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}

func TestPingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pingHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, but got %v", status)
	}

	expected := "Pong"
	if rr.Body.String() != expected {
		t.Errorf("Expected body to be %s, but got %s", expected, rr.Body.String())
	}
}
