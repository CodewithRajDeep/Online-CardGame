package redis

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,    
    })
    return client
}
