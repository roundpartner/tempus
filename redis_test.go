package main

import (
	"github.com/alicebob/miniredis/v2"
	"os"
	"testing"
	"time"
)

func TestConnectToRedis(t *testing.T) {
	MockRedisServer()
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
	MockRedisServer()
	client := New()
	token := &Token{User: 1, Scenario: "email", Token: "hello_world"}
	err := client.Add(token, time.Minute)
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
}

func TestStoreATokenInRedisLater(t *testing.T) {
	MockRedisServer()
	client := New()
	token := &Token{User: 1, Scenario: "email", Token: t.Name()}
	err := client.AddLater(token, time.Minute)
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
	_, err = client.Get(token.Token, passValidation)
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
	MockRedisServer()
	client := New()
	_, err := client.Get("hello_world", passValidation)
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
}

func TestTokenExpiresAfterGet(t *testing.T) {
	MockRedisServer()
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
	MockRedisServer()
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

func MockRedisServer() *miniredis.Miniredis {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	os.Setenv("REDIS_HOST", mr.Addr())
	return mr
}
