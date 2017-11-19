package main

import (
	"gopkg.in/redis.v3"
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
