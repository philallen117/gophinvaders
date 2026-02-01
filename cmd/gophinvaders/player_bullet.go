package main

type PlayerBullet struct {
	LeftX  float32
	TopY   float32
	Active bool // false is the zero value for bool and means bullet is available in the pool
}

// Rectangle returns the rectangular bounds of the player bullet.
func (pb *PlayerBullet) Rectangle() (leftX, topY, width, depth float32) {
	return pb.LeftX, pb.TopY, bulletWidth, bulletHeight
}
