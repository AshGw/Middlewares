package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

func CSPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 32)
		_, err := rand.Read(randomBytes)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		nonce := base64.StdEncoding.EncodeToString(randomBytes)

		// this is arbitrary give or take depending on the service
		directives := fmt.Sprintf(`
			default-src 'self';
			script-src 'nonce-%s' 'strict-dynamic';
			style-src  'nonce-%s';
			img-src 'self' blob: data:;
			font-src 'self';
			object-src 'none';
			base-uri 'self';
			form-action 'self';
			frame-ancestors 'none';
			block-all-mixed-content;
			upgrade-insecure-requests;
		`, nonce, nonce)

		w.Header().Set("Content-Security-Policy", directives)
		w.Header().Set("X-Nonce", nonce)
		next.ServeHTTP(w, r)
	})
}
