package main

import (
	"testing"
)

func TestMoveInvaders(t *testing.T) {
	t.Run("invaders move right when counter reaches delay", func(t *testing.T) {
		game := &Game{
			InvaderDirection:   1.0,
			InvaderMoveCounter: invaderMoveDelay - 1,
			Invaders: []Invader{
				{LeftX: 100, TopY: 50},
				{LeftX: 200, TopY: 50},
			},
		}

		game.MoveInvaders()

		if game.InvaderMoveCounter != 0 {
			t.Errorf("InvaderMoveCounter = %v, want 0", game.InvaderMoveCounter)
		}
		if game.Invaders[0].LeftX != 100+invaderSpeedX {
			t.Errorf("Invader[0].LeftX = %v, want %v", game.Invaders[0].LeftX, 100+invaderSpeedX)
		}
		if game.Invaders[1].LeftX != 200+invaderSpeedX {
			t.Errorf("Invader[1].LeftX = %v, want %v", game.Invaders[1].LeftX, 200+invaderSpeedX)
		}
	})

	t.Run("invaders do not move when counter below delay", func(t *testing.T) {
		game := &Game{
			InvaderDirection:   1.0,
			InvaderMoveCounter: 5,
			Invaders: []Invader{
				{LeftX: 100, TopY: 50},
			},
		}

		game.MoveInvaders()

		if game.InvaderMoveCounter != 6 {
			t.Errorf("InvaderMoveCounter = %v, want 6", game.InvaderMoveCounter)
		}
		if game.Invaders[0].LeftX != 100 {
			t.Errorf("Invader[0].LeftX = %v, want 100", game.Invaders[0].LeftX)
		}
	})

	t.Run("invaders drop and reverse when hitting right edge", func(t *testing.T) {
		game := &Game{
			InvaderDirection:   1.0,
			InvaderMoveCounter: invaderMoveDelay - 1,
			Invaders: []Invader{
				{LeftX: screenWidth - invaderWidth - 2, TopY: 50},
			},
		}

		game.MoveInvaders()

		if game.InvaderDirection != -1.0 {
			t.Errorf("InvaderDirection = %v, want -1.0", game.InvaderDirection)
		}
		if game.Invaders[0].TopY != 50+invaderDropDistance {
			t.Errorf("Invader[0].TopY = %v, want %v", game.Invaders[0].TopY, 50+invaderDropDistance)
		}
		// X position should not change when dropping.
		if game.Invaders[0].LeftX != screenWidth-invaderWidth-2 {
			t.Errorf("Invader[0].LeftX = %v, want %v", game.Invaders[0].LeftX, screenWidth-invaderWidth-2)
		}
	})

	t.Run("invaders drop and reverse when hitting left edge", func(t *testing.T) {
		game := &Game{
			InvaderDirection:   -1.0,
			InvaderMoveCounter: invaderMoveDelay - 1,
			Invaders: []Invader{
				{LeftX: 2, TopY: 50},
			},
		}

		game.MoveInvaders()

		if game.InvaderDirection != 1.0 {
			t.Errorf("InvaderDirection = %v, want 1.0", game.InvaderDirection)
		}
		if game.Invaders[0].TopY != 50+invaderDropDistance {
			t.Errorf("Invader[0].TopY = %v, want %v", game.Invaders[0].TopY, 50+invaderDropDistance)
		}
		if game.Invaders[0].LeftX != 2 {
			t.Errorf("Invader[0].LeftX = %v, want 2", game.Invaders[0].LeftX)
		}
	})

	t.Run("all invaders drop together when one hits edge", func(t *testing.T) {
		game := &Game{
			InvaderDirection:   1.0,
			InvaderMoveCounter: invaderMoveDelay - 1,
			Invaders: []Invader{
				{LeftX: 100, TopY: 50},
				{LeftX: 200, TopY: 75},
				{LeftX: screenWidth - invaderWidth - 2, TopY: 100},
			},
		}
		originalY0 := game.Invaders[0].TopY
		originalY1 := game.Invaders[1].TopY
		originalY2 := game.Invaders[2].TopY

		game.MoveInvaders()

		if game.Invaders[0].TopY != originalY0+invaderDropDistance {
			t.Errorf("Invader[0].TopY = %v, want %v", game.Invaders[0].TopY, originalY0+invaderDropDistance)
		}
		if game.Invaders[1].TopY != originalY1+invaderDropDistance {
			t.Errorf("Invader[1].TopY = %v, want %v", game.Invaders[1].TopY, originalY1+invaderDropDistance)
		}
		if game.Invaders[2].TopY != originalY2+invaderDropDistance {
			t.Errorf("Invader[2].TopY = %v, want %v", game.Invaders[2].TopY, originalY2+invaderDropDistance)
		}
	})

	t.Run("invaders move left when direction is negative", func(t *testing.T) {
		game := &Game{
			InvaderDirection:   -1.0,
			InvaderMoveCounter: invaderMoveDelay - 1,
			Invaders: []Invader{
				{LeftX: 100, TopY: 50},
			},
		}

		game.MoveInvaders()

		if game.Invaders[0].LeftX != 100-invaderSpeedX {
			t.Errorf("Invader[0].LeftX = %v, want %v", game.Invaders[0].LeftX, 100-invaderSpeedX)
		}
	})
}

