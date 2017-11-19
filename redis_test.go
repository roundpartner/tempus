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
	token := &Token{1, "email", "hello_world"}
	err := client.Add(token, time.Second)
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
}

func TestGetTokenFromRedis(t *testing.T) {
	client := New()
	_, err := client.Get("hello_world")
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
}
