package main

import "testing"

func TestTokenNameIsTempus(t *testing.T) {
	token := &Token{1, "email", "hello_world"}
	key := token.Key()
	if key != "tempus_hello_world" {
		t.Fatalf("Unexpected token name: %s", key)
		t.FailNow()
	}
}

func TestTokenImplementsBinaryMarshaler(t *testing.T) {
	token := &Token{1, "email", "hello_world"}
	_, err := token.MarshalBinary()
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
}
