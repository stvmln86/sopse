// Package flag implements comand-line parsing globals and functions.
package flag

import (
	"flag"
)

// flagSet is the default global command-line parser.
var flagSet = flag.NewFlagSet("sopse", flag.ExitOnError)

// Defined command-line arguments.
var (
	Addr    = flagSet.String("addr", "127.0.0.1:8000", "host address")
	Path    = flagSet.String("path", "./sopse.db", "database path")
	BodyMax = flagSet.Int("bodyMax", 4096, "max body size")
	NameMax = flagSet.Int("nameMax", 64, "max name size")
)

// Parse parses a command-line argument slice.
func Parse(elems []string) {
	flagSet.Parse(elems)
}
