package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAddToken(t *testing.T) {
	body := strings.NewReader("{\"user_id\":\"1\",\"scenario\":\"test\"}")
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", body)

	rs := NewRestServer()
	rs.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Service did not return ok status")
		t.FailNow()
	}

	if "application/json; charset=utf-8" != rr.Header().Get("Content-Type") {
		t.Fatalf("Service did not return json header")
		t.FailNow()
	}
}

func TestTokenExists(t *testing.T) {
	rs := NewRestServer()
	token := rs.Generator.Get(2, "test")
	rs.Store.Add(token, time.Hour)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/2/"+token.Token, nil)

	rs.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Service did not return ok status")
		t.FailNow()
	}

	if "application/json; charset=utf-8" != rr.Header().Get("Content-Type") {
		t.Fatalf("Service did not return json header")
		t.FailNow()
	}

	if "{\"user_id\":\"2\",\"scenario\":\"test\",\"token\":\""+token.Token+"\"}" != rr.Body.String() {
		t.Fatalf("Unexpected Response: %s", rr.Body.String())
		t.FailNow()
	}
}
