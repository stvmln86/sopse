package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	// success
	r := Request("GET", "/", "body")
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "/", r.URL.Path)

	// confirm - body
	bytes, err := io.ReadAll(r.Body)
	assert.Equal(t, "body", string(bytes))
	assert.NoError(t, err)
}

func TestResponse(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprint(w, "body")

	// success
	code, body, err := Response(w)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "body", body)
	assert.NoError(t, err)
}
