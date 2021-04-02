package main

import "net/http"

func CCMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, must-revalidate, proxy-revalidate")
		next.ServeHTTP(w, r)
	})
}
