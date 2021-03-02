package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"gosah/piese"
	"image/color"
)

// Draw draws the game screen.
// Draw is called every frame typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	// FIXME: tremura ecranul cand misti
	// Initializare piese
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

	// Deseneaza overlay-ul
	if !piese.Started {
		titlu := piese.Nivele[selected]

		// Masoara dim. textului pt a-l centra
		textWidth := text.BoundString(textFont, titlu).Dx()
		sageataHeight := text.BoundString(bigFont, "<").Dy()
		sageataWidth := text.BoundString(bigFont, ">").Dx()

		// Deseneaza titlul nivelului
		text.Draw(screen, titlu, textFont, (piese.Width-textWidth)/2, 80, color.White)

		// Deseneaza "<" si ">" pt selectia nivelului
		text.Draw(screen, "<", bigFont, piese.Offset-sageataWidth-12,
			(piese.Height+sageataHeight)/2, color.White)
		text.Draw(screen, ">", bigFont, piese.Width-piese.Offset-12,
			(piese.Height+sageataHeight)/2, color.White)
	}

	// Afiseaza statusul jocului dupa ce s-a terminat
	if piese.Mat || piese.Pat {
		//time.Sleep(1 * time.Second)

		// Deseneaza fundalul
		rect := ebiten.NewImage(piese.Width, piese.Height)
		rect.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 50})

		opts := &ebiten.DrawImageOptions{}
		screen.DrawImage(rect, opts)

		var titlu string
		switch piese.Castigator {
		case "W":
			titlu = "Alb a castigat"
		case "B":
			titlu = "Negru a castigat"
		default:
			titlu = "Egalitate"
		}

		// Centreaza textul
		ofsx := text.BoundString(textFont, titlu).Dx()
		ofsy := text.BoundString(textFont, titlu).Dy()

		// Deseneaza textul
		text.Draw(screen, titlu, textFont,
			(piese.Width-ofsx)/2, (piese.Height-ofsy)/2,
			color.RGBA{R: 255, G: 69, B: 0, A: 255})
		return
	} else if piese.Started {
		// Deseneaza fundalul timerului
		rect := ebiten.NewImage(piese.Offset, 2*piese.L)
		opts := &ebiten.DrawImageOptions{}
		rect.Fill(color.RGBA{R: 100, G: 100, B: 100, A: 255})
		screen.DrawImage(rect, opts)

		opts.GeoM.Translate(0, 6*piese.L)
		rect.Fill(color.RGBA{R: 200, G: 200, B: 200, A: 255})
		screen.DrawImage(rect, opts)

		timp := fmt.Sprintf("%02d:%02d", piese.TimpRamas.Negru.Min, piese.TimpRamas.Negru.Sec)
		text.Draw(screen, timp, bigFont, 0, 4 * piese.L / 3, color.White)

		timp = fmt.Sprintf("%02d:%02d", piese.TimpRamas.Alb.Min, piese.TimpRamas.Alb.Sec)
		text.Draw(screen, timp, bigFont, 0, 22 * piese.L / 3, color.White)
	}
}
