package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/goregular"
)

// Game implements ebiten.Game interface.
type Game struct {
	Player             Player
	PlayerBullets      [numPlayerBullets]PlayerBullet
	InvaderBullets     [numInvaderBullets]InvaderBullet
	Invaders           []Invader
	InvaderDirection   float32 // Positive = right, negative = left
	InvaderMoveCounter int     // Counts frames until next movement
	Score              int
	ScoreFontFace      *text.GoTextFace
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

func (g *Game) DrawInvaders(screen *ebiten.Image) {
	for i := range g.Invaders {
		g.Invaders[i].Draw(screen)
	}
}

func (g *Game) DrawScore(screen *ebiten.Image) {
	scoreText := fmt.Sprintf("Score: %d", g.Score)
	// text/v2 uses top-left positioning by default.
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(scoreTextX), float64(scoreTextY))
	op.ColorScale.ScaleWithColor(textColor)
	text.Draw(screen, scoreText, g.ScoreFontFace, op)
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

// CheckCollision returns true if two rectangles overlap using AABB collision detection.
// Takes two 4-tuples: (leftX, topY, width, depth) for each rectangle.
func CheckCollision(x1, y1, w1, h1, x2, y2, w2, h2 float32) bool {
	// Check if rectangles overlap.
	return x1 < x2+w2 && x1+w1 > x2 && y1 < y2+h2 && y1+h1 > y2
}

func (g *Game) MoveInvaders() {
	g.InvaderMoveCounter++
	if g.InvaderMoveCounter < invaderMoveDelay {
		return
	}
	g.InvaderMoveCounter = 0

	// Check if any invader will hit the edge after moving.
	var shouldDrop bool
	for i := range g.Invaders {
		newX := g.Invaders[i].LeftX + g.InvaderDirection*invaderSpeedX
		if newX < 0 || newX+invaderWidth > screenWidth {
			shouldDrop = true
			break
		}
	}

	if shouldDrop {
		// Drop all invaders and reverse direction.
		for i := range g.Invaders {
			g.Invaders[i].TopY += invaderDropDistance
		}
		g.InvaderDirection = -g.InvaderDirection
	} else {
		// Move all invaders horizontally.
		for i := range g.Invaders {
			g.Invaders[i].LeftX += g.InvaderDirection * invaderSpeedX
		}
	}
}

func (g *Game) HandleBulletInvaderCollisions() {
	for i := range g.PlayerBullets {
		bullet := &g.PlayerBullets[i]
		if !bullet.Active {
			continue
		}

		// Get bullet rectangle once for this bullet.
		bx, by, bw, bh := bullet.Rectangle()

		// Check collision with each invader.
		// Loop backwards to safely remove invaders during iteration.
		for j := len(g.Invaders) - 1; j >= 0; j-- {
			ix, iy, iw, ih := g.Invaders[j].Rectangle()
			if CheckCollision(bx, by, bw, bh, ix, iy, iw, ih) {
				// Remove invader.
				g.Invaders = append(g.Invaders[:j], g.Invaders[j+1:]...)
				// Deactivate bullet.
				bullet.Active = false
				// Increment score.
				g.Score += killScore
				// Each bullet kills at most one invader.
				break
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
	g.HandleBulletInvaderCollisions()
	g.MoveInvaders()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawInvaders(screen)
	g.Player.Draw(screen)
	g.DrawPlayerBullets(screen)
	g.DrawScore(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (_ *Game) Layout(_, _ int) (int, int) {
	return int(screenWidth), int(screenHeight)
}

func main() {
	// Initialize the font face for score text using text/v2.
	face, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	scoreFontFace := &text.GoTextFace{
		Source: face,
		Size:   float64(scoreTextFontSize),
	}

	ebiten.SetWindowSize(int(screenWidth), int(screenHeight))
	ebiten.SetWindowTitle("Hello ebiten")

	game := &Game{
		Player:           NewPlayer(),
		InvaderDirection: 1.0, // Start moving right
		ScoreFontFace:    scoreFontFace,
	}

	// Initialize invader grid.
	game.Invaders = make([]Invader, 0, invaderRows*invaderCols)
	for row := 0; row < invaderRows; row++ {
		for col := 0; col < invaderCols; col++ {
			x := invaderStartX + float32(col)*invaderSpacingX
			y := invaderStartY + float32(row)*invaderSpacingY
			game.Invaders = append(game.Invaders, NewInvader(x, y))
		}
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
