package gameloop

func AnimateCharacter(g *Game) {
	//animación del jugador
	g.playerTickCount++

	if g.playerTickCount%20 == 0 {
		g.currentFrame = (g.currentFrame + 1) % (len(g.playerFrames))
	}
	if g.playerTickCount >= 10000 {
		g.playerTickCount = 0
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
