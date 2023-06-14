package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// var square *ebiten.Image

type square struct {
	img *ebiten.Image
	x   int
	y   int
}

type Game struct {
	square *square
}

func (g *Game) Update() error {
	g.square.x++
	g.square.y++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.square.x), float64(g.square.y))
	screen.DrawImage(g.square.img, op)
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Hello, World!")
	s := ebiten.NewImage(100, 100)
	s.Fill(color.RGBA{0xff, 0, 0, 0xff})
	if err := ebiten.RunGame(&Game{
		square: &square{
			img: s,
			x:   0,
			y:   0,
		},
	}); err != nil {
		log.Fatal(err)
	}
}
