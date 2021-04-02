package main

import "net/http"

func RPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Referrer-Policy", "no-referrer, strict-origin-when-cross-origin")
		next.ServeHTTP(w, r)
	})
}