func TestCheckCollision(t *testing.T) {
	t.Run("overlapping rectangles collide", func(t *testing.T) {
		invader := &Invader{LeftX: 100, TopY: 50}
		bullet := &PlayerBullet{LeftX: 110, TopY: 60, Active: true}

		ix, iy, iw, ih := invader.Rectangle()
		bx, by, bw, bh := bullet.Rectangle()
		if !CheckCollision(bx, by, bw, bh, ix, iy, iw, ih) {
			t.Error("Expected collision between overlapping rectangles")
		}
	})

	t.Run("separated rectangles do not collide", func(t *testing.T) {
		invader := &Invader{LeftX: 100, TopY: 50}
		bullet := &PlayerBullet{LeftX: 200, TopY: 200, Active: true}

		ix, iy, iw, ih := invader.Rectangle()
		bx, by, bw, bh := bullet.Rectangle()
		if CheckCollision(bx, by, bw, bh, ix, iy, iw, ih) {
			t.Error("Expected no collision between separated rectangles")
		}
	})

	t.Run("adjacent rectangles do not collide", func(t *testing.T) {
		invader := &Invader{LeftX: 100, TopY: 50}
		// Bullet is exactly to the right of invader (touching but not overlapping).
		bullet := &PlayerBullet{LeftX: 100 + invaderWidth, TopY: 50, Active: true}

		ix, iy, iw, ih := invader.Rectangle()
		bx, by, bw, bh := bullet.Rectangle()
		if CheckCollision(bx, by, bw, bh, ix, iy, iw, ih) {
			t.Error("Expected no collision between adjacent rectangles")
		}
	})

	t.Run("bullet inside invader collides", func(t *testing.T) {
		invader := &Invader{LeftX: 100, TopY: 50}
		// Bullet completely inside invader bounds.
		bullet := &PlayerBullet{LeftX: 110, TopY: 60, Active: true}

		ix, iy, iw, ih := invader.Rectangle()
		bx, by, bw, bh := bullet.Rectangle()
		if !CheckCollision(bx, by, bw, bh, ix, iy, iw, ih) {
			t.Error("Expected collision when bullet is inside invader")
		}
	})

	t.Run("partial overlap collides", func(t *testing.T) {
		invader := &Invader{LeftX: 100, TopY: 50}
		// Invader: X: 100-140, Y: 50-80
		// Bullet overlaps bottom-left corner of invader.
		// Bullet: X: 98-102, Y: 75-85 (overlaps invader X: 100-102, Y: 75-80)
		bullet := &PlayerBullet{LeftX: 98, TopY: 75, Active: true}

		ix, iy, iw, ih := invader.Rectangle()
		bx, by, bw, bh := bullet.Rectangle()
		if !CheckCollision(bx, by, bw, bh, ix, iy, iw, ih) {
			t.Error("Expected collision with partial overlap")
		}
	})
}

