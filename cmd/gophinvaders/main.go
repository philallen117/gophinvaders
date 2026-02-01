package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	// Initialize the font faces using text/v2.
	face, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	scoreFontFace := &text.GoTextFace{
		Source: face,
		Size:   float64(scoreTextFontSize),
	}
	gameOverFontFace := &text.GoTextFace{
		Source: face,
		Size:   float64(gameOverTextFontSize),
	}

	ebiten.SetWindowSize(int(screenWidth), int(screenHeight))
	ebiten.SetWindowTitle("gophinvaders")

	game := NewGame(scoreFontFace, gameOverFontFace)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
