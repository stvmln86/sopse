// Package ware implements HTTP middleware functions.
package ware

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/stvmln86/sopse/sopse/tools/neat"
	"github.com/stvmln86/sopse/sopse/tools/prot"
	"golang.org/x/time/rate"
)

// RateAddrs is a map of active rate limits.
var RateAddrs = make(map[string]*rate.Limiter)

// RateMutex is a mutex for RateAddrs.
var RateMutex = new(sync.Mutex)

// Apply applies all middleware to a HandlerFunc.
func Apply(next http.HandlerFunc, rate int) http.Handler {
	return LogWare(RateWare(http.HandlerFunc(next), rate))
}

// logWriter is a custom ResponseWriter for LogWare.
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

// Write writes a byte slice and records the size.
func (w *logWriter) Write(bytes []byte) (int, error) {
	n, err := w.ResponseWriter.Write(bytes)
	w.Size += n
	return n, err
}

// LogWare is a middleware that logs an outgoing HTTP response.
func LogWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		init := time.Now()
		wrap := &logWriter{w, http.StatusOK, 0}
		next.ServeHTTP(wrap, r)
		secs := time.Since(init).Seconds()
		log.Printf(
			"%s %s %s :: %d %d %1.5f",
			r.RemoteAddr, r.Method, r.URL.Path, wrap.Code, wrap.Size, secs,
		)
	})
}

// RateWare is a middleware that rate limits users.
func RateWare(next http.Handler, size int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := neat.Addr(r)

		RateMutex.Lock()
		lmtr, ok := RateAddrs[addr]
		if !ok {
			lmtr = rate.NewLimiter(rate.Limit(float64(size)/3600.0), size)
			RateAddrs[addr] = lmtr
		}

		RateMutex.Unlock()
		if !lmtr.Allow() {
			prot.WriteError(w, http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
