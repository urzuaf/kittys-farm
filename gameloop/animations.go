package gameloop

func AnimateCharacter(g *Game) {
	//animación del jugador
	g.playerTickCount++

	animationspeed := (5000 - g.playerTickCount) / 100
	if animationspeed < 4 {
		animationspeed = 4
	}
	if animationspeed > 30 {
		animationspeed = 30
	}

	if g.playerTickCount%animationspeed == 0 {
		g.currentFrame = (g.currentFrame + 1) % (len(g.playerFrames))
	}

}
func AnimateLosingCharacter(g *Game) {
	//animación del jugador
	g.playerTickCount++

	if g.playerTickCount%2 == 0 {
		g.currentFrame = (g.currentFrame + 1) % (len(g.playerLosingFrames))
	}
	if g.playerTickCount >= 10000 {
		g.playerTickCount = 0
	}
}

func AnimateEnemies(g *Game, i int) {
	if g.enemies[i].tickCount%20 == 0 {
		g.enemies[i].frame = (g.enemies[i].frame + 1) % (len(g.enemiesFrames))
	}
}
