package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tinne26/etxt"
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
	paused                 bool
	textRenderer           *etxt.Renderer
}

func initGame() *Game {
	g := Game{
		interval: 60,
		throttle: 10,
		board:    makeBoard(),
		// Un/comment for music on/off
		// musicPlayer:  InitMusic(),
		textRenderer: initTextRenderer(),
	}
	g.spawnShape()
	return &g
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
		g.board.clearFullLines()
		if g.updates_since_movement >= g.interval {
			g.updates_since_movement = 0
			g.MoveDown()
			if !g.CanMoveDown() {
				g.board.addShape(g.shape)
				g.spawnShape()
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.DrawFrame(screen)
	if g.paused {
		g.textRenderer.SetTarget(screen)
		millis := time.Now().UnixMilli()
		blue := (millis / 16) % 512
		if blue >= 256 {
			blue = 511 - blue
		}
		changingColor := color.RGBA{0, 255, uint8(blue), 255}
		g.textRenderer.SetColor(changingColor)
		g.textRenderer.Draw("PAUSED", 245, 475)
	} else {
		g.shape.Draw(screen)
		g.board.DrawSquares(screen)
	}
	if g.board.isGameOver() {
		ebitenutil.DebugPrintAt(screen, "Game over!", 500, 0)
		ebitenutil.DebugPrintAt(screen, "Press \"Enter\" to play again", 500, 20)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 1200
}

func (g *Game) spawnShape() {
	index := rand.Intn(len(shapeFuncs))
	g.shape = shapeFuncs[index](coords{x: 5, y: -1})
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
	originalPositions := g.shape.copyPositions()
	// try rotating up to three times
	for i := 0; i < 3; i++ {
		g.shape.Rotate()
		if g.isShapePositionValid() {
			return
		}
	}
	// Restore initial position if no rotations were valid
	g.shape.translate(originalPositions)
}

func (g *Game) MoveLeft() {
	if g.CanMoveLeft() {
		g.shape.MoveLeft()
	}
}

func (g *Game) MoveRight() {
	if g.CanMoveRight() {
		g.shape.MoveRight()
	}
}

func (g *Game) MoveDown() {
	if g.CanMoveDown() {
		g.shape.MoveDown()
	}
}

// CanMoveDown returns true if the shape can move down
func (g *Game) CanMoveDown() bool {
	bottomSquares := g.shape.getBottomSquares()
	return none(bottomSquares, func(s *square) bool {
		return !g.board.isCoordValid(coords{x: s.position.x, y: s.position.y + 1})
	})
}

// CanMoveLeft returns true if the shape can move left
func (g *Game) CanMoveLeft() bool {
	// Select leftmost squares from a shape
	leftSquares := g.shape.getLeftSquares()
	return none(leftSquares, func(s *square) bool {
		return !g.board.isCoordValid(coords{x: s.position.x - 1, y: s.position.y})
	})
}

// CanMoveRight returns true if the shape can move right
func (g *Game) CanMoveRight() bool {
	// Select rightmost squares from a shape
	rightSquares := g.shape.getRightSquares()
	return none(rightSquares, func(s *square) bool {
		return !g.board.isCoordValid(coords{x: s.position.x + 1, y: s.position.y})
	})
}

// isShapePositionValid returns true if the shape is not colliding with any squares
// or the edge of the board
func (g *Game) isShapePositionValid() bool {
	for _, square := range g.shape.squares() {
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
	g.updates_since_movement = 0
	g.spawnShape()
}

func (g *Game) pause() {
	g.paused = true
	g.musicPlayer.Pause()
}

func (g *Game) unpause() {
	g.paused = false
	g.musicPlayer.Play()
}
