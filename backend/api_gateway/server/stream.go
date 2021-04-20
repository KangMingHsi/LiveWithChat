package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

type StreamHandler struct {
	authReq *http.Request
	client *http.Client

	accessibles map[string]int
}

type CheckInfo struct {
	RoleLevel int        
	UserID    string
}

// StreamProcess middleware checks jwt token first
func (h *StreamHandler) StreamProcess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		query := fmt.Sprintf("%s%s", req.Method, req.URL.Path)

		if h.needAuthorization(query) {
			headers := req.Header
			for key := range headers {
				h.authReq.Header.Set(key, headers.Get(key))
			}
			// h.authReq.Header.Set

			resp, err := h.client.Do(h.authReq)
			if err != nil {
				return c.String(http.StatusForbidden, err.Error())
			}
			defer resp.Body.Close()

			if resp.StatusCode < 200 || resp.StatusCode >= 400 {
				return c.String(http.StatusUnauthorized, "token is invalid or expired")
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				c.String(http.StatusForbidden, err.Error())
			}
			var data map[string]interface{} // TopTracks
			err = json.Unmarshal(body, &data)
			if err != nil {
				c.String(http.StatusForbidden, err.Error())
			}

			if !h.passAuthorization(query, int(data["role_level"].(float64))) {
				return c.String(http.StatusForbidden, "the level of this user cannot do this method")
			}
		}
		return next(c)
	}
}

func (h *StreamHandler) needAuthorization(query string) bool {
	if v, ok := h.accessibles[query]; ok {
		return v > 0
	}
	return false
}

func (h *StreamHandler) passAuthorization(query string, level int) bool {
	if v, ok := h.accessibles[query]; ok {
		return level >= v
	}
	return false
}

func NewStreamHandler(authURL *url.URL, accessibles map[string]int) *StreamHandler {
	req, _ := http.NewRequest("POST", authURL.Path, nil)
	return &StreamHandler{
		authReq: req,
		client: &http.Client{},
		accessibles: accessibles,
	}
}