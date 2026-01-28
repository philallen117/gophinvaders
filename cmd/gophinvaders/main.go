package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
	playerWidth  = 50
	playerHeight = 30
	playerSpeed  = 5
)

type Player struct {
	LeftX float32
	TopY  float32
}

func NewPlayer() Player {
	return Player{
		LeftX: screenWidth/2 - playerWidth/2,
		TopY:  screenHeight - playerHeight - 50,
	}
}

// Game implements ebiten.Game interface.
type Game struct {
	Player Player
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Handle left/right arrow keys.
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Player.LeftX -= playerSpeed
		if g.Player.LeftX < 0 {
			g.Player.LeftX = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Player.LeftX += playerSpeed
		if g.Player.LeftX > screenWidth-playerWidth {
			g.Player.LeftX = screenWidth - playerWidth
		}
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw static text.
	ebitenutil.DebugPrint(screen, "hello ebiten")

	// Draw player as a filled rectangle.
	vector.FillRect(screen, g.Player.LeftX, g.Player.TopY, playerWidth, playerHeight, color.White, false)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (_ *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello ebiten")

	game := &Game{
		Player: NewPlayer(),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
