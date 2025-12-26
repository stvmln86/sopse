// Package prot implements HTTP protocol functions.
package prot

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Read returns the body string from a Request.
func Read(r *http.Request) (string, error) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Write writes a text/plain string to a ResponseWriter.
func Write(w http.ResponseWriter, code int, text string) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write([]byte(text))
}

// WriteCode writes a text/plain error code to a ResponseWriter.
func WriteCode(w http.ResponseWriter, code int) {
	stat := strings.ToLower(http.StatusText(code))
	text := fmt.Sprintf("error %d: %s", code, stat)
	Write(w, code, text)
}

// WriteError writes a text/plain error string to a ResponseWriter.
func WriteError(w http.ResponseWriter, code int, text string) {
	text = fmt.Sprintf("error %d: %s", code, text)
	Write(w, code, text)
}
