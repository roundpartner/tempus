package main

import (
	"testing"
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
