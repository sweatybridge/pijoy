package main

import (
	"embed"
	"net"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

var (
	//go:embed certs/raspberrypi.local.crt
	sslCert []byte
	//go:embed certs/raspberrypi.local.key
	sslKey []byte
	//go:embed whip
	whipClient embed.FS
)

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
	e.Use(WithCookie)
	// Register route handlers
	e.StaticFS("/webcam", whipClient)
	return e.StartTLS(":8443", sslCert, sslKey)
}

func WithCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetCookie(&http.Cookie{
			Name:     "whip-url",
			Value:    os.Getenv("WHIP_URL"),
			Path:     "/webcam/whip",
			Secure:   true,
			HttpOnly: true,
		})
		return next(c)
	}
}
