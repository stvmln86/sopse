package main

import (
	"log"
	"os"

	"github.com/stvmln86/sopse/sopse/app"
)

func try(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app, err := app.NewParse(os.Args[1:])
	try(err)

	defer try(app.Close())
	log.Printf("starting Sopse on %q", app.Conf.Addr)
	try(app.Start())
}
