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
	buff := bytes.NewBufferString("body")
	r := httptest.NewRequest("GET", "/", buff)

	// success - normal body
	body := Read(r)
	assert.Equal(t, "body", body)

	// setup
	r = httptest.NewRequest("GET", "/", nil)

	// success - nil body
	body = Read(r)
	assert.Empty(t, body)
}

func TestWrite(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Write(w, http.StatusOK, "%s", "body")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "body", w.Body.String())

	// confirm - headers
	for attr, data := range headers {
		assert.Equal(t, data, w.Header().Get(attr))
	}
}

func TestWriteCode(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteCode(w, http.StatusInternalServerError)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error 500: internal server error", w.Body.String())
}

func TestWriteError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteError(w, http.StatusBadRequest, "%s", "body")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "error 400: body", w.Body.String())
}
