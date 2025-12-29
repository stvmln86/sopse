// Package ware implements HTTP middleware functions.
package ware

import (
	"net/http"
)

// Apply applies all middleware to a HandlerFunc.
func Apply(next http.HandlerFunc, rate int) http.Handler {
	return LogWare(RateWare(http.HandlerFunc(next), rate))
}
