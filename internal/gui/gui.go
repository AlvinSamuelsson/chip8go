package gui

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

const (
	winX         float64 = 64
	winY         float64 = 32
	screenWidth  float64 = 1024
	screenHeight float64 = 768
)

type Window struct {
	*opengl.Window
}

func CreateWindow() (*Window, error) {
	cfg := opengl.WindowConfig{
		Title:  "Alvins CHIP8",
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
	}

	win, err := opengl.NewWindow(cfg)
	if err != nil {
		return nil, err
	}

	window := &Window{
		Window: win,
	}

	return window, nil
}
