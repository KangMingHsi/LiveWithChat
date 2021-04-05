package server

import (
	"authentication"
	"authentication/auth"
	"errors"
	"net/http"
	"strings"

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
	g.POST("/check", h.check)
	g.POST("/refresh", h.refresh)
}

func (h *authHandler) register(c echo.Context) error {
	values, err := c.FormParams()
	if err != nil {
		return toEchoHttpError(err)
	}

	email := values.Get("email")
	gender := values.Get("gender")
	nickname := values.Get("nickname")
	password := values.Get("password")
	
	id, err := h.s.Register(email, gender, nickname, password)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.JSON(http.StatusOK, &struct{ID authentication.MemberID `json:"user_id"`}{ID: id})
}

func (h *authHandler) login(c echo.Context) error {
	values, err := c.FormParams()
	if err != nil {
		return toEchoHttpError(err)
	}

	email := values.Get("email")
	password := values.Get("password")
	ipAddr := c.RealIP()
	accessToken, refreshToken, err := h.s.Login(
		email,
		password,
		ipAddr,
	)

	if err != nil {
		return toEchoHttpError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"AccessToken": accessToken,
		"RefreshToken": refreshToken,
	})
}

func (h *authHandler) logout(c echo.Context) error {
	accessToken, err := h.getToken(c)
	if err != nil {
		return toEchoHttpError(err)
	}
	
	err = h.s.Logout(accessToken)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.String(http.StatusOK, "Successfully logout")
}

func (h *authHandler) check(c echo.Context) error {
	accessToken, err := h.getToken(c)
	if err != nil {
		return toEchoHttpError(err)
	}

	err = h.s.Check(accessToken)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.String(http.StatusOK, "Is valid")
}

func (h *authHandler) refresh(c echo.Context) error {
	refreshToken, err := h.getToken(c)
	if err != nil {
		return toEchoHttpError(err)
	}

	newAccessToken, newRefreshToken, err := h.s.Refresh(refreshToken)
	if err != nil {
		return toEchoHttpError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"AccessToken": newAccessToken,
		"RefreshToken": newRefreshToken,
	})
}

func (h *authHandler) changePassword(c echo.Context) error {
	accessToken, err := h.getToken(c)
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

func (h *authHandler) getToken(c echo.Context) (string, error) {
	authorization := c.Request().Header.Get("Authorization")

	if authorization == "" {
		return "", errors.New("No authorization token")
	}

	accessToken := strings.Split(authorization, " ")[1]
	return accessToken, nil
}

func toEchoHttpError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}