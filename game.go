package main

import (
	"fmt"
	"math/rand"

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
		interval: 30,
		throttle: 10,
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
		g.MoveDown()
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

func (g *Game) spawnShape() {
	index := rand.Intn(7)
	g.shape = shapeFuncs[index](coords{x: 5, y: -1})
}

func (g *Game) ShapeHasBottomContact() bool {
	// Select lowest squares from a shape
	bottomSquares := g.shape.getBottomSquares()
	lowestY := bottomSquares[0].position.y
	if lowestY == (len(g.board.squares) - 1 - HIDDEN_AREA) {
		return true
	}

	// If every bottom square in the shape has a square directly below it, return true
	if some(bottomSquares, func(s *square) bool {
		return g.board.squares[s.position.y+1+HIDDEN_AREA][s.position.x] != nil
	}) {
		return true
	}
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

func (g *Game) Rotate() {
	// TODO: implement me
	// g.shape.rotate()
}

func (g *Game) MoveLeft() {
	// TODO: Prevent from moving left if there are already squares there
	if some(*g.shape, func(sqr *square) bool {
		return sqr.position.x == 0
	}) {
		return
	}
	for _, square := range *g.shape {
		square.position.x--
	}
}

func (g *Game) MoveRight() {
	// TODO: Prevent from moving right if there are already squares there
	if some(*g.shape, func(sqr *square) bool {
		return sqr.position.x == len(g.board.squares[0])-1
	}) {
		return
	}
	for _, square := range *g.shape {
		square.position.x++
	}
}

func (g *Game) MoveDown() {
	if g.ShapeHasBottomContact() {
		return
	}
	for _, sqr := range *g.shape {
		sqr.position.y++
	}
}
