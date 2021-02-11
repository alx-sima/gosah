package main

import (
	"fmt"
	"gosah/piese"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type game struct{}

const (
	width  = 800
	height = 800
	length = width / 8
)

var board [8][8]piese.Piesa
var turn bool

// Update proceeds the game state.
// Update is called every tick (1/0 [s] by default).
func (g *game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	//fmt.Println(ebiten.CursorPosition())
	return nil
}

// Draw draws the game screen.
// Draw is called every frame typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	square, _ := ebiten.NewImage(length, length, ebiten.FilterNearest)
	opts := &ebiten.DrawImageOptions{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {

			x, y := ebiten.CursorPosition()

			if x/length == j && y/length == i && x >= 0 && y >= 0 {
				square.Fill(color.RGBA{0, 230, 64, 255})
			} else {
				if (i+j)%2 == 0 {
					// Patratele Albe
					square.Fill(color.RGBA{205, 133, 63, 170})
				} else {
					// Patratele Negre
					square.Fill(color.RGBA{128, 128, 128, 30})
				}

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
	// FIXME: sa functioneze si pe alte rezolutii
	return outsideWidth, outsideHeight
}

func initializareMatrice( /*gameMode rune*/ ) {
	board[0][4], board[7][4] = piese.NewPiesa(0, 4, 'K', 'N'), piese.NewPiesa(7, 4, 'K', 'A')

	// Initializare rand
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			if !(i == 0 && j == 4) {
				r := rand.Int()
				switch r % 5 {
				case 0:
					// Pion
					board[i][j], board[7-i][j] = piese.NewPiesa(i, j, 'P', 'N'), piese.NewPiesa(7-i, j, 'P', 'A')
				case 1:
					// Nebun
					board[i][j], board[7-i][j] = piese.NewPiesa(i, j, 'B', 'N'), piese.NewPiesa(7-i, j, 'B', 'A')
				case 2:
					// Cal
					board[i][j], board[7-i][j] = piese.NewPiesa(i, j, 'N', 'N'), piese.NewPiesa(7-i, j, 'N', 'N')
				case 3:
					// Tura
					board[i][j], board[7-i][j] = piese.NewPiesa(i, j, 'R', 'N'), piese.NewPiesa(7-i, j, 'R', 'A')
				case 4:
					// Regina
					board[i][j], board[7-i][j] = piese.NewPiesa(i, j, 'Q', 'N'), piese.NewPiesa(7-i, j, 'Q', 'A')
				}
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Printf("\n")
	}

	// Cronometru
	go cronometru()
}

var c chan int

func handle(int) {}

func cronometru() {
	/*select {
	case m := <-c:
		handle(m)
	case <-time.After(5 * time.Second):
		fmt.Println("Ai ramas fara timp zdreanta")
	}*/
	for sec := 10; sec > 0; sec-- {
		fmt.Println(sec)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Ai ramas fara timp cioara")
}
func matriceJoc() {
}

func main() {

	initializareMatrice()
	g := &game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Sah")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
