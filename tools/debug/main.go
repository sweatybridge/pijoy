package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		return err
	}
	defer rpio.Close()

	// Use pin 6, corresponds to physical pin 31
	pinDown := rpio.Pin(6)
	// Set pin to output mode
	pinDown.Output()

	// Toggle pin 20 times
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for i := 0; i < 20; i++ {
		select {
		case <-ticker.C:
			fmt.Fprintln(os.Stderr, "Toggle", i)
			pinDown.Toggle()
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
