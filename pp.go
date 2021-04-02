package main

import (
	"fmt"
	"net/http"
)

func PPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		directives := fmt.Sprintf(
			`"accelerometer=(),
			"ambient-light-sensor=(),
			"autoplay=(),
			"battery=(),
			"camera=(),
			"clipboard-read=(),
			"clipboard-write=(),
			"cross-origin-isolated=(),
			"display-capture=(),
			"document-domain=(),
			"encrypted-media=(),
			"execution-while-not-rendered=(),
			"execution-while-out-of-viewport=(),
			"fullscreen=(),
			"gamepad=(),
			"geolocation=(self 'https://AnyTrustedSite.com'),
			"gyroscope=(),
			"magnetometer=(),
			"microphone=(),
			"midi=(),
			"navigation-override=(),
			"payment=(),
			"picture-in-picture=(),
			"publickey-credentials-get=(),
			"screen-wake-lock=(),
			"speaker=(),
			"speaker-selection=(),
			"sync-xhr=(),
			"usb=(),
			"web-share=(),
			"xr-spatial-tracking=()"`)

		w.Header().Set("Permissions-Policy", directives)
		next.ServeHTTP(w, r)
	})
}
