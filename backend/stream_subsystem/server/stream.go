package server

import (
	"net/http"
	"stream_subsystem/stream"

	"github.com/labstack/echo/v4"
)

type streamHandler struct {
	s  stream.Service
}

func (h *streamHandler) addGroup(e *echo.Group) {
	g := e.Group("/v1/stream")
	g.GET("/videos", h.getVideos)
	g.POST("/videos", h.uploadVideo)
	g.PATCH("/videos", h.updateVideo)
	g.DELETE("/videos", h.deleteVideo)
}

func (h *streamHandler) getVideos(c echo.Context) error {
	videos := h.s.GetVideos()
	return c.JSON(http.StatusOK, videos)
}

func (h *streamHandler) uploadVideo(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	videoType := c.FormValue("video_type")

	uploadedVideo, err := c.FormFile("video")
	if err != nil {
		return toEchoHttpError(err)
	}

	err = h.s.UploadVideo(title, description, "", videoType, uploadedVideo)
	if err != nil {
		return toEchoHttpError(err)
	}
	return c.String(http.StatusOK, "Video uploads successfully")
}

func (h *streamHandler) updateVideo(c echo.Context) error {
	values, err := c.FormParams()
	if err != nil {
		return toEchoHttpError(err)
	}

	var vid string
	data := map[string]interface{}{}
	for key := range values {
		if key == "vid" {
			vid = values.Get(key)
		} else {
			data[key] = values.Get(key)
		}
	}
	err = h.s.UpdateVideo(vid, data)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.String(http.StatusOK, "Video updates successfully")
}

func (h *streamHandler) deleteVideo(c echo.Context) error {
	vid := c.FormValue("vid")
	err := h.s.DeleteVideo(vid)
	if err != nil {
		return toEchoHttpError(err)
	}
	return c.String(http.StatusOK, "Video deletes successfully")
}

func toEchoHttpError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}