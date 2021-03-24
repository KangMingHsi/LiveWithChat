package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func login(c echo.Context) error {
	return nil
}

func register(c echo.Context) error {
	return nil
}

func logout(c echo.Context) error {
	return nil
}

func main() {
	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)
	e.PUT("/register", register)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("/logout", logout)

	e.Logger.Fatal(e.Start(":8080"))
}