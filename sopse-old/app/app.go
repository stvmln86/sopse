// Package app implements the App type and methods.
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stvmln86/sopse/sopse/tools/bolt"
	"github.com/stvmln86/sopse/sopse/tools/conf"
	"github.com/stvmln86/sopse/sopse/tools/ware"
	"go.etcd.io/bbolt"
)

// App is a top-level server controller.
type App struct {
	Conf *conf.Conf
	DB   *bbolt.DB
}

// New returns a new App.
func New(conf *conf.Conf, db *bbolt.DB) *App {
	return &App{conf, db}
}

// NewParse returns a new App from parsed command-line arguments.
func NewParse(elems []string) (*App, error) {
	conf := conf.Parse(elems)
	db, err := bolt.Connect(conf.Dbse)
	if err != nil {
		return nil, fmt.Errorf("cannot open database %q - %w", conf.Dbse, err)
	}

	return New(conf, db), nil
}

// Close closes the App's database connection.
func (a *App) Close() error {
	if err := a.DB.Close(); err != nil {
		return fmt.Errorf("cannot close database %q - %w", a.Conf.Dbse, err)
	}

	return nil
}

// ServeMux returns the App's configured ServeMux.
func (a *App) ServeMux() *http.ServeMux {
	smux := http.NewServeMux()
	for path, hand := range map[string]http.HandlerFunc{
		"GET /":                   a.GetIndexOr404,
		"GET /api/{uuid}":         a.GetUser,
		"POST /api/new":           a.PostNewUser,
		"POST /api/{uuid}/{name}": a.PostPair,
	} {
		smux.Handle(path, ware.Apply(hand, a.Conf.UserRate))
	}

	return smux
}

// Start starts the App's server.
func (a *App) Start() error {
	serv := &http.Server{
		Addr:         a.Conf.Addr,
		Handler:      a.ServeMux(),
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return serv.ListenAndServe()
}
