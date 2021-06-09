package server

import (
	"chat_subsystem"
	"chat_subsystem/chat"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server holds the dependencies for a echo server.
type Server struct {
	Chat	chat.Service
	Host	*echo.Echo
}

// New returns a new echo server.
func New(
	chatService   chat.Service,
	tokenManager chat_subsystem.TokenManager,
) *Server {
	s := &Server{
		Chat: chatService,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	apiG := e.Group("/api")
	chatH := chatHandler{
		s: chatService,
		tokenManager: tokenManager,
	}
	chatH.addGroup(apiG)

	s.Host = e
	return s
}

func toEchoHttpError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}