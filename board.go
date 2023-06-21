package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const HIDDEN_AREA = 3

type Board struct {
	squares    [23][10]*square
	frame      [4]*rectangle
	frameWidth int
}

func (b *Board) AddSquare(s *square) {
	b.squares[s.position.y+HIDDEN_AREA][s.position.x] = s
}

// *******
// DRAWING
// *******

// DrawFrame draws the frame around the board
func (b *Board) DrawFrame(screen *ebiten.Image) {
	for _, rect := range b.frame {
		rect.DrawRectangle(screen)
	}
}

// DrawSquares draws the squares on the board
func (b *Board) DrawSquares(screen *ebiten.Image, updates int, interval int) {
	for rowIndex, row := range b.squares {
		var opacity uint8 = 255
		// If the line is full, flash the line by making it transparent 4 times per interval
		if b.isLineFull(rowIndex) && (uint8(updates/(interval/4))%2) == 0 {
			opacity = 128
		}
		for _, sq := range row {
			if sq != nil {
				sq.Draw(screen, coords{b.frameWidth, b.frameWidth}, opacity)
			}
		}
	}
}

// DrawSquare draws a square on the board
func (b *Board) DrawShape(screen *ebiten.Image, shp shape) {
	shp.Draw(screen, coords{b.frameWidth, b.frameWidth}, 244)
}

// ***********
// END DRAWING
// ***********

// isLineFull returns true if the line is full
func (b *Board) isLineFull(y int) bool {
	return every(b.squares[y][:], func(sq *square) bool {
		return sq != nil
	})
}

func (b *Board) shiftLinesDown(y int) {
	for y > 0 {
		for x := 0; x < 10; x++ {
			b.squares[y][x] = b.squares[y-1][x]
			if b.squares[y][x] != nil {
				b.squares[y][x].position.y++
			}
		}
		y--
	}
}

func makeBoard() *Board {
	frame_color := color.RGBA{184, 184, 184, 0xff}
	frameWidth := 2
	return &Board{
		frameWidth: frameWidth,
		frame: [4]*rectangle{
			makeRectangle(coords{482, frameWidth}, frame_color, coords{0, 0}),
			makeRectangle(coords{frameWidth, 962}, frame_color, coords{482, 0}),
			makeRectangle(coords{482, frameWidth}, frame_color, coords{frameWidth, 962}),
			makeRectangle(coords{frameWidth, 962}, frame_color, coords{0, frameWidth}),
		},
	}
}

// isCoordOccupied returns true if the coordinate is occupied
func (b *Board) isCoordOccupied(c coords) bool {
	return b.squares[c.y+HIDDEN_AREA][c.x] != nil
}

// isCoordOutOfBounds returns true if the coordinate is out of bounds
func (b *Board) isCoordOutOfBounds(c coords) bool {
	return c.x < 0 || c.x >= len(b.squares[0]) || c.y+HIDDEN_AREA >= len(b.squares) || c.y+HIDDEN_AREA < 0
}

// isCoordValid returns true if the coordinate is valid
func (b *Board) isCoordValid(c coords) bool {
	return !b.isCoordOutOfBounds(c) && !b.isCoordOccupied(c)
}

// isGameLost returns true if the game is lost
func (b *Board) isGameOver() bool {
	return some(b.squares[HIDDEN_AREA-1][:], func(sq *square) bool {
		return sq != nil
	})
}

// addShape adds a shape to the board
func (b *Board) addShape(s shape) {
	for _, sq := range s.squares() {
		b.AddSquare(sq)
	}
}

// clearFullLines clears all full lines
func (b *Board) clearFullLines() int {
	var linesCleared int
	for y := len(b.squares) - 1; y >= 0; y-- {
		for b.isLineFull(y) {
			linesCleared++
			b.shiftLinesDown(y)
		}
	}
	return linesCleared
}
