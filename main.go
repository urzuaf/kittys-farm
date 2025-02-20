package main

import (
	"embed"
	"game/gameloop"
	"game/load"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

//go:embed assets/**/*
var assets embed.FS

func main() {
	game := &gameloop.Game{}
	//Cargamos todas las imagenes una vez
	game.FileSystem = &assets
	game.BackgroundImage = load.LoadFromImage("assets/map/map.png", assets)
	game.MenuImage = load.LoadFromImage("assets/map/titlePressed.png", assets)
	game.TryAgainImage = load.LoadFromImage("assets/map/tryAgainPressed.png", assets)
	game.PlayerSpriteSheet = load.LoadFromImage("assets/player/mc.png", assets)
	game.EnemiesSpriteSheet = load.LoadFromImage("assets/enemies/chickens.png", assets)
	//Cargamos la musica
	// Usamos el contexto de Ebiten
	audioContext := audio.CurrentContext()
	game.MenuPlayer = load.LoadMusic(audioContext, assets, "assets/music/Menu.ogg")
	game.GamePlayer = load.LoadMusic(audioContext, assets, "assets/music/Game.ogg")
	game.HitPlayer = load.LoadMusic(audioContext, assets, "assets/music/sfx/Hit.ogg")
	//Cargamos la tipografia
	game.Font = load.LoadFont(assets, "assets/fonts/Jersey15-Regular.ttf")

	game.Reset()
	ebiten.SetWindowSize(int(math.Round(gameloop.Configuration.ScreenWidth*1.2)), int(math.Round(gameloop.Configuration.ScreenHeight*1.2)))
	ebiten.SetVsyncEnabled(true)
	game.InitialMenu = true
	ebiten.SetWindowTitle("Kitty's Farm")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
