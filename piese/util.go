package piese

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// GetSquare returneaza patratul in care se afla mouse-ul
func GetSquare() (i, j, err int) {
	j, i = ebiten.CursorPosition()

	// Returneaza err=-1 daca e prea in stanga sau prea sus
	if i < 0 || j < Offset {
		return 0, 0, -1
	}

	i, j = (i)/L, (j-Offset)/L

	// Returneaza err=1 daca e prea in dreapta sau prea jos
	if i > 7 || j > 7 {
		return 0, 0, 1
	}

	// Returneaza fara erori
	return
}

// Cronometru numara 10 minute pt fiecare jucator
func Cronometru() {

	time.Sleep(time.Second)

	if Turn == 'W' {
		if TimpRamas.Alb.Sec != 0 {
			TimpRamas.Alb.Sec--
		} else {
			TimpRamas.Alb.Min--
			TimpRamas.Alb.Sec = 59
		}
		if TimpRamas.Alb.Min == 0 && TimpRamas.Alb.Sec == 0 {
			Mat = true
			Castigator = "B"
			return
		}
	} else {
		if TimpRamas.Negru.Sec != 0 {
			TimpRamas.Negru.Sec--
		} else {
			TimpRamas.Negru.Min--
			TimpRamas.Negru.Sec = 59
		}
		if TimpRamas.Negru.Min == 0 && TimpRamas.Negru.Sec == 0 {
			Mat = true
			Castigator = "B"
			return
		}
	}

	if !Mat && !Pat {
		Cronometru()
	}
}

// AfisarePatrateAtacate genereaza mutarile posibile pentru piesa din (x, y) si o memoreaza in selected
func AfisarePatrateAtacate(x, y int) {
	if Board[x][y].Culoare == Turn {
		Board[x][y].Move(&Board, x, y, true, SahNegru || SahAlb, false)
		selected = PozitiePiesa{Ref: &Board[x][y], X: x, Y: y}
	}
}

func adaugareMutare(rezultat string) {
	if Turn == 'W' {
		Miscari.Alb = append(Miscari.Alb, rezultat)
	} else {
		Miscari.Negru = append(Miscari.Negru, rezultat)
	}
}

// Mutare muta piesa selectata pe pozitia ceruta
func Mutare() {
	if x, y, err := GetSquare(); err == 0 && Board[x][y].Atacat {
		if Board[x][y].Tip != 0 {
			if Turn == 'W' {
				ramaseNegre.edit(Board[x][y].Tip, +1)
			} else {
				ramaseAlbe.edit(Board[x][y].Tip, +1)
			}
			mutariUltimaCapturare = 0
		} else {
			mutariUltimaCapturare++
		}

		// Translateaza piesa din selected pe pozitia (x, y)
		Board[x][y] = *selected.Ref
		Board[x][y].Mutat = true

		// Verifica daca mutarea provoaca o rocada
		if Board[x][y].Tip == 'K' {
			// In dreapta
			if y-selected.Y == 2 {
				Board[x][y-1], Board[x][y+1] = Board[x][y+1], Board[x][y-1]
				adaugareMutare("O-O")
				// In stanga
			} else if selected.Y-y == 2 {
				Board[x][y+1], Board[x][y-2] = Board[x][y-2], Board[x][y+1]
				adaugareMutare("O-O-O")
			} else {
				adaugareMutare(numire(mutariUltimaCapturare == 0, selected.X, selected.Y, x, y, Board[x][y].Tip, 0))
			}
			// IMPORTANT! aceasta verificare pentru pion trebuie facuta inainte de clear
			// Daca piesa captureaza prin en passant, elimina piesa capturata de pe tabla
		} else if Board[x][y].Tip == 'P' {
			if verifInBound(x-1, y) {
				if Board[x-1][y].EnPassant && selected.X-x == -1 && (selected.Y-y == 1 || selected.Y-y == -1) {
					Board[x-1][y] = Empty()
				}
			}
			if verifInBound(x+1, y) {
				if Board[x+1][y].EnPassant && selected.X-x == 1 && (selected.Y-y == 1 || selected.Y-y == -1) {
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
			if Board[x][y].Tip == 'P' {
				adaugareMutare(numire(mutariUltimaCapturare == 0, selected.X, selected.Y, x, y, 'P', 0))
			} else {
				adaugareMutare(numire(mutariUltimaCapturare == 0, selected.X, selected.Y, x, y, 'P', Board[x][y].Tip))
			}
		} else {
			adaugareMutare(numire(mutariUltimaCapturare == 0, selected.X, selected.Y, x, y, Board[x][y].Tip, 0))
		}

		// Reseteaza tabla de sah si de pozitii atacate
		SahAlb, SahNegru = false, false
		Clear(&Board, true)

		// IMPORTANT!  aceasta verificare pentru pion trebuie facuta dupa de clear
		if Board[x][y].Tip == 'P' {

			// Daca pionul s-a mutat 2 patratele, retine ca e apt pt. en passant
			if selected.X-x == 2 || selected.X-x == -2 {
				Board[x][y].EnPassant = true
			}
		}

		// Stergere pozitia initial selectat
		Board[selected.X][selected.Y] = Empty()
		selected = PozitiePiesa{}

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

		if Mat {
			if Turn == 'W' {
				Miscari.Alb[len(Miscari.Alb)] += "#"
			} else {
				Miscari.Negru[len(Miscari.Negru)] += "#"
			}
		} else if SahNegru || SahAlb {
			if Turn == 'W' {
				Miscari.Alb[len(Miscari.Alb)] += "+"
			} else {
				Miscari.Negru[len(Miscari.Negru)] += "+"
			}
		}

		if len(Miscari.Negru) == len(Miscari.Alb) {
			fmt.Println(fmt.Sprintf("%d", len(Miscari.Alb)) + " " + Miscari.Alb[len(Miscari.Alb) - 1] + " " + Miscari.Negru[len(Miscari.Negru) - 1])
		}

	}
}
