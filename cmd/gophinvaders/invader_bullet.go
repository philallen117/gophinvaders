package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type InvaderBullet struct {
	LeftX  float32
	TopY   float32
	Active bool // false is the zero value for bool and means bullet is available in the pool
}

func (ib *InvaderBullet) Move() {
	if !ib.Active {
		return
	}

	ib.TopY += bulletSpeed
	if ib.TopY >= screenHeight {
		ib.Active = false
	}
}

func (ib *InvaderBullet) Draw(screen *ebiten.Image) {
	if !ib.Active {
		return
	}

	vector.FillRect(screen, ib.LeftX, ib.TopY, bulletWidth, bulletHeight, invaderBulletColor, false)
}

// Rectangle returns the rectangular bounds of the invader bullet.
func (ib *InvaderBullet) Rectangle() (leftX, topY, width, depth float32) {
	return ib.LeftX, ib.TopY, bulletWidth, bulletHeight
}
