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
	code, body, err := test.Response(w)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "body", body)
	assert.NoError(t, err)

	// confirm - headers
	assert.Subset(t, w.Header(), headers)
}

func TestWriteCode(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteCode(w, http.StatusInternalServerError)
	code, body, err := test.Response(w)
	assert.Equal(t, http.StatusInternalServerError, code)
	assert.Equal(t, "error 500: internal server error", body)
	assert.NoError(t, err)
}

func TestWriteError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteError(w, http.StatusBadRequest, "%s", "body")
	code, body, err := test.Response(w)
	assert.Equal(t, http.StatusBadRequest, code)
	assert.Equal(t, "error 400: body", body)
	assert.NoError(t, err)
}
