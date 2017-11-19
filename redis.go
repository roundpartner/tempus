package main

import (
	"gopkg.in/redis.v3"
	"time"
	"fmt"
)

type Store struct {
	client *redis.Client
}

func New() *Store {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Store{
		client: client,
	}
}

func (store *Store) Ping() (string, error) {
	pong, err := store.client.Ping().Result()
	return pong, err
}

func (store *Store) Add(token *Token, duration time.Duration) error {
	success, err := store.client.SetNX(token.Key(), token, duration).Result()
	if err != nil {
		return err
	}
	if !success {
		return fmt.Errorf("token was not stored")
	}
	return nil
}

func (store *Store) Get(id string) (*Token, error) {
	token := Token{Token: id}
	data, err := store.client.Get(token.Key()).Result()
	if err != nil {
		return nil, err
	}
	return store.stringToToken(data)
}

func (store *Store) stringToToken(data string) (*Token, error) {
	token := &Token{}
	b := []byte(data)
	err := token.UnmarshalBinary(b)
	if err != nil {
		return nil, err
	}
	return token, nil
}
