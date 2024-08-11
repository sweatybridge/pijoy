package main

import (
	"embed"
	"html/template"
	"io"
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
	//go:embed assets
	whipClient embed.FS
	//go:embed templates
	htmlTemplates embed.FS
)

type RendererFunc func(w io.Writer, name string, data interface{}, c echo.Context) error

func (f RendererFunc) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return f(w, name, data, c)
}

func start() error {
	tmpl, err := template.ParseFS(htmlTemplates, "**/*")
	if err != nil {
		return err
	}
	// Initialise middleware
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return net.ParseIP(c.RealIP()).IsLoopback()
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// Register route handlers
	e.StaticFS("/", whipClient)
	e.Renderer = RendererFunc(func(w io.Writer, name string, data interface{}, c echo.Context) error {
		return tmpl.ExecuteTemplate(w, name, data)
	})
	e.GET("/webcam/whip", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", os.Getenv("WHIP_URL"))
	})
	return e.StartTLS(":8443", sslCert, sslKey)
}
