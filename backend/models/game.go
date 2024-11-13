package models

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)


type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}


type Game struct {
	Username   string `json:"username"`
	Wins       int    `json:"wins"`
	CurrentCard string `json:"current_card"`
	Deck       []Card `json:"deck"`      
	GameOver   bool   `json:"game_over"`  
	Message    string `json:"message"`  
}


var RedisClient *redis.Client


func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}


func (g *Game) SaveGame() error {
	ctx := context.Background()
	err := RedisClient.HSet(ctx, g.Username, map[string]interface{}{
		"username":    g.Username,
		"wins":        g.Wins,
		"currentCard": g.CurrentCard,
		"gameOver":    g.GameOver,
		"message":     g.Message,
	}).Err()
	return err
}


func GetGame(username string) (*Game, error) {
	ctx := context.Background()
	data, err := RedisClient.HGetAll(ctx, username).Result()
	if err != nil {
		return nil, err
	}

	
	wins, err := strconv.Atoi(data["wins"])
	if err != nil {
		wins = 0
	}

	game := &Game{
		Username:   username,
		Wins:       wins,
		CurrentCard: data["currentCard"],
		GameOver:   data["gameOver"] == "true",
		Message:    data["message"],
	}
	return game, nil
}


func SaveGameState(g *Game) error {
	return g.SaveGame()
}


func GetGameState(username string) (*Game, error) {
	return GetGame(username)
}
