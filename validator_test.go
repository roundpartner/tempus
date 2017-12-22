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

func TestValidatorScenarioPasses(t *testing.T) {
	v := UserScenarioValidator(1, "testing")
	tok := Token{User: 1, Scenario: "testing"}

	if !v(&tok) {
		t.Errorf("Valid token did not validate")
		t.FailNow()
	}
}

func TestValidatorScenarioFails(t *testing.T) {
	v := UserScenarioValidator(1, "invalid")
	tok := Token{User: 1, Scenario: "testing"}

	if v(&tok) {
		t.Errorf("Invalid token validated")
		t.FailNow()
	}
}
