package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface and is composed of one active shape that the player controls and needs to place, and a collection of squares that have already been placed
type Game struct {
	shape   *shape
	squares []*square
}

func (g *Game) Update() error {
	for _, sqr := range *g.shape {
		for _, rect := range *sqr {
			rect.y++
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, sqr := range *g.shape {
		for _, rect := range *sqr {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(rect.x), float64(rect.y))
			screen.DrawImage(rect.img, op)
		}
	}
	// ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func initGame() *Game {
	sqr := MakeSquare(color.RGBA{0xff, 0, 0, 0xff}, coords{0, 0})
	shp := &shape{sqr}
	return &Game{
		shape: shp,
	}
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(initGame()); err != nil {
		log.Fatal(err)
	}
}
