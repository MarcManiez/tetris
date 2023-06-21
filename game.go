package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var SCORE_MATRIX = map[int]int{
	1: 1,
	2: 3,
	3: 5,
	4: 8,
}

// Game implements ebiten.Game interface and is composed of one active shape that the player controls and needs to place, and a collection of squares that have already been placed
type Game struct {
	currentShape shape
	nextShape    shape
	board        *Board
	linesCleared int
	score        int
	// Last update call for which a movement key was pressed
	lastMove int
	// Throttle value for player movement
	throttle int
	// Number of times the update function has run
	updates                int
	updates_since_movement int
	musicPlayer            *audio.Player
	paused                 bool
}

const INTERVAL = 60
const INITIAL_LEVEL = 1

func initGame() *Game {
	g := Game{
		throttle: 10,
		board:    makeBoard(),
		// Un/comment for music on/off
		// musicPlayer:  InitMusic(),
	}
	g.currentShape = makeRandomShape(coords{x: 5, y: -1})
	g.spawnNextShape()
	return &g
}

// spawnNextShape spawns the next shape
func (g *Game) spawnNextShape() {
	g.nextShape = makeRandomShape(coords{x: 0, y: 0})
}

func (g *Game) Update() error {
	if g.paused {
		if includes(getJustPressedKeys(), ebiten.KeyP) {
			g.unpause()
		} else {
			return nil
		}
	}
	if g.board.isGameOver() {
		if includes(getPressedKeys(), ebiten.KeyEnter) {
			g.restart()
		}
	} else {
		g.updates++
		g.updates_since_movement++
		g.HandleInput()
		if g.updates_since_movement == g.interval() {
			linesCleared := g.board.clearFullLines()
			g.linesCleared += linesCleared
			g.score += SCORE_MATRIX[linesCleared] * g.level()
		}
		if g.updates_since_movement >= g.interval() {
			g.updates_since_movement = 0
			g.MoveDown()
			if !g.CanMoveDown() {
				g.cycleShape()
			}
		}
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 1200
}

// cycleShape transfers the current shape to the board, makes the next shape the
// current shape, and spawns a new next shape
func (g *Game) cycleShape() {
	g.board.addShape(g.currentShape)
	g.currentShape = g.nextShape
	g.currentShape.translate(coords{x: 5, y: -1})
	g.spawnNextShape()
}

// HandleInput handles input from the player during active game
func (g *Game) HandleInput() {
	keys := getPressedKeys()
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
	case ebiten.KeyP:
		g.pause()
	}
	g.lastMove = g.updates
}

// ****************
// MOVEMENT METHODS
// ****************

func (g *Game) Rotate() {
	originalPositions := g.currentShape.copyPositions()
	// try rotating up to three times
	for i := 0; i < 3; i++ {
		g.currentShape.Rotate()
		if g.isShapePositionValid() {
			return
		}
	}
	// Restore initial position if no rotations were valid
	g.currentShape.setCoordinates(originalPositions)
}

func (g *Game) MoveLeft() {
	if g.CanMoveLeft() {
		g.currentShape.MoveLeft()
	}
}

func (g *Game) MoveRight() {
	if g.CanMoveRight() {
		g.currentShape.MoveRight()
	}
}

func (g *Game) MoveDown() {
	if g.CanMoveDown() {
		g.currentShape.MoveDown()
	}
}

// CanMoveDown returns true if the shape can move down
func (g *Game) CanMoveDown() bool {
	bottomSquares := g.currentShape.getBottomSquares()
	return none(bottomSquares, func(s *square) bool {
		return !g.board.isCoordValid(coords{x: s.position.x, y: s.position.y + 1})
	})
}

// CanMoveLeft returns true if the shape can move left
func (g *Game) CanMoveLeft() bool {
	// Select leftmost squares from a shape
	leftSquares := g.currentShape.getLeftSquares()
	return none(leftSquares, func(s *square) bool {
		return !g.board.isCoordValid(coords{x: s.position.x - 1, y: s.position.y})
	})
}

// CanMoveRight returns true if the shape can move right
func (g *Game) CanMoveRight() bool {
	// Select rightmost squares from a shape
	rightSquares := g.currentShape.getRightSquares()
	return none(rightSquares, func(s *square) bool {
		return !g.board.isCoordValid(coords{x: s.position.x + 1, y: s.position.y})
	})
}

// isShapePositionValid returns true if the shape is not colliding with any squares
// or the edge of the board
func (g *Game) isShapePositionValid() bool {
	for _, square := range g.currentShape.squares() {
		if !g.board.isCoordValid(square.position) {
			return false
		}
	}
	return true
}

// ******************
// GAME STATE METHODS
// ******************

// restart resets the game after a "game over"
func (g *Game) restart() {
	g.board = makeBoard()
	g.linesCleared = 0
	g.updates_since_movement = 0
	g.score = 0
	g.currentShape = makeRandomShape(coords{x: 5, y: -1})
	g.spawnNextShape()
}

func (g *Game) pause() {
	g.paused = true
	if g.musicPlayer != nil {
		g.musicPlayer.Pause()
	}
}

func (g *Game) unpause() {
	g.paused = false
	if g.musicPlayer != nil {
		g.musicPlayer.Play()
	}
}

func (g *Game) level() int {
	return g.linesCleared/10 + 1
}

// interval represents the number of updates between shape movements
func (g *Game) interval() int {
	level := g.level()
	interval := 0.8 - (float32(level-1) * float32(0.007))
	if level == 1 {
		return int(interval * 60)
	} else {
		return int(math.Pow(float64(interval), float64(level-2)) * float64(60))
	}
}
