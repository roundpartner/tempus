package main

import (
	"testing"
)

func TestGetTokenFromGenerator(t *testing.T) {
	generator := NewTokenGenerator()
	token := generator.Get(1, "test")
	if "" == token.Token {
		t.Fatalf("Token was empty")
		t.FailNow()
	}
	if len(token.Token) != 64 {
		t.Fatalf("Token was length of %d but was expecting 64", len(token.Token))
		t.FailNow()
	}
}

func TestGeneratedTokensAreUnique(t *testing.T) {
	generator := NewTokenGenerator()
	token1 := generator.Get(1, "test")
	token2 := generator.Get(1, "test")
	if token1.Token == token2.Token {
		t.Fatalf("Tokens are not unqiue")
		t.FailNow()
	}
}

func BenchmarkTokenGenerateFast(b *testing.B) {
	generator := NewTokenGenerator()
	b.StartTimer()
	for n := 0; n < 10; n++ {
		generator.Get(1, "test")
	}
	b.StopTimer()
}
