package main

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

func InitMusic() *audio.Player {
	audioContext := audio.NewContext(44100)
	v, err := os.ReadFile("litha.mp3")
	if err != nil {
		log.Fatal(err)
	}
	wavPlayer, err := mp3.DecodeWithSampleRate(44100, bytes.NewReader(v))
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
