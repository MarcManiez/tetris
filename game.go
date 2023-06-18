package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const FRAME_WIDTH = 2

// Game implements ebiten.Game interface and is composed of one active shape that the player controls and needs to place, and a collection of squares that have already been placed
type Game struct {
	shape shape
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
	musicPlayer            *audio.Player
}

func initGame() *Game {
	g := Game{
		interval: 60,
		throttle: 10,
		board:    makeBoard(),
		// Uncomment for music
		// musicPlayer: InitMusic(),
	}
	g.spawnShape()
	return &g
}

func (g *Game) Update() error {
	g.updates++
	g.updates_since_movement++
	g.HandleInput()
	g.ClearFullLines()
	if g.updates_since_movement >= g.interval {
		g.updates_since_movement = 0
		g.MoveDown()
		if !g.CanMoveDown() {
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
	index := rand.Intn(len(shapeFuncs))
	g.shape = shapeFuncs[index](coords{x: 5, y: -1})
}

func (g *Game) CanMoveDown() bool {
	// Select lowest squares from a shape
	bottomSquares := g.shape.getBottomSquares()
	lowestY := bottomSquares[0].position.y
	if lowestY == (len(g.board.squares) - 1 - HIDDEN_AREA) {
		return false
	}

	// If every bottom square in the shape has a square directly below it, return true
	if some(bottomSquares, func(s *square) bool {
		return g.board.squares[s.position.y+1+HIDDEN_AREA][s.position.x] != nil
	}) {
		return false
	}
	return true
}

// CanMoveLeft returns true if the shape can move left
func (g *Game) CanMoveLeft() bool {
	// Select leftmost squares from a shape
	leftSquares := g.shape.getLeftSquares()
	if some(leftSquares, func(s *square) bool {
		return s.position.x == 0
	}) {
		return false
	}
	// If any left square in the shape has a square directly to its left, return true
	if some(leftSquares, func(s *square) bool {
		return g.board.squares[s.position.y+HIDDEN_AREA][s.position.x-1] != nil
	}) {
		return false
	}
	return true
}

// CanMoveRight returns true if the shape can move right
func (g *Game) CanMoveRight() bool {
	// Select rightmost squares from a shape
	rightSquares := g.shape.getRightSquares()
	if some(rightSquares, func(s *square) bool {
		return s.position.x == (len(g.board.squares[0]) - 1)
	}) {
		return false
	}
	// If any right square in the shape has a square directly to its right, return true
	if some(rightSquares, func(s *square) bool {
		return g.board.squares[s.position.y+HIDDEN_AREA][s.position.x+1] != nil
	}) {
		return false
	}
	return true
}

func (g *Game) TransferShapeToSquares() {
	for _, sqr := range g.shape.squares() {
		g.board.AddSquare(sqr)
	}
}

func (g *Game) HandleInput() {
	keys := []ebiten.Key{}
	keys = inpututil.AppendPressedKeys(keys)
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

func (g *Game) ClearFullLines() {
	for y := len(g.board.squares) - 1; y >= 0; y-- {
		for g.board.isLineFull(y) {
			g.board.shiftLinesDown(y)
		}
	}
}

func (g *Game) Rotate() {
	g.shape.Rotate()
}

func (g *Game) MoveLeft() {
	if !g.CanMoveLeft() {
		return
	}
	for _, square := range g.shape.squares() {
		square.position.x--
	}
}

func (g *Game) MoveRight() {
	if !g.CanMoveRight() {
		return
	}
	for _, square := range g.shape.squares() {
		square.position.x++
	}
}

func (g *Game) MoveDown() {
	if !g.CanMoveDown() {
		return
	}
	g.shape.MoveDown()
}
