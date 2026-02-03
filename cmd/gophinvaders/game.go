package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// NewGame creates and initializes a new Game instance.
func NewGame(scoreFontFace, gameOverFontFace *text.GoTextFace) *Game {
	game := &Game{
		ScoreFontFace:    scoreFontFace,
		GameOverFontFace: gameOverFontFace,
	}
	game.Initialize()
	return game
}

// Initialize resets all game state to starting conditions.
func (g *Game) Initialize() {
	// Reset player.
	g.Player = NewPlayer()

	// Clear all bullet pools.
	for i := range g.PlayerBullets {
		g.PlayerBullets[i].Active = false
	}
	for i := range g.InvaderBullets {
		g.InvaderBullets[i].Active = false
	}

	// Reset invader state.
	g.InvaderDirection = 1.0 // Start moving right
	g.InvaderMoveCounter = 0
	g.InvaderShootCounter = 0

	// Initialize invader grid.
	g.Invaders = make([]Invader, 0, invaderRows*invaderCols)
	for row := 0; row < invaderRows; row++ {
		for col := 0; col < invaderCols; col++ {
			x := invaderStartX + float32(col)*invaderSpacingX
			y := invaderStartY + float32(row)*invaderSpacingY
			g.Invaders = append(g.Invaders, NewInvader(x, y))
		}
	}

	// Initialize shields.
	g.Shields = make([]Shield, 0, shieldStartCount)
	for i := 0; i < shieldStartCount; i++ {
		x := shieldStartX + float32(i)*shieldSpacingX
		g.Shields = append(g.Shields, NewShield(x, shieldStartY))
	}

	// Reset score and game state flags.
	g.Score = 0
	g.GameLost = false
	g.PlayerWon = false
}

// Game implements ebiten.Game interface.
type Game struct {
	Player              Player
	PlayerBullets       [numPlayerBullets]PlayerBullet
	InvaderBullets      [numInvaderBullets]InvaderBullet
	Invaders            []Invader
	Shields             []Shield
	InvaderDirection    float32 // Positive = right, negative = left
	InvaderMoveCounter  int     // Counts frames until next movement
	InvaderShootCounter int     // Counts frames until next shooting check
	Score               int
	GameLost            bool
	PlayerWon           bool
	ScoreFontFace       *text.GoTextFace
	GameOverFontFace    *text.GoTextFace
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
		g.InvaderBullets[i].Draw(screen)
	}
}

func (g *Game) DrawInvaders(screen *ebiten.Image) {
	for i := range g.Invaders {
		g.Invaders[i].Draw(screen)
	}
}

func (g *Game) DrawShields(screen *ebiten.Image) {
	for i := range g.Shields {
		g.Shields[i].Draw(screen)
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
		g.InvaderBullets[i].Move()
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

func (g *Game) HandlePlayerBulletInvaderCollisions() {
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
				// Check if all invaders are destroyed.
				if len(g.Invaders) == 0 {
					g.PlayerWon = true
				}
				// Each bullet kills at most one invader.
				break
			}
		}
	}
}

func (g *Game) HandleInvaderShooting() {
	g.InvaderShootCounter++
	if g.InvaderShootCounter < invaderShootDelay {
		return
	}
	g.InvaderShootCounter = 0

	// Each invader has a chance to shoot.
	for i := range g.Invaders {
		if rand.IntN(100) < invaderShootChance {
			// Find an inactive bullet in the pool.
			for j := range g.InvaderBullets {
				bullet := &g.InvaderBullets[j]
				if !bullet.Active {
					x, y := g.Invaders[i].BottomMid()
					bullet.LeftX = x - bulletWidth/2 // Center the bullet
					bullet.TopY = y
					bullet.Active = true
					break
				}
			}
		}
	}
}

func (g *Game) HandleInvaderBulletPlayerCollisions() {
	// Get player rectangle once.
	px, py, pw, ph := g.Player.Rectangle()

	// Check collision with each invader bullet.
	for i := range g.InvaderBullets {
		bullet := &g.InvaderBullets[i]
		if !bullet.Active {
			continue
		}

		// Get bullet rectangle.
		bx, by, bw, bh := bullet.Rectangle()
		if CheckCollision(bx, by, bw, bh, px, py, pw, ph) {
			// Deactivate bullet.
			bullet.Active = false
			// Game is lost.
			g.GameLost = true
			// No need to check other bullets once game is lost.
			return
		}
	}
}

