package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const initData = `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertPlayerSliceEquals(t, got, want)

		// read again
		got = store.GetLeague()
		assertPlayerSliceEquals(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Cleo")
		want := 10
		assertIntEquals(t, got, want)

		//read again
		got = store.GetPlayerScore("Cleo")
		assertIntEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertIntEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertIntEquals(t, got, want)
	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tempFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, writeError := tempFile.Write([]byte(initialData))
	if writeError != nil {
		t.Fatalf("could not write data to temp file %v", err)
	}

	removeFile := func() {
		_ = os.Remove(tempFile.Name())
	}
	return tempFile, removeFile
}
