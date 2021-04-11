package server

import (
	"stream_subsystem/stream"
	"net/http"

	"github.com/labstack/echo/v4"
)

type streamHandler struct {
	s  stream.Service
}

func (h *streamHandler) addGroup(e *echo.Group) {
	g := e.Group("/stream/v1")
	g.GET("videos", h.getVideos)
}

func (h *streamHandler) getVideos(c echo.Context) error {
	return nil
}

func toEchoHttpError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}