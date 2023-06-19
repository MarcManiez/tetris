package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.DrawFrame(screen)
	g.DrawNextShape(screen)
	if g.paused {
		renderer := makeTextRenderer(64)
		renderer.SetTarget(screen)
		renderer.SetColor(makeChangingColor())
		renderer.Draw("PAUSED", 245, 475)
	} else {
		g.board.DrawShape(screen, g.currentShape)
		g.board.DrawSquares(screen)
	}
	if g.board.isGameOver() {
		ebitenutil.DebugPrintAt(screen, "Game over!", 500, 0)
		ebitenutil.DebugPrintAt(screen, "Press \"Enter\" to play again", 500, 20)
	}
}

// DrawNextShape renders the next shape module
func (g *Game) DrawNextShape(screen *ebiten.Image) {
	renderer := makeTextRenderer(24)
	renderer.SetTarget(screen)
	renderer.SetColor(color.White)
	renderer.Draw("Next Shape", 600, 475)
	g.nextShape.Draw(screen, coords{580, 650})
}