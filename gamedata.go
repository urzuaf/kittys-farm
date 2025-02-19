package main

import "github.com/hajimehoshi/ebiten/v2"

type GameData struct {
	screenWidth  int
	screenHeight int
	tileWidth    int
	tileHeight   int
}

func NewGameData() *GameData {
	return &GameData{
		screenWidth:  320,
		screenHeight: 240,
		tileWidth:    16,
		tileHeight:   16,
	}
}

type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

func GetIndexFromXY(x, y int) int {
	gd := NewGameData()
	return (y * gd.screenWidth) + x
}
