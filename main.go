package main

import (
	"context"
	"fmt"
	"time"

	"your-module-path/redisdriver"
)

func main() {
	// Create a new Redis client
	redisClient := redisdriver.NewRedisClient("localhost:6379", "", 0)

	// Use context for timeout/cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Set a value in Redis
	err := redisClient.Set(ctx, "my_key", "my_value", 10*time.Second)
	if err != nil {
		fmt.Println("Failed to set key:", err)
		return
	}

	// Get the value from Redis
	value, err := redisClient.Get(ctx, "my_key")
	if err != nil {
		fmt.Println("Failed to get key:", err)
		return
	}

	fmt.Println("Value from Redis:", value)

	// Delete the key
	err = redisClient.Del(ctx, "my_key")
	if err != nil {
		fmt.Println("Failed to delete key:", err)
		return
	}

	fmt.Println("Key deleted successfully")
}
