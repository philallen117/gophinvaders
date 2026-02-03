package main

import (
	"testing"
)

func TestNewShield(t *testing.T) {
	tests := []struct {
		name       string
		leftX      float32
		topY       float32
		wantLeftX  float32
		wantTopY   float32
		wantHealth int
	}{
		{
			name:       "shield at origin",
			leftX:      0,
			topY:       0,
			wantLeftX:  0,
			wantTopY:   0,
			wantHealth: shieldInitialHealth,
		},
		{
			name:       "shield at arbitrary position",
			leftX:      150,
			topY:       450,
			wantLeftX:  150,
			wantTopY:   450,
			wantHealth: shieldInitialHealth,
		},
		{
			name:       "shield at start position",
			leftX:      shieldStartX,
			topY:       shieldStartY,
			wantLeftX:  shieldStartX,
			wantTopY:   shieldStartY,
			wantHealth: shieldInitialHealth,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shield := NewShield(tt.leftX, tt.topY)

			if shield.LeftX != tt.wantLeftX {
				t.Errorf("NewShield().LeftX = %v, want %v", shield.LeftX, tt.wantLeftX)
			}
			if shield.TopY != tt.wantTopY {
				t.Errorf("NewShield().TopY = %v, want %v", shield.TopY, tt.wantTopY)
			}
			if shield.Health != tt.wantHealth {
				t.Errorf("NewShield().Health = %v, want %v", shield.Health, tt.wantHealth)
			}
		})
	}
}

func TestShieldRectangle(t *testing.T) {
	tests := []struct {
		name      string
		shield    Shield
		wantLeftX float32
		wantTopY  float32
		wantWidth float32
		wantDepth float32
	}{
		{
			name: "shield at origin",
			shield: Shield{
				LeftX:  0,
				TopY:   0,
				Health: shieldInitialHealth,
			},
			wantLeftX: 0,
			wantTopY:  0,
			wantWidth: shieldWidth,
			wantDepth: shieldHeight,
		},
		{
			name: "shield at arbitrary position",
			shield: Shield{
				LeftX:  150,
				TopY:   450,
				Health: 5,
			},
			wantLeftX: 150,
			wantTopY:  450,
			wantWidth: shieldWidth,
			wantDepth: shieldHeight,
		},
		{
			name: "damaged shield",
			shield: Shield{
				LeftX:  200,
				TopY:   400,
				Health: 3,
			},
			wantLeftX: 200,
			wantTopY:  400,
			wantWidth: shieldWidth,
			wantDepth: shieldHeight,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeftX, gotTopY, gotWidth, gotDepth := tt.shield.Rectangle()

			if gotLeftX != tt.wantLeftX {
				t.Errorf("Shield.Rectangle() leftX = %v, want %v", gotLeftX, tt.wantLeftX)
			}
			if gotTopY != tt.wantTopY {
				t.Errorf("Shield.Rectangle() topY = %v, want %v", gotTopY, tt.wantTopY)
			}
			if gotWidth != tt.wantWidth {
				t.Errorf("Shield.Rectangle() width = %v, want %v", gotWidth, tt.wantWidth)
			}
			if gotDepth != tt.wantDepth {
				t.Errorf("Shield.Rectangle() depth = %v, want %v", gotDepth, tt.wantDepth)
			}
		})
	}
}
