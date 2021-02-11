package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	width  = 800
	height = 800
)

type game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	square, _ := ebiten.NewImage(width/8, height/8, ebiten.FilterNearest)
	opts := &ebiten.DrawImageOptions{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				square.Fill(color.White)
			} else {
				square.Fill(color.Black)
			}
			screen.DrawImage(square, opts)
			opts.GeoM.Translate(height/8, 0)
		}
		opts.GeoM.Translate(-9/8*height, height/8)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Sah")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
