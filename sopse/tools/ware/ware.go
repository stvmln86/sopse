// Package ware implements HTTP middleware functions.
package ware

import (
	"log"
	"net/http"
	"time"
)

// Apply applies all middleware to a HandlerFunc.
func Apply(next http.HandlerFunc) http.Handler {
	return LogWare(next)
}

// logWriter is a custom ResponseWriter for logging middleware.
type logWriter struct {
	http.ResponseWriter
	Code int
	Size int
}

// WriteHeader writes and records a status code.
func (w *logWriter) WriteHeader(code int) {
	w.Code = code
	w.ResponseWriter.WriteHeader(code)
}

// Write writes and records a byte slice.
func (w *logWriter) Write(bytes []byte) (int, error) {
	n, err := w.ResponseWriter.Write(bytes)
	w.Size += n
	return n, err
}

// LogWare is a middleware for logging HTTP requests.
func LogWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		init := time.Now()
		wrap := &logWriter{w, http.StatusOK, 0}
		next.ServeHTTP(wrap, r)
		secs := time.Since(init).Seconds()
		log.Printf(
			"%s %s %s :: %d %d %.5f",
			r.RemoteAddr, r.Method, r.URL.Path, wrap.Code, wrap.Size, secs,
		)
	})
}
