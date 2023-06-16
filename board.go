package main

import "github.com/hajimehoshi/ebiten/v2"

type Board [24][10]*square

func (b *Board) AddSquare(s *square) {
	b[s.position.y+4][s.position.x] = s
}

func (b *Board) Draw(screen *ebiten.Image) {
	for _, row := range b {
		for _, sq := range row {
			if sq != nil {
				sq.Draw(screen)
			}
		}
	}
}
