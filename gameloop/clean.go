package gameloop

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func ResetGame(g *Game) {
	ResetGameStruct(g)
	ResetConfiguration(g)
}

func ResetGameStruct(g *Game) {
	g.InitialMenu = false
	g.Paused = false
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
	g.enemiePattern = rand.Intn(10)
	g.hitSoundPlayed = false

}

func ResetConfiguration(g *Game) {
	Configuration.PlayerSpeed = 4
	Configuration.EnemieXSpeed = 3
	Configuration.EnemieYSpeed = 2
}
