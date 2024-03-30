package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/game", getGames)
	http.HandleFunc("/level", getLevels)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
