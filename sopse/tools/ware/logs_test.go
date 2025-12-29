package ware

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "body")
}

func TestLogWare(t *testing.T) {
	// setup
	b := new(bytes.Buffer)
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	log.SetFlags(0)
	log.SetOutput(b)

	// success
	LogWare(http.HandlerFunc(mockHandler)).ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "body", w.Body.String())

	// confirm - logs
	assert.Regexp(t, `192\.0\.2\.1:1234 GET \/ :: 200 4 0\.\d{5}`, b.String())
}
