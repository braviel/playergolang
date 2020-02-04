package main

import (
	"firstapp/playerserver"
	"log"
	"net/http"
)

func main() {
	s := &playerserver.PlayerServer{Store: playerserver.NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", s); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
