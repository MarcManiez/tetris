package main

import "github.com/hajimehoshi/ebiten/v2"

// shape is a collection of squares representing a single tetris piece
type shape []*square

func (s *shape) Draw(screen *ebiten.Image) {
	for _, sqr := range *s {
		sqr.Draw(screen)
	}
}

// getBottomSquares returns the squares that are at the bottom of the shape
func (s *shape) getBottomSquares() []*square {
	return filter(*s, func(sqr *square) bool {
		return !some(*s, func(sqr2 *square) bool {
			if sqr == sqr2 {
				return false
			}
			return sqr2.position.y == sqr.position.y+1 && sqr2.position.x == sqr.position.x
		})
	})
}
