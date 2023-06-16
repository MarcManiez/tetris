package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const FRAME_WIDTH = 2

// Game implements ebiten.Game interface and is composed of one active shape that the player controls and needs to place, and a collection of squares that have already been placed
type Game struct {
	shape *shape
	board *Board
	// Last update call for which a movement key was pressed
	lastMove int
	// Throttle value for player movement
	throttle int
	// Number of times the update function has run
	updates int
	// Number of updates between shape movements
	interval               int
	updates_since_movement int
}

func initGame() *Game {
	g := Game{
		interval: 6,
		throttle: 20,
		board:    makeBoard(),
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
		g.shape.moveDown()
		if g.ShapeHasBottomContact() {
			g.TransferShapeToSquares()
			g.spawnShape()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.shape.Draw(screen)
	g.board.Draw(screen)
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
		square.position.x++
	}
}

func (g *Game) MoveDown() {
	// TODO: implement me
	// g.shape.moveDown()
}

func (g *Game) spawnShape() {
	// index := rand.Intn(7)
	g.shape = shapeFuncs[6](coords{x: 5, y: -1})
}

func (g *Game) ShapeHasBottomContact() bool {
	// // Find lowest y value of shape
	// lowestShapeY := g.shape.getLowestY()
	// // See if bottom frame is directly below lowest y value
	// if lowestShapeY == g.frame[2].y {
	// 	return true
	// }
	// // Search squares to see if any are directly below lowest y value
	// for _, sqr := range g.squares {
	// 	top := sqr[0]
	// 	// TODO: there's a bug here: the shape stops even when there's a gap, because it uses the top-most square on the entire game surface. Instead, it should collect all the bottom squares of the shape and see if any of them are touching the top of any of the squares in the squares collection
	// 	if top.y == lowestShapeY {
	// 		return true
	// 	}
	// }
	return false
}

func (g *Game) TransferShapeToSquares() {
	for _, sqr := range *g.shape {
		g.board.AddSquare(sqr)
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
