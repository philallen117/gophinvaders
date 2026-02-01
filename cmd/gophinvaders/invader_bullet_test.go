package main

import (
	"testing"
)

func TestInvaderBulletRectangle(t *testing.T) {
	tests := []struct {
		name      string
		bullet    InvaderBullet
		wantLeftX float32
		wantTopY  float32
		wantWidth float32
		wantDepth float32
	}{
		{
			name: "inactive bullet at origin",
			bullet: InvaderBullet{
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
			bullet: InvaderBullet{
				LeftX:  75.5,
				TopY:   150.75,
				Active: true,
			},
			wantLeftX: 75.5,
			wantTopY:  150.75,
			wantWidth: bulletWidth,
			wantDepth: bulletHeight,
		},
		{
			name: "bullet at bottom of screen",
			bullet: InvaderBullet{
				LeftX:  screenWidth / 2,
				TopY:   screenHeight - bulletHeight,
				Active: true,
			},
			wantLeftX: screenWidth / 2,
			wantTopY:  screenHeight - bulletHeight,
			wantWidth: bulletWidth,
			wantDepth: bulletHeight,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeftX, gotTopY, gotWidth, gotDepth := tt.bullet.Rectangle()

			if gotLeftX != tt.wantLeftX {
				t.Errorf("InvaderBullet.Rectangle() leftX = %v, want %v", gotLeftX, tt.wantLeftX)
			}
			if gotTopY != tt.wantTopY {
				t.Errorf("InvaderBullet.Rectangle() topY = %v, want %v", gotTopY, tt.wantTopY)
			}
			if gotWidth != tt.wantWidth {
				t.Errorf("InvaderBullet.Rectangle() width = %v, want %v", gotWidth, tt.wantWidth)
			}
			if gotDepth != tt.wantDepth {
				t.Errorf("InvaderBullet.Rectangle() depth = %v, want %v", gotDepth, tt.wantDepth)
			}
		})
	}
}

func TestInvaderBulletMove(t *testing.T) {
	t.Run("active bullet moves down", func(t *testing.T) {
		bullet := InvaderBullet{LeftX: 100, TopY: 200, Active: true}

		bullet.Move()

		if bullet.TopY != 200+bulletSpeed {
			t.Errorf("InvaderBullet.Move() TopY = %v, want %v", bullet.TopY, 200+bulletSpeed)
		}
		if !bullet.Active {
			t.Error("InvaderBullet.Move() deactivated an on-screen bullet")
		}
	})

	t.Run("inactive bullet does not move", func(t *testing.T) {
		bullet := InvaderBullet{LeftX: 100, TopY: 200, Active: false}

		bullet.Move()

		if bullet.TopY != 200 {
			t.Errorf("InvaderBullet.Move() TopY = %v, want %v", bullet.TopY, 200)
		}
		if bullet.Active {
			t.Error("InvaderBullet.Move() activated an inactive bullet")
		}
	})

	t.Run("bullet deactivates when off-screen", func(t *testing.T) {
		bullet := InvaderBullet{LeftX: 100, TopY: screenHeight, Active: true}

		bullet.Move()

		if bullet.Active {
			t.Error("InvaderBullet.Move() did not deactivate an off-screen bullet")
		}
	})
}
