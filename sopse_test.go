////////////////////////////////////////////////////////////////////////////////////////
//             sopse_test.go · unit tests for sopse.go · by Stephen Malone            //
////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

////////////////////////////////////////////////////////////////////////////////////////
//                          part zero · test data and helpers                         //
////////////////////////////////////////////////////////////////////////////////////////

// mockData is mock database data for unit testing.
const mockData = `
	insert into Users (uuid, addr) values ('mockuser', '192.0.2.1');
	insert into Pairs (user, name, body) values (1, 'mockpair',  'body');
	insert into Pairs (user, name, body) values (1, 'mockpair2', 'body');
`

// mockDB initialises DB as an in-memory database populated with mockData.
func mockDB() {
	DB = sqlx.MustConnect("sqlite3", ":memory:")
	DB.MustExec(Pragma + Schema + mockData)
}

// mockHandler is a mock HandlerFunc for middleware testing.
func mockHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "body")
}

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
	b := new(bytes.Buffer)
	w := httptest.NewRecorder()
	log.SetFlags(0)
	log.SetOutput(b)

	// success
	WriteCode(w, http.StatusInternalServerError, errors.New("error"))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error 500: internal server error", w.Body.String())

	// confirm - logs
	assert.Equal(t, "error\n", b.String())
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

// 4.1 · get handlers
//////////////////////

func TestGetIndex(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	mockDB()

	// success
	GetIndex(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())

	// setup
	r = httptest.NewRequest("GET", "/nope", nil)
	w = httptest.NewRecorder()

	// failure - 404 error
	GetIndex(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: path /nope not found", w.Body.String())
}

func TestGetPair(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/mockuser/mockpair", nil)
	w := httptest.NewRecorder()
	r.SetPathValue("uuid", "mockuser")
	r.SetPathValue("name", "mockpair")
	mockDB()

	// success
	GetPair(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "body", w.Body.String())

	// setup
	r = httptest.NewRequest("GET", "/nope/nope", nil)
	w = httptest.NewRecorder()
	r.SetPathValue("uuid", "nope")
	r.SetPathValue("name", "nope")

	// failure - pair not found
	GetPair(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: pair nope/nope not found", w.Body.String())
}

func TestGetUser(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/mockuser", nil)
	w := httptest.NewRecorder()
	r.SetPathValue("uuid", "mockuser")
	mockDB()

	// success
	GetUser(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "mockpair\nmockpair2", w.Body.String())

	// setup
	r = httptest.NewRequest("GET", "/nope", nil)
	w = httptest.NewRecorder()
	r.SetPathValue("uuid", "nope")

	// failure - user not found
	GetUser(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: user nope not found", w.Body.String())
}

// 4.2 · post handlers
///////////////////////

func TestPostCreateUser(t *testing.T) {
	// setup
	r := httptest.NewRequest("POST", "/new", nil)
	w := httptest.NewRecorder()
	mockDB()

	// success
	PostCreateUser(w, r)
	body := w.Body.String()
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Regexp(t, `[0-9a-f]{16}`, body)

	// confirm - database
	var addr string
	err := DB.Get(&addr, "select addr from Users where uuid=?", body)
	assert.Equal(t, "192.0.2.1", addr)
	assert.NoError(t, err)
}

func TestPostSetPair(t *testing.T) {
	// setup
	b := bytes.NewBufferString("body")
	r := httptest.NewRequest("POST", "/mockuser/name", b)
	w := httptest.NewRecorder()
	r.SetPathValue("uuid", "mockuser")
	r.SetPathValue("name", "name")
	mockDB()

	// success
	PostSetPair(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())

	// confirm - database
	var body string
	err := DB.Get(&body, "select body from Pairs where name='name'")
	assert.Equal(t, "body", body)
	assert.NoError(t, err)

	// setup
	r = httptest.NewRequest("POST", "/nope/nope", nil)
	w = httptest.NewRecorder()
	r.SetPathValue("uuid", "nope")
	r.SetPathValue("name", "nope")

	// failure - user not found
	PostSetPair(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: user nope not found", w.Body.String())
}

////////////////////////////////////////////////////////////////////////////////////////
//                        part five · http middleware functions                       //
////////////////////////////////////////////////////////////////////////////////////////

func TestApplyWare(t *testing.T) {
	// success
	hand := ApplyWare(mockHandler)
	assert.NotNil(t, hand)
}

func TestLogWare(t *testing.T) {
	// setup
	b := new(bytes.Buffer)
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	log.SetFlags(0)
	log.SetOutput(b)

	// success
	LogWare(http.HandlerFunc(mockHandler)).ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "body", w.Body.String())

	// confirm - logs
	assert.Regexp(t, `192\.0\.2\.1:1234 GET \/ :: 200 4 0\.\d{5}`, b.String())
}

func TestRateWare(t *testing.T) {
	// setup - clear rate limits
	RateLimits.Addrs = make(map[string]*rate.Limiter)
	*FlagRate = 5

	// success - requests under limit
	for range 5 {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		RateWare(http.HandlerFunc(mockHandler)).ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "body", w.Body.String())
	}

	// failure - rate limit exceeded
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	RateWare(http.HandlerFunc(mockHandler)).ServeHTTP(w, r)
	assert.Equal(t, http.StatusTooManyRequests, w.Code)
	assert.Equal(t, "error 429: rate limit exceeded", w.Body.String())

	// confirm - limiter exists for IP
	assert.NotNil(t, RateLimits.Addrs["192.0.2.1"])
}
