package main

import "testing"

func TestValidatorPasses(t *testing.T) {
	v := UserValidator(1)
	tok := Token{User: 1}

	if !v(&tok) {
		t.Errorf("Valid token did not validate")
		t.FailNow()
	}
}

func TestValidatorFails(t *testing.T) {
	v := UserValidator(1)
	tok := Token{User: 2}

	if v(&tok) {
		t.Errorf("Invalid token validated")
		t.FailNow()
	}
}
