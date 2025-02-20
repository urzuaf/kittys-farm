package gameloop

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	MenuState = iota
	GameplayState
)

var (
	audioContext      *audio.Context
	audioCurrentState int
)

func init() {
	audioContext = audio.NewContext(44100)
}

func SwitchMusic(g *Game, newState int) {
	if newState == GameplayState && audioCurrentState != GameplayState {
		fmt.Println("Playing gameplay music")
		g.MenuPlayer.Pause()
		g.GamePlayer.Rewind()
		g.GamePlayer.Play()
	} else if newState == MenuState && audioCurrentState != MenuState {
		fmt.Println("Playing menu music")
		g.GamePlayer.Pause()
		g.MenuPlayer.Rewind()
		g.MenuPlayer.Play()
	}
	audioCurrentState = newState
}
