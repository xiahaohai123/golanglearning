package main

import (
	"log"
	"net/http"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}