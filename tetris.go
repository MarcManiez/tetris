package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(initGame()); err != nil {
		log.Fatal(err)
	}
}
