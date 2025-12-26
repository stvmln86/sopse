// Package mock implements unit testing mock data and functions.
package mock

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

// Request returns a new mock Request.
func Request(meth, path, body string) *http.Request {
	buff := bytes.NewBufferString(body)
	return httptest.NewRequest(meth, path, buff)
}
