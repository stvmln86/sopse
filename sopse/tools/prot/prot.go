// Package prot implements HTTP protocol functions.
package prot

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// headers is a map of default HTTP response headers.
var headers = map[string]string{
	"Cache-Control":          "no-store",
	"Content-Type":           "text/plain; charset=utf-8",
	"X-Content-Type-Options": "nosniff",
}

// Read returns a Request's body as a string.
func Read(r *http.Request) string {
	if r.Body == nil {
		return ""
	}

	bytes, _ := io.ReadAll(r.Body)
	return string(bytes)
}

// Write writes a formatted text/plain string to a ResponseWriter.
func Write(w http.ResponseWriter, code int, text string, elems ...any) {
	for attr, data := range headers {
		w.Header().Set(attr, data)
	}

	w.WriteHeader(code)
	fmt.Fprintf(w, text, elems...)
}

// WriteCode writes a text/plain error code to a ResponseWriter.
func WriteCode(w http.ResponseWriter, code int) {
	stat := http.StatusText(code)
	stat = strings.ToLower(stat)
	Write(w, code, "error %d: %s", code, stat)
}

// WriteError writes a formatted text/plain error message to a ResponseWriter.
func WriteError(w http.ResponseWriter, code int, text string, elems ...any) {
	text = fmt.Sprintf(text, elems...)
	Write(w, code, "error %d: %s", code, text)
}
