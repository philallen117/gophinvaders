package main

type InvaderBullet struct {
	LeftX  float32
	TopY   float32
	Active bool // false is the zero value for bool and means bullet is available in the pool
}

// Rectangle returns the rectangular bounds of the invader bullet.
func (ib *InvaderBullet) Rectangle() (leftX, topY, width, depth float32) {
	return ib.LeftX, ib.TopY, bulletWidth, bulletHeight
}
