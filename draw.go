package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.DrawFrame(screen)
	g.DrawNextShape(screen)
	g.DrawScore(screen)
	g.DrawLevel(screen)
	if g.paused {
		renderer := makeTextRenderer(textOptions{fontSize: 64})
		renderer.SetTarget(screen)
		renderer.SetColor(makeChangingColor())
		renderer.Draw("PAUSED", 245, 475)
	} else {
		g.board.DrawShape(screen, g.currentShape)
		g.board.DrawSquares(screen, g.updates, g.interval())
	}
	if g.board.isGameOver() {
		ebitenutil.DebugPrintAt(screen, "Game over!", 500, 0)
		ebitenutil.DebugPrintAt(screen, "Press \"Enter\" to play again", 500, 20)
	}
}

// DrawNextShape renders the next shape module
func (g *Game) DrawNextShape(screen *ebiten.Image) {
	renderer := makeTextRenderer(textOptions{fontSize: 24})
	renderer.SetTarget(screen)
	renderer.SetColor(color.White)
	renderer.Draw("Next Shape", 600, 475)
	g.nextShape.Draw(screen, coords{580, 650}, 255)
}

func (g *Game) DrawLevel(screen *ebiten.Image) {
	renderer := makeTextRenderer(textOptions{fontSize: 24})
	renderer.SetTarget(screen)
	renderer.SetColor(color.White)
	level := fmt.Sprintf("Level %d", g.level())
	renderer.Draw(level, 579, 400)
}

func (g *Game) DrawScore(screen *ebiten.Image) {
	renderer := makeTextRenderer(textOptions{fontSize: 24})
	renderer.SetTarget(screen)
	renderer.SetColor(color.White)
	score := fmt.Sprintf("Score %d", g.score)
	renderer.Draw(score, 579, 425)
}
