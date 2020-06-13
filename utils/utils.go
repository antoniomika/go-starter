// Package utils implements utilities for the application.
package utils

import (
	"fmt"
	"io"
	"log"
	"time"
)

var (
	// multiWriter is the writer that can be used for writing to multiple locations.
	multiWriter io.Writer
)

// LogWriter represents a writer that is used for writing logs in multiple locations.
type LogWriter struct {
	TimeFmt     string
	MultiWriter io.Writer
}

// Write implements the write function for the LogWriter. It will add a time in a
// specific format to logs.
func (w LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Fprintf(w.MultiWriter, "%v | %s", time.Now().Format(w.TimeFmt), string(bytes))
}

// Setup main utils. This initializes, whitelists, blacklists,
// and log writers.
func Setup(logWriter io.Writer) {
	multiWriter = logWriter
	_, err := multiWriter.Write([]byte("Test"))
	if err != nil {
		log.Println("Error:", err)
	}
}
