package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/yourproject/models"
)

// DrawCard handles the card drawing logic
func DrawCard(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Retrieve the game state
	game, err := models.GetGameState(username)
	if err != nil {
		http.Error(w, "Failed to retrieve game state", http.StatusInternalServerError)
		return
	}

	// Check if there are cards left in the deck
	if len(game.Deck) == 0 {
		game.GameOver = true
		game.Message = "No more cards left in the deck"
	} else {
		// Draw the top card from the deck
		drawnCard := game.Deck[len(game.Deck)-1]
		game.Deck = game.Deck[:len(game.Deck)-1]
		game.CurrentCard = drawnCard.Suit + " " + drawnCard.Value // Convert Card to string

		// Update the message
		game.Message = "You drew: " + game.CurrentCard
	}

	// Save updated game state
	if err := game.SaveGameState(); err != nil {
		http.Error(w, "Failed to save game state", http.StatusInternalServerError)
		return
	}

	// Respond with the updated game state
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}
