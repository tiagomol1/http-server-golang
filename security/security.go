package security

import (
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if hasPermission(req) {
			next.ServeHTTP(w, req)
		} else {
			http.Error(w, "Unauthorized!", 401)
		}
	}
	return http.HandlerFunc(fn)
}

func hasPermission(req *http.Request) bool {
	authorization := req.Header.Get("Authorization")
	return authorization == "Bearer 1qaz2wsx!@"
}
