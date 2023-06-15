package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface and is composed of one active shape that the player controls and needs to place, and a collection of squares that have already been placed
type Game struct {
	shape                  *shape
	squares                []*square
	frame                  [4]*rectangle
	updates                int
	interval               int
	updates_since_movement int
	shapes                 shapes
}

func InitShapes() shapes {
	arr := shapes{}
	shp := shape{}
	shp[0] = MakeSquare(color.RGBA{0xff, 0, 0, 0xff}, coords{0, 0})
	shp[1] = MakeSquare(color.RGBA{0xff, 0, 0, 0xff}, coords{0, 48})
	shp[2] = MakeSquare(color.RGBA{0xff, 0, 0, 0xff}, coords{0, 2 * 48})
	shp[3] = MakeSquare(color.RGBA{0xff, 0, 0, 0xff}, coords{0, 3 * 48})
	return arr
}

func (g *Game) Update() error {
	g.updates++
	g.updates_since_movement++
	if g.updates_since_movement >= g.interval {
		g.updates_since_movement = 0
		g.shape.move()
		if g.ShapeHasBottomContact() {
			g.TransferShapeToSquares()
			g.spawnShape()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawGameFrame(screen)
	// Draw shape
	for _, sqr := range *g.shape {
		for _, rect := range *sqr {
			DrawRectangle(screen, rect)
		}
	}
	// Draw other squares
	for _, sqr := range g.squares {
		for _, rect := range *sqr {
			DrawRectangle(screen, rect)
		}
	}
	// ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 1200
}

func (g *Game) Rotate() {
	// TODO: implement me
	// g.shape.rotate()
}

func (g *Game) MoveLeft() {
	// TODO: implement me
	// g.shape.moveLeft()
}

func (g *Game) MoveRight() {
	// TODO: implement me
	// g.shape.moveRight()
}

func (g *Game) MoveDown() {
	// TODO: implement me
	// g.shape.moveDown()
}

func (g *Game) spawnShape() {
	index := rand.Intn(7)
	g.shape = shapeFuncs[index](coords{x: 5 * 48, y: -48 + 2})
}

func initGame() *Game {
	frame_clr := color.RGBA{184, 184, 184, 0xff}
	g := Game{
		interval: 6,
		frame: [4]*rectangle{
			makeRectangle(coords{482, 2}, frame_clr, coords{0, 0}),
			makeRectangle(coords{2, 962}, frame_clr, coords{482, 0}),
			makeRectangle(coords{482, 2}, frame_clr, coords{2, 962}),
			makeRectangle(coords{2, 962}, frame_clr, coords{0, 2}),
		},
	}
	g.spawnShape()
	return &g
}

func DrawRectangle(screen *ebiten.Image, rect *rectangle) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.x), float64(rect.y))
	screen.DrawImage(rect.img, op)
}

func (g *Game) DrawGameFrame(screen *ebiten.Image) {
	for _, rect := range g.frame {
		DrawRectangle(screen, rect)
	}
}

func (g *Game) ShapeHasBottomContact() bool {
	// Find lowest y value of shape
	lowestShapeY := g.shape.getLowestY()
	// See if bottom frame is directly below lowest y value
	if lowestShapeY == g.frame[2].y {
		return true
	}
	// Search squares to see if any are directly below lowest y value
	for _, sqr := range g.squares {
		top := sqr[0]
		if top.y == lowestShapeY {
			return true
		}
	}
	return false
}

func (g *Game) TransferShapeToSquares() {
	for _, sqr := range *g.shape {
		g.squares = append(g.squares, sqr)
	}
}
