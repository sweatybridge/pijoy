package main

import (
	"context"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-errors/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/stianeikeland/go-rpio/v4"
	"github.com/sweatybridge/pijoy/internal"
	"github.com/sweatybridge/pijoy/internal/api"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=types.cfg.yml ../../docs/api.yml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yml ../../docs/api.yml

func main() {
	if err := start(context.Background()); err != nil {
		panic(err)
	}
}

func start(ctx context.Context) error {
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
	endpoint, err := listNgrokEndpoints(*swagger)
	if err != nil {
		return err
	}
	if token := os.Getenv("NGROK_AUTHTOKEN"); len(token) > 0 {
		listener, err := ngrok.Listen(ctx, endpoint[0], ngrok.WithAuthtoken(token))
		if err != nil {
			return err
		}
		e.Logger.Info("Ingress established at:", listener.URL())
		e.Listener = listener
	}
	handlers := internal.NewJoystickServer()
	api.RegisterHandlers(e, handlers)
	return e.Start(":8080")
}

func listNgrokEndpoints(swagger openapi3.T) ([]config.Tunnel, error) {
	var result []config.Tunnel
	for _, server := range swagger.Servers {
		if parsed, err := url.Parse(server.URL); err != nil {
			return nil, err
		} else if strings.HasSuffix(parsed.Host, ".ngrok-free.app") {
			endpoint := config.HTTPEndpoint(config.WithDomain(parsed.Host))
			result = append(result, endpoint)
		}
	}
	return result, nil
}
