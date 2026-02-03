package main

// Rectangler is an interface for objects that can provide their rectangular bounds.
type Rectangler interface {
	// Rectangle returns the rectangular bounds as a 4-tuple:
	// leftX, topY, width, depth (height).
	Rectangle() (leftX, topY, width, depth float32)
}
