package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	// Initialize the font face for score text using text/v2.
	face, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	scoreFontFace := &text.GoTextFace{
		Source: face,
		Size:   float64(scoreTextFontSize),
	}

	ebiten.SetWindowSize(int(screenWidth), int(screenHeight))
	ebiten.SetWindowTitle("Hello ebiten")

	game := NewGame(scoreFontFace)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
