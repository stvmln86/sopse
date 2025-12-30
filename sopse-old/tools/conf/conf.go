// Package conf implements the Conf type and methods.
package conf

import (
	"flag"
	"time"
)

// Conf is a container for command-line configuration data.
type Conf struct {
	Addr     string
	Dbse     string
	BodySize int64
	PairLife time.Duration
	UserRate int
	UserSize int
}

// Parse returns a new Conf from parsed command-line arguments.
func Parse(elems []string) *Conf {
	conf := new(Conf)
	fset := flag.NewFlagSet("sopse", flag.ExitOnError)
	fset.StringVar(&conf.Addr, "addr", "127.0.0.1:8000", "host address")
	fset.StringVar(&conf.Dbse, "dbse", "./sopse.db", "database path")
	fset.Int64Var(&conf.BodySize, "bodySize", 4096, "max request body size")
	fset.DurationVar(&conf.PairLife, "pairLife", 24*7*time.Hour, "pair expiry time")
	fset.IntVar(&conf.UserRate, "userRate", 1000, "max requests per hour")
	fset.IntVar(&conf.UserSize, "userSize", 256, "max pairs per user")
	fset.Parse(elems)
	return conf
}
