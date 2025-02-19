package gameloop

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func ControlPlayerMovement(g *Game) {

	// Movimiento del jugador
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) && g.PlayerX > 0 {
		g.PlayerX -= Configuration.PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) && g.PlayerX < Configuration.ScreenWidth-28 {
		g.PlayerX += Configuration.PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) && g.PlayerY > 40 {
		g.PlayerY -= Configuration.PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) && g.PlayerY < Configuration.ScreenHeight-34 {
		g.PlayerY += Configuration.PlayerSpeed
	}

}
