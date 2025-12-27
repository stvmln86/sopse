package ware

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func mockHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "body")
}

func TestApply(t *testing.T) {
	// success
	hand := Apply(mockHandler)
	assert.NotNil(t, hand)
}

func TestLogWare(t *testing.T) {
	// setup
	r := test.Request("GET", "/", "")
	w := httptest.NewRecorder()
	buff := new(bytes.Buffer)
	hand := LogWare(http.HandlerFunc(mockHandler))
	log.SetFlags(0)
	log.SetOutput(buff)

	// success
	hand.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "body", w.Body.String())

	// confirm - logging
	assert.Regexp(t, `192.0.2.1:1234 GET / :: 200 4 0.\d{5}`, buff.String())
}
