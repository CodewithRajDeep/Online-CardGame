package models

type LeaderboardEntry struct {
  Username string  `json:"username"`
  Score    float64 `json:"score"`  
}
