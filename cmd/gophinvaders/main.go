package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Game implements ebiten.Game interface.
type Game struct {
	Player         Player
	PlayerBullets  [numPlayerBullets]PlayerBullet
	InvaderBullets [numInvaderBullets]InvaderBullet
}

func (g *Game) DrawPlayerBullets(screen *ebiten.Image) {
	for i := range g.PlayerBullets {
		bullet := &g.PlayerBullets[i]
		if bullet.Active {
			vector.FillRect(screen, bullet.LeftX, bullet.TopY, bulletWidth, bulletHeight, playerBulletColor, false)
		}
	}
}

func (g *Game) DrawInvaderBullets(screen *ebiten.Image) {
	for i := range g.InvaderBullets {
		bullet := &g.InvaderBullets[i]
		if bullet.Active {
			vector.FillRect(screen, bullet.LeftX, bullet.TopY, bulletWidth, bulletHeight, invaderBulletColor, false)
		}
	}
}

func (g *Game) MovePlayerBullets() {
	for i := range g.PlayerBullets {
		bullet := &g.PlayerBullets[i]
		if bullet.Active {
			bullet.TopY -= bulletSpeed
			if bullet.TopY <= 0 {
				bullet.Active = false // Deactivate bullet when it goes off-screen
			}
		}
	}
}

func (g *Game) MoveInvaderBullets() {
	for i := range g.InvaderBullets {
		bullet := &g.InvaderBullets[i]
		if bullet.Active {
			bullet.TopY += bulletSpeed
			if bullet.TopY >= screenHeight {
				bullet.Active = false // Deactivate bullet when it goes off-screen
			}
		}
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Handle left/right arrow keys.
	var left = ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	var right = ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	g.Player.Move(left, right)

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		for i := range g.PlayerBullets {
			bullet := &g.PlayerBullets[i]
			if !bullet.Active {
				x, y := g.Player.TopMid()
				bullet.LeftX = x - 2 // Center the bullet (assuming bullet width is 4)
				bullet.TopY = y
				bullet.Active = true
				break
			}
		}

	}
	g.MovePlayerBullets()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw static text.
	ebitenutil.DebugPrint(screen, "hello ebiten")

	g.Player.Draw(screen)
	g.DrawPlayerBullets(screen)
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
