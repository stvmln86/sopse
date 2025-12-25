package test

import (
	"fmt"
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
