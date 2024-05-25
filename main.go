package main

import (
	"net"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sweatybridge/pijoy/internal"
	"github.com/sweatybridge/pijoy/internal/api"
)

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=internal/api/types.cfg.yml docs/api.yml
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=internal/api/server.cfg.yml docs/api.yml

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {
	// Initialise middleware
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return net.ParseIP(c.RealIP()).IsLoopback()
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// Start server
	handlers := &internal.JoystickServer{}
	api.RegisterHandlers(e, handlers)
	return e.Start(":8080")
}
