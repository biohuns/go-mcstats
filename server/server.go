package server

import (
	"fmt"

	"github.com/biohuns/go-minecraft-stats/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Server is halding echo instance
type Server struct {
	Instance *echo.Echo
	config   *config.Cfg
}

type configuredContext struct {
	echo.Context
	config *config.Cfg
}

//NewServer return server
func NewServer(cfg *config.Cfg) *Server {
	e := echo.New()
	// set logger
	e.Use(middleware.Logger())
	// set config to context
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &configuredContext{
				c,
				cfg,
			}
			return h(cc)
		}
	})
	// set routing
	e.GET("/stats", statsHandler)

	return &Server{
		config:   cfg,
		Instance: e,
	}
}

//Start start server
func (s *Server) Start() {
	port := fmt.Sprintf(":%d", s.config.Server.Port)
	s.Instance.Logger.Fatal(s.Instance.Start(port))
}
