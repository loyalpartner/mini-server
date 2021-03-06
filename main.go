package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (this *StubPlayerStore) GetPlayerScore(name string) int {
	return this.scores[name]
}

func (this *StubPlayerStore) RecordWin(name string) {
	this.winCalls = append(this.winCalls, name)
}

type PlayerServer struct {
	store PlayerStore
}

func (this *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		this.processWin(w, player)
	case http.MethodGet:
		this.showScore(w, player)
	}
}

func (this *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := this.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, this.store.GetPlayerScore(player))
}

func (this *PlayerServer) processWin(w http.ResponseWriter, player string) {
	this.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (this *InMemoryPlayerStore) RecordWin(name string) {
	this.store[name]++
}

func (this *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return this.store[name]
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
