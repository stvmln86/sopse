package app

import (
	"net/http"

	"github.com/stvmln86/sopse/sopse/tools/prot"
)

// Index is the static index page template.
const Index = `
 ▄██▀█ ▄███▄ ████▄ ▄██▀█ ▄█▀█▄
 ▀███▄ ██ ██ ██ ██ ▀███▄ ██▄█▀
█▄▄██▀▄▀███▀▄████▀█▄▄██▀▄▀█▄▄▄
             ██
             ▀
`

// GetIndex returns the static index page or a 404 error.
func (a *App) GetIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		prot.WriteError(w, http.StatusNotFound, "path %s not found", r.URL.Path)
		return
	}

	prot.Write(w, http.StatusOK, Index)
}
