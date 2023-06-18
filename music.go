package main

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func InitMusic() *audio.Player {
	audioContext := audio.NewContext(44100)
	v, err := os.ReadFile("litha.wav")
	if err != nil {
		log.Fatal(err)
	}
	wavPlayer, err := wav.DecodeWithSampleRate(44100, bytes.NewReader(v))
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
