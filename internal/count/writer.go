package count

import (
	"io"
)

// Writer counts the written bytes
type Writer struct {
	W io.Writer
	N int64
}

// NewWriter generates a new writer
func NewWriter(w io.Writer) *Writer {
	return &Writer{W: w}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	n, err = w.W.Write(p)
	w.N += int64(n)
	return
}
