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

// GetIndex returns the static index page.
func (a *App) GetIndex(w http.ResponseWriter, r *http.Request) {
	prot.Write(w, http.StatusOK, Index)
}
