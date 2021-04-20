package server

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server holds the dependencies for a echo server.
type Server struct {
	Host	*echo.Echo
}

// New returns a new echo server.
func New(authURL, streamURL *url.URL) *Server {
	s := &Server{}

	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authProxy := httputil.NewSingleHostReverseProxy(authURL)
	e.Any("/api/auth/*", echo.WrapHandler(authProxy))

	streamProxy := httputil.NewSingleHostReverseProxy(streamURL)
	v1Group := e.Group("/api/v1")

	sh := NewStreamHandler(
		&url.URL{
			Path: fmt.Sprintf("%s://%s/api/auth/check", authURL.Scheme, authURL.Host),
		},
		map[string]int{
			"GET/api/v1/stream/videos": 0,
			"PATCH/api/v1/stream/videos": 1,
			"POST/api/v1/stream/videos": 1,
			"DELETE/api/v1/stream/videos": 1,
		},
	)
	v1Stream := v1Group.Group("/stream")
	v1Stream.Use(sh.StreamProcess)
	v1Stream.Any("/*", echo.WrapHandler(streamProxy))

	s.Host = e
	return s
}
