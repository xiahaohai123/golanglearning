package main

import (
	"fmt"
	"log"
	"os"
	poker "summersea.top/golanglearning/winscounter"
)

const dbFilename = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	store, err := poker.FileSystemPlayerStoreFromFile(dbFilename)

	if err != nil {
		log.Fatal(err)
	}

	cliGame := poker.NewCLI(store, os.Stdin)
	cliGame.PlayPoker()
}
