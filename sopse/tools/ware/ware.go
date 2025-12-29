// Package ware implements HTTP middleware functions.
package ware

import (
	"net/http"
)

// Apply applies all middleware to a HandlerFunc.
func Apply(next http.HandlerFunc) http.Handler {
	return LogWare(http.HandlerFunc(next))
}
