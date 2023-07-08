package http

import "net/http"

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "appliaction/json; charset=UTF-8")

		next.ServeHTTP(w, r)
	})
}