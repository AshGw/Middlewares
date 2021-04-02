package main

import "net/http"

// in dev use this if u have valid certificates
func STSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains, preload")
		next.ServeHTTP(w, r)
	})
}
