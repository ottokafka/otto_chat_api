package main

import (
	// This json package allows us to encode string into a json format
	"fmt"

	"github.com/go-redis/redis"
)

var Client = redis.NewClient(&redis.Options{
	// Addr:     "10.90.24.205:6379", // TeamCity server
	Addr:     "localhost:6379", // TeamCity server
	Password: "",               // no password set
	DB:       0,                // use default DB
})

// fmt.Println("redis connected")

func RedisTest() {
	redis.NewClient(&redis.Options{
		// Addr:     "10.90.24.205:6379", // TeamCity server
		Addr:     "localhost:6379", // TeamCity server
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	fmt.Println("connected to redis")
}

func ExampleClient() {

	//Save to Redis
	err := Client.Set("myKey", "money", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get from Redis
	val, err := Client.Get("myKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("myKey", val)
}

// err := Client.Set("myKey", "money", 0).Err()
// if err != nil {
// 	panic(err)
// }

// val, err := Client.Get("myKey").Result()
// if err != nil {
// 	panic(err)
// }
// fmt.Println("myKey", val)

// val2, err := Client.Get("key2").Result()
// if err == redis.Nil {
// 	fmt.Println("key2 does not exist")
// } else if err != nil {
// 	panic(err)
// } else {
// 	fmt.Println("key2", val2)
// }
