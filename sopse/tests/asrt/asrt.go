// Package asrt implements unit testing assertion functions.
package asrt

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Response asserts the status code and body from a ResponseRecorder.
func Response(t *testing.T, w *httptest.ResponseRecorder, code int, body string) {
	rslt := w.Result()
	bytes, err := io.ReadAll(rslt.Body)
	assert.Equal(t, code, rslt.StatusCode)
	assert.Equal(t, body, string(bytes))
	assert.NoError(t, err)
}
