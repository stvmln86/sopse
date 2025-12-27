// Package sqls implements SQLite pragma and schema constants.
package sqls

// Pragma is the default always-enabled database pragma.
const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = true;
`

// Schema is the default first-run database schema.
const Schema = `
	create table if not exists Addrs (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		addr text    not null,
		agnt text    not null,

		unique(addr)
	);

	create table if not exists Users (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		uuid text    not null default (lower(hex(randomblob(8)))),
		addr integer not null,

		foreign key (addr) references Addrs(id),
		unique(uuid)
	);

	create table if not exists Pairs (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		user integer not null,
		name text    not null,
		body text    not null,

		foreign key (user) references Users(id),
		unique(user, name)
	);

	create index AddrAddrs on Addrs(addr);
	create index UserUUIDs on Users(uuid);
	create index PairNames on Pairs(user, name);
`
