////////////////////////////////////////////////////////////////////////////////////////
//         sopse · stephen's obsessive pair storage engine · by stephen malone        //
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
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/time/rate"
)

////////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                          //
////////////////////////////////////////////////////////////////////////////////////////

// 1.1 · configuration flags
/////////////////////////////

// Command-line configuration flags.
var (
	FlagAddr     = flag.String("addr", "127.0.0.1:8000", "host address")
	FlagDbse     = flag.String("dbse", "./sopse.db", "database path")
	FlagBodySize = flag.Int64("bodySize", 4096, "max request body size")
	FlagPairLife = flag.Duration("pairLife", 24*7*time.Hour, "pair expiry time")
	FlagTaskWait = flag.Duration("taskWait", 6*time.Hour, "background task wait")
	FlagUserRate = flag.Int("userRate", 1000, "max requests per hour")
	FlagUserSize = flag.Int64("userSize", 256, "max pairs per user")
)

// 1.2 · database schema
/////////////////////////

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
		uuid text    not null default (lower(hex(randomblob(16)))),
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

// 1.3 · static index page
///////////////////////////

// Index is the static index page template.
const Index = `
 ▄██▀█ ▄███▄ ████▄ ▄██▀█ ▄█▀█▄
 ▀███▄ ██ ██ ██ ██ ▀███▄ ██▄█▀
█▄▄██▀▄▀███▀▄████▀█▄▄██▀▄▀█▄▄▄
             ██
             ▀

stephen's obsessive pair storage engine, version v0.0.0:
- system uptime: %s
- github source: https://github.com/stvmln86/sopse
`

// 1.4 · global state
//////////////////////

// Uptime tracks when the server started.
var Uptime = time.Now()

// Rates is a map of active rate limiters by IP address.
var Rates = make(map[string]*rate.Limiter)

// RatesMutex protects concurrent access to RateLimiters.
var RatesMutex = new(sync.Mutex)

// GlobalDB holds the database connection.
var GlobalDB *sqlx.DB

////////////////////////////////////////////////////////////////////////////////////////
//                             part two · database functions                          //
////////////////////////////////////////////////////////////////////////////////////////

// 2.1 · connection and setup
//////////////////////////////

// Connect opens a connection to the SQLite database and initializes the schema.
func Connect(path string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(Pragma + Schema); err != nil {
		return nil, err
	}

	return db, nil
}

// 2.2 · user database functions
//////////////////////////////////

// User represents a single user record.
type User struct {
	ID   int64  `db:"id"`
	Init int64  `db:"init"`
	UUID string `db:"uuid"`
	Addr string `db:"addr"`
}

const (
	queryCreateUser = `
		insert into Users (addr) values (?) returning *
	`
	queryGetUser = `
		select * from Users where uuid=? limit 1
	`
	queryDeleteUser = `
		delete from Users where id=?
	`
	queryListPairNames = `
		select name from Pairs join Users on Pairs.user = Users.id
		where Users.id=? order by name asc
	`
	queryCountUserPairs = `
		select count(*) from Pairs join Users on Pairs.user = Users.id
		where Users.id=? order by name asc
	`
)

// CreateUser creates and returns a new user in the database.
func CreateUser(db *sqlx.DB, addr string) (*User, error) {
	user := &User{}
	if err := db.Get(user, queryCreateUser, addr); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser retrieves an existing user by UUID, or returns nil if not found.
func GetUser(db *sqlx.DB, uuid string) (*User, error) {
	user := &User{}
	err := db.Get(user, queryGetUser, uuid)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return user, nil
	}
}

// DeleteUser removes a user from the database.
func DeleteUser(db *sqlx.DB, userID int64) error {
	_, err := db.Exec(queryDeleteUser, userID)
	return err
}

// ListUserPairNames returns all pair names owned by a user.
func ListUserPairNames(db *sqlx.DB, userID int64) ([]string, error) {
	var names []string
	err := db.Select(&names, queryListPairNames, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return names, nil
}

// CountUserPairs returns the number of pairs owned by a user.
func CountUserPairs(db *sqlx.DB, userID int64) (int64, error) {
	var count int64
	if err := db.Get(&count, queryCountUserPairs, userID); err != nil {
		return 0, err
	}
	return count, nil
}

// 2.3 · pair database functions
//////////////////////////////////

// Pair represents a single key-value pair record.
type Pair struct {
	ID   int64  `db:"id"`
	Init int64  `db:"init"`
	User int64  `db:"user"`
	Name string `db:"name"`
	Body string `db:"body"`
}

const (
	queryCreatePair = `
		insert into Pairs (user, name, body) values (?, ?, ?) returning *
	`
	queryGetPair = `
		select * from Pairs where user=? and name=? limit 1
	`
	queryUpdatePair = `
		update Pairs set body=? where id=?
	`
	queryDeletePair = `
		delete from Pairs where id=?
	`
)

// CreatePair creates and returns a new pair in the database.
func CreatePair(db *sqlx.DB, userID int64, name, body string) (*Pair, error) {
	pair := &Pair{}
	if err := db.Get(pair, queryCreatePair, userID, name, body); err != nil {
		return nil, err
	}
	return pair, nil
}

// GetPair retrieves an existing pair by user ID and name, or returns nil if not found.
func GetPair(db *sqlx.DB, userID int64, name string) (*Pair, error) {
	pair := &Pair{}
	err := db.Get(pair, queryGetPair, userID, name)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return pair, nil
	}
}

