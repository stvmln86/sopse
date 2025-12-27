// Package main implements the main Sopse program.
package main

import (
	"log"
	"os"

	"github.com/stvmln86/sopse/sopse/app"
	"github.com/stvmln86/sopse/sopse/tools/flag"
)

// try fatally logs a non-nil error.
func try(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// main runs the main Sopse program.
func main() {
	flag.Parse(os.Args[1:])
	app, err := app.NewConnect(*flag.Path)
	try(err)

	log.Printf("starting Sopse on %s...", *flag.Addr)
	try(app.Server().ListenAndServe())
	try(app.Close())
}
