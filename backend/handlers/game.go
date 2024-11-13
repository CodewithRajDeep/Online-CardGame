package handlers

import (
    "encoding/json"
    "net/http"
    "log"
    "math/rand"
    "time"
)


type Game struct {
    ID        string  `json:"id"`
    Name      string  `json:"name"`
    Deck      []Card  `json:"deck"`
    CurrentCard string `json:"current_card"`
    GameOver  bool    `json:"game_over"`
    Message   string  `json:"message"`
}


type Card struct {
    Suit  string `json:"suit"`
    Value string `json:"value"`
}


func CreateGame(w http.ResponseWriter, r *http.Request) {
    
    game := Game{
        ID:   "1",
        Name: "Poker",
        Deck: initializeDeck(),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(game)
}


func DrawCard(w http.ResponseWriter, r *http.Request) {
   
    username := "Player1"

    game := Game{
        ID:   "1",
        Name: "Poker",
        Deck: initializeDeck(),
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

 
    log.Printf("Player %s drew a card: %s", username, game.CurrentCard)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(game)
}


func initializeDeck() []Card {
    suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
    values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

    var deck []Card
    for _, suit := range suits {
        for _, value := range values {
            deck = append(deck, Card{Suit: suit, Value: value})
        }
    }

   
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(deck), func(i, j int) {
        deck[i], deck[j] = deck[j], deck[i]
    })

    return deck
}
