package sopse

import (
	"os"
	"time"
)

// project notes //

/*
- json in, jsend out
- data stored in plain json files
- file I/O is done atomically
- user token is hash of addr, salt and time rounded to hour
- Book is a directory, each User is a subdirectory, each Pair is a file
- each User has a hidden "__meta__.json" file parsed as a Meta
*/

// app types //

// top-level app object
type App struct {
	Book *Book
	Conf *Conf
}

// json conf container
type Conf struct {
	Addr     string      `json:"addr"`
	Dire     string      `json:"dire"`
	Extn     string      `json:"extn"`
	Mode     os.FileMode `json:"mode"`
	Salt     string      `json:"salt"`
	RateBody int         `json:"rate_body"`
	// etc
}

// data types //

// root directory containing all data
type Book struct {
	Dire string
	Extn string
	Mode os.FileMode
}

func (b *Book) Create(hash string) (*User, error)
func (b *Book) CreateOrGet(hash string) (*User, error)
func (b *Book) Get(hash string) (*User, error)

// subdirectory containing pairs for one user
type User struct {
	Dire string
	Extn string
	Hash string
}

func (u *User) Get(name string) (*Pair, error)
func (u *User) Set(name, body string) (*Pair, error)
func (u *User) Update() error

// single file containing key-value pair
type Pair struct {
	Orig string    `json:"-"`
	Body string    `json:"body"`
	Hash string    `json:"hash"`
	Init time.Time `json:"init"`
}

func (p *Pair) Delete() error
func (p *Pair) Rename(hash string) error

// special pair file containing user data
type Meta struct {
	Orig string    `json:"-"`
	Addr string    `json:"addr"`
	Init time.Time `json:"init"`
}

// example handlers //

/*
	func (a *App) GetIndex(w http.ResponseWriter, r *http.Request) {
		prot.Write(Index, a.Book.Len())
	}
*/

/*
	func (a *App) GetPair(w http.ResponseWriter, r *http.Request) {
		hash := neat.PathValue("hash", neat.Hash)
		user, err := a.Book.Get(hash)
		switch {
		case err != nil:
			prot.WriteCode(w, http.StatusInternalServerError)
			return
		case user == nil:
			prot.WriteError(w, http.StatusNotFound, "user not found")
			return
		}

		name := neat.PathValue("name", neat.Name)
		pair, ok := user.Get(name)
		if !ok {
			prot.WriteError(w, http.StatusNotFound, "pair not found")
			return
		}

		prot.Write(w, http.StatusOK, pair.Body)
	}
*/
