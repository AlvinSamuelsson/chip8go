package gui

import (
	"image/color"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
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

func (window *Window) DrawGraphics(gfx [64 * 32]byte) {
	window.Clear(color.Black)
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)
	// width, height := screenWidth/winX, screenHeight/winY
	for x := 0; x < 64; x++ {
		for y := 0; y < 32; y++ {
			imd.Push(pixel.V(100, 100))
			imd.Push(pixel.V(900, 900))
			imd.Rectangle(0)
		}
	}
	imd.Draw(window)
	window.Update()
}
