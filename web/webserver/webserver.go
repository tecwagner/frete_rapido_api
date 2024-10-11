package web

import (
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	Router        *gin.Engine
	Handlers      map[string]gin.HandlerFunc
	WebServerPort string
}

func addRoute(router *gin.Engine, method string, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		router.GET(path, handler)
	case "POST":
		router.POST(path, handler)
	case "PUT":
		router.PUT(path, handler)
	case "DELETE":
		router.DELETE(path, handler)
	default:

	}
}

func NewWebServer(webServerPort string) *WebServer {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return &WebServer{
		Router:        r,
		Handlers:      make(map[string]gin.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandlerPublic(method string, path string, handler gin.HandlerFunc) error {
	s.Handlers[path] = handler
	addRoute(s.Router, method, path, handler)
	return nil
}

func (s *WebServer) Start() error {
	err := s.Router.Run(s.WebServerPort)
	if err != nil {
		return err
	}
	return nil
}
