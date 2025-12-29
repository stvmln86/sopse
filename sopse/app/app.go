// Package app implements the App type and methods.
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stvmln86/sopse/sopse/tools/bolt"
	"github.com/stvmln86/sopse/sopse/tools/conf"
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

// ServeMux returns a new ServeMux with the App's handlers.
func (a *App) ServeMux() *http.ServeMux {
	smux := http.NewServeMux()
	return smux
}

// Server returns a new Server with the App's handlers.
func (a *App) Server() *http.Server {
	return &http.Server{
		Addr:         a.Conf.Addr,
		Handler:      a.ServeMux(),
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}
