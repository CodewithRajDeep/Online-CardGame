package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)


var rdb *redis.Client

func init() {
	
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",  
		DB: 0,             
	})
}


func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	
	ctx := context.Background()

	
	leaderboard, err := rdb.ZRevRange(ctx, "leaderboard", 0, 9).Result()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get leaderboard: %v", err), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
}

func AddToLeaderboard(playerName string, score int) error {

	ctx := context.Background()
	_, err := rdb.ZAdd(ctx, "leaderboard", &redis.Z{
		Score:  float64(score),
		Member: playerName,
	}).Result()
	return err
}

func UpdateLeaderboard(playerName string, newScore int) error {
	
	ctx := context.Background()

	
	_, err := rdb.ZAdd(ctx, "leaderboard", &redis.Z{
		Score:  float64(newScore),
		Member: playerName,
	}).Result()
	return err
}


func ClearLeaderboard(w http.ResponseWriter, r *http.Request) {
	
	ctx := context.Background()

	
	_, err := rdb.Del(ctx, "leaderboard").Result()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to clear leaderboard: %v", err), http.StatusInternalServerError)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Leaderboard cleared successfully",
	})
}
