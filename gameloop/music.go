package gameloop

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	MenuState = iota
	GameplayState
)

var (
	audioContext      *audio.Context
	menuPlayer        *audio.Player
	gameplayPlayer    *audio.Player
	shootPlayer       *audio.Player
	audioCurrentState int
)

func init() {
	audioContext = audio.NewContext(44100)
	menuPlayer = loadMusic("assets/music/Menu.wav")
	gameplayPlayer = loadMusic("assets/music/Game.wav")
	loadShootSound("assets/music/sfx/Hit.wav")
}

func loadMusic(filepath string) *audio.Player {
	musicFile, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	musicStream, err := wav.DecodeWithSampleRate(44100, bytes.NewReader(musicFile))
	if err != nil {
		panic(err)
	}

	// FOR LOOP MUSIC
	loopStream := audio.NewInfiniteLoop(musicStream, musicStream.Length())

	player, err := audio.NewPlayer(audioContext, loopStream)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.5) // Ajustar volumen
	return player
}

func loadShootSound(filepath string) {
	shootFile, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	shootStream, err := wav.DecodeWithSampleRate(44100, bytes.NewReader(shootFile))
	if err != nil {
		panic(err)
	}
	shootPlayer, err = audio.NewPlayer(audioContext, shootStream)
	if err != nil {
		panic(err)
	}
	shootPlayer.SetVolume(0.7) // Ajustar volumen del efecto
}

func HitSound() {
	shootFile, err := os.ReadFile("assets/music/sfx/Hit.wav")
	if err != nil {
		panic(err)
	}
	shootStream, err := wav.DecodeWithSampleRate(44100, bytes.NewReader(shootFile))
	if err != nil {
		panic(err)
	}
	newPlayer, err := audio.NewPlayer(audioContext, shootStream)
	if err != nil {
		panic(err)
	}
	newPlayer.SetVolume(0.7)
	newPlayer.Play()
}

func SwitchMusic(newState int) {
	if newState == GameplayState && audioCurrentState != GameplayState {
		fmt.Println("Playing gameplay music")
		menuPlayer.Pause()
		gameplayPlayer.Rewind()
		gameplayPlayer.Play()
	} else if newState == MenuState && audioCurrentState != MenuState {
		fmt.Println("Playing menu music")
		gameplayPlayer.Pause()
		menuPlayer.Rewind()
		menuPlayer.Play()
	}
	audioCurrentState = newState
}
