// Package test implements unit testing data and functions.
package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
)

// Request returns a new mock Request with a body.
func Request(mthd, path, body string) *http.Request {
	buff := bytes.NewBufferString(body)
	return httptest.NewRequest(mthd, path, buff)
}

// Response returns the status code and body from a ResponseRecorder.
func Response(w *httptest.ResponseRecorder) (int, string, error) {
	rslt := w.Result()
	bytes, err := io.ReadAll(rslt.Body)
	return rslt.StatusCode, string(bytes), err
}
