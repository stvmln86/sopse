////////////////////////////////////////////////////////////////////////////////////////
//         sopse · stephen's obsessive pair storage engine · by stephen malone        //
////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                          //
////////////////////////////////////////////////////////////////////////////////////////

// 1.1 · configuration flags
/////////////////////////////

// Command-line configuration flags.
var (
	FlagAddr     = flag.String("addr", ":8000", "host address")
	FlagDbse     = flag.String("dbse", "sopse.db", "database path")
	FlagBodySize = flag.Int64("body_size", 4096, "max request body size")
	FlagPairLife = flag.Duration("pair_life", 24*7*time.Hour, "pair expiry time")
	FlagTaskWait = flag.Duration("task_wait", 6*time.Hour, "background task wait")
	FlagUserRate = flag.Int("user_rate", 1000, "max requests per hour")
	FlagUserSize = flag.Int64("user_size", 256, "max pairs per user")
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
