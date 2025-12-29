// Package prot implements HTTP protocol functions.
package prot

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Read returns a Request's body as a string.
func Read(w http.ResponseWriter, r *http.Request, size int64) string {
	if r.Body == nil {
		return ""
	}

	r.Body = http.MaxBytesReader(w, r.Body, size)
	bytes, _ := io.ReadAll(r.Body)
	return strings.TrimSpace(string(bytes))
}

// Write writes a text/plain string to a ResponseWriter.
func Write(w http.ResponseWriter, code int, text string) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write([]byte(text))
}

// WriteError writes a text/plain error code or string to a ResponseWriter.
func WriteError(w http.ResponseWriter, code int, texts ...string) {
	if len(texts) == 0 {
		stat := http.StatusText(code)
		texts = append(texts, strings.ToLower(stat))
	}

	text := fmt.Sprintf("error %d: %s", code, texts[0])
	Write(w, code, text)
}
