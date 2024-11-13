package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
    "github.com/CodewithRajDeep/Online-CardGame/backend/handlers"
 )

func main() {
	r := mux.NewRouter()

	
	r.HandleFunc("/game", handlers.CreateGame).Methods("POST")
	r.HandleFunc("/draw", handlers.DrawCard).Methods("GET")
	r.HandleFunc("/leaderboard", handlers.GetLeaderboard).Methods("GET")
	r.HandleFunc("/leaderboard/{player}", handlers.UpdateLeaderboard).Methods("POST")

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
