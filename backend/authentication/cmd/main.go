package main

import (
	"authentication"
	"authentication/auth"
	"authentication/inmen"
	"authentication/jwt"
	"authentication/server"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	defaultPort              = "8080"
	defaultSecret            = "secret"
)

func main() {
	var (
		addr   = envString("PORT", defaultPort)
		secretKey = envString("SECRET", defaultSecret)
		inmemory          = flag.Bool("inmem", false, "use in-memory repositories")
	)

	flag.Parse()

	var (
		users authentication.UserRepository
	)

	if *inmemory {
		users = inmen.NewUserRepository()
	} else {
		// TODO: postgresl
	}

	var js authentication.TokenService
	js = jwt.NewJWTService(secretKey, time.Hour * 6)

	var au auth.Service
	au = auth.NewService(users, js)

	srv := server.New(au)
	srv.Host.Logger.Fatal(
		srv.Host.Start(fmt.Sprintf(":%s", addr)))
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}