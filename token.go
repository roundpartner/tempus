package main

import (
	"bytes"
	"encoding/json"
)

type Token struct {
	User     int64  `json:"user_id,string"`
	Scenario string `json:"scenario"`
	Token    string `json:"token"`
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
