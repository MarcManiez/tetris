package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// square is 5 rectangles: 4 for the border, with another one for the middle
type square struct {
	position coords
	color    color.Color
}

const SQUARE_HEIGHT = 48
const SQUARE_BORDER_WIDTH = 2

// Draw draws a square on the screen at the provided origin coordinates
func (s *square) Draw(screen *ebiten.Image, origin coords) {
	insideWidth := SQUARE_HEIGHT - SQUARE_BORDER_WIDTH*2
	arr := []*rectangle{
		// Top
		makeRectangle(
			coords{(SQUARE_HEIGHT - SQUARE_BORDER_WIDTH), SQUARE_BORDER_WIDTH},
			color.RGBA{184, 184, 184, 0xff},
			coords{s.position.x*SQUARE_HEIGHT + origin.x, s.position.y*SQUARE_HEIGHT + origin.y},
		),
		// Bottom
		makeRectangle(
			coords{(SQUARE_HEIGHT - SQUARE_BORDER_WIDTH), SQUARE_BORDER_WIDTH},
			color.RGBA{136, 136, 136, 0xff},
			coords{
				s.position.x*SQUARE_HEIGHT + SQUARE_BORDER_WIDTH + origin.x,
				s.position.y*SQUARE_HEIGHT + (SQUARE_HEIGHT - SQUARE_BORDER_WIDTH) + origin.y,
			},
		),
		// Right
		makeRectangle(
			coords{SQUARE_BORDER_WIDTH, (SQUARE_HEIGHT - SQUARE_BORDER_WIDTH)},
			color.RGBA{200, 200, 200, 0xff},
			coords{s.position.x*SQUARE_HEIGHT + (SQUARE_HEIGHT - SQUARE_BORDER_WIDTH) + origin.x, s.position.y*SQUARE_HEIGHT + origin.y},
		),
		// Left
		makeRectangle(
			coords{SQUARE_BORDER_WIDTH, (SQUARE_HEIGHT - SQUARE_BORDER_WIDTH)},
			color.RGBA{150, 150, 150, 0xff},
			coords{s.position.x*SQUARE_HEIGHT + origin.x, s.position.y*SQUARE_HEIGHT + SQUARE_BORDER_WIDTH + origin.y},
		),
		// Middle
		makeRectangle(
			coords{insideWidth, insideWidth},
			s.color,
			coords{s.position.x*SQUARE_HEIGHT + SQUARE_BORDER_WIDTH + origin.x, s.position.y*SQUARE_HEIGHT + SQUARE_BORDER_WIDTH + origin.y},
		),
	}
	for _, rect := range arr {
		rect.DrawRectangle(screen)
	}
}
