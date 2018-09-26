package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var redisURL = "localhost:6379"
var redisPass = ""
var redisChannelA = "taylor"
var redisChannelB = "swift"

func main() {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPass,
		DB:       0, // use default DB
	})

	_, err := redisdb.Ping().Result()
	if err != nil {
		hasPass := (redisPass != "")
		fmt.Printf("Error: Unable to connect to redis [password: %v]\n  url: %v\n", hasPass, redisURL)
		os.Exit(1)
	}

	fmt.Printf("Connected to redis @ %v\n", redisURL)

	pubsub := redisdb.Subscribe(redisChannelA, redisChannelB)

	// Wait for confirmation that subscription is created before publishing anything.
	_, err = pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// When pubsub is closed channel is closed too.
	defer pubsub.Close()

	fmt.Printf("Listening on '%v' and '%v'...\n", redisChannelA, redisChannelB)
	// Consume messages.
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("Message received on '%v': %v\n", msg.Channel, msg.Payload)
	}
}
