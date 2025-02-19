package gameloop

import "github.com/hajimehoshi/ebiten/v2"

type Rect struct {
	X, Y, W, H float64
}

func CheckCollisions(playerX, playerY float64, playerSprite *ebiten.Image, enemySprites []*ebiten.Image, enemies []Enemie) bool {
	playerBox := Rect{
		X: playerX + float64(playerSprite.Bounds().Dx())/4,
		Y: playerY + float64(playerSprite.Bounds().Dy())/4,
		//Inner padding for more forgiving collisions
		W: float64(playerSprite.Bounds().Dx()) * 2 * 0.7, // Escalado al doble
		H: float64(playerSprite.Bounds().Dy()) * 2 * 0.7,
	}
	for _, enemy := range enemies {

		enemyBox := Rect{
			X: enemy.x + float64(enemySprites[enemy.frame].Bounds().Dx())/2,
			Y: enemy.y + float64(enemySprites[enemy.frame].Bounds().Dy())/2,
			W: (float64(enemySprites[enemy.frame].Bounds().Dx()) * 2) * 0.6, // Escalado al doble y multiplicado por factor de tolerancia
			H: (float64(enemySprites[enemy.frame].Bounds().Dy()) * 2) * 0.6,
		}

		if checkCollisions(playerBox, enemyBox) {
			return true
		}

	}
	return false

}

// Verifica si dos rectángulos se superponen
func checkCollisions(a, b Rect) bool {
	return a.X < b.X+b.W && a.X+a.W > b.X &&
		a.Y < b.Y+b.H && a.Y+a.H > b.Y
}
