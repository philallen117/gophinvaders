package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Shield struct {
	LeftX  float32
	TopY   float32
	Health int
}

func NewShield(leftX, topY float32) Shield {
	return Shield{
		LeftX:  leftX,
		TopY:   topY,
		Health: shieldInitialHealth,
	}
}

func (s *Shield) Draw(screen *ebiten.Image) {
	// Calculate alpha based on health.
	// Alpha ranges from 0 (transparent) to 1 (opaque).
	alpha := float32(s.Health) / float32(shieldInitialHealth)
	
	// Create color with alpha.
	color := shieldColor
	color.A = uint8(alpha * 255)
	
	vector.FillRect(screen, s.LeftX, s.TopY, shieldWidth, shieldHeight, color, false)
}

// Rectangle returns the rectangular bounds of the shield.
func (s *Shield) Rectangle() (leftX, topY, width, depth float32) {
	return s.LeftX, s.TopY, shieldWidth, shieldHeight
}
