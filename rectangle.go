package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// rectangle is the smallest unit of UI. All other elements are made up of rectangles.
type rectangle struct {
	img *ebiten.Image
	x   int
	y   int
	clr color.Color
}

func (r *rectangle) DrawRectangle(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.x), float64(r.y))
	_, _, _, a := r.clr.RGBA()
	op.ColorScale.ScaleAlpha((float32(uint8(a>>8)) / math.MaxUint8))
	screen.DrawImage(r.img, op)
}

func makeRectangle(size coords, clr color.Color, position coords) *rectangle {
	img := ebiten.NewImage(size.x, size.y)
	img.Fill(clr)
	return &rectangle{img, position.x, position.y, clr}
}
