package main

import (
	"gosah/game"
	"gosah/piese"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Activat, deseneaza doar la update
	//ebiten.SetScreenClearedEveryFrame(false)

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(piese.Width, piese.Height)
	ebiten.SetWindowTitle("Sah")

	// Porneste jocul
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
