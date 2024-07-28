package main

import (
	"os"

	"github.com/go-errors/errors"
	"github.com/stianeikeland/go-rpio/v4"
	"golang.org/x/term"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type Joystick struct {
	Up    rpio.Pin
	Down  rpio.Pin
	Left  rpio.Pin
	Right rpio.Pin
}

func NewJoystick() Joystick {
	return Joystick{
		Up:    rpio.Pin(12),
		Down:  rpio.Pin(6),
		Left:  rpio.Pin(5),
		Right: rpio.Pin(13),
	}
}

func run() error {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		return errors.Errorf("failed to open gpio: %w", err)
	}
	defer rpio.Close()

	// Set pin to output mode
	joystick := NewJoystick()
	joystick.Up.Output()
	joystick.Down.Output()
	joystick.Left.Output()
	joystick.Right.Output()

	// switch stdin into 'raw' mode
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return errors.Errorf("failed to set raw mode: %w", err)
	}
	defer term.Restore(fd, oldState)

	b := make([]byte, 1)
	for b[0] != 'q' {
		if _, err := os.Stdin.Read(b); err != nil {
			return errors.Errorf("failed to read char: %w", err)
		}
		switch b[0] {
		case 'w':
			joystick.Up.Toggle()
		case 's':
			joystick.Down.Toggle()
		case 'a':
			joystick.Left.Toggle()
		case 'd':
			joystick.Right.Toggle()
		}
	}

	return nil
}
