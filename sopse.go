///////////////////////////////////////////////////////////////////////////////////////
//        sopse · stephen's obsessive pair storage engine · by Stephen Malone        //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
const Pragma = ``

// Schema is the default first-run database schema.
const Schema = ``

///////////////////////////////////////////////////////////////////////////////////////
//                         part two · http protocol functions                        //
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
//  testing testing testing testing testing testing testing testing testing testing  //
///////////////////////////////////////////////////////////////////////////////////////

// main runs the main Sopse program.
func main() {
	// Parse command-line flags.
	FlagSet.Parse(os.Args[1:])

	// Initialise database.
	DB = sqlx.MustConnect("sqlite3", *FlagPath)
	DB.MustExec(Pragma + Schema)
}
