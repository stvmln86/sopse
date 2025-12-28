////////////////////////////////////////////////////////////////////////////////////////
//             sopse_test.go · unit tests for sopse.go · by Stephen Malone            //
////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

////////////////////////////////////////////////////////////////////////////////////////
//                          part zero · test data and helpers                         //
////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                          //
////////////////////////////////////////////////////////////////////////////////////////

// 1.3 · sqlite constants
//////////////////////////

func TestPragma(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Pragma)
	assert.NoError(t, err)
}

func TestSchema(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Schema)
	assert.NoError(t, err)
}

////////////////////////////////////////////////////////////////////////////////////////
//                        part two · data processing functions                        //
////////////////////////////////////////////////////////////////////////////////////////

func TestAddr(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/", nil)

	// success
	addr := Addr(r)
	assert.Equal(t, "192.0.2.1", addr)
}

func TestBody(t *testing.T) {
	// setup
	b := bytes.NewBufferString("body\n...")
	r := httptest.NewRequest("GET", "/", b)
	w := httptest.NewRecorder()
	*FlagSize = 5

	// success
	body := Body(w, r)
	assert.Equal(t, "body", body)
}

func TestPathValue(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/", nil)
	r.SetPathValue("name", "DATA")

	// success
	data := PathValue(r, "name")
	assert.Equal(t, "data", data)
}

////////////////////////////////////////////////////////////////////////////////////////
//                        part three · http response functions                        //
////////////////////////////////////////////////////////////////////////////////////////

func TestWrite(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Write(w, http.StatusOK, "%s", "body")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "body", w.Body.String())

	// confirm - headers
	assert.Equal(t, "no-store", w.Header().Get("Cache-Control"))
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "nosniff", w.Header().Get("X-Content-Type-Options"))
}

func TestWriteCode(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteCode(w, http.StatusInternalServerError)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error 500: internal server error", w.Body.String())
}

func TestWriteError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteError(w, http.StatusBadRequest, "%s", "body")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "error 400: body", w.Body.String())
}

////////////////////////////////////////////////////////////////////////////////////////
//                         part four · http handler functions                         //
////////////////////////////////////////////////////////////////////////////////////////

func TestGetIndex(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// success
	GetIndex(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, Index, w.Body.String())

	// setup
	r = httptest.NewRequest("GET", "/nope", nil)
	w = httptest.NewRecorder()

	// failure - 404 error
	GetIndex(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: path /nope not found", w.Body.String())
}
