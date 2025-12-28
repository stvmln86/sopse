////////////////////////////////////////////////////////////////////////////////////////
//       sopse.go · stephen's obsessive pair storage engine · by Stephen Malone       //
////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
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

////////////////////////////////////////////////////////////////////////////////////////
//                                    project notes                                   //
////////////////////////////////////////////////////////////////////////////////////////

/*
# 2025-12-28
- don't limit path value sizes (uuid, name, etc)
- use -bodySize as Server.MaxHeaderBytes
*/
