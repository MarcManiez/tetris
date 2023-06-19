package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

//go:embed litha.mp3
var music []byte

func InitMusic() *audio.Player {
	audioContext := audio.NewContext(44100)
	wavPlayer, err := mp3.DecodeWithSampleRate(44100, bytes.NewReader(music))
	if err != nil {
		log.Fatal(err)
	}

	musicPlayer, err := audioContext.NewPlayer(wavPlayer)
	if err != nil {
		log.Fatal(err)
	}

	musicPlayer.Play()
	if err != nil {
		log.Fatal(err)
	}
	return musicPlayer
}
