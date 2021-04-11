package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	defaultPort              = "8080"
	
	defaultAuthHost			 = "auth_subsystem"
	defaultAuthPort			 = "8080"
)

func main() {
	var (
		addr   = envString("PORT", defaultPort)

		authHost = envString("AUTH_HOST", defaultAuthHost)
		authPort = envString("AUTH_PORT", defaultAuthPort)
	)

	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	authProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", authHost, authPort),
	})

	e.Any("/api/auth/*", echo.WrapHandler(authProxy))
	e.Logger.Fatal(
		e.Start(fmt.Sprintf(":%s", addr)))
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}