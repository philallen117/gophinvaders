package main

import (
	"golang.org/x/image/colornames"
)

// Screen configuration.
const (
	screenWidth  float32 = 800
	screenHeight float32 = 600
)

// Player configuration.
const (
	playerWidth  float32 = 50
	playerHeight float32 = 30
	playerSpeed  float32 = 5
)

var playerColor = colornames.Blue

// Bullet configuration.
const (
	numPlayerBullets          = 10
	numInvaderBullets         = 20
	bulletSpeed       float32 = 10 // -ve for upward movement for player bullets
	bulletWidth       float32 = 4
	bulletHeight      float32 = 10
	killScore         int     = 10
)

var (
	playerBulletColor  = colornames.White
	invaderBulletColor = colornames.Red
)

// Invader configuration.
const (
	invaderWidth        float32 = 40
	invaderHeight       float32 = 30
	invaderSpeedX       float32 = 5
	invaderRows                 = 5
	invaderCols                 = 11
	invaderStartX       float32 = 100
	invaderStartY       float32 = 50
	invaderSpacingX     float32 = 60
	invaderSpacingY     float32 = 40
	invaderMoveDelay            = 30 // frames
	invaderDropDistance float32 = 20
	invaderShootDelay           = 60 // frames
	invaderShootChance  float32 = 5  // percent chance per delay interval per live invader
)

var invaderColor = colornames.Red

// Text/Font configuration.
const (
	pointsPerInch             = 72
	dpi                       = 72
	pointPerPixel     float32 = pointsPerInch / dpi
	scoreTextFontSize         = int(25 * pointPerPixel) // points
	scoreTextX        float32 = 20
	scoreTextY        float32 = 10
)

var textColor = colornames.White
