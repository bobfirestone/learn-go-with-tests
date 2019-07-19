package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

// PlayerServer struct
type PlayerServer struct {
	store PlayerStore
}

// ServeHTTP allows PlayerServer handles the requests
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// PlayerStore interface
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// InMemoryPlayerStore temp store values in memory
type InMemoryPlayerStore struct{}

// GetPlayerScore retreives the players score
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	// if name == "Pepper" {
	// 	return "20"
	// }

	// if name == "Floyd" {
	// 	return "10"
	// }

	return 123
}
