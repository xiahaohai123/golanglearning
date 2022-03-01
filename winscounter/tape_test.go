package poker_test

import (
	"io/ioutil"
	poker "summersea.top/golanglearning/winscounter"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "123456")
	defer clean()

	tape := poker.NewTape(file)

	tape.Write([]byte("abc"))

	file.Seek(0, 0)

	newFileContent, _ := ioutil.ReadAll(file)

	got := string(newFileContent)
	want := "abc"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
