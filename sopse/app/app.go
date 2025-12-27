// Package app implements the App type, handlers and methods.
package app

import (
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/sopse/sopse/tools/dbse"
	"github.com/stvmln86/sopse/sopse/tools/flag"
	"github.com/stvmln86/sopse/sopse/tools/ware"
)

// App is a top-level server controller.
type App struct {
	DB *sqlx.DB
}

// New returns a new App.
func New(db *sqlx.DB) *App {
	return &App{DB: db}
}

// NewConnect returns a new App with a connected database.
func NewConnect(path string) (*App, error) {
	db, err := dbse.Connect(path)
	if err != nil {
		return nil, err
	}

	return New(db), nil
}

// Server returns a Server object from the App.
func (a *App) Server() *http.Server {
	smux := http.NewServeMux()
	smux.Handle("GET /", ware.Apply(a.GetIndex))

	return &http.Server{
		Addr:         *flag.Addr,
		Handler:      smux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
