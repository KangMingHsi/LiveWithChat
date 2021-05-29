package server

import (
	"net/http"
	"strconv"
	"stream_subsystem"
	"stream_subsystem/chat"
	"strings"

	"github.com/labstack/echo/v4"
)

type chatHandler struct {
	s  chat.Service
	tokenManager  stream_subsystem.TokenManager
}

func (h *chatHandler) addGroup(e *echo.Group) {
	g := e.Group("/v1/chat")
	g.Use(h.middleware)
	g.GET("/messages", h.getMessages)
	g.POST("/messages", h.createMessage)
	g.PATCH("/messages", h.updateMessage)
	g.DELETE("/messages", h.deleteMessage)
}

func (h *chatHandler) middleware(next echo.HandlerFunc) echo.HandlerFunc {
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

			c.Set("user_id", cliamMap["UserID"].(string))
		}

		return next(c)
	}
}

func (h *chatHandler) getMessages(c echo.Context) error {
	vid := c.FormValue("vid")
	messages := h.s.GetMessages(vid)
	return c.JSON(http.StatusOK, messages)
}

func (h *chatHandler) createMessage(c echo.Context) error {
	text := c.FormValue("text")
	vid := c.FormValue("vid")
	uid := c.Get("user_id").(string)

	err := h.s.CreateMessage(text, vid, uid)
	if err != nil {
		return toEchoHttpError(err)
	}
	return c.String(http.StatusOK, "Message creates successfully")
}

func (h *chatHandler) updateMessage(c echo.Context) error {
	text := c.FormValue("text")
	idString := c.FormValue("id")
	uid := c.Get("user_id").(string)

	id, err := strconv.ParseInt(idString, 0, 64)
	if err != nil {
		return toEchoHttpError(err)
	}

	err = h.s.UpdateMessage(id, text, uid)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.String(http.StatusOK, "Message updates successfully")
}

func (h *chatHandler) deleteMessage(c echo.Context) error {
	idString := c.FormValue("id")
	uid := c.Get("user_id").(string)

	id, err := strconv.ParseInt(idString, 0, 64)
	if err != nil {
		return toEchoHttpError(err)
	}

	err = h.s.DeleteMessage(id, uid)
	if err != nil {
		return toEchoHttpError(err)
	}
	return c.String(http.StatusOK, "Message deletes successfully")
}
