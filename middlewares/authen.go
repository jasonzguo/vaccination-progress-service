package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Authenticate(next httprouter.Handle) httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Check for Auth Header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No Auth Header", http.StatusUnauthorized)
			return
		}

		if next != nil {
			next(w, r, ps)
		}
	}
	return fn
}
