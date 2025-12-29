package ware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestRateWare(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/", nil)
	RateAddrs = make(map[string]*rate.Limiter)

	// success
	hand := RateWare(http.HandlerFunc(mockHandler), 5)
	assert.NotNil(t, hand)

	// confirm - under limit
	for range 5 {
		w := httptest.NewRecorder()
		hand.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "body", w.Body.String())
	}

	// failure - over limit
	w := httptest.NewRecorder()
	hand.ServeHTTP(w, r)
	assert.Equal(t, http.StatusTooManyRequests, w.Code)
	assert.Equal(t, "error 429: too many requests", w.Body.String())

	// confirm - limiter
	assert.NotNil(t, RateAddrs["192.0.2.1"])
}
