package gameloop

import (
	"fmt"
	"game/utils"
	"image"
	_ "image/png"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func EnemiesMovement(g *Game) {
	for i := 0; i < len(g.enemies); i++ {
		g.enemies[i].y += Configuration.EnemieSpeed
		g.enemies[i].tickCount++
		AnimateEnemies(g, i)

		if g.enemies[i].y > Configuration.ScreenHeight {
			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
			i--
		}
	}
}

func EnemiesShooting(g *Game) {
	g.shootTimer++
	if g.shootTimer >= g.shootRate {
		g.shootTimer = 0
		g.enemies = append(g.enemies, Enemie{x: rand.Float64() * (Configuration.ScreenWidth - 30), y: 0})
		if g.shootRate > 10 {
			g.shootRate-- //aumentamos el shootRate con el tiempo
		}
	}
}

func IncreaseEnemiesMovementSpeed(g *Game) {
	if g.score%600 == 0 && Configuration.EnemieSpeed < 15 {
		Configuration.EnemieSpeed++
		fmt.Println("Increased enemies movement speed to: ", Configuration.EnemieSpeed)
	}
}

func GetEnemiesSprites(g *Game) {
	var frameCount int = 4
	var spritewidth int = 12
	var spriteheight int = 12
	var paddingtop int = 12

	for i := 0; i < frameCount; i++ {
		x := i * spritewidth
		frame := g.enemiesSpriteSheet.SubImage(image.Rect(x, paddingtop, x+spritewidth, paddingtop+spriteheight)).(*ebiten.Image)
		g.enemiesFrames = append(g.enemiesFrames, frame)
	}
	utils.RemoveElement(&g.enemiesFrames, 0)
	utils.RemoveElement(&g.enemiesFrames, 1)

}
