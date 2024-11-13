package utils

import "cardgame-backend/models"


var cardTypes = []models.Card{
    {Type: "Cat", Icon: "😼"},
    {Type: "Defuse", Icon: "🙅‍♂️"},
    {Type: "Shuffle", Icon: "🔀"},
    {Type: "Exploding Kitten", Icon: "💣"},
}


func InitializeDeck() []models.Card {
    deck := make([]models.Card, 0, 20)
    for i := 0; i < 5; i++ {
        deck = append(deck, cardTypes[i%len(cardTypes)])
    }
    return deck
}
