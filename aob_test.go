package main

import (
	"fmt"
	"testing"
	"time"
)

func TestItemInserted(t *testing.T) {
	bus = BackUpService{}
	bus.Run()
	token := &Token{1, "test", "Token"}

	err := ItemInserted("a_new_key", token, time.Minute)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
}

func TestEntryEncodesIntoJson(t *testing.T) {
	token := &Token{1, "test", "Token"}
	json, err := token.MarshalBinary()
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	if fmt.Sprintf("%s", json) != "{\"user_id\":\"1\",\"scenario\":\"test\",\"Token\":\"Token\"}" {
		t.Errorf("Entry did not encode: %s", json)
		t.FailNow()
	}
}
