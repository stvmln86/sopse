package app

import (
	"net/http"
	"strings"

	"github.com/stvmln86/sopse/sopse/items/user"
	"github.com/stvmln86/sopse/sopse/tools/prot"
)

// GetUser returns the Pair names for an existing User.
func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("uuid")
	user, err := user.Get(a.DB, uuid)
	switch {
	case user == nil:
		prot.WriteError(w, http.StatusNotFound, "user not found")
		return
	case err != nil:
		prot.WriteError(w, http.StatusInternalServerError)
		return
	}

	pairs, err := user.ListPairs()
	if err != nil {
		prot.WriteError(w, http.StatusInternalServerError)
		return
	}

	var names = make([]string, 0, len(pairs))
	for _, pair := range pairs {
		names = append(names, pair.Name())
	}

	prot.Write(w, http.StatusOK, strings.Join(names, "\n"))
}
