package main

import (
	"fmt"
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
	req, _ := http.NewRequest("GET", "/2/"+token.Scenario+"/"+token.Token, nil)

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

func TestTokenExistsWithoutScenario(t *testing.T) {
	rs := NewRestServer()
	token := rs.Generator.Get(3, "test")
	rs.Store.Add(token, time.Hour)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/3/"+token.Token, nil)

	rs.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Service did not return ok status")
		t.FailNow()
	}

	if "application/json; charset=utf-8" != rr.Header().Get("Content-Type") {
		t.Fatalf("Service did not return json header")
		t.FailNow()
	}

	if "{\"user_id\":\"3\",\"scenario\":\"test\",\"token\":\""+token.Token+"\"}" != rr.Body.String() {
		t.Fatalf("Unexpected Response: %s", rr.Body.String())
		t.FailNow()
	}
}

func TestAddTokenWithMeta(t *testing.T) {
	body := strings.NewReader("{\"user_id\":\"1\",\"scenario\":\"test\",\"meta\":{\"alpha\":\"alfred\",\"beta\":\"bravo\"}}")
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

	if !strings.Contains(rr.Body.String(), "\"meta\":{\"alpha\":\"alfred\",\"beta\":\"bravo\"}") {
		t.Fatalf("Unexpected Response: %s", rr.Body.String())
		t.FailNow()
	}
}

func TestTokenExistsWithMeta(t *testing.T) {
	rs := NewRestServer()
	token := rs.Generator.Get(4, "test")
	token.Meta = make(map[string]string)
	token.Meta["alpha"] = "alfred"
	token.Meta["beta"] = "bravo"

	rs.Store.Add(token, time.Hour)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/4/"+token.Token, nil)

	rs.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Service did not return ok status")
		t.FailNow()
	}

	if "application/json; charset=utf-8" != rr.Header().Get("Content-Type") {
		t.Fatalf("Service did not return json header")
		t.FailNow()
	}

	if !strings.Contains(rr.Body.String(), "\"meta\":{\"alpha\":\"alfred\",\"beta\":\"bravo\"}") {
		t.Fatalf("Unexpected Response: %s", rr.Body.String())
		t.FailNow()
	}
}

func BenchmarkAddToken(b *testing.B) {
	rs := NewRestServer()
	for n := 0; n < b.N; n++ {
		body := strings.NewReader(fmt.Sprintf("{\"user_id\":\"%d\",\"scenario\":\"benchmark %d\"}", n, b.N))
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/", body)

		rs.Router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			b.Fatalf("Service did not return ok status")
			b.FailNow()
		}
	}
}
