package app

import (
	"net/http"

	"github.com/stvmln86/sopse/sopse/items/user"
	"github.com/stvmln86/sopse/sopse/tools/prot"
)

// PostPair sets a new or existing Pair for an existing User.
func (a *App) PostPair(w http.ResponseWriter, r *http.Request) {
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

	name := r.PathValue("name")
	body := prot.Read(w, r, a.Conf.BodySize)
	pair, err := user.GetPair(name)
	if err != nil {
		prot.WriteError(w, http.StatusInternalServerError)
		return
	}

	if pair != nil {
		if err := pair.Update(body); err != nil {
			prot.WriteError(w, http.StatusInternalServerError)
			return
		} else {
			prot.Write(w, http.StatusCreated, "ok")
			return
		}
	}

	size, err := user.Size()
	switch {
	case err != nil:
		prot.WriteError(w, http.StatusInternalServerError)
		return
	case size >= a.Conf.UserSize:
		prot.WriteError(w, http.StatusTooManyRequests, "storage limit reached")
		return
	}

	if _, err := user.AddPair(name, body); err != nil {
		prot.WriteError(w, http.StatusInternalServerError)
		return
	}

	prot.Write(w, http.StatusCreated, "ok")
}
