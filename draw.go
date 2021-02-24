package main

import (
	"gosah/piese"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Draw draws the game screen.
// Draw is called every frame typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {

	// FIXME: tremura ecranul cand misti
	// Deseneaza doar daca a fost efectuata o schimbare
	if piese.Changed == true {
		// Initializare piese
		//piese.Changed = false
		square := ebiten.NewImage(piese.L, piese.L)
		opts := &ebiten.DrawImageOptions{}

		// Muta pozitia initiala a tablei de sah astfel incat aceasta sa fie un patrat centrat
		opts.GeoM.Translate(piese.Offset, 0)

		screen.Clear()

		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				// Coloreaza cu galben patratele in care se poate ajunge cu piesa
				if piese.Board[i][j].Atacat {
					square.Fill(color.RGBA{R: 238, G: 238, A: 255})
					// Coloreaza patratul regelui alb cu rosu daca e in sah
				} else if piese.SahAlb && i == piese.RegeAlb.X && j == piese.RegeAlb.Y {
					square.Fill(color.RGBA{R: 255, A: 255})
					// Coloreaza patratul regelui negru cu rosu daca e in sah
				} else if piese.SahNegru && i == piese.RegeNegru.X && j == piese.RegeNegru.Y {
					square.Fill(color.RGBA{R: 255, A: 255})
				} else {
					if (i+j)%2 == 0 {
						// Coloreaza patratele albe
						square.Fill(color.RGBA{R: 205, G: 133, B: 63, A: 170})
					} else {
						// Coloreaza patratele negre
						square.Fill(color.RGBA{R: 128, G: 128, B: 128, A: 30})
					}
				}

				// Deseneaza patratul
				screen.DrawImage(square, opts)

				// Deseneaza piesa
				img := piese.Board[i][j].DrawPiece()
				if img != nil {
					//opts.GeoM.Scale(0.8, 0.8)
					screen.DrawImage(img, opts)
					//opts.GeoM.Scale(1.25, 1.25)
				}

				// Muta <opts> la dreapta
				opts.GeoM.Translate(piese.Height/8, 0)
			}

			// Muta <opts> in stanga si mai jos
			opts.GeoM.Translate(-9/8*piese.Height, piese.Height/8)
		}
	}

	// Seteaza titlul nivelului (daca e in meniu)
	if !piese.Started {
		titlu := piese.Nivele[selected]
		textWidth := text.BoundString(textFont, titlu).Dx()
		text.Draw(screen, titlu, textFont, (piese.Width-textWidth)/2, 80, color.White)
	}
}
