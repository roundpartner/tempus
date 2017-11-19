package main

import (
	"crypto/rand"
	"encoding/hex"
)

type TokenGenerator struct {
	Tokens chan string
}

func NewTokenGenerator() *TokenGenerator {
	generator := &TokenGenerator{}
	generator.run()
	return generator
}

func (generator *TokenGenerator) run() {
	generator.Tokens = make(chan string, 50)
	go func() {
		for {
			generator.Tokens <- generator.randomToken()
		}
	}()
}

func (generator *TokenGenerator) Get(user int, scenario string) *Token {
	token := &Token{User: user, Scenario: scenario}
	token.Token = <- generator.Tokens
	return token
}

func (generator *TokenGenerator) randomToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
