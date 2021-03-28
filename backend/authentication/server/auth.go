package server

import (
	"authentication"
	"authentication/auth"
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
	g.POST("/change-password", h.changePassword)
	g.POST("/check-and-refresh", h.checkAndRefresh)
}

func (h *authHandler) register(c echo.Context) error {
	password := c.FormValue("password")
	id, err := h.s.Register(password)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.JSON(http.StatusOK, &struct{ID authentication.MemberID `json:"user_id"`}{ID: id})
}

func (h *authHandler) login(c echo.Context) error {
	id := c.FormValue("user_id")
	password := c.FormValue("password")
	ipAddr := c.RealIP()
	accessToken, refreshToken, err := h.s.Login(
		authentication.MemberID(id),
		password,
		ipAddr,
	)

	if err != nil {
		return toEchoHttpError(err)
	}

	c.SetCookie(&http.Cookie{
		Name: "AccessToken",
		Value: accessToken,
	})

	c.SetCookie(&http.Cookie{
		Name: "RefreshToken",
		Value: refreshToken,
	})

	return c.String(http.StatusOK, "Successfully login")
}

func (h *authHandler) logout(c echo.Context) error {
	accessToken, err := h.refreshImpl(c)
	if err != nil {
		return toEchoHttpError(err)
	}

	err = h.s.Logout(accessToken)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.String(http.StatusOK, "Successfully logout")
}

func (h *authHandler) checkAndRefresh(c echo.Context) error {
	_, err := h.refreshImpl(c)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.String(http.StatusOK, "Successfully check and refresh")
}

func (h *authHandler) changePassword (c echo.Context) error {
	accessToken, err := h.refreshImpl(c)
	if err != nil {
		return toEchoHttpError(err)
	}

	newPassword := c.FormValue("newPassword")
	err = h.s.ChangePassword(newPassword, accessToken)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.String(http.StatusOK, "Successfully change password")
}

func (h *authHandler) refreshImpl(c echo.Context) (string, error) {
	accessCookie, err := c.Cookie("AccessToken")
	if err != nil {
		return "", echo.ErrCookieNotFound
	}

	refreshCookie, err := c.Cookie("RefreshToken")
	if err != nil {
		return "", echo.ErrCookieNotFound
	}

	accessToken, refreshToken, err := h.s.CheckAndRefresh(accessCookie.Value, refreshCookie.Value)
	if err != nil {
		return "", err
	}

	c.SetCookie(&http.Cookie{
		Name: "AccessToken",
		Value: accessToken,
	})

	c.SetCookie(&http.Cookie{
		Name: "RefreshToken",
		Value: refreshToken,
	})

	return accessToken, nil
}


func toEchoHttpError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}