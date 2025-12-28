////////////////////////////////////////////////////////////////////////////////////////
//             sopse_test.go · unit tests for sopse.go · by Stephen Malone            //
////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
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
	w := httptest.NewRecorder()
	b := bytes.NewBufferString("body\n...")
	r := httptest.NewRequest("GET", "/", b)
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
