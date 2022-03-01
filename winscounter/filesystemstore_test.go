package poker_test

import (
	"io/ioutil"
	"os"
	poker "summersea.top/golanglearning/winscounter"
	"testing"
)

const initData = `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()
		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()
		// 排序
		want := []poker.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		assertPlayerSliceEquals(t, got, want)

		// read again
		got = store.GetLeague()
		assertPlayerSliceEquals(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()
		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

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
		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertIntEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()
		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertIntEquals(t, got, want)
	})
}

func TestNewFileSystemPlayerStore(t *testing.T) {
	t.Run("works with an empty file", func(t *testing.T) {
		file, removeFile := createTempFile(t, "")
		defer removeFile()

		_, err := poker.NewFileSystemPlayerStore(file)

		assertNoError(t, err)
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

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect an error but got one, %v", err)
	}
}
