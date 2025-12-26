package prot

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
	"github.com/stvmln86/sopse/sopse/tests/mock"
)

func TestRead(t *testing.T) {
	// setup
	r := mock.Request("GET", "/", "body")

	// success
	body, err := Read(r)
	assert.Equal(t, "body", body)
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Write(w, http.StatusOK, "text")
	asrt.Response(t, w, http.StatusOK, "text")

	// confirm - header
	data := w.Header().Get("Content-Type")
	assert.Equal(t, "text/plain; charset=utf-8", data)
}

func TestWriteCode(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteCode(w, http.StatusBadRequest)
	asrt.Response(t, w, http.StatusBadRequest, "error 400: bad request")
}

func TestWriteError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteError(w, http.StatusBadRequest, "text")
	asrt.Response(t, w, http.StatusBadRequest, "error 400: text")
}
