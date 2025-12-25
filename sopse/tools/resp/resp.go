// Package resp implements HTTP response functions.
package resp

import (
	"fmt"
	"net/http"
	"strings"
)

// Write writes a formatted text/plain string to a ResponseWriter.
func Write(w http.ResponseWriter, code int, text string, elems ...any) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, text, elems...)
}

// WriteCode writes a text/plain error code to a ResponseWriter.
func WriteCode(w http.ResponseWriter, code int) {
	text := http.StatusText(code)
	text = strings.ToLower(text)
	Write(w, code, "error %d: %s", code, text)
}

// WriteError writes a formatted text/plain error string to a Response Writer.
func WriteError(w http.ResponseWriter, code int, text string, elems ...any) {
	text = fmt.Sprintf(text, elems...)
	Write(w, code, "error %d: %s", code, text)
}
