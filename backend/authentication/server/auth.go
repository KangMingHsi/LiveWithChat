package server

import (
	"authentication"
	"authentication/auth"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	s  auth.Service
}

func (h *authHandler) addGroup(e *echo.Echo) {
	g := e.Group("/auth")
	g.PUT("/register", h.register)
	g.POST("/login", h.login)
	g.POST("/logout", h.logout)
}

func (h *authHandler) register(c echo.Context) error {
	password := c.FormValue("password")
	id, err := h.s.Register(password)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("%v", err))
	}

	return c.JSON(http.StatusOK, &struct{ID authentication.MemberID `json:"user_id"`}{ID: id})
}

func (h *authHandler) login(c echo.Context) error {
	ID := c.FormValue("user_id")
	password := c.FormValue("password")
	token, err := h.s.Login(authentication.MemberID(ID), password)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("%v", err))
	}

	return c.JSON(http.StatusOK, &struct{Token string `json:"token"`}{Token: token})
}

func (h *authHandler) logout(c echo.Context) error {

	return c.String(http.StatusOK, "Log out.")
}