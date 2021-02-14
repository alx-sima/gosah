package util

import (
	"fmt"
	"gosah/piese"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	Width  = 1920       // Lungimea ecranului
	Height = 1080       // Latimea ecranului
	L      = Height / 8 // Latura unui patrat
)

var (
	Board                                      [8][8]piese.Piesa
	Selected                                   piese.PozitiePiesa
	Clicked, Changed, Mutare, SahAlb, SahNegru bool
	Turn                                       rune
)

// Returneaza patratul in care se afla mouse-ul
func GetSquare() (int, int) {
	j, i := ebiten.CursorPosition()
	i, j = i/L, j/L
	Changed = true
	if i < 0 || j < 0 || i > 7 || j > 7 {
		panic("mouse in afara tablei")
	}
	return i, j
}
func Cronometru() {
	// TODO: adaugat cronometru
	for sec := 10; sec > 0; sec-- {
		fmt.Println(sec)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Ai ramas fara timp cioara")
}
