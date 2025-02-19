package main

import (
	"game/gameloop"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &gameloop.Game{}
	game.Reset()
	ebiten.SetWindowSize(int(math.Round(gameloop.Configuration.ScreenWidth*1.2)), int(math.Round(gameloop.Configuration.ScreenHeight*1.2)))
	ebiten.SetVsyncEnabled(true)
	game.InitialMenu = true
	ebiten.SetWindowTitle("Kitty's Farm")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
