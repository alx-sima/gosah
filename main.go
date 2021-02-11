package main

import (
	"fmt"
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
)

var board [8][8]rune
var turn bool

// Update proceeds the game state.
// Update is called every tick (1/0 [s] by default).
func (g *game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	if turn {

	}
	turn = !turn
	return nil
}

// Draw draws the game screen.
// Draw is called every frame typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	square, _ := ebiten.NewImage(width/8, height/8, ebiten.FilterNearest)
	opts := &ebiten.DrawImageOptions{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				// Patratele Albe
				square.Fill(color.RGBA{205, 133, 63, 170})
			} else {
				// Patratele Negre
				square.Fill(color.RGBA{128, 128, 128, 30})
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
	board[0][4], board[7][4] = 'K', 'K'

	// Initializare rand
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			if !(i == 0 && j == 4) {
				r := rand.Int()
				switch r % 5 {
				case 0:
					// Pion
					board[i][j], board[7-i][j] = 'P', 'P'
				case 1:
					// Nebun
					board[i][j], board[7-i][j] = 'B', 'B'
				case 2:
					// Cal
					board[i][j], board[7-i][j] = 'N', 'N'
				case 3:
					// Tura
					board[i][j], board[7-i][j] = 'R', 'R'
				case 4:
					// Regina
					board[i][j], board[7-i][j] = 'Q', 'Q'
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
	select {
	case m := <-c:
		handle(m)
	case <-time.After(5 * time.Second):
		fmt.Println("timed out")
	}
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
