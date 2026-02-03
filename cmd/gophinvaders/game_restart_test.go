package main

import (
	"testing"
)

func TestRestartAfterGameLoss(t *testing.T) {
	game := NewGame(nil, nil)

	// Set game to lost state with modified values.
	game.GameLost = true
	game.Score = 100
	game.Invaders = nil

	// Call Initialize to restart.
	game.Initialize()

	// Verify state reset.
	if game.GameLost {
		t.Errorf("Expected GameLost to be false after restart")
	}
	if game.Score != 0 {
		t.Errorf("Expected Score to be 0, got %d", game.Score)
	}
	if len(game.Invaders) != invaderRows*invaderCols {
		t.Errorf("Expected %d invaders, got %d", invaderRows*invaderCols, len(game.Invaders))
	}
	if len(game.Shields) != shieldStartCount {
		t.Errorf("Expected %d shields, got %d", shieldStartCount, len(game.Shields))
	}
}

func TestRestartAfterPlayerWon(t *testing.T) {
	game := NewGame(nil, nil)

	// Set game to won state with modified values.
	game.PlayerWon = true
	game.Score = 500
	game.Invaders = nil

	// Call Initialize to restart.
	game.Initialize()

	// Verify state reset.
	if game.PlayerWon {
		t.Errorf("Expected PlayerWon to be false after restart")
	}
	if game.Score != 0 {
		t.Errorf("Expected Score to be 0, got %d", game.Score)
	}
	if len(game.Invaders) != invaderRows*invaderCols {
		t.Errorf("Expected %d invaders, got %d", invaderRows*invaderCols, len(game.Invaders))
	}
	if len(game.Shields) != shieldStartCount {
		t.Errorf("Expected %d shields, got %d", shieldStartCount, len(game.Shields))
	}
}

func TestRestartClearsAllBullets(t *testing.T) {
	game := NewGame(nil, nil)

	// Activate some bullets.
	game.PlayerBullets[0].Active = true
	game.PlayerBullets[1].Active = true
	game.InvaderBullets[0].Active = true
	game.InvaderBullets[5].Active = true

	// Restart game.
	game.Initialize()

	// Verify all bullets inactive.
	for i := range game.PlayerBullets {
		if game.PlayerBullets[i].Active {
			t.Errorf("Expected PlayerBullet[%d] to be inactive", i)
		}
	}
	for i := range game.InvaderBullets {
		if game.InvaderBullets[i].Active {
			t.Errorf("Expected InvaderBullet[%d] to be inactive", i)
		}
	}
}

func TestRestartRecreatesShields(t *testing.T) {
	game := NewGame(nil, nil)

	// Damage and destroy some shields.
	game.Shields[0].Health = 1
	game.Shields = game.Shields[1:] // Remove first shield

	// Restart game.
	game.Initialize()

	// Verify shields recreated.
	if len(game.Shields) != shieldStartCount {
		t.Errorf("Expected %d shields, got %d", shieldStartCount, len(game.Shields))
	}
	for i, shield := range game.Shields {
		if shield.Health != shieldInitialHealth {
			t.Errorf("Expected shield %d to have health %d, got %d", i, shieldInitialHealth, shield.Health)
		}
	}
}

func TestRestartRecreatesInvaders(t *testing.T) {
	game := NewGame(nil, nil)

	// Kill some invaders.
	game.Invaders = game.Invaders[:10]

	// Restart game.
	game.Initialize()

	// Verify full invader formation.
	expectedCount := invaderRows * invaderCols
	if len(game.Invaders) != expectedCount {
		t.Errorf("Expected %d invaders, got %d", expectedCount, len(game.Invaders))
	}
}

func TestRestartResetsCounters(t *testing.T) {
	game := NewGame(nil, nil)

	// Set counters to non-zero values.
	game.InvaderMoveCounter = 50
	game.InvaderShootCounter = 30

	// Restart game.
	game.Initialize()

	// Verify counters reset.
	if game.InvaderMoveCounter != 0 {
		t.Errorf("Expected InvaderMoveCounter to be 0, got %d", game.InvaderMoveCounter)
	}
	if game.InvaderShootCounter != 0 {
		t.Errorf("Expected InvaderShootCounter to be 0, got %d", game.InvaderShootCounter)
	}
}

func TestRestartResetsInvaderDirection(t *testing.T) {
	game := NewGame(nil, nil)

	// Set invader direction to left.
	game.InvaderDirection = -1.0

	// Restart game.
	game.Initialize()

	// Verify direction reset to right.
	if game.InvaderDirection != 1.0 {
		t.Errorf("Expected InvaderDirection to be 1.0, got %f", game.InvaderDirection)
	}
}
