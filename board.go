package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const HIDDEN_AREA = 4

type Board struct {
	squares [24][10]*square
	frame   [4]*rectangle
}

func (b *Board) AddSquare(s *square) {
	b.squares[s.position.y+HIDDEN_AREA][s.position.x] = s
}

func (b *Board) Draw(screen *ebiten.Image) {
	for _, row := range b.squares {
		for _, sq := range row {
			if sq != nil {
				sq.Draw(screen)
			}
		}
	}
	for _, rect := range b.frame {
		rect.DrawRectangle(screen)
	}
}

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
	return &Board{
		frame: [4]*rectangle{
			makeRectangle(coords{482, FRAME_WIDTH}, frame_color, coords{0, 0}),
			makeRectangle(coords{FRAME_WIDTH, 962}, frame_color, coords{482, 0}),
			makeRectangle(coords{482, FRAME_WIDTH}, frame_color, coords{FRAME_WIDTH, 962}),
			makeRectangle(coords{FRAME_WIDTH, 962}, frame_color, coords{0, FRAME_WIDTH}),
		},
	}
}
