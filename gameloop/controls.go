package gameloop

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func ControlPlayerMovement(g *Game) {

	// Movimiento del jugador
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.PlayerX > 0 {
		g.PlayerX -= Configuration.PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.PlayerX < Configuration.ScreenWidth-28 {
		g.PlayerX += Configuration.PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.PlayerY > 100 {
		g.PlayerY -= Configuration.PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.PlayerY < Configuration.ScreenHeight-34 {
		g.PlayerY += Configuration.PlayerSpeed
	}

}
