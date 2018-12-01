package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestThatServiceIsUp(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/check", nil)

	rs := NewRestServer()
	rs.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("Service did not return ok no content status")
		t.FailNow()
	}
}

func TestGetMetrics(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/metrics", nil)

	rs := NewRestServer()
	rs.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("Service did not return ok no content status")
		t.FailNow()
	}
}
