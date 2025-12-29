package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stvmln86/sopse/sopse/tools/prot"
)

// Index is the static index page template.
const Index = `
 ▄██▀█ ▄███▄ ████▄ ▄██▀█ ▄█▀█▄
 ▀███▄ ██ ██ ██ ██ ▀███▄ ██▄█▀
█▄▄██▀▄▀███▀▄████▀█▄▄██▀▄▀█▄▄▄
             ██
             ▀

stephen's obsessive pair storage engine, version v0.0.0:
- system uptime: %s
- github source: https://github.com/stvmln86/sopse
`

// Uptime is the system start time.
var Uptime = time.Now()

// GetIndexOr404 returns the index page or a Not Found error.
func (a *App) GetIndexOr404(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		prot.WriteError(w, http.StatusNotFound)
		return
	}

	dura := time.Since(Uptime).Round(1 * time.Second).String()
	text := fmt.Sprintf(Index, dura)
	prot.Write(w, http.StatusOK, text)
}
