package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game implements ebiten.Game interface and is composed of one active shape that the player controls and needs to place, and a collection of squares that have already been placed
type Game struct {
	shape                  *shape
	squares                []*square
	frame                  [4]*rectangle
	lastMove               int
	throttle               int
	updates                int
	interval               int
	updates_since_movement int
	shapes                 shapes
}

func initGame() *Game {
	frame_clr := color.RGBA{184, 184, 184, 0xff}
	g := Game{
		interval: 6,
		throttle: 20,
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

func (g *Game) Update() error {
	g.updates++
	g.updates_since_movement++
	g.HandleInput()
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
			rect.DrawRectangle(screen)
		}
	}
	// Draw other squares
	for _, sqr := range g.squares {
		for _, rect := range *sqr {
			rect.DrawRectangle(screen)
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
	for _, square := range *g.shape {
		for _, rect := range *square {
			rect.x += 48
		}
	}
}

func (g *Game) MoveDown() {
	// TODO: implement me
	// g.shape.moveDown()
}

func (g *Game) spawnShape() {
	index := rand.Intn(7)
	g.shape = shapeFuncs[index](coords{x: 5 * 48, y: -48 + 2})
}

func (g *Game) DrawGameFrame(screen *ebiten.Image) {
	for _, rect := range g.frame {
		rect.DrawRectangle(screen)
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
		// TODO: there's a bug here: the shape stops even when there's a gap, because it uses the top-most square on the entire game surface. Instead, it should collect all the bottom squares of the shape and see if any of them are touching the top of any of the squares in the squares collection
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

func (g *Game) HandleInput() {
	keys := []ebiten.Key{}
	keys = inpututil.AppendPressedKeys(keys)
	fmt.Println(keys)
	// Do nothing if no keys are pressed or if more than one key is pressed,
	// or if trying to move before the throttle period has passed
	if len(keys) != 1 || g.updates-g.lastMove < g.throttle {
		return
	}
	switch key := keys[0]; key {
	case ebiten.KeyArrowUp:
		g.Rotate()
	case ebiten.KeyArrowRight:
		g.MoveRight()
	case ebiten.KeyArrowDown:
		g.MoveDown()
	case ebiten.KeyArrowLeft:
		g.MoveLeft()
	}
	g.lastMove = g.updates
}
