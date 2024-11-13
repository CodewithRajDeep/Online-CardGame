package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/go-redis/redis/v8"
    "context"
    "log"
)

var rdb *redis.Client
var ctx = context.Background()


func InitializeRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379", 
    })

    
    if err := rdb.Ping(ctx).Err(); err != nil {
        log.Fatalf("Redis connection failed: %v", err)
    }
}


func UpdateLeaderboard(w http.ResponseWriter, r *http.Request) {
    var requestData struct {
        PlayerName string `json:"player_name"`
        NewScore   int    `json:"new_score"`
    }

   
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    
    err := rdb.ZAdd(ctx, "leaderboard", &redis.Z{
		Score:  float64(requestData.NewScore),
		Member: requestData.PlayerName,
	}).Err()
    if err != nil {
        http.Error(w, "Failed to update leaderboard", http.StatusInternalServerError)
        return
    }

    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}


func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
    leaderboard, err := rdb.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
    if err != nil {
        http.Error(w, "Failed to retrieve leaderboard", http.StatusInternalServerError)
        return
    }

    
    var response []struct {
        PlayerName string  `json:"player_name"`
        Score      float64 `json:"score"`
    }

    for _, item := range leaderboard {
        response = append(response, struct {
            PlayerName string  `json:"player_name"`
            Score      float64 `json:"score"`
        }{
            PlayerName: item.Member.(string),
            Score:      item.Score,
        })
    }

   
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
