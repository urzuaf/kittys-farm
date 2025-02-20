package gameloop

import (
	"fmt"
	"game/patterns"
	"game/utils"
	"image"
	_ "image/png"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func EnemiesMovement(g *Game) {
	for i := 0; i < len(g.enemies); i++ {

		//Enemy movement
		g.enemies[i].x, g.enemies[i].y = patterns.GetPattern(g.enemiePattern, g.enemies[i].x, g.enemies[i].y, Configuration.EnemieXSpeed, Configuration.EnemieYSpeed, Configuration.ScreenWidth, Configuration.ScreenHeight, i, g.score, g.enemies[i].tickCount)
		//Enemy Animation
		g.enemies[i].tickCount++
		AnimateEnemies(g, i)

		//Delete enemies offscreen
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

		//for straight patterns we use a random position
		if utils.IsStraight(g.enemiePattern) {
			g.enemies = append(g.enemies, Enemie{x: rand.Float64() * (Configuration.ScreenWidth - 32), y: 0})
		} else {
			g.enemies = append(g.enemies, Enemie{x: Configuration.ScreenWidth/2 + (rand.Float64() * 100) - 80, y: 0})
		}

		if g.shootRate > 15 {
			g.shootRate-- //aumentamos el shootRate con el tiempo
		}
	}
}

func IncreaseEnemiesMovementSpeed(g *Game) {
	if g.score%180 == 0 && Configuration.EnemieYSpeed < 15 {
		Configuration.EnemieYSpeed += 0.1
		fmt.Printf("Increased enemies movement speed to: %.2f\n", Configuration.EnemieYSpeed)
	}
}

func GetEnemiesSprites(g *Game) {
	var frameCount int = 4
	var spritewidth int = 12
	var spriteheight int = 12
	var paddingtop int = 12

	for i := 0; i < frameCount; i++ {
		x := i * spritewidth
		frame := g.EnemiesSpriteSheet.SubImage(image.Rect(x, paddingtop, x+spritewidth, paddingtop+spriteheight)).(*ebiten.Image)
		g.enemiesFrames = append(g.enemiesFrames, frame)
	}
	utils.RemoveElement(&g.enemiesFrames, 0)
	utils.RemoveElement(&g.enemiesFrames, 1)

}
