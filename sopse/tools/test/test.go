// Package test implements unit testing data and functions.
package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

// Request returns a new mock Request with a body.
func Request(mthd, path, body string) *http.Request {
	buff := bytes.NewBufferString(body)
	return httptest.NewRequest(mthd, path, buff)
}
