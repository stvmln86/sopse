// Package test implements unit testing data and functions.
package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// GetResponse returns the status code and body from a ResponseRecorder.
func GetResponse(t *testing.T, w *httptest.ResponseRecorder) (int, string) {
	rslt := w.Result()
	bytes, err := io.ReadAll(rslt.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rslt.StatusCode, string(bytes)
}

// NewRequest returns a new Request.
func NewRequest(meth, path, body string) *http.Request {
	buff := bytes.NewBufferString(body)
	return httptest.NewRequest(meth, path, buff)
}
