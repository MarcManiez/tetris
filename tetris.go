package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func main() {
	initFonts()
	ebiten.SetFullscreen(true)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(initGame()); err != nil {
		log.Fatal(err)
	}
}

func getPressedKeys() []ebiten.Key {
	keys := []ebiten.Key{}
	keys = inpututil.AppendPressedKeys(keys)
	return keys
}

func getJustPressedKeys() []ebiten.Key {
	keys := []ebiten.Key{}
	keys = inpututil.AppendJustPressedKeys(keys)
	return keys
}
