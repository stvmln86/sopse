package ware

import (
	"net/http"
	"sync"

	"github.com/stvmln86/sopse/sopse/tools/neat"
	"github.com/stvmln86/sopse/sopse/tools/prot"
	"golang.org/x/time/rate"
)

// RateAddrs is a map of active rate limits.
var RateAddrs = make(map[string]*rate.Limiter)

// RateMutex is a mutex for RateAddrs.
var RateMutex = new(sync.Mutex)

// RateWare is a middleware that rate limits users.
func RateWare(next http.Handler, size int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := neat.Addr(r)

		RateMutex.Lock()
		lmtr, ok := RateAddrs[addr]
		if !ok {
			lmtr = rate.NewLimiter(rate.Limit(float64(size)/3600.0), size)
			RateAddrs[addr] = lmtr
		}

		RateMutex.Unlock()
		if !lmtr.Allow() {
			prot.WriteError(w, http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
