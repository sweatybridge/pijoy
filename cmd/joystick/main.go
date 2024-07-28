package main

import (
	"net"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-errors/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/stianeikeland/go-rpio/v4"
	"github.com/sweatybridge/pijoy/internal"
	"github.com/sweatybridge/pijoy/internal/api"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=types.cfg.yml ../../docs/api.yml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yml ../../docs/api.yml

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		return errors.Errorf("failed to open gpio: %w", err)
	}
	defer rpio.Close()
	// Initialise middleware
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return net.ParseIP(c.RealIP()).IsLoopback()
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	swagger, err := api.GetSwagger()
	if err != nil {
		return errors.Errorf("failed to parse openapi spec: %w", err)
	}
	e.Use(echomiddleware.OapiRequestValidatorWithOptions(swagger, &echomiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		},
	}))
	// Start server
	handlers := internal.NewJoystickServer()
	api.RegisterHandlers(e, handlers)
	return e.Start(":8080")
}