func (g *Game) HandlePlayerBulletShieldCollisions() {
	for i := range g.PlayerBullets {
		bullet := &g.PlayerBullets[i]
		if !bullet.Active {
			continue
		}
		g.processBulletShieldCollision(bullet)
	}
}

func (g *Game) HandleInvaderBulletShieldCollisions() {
	for i := range g.InvaderBullets {
		bullet := &g.InvaderBullets[i]
		if !bullet.Active {
			continue
		}
		g.processBulletShieldCollision(bullet)
	}
}

// processBulletShieldCollision handles collision between any bullet and shields.
// The bullet interface must provide Rectangle() and have an Active field.
func (g *Game) processBulletShieldCollision(bullet interface {
	Rectangle() (float32, float32, float32, float32)
}) {
	// Get bullet rectangle once for this bullet.
	bx, by, bw, bh := bullet.Rectangle()

	// Check collision with each shield.
	// Loop backwards to safely remove shields during iteration.
	for j := len(g.Shields) - 1; j >= 0; j-- {
		shield := &g.Shields[j]
		sx, sy, sw, sh := shield.Rectangle()
		if CheckCollision(bx, by, bw, bh, sx, sy, sw, sh) {
			// Deactivate bullet.
			switch b := bullet.(type) {
			case *PlayerBullet:
				b.Active = false
			case *InvaderBullet:
				b.Active = false
			}
			// Reduce shield health.
			shield.Health--
			// Remove shield if health reaches zero.
			if shield.Health <= 0 {
				g.Shields = append(g.Shields[:j], g.Shields[j+1:]...)
			}
			// Each bullet hits at most one shield.
			break
		}
	}
}

func (g *Game) HandlePlayerShooting() {
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
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Check for restart when game is over.
	if g.GameLost || g.PlayerWon {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Initialize()
		}
		return nil
	}

	// Handle left/right arrow keys.
	var left = ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	var right = ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	g.Player.Move(left, right)

	g.HandlePlayerShooting()
	g.MovePlayerBullets()
	g.HandlePlayerBulletShieldCollisions()
	g.MoveInvaderBullets()
	g.HandlePlayerBulletInvaderCollisions()
	g.HandleInvaderBulletPlayerCollisions()
	g.HandleInvaderBulletShieldCollisions()
	g.HandleInvaderShooting()
	g.MoveInvaders()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	if g.GameLost {
		g.DrawGameOver(screen, "Game Over")
		return
	}

	if g.PlayerWon {
		g.DrawGameOver(screen, "You won!")
		return
	}

	g.DrawInvaders(screen)
	g.DrawShields(screen)
	g.Player.Draw(screen)
	g.DrawPlayerBullets(screen)
	g.DrawInvaderBullets(screen)
	g.DrawScore(screen)
}

func (g *Game) DrawGameOver(screen *ebiten.Image, message string) {
	// Draw message text centered.
	op := &text.DrawOptions{}
	// Measure the text to center it.
	width, _ := text.Measure(message, g.GameOverFontFace, 0)
	op.GeoM.Translate(float64(screenWidth/2-float32(width)/2), float64(screenHeight/2-70))
	op.ColorScale.ScaleWithColor(textColor)
	text.Draw(screen, message, g.GameOverFontFace, op)

	// Draw score text centered below.
	scoreText := fmt.Sprintf("Score: %d", g.Score)
	op = &text.DrawOptions{}
	width, _ = text.Measure(scoreText, g.GameOverFontFace, 0)
	op.GeoM.Translate(float64(screenWidth/2-float32(width)/2), float64(screenHeight/2-10))
	op.ColorScale.ScaleWithColor(textColor)
	text.Draw(screen, scoreText, g.GameOverFontFace, op)

	// Draw restart instruction centered below.
	restartText := "Press ENTER to restart"
	op = &text.DrawOptions{}
	width, _ = text.Measure(restartText, g.GameOverFontFace, 0)
	op.GeoM.Translate(float64(screenWidth/2-float32(width)/2), float64(screenHeight/2+50))
	op.ColorScale.ScaleWithColor(textColor)
	text.Draw(screen, restartText, g.GameOverFontFace, op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (_ *Game) Layout(_, _ int) (int, int) {
	return int(screenWidth), int(screenHeight)
}
