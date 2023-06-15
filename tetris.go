package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface and is composed of one active shape that the player controls and needs to place, and a collection of squares that have already been placed
type Game struct {
	shape                  *shape
	squares                []*square
	updates                int
	interval               int
	updates_since_movement int
}

func (g *Game) Update() error {
	g.updates++
	g.updates_since_movement++
	if g.updates_since_movement >= g.interval {
		g.updates_since_movement = 0
		g.shape.move()
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
		shape:    shp,
		interval: 60,
	}
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(initGame()); err != nil {
		log.Fatal(err)
	}
}
