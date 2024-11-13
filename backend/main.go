package main

import (
    "cardgame-backend/handlers"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/api/start", handlers.StartGame).Methods("POST")
    r.HandleFunc("/api/draw", handlers.DrawCard).Methods("GET")
    r.HandleFunc("/api/leaderboard", handlers.GetLeaderboard).Methods("GET")

    http.ListenAndServe(":8080", r)
}
