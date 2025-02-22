package gameloop

import (
	"game/utils"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetMcSprites(g *Game) {
	//Move around spritesheet
	var frameCount int = 4
	var spritewidth int = 14
	var spriteheight int = 16
	var paddingtop int = 16

	for i := 0; i < frameCount; i++ {
		x := i * spritewidth
		frame := g.PlayerSpriteSheet.SubImage(image.Rect(x, paddingtop, x+spritewidth, paddingtop+spriteheight)).(*ebiten.Image)
		g.playerFrames = append(g.playerFrames, frame)
	}

	//no queremos el 1r ni 2do frame del spritesheet
	utils.RemoveElement(&g.playerFrames, 1)
	//utils.RemoveElement(&g.playerFrames, 0)
	utils.Swap(&g.playerFrames, 0, 1)
	g.playerFrames = append(g.playerFrames, g.playerFrames[1])

}

func GetLosingMcSprites(g *Game) {
	//Move around spritesheet
	var frameCount int = 4
	var spritewidth int = 14
	var spriteheight int = 16
	var paddingtop int = 32

	for i := 0; i < frameCount; i++ {
		x := i * spritewidth
		frame := g.PlayerSpriteSheet.SubImage(image.Rect(x, paddingtop, x+spritewidth, paddingtop+spriteheight)).(*ebiten.Image)
		g.playerLosingFrames = append(g.playerLosingFrames, frame)
	}
	paddingtop = 48
	for i := 0; i < frameCount; i++ {
		x := i * spritewidth
		frame := g.PlayerSpriteSheet.SubImage(image.Rect(x, paddingtop, x+spritewidth, paddingtop+spriteheight)).(*ebiten.Image)
		g.playerLosingFrames = append(g.playerLosingFrames, frame)
	}

}
