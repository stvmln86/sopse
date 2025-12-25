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
	r := test.NewRequest("GET", "/", "body")

	// success
	body, err := Read(r)
	assert.Equal(t, "body", body)
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Write(w, http.StatusOK, "%s", "body")
	code, body := test.GetResponse(t, w)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "body", body)

	// confirm - header
	data := w.Header().Get("Content-Type")
	assert.Equal(t, "text/plain; charset=utf-8", data)
}

func TestWriteCode(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteCode(w, http.StatusInternalServerError)
	code, body := test.GetResponse(t, w)
	assert.Equal(t, http.StatusInternalServerError, code)
	assert.Equal(t, "error 500: internal server error", body)
}

func TestWriteError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteError(w, http.StatusBadRequest, "%s", "body")
	code, body := test.GetResponse(t, w)
	assert.Equal(t, http.StatusBadRequest, code)
	assert.Equal(t, "error 400: body", body)
}
