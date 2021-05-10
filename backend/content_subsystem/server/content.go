package server

import (
	"content_subsystem/content"
	"net/http"

	"github.com/labstack/echo/v4"
)

type contentHandler struct {
	s  content.Service
}

func (h *contentHandler) addGroup(e *echo.Group) {
	g := e.Group("/v1/content")

	g.GET("/videos", h.getVideoInfo)
	g.POST("/videos", h.saveVideo)
	g.DELETE("/videos", h.deleteVideo)
}

func (h *contentHandler) getVideoInfo(c echo.Context) error {
	vid := c.FormValue("vid")
	info, err := h.s.GetContentInfo(vid)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.JSON(http.StatusOK, info)
}

func (h *contentHandler) saveVideo(c echo.Context) error {
	vid := c.FormValue("vid")
	videoType := c.FormValue("video_type")
	video, err := c.FormFile("video")
	if err != nil {
		return toEchoHttpError(err)
	}

	err = h.s.Save(vid, videoType, video)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.String(http.StatusOK, "Successfully save")
}

func (h *contentHandler) deleteVideo(c echo.Context) error {
	vid := c.Param("id")
	err := h.s.Delete(vid)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.String(http.StatusOK, "Successfully delete")
}

func toEchoHttpError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}