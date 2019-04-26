package main

import "testing"

func TestTokenNameIsTempus(t *testing.T) {
	token := &Token{User: 1, Scenario: "email", Token: "hello_world"}
	key := token.Key()
	if key != "tempus_hello_world" {
		t.Fatalf("Unexpected token name: %s", key)
		t.FailNow()
	}
}

func TestTokenImplementsBinaryMarshaler(t *testing.T) {
	token := &Token{User: 1, Scenario: "email", Token: "hello_world"}
	_, err := token.MarshalBinary()
	if err != nil {
		t.Fatalf("Error %s", err.Error())
		t.FailNow()
	}
}

func TestPasswordTokenSetsExpiry(t *testing.T) {
	token := &Token{User: 1, Scenario: "password", Token: "hello_world"}
	if token.Expires() != PASSWORD {
		t.Fatalf("Error: Token does not expire %d != %d", token.Expires(), PASSWORD)
		t.FailNow()
	}
}

func TestEmailTokenSetsExpiry(t *testing.T) {
	token := &Token{User: 1, Scenario: "email", Token: "hello_world"}
	if token.Expires() != EMAIL {
		t.Fatalf("Error: Token does not expire %d != %d", token.Expires(), EMAIL)
		t.FailNow()
	}
}

func TestRobotTokenSetsExpiry(t *testing.T) {
	token := &Token{User: 0, Scenario: "robot", Token: "hello_world"}
	if token.Expires() != ROBOT {
		t.Fatalf("Error: Token does not expire %d != %d", token.Expires(), ROBOT)
		t.FailNow()
	}
}
