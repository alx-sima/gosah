package main

import (
	"fmt"
	"gosah/piese"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// DO NOT TOUCH THIS IT WORKS
type game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update(_ *ebiten.Image) error {
	// Write your game's logical update.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := piese.GetSquare()

		// Daca apesi pe un patrat atacat apelezi functia de mutare
		if piese.Board[x][y].Atacat == true {
			piese.Mutare()
			piese.Clicked = false
		}
		// Daca ultimul clic a fost pe o piesa, se reseteaza tabla inainte de a inregistra clicul curent
		if piese.Clicked {
			piese.Clear(&piese.Board)
		}
		// Daca clicul a fost pe o piesa afiseaza patratele pe care se poate misca
		if piese.Board[x][y].Tip != 0 {
			piese.Clicked = true
		} else {
			piese.Clicked = false
		}

		if piese.Clicked {
			piese.AfisarePatrateAtacate(x, y)
		}

		// Afisare matrice (doar pt testing)
		for i := 0; i < 8; i++ {
			fmt.Print(i+1, "     ")
			for j := 0; j < 8; j++ {
				fmt.Print(piese.Board[i][j].Control, " ")
			}
			fmt.Print("\n")
		}
		fmt.Println("================")
	}
	//fmt.Println(ebiten.CurrentFPS())
	return nil
}

// Draw draws the game screen.
// Draw is called every frame typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	// FIXME: se deseneaza de doua ori una peste alta
	// Deseneaza doar daca a fost efectuata o schimbare
	if piese.Changed == true {
		piese.Changed = false
		square, _ := ebiten.NewImage(piese.L, piese.L, ebiten.FilterNearest)
		opts := &ebiten.DrawImageOptions{}

		_ = screen.Clear()

		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {

				/*
					// Coloreaza patratul peste care este mouseul
					if x, y := getSquare(); i == x && j == y {
						_ = square.Fill(color.RGBA{G: 230, B: 64, A: 255})
					} else
				*/

				// Coloreaza cu galben patratele in care se poate ajunge cu piesa
				if piese.Board[i][j].Atacat {
					_ = square.Fill(color.RGBA{R: 238, G: 238, A: 255})
					// Coloreaza patratul regelui alb cu rosu daca e in sah
				} else if piese.SahAlb && i == piese.RegeAlb.X && j == piese.RegeAlb.Y {
					_ = square.Fill(color.RGBA{R: 255, A: 255})
					// Coloreaza patratul regelui negru cu rosu daca e in sah
				} else if piese.SahNegru && i == piese.RegeNegru.X && j == piese.RegeNegru.Y {
					_ = square.Fill(color.RGBA{R: 255, A: 255})
				} else {
					if (i+j)%2 == 0 {
						// Coloreaza patratele albe
						_ = square.Fill(color.RGBA{R: 205, G: 133, B: 63, A: 170})
					} else {
						// Coloreaza patratele negre
						_ = square.Fill(color.RGBA{R: 128, G: 128, B: 128, A: 30})
					}
				}

				// Deseneaza patratul
				_ = screen.DrawImage(square, opts)

				img := piese.Board[i][j].DrawPiece()
				if img != nil {
					//opts.GeoM.Scale(0.8, 0.8)
					_ = screen.DrawImage(img, opts)
					//opts.GeoM.Scale(1.25, 1.25)
				}
				// Muta <opts> la dreapta
				opts.GeoM.Translate(piese.Height/8, 0)
			}
			// Muta <opts> in stanga si mai jos
			opts.GeoM.Translate(-9/8*piese.Height, piese.Height/8)
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// FIXME: sa functioneze si pe alte rezolutii
	return outsideHeight, outsideHeight
}

func main() {
	// Initializeaza matricea, jocul si tura
	piese.InitializareMatrice()
	g := &game{}
	piese.Turn = 'W'

	// Nu mai da clear la fiecare frame
	ebiten.SetScreenClearedEveryFrame(false)
	piese.Changed = true

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(piese.Width, piese.Height)
	ebiten.SetWindowTitle("Sah")

	// Porneste jocul
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
