package main

import (
	"encoding/json"
	"net/http"
	models "speedboard/internal"
	"time"

	"github.com/google/uuid"
)

var nowUTC time.Time = time.Now().UTC()

var games = []models.Game{{
	ID:    uuid.New(),
	Title: "Goldeneye 007",
	Description: `GoldenEye 007 is a 1997 first-person shooter video game 
	developed by Rare and published by Nintendo for the Nintendo 64. 
	Based on the 1995 James Bond film GoldenEye, the player controls 
	the secret agent James Bond to prevent a criminal syndicate from 
	using a satellite weapon.`,
	ReleasedAt: time.Date(1997, time.August, 23, 0, 0, 0, 0, time.UTC),
	CreatedAt:  nowUTC,
	UpdatedAt:  nowUTC,
}}

var levels = []models.Level{{
	ID:   uuid.New(),
	Name: "Dam",
	Description: `Dam is the first mission in GoldenEye 007. 
	It is located at the Byelomorye dam, which serves as the setting 
	for the first part of Bond's first mission in GoldenEye.`,
	GameId:    games[0].ID,
	CreatedAt: nowUTC,
	UpdatedAt: nowUTC,
}, {
	ID:   uuid.New(),
	Name: "Facility",
	Description: `Facility is the second mission of GoldenEye 007. 
	Its full name is Chemical Warfare Facility #2, and it is the 
	second part of Bond's first mission.`,
	GameId:    games[0].ID,
	CreatedAt: nowUTC,
	UpdatedAt: nowUTC,
}, {
	ID:   uuid.New(),
	Name: "Runway",
	Description: `Runway is a level in GoldenEye 007. The level is 
	located at the service runway behind the Russian Chemical Warfare 
	Facility #2 and is most likely used by planes/trucks shipping the 
	nerve gas produced in the facility.`,
	GameId:    games[0].ID,
	CreatedAt: nowUTC,
	UpdatedAt: nowUTC,
}}

func getGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func getLevels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(levels)
}
