////////////////////////////////////////////////////////////////////////////////////////
//       sopse.go · stephen's obsessive pair storage engine · by Stephen Malone       //
////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

////////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                          //
////////////////////////////////////////////////////////////////////////////////////////

// 1.1 · system globals
////////////////////////

// DB is the global database connection object.
var DB *sqlx.DB

// Uptime is the system start time.
var Uptime = time.Now()

// 1.2 · configuration flags
/////////////////////////////

// Global command-line flags.
var (
	FlagAddr = flag.String("addr", "127.0.0.1:8000", "host address")
	FlagLife = flag.Duration("life", 24*7*time.Hour, "pair expiry time")
	FlagPath = flag.String("path", "./sopse.db", "database path")
	FlagRate = flag.Int("rate", 1000, "max requests per hour")
	FlagSize = flag.Int("size", 4096, "max request body size")
	FlagUser = flag.Int("user", 256, "max pairs per user")
)

// 1.3 · sqlite constants
//////////////////////////

// Pragma is the default always-enabled database pragma.
const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = true;
`

// Schema is the default first-run database schema.
const Schema = `
	create table if not exists Users (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		uuid text    not null default (lower(hex(randomblob(8)))),
		addr text    not null,

		unique(uuid)
	);

	create table if not exists Pairs (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		user integer not null,
		name text    not null,
		body text    not null,

		foreign key (user) references Users(id) on delete cascade,
		unique(user, name)
	);

	create index if not exists UserUUIDs on Users(uuid);
	create index if not exists PairNames on Pairs(user, name);
`

// SQLite query constants.
const (
	insertUser = `
		insert into Users (addr) values (?) returning uuid
	`

	selectStats = `select
		(select count(*) from Users) as users,
		(select count(*) from Pairs) as pairs;
	`

	upsertPair = `
		insert into Pairs (user, name, body) values (?, ?, ?)
		on conflict (user, name) do update set body = excluded.body
	`
)

// 1.4 · index page template
/////////////////////////////

// Index is the static index page template.
const Index = `
 ▄██▀█ ▄███▄ ████▄ ▄██▀█ ▄█▀█▄
 ▀███▄ ██ ██ ██ ██ ▀███▄ ██▄█▀
█▄▄██▀▄▀███▀▄████▀█▄▄██▀▄▀█▄▄▄
             ██
             ▀

