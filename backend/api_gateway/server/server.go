package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server holds the dependencies for a echo server.
type Server struct {
	Host	*echo.Echo
}

type CheckInfo struct {
	RoleLevel int        
	UserID    string
}

type requestHandler struct{
	authReq *http.Request
	client *http.Client

	accessibles map[string]bool
}

func (h *requestHandler) needAuthorization(query string) bool {
	if v, ok := h.accessibles[query]; ok {
		return v
	}
	return false
}

func AuthTemplate(h *requestHandler) (func(next echo.HandlerFunc) echo.HandlerFunc) {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
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
}

// New returns a new echo server.
func New(authURL, streamURL, chatURL *url.URL) *Server {
	s := &Server{}

	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authProxy := httputil.NewSingleHostReverseProxy(authURL)
	e.Any("/api/auth/*", echo.WrapHandler(authProxy))

	streamProxy := httputil.NewSingleHostReverseProxy(streamURL)
	chatProxy := httputil.NewSingleHostReverseProxy(chatURL)

	v1Group := e.Group("/api/v1")
	authURL = &url.URL{
		Path: fmt.Sprintf("%s://%s/api/auth/check", authURL.Scheme, authURL.Host),
	}
	req, _ := http.NewRequest("POST", authURL.Path, nil)

	streamHandler := &requestHandler{
		authReq: req,
		client: &http.Client{},
		accessibles: map[string]bool{
			"GET/api/v1/stream/videos": false,
			"PATCH/api/v1/stream/videos": true,
			"POST/api/v1/stream/videos": true,
			"DELETE/api/v1/stream/videos": true,
		},
	}
	v1Stream := v1Group.Group("/stream")
	v1Stream.Use(AuthTemplate(streamHandler))
	v1Stream.Any("/*", echo.WrapHandler(streamProxy))

	chatHandler := &requestHandler{
		authReq: req,
		client: &http.Client{},
		accessibles: map[string]bool{
			"GET/api/v1/chat/messages": false,
			"PATCH/api/v1/chat/messages": true,
			"POST/api/v1/chat/messages": true,
			"DELETE/api/v1/chat/messages": true,
		},
	}
	v1Chat := v1Group.Group("/chat")
	v1Chat.Use(AuthTemplate(chatHandler))
	v1Chat.Any("/*", echo.WrapHandler(chatProxy))

	s.Host = e
	return s
}
