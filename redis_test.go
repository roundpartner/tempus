package main

import (
	"testing"
	"time"
)

func TestConnectToRedis(t *testing.T) {
	client := New()
	response, err := client.Ping()
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
	if "PONG" != response {
		t.Fatal("Unexpected response from Redis")
		t.FailNow()
	}
}

func TestStoreATokenInRedis(t *testing.T) {
	client := New()
	token := &Token{User: 1, Scenario: "email", Token: "hello_world"}
	err := client.Add(token, time.Minute)
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
}

func passValidation(_ *Token) bool {
	return true
}

func failValidation(_ *Token) bool {
	return false
}

func TestGetTokenFromRedis(t *testing.T) {
	client := New()
	_, err := client.Get("hello_world", passValidation)
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
}

func TestTokenExpiresAfterGet(t *testing.T) {
	client := New()
	token, err := client.Get("hello_world", passValidation)
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
	if token != nil {
		t.Fatalf("Token has not expired")
		t.FailNow()
	}
}

func TestTokenPersistsIfValidatorFails(t *testing.T) {
	client := New()
	token := &Token{User: 3, Scenario: "email", Token: "persistent_token"}
	client.Add(token, time.Minute)
	token, err := client.Get("persistent_token", failValidation)
	if err == nil {
		t.Fatalf("Expected error")
		t.FailNow()
	}
	if err.Error() != "token did not validate" {
		t.Fatalf("Token has not persisted")
		t.FailNow()
	}
}
