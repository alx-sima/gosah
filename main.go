package main

import (
	"gosah/piese"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// game e o structura care implementeaza interfata ebiten.Game
type game struct{}

var (
	selected int
	textFont font.Face
)

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// FIXME: sa functioneze si pe alte rezolutii
	return outsideWidth, outsideHeight
}

func newGame() *game {
	return &game{}
}

// initializare font
func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	textFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    72,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	// Activat, deseneaza doar la update
	//ebiten.SetScreenClearedEveryFrame(false)
	piese.Changed = true

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(piese.Width, piese.Height)
	ebiten.SetWindowTitle("Sah")

	// Porneste jocul
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
