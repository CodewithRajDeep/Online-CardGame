package handler

import (
	"encoding/json"
	"net/http"
)

// DrawCard handles the card drawing logic
func DrawCard(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	
	game, err := models.GetGameState(username)
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


	if err := game.SaveGameState(); err != nil {
		http.Error(w, "Failed to save game state", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}
