package gameloop

import (
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

// Variables globales para el toque
var (
	touchStartX, touchStartY int // Posición inicial del toque
)

// Función auxiliar para obtener el valor absoluto

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

	// Detectar el arrastre con touch
	touches := ebiten.TouchIDs()

	if len(touches) > 0 { // Si hay un toque activo
		x, y := ebiten.TouchPosition(touches[0])

		if touchStartX == 0 && touchStartY == 0 { // Primer toque: guardar posición inicial
			touchStartX, touchStartY = x, y
		} else {
			// Calcular la diferencia entre la posición actual e inicial
			deltaX := x - touchStartX
			deltaY := y - touchStartY

			// Definir un umbral mínimo para evitar movimientos involuntarios
			threshold := 10

			if utils.AbsValue(deltaX) > utils.AbsValue(deltaY) { // Movimiento horizontal
				if deltaX > threshold && g.PlayerX < Configuration.ScreenWidth-28 { // Derecha
					g.PlayerX += Configuration.PlayerSpeed
				} else if deltaX < -threshold && g.PlayerX > 0 { // Izquierda
					g.PlayerX -= Configuration.PlayerSpeed
				}
			} else { // Movimiento vertical
				if deltaY > threshold && g.PlayerY < Configuration.ScreenHeight-34 { // Abajo
					g.PlayerY += Configuration.PlayerSpeed
				} else if deltaY < -threshold && g.PlayerY > 40 { // Arriba
					g.PlayerY -= Configuration.PlayerSpeed
				}
			}
		}
	} else { // Si no hay toques, reiniciar la posición inicial
		touchStartX, touchStartY = 0, 0
	}
}
