package prot

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func TestRead(t *testing.T) {
	// setup
	r := test.Request("GET", "/", "body")

	// success - normal body
	body := Read(r)
	assert.Equal(t, "body", body)

	// setup
	r.Body = nil

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
	assert.Subset(t, w.Header(), headers)
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
