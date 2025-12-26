///////////////////////////////////////////////////////////////////////////////////////
//        sopse · stephen's obsessive pair storage engine · by Stephen Malone        //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

///////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                         //
///////////////////////////////////////////////////////////////////////////////////////

// 1.1 · system globals
////////////////////////

// DB is the global database connection.
var DB *sqlx.DB

// 1.2 · command-line flags
////////////////////////////

// FlagSet is the default command-line parser.
var FlagSet = flag.NewFlagSet("sopse", flag.ExitOnError)

// Defined command-line flags.
var (
	FlagAddr = FlagSet.String("addr", "127.0.0.1:8000", "host address")
	FlagPath = FlagSet.String("path", "./sopse.db", "database path")
)

// 1.3 · database pragma & schema
//////////////////////////////////

// Pragma is the default always-enabled database pragma.
const Pragma = `
	pragma encoding = utf8;
	pragma foreign_keys = true;
`

// Schema is the default first-run database schema.
const Schema = `
	create table if not exists Users (
		id   integer primary key asc,
		uuid text    not null default (lower(hex(randomblob(16)))),
		init integer not null default (unixepoch()),
		addr text    not null,

		unique(uuid)
	);

	create table if not exists Pairs (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		user integer not null references Users(id),
		name text    not null,
		body text    not null,

		unique(user, name)
	);
`

///////////////////////////////////////////////////////////////////////////////////////
//                       part two · data sanitisation functions                      //
///////////////////////////////////////////////////////////////////////////////////////

// Addr returns the remote IP address from a Request.
func Addr(r *http.Request) string {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

// Trim returns a string trimmed to a maximum size.
func Trim(text string, size int) string {
	if len(text) > size {
		return text[:size]
	}

	return text
}

///////////////////////////////////////////////////////////////////////////////////////
//                        part three · http protocol functions                       //
///////////////////////////////////////////////////////////////////////////////////////

// 2.1 · request functions
///////////////////////////

// Read returns the body string from a Request.
func Read(r *http.Request) (string, error) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// 2.2 · response functions
////////////////////////////

// Write writes a formatted text/plain string to a ResponseWriter.
func Write(w http.ResponseWriter, code int, text string, elems ...any) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, text, elems...)
}

// WriteCode writes a text/plain error code to a ResponseWriter.
func WriteCode(w http.ResponseWriter, code int) {
	text := http.StatusText(code)
	text = strings.ToLower(text)
	Write(w, code, "error %d: %s", code, text)
}

// WriteError writes a formatted text/plain error string to a Response Writer.
func WriteError(w http.ResponseWriter, code int, text string, elems ...any) {
	text = fmt.Sprintf(text, elems...)
	Write(w, code, "error %d: %s", code, text)
}

///////////////////////////////////////////////////////////////////////////////////////
//                         part four · http handler functions                        //
///////////////////////////////////////////////////////////////////////////////////////

// 4.1 · GET handlers
//////////////////////

// GetIndex returns the index page.
func GetIndex(w http.ResponseWriter, r *http.Request) {
	Write(w, http.StatusOK, "hello")
}

// 4.2 · POST handlers
///////////////////////

// PostRegister creates and returns a new user token.
func PostRegister(w http.ResponseWriter, r *http.Request) {
	code := "insert into Users (addr) values (?)"
	if _, err := DB.Exec(code, Addr(r)); err != nil {
		WriteCode(w, http.StatusInternalServerError)
		return
	}

	var uuid string
	code = "select uuid from Users where addr=? order by id desc limit 1"
	if err := DB.Get(&uuid, code, Addr(r)); err != nil {
		WriteCode(w, http.StatusInternalServerError)
		return
	}

	Write(w, http.StatusCreated, "%s", uuid)
}

///////////////////////////////////////////////////////////////////////////////////////
//                          part five · middleware functions                         //
///////////////////////////////////////////////////////////////////////////////////////

// 5.1 · logging middleware
////////////////////////////

// LogWare logs HTTP requests with method, path, and duration.
func LogWare(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
	})
}

///////////////////////////////////////////////////////////////////////////////////////
//  testing testing testing testing testing testing testing testing testing testing  //
///////////////////////////////////////////////////////////////////////////////////////

// main runs the main Sopse program.
func main() {
	// Parse command-line flags.
	FlagSet.Parse(os.Args[1:])

	// Initialise database.
	DB = sqlx.MustConnect("sqlite3", *FlagPath)
	DB.MustExec(Pragma + Schema)

	// Initialise multiplexer.
	mux := http.NewServeMux()
	mux.Handle("GET /", LogWare(GetIndex))
	mux.Handle("POST /new", LogWare(PostRegister))

	// Initialise server.
	srv := &http.Server{
		Addr:              *FlagAddr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	// Run server.
	log.Printf("starting Sopse on %s", *FlagAddr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// Close server.
	if err := srv.Close(); err != nil {
		log.Fatal(err)
	}
}
