package main

import (
	"fmt"
	"gosah/piese"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type game struct{}

type piesaSelectata struct {
	ref  *piese.Piesa
	x, y int
}

const (
	width  = 1080
	height = 1080
	l      = width / 8
)

var (
	board            [8][8]piese.Piesa
	selected         piesaSelectata
	clicked, changed bool
)

func getSquare() (int, int) {
	j, i := ebiten.CursorPosition()
	// Arata unde poti merge la pretul de cateva fps
	changed = true
	return i / l, j / l
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update(_ *ebiten.Image) error {
	// Write your game's logical update.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		clicked = !clicked
		if clicked {
			x, y := getSquare()
			board[x][y].Move(&board, x, y)
			selected = piesaSelectata{&board[x][y], x, y}
		} else {
			if x, y := getSquare(); board[x][y].Atacat {
				changed = true

				board[x][y] = *selected.ref
				board[x][y].Mutat = true
				board[selected.x][selected.y] = piese.Empty()
				selected = piesaSelectata{nil, 0, 0}

			}
			piese.Clear(&board)

			// FOR TESTING PURPOSES
			for i := 0; i < 8; i++ {
				fmt.Print(i+1, "     ")
				for j := 0; j < 8; j++ {
					fmt.Print(board[i][j].Control, " ")
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

	// TODO: se deseneaza de doua ori una peste alta, gofix

	// Deseneaza doar daca a fost efectuata o schimbare
	if changed == true {
		changed = false
		square, _ := ebiten.NewImage(l, l, ebiten.FilterNearest)
		opts := &ebiten.DrawImageOptions{}

		screen.Clear()

		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {

				/*if x, y := getSquare(); i == x && j == y {
					// Patratul selectat
					_ = square.Fill(color.RGBA{G: 230, B: 64, A: 255})
				} else*/if board[i][j].Atacat {
					_ = square.Fill(color.RGBA{R: 238, G: 238, A: 255})
				} else {
					if (i+j)%2 == 0 {
						// Patratele Albe
						_ = square.Fill(color.RGBA{R: 205, G: 133, B: 63, A: 170})
					} else {
						// Patratele Negre
						_ = square.Fill(color.RGBA{R: 128, G: 128, B: 128, A: 30})
					}
				}

				_ = screen.DrawImage(square, opts)
				img := board[i][j].DrawPiece()
				if img != nil {
					//opts.GeoM.Scale(0.8, 0.8)
					_ = screen.DrawImage(img, opts)
					//opts.GeoM.Scale(1.25, 1.25)
				}
				opts.GeoM.Translate(height/8, 0)
			}
			opts.GeoM.Translate(-9/8*height, height/8)
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// FIXME: sa functioneze si pe alte rezolutii
	return outsideWidth, outsideHeight
}

func initializareMatrice( /*gameMode rune*/ ) {
	board[0][4], board[7][4] = piese.NewPiesa('K', 'B'), piese.NewPiesa('K', 'W')

	// Initializare rand
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			if !(i == 0 && j == 4) {
				r := rand.Int()
				switch r % 5 {
				case 0:
					// Pion
					board[i][j], board[7-i][j] = piese.NewPiesa('P', 'B'), piese.NewPiesa('P', 'W')
				case 1:
					// Nebun
					board[i][j], board[7-i][j] = piese.NewPiesa('B', 'B'), piese.NewPiesa('B', 'W')
				case 2:
					// Cal
					board[i][j], board[7-i][j] = piese.NewPiesa('N', 'B'), piese.NewPiesa('N', 'W')
				case 3:
					// Tura
					board[i][j], board[7-i][j] = piese.NewPiesa('R', 'B'), piese.NewPiesa('R', 'W')
				case 4:
					// Regina
					board[i][j], board[7-i][j] = piese.NewPiesa('Q', 'B'), piese.NewPiesa('Q', 'W')
				}
			}
		}
	}

	// FOR TESTING PURPOSES
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", board[i][j].Tip)
		}
		fmt.Print("\n")
	}
	fmt.Println("================")
	// Cronometru
	// go cronometru()
}

func cronometru() {
	for sec := 10; sec > 0; sec-- {
		fmt.Println(sec)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Ai ramas fara timp cioara")
}

func main() {
	initializareMatrice()
	g := &game{}

	// Nu mai da clear la fiecare frame
	ebiten.SetScreenClearedEveryFrame(false)
	changed = true

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Sah")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
