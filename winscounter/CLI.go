package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{store: store, in: bufio.NewScanner(in)}
}

func (cli *CLI) PlayPoker() {
	line := cli.readLine()
	winner := cli.extractWinner(line)
	cli.store.RecordWin(winner)
}

func (CLI) extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
