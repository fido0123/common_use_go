package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type valueEx struct {
	Name  string
	Email string
	Tel   string
}

var (
	client = &redisClient{}
)

type redisClient struct {
	c *redis.Client
}

//GetClient get the redis client for init
func initialize() *redisClient {
	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	if err := c.Ping(context.TODO()).Err(); err != nil {
		panic("Unable to connect to to the nnredis " + err.Error())
	}
	client.c = c
	return client
}

//GetKey get key to fetch
func (client *redisClient) getKey(key string, src interface{}) error {
	val, err := client.c.Get(context.TODO(), key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	return err
}

//SetKey set key
func (client *redisClient) setKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.c.Set(context.TODO(), key, cacheEntry, expiration).Err()

	return err
}
func main() {
	redisClient := initialize()
	key1 := "sampleKey"
	value1 := &valueEx{Name: "someName", Email: "  dev1@abc.com"}
	// err := redisClient.setKey(key1, value1, time.Minute*5)
	// if err != nil {
	// 	log.Fatalf("Error: %v", err.Error())
	// }

	value2 := &valueEx{}
	err := redisClient.getKey(key1, value2)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
	log.Printf("Name: %s", value2.Name)
	log.Printf("Email: %s", value2.Email)

}
