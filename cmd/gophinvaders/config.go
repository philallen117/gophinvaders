package main

import (
	"golang.org/x/image/colornames"
)

// Screen configuration.
const (
	screenWidth  = 800
	screenHeight = 600
)

// Player configuration.
const (
	playerWidth  = 50
	playerHeight = 30
	playerSpeed  = 5
)

var playerColor = colornames.Blue

// Bullet configuration.
const (
	numPlayerBullets  = 10
	numInvaderBullets = 20
	bulletSpeed       = 10 // -ve for upward movement for player bullets
	bulletWidth       = 4
	bulletHeight      = 10
)

var (
	playerBulletColor  = colornames.White
	invaderBulletColor = colornames.Red
)
