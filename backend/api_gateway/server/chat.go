package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	authReq *http.Request
	client *http.Client

	accessibles map[string]bool
}

// ChatProcess middleware checks jwt token first
func (h *ChatHandler) ChatProcess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		query := fmt.Sprintf("%s%s", req.Method, req.URL.Path)

		if h.needAuthorization(query) {
			headers := req.Header
			for key := range headers {
				h.authReq.Header.Set(key, headers.Get(key))
			}

			resp, err := h.client.Do(h.authReq)
			if err != nil {
				return c.String(http.StatusForbidden, err.Error())
			}
			defer resp.Body.Close()

			if resp.StatusCode < 200 || resp.StatusCode >= 400 {
				return c.String(http.StatusUnauthorized, "token is invalid or expired")
			}
			// body, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	c.String(http.StatusForbidden, err.Error())
			// }
			// var data map[string]interface{}
			// err = json.Unmarshal(body, &data)
			// if err != nil {
			// 	c.String(http.StatusForbidden, err.Error())
			// }

			// roleLevel, _ := strconv.P.ParseFloat(data["role_level"].(float64), 64)
			// h.authReq.PostForm.Add("user_id", data["user_id"].(string))
			// h.authReq.PostForm.Add("role_level", roleLevel)
		}
		return next(c)
	}
}

func (h *ChatHandler) needAuthorization(query string) bool {
	if v, ok := h.accessibles[query]; ok {
		return v
	}
	return false
}

func NewChatHandler(authURL *url.URL, accessibles map[string]bool) *ChatHandler {
	req, _ := http.NewRequest("POST", authURL.Path, nil)
	return &ChatHandler{
		authReq: req,
		client: &http.Client{},
		accessibles: accessibles,
	}
}