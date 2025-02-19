package gameloop

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func ResetGameStruct(g *Game) {
	g.PlayerX = Configuration.ScreenWidth / 2
	g.PlayerY = Configuration.ScreenHeight - 50
	fmt.Println(Configuration.ScreenWidth, Configuration.ScreenHeight)
	g.gameOver = false
	g.shootTimer = 0
	g.shootRate = 60
	g.score = 0

	g.playerFrames = []*ebiten.Image{}
	g.currentFrame = 0

	g.playerTickCount = 0

	g.enemiesFrames = []*ebiten.Image{}
	g.enemies = []Enemie{}

	g.hitSoundPlayed = false

}
