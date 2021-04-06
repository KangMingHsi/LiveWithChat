package server

import (
	"authentication/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server holds the dependencies for a echo server.
type Server struct {
	Auth  auth.Service

	Host	*echo.Echo
}

// New returns a new echo server.
func New(au auth.Service) *Server {
	s := &Server{
		Auth:  au,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	apiG := e.Group("/api")
	h := authHandler{s: au}
	h.addGroup(apiG)

	s.Host = e
	return s
}
