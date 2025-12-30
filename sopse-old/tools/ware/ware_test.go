package ware

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func mockHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "body")
}

func TestApply(t *testing.T) {
	// success
	hand := Apply(mockHandler, 1)
	assert.NotNil(t, hand)
}

func TestLogWare(t *testing.T) {
	// setup
	b := new(bytes.Buffer)
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	log.SetFlags(0)
	log.SetOutput(b)

	// success
	hand := LogWare(http.HandlerFunc(mockHandler))
	assert.NotNil(t, hand)

	// confirm - handler
	hand.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "body", w.Body.String())

	// confirm - logs
	assert.Regexp(t, `192\.0\.2\.1:1234 GET \/ :: 200 4 0\.\d{5}`, b.String())
}

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
