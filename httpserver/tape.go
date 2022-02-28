package main

import (
	"io"
)

type ReadWriteSeekTruncate interface {
	io.ReadWriteSeeker
	Truncate(size int64) error
}

type Tape struct {
	file ReadWriteSeekTruncate
}

func (t *Tape) Write(p []byte) (n int, err error) {
	_ = t.file.Truncate(0)
	_, _ = t.file.Seek(0, 0)
	return t.file.Write(p)
}
