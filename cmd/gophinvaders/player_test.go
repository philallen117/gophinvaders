package main

import (
	"testing"
)

func TestPlayerRectangle(t *testing.T) {
	tests := []struct {
		name      string
		player    Player
		wantLeftX float32
		wantTopY  float32
		wantWidth float32
		wantDepth float32
	}{
		{
			name: "default player at origin",
			player: Player{
				LeftX: 0,
				TopY:  0,
			},
			wantLeftX: 0,
			wantTopY:  0,
			wantWidth: playerWidth,
			wantDepth: playerHeight,
		},
		{
			name: "player at arbitrary position",
			player: Player{
				LeftX: 100.5,
				TopY:  200.25,
			},
			wantLeftX: 100.5,
			wantTopY:  200.25,
			wantWidth: playerWidth,
			wantDepth: playerHeight,
		},
		{
			name:      "new player from constructor",
			player:    NewPlayer(),
			wantLeftX: screenWidth/2 - playerWidth/2,
			wantTopY:  screenHeight - playerHeight - 50,
			wantWidth: playerWidth,
			wantDepth: playerHeight,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeftX, gotTopY, gotWidth, gotDepth := tt.player.Rectangle()

			if gotLeftX != tt.wantLeftX {
				t.Errorf("Player.Rectangle() leftX = %v, want %v", gotLeftX, tt.wantLeftX)
			}
			if gotTopY != tt.wantTopY {
				t.Errorf("Player.Rectangle() topY = %v, want %v", gotTopY, tt.wantTopY)
			}
			if gotWidth != tt.wantWidth {
				t.Errorf("Player.Rectangle() width = %v, want %v", gotWidth, tt.wantWidth)
			}
			if gotDepth != tt.wantDepth {
				t.Errorf("Player.Rectangle() depth = %v, want %v", gotDepth, tt.wantDepth)
			}
		})
	}
}
