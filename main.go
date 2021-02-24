package main

import (
	"fmt"
	"gosah/piese"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// game e o structura care implementeaza interfata ebiten.Game
type game struct{}

var selected int

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update() error {
	// Daca jocul nu este inceput, se selecteaza nivelul
	if !piese.Started {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			_, _, parte := piese.GetSquare()
			switch parte {
			// Porneste nivelul
			case 0:
				piese.IncarcaNivel(piese.Nivele[selected])
				piese.Started = true
				piese.Changed = true
			// Schimba nivelul la stanga
			case -1:
				if selected > 0 {
					selected--
				}
			// Schimba nivelul la dreapta
			case 1:
				if selected < len(piese.Nivele)-1 {
					selected++
				}
			}
		}
		return nil
	}

	// Daca jocul este in modul editing, Update este overrided de piese.editor
	if piese.Editing {
		return nil
	}

	// Daca jocul este in pat, se termina
	if piese.Pat {
		piese.Turn = 'X'
		fmt.Println("Ai egalat")
		return nil
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y, err := piese.GetSquare()
		if err != 0 {
			return nil
		}
		// Daca apesi pe un patrat atacat apelezi functia de mutare
		if piese.Board[x][y].Atacat == true {
			piese.Mutare()
			piese.Clicked = false
		}
		// Daca ultimul clic a fost pe o piesa, se reseteaza tabla inainte de a inregistra clicul curent
		if piese.Clicked {
			piese.Clear(&piese.Board, false)
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
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	// FIXME: tremura ecranul cand misti
	// Deseneaza doar daca a fost efectuata o schimbare
	if piese.Changed == true {
		// Initializare piese
		piese.Changed = false
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
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// FIXME: sa functioneze si pe alte rezolutii
	return outsideWidth, outsideHeight
}

func newGame() *game {
	return &game{}
}

func main() {
	// Nu mai da clear la fiecare frame
	ebiten.SetScreenClearedEveryFrame(false)
	piese.Changed = true

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(piese.Width, piese.Height)
	ebiten.SetWindowTitle("Sah")

	// Porneste jocul
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
