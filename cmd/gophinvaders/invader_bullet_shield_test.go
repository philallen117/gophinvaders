package main

import "testing"

func TestInvaderBulletShieldCollision(t *testing.T) {
	game := NewGame(nil, nil)

	// Create a shield.
	shield := NewShield(100, 100)
	game.Shields = []Shield{shield}
	initialHealth := shield.Health

	// Create an active invader bullet that will collide with the shield.
	game.InvaderBullets[0] = InvaderBullet{
		Active: true,
		LeftX:  120,
		TopY:   110,
	}

	game.HandleInvaderBulletShieldCollisions()

	// Bullet should be inactive.
	if game.InvaderBullets[0].Active {
		t.Errorf("Expected bullet to be inactive after shield collision")
	}

	// Shield health should be reduced by 1.
	if game.Shields[0].Health != initialHealth-1 {
		t.Errorf("Expected shield health to be %d, got %d", initialHealth-1, game.Shields[0].Health)
	}
}

func TestShieldDestroyedAtZeroHealth(t *testing.T) {
	game := NewGame(nil, nil)

	// Create a shield with health of 1.
	shield := NewShield(100, 100)
	shield.Health = 1
	game.Shields = []Shield{shield}

	// Create an active invader bullet that will collide with the shield.
	game.InvaderBullets[0] = InvaderBullet{
		Active: true,
		LeftX:  120,
		TopY:   110,
	}

	game.HandleInvaderBulletShieldCollisions()

	// Shield should be destroyed (removed from slice).
	if len(game.Shields) != 0 {
		t.Errorf("Expected shield to be destroyed, but %d shields remain", len(game.Shields))
	}
}

func TestShieldNotDestroyedAboveZeroHealth(t *testing.T) {
	game := NewGame(nil, nil)

	// Create a shield with health of 2.
	shield := NewShield(100, 100)
	shield.Health = 2
	game.Shields = []Shield{shield}

	// Create an active invader bullet that will collide with the shield.
	game.InvaderBullets[0] = InvaderBullet{
		Active: true,
		LeftX:  120,
		TopY:   110,
	}

	game.HandleInvaderBulletShieldCollisions()

	// Shield should still exist.
	if len(game.Shields) != 1 {
		t.Errorf("Expected shield to still exist, got %d shields", len(game.Shields))
	}

	// Shield health should be 1.
	if game.Shields[0].Health != 1 {
		t.Errorf("Expected shield health to be 1, got %d", game.Shields[0].Health)
	}
}

func TestBulletHitsOnlyOneShield(t *testing.T) {
	game := NewGame(nil, nil)

	// Create two overlapping shields.
	shield1 := NewShield(100, 100)
	shield2 := NewShield(110, 110)
	game.Shields = []Shield{shield1, shield2}

	// Create an active invader bullet that will collide with both.
	game.InvaderBullets[0] = InvaderBullet{
		Active: true,
		LeftX:  120,
		TopY:   120,
	}

	game.HandleInvaderBulletShieldCollisions()

	// Bullet should be inactive.
	if game.InvaderBullets[0].Active {
		t.Errorf("Expected bullet to be inactive")
	}

	// Exactly one shield should have taken damage.
	damagedCount := 0
	for _, shield := range game.Shields {
		if shield.Health < shieldInitialHealth {
			damagedCount++
		}
	}

	if damagedCount != 1 {
		t.Errorf("Expected exactly 1 shield to be damaged, got %d", damagedCount)
	}
}

func TestMultipleBulletsCanHitSameShield(t *testing.T) {
	game := NewGame(nil, nil)

	// Create a shield.
	shield := NewShield(100, 100)
	game.Shields = []Shield{shield}

	// Create two active invader bullets that will both collide with the shield.
	game.InvaderBullets[0] = InvaderBullet{
		Active: true,
		LeftX:  120,
		TopY:   110,
	}
	game.InvaderBullets[1] = InvaderBullet{
		Active: true,
		LeftX:  125,
		TopY:   115,
	}

	game.HandleInvaderBulletShieldCollisions()

	// Both bullets should be inactive.
	if game.InvaderBullets[0].Active {
		t.Errorf("Expected bullet 0 to be inactive")
	}
	if game.InvaderBullets[1].Active {
		t.Errorf("Expected bullet 1 to be inactive")
	}

	// Shield health should be reduced by 2.
	expectedHealth := shieldInitialHealth - 2
	if game.Shields[0].Health != expectedHealth {
		t.Errorf("Expected shield health to be %d, got %d", expectedHealth, game.Shields[0].Health)
	}
}

func TestDestroyedShieldNoLongerCollides(t *testing.T) {
	game := NewGame(nil, nil)

	// Create a shield with health of 1 and a second shield.
	shield1 := NewShield(100, 100)
	shield1.Health = 1
	shield2 := NewShield(200, 100)
	game.Shields = []Shield{shield1, shield2}

	// Create an active invader bullet that will destroy shield1.
	game.InvaderBullets[0] = InvaderBullet{
		Active: true,
		LeftX:  120,
		TopY:   110,
	}

	game.HandleInvaderBulletShieldCollisions()

	// First shield should be destroyed.
	if len(game.Shields) != 1 {
		t.Errorf("Expected 1 shield to remain, got %d", len(game.Shields))
	}

	// Create another bullet at the same position.
	game.InvaderBullets[1] = InvaderBullet{
		Active: true,
		LeftX:  120,
		TopY:   110,
	}

	game.HandleInvaderBulletShieldCollisions()

	// Second bullet should still be active (no shield to collide with).
	if !game.InvaderBullets[1].Active {
		t.Errorf("Expected bullet to remain active when no shield at that position")
	}

	// Remaining shield should be undamaged.
	if game.Shields[0].Health != shieldInitialHealth {
		t.Errorf("Expected remaining shield to be undamaged")
	}
}
