package main

import (
	"errors"
	"fmt"
	"gopkg.in/redis.v3"
	"os"
	"time"
)

type Store struct {
	client *redis.Client
	Delete chan *Token
}

func New() *Store {
	host := os.Getenv("REDIS_HOST")

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "",
		DB:       0,
	})
	store := &Store{
		client: client,
	}
	store.run()
	return store
}

func (store *Store) run() {
	store.Delete = make(chan *Token, 50)
	go func() {
		for {
			token := <-store.Delete
			store.client.Del(token.Key())
		}
	}()
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

func (store *Store) Get(id string, fn validate) (*Token, error) {
	token, err := store.get(id)
	if err != nil || token == nil {
		return nil, err
	}
	if !fn(token) {
		return nil, errors.New("token did not validate")
	}
	store.Delete <- token
	return token, err
}

func (store *Store) get(id string) (*Token, error) {
	token := Token{Token: id}
	exists, err := store.client.Exists(token.Key()).Result()
	if !exists {
		return nil, nil
	}
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
