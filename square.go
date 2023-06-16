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

func (s *square) Draw(screen *ebiten.Image) {
	insideWidth := SQUARE_HEIGHT - SQUARE_BORDER_WIDTH*2
	arr := []*rectangle{
		// Top
		makeRectangle(
			coords{(SQUARE_HEIGHT - SQUARE_BORDER_WIDTH), SQUARE_BORDER_WIDTH},
			color.RGBA{184, 184, 184, 0xff},
			coords{s.position.x * SQUARE_HEIGHT, s.position.y*SQUARE_HEIGHT + FRAME_WIDTH},
		),
		// Bottom
		makeRectangle(
			coords{(SQUARE_HEIGHT - SQUARE_BORDER_WIDTH), SQUARE_BORDER_WIDTH},
			color.RGBA{136, 136, 136, 0xff},
			coords{s.position.x*SQUARE_HEIGHT + SQUARE_BORDER_WIDTH, s.position.y*SQUARE_HEIGHT + (SQUARE_HEIGHT - SQUARE_BORDER_WIDTH) + FRAME_WIDTH},
		),
		// Right
		makeRectangle(
			coords{SQUARE_BORDER_WIDTH, (SQUARE_HEIGHT - SQUARE_BORDER_WIDTH)},
			color.RGBA{200, 200, 200, 0xff},
			coords{s.position.x*SQUARE_HEIGHT + (SQUARE_HEIGHT - SQUARE_BORDER_WIDTH), s.position.y*SQUARE_HEIGHT + FRAME_WIDTH},
		),
		// Left
		makeRectangle(
			coords{SQUARE_BORDER_WIDTH, (SQUARE_HEIGHT - SQUARE_BORDER_WIDTH)},
			color.RGBA{150, 150, 150, 0xff},
			coords{s.position.x * SQUARE_HEIGHT, s.position.y*SQUARE_HEIGHT + SQUARE_BORDER_WIDTH + FRAME_WIDTH},
		),
		// Middle
		makeRectangle(
			coords{insideWidth, insideWidth},
			s.color,
			coords{s.position.x*SQUARE_HEIGHT + SQUARE_BORDER_WIDTH, s.position.y*SQUARE_HEIGHT + SQUARE_BORDER_WIDTH + FRAME_WIDTH},
		),
	}
	for _, rect := range arr {
		rect.DrawRectangle(screen)
	}
}
