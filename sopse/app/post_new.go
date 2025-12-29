package app

import (
	"log"
	"net/http"

	"github.com/stvmln86/sopse/sopse/items/user"
	"github.com/stvmln86/sopse/sopse/tools/prot"
)

// PostNewUser creates and returns a new User UUID.
func (a *App) PostNewUser(w http.ResponseWriter, r *http.Request) {
	user, err := user.Create(a.DB, r)
	if err != nil {
		log.Print(err)
		prot.WriteError(w, http.StatusInternalServerError)
		return
	}

	prot.Write(w, http.StatusCreated, user.UUID())
}
