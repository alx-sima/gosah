package main

import (
	"fmt"
	"gosah/piese"
	"gosah/util"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type game struct{}


// Returneaza indicii matricei in care se afla mouse-ul


// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update(_ *ebiten.Image) error {
	// Write your game's logical update.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := util.GetSquare()

		if util.Board[x][y].Atacat == true {
			util.Mutare = true
			util.Clicked = false
		} else {
			util.Mutare = false
		}
		if util.Board[x][y].Tip != 0 {
			util.Clicked = true
		} else {
			if util.Clicked == true {
				piese.Clear(&util.Board)
			}
			util.Clicked = false
		}

		if util.Clicked {
			if util.Board[x][y].Culoare == util.Turn {
				isSah := util.SahNegru || util.SahAlb
				util.Board[x][y].Move(&util.Board, x, y, true, isSah)
				util.Selected = piese.PozitiePiesa{Ref: &util.Board[x][y], X: x, Y: y}
			}
		}
		if util.Mutare {
			if x, y := util.GetSquare(); util.Board[x][y].Atacat {
				util.Changed = true

				util.Board[x][y] = *util.Selected.Ref
				util.Board[x][y].Mutat = true
				util.Board[util.Selected.X][util.Selected.Y] = piese.Empty()
				util.Selected = piese.PozitiePiesa{}

				// Transforma pionul in regina cand ajunge la capat
				if util.Board[x][y].Tip == 'P' {
					if util.Board[x][y].Culoare == 'W' && x == 0 {
						util.Board[x][y].Tip = 'Q'
					}
					if util.Board[x][y].Culoare == 'B' && x == 7 {
						util.Board[x][y].Tip = 'Q'
					}
				}

				// Ia pozitia regelui
				if util.Board[x][y].Tip == 'K' {
					if util.Board[x][y].Culoare == 'W' {
						piese.RegeAlb = piese.PozitiePiesa{Ref: &util.Board[x][y], X: x, Y: y}
					}
					if util.Board[x][y].Culoare == 'B' {
						piese.RegeNegru = piese.PozitiePiesa{Ref: &util.Board[x][y], X: x, Y: y}
					}
				}

				// Schimba tura de joc
				if util.Turn == 'W' {
					// Verifica daca regele negru e in sah
					if util.Board[piese.RegeNegru.X][piese.RegeNegru.Y].Control %2 == 1 {
						util.SahNegru = true
					}
					util.Turn = 'B'
				} else {
					// Verifica daca regele alb e in sah
					if util.Board[piese.RegeAlb.X][piese.RegeAlb.Y].Control > 1 {
						util.SahAlb = true
					}
					util.Turn = 'W'
				}
			}
			piese.Clear(&util.Board)

			// Afisare matrice (doar pt testing)
			for i := 0; i < 8; i++ {
				fmt.Print(i+1, "     ")
				for j := 0; j < 8; j++ {
					fmt.Print(util.Board[i][j].Control, " ")
				}
				fmt.Print("\n")
			}
			fmt.Println("================")
		}
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
	if util.Changed == true {
		util.Changed = false
		square, _ := ebiten.NewImage(util.L, util.L, ebiten.FilterNearest)
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
				if util.Board[i][j].Atacat {
					_ = square.Fill(color.RGBA{R: 238, G: 238, A: 255})
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

				img := util.Board[i][j].DrawPiece()
				if img != nil {
					//opts.GeoM.Scale(0.8, 0.8)
					_ = screen.DrawImage(img, opts)
					//opts.GeoM.Scale(1.25, 1.25)
				}
				// Muta <opts> la dreapta
				opts.GeoM.Translate(util.Height/8, 0)
			}
			// Muta <opts> in stanga si mai jos
			opts.GeoM.Translate(-9/8*util.Height, util.Height/8)
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// FIXME: sa functioneze si pe alte rezolutii
	return outsideHeight, outsideHeight
}

func initializareMatrice() {
	// Initializeaza regii
	util.Board[0][4], util.Board[7][4] = piese.NewPiesa('K', 'B'), piese.NewPiesa('K', 'W')
	piese.RegeAlb = piese.PozitiePiesa{Ref: &util.Board[7][4], X: 7, Y: 4}
	piese.RegeNegru = piese.PozitiePiesa{Ref: &util.Board[0][4], Y: 4}

	// Initializeaza seedul rand-ului
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			// Genereaza piese aleatoriu (mai putin pe pozitia regilor)
			if !(i == 0 && j == 4) {
				r := rand.Int()
				switch r % 5 {
				case 0:
					// Pion
					util.Board[i][j], util.Board[7-i][j] = piese.NewPiesa('P', 'B'), piese.NewPiesa('P', 'W')
				case 1:
					// Nebun
					util.Board[i][j], util.Board[7-i][j] = piese.NewPiesa('B', 'B'), piese.NewPiesa('B', 'W')
				case 2:
					// Cal
					util.Board[i][j], util.Board[7-i][j] = piese.NewPiesa('N', 'B'), piese.NewPiesa('N', 'W')
				case 3:
					// Tura
					util.Board[i][j], util.Board[7-i][j] = piese.NewPiesa('R', 'B'), piese.NewPiesa('R', 'W')
				case 4:
					// Regina
					util.Board[i][j], util.Board[7-i][j] = piese.NewPiesa('Q', 'B'), piese.NewPiesa('Q', 'W')
				}
			}
		}
	}

	// FOR TESTING PURPOSES
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", util.Board[i][j].Tip)
		}
		fmt.Print("\n")
	}
	fmt.Println("================")
	// Cronometru
	// go cronometru()
}



func main() {
	// Initializeaza matricea, jocul si tura
	initializareMatrice()
	g := &game{}
	util.Turn = 'W'

	// Nu mai da clear la fiecare frame
	ebiten.SetScreenClearedEveryFrame(false)
	util.Changed = true

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(util.Width, util.Height)
	ebiten.SetWindowTitle("Sah")

	// Porneste jocul
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
