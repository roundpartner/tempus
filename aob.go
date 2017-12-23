package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

var bus BackUpService

type BackUpService struct {
	entries chan *aobEntry
}

func (bus *BackUpService) Run() {
	bus.entries = make(chan *aobEntry, 50)
	go func() {
		for {
			entry := <-bus.entries
			data, err := entry.MarshalBinary()
			if err != nil {
				log.Printf("Unable to marshal aob entry: %s", err.Error())
				continue
			}
			err = appendToBackUp(data)
			if err != nil {
				log.Printf("Unable to append aob entry to file: %s", err.Error())
				continue
			}
		}
	}()
}

type aobEntry struct {
	Key      string        `json:"key"`
	Token    *Token        `json:"token"`
	Duration time.Duration `json:"duration"`
}

func (entry *aobEntry) MarshalBinary() (data []byte, err error) {
	return json.Marshal(entry)
}

func ItemInserted(key string, token *Token, duration time.Duration) error {
	entry := &aobEntry{Key: key, Token: token, Duration: duration}
	bus.entries <- entry
	return nil
}

func appendToBackUp(data []byte) error {
	f, err := os.OpenFile("test.aob", os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()
	f.Write(data)
	f.WriteString("\n")
	return err
}
