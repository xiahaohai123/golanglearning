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

	file, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFilename, err)
	}

	store, err := poker.NewFileSystemPlayerStore(file)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	cliGame := poker.NewCLI(store, os.Stdin)
	cliGame.PlayPoker()
}
