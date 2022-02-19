package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := ConfigurableSleeper{duration: 1 * time.Second}
	Countdown(os.Stdout, &sleeper)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

const countdownStart = 3
const finalWord = "GO!"

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, _ = fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	_, _ = fmt.Fprint(writer, finalWord)
}
