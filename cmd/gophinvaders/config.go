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

// Invader configuration.
const (
	invaderWidth        = 40
	invaderHeight       = 30
	invaderSpeedX       = 5
	invaderRows         = 5
	invaderCols         = 11
	invaderStartX       = 100
	invaderStartY       = 50
	invaderSpacingX     = 60
	invaderSpacingY     = 40
	invaderMoveDelay    = 30 // frames
	invaderDropDistance = 20
	invaderShootDelay   = 60 // frames
	invaderShootChance  = 5  // percent chance per delay interval per live invader
)

var invaderColor = colornames.Red
