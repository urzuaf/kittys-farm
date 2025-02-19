package gameloop

import (
	"fmt"
	"image/color"
	_ "image/png"
	"time"

	"game/loadsprites"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Config struct {
	ScreenWidth       float64
	ScreenHeight      float64
	PlayerSpeed       float64
	EnemieSpeed       float64
	SpriteScaleFactor float64
}

var Configuration Config = Config{
	ScreenWidth:       480 * 0.9,
	ScreenHeight:      640 * 0.9,
	PlayerSpeed:       4,
	EnemieSpeed:       3,
	SpriteScaleFactor: 2.0,
}

type Game struct {
	PlayerX    float64
	PlayerY    float64
	gameOver   bool
	shootTimer int
	shootRate  int
	score      int

	playerSpriteSheet *ebiten.Image
	playerFrames      []*ebiten.Image
	currentFrame      int

	playerTickCount int

	bulletImage *ebiten.Image

	enemiesSpriteSheet *ebiten.Image
	enemiesFrames      []*ebiten.Image
	enemies            []Enemie

	//flags
	hitSoundPlayed bool
}

type Enemie struct {
	x, y      float64
	tickCount int
	frame     int
}

func (g *Game) Update() error {

	// Comprobar si el juego ha terminado y reiniciar
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			time.Sleep(100 * time.Millisecond)
			g.Reset()
		}
		return nil
	}

	//Aumentar puntuación cada tick
	g.score++

	//Mover al jugador
	ControlPlayerMovement(g)

	//Animar Jugador
	AnimateCharacter(g)

	// Disparo del enemigo
	EnemiesShooting(g)

	//Aumentar dificultad progresivamente
	IncreaseEnemiesMovementSpeed(g)

	// Movimiento de las balas
	EnemiesMovement(g)

	// Colisión con el jugador
	if CheckCollisions(g.PlayerX, g.PlayerY, g.playerFrames[g.currentFrame], g.enemiesFrames, g.enemies) {
		g.gameOver = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//Dibujar menu
	if g.gameOver {
		if !g.hitSoundPlayed {
			HitSound()
		}
		g.hitSoundPlayed = true // Evita que el sonido se reproduzca varias veces
		time.Sleep(300 * time.Millisecond)
		switchMusic(MenuState)
		ebitenutil.DebugPrint(screen, "GAME OVER - Press SPACE to Restart")
		return
	}

	//Si llegamos aqui es porque no estamos en el menu
	switchMusic(GameplayState)

	//dibujar fondo con color verde
	backgroundColor := color.RGBA{192, 212, 112, 255}
	screen.Fill(backgroundColor)

	// Dibujar score
	var str string = fmt.Sprintf("Score: %d ", g.score)
	ebitenutil.DebugPrint(screen, str)

	// Dibujar jugador
	drawer := &ebiten.DrawImageOptions{}
	drawer.GeoM.Scale(2, 2)
	drawer.GeoM.Translate(g.PlayerX, g.PlayerY)
	screen.DrawImage(g.playerFrames[g.currentFrame], drawer)

	// Dibujar balas enemigas
	for _, bullet := range g.enemies {
		b := g.enemiesFrames[bullet.frame]

		var bulletGeoM ebiten.GeoM
		bulletGeoM.Scale(2, 2)
		bulletGeoM.Translate(bullet.x, bullet.y)

		screen.DrawImage(b, &ebiten.DrawImageOptions{
			GeoM: bulletGeoM,
		})
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(Configuration.ScreenWidth), int(Configuration.ScreenHeight)
}

func (g *Game) Reset() {

	//Resetear parametros
	ResetGameStruct(g)

	//Load main character
	g.playerSpriteSheet = loadsprites.LoadFromImage("./assets/player/mc.png")
	GetMcSprites(g)

	//Load enemies
	g.enemiesSpriteSheet = loadsprites.LoadFromImage("./assets/enemies/chickens.png")
	GetEnemiesSprites(g)

	//Pending to DELETE
	img, _, err := ebitenutil.NewImageFromFile("./assets/bullet.png")
	if err != nil {
		panic(err)
	}
	g.bulletImage = img
}
