package router

import (
	"net/http"

	users "github.com/tiagomol1/http-server-golang/controllers/users"
	security "github.com/tiagomol1/http-server-golang/security"
)

func Router() {
	http.HandleFunc("/api/v1/users", security.Auth(users.Index))
}
