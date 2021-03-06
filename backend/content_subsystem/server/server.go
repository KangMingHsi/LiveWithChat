package server

import (
	"content_subsystem/content"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server holds the dependencies for a echo server.
type Server struct {
	Stream  content.Service

	Host	*echo.Echo
}

// New returns a new echo server.
func New(st content.Service) *Server {
	s := &Server{
		Stream:  st,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	srcG := e.Group("/static")
	srcG.Use(middleware.CORS())
	currentDir, _ := filepath.Abs("./")
	srcG.Use(middleware.Static(
		currentDir + "/storage",
	))

	apiG := e.Group("/api")
	h := contentHandler{
		s: st,
	}
	h.addGroup(apiG)

	s.Host = e
	return s
}
