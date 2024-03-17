package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type App struct {
	canvas         *Image
	palette        *Image
	preview        *Image
	selectedColour color.Color
}

func (a *App) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// check pixels
		pixel := a.canvas.WasClicked(x, y)
		if pixel != nil {
			pixel.SetColour(a.selectedColour)
		}

		// check palette
		pixel = a.palette.WasClicked(x, y)
		if pixel != nil {
			a.selectedColour = pixel.colour
		}
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	a.canvas.Render(screen)
	a.palette.Render(screen)
	a.preview.Render(screen)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