stephen's obsessive pair storage engine, version v0.0.0:
- system uptime: %s
- current stats: %d users, %d pairs
- github source: https://github.com/stvmln86/sopse
`

////////////////////////////////////////////////////////////////////////////////////////
//                        part two · data processing functions                        //
////////////////////////////////////////////////////////////////////////////////////////

// Addr returns the remote IP address from a Request.
func Addr(r *http.Request) string {
	addr, _, _ := net.SplitHostPort(r.RemoteAddr)
	return addr
}

// Body returns a Request's body as a whitespace-trimmed string.
func Body(w http.ResponseWriter, r *http.Request) string {
	if r.Body == nil {
		return ""
	}

	r.Body = http.MaxBytesReader(w, r.Body, int64(*FlagSize))
	bytes, _ := io.ReadAll(r.Body)
	return strings.TrimSpace(string(bytes))
}

// PathValue returns a lowercase Request path value.
func PathValue(r *http.Request, name string) string {
	data := r.PathValue(name)
	return strings.ToLower(data)
}

////////////////////////////////////////////////////////////////////////////////////////
//                        part three · http response functions                        //
////////////////////////////////////////////////////////////////////////////////////////

// Write writes a formatted text/plain string to a ResponseWriter.
func Write(w http.ResponseWriter, code int, body string, elems ...any) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintf(w, body, elems...)
}

// WriteCode writes a text/plain error code to a ResponseWriter and logs an error.
func WriteCode(w http.ResponseWriter, code int, err error) {
	log.Print(err)
	stat := http.StatusText(code)
	Write(w, code, "error %d: %s", code, strings.ToLower(stat))
}

// WriteError writes a formatted text/plain error string to a ResponseWriter.
func WriteError(w http.ResponseWriter, code int, body string, elems ...any) {
	body = fmt.Sprintf(body, elems...)
	Write(w, code, "error %d: %s", code, body)
}

////////////////////////////////////////////////////////////////////////////////////////
//                         part four · http handler functions                         //
////////////////////////////////////////////////////////////////////////////////////////

// 4.1 · get handlers
//////////////////////

// GetIndex returns the index page or a 404 error.
func GetIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		WriteError(w, http.StatusNotFound, "path %s not found", r.URL.Path)
		return
	}

	var stats struct {
		Users int `db:"users"`
		Pairs int `db:"pairs"`
	}

	if err := DB.Get(&stats, selectStats); err != nil {
		WriteCode(w, http.StatusInternalServerError, err)
		return
	}

	dura := time.Since(Uptime).Round(1 * time.Second).String()
	Write(w, http.StatusOK, Index, dura, stats.Users, stats.Pairs)
}

// 4.2 · post handlers
///////////////////////

// PostCreateUser creates and returns a new user UUID.
func PostCreateUser(w http.ResponseWriter, r *http.Request) {
	var uuid string
	if err := DB.Get(&uuid, insertUser, Addr(r)); err != nil {
		WriteCode(w, http.StatusInternalServerError, err)
		return
	}

	Write(w, http.StatusCreated, "%s", uuid)
}

// PostSetPair sets a new or existing user pair.
func PostSetPair(w http.ResponseWriter, r *http.Request) {
	body := Body(w, r)
	uuid := PathValue(r, "uuid")
	name := PathValue(r, "name")

	var user int64
	err := DB.Get(&user, `select id from Users where uuid=?`, uuid)

	switch {
	case err == sql.ErrNoRows:
		WriteError(w, http.StatusNotFound, "user %s not found", uuid)
		return
	case err != nil:
		WriteCode(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = DB.Exec(upsertPair, user, name, body); err != nil {
		WriteCode(w, http.StatusInternalServerError, err)
		return
	}

	Write(w, http.StatusOK, "ok")
}

////////////////////////////////////////////////////////////////////////////////////////
//                        part five · http middleware functions                       //
////////////////////////////////////////////////////////////////////////////////////////

// ApplyWare applies all middleware to a HandlerFunc.
func ApplyWare(next http.HandlerFunc) http.Handler {
	return LogWare(next)
}

// logWriter is a custom ResponseWriter for logging middleware.
type logWriter struct {
	http.ResponseWriter
	Code int
	Size int
}

// WriteHeader writes and records a status code.
func (w *logWriter) WriteHeader(code int) {
	w.Code = code
	w.ResponseWriter.WriteHeader(code)
}

// Write writes and records a byte slice.
func (w *logWriter) Write(bytes []byte) (int, error) {
	n, err := w.ResponseWriter.Write(bytes)
	w.Size += n
	return n, err
}

// LogWare is a middleware that logs an outgoing HTTP response.
func LogWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		init := time.Now()
		wrap := &logWriter{ResponseWriter: w, Code: http.StatusOK, Size: 0}
		next.ServeHTTP(wrap, r)
		secs := time.Since(init).Seconds()
		log.Printf(
			"%s %s %s :: %d %d %1.5f",
			r.RemoteAddr, r.Method, r.URL.Path, wrap.Code, wrap.Size, secs,
		)
	})
}

////////////////////////////////////////////////////////////////////////////////////////
//                              part six · main function                              //
////////////////////////////////////////////////////////////////////////////////////////

// try logs a non-nil error.
func try(err error) {
	if err != nil {
		log.Print(err)
	}
}

// main runs the main Sopse program.
func main() {
	// Parse command-line arguments.
	flag.Parse()

	// Initialise database.
	log.Printf("connecting database on %s...", *FlagPath)
	DB = sqlx.MustConnect("sqlite3", *FlagPath)
	DB.MustExec(Pragma + Schema)
	DB.SetMaxIdleConns(1)
	DB.SetMaxOpenConns(1)

	// Initialise multiplexer and handlers.
	smux := http.NewServeMux()
	smux.Handle("POST /new", ApplyWare(PostCreateUser))
	smux.Handle("POST /{uuid}/{name}", ApplyWare(PostSetPair))
	smux.Handle("GET /", ApplyWare(GetIndex))

	// Initialise server.
	serv := &http.Server{
		Addr:           *FlagAddr,
		Handler:        smux,
		MaxHeaderBytes: 8192,
		IdleTimeout:    10 * time.Second,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Start server.
	log.Printf("starting Sopse on %s...", *FlagAddr)
	try(serv.ListenAndServe())

	// Close database and server.
	try(DB.Close())
	try(serv.Close())
}
