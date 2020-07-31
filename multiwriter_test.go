package multiwriter

import (
	"bytes"
	"testing"
)

type testWriter struct {
	lastChunk []byte
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	w.lastChunk = p
	return len(p), nil
}

func assertEqual(t *testing.T, expected []byte, actual []byte) {
	t.Helper()
	if !bytes.Equal(expected, actual) {
		t.Fatal("bytes not equal", expected, actual)
	}
}

func TestMultiWriter(t *testing.T) {
	w0 := testWriter{}
	w1 := testWriter{}

	mw := New(&w0, &w1)

	mw.Write([]byte("first"))

	assertEqual(t, w0.lastChunk, []byte("first"))
	assertEqual(t, w1.lastChunk, []byte("first"))

	mw.Remove(&w0)
	mw.Write([]byte("second"))

	assertEqual(t, w0.lastChunk, []byte("first"))
	assertEqual(t, w1.lastChunk, []byte("second"))

	mw.Remove(&w1)
	mw.Add(&w0)
	mw.Write([]byte("third"))

	assertEqual(t, w0.lastChunk, []byte("third"))
	assertEqual(t, w1.lastChunk, []byte("second"))
}
