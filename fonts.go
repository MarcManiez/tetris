package main

import (
	"embed"
	"image/color"
	"log"
	"time"

	"github.com/tinne26/etxt"
)

//go:embed fonts/*
var fonts embed.FS

func initTextRenderer() *etxt.Renderer {
	// load font library
	fontLib := etxt.NewFontLibrary()
	_, _, err := fontLib.ParseEmbedDirFonts("fonts", fonts)
	if err != nil {
		log.Fatalf("Error while loading fonts: %s", err.Error())
	}

	expectedFonts := []string{"JetBrains Mono Bold"} // !!
	for _, fontName := range expectedFonts {
		if !fontLib.HasFont(fontName) {
			log.Fatal("missing font: " + fontName)
		}
	}

	txtRenderer := etxt.NewStdRenderer()
	glyphsCache := etxt.NewDefaultCache(10 * 1024 * 1024) // 10MB
	txtRenderer.SetCacheHandler(glyphsCache.NewHandler())
	txtRenderer.SetFont(fontLib.GetFont(expectedFonts[0]))
	txtRenderer.SetAlign(etxt.YCenter, etxt.XCenter)
	txtRenderer.SetSizePx(64)
	return txtRenderer
}

// makeChangingColor returns a function that returns a color that changes over time
func makeChangingColor() color.RGBA {
	millis := time.Now().UnixMilli()
	blue := (millis / 16) % 512
	if blue >= 256 {
		blue = 511 - blue
	}
	return color.RGBA{0, 255, uint8(blue), 255}
}
