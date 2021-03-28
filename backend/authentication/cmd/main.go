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
		token authentication.TokenManager
	)

	if *inmemory {
		users = inmen.NewUserRepository()
	} else {
		// TODO: postgresl
	}

	token = jwt.NewTokenManager(
		secretKey,
		time.Second * 5,
		time.Second * 20)

	var au auth.Service
	au = auth.NewService(users, token)

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