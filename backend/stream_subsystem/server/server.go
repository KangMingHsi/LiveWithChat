package server

import (
	"net/http"
	"stream_subsystem"
	"stream_subsystem/stream"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server holds the dependencies for a echo server.
type Server struct {
	Stream  stream.Service
	Host	*echo.Echo
}

// New returns a new echo server.
func New(
	streamService stream.Service,
	tokenManager stream_subsystem.TokenManager,
) *Server {
	s := &Server{
		Stream:  streamService,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	apiG := e.Group("/api")
	streamH := streamHandler{
		s: streamService,
		tokenManager: tokenManager,
	}
	streamH.addGroup(apiG)

	s.Host = e
	return s
}

func toEchoHttpError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}