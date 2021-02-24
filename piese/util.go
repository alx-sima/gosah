package piese

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// GetSquare returneaza patratul in care se afla mouse-ul
func GetSquare() (int, int) {
	j, i := ebiten.CursorPosition()
	i, j = i/L, j/L
	Changed = true
	if i < 0 || j < 0 || i > 7 || j > 7 {
		panic("mouse in afara tablei")
	}
	return i, j
}

// Cronometru ar trebui sa numere 10 minute pt fiecare jucator
func Cronometru() {
	// TODO: adaugat cronometru
	for sec := 10; sec > 0; sec-- {
		fmt.Println(sec)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Ai ramas fara timp!")
}

// AfisarePatrateAtacate genereaza mutarile posibile pentru piesa din (x, y) si o memoreaza in Selected
func AfisarePatrateAtacate(x, y int) {
	if Board[x][y].Culoare == Turn {
		Board[x][y].Move(&Board, x, y, true, SahNegru || SahAlb)
		Selected = PozitiePiesa{Ref: &Board[x][y], X: x, Y: y}
	}
}

// Mutare muta piesa selectata pe pozitia ceruta
func Mutare() {
	if x, y := GetSquare(); Board[x][y].Atacat {

		if Board[x][y].Tip != 0 {
			if Turn == 'W' {
				for i := 0; i < len(PieseNegre); i++ {
					if Board[x][y].Tip == PieseNegre[i] {
						PieseNegre = remove(PieseNegre, i)
						break
					}
				}
			} else {
				for i := 0; i < len(PieseAlbe); i++ {
					if Board[x][y].Tip == PieseAlbe[i] {
						PieseAlbe = remove(PieseAlbe, i)
						break
					}
				}
			}
			MutariUltimaCapturare = 0
		} else {
			MutariUltimaCapturare++
		}

		// Translateaza piesa din selected pe pozitia (x, y)
		Board[x][y] = *Selected.Ref
		Board[x][y].Mutat = true

		// Verifica daca mutarea provoaca o rocada
		if Board[x][y].Tip == 'K' {
			// In dreapta
			if y-Selected.Y == 2 {
				Board[x][y-1], Board[x][y+1] = Board[x][y+1], Board[x][y-1]
				// In stanga
			} else if Selected.Y-y == 2 {
				Board[x][y+1], Board[x][y-2] = Board[x][y-2], Board[x][y+1]
			}
		}
		// IMPORTANT! aceasta verificare pentru pion trebuie facuta inainte de clear
		// Daca piesa captureaza prin en passant, elimina piesa capturata de pe tabla
		if Board[x][y].Tip == 'P' {
			if verifInBound(x-1, y) {
				if Board[x-1][y].EnPassant && Selected.X-x == -1 && (Selected.Y-y == 1 || Selected.Y-y == -1) {
					Board[x-1][y] = Empty()
				}
			}
			if verifInBound(x+1, y) {
				if Board[x+1][y].EnPassant && Selected.X-x == 1 && (Selected.Y-y == 1 || Selected.Y-y == -1) {
					Board[x+1][y] = Empty()
				}
			}

			// Transforma pionul in regina cand ajunge la capat
			if Board[x][y].Culoare == 'W' && x == 0 {
				Board[x][y].Tip = 'Q'
			}
			if Board[x][y].Culoare == 'B' && x == 7 {
				Board[x][y].Tip = 'Q'
			}
		}
		// Stergere pozitia initial selectat
		Board[Selected.X][Selected.Y] = Empty()
		Selected = PozitiePiesa{}

		// Reseteaza tabla de sah si de pozitii atacate
		SahAlb, SahNegru = false, false
		Clear(&Board, true)

		// IMPORTANT!  aceasta verificare pentru pion trebuie facuta dupa de clear
		if Board[x][y].Tip == 'P' {

			// Daca pionul s-a mutat 2 patratele, retine ca e apt pt. en passant
			if Selected.X-x == 2 || Selected.X-x == -2 {
				Board[x][y].EnPassant = true
			}
		}

		// Ia pozitia regelui
		if Board[x][y].Tip == 'K' {
			if Board[x][y].Culoare == 'W' {
				RegeAlb = PozitiePiesa{Ref: &Board[x][y], X: x, Y: y}
			}
			if Board[x][y].Culoare == 'B' {
				RegeNegru = PozitiePiesa{Ref: &Board[x][y], X: x, Y: y}
			}
		}

		// Schimba tura de joc
		if Turn == 'W' {
			// Verifica daca regele negru e in sah
			if Board[RegeNegru.X][RegeNegru.Y].eControlatDeCuloare('W') {
				SahNegru = true
				verifMat(Board[RegeNegru.X][RegeNegru.Y].Culoare)
			}
			Turn = 'B'
		} else {
			// Verifica daca regele alb e in sah
			if Board[RegeAlb.X][RegeAlb.Y].eControlatDeCuloare('B') {
				SahAlb = true
				verifMat(Board[RegeAlb.X][RegeAlb.Y].Culoare)
			}
			Turn = 'W'
		}
		if !Mat {
			VerifPat()
		}

		Changed = true
	}
}

// remove ia sliceul slice si returneaza un nou slice, fara elementul de la pozitia s
func remove(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}
