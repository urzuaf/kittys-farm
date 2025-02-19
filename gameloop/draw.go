package gameloop

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawHitboxes(g *Game, screen *ebiten.Image) {
	//Dibujar hitbox jugador
	playerBox := Rect{
		X: g.PlayerX + (float64(g.playerFrames[g.currentFrame].Bounds().Dx()) / 4),
		Y: g.PlayerY + (float64(g.playerFrames[g.currentFrame].Bounds().Dy()) / 4),
		W: float64(g.playerFrames[g.currentFrame].Bounds().Dx()) * 2 * 0.7, // Escalado al doble
		H: float64(g.playerFrames[g.currentFrame].Bounds().Dy()) * 2 * 0.7,
	}
	DrawHitbox(screen, playerBox)

	// Dibujar hitbox de los enemigos
	for _, enemy := range g.enemies {
		enemyBox := Rect{
			X: enemy.x + float64(g.enemiesFrames[enemy.frame].Bounds().Dx())/2,
			Y: enemy.y + float64(g.enemiesFrames[enemy.frame].Bounds().Dy())/2,
			W: (float64(g.enemiesFrames[enemy.frame].Bounds().Dx()) * 2) * 0.6,
			H: (float64(g.enemiesFrames[enemy.frame].Bounds().Dy()) * 2) * 0.6,
		}

		DrawHitbox(screen, enemyBox)
	}
}

func DrawHitbox(screen *ebiten.Image, box Rect) {
	hitbox := ebiten.NewImage(int(box.W), int(box.H))
	hitbox.Fill(color.RGBA{255, 0, 0, 100}) // Rojo semi-transparente

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(box.X, box.Y) // Posición en pantalla

	screen.DrawImage(hitbox, op)
}
