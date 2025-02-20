package gameloop

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func ControlPlayerMovement(g *Game) {

	// Movimiento del jugador con teclado
	if (ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && g.PlayerX > 0 {
		g.PlayerX -= Configuration.PlayerSpeed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)) && g.PlayerX < Configuration.ScreenWidth-28 {
		g.PlayerX += Configuration.PlayerSpeed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && g.PlayerY > 40 {
		g.PlayerY -= Configuration.PlayerSpeed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && g.PlayerY < Configuration.ScreenHeight-34 {
		g.PlayerY += Configuration.PlayerSpeed
	}

	// Movimiento con touch
	touches := ebiten.TouchIDs()
	if len(touches) > 0 { // Si hay un toque en la pantalla
		x, y := ebiten.TouchPosition(touches[0]) // Obtener coordenadas del primer toque

		// Determinar la dirección según la posición del toque
		if x < int(Configuration.ScreenWidth)/3 { // Izquierda
			g.PlayerX -= Configuration.PlayerSpeed
		} else if x > 2*int(Configuration.ScreenWidth)/3 { // Derecha
			g.PlayerX += Configuration.PlayerSpeed
		}

		if y < int(Configuration.ScreenHeight)/3 { // Arriba
			g.PlayerY -= Configuration.PlayerSpeed
		} else if y > 2*int(Configuration.ScreenHeight)/3 { // Abajo
			g.PlayerY += Configuration.PlayerSpeed
		}
	}

}
