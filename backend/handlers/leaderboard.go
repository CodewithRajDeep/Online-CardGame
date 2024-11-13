package handlers

import (
    "cardgame-backend/redis"
    "net/http"
    "github.com/go-redis/redis/v8"
    "encoding/json"
    "context"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
    client := redis.NewRedisClient()
    defer client.Close()
    ctx := context.Background()

    leaderboard, err := client.ZRevRangeWithScores(ctx, "leaderboard", 0, 10).Result()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(leaderboard)
}

func UpdateLeaderboard(username string, score int64) error {
    client := redis.NewRedisClient()
    defer client.Close()
    ctx := context.Background()

    return client.ZIncrBy(ctx, "leaderboard", float64(score), username).Err()
}