func TestBulletKillsOneInvader(t *testing.T) {
	t.Run("bullet kills only first invader it hits", func(t *testing.T) {
		game := &Game{
			Score: 0,
			Invaders: []Invader{
				{LeftX: 100, TopY: 50},
				{LeftX: 100, TopY: 90}, // Directly below first invader.
			},
		}
		// Bullet positioned to hit first invader.
		game.PlayerBullets[0] = PlayerBullet{LeftX: 110, TopY: 60, Active: true}

		game.HandleBulletInvaderCollisions()

		if len(game.Invaders) != 1 {
			t.Errorf("Expected 1 invader remaining, got %d", len(game.Invaders))
		}
		if !game.PlayerBullets[0].Active {
			// Bullet should be deactivated after hitting an invader.
		} else {
			t.Error("Expected bullet to be deactivated after collision")
		}
		if game.Score != killScore {
			t.Errorf("Expected score to be %d, got %d", killScore, game.Score)
		}
	})

	t.Run("bullet misses all invaders", func(t *testing.T) {
		game := &Game{
			Score: 0,
			Invaders: []Invader{
				{LeftX: 100, TopY: 50},
			},
		}
		// Bullet far from invader.
		game.PlayerBullets[0] = PlayerBullet{LeftX: 500, TopY: 500, Active: true}

		game.HandleBulletInvaderCollisions()

		if len(game.Invaders) != 1 {
			t.Errorf("Expected 1 invader remaining, got %d", len(game.Invaders))
		}
		if !game.PlayerBullets[0].Active {
			t.Error("Expected bullet to remain active when missing")
		}
		if game.Score != 0 {
			t.Errorf("Expected score to remain 0, got %d", game.Score)
		}
	})
}

func TestMultipleBulletsKillMultipleInvaders(t *testing.T) {
	t.Run("multiple bullets kill different invaders in same frame", func(t *testing.T) {
		game := &Game{
			Score: 0,
			Invaders: []Invader{
				{LeftX: 100, TopY: 50},
				{LeftX: 200, TopY: 50},
				{LeftX: 300, TopY: 50},
			},
		}
		// Three bullets hitting three different invaders.
		game.PlayerBullets[0] = PlayerBullet{LeftX: 110, TopY: 60, Active: true}
		game.PlayerBullets[1] = PlayerBullet{LeftX: 210, TopY: 60, Active: true}
		game.PlayerBullets[2] = PlayerBullet{LeftX: 310, TopY: 60, Active: true}

		game.HandleBulletInvaderCollisions()

		if len(game.Invaders) != 0 {
			t.Errorf("Expected 0 invaders remaining, got %d", len(game.Invaders))
		}
		if game.PlayerBullets[0].Active || game.PlayerBullets[1].Active || game.PlayerBullets[2].Active {
			t.Error("Expected all bullets to be deactivated after collisions")
		}
		expectedScore := 3 * killScore
		if game.Score != expectedScore {
			t.Errorf("Expected score to be %d, got %d", expectedScore, game.Score)
		}
	})

	t.Run("two bullets hit one invader - only first bullet registers", func(t *testing.T) {
		game := &Game{
			Score: 0,
			Invaders: []Invader{
				{LeftX: 100, TopY: 50},
			},
		}
		// Two bullets both positioned to hit the same invader.
		game.PlayerBullets[0] = PlayerBullet{LeftX: 110, TopY: 60, Active: true}
		game.PlayerBullets[1] = PlayerBullet{LeftX: 115, TopY: 65, Active: true}

		game.HandleBulletInvaderCollisions()

		// Only one invader, so only one bullet can kill it.
		if len(game.Invaders) != 0 {
			t.Errorf("Expected 0 invaders remaining, got %d", len(game.Invaders))
		}
		// Both bullets checked the invader, but only first one killed it.
		// Second bullet would find no invaders left to collide with.
		deactivatedCount := 0
		if !game.PlayerBullets[0].Active {
			deactivatedCount++
		}
		if !game.PlayerBullets[1].Active {
			deactivatedCount++
		}
		if deactivatedCount != 1 {
			t.Errorf("Expected exactly 1 bullet deactivated, got %d", deactivatedCount)
		}
		if game.Score != killScore {
			t.Errorf("Expected score to be %d, got %d", killScore, game.Score)
		}
	})
}
