package main

import (
	"github.com/go-redis/redis"
)

func getClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	return client, err
}

func setValue(key string, val string) error {
	client, err := getClient()
	if err != nil {
		return err
	}
	err = client.Set(key, val, 0).Err()
	return err
}

func getValue(key string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", err
	}
	return client.Get(key).Result()
}
