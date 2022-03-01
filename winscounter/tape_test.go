package poker

import (
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "123456")
	defer clean()

	tape := &Tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)

	newFileContent, _ := ioutil.ReadAll(file)

	got := string(newFileContent)
	want := "abc"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
