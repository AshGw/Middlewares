package main

import "net/http"

func NoSrvMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Server")
		next.ServeHTTP(w, r)
	})
}
