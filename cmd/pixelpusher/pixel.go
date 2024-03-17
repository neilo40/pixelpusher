package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// Pixel is an singly coloured square image
// x and y are relative to the
type Pixel struct {
	image  *ebiten.Image
	colour color.Color
}

func (p *Pixel) SetColour(colour color.Color) {
	p.colour = colour
	p.image.Fill(p.colour)
}

func (p *Pixel) SetRandomColour() {
	p.SetColour(randomColour())
}

func randomColour() color.Color {
	return color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 0xff}
}
