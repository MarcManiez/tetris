package main

import "github.com/hajimehoshi/ebiten/v2"

// shape is a collection of squares representing a single tetris piece
type shape []*square

func (s *shape) moveDown() {
	for _, sqr := range *s {
		sqr.position.y++
	}
}

func (s *shape) Draw(screen *ebiten.Image) {
	for _, sqr := range *s {
		sqr.Draw(screen)
	}
}

// getLowestY returns the lowest y value of the shape, and the x value of the rectangle with that y value
// func (s *shape) getLowestY() int {
// 	var lowestShapeY int
// 	for _, sqr := range *s {
// 		for _, rect := range *sqr {
// 			height := rect.img.Bounds().Dy()
// 			height += rect.y
// 			if height > lowestShapeY {
// 				lowestShapeY = height
// 			}
// 		}
// 	}
// 	return lowestShapeY
// }
