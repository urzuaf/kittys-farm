package gameloop

import (
	"embed"
	"fmt"
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
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

	PlayerSpriteSheet  *ebiten.Image
	playerFrames       []*ebiten.Image
	playerLosingFrames []*ebiten.Image
	currentFrame       int

	playerTickCount int

	//Enemies sprites and animations
	EnemiesSpriteSheet *ebiten.Image
	enemiesFrames      []*ebiten.Image
	enemies            []Enemie
	enemiePattern      int

	//flags
	hitSoundPlayed bool

	//bgMovement
	BackgroundY float64

	//Images
	BackgroundImage *ebiten.Image
	MenuImage       *ebiten.Image
	TryAgainImage   *ebiten.Image

	//Music
	MenuPlayer *audio.Player
	GamePlayer *audio.Player
	HitPlayer  *audio.Player

	//FS
	FileSystem *embed.FS
}

type Enemie struct {
	x, y      float64
	tickCount int
	frame     int
}

type button struct {
	x, y     float64
	w, h     float64
	hovering bool
}

var menuButton button
var tryAgainButton button

func (g *Game) Update() error {

	//Comprobar si el juego debe iniciar por primera vez
	if g.InitialMenu {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			time.Sleep(100 * time.Millisecond)
			g.Reset()
		}
		//Comprobar el click con el mouse
		x, y := ebiten.CursorPosition()

		// Verificar si el cursor está sobre el rectángulo
		menuButton.hovering = x >= int(menuButton.x) && x <= int(menuButton.x)+int(menuButton.w) && y >= int(menuButton.y) && y <= int(menuButton.y)+int(menuButton.h)
		if menuButton.hovering {
			ebiten.SetCursorShape(ebiten.CursorShapePointer)
		} else {
			ebiten.SetCursorShape(ebiten.CursorShapeDefault)
		}
		if menuButton.hovering && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			time.Sleep(100 * time.Millisecond)
			g.Reset()
		}

		//Comprobar con touch
		touches := ebiten.TouchIDs()
		for _, touchID := range touches {
			tx, ty := ebiten.TouchPosition(touchID)
			if tx >= int(menuButton.x) && tx <= int(menuButton.x+menuButton.w) && ty >= int(menuButton.y) && ty <= int(menuButton.y+menuButton.h) {
				time.Sleep(100 * time.Millisecond)
				g.Reset()
			}
		}

		return nil
	}

	// Comprobar si el juego ha terminado y reiniciar
	if g.gameOver {
		AnimateLosingCharacter(g)
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			time.Sleep(100 * time.Millisecond)
			g.Reset()
		}
		//Comprobar el click con el mouse
		x, y := ebiten.CursorPosition()

		// Verificar si el cursor está sobre el rectángulo
		tryAgainButton.hovering = x >= int(tryAgainButton.x) && x <= int(tryAgainButton.x)+int(tryAgainButton.w) && y >= int(tryAgainButton.y) && y <= int(tryAgainButton.y)+int(tryAgainButton.h)
		if tryAgainButton.hovering {
			ebiten.SetCursorShape(ebiten.CursorShapePointer)
		} else {
			ebiten.SetCursorShape(ebiten.CursorShapeDefault)
		}
		if tryAgainButton.hovering && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			time.Sleep(100 * time.Millisecond)
			g.Reset()
		}

		//Comprobar con touch
		touches := ebiten.TouchIDs()
		for _, touchID := range touches {
			tx, ty := ebiten.TouchPosition(touchID)
			if tx >= int(tryAgainButton.x) && tx <= int(tryAgainButton.x+tryAgainButton.w) && ty >= int(tryAgainButton.y) && ty <= int(tryAgainButton.y+tryAgainButton.h) {
				time.Sleep(100 * time.Millisecond)
				g.Reset()
			}
		}

		return nil
	}

	ebiten.SetCursorShape(ebiten.CursorShapeDefault)

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
		screen.DrawImage(g.MenuImage, &ebiten.DrawImageOptions{})
		return
	}
	//Dibujar menu al perder
	if g.gameOver {
		if !g.hitSoundPlayed {
			g.HitPlayer.Play()
			time.Sleep(200 * time.Millisecond)
			g.HitPlayer.Pause()
		}
		g.hitSoundPlayed = true // Evita que el sonido se reproduzca varias veces
		time.Sleep(300 * time.Millisecond)
		SwitchMusic(g, MenuState)

		//Dibujos
		drawer := &ebiten.DrawImageOptions{}
		drawer.GeoM.Scale(10, 10)
		drawer.GeoM.Translate(Configuration.ScreenWidth/2-64, Configuration.ScreenHeight/2-120)

		screen.DrawImage(g.TryAgainImage, &ebiten.DrawImageOptions{})
		screen.DrawImage(g.playerLosingFrames[g.currentFrame], drawer)

		return
	}

	//Si llegamos aqui es porque no estamos en el menu
	SwitchMusic(g, GameplayState)

	//dibujar fondo con color verde
	bgdrawer := &ebiten.DrawImageOptions{}
	bgdrawer.GeoM.Translate(0, g.BackgroundY)
	screen.DrawImage(g.BackgroundImage, bgdrawer)
	//segundo fondo
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(0, g.BackgroundY-Configuration.ScreenHeight) // Dibuja una segunda copia arriba
	screen.DrawImage(g.BackgroundImage, op2)

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

	menuButton = button{50, 340, 340, 80, false}
	tryAgainButton = button{10, 370, 410, 80, false}
	//Resetear parametros
	ResetGame(g)

	//Load main character
	GetMcSprites(g)

	//Load main character losing animation
	GetLosingMcSprites(g)

	//Load enemies
	GetEnemiesSprites(g)

}
