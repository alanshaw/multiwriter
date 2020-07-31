// Package multiwriter contains an implementation of an "io.Writer" that
// duplicates it's writes to all the provided writers, similar to the Unix
// tee(1) command. Writers can be added and removed dynamically after creation.
//
// Example:
//
// 	package main
//
// 	import (
// 		"os"
// 		"github.com/alanshaw/multiwriter"
// 	)
//
// 	func main() {
// 		w := multiwriter.New(os.Stdout, os.Stderr)
//
// 		w.Write([]byte("written to stdout AND stderr\n"))
//
// 		w.Remove(os.Stderr)
//
// 		w.Write([]byte("written to ONLY stdout\n"))
//
// 		w.Remove(os.Stdout)
// 		w.Add(os.Stderr)
//
// 		w.Write([]byte("written to ONLY stderr\n"))
// 	}
package multiwriter

import (
	"io"
	"sync"
)

// MultiWriter is a writer that writes to multiple other writers.
type MultiWriter struct {
	sync.RWMutex
	writers []io.Writer
}

// New creates a writer that duplicates its writes to all the provided writers,
// similar to the Unix tee(1) command. Writers can be added and removed
// dynamically after creation.
//
// Each write is written to each listed writer, one at a time. If a listed
// writer returns an error, that overall write operation stops and returns the
// error; it does not continue down the list.
func New(writers ...io.Writer) *MultiWriter {
	mw := &MultiWriter{writers: writers}
	return mw
}

// Write writes some bytes to all the writers.
func (mw *MultiWriter) Write(p []byte) (n int, err error) {
	mw.RLock()
	defer mw.RUnlock()

	for _, w := range mw.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}

		if n < len(p) {
			err = io.ErrShortWrite
			return
		}
	}

	return len(p), nil
}

// Add appends a writer to the list of writers this multiwriter writes to.
func (mw *MultiWriter) Add(w io.Writer) {
	mw.Lock()
	mw.writers = append(mw.writers, w)
	mw.Unlock()
}

// Remove will remove a previously added writer from the list of writers.
func (mw *MultiWriter) Remove(w io.Writer) {
	mw.Lock()
	var writers []io.Writer
	for _, ew := range mw.writers {
		if ew != w {
			writers = append(writers, ew)
		}
	}
	mw.writers = writers
	mw.Unlock()
}
