package main

import (
	"testing"
)

func TestInvaderRectangle(t *testing.T) {
	tests := []struct {
		name      string
		invader   Invader
		wantLeftX float32
		wantTopY  float32
		wantWidth float32
		wantDepth float32
	}{
		{
			name: "invader at origin",
			invader: Invader{
				LeftX: 0,
				TopY:  0,
			},
			wantLeftX: 0,
			wantTopY:  0,
			wantWidth: invaderWidth,
			wantDepth: invaderHeight,
		},
		{
			name: "invader at arbitrary position",
			invader: Invader{
				LeftX: 150.5,
				TopY:  75.25,
			},
			wantLeftX: 150.5,
			wantTopY:  75.25,
			wantWidth: invaderWidth,
			wantDepth: invaderHeight,
		},
		{
			name:      "new invader from constructor",
			invader:   NewInvader(invaderStartX, invaderStartY),
			wantLeftX: invaderStartX,
			wantTopY:  invaderStartY,
			wantWidth: invaderWidth,
			wantDepth: invaderHeight,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeftX, gotTopY, gotWidth, gotDepth := tt.invader.Rectangle()

			if gotLeftX != tt.wantLeftX {
				t.Errorf("Invader.Rectangle() leftX = %v, want %v", gotLeftX, tt.wantLeftX)
			}
			if gotTopY != tt.wantTopY {
				t.Errorf("Invader.Rectangle() topY = %v, want %v", gotTopY, tt.wantTopY)
			}
			if gotWidth != tt.wantWidth {
				t.Errorf("Invader.Rectangle() width = %v, want %v", gotWidth, tt.wantWidth)
			}
			if gotDepth != tt.wantDepth {
				t.Errorf("Invader.Rectangle() depth = %v, want %v", gotDepth, tt.wantDepth)
			}
		})
	}
}

func TestNewInvader(t *testing.T) {
	tests := []struct {
		name      string
		leftX     float32
		topY      float32
		wantLeftX float32
		wantTopY  float32
	}{
		{
			name:      "new invader at origin",
			leftX:     0,
			topY:      0,
			wantLeftX: 0,
			wantTopY:  0,
		},
		{
			name:      "new invader at arbitrary position",
			leftX:     150.5,
			topY:      75.25,
			wantLeftX: 150.5,
			wantTopY:  75.25,
		},
		{
			name:      "new invader at start position",
			leftX:     invaderStartX,
			topY:      invaderStartY,
			wantLeftX: invaderStartX,
			wantTopY:  invaderStartY,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invader := NewInvader(tt.leftX, tt.topY)

			if invader.LeftX != tt.wantLeftX {
				t.Errorf("NewInvader().LeftX = %v, want %v", invader.LeftX, tt.wantLeftX)
			}
			if invader.TopY != tt.wantTopY {
				t.Errorf("NewInvader().TopY = %v, want %v", invader.TopY, tt.wantTopY)
			}
		})
	}
}
