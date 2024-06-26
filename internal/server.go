package internal

import (
	"net/http"
	"time"

	"github.com/go-errors/errors"
	"github.com/labstack/echo/v4"
	"github.com/stianeikeland/go-rpio/v4"
	"github.com/sweatybridge/pijoy/internal/api"
)

// Assigned using -ldflags
var version string

type JoystickServer struct {
	btnToPin map[api.Button]rpio.Pin
}

func NewJoystickServer() *JoystickServer {
	joystick := &JoystickServer{
		// TODO: make pin numbers configurable
		btnToPin: map[api.Button]rpio.Pin{
			api.Up:    rpio.Pin(12),
			api.Down:  rpio.Pin(6),
			api.Left:  rpio.Pin(5),
			api.Right: rpio.Pin(13),
		},
	}
	// Setup all pins to output mode
	for _, pin := range joystick.btnToPin {
		pin.Output()
	}
	return joystick
}

// PressJoystick implements api.ServerInterface.
func (vs *JoystickServer) PressJoystick(ctx echo.Context, button api.Button) error {
	// TODO: add unit tests for joystick server
	pin, ok := vs.btnToPin[button]
	if !ok {
		return errors.New("Uninitialized button: " + button)
	}
	pin.Toggle()
	defer pin.Toggle()
	// Wait for at most 10 seconds before release
	timer := time.NewTimer(time.Second * 10)
	select {
	case <-timer.C:
	case <-ctx.Request().Context().Done():
		timer.Stop()
	}
	return ctx.NoContent(http.StatusNoContent)
}

// GetHealth implements api.ServerInterface.
func (vs *JoystickServer) GetHealth(ctx echo.Context) error {
	resp := api.Health{
		Status:  api.Ready,
		Version: version,
	}
	return ctx.JSON(http.StatusOK, resp)
}
