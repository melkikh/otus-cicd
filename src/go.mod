module pocsrv

go 1.13

// +heroku goVersion go1.13
// +heroku install ./cmd/...

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/google/logger v1.0.1
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
)
