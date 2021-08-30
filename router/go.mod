module github.com/tiagomol1/http-server-golang/router

go 1.13

replace github.com/tiagomol1/http-server-golang/security => ../security

replace github.com/tiagomol1/http-server-golang/controllers/users => ../controllers/users

require (
	github.com/tiagomol1/http-server-golang/controllers/users v0.0.0-00010101000000-000000000000
	github.com/tiagomol1/http-server-golang/security v0.0.0-00010101000000-000000000000
)
