package main

import "github.com/hajimehoshi/ebiten/v2"

// Image is a collection of pixels Rows*Columns in size, with an optionally rendered Grid and Outline
type Image struct {
	x, y    int // coord of top left of image
	rows    int
	columns int
	//grid    bool // draw a line between all pixels
	//outline bool // draw an outline around the image
	pixels []*Pixel // expect length to be rows * columns
	scale  int      // native pixels are 1x1.  Scale of e.g. 4 would draw them at 4x4 screen pixels in size
}

// if the image was clicked, return the pixel that was under the mouse, else nil
func (i *Image) WasClicked(mouseX, mouseY int) *Pixel {
	for r := 0; r < i.rows; r++ {
		for c := 0; c < i.columns; c++ {
			pLeft := i.x + c*i.scale
			pRight := pLeft + i.scale
			pTop := i.y + r*i.scale
			pBottom := pTop + i.scale
			if mouseX > pLeft && mouseX < pRight && mouseY > pTop && mouseY < pBottom {
				pNum := r*i.columns + c
				return i.pixels[pNum]
			}
		}
	}
	return nil
}

func (i *Image) Render(screen *ebiten.Image) {
	for r := 0; r < i.rows; r++ {
		for c := 0; c < i.columns; c++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(float64(i.scale), float64(i.scale))
			op.GeoM.Translate(float64(i.x+(c*i.scale)), float64(i.y+(r*i.scale)))
			screen.DrawImage(i.pixels[(r*i.columns)+c].image, op)
		}
	}
}
