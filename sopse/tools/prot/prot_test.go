package prot

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	// setup
	b := bytes.NewBufferString("body...")
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// success - no body
	body := Read(w, r, 4)
	assert.Equal(t, "", body)

	// setup
	r = httptest.NewRequest("GET", "/", b)

	// success - with body
	body = Read(w, r, 4)
	assert.Equal(t, "body", body)
}

func TestWrite(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Write(w, http.StatusOK, "text")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text", w.Body.String())

	// confirm - headers
	assert.Equal(t, "no-store", w.Header().Get("Cache-Control"))
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "nosniff", w.Header().Get("X-Content-Type-Options"))
}

func TestWriteError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success - no text
	WriteError(w, http.StatusNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: not found", w.Body.String())

	// setup
	w = httptest.NewRecorder()

	// success - with text
	WriteError(w, http.StatusNotFound, "text")
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: text", w.Body.String())
}
