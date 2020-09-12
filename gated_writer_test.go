package logger_test

import (
	"bytes"
	"io"
	"testing"

	logger "github.com/da-moon/go-logger"
)

var _ io.Writer = &logger.GatedWriter{}

// TestGatedWriter ...
func TestGatedWriter(t *testing.T) {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	w := logger.NewGatedWriter(buf)
	w.Write([]byte("foo\n"))
	w.Write([]byte("bar\n"))

	if buf.String() != "" {
		t.Fatalf("bad: %s", buf.String())
	}

	w.Flush()

	if buf.String() != "foo\nbar\n" {
		t.Fatalf("bad: %s", buf.String())
	}

	w.Write([]byte("baz\n"))

	if buf.String() != "foo\nbar\nbaz\n" {
		t.Fatalf("bad: %s", buf.String())
	}
}