// UpdatePair updates the body of an existing pair.
func UpdatePair(db *sqlx.DB, pairID int64, body string) error {
	_, err := db.Exec(queryUpdatePair, body, pairID)
	return err
}

// DeletePair removes a pair from the database.
func DeletePair(db *sqlx.DB, pairID int64) error {
	_, err := db.Exec(queryDeletePair, pairID)
	return err
}

////////////////////////////////////////////////////////////////////////////////////////
//                         part three · protocol functions                            //
////////////////////////////////////////////////////////////////////////////////////////

// 3.1 · request reading
/////////////////////////

// ReadRequestBody reads and returns a request's body as a trimmed string.
func ReadRequestBody(w http.ResponseWriter, r *http.Request, maxSize int64) string {
	if r.Body == nil {
		return ""
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxSize)
	bytes, _ := io.ReadAll(r.Body)
	return strings.TrimSpace(string(bytes))
}

// 3.2 · response writing
//////////////////////////

// WriteResponse writes a text/plain string to the response writer.
func WriteResponse(w http.ResponseWriter, code int, text string) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write([]byte(text))
}

// WriteError writes a text/plain error message to the response writer.
func WriteError(w http.ResponseWriter, code int, texts ...string) {
	if len(texts) == 0 {
		stat := http.StatusText(code)
		texts = append(texts, strings.ToLower(stat))
	}

	text := fmt.Sprintf("error %d: %s", code, texts[0])
	WriteResponse(w, code, text)
}

////////////////////////////////////////////////////////////////////////////////////////
//                           part four · utility functions                            //
////////////////////////////////////////////////////////////////////////////////////////

// 4.1 · address extraction
////////////////////////////

// ExtractIPAddress returns the remote IP address from a request.
func ExtractIPAddress(r *http.Request) string {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

// 4.2 · time conversion
/////////////////////////

// UnixToLocalTime converts a Unix UTC timestamp to a local time object.
func UnixToLocalTime(unix int64) time.Time {
	return time.Unix(unix, 0).Local()
}

////////////////////////////////////////////////////////////////////////////////////////
//                          part five · middleware functions                          //
////////////////////////////////////////////////////////////////////////////////////////

// 5.1 · logging middleware
////////////////////////////

// logResponseWriter wraps http.ResponseWriter to capture status code and size.
type logResponseWriter struct {
	http.ResponseWriter
	Code int
	Size int
}

// WriteHeader captures the status code and writes it to the underlying writer.
func (w *logResponseWriter) WriteHeader(code int) {
	w.Code = code
	w.ResponseWriter.WriteHeader(code)
}

// Write captures the response size and writes to the underlying writer.
func (w *logResponseWriter) Write(bytes []byte) (int, error) {
	n, err := w.ResponseWriter.Write(bytes)
	w.Size += n
	return n, err
}

// LoggingMiddleware logs HTTP requests and responses.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &logResponseWriter{w, http.StatusOK, 0}
		next.ServeHTTP(wrapper, r)
		duration := time.Since(start).Seconds()
		log.Printf(
			"%s %s %s :: %d %d %1.5f",
			r.RemoteAddr, r.Method, r.URL.Path, wrapper.Code, wrapper.Size, duration,
		)
	})
}

// 5.2 · rate limiting middleware
//////////////////////////////////

