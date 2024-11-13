package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"
	"time"
)


type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}


type Game struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Deck        []Card  `json:"deck"`
	CurrentCard string  `json:"current_card"`
	GameOver    bool    `json:"game_over"`
	Message     string  `json:"message"`
}


func StartGame(w http.ResponseWriter, r *http.Request) {
	
	game := Game{
		ID:   "1",
		Name: "Poker",
		Deck: generateDeck(), 
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}


func generateDeck() []Card {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	var deck []Card
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{Value: value, Suit: suit})
		}
	}
	
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}


func DrawCard(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	
	game, err := GetGameState(username)
	if err != nil {
		http.Error(w, "Failed to retrieve game state", http.StatusInternalServerError)
		return
	}

	if len(game.Deck) == 0 {
		game.GameOver = true
		game.Message = "No more cards left in the deck"
	} else {
		
		drawnCard := game.Deck[len(game.Deck)-1]
		game.Deck = game.Deck[:len(game.Deck)-1]
		game.CurrentCard = drawnCard.Suit + " " + drawnCard.Value

		
		game.Message = "You drew: " + game.CurrentCard
	}

	
	if err := SaveGameState(game); err != nil {
		http.Error(w, "Failed to save game state", http.StatusInternalServerError)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}


func GetGameState(username string) (*Game, error) {
	
	return &Game{
		ID:          "1",
		Name:        "Poker",
		Deck:        generateDeck(), 
		CurrentCard: "",
		GameOver:    false,
		Message:     "Game started",
	}, nil
}


func SaveGameState(game *Game) error {
	
	fmt.Printf("Game state saved: %+v\n", game)
	return nil
}
