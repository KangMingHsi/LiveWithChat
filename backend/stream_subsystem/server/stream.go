package server

import (
	"net/http"
	"stream_subsystem"
	"stream_subsystem/stream"
	"strings"

	"github.com/labstack/echo/v4"
)

type streamHandler struct {
	s  stream.Service
	tokenManager  stream_subsystem.TokenManager
}

func (h *streamHandler) addGroup(e *echo.Group) {
	g := e.Group("/v1/stream")
	g.Use(h.middleware)
	g.GET("/videos", h.getVideos)
	g.POST("/videos", h.uploadVideo)
	g.PATCH("/videos", h.updateVideo)
	g.DELETE("/videos", h.deleteVideo)
}

func (h *streamHandler) middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		if req.Method != "GET" {
			authorization := req.Header.Get("Authorization")

			if authorization == "" {
				return c.String(http.StatusForbidden, "no token")
			}

			accessToken := strings.Replace(authorization, "Bearer ", "", 1)
			claim, err := h.tokenManager.Verify(accessToken)
			if err != nil {
				return c.String(http.StatusForbidden, err.Error())
			}

			cliamMap := claim.ConvertToMap()
			roleLevel := cliamMap["RoleLevel"].(int)

			if roleLevel < 1 {
				return c.String(http.StatusUnauthorized, "role level is not enough.")
			}

			req.PostForm.Add("user_id", cliamMap["UserID"].(string))
		}

		return next(c)
	}
}

func (h *streamHandler) getVideos(c echo.Context) error {
	videos := h.s.GetVideos()
	return c.JSON(http.StatusOK, videos)
}

func (h *streamHandler) uploadVideo(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	videoType := c.FormValue("video_type")
	uid := c.FormValue("user_id")

	uploadedVideo, err := c.FormFile("video")
	if err != nil {
		return toEchoHttpError(err)
	}

	err = h.s.UploadVideo(title, description, uid, videoType, uploadedVideo)
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