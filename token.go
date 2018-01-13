package main

import (
	"bytes"
	"encoding/json"
	"time"
)

const (
	DAY      = time.Hour * 24
	PASSWORD = DAY * 3
	EMAIL    = DAY * 7
)

type Token struct {
	User     int64             `json:"user_id,string"`
	Scenario string            `json:"scenario"`
	Token    string            `json:"token"`
	Meta     map[string]string `json:"meta,omitempty"`
}

func (token *Token) MarshalBinary() (data []byte, err error) {
	return json.Marshal(token)
}

func (token *Token) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buffer)
	return decoder.Decode(token)
}

func (token *Token) Key() string {
	return "tempus_" + token.Token
}

func (token *Token) Expires() time.Duration {
	if "email" == token.Scenario {
		return EMAIL
	}
	return PASSWORD
}