// RateLimitMiddleware limits requests per IP address.
func RateLimitMiddleware(next http.Handler, maxRate int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := ExtractIPAddress(r)

		RatesMutex.Lock()
		limiter, exists := Rates[addr]
		if !exists {
			limiter = rate.NewLimiter(rate.Limit(float64(maxRate)/3600.0), maxRate)
			Rates[addr] = limiter
		}
		RatesMutex.Unlock()

		if !limiter.Allow() {
			WriteError(w, http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// 5.3 · middleware application
////////////////////////////////

// ApplyMiddleware applies all middleware to a handler function.
func ApplyMiddleware(handler http.HandlerFunc, maxRate int) http.Handler {
	return LoggingMiddleware(RateLimitMiddleware(handler, maxRate))
}

////////////////////////////////////////////////////////////////////////////////////////
//                            part six · http handlers                                //
////////////////////////////////////////////////////////////////////////////////////////

// 6.1 · index handler
///////////////////////

// HandleGetIndex serves the index page or returns 404 for unknown paths.
func HandleGetIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		WriteError(w, http.StatusNotFound, "url not found")
		return
	}

	uptime := time.Since(Uptime).Round(1 * time.Second).String()
	text := fmt.Sprintf(Index, uptime)
	WriteResponse(w, http.StatusOK, text)
}

// 6.2 · get user handler
//////////////////////////

// HandleGetUser returns all pair names for an existing user.
func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("uuid")
	user, err := GetUser(GlobalDB, uuid)

	switch {
	case user == nil:
		WriteError(w, http.StatusNotFound, "user not found")
		return
	case err != nil:
		WriteError(w, http.StatusInternalServerError)
		return
	}

	names, err := ListUserPairNames(GlobalDB, user.ID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError)
		return
	}

	WriteResponse(w, http.StatusOK, strings.Join(names, "\n"))
}

// 6.3 · create user handler
/////////////////////////////

// HandlePostNewUser creates and returns a new user UUID.
func HandlePostNewUser(w http.ResponseWriter, r *http.Request) {
	addr := ExtractIPAddress(r)
	user, err := CreateUser(GlobalDB, addr)
	if err != nil {
		log.Print(err)
		WriteError(w, http.StatusInternalServerError)
		return
	}

	WriteResponse(w, http.StatusCreated, user.UUID)
}

// 6.4 · set pair handler
//////////////////////////

// HandlePostPair creates or updates a pair for an existing user.
func HandlePostPair(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("uuid")
	user, err := GetUser(GlobalDB, uuid)

	switch {
	case user == nil:
		WriteError(w, http.StatusNotFound, "user not found")
		return
	case err != nil:
		WriteError(w, http.StatusInternalServerError)
		return
	}

	name := r.PathValue("name")
	body := ReadRequestBody(w, r, *FlagBodySize)

	// Check if pair already exists
	pair, err := GetPair(GlobalDB, user.ID, name)
	if err != nil {
		WriteError(w, http.StatusInternalServerError)
		return
	}

	// Update existing pair
	if pair != nil {
		if err := UpdatePair(GlobalDB, pair.ID, body); err != nil {
			WriteError(w, http.StatusInternalServerError)
			return
		}
		WriteResponse(w, http.StatusCreated, "ok")
		return
	}

	// Check user storage limit before creating new pair
	count, err := CountUserPairs(GlobalDB, user.ID)
	switch {
	case err != nil:
		WriteError(w, http.StatusInternalServerError)
		return
	case count >= *FlagUserSize:
		WriteError(w, http.StatusTooManyRequests, "storage limit reached")
		return
	}

	// Create new pair
	if _, err := CreatePair(GlobalDB, user.ID, name, body); err != nil {
		WriteError(w, http.StatusInternalServerError)
		return
	}

	WriteResponse(w, http.StatusCreated, "ok")
}

////////////////////////////////////////////////////////////////////////////////////////
//                         part seven · server and main                               //
////////////////////////////////////////////////////////////////////////////////////////

// 7.1 · server configuration
//////////////////////////////

// ConfigureRoutes sets up all HTTP routes with middleware.
func ConfigureRoutes(mux *http.ServeMux, maxRate int) {
	routes := map[string]http.HandlerFunc{
		"GET /":                   HandleGetIndex,
		"GET /api/{uuid}":         HandleGetUser,
		"POST /api/new":           HandlePostNewUser,
		"POST /api/{uuid}/{name}": HandlePostPair,
	}

	for path, handler := range routes {
		mux.Handle(path, ApplyMiddleware(handler, maxRate))
	}
}

// StartServer initializes and starts the HTTP server.
func StartServer(addr string, handler http.Handler) error {
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("starting Sopse on %q", addr)
	return server.ListenAndServe()
}

// 7.2 · main function
///////////////////////

// try exits the program if an error occurred.
func try(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connect to database
	db, err := Connect(*FlagDbse)
	try(err)
	defer try(db.Close())
	GlobalDB = db

	// Configure routes and start server
	mux := http.NewServeMux()
	ConfigureRoutes(mux, *FlagUserRate)
	try(StartServer(*FlagAddr, mux))
}
