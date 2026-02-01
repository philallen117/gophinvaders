package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Invader struct {
	LeftX float32
	TopY  float32
}

func NewInvader(leftX, topY float32) Invader {
	return Invader{
		LeftX: leftX,
		TopY:  topY,
	}
}

func (inv *Invader) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, inv.LeftX, inv.TopY, invaderWidth, invaderHeight, invaderColor, false)
}

// Rectangle returns the rectangular bounds of the invader.
func (inv *Invader) Rectangle() (leftX, topY, width, depth float32) {
	return inv.LeftX, inv.TopY, invaderWidth, invaderHeight
}
