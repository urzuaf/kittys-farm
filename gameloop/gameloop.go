package gameloop

import (
	"fmt"
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
	EnemieXSpeed      float64
	EnemieYSpeed      float64
	SpriteScaleFactor float64
}

var Configuration Config = Config{
	ScreenWidth:       480 * 0.9,
	ScreenHeight:      640 * 0.9,
	PlayerSpeed:       4,
	EnemieXSpeed:      3,
	EnemieYSpeed:      2,
	SpriteScaleFactor: 2.0,
}

type Game struct {
	InitialMenu bool
	PlayerX     float64
	PlayerY     float64
	gameOver    bool
	shootTimer  int
	shootRate   int
	score       int

	//Player sprites and animations

	playerSpriteSheet  *ebiten.Image
	playerFrames       []*ebiten.Image
	playerLosingFrames []*ebiten.Image
	currentFrame       int

	playerTickCount int

	//Enemies sprites and animations
	enemiesSpriteSheet *ebiten.Image
	enemiesFrames      []*ebiten.Image
	enemies            []Enemie
	enemiePattern      int

	//flags
	hitSoundPlayed bool

	//bgMovement
	BackgroundY float64
}

type Enemie struct {
	x, y      float64
	tickCount int
	frame     int
}

func (g *Game) Update() error {

	//Comprobar si el juego debe iniciar por primera vez
	if g.InitialMenu {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			time.Sleep(100 * time.Millisecond)
			g.Reset()
		}
		return nil
	}

	// Comprobar si el juego ha terminado y reiniciar
	if g.gameOver {
		AnimateLosingCharacter(g)
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			time.Sleep(100 * time.Millisecond)
			g.Reset()
		}
		return nil
	}

	//Aumentar puntuación cada tick
	g.score++
	g.BackgroundY += 2

	if g.BackgroundY >= Configuration.ScreenHeight {
		g.BackgroundY = 0
	}

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

	//Dibujar menu al iniciar el juego
	if g.InitialMenu {
		bg, _, err := ebitenutil.NewImageFromFile("assets/map/titlePressed.png")
		if err != nil {
			panic(err)
		}
		screen.DrawImage(bg, &ebiten.DrawImageOptions{})
		return
	}
	//Dibujar menu al perder
	if g.gameOver {
		if !g.hitSoundPlayed {
			HitSound()
		}
		g.hitSoundPlayed = true // Evita que el sonido se reproduzca varias veces
		time.Sleep(300 * time.Millisecond)
		SwitchMusic(MenuState)

		//Dibujos
		drawer := &ebiten.DrawImageOptions{}
		drawer.GeoM.Scale(10, 10)
		drawer.GeoM.Translate(Configuration.ScreenWidth/2-64, Configuration.ScreenHeight/2-120)

		bg, _, err := ebitenutil.NewImageFromFile("assets/map/tryAgainPressed.png")
		if err != nil {
			panic(err)
		}
		screen.DrawImage(bg, &ebiten.DrawImageOptions{})
		screen.DrawImage(g.playerLosingFrames[g.currentFrame], drawer)

		return
	}

	//Si llegamos aqui es porque no estamos en el menu
	SwitchMusic(GameplayState)

	//dibujar fondo con color verde
	background, _, err := ebitenutil.NewImageFromFile("assets/map/map.png")
	if err != nil {
		panic(err)
	}
	bgdrawer := &ebiten.DrawImageOptions{}
	bgdrawer.GeoM.Translate(0, g.BackgroundY)
	screen.DrawImage(background, bgdrawer)
	//segundo fondo
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(0, g.BackgroundY-Configuration.ScreenHeight) // Dibuja una segunda copia arriba
	screen.DrawImage(background, op2)

	// Dibujar score
	var str string = fmt.Sprintf("Score: %d, Pattern: %d", g.score, g.enemiePattern)
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

	//DrawHitboxes(g, screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(Configuration.ScreenWidth), int(Configuration.ScreenHeight)
}

func (g *Game) Reset() {

	//Resetear parametros
	ResetGame(g)

	//Load main character
	g.playerSpriteSheet = loadsprites.LoadFromImage("./assets/player/mc.png")
	GetMcSprites(g)

	//Load main character losing animation
	GetLosingMcSprites(g)

	//Load enemies
	g.enemiesSpriteSheet = loadsprites.LoadFromImage("./assets/enemies/chickens.png")
	GetEnemiesSprites(g)

}
