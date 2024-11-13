package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
    "github.com/CodewithRajDeep/Online-CardGame/backend/handlers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "Hello from Go backend!")
}


func main() {
	
	handlers.InitializeRedis()
	http.HandleFunc("/", Handler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
	
	
	r := mux.NewRouter()

	
	r.HandleFunc("/game", handlers.CreateGame).Methods("POST") 
	r.HandleFunc("/draw", handlers.DrawCard).Methods("GET")   
	r.HandleFunc("/game/draw", handlers.DrawCard).Methods("POST")
	r.HandleFunc("/leaderboard", handlers.GetLeaderboard).Methods("GET") 
	r.HandleFunc("/leaderboard/{player}", handlers.UpdateLeaderboard).Methods("POST") 
	r.HandleFunc("/game/leaderboard", handlers.UpdateLeaderboard).Methods("POST") 
	r.HandleFunc("/game/leaderboard", handlers.GetLeaderboard).Methods("GET") 
	r.HandleFunc("/game/leaderboard", handlers.GetLeaderboard).Methods("GET")
	r.HandleFunc("/game/leaderboard", handlers.UpdateLeaderboard).Methods("POST")
	r.HandleFunc("/leaderboard", handlers.GetLeaderboard).Methods("GET")
    r.HandleFunc("/leaderboard", handlers.UpdateLeaderboard).Methods("POST")

	r.Use(handlers.CORS)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
