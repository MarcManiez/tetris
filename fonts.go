package main

import (
	"log"

	"github.com/tinne26/etxt"
)

func initTextRenderer() *etxt.Renderer {
	// load font library
	fontLib := etxt.NewFontLibrary()
	_, _, err := fontLib.ParseDirFonts("fonts") // !!
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
