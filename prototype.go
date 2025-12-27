package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

///////////////////////////////// constants and globals ////////////////////////////////

var DB *sqlx.DB

const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = true;
`

const Schema = `
	create table if not exists Users (
		id   integer primary key asc,
		hash text    not null default (lower(randomblob(16))),
		init integer not null default (unixepoch()),
		addr text    not null,

		unique(hash)
	);

	create table if not exists Pairs (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		user integer not null,
		name text    not null,
		body text    not null,

		unique(user, name)
	);
`

////////////////////////////////// protocol functions //////////////////////////////////

func writeJSON(w http.ResponseWriter, code int, jmap map[string]any) {
	bytes, err := json.Marshal(jmap)
	if err != nil {
		code = http.StatusInternalServerError
		bytes = []byte(`{"status": "error", "message": "internal server error"`)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(bytes)
}

func Error(w http.ResponseWriter, code int) {
	stat := http.StatusText(code)
	writeJSON(w, code, map[string]any{
		"status":  "error",
		"message": strings.ToLower(stat),
	})
}

func Failure(w http.ResponseWriter, code int, jmap map[string]string) {
	writeJSON(w, code, map[string]any{
		"status": "fail",
		"data":   jmap,
	})
}

func Success(w http.ResponseWriter, code int, jmap map[string]any) {
	writeJSON(w, code, map[string]any{
		"status": "success",
		"data":   jmap,
	})
}

func Unmarshal(r *http.Request) (map[string]any, error) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var jmap map[string]any
	if err := json.Unmarshal(bytes, &jmap); err != nil {
		return nil, err
	}

	return jmap, nil
}

/////////////////////////////////// handler functions //////////////////////////////////

func GetIndex(w http.ResponseWriter, r *http.Request) {
	var size int
	if err := DB.Get(&size, "select count(*) from Pairs"); err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}

	Success(w, http.StatusOK, map[string]any{"pairs": size})
}

///////////////////////////////// middleware functions /////////////////////////////////

func LogWare(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

//////////////////////////////////// main functions ////////////////////////////////////

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// init db
	DB = sqlx.MustConnect("sqlite3", "sopse.db")
	DB.MustExec(Pragma + Schema)

	// init muxer
	mux := http.NewServeMux()
	mux.Handle("GET /", LogWare(GetIndex))

	// init server
	srv := &http.Server{Addr: ":8000", Handler: mux}

	// run server
	log.Printf("running sopse on :8000...")
	try(srv.ListenAndServe())
}
