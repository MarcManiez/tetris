package main

import "github.com/hajimehoshi/ebiten/v2"

// shapeImpl is a collection of squares representing a single tetris piece
type shapeImpl []*square

type shape interface {
	Rotate()
	Draw(screen *ebiten.Image)
	getBottomSquares() []*square
	getLeftSquares() []*square
	getRightSquares() []*square
	squares() []*square
}

func (s *shapeImpl) Draw(screen *ebiten.Image) {
	for _, sqr := range *s {
		sqr.Draw(screen)
	}
}

func (s *shapeImpl) Rotate() {
	var middleSquare *square
	for i, s := range s.squares() {
		if i == 0 {
			// set aside the first square, which will be the center of rotation
			middleSquare = s
		} else {
			// Adaptation of the classic square matrix rotation problem solution
			// (x, y) => [y, matrix.length - 1 - x]
			// This variable compensates for the fact that the rotation does not
			// happen within a matrix starting at 0,0
			recalibration := middleSquare.position.x - middleSquare.position.y
			// Assumes a 3x3 matrix, and this value is the highest x value in the matrix
			maxX := middleSquare.position.x + 1
			originalY := s.position.y
			s.position.y = s.position.x - recalibration
			s.position.x = middleSquare.position.y - originalY - 1 + maxX
		}
	}
}

func (s *shapeImpl) squares() []*square {
	return *s
}

// getBottomSquares returns the squares that are at the bottom of the shape
func (s *shapeImpl) getBottomSquares() []*square {
	return filter(*s, func(sqr *square) bool {
		return !some(*s, func(sqr2 *square) bool {
			if sqr == sqr2 {
				return false
			}
			return sqr2.position.y == sqr.position.y+1 && sqr2.position.x == sqr.position.x
		})
	})
}

// getLeftSquares returns the squares that are at the left of the shape
func (s *shapeImpl) getLeftSquares() []*square {
	return filter(*s, func(sqr *square) bool {
		return !some(*s, func(sqr2 *square) bool {
			if sqr == sqr2 {
				return false
			}
			return sqr2.position.x == sqr.position.x-1 && sqr2.position.y == sqr.position.y
		})
	})
}

// getRightSquares returns the squares that are at the right of the shape
func (s *shapeImpl) getRightSquares() []*square {
	return filter(*s, func(sqr *square) bool {
		return !some(*s, func(sqr2 *square) bool {
			if sqr == sqr2 {
				return false
			}
			return sqr2.position.x == sqr.position.x+1 && sqr2.position.y == sqr.position.y
		})
	})
}
