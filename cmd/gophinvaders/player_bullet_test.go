package main

import (
	"testing"
)

func TestPlayerBulletRectangle(t *testing.T) {
	tests := []struct {
		name      string
		bullet    PlayerBullet
		wantLeftX float32
		wantTopY  float32
		wantWidth float32
		wantDepth float32
	}{
		{
			name: "inactive bullet at origin",
			bullet: PlayerBullet{
				LeftX:  0,
				TopY:   0,
				Active: false,
			},
			wantLeftX: 0,
			wantTopY:  0,
			wantWidth: bulletWidth,
			wantDepth: bulletHeight,
		},
		{
			name: "active bullet at arbitrary position",
			bullet: PlayerBullet{
				LeftX:  50.5,
				TopY:   100.25,
				Active: true,
			},
			wantLeftX: 50.5,
			wantTopY:  100.25,
			wantWidth: bulletWidth,
			wantDepth: bulletHeight,
		},
		{
			name: "bullet at screen edge",
			bullet: PlayerBullet{
				LeftX:  screenWidth - bulletWidth,
				TopY:   0,
				Active: true,
			},
			wantLeftX: screenWidth - bulletWidth,
			wantTopY:  0,
			wantWidth: bulletWidth,
			wantDepth: bulletHeight,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeftX, gotTopY, gotWidth, gotDepth := tt.bullet.Rectangle()

			if gotLeftX != tt.wantLeftX {
				t.Errorf("PlayerBullet.Rectangle() leftX = %v, want %v", gotLeftX, tt.wantLeftX)
			}
			if gotTopY != tt.wantTopY {
				t.Errorf("PlayerBullet.Rectangle() topY = %v, want %v", gotTopY, tt.wantTopY)
			}
			if gotWidth != tt.wantWidth {
				t.Errorf("PlayerBullet.Rectangle() width = %v, want %v", gotWidth, tt.wantWidth)
			}
			if gotDepth != tt.wantDepth {
				t.Errorf("PlayerBullet.Rectangle() depth = %v, want %v", gotDepth, tt.wantDepth)
			}
		})
	}
}
