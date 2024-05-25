package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sweatybridge/pijoy/internal/api"
)

// Assigned using -ldflags
var version string

type JoystickServer struct {
}

// GetHealth implements api.ServerInterface.
func (vs *JoystickServer) GetHealth(ctx echo.Context) error {
	resp := api.Health{
		Status:  api.Ready,
		Version: version,
	}
	return ctx.JSON(http.StatusOK, resp)
}
