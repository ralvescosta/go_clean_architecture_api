package coreinterfaces

import "net/http"

// HeadersMiddleware ...
func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(req http.ResponseWriter, res *http.Request) {
		req.Header().Add("Content-Type", "application/json; charset=utf-8")
		req.Header().Add("X-DNS-Prefetch-Control", "off")
		req.Header().Add("X-Frame-Options", "SAMEORIGIN")
		req.Header().Add("Strict-Transport-Security", "max-age=15552000; includeSubDomains")
		req.Header().Add("X-Download-Options", "noopen")
		req.Header().Add("X-Content-Type-Options", "nosniff")
		req.Header().Add("X-XSS-Protection", "1; mode=block")
		req.Header().Add("Content-Security-Policy", "default-src 'none'")
		req.Header().Add("X-Content-Security-Policy", "default-src 'none'")
		req.Header().Add("X-WebKit-CSP", "default-src 'none'")
		req.Header().Add("X-Permitted-Cross-Domain-Policies", "none")
		req.Header().Add("Referrer-Policy", "origin-when-cross-origin,strict-origin-when-cross-origin")
		req.Header().Add("Access-Control-Allow-Origin	", "*")
		req.Header().Add("Vary", "Accept-Encoding")

		next.ServeHTTP(req, res)
	})
}
