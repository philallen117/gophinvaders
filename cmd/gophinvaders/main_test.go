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
