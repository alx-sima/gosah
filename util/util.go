package util

import (
	"fmt"
	"gosah/piese"
	"time"

	"github.com/hajimehoshi/ebiten"
)
const (
	Width  = 1920
	Height = 1080
	L      = Height / 8
)

var (
	Board                                      [8][8]piese.Piesa
	Selected                                   piese.PozitiePiesa
	Clicked, Changed, Mutare, SahAlb, SahNegru bool
	Turn                                       rune
)
func GetSquare() (int, int) {
	j, i := ebiten.CursorPosition()
	Changed = true
	return i / L, j / L
}
func Cronometru() {
	// TODO
	for sec := 10; sec > 0; sec-- {
		fmt.Println(sec)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Ai ramas fara timp cioara")
}