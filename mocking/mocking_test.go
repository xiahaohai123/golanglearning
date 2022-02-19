package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountdownOperationSpy struct {
	Calls []string
}

const sleep = "sleep"
const write = "write"

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {

	t.Run("print test", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
GO!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
		}
	})

	// 测试行为
	t.Run("sleep after every print", func(t *testing.T) {
		spyProcessPrinter := &CountdownOperationSpy{}
		Countdown(spyProcessPrinter, spyProcessPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spyProcessPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spyProcessPrinter.Calls)
		}
	})
}
