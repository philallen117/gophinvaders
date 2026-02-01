package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	LeftX float32
	TopY  float32
}

func NewPlayer() Player {
	return Player{
		LeftX: screenWidth/2 - playerWidth/2,
		TopY:  screenHeight - playerHeight - 50,
	}
}

func (p *Player) Move(left, right bool) {
	if left {
		p.LeftX -= playerSpeed
		if p.LeftX < 0 {
			p.LeftX = 0
		}
	}
	if right {
		p.LeftX += playerSpeed
		if p.LeftX > screenWidth-playerWidth {
			p.LeftX = screenWidth - playerWidth
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, p.LeftX, p.TopY, playerWidth, playerHeight, playerColor, false)
}

// Rectangle returns the rectangular bounds of the player.
func (p *Player) Rectangle() (leftX, topY, width, depth float32) {
	return p.LeftX, p.TopY, playerWidth, playerHeight
}

// TopMid returns the x and y coordinates of the top-middle point of the player.
func (p *Player) TopMid() (float32, float32) {
	return p.LeftX + playerWidth/2, p.TopY
}
