package main

import (
	"os"
)

type Tape struct {
	file *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	_ = t.file.Truncate(0)
	_, _ = t.file.Seek(0, 0)
	return t.file.Write(p)
}
