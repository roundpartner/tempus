package main

import (
	"gopkg.in/redis.v3"
)

func connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func ping(client *redis.Client) (string, error) {
	pong, err := client.Ping().Result()
	return pong, err
}
