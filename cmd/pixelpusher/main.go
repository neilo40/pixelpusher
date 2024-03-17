package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	canvasPixels := make([]*Pixel, 64)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			p := &Pixel{
				image: ebiten.NewImage(1, 1),
			}
			p.SetColour(color.White)
			canvasPixels[i*8+j] = p
		}
	}
	canvas := &Image{
		x:       0,
		y:       0,
		rows:    8,
		columns: 8,
		pixels:  canvasPixels,
		scale:   20,
	}

	palettePixels := make([]*Pixel, 8)
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			p := &Pixel{
				image: ebiten.NewImage(1, 1),
			}
			p.SetRandomColour()
			palettePixels[i+j*2] = p
		}
	}
	palette := &Image{
		x:       300,
		y:       0,
		rows:    4,
		columns: 2,
		pixels:  palettePixels,
		scale:   10,
	}

	preview := &Image{
		x:       200,
		y:       10,
		rows:    8,
		columns: 8,
		pixels:  canvasPixels,
		scale:   1,
	}

	game := &App{
		canvas:         canvas,
		palette:        palette,
		preview:        preview,
		selectedColour: color.Black,
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pixel Pusher")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
