package main

import (
	"log"
	"net/http"
	"os"
)

const dbFilename = "game.db.json"

func main() {
	file, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFilename, err)
	}

	store, err := NewFileSystemPlayerStore(file)
	if err != nil {
		log.Fatalf("didnt expect an error but got one, %v", err)
	}

	server := NewPlayerServer(store)
	if err = http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
