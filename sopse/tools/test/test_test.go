package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResponse(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprintf(w, "body")

	// success
	code, body := GetResponse(t, w)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "body", body)
}

func TestNewRequest(t *testing.T) {
	// success
	r := NewRequest("GET", "/", "body")
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "/", r.URL.Path)

	// confirm - body
	bytes, err := io.ReadAll(r.Body)
	assert.Equal(t, "body", string(bytes))
	assert.NoError(t, err)
}
